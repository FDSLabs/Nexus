// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nexus/coinswap/v1/coinswap.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
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

// Input defines the properties of order's input
type Input struct {
	Address string     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Coin    types.Coin `protobuf:"bytes,2,opt,name=coin,proto3" json:"coin"`
}

func (m *Input) Reset()         { *m = Input{} }
func (m *Input) String() string { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()    {}
func (*Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_003129b03f6f7188, []int{0}
}
func (m *Input) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Input.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Input) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Input.Merge(m, src)
}
func (m *Input) XXX_Size() int {
	return m.Size()
}
func (m *Input) XXX_DiscardUnknown() {
	xxx_messageInfo_Input.DiscardUnknown(m)
}

var xxx_messageInfo_Input proto.InternalMessageInfo

// Output defines the properties of order's output
type Output struct {
	Address string     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Coin    types.Coin `protobuf:"bytes,2,opt,name=coin,proto3" json:"coin"`
}

func (m *Output) Reset()         { *m = Output{} }
func (m *Output) String() string { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()    {}
func (*Output) Descriptor() ([]byte, []int) {
	return fileDescriptor_003129b03f6f7188, []int{1}
}
func (m *Output) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Output) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Output.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Output) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Output.Merge(m, src)
}
func (m *Output) XXX_Size() int {
	return m.Size()
}
func (m *Output) XXX_DiscardUnknown() {
	xxx_messageInfo_Output.DiscardUnknown(m)
}

var xxx_messageInfo_Output proto.InternalMessageInfo

type Pool struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// denom of base coin of the pool
	StandardDenom string `protobuf:"bytes,2,opt,name=standard_denom,json=standardDenom,proto3" json:"standard_denom,omitempty"`
	// denom of counterparty coin of the pool
	CounterpartyDenom string `protobuf:"bytes,3,opt,name=counterparty_denom,json=counterpartyDenom,proto3" json:"counterparty_denom,omitempty"`
	// escrow account for deposit tokens
	EscrowAddress string `protobuf:"bytes,4,opt,name=escrow_address,json=escrowAddress,proto3" json:"escrow_address,omitempty"`
	// denom of the liquidity pool coin
	LptDenom string `protobuf:"bytes,5,opt,name=lpt_denom,json=lptDenom,proto3" json:"lpt_denom,omitempty"`
}

func (m *Pool) Reset()         { *m = Pool{} }
func (m *Pool) String() string { return proto.CompactTextString(m) }
func (*Pool) ProtoMessage()    {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_003129b03f6f7188, []int{2}
}
func (m *Pool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(m, src)
}
func (m *Pool) XXX_Size() int {
	return m.Size()
}
func (m *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

// Params defines token module's parameters
type Params struct {
	Fee                    cosmossdk_io_math.LegacyDec              `protobuf:"bytes,1,opt,name=fee,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"fee"`
	PoolCreationFee        types.Coin                               `protobuf:"bytes,2,opt,name=pool_creation_fee,json=poolCreationFee,proto3" json:"pool_creation_fee"`
	TaxRate                cosmossdk_io_math.LegacyDec              `protobuf:"bytes,3,opt,name=tax_rate,json=taxRate,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"tax_rate"`
	MaxStandardCoinPerPool cosmossdk_io_math.Int                    `protobuf:"bytes,4,opt,name=max_standard_coin_per_pool,json=maxStandardCoinPerPool,proto3,customtype=cosmossdk.io/math.Int" json:"max_standard_coin_per_pool"`
	MaxSwapAmount          github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,6,rep,name=max_swap_amount,json=maxSwapAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"max_swap_amount"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_003129b03f6f7188, []int{3}
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

func init() {
	proto.RegisterType((*Input)(nil), "nexus.coinswap.v1.Input")
	proto.RegisterType((*Output)(nil), "nexus.coinswap.v1.Output")
	proto.RegisterType((*Pool)(nil), "nexus.coinswap.v1.Pool")
	proto.RegisterType((*Params)(nil), "nexus.coinswap.v1.Params")
}

func init() { proto.RegisterFile("nexus/coinswap/v1/coinswap.proto", fileDescriptor_003129b03f6f7188) }

var fileDescriptor_003129b03f6f7188 = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xcf, 0x6b, 0x13, 0x4f,
	0x14, 0xdf, 0x6d, 0xd3, 0xb4, 0x99, 0xef, 0xb7, 0x2d, 0x59, 0xfc, 0x91, 0xb6, 0xb0, 0x09, 0xc1,
	0x42, 0x0c, 0x64, 0x97, 0x58, 0xf0, 0xe0, 0xad, 0x69, 0x88, 0x16, 0x8a, 0xc6, 0x14, 0x14, 0x3c,
	0xb8, 0xbc, 0xec, 0x8e, 0xe9, 0xd2, 0xec, 0xce, 0xb2, 0x33, 0x49, 0x36, 0x7f, 0x42, 0x3d, 0x79,
	0x14, 0x4f, 0x3d, 0x8a, 0xa7, 0x1c, 0xfc, 0x23, 0x72, 0x2c, 0x9e, 0xc4, 0x43, 0xd5, 0xe4, 0x10,
	0xff, 0x0c, 0x99, 0x1f, 0x89, 0x45, 0x41, 0x41, 0xf0, 0xb2, 0x3b, 0xf3, 0xe6, 0xbd, 0xcf, 0x8f,
	0x37, 0xf3, 0x50, 0x21, 0xc4, 0x49, 0x8f, 0xda, 0x2e, 0xf1, 0x43, 0x3a, 0x80, 0xc8, 0xee, 0x57,
	0x17, 0x6b, 0x2b, 0x8a, 0x09, 0x23, 0x46, 0x56, 0x64, 0x58, 0x8b, 0x68, 0xbf, 0xba, 0x6d, 0xba,
	0x84, 0x06, 0x84, 0xda, 0x6d, 0xa0, 0xd8, 0xee, 0x57, 0xdb, 0x98, 0x81, 0x2c, 0x93, 0x25, 0xdb,
	0xd7, 0x3a, 0xa4, 0x43, 0xc4, 0xd2, 0xe6, 0x2b, 0x15, 0xdd, 0x92, 0x55, 0x8e, 0x3c, 0x90, 0x1b,
	0x75, 0x94, 0x85, 0xc0, 0x0f, 0x89, 0x2d, 0xbe, 0x32, 0x54, 0x7c, 0x82, 0x56, 0x0e, 0xc3, 0xa8,
	0xc7, 0x8c, 0x1c, 0x5a, 0x05, 0xcf, 0x8b, 0x31, 0xa5, 0x39, 0xbd, 0xa0, 0x97, 0x32, 0xad, 0xf9,
	0xd6, 0xd8, 0x43, 0x29, 0x4e, 0x9a, 0x5b, 0x2a, 0xe8, 0xa5, 0xff, 0xee, 0x6c, 0x59, 0x0a, 0x92,
	0xab, 0xb2, 0x94, 0x2a, 0xeb, 0x80, 0xf8, 0x61, 0x2d, 0x35, 0xbe, 0xcc, 0x6b, 0x2d, 0x91, 0x5c,
	0x7c, 0x8a, 0xd2, 0x8f, 0x7a, 0xec, 0x1f, 0x00, 0x8f, 0x74, 0x94, 0x6a, 0x12, 0xd2, 0x35, 0x36,
	0xd0, 0x92, 0xef, 0x29, 0xc8, 0x25, 0xdf, 0x33, 0x76, 0xd1, 0x06, 0x65, 0x10, 0x7a, 0x10, 0x7b,
	0x8e, 0x87, 0x43, 0x12, 0x08, 0xdc, 0x4c, 0x6b, 0x7d, 0x1e, 0xad, 0xf3, 0xa0, 0x51, 0x41, 0x86,
	0x4b, 0x7a, 0x21, 0xc3, 0x71, 0x04, 0x31, 0x1b, 0xaa, 0xd4, 0x65, 0x91, 0x9a, 0xbd, 0x7a, 0x22,
	0xd3, 0x77, 0xd1, 0x06, 0xa6, 0x6e, 0x4c, 0x06, 0xce, 0xdc, 0x44, 0x4a, 0xa2, 0xca, 0xe8, 0xbe,
	0xb2, 0xb2, 0x83, 0x32, 0xdd, 0x88, 0x29, 0xb0, 0x15, 0x91, 0xb1, 0xd6, 0x8d, 0x98, 0xc0, 0x28,
	0x9e, 0xa5, 0x50, 0xba, 0x09, 0x31, 0x04, 0xd4, 0x78, 0x80, 0x96, 0x5f, 0x60, 0x2c, 0x55, 0xd7,
	0xee, 0x72, 0x5b, 0x9f, 0x2e, 0xf3, 0x3b, 0xd2, 0x38, 0xf5, 0x4e, 0x2d, 0x9f, 0xd8, 0x01, 0xb0,
	0x13, 0xeb, 0x08, 0x77, 0xc0, 0x1d, 0xd6, 0xb1, 0xfb, 0xe1, 0x7d, 0x05, 0xa9, 0xbe, 0xd4, 0xb1,
	0xfb, 0x76, 0x36, 0x2a, 0xeb, 0x2d, 0x0e, 0x61, 0x34, 0x51, 0x36, 0x22, 0xa4, 0xeb, 0xb8, 0x31,
	0x06, 0xe6, 0x93, 0xd0, 0xe1, 0xb8, 0x7f, 0xec, 0x64, 0x86, 0x53, 0x4a, 0x94, 0x4d, 0x5e, 0x7e,
	0xa0, 0xaa, 0x1b, 0x18, 0x1b, 0x8f, 0xd1, 0x1a, 0x83, 0xc4, 0x89, 0x81, 0x61, 0xd9, 0x8f, 0xbf,
	0x16, 0xb8, 0xca, 0x20, 0x69, 0x01, 0xc3, 0xc6, 0x73, 0xb4, 0x1d, 0x40, 0xe2, 0x2c, 0xee, 0x85,
	0xdf, 0xa0, 0x13, 0xe1, 0xd8, 0xe1, 0xdc, 0xb2, 0x93, 0xb5, 0xa2, 0x22, 0xb9, 0xfe, 0x2b, 0xc9,
	0x61, 0xc8, 0x24, 0xe0, 0x8d, 0x00, 0x92, 0x63, 0x05, 0xc2, 0x7d, 0x34, 0x71, 0x2c, 0xde, 0xc0,
	0x99, 0x8e, 0x36, 0x05, 0xc1, 0x00, 0x22, 0x07, 0x02, 0x7e, 0x7b, 0xb9, 0x74, 0x61, 0xf9, 0xf7,
	0x3d, 0x68, 0x70, 0xc2, 0x77, 0x9f, 0xf3, 0xa5, 0x8e, 0xcf, 0x4e, 0x7a, 0x6d, 0xcb, 0x25, 0x81,
	0x1a, 0x13, 0xf5, 0xab, 0x50, 0xef, 0xd4, 0x66, 0xc3, 0x08, 0x53, 0x51, 0x40, 0xdf, 0xcc, 0x46,
	0xe5, 0xff, 0xbb, 0xc2, 0xb0, 0x70, 0x40, 0xa5, 0xa8, 0x75, 0x2e, 0x6a, 0x00, 0xd1, 0xbe, 0xe0,
	0xbd, 0x77, 0xeb, 0xf5, 0x79, 0x5e, 0xfb, 0x76, 0x9e, 0xd7, 0x5f, 0xce, 0x46, 0xe5, 0x9b, 0x72,
	0xde, 0x93, 0x1f, 0x13, 0x2f, 0x1f, 0x40, 0xed, 0xfe, 0xf8, 0xab, 0xa9, 0x8d, 0x27, 0xa6, 0x7e,
	0x31, 0x31, 0xf5, 0x2f, 0x13, 0x53, 0x7f, 0x35, 0x35, 0xb5, 0x8b, 0xa9, 0xa9, 0x7d, 0x9c, 0x9a,
	0xda, 0xb3, 0xdb, 0x57, 0x24, 0x35, 0xea, 0xc7, 0x47, 0xd0, 0xa6, 0xf6, 0xc3, 0x9f, 0x91, 0x84,
	0xb2, 0x76, 0x5a, 0xcc, 0xef, 0xde, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf9, 0xb0, 0xbc, 0x30,
	0x5a, 0x04, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Fee.Equal(that1.Fee) {
		return false
	}
	if !this.PoolCreationFee.Equal(&that1.PoolCreationFee) {
		return false
	}
	if !this.TaxRate.Equal(that1.TaxRate) {
		return false
	}
	if !this.MaxStandardCoinPerPool.Equal(that1.MaxStandardCoinPerPool) {
		return false
	}
	if len(this.MaxSwapAmount) != len(that1.MaxSwapAmount) {
		return false
	}
	for i := range this.MaxSwapAmount {
		if !this.MaxSwapAmount[i].Equal(&that1.MaxSwapAmount[i]) {
			return false
		}
	}
	return true
}
func (m *Input) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Input) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Input) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Coin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCoinswap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintCoinswap(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Output) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Output) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Output) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Coin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCoinswap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintCoinswap(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Pool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LptDenom) > 0 {
		i -= len(m.LptDenom)
		copy(dAtA[i:], m.LptDenom)
		i = encodeVarintCoinswap(dAtA, i, uint64(len(m.LptDenom)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.EscrowAddress) > 0 {
		i -= len(m.EscrowAddress)
		copy(dAtA[i:], m.EscrowAddress)
		i = encodeVarintCoinswap(dAtA, i, uint64(len(m.EscrowAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.CounterpartyDenom) > 0 {
		i -= len(m.CounterpartyDenom)
		copy(dAtA[i:], m.CounterpartyDenom)
		i = encodeVarintCoinswap(dAtA, i, uint64(len(m.CounterpartyDenom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.StandardDenom) > 0 {
		i -= len(m.StandardDenom)
		copy(dAtA[i:], m.StandardDenom)
		i = encodeVarintCoinswap(dAtA, i, uint64(len(m.StandardDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintCoinswap(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
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
	if len(m.MaxSwapAmount) > 0 {
		for iNdEx := len(m.MaxSwapAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MaxSwapAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCoinswap(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	{
		size := m.MaxStandardCoinPerPool.Size()
		i -= size
		if _, err := m.MaxStandardCoinPerPool.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCoinswap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.TaxRate.Size()
		i -= size
		if _, err := m.TaxRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCoinswap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.PoolCreationFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCoinswap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Fee.Size()
		i -= size
		if _, err := m.Fee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCoinswap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintCoinswap(dAtA []byte, offset int, v uint64) int {
	offset -= sovCoinswap(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Input) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovCoinswap(uint64(l))
	}
	l = m.Coin.Size()
	n += 1 + l + sovCoinswap(uint64(l))
	return n
}

func (m *Output) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovCoinswap(uint64(l))
	}
	l = m.Coin.Size()
	n += 1 + l + sovCoinswap(uint64(l))
	return n
}

func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovCoinswap(uint64(l))
	}
	l = len(m.StandardDenom)
	if l > 0 {
		n += 1 + l + sovCoinswap(uint64(l))
	}
	l = len(m.CounterpartyDenom)
	if l > 0 {
		n += 1 + l + sovCoinswap(uint64(l))
	}
	l = len(m.EscrowAddress)
	if l > 0 {
		n += 1 + l + sovCoinswap(uint64(l))
	}
	l = len(m.LptDenom)
	if l > 0 {
		n += 1 + l + sovCoinswap(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Fee.Size()
	n += 1 + l + sovCoinswap(uint64(l))
	l = m.PoolCreationFee.Size()
	n += 1 + l + sovCoinswap(uint64(l))
	l = m.TaxRate.Size()
	n += 1 + l + sovCoinswap(uint64(l))
	l = m.MaxStandardCoinPerPool.Size()
	n += 1 + l + sovCoinswap(uint64(l))
	if len(m.MaxSwapAmount) > 0 {
		for _, e := range m.MaxSwapAmount {
			l = e.Size()
			n += 1 + l + sovCoinswap(uint64(l))
		}
	}
	return n
}

func sovCoinswap(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCoinswap(x uint64) (n int) {
	return sovCoinswap(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Input) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoinswap
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
			return fmt.Errorf("proto: Input: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Input: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinswap
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
				return ErrInvalidLengthCoinswap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoinswap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinswap
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
				return ErrInvalidLengthCoinswap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCoinswap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCoinswap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCoinswap
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
func (m *Output) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoinswap
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
			return fmt.Errorf("proto: Output: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Output: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinswap
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
				return ErrInvalidLengthCoinswap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoinswap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinswap
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
				return ErrInvalidLengthCoinswap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCoinswap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCoinswap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCoinswap
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
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoinswap
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
			return fmt.Errorf("proto: Pool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinswap
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
				return ErrInvalidLengthCoinswap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoinswap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StandardDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinswap
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
				return ErrInvalidLengthCoinswap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoinswap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StandardDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CounterpartyDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinswap
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
				return ErrInvalidLengthCoinswap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoinswap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CounterpartyDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EscrowAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinswap
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
				return ErrInvalidLengthCoinswap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoinswap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EscrowAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LptDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinswap
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
				return ErrInvalidLengthCoinswap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoinswap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LptDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCoinswap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCoinswap
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoinswap
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
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinswap
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
				return ErrInvalidLengthCoinswap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoinswap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolCreationFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinswap
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
				return ErrInvalidLengthCoinswap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCoinswap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PoolCreationFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaxRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinswap
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
				return ErrInvalidLengthCoinswap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoinswap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TaxRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxStandardCoinPerPool", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinswap
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
				return ErrInvalidLengthCoinswap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoinswap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxStandardCoinPerPool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSwapAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinswap
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
				return ErrInvalidLengthCoinswap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCoinswap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MaxSwapAmount = append(m.MaxSwapAmount, types.Coin{})
			if err := m.MaxSwapAmount[len(m.MaxSwapAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCoinswap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCoinswap
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
func skipCoinswap(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCoinswap
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
					return 0, ErrIntOverflowCoinswap
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
					return 0, ErrIntOverflowCoinswap
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
				return 0, ErrInvalidLengthCoinswap
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCoinswap
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCoinswap
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCoinswap        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCoinswap          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCoinswap = fmt.Errorf("proto: unexpected end of group")
)
