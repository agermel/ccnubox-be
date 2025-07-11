// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.19.4
// source: feed/v1/feed_error.proto

package feedv1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type ErrorReason int32

const (
	ErrorReason_TOKEN_ALREADY_EXIST          ErrorReason = 0
	ErrorReason_USER_NOT_FOUND               ErrorReason = 1
	ErrorReason_GET_FEED_EVENT_ERROR         ErrorReason = 2
	ErrorReason_CLEAR_FEED_EVENT_ERROR       ErrorReason = 3
	ErrorReason_PUBLIC_FEED_EVENT_ERROR      ErrorReason = 4
	ErrorReason_FIND_CONFIG_OR_TOKEN_ERROR   ErrorReason = 5
	ErrorReason_CHANGE_CONFIG_OR_TOKEN_ERROR ErrorReason = 6
	ErrorReason_REMOVE_CONFIG_OR_TOKEN_ERROR ErrorReason = 7
	ErrorReason_GET_MUXI_FEED_ERROR          ErrorReason = 8
	ErrorReason_INSERT_MUXI_FEED_ERROR       ErrorReason = 9
	ErrorReason_REMOVE_MUXI_FEED_ERROR       ErrorReason = 10
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0:  "TOKEN_ALREADY_EXIST",
		1:  "USER_NOT_FOUND",
		2:  "GET_FEED_EVENT_ERROR",
		3:  "CLEAR_FEED_EVENT_ERROR",
		4:  "PUBLIC_FEED_EVENT_ERROR",
		5:  "FIND_CONFIG_OR_TOKEN_ERROR",
		6:  "CHANGE_CONFIG_OR_TOKEN_ERROR",
		7:  "REMOVE_CONFIG_OR_TOKEN_ERROR",
		8:  "GET_MUXI_FEED_ERROR",
		9:  "INSERT_MUXI_FEED_ERROR",
		10: "REMOVE_MUXI_FEED_ERROR",
	}
	ErrorReason_value = map[string]int32{
		"TOKEN_ALREADY_EXIST":          0,
		"USER_NOT_FOUND":               1,
		"GET_FEED_EVENT_ERROR":         2,
		"CLEAR_FEED_EVENT_ERROR":       3,
		"PUBLIC_FEED_EVENT_ERROR":      4,
		"FIND_CONFIG_OR_TOKEN_ERROR":   5,
		"CHANGE_CONFIG_OR_TOKEN_ERROR": 6,
		"REMOVE_CONFIG_OR_TOKEN_ERROR": 7,
		"GET_MUXI_FEED_ERROR":          8,
		"INSERT_MUXI_FEED_ERROR":       9,
		"REMOVE_MUXI_FEED_ERROR":       10,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_feed_v1_feed_error_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_feed_v1_feed_error_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_feed_v1_feed_error_proto_rawDescGZIP(), []int{0}
}

var File_feed_v1_feed_error_proto protoreflect.FileDescriptor

const file_feed_v1_feed_error_proto_rawDesc = "" +
	"\n" +
	"\x18feed/v1/feed_error.proto\x12\afeed.v1\x1a\x13errors/errors.proto*\x8a\x03\n" +
	"\vErrorReason\x12\x1d\n" +
	"\x13TOKEN_ALREADY_EXIST\x10\x00\x1a\x04\xa8E\xf5\x03\x12\x18\n" +
	"\x0eUSER_NOT_FOUND\x10\x01\x1a\x04\xa8E\xf6\x03\x12\x1e\n" +
	"\x14GET_FEED_EVENT_ERROR\x10\x02\x1a\x04\xa8E\xf7\x03\x12 \n" +
	"\x16CLEAR_FEED_EVENT_ERROR\x10\x03\x1a\x04\xa8E\xf8\x03\x12!\n" +
	"\x17PUBLIC_FEED_EVENT_ERROR\x10\x04\x1a\x04\xa8E\xf9\x03\x12$\n" +
	"\x1aFIND_CONFIG_OR_TOKEN_ERROR\x10\x05\x1a\x04\xa8E\xfa\x03\x12&\n" +
	"\x1cCHANGE_CONFIG_OR_TOKEN_ERROR\x10\x06\x1a\x04\xa8E\xfb\x03\x12&\n" +
	"\x1cREMOVE_CONFIG_OR_TOKEN_ERROR\x10\a\x1a\x04\xa8E\xfc\x03\x12\x1d\n" +
	"\x13GET_MUXI_FEED_ERROR\x10\b\x1a\x04\xa8E\xfd\x03\x12 \n" +
	"\x16INSERT_MUXI_FEED_ERROR\x10\t\x1a\x04\xa8E\xfe\x03\x12 \n" +
	"\x16REMOVE_MUXI_FEED_ERROR\x10\n" +
	"\x1a\x04\xa8E\xff\x03\x1a\x04\xa0E\xf4\x03B@Z>github.com/asynccnu/ccnubox-be/be-api/gen/proto/feed/v1;feedv1b\x06proto3"

var (
	file_feed_v1_feed_error_proto_rawDescOnce sync.Once
	file_feed_v1_feed_error_proto_rawDescData []byte
)

func file_feed_v1_feed_error_proto_rawDescGZIP() []byte {
	file_feed_v1_feed_error_proto_rawDescOnce.Do(func() {
		file_feed_v1_feed_error_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_feed_v1_feed_error_proto_rawDesc), len(file_feed_v1_feed_error_proto_rawDesc)))
	})
	return file_feed_v1_feed_error_proto_rawDescData
}

var file_feed_v1_feed_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_feed_v1_feed_error_proto_goTypes = []any{
	(ErrorReason)(0), // 0: feed.v1.ErrorReason
}
var file_feed_v1_feed_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_feed_v1_feed_error_proto_init() }
func file_feed_v1_feed_error_proto_init() {
	if File_feed_v1_feed_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_feed_v1_feed_error_proto_rawDesc), len(file_feed_v1_feed_error_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_feed_v1_feed_error_proto_goTypes,
		DependencyIndexes: file_feed_v1_feed_error_proto_depIdxs,
		EnumInfos:         file_feed_v1_feed_error_proto_enumTypes,
	}.Build()
	File_feed_v1_feed_error_proto = out.File
	file_feed_v1_feed_error_proto_goTypes = nil
	file_feed_v1_feed_error_proto_depIdxs = nil
}
