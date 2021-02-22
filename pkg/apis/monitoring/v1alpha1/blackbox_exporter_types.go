package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type BlackboxExporter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec BlackboxExporterSpec `json:"spec"`
}

type BlackboxExporterSpec struct {
	// Define resources requests and limits for each blackbox-exporter pod.
	Resources v1.ResourceRequirements `json:"resources,omitempty"`
	// Namespace for a blackbox-exporter deployment.
	Namespace string `json:"namespace,omitempty"`
	// Image to use for a blackbox-exporter deployment.
	Image string `json:"image,omitempty"`
	// Version string used for labeling objects
	Version string `json:"version,omitempty"`
	// Image of config-reloader sidecar container
	// +kubebuilder:default="jimmidyson/configmap-reload:v0.5.0"
	ConfigmapReloaderImage string `json:"configmapReloaderImage,omitempty"`
	// Secure listening port
	// +kubebuilder:default=19115
	Port int `json:"port,omitempty"`
	// Number of replicas of blackbox-exporter
	// +kubebuilder:default=1
	Replicas int `json:"replicas,omitempty"`
	// Configuration of blackbox-exporter in the same form as in
	Modules map[string]string `json:"modules,omitempty"` // TODO: Consider using blackbox-exporter types for this
}
