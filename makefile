run:
	go run ./cmd/main.go

build:
	go build -o bin/interpreter.exe ./cmd/

test:
	go run ./cmd/main.go -json examples/fib.json