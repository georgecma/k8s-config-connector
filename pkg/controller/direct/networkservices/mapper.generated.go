// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +generated:mapper
// krm.group: networkservices.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.networkservices.v1

package networkservices

import (
	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func NetworkServicesServiceBindingObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ServiceBinding) *krm.NetworkServicesServiceBindingObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkServicesServiceBindingObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func NetworkServicesServiceBindingObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkServicesServiceBindingObservedState) *pb.ServiceBinding {
	if in == nil {
		return nil
	}
	out := &pb.ServiceBinding{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
