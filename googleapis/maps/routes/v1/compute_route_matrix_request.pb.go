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
// source: google/maps/routes/v1/compute_route_matrix_request.proto

package routes

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ComputeRouteMatrix request message
type ComputeRouteMatrixRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Array of origins, which determines the rows of the response matrix.
	// Several size restrictions apply to the cardinality of origins and
	// destinations:
	//
	// * The number of elements (origins × destinations) must be no greater than
	// 625 in any case.
	// * The number of elements (origins × destinations) must be no greater than
	// 100 if routing_preference is set to `TRAFFIC_AWARE_OPTIMAL`.
	// * The number of waypoints (origins + destinations) specified as `place_id`
	// must be no greater than 50.
	Origins []*RouteMatrixOrigin `protobuf:"bytes,1,rep,name=origins,proto3" json:"origins,omitempty"`
	// Required. Array of destinations, which determines the columns of the response matrix.
	Destinations []*RouteMatrixDestination `protobuf:"bytes,2,rep,name=destinations,proto3" json:"destinations,omitempty"`
	// Optional. Specifies the mode of transportation.
	TravelMode RouteTravelMode `protobuf:"varint,3,opt,name=travel_mode,json=travelMode,proto3,enum=google.maps.routes.v1.RouteTravelMode" json:"travel_mode,omitempty"`
	// Optional. Specifies how to compute the route. The server attempts to use the selected
	// routing preference to compute the route. If the routing preference results
	// in an error or an extra long latency, an error is returned. In the future,
	// we might implement a fallback mechanism to use a different option when the
	// preferred option does not give a valid result. You can specify this option
	// only when the `travel_mode` is `DRIVE` or `TWO_WHEELER`, otherwise the
	// request fails.
	RoutingPreference RoutingPreference `protobuf:"varint,4,opt,name=routing_preference,json=routingPreference,proto3,enum=google.maps.routes.v1.RoutingPreference" json:"routing_preference,omitempty"`
	// Optional. The departure time. If you don't set this value, this defaults to the time
	// that you made the request. If you set this value to a time that has already
	// occurred, the request fails.
	DepartureTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=departure_time,json=departureTime,proto3" json:"departure_time,omitempty"`
}

func (x *ComputeRouteMatrixRequest) Reset() {
	*x = ComputeRouteMatrixRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_routes_v1_compute_route_matrix_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeRouteMatrixRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeRouteMatrixRequest) ProtoMessage() {}

func (x *ComputeRouteMatrixRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_routes_v1_compute_route_matrix_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeRouteMatrixRequest.ProtoReflect.Descriptor instead.
func (*ComputeRouteMatrixRequest) Descriptor() ([]byte, []int) {
	return file_google_maps_routes_v1_compute_route_matrix_request_proto_rawDescGZIP(), []int{0}
}

func (x *ComputeRouteMatrixRequest) GetOrigins() []*RouteMatrixOrigin {
	if x != nil {
		return x.Origins
	}
	return nil
}

func (x *ComputeRouteMatrixRequest) GetDestinations() []*RouteMatrixDestination {
	if x != nil {
		return x.Destinations
	}
	return nil
}

func (x *ComputeRouteMatrixRequest) GetTravelMode() RouteTravelMode {
	if x != nil {
		return x.TravelMode
	}
	return RouteTravelMode_TRAVEL_MODE_UNSPECIFIED
}

func (x *ComputeRouteMatrixRequest) GetRoutingPreference() RoutingPreference {
	if x != nil {
		return x.RoutingPreference
	}
	return RoutingPreference_ROUTING_PREFERENCE_UNSPECIFIED
}

func (x *ComputeRouteMatrixRequest) GetDepartureTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DepartureTime
	}
	return nil
}

// A single origin for ComputeRouteMatrixRequest
type RouteMatrixOrigin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Origin waypoint
	Waypoint *Waypoint `protobuf:"bytes,1,opt,name=waypoint,proto3" json:"waypoint,omitempty"`
	// Optional. Modifiers for every route that takes this as the origin
	RouteModifiers *RouteModifiers `protobuf:"bytes,2,opt,name=route_modifiers,json=routeModifiers,proto3" json:"route_modifiers,omitempty"`
}

func (x *RouteMatrixOrigin) Reset() {
	*x = RouteMatrixOrigin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_routes_v1_compute_route_matrix_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteMatrixOrigin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteMatrixOrigin) ProtoMessage() {}

func (x *RouteMatrixOrigin) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_routes_v1_compute_route_matrix_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteMatrixOrigin.ProtoReflect.Descriptor instead.
func (*RouteMatrixOrigin) Descriptor() ([]byte, []int) {
	return file_google_maps_routes_v1_compute_route_matrix_request_proto_rawDescGZIP(), []int{1}
}

func (x *RouteMatrixOrigin) GetWaypoint() *Waypoint {
	if x != nil {
		return x.Waypoint
	}
	return nil
}

func (x *RouteMatrixOrigin) GetRouteModifiers() *RouteModifiers {
	if x != nil {
		return x.RouteModifiers
	}
	return nil
}

// A single destination for ComputeRouteMatrixRequest
type RouteMatrixDestination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Destination waypoint
	Waypoint *Waypoint `protobuf:"bytes,1,opt,name=waypoint,proto3" json:"waypoint,omitempty"`
}

func (x *RouteMatrixDestination) Reset() {
	*x = RouteMatrixDestination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_routes_v1_compute_route_matrix_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteMatrixDestination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteMatrixDestination) ProtoMessage() {}

func (x *RouteMatrixDestination) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_routes_v1_compute_route_matrix_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteMatrixDestination.ProtoReflect.Descriptor instead.
func (*RouteMatrixDestination) Descriptor() ([]byte, []int) {
	return file_google_maps_routes_v1_compute_route_matrix_request_proto_rawDescGZIP(), []int{2}
}

func (x *RouteMatrixDestination) GetWaypoint() *Waypoint {
	if x != nil {
		return x.Waypoint
	}
	return nil
}

var File_google_maps_routes_v1_compute_route_matrix_request_proto protoreflect.FileDescriptor

var file_google_maps_routes_v1_compute_route_matrix_request_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d,
	0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x61,
	0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x03,
	0x0a, 0x19, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4d, 0x61,
	0x74, 0x72, 0x69, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x07, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78,
	0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x07, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x73, 0x12, 0x56, 0x0a, 0x0c, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x44, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0c,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4c, 0x0a, 0x0b,
	0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54,
	0x72, 0x61, 0x76, 0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0a,
	0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x5c, 0x0a, 0x12, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x11, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0e, 0x64, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41,
	0x01, 0x52, 0x0d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0xaa, 0x01, 0x0a, 0x11, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78,
	0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x40, 0x0a, 0x08, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x57, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08,
	0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x53, 0x0a, 0x0f, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0e, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x22, 0x5a, 0x0a,
	0x16, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x44, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x08, 0x77, 0x61, 0x79, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x57, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x08, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0xb4, 0x01, 0x0a, 0x19, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x1e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x04, 0x47, 0x4d, 0x52, 0x53,
	0xaa, 0x02, 0x15, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x2e, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x15, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x4d, 0x61, 0x70, 0x73, 0x5c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x5c, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_maps_routes_v1_compute_route_matrix_request_proto_rawDescOnce sync.Once
	file_google_maps_routes_v1_compute_route_matrix_request_proto_rawDescData = file_google_maps_routes_v1_compute_route_matrix_request_proto_rawDesc
)

func file_google_maps_routes_v1_compute_route_matrix_request_proto_rawDescGZIP() []byte {
	file_google_maps_routes_v1_compute_route_matrix_request_proto_rawDescOnce.Do(func() {
		file_google_maps_routes_v1_compute_route_matrix_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_routes_v1_compute_route_matrix_request_proto_rawDescData)
	})
	return file_google_maps_routes_v1_compute_route_matrix_request_proto_rawDescData
}

var file_google_maps_routes_v1_compute_route_matrix_request_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_maps_routes_v1_compute_route_matrix_request_proto_goTypes = []interface{}{
	(*ComputeRouteMatrixRequest)(nil), // 0: google.maps.routes.v1.ComputeRouteMatrixRequest
	(*RouteMatrixOrigin)(nil),         // 1: google.maps.routes.v1.RouteMatrixOrigin
	(*RouteMatrixDestination)(nil),    // 2: google.maps.routes.v1.RouteMatrixDestination
	(RouteTravelMode)(0),              // 3: google.maps.routes.v1.RouteTravelMode
	(RoutingPreference)(0),            // 4: google.maps.routes.v1.RoutingPreference
	(*timestamppb.Timestamp)(nil),     // 5: google.protobuf.Timestamp
	(*Waypoint)(nil),                  // 6: google.maps.routes.v1.Waypoint
	(*RouteModifiers)(nil),            // 7: google.maps.routes.v1.RouteModifiers
}
var file_google_maps_routes_v1_compute_route_matrix_request_proto_depIdxs = []int32{
	1, // 0: google.maps.routes.v1.ComputeRouteMatrixRequest.origins:type_name -> google.maps.routes.v1.RouteMatrixOrigin
	2, // 1: google.maps.routes.v1.ComputeRouteMatrixRequest.destinations:type_name -> google.maps.routes.v1.RouteMatrixDestination
	3, // 2: google.maps.routes.v1.ComputeRouteMatrixRequest.travel_mode:type_name -> google.maps.routes.v1.RouteTravelMode
	4, // 3: google.maps.routes.v1.ComputeRouteMatrixRequest.routing_preference:type_name -> google.maps.routes.v1.RoutingPreference
	5, // 4: google.maps.routes.v1.ComputeRouteMatrixRequest.departure_time:type_name -> google.protobuf.Timestamp
	6, // 5: google.maps.routes.v1.RouteMatrixOrigin.waypoint:type_name -> google.maps.routes.v1.Waypoint
	7, // 6: google.maps.routes.v1.RouteMatrixOrigin.route_modifiers:type_name -> google.maps.routes.v1.RouteModifiers
	6, // 7: google.maps.routes.v1.RouteMatrixDestination.waypoint:type_name -> google.maps.routes.v1.Waypoint
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_google_maps_routes_v1_compute_route_matrix_request_proto_init() }
func file_google_maps_routes_v1_compute_route_matrix_request_proto_init() {
	if File_google_maps_routes_v1_compute_route_matrix_request_proto != nil {
		return
	}
	file_google_maps_routes_v1_compute_routes_request_proto_init()
	file_google_maps_routes_v1_waypoint_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_maps_routes_v1_compute_route_matrix_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeRouteMatrixRequest); i {
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
		file_google_maps_routes_v1_compute_route_matrix_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteMatrixOrigin); i {
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
		file_google_maps_routes_v1_compute_route_matrix_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteMatrixDestination); i {
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
			RawDescriptor: file_google_maps_routes_v1_compute_route_matrix_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_routes_v1_compute_route_matrix_request_proto_goTypes,
		DependencyIndexes: file_google_maps_routes_v1_compute_route_matrix_request_proto_depIdxs,
		MessageInfos:      file_google_maps_routes_v1_compute_route_matrix_request_proto_msgTypes,
	}.Build()
	File_google_maps_routes_v1_compute_route_matrix_request_proto = out.File
	file_google_maps_routes_v1_compute_route_matrix_request_proto_rawDesc = nil
	file_google_maps_routes_v1_compute_route_matrix_request_proto_goTypes = nil
	file_google_maps_routes_v1_compute_route_matrix_request_proto_depIdxs = nil
}
