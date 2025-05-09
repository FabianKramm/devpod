// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	managementv1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// AppLister helps list Apps.
// All objects returned here must be treated as read-only.
type AppLister interface {
	// List lists all Apps in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*managementv1.App, err error)
	// Get retrieves the App from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*managementv1.App, error)
	AppListerExpansion
}

// appLister implements the AppLister interface.
type appLister struct {
	listers.ResourceIndexer[*managementv1.App]
}

// NewAppLister returns a new AppLister.
func NewAppLister(indexer cache.Indexer) AppLister {
	return &appLister{listers.New[*managementv1.App](indexer, managementv1.Resource("app"))}
}
