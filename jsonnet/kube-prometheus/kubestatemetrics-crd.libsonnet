{"apiVersion":"apiextensions.k8s.io/v1","kind":"CustomResourceDefinition","metadata":{"annotations":{"controller-gen.kubebuilder.io/version":"v0.4.1"},"creationTimestamp":null,"name":"kubestatemetrics.monitoring.prometheus-operator.dev"},"spec":{"group":"monitoring.prometheus-operator.dev","names":{"kind":"KubeStateMetrics","listKind":"KubeStateMetricsList","plural":"kubestatemetrics","singular":"kubestatemetrics"},"scope":"Namespaced","versions":[{"name":"v1alpha1","schema":{"openAPIV3Schema":{"properties":{"apiVersion":{"description":"APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources","type":"string"},"kind":{"description":"Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds","type":"string"},"metadata":{"type":"object"},"spec":{"properties":{"image":{"description":"Image to use for a kube-state-metrics deployment.","type":"string"},"namespace":{"description":"Namespace for a kube-state-metrics deployment.","type":"string"},"resources":{"description":"Define resources requests and limits for each kube-state-metrics pod.","properties":{"limits":{"additionalProperties":{"anyOf":[{"type":"integer"},{"type":"string"}],"pattern":"^(\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))))?$","x-kubernetes-int-or-string":true},"description":"Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/","type":"object"},"requests":{"additionalProperties":{"anyOf":[{"type":"integer"},{"type":"string"}],"pattern":"^(\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))))?$","x-kubernetes-int-or-string":true},"description":"Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/","type":"object"}},"type":"object"},"scrapeInterval":{"default":"30s","description":"Scrape Interval for kube-state-metrics ServiceMonitor","type":"string"},"version":{"description":"Version string used for labeling objects","type":"string"}},"type":"object"}},"required":["spec"],"type":"object"}},"served":true,"storage":true}]},"status":{"acceptedNames":{"kind":"","plural":""},"conditions":[],"storedVersions":[]}}