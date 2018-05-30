gen:
	vgo run ./cmd/main.go

test:
	vgo test -v -race ./generator/codegen
	vgo test -v -race ./generator
	vgo test -v -race ./examples
