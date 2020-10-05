MAIN = main.go

default:
	make run

test: lexer parser ast

.PHONY: lexer
lexer:
	go test -v ./lexer

.PHONY: parser
parser:
	go test -v ./parser

.PHONY: ast
ast:
	go test -v ./ast

.PHONY: run
run:
	go run ${MAIN}
