.PHONY: client
client:
	@echo "==> Building GitHub v3 REST API client..."
	@docker run \
		--rm \
		--volume "$(PWD):/local" \
		 openapitools/openapi-generator-cli:v6.0.1 batch --clean /local/config.yaml
