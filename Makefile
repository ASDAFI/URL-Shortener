
shortener:
	@echo "Building shortener"
	$(GOBUILD) ./...
	$(GOBUILD) -o $(BINARY_NAME) -v


clean:
	@echo "Cleanning..."
	-rm -f $(BINARY_NAME)
	-rm -f $(BINARY_UNIX)
	-rm -f $(BINARY_SPEC)
	-rm -rf build
	-rm -rf dist
	@echo "Done cleanning."

deployments:
	docker-compose -f ./deployments/compose/docker-compose.yaml up -d


GOCMD=go
GOMOD=GO111MODULE=on $(GOCMD) mod
GOGET=GO111MODULE=on $(GOCMD) "get"

GOBUILD=$(GOCMD) build
BINARY_NAME=shortener
BINARY_UNIX=$(BINARY_NAME)_unix
BINARY_SPEC=$(BINARY_NAME).specs
