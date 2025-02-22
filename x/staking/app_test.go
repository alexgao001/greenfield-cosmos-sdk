package staking_test

import (
	"encoding/hex"
	"testing"

	"github.com/prysmaticlabs/prysm/crypto/bls"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
)

func checkValidator(t *testing.T, app *simapp.SimApp, addr sdk.AccAddress, expFound bool) types.Validator {
	ctxCheck := app.BaseApp.NewContext(true, tmproto.Header{})
	validator, found := app.StakingKeeper.GetValidator(ctxCheck, addr)

	require.Equal(t, expFound, found)
	return validator
}

func checkDelegation(
	t *testing.T, app *simapp.SimApp, delegatorAddr sdk.AccAddress,
	validatorAddr sdk.AccAddress, expFound bool, expShares sdk.Dec,
) {
	ctxCheck := app.BaseApp.NewContext(true, tmproto.Header{})
	delegation, found := app.StakingKeeper.GetDelegation(ctxCheck, delegatorAddr, validatorAddr)
	if expFound {
		require.True(t, found)
		require.True(sdk.DecEq(t, expShares, delegation.Shares))

		return
	}

	require.False(t, found)
}

func TestStakingMsgs(t *testing.T) {
	genTokens := sdk.TokensFromConsensusPower(42, sdk.DefaultPowerReduction)
	bondTokens := sdk.TokensFromConsensusPower(10, sdk.DefaultPowerReduction)
	genCoin := sdk.NewCoin(sdk.DefaultBondDenom, genTokens)
	bondCoin := sdk.NewCoin(sdk.DefaultBondDenom, bondTokens)

	acc1 := &authtypes.BaseAccount{Address: addr1.String()}
	acc2 := &authtypes.BaseAccount{Address: addr2.String()}
	accs := authtypes.GenesisAccounts{acc1, acc2}
	balances := []banktypes.Balance{
		{
			Address: addr1.String(),
			Coins:   sdk.Coins{genCoin},
		},
		{
			Address: addr2.String(),
			Coins:   sdk.Coins{genCoin},
		},
	}

	app := simapp.SetupWithGenesisAccounts(t, accs, balances...)
	simapp.CheckBalance(t, app, addr1, sdk.Coins{genCoin})
	simapp.CheckBalance(t, app, addr2, sdk.Coins{genCoin})

	// create validator
	description := types.NewDescription("foo_moniker", "", "", "", "")
	blsSecretKey, _ := bls.RandKey()
	blsPubKey := hex.EncodeToString(blsSecretKey.PublicKey().Marshal())
	createValidatorMsg, err := types.NewMsgCreateValidator(
		addr1, valKey.PubKey(),
		bondCoin, description, commissionRates, sdk.OneInt(),
		addr1, addr1, addr1, addr1, blsPubKey,
	)
	require.NoError(t, err)

	header := tmproto.Header{ChainID: simapp.DefaultChainId, Height: app.LastBlockHeight() + 1}
	txGen := simapp.MakeTestEncodingConfig().TxConfig
	_, _, err = simapp.SignCheckDeliver(t, txGen, app.BaseApp, header, []sdk.Msg{createValidatorMsg}, simapp.DefaultChainId, []uint64{0}, []uint64{0}, true, true, []cryptotypes.PrivKey{priv1}, simapp.SetMockHeight(app.BaseApp, 0))
	require.NoError(t, err)
	simapp.CheckBalance(t, app, addr1, sdk.Coins{genCoin.Sub(bondCoin)})

	header = tmproto.Header{ChainID: simapp.DefaultChainId, Height: app.LastBlockHeight() + 1}
	app.BeginBlock(abci.RequestBeginBlock{Header: header})

	validator := checkValidator(t, app, addr1, true)
	require.Equal(t, addr1.String(), validator.OperatorAddress)
	require.Equal(t, types.Bonded, validator.Status)
	require.True(sdk.IntEq(t, bondTokens, validator.BondedTokens()))

	header = tmproto.Header{ChainID: simapp.DefaultChainId, Height: app.LastBlockHeight() + 1}
	app.BeginBlock(abci.RequestBeginBlock{Header: header})

	// edit the validator
	description = types.NewDescription("bar_moniker", "", "", "", "")
	editValidatorMsg := types.NewMsgEditValidator(
		addr1, description, nil, nil,
		sdk.AccAddress(""), sdk.AccAddress(""), "",
	)

	header = tmproto.Header{ChainID: simapp.DefaultChainId, Height: app.LastBlockHeight() + 1}
	_, _, err = simapp.SignCheckDeliver(t, txGen, app.BaseApp, header, []sdk.Msg{editValidatorMsg}, simapp.DefaultChainId, []uint64{0}, []uint64{1}, true, true, []cryptotypes.PrivKey{priv1})
	require.NoError(t, err)

	validator = checkValidator(t, app, addr1, true)
	require.Equal(t, description, validator.Description)

	// delegate
	simapp.CheckBalance(t, app, addr2, sdk.Coins{genCoin})
	delegateMsg := types.NewMsgDelegate(addr2, addr1, bondCoin)

	header = tmproto.Header{ChainID: simapp.DefaultChainId, Height: app.LastBlockHeight() + 1}
	_, _, err = simapp.SignCheckDeliver(t, txGen, app.BaseApp, header, []sdk.Msg{delegateMsg}, simapp.DefaultChainId, []uint64{1}, []uint64{0}, true, true, []cryptotypes.PrivKey{priv2})
	require.NoError(t, err)

	simapp.CheckBalance(t, app, addr2, sdk.Coins{genCoin.Sub(bondCoin)})
	checkDelegation(t, app, addr2, addr1, true, sdk.NewDecFromInt(bondTokens))

	// begin unbonding
	beginUnbondingMsg := types.NewMsgUndelegate(addr2, addr1, bondCoin)
	header = tmproto.Header{ChainID: simapp.DefaultChainId, Height: app.LastBlockHeight() + 1}
	_, _, err = simapp.SignCheckDeliver(t, txGen, app.BaseApp, header, []sdk.Msg{beginUnbondingMsg}, simapp.DefaultChainId, []uint64{1}, []uint64{1}, true, true, []cryptotypes.PrivKey{priv2})
	require.NoError(t, err)

	// delegation should exist anymore
	checkDelegation(t, app, addr2, addr1, false, sdk.Dec{})

	// balance should be the same because bonding not yet complete
	simapp.CheckBalance(t, app, addr2, sdk.Coins{genCoin.Sub(bondCoin)})
}
