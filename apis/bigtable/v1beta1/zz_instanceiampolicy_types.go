/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type InstanceIAMPolicyInitParameters struct {

	// The policy data generated by a google_iam_policy data source.
	PolicyData *string `json:"policyData,omitempty" tf:"policy_data,omitempty"`

	// The project in which the instance belongs.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type InstanceIAMPolicyObservation struct {

	// (Computed) The etag of the instances's IAM policy.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name or relative resource id of the instance to manage IAM policies for.
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// The policy data generated by a google_iam_policy data source.
	PolicyData *string `json:"policyData,omitempty" tf:"policy_data,omitempty"`

	// The project in which the instance belongs.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type InstanceIAMPolicyParameters struct {

	// The name or relative resource id of the instance to manage IAM policies for.
	// +kubebuilder:validation:Required
	Instance *string `json:"instance" tf:"instance,omitempty"`

	// The policy data generated by a google_iam_policy data source.
	// +kubebuilder:validation:Optional
	PolicyData *string `json:"policyData,omitempty" tf:"policy_data,omitempty"`

	// The project in which the instance belongs.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// InstanceIAMPolicySpec defines the desired state of InstanceIAMPolicy
type InstanceIAMPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceIAMPolicyParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider InstanceIAMPolicyInitParameters `json:"initProvider,omitempty"`
}

// InstanceIAMPolicyStatus defines the observed state of InstanceIAMPolicy.
type InstanceIAMPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceIAMPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceIAMPolicy is the Schema for the InstanceIAMPolicys API. Collection of resources to manage IAM policy for a Bigtable instance.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type InstanceIAMPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policyData) || has(self.initProvider.policyData)",message="policyData is a required parameter"
	Spec   InstanceIAMPolicySpec   `json:"spec"`
	Status InstanceIAMPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceIAMPolicyList contains a list of InstanceIAMPolicys
type InstanceIAMPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceIAMPolicy `json:"items"`
}

// Repository type metadata.
var (
	InstanceIAMPolicy_Kind             = "InstanceIAMPolicy"
	InstanceIAMPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstanceIAMPolicy_Kind}.String()
	InstanceIAMPolicy_KindAPIVersion   = InstanceIAMPolicy_Kind + "." + CRDGroupVersion.String()
	InstanceIAMPolicy_GroupVersionKind = CRDGroupVersion.WithKind(InstanceIAMPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&InstanceIAMPolicy{}, &InstanceIAMPolicyList{})
}
