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

type DiskResourcePolicyAttachmentInitParameters struct {

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// A reference to the zone where the disk resides.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type DiskResourcePolicyAttachmentObservation struct {

	// The name of the disk in which the resource policies are attached to.
	Disk *string `json:"disk,omitempty" tf:"disk,omitempty"`

	// an identifier for the resource with format {{project}}/{{zone}}/{{disk}}/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The resource policy to be attached to the disk for scheduling snapshot
	// creation. Do not specify the self link.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// A reference to the zone where the disk resides.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type DiskResourcePolicyAttachmentParameters struct {

	// The name of the disk in which the resource policies are attached to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Disk
	// +kubebuilder:validation:Optional
	Disk *string `json:"disk,omitempty" tf:"disk,omitempty"`

	// Reference to a Disk in compute to populate disk.
	// +kubebuilder:validation:Optional
	DiskRef *v1.Reference `json:"diskRef,omitempty" tf:"-"`

	// Selector for a Disk in compute to populate disk.
	// +kubebuilder:validation:Optional
	DiskSelector *v1.Selector `json:"diskSelector,omitempty" tf:"-"`

	// The resource policy to be attached to the disk for scheduling snapshot
	// creation. Do not specify the self link.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.ResourcePolicy
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a ResourcePolicy in compute to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a ResourcePolicy in compute to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// A reference to the zone where the disk resides.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// DiskResourcePolicyAttachmentSpec defines the desired state of DiskResourcePolicyAttachment
type DiskResourcePolicyAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DiskResourcePolicyAttachmentParameters `json:"forProvider"`
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
	InitProvider DiskResourcePolicyAttachmentInitParameters `json:"initProvider,omitempty"`
}

// DiskResourcePolicyAttachmentStatus defines the observed state of DiskResourcePolicyAttachment.
type DiskResourcePolicyAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DiskResourcePolicyAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DiskResourcePolicyAttachment is the Schema for the DiskResourcePolicyAttachments API. Adds existing resource policies to a disk.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type DiskResourcePolicyAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DiskResourcePolicyAttachmentSpec   `json:"spec"`
	Status            DiskResourcePolicyAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DiskResourcePolicyAttachmentList contains a list of DiskResourcePolicyAttachments
type DiskResourcePolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiskResourcePolicyAttachment `json:"items"`
}

// Repository type metadata.
var (
	DiskResourcePolicyAttachment_Kind             = "DiskResourcePolicyAttachment"
	DiskResourcePolicyAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DiskResourcePolicyAttachment_Kind}.String()
	DiskResourcePolicyAttachment_KindAPIVersion   = DiskResourcePolicyAttachment_Kind + "." + CRDGroupVersion.String()
	DiskResourcePolicyAttachment_GroupVersionKind = CRDGroupVersion.WithKind(DiskResourcePolicyAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&DiskResourcePolicyAttachment{}, &DiskResourcePolicyAttachmentList{})
}
