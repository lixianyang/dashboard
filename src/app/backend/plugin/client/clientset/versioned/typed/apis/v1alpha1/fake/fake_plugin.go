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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/lixianyang/dashboard/v2/src/app/backend/plugin/apis/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePlugins implements PluginInterface
type FakePlugins struct {
	Fake *FakeDashboardV1alpha1
	ns   string
}

var pluginsResource = schema.GroupVersionResource{Group: "dashboard.k8s.io", Version: "v1alpha1", Resource: "plugins"}

var pluginsKind = schema.GroupVersionKind{Group: "dashboard.k8s.io", Version: "v1alpha1", Kind: "Plugin"}

// Get takes name of the plugin, and returns the corresponding plugin object, and an error if there is any.
func (c *FakePlugins) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Plugin, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pluginsResource, c.ns, name), &v1alpha1.Plugin{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Plugin), err
}

// List takes label and field selectors, and returns the list of Plugins that match those selectors.
func (c *FakePlugins) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PluginList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pluginsResource, pluginsKind, c.ns, opts), &v1alpha1.PluginList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PluginList{ListMeta: obj.(*v1alpha1.PluginList).ListMeta}
	for _, item := range obj.(*v1alpha1.PluginList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested plugins.
func (c *FakePlugins) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pluginsResource, c.ns, opts))

}

// Create takes the representation of a plugin and creates it.  Returns the server's representation of the plugin, and an error, if there is any.
func (c *FakePlugins) Create(ctx context.Context, plugin *v1alpha1.Plugin, opts v1.CreateOptions) (result *v1alpha1.Plugin, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pluginsResource, c.ns, plugin), &v1alpha1.Plugin{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Plugin), err
}

// Update takes the representation of a plugin and updates it. Returns the server's representation of the plugin, and an error, if there is any.
func (c *FakePlugins) Update(ctx context.Context, plugin *v1alpha1.Plugin, opts v1.UpdateOptions) (result *v1alpha1.Plugin, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pluginsResource, c.ns, plugin), &v1alpha1.Plugin{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Plugin), err
}

// Delete takes name of the plugin and deletes it. Returns an error if one occurs.
func (c *FakePlugins) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pluginsResource, c.ns, name), &v1alpha1.Plugin{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePlugins) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pluginsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PluginList{})
	return err
}

// Patch applies the patch and returns the patched plugin.
func (c *FakePlugins) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Plugin, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pluginsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Plugin{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Plugin), err
}
