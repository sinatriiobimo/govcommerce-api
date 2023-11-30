.PHONY: prettify cover generate

prettify:
	gofmt -s -w .

cover:
	@go test ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out

generate:
	@go generate -x ./...