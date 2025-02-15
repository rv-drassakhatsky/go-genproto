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
// source: google/maps/routes/v1alpha/route_service.proto

package routes

import (
	context "context"
	reflect "reflect"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	v1 "google.golang.org/genproto/googleapis/maps/routes/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_google_maps_routes_v1alpha_route_service_proto protoreflect.FileDescriptor

var file_google_maps_routes_v1alpha_route_service_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61,
	0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70,
	0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x5f, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x5f, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xc6, 0x04, 0x0a, 0x0b, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x12,
	0x8d, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x12, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3a, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12,
	0x9b, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x12, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x72, 0x69,
	0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x45, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22, 0x1b, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x3a, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x3a, 0x01, 0x2a, 0x30, 0x01, 0x12, 0xa5, 0x01,
	0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d,
	0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x21, 0x22, 0x1c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3a, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x73, 0x3a, 0x01, 0x2a, 0x1a, 0x61, 0xca, 0x41, 0x1e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x3d, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2d,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x70,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x42, 0xc1, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x17, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x61,
	0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x3b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x04, 0x47, 0x4d,
	0x52, 0x53, 0xaa, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x73,
	0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x56, 0x31, 0x41, 0x6c, 0x70, 0x68, 0x61, 0xca,
	0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x4d, 0x61, 0x70, 0x73, 0x5c, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_google_maps_routes_v1alpha_route_service_proto_goTypes = []interface{}{
	(*v1.ComputeRoutesRequest)(nil),        // 0: google.maps.routes.v1.ComputeRoutesRequest
	(*v1.ComputeRouteMatrixRequest)(nil),   // 1: google.maps.routes.v1.ComputeRouteMatrixRequest
	(*v1.ComputeCustomRoutesRequest)(nil),  // 2: google.maps.routes.v1.ComputeCustomRoutesRequest
	(*v1.ComputeRoutesResponse)(nil),       // 3: google.maps.routes.v1.ComputeRoutesResponse
	(*v1.RouteMatrixElement)(nil),          // 4: google.maps.routes.v1.RouteMatrixElement
	(*v1.ComputeCustomRoutesResponse)(nil), // 5: google.maps.routes.v1.ComputeCustomRoutesResponse
}
var file_google_maps_routes_v1alpha_route_service_proto_depIdxs = []int32{
	0, // 0: google.maps.routes.v1alpha.RoutesAlpha.ComputeRoutes:input_type -> google.maps.routes.v1.ComputeRoutesRequest
	1, // 1: google.maps.routes.v1alpha.RoutesAlpha.ComputeRouteMatrix:input_type -> google.maps.routes.v1.ComputeRouteMatrixRequest
	2, // 2: google.maps.routes.v1alpha.RoutesAlpha.ComputeCustomRoutes:input_type -> google.maps.routes.v1.ComputeCustomRoutesRequest
	3, // 3: google.maps.routes.v1alpha.RoutesAlpha.ComputeRoutes:output_type -> google.maps.routes.v1.ComputeRoutesResponse
	4, // 4: google.maps.routes.v1alpha.RoutesAlpha.ComputeRouteMatrix:output_type -> google.maps.routes.v1.RouteMatrixElement
	5, // 5: google.maps.routes.v1alpha.RoutesAlpha.ComputeCustomRoutes:output_type -> google.maps.routes.v1.ComputeCustomRoutesResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_maps_routes_v1alpha_route_service_proto_init() }
func file_google_maps_routes_v1alpha_route_service_proto_init() {
	if File_google_maps_routes_v1alpha_route_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_maps_routes_v1alpha_route_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_maps_routes_v1alpha_route_service_proto_goTypes,
		DependencyIndexes: file_google_maps_routes_v1alpha_route_service_proto_depIdxs,
	}.Build()
	File_google_maps_routes_v1alpha_route_service_proto = out.File
	file_google_maps_routes_v1alpha_route_service_proto_rawDesc = nil
	file_google_maps_routes_v1alpha_route_service_proto_goTypes = nil
	file_google_maps_routes_v1alpha_route_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RoutesAlphaClient is the client API for RoutesAlpha service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RoutesAlphaClient interface {
	// Returns the primary route along with optional alternate routes, given a set
	// of terminal and intermediate waypoints.
	//
	// **NOTE:** This method requires that you specify a response field mask in
	// the input. You can provide the response field mask by using the URL
	// parameter `$fields` or `fields`, or by using the HTTP/gRPC header
	// `X-Goog-FieldMask` (see the [available URL parameters and
	// headers](https://cloud.google.com/apis/docs/system-parameters). The value
	// is a comma separated list of field paths. See this detailed documentation
	// about [how to construct the field
	// paths](https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto).
	//
	// For example, in this method:
	//
	// * Field mask of all available fields (for manual inspection):
	//   `X-Goog-FieldMask: *`
	// * Field mask of Route-level duration, distance, and polyline (an example
	// production setup):
	//   `X-Goog-FieldMask:
	//   routes.duration,routes.distanceMeters,routes.polyline.encodedPolyline`
	//
	// Google discourages the use of the wildcard (`*`) response field mask, or
	// specifying the field mask at the top level (`routes`), because:
	//
	// * Selecting only the fields that you need helps our server save computation
	// cycles, allowing us to return the result to you with a lower latency.
	// * Selecting only the fields that you need in your production job ensures
	// stable latency performance. We might add more response fields in the
	// future, and those new fields might require extra computation time. If you
	// select all fields, or if you select all fields at the top level, then you
	// might experience performance degradation because any new field we add will
	// be automatically included in the response.
	// * Selecting only the fields that you need results in a smaller response
	// size, and thus higher network throughput.
	ComputeRoutes(ctx context.Context, in *v1.ComputeRoutesRequest, opts ...grpc.CallOption) (*v1.ComputeRoutesResponse, error)
	// Takes in a list of origins and destinations and returns a stream containing
	// route information for each combination of origin and destination.
	//
	// **NOTE:** This method requires that you specify a response field mask in
	// the input. You can provide the response field mask by using the URL
	// parameter `$fields` or `fields`, or by using the HTTP/gRPC header
	// `X-Goog-FieldMask` (see the [available URL parameters and
	// headers](https://cloud.google.com/apis/docs/system-parameters). The value
	// is a comma separated list of field paths. See this detailed documentation
	// about [how to construct the field
	// paths](https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto).
	//
	// For example, in this method:
	//
	// * Field mask of all available fields (for manual inspection):
	//   `X-Goog-FieldMask: *`
	// * Field mask of route durations, distances, element status, condition, and
	//   element indices (an example production setup):
	//   `X-Goog-FieldMask:
	//   originIndex,destinationIndex,status,condition,distanceMeters,duration`
	//
	// It is critical that you include `status` in your field mask as otherwise
	// all messages will appear to be OK. Google discourages the use of the
	// wildcard (`*`) response field mask, because:
	//
	// * Selecting only the fields that you need helps our server save computation
	// cycles, allowing us to return the result to you with a lower latency.
	// * Selecting only the fields that you need in your production job ensures
	// stable latency performance. We might add more response fields in the
	// future, and those new fields might require extra computation time. If you
	// select all fields, or if you select all fields at the top level, then you
	// might experience performance degradation because any new field we add will
	// be automatically included in the response.
	// * Selecting only the fields that you need results in a smaller response
	// size, and thus higher network throughput.
	ComputeRouteMatrix(ctx context.Context, in *v1.ComputeRouteMatrixRequest, opts ...grpc.CallOption) (RoutesAlpha_ComputeRouteMatrixClient, error)
	// Given a set of terminal and intermediate waypoints, and a route objective,
	// computes the best route for the route objective. Also returns fastest route
	// and shortest route as reference routes.
	//
	// **NOTE:** This method requires that you specify a response field mask in
	// the input. You can provide the response field mask by using the URL
	// parameter `$fields` or `fields`, or by using the HTTP/gRPC header
	// `X-Goog-FieldMask` (see the [available URL parameters and
	// headers](https://cloud.google.com/apis/docs/system-parameters). The value
	// is a comma separated list of field paths. See this detailed documentation
	// about [how to construct the field
	// paths](https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto).
	//
	// For example, in this method:
	//
	// * Field mask of all available fields (for manual inspection):
	//   `X-Goog-FieldMask: *`
	// * Field mask of route distances, durations, token and toll info:
	//   `X-Goog-FieldMask:
	//   routes.route.distanceMeters,routes.route.duration,routes.token,routes.route.travelAdvisory.tollInfo`
	//
	// Google discourages the use of the wildcard (`*`) response field mask, or
	// specifying the field mask at the top level (`routes`), because:
	//
	// * Selecting only the fields that you need helps our server save computation
	// cycles, allowing us to return the result to you with a lower latency.
	// * Selecting only the fields that you need in your production job ensures
	// stable latency performance. We might add more response fields in the
	// future, and those new fields might require extra computation time. If you
	// select all fields, or if you select all fields at the top level, then you
	// might experience performance degradation because any new field we add will
	// be automatically included in the response.
	// * Selecting only the fields that you need results in a smaller response
	// size, and thus higher network throughput.
	ComputeCustomRoutes(ctx context.Context, in *v1.ComputeCustomRoutesRequest, opts ...grpc.CallOption) (*v1.ComputeCustomRoutesResponse, error)
}

type routesAlphaClient struct {
	cc grpc.ClientConnInterface
}

func NewRoutesAlphaClient(cc grpc.ClientConnInterface) RoutesAlphaClient {
	return &routesAlphaClient{cc}
}

func (c *routesAlphaClient) ComputeRoutes(ctx context.Context, in *v1.ComputeRoutesRequest, opts ...grpc.CallOption) (*v1.ComputeRoutesResponse, error) {
	out := new(v1.ComputeRoutesResponse)
	err := c.cc.Invoke(ctx, "/google.maps.routes.v1alpha.RoutesAlpha/ComputeRoutes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routesAlphaClient) ComputeRouteMatrix(ctx context.Context, in *v1.ComputeRouteMatrixRequest, opts ...grpc.CallOption) (RoutesAlpha_ComputeRouteMatrixClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RoutesAlpha_serviceDesc.Streams[0], "/google.maps.routes.v1alpha.RoutesAlpha/ComputeRouteMatrix", opts...)
	if err != nil {
		return nil, err
	}
	x := &routesAlphaComputeRouteMatrixClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RoutesAlpha_ComputeRouteMatrixClient interface {
	Recv() (*v1.RouteMatrixElement, error)
	grpc.ClientStream
}

type routesAlphaComputeRouteMatrixClient struct {
	grpc.ClientStream
}

func (x *routesAlphaComputeRouteMatrixClient) Recv() (*v1.RouteMatrixElement, error) {
	m := new(v1.RouteMatrixElement)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routesAlphaClient) ComputeCustomRoutes(ctx context.Context, in *v1.ComputeCustomRoutesRequest, opts ...grpc.CallOption) (*v1.ComputeCustomRoutesResponse, error) {
	out := new(v1.ComputeCustomRoutesResponse)
	err := c.cc.Invoke(ctx, "/google.maps.routes.v1alpha.RoutesAlpha/ComputeCustomRoutes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoutesAlphaServer is the server API for RoutesAlpha service.
type RoutesAlphaServer interface {
	// Returns the primary route along with optional alternate routes, given a set
	// of terminal and intermediate waypoints.
	//
	// **NOTE:** This method requires that you specify a response field mask in
	// the input. You can provide the response field mask by using the URL
	// parameter `$fields` or `fields`, or by using the HTTP/gRPC header
	// `X-Goog-FieldMask` (see the [available URL parameters and
	// headers](https://cloud.google.com/apis/docs/system-parameters). The value
	// is a comma separated list of field paths. See this detailed documentation
	// about [how to construct the field
	// paths](https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto).
	//
	// For example, in this method:
	//
	// * Field mask of all available fields (for manual inspection):
	//   `X-Goog-FieldMask: *`
	// * Field mask of Route-level duration, distance, and polyline (an example
	// production setup):
	//   `X-Goog-FieldMask:
	//   routes.duration,routes.distanceMeters,routes.polyline.encodedPolyline`
	//
	// Google discourages the use of the wildcard (`*`) response field mask, or
	// specifying the field mask at the top level (`routes`), because:
	//
	// * Selecting only the fields that you need helps our server save computation
	// cycles, allowing us to return the result to you with a lower latency.
	// * Selecting only the fields that you need in your production job ensures
	// stable latency performance. We might add more response fields in the
	// future, and those new fields might require extra computation time. If you
	// select all fields, or if you select all fields at the top level, then you
	// might experience performance degradation because any new field we add will
	// be automatically included in the response.
	// * Selecting only the fields that you need results in a smaller response
	// size, and thus higher network throughput.
	ComputeRoutes(context.Context, *v1.ComputeRoutesRequest) (*v1.ComputeRoutesResponse, error)
	// Takes in a list of origins and destinations and returns a stream containing
	// route information for each combination of origin and destination.
	//
	// **NOTE:** This method requires that you specify a response field mask in
	// the input. You can provide the response field mask by using the URL
	// parameter `$fields` or `fields`, or by using the HTTP/gRPC header
	// `X-Goog-FieldMask` (see the [available URL parameters and
	// headers](https://cloud.google.com/apis/docs/system-parameters). The value
	// is a comma separated list of field paths. See this detailed documentation
	// about [how to construct the field
	// paths](https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto).
	//
	// For example, in this method:
	//
	// * Field mask of all available fields (for manual inspection):
	//   `X-Goog-FieldMask: *`
	// * Field mask of route durations, distances, element status, condition, and
	//   element indices (an example production setup):
	//   `X-Goog-FieldMask:
	//   originIndex,destinationIndex,status,condition,distanceMeters,duration`
	//
	// It is critical that you include `status` in your field mask as otherwise
	// all messages will appear to be OK. Google discourages the use of the
	// wildcard (`*`) response field mask, because:
	//
	// * Selecting only the fields that you need helps our server save computation
	// cycles, allowing us to return the result to you with a lower latency.
	// * Selecting only the fields that you need in your production job ensures
	// stable latency performance. We might add more response fields in the
	// future, and those new fields might require extra computation time. If you
	// select all fields, or if you select all fields at the top level, then you
	// might experience performance degradation because any new field we add will
	// be automatically included in the response.
	// * Selecting only the fields that you need results in a smaller response
	// size, and thus higher network throughput.
	ComputeRouteMatrix(*v1.ComputeRouteMatrixRequest, RoutesAlpha_ComputeRouteMatrixServer) error
	// Given a set of terminal and intermediate waypoints, and a route objective,
	// computes the best route for the route objective. Also returns fastest route
	// and shortest route as reference routes.
	//
	// **NOTE:** This method requires that you specify a response field mask in
	// the input. You can provide the response field mask by using the URL
	// parameter `$fields` or `fields`, or by using the HTTP/gRPC header
	// `X-Goog-FieldMask` (see the [available URL parameters and
	// headers](https://cloud.google.com/apis/docs/system-parameters). The value
	// is a comma separated list of field paths. See this detailed documentation
	// about [how to construct the field
	// paths](https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto).
	//
	// For example, in this method:
	//
	// * Field mask of all available fields (for manual inspection):
	//   `X-Goog-FieldMask: *`
	// * Field mask of route distances, durations, token and toll info:
	//   `X-Goog-FieldMask:
	//   routes.route.distanceMeters,routes.route.duration,routes.token,routes.route.travelAdvisory.tollInfo`
	//
	// Google discourages the use of the wildcard (`*`) response field mask, or
	// specifying the field mask at the top level (`routes`), because:
	//
	// * Selecting only the fields that you need helps our server save computation
	// cycles, allowing us to return the result to you with a lower latency.
	// * Selecting only the fields that you need in your production job ensures
	// stable latency performance. We might add more response fields in the
	// future, and those new fields might require extra computation time. If you
	// select all fields, or if you select all fields at the top level, then you
	// might experience performance degradation because any new field we add will
	// be automatically included in the response.
	// * Selecting only the fields that you need results in a smaller response
	// size, and thus higher network throughput.
	ComputeCustomRoutes(context.Context, *v1.ComputeCustomRoutesRequest) (*v1.ComputeCustomRoutesResponse, error)
}

// UnimplementedRoutesAlphaServer can be embedded to have forward compatible implementations.
type UnimplementedRoutesAlphaServer struct {
}

func (*UnimplementedRoutesAlphaServer) ComputeRoutes(context.Context, *v1.ComputeRoutesRequest) (*v1.ComputeRoutesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComputeRoutes not implemented")
}
func (*UnimplementedRoutesAlphaServer) ComputeRouteMatrix(*v1.ComputeRouteMatrixRequest, RoutesAlpha_ComputeRouteMatrixServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeRouteMatrix not implemented")
}
func (*UnimplementedRoutesAlphaServer) ComputeCustomRoutes(context.Context, *v1.ComputeCustomRoutesRequest) (*v1.ComputeCustomRoutesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComputeCustomRoutes not implemented")
}

func RegisterRoutesAlphaServer(s *grpc.Server, srv RoutesAlphaServer) {
	s.RegisterService(&_RoutesAlpha_serviceDesc, srv)
}

func _RoutesAlpha_ComputeRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ComputeRoutesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutesAlphaServer).ComputeRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.maps.routes.v1alpha.RoutesAlpha/ComputeRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutesAlphaServer).ComputeRoutes(ctx, req.(*v1.ComputeRoutesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutesAlpha_ComputeRouteMatrix_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(v1.ComputeRouteMatrixRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RoutesAlphaServer).ComputeRouteMatrix(m, &routesAlphaComputeRouteMatrixServer{stream})
}

type RoutesAlpha_ComputeRouteMatrixServer interface {
	Send(*v1.RouteMatrixElement) error
	grpc.ServerStream
}

type routesAlphaComputeRouteMatrixServer struct {
	grpc.ServerStream
}

func (x *routesAlphaComputeRouteMatrixServer) Send(m *v1.RouteMatrixElement) error {
	return x.ServerStream.SendMsg(m)
}

func _RoutesAlpha_ComputeCustomRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ComputeCustomRoutesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutesAlphaServer).ComputeCustomRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.maps.routes.v1alpha.RoutesAlpha/ComputeCustomRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutesAlphaServer).ComputeCustomRoutes(ctx, req.(*v1.ComputeCustomRoutesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoutesAlpha_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.maps.routes.v1alpha.RoutesAlpha",
	HandlerType: (*RoutesAlphaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ComputeRoutes",
			Handler:    _RoutesAlpha_ComputeRoutes_Handler,
		},
		{
			MethodName: "ComputeCustomRoutes",
			Handler:    _RoutesAlpha_ComputeCustomRoutes_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ComputeRouteMatrix",
			Handler:       _RoutesAlpha_ComputeRouteMatrix_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "google/maps/routes/v1alpha/route_service.proto",
}
