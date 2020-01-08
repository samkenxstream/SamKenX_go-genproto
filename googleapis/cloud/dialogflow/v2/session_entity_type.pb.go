// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/dialogflow/v2/session_entity_type.proto

package dialogflow

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The types of modifications for a session entity type.
type SessionEntityType_EntityOverrideMode int32

const (
	// Not specified. This value should be never used.
	SessionEntityType_ENTITY_OVERRIDE_MODE_UNSPECIFIED SessionEntityType_EntityOverrideMode = 0
	// The collection of session entities overrides the collection of entities
	// in the corresponding developer entity type.
	SessionEntityType_ENTITY_OVERRIDE_MODE_OVERRIDE SessionEntityType_EntityOverrideMode = 1
	// The collection of session entities extends the collection of entities in
	// the corresponding developer entity type.
	//
	// Note: Even in this override mode calls to `ListSessionEntityTypes`,
	// `GetSessionEntityType`, `CreateSessionEntityType` and
	// `UpdateSessionEntityType` only return the additional entities added in
	// this session entity type. If you want to get the supplemented list,
	// please call [EntityTypes.GetEntityType][google.cloud.dialogflow.v2.EntityTypes.GetEntityType] on the developer entity type
	// and merge.
	SessionEntityType_ENTITY_OVERRIDE_MODE_SUPPLEMENT SessionEntityType_EntityOverrideMode = 2
)

var SessionEntityType_EntityOverrideMode_name = map[int32]string{
	0: "ENTITY_OVERRIDE_MODE_UNSPECIFIED",
	1: "ENTITY_OVERRIDE_MODE_OVERRIDE",
	2: "ENTITY_OVERRIDE_MODE_SUPPLEMENT",
}

var SessionEntityType_EntityOverrideMode_value = map[string]int32{
	"ENTITY_OVERRIDE_MODE_UNSPECIFIED": 0,
	"ENTITY_OVERRIDE_MODE_OVERRIDE":    1,
	"ENTITY_OVERRIDE_MODE_SUPPLEMENT":  2,
}

func (x SessionEntityType_EntityOverrideMode) String() string {
	return proto.EnumName(SessionEntityType_EntityOverrideMode_name, int32(x))
}

func (SessionEntityType_EntityOverrideMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_841e168133818181, []int{0, 0}
}

// Represents a session entity type.
//
// Extends or replaces a developer entity type at the user session level (we
// refer to the entity types defined at the agent level as "developer entity
// types").
//
// Note: session entity types apply to all queries, regardless of the language.
type SessionEntityType struct {
	// Required. The unique identifier of this session entity type. Format:
	// `projects/<Project ID>/agent/sessions/<Session ID>/entityTypes/<Entity Type
	// Display Name>`.
	//
	// `<Entity Type Display Name>` must be the display name of an existing entity
	// type in the same agent that will be overridden or supplemented.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. Indicates whether the additional data should override or
	// supplement the developer entity type definition.
	EntityOverrideMode SessionEntityType_EntityOverrideMode `protobuf:"varint,2,opt,name=entity_override_mode,json=entityOverrideMode,proto3,enum=google.cloud.dialogflow.v2.SessionEntityType_EntityOverrideMode" json:"entity_override_mode,omitempty"`
	// Required. The collection of entities associated with this session entity
	// type.
	Entities             []*EntityType_Entity `protobuf:"bytes,3,rep,name=entities,proto3" json:"entities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SessionEntityType) Reset()         { *m = SessionEntityType{} }
func (m *SessionEntityType) String() string { return proto.CompactTextString(m) }
func (*SessionEntityType) ProtoMessage()    {}
func (*SessionEntityType) Descriptor() ([]byte, []int) {
	return fileDescriptor_841e168133818181, []int{0}
}

func (m *SessionEntityType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionEntityType.Unmarshal(m, b)
}
func (m *SessionEntityType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionEntityType.Marshal(b, m, deterministic)
}
func (m *SessionEntityType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionEntityType.Merge(m, src)
}
func (m *SessionEntityType) XXX_Size() int {
	return xxx_messageInfo_SessionEntityType.Size(m)
}
func (m *SessionEntityType) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionEntityType.DiscardUnknown(m)
}

var xxx_messageInfo_SessionEntityType proto.InternalMessageInfo

func (m *SessionEntityType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SessionEntityType) GetEntityOverrideMode() SessionEntityType_EntityOverrideMode {
	if m != nil {
		return m.EntityOverrideMode
	}
	return SessionEntityType_ENTITY_OVERRIDE_MODE_UNSPECIFIED
}

func (m *SessionEntityType) GetEntities() []*EntityType_Entity {
	if m != nil {
		return m.Entities
	}
	return nil
}

// The request message for [SessionEntityTypes.ListSessionEntityTypes][google.cloud.dialogflow.v2.SessionEntityTypes.ListSessionEntityTypes].
type ListSessionEntityTypesRequest struct {
	// Required. The session to list all session entity types from.
	// Format: `projects/<Project ID>/agent/sessions/<Session ID>`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional. The maximum number of items to return in a single page. By
	// default 100 and at most 1000.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. The next_page_token value returned from a previous list request.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListSessionEntityTypesRequest) Reset()         { *m = ListSessionEntityTypesRequest{} }
func (m *ListSessionEntityTypesRequest) String() string { return proto.CompactTextString(m) }
func (*ListSessionEntityTypesRequest) ProtoMessage()    {}
func (*ListSessionEntityTypesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_841e168133818181, []int{1}
}

func (m *ListSessionEntityTypesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSessionEntityTypesRequest.Unmarshal(m, b)
}
func (m *ListSessionEntityTypesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSessionEntityTypesRequest.Marshal(b, m, deterministic)
}
func (m *ListSessionEntityTypesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSessionEntityTypesRequest.Merge(m, src)
}
func (m *ListSessionEntityTypesRequest) XXX_Size() int {
	return xxx_messageInfo_ListSessionEntityTypesRequest.Size(m)
}
func (m *ListSessionEntityTypesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSessionEntityTypesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListSessionEntityTypesRequest proto.InternalMessageInfo

func (m *ListSessionEntityTypesRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListSessionEntityTypesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListSessionEntityTypesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// The response message for [SessionEntityTypes.ListSessionEntityTypes][google.cloud.dialogflow.v2.SessionEntityTypes.ListSessionEntityTypes].
type ListSessionEntityTypesResponse struct {
	// The list of session entity types. There will be a maximum number of items
	// returned based on the page_size field in the request.
	SessionEntityTypes []*SessionEntityType `protobuf:"bytes,1,rep,name=session_entity_types,json=sessionEntityTypes,proto3" json:"session_entity_types,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no
	// more results in the list.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListSessionEntityTypesResponse) Reset()         { *m = ListSessionEntityTypesResponse{} }
func (m *ListSessionEntityTypesResponse) String() string { return proto.CompactTextString(m) }
func (*ListSessionEntityTypesResponse) ProtoMessage()    {}
func (*ListSessionEntityTypesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_841e168133818181, []int{2}
}

func (m *ListSessionEntityTypesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSessionEntityTypesResponse.Unmarshal(m, b)
}
func (m *ListSessionEntityTypesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSessionEntityTypesResponse.Marshal(b, m, deterministic)
}
func (m *ListSessionEntityTypesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSessionEntityTypesResponse.Merge(m, src)
}
func (m *ListSessionEntityTypesResponse) XXX_Size() int {
	return xxx_messageInfo_ListSessionEntityTypesResponse.Size(m)
}
func (m *ListSessionEntityTypesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSessionEntityTypesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListSessionEntityTypesResponse proto.InternalMessageInfo

func (m *ListSessionEntityTypesResponse) GetSessionEntityTypes() []*SessionEntityType {
	if m != nil {
		return m.SessionEntityTypes
	}
	return nil
}

func (m *ListSessionEntityTypesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The request message for [SessionEntityTypes.GetSessionEntityType][google.cloud.dialogflow.v2.SessionEntityTypes.GetSessionEntityType].
type GetSessionEntityTypeRequest struct {
	// Required. The name of the session entity type. Format:
	// `projects/<Project ID>/agent/sessions/<Session ID>/entityTypes/<Entity Type
	// Display Name>`.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSessionEntityTypeRequest) Reset()         { *m = GetSessionEntityTypeRequest{} }
func (m *GetSessionEntityTypeRequest) String() string { return proto.CompactTextString(m) }
func (*GetSessionEntityTypeRequest) ProtoMessage()    {}
func (*GetSessionEntityTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_841e168133818181, []int{3}
}

func (m *GetSessionEntityTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSessionEntityTypeRequest.Unmarshal(m, b)
}
func (m *GetSessionEntityTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSessionEntityTypeRequest.Marshal(b, m, deterministic)
}
func (m *GetSessionEntityTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSessionEntityTypeRequest.Merge(m, src)
}
func (m *GetSessionEntityTypeRequest) XXX_Size() int {
	return xxx_messageInfo_GetSessionEntityTypeRequest.Size(m)
}
func (m *GetSessionEntityTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSessionEntityTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSessionEntityTypeRequest proto.InternalMessageInfo

func (m *GetSessionEntityTypeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request message for [SessionEntityTypes.CreateSessionEntityType][google.cloud.dialogflow.v2.SessionEntityTypes.CreateSessionEntityType].
type CreateSessionEntityTypeRequest struct {
	// Required. The session to create a session entity type for.
	// Format: `projects/<Project ID>/agent/sessions/<Session ID>`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The session entity type to create.
	SessionEntityType    *SessionEntityType `protobuf:"bytes,2,opt,name=session_entity_type,json=sessionEntityType,proto3" json:"session_entity_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CreateSessionEntityTypeRequest) Reset()         { *m = CreateSessionEntityTypeRequest{} }
func (m *CreateSessionEntityTypeRequest) String() string { return proto.CompactTextString(m) }
func (*CreateSessionEntityTypeRequest) ProtoMessage()    {}
func (*CreateSessionEntityTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_841e168133818181, []int{4}
}

func (m *CreateSessionEntityTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSessionEntityTypeRequest.Unmarshal(m, b)
}
func (m *CreateSessionEntityTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSessionEntityTypeRequest.Marshal(b, m, deterministic)
}
func (m *CreateSessionEntityTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSessionEntityTypeRequest.Merge(m, src)
}
func (m *CreateSessionEntityTypeRequest) XXX_Size() int {
	return xxx_messageInfo_CreateSessionEntityTypeRequest.Size(m)
}
func (m *CreateSessionEntityTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSessionEntityTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSessionEntityTypeRequest proto.InternalMessageInfo

func (m *CreateSessionEntityTypeRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateSessionEntityTypeRequest) GetSessionEntityType() *SessionEntityType {
	if m != nil {
		return m.SessionEntityType
	}
	return nil
}

// The request message for [SessionEntityTypes.UpdateSessionEntityType][google.cloud.dialogflow.v2.SessionEntityTypes.UpdateSessionEntityType].
type UpdateSessionEntityTypeRequest struct {
	// Required. The entity type to update. Format:
	// `projects/<Project ID>/agent/sessions/<Session ID>/entityTypes/<Entity Type
	// Display Name>`.
	SessionEntityType *SessionEntityType `protobuf:"bytes,1,opt,name=session_entity_type,json=sessionEntityType,proto3" json:"session_entity_type,omitempty"`
	// Optional. The mask to control which fields get updated.
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateSessionEntityTypeRequest) Reset()         { *m = UpdateSessionEntityTypeRequest{} }
func (m *UpdateSessionEntityTypeRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateSessionEntityTypeRequest) ProtoMessage()    {}
func (*UpdateSessionEntityTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_841e168133818181, []int{5}
}

func (m *UpdateSessionEntityTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateSessionEntityTypeRequest.Unmarshal(m, b)
}
func (m *UpdateSessionEntityTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateSessionEntityTypeRequest.Marshal(b, m, deterministic)
}
func (m *UpdateSessionEntityTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateSessionEntityTypeRequest.Merge(m, src)
}
func (m *UpdateSessionEntityTypeRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateSessionEntityTypeRequest.Size(m)
}
func (m *UpdateSessionEntityTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateSessionEntityTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateSessionEntityTypeRequest proto.InternalMessageInfo

func (m *UpdateSessionEntityTypeRequest) GetSessionEntityType() *SessionEntityType {
	if m != nil {
		return m.SessionEntityType
	}
	return nil
}

func (m *UpdateSessionEntityTypeRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// The request message for [SessionEntityTypes.DeleteSessionEntityType][google.cloud.dialogflow.v2.SessionEntityTypes.DeleteSessionEntityType].
type DeleteSessionEntityTypeRequest struct {
	// Required. The name of the entity type to delete. Format:
	// `projects/<Project ID>/agent/sessions/<Session ID>/entityTypes/<Entity Type
	// Display Name>`.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteSessionEntityTypeRequest) Reset()         { *m = DeleteSessionEntityTypeRequest{} }
func (m *DeleteSessionEntityTypeRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteSessionEntityTypeRequest) ProtoMessage()    {}
func (*DeleteSessionEntityTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_841e168133818181, []int{6}
}

func (m *DeleteSessionEntityTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSessionEntityTypeRequest.Unmarshal(m, b)
}
func (m *DeleteSessionEntityTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSessionEntityTypeRequest.Marshal(b, m, deterministic)
}
func (m *DeleteSessionEntityTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSessionEntityTypeRequest.Merge(m, src)
}
func (m *DeleteSessionEntityTypeRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteSessionEntityTypeRequest.Size(m)
}
func (m *DeleteSessionEntityTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSessionEntityTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSessionEntityTypeRequest proto.InternalMessageInfo

func (m *DeleteSessionEntityTypeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterEnum("google.cloud.dialogflow.v2.SessionEntityType_EntityOverrideMode", SessionEntityType_EntityOverrideMode_name, SessionEntityType_EntityOverrideMode_value)
	proto.RegisterType((*SessionEntityType)(nil), "google.cloud.dialogflow.v2.SessionEntityType")
	proto.RegisterType((*ListSessionEntityTypesRequest)(nil), "google.cloud.dialogflow.v2.ListSessionEntityTypesRequest")
	proto.RegisterType((*ListSessionEntityTypesResponse)(nil), "google.cloud.dialogflow.v2.ListSessionEntityTypesResponse")
	proto.RegisterType((*GetSessionEntityTypeRequest)(nil), "google.cloud.dialogflow.v2.GetSessionEntityTypeRequest")
	proto.RegisterType((*CreateSessionEntityTypeRequest)(nil), "google.cloud.dialogflow.v2.CreateSessionEntityTypeRequest")
	proto.RegisterType((*UpdateSessionEntityTypeRequest)(nil), "google.cloud.dialogflow.v2.UpdateSessionEntityTypeRequest")
	proto.RegisterType((*DeleteSessionEntityTypeRequest)(nil), "google.cloud.dialogflow.v2.DeleteSessionEntityTypeRequest")
}

func init() {
	proto.RegisterFile("google/cloud/dialogflow/v2/session_entity_type.proto", fileDescriptor_841e168133818181)
}

var fileDescriptor_841e168133818181 = []byte{
	// 983 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xcf, 0x6f, 0xdb, 0x54,
	0x1c, 0xe7, 0xb9, 0x30, 0xad, 0x6f, 0x02, 0xba, 0xb7, 0xaa, 0xc9, 0x5c, 0x96, 0x65, 0x1e, 0x42,
	0xa5, 0xb4, 0xb6, 0x94, 0x4d, 0x20, 0x0a, 0x93, 0xe6, 0x26, 0x6e, 0xa9, 0x58, 0xdb, 0x28, 0x4d,
	0x27, 0x8d, 0x8b, 0xe5, 0x26, 0xdf, 0xb8, 0xa6, 0x8e, 0x9f, 0xf1, 0x7b, 0x69, 0xda, 0x45, 0x3d,
	0xb0, 0x7f, 0x81, 0xbf, 0x00, 0x24, 0x0e, 0x70, 0xe4, 0xcc, 0x99, 0x03, 0x47, 0x90, 0x76, 0xe8,
	0x69, 0x07, 0x4e, 0x5c, 0xe1, 0x80, 0x38, 0x21, 0xfb, 0x39, 0x4d, 0x32, 0xff, 0x98, 0x52, 0x71,
	0x8b, 0xbf, 0x3f, 0x3f, 0xdf, 0xcf, 0xf7, 0x47, 0x1e, 0xbe, 0x6f, 0x53, 0x6a, 0xbb, 0xa0, 0xb5,
	0x5c, 0xda, 0x6b, 0x6b, 0x6d, 0xc7, 0x72, 0xa9, 0xdd, 0x71, 0x69, 0x5f, 0x3b, 0xae, 0x68, 0x0c,
	0x18, 0x73, 0xa8, 0x67, 0x82, 0xc7, 0x1d, 0x7e, 0x6a, 0xf2, 0x53, 0x1f, 0x54, 0x3f, 0xa0, 0x9c,
	0x12, 0x59, 0x78, 0xa9, 0x91, 0x97, 0x3a, 0xf2, 0x52, 0x8f, 0x2b, 0xf2, 0x3b, 0x71, 0x44, 0xcb,
	0x77, 0x34, 0xcb, 0xf3, 0x28, 0xb7, 0xb8, 0x43, 0x3d, 0x26, 0x3c, 0xe5, 0xc2, 0x98, 0xb6, 0xe5,
	0x3a, 0xe0, 0xf1, 0x58, 0x71, 0x7b, 0x4c, 0xd1, 0x71, 0xc0, 0x6d, 0x9b, 0x07, 0x70, 0x68, 0x1d,
	0x3b, 0x34, 0x88, 0x0d, 0x6e, 0x8e, 0x19, 0x04, 0xc0, 0x68, 0x2f, 0x68, 0xc5, 0x70, 0xe4, 0x95,
	0x9c, 0x22, 0x12, 0xe0, 0xe5, 0xc5, 0xd8, 0x3a, 0xfa, 0x3a, 0xe8, 0x75, 0x34, 0xe8, 0xfa, 0xfc,
	0x34, 0x56, 0x96, 0x5f, 0x56, 0x0a, 0x2c, 0x5d, 0x8b, 0x1d, 0x09, 0x0b, 0xe5, 0xaf, 0x19, 0x7c,
	0x7d, 0x4f, 0x30, 0x63, 0x44, 0xb1, 0x9b, 0xa7, 0x3e, 0x90, 0x02, 0x7e, 0xdd, 0xb3, 0xba, 0x50,
	0x44, 0x65, 0xb4, 0x34, 0xbb, 0x3e, 0xf3, 0x42, 0x97, 0x1a, 0x91, 0x80, 0xf4, 0xf1, 0x7c, 0x0c,
	0x81, 0x1e, 0x43, 0x10, 0x38, 0x6d, 0x30, 0xbb, 0xb4, 0x0d, 0x45, 0xa9, 0x8c, 0x96, 0xde, 0xaa,
	0x3c, 0x54, 0xb3, 0x99, 0x54, 0x13, 0x59, 0x54, 0xf1, 0x73, 0x37, 0x0e, 0xb4, 0x4d, 0xdb, 0x20,
	0x52, 0x11, 0x48, 0x28, 0xc8, 0x0e, 0xbe, 0x1a, 0x49, 0x1d, 0x60, 0xc5, 0x99, 0xf2, 0xcc, 0xd2,
	0xb5, 0xca, 0x6a, 0x5e, 0xb2, 0x44, 0x16, 0x11, 0xf9, 0x22, 0x86, 0xf2, 0x0c, 0x61, 0x92, 0xcc,
	0x4f, 0xde, 0xc5, 0x65, 0x63, 0xa7, 0xb9, 0xd5, 0x7c, 0x62, 0xee, 0x3e, 0x36, 0x1a, 0x8d, 0xad,
	0x9a, 0x61, 0x6e, 0xef, 0xd6, 0x0c, 0x73, 0x7f, 0x67, 0xaf, 0x6e, 0x54, 0xb7, 0x36, 0xb6, 0x8c,
	0xda, 0xdc, 0x6b, 0xe4, 0x0e, 0xbe, 0x95, 0x6a, 0x35, 0xfc, 0x9a, 0x43, 0xe4, 0x2e, 0xbe, 0x9d,
	0x6a, 0xb2, 0xb7, 0x5f, 0xaf, 0x3f, 0x32, 0xb6, 0x8d, 0x9d, 0xe6, 0x9c, 0xb4, 0xd6, 0xff, 0x53,
	0xe7, 0xf8, 0x83, 0x31, 0xe4, 0xa2, 0x22, 0xcb, 0x77, 0x98, 0xda, 0xa2, 0x5d, 0x2d, 0xd9, 0x18,
	0xc3, 0x0f, 0xe8, 0x97, 0xd0, 0xe2, 0x4c, 0x1b, 0xc4, 0xbf, 0xce, 0x34, 0xcb, 0x06, 0x8f, 0x0f,
	0xc7, 0x9b, 0x69, 0x83, 0xf8, 0xd7, 0x59, 0x3c, 0x2c, 0xa1, 0x1f, 0xd3, 0x06, 0x63, 0x93, 0x73,
	0x16, 0x56, 0x7f, 0xeb, 0x91, 0xc3, 0x78, 0x22, 0x01, 0x6b, 0xc0, 0x57, 0x3d, 0x60, 0x9c, 0x2c,
	0xe2, 0x2b, 0xbe, 0x15, 0x80, 0xc7, 0xc7, 0x67, 0x20, 0x16, 0x91, 0x32, 0x9e, 0xf5, 0x2d, 0x1b,
	0x4c, 0xe6, 0x3c, 0x15, 0xad, 0x7f, 0x23, 0xd4, 0xa3, 0xc6, 0xd5, 0x50, 0xba, 0xe7, 0x3c, 0x05,
	0xa2, 0x60, 0x1c, 0x59, 0x70, 0x7a, 0x04, 0x5e, 0x71, 0x66, 0x18, 0x02, 0x35, 0x22, 0xc7, 0x66,
	0x28, 0x55, 0x7e, 0x40, 0xb8, 0x94, 0x05, 0x82, 0xf9, 0xd4, 0x63, 0x40, 0x4c, 0x3c, 0x9f, 0xb2,
	0xb6, 0xac, 0x88, 0x5e, 0x3d, 0x01, 0x89, 0xa8, 0x0d, 0xc2, 0x12, 0x89, 0xc8, 0x7b, 0xf8, 0x6d,
	0x0f, 0x4e, 0xb8, 0x39, 0x06, 0x36, 0xac, 0x67, 0xb6, 0xf1, 0x66, 0x28, 0xae, 0x5f, 0x60, 0xed,
	0xe0, 0xc5, 0x4d, 0x48, 0x22, 0x1d, 0xb2, 0xb5, 0x39, 0xb1, 0x2f, 0xf7, 0x5e, 0xe8, 0xd2, 0xbf,
	0xfa, 0xea, 0x54, 0x9d, 0x15, 0xfb, 0xa5, 0x7c, 0x8b, 0x70, 0xa9, 0x1a, 0x80, 0xc5, 0x21, 0x33,
	0x57, 0x6e, 0x67, 0x0e, 0xf0, 0x8d, 0x14, 0xc2, 0xa2, 0x9a, 0xa6, 0xe5, 0x4b, 0x04, 0xbe, 0x9e,
	0x20, 0x4d, 0xf9, 0x05, 0xe1, 0xd2, 0xbe, 0xdf, 0xce, 0xc3, 0x98, 0x01, 0x03, 0xfd, 0x8f, 0x30,
	0xc8, 0x43, 0x7c, 0xad, 0x17, 0xa1, 0x88, 0xce, 0x59, 0x5c, 0xa2, 0x3c, 0x8c, 0x3d, 0xbc, 0x78,
	0xea, 0x46, 0x78, 0xf1, 0xb6, 0x2d, 0x76, 0x24, 0xe6, 0x0f, 0x0b, 0x9f, 0x50, 0xa0, 0x98, 0xb8,
	0x54, 0x03, 0x17, 0x72, 0xea, 0x78, 0x30, 0xd1, 0xd7, 0xf7, 0xa3, 0xbe, 0xde, 0xc5, 0x77, 0xb2,
	0xfb, 0x5a, 0xa5, 0x1e, 0x87, 0x13, 0x2e, 0xba, 0x59, 0x79, 0x3e, 0x8b, 0x49, 0x72, 0xba, 0xc9,
	0x73, 0x84, 0x17, 0xd2, 0x07, 0x9f, 0x7c, 0x9c, 0xc7, 0x4d, 0xee, 0xc6, 0xca, 0x6b, 0x97, 0x71,
	0x15, 0x7b, 0xa6, 0x18, 0xe7, 0x7a, 0x3c, 0x41, 0xcf, 0x7e, 0xff, 0xe3, 0x1b, 0xe9, 0x43, 0x72,
	0x3f, 0xfc, 0xb3, 0x19, 0x08, 0xd1, 0x83, 0x8b, 0x9b, 0xb3, 0xfc, 0xf2, 0xad, 0x59, 0x9e, 0x38,
	0x32, 0xe4, 0x67, 0x84, 0xe7, 0xd3, 0xd6, 0x84, 0x7c, 0x94, 0x87, 0x2d, 0x67, 0xb1, 0xe4, 0xe9,
	0x66, 0x45, 0xf9, 0x74, 0x02, 0x7e, 0xd8, 0x83, 0x3c, 0xf0, 0x13, 0x07, 0x72, 0xf9, 0x8c, 0xfc,
	0x8d, 0x70, 0x21, 0x63, 0xf9, 0x48, 0x2e, 0xbb, 0xf9, 0x1b, 0x3b, 0x6d, 0x11, 0xde, 0xb9, 0x2e,
	0x0b, 0xe6, 0x57, 0x52, 0x96, 0x28, 0xaa, 0xf0, 0x73, 0xe5, 0x52, 0x0d, 0x5a, 0x4b, 0xdb, 0x4a,
	0xf2, 0xb5, 0x84, 0x0b, 0x19, 0xfb, 0x9c, 0x5f, 0x76, 0xfe, 0x11, 0x98, 0xb6, 0xec, 0xc1, 0xb9,
	0x7e, 0x23, 0xab, 0x5e, 0xb3, 0xf2, 0x59, 0x54, 0x6f, 0xda, 0x1b, 0x6e, 0xca, 0x2e, 0xa7, 0x73,
	0xf0, 0x13, 0xc2, 0x85, 0x8c, 0x5b, 0x90, 0xcf, 0x41, 0xfe, 0x01, 0x91, 0x17, 0x12, 0xf7, 0xc8,
	0x08, 0x9f, 0x67, 0x4a, 0xf5, 0x5c, 0x8f, 0x4e, 0x84, 0x98, 0xd7, 0xe5, 0x4b, 0xcd, 0xab, 0x7c,
	0xf2, 0xab, 0x7e, 0x33, 0xf3, 0x14, 0xfd, 0xa6, 0x3f, 0x39, 0xe4, 0xdc, 0x67, 0x6b, 0x9a, 0xd6,
	0xef, 0x27, 0xee, 0x94, 0xd5, 0xe3, 0x87, 0xe2, 0x81, 0xb9, 0xea, 0xbb, 0x16, 0xef, 0xd0, 0xa0,
	0xbb, 0xf2, 0x2a, 0xf3, 0x51, 0xaa, 0xf5, 0xef, 0x11, 0x2e, 0xb5, 0x68, 0x37, 0x87, 0x98, 0xf5,
	0x85, 0x04, 0x27, 0xf5, 0x90, 0x82, 0x3a, 0xfa, 0xa2, 0x16, 0x7b, 0xd9, 0xd4, 0xb5, 0x3c, 0x5b,
	0xa5, 0x81, 0xad, 0xd9, 0xe0, 0x45, 0x04, 0x69, 0xa3, 0xbc, 0x69, 0xcf, 0xdf, 0x4f, 0x46, 0x5f,
	0xff, 0x20, 0xf4, 0x9d, 0x24, 0xd5, 0x36, 0x7e, 0x94, 0xe4, 0x4d, 0x11, 0xae, 0x1a, 0x81, 0xa8,
	0x8d, 0x40, 0x3c, 0xae, 0x1c, 0x5c, 0x89, 0xa2, 0xde, 0xfb, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xd3,
	0xae, 0x29, 0xb4, 0x18, 0x0c, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SessionEntityTypesClient is the client API for SessionEntityTypes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SessionEntityTypesClient interface {
	// Returns the list of all session entity types in the specified session.
	//
	// This method doesn't work with Google Assistant integration.
	// Contact Dialogflow support if you need to use session entities
	// with Google Assistant integration.
	ListSessionEntityTypes(ctx context.Context, in *ListSessionEntityTypesRequest, opts ...grpc.CallOption) (*ListSessionEntityTypesResponse, error)
	// Retrieves the specified session entity type.
	//
	// This method doesn't work with Google Assistant integration.
	// Contact Dialogflow support if you need to use session entities
	// with Google Assistant integration.
	GetSessionEntityType(ctx context.Context, in *GetSessionEntityTypeRequest, opts ...grpc.CallOption) (*SessionEntityType, error)
	// Creates a session entity type.
	//
	// If the specified session entity type already exists, overrides the session
	// entity type.
	//
	// This method doesn't work with Google Assistant integration.
	// Contact Dialogflow support if you need to use session entities
	// with Google Assistant integration.
	CreateSessionEntityType(ctx context.Context, in *CreateSessionEntityTypeRequest, opts ...grpc.CallOption) (*SessionEntityType, error)
	// Updates the specified session entity type.
	//
	// This method doesn't work with Google Assistant integration.
	// Contact Dialogflow support if you need to use session entities
	// with Google Assistant integration.
	UpdateSessionEntityType(ctx context.Context, in *UpdateSessionEntityTypeRequest, opts ...grpc.CallOption) (*SessionEntityType, error)
	// Deletes the specified session entity type.
	//
	// This method doesn't work with Google Assistant integration.
	// Contact Dialogflow support if you need to use session entities
	// with Google Assistant integration.
	DeleteSessionEntityType(ctx context.Context, in *DeleteSessionEntityTypeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type sessionEntityTypesClient struct {
	cc *grpc.ClientConn
}

func NewSessionEntityTypesClient(cc *grpc.ClientConn) SessionEntityTypesClient {
	return &sessionEntityTypesClient{cc}
}

func (c *sessionEntityTypesClient) ListSessionEntityTypes(ctx context.Context, in *ListSessionEntityTypesRequest, opts ...grpc.CallOption) (*ListSessionEntityTypesResponse, error) {
	out := new(ListSessionEntityTypesResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.dialogflow.v2.SessionEntityTypes/ListSessionEntityTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionEntityTypesClient) GetSessionEntityType(ctx context.Context, in *GetSessionEntityTypeRequest, opts ...grpc.CallOption) (*SessionEntityType, error) {
	out := new(SessionEntityType)
	err := c.cc.Invoke(ctx, "/google.cloud.dialogflow.v2.SessionEntityTypes/GetSessionEntityType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionEntityTypesClient) CreateSessionEntityType(ctx context.Context, in *CreateSessionEntityTypeRequest, opts ...grpc.CallOption) (*SessionEntityType, error) {
	out := new(SessionEntityType)
	err := c.cc.Invoke(ctx, "/google.cloud.dialogflow.v2.SessionEntityTypes/CreateSessionEntityType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionEntityTypesClient) UpdateSessionEntityType(ctx context.Context, in *UpdateSessionEntityTypeRequest, opts ...grpc.CallOption) (*SessionEntityType, error) {
	out := new(SessionEntityType)
	err := c.cc.Invoke(ctx, "/google.cloud.dialogflow.v2.SessionEntityTypes/UpdateSessionEntityType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionEntityTypesClient) DeleteSessionEntityType(ctx context.Context, in *DeleteSessionEntityTypeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.cloud.dialogflow.v2.SessionEntityTypes/DeleteSessionEntityType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionEntityTypesServer is the server API for SessionEntityTypes service.
type SessionEntityTypesServer interface {
	// Returns the list of all session entity types in the specified session.
	//
	// This method doesn't work with Google Assistant integration.
	// Contact Dialogflow support if you need to use session entities
	// with Google Assistant integration.
	ListSessionEntityTypes(context.Context, *ListSessionEntityTypesRequest) (*ListSessionEntityTypesResponse, error)
	// Retrieves the specified session entity type.
	//
	// This method doesn't work with Google Assistant integration.
	// Contact Dialogflow support if you need to use session entities
	// with Google Assistant integration.
	GetSessionEntityType(context.Context, *GetSessionEntityTypeRequest) (*SessionEntityType, error)
	// Creates a session entity type.
	//
	// If the specified session entity type already exists, overrides the session
	// entity type.
	//
	// This method doesn't work with Google Assistant integration.
	// Contact Dialogflow support if you need to use session entities
	// with Google Assistant integration.
	CreateSessionEntityType(context.Context, *CreateSessionEntityTypeRequest) (*SessionEntityType, error)
	// Updates the specified session entity type.
	//
	// This method doesn't work with Google Assistant integration.
	// Contact Dialogflow support if you need to use session entities
	// with Google Assistant integration.
	UpdateSessionEntityType(context.Context, *UpdateSessionEntityTypeRequest) (*SessionEntityType, error)
	// Deletes the specified session entity type.
	//
	// This method doesn't work with Google Assistant integration.
	// Contact Dialogflow support if you need to use session entities
	// with Google Assistant integration.
	DeleteSessionEntityType(context.Context, *DeleteSessionEntityTypeRequest) (*empty.Empty, error)
}

// UnimplementedSessionEntityTypesServer can be embedded to have forward compatible implementations.
type UnimplementedSessionEntityTypesServer struct {
}

func (*UnimplementedSessionEntityTypesServer) ListSessionEntityTypes(ctx context.Context, req *ListSessionEntityTypesRequest) (*ListSessionEntityTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSessionEntityTypes not implemented")
}
func (*UnimplementedSessionEntityTypesServer) GetSessionEntityType(ctx context.Context, req *GetSessionEntityTypeRequest) (*SessionEntityType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSessionEntityType not implemented")
}
func (*UnimplementedSessionEntityTypesServer) CreateSessionEntityType(ctx context.Context, req *CreateSessionEntityTypeRequest) (*SessionEntityType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSessionEntityType not implemented")
}
func (*UnimplementedSessionEntityTypesServer) UpdateSessionEntityType(ctx context.Context, req *UpdateSessionEntityTypeRequest) (*SessionEntityType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSessionEntityType not implemented")
}
func (*UnimplementedSessionEntityTypesServer) DeleteSessionEntityType(ctx context.Context, req *DeleteSessionEntityTypeRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSessionEntityType not implemented")
}

func RegisterSessionEntityTypesServer(s *grpc.Server, srv SessionEntityTypesServer) {
	s.RegisterService(&_SessionEntityTypes_serviceDesc, srv)
}

func _SessionEntityTypes_ListSessionEntityTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSessionEntityTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionEntityTypesServer).ListSessionEntityTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2.SessionEntityTypes/ListSessionEntityTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionEntityTypesServer).ListSessionEntityTypes(ctx, req.(*ListSessionEntityTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionEntityTypes_GetSessionEntityType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionEntityTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionEntityTypesServer).GetSessionEntityType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2.SessionEntityTypes/GetSessionEntityType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionEntityTypesServer).GetSessionEntityType(ctx, req.(*GetSessionEntityTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionEntityTypes_CreateSessionEntityType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionEntityTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionEntityTypesServer).CreateSessionEntityType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2.SessionEntityTypes/CreateSessionEntityType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionEntityTypesServer).CreateSessionEntityType(ctx, req.(*CreateSessionEntityTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionEntityTypes_UpdateSessionEntityType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSessionEntityTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionEntityTypesServer).UpdateSessionEntityType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2.SessionEntityTypes/UpdateSessionEntityType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionEntityTypesServer).UpdateSessionEntityType(ctx, req.(*UpdateSessionEntityTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionEntityTypes_DeleteSessionEntityType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSessionEntityTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionEntityTypesServer).DeleteSessionEntityType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2.SessionEntityTypes/DeleteSessionEntityType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionEntityTypesServer).DeleteSessionEntityType(ctx, req.(*DeleteSessionEntityTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SessionEntityTypes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.dialogflow.v2.SessionEntityTypes",
	HandlerType: (*SessionEntityTypesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSessionEntityTypes",
			Handler:    _SessionEntityTypes_ListSessionEntityTypes_Handler,
		},
		{
			MethodName: "GetSessionEntityType",
			Handler:    _SessionEntityTypes_GetSessionEntityType_Handler,
		},
		{
			MethodName: "CreateSessionEntityType",
			Handler:    _SessionEntityTypes_CreateSessionEntityType_Handler,
		},
		{
			MethodName: "UpdateSessionEntityType",
			Handler:    _SessionEntityTypes_UpdateSessionEntityType_Handler,
		},
		{
			MethodName: "DeleteSessionEntityType",
			Handler:    _SessionEntityTypes_DeleteSessionEntityType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/dialogflow/v2/session_entity_type.proto",
}
