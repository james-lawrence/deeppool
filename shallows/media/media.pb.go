// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.2
// source: media.proto

package media

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Media struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Mimetype      string                 `protobuf:"bytes,3,opt,name=mimetype,proto3" json:"mimetype,omitempty"`
	Image         string                 `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Media) Reset() {
	*x = Media{}
	mi := &file_media_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Media) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Media) ProtoMessage() {}

func (x *Media) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Media.ProtoReflect.Descriptor instead.
func (*Media) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{0}
}

func (x *Media) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Media) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Media) GetMimetype() string {
	if x != nil {
		return x.Mimetype
	}
	return ""
}

func (x *Media) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type MediaRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         string                 `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Offset        uint64                 `protobuf:"varint,900,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit         uint64                 `protobuf:"varint,901,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MediaRequest) Reset() {
	*x = MediaRequest{}
	mi := &file_media_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MediaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaRequest) ProtoMessage() {}

func (x *MediaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaRequest.ProtoReflect.Descriptor instead.
func (*MediaRequest) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{1}
}

func (x *MediaRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *MediaRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *MediaRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type MediaResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Next          *MediaRequest          `protobuf:"bytes,1,opt,name=next,proto3" json:"next,omitempty"`
	Items         []*Media               `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MediaResponse) Reset() {
	*x = MediaResponse{}
	mi := &file_media_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MediaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaResponse) ProtoMessage() {}

func (x *MediaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaResponse.ProtoReflect.Descriptor instead.
func (*MediaResponse) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{2}
}

func (x *MediaResponse) GetNext() *MediaRequest {
	if x != nil {
		return x.Next
	}
	return nil
}

func (x *MediaResponse) GetItems() []*Media {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_media_proto protoreflect.FileDescriptor

var file_media_proto_rawDesc = string([]byte{
	0x0a, 0x0b, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x22, 0x71, 0x0a, 0x05, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x63, 0x0a, 0x0c, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x17, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x84, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x15, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x85, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x4a, 0x05, 0x08,
	0x02, 0x10, 0x84, 0x07, 0x4a, 0x06, 0x08, 0x86, 0x07, 0x10, 0xe8, 0x07, 0x22, 0x5c, 0x0a, 0x0d,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a,
	0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_media_proto_rawDescOnce sync.Once
	file_media_proto_rawDescData []byte
)

func file_media_proto_rawDescGZIP() []byte {
	file_media_proto_rawDescOnce.Do(func() {
		file_media_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_media_proto_rawDesc), len(file_media_proto_rawDesc)))
	})
	return file_media_proto_rawDescData
}

var file_media_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_media_proto_goTypes = []any{
	(*Media)(nil),         // 0: media.Media
	(*MediaRequest)(nil),  // 1: media.MediaRequest
	(*MediaResponse)(nil), // 2: media.MediaResponse
}
var file_media_proto_depIdxs = []int32{
	1, // 0: media.MediaResponse.next:type_name -> media.MediaRequest
	0, // 1: media.MediaResponse.items:type_name -> media.Media
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_media_proto_init() }
func file_media_proto_init() {
	if File_media_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_media_proto_rawDesc), len(file_media_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_media_proto_goTypes,
		DependencyIndexes: file_media_proto_depIdxs,
		MessageInfos:      file_media_proto_msgTypes,
	}.Build()
	File_media_proto = out.File
	file_media_proto_goTypes = nil
	file_media_proto_depIdxs = nil
}
