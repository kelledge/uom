PKG          ?= ./...
REPORTS_DIR  ?= reports
COVERPROFILE ?= $(REPORTS_DIR)/unit.coverage.out

JUNIT_UNIT_TEST_FILE ?= $(REPORTS_DIR)/unit.xml
JUNIT_LINT_FILE      ?= $(REPORTS_DIR)/lint.xml

.PHONY: test-unit
test-unit:
	mkdir -p $(REPORTS_DIR)
	gotestsum --format=testname --junitfile=$(JUNIT_UNIT_TEST_FILE) -- \
		-race -coverprofile=$(COVERPROFILE) -covermode=atomic $(PKG)

.PHONY: test-lint
test-lint:
	mkdir -p $(REPORTS_DIR)
	golangci-lint run

.PHONY: test-fmt
test-fmt:
	@out=$$(gofmt -l .); \
	if [ -n "$$out" ]; then \
		echo "Files need gofmt:"; echo "$$out"; \
		exit 1; \
	fi

.PHONY: test
test: test-fmt test-lint test-unit
	@echo "All tests passed."