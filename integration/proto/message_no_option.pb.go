// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: integration/proto/message_no_option.proto

package proto

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type MessageNoOption struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageNoOption) Reset()         { *m = MessageNoOption{} }
func (m *MessageNoOption) String() string { return proto.CompactTextString(m) }
func (*MessageNoOption) ProtoMessage()    {}
func (*MessageNoOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_765bff6061906f37, []int{0}
}
func (m *MessageNoOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageNoOption.Unmarshal(m, b)
}
func (m *MessageNoOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageNoOption.Marshal(b, m, deterministic)
}
func (m *MessageNoOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageNoOption.Merge(m, src)
}
func (m *MessageNoOption) XXX_Size() int {
	return xxx_messageInfo_MessageNoOption.Size(m)
}
func (m *MessageNoOption) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageNoOption.DiscardUnknown(m)
}

var xxx_messageInfo_MessageNoOption proto.InternalMessageInfo

func (m *MessageNoOption) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*MessageNoOption)(nil), "proto.MessageNoOption")
}

func init() {
	proto.RegisterFile("integration/proto/message_no_option.proto", fileDescriptor_765bff6061906f37)
}

var fileDescriptor_765bff6061906f37 = []byte{
	// 98 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xcc, 0xcc, 0x2b, 0x49,
	0x4d, 0x2f, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xcf, 0x4d,
	0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x8d, 0xcf, 0xcb, 0x8f, 0xcf, 0x2f, 0x00, 0x89, 0xeb, 0x81, 0xc5,
	0x85, 0x58, 0xc1, 0x94, 0x92, 0x22, 0x17, 0xbf, 0x2f, 0x44, 0x85, 0x5f, 0xbe, 0x3f, 0x58, 0x5e,
	0x88, 0x8f, 0x8b, 0x29, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x88, 0x29, 0x33, 0x25,
	0x89, 0x0d, 0xac, 0xd2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x63, 0x27, 0x69, 0x5d, 0x00,
	0x00, 0x00,
}