// Copyright 2021 Google LLC
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/cloud/networksecurity/v1beta1/client_tls_policy.proto

package networksecurity

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ClientTlsPolicy is a resource that specifies how a client should authenticate
// connections to backends of a service. This resource itself does not affect
// configuration unless it is attached to a backend service resource.
type ClientTlsPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Name of the ClientTlsPolicy resource. It matches the pattern
	// `projects/*/locations/{location}/clientTlsPolicies/{client_tls_policy}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. Free-text description of the resource.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Output only. The timestamp when the resource was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The timestamp when the resource was updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Optional. Set of label tags associated with the resource.
	Labels map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Optional. Server Name Indication string to present to the server during TLS
	// handshake. E.g: "secure.example.com".
	Sni string `protobuf:"bytes,6,opt,name=sni,proto3" json:"sni,omitempty"`
	// Optional. Defines a mechanism to provision client identity (public and private keys)
	// for peer to peer authentication. The presence of this dictates mTLS.
	ClientCertificate *CertificateProvider `protobuf:"bytes,7,opt,name=client_certificate,json=clientCertificate,proto3" json:"client_certificate,omitempty"`
	// Optional. Defines the mechanism to obtain the Certificate Authority certificate to
	// validate the server certificate. If empty, client does not validate the
	// server certificate.
	ServerValidationCa []*ValidationCA `protobuf:"bytes,8,rep,name=server_validation_ca,json=serverValidationCa,proto3" json:"server_validation_ca,omitempty"`
}

func (x *ClientTlsPolicy) Reset() {
	*x = ClientTlsPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientTlsPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientTlsPolicy) ProtoMessage() {}

func (x *ClientTlsPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientTlsPolicy.ProtoReflect.Descriptor instead.
func (*ClientTlsPolicy) Descriptor() ([]byte, []int) {
	return file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_rawDescGZIP(), []int{0}
}

func (x *ClientTlsPolicy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClientTlsPolicy) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ClientTlsPolicy) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *ClientTlsPolicy) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *ClientTlsPolicy) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *ClientTlsPolicy) GetSni() string {
	if x != nil {
		return x.Sni
	}
	return ""
}

func (x *ClientTlsPolicy) GetClientCertificate() *CertificateProvider {
	if x != nil {
		return x.ClientCertificate
	}
	return nil
}

func (x *ClientTlsPolicy) GetServerValidationCa() []*ValidationCA {
	if x != nil {
		return x.ServerValidationCa
	}
	return nil
}

// Request used by the ListClientTlsPolicies method.
type ListClientTlsPoliciesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The project and location from which the ClientTlsPolicies should
	// be listed, specified in the format `projects/*/locations/{location}`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Maximum number of ClientTlsPolicies to return per call.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The value returned by the last `ListClientTlsPoliciesResponse`
	// Indicates that this is a continuation of a prior
	// `ListClientTlsPolicies` call, and that the system
	// should return the next page of data.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListClientTlsPoliciesRequest) Reset() {
	*x = ListClientTlsPoliciesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClientTlsPoliciesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClientTlsPoliciesRequest) ProtoMessage() {}

func (x *ListClientTlsPoliciesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClientTlsPoliciesRequest.ProtoReflect.Descriptor instead.
func (*ListClientTlsPoliciesRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_rawDescGZIP(), []int{1}
}

func (x *ListClientTlsPoliciesRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListClientTlsPoliciesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListClientTlsPoliciesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Response returned by the ListClientTlsPolicies method.
type ListClientTlsPoliciesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of ClientTlsPolicy resources.
	ClientTlsPolicies []*ClientTlsPolicy `protobuf:"bytes,1,rep,name=client_tls_policies,json=clientTlsPolicies,proto3" json:"client_tls_policies,omitempty"`
	// If there might be more results than those appearing in this response, then
	// `next_page_token` is included. To get the next set of results, call this
	// method again using the value of `next_page_token` as `page_token`.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListClientTlsPoliciesResponse) Reset() {
	*x = ListClientTlsPoliciesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClientTlsPoliciesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClientTlsPoliciesResponse) ProtoMessage() {}

func (x *ListClientTlsPoliciesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClientTlsPoliciesResponse.ProtoReflect.Descriptor instead.
func (*ListClientTlsPoliciesResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_rawDescGZIP(), []int{2}
}

func (x *ListClientTlsPoliciesResponse) GetClientTlsPolicies() []*ClientTlsPolicy {
	if x != nil {
		return x.ClientTlsPolicies
	}
	return nil
}

func (x *ListClientTlsPoliciesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// Request used by the GetClientTlsPolicy method.
type GetClientTlsPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. A name of the ClientTlsPolicy to get. Must be in the format
	// `projects/*/locations/{location}/clientTlsPolicies/*`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetClientTlsPolicyRequest) Reset() {
	*x = GetClientTlsPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientTlsPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientTlsPolicyRequest) ProtoMessage() {}

func (x *GetClientTlsPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClientTlsPolicyRequest.ProtoReflect.Descriptor instead.
func (*GetClientTlsPolicyRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_rawDescGZIP(), []int{3}
}

func (x *GetClientTlsPolicyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request used by the CreateClientTlsPolicy method.
type CreateClientTlsPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The parent resource of the ClientTlsPolicy. Must be in
	// the format `projects/*/locations/{location}`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. Short name of the ClientTlsPolicy resource to be created. This value should
	// be 1-63 characters long, containing only letters, numbers, hyphens, and
	// underscores, and should not start with a number. E.g. "client_mtls_policy".
	ClientTlsPolicyId string `protobuf:"bytes,2,opt,name=client_tls_policy_id,json=clientTlsPolicyId,proto3" json:"client_tls_policy_id,omitempty"`
	// Required. ClientTlsPolicy resource to be created.
	ClientTlsPolicy *ClientTlsPolicy `protobuf:"bytes,3,opt,name=client_tls_policy,json=clientTlsPolicy,proto3" json:"client_tls_policy,omitempty"`
}

func (x *CreateClientTlsPolicyRequest) Reset() {
	*x = CreateClientTlsPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClientTlsPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClientTlsPolicyRequest) ProtoMessage() {}

func (x *CreateClientTlsPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClientTlsPolicyRequest.ProtoReflect.Descriptor instead.
func (*CreateClientTlsPolicyRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_rawDescGZIP(), []int{4}
}

func (x *CreateClientTlsPolicyRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateClientTlsPolicyRequest) GetClientTlsPolicyId() string {
	if x != nil {
		return x.ClientTlsPolicyId
	}
	return ""
}

func (x *CreateClientTlsPolicyRequest) GetClientTlsPolicy() *ClientTlsPolicy {
	if x != nil {
		return x.ClientTlsPolicy
	}
	return nil
}

// Request used by UpdateClientTlsPolicy method.
type UpdateClientTlsPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. Field mask is used to specify the fields to be overwritten in the
	// ClientTlsPolicy resource by the update.  The fields
	// specified in the update_mask are relative to the resource, not
	// the full request. A field will be overwritten if it is in the
	// mask. If the user does not provide a mask then all fields will be
	// overwritten.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,1,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// Required. Updated ClientTlsPolicy resource.
	ClientTlsPolicy *ClientTlsPolicy `protobuf:"bytes,2,opt,name=client_tls_policy,json=clientTlsPolicy,proto3" json:"client_tls_policy,omitempty"`
}

func (x *UpdateClientTlsPolicyRequest) Reset() {
	*x = UpdateClientTlsPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateClientTlsPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateClientTlsPolicyRequest) ProtoMessage() {}

func (x *UpdateClientTlsPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateClientTlsPolicyRequest.ProtoReflect.Descriptor instead.
func (*UpdateClientTlsPolicyRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateClientTlsPolicyRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (x *UpdateClientTlsPolicyRequest) GetClientTlsPolicy() *ClientTlsPolicy {
	if x != nil {
		return x.ClientTlsPolicy
	}
	return nil
}

// Request used by the DeleteClientTlsPolicy method.
type DeleteClientTlsPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. A name of the ClientTlsPolicy to delete. Must be in
	// the format `projects/*/locations/{location}/clientTlsPolicies/*`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteClientTlsPolicyRequest) Reset() {
	*x = DeleteClientTlsPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteClientTlsPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClientTlsPolicyRequest) ProtoMessage() {}

func (x *DeleteClientTlsPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClientTlsPolicyRequest.ProtoReflect.Descriptor instead.
func (*DeleteClientTlsPolicyRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteClientTlsPolicyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_google_cloud_networksecurity_v1beta1_client_tls_policy_proto protoreflect.FileDescriptor

var file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x6c,
	0x73, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe6, 0x05, 0x0a, 0x0f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6c, 0x73, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x5e, 0x0a, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6c, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41,
	0x01, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x15, 0x0a, 0x03, 0x73, 0x6e, 0x69,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x03, 0x73, 0x6e, 0x69,
	0x12, 0x6d, 0x0a, 0x12, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x11, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12,
	0x69, 0x0a, 0x14, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x61, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x41, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x12, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x82, 0x01, 0xea, 0x41, 0x7f, 0x0a, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x54, 0x6c, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x4d, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6c, 0x73, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x6c, 0x73, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x7d, 0x22, 0x9d, 0x01, 0x0a, 0x1c, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6c, 0x73, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x02,
	0xfa, 0x41, 0x23, 0x0a, 0x21, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xae, 0x01, 0x0a, 0x1d, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6c, 0x73, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a, 0x13,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x6c, 0x73, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6c, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x52, 0x11, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6c, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65,
	0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x67, 0x0a, 0x19, 0x47,
	0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6c, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x36, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x30, 0x0a, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6c, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x8c, 0x02, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6c, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4e, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x36, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x30, 0x0a, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x54, 0x6c, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x14, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x6c, 0x73, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x11, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x54, 0x6c, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x12, 0x66, 0x0a, 0x11, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x6c, 0x73, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x54, 0x6c, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6c, 0x73, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x22, 0xc8, 0x01, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x54, 0x6c, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d,
	0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x66, 0x0a, 0x11, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x6c, 0x73, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54,
	0x6c, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6c, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x6a,
	0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6c,
	0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x36, 0xe0, 0x41,
	0x02, 0xfa, 0x41, 0x30, 0x0a, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6c, 0x73, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x8f, 0x02, 0x0a, 0x28, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x14, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54,
	0x6c, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x53, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x3b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0xaa, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x24, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xea, 0x02, 0x27, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_rawDescOnce sync.Once
	file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_rawDescData = file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_rawDesc
)

func file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_rawDescGZIP() []byte {
	file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_rawDescOnce.Do(func() {
		file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_rawDescData)
	})
	return file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_rawDescData
}

var file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_goTypes = []interface{}{
	(*ClientTlsPolicy)(nil),               // 0: google.cloud.networksecurity.v1beta1.ClientTlsPolicy
	(*ListClientTlsPoliciesRequest)(nil),  // 1: google.cloud.networksecurity.v1beta1.ListClientTlsPoliciesRequest
	(*ListClientTlsPoliciesResponse)(nil), // 2: google.cloud.networksecurity.v1beta1.ListClientTlsPoliciesResponse
	(*GetClientTlsPolicyRequest)(nil),     // 3: google.cloud.networksecurity.v1beta1.GetClientTlsPolicyRequest
	(*CreateClientTlsPolicyRequest)(nil),  // 4: google.cloud.networksecurity.v1beta1.CreateClientTlsPolicyRequest
	(*UpdateClientTlsPolicyRequest)(nil),  // 5: google.cloud.networksecurity.v1beta1.UpdateClientTlsPolicyRequest
	(*DeleteClientTlsPolicyRequest)(nil),  // 6: google.cloud.networksecurity.v1beta1.DeleteClientTlsPolicyRequest
	nil,                                   // 7: google.cloud.networksecurity.v1beta1.ClientTlsPolicy.LabelsEntry
	(*timestamppb.Timestamp)(nil),         // 8: google.protobuf.Timestamp
	(*CertificateProvider)(nil),           // 9: google.cloud.networksecurity.v1beta1.CertificateProvider
	(*ValidationCA)(nil),                  // 10: google.cloud.networksecurity.v1beta1.ValidationCA
	(*fieldmaskpb.FieldMask)(nil),         // 11: google.protobuf.FieldMask
}
var file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_depIdxs = []int32{
	8,  // 0: google.cloud.networksecurity.v1beta1.ClientTlsPolicy.create_time:type_name -> google.protobuf.Timestamp
	8,  // 1: google.cloud.networksecurity.v1beta1.ClientTlsPolicy.update_time:type_name -> google.protobuf.Timestamp
	7,  // 2: google.cloud.networksecurity.v1beta1.ClientTlsPolicy.labels:type_name -> google.cloud.networksecurity.v1beta1.ClientTlsPolicy.LabelsEntry
	9,  // 3: google.cloud.networksecurity.v1beta1.ClientTlsPolicy.client_certificate:type_name -> google.cloud.networksecurity.v1beta1.CertificateProvider
	10, // 4: google.cloud.networksecurity.v1beta1.ClientTlsPolicy.server_validation_ca:type_name -> google.cloud.networksecurity.v1beta1.ValidationCA
	0,  // 5: google.cloud.networksecurity.v1beta1.ListClientTlsPoliciesResponse.client_tls_policies:type_name -> google.cloud.networksecurity.v1beta1.ClientTlsPolicy
	0,  // 6: google.cloud.networksecurity.v1beta1.CreateClientTlsPolicyRequest.client_tls_policy:type_name -> google.cloud.networksecurity.v1beta1.ClientTlsPolicy
	11, // 7: google.cloud.networksecurity.v1beta1.UpdateClientTlsPolicyRequest.update_mask:type_name -> google.protobuf.FieldMask
	0,  // 8: google.cloud.networksecurity.v1beta1.UpdateClientTlsPolicyRequest.client_tls_policy:type_name -> google.cloud.networksecurity.v1beta1.ClientTlsPolicy
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_init() }
func file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_init() {
	if File_google_cloud_networksecurity_v1beta1_client_tls_policy_proto != nil {
		return
	}
	file_google_cloud_networksecurity_v1beta1_tls_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientTlsPolicy); i {
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
		file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClientTlsPoliciesRequest); i {
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
		file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClientTlsPoliciesResponse); i {
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
		file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClientTlsPolicyRequest); i {
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
		file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClientTlsPolicyRequest); i {
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
		file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateClientTlsPolicyRequest); i {
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
		file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteClientTlsPolicyRequest); i {
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
			RawDescriptor: file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_goTypes,
		DependencyIndexes: file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_depIdxs,
		MessageInfos:      file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_msgTypes,
	}.Build()
	File_google_cloud_networksecurity_v1beta1_client_tls_policy_proto = out.File
	file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_rawDesc = nil
	file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_goTypes = nil
	file_google_cloud_networksecurity_v1beta1_client_tls_policy_proto_depIdxs = nil
}
