// Code generated by protoc-gen-go. DO NOT EDIT.
// source: version_metadata.proto

package statecouchdb

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

type VersionAndMetadata struct {
	Version              []byte   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Metadata             []byte   `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionAndMetadata) Reset()         { *m = VersionAndMetadata{} }
func (m *VersionAndMetadata) String() string { return proto.CompactTextString(m) }
func (*VersionAndMetadata) ProtoMessage()    {}
func (*VersionAndMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_96eb34fdc42aad00, []int{0}
}

func (m *VersionAndMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionAndMetadata.Unmarshal(m, b)
}
func (m *VersionAndMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionAndMetadata.Marshal(b, m, deterministic)
}
func (m *VersionAndMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionAndMetadata.Merge(m, src)
}
func (m *VersionAndMetadata) XXX_Size() int {
	return xxx_messageInfo_VersionAndMetadata.Size(m)
}
func (m *VersionAndMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionAndMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_VersionAndMetadata proto.InternalMessageInfo

func (m *VersionAndMetadata) GetVersion() []byte {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *VersionAndMetadata) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*VersionAndMetadata)(nil), "statecouchdb.VersionAndMetadata")
}

func init() { proto.RegisterFile("version_metadata.proto", fileDescriptor_96eb34fdc42aad00) }

var fileDescriptor_96eb34fdc42aad00 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x4b, 0x2d, 0x2a,
	0xce, 0xcc, 0xcf, 0x8b, 0xcf, 0x4d, 0x2d, 0x49, 0x4c, 0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x29, 0x2e, 0x49, 0x2c, 0x49, 0x4d, 0xce, 0x2f, 0x4d, 0xce, 0x48, 0x49,
	0x52, 0xf2, 0xe2, 0x12, 0x0a, 0x83, 0xa8, 0x73, 0xcc, 0x4b, 0xf1, 0x85, 0xaa, 0x14, 0x92, 0xe0,
	0x62, 0x87, 0xea, 0x96, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x82, 0x71, 0x85, 0xa4, 0xb8, 0x38,
	0x60, 0xe6, 0x49, 0x30, 0x81, 0xa5, 0xe0, 0x7c, 0xa7, 0x80, 0x28, 0xbf, 0xf4, 0xcc, 0x92, 0x8c,
	0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0x8c, 0xca, 0x82, 0xd4, 0xa2, 0x9c, 0xd4, 0x94, 0xf4,
	0xd4, 0x22, 0xfd, 0xb4, 0xc4, 0xa4, 0xa2, 0xcc, 0x64, 0xfd, 0xe4, 0xfc, 0xa2, 0x54, 0x7d, 0xa8,
	0x50, 0x76, 0x19, 0x94, 0x51, 0x52, 0x91, 0x9b, 0x9e, 0x5b, 0xa2, 0x0f, 0x76, 0x55, 0x4a, 0x92,
	0x3e, 0xb2, 0xeb, 0x92, 0xd8, 0xc0, 0x4e, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x42,
	0x95, 0xd8, 0xcc, 0x00, 0x00, 0x00,
}