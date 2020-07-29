
.PHONY: default test

# go test
test:
	go test ./client ./mdblock

lint:
	golangci-lint run

fmt:
	go fmt .