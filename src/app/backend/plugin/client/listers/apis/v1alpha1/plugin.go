// Copyright 2017 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/lixianyang/dashboard/v2/src/app/backend/plugin/apis/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PluginLister helps list Plugins.
// All objects returned here must be treated as read-only.
type PluginLister interface {
	// List lists all Plugins in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Plugin, err error)
	// Plugins returns an object that can list and get Plugins.
	Plugins(namespace string) PluginNamespaceLister
	PluginListerExpansion
}

// pluginLister implements the PluginLister interface.
type pluginLister struct {
	indexer cache.Indexer
}

// NewPluginLister returns a new PluginLister.
func NewPluginLister(indexer cache.Indexer) PluginLister {
	return &pluginLister{indexer: indexer}
}

// List lists all Plugins in the indexer.
func (s *pluginLister) List(selector labels.Selector) (ret []*v1alpha1.Plugin, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Plugin))
	})
	return ret, err
}

// Plugins returns an object that can list and get Plugins.
func (s *pluginLister) Plugins(namespace string) PluginNamespaceLister {
	return pluginNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PluginNamespaceLister helps list and get Plugins.
// All objects returned here must be treated as read-only.
type PluginNamespaceLister interface {
	// List lists all Plugins in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Plugin, err error)
	// Get retrieves the Plugin from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Plugin, error)
	PluginNamespaceListerExpansion
}

// pluginNamespaceLister implements the PluginNamespaceLister
// interface.
type pluginNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Plugins in the indexer for a given namespace.
func (s pluginNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Plugin, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Plugin))
	})
	return ret, err
}

// Get retrieves the Plugin from the indexer for a given namespace and name.
func (s pluginNamespaceLister) Get(name string) (*v1alpha1.Plugin, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("plugin"), name)
	}
	return obj.(*v1alpha1.Plugin), nil
}
