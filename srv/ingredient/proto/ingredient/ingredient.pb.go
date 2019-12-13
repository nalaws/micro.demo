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

func init() {
	proto.RegisterEnum("Species", Species_name, Species_value)
	proto.RegisterType((*Response)(nil), "Response")
	proto.RegisterType((*NewIngredient)(nil), "NewIngredient")
}

func init() { proto.RegisterFile("ingredient.proto", fileDescriptor_578889e3af7b2ab9) }

var fileDescriptor_578889e3af7b2ab9 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x4f, 0x6b, 0x84, 0x30,
	0x14, 0xc4, 0x4d, 0x6d, 0xfd, 0xf3, 0xa0, 0x22, 0x39, 0x85, 0x9e, 0x6c, 0x4e, 0xd2, 0x43, 0x28,
	0xf6, 0x13, 0xd8, 0x4b, 0x29, 0x94, 0x1e, 0xd2, 0x4f, 0x90, 0x9a, 0x87, 0x04, 0xd6, 0x44, 0x4c,
	0x60, 0xd9, 0x6f, 0xbf, 0x18, 0xd7, 0x65, 0xf7, 0xf6, 0x7b, 0xf3, 0x86, 0x19, 0x06, 0x6a, 0x63,
	0xc7, 0x05, 0xb5, 0x41, 0x1b, 0xc4, 0xbc, 0xb8, 0xe0, 0xf8, 0x0f, 0x14, 0x12, 0xfd, 0xec, 0xac,
	0x47, 0x4a, 0xe1, 0x71, 0x70, 0x1a, 0x19, 0x69, 0x48, 0xfb, 0x24, 0x23, 0x53, 0x06, 0xf9, 0x84,
	0xde, 0xab, 0x11, 0xd9, 0x43, 0x43, 0xda, 0x52, 0xee, 0xe7, 0xea, 0xd6, 0x2a, 0x28, 0x96, 0x46,
	0x39, 0x32, 0xff, 0x82, 0xe7, 0x5f, 0x3c, 0x7e, 0x5f, 0x4b, 0x56, 0x93, 0x55, 0xd3, 0x16, 0x59,
	0xca, 0xc8, 0x94, 0x43, 0xee, 0x67, 0x1c, 0x0c, 0xfa, 0x18, 0x59, 0x75, 0x85, 0xf8, 0xdb, 0x6e,
	0xb9, 0x3f, 0xde, 0x5e, 0x21, 0xbf, 0x68, 0x14, 0x20, 0xeb, 0xad, 0x99, 0xd4, 0xa1, 0x4e, 0x56,
	0xfe, 0x74, 0x41, 0xd9, 0x53, 0x4d, 0xba, 0x77, 0x80, 0x9b, 0x22, 0x0e, 0x69, 0xaf, 0x35, 0xad,
	0xc4, 0x5d, 0xff, 0x4b, 0x29, 0xf6, 0x75, 0x3c, 0xf9, 0xcf, 0xe2, 0xe4, 0x8f, 0x73, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x60, 0x81, 0xce, 0x11, 0x06, 0x01, 0x00, 0x00,
}
