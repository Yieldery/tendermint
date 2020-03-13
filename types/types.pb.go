// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/types.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// BlockIdFlag indicates which BlcokID the signature is for
type BlockIDFlag int32

const (
	BLOCKD_ID_FLAG_UNKNOWN BlockIDFlag = 0
	BlockIDFlagAbsent      BlockIDFlag = 1
	BlockIDFlagCommit      BlockIDFlag = 2
	BlockIDFlagNil         BlockIDFlag = 3
)

var BlockIDFlag_name = map[int32]string{
	0: "BLOCKD_ID_FLAG_UNKNOWN",
	1: "BLOCK_ID_FLAG_ABSENT",
	2: "BLOCK_ID_FLAG_COMMIT",
	3: "BLOCK_ID_FLAG_NIL",
}

var BlockIDFlag_value = map[string]int32{
	"BLOCKD_ID_FLAG_UNKNOWN": 0,
	"BLOCK_ID_FLAG_ABSENT":   1,
	"BLOCK_ID_FLAG_COMMIT":   2,
	"BLOCK_ID_FLAG_NIL":      3,
}

func (BlockIDFlag) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2c0f90c600ad7e2e, []int{0}
}

// SignedMsgType is a type of signed message in the consensus.
type SignedMsgType int32

const (
	SIGNED_MSG_TYPE_UNKNOWN SignedMsgType = 0
	PrevoteType             SignedMsgType = 1
	PrecommitType           SignedMsgType = 2
	ProposalType            SignedMsgType = 3
)

var SignedMsgType_name = map[int32]string{
	0: "SIGNED_MSG_TYPE_UNKNOWN",
	1: "PREVOTE_TYPE",
	2: "PRECOMMIT_TYPE",
	3: "PROPOSAL_TYPE",
}

var SignedMsgType_value = map[string]int32{
	"SIGNED_MSG_TYPE_UNKNOWN": 0,
	"PREVOTE_TYPE":            1,
	"PRECOMMIT_TYPE":          2,
	"PROPOSAL_TYPE":           3,
}

func (SignedMsgType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2c0f90c600ad7e2e, []int{1}
}

// PartsetHeader
type PartSetHeader struct {
	Total uint32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Hash  []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *PartSetHeader) Reset()         { *m = PartSetHeader{} }
func (m *PartSetHeader) String() string { return proto.CompactTextString(m) }
func (*PartSetHeader) ProtoMessage()    {}
func (*PartSetHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c0f90c600ad7e2e, []int{0}
}
func (m *PartSetHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PartSetHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PartSetHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PartSetHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartSetHeader.Merge(m, src)
}
func (m *PartSetHeader) XXX_Size() int {
	return m.Size()
}
func (m *PartSetHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_PartSetHeader.DiscardUnknown(m)
}

var xxx_messageInfo_PartSetHeader proto.InternalMessageInfo

func (m *PartSetHeader) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *PartSetHeader) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

// BlockID
type BlockID struct {
	Hash        []byte        `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	PartsHeader PartSetHeader `protobuf:"bytes,2,opt,name=parts_header,json=partsHeader,proto3" json:"parts_header"`
}

func (m *BlockID) Reset()         { *m = BlockID{} }
func (m *BlockID) String() string { return proto.CompactTextString(m) }
func (*BlockID) ProtoMessage()    {}
func (*BlockID) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c0f90c600ad7e2e, []int{1}
}
func (m *BlockID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockID.Merge(m, src)
}
func (m *BlockID) XXX_Size() int {
	return m.Size()
}
func (m *BlockID) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockID.DiscardUnknown(m)
}

var xxx_messageInfo_BlockID proto.InternalMessageInfo

func (m *BlockID) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *BlockID) GetPartsHeader() PartSetHeader {
	if m != nil {
		return m.PartsHeader
	}
	return PartSetHeader{}
}

func init() {
	proto.RegisterEnum("tendermint.types.BlockIDFlag", BlockIDFlag_name, BlockIDFlag_value)
	proto.RegisterEnum("tendermint.types.SignedMsgType", SignedMsgType_name, SignedMsgType_value)
	proto.RegisterType((*PartSetHeader)(nil), "tendermint.types.PartSetHeader")
	proto.RegisterType((*BlockID)(nil), "tendermint.types.BlockID")
}

func init() { proto.RegisterFile("types/types.proto", fileDescriptor_2c0f90c600ad7e2e) }

var fileDescriptor_2c0f90c600ad7e2e = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x4f, 0x8f, 0x93, 0x40,
	0x18, 0xc6, 0x99, 0x6e, 0xfd, 0x93, 0x69, 0x59, 0xe9, 0x64, 0xd5, 0x0d, 0x26, 0x14, 0xd7, 0xb8,
	0x59, 0xf7, 0x00, 0x89, 0xde, 0xbc, 0x68, 0x69, 0xd9, 0x2e, 0xd9, 0x16, 0x08, 0xa0, 0x46, 0x2f,
	0x84, 0xb6, 0x13, 0x20, 0xd2, 0x0e, 0x81, 0xd1, 0xa4, 0xdf, 0xc0, 0x70, 0xf2, 0x0b, 0x90, 0x98,
	0xa8, 0x89, 0x1f, 0xc3, 0xe3, 0x1e, 0xf7, 0xe8, 0xc9, 0x98, 0xf6, 0xe2, 0xc7, 0x30, 0x1d, 0xcc,
	0xb6, 0xdd, 0xbd, 0x90, 0x97, 0xf7, 0xfd, 0x3d, 0xcf, 0x3c, 0x93, 0x77, 0x60, 0x8b, 0xce, 0x53,
	0x9c, 0xab, 0xec, 0xab, 0xa4, 0x19, 0xa1, 0x04, 0x09, 0x14, 0xcf, 0x26, 0x38, 0x9b, 0xc6, 0x33,
	0xaa, 0xb0, 0xbe, 0x78, 0x48, 0xa3, 0x38, 0x9b, 0xf8, 0x69, 0x90, 0xd1, 0xb9, 0xca, 0x20, 0x35,
	0x24, 0x21, 0x59, 0x57, 0x95, 0xf2, 0xe0, 0x05, 0xe4, 0xed, 0x20, 0xa3, 0x2e, 0xa6, 0xa7, 0x38,
	0x98, 0xe0, 0x0c, 0xed, 0xc1, 0x1b, 0x94, 0xd0, 0x20, 0xd9, 0x07, 0x32, 0x38, 0xe2, 0x9d, 0xea,
	0x07, 0x21, 0x58, 0x8f, 0x82, 0x3c, 0xda, 0xaf, 0xc9, 0xe0, 0xa8, 0xe9, 0xb0, 0xfa, 0x79, 0xfd,
	0xef, 0x97, 0x36, 0x38, 0x08, 0xe1, 0x2d, 0x2d, 0x21, 0xe3, 0xf7, 0x46, 0xef, 0x12, 0x02, 0x6b,
	0x08, 0x9d, 0xc2, 0xe6, 0x2a, 0x43, 0xee, 0x47, 0xcc, 0x9e, 0x19, 0x34, 0x9e, 0xb6, 0x95, 0xab,
	0x81, 0x95, 0xad, 0x14, 0x5a, 0xfd, 0xfc, 0x77, 0x9b, 0x73, 0x1a, 0x4c, 0x5a, 0xb5, 0x8e, 0x7f,
	0x02, 0xd8, 0xf8, 0x7f, 0xd2, 0x49, 0x12, 0x84, 0x48, 0x84, 0xf7, 0xb4, 0x81, 0xd5, 0x3d, 0xeb,
	0xf9, 0x46, 0xcf, 0x3f, 0x19, 0x74, 0xfa, 0xfe, 0x2b, 0xf3, 0xcc, 0xb4, 0xde, 0x98, 0x02, 0x87,
	0x54, 0xb8, 0xc7, 0x66, 0x97, 0xa3, 0x8e, 0xe6, 0xea, 0xa6, 0x27, 0x00, 0xf1, 0x6e, 0x51, 0xca,
	0xad, 0x0d, 0x9b, 0xce, 0x28, 0xc7, 0x33, 0x7a, 0x5d, 0xd0, 0xb5, 0x86, 0x43, 0xc3, 0x13, 0x6a,
	0xd7, 0x04, 0x5d, 0x32, 0x9d, 0xc6, 0x14, 0x3d, 0x81, 0xad, 0x6d, 0x81, 0x69, 0x0c, 0x84, 0x1d,
	0x11, 0x15, 0xa5, 0xbc, 0xbb, 0x41, 0x9b, 0x71, 0x22, 0xde, 0xfe, 0xf4, 0x55, 0xe2, 0x7e, 0x7c,
	0x93, 0xb8, 0xe3, 0xef, 0x00, 0xf2, 0x6e, 0x1c, 0xce, 0xf0, 0x64, 0x98, 0x87, 0xde, 0x3c, 0xc5,
	0xe8, 0x01, 0xbc, 0xef, 0x1a, 0x7d, 0x53, 0xef, 0xf9, 0x43, 0xb7, 0xef, 0x7b, 0x6f, 0x6d, 0x7d,
	0xe3, 0x16, 0x0f, 0x61, 0xd3, 0x76, 0xf4, 0xd7, 0x96, 0xa7, 0xb3, 0x89, 0x00, 0xc4, 0x3b, 0x45,
	0x29, 0x37, 0xec, 0x0c, 0x7f, 0x24, 0x14, 0x33, 0xfd, 0x63, 0xb8, 0x6b, 0x3b, 0x7a, 0x15, 0xb6,
	0x82, 0x6a, 0x62, 0xab, 0x28, 0x65, 0xde, 0xce, 0xf0, 0x98, 0x25, 0x65, 0xd8, 0x23, 0xc8, 0xdb,
	0x8e, 0x65, 0x5b, 0x6e, 0x67, 0x50, 0x51, 0x3b, 0xa2, 0x50, 0x94, 0x72, 0xd3, 0xce, 0x48, 0x4a,
	0xf2, 0x20, 0x59, 0x41, 0xeb, 0x9c, 0xda, 0xcb, 0xf3, 0x85, 0x04, 0x2e, 0x16, 0x12, 0xf8, 0xb3,
	0x90, 0xc0, 0xe7, 0xa5, 0xc4, 0x5d, 0x2c, 0x25, 0xee, 0xd7, 0x52, 0xe2, 0xde, 0x1d, 0x86, 0x31,
	0x8d, 0x3e, 0x8c, 0x94, 0x31, 0x99, 0xaa, 0xeb, 0x15, 0x6e, 0x95, 0xab, 0x6d, 0x8e, 0x6e, 0xb2,
	0xd7, 0xf5, 0xec, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x66, 0x0a, 0x91, 0x0d, 0xac, 0x02, 0x00,
	0x00,
}

func (this *PartSetHeader) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PartSetHeader)
	if !ok {
		that2, ok := that.(PartSetHeader)
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
	if this.Total != that1.Total {
		return false
	}
	if !bytes.Equal(this.Hash, that1.Hash) {
		return false
	}
	return true
}
func (m *PartSetHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PartSetHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PartSetHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x12
	}
	if m.Total != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Total))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BlockID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.PartsHeader.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
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
func (m *PartSetHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Total != 0 {
		n += 1 + sovTypes(uint64(m.Total))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *BlockID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.PartsHeader.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PartSetHeader) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: PartSetHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PartSetHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			m.Total = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Total |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
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
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
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
func (m *BlockID) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: BlockID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
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
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PartsHeader", wireType)
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
			if err := m.PartsHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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