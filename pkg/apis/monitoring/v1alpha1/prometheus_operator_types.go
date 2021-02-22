package v1alpha1

import (
	prom "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
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

type PrometheusProvisioning struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec PrometheusProvisioningSpec `json:"spec"`
}

// PrometheusProvisioningSpec extends PrometheusSpec from prometheus-operator by adding fields necessary for deployment
type PrometheusProvisioningSpec struct {
	prom.PrometheusSpec `json:",inline"`
	// Namespace to deploy prometheus statefulset in.
	Namespace string `json:"namespace,omitempty"`
	// Namespaces to allowed to be watched by prometheus
	// +optional
	// kubebuilder:default=["default","kube-system"] // TODO: how can we include a field value that needs to be defined by user? // FIXME(paulfantom): controller-gen doesn't like defaulting arrays
	Namespaces []string `json:"namespaces,omitempty"`
}

type AlertmanagerProvisioning struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec AlertmanagerProvisioningSpec `json:"spec"`
}

type AlertmanagerProvisioningSpec struct {
	prom.AlertmanagerSpec `json:",inline"`
	// Namespace to deploy alertmanager statefulset in.
	Namespace string `json:"namespace,omitempty"`
	// Embedded alertmanager configuration file. Spec as in https://prometheus.io/docs/alerting/latest/configuration/#configuration-file
	// +optional
	Config map[string]string `json:"config,omitempty"`
}
