package core

import (
    "testing"
)

func TestLexer(t *testing.T) {
    lexer := NewLexer("SELECT * FROM table WHERE id = 1")
    lexer.Tokenize()
    tokens := lexer.GetTokens()
    println(tokens)
}
