/*
Copyright 2025.

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
	"errors"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var llamastackdistributionlog = logf.Log.WithName("llamastackdistribution-resource")

func (r *LlamaStackDistribution) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//nolint:lll // Webhook marker must be on a single line for controller-gen to parse correctly
//+kubebuilder:webhook:path=/validate-llama-x-k8s-io-v1alpha1-llamastackdistribution,mutating=false,failurePolicy=fail,sideEffects=None,groups=llama.x-k8s.io,resources=llamastackdistributions,verbs=create;update,versions=v1alpha1,name=llamastackdistribution.x-k8s.io,admissionReviewVersions=v1

var _ webhook.Validator = &LlamaStackDistribution{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type.
func (r *LlamaStackDistribution) ValidateCreate() error {
	llamastackdistributionlog.Info("validate create", "name", r.Name)
	return r.validateDistribution()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type.
func (r *LlamaStackDistribution) ValidateUpdate(old runtime.Object) error {
	llamastackdistributionlog.Info("validate update", "name", r.Name)
	return r.validateDistribution()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type.
func (r *LlamaStackDistribution) ValidateDelete() error {
	llamastackdistributionlog.Info("validate delete", "name", r.Name)
	return nil
}

func (r *LlamaStackDistribution) validateDistribution() error {
	// Validate that only one of name or image is set
	if r.Spec.Server.Distribution.Name != "" && r.Spec.Server.Distribution.Image != "" {
		return errors.New("only one of distribution.name or distribution.image can be set")
	}

	// Validate that at least one of name or image is set
	if r.Spec.Server.Distribution.Name == "" && r.Spec.Server.Distribution.Image == "" {
		return errors.New("either distribution.name or distribution.image must be set")
	}

	// If name is set, validate it exists in imageMap
	if r.Spec.Server.Distribution.Name != "" {
		if _, exists := ImageMap[r.Spec.Server.Distribution.Name]; !exists {
			return fmt.Errorf("invalid distribution name: %s", r.Spec.Server.Distribution.Name)
		}
	}

	return nil
}
