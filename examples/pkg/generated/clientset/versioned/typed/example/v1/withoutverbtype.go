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

// Code generated by client-gen-v0.24.0. DO NOT EDIT.

package v1

import (
	rest "k8s.io/client-go/rest"
)

// WithoutVerbTypesGetter has a method to return a WithoutVerbTypeInterface.
// A group's client should implement this interface.
type WithoutVerbTypesGetter interface {
	WithoutVerbTypes(namespace string) WithoutVerbTypeInterface
}

// WithoutVerbTypeInterface has methods to work with WithoutVerbType resources.
type WithoutVerbTypeInterface interface {
	WithoutVerbTypeExpansion
}

// withoutVerbTypes implements WithoutVerbTypeInterface
type withoutVerbTypes struct {
	client rest.Interface
	ns     string
}

// newWithoutVerbTypes returns a WithoutVerbTypes
func newWithoutVerbTypes(c *ExampleV1Client, namespace string) *withoutVerbTypes {
	return &withoutVerbTypes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}