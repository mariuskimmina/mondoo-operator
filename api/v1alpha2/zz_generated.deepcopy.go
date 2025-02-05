//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Mondoo, Inc.

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha2

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Admission) DeepCopyInto(out *Admission) {
	*out = *in
	out.Image = in.Image
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	out.CertificateProvisioning = in.CertificateProvisioning
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Admission.
func (in *Admission) DeepCopy() *Admission {
	if in == nil {
		return nil
	}
	out := new(Admission)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateProvisioning) DeepCopyInto(out *CertificateProvisioning) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateProvisioning.
func (in *CertificateProvisioning) DeepCopy() *CertificateProvisioning {
	if in == nil {
		return nil
	}
	out := new(CertificateProvisioning)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleIntegration) DeepCopyInto(out *ConsoleIntegration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleIntegration.
func (in *ConsoleIntegration) DeepCopy() *ConsoleIntegration {
	if in == nil {
		return nil
	}
	out := new(ConsoleIntegration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Containers) DeepCopyInto(out *Containers) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Containers.
func (in *Containers) DeepCopy() *Containers {
	if in == nil {
		return nil
	}
	out := new(Containers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Filtering) DeepCopyInto(out *Filtering) {
	*out = *in
	in.Namespaces.DeepCopyInto(&out.Namespaces)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Filtering.
func (in *Filtering) DeepCopy() *Filtering {
	if in == nil {
		return nil
	}
	out := new(Filtering)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilteringSpec) DeepCopyInto(out *FilteringSpec) {
	*out = *in
	if in.Include != nil {
		in, out := &in.Include, &out.Include
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Exclude != nil {
		in, out := &in.Exclude, &out.Exclude
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilteringSpec.
func (in *FilteringSpec) DeepCopy() *FilteringSpec {
	if in == nil {
		return nil
	}
	out := new(FilteringSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesResources) DeepCopyInto(out *KubernetesResources) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesResources.
func (in *KubernetesResources) DeepCopy() *KubernetesResources {
	if in == nil {
		return nil
	}
	out := new(KubernetesResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metrics) DeepCopyInto(out *Metrics) {
	*out = *in
	if in.ResourceLabels != nil {
		in, out := &in.ResourceLabels, &out.ResourceLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metrics.
func (in *Metrics) DeepCopy() *Metrics {
	if in == nil {
		return nil
	}
	out := new(Metrics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MondooAuditConfig) DeepCopyInto(out *MondooAuditConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MondooAuditConfig.
func (in *MondooAuditConfig) DeepCopy() *MondooAuditConfig {
	if in == nil {
		return nil
	}
	out := new(MondooAuditConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MondooAuditConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MondooAuditConfigCondition) DeepCopyInto(out *MondooAuditConfigCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MondooAuditConfigCondition.
func (in *MondooAuditConfigCondition) DeepCopy() *MondooAuditConfigCondition {
	if in == nil {
		return nil
	}
	out := new(MondooAuditConfigCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MondooAuditConfigList) DeepCopyInto(out *MondooAuditConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MondooAuditConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MondooAuditConfigList.
func (in *MondooAuditConfigList) DeepCopy() *MondooAuditConfigList {
	if in == nil {
		return nil
	}
	out := new(MondooAuditConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MondooAuditConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MondooAuditConfigSpec) DeepCopyInto(out *MondooAuditConfigSpec) {
	*out = *in
	out.MondooCredsSecretRef = in.MondooCredsSecretRef
	out.MondooTokenSecretRef = in.MondooTokenSecretRef
	in.Scanner.DeepCopyInto(&out.Scanner)
	out.KubernetesResources = in.KubernetesResources
	in.Nodes.DeepCopyInto(&out.Nodes)
	in.Admission.DeepCopyInto(&out.Admission)
	out.ConsoleIntegration = in.ConsoleIntegration
	in.Filtering.DeepCopyInto(&out.Filtering)
	in.Containers.DeepCopyInto(&out.Containers)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MondooAuditConfigSpec.
func (in *MondooAuditConfigSpec) DeepCopy() *MondooAuditConfigSpec {
	if in == nil {
		return nil
	}
	out := new(MondooAuditConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MondooAuditConfigStatus) DeepCopyInto(out *MondooAuditConfigStatus) {
	*out = *in
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]MondooAuditConfigCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MondooAuditConfigStatus.
func (in *MondooAuditConfigStatus) DeepCopy() *MondooAuditConfigStatus {
	if in == nil {
		return nil
	}
	out := new(MondooAuditConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MondooOperatorConfig) DeepCopyInto(out *MondooOperatorConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MondooOperatorConfig.
func (in *MondooOperatorConfig) DeepCopy() *MondooOperatorConfig {
	if in == nil {
		return nil
	}
	out := new(MondooOperatorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MondooOperatorConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MondooOperatorConfigCondition) DeepCopyInto(out *MondooOperatorConfigCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MondooOperatorConfigCondition.
func (in *MondooOperatorConfigCondition) DeepCopy() *MondooOperatorConfigCondition {
	if in == nil {
		return nil
	}
	out := new(MondooOperatorConfigCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MondooOperatorConfigList) DeepCopyInto(out *MondooOperatorConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MondooOperatorConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MondooOperatorConfigList.
func (in *MondooOperatorConfigList) DeepCopy() *MondooOperatorConfigList {
	if in == nil {
		return nil
	}
	out := new(MondooOperatorConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MondooOperatorConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MondooOperatorConfigSpec) DeepCopyInto(out *MondooOperatorConfigSpec) {
	*out = *in
	in.Metrics.DeepCopyInto(&out.Metrics)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MondooOperatorConfigSpec.
func (in *MondooOperatorConfigSpec) DeepCopy() *MondooOperatorConfigSpec {
	if in == nil {
		return nil
	}
	out := new(MondooOperatorConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MondooOperatorConfigStatus) DeepCopyInto(out *MondooOperatorConfigStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]MondooOperatorConfigCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MondooOperatorConfigStatus.
func (in *MondooOperatorConfigStatus) DeepCopy() *MondooOperatorConfigStatus {
	if in == nil {
		return nil
	}
	out := new(MondooOperatorConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Nodes) DeepCopyInto(out *Nodes) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Nodes.
func (in *Nodes) DeepCopy() *Nodes {
	if in == nil {
		return nil
	}
	out := new(Nodes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Scanner) DeepCopyInto(out *Scanner) {
	*out = *in
	out.Image = in.Image
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	out.PrivateRegistriesPullSecretRef = in.PrivateRegistriesPullSecretRef
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Scanner.
func (in *Scanner) DeepCopy() *Scanner {
	if in == nil {
		return nil
	}
	out := new(Scanner)
	in.DeepCopyInto(out)
	return out
}
