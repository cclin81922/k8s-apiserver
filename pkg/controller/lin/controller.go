
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


package lin

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"

	"github.com/cclin81922/k8s-apiserver/pkg/apis/cclin/v1"
	"github.com/cclin81922/k8s-apiserver/pkg/controller/sharedinformers"
	listers "github.com/cclin81922/k8s-apiserver/pkg/client/listers_generated/cclin/v1"
)

// +controller:group=cclin,version=v1,kind=Lin,resource=lins
type LinControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about Lin
	lister listers.LinLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *LinControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing lins labels
	c.lister = arguments.GetSharedInformers().Factory.Cclin().V1().Lins().Lister()
}

// Reconcile handles enqueued messages
func (c *LinControllerImpl) Reconcile(u *v1.Lin) error {
	// Implement controller logic here
	log.Printf("Running reconcile Lin for %s\n", u.Name)
	return nil
}

func (c *LinControllerImpl) Get(namespace, name string) (*v1.Lin, error) {
	return c.lister.Lins(namespace).Get(name)
}
