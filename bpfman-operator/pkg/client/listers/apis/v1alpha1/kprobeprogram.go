/*
Copyright 2023 The bpfman Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/bpfman/bpfman/bpfman-operator/apis/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KprobeProgramLister helps list KprobePrograms.
// All objects returned here must be treated as read-only.
type KprobeProgramLister interface {
	// List lists all KprobePrograms in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.KprobeProgram, err error)
	// Get retrieves the KprobeProgram from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.KprobeProgram, error)
	KprobeProgramListerExpansion
}

// kprobeProgramLister implements the KprobeProgramLister interface.
type kprobeProgramLister struct {
	indexer cache.Indexer
}

// NewKprobeProgramLister returns a new KprobeProgramLister.
func NewKprobeProgramLister(indexer cache.Indexer) KprobeProgramLister {
	return &kprobeProgramLister{indexer: indexer}
}

// List lists all KprobePrograms in the indexer.
func (s *kprobeProgramLister) List(selector labels.Selector) (ret []*v1alpha1.KprobeProgram, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KprobeProgram))
	})
	return ret, err
}

// Get retrieves the KprobeProgram from the index for a given name.
func (s *kprobeProgramLister) Get(name string) (*v1alpha1.KprobeProgram, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("kprobeprogram"), name)
	}
	return obj.(*v1alpha1.KprobeProgram), nil
}
