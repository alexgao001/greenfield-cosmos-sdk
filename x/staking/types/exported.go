package types

import (
	"cosmossdk.io/math"
	tmprotocrypto "github.com/tendermint/tendermint/proto/tendermint/crypto"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DelegationI delegation bond for a delegated proof of stake system
type DelegationI interface {
	GetDelegatorAddr() sdk.AccAddress // delegator sdk.AccAddress for the bond
	GetValidatorAddr() sdk.AccAddress // validator operator address
	GetShares() sdk.Dec               // amount of validator's shares held in this delegation
}

// ValidatorI expected validator functions
type ValidatorI interface {
	IsJailed() bool                                          // whether the validator is jailed
	GetMoniker() string                                      // moniker of the validator
	GetStatus() BondStatus                                   // status of the validator
	IsBonded() bool                                          // check if has a bonded status
	IsUnbonded() bool                                        // check if has status unbonded
	IsUnbonding() bool                                       // check if has status unbonding
	GetOperator() sdk.AccAddress                             // operator address to receive/return validators coins
	GetSelfDelegator() sdk.AccAddress                        // validator address used for self delegation
	GetRelayer() sdk.AccAddress                              // validator authorized relayer/operator address
	GetChallenger() sdk.AccAddress                           // validator authorized challenger/operator address
	GetBlsKey() []byte                                       // validator authorized relayer/challenger's bls pubkey
	ConsPubKey() (cryptotypes.PubKey, error)                 // validation consensus pubkey (cryptotypes.PubKey)
	TmConsPublicKey() (tmprotocrypto.PublicKey, error)       // validation consensus pubkey (Tendermint)
	GetConsAddr() (sdk.ConsAddress, error)                   // validation consensus address
	GetTokens() math.Int                                     // validation tokens
	GetBondedTokens() math.Int                               // validator bonded tokens
	GetConsensusPower(math.Int) int64                        // validation power in tendermint
	GetCommission() sdk.Dec                                  // validator commission rate
	GetMinSelfDelegation() math.Int                          // validator minimum self delegation
	GetDelegatorShares() sdk.Dec                             // total outstanding delegator shares
	TokensFromShares(sdk.Dec) sdk.Dec                        // token worth of provided delegator shares
	TokensFromSharesTruncated(sdk.Dec) sdk.Dec               // token worth of provided delegator shares, truncated
	TokensFromSharesRoundUp(sdk.Dec) sdk.Dec                 // token worth of provided delegator shares, rounded up
	SharesFromTokens(amt math.Int) (sdk.Dec, error)          // shares worth of delegator's bond
	SharesFromTokensTruncated(amt math.Int) (sdk.Dec, error) // truncated shares worth of delegator's bond
}
