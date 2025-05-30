// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DatastreamRouteGVK = GroupVersion.WithKind("DatastreamRoute")

// DatastreamRouteSpec defines the desired state of DatastreamRoute
// +kcc:spec:proto=google.cloud.datastream.v1.Route
type DatastreamRouteSpec struct {
	// The DatastreamRoute name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The parent that owns the collection of Routes.
	// +required
	PrivateConnectionRef *PrivateConnectionRef `json:"privateConnectionRef,omitempty"`

	// Labels.
	// +kcc:proto:field=google.cloud.datastream.v1.Route.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Display name.
	// +kcc:proto:field=google.cloud.datastream.v1.Route.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Destination address for connection
	// +kcc:proto:field=google.cloud.datastream.v1.Route.destination_address
	DestinationAddress *string `json:"destinationAddress,omitempty"`

	// Destination port for connection
	// +kcc:proto:field=google.cloud.datastream.v1.Route.destination_port
	DestinationPort *int32 `json:"destinationPort,omitempty"`
}

// DatastreamRouteStatus defines the config connector machine state of DatastreamRoute
type DatastreamRouteStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DatastreamRoute resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DatastreamRouteObservedState `json:"observedState,omitempty"`
}

// DatastreamRouteObservedState is the state of the DatastreamRoute resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.datastream.v1.Route
type DatastreamRouteObservedState struct {
	// Output only. The resource's name.
	// +kcc:proto:field=google.cloud.datastream.v1.Route.name
	// NOTYET: this field serves the same purpose as externalRef
	// Name *string `json:"name,omitempty"`

	// Output only. The create time of the resource.
	// +kcc:proto:field=google.cloud.datastream.v1.Route.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update time of the resource.
	// +kcc:proto:field=google.cloud.datastream.v1.Route.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdatastreamroute;gcpdatastreamroutes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DatastreamRoute is the Schema for the DatastreamRoute API
// +k8s:openapi-gen=true
type DatastreamRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DatastreamRouteSpec   `json:"spec,omitempty"`
	Status DatastreamRouteStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DatastreamRouteList contains a list of DatastreamRoute
type DatastreamRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatastreamRoute `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DatastreamRoute{}, &DatastreamRouteList{})
}
