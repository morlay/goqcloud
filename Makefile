gen:
	vgo run ./cmd/main.go

test:
	vgo test -v -race ./...
