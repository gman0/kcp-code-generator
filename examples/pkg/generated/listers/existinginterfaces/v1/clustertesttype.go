/*
Copyright The KCP Authors.

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

// Code generated by lister-gen-v0.32. DO NOT EDIT.

package v1

import (
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"

	existinginterfacesv1 "acme.corp/pkg/apis/existinginterfaces/v1"
)

// ClusterTestTypeLister helps list ClusterTestTypes.
// All objects returned here must be treated as read-only.
type ClusterTestTypeLister interface {
	// List lists all ClusterTestTypes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*existinginterfacesv1.ClusterTestType, err error)
	// Get retrieves the ClusterTestType from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*existinginterfacesv1.ClusterTestType, error)
	ClusterTestTypeListerExpansion
}

// clusterTestTypeLister implements the ClusterTestTypeLister interface.
type clusterTestTypeLister struct {
	listers.ResourceIndexer[*existinginterfacesv1.ClusterTestType]
}

// NewClusterTestTypeLister returns a new ClusterTestTypeLister.
func NewClusterTestTypeLister(indexer cache.Indexer) ClusterTestTypeLister {
	return &clusterTestTypeLister{listers.New[*existinginterfacesv1.ClusterTestType](indexer, existinginterfacesv1.Resource("clustertesttype"))}
}
