package core

import (
    "strings"
)

type Token struct {
    Token string
    Type string
}

type Lexer struct {
    Source string
    Position int
    CurrentChar string
    Tokens []Token
    chars []string
}

func NewLexer(source string) *Lexer {
    l := &Lexer{Source: source, Position: 0, CurrentChar: string(source[0]), Tokens: []Token{}}
    return l
}

// sample query: SELECT * FROM table WHERE id = 1
func (l *Lexer) Tokenize() {
    l.chars = strings.Split(l.Source, "")
    l.Position = 0

    if len(l.chars) == 0 {
        return
    }

    for l.Position < len(l.chars) {

    }

}
