// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/question_proto.proto

package question

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

type Question struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Difficulty           string   `protobuf:"bytes,2,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	Question             string   `protobuf:"bytes,3,opt,name=question,proto3" json:"question,omitempty"`
	Options              []string `protobuf:"bytes,4,rep,name=options,proto3" json:"options,omitempty"`
	Answer               string   `protobuf:"bytes,5,opt,name=answer,proto3" json:"answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Question) Reset()         { *m = Question{} }
func (m *Question) String() string { return proto.CompactTextString(m) }
func (*Question) ProtoMessage()    {}
func (*Question) Descriptor() ([]byte, []int) {
	return fileDescriptor_0789342f327d6024, []int{0}
}

func (m *Question) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Question.Unmarshal(m, b)
}
func (m *Question) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Question.Marshal(b, m, deterministic)
}
func (m *Question) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Question.Merge(m, src)
}
func (m *Question) XXX_Size() int {
	return xxx_messageInfo_Question.Size(m)
}
func (m *Question) XXX_DiscardUnknown() {
	xxx_messageInfo_Question.DiscardUnknown(m)
}

var xxx_messageInfo_Question proto.InternalMessageInfo

func (m *Question) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Question) GetDifficulty() string {
	if m != nil {
		return m.Difficulty
	}
	return ""
}

func (m *Question) GetQuestion() string {
	if m != nil {
		return m.Question
	}
	return ""
}

func (m *Question) GetOptions() []string {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *Question) GetAnswer() string {
	if m != nil {
		return m.Answer
	}
	return ""
}

type QuestionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuestionRequest) Reset()         { *m = QuestionRequest{} }
func (m *QuestionRequest) String() string { return proto.CompactTextString(m) }
func (*QuestionRequest) ProtoMessage()    {}
func (*QuestionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0789342f327d6024, []int{1}
}

func (m *QuestionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuestionRequest.Unmarshal(m, b)
}
func (m *QuestionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuestionRequest.Marshal(b, m, deterministic)
}
func (m *QuestionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuestionRequest.Merge(m, src)
}
func (m *QuestionRequest) XXX_Size() int {
	return xxx_messageInfo_QuestionRequest.Size(m)
}
func (m *QuestionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuestionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuestionRequest proto.InternalMessageInfo

type QuestionResponse struct {
	Data                 []*Question `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *QuestionResponse) Reset()         { *m = QuestionResponse{} }
func (m *QuestionResponse) String() string { return proto.CompactTextString(m) }
func (*QuestionResponse) ProtoMessage()    {}
func (*QuestionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0789342f327d6024, []int{2}
}

func (m *QuestionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuestionResponse.Unmarshal(m, b)
}
func (m *QuestionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuestionResponse.Marshal(b, m, deterministic)
}
func (m *QuestionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuestionResponse.Merge(m, src)
}
func (m *QuestionResponse) XXX_Size() int {
	return xxx_messageInfo_QuestionResponse.Size(m)
}
func (m *QuestionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuestionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuestionResponse proto.InternalMessageInfo

func (m *QuestionResponse) GetData() []*Question {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Question)(nil), "whatw_golang_laravel.Question")
	proto.RegisterType((*QuestionRequest)(nil), "whatw_golang_laravel.QuestionRequest")
	proto.RegisterType((*QuestionResponse)(nil), "whatw_golang_laravel.QuestionResponse")
}

func init() {
	proto.RegisterFile("proto/question_proto.proto", fileDescriptor_0789342f327d6024)
}

var fileDescriptor_0789342f327d6024 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x2c, 0x4d, 0x2d, 0x2e, 0xc9, 0xcc, 0xcf, 0x8b, 0x07, 0x73, 0xf5, 0xc0, 0xa4,
	0x90, 0x48, 0x79, 0x46, 0x62, 0x49, 0x79, 0x7c, 0x7a, 0x7e, 0x4e, 0x62, 0x5e, 0x7a, 0x7c, 0x4e,
	0x62, 0x51, 0x62, 0x59, 0x6a, 0x8e, 0x52, 0x07, 0x23, 0x17, 0x47, 0x20, 0x54, 0xb9, 0x10, 0x1f,
	0x17, 0x53, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x53, 0x66, 0x8a, 0x90, 0x1c,
	0x17, 0x57, 0x4a, 0x66, 0x5a, 0x5a, 0x66, 0x72, 0x69, 0x4e, 0x49, 0xa5, 0x04, 0x93, 0x02, 0xa3,
	0x06, 0x67, 0x10, 0x92, 0x88, 0x90, 0x14, 0x17, 0x07, 0xcc, 0x2a, 0x09, 0x66, 0xb0, 0x2c, 0x9c,
	0x2f, 0x24, 0xc1, 0xc5, 0x9e, 0x5f, 0x00, 0x62, 0x15, 0x4b, 0xb0, 0x28, 0x30, 0x6b, 0x70, 0x06,
	0xc1, 0xb8, 0x42, 0x62, 0x5c, 0x6c, 0x89, 0x79, 0xc5, 0xe5, 0xa9, 0x45, 0x12, 0xac, 0x60, 0x3d,
	0x50, 0x9e, 0x92, 0x20, 0x17, 0x3f, 0xcc, 0x25, 0x41, 0xa9, 0x60, 0x73, 0x94, 0xdc, 0xb8, 0x04,
	0x10, 0x42, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x46, 0x5c, 0x2c, 0x29, 0x89, 0x25, 0x89,
	0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x72, 0x7a, 0xd8, 0xbc, 0xa5, 0x07, 0xd7, 0x05, 0x56,
	0x6b, 0x54, 0x88, 0x30, 0x3a, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0x28, 0x8e, 0x8b, 0xdb,
	0x3d, 0xb5, 0x04, 0xee, 0x75, 0x55, 0x02, 0xe6, 0x40, 0x1c, 0x24, 0xa5, 0x46, 0x48, 0x19, 0xc4,
	0x91, 0x4a, 0x0c, 0x4e, 0xa2, 0x51, 0xc2, 0x05, 0x49, 0xfa, 0x50, 0x05, 0xf0, 0x18, 0x49, 0x62,
	0x03, 0x47, 0x86, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x8f, 0x51, 0x33, 0xaa, 0x01, 0x00,
	0x00,
}
