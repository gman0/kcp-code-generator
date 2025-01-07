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

package fake

import (
	context "context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gentype "k8s.io/client-go/gentype"
	testing "k8s.io/client-go/testing"

	examplev1 "acme.corp/pkg/apis/example/v1"
	v1 "acme.corp/pkg/apis/exampledashed/v1"
	exampledashedv1 "acme.corp/pkg/generated/clientset/versioned/typed/exampledashed/v1"
)

// fakeTestTypes implements TestTypeInterface
type fakeTestTypes struct {
	*gentype.FakeClientWithList[*v1.TestType, *v1.TestTypeList]
	Fake *FakeExampleDashedV1
}

func newFakeTestTypes(fake *FakeExampleDashedV1, namespace string) exampledashedv1.TestTypeInterface {
	return &fakeTestTypes{
		gentype.NewFakeClientWithList[*v1.TestType, *v1.TestTypeList](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("testtypes"),
			v1.SchemeGroupVersion.WithKind("TestType"),
			func() *v1.TestType { return &v1.TestType{} },
			func() *v1.TestTypeList { return &v1.TestTypeList{} },
			func(dst, src *v1.TestTypeList) { dst.ListMeta = src.ListMeta },
			func(list *v1.TestTypeList) []*v1.TestType { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.TestTypeList, items []*v1.TestType) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}

// CreateField takes the representation of a field and creates it.  Returns the server's representation of the field, and an error, if there is any.
func (c *fakeTestTypes) CreateField(ctx context.Context, testTypeName string, field *examplev1.Field, opts metav1.CreateOptions) (result *examplev1.Field, err error) {
	emptyResult := &examplev1.Field{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateSubresourceActionWithOptions(c.Resource(), testTypeName, "field", c.Namespace(), field, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*examplev1.Field), err
}

// UpdateField takes the representation of a field and updates it. Returns the server's representation of the field, and an error, if there is any.
func (c *fakeTestTypes) UpdateField(ctx context.Context, testTypeName string, field *examplev1.Field, opts metav1.UpdateOptions) (result *examplev1.Field, err error) {
	emptyResult := &examplev1.Field{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(c.Resource(), "field", c.Namespace(), field, opts), &examplev1.Field{})

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*examplev1.Field), err
}

// GetField takes name of the testType, and returns the corresponding field object, and an error if there is any.
func (c *fakeTestTypes) GetField(ctx context.Context, testTypeName string, options metav1.GetOptions) (result *examplev1.Field, err error) {
	emptyResult := &examplev1.Field{}
	obj, err := c.Fake.
		Invokes(testing.NewGetSubresourceActionWithOptions(c.Resource(), c.Namespace(), "field", testTypeName, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*examplev1.Field), err
}
