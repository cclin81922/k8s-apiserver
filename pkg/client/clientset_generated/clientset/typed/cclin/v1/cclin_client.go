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
	"github.com/cclin81922/k8s-apiserver/pkg/client/clientset_generated/clientset/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type CclinV1Interface interface {
	RESTClient() rest.Interface
	LinsGetter
}

// CclinV1Client is used to interact with features provided by the cclin.cclin group.
type CclinV1Client struct {
	restClient rest.Interface
}

func (c *CclinV1Client) Lins(namespace string) LinInterface {
	return newLins(c, namespace)
}

// NewForConfig creates a new CclinV1Client for the given config.
func NewForConfig(c *rest.Config) (*CclinV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &CclinV1Client{client}, nil
}

// NewForConfigOrDie creates a new CclinV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *CclinV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new CclinV1Client for the given RESTClient.
func New(c rest.Interface) *CclinV1Client {
	return &CclinV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *CclinV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
