// Copyright 2025 Google LLC
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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: mockgcp/cloud/resourcemanager/v3/tag_holds.proto

package resourcemanagerpb

import (
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	_ "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A TagHold represents the use of a TagValue that is not captured by
// TagBindings. If a TagValue has any TagHolds, deletion will be blocked.
// This resource is intended to be created in the same cloud location as the
// `holder`.
type TagHold struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of a TagHold. This is a String of the form:
	// `tagValues/{tag-value-id}/tagHolds/{tag-hold-id}`
	// (e.g. `tagValues/123/tagHolds/456`). This resource name is generated by
	// the server.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The name of the resource where the TagValue is being used. Must
	// be less than 200 characters. E.g.
	// `//compute.googleapis.com/compute/projects/myproject/regions/us-east-1/instanceGroupManagers/instance-group`
	Holder string `protobuf:"bytes,2,opt,name=holder,proto3" json:"holder,omitempty"`
	// Optional. An optional string representing the origin of this request. This
	// field should include human-understandable information to distinguish
	// origins from each other. Must be less than 200 characters. E.g.
	// `migs-35678234`
	Origin string `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	// Optional. A URL where an end user can learn more about removing this hold.
	// E.g.
	// `https://cloud.google.com/resource-manager/docs/tags/tags-creating-and-managing`
	HelpLink string `protobuf:"bytes,4,opt,name=help_link,json=helpLink,proto3" json:"help_link,omitempty"`
	// Output only. The time this TagHold was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *TagHold) Reset() {
	*x = TagHold{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagHold) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagHold) ProtoMessage() {}

func (x *TagHold) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagHold.ProtoReflect.Descriptor instead.
func (*TagHold) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_rawDescGZIP(), []int{0}
}

func (x *TagHold) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TagHold) GetHolder() string {
	if x != nil {
		return x.Holder
	}
	return ""
}

func (x *TagHold) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *TagHold) GetHelpLink() string {
	if x != nil {
		return x.HelpLink
	}
	return ""
}

func (x *TagHold) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

// The request message to create a TagHold.
type CreateTagHoldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the TagHold's parent TagValue. Must be of
	// the form: `tagValues/{tag-value-id}`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The TagHold to be created.
	TagHold *TagHold `protobuf:"bytes,2,opt,name=tag_hold,json=tagHold,proto3" json:"tag_hold,omitempty"`
	// Optional. Set to true to perform the validations necessary for creating the
	// resource, but not actually perform the action.
	ValidateOnly bool `protobuf:"varint,3,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
}

func (x *CreateTagHoldRequest) Reset() {
	*x = CreateTagHoldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTagHoldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTagHoldRequest) ProtoMessage() {}

func (x *CreateTagHoldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTagHoldRequest.ProtoReflect.Descriptor instead.
func (*CreateTagHoldRequest) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTagHoldRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateTagHoldRequest) GetTagHold() *TagHold {
	if x != nil {
		return x.TagHold
	}
	return nil
}

func (x *CreateTagHoldRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

// Runtime operation information for creating a TagHold.
// (-- The metadata is currently empty, but may include information in the
// future. --)
type CreateTagHoldMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateTagHoldMetadata) Reset() {
	*x = CreateTagHoldMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTagHoldMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTagHoldMetadata) ProtoMessage() {}

func (x *CreateTagHoldMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTagHoldMetadata.ProtoReflect.Descriptor instead.
func (*CreateTagHoldMetadata) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_rawDescGZIP(), []int{2}
}

// The request message to delete a TagHold.
type DeleteTagHoldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the TagHold to delete. Must be of the form:
	// `tagValues/{tag-value-id}/tagHolds/{tag-hold-id}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. Set to true to perform the validations necessary for deleting the
	// resource, but not actually perform the action.
	ValidateOnly bool `protobuf:"varint,2,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
}

func (x *DeleteTagHoldRequest) Reset() {
	*x = DeleteTagHoldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTagHoldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTagHoldRequest) ProtoMessage() {}

func (x *DeleteTagHoldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTagHoldRequest.ProtoReflect.Descriptor instead.
func (*DeleteTagHoldRequest) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteTagHoldRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeleteTagHoldRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

// Runtime operation information for deleting a TagHold.
// (-- The metadata is currently empty, but may include information in the
// future. --)
type DeleteTagHoldMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTagHoldMetadata) Reset() {
	*x = DeleteTagHoldMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTagHoldMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTagHoldMetadata) ProtoMessage() {}

func (x *DeleteTagHoldMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTagHoldMetadata.ProtoReflect.Descriptor instead.
func (*DeleteTagHoldMetadata) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_rawDescGZIP(), []int{4}
}

// The request message for listing the TagHolds under a TagValue.
type ListTagHoldsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the parent TagValue. Must be of the form:
	// `tagValues/{tag-value-id}`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional. The maximum number of TagHolds to return in the response. The
	// server allows a maximum of 300 TagHolds to return. If unspecified, the
	// server will use 100 as the default.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. A pagination token returned from a previous call to
	// `ListTagHolds` that indicates where this listing should continue from.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Optional. Criteria used to select a subset of TagHolds parented by the
	// TagValue to return. This field follows the syntax defined by aip.dev/160;
	// the `holder` and `origin` fields are supported for filtering. Currently
	// only `AND` syntax is supported. Some example queries are:
	//
	//   - `holder =
	//     //compute.googleapis.com/compute/projects/myproject/regions/us-east-1/instanceGroupManagers/instance-group`
	//   - `origin = 35678234`
	//   - `holder =
	//     //compute.googleapis.com/compute/projects/myproject/regions/us-east-1/instanceGroupManagers/instance-group
	//     AND origin = 35678234`
	Filter string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListTagHoldsRequest) Reset() {
	*x = ListTagHoldsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTagHoldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTagHoldsRequest) ProtoMessage() {}

func (x *ListTagHoldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTagHoldsRequest.ProtoReflect.Descriptor instead.
func (*ListTagHoldsRequest) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_rawDescGZIP(), []int{5}
}

func (x *ListTagHoldsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListTagHoldsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListTagHoldsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListTagHoldsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// The ListTagHolds response.
type ListTagHoldsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A possibly paginated list of TagHolds.
	TagHolds []*TagHold `protobuf:"bytes,1,rep,name=tag_holds,json=tagHolds,proto3" json:"tag_holds,omitempty"`
	// Pagination token.
	//
	// If the result set is too large to fit in a single response, this token
	// is returned. It encodes the position of the current result cursor.
	// Feeding this value into a new list request with the `page_token` parameter
	// gives the next page of the results.
	//
	// When `next_page_token` is not filled in, there is no next page and
	// the list returned is the last page in the result set.
	//
	// Pagination tokens have a limited lifetime.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListTagHoldsResponse) Reset() {
	*x = ListTagHoldsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTagHoldsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTagHoldsResponse) ProtoMessage() {}

func (x *ListTagHoldsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTagHoldsResponse.ProtoReflect.Descriptor instead.
func (*ListTagHoldsResponse) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_rawDescGZIP(), []int{6}
}

func (x *ListTagHoldsResponse) GetTagHolds() []*TagHold {
	if x != nil {
		return x.TagHolds
	}
	return nil
}

func (x *ListTagHoldsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_mockgcp_cloud_resourcemanager_v3_tag_holds_proto protoreflect.FileDescriptor

var file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_rawDesc = []byte{
	0x0a, 0x30, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f,
	0x76, 0x33, 0x2f, 0x74, 0x61, 0x67, 0x5f, 0x68, 0x6f, 0x6c, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x20, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x76, 0x33, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x02, 0x0a, 0x07, 0x54,
	0x61, 0x67, 0x48, 0x6f, 0x6c, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x06, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x06,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x01, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x09, 0x68, 0x65, 0x6c,
	0x70, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x01, 0x52, 0x08, 0x68, 0x65, 0x6c, 0x70, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x40, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x3a, 0x5b, 0xea,
	0x41, 0x58, 0x0a, 0x2b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x61, 0x67, 0x48, 0x6f, 0x6c, 0x64, 0x12,
	0x29, 0x74, 0x61, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2f, 0x7b, 0x74, 0x61, 0x67, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x7d, 0x2f, 0x74, 0x61, 0x67, 0x48, 0x6f, 0x6c, 0x64, 0x73, 0x2f,
	0x7b, 0x74, 0x61, 0x67, 0x5f, 0x68, 0x6f, 0x6c, 0x64, 0x7d, 0x22, 0xd8, 0x01, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x48, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x33, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2d, 0x12, 0x2b, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x54, 0x61, 0x67, 0x48, 0x6f, 0x6c, 0x64, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x12, 0x49, 0x0a, 0x08, 0x74, 0x61, 0x67, 0x5f, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x61, 0x67, 0x48, 0x6f, 0x6c, 0x64, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x07, 0x74, 0x61, 0x67, 0x48, 0x6f, 0x6c, 0x64, 0x12, 0x28, 0x0a, 0x0d, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0x17, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x67, 0x48, 0x6f, 0x6c, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x89,
	0x01, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x67, 0x48, 0x6f, 0x6c, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x33, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2d, 0x0a, 0x2b, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x54, 0x61, 0x67, 0x48, 0x6f, 0x6c, 0x64, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x28, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x6e, 0x6c,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0c, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0x17, 0x0a, 0x15, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x67, 0x48, 0x6f, 0x6c, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x22, 0xc5, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x48,
	0x6f, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x33, 0xe0, 0x41, 0x02,
	0xfa, 0x41, 0x2d, 0x12, 0x2b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x61, 0x67, 0x48, 0x6f, 0x6c, 0x64,
	0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x01, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x01, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x86, 0x01, 0x0a, 0x14,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x48, 0x6f, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x09, 0x74, 0x61, 0x67, 0x5f, 0x68, 0x6f, 0x6c, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63,
	0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x61, 0x67, 0x48, 0x6f,
	0x6c, 0x64, 0x52, 0x08, 0x74, 0x61, 0x67, 0x48, 0x6f, 0x6c, 0x64, 0x73, 0x12, 0x26, 0x0a, 0x0f,
	0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xf0, 0x05, 0x0a, 0x08, 0x54, 0x61, 0x67, 0x48, 0x6f, 0x6c, 0x64,
	0x73, 0x12, 0xd0, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x48,
	0x6f, 0x6c, 0x64, 0x12, 0x36, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67,
	0x48, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x68, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x2d, 0x22, 0x21, 0x2f, 0x76, 0x33, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d,
	0x74, 0x61, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x74, 0x61, 0x67,
	0x48, 0x6f, 0x6c, 0x64, 0x73, 0x3a, 0x08, 0x74, 0x61, 0x67, 0x5f, 0x68, 0x6f, 0x6c, 0x64, 0xda,
	0x41, 0x0f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x2c, 0x74, 0x61, 0x67, 0x5f, 0x68, 0x6f, 0x6c,
	0x64, 0xca, 0x41, 0x20, 0x0a, 0x07, 0x54, 0x61, 0x67, 0x48, 0x6f, 0x6c, 0x64, 0x12, 0x15, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x48, 0x6f, 0x6c, 0x64, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0xc9, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x61, 0x67, 0x48, 0x6f, 0x6c, 0x64, 0x12, 0x36, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x61, 0x67, 0x48, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x61, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x23, 0x2a, 0x21, 0x2f, 0x76, 0x33, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65,
	0x3d, 0x74, 0x61, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x74, 0x61, 0x67,
	0x48, 0x6f, 0x6c, 0x64, 0x73, 0x2f, 0x2a, 0x7d, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xca,
	0x41, 0x2e, 0x0a, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x61, 0x67, 0x48, 0x6f, 0x6c, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0xb1, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x48, 0x6f, 0x6c, 0x64,
	0x73, 0x12, 0x35, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x48, 0x6f, 0x6c, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67,
	0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x61, 0x67, 0x48, 0x6f, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f, 0x76, 0x33, 0x2f, 0x7b, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x74, 0x61, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2f,
	0x2a, 0x7d, 0x2f, 0x74, 0x61, 0x67, 0x48, 0x6f, 0x6c, 0x64, 0x73, 0xda, 0x41, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x1a, 0x90, 0x01, 0xca, 0x41, 0x23, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x67,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2c, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x72, 0x65,
	0x61, 0x64, 0x2d, 0x6f, 0x6e, 0x6c, 0x79, 0x42, 0xef, 0x01, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e,
	0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x33,
	0x42, 0x0d, 0x54, 0x61, 0x67, 0x48, 0x6f, 0x6c, 0x64, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x33, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x70, 0x62,
	0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x56, 0x33, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x5c, 0x56, 0x33, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_rawDescOnce sync.Once
	file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_rawDescData = file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_rawDesc
)

func file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_rawDescGZIP() []byte {
	file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_rawDescOnce.Do(func() {
		file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_rawDescData = protoimpl.X.CompressGZIP(file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_rawDescData)
	})
	return file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_rawDescData
}

var file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_goTypes = []interface{}{
	(*TagHold)(nil),                 // 0: mockgcp.cloud.resourcemanager.v3.TagHold
	(*CreateTagHoldRequest)(nil),    // 1: mockgcp.cloud.resourcemanager.v3.CreateTagHoldRequest
	(*CreateTagHoldMetadata)(nil),   // 2: mockgcp.cloud.resourcemanager.v3.CreateTagHoldMetadata
	(*DeleteTagHoldRequest)(nil),    // 3: mockgcp.cloud.resourcemanager.v3.DeleteTagHoldRequest
	(*DeleteTagHoldMetadata)(nil),   // 4: mockgcp.cloud.resourcemanager.v3.DeleteTagHoldMetadata
	(*ListTagHoldsRequest)(nil),     // 5: mockgcp.cloud.resourcemanager.v3.ListTagHoldsRequest
	(*ListTagHoldsResponse)(nil),    // 6: mockgcp.cloud.resourcemanager.v3.ListTagHoldsResponse
	(*timestamp.Timestamp)(nil),     // 7: google.protobuf.Timestamp
	(*longrunningpb.Operation)(nil), // 8: google.longrunning.Operation
}
var file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_depIdxs = []int32{
	7, // 0: mockgcp.cloud.resourcemanager.v3.TagHold.create_time:type_name -> google.protobuf.Timestamp
	0, // 1: mockgcp.cloud.resourcemanager.v3.CreateTagHoldRequest.tag_hold:type_name -> mockgcp.cloud.resourcemanager.v3.TagHold
	0, // 2: mockgcp.cloud.resourcemanager.v3.ListTagHoldsResponse.tag_holds:type_name -> mockgcp.cloud.resourcemanager.v3.TagHold
	1, // 3: mockgcp.cloud.resourcemanager.v3.TagHolds.CreateTagHold:input_type -> mockgcp.cloud.resourcemanager.v3.CreateTagHoldRequest
	3, // 4: mockgcp.cloud.resourcemanager.v3.TagHolds.DeleteTagHold:input_type -> mockgcp.cloud.resourcemanager.v3.DeleteTagHoldRequest
	5, // 5: mockgcp.cloud.resourcemanager.v3.TagHolds.ListTagHolds:input_type -> mockgcp.cloud.resourcemanager.v3.ListTagHoldsRequest
	8, // 6: mockgcp.cloud.resourcemanager.v3.TagHolds.CreateTagHold:output_type -> google.longrunning.Operation
	8, // 7: mockgcp.cloud.resourcemanager.v3.TagHolds.DeleteTagHold:output_type -> google.longrunning.Operation
	6, // 8: mockgcp.cloud.resourcemanager.v3.TagHolds.ListTagHolds:output_type -> mockgcp.cloud.resourcemanager.v3.ListTagHoldsResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_init() }
func file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_init() {
	if File_mockgcp_cloud_resourcemanager_v3_tag_holds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagHold); i {
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
		file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTagHoldRequest); i {
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
		file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTagHoldMetadata); i {
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
		file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTagHoldRequest); i {
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
		file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTagHoldMetadata); i {
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
		file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTagHoldsRequest); i {
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
		file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTagHoldsResponse); i {
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
			RawDescriptor: file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_goTypes,
		DependencyIndexes: file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_depIdxs,
		MessageInfos:      file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_msgTypes,
	}.Build()
	File_mockgcp_cloud_resourcemanager_v3_tag_holds_proto = out.File
	file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_rawDesc = nil
	file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_goTypes = nil
	file_mockgcp_cloud_resourcemanager_v3_tag_holds_proto_depIdxs = nil
}
