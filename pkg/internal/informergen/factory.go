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

	"github.com/kcp-dev/code-generator/pkg/parser"
)

type Factory struct {
	// Groups are the groups in this informer factory.
	Groups []parser.Group

	// PackagePath is the package under which these informers will be exposed.
	// e.g. "github.com/kcp-dev/client-go/clients/informers"
	// TODO(skuznets) we should be able to figure this out from the output dir, ideally
	PackagePath string

	// ClientsetPackagePath is the package under which the cluster-aware client-set will be exposed.
	// e.g. "github.com/kcp-dev/client-go/clients/clientset/versioned"
	// TODO(skuznets) we should be able to figure this out from the output dir, ideally
	ClientsetPackagePath string

	// SingleClusterInformerPackagePath is the package under which the cluster-unaware listers are exposed.
	// e.g. "k8s.io/client-go/informers"
	SingleClusterInformerPackagePath string
}

func (f *Factory) WriteContent(w io.Writer) error {
	templ, err := template.New("factory").Funcs(templateFuncs).Parse(sharedInformerFactoryStruct)
	if err != nil {
		return err
	}

	m := map[string]interface{}{
		"groups":                           f.Groups,
		"packagePath":                      f.PackagePath,
		"clientsetPackagePath":             f.ClientsetPackagePath,
		"singleClusterInformerPackagePath": f.SingleClusterInformerPackagePath,
		"useUpstreamInterfaces":            f.SingleClusterInformerPackagePath != "",
	}
	return templ.Execute(w, m)
}

var sharedInformerFactoryStruct = `
//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by kcp code-generator. DO NOT EDIT.

package informers

import (
	"reflect"
	"sync"
	"time"

	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v2"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/cache"

	clientset "{{.clientsetPackagePath}}"
	{{if .useUpstreamInterfaces -}}
	upstreaminformers "{{.singleClusterInformerPackagePath}}"
	{{end -}}

{{range .groups}}	{{.Group.PackageName}}informers "{{$.packagePath}}/{{.Group.PackageName}}"
{{end -}}

	"{{.packagePath}}/internalinterfaces"
)

// SharedInformerOption defines the functional option type for SharedInformerFactory.
type SharedInformerOption func(*sharedInformerFactory) *sharedInformerFactory

type sharedInformerFactory struct {
	client clientset.ClusterInterface
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	lock sync.Mutex
	defaultResync time.Duration
	customResync map[reflect.Type]time.Duration

	informers map[reflect.Type]kcpcache.ScopeableSharedIndexInformer
	// startedInformers is used for tracking which informers have been started.
	// This allows Start() to be called multiple times safely.
	startedInformers map[reflect.Type]bool
}

// WithCustomResyncConfig sets a custom resync period for the specified informer types.
func WithCustomResyncConfig(resyncConfig map[metav1.Object]time.Duration) SharedInformerOption {
	return func(factory *sharedInformerFactory) *sharedInformerFactory {
		for k, v := range resyncConfig {
			factory.customResync[reflect.TypeOf(k)] = v
		}
		return factory
	}
}

// WithTweakListOptions sets a custom filter on all listers of the configured SharedInformerFactory.
func WithTweakListOptions(tweakListOptions internalinterfaces.TweakListOptionsFunc) SharedInformerOption {
	return func(factory *sharedInformerFactory) *sharedInformerFactory {
		factory.tweakListOptions = tweakListOptions
		return factory
	}
}

// NewSharedInformerFactory constructs a new instance of SharedInformerFactory for all namespaces.
func NewSharedInformerFactory(client clientset.ClusterInterface, defaultResync time.Duration) SharedInformerFactory {
	return NewSharedInformerFactoryWithOptions(client, defaultResync)
}

// NewSharedInformerFactoryWithOptions constructs a new instance of a SharedInformerFactory with additional options.
func NewSharedInformerFactoryWithOptions(client clientset.ClusterInterface, defaultResync time.Duration, options ...SharedInformerOption) SharedInformerFactory {
	factory := &sharedInformerFactory{
		client:           client,
		defaultResync:    defaultResync,
		informers:        make(map[reflect.Type]kcpcache.ScopeableSharedIndexInformer),
		startedInformers: make(map[reflect.Type]bool),
		customResync:     make(map[reflect.Type]time.Duration),
	}

	// Apply all options
	for _, opt := range options {
		factory = opt(factory)
	}

	return factory
}

// Start initializes all requested informers.
func (f *sharedInformerFactory) Start(stopCh <-chan struct{}) {
  f.lock.Lock()
  defer f.lock.Unlock()

  for informerType, informer := range f.informers {
    if !f.startedInformers[informerType] {
      go informer.Run(stopCh)
      f.startedInformers[informerType] = true
    }
  }
}

// WaitForCacheSync waits for all started informers' cache were synced.
func (f *sharedInformerFactory) WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool {
	informers := func()map[reflect.Type]kcpcache.ScopeableSharedIndexInformer{
               f.lock.Lock()
               defer f.lock.Unlock()

               informers := map[reflect.Type]kcpcache.ScopeableSharedIndexInformer{}
               for informerType, informer := range f.informers {
                       if f.startedInformers[informerType] {
                               informers[informerType] = informer
                       }
               }
               return informers
       }()

       res := map[reflect.Type]bool{}
       for informType, informer := range informers {
               res[informType] = cache.WaitForCacheSync(stopCh, informer.HasSynced)
       }
       return res
}

// InternalInformerFor returns the SharedIndexInformer for obj using an internal
// client.
func (f *sharedInformerFactory) InformerFor(obj runtime.Object, newFunc internalinterfaces.NewInformerFunc) kcpcache.ScopeableSharedIndexInformer {
  f.lock.Lock()
  defer f.lock.Unlock()

  informerType := reflect.TypeOf(obj)
  informer, exists := f.informers[informerType]
  if exists {
    return informer
  }

  resyncPeriod, exists := f.customResync[informerType]
  if !exists {
    resyncPeriod = f.defaultResync
  }

  informer = newFunc(f.client, resyncPeriod)
  f.informers[informerType] = informer

  return informer
}

type ScopedDynamicSharedInformerFactory interface {
	ForResource(resource schema.GroupVersionResource) ({{if .useUpstreamInterfaces}}upstreaminformers.{{end}}GenericInformer, error)
	Start(stopCh <-chan struct{})
}

// SharedInformerFactory provides shared informers for resources in all known
// API group versions.
type SharedInformerFactory interface {
	internalinterfaces.SharedInformerFactory
	Cluster(logicalcluster.Name) ScopedDynamicSharedInformerFactory
	ForResource(resource schema.GroupVersionResource) (GenericClusterInformer, error)
	WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool

{{range .groups}}	{{.GoName}}() {{.Group.PackageName}}informers.ClusterInterface
{{end -}}
}

{{range .groups}}
func (f *sharedInformerFactory) {{.GoName}}() {{.Group.PackageName}}informers.ClusterInterface {
  return {{.Group.PackageName}}informers.New(f, f.tweakListOptions)
}
{{end}}

func (f *sharedInformerFactory) Cluster(cluster logicalcluster.Name) ScopedDynamicSharedInformerFactory {
	return &scopedDynamicSharedInformerFactory{
		sharedInformerFactory: f,
		cluster: cluster,
	}
}

type scopedDynamicSharedInformerFactory struct {
	*sharedInformerFactory
	cluster logicalcluster.Name
}

func (f *scopedDynamicSharedInformerFactory) ForResource(resource schema.GroupVersionResource) ({{if .useUpstreamInterfaces}}upstreaminformers.{{end}}GenericInformer, error) {
	clusterInformer, err := f.sharedInformerFactory.ForResource(resource)
	if err != nil {
		return nil, err
	}
	return clusterInformer.Cluster(f.cluster), nil 
}

func (f *scopedDynamicSharedInformerFactory) Start(stopCh <-chan struct{}) {
	f.sharedInformerFactory.Start(stopCh)
}
`
