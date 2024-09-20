build:
	@go build -o ./bin/bpapi ./cmd/api/main.go

run: build
	@./bin/bpapi