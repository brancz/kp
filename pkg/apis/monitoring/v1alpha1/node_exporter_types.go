package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NodeExporter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec NodeExporterSpec `json:"spec"`
}

type NodeExporterSpec struct {
	// Define resources requests and limits for each node-exporter pod.
	Resources v1.ResourceRequirements `json:"resources,omitempty"`
}
