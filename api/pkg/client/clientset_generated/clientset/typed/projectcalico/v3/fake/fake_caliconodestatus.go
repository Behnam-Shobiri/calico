// Copyright (c) 2025 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v3 "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCalicoNodeStatuses implements CalicoNodeStatusInterface
type FakeCalicoNodeStatuses struct {
	Fake *FakeProjectcalicoV3
}

var caliconodestatusesResource = v3.SchemeGroupVersion.WithResource("caliconodestatuses")

var caliconodestatusesKind = v3.SchemeGroupVersion.WithKind("CalicoNodeStatus")

// Get takes name of the calicoNodeStatus, and returns the corresponding calicoNodeStatus object, and an error if there is any.
func (c *FakeCalicoNodeStatuses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.CalicoNodeStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(caliconodestatusesResource, name), &v3.CalicoNodeStatus{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.CalicoNodeStatus), err
}

// List takes label and field selectors, and returns the list of CalicoNodeStatuses that match those selectors.
func (c *FakeCalicoNodeStatuses) List(ctx context.Context, opts v1.ListOptions) (result *v3.CalicoNodeStatusList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(caliconodestatusesResource, caliconodestatusesKind, opts), &v3.CalicoNodeStatusList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.CalicoNodeStatusList{ListMeta: obj.(*v3.CalicoNodeStatusList).ListMeta}
	for _, item := range obj.(*v3.CalicoNodeStatusList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested calicoNodeStatuses.
func (c *FakeCalicoNodeStatuses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(caliconodestatusesResource, opts))
}

// Create takes the representation of a calicoNodeStatus and creates it.  Returns the server's representation of the calicoNodeStatus, and an error, if there is any.
func (c *FakeCalicoNodeStatuses) Create(ctx context.Context, calicoNodeStatus *v3.CalicoNodeStatus, opts v1.CreateOptions) (result *v3.CalicoNodeStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(caliconodestatusesResource, calicoNodeStatus), &v3.CalicoNodeStatus{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.CalicoNodeStatus), err
}

// Update takes the representation of a calicoNodeStatus and updates it. Returns the server's representation of the calicoNodeStatus, and an error, if there is any.
func (c *FakeCalicoNodeStatuses) Update(ctx context.Context, calicoNodeStatus *v3.CalicoNodeStatus, opts v1.UpdateOptions) (result *v3.CalicoNodeStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(caliconodestatusesResource, calicoNodeStatus), &v3.CalicoNodeStatus{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.CalicoNodeStatus), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCalicoNodeStatuses) UpdateStatus(ctx context.Context, calicoNodeStatus *v3.CalicoNodeStatus, opts v1.UpdateOptions) (*v3.CalicoNodeStatus, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(caliconodestatusesResource, "status", calicoNodeStatus), &v3.CalicoNodeStatus{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.CalicoNodeStatus), err
}

// Delete takes name of the calicoNodeStatus and deletes it. Returns an error if one occurs.
func (c *FakeCalicoNodeStatuses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(caliconodestatusesResource, name, opts), &v3.CalicoNodeStatus{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCalicoNodeStatuses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(caliconodestatusesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v3.CalicoNodeStatusList{})
	return err
}

// Patch applies the patch and returns the patched calicoNodeStatus.
func (c *FakeCalicoNodeStatuses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.CalicoNodeStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(caliconodestatusesResource, name, pt, data, subresources...), &v3.CalicoNodeStatus{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.CalicoNodeStatus), err
}
