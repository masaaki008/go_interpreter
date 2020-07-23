default:
	make run

test: lexer

.PHONY: lexer
lexer:
	go test ./lexer

.PHONY: run
run:
	go run main.go
