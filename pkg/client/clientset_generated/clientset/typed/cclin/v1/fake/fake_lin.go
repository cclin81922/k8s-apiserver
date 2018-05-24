/*
Copyright 2018 CC LIN.

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
package fake

import (
	cclin_v1 "github.com/cclin81922/k8s-apiserver/pkg/apis/cclin/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLins implements LinInterface
type FakeLins struct {
	Fake *FakeCclinV1
	ns   string
}

var linsResource = schema.GroupVersionResource{Group: "cclin.cclin", Version: "v1", Resource: "lins"}

var linsKind = schema.GroupVersionKind{Group: "cclin.cclin", Version: "v1", Kind: "Lin"}

// Get takes name of the lin, and returns the corresponding lin object, and an error if there is any.
func (c *FakeLins) Get(name string, options v1.GetOptions) (result *cclin_v1.Lin, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(linsResource, c.ns, name), &cclin_v1.Lin{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cclin_v1.Lin), err
}

// List takes label and field selectors, and returns the list of Lins that match those selectors.
func (c *FakeLins) List(opts v1.ListOptions) (result *cclin_v1.LinList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(linsResource, linsKind, c.ns, opts), &cclin_v1.LinList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &cclin_v1.LinList{}
	for _, item := range obj.(*cclin_v1.LinList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested lins.
func (c *FakeLins) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(linsResource, c.ns, opts))

}

// Create takes the representation of a lin and creates it.  Returns the server's representation of the lin, and an error, if there is any.
func (c *FakeLins) Create(lin *cclin_v1.Lin) (result *cclin_v1.Lin, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(linsResource, c.ns, lin), &cclin_v1.Lin{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cclin_v1.Lin), err
}

// Update takes the representation of a lin and updates it. Returns the server's representation of the lin, and an error, if there is any.
func (c *FakeLins) Update(lin *cclin_v1.Lin) (result *cclin_v1.Lin, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(linsResource, c.ns, lin), &cclin_v1.Lin{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cclin_v1.Lin), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLins) UpdateStatus(lin *cclin_v1.Lin) (*cclin_v1.Lin, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(linsResource, "status", c.ns, lin), &cclin_v1.Lin{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cclin_v1.Lin), err
}

// Delete takes name of the lin and deletes it. Returns an error if one occurs.
func (c *FakeLins) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(linsResource, c.ns, name), &cclin_v1.Lin{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLins) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(linsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &cclin_v1.LinList{})
	return err
}

// Patch applies the patch and returns the patched lin.
func (c *FakeLins) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *cclin_v1.Lin, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(linsResource, c.ns, name, data, subresources...), &cclin_v1.Lin{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cclin_v1.Lin), err
}
