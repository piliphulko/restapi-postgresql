// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: server-account-aut.proto

package aaGrpc

import (
	basic "github.com/piliphulko/marketplace-example/internal/proto-genr/basic"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_server_account_aut_proto protoreflect.FileDescriptor

var file_server_account_aut_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2d, 0x61, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x61, 0x67, 0x72,
	0x70, 0x63, 0x1a, 0x0b, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xa0, 0x01, 0x0a, 0x0a, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x12, 0x2c,
	0x0a, 0x0a, 0x41, 0x75, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x2e, 0x62,
	0x61, 0x73, 0x69, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x73, 0x73, 0x1a, 0x0c,
	0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x31, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x2e,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x1a, 0x0c, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x31, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x12, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0c, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x69, 0x6c, 0x69, 0x70, 0x68, 0x75, 0x6c, 0x6b, 0x6f, 0x2f, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67,
	0x65, 0x6e, 0x72, 0x2f, 0x61, 0x61, 0x47, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_server_account_aut_proto_goTypes = []interface{}{
	(*basic.LoginPass)(nil),   // 0: basic.LoginPass
	(*basic.AccountInfo)(nil), // 1: basic.AccountInfo
	(*basic.Reply)(nil),       // 2: basic.Reply
}
var file_server_account_aut_proto_depIdxs = []int32{
	0, // 0: aagrpc.AccountAut.AutAccount:input_type -> basic.LoginPass
	1, // 1: aagrpc.AccountAut.CreateAccount:input_type -> basic.AccountInfo
	1, // 2: aagrpc.AccountAut.UpdateAccount:input_type -> basic.AccountInfo
	2, // 3: aagrpc.AccountAut.AutAccount:output_type -> basic.Reply
	2, // 4: aagrpc.AccountAut.CreateAccount:output_type -> basic.Reply
	2, // 5: aagrpc.AccountAut.UpdateAccount:output_type -> basic.Reply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_server_account_aut_proto_init() }
func file_server_account_aut_proto_init() {
	if File_server_account_aut_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_server_account_aut_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_account_aut_proto_goTypes,
		DependencyIndexes: file_server_account_aut_proto_depIdxs,
	}.Build()
	File_server_account_aut_proto = out.File
	file_server_account_aut_proto_rawDesc = nil
	file_server_account_aut_proto_goTypes = nil
	file_server_account_aut_proto_depIdxs = nil
}
