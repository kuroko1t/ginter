package lexer

import (
	"testing"
	"github.com/kuroko1t/ginter/token"
)


func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;

def add(x, y):
  return x + y ;
};

let result = add(five, ten);
`

	tests := []struct{
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INDENT, "\n"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INDENT, "\n"},
		{token.INDENT, "\n"},
		{token.FUNCTION, "def"},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.COLON, ":"},
		{token.INDENT, "\n  "},
		{token.RETURN, "return"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		 {token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.INDENT, "\n"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.INDENT, "\n"},
		{token.INDENT, "\n"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.INDENT, "\n"},
		{token.EOF, ""},
	}
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] = tokentype worng. expected=%q, got=%q", i , tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] = literal worng. expected=%q, got=%q",i , tt.expectedLiteral, tok.Literal)
		}
	}
}
