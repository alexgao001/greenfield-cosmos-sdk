// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/crosschain/v1/event.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// EventCrossChain is emitted when there is a cross chain package created
type EventCrossChain struct {
	// Source chain id of the cross chain package
	SrcChainId uint32 `protobuf:"varint,1,opt,name=src_chain_id,json=srcChainId,proto3" json:"src_chain_id,omitempty"`
	// Destination chain id of the cross chainpackage
	DestChainId uint32 `protobuf:"varint,2,opt,name=dest_chain_id,json=destChainId,proto3" json:"dest_chain_id,omitempty"`
	// Channel id of the cross chain package
	ChannelId uint32 `protobuf:"varint,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// Sequence of the cross chain package
	Sequence uint64 `protobuf:"varint,4,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// Package type of the cross chain package, like SYN, ACK and FAIL_ACK
	PackageType uint32 `protobuf:"varint,5,opt,name=package_type,json=packageType,proto3" json:"package_type,omitempty"`
	// Timestamp of the cross chain package
	Timestamp uint64 `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Payload of the cross chain package
	PackageLoad string `protobuf:"bytes,7,opt,name=package_load,json=packageLoad,proto3" json:"package_load,omitempty"`
	// Relayer fee for the cross chain package
	RelayerFee string `protobuf:"bytes,8,opt,name=relayer_fee,json=relayerFee,proto3" json:"relayer_fee,omitempty"`
	// Relayer fee for the ACK or FAIL_ACK package of this cross chain package
	AckRelayerFee string `protobuf:"bytes,9,opt,name=ack_relayer_fee,json=ackRelayerFee,proto3" json:"ack_relayer_fee,omitempty"`
}

func (m *EventCrossChain) Reset()         { *m = EventCrossChain{} }
func (m *EventCrossChain) String() string { return proto.CompactTextString(m) }
func (*EventCrossChain) ProtoMessage()    {}
func (*EventCrossChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a5a3ba75f5dd2c3, []int{0}
}
func (m *EventCrossChain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCrossChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCrossChain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCrossChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCrossChain.Merge(m, src)
}
func (m *EventCrossChain) XXX_Size() int {
	return m.Size()
}
func (m *EventCrossChain) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCrossChain.DiscardUnknown(m)
}

var xxx_messageInfo_EventCrossChain proto.InternalMessageInfo

func (m *EventCrossChain) GetSrcChainId() uint32 {
	if m != nil {
		return m.SrcChainId
	}
	return 0
}

func (m *EventCrossChain) GetDestChainId() uint32 {
	if m != nil {
		return m.DestChainId
	}
	return 0
}

func (m *EventCrossChain) GetChannelId() uint32 {
	if m != nil {
		return m.ChannelId
	}
	return 0
}

func (m *EventCrossChain) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *EventCrossChain) GetPackageType() uint32 {
	if m != nil {
		return m.PackageType
	}
	return 0
}

func (m *EventCrossChain) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *EventCrossChain) GetPackageLoad() string {
	if m != nil {
		return m.PackageLoad
	}
	return ""
}

func (m *EventCrossChain) GetRelayerFee() string {
	if m != nil {
		return m.RelayerFee
	}
	return ""
}

func (m *EventCrossChain) GetAckRelayerFee() string {
	if m != nil {
		return m.AckRelayerFee
	}
	return ""
}

func init() {
	proto.RegisterType((*EventCrossChain)(nil), "cosmos.crosschain.v1.EventCrossChain")
}

func init() { proto.RegisterFile("cosmos/crosschain/v1/event.proto", fileDescriptor_2a5a3ba75f5dd2c3) }

var fileDescriptor_2a5a3ba75f5dd2c3 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xd1, 0x4f, 0x4f, 0xfa, 0x30,
	0x18, 0xc0, 0x71, 0xca, 0x8f, 0x1f, 0xb2, 0x07, 0x08, 0x49, 0xe3, 0x61, 0x31, 0x3a, 0x27, 0x07,
	0xc3, 0xc5, 0x2d, 0xc4, 0x77, 0x20, 0xd1, 0x84, 0xc4, 0x13, 0xf1, 0xe4, 0x65, 0x29, 0xed, 0x23,
	0x2c, 0x63, 0xeb, 0x6c, 0x0b, 0x91, 0x77, 0xe1, 0x3b, 0xf2, 0xea, 0x91, 0xa3, 0x47, 0x03, 0x6f,
	0xc4, 0xb4, 0xfc, 0x3d, 0x2d, 0xfb, 0xf6, 0xd3, 0x27, 0x69, 0x1e, 0x08, 0xb9, 0xd4, 0xb9, 0xd4,
	0x31, 0x57, 0x52, 0x6b, 0x3e, 0x65, 0x69, 0x11, 0x2f, 0xfa, 0x31, 0x2e, 0xb0, 0x30, 0x51, 0xa9,
	0xa4, 0x91, 0xf4, 0x7c, 0x2b, 0xa2, 0xa3, 0x88, 0x16, 0xfd, 0xee, 0x57, 0x15, 0x3a, 0x8f, 0x56,
	0x0d, 0x6c, 0x1e, 0xd8, 0x4c, 0x43, 0x68, 0x69, 0xc5, 0x13, 0x67, 0x92, 0x54, 0xf8, 0x24, 0x24,
	0xbd, 0xf6, 0x08, 0xb4, 0xe2, 0xee, 0x7c, 0x28, 0x68, 0x17, 0xda, 0x02, 0xb5, 0x39, 0x92, 0xaa,
	0x23, 0x4d, 0x1b, 0xf7, 0xe6, 0x0a, 0x80, 0x4f, 0x59, 0x51, 0xe0, 0xcc, 0x82, 0x7f, 0x0e, 0x78,
	0xbb, 0x32, 0x14, 0xf4, 0x02, 0x1a, 0x1a, 0xdf, 0xe7, 0x58, 0x70, 0xf4, 0x6b, 0x21, 0xe9, 0xd5,
	0x46, 0x87, 0x7f, 0x7a, 0x03, 0xad, 0x92, 0xf1, 0x8c, 0x4d, 0x30, 0x31, 0xcb, 0x12, 0xfd, 0xff,
	0xdb, 0xe9, 0xbb, 0xf6, 0xb2, 0x2c, 0x91, 0x5e, 0x82, 0x67, 0xd2, 0x1c, 0xb5, 0x61, 0x79, 0xe9,
	0xd7, 0xdd, 0xfd, 0x63, 0x38, 0x1d, 0x30, 0x93, 0x4c, 0xf8, 0x67, 0x21, 0xe9, 0x79, 0x87, 0x01,
	0xcf, 0x92, 0x09, 0x7a, 0x0d, 0x4d, 0x85, 0x33, 0xb6, 0x44, 0x95, 0xbc, 0x21, 0xfa, 0x0d, 0x27,
	0x60, 0x97, 0x9e, 0x10, 0xe9, 0x2d, 0x74, 0x18, 0xcf, 0x92, 0x53, 0xe4, 0x39, 0xd4, 0x66, 0x3c,
	0x1b, 0x1d, 0xdc, 0xc3, 0xf0, 0x7b, 0x1d, 0x90, 0xd5, 0x3a, 0x20, 0xbf, 0xeb, 0x80, 0x7c, 0x6e,
	0x82, 0xca, 0x6a, 0x13, 0x54, 0x7e, 0x36, 0x41, 0xe5, 0x35, 0x9e, 0xa4, 0x66, 0x3a, 0x1f, 0x47,
	0x5c, 0xe6, 0xf1, 0x7e, 0x3d, 0xee, 0x73, 0xa7, 0x45, 0x16, 0x7f, 0x9c, 0xee, 0xca, 0x3e, 0x53,
	0x8f, 0xeb, 0x6e, 0x53, 0xf7, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x93, 0x08, 0x3b, 0x98, 0xcd,
	0x01, 0x00, 0x00,
}

func (m *EventCrossChain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCrossChain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCrossChain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AckRelayerFee) > 0 {
		i -= len(m.AckRelayerFee)
		copy(dAtA[i:], m.AckRelayerFee)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.AckRelayerFee)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.RelayerFee) > 0 {
		i -= len(m.RelayerFee)
		copy(dAtA[i:], m.RelayerFee)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.RelayerFee)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.PackageLoad) > 0 {
		i -= len(m.PackageLoad)
		copy(dAtA[i:], m.PackageLoad)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.PackageLoad)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Timestamp != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x30
	}
	if m.PackageType != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.PackageType))
		i--
		dAtA[i] = 0x28
	}
	if m.Sequence != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x20
	}
	if m.ChannelId != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.ChannelId))
		i--
		dAtA[i] = 0x18
	}
	if m.DestChainId != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.DestChainId))
		i--
		dAtA[i] = 0x10
	}
	if m.SrcChainId != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.SrcChainId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventCrossChain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SrcChainId != 0 {
		n += 1 + sovEvent(uint64(m.SrcChainId))
	}
	if m.DestChainId != 0 {
		n += 1 + sovEvent(uint64(m.DestChainId))
	}
	if m.ChannelId != 0 {
		n += 1 + sovEvent(uint64(m.ChannelId))
	}
	if m.Sequence != 0 {
		n += 1 + sovEvent(uint64(m.Sequence))
	}
	if m.PackageType != 0 {
		n += 1 + sovEvent(uint64(m.PackageType))
	}
	if m.Timestamp != 0 {
		n += 1 + sovEvent(uint64(m.Timestamp))
	}
	l = len(m.PackageLoad)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.RelayerFee)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.AckRelayerFee)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventCrossChain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventCrossChain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCrossChain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcChainId", wireType)
			}
			m.SrcChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SrcChainId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestChainId", wireType)
			}
			m.DestChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DestChainId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			m.ChannelId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChannelId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PackageType", wireType)
			}
			m.PackageType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PackageType |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PackageLoad", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PackageLoad = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RelayerFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RelayerFee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AckRelayerFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AckRelayerFee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)
