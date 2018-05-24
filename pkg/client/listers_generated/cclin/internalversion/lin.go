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

// This file was automatically generated by lister-gen

package internalversion

import (
	cclin "github.com/cclin81922/k8s-apiserver/pkg/apis/cclin"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LinLister helps list Lins.
type LinLister interface {
	// List lists all Lins in the indexer.
	List(selector labels.Selector) (ret []*cclin.Lin, err error)
	// Lins returns an object that can list and get Lins.
	Lins(namespace string) LinNamespaceLister
	LinListerExpansion
}

// linLister implements the LinLister interface.
type linLister struct {
	indexer cache.Indexer
}

// NewLinLister returns a new LinLister.
func NewLinLister(indexer cache.Indexer) LinLister {
	return &linLister{indexer: indexer}
}

// List lists all Lins in the indexer.
func (s *linLister) List(selector labels.Selector) (ret []*cclin.Lin, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*cclin.Lin))
	})
	return ret, err
}

// Lins returns an object that can list and get Lins.
func (s *linLister) Lins(namespace string) LinNamespaceLister {
	return linNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LinNamespaceLister helps list and get Lins.
type LinNamespaceLister interface {
	// List lists all Lins in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*cclin.Lin, err error)
	// Get retrieves the Lin from the indexer for a given namespace and name.
	Get(name string) (*cclin.Lin, error)
	LinNamespaceListerExpansion
}

// linNamespaceLister implements the LinNamespaceLister
// interface.
type linNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Lins in the indexer for a given namespace.
func (s linNamespaceLister) List(selector labels.Selector) (ret []*cclin.Lin, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*cclin.Lin))
	})
	return ret, err
}

// Get retrieves the Lin from the indexer for a given namespace and name.
func (s linNamespaceLister) Get(name string) (*cclin.Lin, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(cclin.Resource("lin"), name)
	}
	return obj.(*cclin.Lin), nil
}