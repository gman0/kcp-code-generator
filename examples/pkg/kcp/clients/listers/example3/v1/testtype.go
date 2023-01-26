//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by kcp code-generator. DO NOT EDIT.

package v1

import (
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	example3v1 "acme.corp/pkg/apis/example3/v1"
)

// TestTypeClusterLister can list TestTypes across all workspaces, or scope down to a TestTypeLister for one workspace.
// All objects returned here must be treated as read-only.
type TestTypeClusterLister interface {
	// List lists all TestTypes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*example3v1.TestType, err error)
	// Cluster returns a lister that can list and get TestTypes in one workspace.
	Cluster(clusterName logicalcluster.Name) TestTypeLister
	TestTypeClusterListerExpansion
}

type testTypeClusterLister struct {
	indexer cache.Indexer
}

// NewTestTypeClusterLister returns a new TestTypeClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
// - has the kcpcache.ClusterAndNamespaceIndex as an index
func NewTestTypeClusterLister(indexer cache.Indexer) *testTypeClusterLister {
	return &testTypeClusterLister{indexer: indexer}
}

// List lists all TestTypes in the indexer across all workspaces.
func (s *testTypeClusterLister) List(selector labels.Selector) (ret []*example3v1.TestType, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*example3v1.TestType))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get TestTypes.
func (s *testTypeClusterLister) Cluster(clusterName logicalcluster.Name) TestTypeLister {
	return &testTypeLister{indexer: s.indexer, clusterName: clusterName}
}

// TestTypeLister can list TestTypes across all namespaces, or scope down to a TestTypeNamespaceLister for one namespace.
// All objects returned here must be treated as read-only.
type TestTypeLister interface {
	// List lists all TestTypes in the workspace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*example3v1.TestType, err error)
	// TestTypes returns a lister that can list and get TestTypes in one workspace and namespace.
	TestTypes(namespace string) TestTypeNamespaceLister
	TestTypeListerExpansion
}

// testTypeLister can list all TestTypes inside a workspace or scope down to a TestTypeLister for one namespace.
type testTypeLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all TestTypes in the indexer for a workspace.
func (s *testTypeLister) List(selector labels.Selector) (ret []*example3v1.TestType, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*example3v1.TestType))
	})
	return ret, err
}

// TestTypes returns an object that can list and get TestTypes in one namespace.
func (s *testTypeLister) TestTypes(namespace string) TestTypeNamespaceLister {
	return &testTypeNamespaceLister{indexer: s.indexer, clusterName: s.clusterName, namespace: namespace}
}

// testTypeNamespaceLister helps list and get TestTypes.
// All objects returned here must be treated as read-only.
type TestTypeNamespaceLister interface {
	// List lists all TestTypes in the workspace and namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*example3v1.TestType, err error)
	// Get retrieves the TestType from the indexer for a given workspace, namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*example3v1.TestType, error)
	TestTypeNamespaceListerExpansion
}

// testTypeNamespaceLister helps list and get TestTypes.
// All objects returned here must be treated as read-only.
type testTypeNamespaceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
	namespace   string
}

// List lists all TestTypes in the indexer for a given workspace and namespace.
func (s *testTypeNamespaceLister) List(selector labels.Selector) (ret []*example3v1.TestType, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.clusterName, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*example3v1.TestType))
	})
	return ret, err
}

// Get retrieves the TestType from the indexer for a given workspace, namespace and name.
func (s *testTypeNamespaceLister) Get(name string) (*example3v1.TestType, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(example3v1.Resource("testtypes"), name)
	}
	return obj.(*example3v1.TestType), nil
}

// NewTestTypeLister returns a new TestTypeLister.
// We assume that the indexer:
// - is fed by a workspace-scoped LIST+WATCH
// - uses cache.MetaNamespaceKeyFunc as the key function
// - has the cache.NamespaceIndex as an index
func NewTestTypeLister(indexer cache.Indexer) *testTypeScopedLister {
	return &testTypeScopedLister{indexer: indexer}
}

// testTypeScopedLister can list all TestTypes inside a workspace or scope down to a TestTypeLister for one namespace.
type testTypeScopedLister struct {
	indexer cache.Indexer
}

// List lists all TestTypes in the indexer for a workspace.
func (s *testTypeScopedLister) List(selector labels.Selector) (ret []*example3v1.TestType, err error) {
	err = cache.ListAll(s.indexer, selector, func(i interface{}) {
		ret = append(ret, i.(*example3v1.TestType))
	})
	return ret, err
}

// TestTypes returns an object that can list and get TestTypes in one namespace.
func (s *testTypeScopedLister) TestTypes(namespace string) TestTypeNamespaceLister {
	return &testTypeScopedNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// testTypeScopedNamespaceLister helps list and get TestTypes.
type testTypeScopedNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TestTypes in the indexer for a given workspace and namespace.
func (s *testTypeScopedNamespaceLister) List(selector labels.Selector) (ret []*example3v1.TestType, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*example3v1.TestType))
	})
	return ret, err
}

// Get retrieves the TestType from the indexer for a given workspace, namespace and name.
func (s *testTypeScopedNamespaceLister) Get(name string) (*example3v1.TestType, error) {
	key := s.namespace + "/" + name
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(example3v1.Resource("testtypes"), name)
	}
	return obj.(*example3v1.TestType), nil
}
