// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	operatorv1alpha1 "github.com/openshift/azure-disk-csi-driver-operator/pkg/apis/operator/v1alpha1"
	versioned "github.com/openshift/azure-disk-csi-driver-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/openshift/azure-disk-csi-driver-operator/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/openshift/azure-disk-csi-driver-operator/pkg/generated/listers/operator/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AzureDiskDriverInformer provides access to a shared informer and lister for
// AzureDiskDrivers.
type AzureDiskDriverInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.AzureDiskDriverLister
}

type azureDiskDriverInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewAzureDiskDriverInformer constructs a new informer for AzureDiskDriver type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAzureDiskDriverInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAzureDiskDriverInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredAzureDiskDriverInformer constructs a new informer for AzureDiskDriver type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAzureDiskDriverInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CsiV1alpha1().AzureDiskDrivers().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CsiV1alpha1().AzureDiskDrivers().Watch(context.TODO(), options)
			},
		},
		&operatorv1alpha1.AzureDiskDriver{},
		resyncPeriod,
		indexers,
	)
}

func (f *azureDiskDriverInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAzureDiskDriverInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *azureDiskDriverInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&operatorv1alpha1.AzureDiskDriver{}, f.defaultInformer)
}

func (f *azureDiskDriverInformer) Lister() v1alpha1.AzureDiskDriverLister {
	return v1alpha1.NewAzureDiskDriverLister(f.Informer().GetIndexer())
}
