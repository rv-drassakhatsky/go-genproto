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
// source: google/appengine/v1beta/deploy.proto

package appengine

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Code and application artifacts used to deploy a version to App Engine.
type Deployment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Manifest of the files stored in Google Cloud Storage that are included
	// as part of this version. All files must be readable using the
	// credentials supplied with this call.
	Files map[string]*FileInfo `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The Docker image for the container that runs the version.
	// Only applicable for instances running in the App Engine flexible environment.
	Container *ContainerInfo `protobuf:"bytes,2,opt,name=container,proto3" json:"container,omitempty"`
	// The zip file for this deployment, if this is a zip deployment.
	Zip *ZipInfo `protobuf:"bytes,3,opt,name=zip,proto3" json:"zip,omitempty"`
	// Google Cloud Build build information. Only applicable for instances running
	// in the App Engine flexible environment.
	Build *BuildInfo `protobuf:"bytes,5,opt,name=build,proto3" json:"build,omitempty"`
	// Options for any Google Cloud Build builds created as a part of this
	// deployment.
	//
	// These options will only be used if a new build is created, such as when
	// deploying to the App Engine flexible environment using files or zip.
	CloudBuildOptions *CloudBuildOptions `protobuf:"bytes,6,opt,name=cloud_build_options,json=cloudBuildOptions,proto3" json:"cloud_build_options,omitempty"`
}

func (x *Deployment) Reset() {
	*x = Deployment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_appengine_v1beta_deploy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deployment) ProtoMessage() {}

func (x *Deployment) ProtoReflect() protoreflect.Message {
	mi := &file_google_appengine_v1beta_deploy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deployment.ProtoReflect.Descriptor instead.
func (*Deployment) Descriptor() ([]byte, []int) {
	return file_google_appengine_v1beta_deploy_proto_rawDescGZIP(), []int{0}
}

func (x *Deployment) GetFiles() map[string]*FileInfo {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *Deployment) GetContainer() *ContainerInfo {
	if x != nil {
		return x.Container
	}
	return nil
}

func (x *Deployment) GetZip() *ZipInfo {
	if x != nil {
		return x.Zip
	}
	return nil
}

func (x *Deployment) GetBuild() *BuildInfo {
	if x != nil {
		return x.Build
	}
	return nil
}

func (x *Deployment) GetCloudBuildOptions() *CloudBuildOptions {
	if x != nil {
		return x.CloudBuildOptions
	}
	return nil
}

// Single source file that is part of the version to be deployed. Each source
// file that is deployed must be specified separately.
type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// URL source to use to fetch this file. Must be a URL to a resource in
	// Google Cloud Storage in the form
	// 'http(s)://storage.googleapis.com/\<bucket\>/\<object\>'.
	SourceUrl string `protobuf:"bytes,1,opt,name=source_url,json=sourceUrl,proto3" json:"source_url,omitempty"`
	// The SHA1 hash of the file, in hex.
	Sha1Sum string `protobuf:"bytes,2,opt,name=sha1_sum,json=sha1Sum,proto3" json:"sha1_sum,omitempty"`
	// The MIME type of the file.
	//
	// Defaults to the value from Google Cloud Storage.
	MimeType string `protobuf:"bytes,3,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_appengine_v1beta_deploy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_appengine_v1beta_deploy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_google_appengine_v1beta_deploy_proto_rawDescGZIP(), []int{1}
}

func (x *FileInfo) GetSourceUrl() string {
	if x != nil {
		return x.SourceUrl
	}
	return ""
}

func (x *FileInfo) GetSha1Sum() string {
	if x != nil {
		return x.Sha1Sum
	}
	return ""
}

func (x *FileInfo) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

// Docker image that is used to create a container and start a VM instance for
// the version that you deploy. Only applicable for instances running in the App
// Engine flexible environment.
type ContainerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// URI to the hosted container image in Google Container Registry. The URI
	// must be fully qualified and include a tag or digest.
	// Examples: "gcr.io/my-project/image:tag" or "gcr.io/my-project/image@digest"
	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *ContainerInfo) Reset() {
	*x = ContainerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_appengine_v1beta_deploy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerInfo) ProtoMessage() {}

func (x *ContainerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_appengine_v1beta_deploy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerInfo.ProtoReflect.Descriptor instead.
func (*ContainerInfo) Descriptor() ([]byte, []int) {
	return file_google_appengine_v1beta_deploy_proto_rawDescGZIP(), []int{2}
}

func (x *ContainerInfo) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

// Google Cloud Build information.
type BuildInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Google Cloud Build id.
	// Example: "f966068f-08b2-42c8-bdfe-74137dff2bf9"
	CloudBuildId string `protobuf:"bytes,1,opt,name=cloud_build_id,json=cloudBuildId,proto3" json:"cloud_build_id,omitempty"`
}

func (x *BuildInfo) Reset() {
	*x = BuildInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_appengine_v1beta_deploy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildInfo) ProtoMessage() {}

func (x *BuildInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_appengine_v1beta_deploy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildInfo.ProtoReflect.Descriptor instead.
func (*BuildInfo) Descriptor() ([]byte, []int) {
	return file_google_appengine_v1beta_deploy_proto_rawDescGZIP(), []int{3}
}

func (x *BuildInfo) GetCloudBuildId() string {
	if x != nil {
		return x.CloudBuildId
	}
	return ""
}

// Options for the build operations performed as a part of the version
// deployment. Only applicable for App Engine flexible environment when creating
// a version using source code directly.
type CloudBuildOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Path to the yaml file used in deployment, used to determine runtime
	// configuration details.
	//
	// Required for flexible environment builds.
	//
	// See https://cloud.google.com/appengine/docs/standard/python/config/appref
	// for more details.
	AppYamlPath string `protobuf:"bytes,1,opt,name=app_yaml_path,json=appYamlPath,proto3" json:"app_yaml_path,omitempty"`
	// The Cloud Build timeout used as part of any dependent builds performed by
	// version creation. Defaults to 10 minutes.
	CloudBuildTimeout *durationpb.Duration `protobuf:"bytes,2,opt,name=cloud_build_timeout,json=cloudBuildTimeout,proto3" json:"cloud_build_timeout,omitempty"`
}

func (x *CloudBuildOptions) Reset() {
	*x = CloudBuildOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_appengine_v1beta_deploy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudBuildOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudBuildOptions) ProtoMessage() {}

func (x *CloudBuildOptions) ProtoReflect() protoreflect.Message {
	mi := &file_google_appengine_v1beta_deploy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudBuildOptions.ProtoReflect.Descriptor instead.
func (*CloudBuildOptions) Descriptor() ([]byte, []int) {
	return file_google_appengine_v1beta_deploy_proto_rawDescGZIP(), []int{4}
}

func (x *CloudBuildOptions) GetAppYamlPath() string {
	if x != nil {
		return x.AppYamlPath
	}
	return ""
}

func (x *CloudBuildOptions) GetCloudBuildTimeout() *durationpb.Duration {
	if x != nil {
		return x.CloudBuildTimeout
	}
	return nil
}

// The zip file information for a zip deployment.
type ZipInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// URL of the zip file to deploy from. Must be a URL to a resource in
	// Google Cloud Storage in the form
	// 'http(s)://storage.googleapis.com/\<bucket\>/\<object\>'.
	SourceUrl string `protobuf:"bytes,3,opt,name=source_url,json=sourceUrl,proto3" json:"source_url,omitempty"`
	// An estimate of the number of files in a zip for a zip deployment.
	// If set, must be greater than or equal to the actual number of files.
	// Used for optimizing performance; if not provided, deployment may be slow.
	FilesCount int32 `protobuf:"varint,4,opt,name=files_count,json=filesCount,proto3" json:"files_count,omitempty"`
}

func (x *ZipInfo) Reset() {
	*x = ZipInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_appengine_v1beta_deploy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZipInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZipInfo) ProtoMessage() {}

func (x *ZipInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_appengine_v1beta_deploy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZipInfo.ProtoReflect.Descriptor instead.
func (*ZipInfo) Descriptor() ([]byte, []int) {
	return file_google_appengine_v1beta_deploy_proto_rawDescGZIP(), []int{5}
}

func (x *ZipInfo) GetSourceUrl() string {
	if x != nil {
		return x.SourceUrl
	}
	return ""
}

func (x *ZipInfo) GetFilesCount() int32 {
	if x != nil {
		return x.FilesCount
	}
	return 0
}

var File_google_appengine_v1beta_deploy_proto protoreflect.FileDescriptor

var file_google_appengine_v1beta_deploy_proto_rawDesc = []byte{
	0x0a, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x03,
	0x0a, 0x0a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x44, 0x0a, 0x05,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x12, 0x44, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x03, 0x7a, 0x69, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e,
	0x5a, 0x69, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x7a, 0x69, 0x70, 0x12, 0x38, 0x0a, 0x05,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x5a, 0x0a, 0x13, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x11, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x1a, 0x5b, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x37, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x61, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x68,
	0x61, 0x31, 0x5f, 0x73, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x68,
	0x61, 0x31, 0x53, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x25, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x31, 0x0a, 0x09, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x22, 0x82, 0x01, 0x0a,
	0x11, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x61, 0x70, 0x70, 0x5f, 0x79, 0x61, 0x6d, 0x6c, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x59, 0x61,
	0x6d, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x12, 0x49, 0x0a, 0x13, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x22, 0x49, 0x0a, 0x07, 0x5a, 0x69, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0xd1, 0x01, 0x0a,
	0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x42, 0x0b, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x3b, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0xaa, 0x02, 0x1d,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x70, 0x70,
	0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0xca, 0x02, 0x1d,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x70, 0x70,
	0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0xea, 0x02, 0x20,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41,
	0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_appengine_v1beta_deploy_proto_rawDescOnce sync.Once
	file_google_appengine_v1beta_deploy_proto_rawDescData = file_google_appengine_v1beta_deploy_proto_rawDesc
)

func file_google_appengine_v1beta_deploy_proto_rawDescGZIP() []byte {
	file_google_appengine_v1beta_deploy_proto_rawDescOnce.Do(func() {
		file_google_appengine_v1beta_deploy_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_appengine_v1beta_deploy_proto_rawDescData)
	})
	return file_google_appengine_v1beta_deploy_proto_rawDescData
}

var file_google_appengine_v1beta_deploy_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_google_appengine_v1beta_deploy_proto_goTypes = []interface{}{
	(*Deployment)(nil),          // 0: google.appengine.v1beta.Deployment
	(*FileInfo)(nil),            // 1: google.appengine.v1beta.FileInfo
	(*ContainerInfo)(nil),       // 2: google.appengine.v1beta.ContainerInfo
	(*BuildInfo)(nil),           // 3: google.appengine.v1beta.BuildInfo
	(*CloudBuildOptions)(nil),   // 4: google.appengine.v1beta.CloudBuildOptions
	(*ZipInfo)(nil),             // 5: google.appengine.v1beta.ZipInfo
	nil,                         // 6: google.appengine.v1beta.Deployment.FilesEntry
	(*durationpb.Duration)(nil), // 7: google.protobuf.Duration
}
var file_google_appengine_v1beta_deploy_proto_depIdxs = []int32{
	6, // 0: google.appengine.v1beta.Deployment.files:type_name -> google.appengine.v1beta.Deployment.FilesEntry
	2, // 1: google.appengine.v1beta.Deployment.container:type_name -> google.appengine.v1beta.ContainerInfo
	5, // 2: google.appengine.v1beta.Deployment.zip:type_name -> google.appengine.v1beta.ZipInfo
	3, // 3: google.appengine.v1beta.Deployment.build:type_name -> google.appengine.v1beta.BuildInfo
	4, // 4: google.appengine.v1beta.Deployment.cloud_build_options:type_name -> google.appengine.v1beta.CloudBuildOptions
	7, // 5: google.appengine.v1beta.CloudBuildOptions.cloud_build_timeout:type_name -> google.protobuf.Duration
	1, // 6: google.appengine.v1beta.Deployment.FilesEntry.value:type_name -> google.appengine.v1beta.FileInfo
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_appengine_v1beta_deploy_proto_init() }
func file_google_appengine_v1beta_deploy_proto_init() {
	if File_google_appengine_v1beta_deploy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_appengine_v1beta_deploy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deployment); i {
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
		file_google_appengine_v1beta_deploy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfo); i {
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
		file_google_appengine_v1beta_deploy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerInfo); i {
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
		file_google_appengine_v1beta_deploy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildInfo); i {
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
		file_google_appengine_v1beta_deploy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudBuildOptions); i {
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
		file_google_appengine_v1beta_deploy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZipInfo); i {
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
			RawDescriptor: file_google_appengine_v1beta_deploy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_appengine_v1beta_deploy_proto_goTypes,
		DependencyIndexes: file_google_appengine_v1beta_deploy_proto_depIdxs,
		MessageInfos:      file_google_appengine_v1beta_deploy_proto_msgTypes,
	}.Build()
	File_google_appengine_v1beta_deploy_proto = out.File
	file_google_appengine_v1beta_deploy_proto_rawDesc = nil
	file_google_appengine_v1beta_deploy_proto_goTypes = nil
	file_google_appengine_v1beta_deploy_proto_depIdxs = nil
}
