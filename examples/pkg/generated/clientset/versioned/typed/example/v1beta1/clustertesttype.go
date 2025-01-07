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

// Code generated by client-gen-v0.32. DO NOT EDIT.

package v1beta1

import (
	context "context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"

	examplev1beta1 "acme.corp/pkg/apis/example/v1beta1"
	scheme "acme.corp/pkg/generated/clientset/versioned/scheme"
)

// ClusterTestTypesGetter has a method to return a ClusterTestTypeInterface.
// A group's client should implement this interface.
type ClusterTestTypesGetter interface {
	ClusterTestTypes() ClusterTestTypeInterface
}

// ClusterTestTypeInterface has methods to work with ClusterTestType resources.
type ClusterTestTypeInterface interface {
	Create(ctx context.Context, clusterTestType *examplev1beta1.ClusterTestType, opts v1.CreateOptions) (*examplev1beta1.ClusterTestType, error)
	Update(ctx context.Context, clusterTestType *examplev1beta1.ClusterTestType, opts v1.UpdateOptions) (*examplev1beta1.ClusterTestType, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, clusterTestType *examplev1beta1.ClusterTestType, opts v1.UpdateOptions) (*examplev1beta1.ClusterTestType, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*examplev1beta1.ClusterTestType, error)
	List(ctx context.Context, opts v1.ListOptions) (*examplev1beta1.ClusterTestTypeList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *examplev1beta1.ClusterTestType, err error)
	ClusterTestTypeExpansion
}

// clusterTestTypes implements ClusterTestTypeInterface
type clusterTestTypes struct {
	*gentype.ClientWithList[*examplev1beta1.ClusterTestType, *examplev1beta1.ClusterTestTypeList]
}

// newClusterTestTypes returns a ClusterTestTypes
func newClusterTestTypes(c *ExampleV1beta1Client) *clusterTestTypes {
	return &clusterTestTypes{
		gentype.NewClientWithList[*examplev1beta1.ClusterTestType, *examplev1beta1.ClusterTestTypeList](
			"clustertesttypes",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *examplev1beta1.ClusterTestType { return &examplev1beta1.ClusterTestType{} },
			func() *examplev1beta1.ClusterTestTypeList { return &examplev1beta1.ClusterTestTypeList{} },
		),
	}
}
