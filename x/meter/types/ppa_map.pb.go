// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: meter/ppa_map.proto

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

type PpaMap struct {
	ConsumerDeviceID   string `protobuf:"bytes,1,opt,name=consumerDeviceID,proto3" json:"consumerDeviceID,omitempty"`
	AgreementID        string `protobuf:"bytes,2,opt,name=agreementID,proto3" json:"agreementID,omitempty"`
	AgreementActive    bool   `protobuf:"varint,3,opt,name=agreementActive,proto3" json:"agreementActive,omitempty"`
	ContractID         string `protobuf:"bytes,4,opt,name=contractID,proto3" json:"contractID,omitempty"`
	ProducerDeviceID   string `protobuf:"bytes,5,opt,name=producerDeviceID,proto3" json:"producerDeviceID,omitempty"`
	AgreementStartDate uint64 `protobuf:"varint,6,opt,name=agreementStartDate,proto3" json:"agreementStartDate,omitempty"`
	AgreementEndDate   uint64 `protobuf:"varint,7,opt,name=agreementEndDate,proto3" json:"agreementEndDate,omitempty"`
	Creator            string `protobuf:"bytes,8,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *PpaMap) Reset()         { *m = PpaMap{} }
func (m *PpaMap) String() string { return proto.CompactTextString(m) }
func (*PpaMap) ProtoMessage()    {}
func (*PpaMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdb29dd3669abb05, []int{0}
}
func (m *PpaMap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PpaMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PpaMap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PpaMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PpaMap.Merge(m, src)
}
func (m *PpaMap) XXX_Size() int {
	return m.Size()
}
func (m *PpaMap) XXX_DiscardUnknown() {
	xxx_messageInfo_PpaMap.DiscardUnknown(m)
}

var xxx_messageInfo_PpaMap proto.InternalMessageInfo

func (m *PpaMap) GetConsumerDeviceID() string {
	if m != nil {
		return m.ConsumerDeviceID
	}
	return ""
}

func (m *PpaMap) GetAgreementID() string {
	if m != nil {
		return m.AgreementID
	}
	return ""
}

func (m *PpaMap) GetAgreementActive() bool {
	if m != nil {
		return m.AgreementActive
	}
	return false
}

func (m *PpaMap) GetContractID() string {
	if m != nil {
		return m.ContractID
	}
	return ""
}

func (m *PpaMap) GetProducerDeviceID() string {
	if m != nil {
		return m.ProducerDeviceID
	}
	return ""
}

func (m *PpaMap) GetAgreementStartDate() uint64 {
	if m != nil {
		return m.AgreementStartDate
	}
	return 0
}

func (m *PpaMap) GetAgreementEndDate() uint64 {
	if m != nil {
		return m.AgreementEndDate
	}
	return 0
}

func (m *PpaMap) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*PpaMap)(nil), "electra.meter.PpaMap")
}

func init() { proto.RegisterFile("meter/ppa_map.proto", fileDescriptor_fdb29dd3669abb05) }

var fileDescriptor_fdb29dd3669abb05 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xbb, 0xb5, 0xa6, 0x75, 0x45, 0x94, 0x15, 0x61, 0x4f, 0x4b, 0xf0, 0x14, 0x3c, 0x24,
	0x07, 0x9f, 0x40, 0x89, 0x87, 0x1e, 0x04, 0x89, 0x37, 0x2f, 0x32, 0x6e, 0x07, 0x29, 0x98, 0xec,
	0x32, 0x9d, 0x16, 0x7d, 0x0b, 0x9f, 0xc9, 0x93, 0xc7, 0x1e, 0x3d, 0x4a, 0xf2, 0x22, 0x92, 0xad,
	0x86, 0xd0, 0x7a, 0x9c, 0x6f, 0x7e, 0xbe, 0xfd, 0xd9, 0x91, 0xa7, 0x25, 0x32, 0x52, 0xe6, 0x3d,
	0x3c, 0x96, 0xe0, 0x53, 0x4f, 0x8e, 0x9d, 0x3a, 0xc2, 0x17, 0xb4, 0x4c, 0x90, 0x86, 0xe5, 0xf9,
	0xc7, 0x50, 0x46, 0x77, 0x1e, 0x6e, 0xc1, 0xab, 0x0b, 0x79, 0x62, 0x5d, 0xb5, 0x58, 0x96, 0x48,
	0x39, 0xae, 0xe6, 0x16, 0xa7, 0xb9, 0x16, 0xb1, 0x48, 0x0e, 0x8a, 0x1d, 0xae, 0x62, 0x79, 0x08,
	0xcf, 0x84, 0x58, 0x62, 0xc5, 0xd3, 0x5c, 0x0f, 0x43, 0xac, 0x8f, 0x54, 0x22, 0x8f, 0xbb, 0xf1,
	0xca, 0xf2, 0x7c, 0x85, 0x7a, 0x2f, 0x16, 0xc9, 0xa4, 0xd8, 0xc6, 0xca, 0x48, 0x69, 0x5d, 0xc5,
	0x04, 0xb6, 0x55, 0x8d, 0x82, 0xaa, 0x47, 0xda, 0x5e, 0x9e, 0xdc, 0x6c, 0x69, 0x7b, 0xbd, 0xf6,
	0x37, 0xbd, 0xb6, 0xb9, 0x4a, 0xa5, 0xea, 0xf4, 0xf7, 0x0c, 0xc4, 0x39, 0x30, 0xea, 0x28, 0x16,
	0xc9, 0xa8, 0xf8, 0x67, 0xd3, 0xba, 0x3b, 0x7a, 0x53, 0xcd, 0x42, 0x7a, 0x1c, 0xd2, 0x3b, 0x5c,
	0x69, 0x39, 0xb6, 0x84, 0xc0, 0x8e, 0xf4, 0x24, 0x3c, 0xff, 0x37, 0x5e, 0x67, 0x9f, 0xb5, 0x11,
	0xeb, 0xda, 0x88, 0xef, 0xda, 0x88, 0xf7, 0xc6, 0x0c, 0xd6, 0x8d, 0x19, 0x7c, 0x35, 0x66, 0xf0,
	0x70, 0xf6, 0xfb, 0xdb, 0xd9, 0x6b, 0xb6, 0x39, 0x06, 0xbf, 0x79, 0x5c, 0x3c, 0x45, 0xe1, 0x16,
	0x97, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x37, 0x5f, 0xe6, 0xa2, 0x01, 0x00, 0x00,
}

func (m *PpaMap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PpaMap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PpaMap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintPpaMap(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x42
	}
	if m.AgreementEndDate != 0 {
		i = encodeVarintPpaMap(dAtA, i, uint64(m.AgreementEndDate))
		i--
		dAtA[i] = 0x38
	}
	if m.AgreementStartDate != 0 {
		i = encodeVarintPpaMap(dAtA, i, uint64(m.AgreementStartDate))
		i--
		dAtA[i] = 0x30
	}
	if len(m.ProducerDeviceID) > 0 {
		i -= len(m.ProducerDeviceID)
		copy(dAtA[i:], m.ProducerDeviceID)
		i = encodeVarintPpaMap(dAtA, i, uint64(len(m.ProducerDeviceID)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ContractID) > 0 {
		i -= len(m.ContractID)
		copy(dAtA[i:], m.ContractID)
		i = encodeVarintPpaMap(dAtA, i, uint64(len(m.ContractID)))
		i--
		dAtA[i] = 0x22
	}
	if m.AgreementActive {
		i--
		if m.AgreementActive {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.AgreementID) > 0 {
		i -= len(m.AgreementID)
		copy(dAtA[i:], m.AgreementID)
		i = encodeVarintPpaMap(dAtA, i, uint64(len(m.AgreementID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ConsumerDeviceID) > 0 {
		i -= len(m.ConsumerDeviceID)
		copy(dAtA[i:], m.ConsumerDeviceID)
		i = encodeVarintPpaMap(dAtA, i, uint64(len(m.ConsumerDeviceID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPpaMap(dAtA []byte, offset int, v uint64) int {
	offset -= sovPpaMap(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PpaMap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ConsumerDeviceID)
	if l > 0 {
		n += 1 + l + sovPpaMap(uint64(l))
	}
	l = len(m.AgreementID)
	if l > 0 {
		n += 1 + l + sovPpaMap(uint64(l))
	}
	if m.AgreementActive {
		n += 2
	}
	l = len(m.ContractID)
	if l > 0 {
		n += 1 + l + sovPpaMap(uint64(l))
	}
	l = len(m.ProducerDeviceID)
	if l > 0 {
		n += 1 + l + sovPpaMap(uint64(l))
	}
	if m.AgreementStartDate != 0 {
		n += 1 + sovPpaMap(uint64(m.AgreementStartDate))
	}
	if m.AgreementEndDate != 0 {
		n += 1 + sovPpaMap(uint64(m.AgreementEndDate))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovPpaMap(uint64(l))
	}
	return n
}

func sovPpaMap(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPpaMap(x uint64) (n int) {
	return sovPpaMap(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PpaMap) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPpaMap
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
			return fmt.Errorf("proto: PpaMap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PpaMap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsumerDeviceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPpaMap
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
				return ErrInvalidLengthPpaMap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPpaMap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConsumerDeviceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AgreementID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPpaMap
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
				return ErrInvalidLengthPpaMap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPpaMap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AgreementID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AgreementActive", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPpaMap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AgreementActive = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPpaMap
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
				return ErrInvalidLengthPpaMap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPpaMap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProducerDeviceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPpaMap
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
				return ErrInvalidLengthPpaMap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPpaMap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProducerDeviceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AgreementStartDate", wireType)
			}
			m.AgreementStartDate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPpaMap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AgreementStartDate |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AgreementEndDate", wireType)
			}
			m.AgreementEndDate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPpaMap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AgreementEndDate |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPpaMap
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
				return ErrInvalidLengthPpaMap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPpaMap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPpaMap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPpaMap
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
func skipPpaMap(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPpaMap
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
					return 0, ErrIntOverflowPpaMap
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
					return 0, ErrIntOverflowPpaMap
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
				return 0, ErrInvalidLengthPpaMap
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPpaMap
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPpaMap
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPpaMap        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPpaMap          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPpaMap = fmt.Errorf("proto: unexpected end of group")
)
