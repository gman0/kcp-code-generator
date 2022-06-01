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
	"context"
	"fmt"
	exampleapiv1 "github.com/kcp-dev/code-generator/examples/pkg/apis/example/v1"
	examplev1 "github.com/kcp-dev/code-generator/examples/pkg/generated/clientset/versioned/typed/example/v1"

	kcp "github.com/kcp-dev/apimachinery/pkg/client"
	"github.com/kcp-dev/logicalcluster"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
)

// WrappedExampleV1 wraps the client interface with a
// logical cluster.
type WrappedExampleV1 struct {
	cluster  logicalcluster.Name
	delegate examplev1.ExampleV1Interface
}

// New creates a WrappedExampleV1 with the given logical cluster and client interface.
func New(cluster logicalcluster.Name, delegate examplev1.ExampleV1Interface) *WrappedExampleV1 {
	return &WrappedExampleV1{cluster: cluster, delegate: delegate}
}

// RESTClient returns the underlying RESTClient.
func (w *WrappedExampleV1) RESTClient() rest.Interface {
	return w.delegate.RESTClient()
}

// WrappedExampleV1 contains the wrapped logical cluster and interface.
func (w *WrappedExampleV1) ClusterTestTypes() examplev1.ClusterTestTypeInterface {
	return &wrappedClusterTestType{
		cluster:  w.cluster,
		delegate: w.delegate.ClusterTestTypes(),
	}
}

type wrappedClusterTestType struct {
	cluster  logicalcluster.Name
	delegate examplev1.ClusterTestTypeInterface
}

// checkCluster retrieves the logical cluster name from the given context and checks
// if it is the same as the one passed while creating a wrappedClusterTestType. It errors when
// there is a mismatch.
func (w *wrappedClusterTestType) checkCluster(ctx context.Context) (context.Context, error) {
	ctxCluster, ok := kcp.ClusterFromContext(ctx)
	if !ok {
		return kcp.WithCluster(ctx, w.cluster), nil
	} else if ctxCluster != w.cluster {
		return ctx, fmt.Errorf("cluster mismatch: context=%q, client=%q", ctxCluster, w.cluster)
	}
	return ctx, nil
}

// Create implements ClusterTestTypeInterface.
func (w *wrappedClusterTestType) Create(ctx context.Context, clusterTestType *exampleapiv1.ClusterTestType, opts metav1.CreateOptions) (*exampleapiv1.ClusterTestType, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Create(ctx, clusterTestType, opts)
}

// Update implements ClusterTestTypeInterface.
func (w *wrappedClusterTestType) Update(ctx context.Context, clusterTestType *exampleapiv1.ClusterTestType, opts metav1.UpdateOptions) (*exampleapiv1.ClusterTestType, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Update(ctx, clusterTestType, opts)
}

// UpdateStatus implements ClusterTestTypeInterface. It was generated because the type contains a Status member.
func (w *wrappedClusterTestType) UpdateStatus(ctx context.Context, clusterTestType *exampleapiv1.ClusterTestType, opts metav1.UpdateOptions) (*exampleapiv1.ClusterTestType, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.UpdateStatus(ctx, clusterTestType, opts)
}

// Update implements ClusterTestTypeInterface.
func (w *wrappedClusterTestType) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return err
	}
	return w.delegate.Delete(ctx, name, opts)
}

// DeleteCollection implements ClusterTestTypeInterface.
func (w *wrappedClusterTestType) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listopts metav1.ListOptions) error {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return err
	}
	return w.delegate.DeleteCollection(ctx, opts, listopts)
}

// Get implements ClusterTestTypeInterface.
func (w *wrappedClusterTestType) Get(ctx context.Context, name string, opts metav1.GetOptions) (*exampleapiv1.ClusterTestType, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Get(ctx, name, opts)
}

// List implements ClusterTestTypeInterface.
func (w *wrappedClusterTestType) List(ctx context.Context, opts metav1.ListOptions) (*exampleapiv1.ClusterTestTypeList, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.List(ctx, opts)
}

// Watch implements ClusterTestTypeInterface.
func (w *wrappedClusterTestType) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Watch(ctx, opts)
}

// Patch implements ClusterTestTypeInterface.
func (w *wrappedClusterTestType) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *exampleapiv1.ClusterTestType, err error) {
	ctx, err = w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Patch(ctx, name, pt, data, opts, subresources...)
}

// WrappedExampleV1 contains the wrapped logical cluster and interface.
func (w *WrappedExampleV1) TestTypes(namespace string) examplev1.TestTypeInterface {
	return &wrappedTestType{
		cluster:  w.cluster,
		delegate: w.delegate.TestTypes(namespace),
	}
}

type wrappedTestType struct {
	cluster  logicalcluster.Name
	delegate examplev1.TestTypeInterface
}

// checkCluster retrieves the logical cluster name from the given context and checks
// if it is the same as the one passed while creating a wrappedTestType. It errors when
// there is a mismatch.
func (w *wrappedTestType) checkCluster(ctx context.Context) (context.Context, error) {
	ctxCluster, ok := kcp.ClusterFromContext(ctx)
	if !ok {
		return kcp.WithCluster(ctx, w.cluster), nil
	} else if ctxCluster != w.cluster {
		return ctx, fmt.Errorf("cluster mismatch: context=%q, client=%q", ctxCluster, w.cluster)
	}
	return ctx, nil
}

// Create implements TestTypeInterface.
func (w *wrappedTestType) Create(ctx context.Context, testType *exampleapiv1.TestType, opts metav1.CreateOptions) (*exampleapiv1.TestType, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Create(ctx, testType, opts)
}

// Update implements TestTypeInterface.
func (w *wrappedTestType) Update(ctx context.Context, testType *exampleapiv1.TestType, opts metav1.UpdateOptions) (*exampleapiv1.TestType, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Update(ctx, testType, opts)
}

// Update implements TestTypeInterface.
func (w *wrappedTestType) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return err
	}
	return w.delegate.Delete(ctx, name, opts)
}

// DeleteCollection implements TestTypeInterface.
func (w *wrappedTestType) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listopts metav1.ListOptions) error {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return err
	}
	return w.delegate.DeleteCollection(ctx, opts, listopts)
}

// Get implements TestTypeInterface.
func (w *wrappedTestType) Get(ctx context.Context, name string, opts metav1.GetOptions) (*exampleapiv1.TestType, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Get(ctx, name, opts)
}

// List implements TestTypeInterface.
func (w *wrappedTestType) List(ctx context.Context, opts metav1.ListOptions) (*exampleapiv1.TestTypeList, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.List(ctx, opts)
}

// Watch implements TestTypeInterface.
func (w *wrappedTestType) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	ctx, err := w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Watch(ctx, opts)
}

// Patch implements TestTypeInterface.
func (w *wrappedTestType) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *exampleapiv1.TestType, err error) {
	ctx, err = w.checkCluster(ctx)
	if err != nil {
		return nil, err
	}
	return w.delegate.Patch(ctx, name, pt, data, opts, subresources...)
}
