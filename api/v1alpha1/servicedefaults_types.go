/*


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
	consulapi "github.com/hashicorp/consul/api"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ServiceDefaultsSpec defines the desired state of ServiceDefaults
type ServiceDefaultsSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Protocol    string            `json:"protocol,omitempty"`
	MeshGateway MeshGatewayConfig `json:"meshGateway,omitempty"`
	Expose      ExposeConfig      `json:"expose,omitempty"`
	ExternalSNI string            `json:"externalSNI,omitempty"`
}

// ServiceDefaultsStatus defines the observed state of ServiceDefaults
type ServiceDefaultsStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ServiceDefaults is the Schema for the servicedefaults API
type ServiceDefaults struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceDefaultsSpec   `json:"spec,omitempty"`
	Status ServiceDefaultsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceDefaultsList contains a list of ServiceDefaults
type ServiceDefaultsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceDefaults `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ServiceDefaults{}, &ServiceDefaultsList{})
}

func (s *ServiceDefaults) ToConsul() *consulapi.ServiceConfigEntry {
	return &consulapi.ServiceConfigEntry{
		Kind:      s.Kind,
		Name:      s.Name,
		Namespace: s.Namespace, //this is subject to change
		Protocol:  s.Spec.Protocol,
		MeshGateway: consulapi.MeshGatewayConfig{
			Mode: consulapi.MeshGatewayModeDefault, //this will change. forcing it to default for now.
		},
		Expose: consulapi.ExposeConfig{
			Checks: s.Spec.Expose.Checks,
			Paths:  []consulapi.ExposePath{}, //will create a helper on our expose paths to translate to consul expose paths
		},
		ExternalSNI: s.Spec.ExternalSNI,
	}
}

// this will check if the consul struct shares the same spec as the spec of the resource
func (s *ServiceDefaults) MatchesConsul(entry consulapi.ConfigEntry) bool {
	return true
}
