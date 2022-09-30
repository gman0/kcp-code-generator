/*
Copyright 2022 The KCP Authors.

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

package informergen

import (
	"io"
	"text/template"
)

type FactoryInterface struct {
	// ClientsetPackagePath is the package under which the cluster-aware client-set will be exposed.
	// TODO(skuznets) we should be able to figure this out from the output dir, ideally
	ClientsetPackagePath string
}

func (f *FactoryInterface) WriteContent(w io.Writer) error {
	templ, err := template.New("factoryInterface").Funcs(templateFuncs).Parse(externalSharedInformerFactoryInterface)
	if err != nil {
		return err
	}

	m := map[string]interface{}{
		"clientsetPackagePath": f.ClientsetPackagePath,
	}
	return templ.Execute(w, m)
}

var externalSharedInformerFactoryInterface = `
//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by kcp code-generator. DO NOT EDIT.

package internalinterfaces

import (
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	cache "k8s.io/client-go/tools/cache"

	client "{{.clientsetPackagePath}}"
)

// NewInformerFunc takes client.ClusterInterface and time.Duration to return a SharedIndexInformer.
type NewInformerFunc func(client.ClusterInterface, time.Duration) cache.SharedIndexInformer

// SharedInformerFactory a small interface to allow for adding an informer without an import cycle
type SharedInformerFactory interface {
	Start(stopCh <-chan struct{})
	InformerFor(obj runtime.Object, newFunc NewInformerFunc) cache.SharedIndexInformer
}

// TweakListOptionsFunc is a function that transforms a v1.ListOptions.
type TweakListOptionsFunc func(*v1.ListOptions)
`
