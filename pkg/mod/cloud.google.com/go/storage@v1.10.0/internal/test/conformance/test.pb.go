// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package google_cloud_conformance_storage_v1

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type UrlStyle int32

const (
	UrlStyle_PATH_STYLE            UrlStyle = 0
	UrlStyle_VIRTUAL_HOSTED_STYLE  UrlStyle = 1
	UrlStyle_BUCKET_BOUND_HOSTNAME UrlStyle = 2
)

var UrlStyle_name = map[int32]string{
	0: "PATH_STYLE",
	1: "VIRTUAL_HOSTED_STYLE",
	2: "BUCKET_BOUND_HOSTNAME",
}

var UrlStyle_value = map[string]int32{
	"PATH_STYLE":            0,
	"VIRTUAL_HOSTED_STYLE":  1,
	"BUCKET_BOUND_HOSTNAME": 2,
}

func (x UrlStyle) String() string {
	return proto.EnumName(UrlStyle_name, int32(x))
}

func (UrlStyle) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

type TestFile struct {
	SigningV4Tests       []*SigningV4Test    `protobuf:"bytes,1,rep,name=signing_v4_tests,json=signingV4Tests,proto3" json:"signing_v4_tests,omitempty"`
	PostPolicyV4Tests    []*PostPolicyV4Test `protobuf:"bytes,2,rep,name=post_policy_v4_tests,json=postPolicyV4Tests,proto3" json:"post_policy_v4_tests,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *TestFile) Reset()         { *m = TestFile{} }
func (m *TestFile) String() string { return proto.CompactTextString(m) }
func (*TestFile) ProtoMessage()    {}
func (*TestFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *TestFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestFile.Unmarshal(m, b)
}
func (m *TestFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestFile.Marshal(b, m, deterministic)
}
func (m *TestFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestFile.Merge(m, src)
}
func (m *TestFile) XXX_Size() int {
	return xxx_messageInfo_TestFile.Size(m)
}
func (m *TestFile) XXX_DiscardUnknown() {
	xxx_messageInfo_TestFile.DiscardUnknown(m)
}

var xxx_messageInfo_TestFile proto.InternalMessageInfo

func (m *TestFile) GetSigningV4Tests() []*SigningV4Test {
	if m != nil {
		return m.SigningV4Tests
	}
	return nil
}

func (m *TestFile) GetPostPolicyV4Tests() []*PostPolicyV4Test {
	if m != nil {
		return m.PostPolicyV4Tests
	}
	return nil
}

type SigningV4Test struct {
	FileName                 string               `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	Description              string               `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Bucket                   string               `protobuf:"bytes,3,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Object                   string               `protobuf:"bytes,4,opt,name=object,proto3" json:"object,omitempty"`
	Method                   string               `protobuf:"bytes,5,opt,name=method,proto3" json:"method,omitempty"`
	Expiration               int64                `protobuf:"varint,6,opt,name=expiration,proto3" json:"expiration,omitempty"`
	Timestamp                *timestamp.Timestamp `protobuf:"bytes,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ExpectedUrl              string               `protobuf:"bytes,8,opt,name=expectedUrl,proto3" json:"expectedUrl,omitempty"`
	Headers                  map[string]string    `protobuf:"bytes,9,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	QueryParameters          map[string]string    `protobuf:"bytes,10,rep,name=query_parameters,json=queryParameters,proto3" json:"query_parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Scheme                   string               `protobuf:"bytes,11,opt,name=scheme,proto3" json:"scheme,omitempty"`
	UrlStyle                 UrlStyle             `protobuf:"varint,12,opt,name=urlStyle,proto3,enum=google.cloud.conformance.storage.v1.UrlStyle" json:"urlStyle,omitempty"`
	BucketBoundHostname      string               `protobuf:"bytes,13,opt,name=bucketBoundHostname,proto3" json:"bucketBoundHostname,omitempty"`
	ExpectedCanonicalRequest string               `protobuf:"bytes,14,opt,name=expectedCanonicalRequest,proto3" json:"expectedCanonicalRequest,omitempty"`
	ExpectedStringToSign     string               `protobuf:"bytes,15,opt,name=expectedStringToSign,proto3" json:"expectedStringToSign,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}             `json:"-"`
	XXX_unrecognized         []byte               `json:"-"`
	XXX_sizecache            int32                `json:"-"`
}

func (m *SigningV4Test) Reset()         { *m = SigningV4Test{} }
func (m *SigningV4Test) String() string { return proto.CompactTextString(m) }
func (*SigningV4Test) ProtoMessage()    {}
func (*SigningV4Test) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}

func (m *SigningV4Test) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SigningV4Test.Unmarshal(m, b)
}
func (m *SigningV4Test) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SigningV4Test.Marshal(b, m, deterministic)
}
func (m *SigningV4Test) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SigningV4Test.Merge(m, src)
}
func (m *SigningV4Test) XXX_Size() int {
	return xxx_messageInfo_SigningV4Test.Size(m)
}
func (m *SigningV4Test) XXX_DiscardUnknown() {
	xxx_messageInfo_SigningV4Test.DiscardUnknown(m)
}

var xxx_messageInfo_SigningV4Test proto.InternalMessageInfo

func (m *SigningV4Test) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *SigningV4Test) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SigningV4Test) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *SigningV4Test) GetObject() string {
	if m != nil {
		return m.Object
	}
	return ""
}

func (m *SigningV4Test) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *SigningV4Test) GetExpiration() int64 {
	if m != nil {
		return m.Expiration
	}
	return 0
}

func (m *SigningV4Test) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *SigningV4Test) GetExpectedUrl() string {
	if m != nil {
		return m.ExpectedUrl
	}
	return ""
}

func (m *SigningV4Test) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *SigningV4Test) GetQueryParameters() map[string]string {
	if m != nil {
		return m.QueryParameters
	}
	return nil
}

func (m *SigningV4Test) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *SigningV4Test) GetUrlStyle() UrlStyle {
	if m != nil {
		return m.UrlStyle
	}
	return UrlStyle_PATH_STYLE
}

func (m *SigningV4Test) GetBucketBoundHostname() string {
	if m != nil {
		return m.BucketBoundHostname
	}
	return ""
}

func (m *SigningV4Test) GetExpectedCanonicalRequest() string {
	if m != nil {
		return m.ExpectedCanonicalRequest
	}
	return ""
}

func (m *SigningV4Test) GetExpectedStringToSign() string {
	if m != nil {
		return m.ExpectedStringToSign
	}
	return ""
}

type ConditionalMatches struct {
	Expression           []string `protobuf:"bytes,1,rep,name=expression,proto3" json:"expression,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConditionalMatches) Reset()         { *m = ConditionalMatches{} }
func (m *ConditionalMatches) String() string { return proto.CompactTextString(m) }
func (*ConditionalMatches) ProtoMessage()    {}
func (*ConditionalMatches) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{2}
}

func (m *ConditionalMatches) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConditionalMatches.Unmarshal(m, b)
}
func (m *ConditionalMatches) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConditionalMatches.Marshal(b, m, deterministic)
}
func (m *ConditionalMatches) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConditionalMatches.Merge(m, src)
}
func (m *ConditionalMatches) XXX_Size() int {
	return xxx_messageInfo_ConditionalMatches.Size(m)
}
func (m *ConditionalMatches) XXX_DiscardUnknown() {
	xxx_messageInfo_ConditionalMatches.DiscardUnknown(m)
}

var xxx_messageInfo_ConditionalMatches proto.InternalMessageInfo

func (m *ConditionalMatches) GetExpression() []string {
	if m != nil {
		return m.Expression
	}
	return nil
}

type PolicyConditions struct {
	ContentLengthRange   []int32  `protobuf:"varint,1,rep,packed,name=contentLengthRange,proto3" json:"contentLengthRange,omitempty"`
	StartsWith           []string `protobuf:"bytes,2,rep,name=startsWith,proto3" json:"startsWith,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PolicyConditions) Reset()         { *m = PolicyConditions{} }
func (m *PolicyConditions) String() string { return proto.CompactTextString(m) }
func (*PolicyConditions) ProtoMessage()    {}
func (*PolicyConditions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{3}
}

func (m *PolicyConditions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolicyConditions.Unmarshal(m, b)
}
func (m *PolicyConditions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolicyConditions.Marshal(b, m, deterministic)
}
func (m *PolicyConditions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyConditions.Merge(m, src)
}
func (m *PolicyConditions) XXX_Size() int {
	return xxx_messageInfo_PolicyConditions.Size(m)
}
func (m *PolicyConditions) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyConditions.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyConditions proto.InternalMessageInfo

func (m *PolicyConditions) GetContentLengthRange() []int32 {
	if m != nil {
		return m.ContentLengthRange
	}
	return nil
}

func (m *PolicyConditions) GetStartsWith() []string {
	if m != nil {
		return m.StartsWith
	}
	return nil
}

type PolicyInput struct {
	Scheme               string               `protobuf:"bytes,1,opt,name=scheme,proto3" json:"scheme,omitempty"`
	UrlStyle             UrlStyle             `protobuf:"varint,2,opt,name=urlStyle,proto3,enum=google.cloud.conformance.storage.v1.UrlStyle" json:"urlStyle,omitempty"`
	BucketBoundHostname  string               `protobuf:"bytes,3,opt,name=bucketBoundHostname,proto3" json:"bucketBoundHostname,omitempty"`
	Bucket               string               `protobuf:"bytes,4,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Object               string               `protobuf:"bytes,5,opt,name=object,proto3" json:"object,omitempty"`
	Expiration           int32                `protobuf:"varint,6,opt,name=expiration,proto3" json:"expiration,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Fields               map[string]string    `protobuf:"bytes,8,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Conditions           *PolicyConditions    `protobuf:"bytes,9,opt,name=conditions,proto3" json:"conditions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PolicyInput) Reset()         { *m = PolicyInput{} }
func (m *PolicyInput) String() string { return proto.CompactTextString(m) }
func (*PolicyInput) ProtoMessage()    {}
func (*PolicyInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{4}
}

func (m *PolicyInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolicyInput.Unmarshal(m, b)
}
func (m *PolicyInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolicyInput.Marshal(b, m, deterministic)
}
func (m *PolicyInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyInput.Merge(m, src)
}
func (m *PolicyInput) XXX_Size() int {
	return xxx_messageInfo_PolicyInput.Size(m)
}
func (m *PolicyInput) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyInput.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyInput proto.InternalMessageInfo

func (m *PolicyInput) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *PolicyInput) GetUrlStyle() UrlStyle {
	if m != nil {
		return m.UrlStyle
	}
	return UrlStyle_PATH_STYLE
}

func (m *PolicyInput) GetBucketBoundHostname() string {
	if m != nil {
		return m.BucketBoundHostname
	}
	return ""
}

func (m *PolicyInput) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *PolicyInput) GetObject() string {
	if m != nil {
		return m.Object
	}
	return ""
}

func (m *PolicyInput) GetExpiration() int32 {
	if m != nil {
		return m.Expiration
	}
	return 0
}

func (m *PolicyInput) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *PolicyInput) GetFields() map[string]string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *PolicyInput) GetConditions() *PolicyConditions {
	if m != nil {
		return m.Conditions
	}
	return nil
}

type PolicyOutput struct {
	Url                   string            `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Fields                map[string]string `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ExpectedDecodedPolicy string            `protobuf:"bytes,3,opt,name=expectedDecodedPolicy,proto3" json:"expectedDecodedPolicy,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}          `json:"-"`
	XXX_unrecognized      []byte            `json:"-"`
	XXX_sizecache         int32             `json:"-"`
}

func (m *PolicyOutput) Reset()         { *m = PolicyOutput{} }
func (m *PolicyOutput) String() string { return proto.CompactTextString(m) }
func (*PolicyOutput) ProtoMessage()    {}
func (*PolicyOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{5}
}

func (m *PolicyOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolicyOutput.Unmarshal(m, b)
}
func (m *PolicyOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolicyOutput.Marshal(b, m, deterministic)
}
func (m *PolicyOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyOutput.Merge(m, src)
}
func (m *PolicyOutput) XXX_Size() int {
	return xxx_messageInfo_PolicyOutput.Size(m)
}
func (m *PolicyOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyOutput.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyOutput proto.InternalMessageInfo

func (m *PolicyOutput) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *PolicyOutput) GetFields() map[string]string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *PolicyOutput) GetExpectedDecodedPolicy() string {
	if m != nil {
		return m.ExpectedDecodedPolicy
	}
	return ""
}

type PostPolicyV4Test struct {
	Description          string        `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	PolicyInput          *PolicyInput  `protobuf:"bytes,2,opt,name=policyInput,proto3" json:"policyInput,omitempty"`
	PolicyOutput         *PolicyOutput `protobuf:"bytes,3,opt,name=policyOutput,proto3" json:"policyOutput,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PostPolicyV4Test) Reset()         { *m = PostPolicyV4Test{} }
func (m *PostPolicyV4Test) String() string { return proto.CompactTextString(m) }
func (*PostPolicyV4Test) ProtoMessage()    {}
func (*PostPolicyV4Test) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{6}
}

func (m *PostPolicyV4Test) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostPolicyV4Test.Unmarshal(m, b)
}
func (m *PostPolicyV4Test) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostPolicyV4Test.Marshal(b, m, deterministic)
}
func (m *PostPolicyV4Test) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostPolicyV4Test.Merge(m, src)
}
func (m *PostPolicyV4Test) XXX_Size() int {
	return xxx_messageInfo_PostPolicyV4Test.Size(m)
}
func (m *PostPolicyV4Test) XXX_DiscardUnknown() {
	xxx_messageInfo_PostPolicyV4Test.DiscardUnknown(m)
}

var xxx_messageInfo_PostPolicyV4Test proto.InternalMessageInfo

func (m *PostPolicyV4Test) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *PostPolicyV4Test) GetPolicyInput() *PolicyInput {
	if m != nil {
		return m.PolicyInput
	}
	return nil
}

func (m *PostPolicyV4Test) GetPolicyOutput() *PolicyOutput {
	if m != nil {
		return m.PolicyOutput
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.cloud.conformance.storage.v1.UrlStyle", UrlStyle_name, UrlStyle_value)
	proto.RegisterType((*TestFile)(nil), "google.cloud.conformance.storage.v1.TestFile")
	proto.RegisterType((*SigningV4Test)(nil), "google.cloud.conformance.storage.v1.SigningV4Test")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.conformance.storage.v1.SigningV4Test.HeadersEntry")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.conformance.storage.v1.SigningV4Test.QueryParametersEntry")
	proto.RegisterType((*ConditionalMatches)(nil), "google.cloud.conformance.storage.v1.ConditionalMatches")
	proto.RegisterType((*PolicyConditions)(nil), "google.cloud.conformance.storage.v1.PolicyConditions")
	proto.RegisterType((*PolicyInput)(nil), "google.cloud.conformance.storage.v1.PolicyInput")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.conformance.storage.v1.PolicyInput.FieldsEntry")
	proto.RegisterType((*PolicyOutput)(nil), "google.cloud.conformance.storage.v1.PolicyOutput")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.conformance.storage.v1.PolicyOutput.FieldsEntry")
	proto.RegisterType((*PostPolicyV4Test)(nil), "google.cloud.conformance.storage.v1.PostPolicyV4Test")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 901 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xef, 0x6e, 0xe2, 0x46,
	0x10, 0xaf, 0x21, 0x70, 0x30, 0xe4, 0x72, 0x74, 0xcb, 0x55, 0x2e, 0x1f, 0x5a, 0x44, 0x3f, 0x94,
	0x56, 0xaa, 0xef, 0x42, 0x73, 0xd2, 0x35, 0x6a, 0x55, 0x05, 0xc2, 0x5d, 0xa2, 0xe6, 0x12, 0x6a,
	0x4c, 0x4e, 0x27, 0x55, 0x42, 0xc6, 0x1e, 0xc0, 0x3d, 0xe3, 0x75, 0xbc, 0xeb, 0xe8, 0x78, 0x8f,
	0x3e, 0x45, 0x9f, 0xa6, 0x6f, 0xd0, 0x4f, 0x95, 0xfa, 0x18, 0xd5, 0xee, 0xda, 0x60, 0x2e, 0x44,
	0x82, 0xfe, 0xf9, 0xc6, 0xcc, 0xec, 0xfc, 0x7e, 0xeb, 0x99, 0xdf, 0xcc, 0x02, 0xc0, 0x91, 0x71,
	0x23, 0x8c, 0x28, 0xa7, 0xe4, 0xf3, 0x29, 0xa5, 0x53, 0x1f, 0x0d, 0xc7, 0xa7, 0xb1, 0x6b, 0x38,
	0x34, 0x98, 0xd0, 0x68, 0x6e, 0x07, 0x0e, 0x1a, 0x8c, 0xd3, 0xc8, 0x9e, 0xa2, 0x71, 0x7b, 0x58,
	0xff, 0x4c, 0x1d, 0x7a, 0x22, 0x53, 0xc6, 0xf1, 0xe4, 0x09, 0xf7, 0xe6, 0xc8, 0xb8, 0x3d, 0x0f,
	0x15, 0x4a, 0xf3, 0x77, 0x0d, 0x4a, 0x16, 0x32, 0xfe, 0xc2, 0xf3, 0x91, 0xfc, 0x0c, 0x55, 0xe6,
	0x4d, 0x03, 0x2f, 0x98, 0x8e, 0x6e, 0x8f, 0x46, 0x82, 0x8b, 0xe9, 0x5a, 0x23, 0xdf, 0xaa, 0xb4,
	0xdb, 0xc6, 0x16, 0x6c, 0xc6, 0x40, 0x25, 0x5f, 0x1f, 0x09, 0x44, 0xf3, 0x80, 0x65, 0x4d, 0x46,
	0x26, 0x50, 0x0b, 0x29, 0xe3, 0xa3, 0x90, 0xfa, 0x9e, 0xb3, 0x58, 0x31, 0xe4, 0x24, 0xc3, 0xb3,
	0xad, 0x18, 0xfa, 0x94, 0xf1, 0xbe, 0xcc, 0x4f, 0x48, 0x3e, 0x0c, 0xdf, 0xf3, 0xb0, 0xe6, 0x9f,
	0x45, 0x78, 0xb8, 0x76, 0x13, 0x52, 0x87, 0xd2, 0xc4, 0xf3, 0xf1, 0xd2, 0x9e, 0xa3, 0xae, 0x35,
	0xb4, 0x56, 0xd9, 0x5c, 0xda, 0xa4, 0x01, 0x15, 0x17, 0x99, 0x13, 0x79, 0x21, 0xf7, 0x68, 0xa0,
	0xe7, 0x64, 0x38, 0xeb, 0x22, 0x1f, 0x43, 0x71, 0x1c, 0x3b, 0x6f, 0x91, 0xeb, 0x79, 0x19, 0x4c,
	0x2c, 0xe1, 0xa7, 0xe3, 0x5f, 0xd0, 0xe1, 0xfa, 0x9e, 0xf2, 0x2b, 0x4b, 0xf8, 0xe7, 0xc8, 0x67,
	0xd4, 0xd5, 0x0b, 0xca, 0xaf, 0x2c, 0xf2, 0x29, 0x00, 0xbe, 0x0b, 0xbd, 0xc8, 0x96, 0x44, 0xc5,
	0x86, 0xd6, 0xca, 0x9b, 0x19, 0x0f, 0x79, 0x0e, 0xe5, 0x65, 0x77, 0xf4, 0x07, 0x0d, 0xad, 0x55,
	0x69, 0xd7, 0xd3, 0xa2, 0xa4, 0xfd, 0x33, 0xac, 0xf4, 0x84, 0xb9, 0x3a, 0x2c, 0xbe, 0x01, 0xdf,
	0x85, 0xe8, 0x70, 0x74, 0x87, 0x91, 0xaf, 0x97, 0xd4, 0x37, 0x64, 0x5c, 0xe4, 0x0d, 0x3c, 0x98,
	0xa1, 0xed, 0x62, 0xc4, 0xf4, 0xb2, 0x2c, 0xf7, 0x0f, 0xbb, 0x37, 0xd4, 0x38, 0x53, 0x08, 0xbd,
	0x80, 0x47, 0x0b, 0x33, 0xc5, 0x23, 0x11, 0x54, 0x6f, 0x62, 0x8c, 0x16, 0xa3, 0xd0, 0x8e, 0xec,
	0x39, 0x72, 0xc1, 0x01, 0x92, 0xe3, 0xe5, 0x3f, 0xe0, 0xf8, 0x49, 0x40, 0xf5, 0x97, 0x48, 0x8a,
	0xeb, 0xd1, 0xcd, 0xba, 0x57, 0x94, 0x98, 0x39, 0x33, 0x9c, 0xa3, 0x5e, 0x51, 0x25, 0x56, 0x16,
	0x39, 0x87, 0x52, 0x1c, 0xf9, 0x03, 0xbe, 0xf0, 0x51, 0xdf, 0x6f, 0x68, 0xad, 0x83, 0xf6, 0xd7,
	0x5b, 0xdd, 0x61, 0x98, 0x24, 0x99, 0xcb, 0x74, 0xf2, 0x14, 0x3e, 0x52, 0x7d, 0xee, 0xd0, 0x38,
	0x70, 0xcf, 0x28, 0xe3, 0x81, 0x90, 0xcf, 0x43, 0xc9, 0xb7, 0x29, 0x44, 0x8e, 0x41, 0x4f, 0x4b,
	0xde, 0xb5, 0x03, 0x1a, 0x78, 0x8e, 0xed, 0x9b, 0x78, 0x13, 0x23, 0xe3, 0xfa, 0x81, 0x4c, 0xbb,
	0x37, 0x4e, 0xda, 0x50, 0x4b, 0x63, 0x03, 0x1e, 0x79, 0xc1, 0xd4, 0xa2, 0xa2, 0x2e, 0xfa, 0x23,
	0x99, 0xb7, 0x31, 0x56, 0x3f, 0x86, 0xfd, 0x6c, 0x47, 0x48, 0x15, 0xf2, 0x6f, 0x71, 0x91, 0x08,
	0x5c, 0xfc, 0x24, 0x35, 0x28, 0xdc, 0xda, 0x7e, 0x8c, 0x89, 0xaa, 0x95, 0x71, 0x9c, 0x7b, 0xae,
	0xd5, 0x3b, 0x50, 0xdb, 0x54, 0xe9, 0x5d, 0x30, 0x9a, 0x47, 0x40, 0xba, 0x34, 0x70, 0x3d, 0x21,
	0x5e, 0xdb, 0x7f, 0x65, 0x73, 0x67, 0x86, 0x2c, 0x51, 0x79, 0x84, 0x8c, 0x09, 0x95, 0x8b, 0xed,
	0x51, 0x36, 0x33, 0x9e, 0xe6, 0x18, 0xaa, 0x6a, 0x5c, 0x97, 0xb9, 0x8c, 0x18, 0x40, 0x1c, 0x1a,
	0x70, 0x0c, 0xf8, 0x05, 0x06, 0x53, 0x3e, 0x33, 0xed, 0x60, 0x8a, 0x32, 0xb7, 0x60, 0x6e, 0x88,
	0x08, 0x0e, 0xc6, 0xed, 0x88, 0xb3, 0xd7, 0x1e, 0x9f, 0xc9, 0xfd, 0x51, 0x36, 0x33, 0x9e, 0xe6,
	0xaf, 0x7b, 0x50, 0x51, 0x24, 0xe7, 0x41, 0x18, 0xf3, 0x8c, 0x5c, 0xb4, 0x7b, 0xe5, 0x92, 0xfb,
	0x5f, 0xe4, 0x92, 0xbf, 0x5f, 0x2e, 0xab, 0xb5, 0xb2, 0x77, 0xcf, 0x5a, 0x29, 0xac, 0xad, 0x95,
	0xbb, 0xeb, 0xa3, 0xf0, 0x1f, 0xad, 0x0f, 0x0b, 0x8a, 0x13, 0x0f, 0x7d, 0x97, 0xe9, 0x25, 0x39,
	0xb7, 0xdf, 0x6d, 0xb9, 0x8a, 0x97, 0x05, 0x36, 0x5e, 0xc8, 0x74, 0x35, 0xac, 0x09, 0x16, 0x19,
	0x02, 0x38, 0xcb, 0x16, 0xeb, 0x65, 0x79, 0xa1, 0x67, 0x3b, 0x20, 0xaf, 0xf4, 0x61, 0x66, 0x80,
	0xea, 0xdf, 0x42, 0x25, 0xc3, 0xb6, 0x93, 0x60, 0xff, 0xd2, 0x60, 0x5f, 0x61, 0x5f, 0xc5, 0x5c,
	0xe8, 0xa2, 0x0a, 0xf9, 0x38, 0xf2, 0xd3, 0xe4, 0x38, 0xf2, 0xc9, 0x70, 0x59, 0x0a, 0xf5, 0x2a,
	0x7d, 0xbf, 0xc3, 0x85, 0x15, 0xe8, 0xc6, 0x5a, 0x1c, 0xc1, 0xe3, 0x74, 0x84, 0x4f, 0xd1, 0xa1,
	0x2e, 0xba, 0x2a, 0x25, 0xd1, 0xc7, 0xe6, 0xe0, 0xbf, 0xf9, 0xd4, 0x3f, 0x34, 0x31, 0x66, 0xeb,
	0x2f, 0xe3, 0xfb, 0x4f, 0x9d, 0x76, 0xf7, 0xa9, 0x33, 0xa1, 0x12, 0xae, 0xda, 0x2a, 0x61, 0x2b,
	0xed, 0xa7, 0xbb, 0xca, 0xc1, 0xcc, 0x82, 0x90, 0x21, 0xec, 0x87, 0x99, 0xfa, 0xc8, 0x4f, 0xae,
	0xb4, 0x0f, 0x77, 0x2e, 0xac, 0xb9, 0x06, 0xf3, 0xd5, 0x15, 0x94, 0xd2, 0x31, 0x24, 0x07, 0x00,
	0xfd, 0x13, 0xeb, 0x6c, 0x34, 0xb0, 0xde, 0x5c, 0xf4, 0xaa, 0x1f, 0x10, 0x1d, 0x6a, 0xd7, 0xe7,
	0xa6, 0x35, 0x3c, 0xb9, 0x18, 0x9d, 0x5d, 0x0d, 0xac, 0xde, 0x69, 0x12, 0xd1, 0xc8, 0x27, 0xf0,
	0xb8, 0x33, 0xec, 0xfe, 0xd8, 0xb3, 0x46, 0x9d, 0xab, 0xe1, 0xe5, 0xa9, 0x0c, 0x5f, 0x9e, 0xbc,
	0xea, 0x55, 0x73, 0x9d, 0xd7, 0xf0, 0x85, 0x43, 0xe7, 0xdb, 0x5c, 0xab, 0xaf, 0xfd, 0x96, 0xfb,
	0xf2, 0xa5, 0x3a, 0xd7, 0x95, 0xe7, 0x06, 0x49, 0xec, 0xfa, 0xd0, 0x90, 0xff, 0x41, 0x8c, 0xee,
	0x2a, 0x71, 0x5c, 0x94, 0xd3, 0xf7, 0xcd, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x35, 0x7d,
	0xfd, 0xbd, 0x09, 0x00, 0x00,
}
