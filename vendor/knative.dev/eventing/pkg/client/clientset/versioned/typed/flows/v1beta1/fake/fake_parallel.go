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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1beta1 "knative.dev/eventing/pkg/apis/flows/v1beta1"
)

// FakeParallels implements ParallelInterface
type FakeParallels struct {
	Fake *FakeFlowsV1beta1
	ns   string
}

var parallelsResource = schema.GroupVersionResource{Group: "flows.knative.dev", Version: "v1beta1", Resource: "parallels"}

var parallelsKind = schema.GroupVersionKind{Group: "flows.knative.dev", Version: "v1beta1", Kind: "Parallel"}

// Get takes name of the parallel, and returns the corresponding parallel object, and an error if there is any.
func (c *FakeParallels) Get(name string, options v1.GetOptions) (result *v1beta1.Parallel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(parallelsResource, c.ns, name), &v1beta1.Parallel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Parallel), err
}

// List takes label and field selectors, and returns the list of Parallels that match those selectors.
func (c *FakeParallels) List(opts v1.ListOptions) (result *v1beta1.ParallelList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(parallelsResource, parallelsKind, c.ns, opts), &v1beta1.ParallelList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ParallelList{ListMeta: obj.(*v1beta1.ParallelList).ListMeta}
	for _, item := range obj.(*v1beta1.ParallelList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested parallels.
func (c *FakeParallels) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(parallelsResource, c.ns, opts))

}

// Create takes the representation of a parallel and creates it.  Returns the server's representation of the parallel, and an error, if there is any.
func (c *FakeParallels) Create(parallel *v1beta1.Parallel) (result *v1beta1.Parallel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(parallelsResource, c.ns, parallel), &v1beta1.Parallel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Parallel), err
}

// Update takes the representation of a parallel and updates it. Returns the server's representation of the parallel, and an error, if there is any.
func (c *FakeParallels) Update(parallel *v1beta1.Parallel) (result *v1beta1.Parallel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(parallelsResource, c.ns, parallel), &v1beta1.Parallel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Parallel), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeParallels) UpdateStatus(parallel *v1beta1.Parallel) (*v1beta1.Parallel, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(parallelsResource, "status", c.ns, parallel), &v1beta1.Parallel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Parallel), err
}

// Delete takes name of the parallel and deletes it. Returns an error if one occurs.
func (c *FakeParallels) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(parallelsResource, c.ns, name), &v1beta1.Parallel{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeParallels) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(parallelsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.ParallelList{})
	return err
}

// Patch applies the patch and returns the patched parallel.
func (c *FakeParallels) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Parallel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(parallelsResource, c.ns, name, pt, data, subresources...), &v1beta1.Parallel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Parallel), err
}
