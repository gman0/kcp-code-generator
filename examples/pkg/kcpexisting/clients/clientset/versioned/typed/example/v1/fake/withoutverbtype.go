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
	"github.com/kcp-dev/logicalcluster/v3"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	"k8s.io/apimachinery/pkg/runtime/schema"

	examplev1client "acme.corp/pkg/generated/clientset/versioned/typed/example/v1"
	kcpexamplev1 "acme.corp/pkg/kcpexisting/clients/clientset/versioned/typed/example/v1"
)

var withoutVerbTypesResource = schema.GroupVersionResource{Group: "example", Version: "v1", Resource: "withoutverbtypes"}
var withoutVerbTypesKind = schema.GroupVersionKind{Group: "example", Version: "v1", Kind: "WithoutVerbType"}

type withoutVerbTypesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *withoutVerbTypesClusterClient) Cluster(clusterPath logicalcluster.Path) kcpexamplev1.WithoutVerbTypesNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &withoutVerbTypesNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

type withoutVerbTypesNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *withoutVerbTypesNamespacer) Namespace(namespace string) examplev1client.WithoutVerbTypeInterface {
	return &withoutVerbTypesClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type withoutVerbTypesClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}
