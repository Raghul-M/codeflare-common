# CODEFLARE_SDK_VERSION defines the default version of the CodeFlare SDK
CODEFLARE_SDK_VERSION ?= 0.9.0

# RAY_VERSION defines the default version of Ray (used for testing)
RAY_VERSION ?= 2.5.0

# RAY_IMAGE defines the default container image for Ray (used for testing)
RAY_IMAGE ?= rayproject/ray:$(RAY_VERSION)

##@ Development

DEFAULTS_TEST_FILE := support/defaults.go

.PHONY: defaults
defaults:
	$(info Regenerating $(DEFAULTS_TEST_FILE))
	@echo "package support" > $(DEFAULTS_TEST_FILE)
	@echo "" >> $(DEFAULTS_TEST_FILE)
	@echo "// ***********************" >> $(DEFAULTS_TEST_FILE)
	@echo "//  DO NOT EDIT THIS FILE"  >> $(DEFAULTS_TEST_FILE)
	@echo "// ***********************" >> $(DEFAULTS_TEST_FILE)
	@echo "" >> $(DEFAULTS_TEST_FILE)
	@echo "const (" >> $(DEFAULTS_TEST_FILE)
	@echo "  CodeFlareSDKVersion = \"$(CODEFLARE_SDK_VERSION)\"" >> $(DEFAULTS_TEST_FILE)
	@echo "  RayVersion = \"$(RAY_VERSION)\"" >> $(DEFAULTS_TEST_FILE)
	@echo "  RayImage = \"$(RAY_IMAGE)\"" >> $(DEFAULTS_TEST_FILE)
	@echo "" >> $(DEFAULTS_TEST_FILE)
	@echo ")" >> $(DEFAULTS_TEST_FILE)
	@echo "" >> $(DEFAULTS_TEST_FILE)

	gofmt -w $(DEFAULTS_TEST_FILE)

## Tool Binaries
CONTROLLER_GEN ?= $(LOCALBIN)/controller-gen

## Tool Versions
CONTROLLER_TOOLS_VERSION ?= v0.9.2


## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

OPENSHIFT-GOIMPORTS ?= $(LOCALBIN)/openshift-goimports

.PHONY: openshift-goimports
openshift-goimports: $(OPENSHIFT-GOIMPORTS) ## Download openshift-goimports locally if necessary.
$(OPENSHIFT-GOIMPORTS): $(LOCALBIN)
	test -s $(LOCALBIN)/openshift-goimports || GOBIN=$(LOCALBIN) go install github.com/openshift-eng/openshift-goimports@latest

.PHONY: imports
imports: openshift-goimports ## Organize imports in go files using openshift-goimports. Example: make imports
	$(OPENSHIFT-GOIMPORTS)

.PHONY: verify-imports
verify-imports: openshift-goimports ## Run import verifications.
	./hack/verify-imports.sh $(OPENSHIFT-GOIMPORTS)

.PHONY: controller-gen
controller-gen: $(CONTROLLER_GEN) ## Download controller-gen locally if necessary.
$(CONTROLLER_GEN): $(LOCALBIN)
	test -s $(LOCALBIN)/controller-gen || GOBIN=$(LOCALBIN) go install sigs.k8s.io/controller-tools/cmd/controller-gen@$(CONTROLLER_TOOLS_VERSION)

.PHONY: manifests
manifests: controller-gen ## Generate RBAC objects.
	$(CONTROLLER_GEN) rbac:roleName=manager-role webhook paths="./..."