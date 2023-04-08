up:
	docker compose up -d

lint:
	gofumpt -w .
	go mod tidy
	golangci-lint run ./...