// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: integration/proto/message_with_option.proto

package proto

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/travisjeffery/proto-go-sql"
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

type MessageWithOption struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageWithOption) Reset()         { *m = MessageWithOption{} }
func (m *MessageWithOption) String() string { return proto.CompactTextString(m) }
func (*MessageWithOption) ProtoMessage()    {}
func (*MessageWithOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0aebaba4f6f5191, []int{0}
}
func (m *MessageWithOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageWithOption.Unmarshal(m, b)
}
func (m *MessageWithOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageWithOption.Marshal(b, m, deterministic)
}
func (m *MessageWithOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageWithOption.Merge(m, src)
}
func (m *MessageWithOption) XXX_Size() int {
	return xxx_messageInfo_MessageWithOption.Size(m)
}
func (m *MessageWithOption) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageWithOption.DiscardUnknown(m)
}

var xxx_messageInfo_MessageWithOption proto.InternalMessageInfo

func (m *MessageWithOption) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*MessageWithOption)(nil), "proto.MessageWithOption")
}

func init() {
	proto.RegisterFile("integration/proto/message_with_option.proto", fileDescriptor_e0aebaba4f6f5191)
}

var fileDescriptor_e0aebaba4f6f5191 = []byte{
	// 119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xce, 0xcc, 0x2b, 0x49,
	0x4d, 0x2f, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xcf, 0x4d,
	0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x8d, 0x2f, 0xcf, 0x2c, 0xc9, 0x88, 0xcf, 0x2f, 0x00, 0xc9, 0xe8,
	0x81, 0x65, 0x84, 0x58, 0xc1, 0x94, 0x14, 0x67, 0x71, 0x61, 0x0e, 0x44, 0x44, 0x49, 0x97, 0x4b,
	0xd0, 0x17, 0xa2, 0x3c, 0x3c, 0xb3, 0x24, 0xc3, 0x1f, 0xac, 0x58, 0x88, 0x8f, 0x8b, 0x29, 0x33,
	0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x88, 0x29, 0x33, 0xc5, 0x8a, 0xa3, 0xeb, 0xa7, 0x36,
	0x4b, 0x56, 0x71, 0x7e, 0x5e, 0x12, 0x1b, 0x58, 0x97, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xda,
	0x63, 0xa2, 0x9b, 0x76, 0x00, 0x00, 0x00,
}
