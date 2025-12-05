
.PHONY: test
test: functest unittest

.PHONY: functest
functest:
	bash tests/functest.sh

.PHONY: unittest
unittest:
	go test ./...
