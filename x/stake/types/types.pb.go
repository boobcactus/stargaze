// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/stake/types/types.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Post struct {
	ID       uint64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	VendorID uint64    `protobuf:"varint,2,opt,name=vendor_id,json=vendorId,proto3" json:"vendor_id,omitempty" yaml:"vendor_id"`
	Body     string    `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty" yaml:"body"`
	VoteEnd  time.Time `protobuf:"bytes,5,opt,name=vote_end,json=voteEnd,proto3,stdtime" json:"vote_end" yaml:"vote_end"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc33dc9db3f2c3c8, []int{0}
}
func (m *Post) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Post.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return m.Size()
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Post) GetVendorID() uint64 {
	if m != nil {
		return m.VendorID
	}
	return 0
}

func (m *Post) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Post) GetVoteEnd() time.Time {
	if m != nil {
		return m.VoteEnd
	}
	return time.Time{}
}

type MsgDelegate struct {
	VendorID      uint64                                        `protobuf:"varint,1,opt,name=vendor_id,json=vendorId,proto3" json:"vendor_id" yaml:"vendor_id"`
	PostID        uint64                                        `protobuf:"varint,2,opt,name=post_id,json=postId,proto3" json:"post_id" yaml:"post_id"`
	DelegatorAddr github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,3,opt,name=delegator_addr,json=delegatorAddr,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"delegator_addr,omitempty" yaml:"delegator_addr"`
	ValidatorAddr github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,4,opt,name=validator_addr,json=validatorAddr,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"validator_addr,omitempty" yaml:"validator_addr"`
	Amount        types.Coin                                    `protobuf:"bytes,5,opt,name=amount,proto3" json:"amount"`
}

func (m *MsgDelegate) Reset()         { *m = MsgDelegate{} }
func (m *MsgDelegate) String() string { return proto.CompactTextString(m) }
func (*MsgDelegate) ProtoMessage()    {}
func (*MsgDelegate) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc33dc9db3f2c3c8, []int{1}
}
func (m *MsgDelegate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDelegate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDelegate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDelegate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDelegate.Merge(m, src)
}
func (m *MsgDelegate) XXX_Size() int {
	return m.Size()
}
func (m *MsgDelegate) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDelegate.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDelegate proto.InternalMessageInfo

func (m *MsgDelegate) GetVendorID() uint64 {
	if m != nil {
		return m.VendorID
	}
	return 0
}

func (m *MsgDelegate) GetPostID() uint64 {
	if m != nil {
		return m.PostID
	}
	return 0
}

func (m *MsgDelegate) GetDelegatorAddr() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.DelegatorAddr
	}
	return nil
}

func (m *MsgDelegate) GetValidatorAddr() github_com_cosmos_cosmos_sdk_types.ValAddress {
	if m != nil {
		return m.ValidatorAddr
	}
	return nil
}

func (m *MsgDelegate) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*Post)(nil), "rocket_protocol.stakebird.x.stake.v1.Post")
	proto.RegisterType((*MsgDelegate)(nil), "rocket_protocol.stakebird.x.stake.v1.MsgDelegate")
}

func init() { proto.RegisterFile("x/stake/types/types.proto", fileDescriptor_bc33dc9db3f2c3c8) }

var fileDescriptor_bc33dc9db3f2c3c8 = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x3f, 0x6f, 0xd3, 0x40,
	0x14, 0x8f, 0x53, 0x93, 0x26, 0x0e, 0xb4, 0xc8, 0x15, 0x52, 0x08, 0x92, 0x2f, 0x72, 0x11, 0x64,
	0xc9, 0x59, 0x29, 0x5b, 0x25, 0x24, 0x6a, 0xca, 0x90, 0x0a, 0x24, 0x64, 0xa1, 0x0e, 0x2c, 0x91,
	0xed, 0x3b, 0xcc, 0x29, 0xb6, 0x2f, 0xf8, 0x2e, 0x51, 0xf3, 0x2d, 0xfa, 0x11, 0xf8, 0x1c, 0x7c,
	0x82, 0x8e, 0x1d, 0x99, 0x0e, 0xe4, 0x2c, 0x88, 0x31, 0x23, 0x13, 0xf2, 0x9d, 0xe3, 0x24, 0x2c,
	0x74, 0xb1, 0xde, 0x9f, 0xdf, 0x7b, 0x3f, 0xff, 0xde, 0xbb, 0x67, 0x3c, 0xbe, 0x72, 0x18, 0xf7,
	0x27, 0xd8, 0xe1, 0x8b, 0x29, 0x66, 0xea, 0x0b, 0xa7, 0x19, 0xe5, 0xd4, 0x7c, 0x9a, 0xd1, 0x70,
	0x82, 0xf9, 0x58, 0x7a, 0x21, 0x8d, 0xa1, 0x04, 0x06, 0x24, 0x43, 0xf0, 0x4a, 0xd9, 0x70, 0x3e,
	0xec, 0xda, 0x21, 0x4d, 0x79, 0x46, 0x02, 0x47, 0xc2, 0x9c, 0x88, 0x46, 0x74, 0x63, 0xa9, 0x4e,
	0xdd, 0x67, 0xbb, 0x98, 0x90, 0xb2, 0x84, 0xb2, 0x81, 0x72, 0xb6, 0x18, 0xbb, 0x20, 0xa2, 0x34,
	0x8a, 0xb1, 0x82, 0x05, 0xb3, 0x4f, 0x0e, 0x27, 0x09, 0x66, 0xdc, 0x4f, 0xa6, 0x0a, 0x60, 0x2f,
	0x35, 0x43, 0x7f, 0x4f, 0x19, 0x37, 0x8f, 0x8d, 0x3a, 0x41, 0x1d, 0xad, 0xa7, 0xf5, 0x75, 0xf7,
	0x28, 0x17, 0xa0, 0x3e, 0x3a, 0x5f, 0x09, 0xd0, 0x5a, 0xf8, 0x49, 0x7c, 0x6a, 0x13, 0x64, 0x7b,
	0x75, 0x82, 0xcc, 0x97, 0x46, 0x6b, 0x8e, 0x53, 0x44, 0xb3, 0x31, 0x41, 0x9d, 0xba, 0xc4, 0xf6,
	0x72, 0x01, 0x9a, 0x97, 0x32, 0x28, 0x2b, 0x1e, 0xaa, 0x8a, 0x0a, 0x66, 0x7b, 0x4d, 0x65, 0x8f,
	0x90, 0x79, 0x6c, 0xe8, 0x01, 0x45, 0x8b, 0xce, 0x5e, 0x4f, 0xeb, 0xb7, 0xdc, 0xc3, 0x95, 0x00,
	0x6d, 0x85, 0x2e, 0xa2, 0xb6, 0x27, 0x93, 0xa6, 0x67, 0x34, 0xe7, 0x94, 0xe3, 0x31, 0x4e, 0x51,
	0xe7, 0x5e, 0x4f, 0xeb, 0xb7, 0x4f, 0xba, 0x50, 0xa9, 0x80, 0x6b, 0x15, 0xf0, 0xc3, 0x5a, 0x85,
	0xfb, 0xe4, 0x46, 0x80, 0xda, 0x4a, 0x80, 0xc3, 0x92, 0xb6, 0xac, 0xb4, 0xaf, 0x7f, 0x00, 0xcd,
	0xdb, 0x2f, 0xdc, 0x37, 0x29, 0xb2, 0xbf, 0xed, 0x19, 0xed, 0x77, 0x2c, 0x3a, 0xc7, 0x31, 0x8e,
	0x7c, 0x8e, 0xcd, 0x8b, 0x6d, 0x1d, 0x4a, 0xf3, 0x60, 0x5b, 0xc7, 0x6f, 0x01, 0x36, 0x80, 0xff,
	0x88, 0x7a, 0x65, 0xec, 0x4f, 0x29, 0xe3, 0x9b, 0x89, 0x3c, 0xcf, 0x05, 0x68, 0x14, 0x33, 0x95,
	0x7d, 0xd6, 0xc9, 0x95, 0x00, 0x07, 0xaa, 0x4b, 0x19, 0xb0, 0xbd, 0x46, 0x61, 0x8d, 0x90, 0xf9,
	0xc5, 0x38, 0x40, 0xea, 0xcf, 0x68, 0x36, 0xf6, 0x11, 0xca, 0xe4, 0x80, 0xee, 0xbb, 0x17, 0x2b,
	0x01, 0x1e, 0xa9, 0x9a, 0xdd, 0xbc, 0xfd, 0x47, 0x80, 0x41, 0x44, 0xf8, 0xe7, 0x59, 0x00, 0x43,
	0x9a, 0x94, 0xeb, 0x5f, 0xbf, 0x02, 0x86, 0x26, 0xe5, 0x1b, 0x38, 0x0b, 0xc3, 0x33, 0x84, 0x32,
	0xcc, 0x98, 0xf7, 0xa0, 0xea, 0x50, 0x44, 0x0a, 0xca, 0xb9, 0x1f, 0x13, 0xb4, 0xa1, 0xd4, 0xff,
	0xa5, 0xdc, 0xcd, 0xdf, 0x95, 0xf2, 0xd2, 0x8f, 0x2b, 0xca, 0xaa, 0x83, 0xa4, 0x1c, 0x1a, 0x0d,
	0x3f, 0xa1, 0xb3, 0x94, 0x97, 0x5b, 0x3d, 0x82, 0xaa, 0x78, 0xcc, 0xd0, 0x04, 0xce, 0x87, 0xf0,
	0x35, 0x25, 0xa9, 0xab, 0x17, 0xeb, 0xf4, 0x4a, 0xe0, 0xa9, 0xfe, 0xeb, 0x2b, 0xd0, 0xdc, 0xb7,
	0x37, 0xb9, 0xa5, 0xdd, 0xe6, 0x96, 0xf6, 0x33, 0xb7, 0xb4, 0xeb, 0xa5, 0x55, 0xbb, 0x5d, 0x5a,
	0xb5, 0xef, 0x4b, 0xab, 0xf6, 0xf1, 0x64, 0xeb, 0x87, 0xd4, 0x69, 0x0d, 0xd6, 0xa7, 0xe5, 0x54,
	0xa7, 0xe5, 0xec, 0xdc, 0x63, 0xd0, 0x90, 0x88, 0x17, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x32,
	0x4f, 0xd8, 0xf0, 0xa7, 0x03, 0x00, 0x00,
}

func (this *MsgDelegate) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgDelegate)
	if !ok {
		that2, ok := that.(MsgDelegate)
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
	if this.VendorID != that1.VendorID {
		return false
	}
	if this.PostID != that1.PostID {
		return false
	}
	if !bytes.Equal(this.DelegatorAddr, that1.DelegatorAddr) {
		return false
	}
	if !bytes.Equal(this.ValidatorAddr, that1.ValidatorAddr) {
		return false
	}
	if !this.Amount.Equal(&that1.Amount) {
		return false
	}
	return true
}
func (m *Post) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Post) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Post) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.VoteEnd, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.VoteEnd):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintTypes(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x1a
	}
	if m.VendorID != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.VendorID))
		i--
		dAtA[i] = 0x10
	}
	if m.ID != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgDelegate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDelegate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDelegate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.ValidatorAddr) > 0 {
		i -= len(m.ValidatorAddr)
		copy(dAtA[i:], m.ValidatorAddr)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ValidatorAddr)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.DelegatorAddr) > 0 {
		i -= len(m.DelegatorAddr)
		copy(dAtA[i:], m.DelegatorAddr)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.DelegatorAddr)))
		i--
		dAtA[i] = 0x1a
	}
	if m.PostID != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.PostID))
		i--
		dAtA[i] = 0x10
	}
	if m.VendorID != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.VendorID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Post) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovTypes(uint64(m.ID))
	}
	if m.VendorID != 0 {
		n += 1 + sovTypes(uint64(m.VendorID))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.VoteEnd)
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func (m *MsgDelegate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VendorID != 0 {
		n += 1 + sovTypes(uint64(m.VendorID))
	}
	if m.PostID != 0 {
		n += 1 + sovTypes(uint64(m.PostID))
	}
	l = len(m.DelegatorAddr)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ValidatorAddr)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Post) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Post: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Post: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VendorID", wireType)
			}
			m.VendorID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VendorID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteEnd", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.VoteEnd, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *MsgDelegate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: MsgDelegate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDelegate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VendorID", wireType)
			}
			m.VendorID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VendorID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostID", wireType)
			}
			m.PostID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PostID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatorAddr = append(m.DelegatorAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.DelegatorAddr == nil {
				m.DelegatorAddr = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddr = append(m.ValidatorAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.ValidatorAddr == nil {
				m.ValidatorAddr = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
