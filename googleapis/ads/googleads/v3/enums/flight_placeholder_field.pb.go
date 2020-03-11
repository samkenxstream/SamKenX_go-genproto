// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/flight_placeholder_field.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Possible values for Flight placeholder fields.
type FlightPlaceholderFieldEnum_FlightPlaceholderField int32

const (
	// Not specified.
	FlightPlaceholderFieldEnum_UNSPECIFIED FlightPlaceholderFieldEnum_FlightPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	FlightPlaceholderFieldEnum_UNKNOWN FlightPlaceholderFieldEnum_FlightPlaceholderField = 1
	// Data Type: STRING. Required. Destination id. Example: PAR, LON.
	// For feed items that only have destination id, destination id must be a
	// unique key. For feed items that have both destination id and origin id,
	// then the combination must be a unique key.
	FlightPlaceholderFieldEnum_DESTINATION_ID FlightPlaceholderFieldEnum_FlightPlaceholderField = 2
	// Data Type: STRING. Origin id. Example: PAR, LON.
	// Optional. Combination of destination id and origin id must be unique per
	// offer.
	FlightPlaceholderFieldEnum_ORIGIN_ID FlightPlaceholderFieldEnum_FlightPlaceholderField = 3
	// Data Type: STRING. Required. Main headline with product name to be shown
	// in dynamic ad.
	FlightPlaceholderFieldEnum_FLIGHT_DESCRIPTION FlightPlaceholderFieldEnum_FlightPlaceholderField = 4
	// Data Type: STRING. Shorter names are recommended.
	FlightPlaceholderFieldEnum_ORIGIN_NAME FlightPlaceholderFieldEnum_FlightPlaceholderField = 5
	// Data Type: STRING. Shorter names are recommended.
	FlightPlaceholderFieldEnum_DESTINATION_NAME FlightPlaceholderFieldEnum_FlightPlaceholderField = 6
	// Data Type: STRING. Price to be shown in the ad.
	// Example: "100.00 USD"
	FlightPlaceholderFieldEnum_FLIGHT_PRICE FlightPlaceholderFieldEnum_FlightPlaceholderField = 7
	// Data Type: STRING. Formatted price to be shown in the ad.
	// Example: "Starting at $100.00 USD", "$80 - $100"
	FlightPlaceholderFieldEnum_FORMATTED_PRICE FlightPlaceholderFieldEnum_FlightPlaceholderField = 8
	// Data Type: STRING. Sale price to be shown in the ad.
	// Example: "80.00 USD"
	FlightPlaceholderFieldEnum_FLIGHT_SALE_PRICE FlightPlaceholderFieldEnum_FlightPlaceholderField = 9
	// Data Type: STRING. Formatted sale price to be shown in the ad.
	// Example: "On sale for $80.00", "$60 - $80"
	FlightPlaceholderFieldEnum_FORMATTED_SALE_PRICE FlightPlaceholderFieldEnum_FlightPlaceholderField = 10
	// Data Type: URL. Image to be displayed in the ad.
	FlightPlaceholderFieldEnum_IMAGE_URL FlightPlaceholderFieldEnum_FlightPlaceholderField = 11
	// Data Type: URL_LIST. Required. Final URLs for the ad when using Upgraded
	// URLs. User will be redirected to these URLs when they click on an ad, or
	// when they click on a specific flight for ads that show multiple
	// flights.
	FlightPlaceholderFieldEnum_FINAL_URLS FlightPlaceholderFieldEnum_FlightPlaceholderField = 12
	// Data Type: URL_LIST. Final mobile URLs for the ad when using Upgraded
	// URLs.
	FlightPlaceholderFieldEnum_FINAL_MOBILE_URLS FlightPlaceholderFieldEnum_FlightPlaceholderField = 13
	// Data Type: URL. Tracking template for the ad when using Upgraded URLs.
	FlightPlaceholderFieldEnum_TRACKING_URL FlightPlaceholderFieldEnum_FlightPlaceholderField = 14
	// Data Type: STRING. Android app link. Must be formatted as:
	// android-app://{package_id}/{scheme}/{host_path}.
	// The components are defined as follows:
	// package_id: app ID as specified in Google Play.
	// scheme: the scheme to pass to the application. Can be HTTP, or a custom
	//   scheme.
	// host_path: identifies the specific content within your application.
	FlightPlaceholderFieldEnum_ANDROID_APP_LINK FlightPlaceholderFieldEnum_FlightPlaceholderField = 15
	// Data Type: STRING_LIST. List of recommended destination IDs to show
	// together with this item.
	FlightPlaceholderFieldEnum_SIMILAR_DESTINATION_IDS FlightPlaceholderFieldEnum_FlightPlaceholderField = 16
	// Data Type: STRING. iOS app link.
	FlightPlaceholderFieldEnum_IOS_APP_LINK FlightPlaceholderFieldEnum_FlightPlaceholderField = 17
	// Data Type: INT64. iOS app store ID.
	FlightPlaceholderFieldEnum_IOS_APP_STORE_ID FlightPlaceholderFieldEnum_FlightPlaceholderField = 18
)

var FlightPlaceholderFieldEnum_FlightPlaceholderField_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "DESTINATION_ID",
	3:  "ORIGIN_ID",
	4:  "FLIGHT_DESCRIPTION",
	5:  "ORIGIN_NAME",
	6:  "DESTINATION_NAME",
	7:  "FLIGHT_PRICE",
	8:  "FORMATTED_PRICE",
	9:  "FLIGHT_SALE_PRICE",
	10: "FORMATTED_SALE_PRICE",
	11: "IMAGE_URL",
	12: "FINAL_URLS",
	13: "FINAL_MOBILE_URLS",
	14: "TRACKING_URL",
	15: "ANDROID_APP_LINK",
	16: "SIMILAR_DESTINATION_IDS",
	17: "IOS_APP_LINK",
	18: "IOS_APP_STORE_ID",
}

var FlightPlaceholderFieldEnum_FlightPlaceholderField_value = map[string]int32{
	"UNSPECIFIED":             0,
	"UNKNOWN":                 1,
	"DESTINATION_ID":          2,
	"ORIGIN_ID":               3,
	"FLIGHT_DESCRIPTION":      4,
	"ORIGIN_NAME":             5,
	"DESTINATION_NAME":        6,
	"FLIGHT_PRICE":            7,
	"FORMATTED_PRICE":         8,
	"FLIGHT_SALE_PRICE":       9,
	"FORMATTED_SALE_PRICE":    10,
	"IMAGE_URL":               11,
	"FINAL_URLS":              12,
	"FINAL_MOBILE_URLS":       13,
	"TRACKING_URL":            14,
	"ANDROID_APP_LINK":        15,
	"SIMILAR_DESTINATION_IDS": 16,
	"IOS_APP_LINK":            17,
	"IOS_APP_STORE_ID":        18,
}

func (x FlightPlaceholderFieldEnum_FlightPlaceholderField) String() string {
	return proto.EnumName(FlightPlaceholderFieldEnum_FlightPlaceholderField_name, int32(x))
}

func (FlightPlaceholderFieldEnum_FlightPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_05562ffffd9d8d69, []int{0, 0}
}

// Values for Flight placeholder fields.
// For more information about dynamic remarketing feeds, see
// https://support.google.com/google-ads/answer/6053288.
type FlightPlaceholderFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlightPlaceholderFieldEnum) Reset()         { *m = FlightPlaceholderFieldEnum{} }
func (m *FlightPlaceholderFieldEnum) String() string { return proto.CompactTextString(m) }
func (*FlightPlaceholderFieldEnum) ProtoMessage()    {}
func (*FlightPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_05562ffffd9d8d69, []int{0}
}

func (m *FlightPlaceholderFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlightPlaceholderFieldEnum.Unmarshal(m, b)
}
func (m *FlightPlaceholderFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlightPlaceholderFieldEnum.Marshal(b, m, deterministic)
}
func (m *FlightPlaceholderFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlightPlaceholderFieldEnum.Merge(m, src)
}
func (m *FlightPlaceholderFieldEnum) XXX_Size() int {
	return xxx_messageInfo_FlightPlaceholderFieldEnum.Size(m)
}
func (m *FlightPlaceholderFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FlightPlaceholderFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FlightPlaceholderFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.FlightPlaceholderFieldEnum_FlightPlaceholderField", FlightPlaceholderFieldEnum_FlightPlaceholderField_name, FlightPlaceholderFieldEnum_FlightPlaceholderField_value)
	proto.RegisterType((*FlightPlaceholderFieldEnum)(nil), "google.ads.googleads.v3.enums.FlightPlaceholderFieldEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/flight_placeholder_field.proto", fileDescriptor_05562ffffd9d8d69)
}

var fileDescriptor_05562ffffd9d8d69 = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x6e, 0xd3, 0x40,
	0x14, 0x26, 0x09, 0xb4, 0x74, 0xd2, 0x26, 0xd3, 0xa1, 0x14, 0x54, 0xda, 0x45, 0x7b, 0x00, 0x7b,
	0xe1, 0x9d, 0x61, 0x33, 0x89, 0xc7, 0x66, 0x14, 0x7b, 0x6c, 0xd9, 0x4e, 0x90, 0x50, 0x24, 0xcb,
	0xd4, 0xae, 0x6b, 0xc9, 0xf1, 0x44, 0x99, 0xb4, 0x17, 0xe1, 0x06, 0x2c, 0x59, 0x72, 0x0c, 0x8e,
	0xc2, 0x8e, 0x1b, 0xa0, 0x19, 0x3b, 0x3f, 0x20, 0x60, 0x33, 0x7a, 0xef, 0xfb, 0xd3, 0xd3, 0xbc,
	0x07, 0xde, 0x15, 0x9c, 0x17, 0x55, 0xae, 0xa7, 0x99, 0xd0, 0x9b, 0x52, 0x56, 0x8f, 0x86, 0x9e,
	0xd7, 0x0f, 0x0b, 0xa1, 0xdf, 0x55, 0x65, 0x71, 0xbf, 0x4e, 0x96, 0x55, 0x7a, 0x9b, 0xdf, 0xf3,
	0x2a, 0xcb, 0x57, 0xc9, 0x5d, 0x99, 0x57, 0x99, 0xb6, 0x5c, 0xf1, 0x35, 0x47, 0x57, 0x8d, 0x45,
	0x4b, 0x33, 0xa1, 0x6d, 0xdd, 0xda, 0xa3, 0xa1, 0x29, 0xf7, 0xc5, 0xe5, 0x26, 0x7c, 0x59, 0xea,
	0x69, 0x5d, 0xf3, 0x75, 0xba, 0x2e, 0x79, 0x2d, 0x1a, 0xf3, 0xcd, 0xb7, 0x1e, 0xb8, 0xb0, 0x55,
	0x7e, 0xb0, 0x8b, 0xb7, 0x65, 0x3a, 0xa9, 0x1f, 0x16, 0x37, 0x9f, 0x7b, 0xe0, 0xfc, 0xef, 0x34,
	0x1a, 0x82, 0xfe, 0x94, 0x45, 0x01, 0x19, 0x53, 0x9b, 0x12, 0x0b, 0x3e, 0x41, 0x7d, 0x70, 0x38,
	0x65, 0x13, 0xe6, 0x7f, 0x60, 0xb0, 0x83, 0x10, 0x18, 0x58, 0x24, 0x8a, 0x29, 0xc3, 0x31, 0xf5,
	0x59, 0x42, 0x2d, 0xd8, 0x45, 0x27, 0xe0, 0xc8, 0x0f, 0xa9, 0x43, 0x55, 0xdb, 0x43, 0xe7, 0x00,
	0xd9, 0x2e, 0x75, 0xde, 0xc7, 0x89, 0x45, 0xa2, 0x71, 0x48, 0x03, 0xa9, 0x84, 0x4f, 0x65, 0x70,
	0x2b, 0x63, 0xd8, 0x23, 0xf0, 0x19, 0x3a, 0x03, 0x70, 0x3f, 0x4b, 0xa1, 0x07, 0x08, 0x82, 0xe3,
	0xd6, 0x1e, 0x84, 0x74, 0x4c, 0xe0, 0x21, 0x7a, 0x01, 0x86, 0xb6, 0x1f, 0x7a, 0x38, 0x8e, 0x89,
	0xd5, 0x82, 0xcf, 0xd1, 0x4b, 0x70, 0xda, 0xca, 0x22, 0xec, 0x92, 0x16, 0x3e, 0x42, 0xaf, 0xc1,
	0xd9, 0x4e, 0xbb, 0xc7, 0x00, 0x39, 0x25, 0xf5, 0xb0, 0x43, 0x92, 0x69, 0xe8, 0xc2, 0x3e, 0x1a,
	0x00, 0x60, 0x53, 0x86, 0x5d, 0xd9, 0x46, 0xf0, 0x58, 0xe5, 0xa9, 0xde, 0xf3, 0x47, 0xd4, 0x25,
	0x0d, 0x7c, 0x22, 0xa7, 0x89, 0x43, 0x3c, 0x9e, 0x50, 0xe6, 0x28, 0xe3, 0x40, 0x4e, 0x8d, 0x99,
	0x15, 0xfa, 0xd4, 0x4a, 0x70, 0x10, 0x24, 0x2e, 0x65, 0x13, 0x38, 0x44, 0x6f, 0xc0, 0xab, 0x88,
	0x7a, 0xd4, 0xc5, 0x61, 0xf2, 0xfb, 0xff, 0x44, 0x10, 0xca, 0x10, 0xea, 0x47, 0x3b, 0xf9, 0xa9,
	0x0c, 0xd9, 0x20, 0x51, 0xec, 0x87, 0x44, 0xfe, 0x1c, 0x1a, 0xfd, 0xec, 0x80, 0xeb, 0x5b, 0xbe,
	0xd0, 0xfe, 0xbb, 0xf8, 0xd1, 0x65, 0xb3, 0x38, 0xf1, 0xe7, 0xe6, 0x02, 0xb9, 0xf8, 0xa0, 0xf3,
	0x71, 0xd4, 0xda, 0x0b, 0x5e, 0xa5, 0x75, 0xa1, 0xf1, 0x55, 0xa1, 0x17, 0x79, 0xad, 0xce, 0x62,
	0x73, 0x85, 0xcb, 0x52, 0xfc, 0xe3, 0x28, 0xdf, 0xaa, 0xf7, 0x4b, 0xb7, 0xe7, 0x60, 0xfc, 0xb5,
	0x7b, 0xe5, 0x34, 0x51, 0x38, 0x13, 0x5a, 0x53, 0xca, 0x6a, 0x66, 0x68, 0xf2, 0x88, 0xc4, 0xf7,
	0x0d, 0x3f, 0xc7, 0x99, 0x98, 0x6f, 0xf9, 0xf9, 0xcc, 0x98, 0x2b, 0xfe, 0x47, 0xf7, 0xba, 0x01,
	0x4d, 0x13, 0x67, 0xc2, 0x34, 0xb7, 0x0a, 0xd3, 0x9c, 0x19, 0xa6, 0xa9, 0x34, 0x9f, 0x0e, 0xd4,
	0x60, 0xc6, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x63, 0xf9, 0xca, 0x7e, 0x2c, 0x03, 0x00, 0x00,
}