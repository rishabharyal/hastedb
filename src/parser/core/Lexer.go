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
    isString bool
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
        l.CurrentChar = l.chars[l.Position]
        // if the instance in a string, any character is candidate for a string literal
        if l.isString {
            if l.CurrentChar != "\"" {
                // Add the current character to the chars array
                l.chars = append(l.chars, l.CurrentChar)
                l.Position++
                continue
            }
        }
        // Sinc this is already a string, and we also got a closing double quote, we can add the string to the tokens array
        l.Tokens = append(l.Tokens, Token{Token: strings.Join(l.chars, ""), Type: "string"})
        l.chars = []string{}
        l.isString = false
        l.Position++
        continue
    }

}
