// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: meter/producerbillingline.proto

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

type Producerbillingline struct {
	ProducerDeviceID string `protobuf:"bytes,1,opt,name=producerDeviceID,proto3" json:"producerDeviceID,omitempty"`
	CycleID          uint64 `protobuf:"varint,2,opt,name=cycleID,proto3" json:"cycleID,omitempty"`
	Lineid           uint64 `protobuf:"varint,3,opt,name=lineid,proto3" json:"lineid,omitempty"`
	CustomerdeviceID string `protobuf:"bytes,4,opt,name=customerdeviceID,proto3" json:"customerdeviceID,omitempty"`
	BillContractID   string `protobuf:"bytes,5,opt,name=billContractID,proto3" json:"billContractID,omitempty"`
	LineWh           uint64 `protobuf:"varint,6,opt,name=lineWh,proto3" json:"lineWh,omitempty"`
	LineWhPrice      uint64 `protobuf:"varint,7,opt,name=lineWhPrice,proto3" json:"lineWhPrice,omitempty"`
	Curency          string `protobuf:"bytes,8,opt,name=curency,proto3" json:"curency,omitempty"`
	Decremented      uint64 `protobuf:"varint,9,opt,name=decremented,proto3" json:"decremented,omitempty"`
	Phase            uint64 `protobuf:"varint,10,opt,name=phase,proto3" json:"phase,omitempty"`
	Creator          string `protobuf:"bytes,11,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Producerbillingline) Reset()         { *m = Producerbillingline{} }
func (m *Producerbillingline) String() string { return proto.CompactTextString(m) }
func (*Producerbillingline) ProtoMessage()    {}
func (*Producerbillingline) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a3eb9d4cc8b8029, []int{0}
}
func (m *Producerbillingline) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Producerbillingline) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Producerbillingline.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Producerbillingline) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Producerbillingline.Merge(m, src)
}
func (m *Producerbillingline) XXX_Size() int {
	return m.Size()
}
func (m *Producerbillingline) XXX_DiscardUnknown() {
	xxx_messageInfo_Producerbillingline.DiscardUnknown(m)
}

var xxx_messageInfo_Producerbillingline proto.InternalMessageInfo

func (m *Producerbillingline) GetProducerDeviceID() string {
	if m != nil {
		return m.ProducerDeviceID
	}
	return ""
}

func (m *Producerbillingline) GetCycleID() uint64 {
	if m != nil {
		return m.CycleID
	}
	return 0
}

func (m *Producerbillingline) GetLineid() uint64 {
	if m != nil {
		return m.Lineid
	}
	return 0
}

func (m *Producerbillingline) GetCustomerdeviceID() string {
	if m != nil {
		return m.CustomerdeviceID
	}
	return ""
}

func (m *Producerbillingline) GetBillContractID() string {
	if m != nil {
		return m.BillContractID
	}
	return ""
}

func (m *Producerbillingline) GetLineWh() uint64 {
	if m != nil {
		return m.LineWh
	}
	return 0
}

func (m *Producerbillingline) GetLineWhPrice() uint64 {
	if m != nil {
		return m.LineWhPrice
	}
	return 0
}

func (m *Producerbillingline) GetCurency() string {
	if m != nil {
		return m.Curency
	}
	return ""
}

func (m *Producerbillingline) GetDecremented() uint64 {
	if m != nil {
		return m.Decremented
	}
	return 0
}

func (m *Producerbillingline) GetPhase() uint64 {
	if m != nil {
		return m.Phase
	}
	return 0
}

func (m *Producerbillingline) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Producerbillingline)(nil), "electra.meter.Producerbillingline")
}

func init() { proto.RegisterFile("meter/producerbillingline.proto", fileDescriptor_5a3eb9d4cc8b8029) }

var fileDescriptor_5a3eb9d4cc8b8029 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xbd, 0x4e, 0xfb, 0x30,
	0x14, 0xc5, 0xeb, 0xfe, 0xfb, 0xf1, 0xaf, 0x2b, 0x10, 0x32, 0x1f, 0xf2, 0x64, 0x22, 0x06, 0x54,
	0x31, 0x34, 0x03, 0x6f, 0x00, 0x59, 0xba, 0x55, 0x5d, 0x2a, 0xb1, 0xa5, 0xce, 0x15, 0xb1, 0x94,
	0xc4, 0xd1, 0x8d, 0x83, 0xc8, 0x5b, 0xf0, 0x58, 0x8c, 0x1d, 0x19, 0x51, 0x32, 0xf3, 0x0e, 0x28,
	0x76, 0x22, 0xaa, 0x76, 0xf3, 0x39, 0xf7, 0xf8, 0x77, 0xaf, 0x74, 0xe8, 0x6d, 0x0a, 0x06, 0xd0,
	0xcf, 0x51, 0x47, 0xa5, 0x04, 0xdc, 0xa9, 0x24, 0x51, 0xd9, 0x6b, 0xa2, 0x32, 0x58, 0xe6, 0xa8,
	0x8d, 0x66, 0x67, 0x90, 0x80, 0x34, 0x18, 0x2e, 0x6d, 0xf0, 0xee, 0x67, 0x48, 0x2f, 0xd7, 0xa7,
	0x61, 0xf6, 0x40, 0x2f, 0x7a, 0x46, 0x00, 0x6f, 0x4a, 0xc2, 0x2a, 0xe0, 0xc4, 0x23, 0x8b, 0xd9,
	0xe6, 0xc4, 0x67, 0x9c, 0x4e, 0x65, 0x25, 0x93, 0x36, 0x32, 0xf4, 0xc8, 0x62, 0xb4, 0xe9, 0x25,
	0xbb, 0xa1, 0x93, 0x96, 0xa6, 0x22, 0xfe, 0xcf, 0x0e, 0x3a, 0xd5, 0xd2, 0x65, 0x59, 0x18, 0x9d,
	0x02, 0x46, 0x3d, 0x7d, 0xe4, 0xe8, 0xc7, 0x3e, 0xbb, 0xa7, 0xe7, 0xed, 0x61, 0xcf, 0x3a, 0x33,
	0x18, 0x4a, 0xb3, 0x0a, 0xf8, 0xd8, 0x26, 0x8f, 0xdc, 0x7e, 0xd7, 0x36, 0xe6, 0x93, 0xbf, 0x5d,
	0xdb, 0x98, 0x79, 0x74, 0xee, 0x5e, 0x6b, 0x54, 0x12, 0xf8, 0xd4, 0x0e, 0x0f, 0x2d, 0x7b, 0x7f,
	0x89, 0x90, 0xc9, 0x8a, 0xff, 0xb7, 0xe8, 0x5e, 0xb6, 0x7f, 0x23, 0x90, 0x08, 0x29, 0x64, 0x06,
	0x22, 0x3e, 0x73, 0x7f, 0x0f, 0x2c, 0x76, 0x45, 0xc7, 0x79, 0x1c, 0x16, 0xc0, 0xa9, 0x9d, 0x39,
	0x61, 0x89, 0x08, 0xa1, 0xd1, 0xc8, 0xe7, 0x1d, 0xd1, 0xc9, 0x27, 0xff, 0xb3, 0x16, 0x64, 0x5f,
	0x0b, 0xf2, 0x5d, 0x0b, 0xf2, 0xd1, 0x88, 0xc1, 0xbe, 0x11, 0x83, 0xaf, 0x46, 0x0c, 0x5e, 0xae,
	0xbb, 0x62, 0xfc, 0x77, 0xdf, 0x75, 0x68, 0xaa, 0x1c, 0x8a, 0xdd, 0xc4, 0xd6, 0xf6, 0xf8, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0x86, 0x22, 0x25, 0x70, 0xd9, 0x01, 0x00, 0x00,
}

func (m *Producerbillingline) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Producerbillingline) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Producerbillingline) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintProducerbillingline(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x5a
	}
	if m.Phase != 0 {
		i = encodeVarintProducerbillingline(dAtA, i, uint64(m.Phase))
		i--
		dAtA[i] = 0x50
	}
	if m.Decremented != 0 {
		i = encodeVarintProducerbillingline(dAtA, i, uint64(m.Decremented))
		i--
		dAtA[i] = 0x48
	}
	if len(m.Curency) > 0 {
		i -= len(m.Curency)
		copy(dAtA[i:], m.Curency)
		i = encodeVarintProducerbillingline(dAtA, i, uint64(len(m.Curency)))
		i--
		dAtA[i] = 0x42
	}
	if m.LineWhPrice != 0 {
		i = encodeVarintProducerbillingline(dAtA, i, uint64(m.LineWhPrice))
		i--
		dAtA[i] = 0x38
	}
	if m.LineWh != 0 {
		i = encodeVarintProducerbillingline(dAtA, i, uint64(m.LineWh))
		i--
		dAtA[i] = 0x30
	}
	if len(m.BillContractID) > 0 {
		i -= len(m.BillContractID)
		copy(dAtA[i:], m.BillContractID)
		i = encodeVarintProducerbillingline(dAtA, i, uint64(len(m.BillContractID)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.CustomerdeviceID) > 0 {
		i -= len(m.CustomerdeviceID)
		copy(dAtA[i:], m.CustomerdeviceID)
		i = encodeVarintProducerbillingline(dAtA, i, uint64(len(m.CustomerdeviceID)))
		i--
		dAtA[i] = 0x22
	}
	if m.Lineid != 0 {
		i = encodeVarintProducerbillingline(dAtA, i, uint64(m.Lineid))
		i--
		dAtA[i] = 0x18
	}
	if m.CycleID != 0 {
		i = encodeVarintProducerbillingline(dAtA, i, uint64(m.CycleID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ProducerDeviceID) > 0 {
		i -= len(m.ProducerDeviceID)
		copy(dAtA[i:], m.ProducerDeviceID)
		i = encodeVarintProducerbillingline(dAtA, i, uint64(len(m.ProducerDeviceID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProducerbillingline(dAtA []byte, offset int, v uint64) int {
	offset -= sovProducerbillingline(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Producerbillingline) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ProducerDeviceID)
	if l > 0 {
		n += 1 + l + sovProducerbillingline(uint64(l))
	}
	if m.CycleID != 0 {
		n += 1 + sovProducerbillingline(uint64(m.CycleID))
	}
	if m.Lineid != 0 {
		n += 1 + sovProducerbillingline(uint64(m.Lineid))
	}
	l = len(m.CustomerdeviceID)
	if l > 0 {
		n += 1 + l + sovProducerbillingline(uint64(l))
	}
	l = len(m.BillContractID)
	if l > 0 {
		n += 1 + l + sovProducerbillingline(uint64(l))
	}
	if m.LineWh != 0 {
		n += 1 + sovProducerbillingline(uint64(m.LineWh))
	}
	if m.LineWhPrice != 0 {
		n += 1 + sovProducerbillingline(uint64(m.LineWhPrice))
	}
	l = len(m.Curency)
	if l > 0 {
		n += 1 + l + sovProducerbillingline(uint64(l))
	}
	if m.Decremented != 0 {
		n += 1 + sovProducerbillingline(uint64(m.Decremented))
	}
	if m.Phase != 0 {
		n += 1 + sovProducerbillingline(uint64(m.Phase))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovProducerbillingline(uint64(l))
	}
	return n
}

func sovProducerbillingline(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProducerbillingline(x uint64) (n int) {
	return sovProducerbillingline(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Producerbillingline) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProducerbillingline
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
			return fmt.Errorf("proto: Producerbillingline: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Producerbillingline: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProducerDeviceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProducerbillingline
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
				return ErrInvalidLengthProducerbillingline
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProducerbillingline
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProducerDeviceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CycleID", wireType)
			}
			m.CycleID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProducerbillingline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CycleID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lineid", wireType)
			}
			m.Lineid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProducerbillingline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Lineid |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CustomerdeviceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProducerbillingline
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
				return ErrInvalidLengthProducerbillingline
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProducerbillingline
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CustomerdeviceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BillContractID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProducerbillingline
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
				return ErrInvalidLengthProducerbillingline
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProducerbillingline
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BillContractID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LineWh", wireType)
			}
			m.LineWh = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProducerbillingline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LineWh |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LineWhPrice", wireType)
			}
			m.LineWhPrice = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProducerbillingline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LineWhPrice |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Curency", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProducerbillingline
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
				return ErrInvalidLengthProducerbillingline
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProducerbillingline
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Curency = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decremented", wireType)
			}
			m.Decremented = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProducerbillingline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decremented |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Phase", wireType)
			}
			m.Phase = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProducerbillingline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Phase |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProducerbillingline
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
				return ErrInvalidLengthProducerbillingline
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProducerbillingline
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProducerbillingline(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProducerbillingline
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
func skipProducerbillingline(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProducerbillingline
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
					return 0, ErrIntOverflowProducerbillingline
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
					return 0, ErrIntOverflowProducerbillingline
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
				return 0, ErrInvalidLengthProducerbillingline
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProducerbillingline
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProducerbillingline
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProducerbillingline        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProducerbillingline          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProducerbillingline = fmt.Errorf("proto: unexpected end of group")
)
