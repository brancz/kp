module github.com/brancz/kp

go 1.16

require (
	github.com/kubernetes-sigs/prometheus-adapter v0.8.3
	github.com/pkg/errors v0.9.1
	k8s.io/api v0.20.4
	k8s.io/apimachinery v0.20.4
)

replace github.com/kubernetes-sigs/prometheus-adapter => github.com/directxman12/k8s-prometheus-adapter v0.8.3
