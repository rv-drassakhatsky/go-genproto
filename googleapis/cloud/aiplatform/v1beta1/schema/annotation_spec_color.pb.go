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
// source: google/cloud/aiplatform/v1beta1/schema/annotation_spec_color.proto

package schema

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	color "google.golang.org/genproto/googleapis/type/color"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// An entry of mapping between color and AnnotationSpec. The mapping is used in
// segmentation mask.
type AnnotationSpecColor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The color of the AnnotationSpec in a segmentation mask.
	Color *color.Color `protobuf:"bytes,1,opt,name=color,proto3" json:"color,omitempty"`
	// The display name of the AnnotationSpec represented by the color in the
	// segmentation mask.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The ID of the AnnotationSpec represented by the color in the segmentation
	// mask.
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AnnotationSpecColor) Reset() {
	*x = AnnotationSpecColor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_schema_annotation_spec_color_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnotationSpecColor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnotationSpecColor) ProtoMessage() {}

func (x *AnnotationSpecColor) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_schema_annotation_spec_color_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnotationSpecColor.ProtoReflect.Descriptor instead.
func (*AnnotationSpecColor) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_schema_annotation_spec_color_proto_rawDescGZIP(), []int{0}
}

func (x *AnnotationSpecColor) GetColor() *color.Color {
	if x != nil {
		return x.Color
	}
	return nil
}

func (x *AnnotationSpecColor) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *AnnotationSpecColor) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_google_cloud_aiplatform_v1beta1_schema_annotation_spec_color_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1beta1_schema_annotation_spec_color_proto_rawDesc = []byte{
	0x0a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x1a, 0x17, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x13, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x70, 0x65, 0x63, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x28, 0x0a, 0x05, 0x63, 0x6f,
	0x6c, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x05, 0x63,
	0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x42, 0x96, 0x01, 0x0a, 0x2a, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x42, 0x18, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x4c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x3b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1beta1_schema_annotation_spec_color_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1beta1_schema_annotation_spec_color_proto_rawDescData = file_google_cloud_aiplatform_v1beta1_schema_annotation_spec_color_proto_rawDesc
)

func file_google_cloud_aiplatform_v1beta1_schema_annotation_spec_color_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1beta1_schema_annotation_spec_color_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1beta1_schema_annotation_spec_color_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1beta1_schema_annotation_spec_color_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1beta1_schema_annotation_spec_color_proto_rawDescData
}

var file_google_cloud_aiplatform_v1beta1_schema_annotation_spec_color_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_aiplatform_v1beta1_schema_annotation_spec_color_proto_goTypes = []interface{}{
	(*AnnotationSpecColor)(nil), // 0: google.cloud.aiplatform.v1beta1.schema.AnnotationSpecColor
	(*color.Color)(nil),         // 1: google.type.Color
}
var file_google_cloud_aiplatform_v1beta1_schema_annotation_spec_color_proto_depIdxs = []int32{
	1, // 0: google.cloud.aiplatform.v1beta1.schema.AnnotationSpecColor.color:type_name -> google.type.Color
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1beta1_schema_annotation_spec_color_proto_init() }
func file_google_cloud_aiplatform_v1beta1_schema_annotation_spec_color_proto_init() {
	if File_google_cloud_aiplatform_v1beta1_schema_annotation_spec_color_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1beta1_schema_annotation_spec_color_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnotationSpecColor); i {
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
			RawDescriptor: file_google_cloud_aiplatform_v1beta1_schema_annotation_spec_color_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1beta1_schema_annotation_spec_color_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1beta1_schema_annotation_spec_color_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1beta1_schema_annotation_spec_color_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1beta1_schema_annotation_spec_color_proto = out.File
	file_google_cloud_aiplatform_v1beta1_schema_annotation_spec_color_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1beta1_schema_annotation_spec_color_proto_goTypes = nil
	file_google_cloud_aiplatform_v1beta1_schema_annotation_spec_color_proto_depIdxs = nil
}
