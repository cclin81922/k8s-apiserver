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
package v1

import (
	v1 "github.com/cclin81922/k8s-apiserver/pkg/apis/cclin/v1"
	scheme "github.com/cclin81922/k8s-apiserver/pkg/client/clientset_generated/clientset/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// LinsGetter has a method to return a LinInterface.
// A group's client should implement this interface.
type LinsGetter interface {
	Lins(namespace string) LinInterface
}

// LinInterface has methods to work with Lin resources.
type LinInterface interface {
	Create(*v1.Lin) (*v1.Lin, error)
	Update(*v1.Lin) (*v1.Lin, error)
	UpdateStatus(*v1.Lin) (*v1.Lin, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.Lin, error)
	List(opts meta_v1.ListOptions) (*v1.LinList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Lin, err error)
	LinExpansion
}

// lins implements LinInterface
type lins struct {
	client rest.Interface
	ns     string
}

// newLins returns a Lins
func newLins(c *CclinV1Client, namespace string) *lins {
	return &lins{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the lin, and returns the corresponding lin object, and an error if there is any.
func (c *lins) Get(name string, options meta_v1.GetOptions) (result *v1.Lin, err error) {
	result = &v1.Lin{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("lins").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Lins that match those selectors.
func (c *lins) List(opts meta_v1.ListOptions) (result *v1.LinList, err error) {
	result = &v1.LinList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("lins").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested lins.
func (c *lins) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("lins").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a lin and creates it.  Returns the server's representation of the lin, and an error, if there is any.
func (c *lins) Create(lin *v1.Lin) (result *v1.Lin, err error) {
	result = &v1.Lin{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("lins").
		Body(lin).
		Do().
		Into(result)
	return
}

// Update takes the representation of a lin and updates it. Returns the server's representation of the lin, and an error, if there is any.
func (c *lins) Update(lin *v1.Lin) (result *v1.Lin, err error) {
	result = &v1.Lin{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("lins").
		Name(lin.Name).
		Body(lin).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *lins) UpdateStatus(lin *v1.Lin) (result *v1.Lin, err error) {
	result = &v1.Lin{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("lins").
		Name(lin.Name).
		SubResource("status").
		Body(lin).
		Do().
		Into(result)
	return
}

// Delete takes name of the lin and deletes it. Returns an error if one occurs.
func (c *lins) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("lins").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *lins) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("lins").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched lin.
func (c *lins) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Lin, err error) {
	result = &v1.Lin{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("lins").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
