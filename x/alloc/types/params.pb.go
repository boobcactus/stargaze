// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: publicawesome/stargaze/alloc/v1beta1/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

type WeightedAddress struct {
	Address string                                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	Weight  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=weight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"weight" yaml:"weight"`
}

func (m *WeightedAddress) Reset()         { *m = WeightedAddress{} }
func (m *WeightedAddress) String() string { return proto.CompactTextString(m) }
func (*WeightedAddress) ProtoMessage()    {}
func (*WeightedAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_7716a10c05d88367, []int{0}
}
func (m *WeightedAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WeightedAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WeightedAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WeightedAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeightedAddress.Merge(m, src)
}
func (m *WeightedAddress) XXX_Size() int {
	return m.Size()
}
func (m *WeightedAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_WeightedAddress.DiscardUnknown(m)
}

var xxx_messageInfo_WeightedAddress proto.InternalMessageInfo

func (m *WeightedAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type DistributionProportions struct {
	NftIncentives    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=nft_incentives,json=nftIncentives,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"nft_incentives" yaml:"nft_incentives"`
	DeveloperRewards github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=developer_rewards,json=developerRewards,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"developer_rewards" yaml:"developer_rewards"`
	CommunityPool    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=community_pool,json=communityPool,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"community_pool" yaml:"community_pool"`
}

func (m *DistributionProportions) Reset()         { *m = DistributionProportions{} }
func (m *DistributionProportions) String() string { return proto.CompactTextString(m) }
func (*DistributionProportions) ProtoMessage()    {}
func (*DistributionProportions) Descriptor() ([]byte, []int) {
	return fileDescriptor_7716a10c05d88367, []int{1}
}
func (m *DistributionProportions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DistributionProportions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DistributionProportions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DistributionProportions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributionProportions.Merge(m, src)
}
func (m *DistributionProportions) XXX_Size() int {
	return m.Size()
}
func (m *DistributionProportions) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributionProportions.DiscardUnknown(m)
}

var xxx_messageInfo_DistributionProportions proto.InternalMessageInfo

type Params struct {
	// distribution_proportions defines the proportion of the minted denom
	DistributionProportions DistributionProportions `protobuf:"bytes,1,opt,name=distribution_proportions,json=distributionProportions,proto3" json:"distribution_proportions"`
	// addresses to receive developer rewards
	WeightedDeveloperRewardsReceivers []WeightedAddress `protobuf:"bytes,2,rep,name=weighted_developer_rewards_receivers,json=weightedDeveloperRewardsReceivers,proto3" json:"weighted_developer_rewards_receivers" yaml:"developer_rewards_receiver"`
	// addresses to receive incentive rewards
	WeightedIncentivesRewardsReceivers []WeightedAddress `protobuf:"bytes,3,rep,name=weighted_incentives_rewards_receivers,json=weightedIncentivesRewardsReceivers,proto3" json:"weighted_incentives_rewards_receivers" yaml:"incentives_rewards_receiver"`
	// SupplementAmount is the amount to be supplemented from the pool on top of
	// newly minted coins.
	SupplementAmount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=supplement_amount,json=supplementAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"supplement_amount" yaml:"supplement_amount"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_7716a10c05d88367, []int{2}
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

func (m *Params) GetDistributionProportions() DistributionProportions {
	if m != nil {
		return m.DistributionProportions
	}
	return DistributionProportions{}
}

func (m *Params) GetWeightedDeveloperRewardsReceivers() []WeightedAddress {
	if m != nil {
		return m.WeightedDeveloperRewardsReceivers
	}
	return nil
}

func (m *Params) GetWeightedIncentivesRewardsReceivers() []WeightedAddress {
	if m != nil {
		return m.WeightedIncentivesRewardsReceivers
	}
	return nil
}

func (m *Params) GetSupplementAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.SupplementAmount
	}
	return nil
}

func init() {
	proto.RegisterType((*WeightedAddress)(nil), "publicawesome.stargaze.alloc.v1beta1.WeightedAddress")
	proto.RegisterType((*DistributionProportions)(nil), "publicawesome.stargaze.alloc.v1beta1.DistributionProportions")
	proto.RegisterType((*Params)(nil), "publicawesome.stargaze.alloc.v1beta1.Params")
}

func init() {
	proto.RegisterFile("publicawesome/stargaze/alloc/v1beta1/params.proto", fileDescriptor_7716a10c05d88367)
}

var fileDescriptor_7716a10c05d88367 = []byte{
	// 602 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x41, 0x6f, 0xd3, 0x3e,
	0x18, 0xc6, 0x9b, 0x75, 0xff, 0xfe, 0x85, 0xa7, 0x8e, 0x2d, 0x02, 0x2d, 0xec, 0x90, 0x30, 0x6b,
	0xa0, 0x82, 0x58, 0xa2, 0x8e, 0x71, 0x41, 0x42, 0x68, 0xa5, 0x12, 0x02, 0x21, 0x54, 0xe5, 0x32,
	0x89, 0x4b, 0xe4, 0x24, 0x5e, 0x67, 0x91, 0xc4, 0x91, 0xed, 0xb4, 0x94, 0x03, 0x1f, 0x80, 0x13,
	0x07, 0x0e, 0x7c, 0x06, 0x6e, 0x48, 0x7c, 0x88, 0x1d, 0x77, 0x44, 0x1c, 0x0a, 0x6a, 0xbf, 0x41,
	0x3f, 0x01, 0x4a, 0x9c, 0xa4, 0xeb, 0x4a, 0x51, 0xa7, 0x9d, 0xe2, 0x38, 0x79, 0x9e, 0xf7, 0xe7,
	0xf7, 0x7d, 0x64, 0xd0, 0x8c, 0x13, 0x37, 0x20, 0x1e, 0xea, 0x63, 0x4e, 0x43, 0x6c, 0x71, 0x81,
	0x58, 0x17, 0xbd, 0xc7, 0x16, 0x0a, 0x02, 0xea, 0x59, 0xbd, 0xa6, 0x8b, 0x05, 0x6a, 0x5a, 0x31,
	0x62, 0x28, 0xe4, 0x66, 0xcc, 0xa8, 0xa0, 0xea, 0xee, 0x8c, 0xc4, 0x2c, 0x24, 0x66, 0x26, 0x31,
	0x73, 0xc9, 0xf6, 0x8d, 0x2e, 0xed, 0xd2, 0x4c, 0x60, 0xa5, 0x2b, 0xa9, 0xdd, 0xd6, 0x3d, 0xca,
	0x43, 0xca, 0x2d, 0x17, 0x71, 0x5c, 0xba, 0x7b, 0x94, 0x44, 0xf2, 0x3b, 0xfc, 0xa2, 0x80, 0xeb,
	0x47, 0x98, 0x74, 0x4f, 0x04, 0xf6, 0x0f, 0x7d, 0x9f, 0x61, 0xce, 0xd5, 0x07, 0xe0, 0x7f, 0x24,
	0x97, 0x9a, 0x72, 0x5b, 0x69, 0x5c, 0x6b, 0xa9, 0x93, 0xa1, 0xb1, 0x3e, 0x40, 0x61, 0xf0, 0x18,
	0xe6, 0x1f, 0xa0, 0x5d, 0xfc, 0xa2, 0x1e, 0x81, 0x5a, 0x3f, 0x33, 0xd0, 0x56, 0xb2, 0x9f, 0x9f,
	0x9e, 0x0e, 0x8d, 0xca, 0xcf, 0xa1, 0x71, 0xb7, 0x4b, 0xc4, 0x49, 0xe2, 0x9a, 0x1e, 0x0d, 0xad,
	0x1c, 0x42, 0x3e, 0xf6, 0xb8, 0xff, 0xd6, 0x12, 0x83, 0x18, 0x73, 0xb3, 0x8d, 0xbd, 0xc9, 0xd0,
	0xa8, 0x4b, 0x6b, 0xe9, 0x02, 0xed, 0xdc, 0x0e, 0x4e, 0x56, 0xc0, 0x56, 0x9b, 0x70, 0xc1, 0x88,
	0x9b, 0x08, 0x42, 0xa3, 0x0e, 0xa3, 0x31, 0x65, 0xe9, 0x8a, 0xab, 0x11, 0x58, 0x8f, 0x8e, 0x85,
	0x43, 0x22, 0x0f, 0x47, 0x82, 0xf4, 0x70, 0x41, 0xfa, 0xfc, 0xd2, 0xc5, 0x6f, 0xca, 0xe2, 0xb3,
	0x6e, 0xd0, 0xae, 0x47, 0xc7, 0xe2, 0x45, 0xf9, 0xae, 0xf6, 0xc1, 0xa6, 0x8f, 0x7b, 0x38, 0xa0,
	0x31, 0x66, 0x0e, 0xc3, 0x7d, 0xc4, 0x7c, 0x9e, 0x9f, 0xf7, 0xe5, 0xa5, 0x4b, 0x6a, 0xb2, 0xe4,
	0x9c, 0x21, 0xb4, 0x37, 0xca, 0x3d, 0x5b, 0x6e, 0xa5, 0x07, 0xf5, 0x68, 0x18, 0x26, 0x11, 0x11,
	0x03, 0x27, 0xa6, 0x34, 0xd0, 0xaa, 0x57, 0x3b, 0xe8, 0xac, 0x1b, 0xb4, 0xeb, 0xe5, 0x46, 0x27,
	0x7d, 0xff, 0xf8, 0x1f, 0xa8, 0x75, 0xb2, 0xf0, 0xa9, 0x1f, 0x80, 0xe6, 0x9f, 0x6b, 0xbf, 0x13,
	0x4f, 0xfb, 0x9f, 0x75, 0x7b, 0x6d, 0xff, 0x89, 0xb9, 0x4c, 0x32, 0xcd, 0x05, 0x43, 0x6c, 0xad,
	0xa6, 0x67, 0xb0, 0xb7, 0xfc, 0x05, 0x33, 0xfe, 0xa6, 0x80, 0xdd, 0x7e, 0x1e, 0x4d, 0x67, 0xae,
	0x59, 0x0e, 0xc3, 0x1e, 0x26, 0x3d, 0xcc, 0xd2, 0x39, 0x54, 0x1b, 0x6b, 0xfb, 0x8f, 0x96, 0x83,
	0xb9, 0x10, 0xf6, 0xd6, 0xbd, 0x14, 0x62, 0x32, 0x34, 0x76, 0x16, 0x0c, 0xa5, 0xac, 0x03, 0xed,
	0x9d, 0x82, 0xa6, 0x7d, 0x61, 0x4a, 0x76, 0x81, 0xa2, 0x7e, 0x57, 0xc0, 0x9d, 0x92, 0x79, 0x9a,
	0xa7, 0xbf, 0x40, 0x57, 0xaf, 0x02, 0x7d, 0x3f, 0x87, 0x86, 0x12, 0xfa, 0x1f, 0x85, 0xa0, 0x0d,
	0x0b, 0x9e, 0x69, 0x9c, 0xe7, 0xb0, 0x3f, 0x2b, 0x60, 0x93, 0x27, 0x71, 0x1c, 0xe0, 0x10, 0x47,
	0xc2, 0x41, 0x21, 0x4d, 0x22, 0xa1, 0xad, 0x66, 0x88, 0xb7, 0x4c, 0x19, 0x28, 0x33, 0xbd, 0x42,
	0x4a, 0xa2, 0x67, 0x94, 0x44, 0xad, 0x57, 0x39, 0x46, 0x1e, 0xe8, 0x39, 0x07, 0xf8, 0xf5, 0x97,
	0xd1, 0x58, 0x22, 0xa0, 0xa9, 0x19, 0xb7, 0x37, 0xa6, 0xfa, 0xc3, 0x4c, 0xde, 0x7a, 0x7d, 0x3a,
	0xd2, 0x95, 0xb3, 0x91, 0xae, 0xfc, 0x1e, 0xe9, 0xca, 0xa7, 0xb1, 0x5e, 0x39, 0x1b, 0xeb, 0x95,
	0x1f, 0x63, 0xbd, 0xf2, 0xe6, 0xe0, 0x9c, 0xab, 0xec, 0xe0, 0xde, 0xdc, 0x8d, 0xda, 0x6b, 0x1e,
	0x58, 0xef, 0xf2, 0x7b, 0x35, 0xab, 0xe3, 0xd6, 0xb2, 0x3b, 0xef, 0xe1, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x05, 0xea, 0x3c, 0x97, 0x84, 0x05, 0x00, 0x00,
}

func (m *WeightedAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WeightedAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WeightedAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Weight.Size()
		i -= size
		if _, err := m.Weight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DistributionProportions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DistributionProportions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DistributionProportions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.CommunityPool.Size()
		i -= size
		if _, err := m.CommunityPool.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.DeveloperRewards.Size()
		i -= size
		if _, err := m.DeveloperRewards.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.NftIncentives.Size()
		i -= size
		if _, err := m.NftIncentives.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
	if len(m.SupplementAmount) > 0 {
		for iNdEx := len(m.SupplementAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SupplementAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.WeightedIncentivesRewardsReceivers) > 0 {
		for iNdEx := len(m.WeightedIncentivesRewardsReceivers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WeightedIncentivesRewardsReceivers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.WeightedDeveloperRewardsReceivers) > 0 {
		for iNdEx := len(m.WeightedDeveloperRewardsReceivers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WeightedDeveloperRewardsReceivers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.DistributionProportions.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WeightedAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.Weight.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func (m *DistributionProportions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.NftIncentives.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.DeveloperRewards.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.CommunityPool.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.DistributionProportions.Size()
	n += 1 + l + sovParams(uint64(l))
	if len(m.WeightedDeveloperRewardsReceivers) > 0 {
		for _, e := range m.WeightedDeveloperRewardsReceivers {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.WeightedIncentivesRewardsReceivers) > 0 {
		for _, e := range m.WeightedIncentivesRewardsReceivers {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.SupplementAmount) > 0 {
		for _, e := range m.SupplementAmount {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WeightedAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: WeightedAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WeightedAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Weight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *DistributionProportions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: DistributionProportions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DistributionProportions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftIncentives", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NftIncentives.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeveloperRewards", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DeveloperRewards.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommunityPool", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommunityPool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
				return ErrIntOverflowParams
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
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionProportions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DistributionProportions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WeightedDeveloperRewardsReceivers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WeightedDeveloperRewardsReceivers = append(m.WeightedDeveloperRewardsReceivers, WeightedAddress{})
			if err := m.WeightedDeveloperRewardsReceivers[len(m.WeightedDeveloperRewardsReceivers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WeightedIncentivesRewardsReceivers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WeightedIncentivesRewardsReceivers = append(m.WeightedIncentivesRewardsReceivers, WeightedAddress{})
			if err := m.WeightedIncentivesRewardsReceivers[len(m.WeightedIncentivesRewardsReceivers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SupplementAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SupplementAmount = append(m.SupplementAmount, types.Coin{})
			if err := m.SupplementAmount[len(m.SupplementAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
