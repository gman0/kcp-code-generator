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

package v1beta1

import (
	"github.com/kcp-dev/logicalcluster/v2"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	"k8s.io/client-go/rest"

	examplev1beta1 "acme.corp/pkg/generated/clientset/versioned/typed/example/v1beta1"
	kcpexamplev1beta1 "acme.corp/pkg/kcpexisting/clients/clientset/versioned/typed/example/v1beta1"
)

var _ kcpexamplev1beta1.ExampleV1beta1ClusterInterface = (*ExampleV1beta1ClusterClient)(nil)

type ExampleV1beta1ClusterClient struct {
	*kcptesting.Fake
}

func (c *ExampleV1beta1ClusterClient) Cluster(cluster logicalcluster.Name) examplev1beta1.ExampleV1beta1Interface {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &ExampleV1beta1Client{Fake: c.Fake, Cluster: cluster}
}

func (c *ExampleV1beta1ClusterClient) TestTypes() kcpexamplev1beta1.TestTypeClusterInterface {
	return &testTypesClusterClient{Fake: c.Fake}
}

func (c *ExampleV1beta1ClusterClient) ClusterTestTypes() kcpexamplev1beta1.ClusterTestTypeClusterInterface {
	return &clusterTestTypesClusterClient{Fake: c.Fake}
}

var _ examplev1beta1.ExampleV1beta1Interface = (*ExampleV1beta1Client)(nil)

type ExampleV1beta1Client struct {
	*kcptesting.Fake
	Cluster logicalcluster.Name
}

func (c *ExampleV1beta1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *ExampleV1beta1Client) TestTypes(namespace string) examplev1beta1.TestTypeInterface {
	return &testTypesClient{Fake: c.Fake, Cluster: c.Cluster, Namespace: namespace}
}

func (c *ExampleV1beta1Client) ClusterTestTypes() examplev1beta1.ClusterTestTypeInterface {
	return &clusterTestTypesClient{Fake: c.Fake, Cluster: c.Cluster}
}
