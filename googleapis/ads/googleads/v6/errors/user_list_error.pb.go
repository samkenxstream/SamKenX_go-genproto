// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: google/ads/googleads/v6/errors/user_list_error.proto

package errors

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Enum describing possible user list errors.
type UserListErrorEnum_UserListError int32

const (
	// Enum unspecified.
	UserListErrorEnum_UNSPECIFIED UserListErrorEnum_UserListError = 0
	// The received error code is not known in this version.
	UserListErrorEnum_UNKNOWN UserListErrorEnum_UserListError = 1
	// Creating and updating external remarketing user lists is not supported.
	UserListErrorEnum_EXTERNAL_REMARKETING_USER_LIST_MUTATE_NOT_SUPPORTED UserListErrorEnum_UserListError = 2
	// Concrete type of user list is required.
	UserListErrorEnum_CONCRETE_TYPE_REQUIRED UserListErrorEnum_UserListError = 3
	// Creating/updating user list conversion types requires specifying the
	// conversion type Id.
	UserListErrorEnum_CONVERSION_TYPE_ID_REQUIRED UserListErrorEnum_UserListError = 4
	// Remarketing user list cannot have duplicate conversion types.
	UserListErrorEnum_DUPLICATE_CONVERSION_TYPES UserListErrorEnum_UserListError = 5
	// Conversion type is invalid/unknown.
	UserListErrorEnum_INVALID_CONVERSION_TYPE UserListErrorEnum_UserListError = 6
	// User list description is empty or invalid.
	UserListErrorEnum_INVALID_DESCRIPTION UserListErrorEnum_UserListError = 7
	// User list name is empty or invalid.
	UserListErrorEnum_INVALID_NAME UserListErrorEnum_UserListError = 8
	// Type of the UserList does not match.
	UserListErrorEnum_INVALID_TYPE UserListErrorEnum_UserListError = 9
	// Embedded logical user lists are not allowed.
	UserListErrorEnum_CAN_NOT_ADD_LOGICAL_LIST_AS_LOGICAL_LIST_OPERAND UserListErrorEnum_UserListError = 10
	// User list rule operand is invalid.
	UserListErrorEnum_INVALID_USER_LIST_LOGICAL_RULE_OPERAND UserListErrorEnum_UserListError = 11
	// Name is already being used for another user list for the account.
	UserListErrorEnum_NAME_ALREADY_USED UserListErrorEnum_UserListError = 12
	// Name is required when creating a new conversion type.
	UserListErrorEnum_NEW_CONVERSION_TYPE_NAME_REQUIRED UserListErrorEnum_UserListError = 13
	// The given conversion type name has been used.
	UserListErrorEnum_CONVERSION_TYPE_NAME_ALREADY_USED UserListErrorEnum_UserListError = 14
	// Only an owner account may edit a user list.
	UserListErrorEnum_OWNERSHIP_REQUIRED_FOR_SET UserListErrorEnum_UserListError = 15
	// Creating user list without setting type in oneof user_list field, or
	// creating/updating read-only user list types is not allowed.
	UserListErrorEnum_USER_LIST_MUTATE_NOT_SUPPORTED UserListErrorEnum_UserListError = 16
	// Rule is invalid.
	UserListErrorEnum_INVALID_RULE UserListErrorEnum_UserListError = 17
	// The specified date range is empty.
	UserListErrorEnum_INVALID_DATE_RANGE UserListErrorEnum_UserListError = 27
	// A UserList which is privacy sensitive or legal rejected cannot be mutated
	// by external users.
	UserListErrorEnum_CAN_NOT_MUTATE_SENSITIVE_USERLIST UserListErrorEnum_UserListError = 28
	// Maximum number of rulebased user lists a customer can have.
	UserListErrorEnum_MAX_NUM_RULEBASED_USERLISTS UserListErrorEnum_UserListError = 29
	// BasicUserList's billable record field cannot be modified once it is set.
	UserListErrorEnum_CANNOT_MODIFY_BILLABLE_RECORD_COUNT UserListErrorEnum_UserListError = 30
	// crm_based_user_list.app_id field must be set when upload_key_type is
	// MOBILE_ADVERTISING_ID.
	UserListErrorEnum_APP_ID_NOT_SET UserListErrorEnum_UserListError = 31
	// Name of the user list is reserved for system generated lists and cannot
	// be used.
	UserListErrorEnum_USERLIST_NAME_IS_RESERVED_FOR_SYSTEM_LIST UserListErrorEnum_UserListError = 32
	// Advertiser needs to be on the allow-list to use remarketing lists created
	// from advertiser uploaded data (e.g., Customer Match lists).
	UserListErrorEnum_ADVERTISER_NOT_ON_ALLOWLIST_FOR_USING_UPLOADED_DATA UserListErrorEnum_UserListError = 37
	// The provided rule_type is not supported for the user list.
	UserListErrorEnum_RULE_TYPE_IS_NOT_SUPPORTED UserListErrorEnum_UserListError = 34
	// Similar user list cannot be used as a logical user list operand.
	UserListErrorEnum_CAN_NOT_ADD_A_SIMILAR_USERLIST_AS_LOGICAL_LIST_OPERAND UserListErrorEnum_UserListError = 35
	// Logical user list should not have a mix of CRM based user list and other
	// types of lists in its rules.
	UserListErrorEnum_CAN_NOT_MIX_CRM_BASED_IN_LOGICAL_LIST_WITH_OTHER_LISTS UserListErrorEnum_UserListError = 36
)

// Enum value maps for UserListErrorEnum_UserListError.
var (
	UserListErrorEnum_UserListError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "EXTERNAL_REMARKETING_USER_LIST_MUTATE_NOT_SUPPORTED",
		3:  "CONCRETE_TYPE_REQUIRED",
		4:  "CONVERSION_TYPE_ID_REQUIRED",
		5:  "DUPLICATE_CONVERSION_TYPES",
		6:  "INVALID_CONVERSION_TYPE",
		7:  "INVALID_DESCRIPTION",
		8:  "INVALID_NAME",
		9:  "INVALID_TYPE",
		10: "CAN_NOT_ADD_LOGICAL_LIST_AS_LOGICAL_LIST_OPERAND",
		11: "INVALID_USER_LIST_LOGICAL_RULE_OPERAND",
		12: "NAME_ALREADY_USED",
		13: "NEW_CONVERSION_TYPE_NAME_REQUIRED",
		14: "CONVERSION_TYPE_NAME_ALREADY_USED",
		15: "OWNERSHIP_REQUIRED_FOR_SET",
		16: "USER_LIST_MUTATE_NOT_SUPPORTED",
		17: "INVALID_RULE",
		27: "INVALID_DATE_RANGE",
		28: "CAN_NOT_MUTATE_SENSITIVE_USERLIST",
		29: "MAX_NUM_RULEBASED_USERLISTS",
		30: "CANNOT_MODIFY_BILLABLE_RECORD_COUNT",
		31: "APP_ID_NOT_SET",
		32: "USERLIST_NAME_IS_RESERVED_FOR_SYSTEM_LIST",
		37: "ADVERTISER_NOT_ON_ALLOWLIST_FOR_USING_UPLOADED_DATA",
		34: "RULE_TYPE_IS_NOT_SUPPORTED",
		35: "CAN_NOT_ADD_A_SIMILAR_USERLIST_AS_LOGICAL_LIST_OPERAND",
		36: "CAN_NOT_MIX_CRM_BASED_IN_LOGICAL_LIST_WITH_OTHER_LISTS",
	}
	UserListErrorEnum_UserListError_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"EXTERNAL_REMARKETING_USER_LIST_MUTATE_NOT_SUPPORTED":    2,
		"CONCRETE_TYPE_REQUIRED":                                 3,
		"CONVERSION_TYPE_ID_REQUIRED":                            4,
		"DUPLICATE_CONVERSION_TYPES":                             5,
		"INVALID_CONVERSION_TYPE":                                6,
		"INVALID_DESCRIPTION":                                    7,
		"INVALID_NAME":                                           8,
		"INVALID_TYPE":                                           9,
		"CAN_NOT_ADD_LOGICAL_LIST_AS_LOGICAL_LIST_OPERAND":       10,
		"INVALID_USER_LIST_LOGICAL_RULE_OPERAND":                 11,
		"NAME_ALREADY_USED":                                      12,
		"NEW_CONVERSION_TYPE_NAME_REQUIRED":                      13,
		"CONVERSION_TYPE_NAME_ALREADY_USED":                      14,
		"OWNERSHIP_REQUIRED_FOR_SET":                             15,
		"USER_LIST_MUTATE_NOT_SUPPORTED":                         16,
		"INVALID_RULE":                                           17,
		"INVALID_DATE_RANGE":                                     27,
		"CAN_NOT_MUTATE_SENSITIVE_USERLIST":                      28,
		"MAX_NUM_RULEBASED_USERLISTS":                            29,
		"CANNOT_MODIFY_BILLABLE_RECORD_COUNT":                    30,
		"APP_ID_NOT_SET":                                         31,
		"USERLIST_NAME_IS_RESERVED_FOR_SYSTEM_LIST":              32,
		"ADVERTISER_NOT_ON_ALLOWLIST_FOR_USING_UPLOADED_DATA":    37,
		"RULE_TYPE_IS_NOT_SUPPORTED":                             34,
		"CAN_NOT_ADD_A_SIMILAR_USERLIST_AS_LOGICAL_LIST_OPERAND": 35,
		"CAN_NOT_MIX_CRM_BASED_IN_LOGICAL_LIST_WITH_OTHER_LISTS": 36,
	}
)

func (x UserListErrorEnum_UserListError) Enum() *UserListErrorEnum_UserListError {
	p := new(UserListErrorEnum_UserListError)
	*p = x
	return p
}

func (x UserListErrorEnum_UserListError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserListErrorEnum_UserListError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v6_errors_user_list_error_proto_enumTypes[0].Descriptor()
}

func (UserListErrorEnum_UserListError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v6_errors_user_list_error_proto_enumTypes[0]
}

func (x UserListErrorEnum_UserListError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserListErrorEnum_UserListError.Descriptor instead.
func (UserListErrorEnum_UserListError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v6_errors_user_list_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible user list errors.
type UserListErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserListErrorEnum) Reset() {
	*x = UserListErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v6_errors_user_list_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListErrorEnum) ProtoMessage() {}

func (x *UserListErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v6_errors_user_list_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListErrorEnum.ProtoReflect.Descriptor instead.
func (*UserListErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v6_errors_user_list_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v6_errors_user_list_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v6_errors_user_list_error_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x36, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x36, 0x2e,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x07, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xd7, 0x07, 0x0a, 0x0d, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x37, 0x0a, 0x33, 0x45, 0x58,
	0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x52, 0x45, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x49,
	0x4e, 0x47, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x4d, 0x55, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45,
	0x44, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x4f, 0x4e, 0x43, 0x52, 0x45, 0x54, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x03, 0x12,
	0x1f, 0x0a, 0x1b, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x49, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x04,
	0x12, 0x1e, 0x0a, 0x1a, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4f,
	0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x53, 0x10, 0x05,
	0x12, 0x1b, 0x0a, 0x17, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x56,
	0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x06, 0x12, 0x17, 0x0a,
	0x13, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x08, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x09, 0x12, 0x34, 0x0a, 0x30, 0x43, 0x41,
	0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x43, 0x41,
	0x4c, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x41, 0x53, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x43, 0x41,
	0x4c, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x4e, 0x44, 0x10, 0x0a,
	0x12, 0x2a, 0x0a, 0x26, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x55, 0x53, 0x45, 0x52,
	0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x43, 0x41, 0x4c, 0x5f, 0x52, 0x55,
	0x4c, 0x45, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x4e, 0x44, 0x10, 0x0b, 0x12, 0x15, 0x0a, 0x11,
	0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x55, 0x53, 0x45,
	0x44, 0x10, 0x0c, 0x12, 0x25, 0x0a, 0x21, 0x4e, 0x45, 0x57, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45,
	0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f,
	0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x0d, 0x12, 0x25, 0x0a, 0x21, 0x43, 0x4f,
	0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x41,
	0x4d, 0x45, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x55, 0x53, 0x45, 0x44, 0x10,
	0x0e, 0x12, 0x1e, 0x0a, 0x1a, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x52,
	0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x53, 0x45, 0x54, 0x10,
	0x0f, 0x12, 0x22, 0x0a, 0x1e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x4d,
	0x55, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52,
	0x54, 0x45, 0x44, 0x10, 0x10, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x5f, 0x52, 0x55, 0x4c, 0x45, 0x10, 0x11, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x1b, 0x12,
	0x25, 0x0a, 0x21, 0x43, 0x41, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x55, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x53, 0x45, 0x4e, 0x53, 0x49, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x52,
	0x4c, 0x49, 0x53, 0x54, 0x10, 0x1c, 0x12, 0x1f, 0x0a, 0x1b, 0x4d, 0x41, 0x58, 0x5f, 0x4e, 0x55,
	0x4d, 0x5f, 0x52, 0x55, 0x4c, 0x45, 0x42, 0x41, 0x53, 0x45, 0x44, 0x5f, 0x55, 0x53, 0x45, 0x52,
	0x4c, 0x49, 0x53, 0x54, 0x53, 0x10, 0x1d, 0x12, 0x27, 0x0a, 0x23, 0x43, 0x41, 0x4e, 0x4e, 0x4f,
	0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59, 0x5f, 0x42, 0x49, 0x4c, 0x4c, 0x41, 0x42, 0x4c,
	0x45, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x1e,
	0x12, 0x12, 0x0a, 0x0e, 0x41, 0x50, 0x50, 0x5f, 0x49, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53,
	0x45, 0x54, 0x10, 0x1f, 0x12, 0x2d, 0x0a, 0x29, 0x55, 0x53, 0x45, 0x52, 0x4c, 0x49, 0x53, 0x54,
	0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x53, 0x45, 0x52, 0x56, 0x45,
	0x44, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f, 0x4c, 0x49, 0x53,
	0x54, 0x10, 0x20, 0x12, 0x37, 0x0a, 0x33, 0x41, 0x44, 0x56, 0x45, 0x52, 0x54, 0x49, 0x53, 0x45,
	0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4f, 0x4e, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x4c, 0x49,
	0x53, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x55, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x50, 0x4c,
	0x4f, 0x41, 0x44, 0x45, 0x44, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x25, 0x12, 0x1e, 0x0a, 0x1a,
	0x52, 0x55, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x53, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x22, 0x12, 0x3a, 0x0a, 0x36,
	0x43, 0x41, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x41, 0x5f, 0x53, 0x49,
	0x4d, 0x49, 0x4c, 0x41, 0x52, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x41,
	0x53, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x43, 0x41, 0x4c, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x4f,
	0x50, 0x45, 0x52, 0x41, 0x4e, 0x44, 0x10, 0x23, 0x12, 0x3a, 0x0a, 0x36, 0x43, 0x41, 0x4e, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x49, 0x58, 0x5f, 0x43, 0x52, 0x4d, 0x5f, 0x42, 0x41, 0x53, 0x45,
	0x44, 0x5f, 0x49, 0x4e, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x43, 0x41, 0x4c, 0x5f, 0x4c, 0x49, 0x53,
	0x54, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x5f, 0x4c, 0x49, 0x53,
	0x54, 0x53, 0x10, 0x24, 0x42, 0xed, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x36, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x12, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x44, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x36, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x36, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02,
	0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x36, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea,
	0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x36, 0x3a, 0x3a, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v6_errors_user_list_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v6_errors_user_list_error_proto_rawDescData = file_google_ads_googleads_v6_errors_user_list_error_proto_rawDesc
)

func file_google_ads_googleads_v6_errors_user_list_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v6_errors_user_list_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v6_errors_user_list_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v6_errors_user_list_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v6_errors_user_list_error_proto_rawDescData
}

var file_google_ads_googleads_v6_errors_user_list_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v6_errors_user_list_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v6_errors_user_list_error_proto_goTypes = []interface{}{
	(UserListErrorEnum_UserListError)(0), // 0: google.ads.googleads.v6.errors.UserListErrorEnum.UserListError
	(*UserListErrorEnum)(nil),            // 1: google.ads.googleads.v6.errors.UserListErrorEnum
}
var file_google_ads_googleads_v6_errors_user_list_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v6_errors_user_list_error_proto_init() }
func file_google_ads_googleads_v6_errors_user_list_error_proto_init() {
	if File_google_ads_googleads_v6_errors_user_list_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v6_errors_user_list_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v6_errors_user_list_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v6_errors_user_list_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v6_errors_user_list_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v6_errors_user_list_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v6_errors_user_list_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v6_errors_user_list_error_proto = out.File
	file_google_ads_googleads_v6_errors_user_list_error_proto_rawDesc = nil
	file_google_ads_googleads_v6_errors_user_list_error_proto_goTypes = nil
	file_google_ads_googleads_v6_errors_user_list_error_proto_depIdxs = nil
}
