// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// PrometheusLister helps list Prometheuses.
// All objects returned here must be treated as read-only.
type PrometheusLister interface {
	// List lists all Prometheuses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Prometheus, err error)
	// Prometheuses returns an object that can list and get Prometheuses.
	Prometheuses(namespace string) PrometheusNamespaceLister
	PrometheusListerExpansion
}

// prometheusLister implements the PrometheusLister interface.
type prometheusLister struct {
	listers.ResourceIndexer[*v1.Prometheus]
}

// NewPrometheusLister returns a new PrometheusLister.
func NewPrometheusLister(indexer cache.Indexer) PrometheusLister {
	return &prometheusLister{listers.New[*v1.Prometheus](indexer, v1.Resource("prometheus"))}
}

// Prometheuses returns an object that can list and get Prometheuses.
func (s *prometheusLister) Prometheuses(namespace string) PrometheusNamespaceLister {
	return prometheusNamespaceLister{listers.NewNamespaced[*v1.Prometheus](s.ResourceIndexer, namespace)}
}

// PrometheusNamespaceLister helps list and get Prometheuses.
// All objects returned here must be treated as read-only.
type PrometheusNamespaceLister interface {
	// List lists all Prometheuses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Prometheus, err error)
	// Get retrieves the Prometheus from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Prometheus, error)
	PrometheusNamespaceListerExpansion
}

// prometheusNamespaceLister implements the PrometheusNamespaceLister
// interface.
type prometheusNamespaceLister struct {
	listers.ResourceIndexer[*v1.Prometheus]
}
