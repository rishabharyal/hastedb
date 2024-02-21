package core

import (
    "strings"
)


var TOKEN_TYPES = map[string]string{
    "SELECT": "select",
    "FROM": "from",
    "WHERE": "where",
    "AND": "and",
    "OR": "or",
    "NOT": "not",
    "IN": "in",
    "LIKE": "like",
    "BETWEEN": "between",
    "IS": "is",
    "NULL": "null",
    "ORDER": "order",
    "BY": "by",
    "ASC": "asc",
    "DESC": "desc",
    "LIMIT": "limit",
    "OFFSET": "offset",
    "INSERT": "insert",
    "INTO": "into",
    "VALUES": "values",
    "UPDATE": "update",
    "SET": "set",
    "DELETE": "delete",
    "CREATE": "create",
    "TABLE": "table",
    "DROP": "drop",
    "ALTER": "alter",
    "ADD": "add",
    "PRIMARY": "primary",
    "KEY": "key",
    "FOREIGN": "foreign",
    "REFERENCES": "references",
    "CONSTRAINT": "constraint",
    "UNIQUE": "unique",
}

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
        } else {
            // Sinc this is already a string, and we also got a closing double quote, we can add the string to the tokens array
            l.Tokens = append(l.Tokens, Token{Token: strings.Join(l.chars, ""), Type: "string"})
            l.chars = []string{}
            l.isString = false
            l.Position++
            continue
        }
        if l.CurrentChar == "\"" {
            l.isString = true
            l.Position++
            continue
        }

        // if the current character is a space, we need to save current token and reset the chars array
        if l.CurrentChar == " " {
            l.HandleToken()
            l.chars = []string{}
            l.Position++
            continue
        }

        l.chars = append(l.chars, l.CurrentChar)
        l.Position++

    }
    

}

func (l *Lexer) HandleToken() {
}
