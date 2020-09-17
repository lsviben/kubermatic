// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	versioned "k8c.io/kubermatic/v2/pkg/crd/client/clientset/versioned"
	internalinterfaces "k8c.io/kubermatic/v2/pkg/crd/client/informers/externalversions/internalinterfaces"
	v1 "k8c.io/kubermatic/v2/pkg/crd/client/listers/kubermatic/v1"
	kubermaticv1 "k8c.io/kubermatic/v2/pkg/crd/kubermatic/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ConstraintTemplateInformer provides access to a shared informer and lister for
// ConstraintTemplates.
type ConstraintTemplateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ConstraintTemplateLister
}

type constraintTemplateInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewConstraintTemplateInformer constructs a new informer for ConstraintTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewConstraintTemplateInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredConstraintTemplateInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredConstraintTemplateInformer constructs a new informer for ConstraintTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredConstraintTemplateInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubermaticV1().ConstraintTemplates().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubermaticV1().ConstraintTemplates().Watch(context.TODO(), options)
			},
		},
		&kubermaticv1.ConstraintTemplate{},
		resyncPeriod,
		indexers,
	)
}

func (f *constraintTemplateInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredConstraintTemplateInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *constraintTemplateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubermaticv1.ConstraintTemplate{}, f.defaultInformer)
}

func (f *constraintTemplateInformer) Lister() v1.ConstraintTemplateLister {
	return v1.NewConstraintTemplateLister(f.Informer().GetIndexer())
}
