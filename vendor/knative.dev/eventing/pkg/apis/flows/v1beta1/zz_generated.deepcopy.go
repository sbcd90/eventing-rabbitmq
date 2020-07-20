// +build !ignore_autogenerated

/*
Copyright 2020 The Knative Authors

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	duckv1beta1 "knative.dev/eventing/pkg/apis/duck/v1beta1"
	messagingv1beta1 "knative.dev/eventing/pkg/apis/messaging/v1beta1"
	v1 "knative.dev/pkg/apis/duck/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Parallel) DeepCopyInto(out *Parallel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Parallel.
func (in *Parallel) DeepCopy() *Parallel {
	if in == nil {
		return nil
	}
	out := new(Parallel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Parallel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParallelBranch) DeepCopyInto(out *ParallelBranch) {
	*out = *in
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(v1.Destination)
		(*in).DeepCopyInto(*out)
	}
	in.Subscriber.DeepCopyInto(&out.Subscriber)
	if in.Reply != nil {
		in, out := &in.Reply, &out.Reply
		*out = new(v1.Destination)
		(*in).DeepCopyInto(*out)
	}
	if in.Delivery != nil {
		in, out := &in.Delivery, &out.Delivery
		*out = new(duckv1beta1.DeliverySpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParallelBranch.
func (in *ParallelBranch) DeepCopy() *ParallelBranch {
	if in == nil {
		return nil
	}
	out := new(ParallelBranch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParallelBranchStatus) DeepCopyInto(out *ParallelBranchStatus) {
	*out = *in
	in.FilterSubscriptionStatus.DeepCopyInto(&out.FilterSubscriptionStatus)
	in.FilterChannelStatus.DeepCopyInto(&out.FilterChannelStatus)
	in.SubscriptionStatus.DeepCopyInto(&out.SubscriptionStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParallelBranchStatus.
func (in *ParallelBranchStatus) DeepCopy() *ParallelBranchStatus {
	if in == nil {
		return nil
	}
	out := new(ParallelBranchStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParallelChannelStatus) DeepCopyInto(out *ParallelChannelStatus) {
	*out = *in
	out.Channel = in.Channel
	in.ReadyCondition.DeepCopyInto(&out.ReadyCondition)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParallelChannelStatus.
func (in *ParallelChannelStatus) DeepCopy() *ParallelChannelStatus {
	if in == nil {
		return nil
	}
	out := new(ParallelChannelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParallelList) DeepCopyInto(out *ParallelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Parallel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParallelList.
func (in *ParallelList) DeepCopy() *ParallelList {
	if in == nil {
		return nil
	}
	out := new(ParallelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ParallelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParallelSpec) DeepCopyInto(out *ParallelSpec) {
	*out = *in
	if in.Branches != nil {
		in, out := &in.Branches, &out.Branches
		*out = make([]ParallelBranch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ChannelTemplate != nil {
		in, out := &in.ChannelTemplate, &out.ChannelTemplate
		*out = new(messagingv1beta1.ChannelTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Reply != nil {
		in, out := &in.Reply, &out.Reply
		*out = new(v1.Destination)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParallelSpec.
func (in *ParallelSpec) DeepCopy() *ParallelSpec {
	if in == nil {
		return nil
	}
	out := new(ParallelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParallelStatus) DeepCopyInto(out *ParallelStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	in.IngressChannelStatus.DeepCopyInto(&out.IngressChannelStatus)
	if in.BranchStatuses != nil {
		in, out := &in.BranchStatuses, &out.BranchStatuses
		*out = make([]ParallelBranchStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.AddressStatus.DeepCopyInto(&out.AddressStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParallelStatus.
func (in *ParallelStatus) DeepCopy() *ParallelStatus {
	if in == nil {
		return nil
	}
	out := new(ParallelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParallelSubscriptionStatus) DeepCopyInto(out *ParallelSubscriptionStatus) {
	*out = *in
	out.Subscription = in.Subscription
	in.ReadyCondition.DeepCopyInto(&out.ReadyCondition)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParallelSubscriptionStatus.
func (in *ParallelSubscriptionStatus) DeepCopy() *ParallelSubscriptionStatus {
	if in == nil {
		return nil
	}
	out := new(ParallelSubscriptionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sequence) DeepCopyInto(out *Sequence) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sequence.
func (in *Sequence) DeepCopy() *Sequence {
	if in == nil {
		return nil
	}
	out := new(Sequence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Sequence) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SequenceChannelStatus) DeepCopyInto(out *SequenceChannelStatus) {
	*out = *in
	out.Channel = in.Channel
	in.ReadyCondition.DeepCopyInto(&out.ReadyCondition)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SequenceChannelStatus.
func (in *SequenceChannelStatus) DeepCopy() *SequenceChannelStatus {
	if in == nil {
		return nil
	}
	out := new(SequenceChannelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SequenceList) DeepCopyInto(out *SequenceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Sequence, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SequenceList.
func (in *SequenceList) DeepCopy() *SequenceList {
	if in == nil {
		return nil
	}
	out := new(SequenceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SequenceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SequenceSpec) DeepCopyInto(out *SequenceSpec) {
	*out = *in
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]SequenceStep, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ChannelTemplate != nil {
		in, out := &in.ChannelTemplate, &out.ChannelTemplate
		*out = new(messagingv1beta1.ChannelTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Reply != nil {
		in, out := &in.Reply, &out.Reply
		*out = new(v1.Destination)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SequenceSpec.
func (in *SequenceSpec) DeepCopy() *SequenceSpec {
	if in == nil {
		return nil
	}
	out := new(SequenceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SequenceStatus) DeepCopyInto(out *SequenceStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.SubscriptionStatuses != nil {
		in, out := &in.SubscriptionStatuses, &out.SubscriptionStatuses
		*out = make([]SequenceSubscriptionStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ChannelStatuses != nil {
		in, out := &in.ChannelStatuses, &out.ChannelStatuses
		*out = make([]SequenceChannelStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.AddressStatus.DeepCopyInto(&out.AddressStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SequenceStatus.
func (in *SequenceStatus) DeepCopy() *SequenceStatus {
	if in == nil {
		return nil
	}
	out := new(SequenceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SequenceStep) DeepCopyInto(out *SequenceStep) {
	*out = *in
	in.Destination.DeepCopyInto(&out.Destination)
	if in.Delivery != nil {
		in, out := &in.Delivery, &out.Delivery
		*out = new(duckv1beta1.DeliverySpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SequenceStep.
func (in *SequenceStep) DeepCopy() *SequenceStep {
	if in == nil {
		return nil
	}
	out := new(SequenceStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SequenceSubscriptionStatus) DeepCopyInto(out *SequenceSubscriptionStatus) {
	*out = *in
	out.Subscription = in.Subscription
	in.ReadyCondition.DeepCopyInto(&out.ReadyCondition)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SequenceSubscriptionStatus.
func (in *SequenceSubscriptionStatus) DeepCopy() *SequenceSubscriptionStatus {
	if in == nil {
		return nil
	}
	out := new(SequenceSubscriptionStatus)
	in.DeepCopyInto(out)
	return out
}
