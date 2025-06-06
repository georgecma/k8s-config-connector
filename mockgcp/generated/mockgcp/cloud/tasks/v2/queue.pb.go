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
// source: mockgcp/cloud/tasks/v2/queue.proto

package cloudtaskspb

import (
	duration "github.com/golang/protobuf/ptypes/duration"
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

// State of the queue.
type Queue_State int32

const (
	// Unspecified state.
	Queue_STATE_UNSPECIFIED Queue_State = 0
	// The queue is running. Tasks can be dispatched.
	//
	// If the queue was created using Cloud Tasks and the queue has
	// had no activity (method calls or task dispatches) for 30 days,
	// the queue may take a few minutes to re-activate. Some method
	// calls may return [NOT_FOUND][google.rpc.Code.NOT_FOUND] and
	// tasks may not be dispatched for a few minutes until the queue
	// has been re-activated.
	Queue_RUNNING Queue_State = 1
	// Tasks are paused by the user. If the queue is paused then Cloud
	// Tasks will stop delivering tasks from it, but more tasks can
	// still be added to it by the user.
	Queue_PAUSED Queue_State = 2
	// The queue is disabled.
	//
	// A queue becomes `DISABLED` when
	// [queue.yaml](https://cloud.google.com/appengine/docs/python/config/queueref)
	// or
	// [queue.xml](https://cloud.google.com/appengine/docs/standard/java/config/queueref)
	// is uploaded which does not contain the queue. You cannot directly disable
	// a queue.
	//
	// When a queue is disabled, tasks can still be added to a queue
	// but the tasks are not dispatched.
	//
	// To permanently delete this queue and all of its tasks, call
	// [DeleteQueue][mockgcp.cloud.tasks.v2.CloudTasks.DeleteQueue].
	Queue_DISABLED Queue_State = 3
)

// Enum value maps for Queue_State.
var (
	Queue_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "RUNNING",
		2: "PAUSED",
		3: "DISABLED",
	}
	Queue_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"RUNNING":           1,
		"PAUSED":            2,
		"DISABLED":          3,
	}
)

func (x Queue_State) Enum() *Queue_State {
	p := new(Queue_State)
	*p = x
	return p
}

func (x Queue_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Queue_State) Descriptor() protoreflect.EnumDescriptor {
	return file_mockgcp_cloud_tasks_v2_queue_proto_enumTypes[0].Descriptor()
}

func (Queue_State) Type() protoreflect.EnumType {
	return &file_mockgcp_cloud_tasks_v2_queue_proto_enumTypes[0]
}

func (x Queue_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Queue_State.Descriptor instead.
func (Queue_State) EnumDescriptor() ([]byte, []int) {
	return file_mockgcp_cloud_tasks_v2_queue_proto_rawDescGZIP(), []int{0, 0}
}

// A queue is a container of related tasks. Queues are configured to manage
// how those tasks are dispatched. Configurable properties include rate limits,
// retry options, queue types, and others.
type Queue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Caller-specified and required in
	// [CreateQueue][mockgcp.cloud.tasks.v2.CloudTasks.CreateQueue], after which it
	// becomes output only.
	//
	// The queue name.
	//
	// The queue name must have the following format:
	// `projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID`
	//
	//   - `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]),
	//     hyphens (-), colons (:), or periods (.).
	//     For more information, see
	//     [Identifying
	//     projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects)
	//   - `LOCATION_ID` is the canonical ID for the queue's location.
	//     The list of available locations can be obtained by calling
	//     [ListLocations][mockgcp.cloud.location.Locations.ListLocations].
	//     For more information, see https://cloud.google.com/about/locations/.
	//   - `QUEUE_ID` can contain letters ([A-Za-z]), numbers ([0-9]), or
	//     hyphens (-). The maximum length is 100 characters.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Overrides for
	// [task-level
	// app_engine_routing][mockgcp.cloud.tasks.v2.AppEngineHttpRequest.app_engine_routing].
	// These settings apply only to
	// [App Engine tasks][mockgcp.cloud.tasks.v2.AppEngineHttpRequest] in this
	// queue. [Http tasks][mockgcp.cloud.tasks.v2.HttpRequest] are not affected.
	//
	// If set, `app_engine_routing_override` is used for all
	// [App Engine tasks][mockgcp.cloud.tasks.v2.AppEngineHttpRequest] in the
	// queue, no matter what the setting is for the [task-level
	// app_engine_routing][mockgcp.cloud.tasks.v2.AppEngineHttpRequest.app_engine_routing].
	AppEngineRoutingOverride *AppEngineRouting `protobuf:"bytes,2,opt,name=app_engine_routing_override,json=appEngineRoutingOverride,proto3" json:"app_engine_routing_override,omitempty"`
	// Rate limits for task dispatches.
	//
	// [rate_limits][mockgcp.cloud.tasks.v2.Queue.rate_limits] and
	// [retry_config][mockgcp.cloud.tasks.v2.Queue.retry_config] are related
	// because they both control task attempts. However they control task attempts
	// in different ways:
	//
	// * [rate_limits][mockgcp.cloud.tasks.v2.Queue.rate_limits] controls the total
	// rate of
	//
	//	dispatches from a queue (i.e. all traffic dispatched from the
	//	queue, regardless of whether the dispatch is from a first
	//	attempt or a retry).
	//
	// * [retry_config][mockgcp.cloud.tasks.v2.Queue.retry_config] controls what
	// happens to
	//
	//	particular a task after its first attempt fails. That is,
	//	[retry_config][mockgcp.cloud.tasks.v2.Queue.retry_config] controls task
	//	retries (the second attempt, third attempt, etc).
	//
	// The queue's actual dispatch rate is the result of:
	//
	// * Number of tasks in the queue
	// * User-specified throttling:
	// [rate_limits][mockgcp.cloud.tasks.v2.Queue.rate_limits],
	//
	//	[retry_config][mockgcp.cloud.tasks.v2.Queue.retry_config], and the
	//	[queue's state][mockgcp.cloud.tasks.v2.Queue.state].
	//   - System throttling due to `429` (Too Many Requests) or `503` (Service
	//     Unavailable) responses from the worker, high error rates, or to smooth
	//     sudden large traffic spikes.
	RateLimits *RateLimits `protobuf:"bytes,3,opt,name=rate_limits,json=rateLimits,proto3" json:"rate_limits,omitempty"`
	// Settings that determine the retry behavior.
	//
	//   - For tasks created using Cloud Tasks: the queue-level retry settings
	//     apply to all tasks in the queue that were created using Cloud Tasks.
	//     Retry settings cannot be set on individual tasks.
	//   - For tasks created using the App Engine SDK: the queue-level retry
	//     settings apply to all tasks in the queue which do not have retry settings
	//     explicitly set on the task and were created by the App Engine SDK. See
	//     [App Engine
	//     documentation](https://cloud.google.com/appengine/docs/standard/python/taskqueue/push/retrying-tasks).
	RetryConfig *RetryConfig `protobuf:"bytes,4,opt,name=retry_config,json=retryConfig,proto3" json:"retry_config,omitempty"`
	// Output only. The state of the queue.
	//
	// `state` can only be changed by calling
	// [PauseQueue][mockgcp.cloud.tasks.v2.CloudTasks.PauseQueue],
	// [ResumeQueue][mockgcp.cloud.tasks.v2.CloudTasks.ResumeQueue], or uploading
	// [queue.yaml/xml](https://cloud.google.com/appengine/docs/python/config/queueref).
	// [UpdateQueue][mockgcp.cloud.tasks.v2.CloudTasks.UpdateQueue] cannot be used
	// to change `state`.
	State Queue_State `protobuf:"varint,5,opt,name=state,proto3,enum=mockgcp.cloud.tasks.v2.Queue_State" json:"state,omitempty"`
	// Output only. The last time this queue was purged.
	//
	// All tasks that were [created][mockgcp.cloud.tasks.v2.Task.create_time]
	// before this time were purged.
	//
	// A queue can be purged using
	// [PurgeQueue][mockgcp.cloud.tasks.v2.CloudTasks.PurgeQueue], the [App Engine
	// Task Queue SDK, or the Cloud
	// Console](https://cloud.google.com/appengine/docs/standard/python/taskqueue/push/deleting-tasks-and-queues#purging_all_tasks_from_a_queue).
	//
	// Purge time will be truncated to the nearest microsecond. Purge
	// time will be unset if the queue has never been purged.
	PurgeTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=purge_time,json=purgeTime,proto3" json:"purge_time,omitempty"`
	// Configuration options for writing logs to
	// [Stackdriver Logging](https://cloud.google.com/logging/docs/). If this
	// field is unset, then no logs are written.
	StackdriverLoggingConfig *StackdriverLoggingConfig `protobuf:"bytes,9,opt,name=stackdriver_logging_config,json=stackdriverLoggingConfig,proto3" json:"stackdriver_logging_config,omitempty"`
}

func (x *Queue) Reset() {
	*x = Queue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_tasks_v2_queue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Queue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Queue) ProtoMessage() {}

func (x *Queue) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_tasks_v2_queue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Queue.ProtoReflect.Descriptor instead.
func (*Queue) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_tasks_v2_queue_proto_rawDescGZIP(), []int{0}
}

func (x *Queue) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Queue) GetAppEngineRoutingOverride() *AppEngineRouting {
	if x != nil {
		return x.AppEngineRoutingOverride
	}
	return nil
}

func (x *Queue) GetRateLimits() *RateLimits {
	if x != nil {
		return x.RateLimits
	}
	return nil
}

func (x *Queue) GetRetryConfig() *RetryConfig {
	if x != nil {
		return x.RetryConfig
	}
	return nil
}

func (x *Queue) GetState() Queue_State {
	if x != nil {
		return x.State
	}
	return Queue_STATE_UNSPECIFIED
}

func (x *Queue) GetPurgeTime() *timestamp.Timestamp {
	if x != nil {
		return x.PurgeTime
	}
	return nil
}

func (x *Queue) GetStackdriverLoggingConfig() *StackdriverLoggingConfig {
	if x != nil {
		return x.StackdriverLoggingConfig
	}
	return nil
}

// Rate limits.
//
// This message determines the maximum rate that tasks can be dispatched by a
// queue, regardless of whether the dispatch is a first task attempt or a retry.
//
// Note: The debugging command,
// [RunTask][mockgcp.cloud.tasks.v2.CloudTasks.RunTask], will run a task even if
// the queue has reached its [RateLimits][mockgcp.cloud.tasks.v2.RateLimits].
type RateLimits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum rate at which tasks are dispatched from this queue.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	// * The maximum allowed value is 500.
	//
	// This field has the same meaning as
	// [rate in
	// queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#rate).
	MaxDispatchesPerSecond float64 `protobuf:"fixed64,1,opt,name=max_dispatches_per_second,json=maxDispatchesPerSecond,proto3" json:"max_dispatches_per_second,omitempty"`
	// Output only. The max burst size.
	//
	// Max burst size limits how fast tasks in queue are processed when
	// many tasks are in the queue and the rate is high. This field
	// allows the queue to have a high rate so processing starts shortly
	// after a task is enqueued, but still limits resource usage when
	// many tasks are enqueued in a short period of time.
	//
	// The [token bucket](https://wikipedia.org/wiki/Token_Bucket)
	// algorithm is used to control the rate of task dispatches. Each
	// queue has a token bucket that holds tokens, up to the maximum
	// specified by `max_burst_size`. Each time a task is dispatched, a
	// token is removed from the bucket. Tasks will be dispatched until
	// the queue's bucket runs out of tokens. The bucket will be
	// continuously refilled with new tokens based on
	// [max_dispatches_per_second][mockgcp.cloud.tasks.v2.RateLimits.max_dispatches_per_second].
	//
	// Cloud Tasks will pick the value of `max_burst_size` based on the
	// value of
	// [max_dispatches_per_second][mockgcp.cloud.tasks.v2.RateLimits.max_dispatches_per_second].
	//
	// For queues that were created or updated using
	// `queue.yaml/xml`, `max_burst_size` is equal to
	// [bucket_size](https://cloud.google.com/appengine/docs/standard/python/config/queueref#bucket_size).
	// Since `max_burst_size` is output only, if
	// [UpdateQueue][mockgcp.cloud.tasks.v2.CloudTasks.UpdateQueue] is called on a
	// queue created by `queue.yaml/xml`, `max_burst_size` will be reset based on
	// the value of
	// [max_dispatches_per_second][mockgcp.cloud.tasks.v2.RateLimits.max_dispatches_per_second],
	// regardless of whether
	// [max_dispatches_per_second][mockgcp.cloud.tasks.v2.RateLimits.max_dispatches_per_second]
	// is updated.
	MaxBurstSize int32 `protobuf:"varint,2,opt,name=max_burst_size,json=maxBurstSize,proto3" json:"max_burst_size,omitempty"`
	// The maximum number of concurrent tasks that Cloud Tasks allows
	// to be dispatched for this queue. After this threshold has been
	// reached, Cloud Tasks stops dispatching tasks until the number of
	// concurrent requests decreases.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	// The maximum allowed value is 5,000.
	//
	// This field has the same meaning as
	// [max_concurrent_requests in
	// queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#max_concurrent_requests).
	MaxConcurrentDispatches int32 `protobuf:"varint,3,opt,name=max_concurrent_dispatches,json=maxConcurrentDispatches,proto3" json:"max_concurrent_dispatches,omitempty"`
}

func (x *RateLimits) Reset() {
	*x = RateLimits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_tasks_v2_queue_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimits) ProtoMessage() {}

func (x *RateLimits) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_tasks_v2_queue_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimits.ProtoReflect.Descriptor instead.
func (*RateLimits) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_tasks_v2_queue_proto_rawDescGZIP(), []int{1}
}

func (x *RateLimits) GetMaxDispatchesPerSecond() float64 {
	if x != nil {
		return x.MaxDispatchesPerSecond
	}
	return 0
}

func (x *RateLimits) GetMaxBurstSize() int32 {
	if x != nil {
		return x.MaxBurstSize
	}
	return 0
}

func (x *RateLimits) GetMaxConcurrentDispatches() int32 {
	if x != nil {
		return x.MaxConcurrentDispatches
	}
	return 0
}

// Retry config.
//
// These settings determine when a failed task attempt is retried.
type RetryConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of attempts per task.
	//
	// Cloud Tasks will attempt the task `max_attempts` times (that is, if the
	// first attempt fails, then there will be `max_attempts - 1` retries). Must
	// be >= -1.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	// -1 indicates unlimited attempts.
	//
	// This field has the same meaning as
	// [task_retry_limit in
	// queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#retry_parameters).
	MaxAttempts int32 `protobuf:"varint,1,opt,name=max_attempts,json=maxAttempts,proto3" json:"max_attempts,omitempty"`
	// If positive, `max_retry_duration` specifies the time limit for
	// retrying a failed task, measured from when the task was first
	// attempted. Once `max_retry_duration` time has passed *and* the
	// task has been attempted
	// [max_attempts][mockgcp.cloud.tasks.v2.RetryConfig.max_attempts] times, no
	// further attempts will be made and the task will be deleted.
	//
	// If zero, then the task age is unlimited.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	// `max_retry_duration` will be truncated to the nearest second.
	//
	// This field has the same meaning as
	// [task_age_limit in
	// queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#retry_parameters).
	MaxRetryDuration *duration.Duration `protobuf:"bytes,2,opt,name=max_retry_duration,json=maxRetryDuration,proto3" json:"max_retry_duration,omitempty"`
	// A task will be [scheduled][mockgcp.cloud.tasks.v2.Task.schedule_time] for
	// retry between [min_backoff][mockgcp.cloud.tasks.v2.RetryConfig.min_backoff]
	// and [max_backoff][mockgcp.cloud.tasks.v2.RetryConfig.max_backoff] duration
	// after it fails, if the queue's
	// [RetryConfig][mockgcp.cloud.tasks.v2.RetryConfig] specifies that the task
	// should be retried.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	// `min_backoff` will be truncated to the nearest second.
	//
	// This field has the same meaning as
	// [min_backoff_seconds in
	// queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#retry_parameters).
	MinBackoff *duration.Duration `protobuf:"bytes,3,opt,name=min_backoff,json=minBackoff,proto3" json:"min_backoff,omitempty"`
	// A task will be [scheduled][mockgcp.cloud.tasks.v2.Task.schedule_time] for
	// retry between [min_backoff][mockgcp.cloud.tasks.v2.RetryConfig.min_backoff]
	// and [max_backoff][mockgcp.cloud.tasks.v2.RetryConfig.max_backoff] duration
	// after it fails, if the queue's
	// [RetryConfig][mockgcp.cloud.tasks.v2.RetryConfig] specifies that the task
	// should be retried.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	// `max_backoff` will be truncated to the nearest second.
	//
	// This field has the same meaning as
	// [max_backoff_seconds in
	// queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#retry_parameters).
	MaxBackoff *duration.Duration `protobuf:"bytes,4,opt,name=max_backoff,json=maxBackoff,proto3" json:"max_backoff,omitempty"`
	// The time between retries will double `max_doublings` times.
	//
	// A task's retry interval starts at
	// [min_backoff][mockgcp.cloud.tasks.v2.RetryConfig.min_backoff], then doubles
	// `max_doublings` times, then increases linearly, and finally
	// retries at intervals of
	// [max_backoff][mockgcp.cloud.tasks.v2.RetryConfig.max_backoff] up to
	// [max_attempts][mockgcp.cloud.tasks.v2.RetryConfig.max_attempts] times.
	//
	// For example, if
	// [min_backoff][mockgcp.cloud.tasks.v2.RetryConfig.min_backoff] is 10s,
	// [max_backoff][mockgcp.cloud.tasks.v2.RetryConfig.max_backoff] is 300s, and
	// `max_doublings` is 3, then the a task will first be retried in
	// 10s. The retry interval will double three times, and then
	// increase linearly by 2^3 * 10s.  Finally, the task will retry at
	// intervals of [max_backoff][mockgcp.cloud.tasks.v2.RetryConfig.max_backoff]
	// until the task has been attempted
	// [max_attempts][mockgcp.cloud.tasks.v2.RetryConfig.max_attempts] times. Thus,
	// the requests will retry at 10s, 20s, 40s, 80s, 160s, 240s, 300s, 300s, ....
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the
	// default.
	//
	// This field has the same meaning as
	// [max_doublings in
	// queue.yaml/xml](https://cloud.google.com/appengine/docs/standard/python/config/queueref#retry_parameters).
	MaxDoublings int32 `protobuf:"varint,5,opt,name=max_doublings,json=maxDoublings,proto3" json:"max_doublings,omitempty"`
}

func (x *RetryConfig) Reset() {
	*x = RetryConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_tasks_v2_queue_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetryConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryConfig) ProtoMessage() {}

func (x *RetryConfig) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_tasks_v2_queue_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetryConfig.ProtoReflect.Descriptor instead.
func (*RetryConfig) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_tasks_v2_queue_proto_rawDescGZIP(), []int{2}
}

func (x *RetryConfig) GetMaxAttempts() int32 {
	if x != nil {
		return x.MaxAttempts
	}
	return 0
}

func (x *RetryConfig) GetMaxRetryDuration() *duration.Duration {
	if x != nil {
		return x.MaxRetryDuration
	}
	return nil
}

func (x *RetryConfig) GetMinBackoff() *duration.Duration {
	if x != nil {
		return x.MinBackoff
	}
	return nil
}

func (x *RetryConfig) GetMaxBackoff() *duration.Duration {
	if x != nil {
		return x.MaxBackoff
	}
	return nil
}

func (x *RetryConfig) GetMaxDoublings() int32 {
	if x != nil {
		return x.MaxDoublings
	}
	return 0
}

// Configuration options for writing logs to
// [Stackdriver Logging](https://cloud.google.com/logging/docs/).
type StackdriverLoggingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the fraction of operations to write to
	// [Stackdriver Logging](https://cloud.google.com/logging/docs/).
	// This field may contain any value between 0.0 and 1.0, inclusive.
	// 0.0 is the default and means that no operations are logged.
	SamplingRatio float64 `protobuf:"fixed64,1,opt,name=sampling_ratio,json=samplingRatio,proto3" json:"sampling_ratio,omitempty"`
}

func (x *StackdriverLoggingConfig) Reset() {
	*x = StackdriverLoggingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_cloud_tasks_v2_queue_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StackdriverLoggingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StackdriverLoggingConfig) ProtoMessage() {}

func (x *StackdriverLoggingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_cloud_tasks_v2_queue_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StackdriverLoggingConfig.ProtoReflect.Descriptor instead.
func (*StackdriverLoggingConfig) Descriptor() ([]byte, []int) {
	return file_mockgcp_cloud_tasks_v2_queue_proto_rawDescGZIP(), []int{3}
}

func (x *StackdriverLoggingConfig) GetSamplingRatio() float64 {
	if x != nil {
		return x.SamplingRatio
	}
	return 0
}

var File_mockgcp_cloud_tasks_v2_queue_proto protoreflect.FileDescriptor

var file_mockgcp_cloud_tasks_v2_queue_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x32, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x76, 0x32, 0x2f,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x05,
	0x0a, 0x05, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x67, 0x0a, 0x1b, 0x61,
	0x70, 0x70, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x67, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x18, 0x61, 0x70, 0x70, 0x45,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x4f, 0x76, 0x65, 0x72,
	0x72, 0x69, 0x64, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x6f, 0x63, 0x6b,
	0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e,
	0x76, 0x32, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x52, 0x0a, 0x72,
	0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x46, 0x0a, 0x0c, 0x72, 0x65, 0x74,
	0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x79, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x72, 0x65, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x39, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x23, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x0a,
	0x70, 0x75, 0x72, 0x67, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x70, 0x75,
	0x72, 0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x6e, 0x0a, 0x1a, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6d, 0x6f,
	0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x18, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e,
	0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x45, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49,
	0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x02,
	0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x3a, 0x5c,
	0xea, 0x41, 0x59, 0x0a, 0x1f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x51,
	0x75, 0x65, 0x75, 0x65, 0x12, 0x36, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x73, 0x2f, 0x7b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x7d, 0x22, 0xa9, 0x01, 0x0a,
	0x0a, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x39, 0x0a, 0x19, 0x6d,
	0x61, 0x78, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x5f, 0x70, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x16,
	0x6d, 0x61, 0x78, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x50, 0x65, 0x72,
	0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x75,
	0x72, 0x73, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x6d, 0x61, 0x78, 0x42, 0x75, 0x72, 0x73, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x3a, 0x0a, 0x19,
	0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x64,
	0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x17, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x44, 0x69,
	0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x22, 0x96, 0x02, 0x0a, 0x0b, 0x52, 0x65, 0x74,
	0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x5f,
	0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x6d, 0x61, 0x78, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x12, 0x47, 0x0a, 0x12, 0x6d,
	0x61, 0x78, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x10, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x74, 0x72, 0x79, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x0b, 0x6d, 0x69, 0x6e, 0x5f, 0x62, 0x61, 0x63, 0x6b,
	0x6f, 0x66, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6d, 0x69, 0x6e, 0x42, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66,
	0x12, 0x3a, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x6d, 0x61, 0x78, 0x42, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x12, 0x23, 0x0a, 0x0d,
	0x6d, 0x61, 0x78, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x69, 0x6e, 0x67,
	0x73, 0x22, 0x41, 0x0a, 0x18, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x25, 0x0a,
	0x0e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x52,
	0x61, 0x74, 0x69, 0x6f, 0x42, 0x6a, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x63, 0x6b,
	0x67, 0x63, 0x70, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e,
	0x76, 0x32, 0x42, 0x0a, 0x51, 0x75, 0x65, 0x75, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x70, 0x62, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mockgcp_cloud_tasks_v2_queue_proto_rawDescOnce sync.Once
	file_mockgcp_cloud_tasks_v2_queue_proto_rawDescData = file_mockgcp_cloud_tasks_v2_queue_proto_rawDesc
)

func file_mockgcp_cloud_tasks_v2_queue_proto_rawDescGZIP() []byte {
	file_mockgcp_cloud_tasks_v2_queue_proto_rawDescOnce.Do(func() {
		file_mockgcp_cloud_tasks_v2_queue_proto_rawDescData = protoimpl.X.CompressGZIP(file_mockgcp_cloud_tasks_v2_queue_proto_rawDescData)
	})
	return file_mockgcp_cloud_tasks_v2_queue_proto_rawDescData
}

var file_mockgcp_cloud_tasks_v2_queue_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mockgcp_cloud_tasks_v2_queue_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mockgcp_cloud_tasks_v2_queue_proto_goTypes = []interface{}{
	(Queue_State)(0),                 // 0: mockgcp.cloud.tasks.v2.Queue.State
	(*Queue)(nil),                    // 1: mockgcp.cloud.tasks.v2.Queue
	(*RateLimits)(nil),               // 2: mockgcp.cloud.tasks.v2.RateLimits
	(*RetryConfig)(nil),              // 3: mockgcp.cloud.tasks.v2.RetryConfig
	(*StackdriverLoggingConfig)(nil), // 4: mockgcp.cloud.tasks.v2.StackdriverLoggingConfig
	(*AppEngineRouting)(nil),         // 5: mockgcp.cloud.tasks.v2.AppEngineRouting
	(*timestamp.Timestamp)(nil),      // 6: google.protobuf.Timestamp
	(*duration.Duration)(nil),        // 7: google.protobuf.Duration
}
var file_mockgcp_cloud_tasks_v2_queue_proto_depIdxs = []int32{
	5, // 0: mockgcp.cloud.tasks.v2.Queue.app_engine_routing_override:type_name -> mockgcp.cloud.tasks.v2.AppEngineRouting
	2, // 1: mockgcp.cloud.tasks.v2.Queue.rate_limits:type_name -> mockgcp.cloud.tasks.v2.RateLimits
	3, // 2: mockgcp.cloud.tasks.v2.Queue.retry_config:type_name -> mockgcp.cloud.tasks.v2.RetryConfig
	0, // 3: mockgcp.cloud.tasks.v2.Queue.state:type_name -> mockgcp.cloud.tasks.v2.Queue.State
	6, // 4: mockgcp.cloud.tasks.v2.Queue.purge_time:type_name -> google.protobuf.Timestamp
	4, // 5: mockgcp.cloud.tasks.v2.Queue.stackdriver_logging_config:type_name -> mockgcp.cloud.tasks.v2.StackdriverLoggingConfig
	7, // 6: mockgcp.cloud.tasks.v2.RetryConfig.max_retry_duration:type_name -> google.protobuf.Duration
	7, // 7: mockgcp.cloud.tasks.v2.RetryConfig.min_backoff:type_name -> google.protobuf.Duration
	7, // 8: mockgcp.cloud.tasks.v2.RetryConfig.max_backoff:type_name -> google.protobuf.Duration
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_mockgcp_cloud_tasks_v2_queue_proto_init() }
func file_mockgcp_cloud_tasks_v2_queue_proto_init() {
	if File_mockgcp_cloud_tasks_v2_queue_proto != nil {
		return
	}
	file_mockgcp_cloud_tasks_v2_target_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mockgcp_cloud_tasks_v2_queue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Queue); i {
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
		file_mockgcp_cloud_tasks_v2_queue_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimits); i {
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
		file_mockgcp_cloud_tasks_v2_queue_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetryConfig); i {
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
		file_mockgcp_cloud_tasks_v2_queue_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StackdriverLoggingConfig); i {
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
			RawDescriptor: file_mockgcp_cloud_tasks_v2_queue_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mockgcp_cloud_tasks_v2_queue_proto_goTypes,
		DependencyIndexes: file_mockgcp_cloud_tasks_v2_queue_proto_depIdxs,
		EnumInfos:         file_mockgcp_cloud_tasks_v2_queue_proto_enumTypes,
		MessageInfos:      file_mockgcp_cloud_tasks_v2_queue_proto_msgTypes,
	}.Build()
	File_mockgcp_cloud_tasks_v2_queue_proto = out.File
	file_mockgcp_cloud_tasks_v2_queue_proto_rawDesc = nil
	file_mockgcp_cloud_tasks_v2_queue_proto_goTypes = nil
	file_mockgcp_cloud_tasks_v2_queue_proto_depIdxs = nil
}
