/*
Copyright 2020 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/rancher/fleet/pkg/apis/fleet.cattle.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BundleDeploymentLister helps list BundleDeployments.
type BundleDeploymentLister interface {
	// List lists all BundleDeployments in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.BundleDeployment, err error)
	// BundleDeployments returns an object that can list and get BundleDeployments.
	BundleDeployments(namespace string) BundleDeploymentNamespaceLister
	BundleDeploymentListerExpansion
}

// bundleDeploymentLister implements the BundleDeploymentLister interface.
type bundleDeploymentLister struct {
	indexer cache.Indexer
}

// NewBundleDeploymentLister returns a new BundleDeploymentLister.
func NewBundleDeploymentLister(indexer cache.Indexer) BundleDeploymentLister {
	return &bundleDeploymentLister{indexer: indexer}
}

// List lists all BundleDeployments in the indexer.
func (s *bundleDeploymentLister) List(selector labels.Selector) (ret []*v1alpha1.BundleDeployment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BundleDeployment))
	})
	return ret, err
}

// BundleDeployments returns an object that can list and get BundleDeployments.
func (s *bundleDeploymentLister) BundleDeployments(namespace string) BundleDeploymentNamespaceLister {
	return bundleDeploymentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BundleDeploymentNamespaceLister helps list and get BundleDeployments.
type BundleDeploymentNamespaceLister interface {
	// List lists all BundleDeployments in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.BundleDeployment, err error)
	// Get retrieves the BundleDeployment from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.BundleDeployment, error)
	BundleDeploymentNamespaceListerExpansion
}

// bundleDeploymentNamespaceLister implements the BundleDeploymentNamespaceLister
// interface.
type bundleDeploymentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BundleDeployments in the indexer for a given namespace.
func (s bundleDeploymentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BundleDeployment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BundleDeployment))
	})
	return ret, err
}

// Get retrieves the BundleDeployment from the indexer for a given namespace and name.
func (s bundleDeploymentNamespaceLister) Get(name string) (*v1alpha1.BundleDeployment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("bundledeployment"), name)
	}
	return obj.(*v1alpha1.BundleDeployment), nil
}
