// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: microserver-account-authentication.proto

package core

import (
	basic "github.com/piliphulko/marketplace-example/api/basic"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_microserver_account_authentication_proto protoreflect.FileDescriptor

var file_microserver_account_authentication_proto_rawDesc = []byte{
	0x0a, 0x28, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0b,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb8, 0x01, 0x0a, 0x0a, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x12, 0x30, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x50, 0x61, 0x73, 0x73, 0x1a, 0x10, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4a, 0x57, 0x54, 0x12, 0x3b, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x2e, 0x62, 0x61, 0x73,
	0x69, 0x63, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3b, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x42, 0x64, 0x5a, 0x62, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x69, 0x6c, 0x69, 0x70, 0x68, 0x75, 0x6c, 0x6b, 0x6f, 0x2f, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_microserver_account_authentication_proto_goTypes = []interface{}{
	(*basic.LoginPass)(nil),   // 0: basic.LoginPass
	(*basic.AccountInfo)(nil), // 1: basic.AccountInfo
	(*basic.StringJWT)(nil),   // 2: basic.StringJWT
	(*emptypb.Empty)(nil),     // 3: google.protobuf.Empty
}
var file_microserver_account_authentication_proto_depIdxs = []int32{
	0, // 0: microserver_account_authentication.AccountAut.AutAccount:input_type -> basic.LoginPass
	1, // 1: microserver_account_authentication.AccountAut.CreateAccount:input_type -> basic.AccountInfo
	1, // 2: microserver_account_authentication.AccountAut.UpdateAccount:input_type -> basic.AccountInfo
	2, // 3: microserver_account_authentication.AccountAut.AutAccount:output_type -> basic.StringJWT
	3, // 4: microserver_account_authentication.AccountAut.CreateAccount:output_type -> google.protobuf.Empty
	3, // 5: microserver_account_authentication.AccountAut.UpdateAccount:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_microserver_account_authentication_proto_init() }
func file_microserver_account_authentication_proto_init() {
	if File_microserver_account_authentication_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_microserver_account_authentication_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_microserver_account_authentication_proto_goTypes,
		DependencyIndexes: file_microserver_account_authentication_proto_depIdxs,
	}.Build()
	File_microserver_account_authentication_proto = out.File
	file_microserver_account_authentication_proto_rawDesc = nil
	file_microserver_account_authentication_proto_goTypes = nil
	file_microserver_account_authentication_proto_depIdxs = nil
}
