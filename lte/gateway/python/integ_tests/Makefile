# By default just setup the env.
.PHONY: default
default: $(PYTHON_BUILD)/setupinteg_env

include $(MAGMA_ROOT)/orc8r/gateway/python/python.mk
include defs.mk

# Pull in dependencies from the magma dev build to create virtualenv and
# generate the python protobufs and the common rpc wrappers.
# Also, copy over test config files and update PYTHONPATH to include s1aptester
# lib directory, as well as Swagger-generated API bindings.
$(PYTHON_BUILD)/setupinteg_env:
	make -C .. buildenv
	#@echo "Installing swagger-codegen requirements"
	#$(VIRT_ENV_PIP_INSTALL) -r $(SWAGGER_CODEGEN_OUTPUT)/requirements.txt
	@echo "Copying s1aptester config files"
	cp $(S1AP_TESTER_CFG)/* $(S1AP_TESTER_ROOT)
	@echo "Install MySQL for upstreaming"
	$(VIRT_ENV_PIP_INSTALL) mysqlclient==1.3.13
	@echo "Install flaky for test retries support"
	$(VIRT_ENV_PIP_INSTALL) flaky
	@echo export PYTHONPATH=$(PYTHONPATH):$(S1AP_TESTER_PYTHON_PATH) >> $(PYTHON_BUILD)/bin/activate
	touch $(PYTHON_BUILD)/setupinteg_env

RESULTS_DIR := /var/tmp/test_results
ifdef enable-flaky-retry
	FLAKY_CMD_ARGS := --force-flaky --no-flaky-report --max-runs=3 --min-passes=1
else
	FLAKY_CMD_ARGS :=
endif

define apply_final_status
	if [ ! -z `grep -s pass $(MAGMA_ROOT)/test_status.txt` ]; \
	then echo "Final status: Passed"; \
	else echo "Final status: Failed"; exit 1; \
	fi
endef

# Don't remove the ending comment from function execute_test. Somehow without
# any ending statement the failing testcases are missing status update
define execute_test
	echo "Running test: $(1)"
	timeout --foreground -k 930s 900s \
	sudo -E PATH=$(PATH) PYTHONPATH=$(PYTHONPATH):$(S1AP_TESTER_PYTHON_PATH) \
	$(BIN)/pytest \
	$(FLAKY_CMD_ARGS) \
	--capture=tee-sys \
	--junit-xml=$(RESULTS_DIR)/$(basename $(notdir $(1))).xml -x $(1) \
	|| (echo "fail" > $(MAGMA_ROOT)/test_status.txt && exit 1)
	# Testcase $(1) execution completed
endef

.PHONY: prepare_environment
prepare_environment: $(PYTHON_BUILD)/setupinteg_env $(BIN)/pytest
	. $(PYTHON_BUILD)/bin/activate
	echo "pass" > $(MAGMA_ROOT)/test_status.txt
	echo "Flaky test retries enabled? $(enable-flaky-retry)"

.PHONY: prepare_federation
prepare_federation:
	$(eval export FEDERATED_MODE = True)

.PHONY: selected_tests
selected_tests: prepare_environment
	-$(foreach test,$(TESTS),$(call execute_test,$(test));)

.PHONY: precommit
precommit: prepare_environment
	-$(foreach test,$(PRECOMMIT_TESTS),$(call execute_test,$(test));)

.PHONY: precommit_tests_noncontainer
precommit_tests_noncontainer: prepare_environment
	-$(foreach test,$(PRECOMMIT_TESTS_NOT_CONTAINER) $(PRECOMMIT_TESTS_IPV6),$(call execute_test,$(test));)

.PHONY: extended_tests
extended_tests: prepare_environment
	-$(foreach test,$(EXTENDED_TESTS),$(call execute_test,$(test));)

.PHONY: extended_tests_noncontainer
extended_tests_noncontainer: prepare_environment
	-$(foreach test,$(EXTENDED_TESTS_NOT_CONTAINER),$(call execute_test,$(test));)

.PHONY: extended_tests_long
extended_tests_long: prepare_environment
	-$(foreach test,$(EXTENDED_TESTS_LONG),$(call execute_test,$(test));)

.PHONY: integ_test
integ_test: precommit precommit_tests_noncontainer extended_tests extended_tests_noncontainer extended_tests_long

.PHONY: integ_test_containerized
integ_test_containerized: precommit extended_tests

.PHONY: federated_integ_test
federated_integ_test: prepare_environment prepare_federation
	-$(foreach test,$(FEDERATED_TESTS),$(call execute_test,$(test));)

.PHONY: nonsanity
nonsanity: prepare_environment
	-$(foreach test,$(NON_SANITY_TESTS),$(call execute_test,$(test));)

.PHONY: evaluate_result
evaluate_result:
	$(call apply_final_status)
