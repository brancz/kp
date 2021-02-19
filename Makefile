SHELL=/bin/bash -o pipefail
TOOLS_BIN_DIR ?= $(shell pwd)/tmp/bin
export PATH := $(TOOLS_BIN_DIR):$(PATH)

TYPES_V1ALPHA1_TARGET := pkg/apis/monitoring/v1alpha1/node_exporter_types.go

CONTROLLER_GEN_BINARY := $(TOOLS_BIN_DIR)/controller-gen
GOJSONTOYAML_BINARY := $(TOOLS_BIN_DIR)/gojsontoyaml

all: k8s-gen generate-crds

DEEPCOPY_TARGETS := pkg/apis/monitoring/v1alpha1/zz_generated.deepcopy.go
$(DEEPCOPY_TARGETS): $(CONTROLLER_GEN_BINARY)
	cd ./pkg/apis/monitoring/v1alpha1 && $(CONTROLLER_GEN_BINARY) object:headerFile=$(CURDIR)/.header \
		paths=.

.PHONY: k8s-gen
k8s-gen: \
	$(DEEPCOPY_TARGETS)

.PHONY: generate-crds
generate-crds: $(CONTROLLER_GEN_BINARY) $(GOJSONTOYAML_BINARY) $(TYPES_V1ALPHA1_TARGET)
	go run -v ./generate-crds.go --controller-gen=$(CONTROLLER_GEN_BINARY) --gojsontoyaml=$(GOJSONTOYAML_BINARY)

############
# Binaries #
############

$(TOOLS_BIN_DIR):
	mkdir -p $(TOOLS_BIN_DIR)

$(CONTROLLER_GEN_BINARY): $(TOOLS_BIN_DIR)
	# Requires Go >=1.16
	GOBIN=$(TOOLS_BIN_DIR) go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.4.1

$(GOJSONTOYAML_BINARY): $(TOOLS_BIN_DIR)
	# Requires Go >=1.16
	GOBIN=$(TOOLS_BIN_DIR) go install github.com/brancz/gojsontoyaml@202f76bf8c1f8fb74941a845b349941064603185
