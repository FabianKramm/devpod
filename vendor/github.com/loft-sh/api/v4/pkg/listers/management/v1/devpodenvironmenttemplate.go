// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	managementv1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// DevPodEnvironmentTemplateLister helps list DevPodEnvironmentTemplates.
// All objects returned here must be treated as read-only.
type DevPodEnvironmentTemplateLister interface {
	// List lists all DevPodEnvironmentTemplates in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*managementv1.DevPodEnvironmentTemplate, err error)
	// Get retrieves the DevPodEnvironmentTemplate from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*managementv1.DevPodEnvironmentTemplate, error)
	DevPodEnvironmentTemplateListerExpansion
}

// devPodEnvironmentTemplateLister implements the DevPodEnvironmentTemplateLister interface.
type devPodEnvironmentTemplateLister struct {
	listers.ResourceIndexer[*managementv1.DevPodEnvironmentTemplate]
}

// NewDevPodEnvironmentTemplateLister returns a new DevPodEnvironmentTemplateLister.
func NewDevPodEnvironmentTemplateLister(indexer cache.Indexer) DevPodEnvironmentTemplateLister {
	return &devPodEnvironmentTemplateLister{listers.New[*managementv1.DevPodEnvironmentTemplate](indexer, managementv1.Resource("devpodenvironmenttemplate"))}
}
