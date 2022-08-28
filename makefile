build:
	@echo "Formatting.."
	@go fmt ./...
	@echo "Building.."
	@go build -v -o boilerplate-service cmd/main.go

run:
	@echo "Running.."
	make build
	@./boilerplate-service