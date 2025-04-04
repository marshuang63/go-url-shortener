// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/url.proto

package urlshortenerpb

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

type ShortenRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LongUrl       string                 `protobuf:"bytes,1,opt,name=long_url,json=longUrl,proto3" json:"long_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShortenRequest) Reset() {
	*x = ShortenRequest{}
	mi := &file_proto_url_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShortenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenRequest) ProtoMessage() {}

func (x *ShortenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenRequest.ProtoReflect.Descriptor instead.
func (*ShortenRequest) Descriptor() ([]byte, []int) {
	return file_proto_url_proto_rawDescGZIP(), []int{0}
}

func (x *ShortenRequest) GetLongUrl() string {
	if x != nil {
		return x.LongUrl
	}
	return ""
}

type ShortenResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ShortUrl      string                 `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShortenResponse) Reset() {
	*x = ShortenResponse{}
	mi := &file_proto_url_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShortenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenResponse) ProtoMessage() {}

func (x *ShortenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenResponse.ProtoReflect.Descriptor instead.
func (*ShortenResponse) Descriptor() ([]byte, []int) {
	return file_proto_url_proto_rawDescGZIP(), []int{1}
}

func (x *ShortenResponse) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type ExpandRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ShortUrl      string                 `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExpandRequest) Reset() {
	*x = ExpandRequest{}
	mi := &file_proto_url_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExpandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandRequest) ProtoMessage() {}

func (x *ExpandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandRequest.ProtoReflect.Descriptor instead.
func (*ExpandRequest) Descriptor() ([]byte, []int) {
	return file_proto_url_proto_rawDescGZIP(), []int{2}
}

func (x *ExpandRequest) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type ExpandResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LongUrl       string                 `protobuf:"bytes,1,opt,name=long_url,json=longUrl,proto3" json:"long_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExpandResponse) Reset() {
	*x = ExpandResponse{}
	mi := &file_proto_url_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExpandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandResponse) ProtoMessage() {}

func (x *ExpandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandResponse.ProtoReflect.Descriptor instead.
func (*ExpandResponse) Descriptor() ([]byte, []int) {
	return file_proto_url_proto_rawDescGZIP(), []int{3}
}

func (x *ExpandResponse) GetLongUrl() string {
	if x != nil {
		return x.LongUrl
	}
	return ""
}

var File_proto_url_proto protoreflect.FileDescriptor

const file_proto_url_proto_rawDesc = "" +
	"\n" +
	"\x0fproto/url.proto\x12\x0eurlshortenerpb\"+\n" +
	"\x0eShortenRequest\x12\x19\n" +
	"\blong_url\x18\x01 \x01(\tR\alongUrl\".\n" +
	"\x0fShortenResponse\x12\x1b\n" +
	"\tshort_url\x18\x01 \x01(\tR\bshortUrl\",\n" +
	"\rExpandRequest\x12\x1b\n" +
	"\tshort_url\x18\x01 \x01(\tR\bshortUrl\"+\n" +
	"\x0eExpandResponse\x12\x19\n" +
	"\blong_url\x18\x01 \x01(\tR\alongUrl2\xa9\x01\n" +
	"\fURLShortener\x12M\n" +
	"\n" +
	"ShortenURL\x12\x1e.urlshortenerpb.ShortenRequest\x1a\x1f.urlshortenerpb.ShortenResponse\x12J\n" +
	"\tExpandURL\x12\x1d.urlshortenerpb.ExpandRequest\x1a\x1e.urlshortenerpb.ExpandResponseB\x16Z\x14proto/urlshortenerpbb\x06proto3"

var (
	file_proto_url_proto_rawDescOnce sync.Once
	file_proto_url_proto_rawDescData []byte
)

func file_proto_url_proto_rawDescGZIP() []byte {
	file_proto_url_proto_rawDescOnce.Do(func() {
		file_proto_url_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_url_proto_rawDesc), len(file_proto_url_proto_rawDesc)))
	})
	return file_proto_url_proto_rawDescData
}

var file_proto_url_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_url_proto_goTypes = []any{
	(*ShortenRequest)(nil),  // 0: urlshortenerpb.ShortenRequest
	(*ShortenResponse)(nil), // 1: urlshortenerpb.ShortenResponse
	(*ExpandRequest)(nil),   // 2: urlshortenerpb.ExpandRequest
	(*ExpandResponse)(nil),  // 3: urlshortenerpb.ExpandResponse
}
var file_proto_url_proto_depIdxs = []int32{
	0, // 0: urlshortenerpb.URLShortener.ShortenURL:input_type -> urlshortenerpb.ShortenRequest
	2, // 1: urlshortenerpb.URLShortener.ExpandURL:input_type -> urlshortenerpb.ExpandRequest
	1, // 2: urlshortenerpb.URLShortener.ShortenURL:output_type -> urlshortenerpb.ShortenResponse
	3, // 3: urlshortenerpb.URLShortener.ExpandURL:output_type -> urlshortenerpb.ExpandResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_url_proto_init() }
func file_proto_url_proto_init() {
	if File_proto_url_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_url_proto_rawDesc), len(file_proto_url_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_url_proto_goTypes,
		DependencyIndexes: file_proto_url_proto_depIdxs,
		MessageInfos:      file_proto_url_proto_msgTypes,
	}.Build()
	File_proto_url_proto = out.File
	file_proto_url_proto_goTypes = nil
	file_proto_url_proto_depIdxs = nil
}
