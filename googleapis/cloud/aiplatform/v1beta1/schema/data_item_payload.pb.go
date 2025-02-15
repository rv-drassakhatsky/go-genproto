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
// source: google/cloud/aiplatform/v1beta1/schema/data_item_payload.proto

package schema

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/durationpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Payload of Image DataItem.
type ImageDataItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Google Cloud Storage URI points to the original image in user's bucket.
	// The image is up to 30MB in size.
	GcsUri string `protobuf:"bytes,1,opt,name=gcs_uri,json=gcsUri,proto3" json:"gcs_uri,omitempty"`
	// Output only. The mime type of the content of the image. Only the images in below listed
	// mime types are supported.
	// - image/jpeg
	// - image/gif
	// - image/png
	// - image/webp
	// - image/bmp
	// - image/tiff
	// - image/vnd.microsoft.icon
	MimeType string `protobuf:"bytes,2,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
}

func (x *ImageDataItem) Reset() {
	*x = ImageDataItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageDataItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageDataItem) ProtoMessage() {}

func (x *ImageDataItem) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageDataItem.ProtoReflect.Descriptor instead.
func (*ImageDataItem) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_rawDescGZIP(), []int{0}
}

func (x *ImageDataItem) GetGcsUri() string {
	if x != nil {
		return x.GcsUri
	}
	return ""
}

func (x *ImageDataItem) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

// Payload of Video DataItem.
type VideoDataItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Google Cloud Storage URI points to the original video in user's bucket.
	// The video is up to 50 GB in size and up to 3 hour in duration.
	GcsUri string `protobuf:"bytes,1,opt,name=gcs_uri,json=gcsUri,proto3" json:"gcs_uri,omitempty"`
	// Output only. The mime type of the content of the video. Only the videos in below listed
	// mime types are supported.
	// Supported mime_type:
	// - video/mp4
	// - video/avi
	// - video/quicktime
	MimeType string `protobuf:"bytes,2,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
}

func (x *VideoDataItem) Reset() {
	*x = VideoDataItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoDataItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoDataItem) ProtoMessage() {}

func (x *VideoDataItem) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoDataItem.ProtoReflect.Descriptor instead.
func (*VideoDataItem) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_rawDescGZIP(), []int{1}
}

func (x *VideoDataItem) GetGcsUri() string {
	if x != nil {
		return x.GcsUri
	}
	return ""
}

func (x *VideoDataItem) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

// Payload of Text DataItem.
type TextDataItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Google Cloud Storage URI points to the original text in user's bucket.
	// The text file is up to 10MB in size.
	GcsUri string `protobuf:"bytes,1,opt,name=gcs_uri,json=gcsUri,proto3" json:"gcs_uri,omitempty"`
}

func (x *TextDataItem) Reset() {
	*x = TextDataItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextDataItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextDataItem) ProtoMessage() {}

func (x *TextDataItem) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextDataItem.ProtoReflect.Descriptor instead.
func (*TextDataItem) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_rawDescGZIP(), []int{2}
}

func (x *TextDataItem) GetGcsUri() string {
	if x != nil {
		return x.GcsUri
	}
	return ""
}

var File_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x74,
	0x65, 0x6d, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x0d, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1c, 0x0a, 0x07, 0x67, 0x63, 0x73, 0x5f,
	0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06,
	0x67, 0x63, 0x73, 0x55, 0x72, 0x69, 0x12, 0x20, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x08,
	0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4f, 0x0a, 0x0d, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x44, 0x61, 0x74, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1c, 0x0a, 0x07, 0x67, 0x63, 0x73,
	0x5f, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x06, 0x67, 0x63, 0x73, 0x55, 0x72, 0x69, 0x12, 0x20, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x2c, 0x0a, 0x0c, 0x54, 0x65, 0x78,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1c, 0x0a, 0x07, 0x67, 0x63, 0x73,
	0x5f, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x06, 0x67, 0x63, 0x73, 0x55, 0x72, 0x69, 0x42, 0x92, 0x01, 0x0a, 0x2a, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x42, 0x14, 0x44, 0x61, 0x74, 0x61, 0x49, 0x74, 0x65, 0x6d,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x3b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_rawDescData = file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_rawDesc
)

func file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_rawDescData
}

var file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_goTypes = []interface{}{
	(*ImageDataItem)(nil), // 0: google.cloud.aiplatform.v1beta1.schema.ImageDataItem
	(*VideoDataItem)(nil), // 1: google.cloud.aiplatform.v1beta1.schema.VideoDataItem
	(*TextDataItem)(nil),  // 2: google.cloud.aiplatform.v1beta1.schema.TextDataItem
}
var file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_init() }
func file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_init() {
	if File_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageDataItem); i {
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
		file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoDataItem); i {
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
		file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextDataItem); i {
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
			RawDescriptor: file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto = out.File
	file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_goTypes = nil
	file_google_cloud_aiplatform_v1beta1_schema_data_item_payload_proto_depIdxs = nil
}
