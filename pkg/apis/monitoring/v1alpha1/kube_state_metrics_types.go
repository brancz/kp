package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type KubeStateMetrics struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec KubeStateMetricsSpec `json:"spec"`
}

type KubeStateMetricsSpec struct {
	// Define resources requests and limits for each kube-state-metrics pod.
	Resources v1.ResourceRequirements `json:"resources,omitempty"`
	// Namespace for a kube-state-metrics deployment.
	Namespace string `json:"namespace,omitempty"`
	// Image to use for a kube-state-metrics deployment.
	Image string `json:"image,omitempty"`
	// Version string used for labeling objects
	Version string `json:"version,omitempty"`
	// Scrape Interval for kube-state-metrics ServiceMonitor
	// +kubebuilder:default="30s"
	ScrapeInterval string `json:"scrapeInterval,omitempty"`
}
