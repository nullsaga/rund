rund: build
	@./bin/rund

build:
	@go build -o bin/rund cmd/rund/main.go
