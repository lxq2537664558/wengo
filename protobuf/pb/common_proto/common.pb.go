// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package common_proto

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

//结果码
type RestInt32CodeMsg struct {
	ResCode              int32    `protobuf:"varint,1,opt,name=res_code,json=resCode,proto3" json:"res_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RestInt32CodeMsg) Reset()         { *m = RestInt32CodeMsg{} }
func (m *RestInt32CodeMsg) String() string { return proto.CompactTextString(m) }
func (*RestInt32CodeMsg) ProtoMessage()    {}
func (*RestInt32CodeMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *RestInt32CodeMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RestInt32CodeMsg.Unmarshal(m, b)
}
func (m *RestInt32CodeMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RestInt32CodeMsg.Marshal(b, m, deterministic)
}
func (m *RestInt32CodeMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestInt32CodeMsg.Merge(m, src)
}
func (m *RestInt32CodeMsg) XXX_Size() int {
	return xxx_messageInfo_RestInt32CodeMsg.Size(m)
}
func (m *RestInt32CodeMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_RestInt32CodeMsg.DiscardUnknown(m)
}

var xxx_messageInfo_RestInt32CodeMsg proto.InternalMessageInfo

func (m *RestInt32CodeMsg) GetResCode() int32 {
	if m != nil {
		return m.ResCode
	}
	return 0
}

//手机
type Test struct {
	Type                 int32    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Number               string   `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Test) Reset()         { *m = Test{} }
func (m *Test) String() string { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()    {}
func (*Test) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *Test) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test.Unmarshal(m, b)
}
func (m *Test) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test.Marshal(b, m, deterministic)
}
func (m *Test) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test.Merge(m, src)
}
func (m *Test) XXX_Size() int {
	return xxx_messageInfo_Test.Size(m)
}
func (m *Test) XXX_DiscardUnknown() {
	xxx_messageInfo_Test.DiscardUnknown(m)
}

var xxx_messageInfo_Test proto.InternalMessageInfo

func (m *Test) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Test) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func init() {
	proto.RegisterType((*RestInt32CodeMsg)(nil), "common_proto.RestInt32CodeMsg")
	proto.RegisterType((*Test)(nil), "common_proto.test")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x82, 0xf2, 0xe2, 0xc1, 0x3c, 0x25, 0x5d,
	0x2e, 0x81, 0xa0, 0xd4, 0xe2, 0x12, 0xcf, 0xbc, 0x12, 0x63, 0x23, 0xe7, 0xfc, 0x94, 0x54, 0xdf,
	0xe2, 0x74, 0x21, 0x49, 0x2e, 0x8e, 0xa2, 0xd4, 0xe2, 0xf8, 0xe4, 0xfc, 0x94, 0x54, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xd6, 0x20, 0xf6, 0xa2, 0xd4, 0x62, 0x90, 0xac, 0x92, 0x11, 0x17, 0x4b, 0x49,
	0x6a, 0x71, 0x89, 0x90, 0x10, 0x17, 0x4b, 0x49, 0x65, 0x01, 0x4c, 0x1a, 0xcc, 0x16, 0x12, 0xe3,
	0x62, 0xcb, 0x2b, 0xcd, 0x4d, 0x4a, 0x2d, 0x92, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf2,
	0x9c, 0x04, 0xa2, 0xf8, 0xf4, 0xf4, 0x91, 0x2d, 0x4d, 0x62, 0x03, 0x53, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x91, 0x4e, 0xa4, 0x4c, 0x99, 0x00, 0x00, 0x00,
}
