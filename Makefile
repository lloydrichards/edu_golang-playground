build:
	go build -o calculator-api/bin/calculator-api calculator-api/cmd/main.go
	go build -o todo-cli/bin/todo-cli todo-cli/cmd/main.go
dev-api: build
	cd ./calculator-api && go run ./cmd/main.go
dev-cli: build
	cd ./todo-cli && go run ./cmd/main.go