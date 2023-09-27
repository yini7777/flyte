// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/flyteorg/flyte/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FlyteWorkflowLister helps list FlyteWorkflows.
// All objects returned here must be treated as read-only.
type FlyteWorkflowLister interface {
	// List lists all FlyteWorkflows in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FlyteWorkflow, err error)
	// FlyteWorkflows returns an object that can list and get FlyteWorkflows.
	FlyteWorkflows(namespace string) FlyteWorkflowNamespaceLister
	FlyteWorkflowListerExpansion
}

// flyteWorkflowLister implements the FlyteWorkflowLister interface.
type flyteWorkflowLister struct {
	indexer cache.Indexer
}

// NewFlyteWorkflowLister returns a new FlyteWorkflowLister.
func NewFlyteWorkflowLister(indexer cache.Indexer) FlyteWorkflowLister {
	return &flyteWorkflowLister{indexer: indexer}
}

// List lists all FlyteWorkflows in the indexer.
func (s *flyteWorkflowLister) List(selector labels.Selector) (ret []*v1alpha1.FlyteWorkflow, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FlyteWorkflow))
	})
	return ret, err
}

// FlyteWorkflows returns an object that can list and get FlyteWorkflows.
func (s *flyteWorkflowLister) FlyteWorkflows(namespace string) FlyteWorkflowNamespaceLister {
	return flyteWorkflowNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FlyteWorkflowNamespaceLister helps list and get FlyteWorkflows.
// All objects returned here must be treated as read-only.
type FlyteWorkflowNamespaceLister interface {
	// List lists all FlyteWorkflows in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FlyteWorkflow, err error)
	// Get retrieves the FlyteWorkflow from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FlyteWorkflow, error)
	FlyteWorkflowNamespaceListerExpansion
}

// flyteWorkflowNamespaceLister implements the FlyteWorkflowNamespaceLister
// interface.
type flyteWorkflowNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FlyteWorkflows in the indexer for a given namespace.
func (s flyteWorkflowNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FlyteWorkflow, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FlyteWorkflow))
	})
	return ret, err
}

// Get retrieves the FlyteWorkflow from the indexer for a given namespace and name.
func (s flyteWorkflowNamespaceLister) Get(name string) (*v1alpha1.FlyteWorkflow, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("flyteworkflow"), name)
	}
	return obj.(*v1alpha1.FlyteWorkflow), nil
}