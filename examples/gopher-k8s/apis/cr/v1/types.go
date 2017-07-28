/*
Copyright 2017 The Kubernetes Authors.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Thank you for Asta Xie invite me to Gopher meetup!
// In return, I will invite Asta Xie into Kubernetes, literally.
const AstaXieResourcePlural = "astaxies"

// +k8s:deepcopy-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type AstaXie struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	// Describes Asta Xie
	Spec AstaXieSpec `json:"spec"`
	// Status of Asta Xie
	Status AstaXieStatus `json:"status,omitempty"`
}

type AstaXieSpec struct {
	Language string `json:"language"`
	Handsome bool   `json:"handsome"`
}

type AstaXieStatus struct {
	State   AstaXieState `json:"state,omitempty"`
	Message string       `json:"message,omitempty"`
}

type AstaXieState string

const (
	AstaXieStateInvited  AstaXieState = "Invited"
	AstaXieStateAccepted AstaXieState = "Accepted"
)

// +k8s:deepcopy-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type AstaXieList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []AstaXie `json:"items"`
}
