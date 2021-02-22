package v1alpha1

import (
	promadapter "github.com/kubernetes-sigs/prometheus-adapter/pkg/config"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PrometheusAdapter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec PrometheusAdapterSpec `json:"spec"`
}

type PrometheusAdapterSpec struct {
	// Namespace for a prometheus-adapter deployment.
	Namespace string `json:"namespace,omitempty"`
	// Image to use for a prometheus-adapter deployment.
	Image string `json:"image,omitempty"`
	// Version string used for labeling objects
	Version string `json:"version,omitempty"`
	// Define resources requests and limits for each prometheus-adapter pod.
	// +optional
	// +kubebuilder:default={requests:{cpu:"102m",memory:"180Mi"},limits:{cpu:"250m",memory:"180Mi"}}
	Resources v1.ResourceRequirements `json:"resources,omitempty"`
	// Secure listening port
	// +optional
	// +kubebuilder:default=9100
	Port int `json:"port,omitempty"`
	// Configuration of prometheus-adapter in the same form as in
	// +optional
	Config promadapter.MetricsDiscoveryConfig `json:"config,omitempty"` // TODO: requires https://github.com/kubernetes-sigs/prometheus-adapter/pull/372
}
