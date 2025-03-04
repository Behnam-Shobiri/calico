// Copyright (c) 2025 Tigera, Inc. All rights reserved.

// Code generated by lister-gen. DO NOT EDIT.

package v3

import (
	v3 "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterInformationLister helps list ClusterInformations.
// All objects returned here must be treated as read-only.
type ClusterInformationLister interface {
	// List lists all ClusterInformations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v3.ClusterInformation, err error)
	// Get retrieves the ClusterInformation from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v3.ClusterInformation, error)
	ClusterInformationListerExpansion
}

// clusterInformationLister implements the ClusterInformationLister interface.
type clusterInformationLister struct {
	indexer cache.Indexer
}

// NewClusterInformationLister returns a new ClusterInformationLister.
func NewClusterInformationLister(indexer cache.Indexer) ClusterInformationLister {
	return &clusterInformationLister{indexer: indexer}
}

// List lists all ClusterInformations in the indexer.
func (s *clusterInformationLister) List(selector labels.Selector) (ret []*v3.ClusterInformation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.ClusterInformation))
	})
	return ret, err
}

// Get retrieves the ClusterInformation from the index for a given name.
func (s *clusterInformationLister) Get(name string) (*v3.ClusterInformation, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v3.Resource("clusterinformation"), name)
	}
	return obj.(*v3.ClusterInformation), nil
}
