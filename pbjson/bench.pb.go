// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bench.proto

package pbjson

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Small struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Height               float64  `protobuf:"fixed64,2,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Small) Reset()         { *m = Small{} }
func (m *Small) String() string { return proto.CompactTextString(m) }
func (*Small) ProtoMessage()    {}
func (*Small) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c1af4c375b9b7bb, []int{0}
}

func (m *Small) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Small.Unmarshal(m, b)
}
func (m *Small) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Small.Marshal(b, m, deterministic)
}
func (m *Small) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Small.Merge(m, src)
}
func (m *Small) XXX_Size() int {
	return xxx_messageInfo_Small.Size(m)
}
func (m *Small) XXX_DiscardUnknown() {
	xxx_messageInfo_Small.DiscardUnknown(m)
}

var xxx_messageInfo_Small proto.InternalMessageInfo

func (m *Small) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Small) GetHeight() float64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type Medium struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Height               float64  `protobuf:"fixed64,2,opt,name=height,proto3" json:"height,omitempty"`
	Description          []byte   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Num                  float32  `protobuf:"fixed32,4,opt,name=num,proto3" json:"num,omitempty"`
	Tru                  bool     `protobuf:"varint,5,opt,name=tru,proto3" json:"tru,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Medium) Reset()         { *m = Medium{} }
func (m *Medium) String() string { return proto.CompactTextString(m) }
func (*Medium) ProtoMessage()    {}
func (*Medium) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c1af4c375b9b7bb, []int{1}
}

func (m *Medium) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Medium.Unmarshal(m, b)
}
func (m *Medium) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Medium.Marshal(b, m, deterministic)
}
func (m *Medium) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Medium.Merge(m, src)
}
func (m *Medium) XXX_Size() int {
	return xxx_messageInfo_Medium.Size(m)
}
func (m *Medium) XXX_DiscardUnknown() {
	xxx_messageInfo_Medium.DiscardUnknown(m)
}

var xxx_messageInfo_Medium proto.InternalMessageInfo

func (m *Medium) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Medium) GetHeight() float64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Medium) GetDescription() []byte {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *Medium) GetNum() float32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *Medium) GetTru() bool {
	if m != nil {
		return m.Tru
	}
	return false
}

type Large struct {
	Name                 string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Height               float64       `protobuf:"fixed64,2,opt,name=height,proto3" json:"height,omitempty"`
	Description          []byte        `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Num                  float32       `protobuf:"fixed32,4,opt,name=num,proto3" json:"num,omitempty"`
	Tru                  bool          `protobuf:"varint,5,opt,name=tru,proto3" json:"tru,omitempty"`
	Data                 []*Large_Info `protobuf:"bytes,6,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Large) Reset()         { *m = Large{} }
func (m *Large) String() string { return proto.CompactTextString(m) }
func (*Large) ProtoMessage()    {}
func (*Large) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c1af4c375b9b7bb, []int{2}
}

func (m *Large) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Large.Unmarshal(m, b)
}
func (m *Large) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Large.Marshal(b, m, deterministic)
}
func (m *Large) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Large.Merge(m, src)
}
func (m *Large) XXX_Size() int {
	return xxx_messageInfo_Large.Size(m)
}
func (m *Large) XXX_DiscardUnknown() {
	xxx_messageInfo_Large.DiscardUnknown(m)
}

var xxx_messageInfo_Large proto.InternalMessageInfo

func (m *Large) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Large) GetHeight() float64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Large) GetDescription() []byte {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *Large) GetNum() float32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *Large) GetTru() bool {
	if m != nil {
		return m.Tru
	}
	return false
}

func (m *Large) GetData() []*Large_Info {
	if m != nil {
		return m.Data
	}
	return nil
}

type Large_Info struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Lit                  bool     `protobuf:"varint,3,opt,name=lit,proto3" json:"lit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Large_Info) Reset()         { *m = Large_Info{} }
func (m *Large_Info) String() string { return proto.CompactTextString(m) }
func (*Large_Info) ProtoMessage()    {}
func (*Large_Info) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c1af4c375b9b7bb, []int{2, 0}
}

func (m *Large_Info) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Large_Info.Unmarshal(m, b)
}
func (m *Large_Info) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Large_Info.Marshal(b, m, deterministic)
}
func (m *Large_Info) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Large_Info.Merge(m, src)
}
func (m *Large_Info) XXX_Size() int {
	return xxx_messageInfo_Large_Info.Size(m)
}
func (m *Large_Info) XXX_DiscardUnknown() {
	xxx_messageInfo_Large_Info.DiscardUnknown(m)
}

var xxx_messageInfo_Large_Info proto.InternalMessageInfo

func (m *Large_Info) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Large_Info) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Large_Info) GetLit() bool {
	if m != nil {
		return m.Lit
	}
	return false
}

func init() {
	proto.RegisterType((*Small)(nil), "pbjson.Small")
	proto.RegisterType((*Medium)(nil), "pbjson.Medium")
	proto.RegisterType((*Large)(nil), "pbjson.Large")
	proto.RegisterType((*Large_Info)(nil), "pbjson.Large.Info")
}

func init() { proto.RegisterFile("bench.proto", fileDescriptor_8c1af4c375b9b7bb) }

var fileDescriptor_8c1af4c375b9b7bb = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x91, 0xc1, 0x4e, 0xc4, 0x20,
	0x10, 0x86, 0x03, 0x6d, 0xc9, 0x3a, 0x35, 0x66, 0x33, 0x07, 0x43, 0x3c, 0x91, 0x1e, 0x0c, 0xa7,
	0x1e, 0xdc, 0xab, 0x2f, 0x60, 0xa2, 0x17, 0x7c, 0x02, 0xba, 0xe0, 0x16, 0xd3, 0x42, 0xc3, 0xd2,
	0x8b, 0x2f, 0xeb, 0xab, 0x18, 0x70, 0x63, 0x7a, 0xf5, 0xb2, 0xb7, 0xef, 0xff, 0x09, 0x99, 0x2f,
	0x33, 0xd0, 0x0e, 0xd6, 0x1f, 0xc7, 0x7e, 0x89, 0x21, 0x05, 0x64, 0xcb, 0xf0, 0x79, 0x0e, 0xbe,
	0x3b, 0x40, 0xf3, 0x3e, 0xeb, 0x69, 0x42, 0x84, 0xda, 0xeb, 0xd9, 0x72, 0x22, 0x88, 0xbc, 0x51,
	0x85, 0xf1, 0x1e, 0xd8, 0x68, 0xdd, 0x69, 0x4c, 0x9c, 0x0a, 0x22, 0x89, 0xba, 0xa4, 0xee, 0x0b,
	0xd8, 0x9b, 0x35, 0x6e, 0x9d, 0xff, 0xf3, 0x0b, 0x05, 0xb4, 0xc6, 0x9e, 0x8f, 0xd1, 0x2d, 0xc9,
	0x05, 0xcf, 0x2b, 0x41, 0xe4, 0xad, 0xda, 0x56, 0xb8, 0x87, 0xca, 0xaf, 0x33, 0xaf, 0x05, 0x91,
	0x54, 0x65, 0xcc, 0x4d, 0x8a, 0x2b, 0x6f, 0x04, 0x91, 0x3b, 0x95, 0xb1, 0xfb, 0x26, 0xd0, 0xbc,
	0xea, 0x78, 0xb2, 0xd7, 0x9f, 0x8d, 0x8f, 0x50, 0x1b, 0x9d, 0x34, 0x67, 0xa2, 0x92, 0xed, 0x13,
	0xf6, 0xbf, 0x3b, 0xec, 0x8b, 0x4e, 0xff, 0xe2, 0x3f, 0x82, 0x2a, 0xef, 0x0f, 0xcf, 0x50, 0xe7,
	0x84, 0x77, 0x40, 0x9d, 0xb9, 0xf8, 0x51, 0x67, 0xfe, 0x8c, 0xe9, 0xc6, 0x78, 0x0f, 0xd5, 0xe4,
	0x52, 0x31, 0xda, 0xa9, 0x8c, 0x03, 0x2b, 0x17, 0x3a, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0xca,
	0xbf, 0x7d, 0x3d, 0xb0, 0x01, 0x00, 0x00,
}