/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/LeoNumber1/k8s-controller-custom-resource/pkg/apis/samplecrd/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NetworkLister helps list Networks.
type NetworkLister interface {
	// List lists all Networks in the indexer.
	List(selector labels.Selector) (ret []*v1.Network, err error)
	// Networks returns an object that can list and get Networks.
	Networks(namespace string) NetworkNamespaceLister
	NetworkListerExpansion
}

// networkLister implements the NetworkLister interface.
type networkLister struct {
	indexer cache.Indexer
}

// NewNetworkLister returns a new NetworkLister.
func NewNetworkLister(indexer cache.Indexer) NetworkLister {
	return &networkLister{indexer: indexer}
}

// List lists all Networks in the indexer.
func (s *networkLister) List(selector labels.Selector) (ret []*v1.Network, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Network))
	})
	return ret, err
}

// Networks returns an object that can list and get Networks.
func (s *networkLister) Networks(namespace string) NetworkNamespaceLister {
	return networkNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NetworkNamespaceLister helps list and get Networks.
type NetworkNamespaceLister interface {
	// List lists all Networks in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.Network, err error)
	// Get retrieves the Network from the indexer for a given namespace and name.
	Get(name string) (*v1.Network, error)
	NetworkNamespaceListerExpansion
}

// networkNamespaceLister implements the NetworkNamespaceLister
// interface.
type networkNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Networks in the indexer for a given namespace.
func (s networkNamespaceLister) List(selector labels.Selector) (ret []*v1.Network, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Network))
	})
	return ret, err
}

// Get retrieves the Network from the indexer for a given namespace and name.
func (s networkNamespaceLister) Get(name string) (*v1.Network, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("network"), name)
	}
	return obj.(*v1.Network), nil
}
