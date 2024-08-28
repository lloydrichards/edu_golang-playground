build:
	go build -o build/playground main.go
	go build -o build/calculator-api calculator-api/main.go
	go build -o build/todo-cli todo-cli/main.go
dev: build
	go run .
dev-api: build
	cd calculator-api && go run *.go