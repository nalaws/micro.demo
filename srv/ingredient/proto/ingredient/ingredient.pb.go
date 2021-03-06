// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ingredient.proto

package ingredient

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

type Species int32

const (
	Species_Animal Species = 0
	Species_Botany Species = 1
)

var Species_name = map[int32]string{
	0: "Animal",
	1: "Botany",
}

var Species_value = map[string]int32{
	"Animal": 0,
	"Botany": 1,
}

func (x Species) String() string {
	return proto.EnumName(Species_name, int32(x))
}

func (Species) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_578889e3af7b2ab9, []int{0}
}

type Response struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_578889e3af7b2ab9, []int{0}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Response) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type NewIngredient struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Species              Species  `protobuf:"varint,2,opt,name=species,proto3,enum=Species" json:"species,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewIngredient) Reset()         { *m = NewIngredient{} }
func (m *NewIngredient) String() string { return proto.CompactTextString(m) }
func (*NewIngredient) ProtoMessage()    {}
func (*NewIngredient) Descriptor() ([]byte, []int) {
	return fileDescriptor_578889e3af7b2ab9, []int{1}
}

func (m *NewIngredient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewIngredient.Unmarshal(m, b)
}
func (m *NewIngredient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewIngredient.Marshal(b, m, deterministic)
}
func (m *NewIngredient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewIngredient.Merge(m, src)
}
func (m *NewIngredient) XXX_Size() int {
	return xxx_messageInfo_NewIngredient.Size(m)
}
func (m *NewIngredient) XXX_DiscardUnknown() {
	xxx_messageInfo_NewIngredient.DiscardUnknown(m)
}

var xxx_messageInfo_NewIngredient proto.InternalMessageInfo

func (m *NewIngredient) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NewIngredient) GetSpecies() Species {
	if m != nil {
		return m.Species
	}
	return Species_Animal
}

type IngredientName struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IngredientName) Reset()         { *m = IngredientName{} }
func (m *IngredientName) String() string { return proto.CompactTextString(m) }
func (*IngredientName) ProtoMessage()    {}
func (*IngredientName) Descriptor() ([]byte, []int) {
	return fileDescriptor_578889e3af7b2ab9, []int{2}
}

func (m *IngredientName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IngredientName.Unmarshal(m, b)
}
func (m *IngredientName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IngredientName.Marshal(b, m, deterministic)
}
func (m *IngredientName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IngredientName.Merge(m, src)
}
func (m *IngredientName) XXX_Size() int {
	return xxx_messageInfo_IngredientName.Size(m)
}
func (m *IngredientName) XXX_DiscardUnknown() {
	xxx_messageInfo_IngredientName.DiscardUnknown(m)
}

var xxx_messageInfo_IngredientName proto.InternalMessageInfo

func (m *IngredientName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type IngredientInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IngredientInfo) Reset()         { *m = IngredientInfo{} }
func (m *IngredientInfo) String() string { return proto.CompactTextString(m) }
func (*IngredientInfo) ProtoMessage()    {}
func (*IngredientInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_578889e3af7b2ab9, []int{3}
}

func (m *IngredientInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IngredientInfo.Unmarshal(m, b)
}
func (m *IngredientInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IngredientInfo.Marshal(b, m, deterministic)
}
func (m *IngredientInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IngredientInfo.Merge(m, src)
}
func (m *IngredientInfo) XXX_Size() int {
	return xxx_messageInfo_IngredientInfo.Size(m)
}
func (m *IngredientInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_IngredientInfo.DiscardUnknown(m)
}

var xxx_messageInfo_IngredientInfo proto.InternalMessageInfo

func (m *IngredientInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterEnum("Species", Species_name, Species_value)
	proto.RegisterType((*Response)(nil), "Response")
	proto.RegisterType((*NewIngredient)(nil), "NewIngredient")
	proto.RegisterType((*IngredientName)(nil), "IngredientName")
	proto.RegisterType((*IngredientInfo)(nil), "IngredientInfo")
}

func init() { proto.RegisterFile("ingredient.proto", fileDescriptor_578889e3af7b2ab9) }

var fileDescriptor_578889e3af7b2ab9 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x63, 0x0a, 0x4d, 0x73, 0x12, 0x21, 0x32, 0x8b, 0xd5, 0xa9, 0x58, 0x0c, 0x15, 0x83,
	0x87, 0x32, 0x31, 0xb6, 0x4b, 0x55, 0x09, 0x75, 0x30, 0xbf, 0xc0, 0xd4, 0x47, 0x65, 0x41, 0xec,
	0xa8, 0xb6, 0x84, 0xfa, 0xef, 0x91, 0x5d, 0xd2, 0x24, 0x52, 0xb6, 0x77, 0x77, 0xcf, 0x9f, 0xdf,
	0x1d, 0x54, 0xc6, 0x1e, 0x4f, 0xa8, 0x0d, 0xda, 0x20, 0x9a, 0x93, 0x0b, 0x8e, 0xbf, 0xc3, 0x4c,
	0xa2, 0x6f, 0x9c, 0xf5, 0x48, 0x29, 0xdc, 0x1e, 0x9c, 0x46, 0x46, 0x16, 0x64, 0x79, 0x27, 0x93,
	0xa6, 0x0c, 0xf2, 0x1a, 0xbd, 0x57, 0x47, 0x64, 0x37, 0x0b, 0xb2, 0x2c, 0x64, 0x5b, 0x46, 0xb7,
	0x56, 0x41, 0xb1, 0x49, 0x6a, 0x27, 0xcd, 0xb7, 0x70, 0xbf, 0xc7, 0xdf, 0xdd, 0xf5, 0x93, 0x68,
	0xb2, 0xaa, 0xbe, 0x20, 0x0b, 0x99, 0x34, 0xe5, 0x90, 0xfb, 0x06, 0x0f, 0x06, 0x7d, 0x42, 0x96,
	0xab, 0x99, 0xf8, 0xb8, 0xd4, 0xb2, 0x1d, 0xf0, 0x67, 0x28, 0x3b, 0xca, 0x3e, 0xbe, 0x1a, 0x21,
	0x0d, 0x5d, 0x3b, 0xfb, 0xe5, 0xc6, 0x5c, 0x2f, 0x4f, 0x90, 0xff, 0xf3, 0x29, 0xc0, 0x74, 0x6d,
	0x4d, 0xad, 0x7e, 0xaa, 0x2c, 0xea, 0x8d, 0x0b, 0xca, 0x9e, 0x2b, 0xb2, 0xfa, 0x06, 0xe8, 0x85,
	0xe6, 0x30, 0x59, 0x6b, 0x4d, 0x4b, 0x31, 0xd8, 0x65, 0x5e, 0x88, 0xf6, 0x52, 0x3c, 0xa3, 0x6f,
	0xf0, 0xb8, 0xc5, 0xd0, 0x4d, 0x37, 0xe7, 0x94, 0xf2, 0x41, 0x0c, 0x63, 0xcf, 0xfb, 0x8d, 0x98,
	0x90, 0x67, 0x9f, 0xd3, 0x74, 0xf9, 0xd7, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x93, 0x61,
	0x21, 0x8d, 0x01, 0x00, 0x00,
}
