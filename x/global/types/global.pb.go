// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kyve/global/v1beta1/global.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Params defines the global module parameters.
type Params struct {
	// min_gas_price defines the minimum gas price value for all transactions.
	MinGasPrice cosmossdk_io_math.LegacyDec `protobuf:"bytes,1,opt,name=min_gas_price,json=minGasPrice,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"min_gas_price"`
	// burn_ratio defines the ratio of transaction fees burnt.
	BurnRatio cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=burn_ratio,json=burnRatio,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"burn_ratio"`
	// gas_adjustments can add a constant amount of gas to a specific message type.
	// This gives more control to make certain messages more expensive to avoid spamming
	// of certain types of messages.
	GasAdjustments []GasAdjustment `protobuf:"bytes,3,rep,name=gas_adjustments,json=gasAdjustments,proto3" json:"gas_adjustments"`
	// gas_refunds lets the governance specify a fraction of how much gas
	// a user gets refunded for a certain type of transaction.
	// This could be used to make transactions which support to network cheaper.
	// Gas refunds only work if the transaction only included one message.
	GasRefunds []GasRefund `protobuf:"bytes,4,rep,name=gas_refunds,json=gasRefunds,proto3" json:"gas_refunds"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1b5d4c0bbdf8bfb, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetGasAdjustments() []GasAdjustment {
	if m != nil {
		return m.GasAdjustments
	}
	return nil
}

func (m *Params) GetGasRefunds() []GasRefund {
	if m != nil {
		return m.GasRefunds
	}
	return nil
}

// GasAdjustment stores for every message type a fixed amount
// of gas which is added to the message
type GasAdjustment struct {
	// type of the sdk-message
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// amount of gas which is added to the message
	Amount uint64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *GasAdjustment) Reset()         { *m = GasAdjustment{} }
func (m *GasAdjustment) String() string { return proto.CompactTextString(m) }
func (*GasAdjustment) ProtoMessage()    {}
func (*GasAdjustment) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1b5d4c0bbdf8bfb, []int{1}
}
func (m *GasAdjustment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GasAdjustment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GasAdjustment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GasAdjustment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GasAdjustment.Merge(m, src)
}
func (m *GasAdjustment) XXX_Size() int {
	return m.Size()
}
func (m *GasAdjustment) XXX_DiscardUnknown() {
	xxx_messageInfo_GasAdjustment.DiscardUnknown(m)
}

var xxx_messageInfo_GasAdjustment proto.InternalMessageInfo

func (m *GasAdjustment) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GasAdjustment) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// GasRefund stores the fraction of gas which will be refunded for a given
// type of message.
// This only works if the transaction only includes one message.
type GasRefund struct {
	// type of the sdk-message
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// fraction in decimal representation between 0 and 1
	Fraction cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=fraction,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"fraction"`
}

func (m *GasRefund) Reset()         { *m = GasRefund{} }
func (m *GasRefund) String() string { return proto.CompactTextString(m) }
func (*GasRefund) ProtoMessage()    {}
func (*GasRefund) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1b5d4c0bbdf8bfb, []int{2}
}
func (m *GasRefund) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GasRefund) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GasRefund.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GasRefund) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GasRefund.Merge(m, src)
}
func (m *GasRefund) XXX_Size() int {
	return m.Size()
}
func (m *GasRefund) XXX_DiscardUnknown() {
	xxx_messageInfo_GasRefund.DiscardUnknown(m)
}

var xxx_messageInfo_GasRefund proto.InternalMessageInfo

func (m *GasRefund) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "kyve.global.v1beta1.Params")
	proto.RegisterType((*GasAdjustment)(nil), "kyve.global.v1beta1.GasAdjustment")
	proto.RegisterType((*GasRefund)(nil), "kyve.global.v1beta1.GasRefund")
}

func init() { proto.RegisterFile("kyve/global/v1beta1/global.proto", fileDescriptor_d1b5d4c0bbdf8bfb) }

var fileDescriptor_d1b5d4c0bbdf8bfb = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x6b, 0xe2, 0x40,
	0x18, 0xc6, 0x13, 0x0d, 0xb2, 0x8e, 0xb8, 0x0b, 0xb3, 0xcb, 0x12, 0x76, 0x21, 0x4a, 0xf6, 0x22,
	0x2c, 0x24, 0xb8, 0x7b, 0xdc, 0xc3, 0xd2, 0x50, 0xf1, 0xd0, 0x52, 0x6c, 0x0e, 0x85, 0xf6, 0x62,
	0x27, 0x71, 0x1c, 0xa7, 0x3a, 0x33, 0x32, 0x33, 0xb1, 0xf5, 0x5b, 0xf4, 0x9b, 0xf4, 0x6b, 0x78,
	0xf4, 0x58, 0x7a, 0x90, 0xa2, 0x5f, 0xa4, 0x4c, 0x1a, 0xa5, 0x05, 0x0b, 0xed, 0x6d, 0xde, 0x77,
	0x9e, 0xe7, 0xf7, 0xfe, 0xe1, 0x05, 0xcd, 0xf1, 0x7c, 0x86, 0x43, 0x32, 0x11, 0x09, 0x9a, 0x84,
	0xb3, 0x76, 0x82, 0x35, 0x6a, 0x17, 0x61, 0x30, 0x95, 0x42, 0x0b, 0xf8, 0xd5, 0x28, 0x82, 0x22,
	0x55, 0x28, 0x7e, 0x7c, 0x23, 0x82, 0x88, 0xfc, 0x3f, 0x34, 0xaf, 0x67, 0xa9, 0x7f, 0x57, 0x02,
	0x95, 0x1e, 0x92, 0x88, 0x29, 0xd8, 0x05, 0x75, 0x46, 0x79, 0x9f, 0x20, 0xd5, 0x9f, 0x4a, 0x9a,
	0x62, 0xd7, 0x6e, 0xda, 0xad, 0x6a, 0xf4, 0x6b, 0xb1, 0x6a, 0x58, 0x0f, 0xab, 0xc6, 0xcf, 0x54,
	0x28, 0x26, 0x94, 0x1a, 0x8c, 0x03, 0x2a, 0x42, 0x86, 0xf4, 0x28, 0x38, 0xc6, 0x04, 0xa5, 0xf3,
	0x43, 0x9c, 0xc6, 0x35, 0x46, 0x79, 0x17, 0xa9, 0x9e, 0xf1, 0xc1, 0x08, 0x80, 0x24, 0x93, 0xbc,
	0x2f, 0x91, 0xa6, 0xc2, 0x2d, 0xbd, 0x9f, 0x52, 0x35, 0xb6, 0xd8, 0xb8, 0xe0, 0x29, 0xf8, 0x62,
	0x1a, 0x41, 0x83, 0xab, 0x4c, 0x69, 0x86, 0xb9, 0x56, 0x6e, 0xb9, 0x59, 0x6e, 0xd5, 0xfe, 0xf8,
	0xc1, 0x9e, 0xe1, 0x82, 0x2e, 0x52, 0x07, 0x3b, 0x69, 0xe4, 0x98, 0x62, 0xf1, 0x67, 0xf2, 0x32,
	0xa9, 0x60, 0x07, 0xd4, 0x0c, 0x52, 0xe2, 0x61, 0xc6, 0x07, 0xca, 0x75, 0x72, 0x9c, 0xf7, 0x16,
	0x2e, 0xce, 0x65, 0x05, 0x0a, 0x90, 0x6d, 0x42, 0xf9, 0xff, 0x40, 0xfd, 0x55, 0x35, 0x08, 0x81,
	0xa3, 0xe7, 0xd3, 0x62, 0x5d, 0x71, 0xfe, 0x86, 0xdf, 0x41, 0x05, 0x31, 0x91, 0x71, 0x9d, 0x8f,
	0xef, 0xc4, 0x45, 0xe4, 0x5f, 0x82, 0xea, 0x8e, 0xbd, 0xd7, 0xf8, 0x1f, 0x7c, 0x1a, 0x4a, 0x94,
	0x6a, 0x2a, 0xf8, 0x47, 0x36, 0xb7, 0x33, 0x45, 0x9d, 0xc5, 0xda, 0xb3, 0x97, 0x6b, 0xcf, 0x7e,
	0x5c, 0x7b, 0xf6, 0xed, 0xc6, 0xb3, 0x96, 0x1b, 0xcf, 0xba, 0xdf, 0x78, 0xd6, 0xc5, 0x6f, 0x42,
	0xf5, 0x28, 0x4b, 0x82, 0x54, 0xb0, 0xf0, 0xe8, 0xfc, 0xac, 0x73, 0x82, 0xf5, 0xb5, 0x90, 0xe3,
	0x30, 0x1d, 0x21, 0xca, 0xc3, 0x9b, 0xed, 0x45, 0x99, 0x36, 0x54, 0x52, 0xc9, 0xcf, 0xe3, 0xef,
	0x53, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x5d, 0xec, 0x6f, 0x6d, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GasRefunds) > 0 {
		for iNdEx := len(m.GasRefunds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GasRefunds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGlobal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.GasAdjustments) > 0 {
		for iNdEx := len(m.GasAdjustments) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GasAdjustments[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGlobal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size := m.BurnRatio.Size()
		i -= size
		if _, err := m.BurnRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGlobal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.MinGasPrice.Size()
		i -= size
		if _, err := m.MinGasPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGlobal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *GasAdjustment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GasAdjustment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GasAdjustment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != 0 {
		i = encodeVarintGlobal(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintGlobal(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GasRefund) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GasRefund) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GasRefund) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Fraction.Size()
		i -= size
		if _, err := m.Fraction.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGlobal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintGlobal(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGlobal(dAtA []byte, offset int, v uint64) int {
	offset -= sovGlobal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MinGasPrice.Size()
	n += 1 + l + sovGlobal(uint64(l))
	l = m.BurnRatio.Size()
	n += 1 + l + sovGlobal(uint64(l))
	if len(m.GasAdjustments) > 0 {
		for _, e := range m.GasAdjustments {
			l = e.Size()
			n += 1 + l + sovGlobal(uint64(l))
		}
	}
	if len(m.GasRefunds) > 0 {
		for _, e := range m.GasRefunds {
			l = e.Size()
			n += 1 + l + sovGlobal(uint64(l))
		}
	}
	return n
}

func (m *GasAdjustment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovGlobal(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovGlobal(uint64(m.Amount))
	}
	return n
}

func (m *GasRefund) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovGlobal(uint64(l))
	}
	l = m.Fraction.Size()
	n += 1 + l + sovGlobal(uint64(l))
	return n
}

func sovGlobal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGlobal(x uint64) (n int) {
	return sovGlobal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGlobal
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinGasPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGlobal
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
				return ErrInvalidLengthGlobal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGlobal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinGasPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGlobal
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
				return ErrInvalidLengthGlobal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGlobal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BurnRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasAdjustments", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGlobal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGlobal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGlobal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GasAdjustments = append(m.GasAdjustments, GasAdjustment{})
			if err := m.GasAdjustments[len(m.GasAdjustments)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasRefunds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGlobal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGlobal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGlobal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GasRefunds = append(m.GasRefunds, GasRefund{})
			if err := m.GasRefunds[len(m.GasRefunds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGlobal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGlobal
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
func (m *GasAdjustment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGlobal
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
			return fmt.Errorf("proto: GasAdjustment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GasAdjustment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGlobal
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
				return ErrInvalidLengthGlobal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGlobal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGlobal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGlobal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGlobal
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
func (m *GasRefund) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGlobal
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
			return fmt.Errorf("proto: GasRefund: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GasRefund: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGlobal
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
				return ErrInvalidLengthGlobal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGlobal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fraction", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGlobal
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
				return ErrInvalidLengthGlobal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGlobal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Fraction.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGlobal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGlobal
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
func skipGlobal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGlobal
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
					return 0, ErrIntOverflowGlobal
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
					return 0, ErrIntOverflowGlobal
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
				return 0, ErrInvalidLengthGlobal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGlobal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGlobal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGlobal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGlobal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGlobal = fmt.Errorf("proto: unexpected end of group")
)
