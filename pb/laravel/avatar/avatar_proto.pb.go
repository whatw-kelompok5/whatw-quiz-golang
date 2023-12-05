// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/avatar_proto.proto

package avatar

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type Avatar struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Image                string   `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Price                *any.Any `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Avatar) Reset()         { *m = Avatar{} }
func (m *Avatar) String() string { return proto.CompactTextString(m) }
func (*Avatar) ProtoMessage()    {}
func (*Avatar) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0c872a618e0c11a, []int{0}
}

func (m *Avatar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Avatar.Unmarshal(m, b)
}
func (m *Avatar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Avatar.Marshal(b, m, deterministic)
}
func (m *Avatar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Avatar.Merge(m, src)
}
func (m *Avatar) XXX_Size() int {
	return xxx_messageInfo_Avatar.Size(m)
}
func (m *Avatar) XXX_DiscardUnknown() {
	xxx_messageInfo_Avatar.DiscardUnknown(m)
}

var xxx_messageInfo_Avatar proto.InternalMessageInfo

func (m *Avatar) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Avatar) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Avatar) GetPrice() *any.Any {
	if m != nil {
		return m.Price
	}
	return nil
}

type AvatarRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AvatarRequest) Reset()         { *m = AvatarRequest{} }
func (m *AvatarRequest) String() string { return proto.CompactTextString(m) }
func (*AvatarRequest) ProtoMessage()    {}
func (*AvatarRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0c872a618e0c11a, []int{1}
}

func (m *AvatarRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvatarRequest.Unmarshal(m, b)
}
func (m *AvatarRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvatarRequest.Marshal(b, m, deterministic)
}
func (m *AvatarRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvatarRequest.Merge(m, src)
}
func (m *AvatarRequest) XXX_Size() int {
	return xxx_messageInfo_AvatarRequest.Size(m)
}
func (m *AvatarRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AvatarRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AvatarRequest proto.InternalMessageInfo

type AvatarResponse struct {
	Data                 []*Avatar `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AvatarResponse) Reset()         { *m = AvatarResponse{} }
func (m *AvatarResponse) String() string { return proto.CompactTextString(m) }
func (*AvatarResponse) ProtoMessage()    {}
func (*AvatarResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0c872a618e0c11a, []int{2}
}

func (m *AvatarResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvatarResponse.Unmarshal(m, b)
}
func (m *AvatarResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvatarResponse.Marshal(b, m, deterministic)
}
func (m *AvatarResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvatarResponse.Merge(m, src)
}
func (m *AvatarResponse) XXX_Size() int {
	return xxx_messageInfo_AvatarResponse.Size(m)
}
func (m *AvatarResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AvatarResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AvatarResponse proto.InternalMessageInfo

func (m *AvatarResponse) GetData() []*Avatar {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Avatar)(nil), "whatw_golang_laravel.Avatar")
	proto.RegisterType((*AvatarRequest)(nil), "whatw_golang_laravel.AvatarRequest")
	proto.RegisterType((*AvatarResponse)(nil), "whatw_golang_laravel.AvatarResponse")
}

func init() {
	proto.RegisterFile("proto/avatar_proto.proto", fileDescriptor_b0c872a618e0c11a)
}

var fileDescriptor_b0c872a618e0c11a = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4f, 0x83, 0x30,
	0x14, 0xc7, 0x2d, 0x93, 0x25, 0xeb, 0xe2, 0x8c, 0x95, 0x43, 0x5d, 0x3c, 0x10, 0xf4, 0x40, 0x3c,
	0x14, 0x83, 0x9f, 0x60, 0xbb, 0x78, 0xc7, 0x8b, 0xd9, 0x85, 0x3c, 0xc6, 0xb3, 0x36, 0x41, 0x8a,
	0xa5, 0x63, 0xd9, 0xb7, 0x37, 0x69, 0xe1, 0x66, 0xb8, 0x34, 0x79, 0xed, 0xef, 0xfd, 0xf2, 0xff,
	0x97, 0xf2, 0xce, 0x68, 0xab, 0x33, 0x18, 0xc0, 0x82, 0x29, 0xdd, 0x20, 0xdc, 0xc9, 0xa2, 0xf3,
	0x37, 0xd8, 0x73, 0x29, 0x75, 0x03, 0xad, 0x2c, 0x1b, 0x30, 0x30, 0x60, 0xb3, 0x7d, 0x90, 0x5a,
	0xcb, 0x06, 0x33, 0xc7, 0x54, 0xa7, 0xaf, 0x0c, 0xda, 0x8b, 0x5f, 0x48, 0x0e, 0x74, 0xb9, 0x73,
	0x1a, 0xb6, 0xa1, 0x81, 0xaa, 0x39, 0x89, 0x49, 0x1a, 0x16, 0x81, 0xaa, 0x59, 0x44, 0x43, 0xf5,
	0x03, 0x12, 0x79, 0x10, 0x93, 0x74, 0x55, 0xf8, 0x81, 0xbd, 0xd0, 0xb0, 0x33, 0xea, 0x88, 0x7c,
	0x11, 0x93, 0x74, 0x9d, 0x47, 0xc2, 0xab, 0xc5, 0xa4, 0x16, 0xbb, 0xf6, 0x52, 0x78, 0x24, 0xb9,
	0xa5, 0x37, 0xde, 0x5d, 0xe0, 0xef, 0x09, 0x7b, 0x9b, 0xec, 0xe9, 0x66, 0xba, 0xe8, 0x3b, 0xdd,
	0xf6, 0xc8, 0x5e, 0xe9, 0x75, 0x0d, 0x16, 0x38, 0x89, 0x17, 0xe9, 0x3a, 0x7f, 0x14, 0xff, 0xc5,
	0x17, 0xe3, 0x8e, 0x23, 0x73, 0x35, 0x49, 0x3f, 0xd0, 0x0c, 0xea, 0x88, 0xec, 0x93, 0xae, 0xde,
	0xd1, 0x8e, 0x25, 0x9e, 0x66, 0x0d, 0x3e, 0xc6, 0xf6, 0x79, 0x1e, 0xf2, 0xd1, 0x92, 0xab, 0xfd,
	0xfd, 0xe1, 0xae, 0xab, 0xb2, 0xf1, 0x79, 0xfc, 0xed, 0x6a, 0xe9, 0x9a, 0xbe, 0xfd, 0x05, 0x00,
	0x00, 0xff, 0xff, 0x4e, 0x12, 0x6c, 0x16, 0x84, 0x01, 0x00, 0x00,
}
