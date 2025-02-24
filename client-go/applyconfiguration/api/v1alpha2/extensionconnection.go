/*
Copyright 2024 The Kubernetes Authors.

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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha2

import (
	apiv1alpha2 "sigs.k8s.io/gateway-api-inference-extension/api/v1alpha2"
)

// ExtensionConnectionApplyConfiguration represents a declarative configuration of the ExtensionConnection type for use
// with apply.
type ExtensionConnectionApplyConfiguration struct {
	FailureMode *apiv1alpha2.ExtensionFailureMode `json:"failureMode,omitempty"`
}

// ExtensionConnectionApplyConfiguration constructs a declarative configuration of the ExtensionConnection type for use with
// apply.
func ExtensionConnection() *ExtensionConnectionApplyConfiguration {
	return &ExtensionConnectionApplyConfiguration{}
}

// WithFailureMode sets the FailureMode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FailureMode field is set to the value of the last call.
func (b *ExtensionConnectionApplyConfiguration) WithFailureMode(value apiv1alpha2.ExtensionFailureMode) *ExtensionConnectionApplyConfiguration {
	b.FailureMode = &value
	return b
}
