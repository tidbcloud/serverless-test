.PHONY: install-openapi-generator
install-openapi-generator:
	cd tools/openapi-generator && npm install

.PHONY: generate-coreportalapi-client
generate-coreportalapi-client: ## Generate corePortalAPI client
	@echo "==> Generating coreportalapi client"
	rm -rf pkg/coreportalapi
	cd tools/openapi-generator && npx openapi-generator-cli generate --inline-schema-options RESOLVE_INLINE_ENUMS=true --additional-properties=withGoMod=false,enumClassPrefix=true --global-property=apiTests=false,apiDocs=false,modelDocs=false,modelTests=false -i ../../pkg/coreportalapi.swagger.json -g go -o ../../pkg/coreportalapi --package-name coreportalapi
	cd pkg && go fmt ./coreportalapi/... && goimports -w .

.PHONY: generate-importconsole-client
generate-importconsole-client: ## Generate importconsole client
	@echo "==> Generating importconsole client"
	rm -rf pkg/console/import
	cd tools/openapi-generator && npx openapi-generator-cli generate --inline-schema-options RESOLVE_INLINE_ENUMS=true --additional-properties=withGoMod=false,enumClassPrefix=true --global-property=apiTests=false,apiDocs=false,modelDocs=false,modelTests=false -i ../../pkg/console/import.swagger.json -g go -o ../../pkg/console/import --package-name consoleimportapi
	cd pkg && go fmt ./console/import/... && goimports -w .
