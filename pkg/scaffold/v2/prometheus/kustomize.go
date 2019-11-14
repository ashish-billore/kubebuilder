/*
Copyright 2019 The Kubernetes Authors.

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

package prometheus

import (
	"path/filepath"
	"sigs.k8s.io/kubebuilder/pkg/scaffold/input"
)

// Kustomization scaffolds the kustomizaiton in the prometheus folder
type Kustomization struct {
	input.Input
}

// GetInput implements input.File
func (p *Kustomization) GetInput() (input.Input, error) {
	if p.Path == "" {
		p.Path = filepath.Join("config", "prometheus", "kustomization.yaml")
	}
	p.TemplateBody = kustomizationTemplate
	return p.Input, nil
}

var kustomizationTemplate = `resources:
- monitor.yaml
`
