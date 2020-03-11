// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/keyword_plan_idea_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v3/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v3/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Request message for [KeywordIdeaService.GenerateKeywordIdeas][].
type GenerateKeywordIdeasRequest struct {
	// The ID of the customer with the recommendation.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The resource name of the language to target.
	// Required
	Language *wrappers.StringValue `protobuf:"bytes,7,opt,name=language,proto3" json:"language,omitempty"`
	// The resource names of the location to target.
	// Max 10
	GeoTargetConstants []*wrappers.StringValue `protobuf:"bytes,8,rep,name=geo_target_constants,json=geoTargetConstants,proto3" json:"geo_target_constants,omitempty"`
	// Targeting network.
	KeywordPlanNetwork enums.KeywordPlanNetworkEnum_KeywordPlanNetwork `protobuf:"varint,9,opt,name=keyword_plan_network,json=keywordPlanNetwork,proto3,enum=google.ads.googleads.v3.enums.KeywordPlanNetworkEnum_KeywordPlanNetwork" json:"keyword_plan_network,omitempty"`
	// The type of seed to generate keyword ideas.
	//
	// Types that are valid to be assigned to Seed:
	//	*GenerateKeywordIdeasRequest_KeywordAndUrlSeed
	//	*GenerateKeywordIdeasRequest_KeywordSeed
	//	*GenerateKeywordIdeasRequest_UrlSeed
	Seed                 isGenerateKeywordIdeasRequest_Seed `protobuf_oneof:"seed"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *GenerateKeywordIdeasRequest) Reset()         { *m = GenerateKeywordIdeasRequest{} }
func (m *GenerateKeywordIdeasRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateKeywordIdeasRequest) ProtoMessage()    {}
func (*GenerateKeywordIdeasRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_26269aab04477620, []int{0}
}

func (m *GenerateKeywordIdeasRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateKeywordIdeasRequest.Unmarshal(m, b)
}
func (m *GenerateKeywordIdeasRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateKeywordIdeasRequest.Marshal(b, m, deterministic)
}
func (m *GenerateKeywordIdeasRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateKeywordIdeasRequest.Merge(m, src)
}
func (m *GenerateKeywordIdeasRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateKeywordIdeasRequest.Size(m)
}
func (m *GenerateKeywordIdeasRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateKeywordIdeasRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateKeywordIdeasRequest proto.InternalMessageInfo

func (m *GenerateKeywordIdeasRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *GenerateKeywordIdeasRequest) GetLanguage() *wrappers.StringValue {
	if m != nil {
		return m.Language
	}
	return nil
}

func (m *GenerateKeywordIdeasRequest) GetGeoTargetConstants() []*wrappers.StringValue {
	if m != nil {
		return m.GeoTargetConstants
	}
	return nil
}

func (m *GenerateKeywordIdeasRequest) GetKeywordPlanNetwork() enums.KeywordPlanNetworkEnum_KeywordPlanNetwork {
	if m != nil {
		return m.KeywordPlanNetwork
	}
	return enums.KeywordPlanNetworkEnum_UNSPECIFIED
}

type isGenerateKeywordIdeasRequest_Seed interface {
	isGenerateKeywordIdeasRequest_Seed()
}

type GenerateKeywordIdeasRequest_KeywordAndUrlSeed struct {
	KeywordAndUrlSeed *KeywordAndUrlSeed `protobuf:"bytes,2,opt,name=keyword_and_url_seed,json=keywordAndUrlSeed,proto3,oneof"`
}

type GenerateKeywordIdeasRequest_KeywordSeed struct {
	KeywordSeed *KeywordSeed `protobuf:"bytes,3,opt,name=keyword_seed,json=keywordSeed,proto3,oneof"`
}

type GenerateKeywordIdeasRequest_UrlSeed struct {
	UrlSeed *UrlSeed `protobuf:"bytes,5,opt,name=url_seed,json=urlSeed,proto3,oneof"`
}

func (*GenerateKeywordIdeasRequest_KeywordAndUrlSeed) isGenerateKeywordIdeasRequest_Seed() {}

func (*GenerateKeywordIdeasRequest_KeywordSeed) isGenerateKeywordIdeasRequest_Seed() {}

func (*GenerateKeywordIdeasRequest_UrlSeed) isGenerateKeywordIdeasRequest_Seed() {}

func (m *GenerateKeywordIdeasRequest) GetSeed() isGenerateKeywordIdeasRequest_Seed {
	if m != nil {
		return m.Seed
	}
	return nil
}

func (m *GenerateKeywordIdeasRequest) GetKeywordAndUrlSeed() *KeywordAndUrlSeed {
	if x, ok := m.GetSeed().(*GenerateKeywordIdeasRequest_KeywordAndUrlSeed); ok {
		return x.KeywordAndUrlSeed
	}
	return nil
}

func (m *GenerateKeywordIdeasRequest) GetKeywordSeed() *KeywordSeed {
	if x, ok := m.GetSeed().(*GenerateKeywordIdeasRequest_KeywordSeed); ok {
		return x.KeywordSeed
	}
	return nil
}

func (m *GenerateKeywordIdeasRequest) GetUrlSeed() *UrlSeed {
	if x, ok := m.GetSeed().(*GenerateKeywordIdeasRequest_UrlSeed); ok {
		return x.UrlSeed
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GenerateKeywordIdeasRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GenerateKeywordIdeasRequest_KeywordAndUrlSeed)(nil),
		(*GenerateKeywordIdeasRequest_KeywordSeed)(nil),
		(*GenerateKeywordIdeasRequest_UrlSeed)(nil),
	}
}

// Keyword And Url Seed
type KeywordAndUrlSeed struct {
	// The URL to crawl in order to generate keyword ideas.
	Url *wrappers.StringValue `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// Requires at least one keyword.
	Keywords             []*wrappers.StringValue `protobuf:"bytes,2,rep,name=keywords,proto3" json:"keywords,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *KeywordAndUrlSeed) Reset()         { *m = KeywordAndUrlSeed{} }
func (m *KeywordAndUrlSeed) String() string { return proto.CompactTextString(m) }
func (*KeywordAndUrlSeed) ProtoMessage()    {}
func (*KeywordAndUrlSeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_26269aab04477620, []int{1}
}

func (m *KeywordAndUrlSeed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordAndUrlSeed.Unmarshal(m, b)
}
func (m *KeywordAndUrlSeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordAndUrlSeed.Marshal(b, m, deterministic)
}
func (m *KeywordAndUrlSeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordAndUrlSeed.Merge(m, src)
}
func (m *KeywordAndUrlSeed) XXX_Size() int {
	return xxx_messageInfo_KeywordAndUrlSeed.Size(m)
}
func (m *KeywordAndUrlSeed) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordAndUrlSeed.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordAndUrlSeed proto.InternalMessageInfo

func (m *KeywordAndUrlSeed) GetUrl() *wrappers.StringValue {
	if m != nil {
		return m.Url
	}
	return nil
}

func (m *KeywordAndUrlSeed) GetKeywords() []*wrappers.StringValue {
	if m != nil {
		return m.Keywords
	}
	return nil
}

// Keyword Seed
type KeywordSeed struct {
	// Requires at least one keyword.
	Keywords             []*wrappers.StringValue `protobuf:"bytes,1,rep,name=keywords,proto3" json:"keywords,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *KeywordSeed) Reset()         { *m = KeywordSeed{} }
func (m *KeywordSeed) String() string { return proto.CompactTextString(m) }
func (*KeywordSeed) ProtoMessage()    {}
func (*KeywordSeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_26269aab04477620, []int{2}
}

func (m *KeywordSeed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordSeed.Unmarshal(m, b)
}
func (m *KeywordSeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordSeed.Marshal(b, m, deterministic)
}
func (m *KeywordSeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordSeed.Merge(m, src)
}
func (m *KeywordSeed) XXX_Size() int {
	return xxx_messageInfo_KeywordSeed.Size(m)
}
func (m *KeywordSeed) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordSeed.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordSeed proto.InternalMessageInfo

func (m *KeywordSeed) GetKeywords() []*wrappers.StringValue {
	if m != nil {
		return m.Keywords
	}
	return nil
}

// Url Seed
type UrlSeed struct {
	// The URL to crawl in order to generate keyword ideas.
	Url                  *wrappers.StringValue `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UrlSeed) Reset()         { *m = UrlSeed{} }
func (m *UrlSeed) String() string { return proto.CompactTextString(m) }
func (*UrlSeed) ProtoMessage()    {}
func (*UrlSeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_26269aab04477620, []int{3}
}

func (m *UrlSeed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UrlSeed.Unmarshal(m, b)
}
func (m *UrlSeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UrlSeed.Marshal(b, m, deterministic)
}
func (m *UrlSeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UrlSeed.Merge(m, src)
}
func (m *UrlSeed) XXX_Size() int {
	return xxx_messageInfo_UrlSeed.Size(m)
}
func (m *UrlSeed) XXX_DiscardUnknown() {
	xxx_messageInfo_UrlSeed.DiscardUnknown(m)
}

var xxx_messageInfo_UrlSeed proto.InternalMessageInfo

func (m *UrlSeed) GetUrl() *wrappers.StringValue {
	if m != nil {
		return m.Url
	}
	return nil
}

// Response message for [KeywordIdeaService.GenerateKeywordIdeas][].
type GenerateKeywordIdeaResponse struct {
	// Results of generating keyword ideas.
	Results              []*GenerateKeywordIdeaResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *GenerateKeywordIdeaResponse) Reset()         { *m = GenerateKeywordIdeaResponse{} }
func (m *GenerateKeywordIdeaResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateKeywordIdeaResponse) ProtoMessage()    {}
func (*GenerateKeywordIdeaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_26269aab04477620, []int{4}
}

func (m *GenerateKeywordIdeaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateKeywordIdeaResponse.Unmarshal(m, b)
}
func (m *GenerateKeywordIdeaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateKeywordIdeaResponse.Marshal(b, m, deterministic)
}
func (m *GenerateKeywordIdeaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateKeywordIdeaResponse.Merge(m, src)
}
func (m *GenerateKeywordIdeaResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateKeywordIdeaResponse.Size(m)
}
func (m *GenerateKeywordIdeaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateKeywordIdeaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateKeywordIdeaResponse proto.InternalMessageInfo

func (m *GenerateKeywordIdeaResponse) GetResults() []*GenerateKeywordIdeaResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result of generating keyword ideas.
type GenerateKeywordIdeaResult struct {
	// Text of the keyword idea.
	// As in Keyword Plan historical metrics, this text may not be an actual
	// keyword, but the canonical form of multiple keywords.
	// See KeywordPlanKeywordHistoricalMetrics message in KeywordPlanService.
	Text *wrappers.StringValue `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	// The historical metrics for the keyword
	KeywordIdeaMetrics   *common.KeywordPlanHistoricalMetrics `protobuf:"bytes,3,opt,name=keyword_idea_metrics,json=keywordIdeaMetrics,proto3" json:"keyword_idea_metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *GenerateKeywordIdeaResult) Reset()         { *m = GenerateKeywordIdeaResult{} }
func (m *GenerateKeywordIdeaResult) String() string { return proto.CompactTextString(m) }
func (*GenerateKeywordIdeaResult) ProtoMessage()    {}
func (*GenerateKeywordIdeaResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_26269aab04477620, []int{5}
}

func (m *GenerateKeywordIdeaResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateKeywordIdeaResult.Unmarshal(m, b)
}
func (m *GenerateKeywordIdeaResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateKeywordIdeaResult.Marshal(b, m, deterministic)
}
func (m *GenerateKeywordIdeaResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateKeywordIdeaResult.Merge(m, src)
}
func (m *GenerateKeywordIdeaResult) XXX_Size() int {
	return xxx_messageInfo_GenerateKeywordIdeaResult.Size(m)
}
func (m *GenerateKeywordIdeaResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateKeywordIdeaResult.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateKeywordIdeaResult proto.InternalMessageInfo

func (m *GenerateKeywordIdeaResult) GetText() *wrappers.StringValue {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *GenerateKeywordIdeaResult) GetKeywordIdeaMetrics() *common.KeywordPlanHistoricalMetrics {
	if m != nil {
		return m.KeywordIdeaMetrics
	}
	return nil
}

func init() {
	proto.RegisterType((*GenerateKeywordIdeasRequest)(nil), "google.ads.googleads.v3.services.GenerateKeywordIdeasRequest")
	proto.RegisterType((*KeywordAndUrlSeed)(nil), "google.ads.googleads.v3.services.KeywordAndUrlSeed")
	proto.RegisterType((*KeywordSeed)(nil), "google.ads.googleads.v3.services.KeywordSeed")
	proto.RegisterType((*UrlSeed)(nil), "google.ads.googleads.v3.services.UrlSeed")
	proto.RegisterType((*GenerateKeywordIdeaResponse)(nil), "google.ads.googleads.v3.services.GenerateKeywordIdeaResponse")
	proto.RegisterType((*GenerateKeywordIdeaResult)(nil), "google.ads.googleads.v3.services.GenerateKeywordIdeaResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/keyword_plan_idea_service.proto", fileDescriptor_26269aab04477620)
}

var fileDescriptor_26269aab04477620 = []byte{
	// 748 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4f, 0x4f, 0xdb, 0x48,
	0x14, 0x5f, 0x3b, 0x2c, 0x81, 0xc9, 0x6a, 0x25, 0x46, 0x68, 0x37, 0x0b, 0x68, 0x37, 0x8a, 0x38,
	0xb0, 0x48, 0xb5, 0xab, 0xe4, 0x50, 0x6a, 0x1a, 0xa9, 0xa1, 0x6a, 0x13, 0x54, 0x15, 0x21, 0x53,
	0x72, 0xa8, 0x22, 0x59, 0x83, 0xfd, 0xb0, 0xac, 0x38, 0x33, 0xe9, 0xcc, 0x18, 0xfa, 0x47, 0x5c,
	0xb8, 0xf6, 0xd8, 0x6f, 0xd0, 0x63, 0xbf, 0x43, 0xbf, 0x00, 0xa7, 0x4a, 0x7c, 0x85, 0x9e, 0xfa,
	0x29, 0x2a, 0xdb, 0x33, 0x49, 0x80, 0xa4, 0xa1, 0xdc, 0x9e, 0xdf, 0xfb, 0xbd, 0xdf, 0xfb, 0xcd,
	0x9b, 0x37, 0xcf, 0xe8, 0x71, 0xc8, 0x58, 0x18, 0x83, 0x4d, 0x02, 0x61, 0xe7, 0x66, 0x6a, 0x9d,
	0xd4, 0x6d, 0x01, 0xfc, 0x24, 0xf2, 0x41, 0xd8, 0x3d, 0x78, 0x7b, 0xca, 0x78, 0xe0, 0x0d, 0x62,
	0x42, 0xbd, 0x28, 0x00, 0xe2, 0xa9, 0x90, 0x35, 0xe0, 0x4c, 0x32, 0x5c, 0xc9, 0xd3, 0x2c, 0x12,
	0x08, 0x6b, 0xc8, 0x60, 0x9d, 0xd4, 0x2d, 0xcd, 0xb0, 0xb2, 0x35, 0xad, 0x86, 0xcf, 0xfa, 0x7d,
	0x46, 0xaf, 0x56, 0xc8, 0x7d, 0x39, 0xf7, 0xf4, 0x4c, 0xa0, 0x49, 0xff, 0x9a, 0x34, 0x0a, 0xf2,
	0x94, 0xf1, 0x9e, 0xca, 0x5c, 0xd3, 0x99, 0x83, 0xc8, 0x26, 0x94, 0x32, 0x49, 0x64, 0xc4, 0xa8,
	0x50, 0xd1, 0x7f, 0x55, 0x34, 0xfb, 0x3a, 0x4a, 0x8e, 0xed, 0x53, 0x4e, 0x06, 0x03, 0xe0, 0x3a,
	0xfe, 0xf7, 0x58, 0xb6, 0x1f, 0x47, 0x40, 0x65, 0x1e, 0xa8, 0x7e, 0x9d, 0x43, 0xab, 0x2d, 0xa0,
	0xc0, 0x89, 0x84, 0xe7, 0x79, 0xf5, 0xdd, 0x00, 0x88, 0x70, 0xe1, 0x75, 0x02, 0x42, 0xe2, 0xff,
	0x50, 0xc9, 0x4f, 0x84, 0x64, 0x7d, 0xe0, 0x5e, 0x14, 0x94, 0x8d, 0x8a, 0xb1, 0xb1, 0xe8, 0x22,
	0xed, 0xda, 0x0d, 0xf0, 0x16, 0x5a, 0x88, 0x09, 0x0d, 0x13, 0x12, 0x42, 0xb9, 0x58, 0x31, 0x36,
	0x4a, 0xb5, 0x35, 0xd5, 0x35, 0x4b, 0x8b, 0xb1, 0x0e, 0x24, 0x8f, 0x68, 0xd8, 0x21, 0x71, 0x02,
	0xee, 0x10, 0x8d, 0xf7, 0xd0, 0x72, 0x08, 0xcc, 0x93, 0x84, 0x87, 0x20, 0x3d, 0x9f, 0x51, 0x21,
	0x09, 0x95, 0xa2, 0xbc, 0x50, 0x29, 0xcc, 0x64, 0xc1, 0x21, 0xb0, 0x97, 0x59, 0xe2, 0x13, 0x9d,
	0x87, 0xdf, 0xa1, 0xe5, 0x49, 0xfd, 0x2b, 0x2f, 0x56, 0x8c, 0x8d, 0x3f, 0x6b, 0x6d, 0x6b, 0xda,
	0xb5, 0x66, 0xad, 0xb7, 0xd4, 0xe1, 0xf7, 0x63, 0x42, 0xf7, 0xf2, 0xc4, 0xa7, 0x34, 0xe9, 0x4f,
	0x70, 0xbb, 0xb8, 0x77, 0xc3, 0x87, 0x8f, 0x47, 0xb5, 0x09, 0x0d, 0xbc, 0x84, 0xc7, 0x9e, 0x00,
	0x08, 0xca, 0x66, 0xd6, 0x91, 0xba, 0x35, 0x6b, 0xa4, 0x74, 0x9d, 0x26, 0x0d, 0x0e, 0x79, 0x7c,
	0x00, 0x10, 0xb4, 0x7f, 0x73, 0x97, 0x7a, 0xd7, 0x9d, 0xd8, 0x45, 0x7f, 0xe8, 0x3a, 0x19, 0x7f,
	0x21, 0xe3, 0xbf, 0x77, 0x6b, 0x7e, 0xc5, 0x5c, 0xea, 0x8d, 0x3e, 0xf1, 0x33, 0xb4, 0x30, 0xd4,
	0xfb, 0x7b, 0xc6, 0xf7, 0xff, 0x6c, 0xbe, 0x91, 0xca, 0x62, 0x92, 0x9b, 0x3b, 0xf3, 0x68, 0x2e,
	0xe5, 0xa8, 0x9e, 0xa1, 0xa5, 0x1b, 0xa7, 0xc1, 0x16, 0x2a, 0x24, 0x3c, 0xce, 0xe6, 0x67, 0xd6,
	0xdd, 0xa6, 0xc0, 0x74, 0xac, 0x94, 0x46, 0x51, 0x36, 0x6f, 0x31, 0x10, 0x43, 0x74, 0xb5, 0x85,
	0x4a, 0x63, 0x87, 0xbd, 0x42, 0x64, 0xfc, 0x12, 0xd1, 0x43, 0x54, 0xbc, 0xa3, 0xfa, 0xaa, 0x9c,
	0xf8, 0xa8, 0x5c, 0x10, 0x03, 0x46, 0x05, 0xe0, 0x43, 0x54, 0xe4, 0x20, 0x92, 0x58, 0x6a, 0x49,
	0xdb, 0xb3, 0x1b, 0x3e, 0x99, 0x2f, 0x89, 0xa5, 0xab, 0xb9, 0xaa, 0x5f, 0x0c, 0xf4, 0xcf, 0x54,
	0x18, 0xbe, 0x8f, 0xe6, 0x24, 0xbc, 0x91, 0x6a, 0x24, 0x7f, 0x7e, 0x88, 0x0c, 0x89, 0xe9, 0x68,
	0xa8, 0xb3, 0x35, 0xd9, 0x07, 0xc9, 0x23, 0x5f, 0xa8, 0xa1, 0x7b, 0x34, 0x55, 0xb3, 0xda, 0x78,
	0x63, 0x4f, 0xa7, 0x1d, 0x09, 0xc9, 0x78, 0xe4, 0x93, 0xf8, 0x45, 0xce, 0x31, 0x7c, 0x44, 0xa9,
	0x40, 0xe5, 0xab, 0x7d, 0x30, 0xd1, 0x5f, 0x63, 0x49, 0x69, 0xe8, 0x20, 0x3f, 0x3e, 0xbe, 0x34,
	0xd0, 0xf2, 0xa4, 0x35, 0x85, 0x1b, 0x77, 0xea, 0x9c, 0x5e, 0x6f, 0x2b, 0x8d, 0xbb, 0x36, 0x3e,
	0xbb, 0xc8, 0x6a, 0xe3, 0xfc, 0xf2, 0xdb, 0x47, 0xf3, 0x41, 0xb5, 0x96, 0x2d, 0x7f, 0xb5, 0x14,
	0x85, 0xfd, 0x7e, 0x6c, 0x65, 0x36, 0x36, 0xcf, 0x9c, 0x70, 0x82, 0x02, 0xc7, 0xd8, 0x5c, 0x59,
	0xbd, 0x68, 0x96, 0x47, 0x45, 0x95, 0x35, 0x88, 0x44, 0xda, 0xc1, 0x9d, 0x73, 0x13, 0xad, 0xfb,
	0xac, 0x3f, 0x53, 0xe0, 0xce, 0xea, 0xe4, 0x9e, 0xed, 0xa7, 0x17, 0xbb, 0x6f, 0xbc, 0x6a, 0x2b,
	0x82, 0x90, 0xa5, 0x9b, 0xd7, 0x62, 0x3c, 0xb4, 0x43, 0xa0, 0xd9, 0xb5, 0xdb, 0xa3, 0x92, 0xd3,
	0xff, 0x97, 0xdb, 0xda, 0xf8, 0x64, 0x16, 0x5a, 0xcd, 0xe6, 0x67, 0xb3, 0xd2, 0xca, 0x09, 0x9b,
	0x81, 0xb0, 0x72, 0x33, 0xb5, 0x3a, 0x75, 0x4b, 0x15, 0x16, 0x17, 0x1a, 0xd2, 0x6d, 0x06, 0xa2,
	0x3b, 0x84, 0x74, 0x3b, 0xf5, 0xae, 0x86, 0x7c, 0x37, 0xd7, 0x73, 0xbf, 0xe3, 0x34, 0x03, 0xe1,
	0x38, 0x43, 0x90, 0xe3, 0x74, 0xea, 0x8e, 0xa3, 0x61, 0x47, 0xf3, 0x99, 0xce, 0xfa, 0x8f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x97, 0x74, 0x32, 0x68, 0xd6, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KeywordPlanIdeaServiceClient is the client API for KeywordPlanIdeaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeywordPlanIdeaServiceClient interface {
	// Returns a list of keyword ideas.
	GenerateKeywordIdeas(ctx context.Context, in *GenerateKeywordIdeasRequest, opts ...grpc.CallOption) (*GenerateKeywordIdeaResponse, error)
}

type keywordPlanIdeaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeywordPlanIdeaServiceClient(cc grpc.ClientConnInterface) KeywordPlanIdeaServiceClient {
	return &keywordPlanIdeaServiceClient{cc}
}

func (c *keywordPlanIdeaServiceClient) GenerateKeywordIdeas(ctx context.Context, in *GenerateKeywordIdeasRequest, opts ...grpc.CallOption) (*GenerateKeywordIdeaResponse, error) {
	out := new(GenerateKeywordIdeaResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.KeywordPlanIdeaService/GenerateKeywordIdeas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordPlanIdeaServiceServer is the server API for KeywordPlanIdeaService service.
type KeywordPlanIdeaServiceServer interface {
	// Returns a list of keyword ideas.
	GenerateKeywordIdeas(context.Context, *GenerateKeywordIdeasRequest) (*GenerateKeywordIdeaResponse, error)
}

// UnimplementedKeywordPlanIdeaServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKeywordPlanIdeaServiceServer struct {
}

func (*UnimplementedKeywordPlanIdeaServiceServer) GenerateKeywordIdeas(ctx context.Context, req *GenerateKeywordIdeasRequest) (*GenerateKeywordIdeaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateKeywordIdeas not implemented")
}

func RegisterKeywordPlanIdeaServiceServer(s *grpc.Server, srv KeywordPlanIdeaServiceServer) {
	s.RegisterService(&_KeywordPlanIdeaService_serviceDesc, srv)
}

func _KeywordPlanIdeaService_GenerateKeywordIdeas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateKeywordIdeasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanIdeaServiceServer).GenerateKeywordIdeas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.KeywordPlanIdeaService/GenerateKeywordIdeas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanIdeaServiceServer).GenerateKeywordIdeas(ctx, req.(*GenerateKeywordIdeasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeywordPlanIdeaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.KeywordPlanIdeaService",
	HandlerType: (*KeywordPlanIdeaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateKeywordIdeas",
			Handler:    _KeywordPlanIdeaService_GenerateKeywordIdeas_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/keyword_plan_idea_service.proto",
}