// Copyright 2024 Google LLC
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
// source: mockgcp/cloud/aiplatform/v1/artifact.proto

package aiplatformpb

import (
	_struct "github.com/golang/protobuf/ptypes/struct"
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

// Describes the state of the Artifact.
type Artifact_State int32

const (
	// Unspecified state for the Artifact.
	Artifact_STATE_UNSPECIFIED Artifact_State = 0
	// A state used by systems like Vertex AI Pipelines to indicate that the
	// underlying data item represented by this Artifact is being created.
	Artifact_PENDING Artifact_State = 1
	// A state indicating that the Artifact should exist, unless something
	// external to the system deletes it.
	Artifact_LIVE Artifact_State = 2
)

// Enum value maps for Artifact_State.
var (
	Artifact_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "PENDING",
		2: "LIVE",
	}
	Artifact_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"PENDING":           1,
		"LIVE":              2,
	}
)

func (x Artifact_State) Enum() *Artifact_State {
	p := new(Artifact_State)
	*p = x
	return p
}

func (x Artifact_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Artifact_State) Descriptor() protoreflect.EnumDescriptor {
	return file_mockgcp_cloud_aiplatform_v1_artifact_proto_enumTypes[0].Descriptor()
}

func (Artifact_State) Type() protoreflect.EnumType {
	return &file_mockgcp_cloud_aiplatform_v1_artifact_proto_enumTypes[0]
}

func (x Artifact_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Artifact_State.Descriptor instead.
func (Artifact_State) EnumDescriptor() ([]byte, []int) {
	return file_mockgcp_cloud_aiplatform_v1_artifact_proto_rawDescGZIP(), []int{0, 0}
}

// Instance of a general artifact.
type Artifact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the Artifact.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// User provided display name of the Artifact.
	// May be up to 128 Unicode characters.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The uniform resource identifier of the artifact file.
	// May be empty if there is no actual artifact file.
	Uri string `protobuf:"bytes,6,opt,name=uri,proto3" json:"uri,omitempty"`
	// An eTag used to perform consistent read-modify-write updates. If not set, a
	// blind "overwrite" update happens.
	Etag string `protobuf:"bytes,9,opt,name=etag,proto3" json:"etag,omitempty"`
	// The labels with user-defined metadata to organize your Artifacts.
	//
	// Label keys and values can be no longer than 64 characters
	// (Unicode codepoints), can only contain lowercase letters, numeric
	// characters, underscores and dashes. International characters are allowed.
	// No more than 64 user labels can be associated with one Artifact (System
	// labels are excluded).
	Labels map[string]string `protobuf:"bytes,10,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Output only. Timestamp when this Artifact was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,11,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Timestamp when this Artifact was last updated.
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,12,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The state of this Artifact. This is a property of the Artifact, and does
	// not imply or capture any ongoing process. This property is managed by
	// clients (such as Vertex AI Pipelines), and the system does not prescribe
	// or check the validity of state transitions.
	State Artifact_State `protobuf:"varint,13,opt,name=state,proto3,enum=mockgcp.cloud.aiplatform.v1.Artifact_State" json:"state,omitempty"`
	// The title of the schema describing the metadata.
	//
	// Schema title and version is expected to be registered in earlier Create
	// Schema calls. And both are used together as unique identifiers to identify
	// schemas within the local metadata store.
	SchemaTitle string `protobuf:"bytes,14,opt,name=schema_title,json=schemaTitle,proto3" json:"schema_title,omitempty"`
	// The version of the schema in schema_name to use.
	//
	// Schema title and version is expected to be registered in earlier Create
	// Schema calls. And both are used together as unique identifiers to identify
	// schemas within the local metadata store.
	SchemaVersion string `protobuf:"bytes,15,opt,name=schema_version,json=schemaVersion,proto3" json:"schema_version,omitempty"`
	// Properties of the Artifact.
	// Top level metadata keys' heading and trailing spaces will be trimmed.
	// The size of this field should not exceed 200KB.
	Metadata *_struct.Struct `protobuf:"bytes,16,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Description of the Artifact
	Description string `protobuf:"bytes,17,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Artifact) Reset() {
	*x = Artifact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_aiplatform_v1_artifact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Artifact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Artifact) ProtoMessage() {}

func (x *Artifact) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_aiplatform_v1_artifact_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Artifact.ProtoReflect.Descriptor instead.
func (*Artifact) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_aiplatform_v1_artifact_proto_rawDescGZIP(), []int{0}
}

func (x *Artifact) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Artifact) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Artifact) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Artifact) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *Artifact) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Artifact) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Artifact) GetUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Artifact) GetState() Artifact_State {
	if x != nil {
		return x.State
	}
	return Artifact_STATE_UNSPECIFIED
}

func (x *Artifact) GetSchemaTitle() string {
	if x != nil {
		return x.SchemaTitle
	}
	return ""
}

func (x *Artifact) GetSchemaVersion() string {
	if x != nil {
		return x.SchemaVersion
	}
	return ""
}

func (x *Artifact) GetMetadata() *_struct.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Artifact) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_mockgcp_cloud_aiplatform_v1_artifact_proto protoreflect.FileDescriptor

var file_mockgcp_cloud_aiplatform_v1_artifact_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x6d, 0x6f,
	0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x06, 0x0a, 0x08, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x69, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12,
	0x12, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65,
	0x74, 0x61, 0x67, 0x12, 0x49, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x40,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x41, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2b, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x33, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x35, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x4c, 0x49, 0x56, 0x45, 0x10, 0x02, 0x3a, 0x86, 0x01, 0xea, 0x41, 0x82, 0x01, 0x0a,
	0x22, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x12, 0x5c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x7b, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7d, 0x2f, 0x61, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x7d, 0x42, 0xcc, 0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63,
	0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x69, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0x3b, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0xaa, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31,
	0xea, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mockgcp_cloud_aiplatform_v1_artifact_proto_rawDescOnce sync.Once
	file_mockgcp_cloud_aiplatform_v1_artifact_proto_rawDescData = file_mockgcp_cloud_aiplatform_v1_artifact_proto_rawDesc
)

func file_mockgcp_cloud_aiplatform_v1_artifact_proto_rawDescGZIP() []byte {
	file_mockgcp_cloud_aiplatform_v1_artifact_proto_rawDescOnce.Do(func() {
		file_mockgcp_cloud_aiplatform_v1_artifact_proto_rawDescData = protoimpl.X.CompressGZIP(file_mockgcp_cloud_aiplatform_v1_artifact_proto_rawDescData)
	})
	return file_mockgcp_cloud_aiplatform_v1_artifact_proto_rawDescData
}

var file_mockgcp_cloud_aiplatform_v1_artifact_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mockgcp_cloud_aiplatform_v1_artifact_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mockgcp_cloud_aiplatform_v1_artifact_proto_goTypes = []interface{}{
	(Artifact_State)(0),         // 0: mockgcp.cloud.aiplatform.v1.Artifact.State
	(*Artifact)(nil),            // 1: mockgcp.cloud.aiplatform.v1.Artifact
	nil,                         // 2: mockgcp.cloud.aiplatform.v1.Artifact.LabelsEntry
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*_struct.Struct)(nil),      // 4: google.protobuf.Struct
}
var file_mockgcp_cloud_aiplatform_v1_artifact_proto_depIdxs = []int32{
	2, // 0: mockgcp.cloud.aiplatform.v1.Artifact.labels:type_name -> mockgcp.cloud.aiplatform.v1.Artifact.LabelsEntry
	3, // 1: mockgcp.cloud.aiplatform.v1.Artifact.create_time:type_name -> google.protobuf.Timestamp
	3, // 2: mockgcp.cloud.aiplatform.v1.Artifact.update_time:type_name -> google.protobuf.Timestamp
	0, // 3: mockgcp.cloud.aiplatform.v1.Artifact.state:type_name -> mockgcp.cloud.aiplatform.v1.Artifact.State
	4, // 4: mockgcp.cloud.aiplatform.v1.Artifact.metadata:type_name -> google.protobuf.Struct
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_mockgcp_cloud_aiplatform_v1_artifact_proto_init() }
func file_mockgcp_cloud_aiplatform_v1_artifact_proto_init() {
	if File_mockgcp_cloud_aiplatform_v1_artifact_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mockgcp_cloud_aiplatform_v1_artifact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Artifact); i {
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
			RawDescriptor: file_mockgcp_cloud_aiplatform_v1_artifact_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mockgcp_cloud_aiplatform_v1_artifact_proto_goTypes,
		DependencyIndexes: file_mockgcp_cloud_aiplatform_v1_artifact_proto_depIdxs,
		EnumInfos:         file_mockgcp_cloud_aiplatform_v1_artifact_proto_enumTypes,
		MessageInfos:      file_mockgcp_cloud_aiplatform_v1_artifact_proto_msgTypes,
	}.Build()
	File_mockgcp_cloud_aiplatform_v1_artifact_proto = out.File
	file_mockgcp_cloud_aiplatform_v1_artifact_proto_rawDesc = nil
	file_mockgcp_cloud_aiplatform_v1_artifact_proto_goTypes = nil
	file_mockgcp_cloud_aiplatform_v1_artifact_proto_depIdxs = nil
}
