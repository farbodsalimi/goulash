.DEFAULT_GOAL=help

.PHONY: fix
fix: ## Modernize go code
	go fix ./...

.PHONY: test
test: ## Test all the test files recursively
	gotestsum -- -v -covermode=atomic ./...

.PHONY: --help
--help: ##
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-16s\033[0m %s\n", $$1, $$2}'

.PHONY: help
help: --help
