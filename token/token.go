package token

import (
	//"fmt"
)

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}


const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"
	IDENT = "IDENT"
	INT = "INT"
	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"
	TRUE = "TRUE"
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"

	INDENT = "INDENT"

	EQ = "=="
	NOT_EQ = "!="

	LT = "<"
	GT = ">"

	COMMA = ","
	SEMICOLON = ";"
	COLON = ":"

	LPAREN = "("
	RPAREN = "("
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET = "LET"

	FILE = "FILE"
	LINE = "LINE"
)

var keywords = map[string]TokenType{
	"def": FUNCTION,
	"let": LET,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	//fmt.Println("Lookupident",ident)
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
