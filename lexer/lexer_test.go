package lexer

import (
	"monkey-ball/token"
	"testing"
)

func TestSimpleToken(t *testing.T) {
	input := `=`
	l := New(input)
	tok := l.NextToken()

	if tok.Type != token.ASSIGN {
		t.Fatalf("SimpleTokenTest failed - Expected=%q and got=%q",
			token.ASSIGN, tok.Type)
	}
}

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokenType wrong.  expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong.  expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
