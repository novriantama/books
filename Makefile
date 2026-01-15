CMD_PATH=cmd/main.go

deps:
	@echo "Checking and downloading dependencies..."
	go mod download

tidy:
	@echo "Tidying up modules..."
	go mod tidy

run:
	@echo "Running the application..."
	go run $(CMD_PATH)