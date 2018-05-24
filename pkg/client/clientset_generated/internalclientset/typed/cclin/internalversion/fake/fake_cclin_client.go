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
	internalversion "github.com/cclin81922/k8s-apiserver/pkg/client/clientset_generated/internalclientset/typed/cclin/internalversion"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCclin struct {
	*testing.Fake
}

func (c *FakeCclin) Lins(namespace string) internalversion.LinInterface {
	return &FakeLins{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCclin) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
