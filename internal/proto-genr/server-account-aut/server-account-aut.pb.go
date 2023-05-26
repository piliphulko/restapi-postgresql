// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: server-account-aut.proto

package server_account_aut

import (
	basic "github.com/piliphulko/marketplace-example/internal/proto-genr/basic"
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

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply basic.REPLY `protobuf:"varint,1,opt,name=reply,proto3,enum=basic.REPLY" json:"reply,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_account_aut_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_server_account_aut_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_server_account_aut_proto_rawDescGZIP(), []int{0}
}

func (x *Reply) GetReply() basic.REPLY {
	if x != nil {
		return x.Reply
	}
	return basic.REPLY(0)
}

type LoginPass struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to AccountChoice:
	//
	//	*LoginPass_CustomerLoginPass
	//	*LoginPass_WarehouseLoginPass
	//	*LoginPass_VendorLoginPass
	AccountChoice isLoginPass_AccountChoice `protobuf_oneof:"account_choice"`
}

func (x *LoginPass) Reset() {
	*x = LoginPass{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_account_aut_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginPass) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginPass) ProtoMessage() {}

func (x *LoginPass) ProtoReflect() protoreflect.Message {
	mi := &file_server_account_aut_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginPass.ProtoReflect.Descriptor instead.
func (*LoginPass) Descriptor() ([]byte, []int) {
	return file_server_account_aut_proto_rawDescGZIP(), []int{1}
}

func (m *LoginPass) GetAccountChoice() isLoginPass_AccountChoice {
	if m != nil {
		return m.AccountChoice
	}
	return nil
}

func (x *LoginPass) GetCustomerLoginPass() *basic.CustomerAut {
	if x, ok := x.GetAccountChoice().(*LoginPass_CustomerLoginPass); ok {
		return x.CustomerLoginPass
	}
	return nil
}

func (x *LoginPass) GetWarehouseLoginPass() *basic.WarehouseAut {
	if x, ok := x.GetAccountChoice().(*LoginPass_WarehouseLoginPass); ok {
		return x.WarehouseLoginPass
	}
	return nil
}

func (x *LoginPass) GetVendorLoginPass() *basic.VendorAut {
	if x, ok := x.GetAccountChoice().(*LoginPass_VendorLoginPass); ok {
		return x.VendorLoginPass
	}
	return nil
}

type isLoginPass_AccountChoice interface {
	isLoginPass_AccountChoice()
}

type LoginPass_CustomerLoginPass struct {
	CustomerLoginPass *basic.CustomerAut `protobuf:"bytes,1,opt,name=customer_login_pass,json=customerLoginPass,proto3,oneof"`
}

type LoginPass_WarehouseLoginPass struct {
	WarehouseLoginPass *basic.WarehouseAut `protobuf:"bytes,2,opt,name=warehouse_login_pass,json=warehouseLoginPass,proto3,oneof"`
}

type LoginPass_VendorLoginPass struct {
	VendorLoginPass *basic.VendorAut `protobuf:"bytes,3,opt,name=vendor_login_pass,json=vendorLoginPass,proto3,oneof"`
}

func (*LoginPass_CustomerLoginPass) isLoginPass_AccountChoice() {}

func (*LoginPass_WarehouseLoginPass) isLoginPass_AccountChoice() {}

func (*LoginPass_VendorLoginPass) isLoginPass_AccountChoice() {}

type AccountInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to AccountInfo:
	//
	//	*AccountInfo_CustomerInfo
	//	*AccountInfo_WarehouseInfo
	//	*AccountInfo_VendorInfo
	AccountInfo isAccountInfo_AccountInfo `protobuf_oneof:"account_info"`
}

func (x *AccountInfo) Reset() {
	*x = AccountInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_account_aut_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountInfo) ProtoMessage() {}

func (x *AccountInfo) ProtoReflect() protoreflect.Message {
	mi := &file_server_account_aut_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountInfo.ProtoReflect.Descriptor instead.
func (*AccountInfo) Descriptor() ([]byte, []int) {
	return file_server_account_aut_proto_rawDescGZIP(), []int{2}
}

func (m *AccountInfo) GetAccountInfo() isAccountInfo_AccountInfo {
	if m != nil {
		return m.AccountInfo
	}
	return nil
}

func (x *AccountInfo) GetCustomerInfo() *basic.CustomerInfo {
	if x, ok := x.GetAccountInfo().(*AccountInfo_CustomerInfo); ok {
		return x.CustomerInfo
	}
	return nil
}

func (x *AccountInfo) GetWarehouseInfo() *basic.WarehouseInfo {
	if x, ok := x.GetAccountInfo().(*AccountInfo_WarehouseInfo); ok {
		return x.WarehouseInfo
	}
	return nil
}

func (x *AccountInfo) GetVendorInfo() *basic.VendorInfo {
	if x, ok := x.GetAccountInfo().(*AccountInfo_VendorInfo); ok {
		return x.VendorInfo
	}
	return nil
}

type isAccountInfo_AccountInfo interface {
	isAccountInfo_AccountInfo()
}

type AccountInfo_CustomerInfo struct {
	CustomerInfo *basic.CustomerInfo `protobuf:"bytes,1,opt,name=customer_info,json=customerInfo,proto3,oneof"`
}

type AccountInfo_WarehouseInfo struct {
	WarehouseInfo *basic.WarehouseInfo `protobuf:"bytes,2,opt,name=warehouse_info,json=warehouseInfo,proto3,oneof"`
}

type AccountInfo_VendorInfo struct {
	VendorInfo *basic.VendorInfo `protobuf:"bytes,3,opt,name=vendor_info,json=vendorInfo,proto3,oneof"`
}

func (*AccountInfo_CustomerInfo) isAccountInfo_AccountInfo() {}

func (*AccountInfo_WarehouseInfo) isAccountInfo_AccountInfo() {}

func (*AccountInfo_VendorInfo) isAccountInfo_AccountInfo() {}

var File_server_account_aut_proto protoreflect.FileDescriptor

var file_server_account_aut_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2d, 0x61, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x61, 0x75, 0x74, 0x1a, 0x0b, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x05, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x22, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0c, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x52, 0x45, 0x50, 0x4c, 0x59, 0x52,
	0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x22, 0xec, 0x01, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x50, 0x61, 0x73, 0x73, 0x12, 0x44, 0x0a, 0x13, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x41, 0x75, 0x74, 0x48, 0x00, 0x52, 0x11, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x73, 0x73, 0x12, 0x47, 0x0a, 0x14, 0x77, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x70, 0x61,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x41, 0x75, 0x74, 0x48, 0x00, 0x52,
	0x12, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50,
	0x61, 0x73, 0x73, 0x12, 0x3e, 0x0a, 0x11, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x41, 0x75, 0x74,
	0x48, 0x00, 0x52, 0x0f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50,
	0x61, 0x73, 0x73, 0x42, 0x10, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63,
	0x68, 0x6f, 0x69, 0x63, 0x65, 0x22, 0xce, 0x01, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3a, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62,
	0x61, 0x73, 0x69, 0x63, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x48, 0x00, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x3d, 0x0a, 0x0e, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48,
	0x00, 0x52, 0x0d, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x34, 0x0a, 0x0b, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x56, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0a, 0x76, 0x65, 0x6e, 0x64,
	0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x0e, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xe2, 0x01, 0x0a, 0x0a, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x41, 0x75, 0x74, 0x12, 0x42, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x61, 0x75, 0x74, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x73, 0x73,
	0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x61, 0x75, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x47, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x61, 0x75, 0x74, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x61, 0x75, 0x74, 0x2e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x47, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x61, 0x75, 0x74, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x61, 0x75, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x52, 0x5a, 0x50, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x69, 0x6c, 0x69, 0x70, 0x68,
	0x75, 0x6c, 0x6b, 0x6f, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x72, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2d, 0x61, 0x75, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_account_aut_proto_rawDescOnce sync.Once
	file_server_account_aut_proto_rawDescData = file_server_account_aut_proto_rawDesc
)

func file_server_account_aut_proto_rawDescGZIP() []byte {
	file_server_account_aut_proto_rawDescOnce.Do(func() {
		file_server_account_aut_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_account_aut_proto_rawDescData)
	})
	return file_server_account_aut_proto_rawDescData
}

var file_server_account_aut_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_server_account_aut_proto_goTypes = []interface{}{
	(*Reply)(nil),               // 0: serveraccountaut.Reply
	(*LoginPass)(nil),           // 1: serveraccountaut.LoginPass
	(*AccountInfo)(nil),         // 2: serveraccountaut.AccountInfo
	(basic.REPLY)(0),            // 3: basic.REPLY
	(*basic.CustomerAut)(nil),   // 4: basic.CustomerAut
	(*basic.WarehouseAut)(nil),  // 5: basic.WarehouseAut
	(*basic.VendorAut)(nil),     // 6: basic.VendorAut
	(*basic.CustomerInfo)(nil),  // 7: basic.CustomerInfo
	(*basic.WarehouseInfo)(nil), // 8: basic.WarehouseInfo
	(*basic.VendorInfo)(nil),    // 9: basic.VendorInfo
}
var file_server_account_aut_proto_depIdxs = []int32{
	3,  // 0: serveraccountaut.Reply.reply:type_name -> basic.REPLY
	4,  // 1: serveraccountaut.LoginPass.customer_login_pass:type_name -> basic.CustomerAut
	5,  // 2: serveraccountaut.LoginPass.warehouse_login_pass:type_name -> basic.WarehouseAut
	6,  // 3: serveraccountaut.LoginPass.vendor_login_pass:type_name -> basic.VendorAut
	7,  // 4: serveraccountaut.AccountInfo.customer_info:type_name -> basic.CustomerInfo
	8,  // 5: serveraccountaut.AccountInfo.warehouse_info:type_name -> basic.WarehouseInfo
	9,  // 6: serveraccountaut.AccountInfo.vendor_info:type_name -> basic.VendorInfo
	1,  // 7: serveraccountaut.AccountAut.AutAccount:input_type -> serveraccountaut.LoginPass
	2,  // 8: serveraccountaut.AccountAut.CreateAccount:input_type -> serveraccountaut.AccountInfo
	2,  // 9: serveraccountaut.AccountAut.UpdateAccount:input_type -> serveraccountaut.AccountInfo
	0,  // 10: serveraccountaut.AccountAut.AutAccount:output_type -> serveraccountaut.Reply
	0,  // 11: serveraccountaut.AccountAut.CreateAccount:output_type -> serveraccountaut.Reply
	0,  // 12: serveraccountaut.AccountAut.UpdateAccount:output_type -> serveraccountaut.Reply
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_server_account_aut_proto_init() }
func file_server_account_aut_proto_init() {
	if File_server_account_aut_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_account_aut_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
		file_server_account_aut_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginPass); i {
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
		file_server_account_aut_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountInfo); i {
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
	file_server_account_aut_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*LoginPass_CustomerLoginPass)(nil),
		(*LoginPass_WarehouseLoginPass)(nil),
		(*LoginPass_VendorLoginPass)(nil),
	}
	file_server_account_aut_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*AccountInfo_CustomerInfo)(nil),
		(*AccountInfo_WarehouseInfo)(nil),
		(*AccountInfo_VendorInfo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_server_account_aut_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_account_aut_proto_goTypes,
		DependencyIndexes: file_server_account_aut_proto_depIdxs,
		MessageInfos:      file_server_account_aut_proto_msgTypes,
	}.Build()
	File_server_account_aut_proto = out.File
	file_server_account_aut_proto_rawDesc = nil
	file_server_account_aut_proto_goTypes = nil
	file_server_account_aut_proto_depIdxs = nil
}
