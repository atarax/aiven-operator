// Copyright (c) 2022 Aiven, Helsinki, Finland. https://aiven.io/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	clickhouseuserconfig "github.com/aiven/aiven-operator/api/v1alpha1/userconfig/service/clickhouse"
)

// ClickhouseSpec defines the desired state of Clickhouse
type ClickhouseSpec struct {
	ServiceCommonSpec `json:",inline"`

	// +kubebuilder:validation:Format="^[1-9][0-9]*(GiB|G)*"
	// The disk space of the service, possible values depend on the service type, the cloud provider and the project. Reducing will result in the service re-balancing.
	DiskSpace string `json:"disk_space,omitempty"`

	// Authentication reference to Aiven token in a secret
	AuthSecretRef *AuthSecretReference `json:"authSecretRef,omitempty"`

	// Information regarding secret creation.
	// Exposed keys: `CLICKHOUSE_HOST`, `CLICKHOUSE_PORT`, `CLICKHOUSE_USER`, `CLICKHOUSE_PASSWORD`
	ConnInfoSecretTarget ConnInfoSecretTarget `json:"connInfoSecretTarget,omitempty"`

	// OpenSearch specific user configuration options
	UserConfig *clickhouseuserconfig.ClickhouseUserConfig `json:"userConfig,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Clickhouse is the Schema for the clickhouses API
type Clickhouse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClickhouseSpec `json:"spec,omitempty"`
	Status ServiceStatus  `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ClickhouseList contains a list of Clickhouse
type ClickhouseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Clickhouse `json:"items"`
}

func (in *Clickhouse) AuthSecretRef() *AuthSecretReference {
	return in.Spec.AuthSecretRef
}

func (in *Clickhouse) GetRefs() []*ResourceReferenceObject {
	return in.Spec.GetRefs(in.GetNamespace())
}

func (in *Clickhouse) GetConnInfoSecretTarget() ConnInfoSecretTarget {
	return in.Spec.ConnInfoSecretTarget
}

func init() {
	SchemeBuilder.Register(&Clickhouse{}, &ClickhouseList{})
}
