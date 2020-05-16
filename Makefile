.PHONY: default
default: help

.PHONY: help
help: ## help information about make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: opengl-version
opengl-version: ## Print out the OpenGl version installed
	glxinfo | grep "OpenGL version"

.PHONY: run
run: ## Run program
	go run ./main.go
