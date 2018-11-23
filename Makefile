gen:
	go run ./__generator__/main.go

test:
	go test -v -race ./...
