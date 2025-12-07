.PHONY: test
test: functest unittest

.PHONY: functest
functest:
	pnpm test

.PHONY: unittest
unittest:
	go test ./...
