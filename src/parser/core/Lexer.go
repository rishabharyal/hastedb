package core

import (
    "strings"
    "strconv"
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
    Value string
    Type string
}

type Lexer struct {
    Source string
    Position int
    CurrentChar string
    Tokens []Token
    chars []string
    isString bool
    isNumber bool
    isFloat bool
    currentTokenString string
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
        println(l.CurrentChar)
        // if the instance in a string, any character is candidate for a string literal
        if l.isString {
            if l.CurrentChar != "\"" {
                // Add the current character to the chars array
                l.currentTokenString += l.CurrentChar
                l.Position++
                continue
            }
        } else {
            // Sinc this is already a string, and we also got a closing double quote, we can add the string to the tokens array
            l.Tokens = append(l.Tokens, Token{Value: l.currentTokenString, Type: "string"})
            l.currentTokenString = ""
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
            l.currentTokenString = ""
            l.Position++
            continue
        }

        if l.isNumber {
            if l.CurrentChar == "." {
                if l.isFloat {
                    // Its already a float, so we need to throw error
                    println("Error: Invalid float")
                    return
                }
                l.isFloat = true
            } else if !isStringANumber(l.currentTokenString) {
                // Its not a valid number, so we need to throw error
                println("Error: Invalid number")
                return
            }
            l.currentTokenString += l.CurrentChar
            l.Position++
            continue
        }

        l.currentTokenString += l.CurrentChar
        
        if isStringANumber(l.currentTokenString) {
            l.isNumber = true 
        }

        l.Position++

    }
    

}

func isStringANumber(s string) bool {
    // Try to convert the string to a float64.
    // The function returns an error if the conversion fails.
    _, err := strconv.ParseFloat(s, 64)
    return err == nil
}

func (l *Lexer) GetTokens() []Token {
    return l.Tokens
}

func (l *Lexer) HandleToken() {
}
