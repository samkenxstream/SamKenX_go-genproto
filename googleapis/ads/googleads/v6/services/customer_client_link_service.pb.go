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
// source: google/ads/googleads/v6/services/customer_client_link_service.proto

package services

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v6/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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

// Request message for [CustomerClientLinkService.GetCustomerClientLink][google.ads.googleads.v6.services.CustomerClientLinkService.GetCustomerClientLink].
type GetCustomerClientLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the customer client link to fetch.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *GetCustomerClientLinkRequest) Reset() {
	*x = GetCustomerClientLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v6_services_customer_client_link_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerClientLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerClientLinkRequest) ProtoMessage() {}

func (x *GetCustomerClientLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v6_services_customer_client_link_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerClientLinkRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerClientLinkRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v6_services_customer_client_link_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetCustomerClientLinkRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

// Request message for [CustomerClientLinkService.MutateCustomerClientLink][google.ads.googleads.v6.services.CustomerClientLinkService.MutateCustomerClientLink].
type MutateCustomerClientLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ID of the customer whose customer link are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The operation to perform on the individual CustomerClientLink.
	Operation *CustomerClientLinkOperation `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
}

func (x *MutateCustomerClientLinkRequest) Reset() {
	*x = MutateCustomerClientLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v6_services_customer_client_link_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCustomerClientLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCustomerClientLinkRequest) ProtoMessage() {}

func (x *MutateCustomerClientLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v6_services_customer_client_link_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCustomerClientLinkRequest.ProtoReflect.Descriptor instead.
func (*MutateCustomerClientLinkRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v6_services_customer_client_link_service_proto_rawDescGZIP(), []int{1}
}

func (x *MutateCustomerClientLinkRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateCustomerClientLinkRequest) GetOperation() *CustomerClientLinkOperation {
	if x != nil {
		return x.Operation
	}
	return nil
}

// A single operation (create, update) on a CustomerClientLink.
type CustomerClientLinkOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are assignable to Operation:
	//	*CustomerClientLinkOperation_Create
	//	*CustomerClientLinkOperation_Update
	Operation isCustomerClientLinkOperation_Operation `protobuf_oneof:"operation"`
}

func (x *CustomerClientLinkOperation) Reset() {
	*x = CustomerClientLinkOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v6_services_customer_client_link_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerClientLinkOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerClientLinkOperation) ProtoMessage() {}

func (x *CustomerClientLinkOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v6_services_customer_client_link_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerClientLinkOperation.ProtoReflect.Descriptor instead.
func (*CustomerClientLinkOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v6_services_customer_client_link_service_proto_rawDescGZIP(), []int{2}
}

func (x *CustomerClientLinkOperation) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (m *CustomerClientLinkOperation) GetOperation() isCustomerClientLinkOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (x *CustomerClientLinkOperation) GetCreate() *resources.CustomerClientLink {
	if x, ok := x.GetOperation().(*CustomerClientLinkOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (x *CustomerClientLinkOperation) GetUpdate() *resources.CustomerClientLink {
	if x, ok := x.GetOperation().(*CustomerClientLinkOperation_Update); ok {
		return x.Update
	}
	return nil
}

type isCustomerClientLinkOperation_Operation interface {
	isCustomerClientLinkOperation_Operation()
}

type CustomerClientLinkOperation_Create struct {
	// Create operation: No resource name is expected for the new link.
	Create *resources.CustomerClientLink `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CustomerClientLinkOperation_Update struct {
	// Update operation: The link is expected to have a valid resource name.
	Update *resources.CustomerClientLink `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

func (*CustomerClientLinkOperation_Create) isCustomerClientLinkOperation_Operation() {}

func (*CustomerClientLinkOperation_Update) isCustomerClientLinkOperation_Operation() {}

// Response message for a CustomerClientLink mutate.
type MutateCustomerClientLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A result that identifies the resource affected by the mutate request.
	Result *MutateCustomerClientLinkResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *MutateCustomerClientLinkResponse) Reset() {
	*x = MutateCustomerClientLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v6_services_customer_client_link_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCustomerClientLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCustomerClientLinkResponse) ProtoMessage() {}

func (x *MutateCustomerClientLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v6_services_customer_client_link_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCustomerClientLinkResponse.ProtoReflect.Descriptor instead.
func (*MutateCustomerClientLinkResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v6_services_customer_client_link_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateCustomerClientLinkResponse) GetResult() *MutateCustomerClientLinkResult {
	if x != nil {
		return x.Result
	}
	return nil
}

// The result for a single customer client link mutate.
type MutateCustomerClientLinkResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *MutateCustomerClientLinkResult) Reset() {
	*x = MutateCustomerClientLinkResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v6_services_customer_client_link_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCustomerClientLinkResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCustomerClientLinkResult) ProtoMessage() {}

func (x *MutateCustomerClientLinkResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v6_services_customer_client_link_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCustomerClientLinkResult.ProtoReflect.Descriptor instead.
func (*MutateCustomerClientLinkResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v6_services_customer_client_link_service_proto_rawDescGZIP(), []int{4}
}

func (x *MutateCustomerClientLinkResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

var File_google_ads_googleads_v6_services_customer_client_link_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v6_services_customer_client_link_service_proto_rawDesc = []byte{
	0x0a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x36, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x36, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x36,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x78, 0x0a, 0x1c, 0x47, 0x65,
	0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c,
	0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x58, 0x0a, 0x0d, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x33, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2d, 0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0xa9, 0x01, 0x0a, 0x1f, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x60,
	0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x89, 0x02, 0x0a, 0x1b, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73,
	0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x4f, 0x0a,
	0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x36, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x4c, 0x69, 0x6e, 0x6b, 0x48, 0x00, 0x52, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x4f,
	0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x36, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x48, 0x00, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x0b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7c, 0x0a, 0x20,
	0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x58, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x45, 0x0a, 0x1e, 0x4d, 0x75,
	0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x32, 0x99, 0x04, 0x0a, 0x19, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0xdd, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69,
	0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x36, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b,
	0x22, 0x4d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x37, 0x12, 0x35, 0x2f, 0x76, 0x36, 0x2f, 0x7b, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x2f, 0x2a, 0x7d, 0xda,
	0x41, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0xfe, 0x01, 0x0a, 0x18, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x41, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x5b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3d, 0x22, 0x38, 0x2f, 0x76, 0x36,
	0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x3a, 0x6d,
	0x75, 0x74, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0xda, 0x41, 0x15, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x1b, 0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x42, 0x85, 0x02,
	0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x1e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x36,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e,
	0x56, 0x36, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x20, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x5c, 0x56, 0x36, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02,
	0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x36, 0x3a, 0x3a, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v6_services_customer_client_link_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v6_services_customer_client_link_service_proto_rawDescData = file_google_ads_googleads_v6_services_customer_client_link_service_proto_rawDesc
)

func file_google_ads_googleads_v6_services_customer_client_link_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v6_services_customer_client_link_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v6_services_customer_client_link_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v6_services_customer_client_link_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v6_services_customer_client_link_service_proto_rawDescData
}

var file_google_ads_googleads_v6_services_customer_client_link_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_ads_googleads_v6_services_customer_client_link_service_proto_goTypes = []interface{}{
	(*GetCustomerClientLinkRequest)(nil),     // 0: google.ads.googleads.v6.services.GetCustomerClientLinkRequest
	(*MutateCustomerClientLinkRequest)(nil),  // 1: google.ads.googleads.v6.services.MutateCustomerClientLinkRequest
	(*CustomerClientLinkOperation)(nil),      // 2: google.ads.googleads.v6.services.CustomerClientLinkOperation
	(*MutateCustomerClientLinkResponse)(nil), // 3: google.ads.googleads.v6.services.MutateCustomerClientLinkResponse
	(*MutateCustomerClientLinkResult)(nil),   // 4: google.ads.googleads.v6.services.MutateCustomerClientLinkResult
	(*fieldmaskpb.FieldMask)(nil),            // 5: google.protobuf.FieldMask
	(*resources.CustomerClientLink)(nil),     // 6: google.ads.googleads.v6.resources.CustomerClientLink
}
var file_google_ads_googleads_v6_services_customer_client_link_service_proto_depIdxs = []int32{
	2, // 0: google.ads.googleads.v6.services.MutateCustomerClientLinkRequest.operation:type_name -> google.ads.googleads.v6.services.CustomerClientLinkOperation
	5, // 1: google.ads.googleads.v6.services.CustomerClientLinkOperation.update_mask:type_name -> google.protobuf.FieldMask
	6, // 2: google.ads.googleads.v6.services.CustomerClientLinkOperation.create:type_name -> google.ads.googleads.v6.resources.CustomerClientLink
	6, // 3: google.ads.googleads.v6.services.CustomerClientLinkOperation.update:type_name -> google.ads.googleads.v6.resources.CustomerClientLink
	4, // 4: google.ads.googleads.v6.services.MutateCustomerClientLinkResponse.result:type_name -> google.ads.googleads.v6.services.MutateCustomerClientLinkResult
	0, // 5: google.ads.googleads.v6.services.CustomerClientLinkService.GetCustomerClientLink:input_type -> google.ads.googleads.v6.services.GetCustomerClientLinkRequest
	1, // 6: google.ads.googleads.v6.services.CustomerClientLinkService.MutateCustomerClientLink:input_type -> google.ads.googleads.v6.services.MutateCustomerClientLinkRequest
	6, // 7: google.ads.googleads.v6.services.CustomerClientLinkService.GetCustomerClientLink:output_type -> google.ads.googleads.v6.resources.CustomerClientLink
	3, // 8: google.ads.googleads.v6.services.CustomerClientLinkService.MutateCustomerClientLink:output_type -> google.ads.googleads.v6.services.MutateCustomerClientLinkResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v6_services_customer_client_link_service_proto_init() }
func file_google_ads_googleads_v6_services_customer_client_link_service_proto_init() {
	if File_google_ads_googleads_v6_services_customer_client_link_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v6_services_customer_client_link_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerClientLinkRequest); i {
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
		file_google_ads_googleads_v6_services_customer_client_link_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateCustomerClientLinkRequest); i {
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
		file_google_ads_googleads_v6_services_customer_client_link_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerClientLinkOperation); i {
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
		file_google_ads_googleads_v6_services_customer_client_link_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateCustomerClientLinkResponse); i {
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
		file_google_ads_googleads_v6_services_customer_client_link_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateCustomerClientLinkResult); i {
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
	file_google_ads_googleads_v6_services_customer_client_link_service_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*CustomerClientLinkOperation_Create)(nil),
		(*CustomerClientLinkOperation_Update)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v6_services_customer_client_link_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v6_services_customer_client_link_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v6_services_customer_client_link_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v6_services_customer_client_link_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v6_services_customer_client_link_service_proto = out.File
	file_google_ads_googleads_v6_services_customer_client_link_service_proto_rawDesc = nil
	file_google_ads_googleads_v6_services_customer_client_link_service_proto_goTypes = nil
	file_google_ads_googleads_v6_services_customer_client_link_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CustomerClientLinkServiceClient is the client API for CustomerClientLinkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerClientLinkServiceClient interface {
	// Returns the requested CustomerClientLink in full detail.
	GetCustomerClientLink(ctx context.Context, in *GetCustomerClientLinkRequest, opts ...grpc.CallOption) (*resources.CustomerClientLink, error)
	// Creates or updates a customer client link. Operation statuses are returned.
	MutateCustomerClientLink(ctx context.Context, in *MutateCustomerClientLinkRequest, opts ...grpc.CallOption) (*MutateCustomerClientLinkResponse, error)
}

type customerClientLinkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerClientLinkServiceClient(cc grpc.ClientConnInterface) CustomerClientLinkServiceClient {
	return &customerClientLinkServiceClient{cc}
}

func (c *customerClientLinkServiceClient) GetCustomerClientLink(ctx context.Context, in *GetCustomerClientLinkRequest, opts ...grpc.CallOption) (*resources.CustomerClientLink, error) {
	out := new(resources.CustomerClientLink)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v6.services.CustomerClientLinkService/GetCustomerClientLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClientLinkServiceClient) MutateCustomerClientLink(ctx context.Context, in *MutateCustomerClientLinkRequest, opts ...grpc.CallOption) (*MutateCustomerClientLinkResponse, error) {
	out := new(MutateCustomerClientLinkResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v6.services.CustomerClientLinkService/MutateCustomerClientLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerClientLinkServiceServer is the server API for CustomerClientLinkService service.
type CustomerClientLinkServiceServer interface {
	// Returns the requested CustomerClientLink in full detail.
	GetCustomerClientLink(context.Context, *GetCustomerClientLinkRequest) (*resources.CustomerClientLink, error)
	// Creates or updates a customer client link. Operation statuses are returned.
	MutateCustomerClientLink(context.Context, *MutateCustomerClientLinkRequest) (*MutateCustomerClientLinkResponse, error)
}

// UnimplementedCustomerClientLinkServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerClientLinkServiceServer struct {
}

func (*UnimplementedCustomerClientLinkServiceServer) GetCustomerClientLink(context.Context, *GetCustomerClientLinkRequest) (*resources.CustomerClientLink, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerClientLink not implemented")
}
func (*UnimplementedCustomerClientLinkServiceServer) MutateCustomerClientLink(context.Context, *MutateCustomerClientLinkRequest) (*MutateCustomerClientLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCustomerClientLink not implemented")
}

func RegisterCustomerClientLinkServiceServer(s *grpc.Server, srv CustomerClientLinkServiceServer) {
	s.RegisterService(&_CustomerClientLinkService_serviceDesc, srv)
}

func _CustomerClientLinkService_GetCustomerClientLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerClientLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerClientLinkServiceServer).GetCustomerClientLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v6.services.CustomerClientLinkService/GetCustomerClientLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerClientLinkServiceServer).GetCustomerClientLink(ctx, req.(*GetCustomerClientLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerClientLinkService_MutateCustomerClientLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerClientLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerClientLinkServiceServer).MutateCustomerClientLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v6.services.CustomerClientLinkService/MutateCustomerClientLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerClientLinkServiceServer).MutateCustomerClientLink(ctx, req.(*MutateCustomerClientLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerClientLinkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v6.services.CustomerClientLinkService",
	HandlerType: (*CustomerClientLinkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomerClientLink",
			Handler:    _CustomerClientLinkService_GetCustomerClientLink_Handler,
		},
		{
			MethodName: "MutateCustomerClientLink",
			Handler:    _CustomerClientLinkService_MutateCustomerClientLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v6/services/customer_client_link_service.proto",
}
