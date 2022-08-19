default: testAll

TEST_COVERAGE_THRESHOLD := 100

testAll:
	go test -race github.com/aabishkaryal/DSA/problems -json -cover | tparse -smallscreen

test:
	go test -race github.com/aabishkaryal/DSA/problems -failfast -run Test$(FUNC)

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