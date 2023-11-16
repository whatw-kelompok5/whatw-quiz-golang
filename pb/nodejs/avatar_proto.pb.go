// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.0
// source: avatar_proto.proto

package nodejs

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Avatars struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Image string `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Price int32  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Avatars) Reset() {
	*x = Avatars{}
	if protoimpl.UnsafeEnabled {
		mi := &file_avatar_proto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Avatars) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Avatars) ProtoMessage() {}

func (x *Avatars) ProtoReflect() protoreflect.Message {
	mi := &file_avatar_proto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Avatars.ProtoReflect.Descriptor instead.
func (*Avatars) Descriptor() ([]byte, []int) {
	return file_avatar_proto_proto_rawDescGZIP(), []int{0}
}

func (x *Avatars) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Avatars) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Avatars) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type AvatarsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AvatarsRequest) Reset() {
	*x = AvatarsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_avatar_proto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarsRequest) ProtoMessage() {}

func (x *AvatarsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_avatar_proto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarsRequest.ProtoReflect.Descriptor instead.
func (*AvatarsRequest) Descriptor() ([]byte, []int) {
	return file_avatar_proto_proto_rawDescGZIP(), []int{1}
}

type AvatarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Avatars `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *AvatarResponse) Reset() {
	*x = AvatarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_avatar_proto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarResponse) ProtoMessage() {}

func (x *AvatarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_avatar_proto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarResponse.ProtoReflect.Descriptor instead.
func (*AvatarResponse) Descriptor() ([]byte, []int) {
	return file_avatar_proto_proto_rawDescGZIP(), []int{2}
}

func (x *AvatarResponse) GetData() []*Avatars {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_avatar_proto_proto protoreflect.FileDescriptor

var file_avatar_proto_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x77, 0x68, 0x61, 0x74, 0x77, 0x5f, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x6a, 0x73, 0x22, 0x45, 0x0a, 0x07, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x22, 0x10, 0x0a, 0x0e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x42, 0x0a, 0x0e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x68, 0x61, 0x74, 0x77, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x6a, 0x73, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x73,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x69, 0x0a, 0x0d, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x73, 0x12, 0x23, 0x2e, 0x77, 0x68, 0x61, 0x74, 0x77, 0x5f, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x6a, 0x73, 0x2e, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x77, 0x68, 0x61,
	0x74, 0x77, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x6a, 0x73,
	0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0b, 0x5a, 0x09, 0x70, 0x62, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x6a, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_avatar_proto_proto_rawDescOnce sync.Once
	file_avatar_proto_proto_rawDescData = file_avatar_proto_proto_rawDesc
)

func file_avatar_proto_proto_rawDescGZIP() []byte {
	file_avatar_proto_proto_rawDescOnce.Do(func() {
		file_avatar_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_avatar_proto_proto_rawDescData)
	})
	return file_avatar_proto_proto_rawDescData
}

var file_avatar_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_avatar_proto_proto_goTypes = []interface{}{
	(*Avatars)(nil),        // 0: whatw_golang_nodejs.Avatars
	(*AvatarsRequest)(nil), // 1: whatw_golang_nodejs.AvatarsRequest
	(*AvatarResponse)(nil), // 2: whatw_golang_nodejs.AvatarResponse
}
var file_avatar_proto_proto_depIdxs = []int32{
	0, // 0: whatw_golang_nodejs.AvatarResponse.data:type_name -> whatw_golang_nodejs.Avatars
	1, // 1: whatw_golang_nodejs.AvatarService.GetAvatars:input_type -> whatw_golang_nodejs.AvatarsRequest
	2, // 2: whatw_golang_nodejs.AvatarService.GetAvatars:output_type -> whatw_golang_nodejs.AvatarResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_avatar_proto_proto_init() }
func file_avatar_proto_proto_init() {
	if File_avatar_proto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_avatar_proto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Avatars); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_avatar_proto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_avatar_proto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_avatar_proto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_avatar_proto_proto_goTypes,
		DependencyIndexes: file_avatar_proto_proto_depIdxs,
		MessageInfos:      file_avatar_proto_proto_msgTypes,
	}.Build()
	File_avatar_proto_proto = out.File
	file_avatar_proto_proto_rawDesc = nil
	file_avatar_proto_proto_goTypes = nil
	file_avatar_proto_proto_depIdxs = nil
}
