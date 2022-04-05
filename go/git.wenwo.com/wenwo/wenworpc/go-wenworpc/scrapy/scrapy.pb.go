// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: wenworpc/scrapy.proto

package scrapy

import (
	common "git.wenwo.com/wenwo/wenworpc/go-wenworpc/common"
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

type ArticleScrapy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Author   string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Platform string `protobuf:"bytes,3,opt,name=platform,proto3" json:"platform,omitempty"`
	Url      string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	Abstruct string `protobuf:"bytes,5,opt,name=abstruct,proto3" json:"abstruct,omitempty"`
	Content  string `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ArticleScrapy) Reset() {
	*x = ArticleScrapy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wenworpc_scrapy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleScrapy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleScrapy) ProtoMessage() {}

func (x *ArticleScrapy) ProtoReflect() protoreflect.Message {
	mi := &file_wenworpc_scrapy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleScrapy.ProtoReflect.Descriptor instead.
func (*ArticleScrapy) Descriptor() ([]byte, []int) {
	return file_wenworpc_scrapy_proto_rawDescGZIP(), []int{0}
}

func (x *ArticleScrapy) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ArticleScrapy) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *ArticleScrapy) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *ArticleScrapy) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ArticleScrapy) GetAbstruct() string {
	if x != nil {
		return x.Abstruct
	}
	return ""
}

func (x *ArticleScrapy) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type ArticleScrapyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerToken string `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
	Page          uint32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Pagesize      uint32 `protobuf:"varint,3,opt,name=pagesize,proto3" json:"pagesize,omitempty"`
}

func (x *ArticleScrapyRequest) Reset() {
	*x = ArticleScrapyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wenworpc_scrapy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleScrapyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleScrapyRequest) ProtoMessage() {}

func (x *ArticleScrapyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wenworpc_scrapy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleScrapyRequest.ProtoReflect.Descriptor instead.
func (*ArticleScrapyRequest) Descriptor() ([]byte, []int) {
	return file_wenworpc_scrapy_proto_rawDescGZIP(), []int{1}
}

func (x *ArticleScrapyRequest) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

func (x *ArticleScrapyRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ArticleScrapyRequest) GetPagesize() uint32 {
	if x != nil {
		return x.Pagesize
	}
	return 0
}

type ArticleScrapyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *common.Error    `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Aes   []*ArticleScrapy `protobuf:"bytes,2,rep,name=aes,proto3" json:"aes,omitempty"`
}

func (x *ArticleScrapyResponse) Reset() {
	*x = ArticleScrapyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wenworpc_scrapy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleScrapyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleScrapyResponse) ProtoMessage() {}

func (x *ArticleScrapyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wenworpc_scrapy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleScrapyResponse.ProtoReflect.Descriptor instead.
func (*ArticleScrapyResponse) Descriptor() ([]byte, []int) {
	return file_wenworpc_scrapy_proto_rawDescGZIP(), []int{2}
}

func (x *ArticleScrapyResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *ArticleScrapyResponse) GetAes() []*ArticleScrapy {
	if x != nil {
		return x.Aes
	}
	return nil
}

var File_wenworpc_scrapy_proto protoreflect.FileDescriptor

var file_wenworpc_scrapy_proto_rawDesc = []byte{
	0x0a, 0x15, 0x77, 0x65, 0x6e, 0x77, 0x6f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x63, 0x72, 0x61, 0x70,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x77, 0x65, 0x6e, 0x77, 0x6f, 0x72, 0x70,
	0x63, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x70, 0x79, 0x1a, 0x15, 0x77, 0x65, 0x6e, 0x77, 0x6f, 0x72,
	0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa1, 0x01, 0x0a, 0x0d, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x53, 0x63, 0x72, 0x61, 0x70,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x61, 0x62, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x61, 0x62, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0x6d, 0x0a, 0x14, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x53, 0x63,
	0x72, 0x61, 0x70, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x73, 0x69,
	0x7a, 0x65, 0x22, 0x70, 0x0a, 0x15, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x53, 0x63, 0x72,
	0x61, 0x70, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x77, 0x65, 0x6e,
	0x77, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x30, 0x0a, 0x03, 0x61, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x77, 0x65, 0x6e, 0x77, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x70,
	0x79, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x53, 0x63, 0x72, 0x61, 0x70, 0x79, 0x52,
	0x03, 0x61, 0x65, 0x73, 0x32, 0x68, 0x0a, 0x06, 0x53, 0x63, 0x72, 0x61, 0x70, 0x79, 0x12, 0x5e,
	0x0a, 0x0b, 0x67, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x25, 0x2e,
	0x77, 0x65, 0x6e, 0x77, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x70, 0x79, 0x2e,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x53, 0x63, 0x72, 0x61, 0x70, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x77, 0x65, 0x6e, 0x77, 0x6f, 0x72, 0x70, 0x63, 0x2e,
	0x73, 0x63, 0x72, 0x61, 0x70, 0x79, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x53, 0x63,
	0x72, 0x61, 0x70, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x46,
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x65, 0x6e, 0x77, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x73,
	0x63, 0x72, 0x61, 0x70, 0x79, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x2e, 0x77, 0x65, 0x6e, 0x77, 0x6f,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x6e, 0x77, 0x6f, 0x2f, 0x77, 0x65, 0x6e, 0x77, 0x6f,
	0x72, 0x70, 0x63, 0x2f, 0x67, 0x6f, 0x2d, 0x77, 0x65, 0x6e, 0x77, 0x6f, 0x72, 0x70, 0x63, 0x2f,
	0x73, 0x63, 0x72, 0x61, 0x70, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wenworpc_scrapy_proto_rawDescOnce sync.Once
	file_wenworpc_scrapy_proto_rawDescData = file_wenworpc_scrapy_proto_rawDesc
)

func file_wenworpc_scrapy_proto_rawDescGZIP() []byte {
	file_wenworpc_scrapy_proto_rawDescOnce.Do(func() {
		file_wenworpc_scrapy_proto_rawDescData = protoimpl.X.CompressGZIP(file_wenworpc_scrapy_proto_rawDescData)
	})
	return file_wenworpc_scrapy_proto_rawDescData
}

var file_wenworpc_scrapy_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_wenworpc_scrapy_proto_goTypes = []interface{}{
	(*ArticleScrapy)(nil),         // 0: wenworpc.scrapy.ArticleScrapy
	(*ArticleScrapyRequest)(nil),  // 1: wenworpc.scrapy.ArticleScrapyRequest
	(*ArticleScrapyResponse)(nil), // 2: wenworpc.scrapy.ArticleScrapyResponse
	(*common.Error)(nil),          // 3: wenworpc.Error
}
var file_wenworpc_scrapy_proto_depIdxs = []int32{
	3, // 0: wenworpc.scrapy.ArticleScrapyResponse.error:type_name -> wenworpc.Error
	0, // 1: wenworpc.scrapy.ArticleScrapyResponse.aes:type_name -> wenworpc.scrapy.ArticleScrapy
	1, // 2: wenworpc.scrapy.Scrapy.getArticles:input_type -> wenworpc.scrapy.ArticleScrapyRequest
	2, // 3: wenworpc.scrapy.Scrapy.getArticles:output_type -> wenworpc.scrapy.ArticleScrapyResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_wenworpc_scrapy_proto_init() }
func file_wenworpc_scrapy_proto_init() {
	if File_wenworpc_scrapy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wenworpc_scrapy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleScrapy); i {
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
		file_wenworpc_scrapy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleScrapyRequest); i {
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
		file_wenworpc_scrapy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleScrapyResponse); i {
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
			RawDescriptor: file_wenworpc_scrapy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wenworpc_scrapy_proto_goTypes,
		DependencyIndexes: file_wenworpc_scrapy_proto_depIdxs,
		MessageInfos:      file_wenworpc_scrapy_proto_msgTypes,
	}.Build()
	File_wenworpc_scrapy_proto = out.File
	file_wenworpc_scrapy_proto_rawDesc = nil
	file_wenworpc_scrapy_proto_goTypes = nil
	file_wenworpc_scrapy_proto_depIdxs = nil
}