
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
	"log"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/endpoints/request"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/cclin81922/k8s-apiserver/pkg/apis/cclin"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Lin
// +k8s:openapi-gen=true
// +resource:path=lins,strategy=LinStrategy
type Lin struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LinSpec   `json:"spec,omitempty"`
	Status LinStatus `json:"status,omitempty"`
}

// LinSpec defines the desired state of Lin
type LinSpec struct {
}

// LinStatus defines the observed state of Lin
type LinStatus struct {
}

// Validate checks that an instance of Lin is well formed
func (LinStrategy) Validate(ctx request.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*cclin.Lin)
	log.Printf("Validating fields for Lin %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// DefaultingFunction sets default Lin field values
func (LinSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*Lin)
	// set default field values here
	log.Printf("Defaulting fields for Lin %s\n", obj.Name)
}
