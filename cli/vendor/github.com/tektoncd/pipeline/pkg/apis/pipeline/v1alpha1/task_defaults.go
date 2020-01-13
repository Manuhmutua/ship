/*
Copyright 2019 The Tekton Authors

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

package v1alpha1

import (
	"context"

	"knative.dev/pkg/apis"
)

var _ apis.Defaultable = (*Task)(nil)

func (t *Task) SetDefaults(ctx context.Context) {
	t.Spec.SetDefaults(ctx)
}

// SetDefaults set any defaults for the task spec
func (ts *TaskSpec) SetDefaults(ctx context.Context) {
	if ts.Inputs != nil {
		ts.Inputs.SetDefaults(ctx)
	}
}

func (inputs *Inputs) SetDefaults(ctx context.Context) {
	for i := range inputs.Params {
		inputs.Params[i].SetDefaults(ctx)
	}
}
