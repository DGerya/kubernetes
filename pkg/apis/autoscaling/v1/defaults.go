/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

package v1

import (
	"k8s.io/kubernetes/pkg/runtime"
)

func addDefaultingFuncs(scheme *runtime.Scheme) {
	scheme.AddDefaultingFuncs(
		func(obj *HorizontalPodAutoscaler) {
			if obj.Spec.MinReplicas == nil {
				minReplicas := int32(1)
				obj.Spec.MinReplicas = &minReplicas
			}
			if obj.Spec.CPUUtilization == nil {
				obj.Spec.CPUUtilization = &CPUTargetUtilization{TargetPercentage: 80}
			}
		},
	)
}
