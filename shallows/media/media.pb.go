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
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
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

func (x *Media) GetId() string {
	if x != nil {
		return x.Id
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

type MediaSearchRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         string                 `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Offset        uint64                 `protobuf:"varint,900,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit         uint64                 `protobuf:"varint,901,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MediaSearchRequest) Reset() {
	*x = MediaSearchRequest{}
	mi := &file_media_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MediaSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaSearchRequest) ProtoMessage() {}

func (x *MediaSearchRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use MediaSearchRequest.ProtoReflect.Descriptor instead.
func (*MediaSearchRequest) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{1}
}

func (x *MediaSearchRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *MediaSearchRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *MediaSearchRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type MediaSearchResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Next          *MediaSearchRequest    `protobuf:"bytes,1,opt,name=next,proto3" json:"next,omitempty"`
	Items         []*Media               `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MediaSearchResponse) Reset() {
	*x = MediaSearchResponse{}
	mi := &file_media_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MediaSearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaSearchResponse) ProtoMessage() {}

func (x *MediaSearchResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use MediaSearchResponse.ProtoReflect.Descriptor instead.
func (*MediaSearchResponse) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{2}
}

func (x *MediaSearchResponse) GetNext() *MediaSearchRequest {
	if x != nil {
		return x.Next
	}
	return nil
}

func (x *MediaSearchResponse) GetItems() []*Media {
	if x != nil {
		return x.Items
	}
	return nil
}

type Download struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Media         *Media                 `protobuf:"bytes,1,opt,name=media,proto3" json:"media,omitempty"`
	Bytes         uint64                 `protobuf:"varint,2,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Downloaded    uint64                 `protobuf:"varint,3,opt,name=downloaded,proto3" json:"downloaded,omitempty"`
	InitiatedAt   string                 `protobuf:"bytes,4,opt,name=initiated_at,proto3" json:"initiated_at,omitempty"`
	PausedAt      string                 `protobuf:"bytes,5,opt,name=paused_at,proto3" json:"paused_at,omitempty"`
	Peers         uint32                 `protobuf:"varint,6,opt,name=peers,proto3" json:"peers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Download) Reset() {
	*x = Download{}
	mi := &file_media_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Download) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Download) ProtoMessage() {}

func (x *Download) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Download.ProtoReflect.Descriptor instead.
func (*Download) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{3}
}

func (x *Download) GetMedia() *Media {
	if x != nil {
		return x.Media
	}
	return nil
}

func (x *Download) GetBytes() uint64 {
	if x != nil {
		return x.Bytes
	}
	return 0
}

func (x *Download) GetDownloaded() uint64 {
	if x != nil {
		return x.Downloaded
	}
	return 0
}

func (x *Download) GetInitiatedAt() string {
	if x != nil {
		return x.InitiatedAt
	}
	return ""
}

func (x *Download) GetPausedAt() string {
	if x != nil {
		return x.PausedAt
	}
	return ""
}

func (x *Download) GetPeers() uint32 {
	if x != nil {
		return x.Peers
	}
	return 0
}

type DownloadSearchRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         string                 `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Offset        uint64                 `protobuf:"varint,900,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit         uint64                 `protobuf:"varint,901,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadSearchRequest) Reset() {
	*x = DownloadSearchRequest{}
	mi := &file_media_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadSearchRequest) ProtoMessage() {}

func (x *DownloadSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadSearchRequest.ProtoReflect.Descriptor instead.
func (*DownloadSearchRequest) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{4}
}

func (x *DownloadSearchRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *DownloadSearchRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *DownloadSearchRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type DownloadSearchResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Next          *DownloadSearchRequest `protobuf:"bytes,1,opt,name=next,proto3" json:"next,omitempty"`
	Items         []*Download            `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadSearchResponse) Reset() {
	*x = DownloadSearchResponse{}
	mi := &file_media_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadSearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadSearchResponse) ProtoMessage() {}

func (x *DownloadSearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadSearchResponse.ProtoReflect.Descriptor instead.
func (*DownloadSearchResponse) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{5}
}

func (x *DownloadSearchResponse) GetNext() *DownloadSearchRequest {
	if x != nil {
		return x.Next
	}
	return nil
}

func (x *DownloadSearchResponse) GetItems() []*Download {
	if x != nil {
		return x.Items
	}
	return nil
}

type DownloadBeginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadBeginRequest) Reset() {
	*x = DownloadBeginRequest{}
	mi := &file_media_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadBeginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadBeginRequest) ProtoMessage() {}

func (x *DownloadBeginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadBeginRequest.ProtoReflect.Descriptor instead.
func (*DownloadBeginRequest) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{6}
}

type DownloadBeginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Download      *Download              `protobuf:"bytes,1,opt,name=download,proto3" json:"download,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadBeginResponse) Reset() {
	*x = DownloadBeginResponse{}
	mi := &file_media_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadBeginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadBeginResponse) ProtoMessage() {}

func (x *DownloadBeginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadBeginResponse.ProtoReflect.Descriptor instead.
func (*DownloadBeginResponse) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{7}
}

func (x *DownloadBeginResponse) GetDownload() *Download {
	if x != nil {
		return x.Download
	}
	return nil
}

type DownloadPauseRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadPauseRequest) Reset() {
	*x = DownloadPauseRequest{}
	mi := &file_media_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadPauseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadPauseRequest) ProtoMessage() {}

func (x *DownloadPauseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadPauseRequest.ProtoReflect.Descriptor instead.
func (*DownloadPauseRequest) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{8}
}

type DownloadPauseResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Download      *Download              `protobuf:"bytes,1,opt,name=download,proto3" json:"download,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadPauseResponse) Reset() {
	*x = DownloadPauseResponse{}
	mi := &file_media_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadPauseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadPauseResponse) ProtoMessage() {}

func (x *DownloadPauseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadPauseResponse.ProtoReflect.Descriptor instead.
func (*DownloadPauseResponse) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{9}
}

func (x *DownloadPauseResponse) GetDownload() *Download {
	if x != nil {
		return x.Download
	}
	return nil
}

var File_media_proto protoreflect.FileDescriptor

var file_media_proto_rawDesc = string([]byte{
	0x0a, 0x0b, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x22, 0x6b, 0x0a, 0x05, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x22, 0x69, 0x0a, 0x12, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x17, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x84, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x15, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x85, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x4a, 0x05, 0x08,
	0x02, 0x10, 0x84, 0x07, 0x4a, 0x06, 0x08, 0x86, 0x07, 0x10, 0xe8, 0x07, 0x22, 0x68, 0x0a, 0x13,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x6e, 0x65,
	0x78, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xbc, 0x01, 0x0a, 0x08, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x22, 0x0a, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x52, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0a, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x12, 0x22, 0x0a,
	0x0c, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x70, 0x65, 0x65, 0x72, 0x73, 0x22, 0x6c, 0x0a, 0x15, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x84,
	0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x15, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x85, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x4a, 0x05, 0x08, 0x02, 0x10, 0x84, 0x07, 0x4a, 0x06, 0x08, 0x86, 0x07,
	0x10, 0xe8, 0x07, 0x22, 0x71, 0x0a, 0x16, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a,
	0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x12,
	0x25, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x44,
	0x0a, 0x15, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x64, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x08, 0x64, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x50, 0x61, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x44, 0x0a, 0x15,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x61, 0x75, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x08, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_media_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_media_proto_goTypes = []any{
	(*Media)(nil),                  // 0: media.Media
	(*MediaSearchRequest)(nil),     // 1: media.MediaSearchRequest
	(*MediaSearchResponse)(nil),    // 2: media.MediaSearchResponse
	(*Download)(nil),               // 3: media.Download
	(*DownloadSearchRequest)(nil),  // 4: media.DownloadSearchRequest
	(*DownloadSearchResponse)(nil), // 5: media.DownloadSearchResponse
	(*DownloadBeginRequest)(nil),   // 6: media.DownloadBeginRequest
	(*DownloadBeginResponse)(nil),  // 7: media.DownloadBeginResponse
	(*DownloadPauseRequest)(nil),   // 8: media.DownloadPauseRequest
	(*DownloadPauseResponse)(nil),  // 9: media.DownloadPauseResponse
}
var file_media_proto_depIdxs = []int32{
	1, // 0: media.MediaSearchResponse.next:type_name -> media.MediaSearchRequest
	0, // 1: media.MediaSearchResponse.items:type_name -> media.Media
	0, // 2: media.Download.media:type_name -> media.Media
	4, // 3: media.DownloadSearchResponse.next:type_name -> media.DownloadSearchRequest
	3, // 4: media.DownloadSearchResponse.items:type_name -> media.Download
	3, // 5: media.DownloadBeginResponse.download:type_name -> media.Download
	3, // 6: media.DownloadPauseResponse.download:type_name -> media.Download
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
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
			NumMessages:   10,
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
