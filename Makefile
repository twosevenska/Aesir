.PHONY: go-get
go-get:
	@echo "Fetching dependencies"
	glide up

.PHONY: go-test
go-test:
	@echo "Run all project tests..."
	go test -v ./...
