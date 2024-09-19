build:
	@go build -o ./bin/bpapi ./cmd/main.go

run: build
	@./bin/bpapi