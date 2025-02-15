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
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: google/appengine/v1beta/operation.proto

package appengine

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Metadata for the given [google.longrunning.Operation][google.longrunning.Operation].
type OperationMetadataV1Beta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// API method that initiated this operation. Example:
	// `google.appengine.v1beta.Versions.CreateVersion`.
	//
	// @OutputOnly
	Method string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	// Time that this operation was created.
	//
	// @OutputOnly
	InsertTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=insert_time,json=insertTime,proto3" json:"insert_time,omitempty"`
	// Time that this operation completed.
	//
	// @OutputOnly
	EndTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// User who requested this operation.
	//
	// @OutputOnly
	User string `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	// Name of the resource that this operation is acting on. Example:
	// `apps/myapp/services/default`.
	//
	// @OutputOnly
	Target string `protobuf:"bytes,5,opt,name=target,proto3" json:"target,omitempty"`
	// Ephemeral message that may change every time the operation is polled.
	// @OutputOnly
	EphemeralMessage string `protobuf:"bytes,6,opt,name=ephemeral_message,json=ephemeralMessage,proto3" json:"ephemeral_message,omitempty"`
	// Durable messages that persist on every operation poll.
	// @OutputOnly
	Warning []string `protobuf:"bytes,7,rep,name=warning,proto3" json:"warning,omitempty"`
	// Metadata specific to the type of operation in progress.
	// @OutputOnly
	//
	// Types that are assignable to MethodMetadata:
	//	*OperationMetadataV1Beta_CreateVersionMetadata
	MethodMetadata isOperationMetadataV1Beta_MethodMetadata `protobuf_oneof:"method_metadata"`
}

func (x *OperationMetadataV1Beta) Reset() {
	*x = OperationMetadataV1Beta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_appengine_v1beta_operation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationMetadataV1Beta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationMetadataV1Beta) ProtoMessage() {}

func (x *OperationMetadataV1Beta) ProtoReflect() protoreflect.Message {
	mi := &file_google_appengine_v1beta_operation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationMetadataV1Beta.ProtoReflect.Descriptor instead.
func (*OperationMetadataV1Beta) Descriptor() ([]byte, []int) {
	return file_google_appengine_v1beta_operation_proto_rawDescGZIP(), []int{0}
}

func (x *OperationMetadataV1Beta) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *OperationMetadataV1Beta) GetInsertTime() *timestamppb.Timestamp {
	if x != nil {
		return x.InsertTime
	}
	return nil
}

func (x *OperationMetadataV1Beta) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *OperationMetadataV1Beta) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *OperationMetadataV1Beta) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *OperationMetadataV1Beta) GetEphemeralMessage() string {
	if x != nil {
		return x.EphemeralMessage
	}
	return ""
}

func (x *OperationMetadataV1Beta) GetWarning() []string {
	if x != nil {
		return x.Warning
	}
	return nil
}

func (m *OperationMetadataV1Beta) GetMethodMetadata() isOperationMetadataV1Beta_MethodMetadata {
	if m != nil {
		return m.MethodMetadata
	}
	return nil
}

func (x *OperationMetadataV1Beta) GetCreateVersionMetadata() *CreateVersionMetadataV1Beta {
	if x, ok := x.GetMethodMetadata().(*OperationMetadataV1Beta_CreateVersionMetadata); ok {
		return x.CreateVersionMetadata
	}
	return nil
}

type isOperationMetadataV1Beta_MethodMetadata interface {
	isOperationMetadataV1Beta_MethodMetadata()
}

type OperationMetadataV1Beta_CreateVersionMetadata struct {
	CreateVersionMetadata *CreateVersionMetadataV1Beta `protobuf:"bytes,8,opt,name=create_version_metadata,json=createVersionMetadata,proto3,oneof"`
}

func (*OperationMetadataV1Beta_CreateVersionMetadata) isOperationMetadataV1Beta_MethodMetadata() {}

// Metadata for the given [google.longrunning.Operation][google.longrunning.Operation] during a
// [google.appengine.v1beta.CreateVersionRequest][google.appengine.v1beta.CreateVersionRequest].
type CreateVersionMetadataV1Beta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Cloud Build ID if one was created as part of the version create.
	// @OutputOnly
	CloudBuildId string `protobuf:"bytes,1,opt,name=cloud_build_id,json=cloudBuildId,proto3" json:"cloud_build_id,omitempty"`
}

func (x *CreateVersionMetadataV1Beta) Reset() {
	*x = CreateVersionMetadataV1Beta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_appengine_v1beta_operation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVersionMetadataV1Beta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVersionMetadataV1Beta) ProtoMessage() {}

func (x *CreateVersionMetadataV1Beta) ProtoReflect() protoreflect.Message {
	mi := &file_google_appengine_v1beta_operation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVersionMetadataV1Beta.ProtoReflect.Descriptor instead.
func (*CreateVersionMetadataV1Beta) Descriptor() ([]byte, []int) {
	return file_google_appengine_v1beta_operation_proto_rawDescGZIP(), []int{1}
}

func (x *CreateVersionMetadataV1Beta) GetCloudBuildId() string {
	if x != nil {
		return x.CloudBuildId
	}
	return ""
}

var File_google_appengine_v1beta_operation_proto protoreflect.FileDescriptor

var file_google_appengine_v1beta_operation_proto_rawDesc = []byte{
	0x0a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x9b, 0x03, 0x0a, 0x17, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x65, 0x70, 0x68, 0x65, 0x6d, 0x65, 0x72,
	0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x65, 0x70, 0x68, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x6e, 0x0a, 0x17,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x56, 0x31, 0x42,
	0x65, 0x74, 0x61, 0x48, 0x00, 0x52, 0x15, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x11, 0x0a, 0x0f,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x43, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x12, 0x24,
	0x0a, 0x0e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x49, 0x64, 0x42, 0xd4, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x42, 0x0e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70,
	0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x3b, 0x61,
	0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0xaa, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0xca, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0xea, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_appengine_v1beta_operation_proto_rawDescOnce sync.Once
	file_google_appengine_v1beta_operation_proto_rawDescData = file_google_appengine_v1beta_operation_proto_rawDesc
)

func file_google_appengine_v1beta_operation_proto_rawDescGZIP() []byte {
	file_google_appengine_v1beta_operation_proto_rawDescOnce.Do(func() {
		file_google_appengine_v1beta_operation_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_appengine_v1beta_operation_proto_rawDescData)
	})
	return file_google_appengine_v1beta_operation_proto_rawDescData
}

var file_google_appengine_v1beta_operation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_appengine_v1beta_operation_proto_goTypes = []interface{}{
	(*OperationMetadataV1Beta)(nil),     // 0: google.appengine.v1beta.OperationMetadataV1Beta
	(*CreateVersionMetadataV1Beta)(nil), // 1: google.appengine.v1beta.CreateVersionMetadataV1Beta
	(*timestamppb.Timestamp)(nil),       // 2: google.protobuf.Timestamp
}
var file_google_appengine_v1beta_operation_proto_depIdxs = []int32{
	2, // 0: google.appengine.v1beta.OperationMetadataV1Beta.insert_time:type_name -> google.protobuf.Timestamp
	2, // 1: google.appengine.v1beta.OperationMetadataV1Beta.end_time:type_name -> google.protobuf.Timestamp
	1, // 2: google.appengine.v1beta.OperationMetadataV1Beta.create_version_metadata:type_name -> google.appengine.v1beta.CreateVersionMetadataV1Beta
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_appengine_v1beta_operation_proto_init() }
func file_google_appengine_v1beta_operation_proto_init() {
	if File_google_appengine_v1beta_operation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_appengine_v1beta_operation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationMetadataV1Beta); i {
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
		file_google_appengine_v1beta_operation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVersionMetadataV1Beta); i {
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
	file_google_appengine_v1beta_operation_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*OperationMetadataV1Beta_CreateVersionMetadata)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_appengine_v1beta_operation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_appengine_v1beta_operation_proto_goTypes,
		DependencyIndexes: file_google_appengine_v1beta_operation_proto_depIdxs,
		MessageInfos:      file_google_appengine_v1beta_operation_proto_msgTypes,
	}.Build()
	File_google_appengine_v1beta_operation_proto = out.File
	file_google_appengine_v1beta_operation_proto_rawDesc = nil
	file_google_appengine_v1beta_operation_proto_goTypes = nil
	file_google_appengine_v1beta_operation_proto_depIdxs = nil
}
