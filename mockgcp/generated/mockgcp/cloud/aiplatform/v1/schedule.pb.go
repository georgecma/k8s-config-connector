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
// source: mockgcp/cloud/aiplatform/v1/schedule.proto

package aiplatformpb

import (
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

// Possible state of the schedule.
type Schedule_State int32

const (
	// Unspecified.
	Schedule_STATE_UNSPECIFIED Schedule_State = 0
	// The Schedule is active. Runs are being scheduled on the user-specified
	// timespec.
	Schedule_ACTIVE Schedule_State = 1
	// The schedule is paused. No new runs will be created until the schedule
	// is resumed. Already started runs will be allowed to complete.
	Schedule_PAUSED Schedule_State = 2
	// The Schedule is completed. No new runs will be scheduled. Already started
	// runs will be allowed to complete. Schedules in completed state cannot be
	// paused or resumed.
	Schedule_COMPLETED Schedule_State = 3
)

// Enum value maps for Schedule_State.
var (
	Schedule_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "ACTIVE",
		2: "PAUSED",
		3: "COMPLETED",
	}
	Schedule_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"ACTIVE":            1,
		"PAUSED":            2,
		"COMPLETED":         3,
	}
)

func (x Schedule_State) Enum() *Schedule_State {
	p := new(Schedule_State)
	*p = x
	return p
}

func (x Schedule_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Schedule_State) Descriptor() protoreflect.EnumDescriptor {
	return file_mockgcp_cloud_aiplatform_v1_schedule_proto_enumTypes[0].Descriptor()
}

func (Schedule_State) Type() protoreflect.EnumType {
	return &file_mockgcp_cloud_aiplatform_v1_schedule_proto_enumTypes[0]
}

func (x Schedule_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Schedule_State.Descriptor instead.
func (Schedule_State) EnumDescriptor() ([]byte, []int) {
	return file_mockgcp_cloud_aiplatform_v1_schedule_proto_rawDescGZIP(), []int{0, 0}
}

// An instance of a Schedule periodically schedules runs to make API calls based
// on user specified time specification and API request type.
type Schedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required.
	// The time specification to launch scheduled runs.
	//
	// Types that are assignable to TimeSpecification:
	//
	//	*Schedule_Cron
	TimeSpecification isSchedule_TimeSpecification `protobuf_oneof:"time_specification"`
	// Required.
	// The API request template to launch the scheduled runs.
	// User-specified ID is not supported in the request template.
	//
	// Types that are assignable to Request:
	//
	//	*Schedule_CreatePipelineJobRequest
	//	*Schedule_CreateNotebookExecutionJobRequest
	Request isSchedule_Request `protobuf_oneof:"request"`
	// Immutable. The resource name of the Schedule.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. User provided name of the Schedule.
	// The name can be up to 128 characters long and can consist of any UTF-8
	// characters.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Optional. Timestamp after which the first run can be scheduled.
	// Default to Schedule create time if not specified.
	StartTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Optional. Timestamp after which no new runs can be scheduled.
	// If specified, The schedule will be completed when either
	// end_time is reached or when scheduled_run_count >= max_run_count.
	// If not specified, new runs will keep getting scheduled until this Schedule
	// is paused or deleted. Already scheduled runs will be allowed to complete.
	// Unset if not specified.
	EndTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Optional. Maximum run count of the schedule.
	// If specified, The schedule will be completed when either
	// started_run_count >= max_run_count or when end_time is reached.
	// If not specified, new runs will keep getting scheduled until this Schedule
	// is paused or deleted. Already scheduled runs will be allowed to complete.
	// Unset if not specified.
	MaxRunCount int64 `protobuf:"varint,16,opt,name=max_run_count,json=maxRunCount,proto3" json:"max_run_count,omitempty"`
	// Output only. The number of runs started by this schedule.
	StartedRunCount int64 `protobuf:"varint,17,opt,name=started_run_count,json=startedRunCount,proto3" json:"started_run_count,omitempty"`
	// Output only. The state of this Schedule.
	State Schedule_State `protobuf:"varint,5,opt,name=state,proto3,enum=mockgcp.cloud.aiplatform.v1.Schedule_State" json:"state,omitempty"`
	// Output only. Timestamp when this Schedule was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Timestamp when this Schedule was updated.
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,19,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Output only. Timestamp when this Schedule should schedule the next run.
	// Having a next_run_time in the past means the runs are being started
	// behind schedule.
	NextRunTime *timestamp.Timestamp `protobuf:"bytes,7,opt,name=next_run_time,json=nextRunTime,proto3" json:"next_run_time,omitempty"`
	// Output only. Timestamp when this Schedule was last paused.
	// Unset if never paused.
	LastPauseTime *timestamp.Timestamp `protobuf:"bytes,8,opt,name=last_pause_time,json=lastPauseTime,proto3" json:"last_pause_time,omitempty"`
	// Output only. Timestamp when this Schedule was last resumed.
	// Unset if never resumed from pause.
	LastResumeTime *timestamp.Timestamp `protobuf:"bytes,9,opt,name=last_resume_time,json=lastResumeTime,proto3" json:"last_resume_time,omitempty"`
	// Required. Maximum number of runs that can be started concurrently for this
	// Schedule. This is the limit for starting the scheduled requests and not the
	// execution of the operations/jobs created by the requests (if applicable).
	MaxConcurrentRunCount int64 `protobuf:"varint,11,opt,name=max_concurrent_run_count,json=maxConcurrentRunCount,proto3" json:"max_concurrent_run_count,omitempty"`
	// Optional. Whether new scheduled runs can be queued when max_concurrent_runs
	// limit is reached. If set to true, new runs will be queued instead of
	// skipped. Default to false.
	AllowQueueing bool `protobuf:"varint,12,opt,name=allow_queueing,json=allowQueueing,proto3" json:"allow_queueing,omitempty"`
	// Output only. Whether to backfill missed runs when the schedule is resumed
	// from PAUSED state. If set to true, all missed runs will be scheduled. New
	// runs will be scheduled after the backfill is complete. Default to false.
	CatchUp bool `protobuf:"varint,13,opt,name=catch_up,json=catchUp,proto3" json:"catch_up,omitempty"`
	// Output only. Response of the last scheduled run.
	// This is the response for starting the scheduled requests and not the
	// execution of the operations/jobs created by the requests (if applicable).
	// Unset if no run has been scheduled yet.
	LastScheduledRunResponse *Schedule_RunResponse `protobuf:"bytes,18,opt,name=last_scheduled_run_response,json=lastScheduledRunResponse,proto3" json:"last_scheduled_run_response,omitempty"`
}

func (x *Schedule) Reset() {
	*x = Schedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_aiplatform_v1_schedule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedule) ProtoMessage() {}

func (x *Schedule) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_aiplatform_v1_schedule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedule.ProtoReflect.Descriptor instead.
func (*Schedule) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_aiplatform_v1_schedule_proto_rawDescGZIP(), []int{0}
}

func (m *Schedule) GetTimeSpecification() isSchedule_TimeSpecification {
	if m != nil {
		return m.TimeSpecification
	}
	return nil
}

func (x *Schedule) GetCron() string {
	if x, ok := x.GetTimeSpecification().(*Schedule_Cron); ok {
		return x.Cron
	}
	return ""
}

func (m *Schedule) GetRequest() isSchedule_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *Schedule) GetCreatePipelineJobRequest() *CreatePipelineJobRequest {
	if x, ok := x.GetRequest().(*Schedule_CreatePipelineJobRequest); ok {
		return x.CreatePipelineJobRequest
	}
	return nil
}

func (x *Schedule) GetCreateNotebookExecutionJobRequest() *CreateNotebookExecutionJobRequest {
	if x, ok := x.GetRequest().(*Schedule_CreateNotebookExecutionJobRequest); ok {
		return x.CreateNotebookExecutionJobRequest
	}
	return nil
}

func (x *Schedule) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Schedule) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Schedule) GetStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Schedule) GetEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *Schedule) GetMaxRunCount() int64 {
	if x != nil {
		return x.MaxRunCount
	}
	return 0
}

func (x *Schedule) GetStartedRunCount() int64 {
	if x != nil {
		return x.StartedRunCount
	}
	return 0
}

func (x *Schedule) GetState() Schedule_State {
	if x != nil {
		return x.State
	}
	return Schedule_STATE_UNSPECIFIED
}

func (x *Schedule) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Schedule) GetUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Schedule) GetNextRunTime() *timestamp.Timestamp {
	if x != nil {
		return x.NextRunTime
	}
	return nil
}

func (x *Schedule) GetLastPauseTime() *timestamp.Timestamp {
	if x != nil {
		return x.LastPauseTime
	}
	return nil
}

func (x *Schedule) GetLastResumeTime() *timestamp.Timestamp {
	if x != nil {
		return x.LastResumeTime
	}
	return nil
}

func (x *Schedule) GetMaxConcurrentRunCount() int64 {
	if x != nil {
		return x.MaxConcurrentRunCount
	}
	return 0
}

func (x *Schedule) GetAllowQueueing() bool {
	if x != nil {
		return x.AllowQueueing
	}
	return false
}

func (x *Schedule) GetCatchUp() bool {
	if x != nil {
		return x.CatchUp
	}
	return false
}

func (x *Schedule) GetLastScheduledRunResponse() *Schedule_RunResponse {
	if x != nil {
		return x.LastScheduledRunResponse
	}
	return nil
}

type isSchedule_TimeSpecification interface {
	isSchedule_TimeSpecification()
}

type Schedule_Cron struct {
	// Cron schedule (https://en.wikipedia.org/wiki/Cron) to launch scheduled
	// runs. To explicitly set a timezone to the cron tab, apply a prefix in the
	// cron tab: "CRON_TZ=${IANA_TIME_ZONE}" or "TZ=${IANA_TIME_ZONE}".
	// The ${IANA_TIME_ZONE} may only be a valid string from IANA time zone
	// database. For example, "CRON_TZ=America/New_York 1 * * * *", or
	// "TZ=America/New_York 1 * * * *".
	Cron string `protobuf:"bytes,10,opt,name=cron,proto3,oneof"`
}

func (*Schedule_Cron) isSchedule_TimeSpecification() {}

type isSchedule_Request interface {
	isSchedule_Request()
}

type Schedule_CreatePipelineJobRequest struct {
	// Request for
	// [PipelineService.CreatePipelineJob][mockgcp.cloud.aiplatform.v1.PipelineService.CreatePipelineJob].
	// CreatePipelineJobRequest.parent field is required (format:
	// projects/{project}/locations/{location}).
	CreatePipelineJobRequest *CreatePipelineJobRequest `protobuf:"bytes,14,opt,name=create_pipeline_job_request,json=createPipelineJobRequest,proto3,oneof"`
}

type Schedule_CreateNotebookExecutionJobRequest struct {
	// Request for
	// [NotebookService.CreateNotebookExecutionJob][mockgcp.cloud.aiplatform.v1.NotebookService.CreateNotebookExecutionJob].
	CreateNotebookExecutionJobRequest *CreateNotebookExecutionJobRequest `protobuf:"bytes,20,opt,name=create_notebook_execution_job_request,json=createNotebookExecutionJobRequest,proto3,oneof"`
}

func (*Schedule_CreatePipelineJobRequest) isSchedule_Request() {}

func (*Schedule_CreateNotebookExecutionJobRequest) isSchedule_Request() {}

// Status of a scheduled run.
type Schedule_RunResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The scheduled run time based on the user-specified schedule.
	ScheduledRunTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=scheduled_run_time,json=scheduledRunTime,proto3" json:"scheduled_run_time,omitempty"`
	// The response of the scheduled run.
	RunResponse string `protobuf:"bytes,2,opt,name=run_response,json=runResponse,proto3" json:"run_response,omitempty"`
}

func (x *Schedule_RunResponse) Reset() {
	*x = Schedule_RunResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_aiplatform_v1_schedule_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schedule_RunResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedule_RunResponse) ProtoMessage() {}

func (x *Schedule_RunResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_aiplatform_v1_schedule_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedule_RunResponse.ProtoReflect.Descriptor instead.
func (*Schedule_RunResponse) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_aiplatform_v1_schedule_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Schedule_RunResponse) GetScheduledRunTime() *timestamp.Timestamp {
	if x != nil {
		return x.ScheduledRunTime
	}
	return nil
}

func (x *Schedule_RunResponse) GetRunResponse() string {
	if x != nil {
		return x.RunResponse
	}
	return ""
}

var File_mockgcp_cloud_aiplatform_v1_schedule_proto protoreflect.FileDescriptor

var file_mockgcp_cloud_aiplatform_v1_schedule_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x6d, 0x6f,
	0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x6d, 0x6f, 0x63, 0x6b, 0x67,
	0x63, 0x70, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5,
	0x0c, 0x0a, 0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x04, 0x63,
	0x72, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x63, 0x72, 0x6f,
	0x6e, 0x12, 0x76, 0x0a, 0x1b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x01, 0x52,
	0x18, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x92, 0x01, 0x0a, 0x25, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x6d, 0x6f, 0x63, 0x6b,
	0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f,
	0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x01, 0x52, 0x21, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x05, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x3e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x03, 0xe0, 0x41, 0x01, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x3a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0,
	0x41, 0x01, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0d, 0x6d,
	0x61, 0x78, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x52, 0x75, 0x6e, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x11, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f,
	0x72, 0x75, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x52, 0x75, 0x6e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x46, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x40, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x43, 0x0a, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x52,
	0x75, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x47, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x70,
	0x61, 0x75, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x50, 0x61, 0x75, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x49, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x18, 0x6d, 0x61,
	0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x75, 0x6e,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x15, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x52, 0x75, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x0e, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x69, 0x6e, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08,
	0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x75, 0x70,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x07, 0x63, 0x61, 0x74,
	0x63, 0x68, 0x55, 0x70, 0x12, 0x75, 0x0a, 0x1b, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x6d, 0x6f, 0x63, 0x6b,
	0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x18, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64,
	0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x7a, 0x0a, 0x0b, 0x52,
	0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x12, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x10, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x52, 0x75, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x75, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x75, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x45, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56,
	0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x02, 0x12,
	0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x03, 0x3a, 0x65,
	0xea, 0x41, 0x62, 0x0a, 0x22, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x3c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d,
	0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x7d, 0x42, 0x14, 0x0a, 0x12, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0xcc, 0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x6d,
	0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x76,
	0x31, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0x3b, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0xaa, 0x02, 0x1a, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mockgcp_cloud_aiplatform_v1_schedule_proto_rawDescOnce sync.Once
	file_mockgcp_cloud_aiplatform_v1_schedule_proto_rawDescData = file_mockgcp_cloud_aiplatform_v1_schedule_proto_rawDesc
)

func file_mockgcp_cloud_aiplatform_v1_schedule_proto_rawDescGZIP() []byte {
	file_mockgcp_cloud_aiplatform_v1_schedule_proto_rawDescOnce.Do(func() {
		file_mockgcp_cloud_aiplatform_v1_schedule_proto_rawDescData = protoimpl.X.CompressGZIP(file_mockgcp_cloud_aiplatform_v1_schedule_proto_rawDescData)
	})
	return file_mockgcp_cloud_aiplatform_v1_schedule_proto_rawDescData
}

var file_mockgcp_cloud_aiplatform_v1_schedule_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mockgcp_cloud_aiplatform_v1_schedule_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mockgcp_cloud_aiplatform_v1_schedule_proto_goTypes = []interface{}{
	(Schedule_State)(0),                       // 0: mockgcp.cloud.aiplatform.v1.Schedule.State
	(*Schedule)(nil),                          // 1: mockgcp.cloud.aiplatform.v1.Schedule
	(*Schedule_RunResponse)(nil),              // 2: mockgcp.cloud.aiplatform.v1.Schedule.RunResponse
	(*CreatePipelineJobRequest)(nil),          // 3: mockgcp.cloud.aiplatform.v1.CreatePipelineJobRequest
	(*CreateNotebookExecutionJobRequest)(nil), // 4: mockgcp.cloud.aiplatform.v1.CreateNotebookExecutionJobRequest
	(*timestamp.Timestamp)(nil),               // 5: google.protobuf.Timestamp
}
var file_mockgcp_cloud_aiplatform_v1_schedule_proto_depIdxs = []int32{
	3,  // 0: mockgcp.cloud.aiplatform.v1.Schedule.create_pipeline_job_request:type_name -> mockgcp.cloud.aiplatform.v1.CreatePipelineJobRequest
	4,  // 1: mockgcp.cloud.aiplatform.v1.Schedule.create_notebook_execution_job_request:type_name -> mockgcp.cloud.aiplatform.v1.CreateNotebookExecutionJobRequest
	5,  // 2: mockgcp.cloud.aiplatform.v1.Schedule.start_time:type_name -> google.protobuf.Timestamp
	5,  // 3: mockgcp.cloud.aiplatform.v1.Schedule.end_time:type_name -> google.protobuf.Timestamp
	0,  // 4: mockgcp.cloud.aiplatform.v1.Schedule.state:type_name -> mockgcp.cloud.aiplatform.v1.Schedule.State
	5,  // 5: mockgcp.cloud.aiplatform.v1.Schedule.create_time:type_name -> google.protobuf.Timestamp
	5,  // 6: mockgcp.cloud.aiplatform.v1.Schedule.update_time:type_name -> google.protobuf.Timestamp
	5,  // 7: mockgcp.cloud.aiplatform.v1.Schedule.next_run_time:type_name -> google.protobuf.Timestamp
	5,  // 8: mockgcp.cloud.aiplatform.v1.Schedule.last_pause_time:type_name -> google.protobuf.Timestamp
	5,  // 9: mockgcp.cloud.aiplatform.v1.Schedule.last_resume_time:type_name -> google.protobuf.Timestamp
	2,  // 10: mockgcp.cloud.aiplatform.v1.Schedule.last_scheduled_run_response:type_name -> mockgcp.cloud.aiplatform.v1.Schedule.RunResponse
	5,  // 11: mockgcp.cloud.aiplatform.v1.Schedule.RunResponse.scheduled_run_time:type_name -> google.protobuf.Timestamp
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_mockgcp_cloud_aiplatform_v1_schedule_proto_init() }
func file_mockgcp_cloud_aiplatform_v1_schedule_proto_init() {
	if File_mockgcp_cloud_aiplatform_v1_schedule_proto != nil {
		return
	}
	file_mockgcp_cloud_aiplatform_v1_notebook_service_proto_init()
	file_mockgcp_cloud_aiplatform_v1_pipeline_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mockgcp_cloud_aiplatform_v1_schedule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schedule); i {
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
		file_mockgcp_cloud_aiplatform_v1_schedule_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schedule_RunResponse); i {
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
	file_mockgcp_cloud_aiplatform_v1_schedule_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Schedule_Cron)(nil),
		(*Schedule_CreatePipelineJobRequest)(nil),
		(*Schedule_CreateNotebookExecutionJobRequest)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mockgcp_cloud_aiplatform_v1_schedule_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mockgcp_cloud_aiplatform_v1_schedule_proto_goTypes,
		DependencyIndexes: file_mockgcp_cloud_aiplatform_v1_schedule_proto_depIdxs,
		EnumInfos:         file_mockgcp_cloud_aiplatform_v1_schedule_proto_enumTypes,
		MessageInfos:      file_mockgcp_cloud_aiplatform_v1_schedule_proto_msgTypes,
	}.Build()
	File_mockgcp_cloud_aiplatform_v1_schedule_proto = out.File
	file_mockgcp_cloud_aiplatform_v1_schedule_proto_rawDesc = nil
	file_mockgcp_cloud_aiplatform_v1_schedule_proto_goTypes = nil
	file_mockgcp_cloud_aiplatform_v1_schedule_proto_depIdxs = nil
}
