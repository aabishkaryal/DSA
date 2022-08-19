default: test-parse

TEST_COVERAGE_THRESHOLD := 100
FUZZ_TIME := 10
NUM_PROCESSORS := $(shell nproc)

test-parse:
	go test -race github.com/aabishkaryal/DSA/problems -failfast -json -cover | tparse -smallscreen

test-ci:
	go test -race -v github.com/aabishkaryal/DSA/problems -failfast

test-func:
	@test -n "$(FUNC)" || (echo "FUNC env var required" && exit 1)
	go test -race github.com/aabishkaryal/DSA/problems -failfast -run Test$(FUNC) 
	go test -race github.com/aabishkaryal/DSA/problems -failfast -run Fuzz$(FUNC) -fuzz=Fuzz -fuzztime $(FUZZ_TIME)s -parallel $$((${NUM_PROCESSORS} / 2))

test-deploy: lint test-parse coverage

clean:
	go clean -testcache

coverage:
	scripts/test_coverage.sh $(TEST_COVERAGE_THRESHOLD)

lint:
	golangci-lint run
	go mod tidy -v && git --no-pager diff --quiet go.mod

install-linter:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

install-fumpt:
	go install mvdan.cc/gofumpt@latest

install-tools: install-linter install-fumpt