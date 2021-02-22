package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PrometheusOperator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec PrometheusOperatorSpec `json:"spec"`
}

type PrometheusOperatorSpec struct {
	// Namespace for a prometheus-operator deployment.
	Namespace string `json:"namespace,omitempty"`
	// Image to use for a prometheus-operator deployment.
	Image string `json:"image,omitempty"`
	// Version string used for labeling objects
	Version string `json:"version,omitempty"`
	// Define resources requests and limits for each prometheus-operator pod.
	// +optional
	// +kubebuilder:default={requests:{cpu:"100m",memory:"100Mi"},limits:{cpu:"200m",memory:"200Mi"}}
	Resources v1.ResourceRequirements `json:"resources,omitempty"`
	// Image of config-reloader sidecar container
	ConfigmapReloaderImage string `json:"configmapReloaderImage,omitempty"`
}
