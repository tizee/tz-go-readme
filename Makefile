
.PHONY: default test

# go test
test:
	go test ./client/util ./mdblock

lint:
	golangci-lint run

fmt:
	go fmt .