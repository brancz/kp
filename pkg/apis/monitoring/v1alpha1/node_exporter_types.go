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
	// The version of node-exporter the specified image represent.
	Version string `json:"version"`
	// The image of node-exporter to use.
	Image string `json:"image"`
	// ListenAddress is the host node-exporter processes will bind to.
	// +optional
	// +kubebuilder:default="127.0.0.1"
	ListenAddress string `json:"listenAddress"`
	// Port is the port node-exporter processes will bind to.
	// +optional
	// +kubebuilder:default=9100
	Port int `json:"port"`
	// Resources defines the resources requests and limits for each
	// node-exporter pod.
	// +optional
	// +kubebuilder:default={requests:{cpu:"102m",memory:"180Mi"},limits:{cpu:"250m",memory:"180Mi"}}
	Resources v1.ResourceRequirements `json:"resources,omitempty"`
}
