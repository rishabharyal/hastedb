package core

import (
	"errors"
	"strconv"
	"strings"
)


var TOKEN_TYPES = map[string]string{
    "SELECT": "SELECT",
    "FROM": "FROM",
    "WHERE": "WHERE",
    "AND": "AND",
    "OR": "OR",
    "NOT": "NOT",
    "IN": "IN",
    "LIKE": "LIKE",
    "BETWEEN": "BETWEEN",
    "IS": "IS",
    "NULL": "NULL",
    "ORDER": "ORDER",
    "BY": "BY",
    "ASC": "ASC",
    "DESC": "DESC",
    "LIMIT": "LIMIT",
    "OFFSET": "OFFSET",
    "INSERT": "INSERT",
    "INTO": "INTO",
    "VALUES": "VALUES",
    "UPDATE": "UPDATE",
    "SET": "SET",
    "DELETE": "DELETE",
    "CREATE": "CREATE",
    "TABLE": "TABLE",
    "DROP": "DROP",
    "ALTER": "ALTER",
    "ADD": "ADD",
    "PRIMARY": "PRIMARY",
    "KEY": "KEY",
    "FOREIGN": "FOREIGN",
    "REFERENCES": "REFERENCES",
    "CONSTRAINT": "CONSTRAINT",
    "UNIQUE": "UNIQUE",
}

var OPERATORS = map[string]string{
    "=": "EQUAL",
    "!=": "NOTEQUAL",
    ">": "GREATERTHAN",
    "<": "LESSTHAN",
    ">=": "GREATERTHANOREQUAL",
    "<=": "LESSTHANOREQUAL",
    "&&": "AND",
    "||": "OR",
    "+": "PLUS",
    "-": "MINUS",
    "*": "MULTIPLY",
    "/": "DIVIDE",
    "%": "MODULUS",
    "(": "LEFTPARENTHESIS",
    ")": "RIGHTPARENTHESIS",
    ",": "COMMA",
    ".": "DOT",
    ";": "SEMICOLON",
    ":": "COLON",
    "?": "WILDCARD",
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
        // if the instance in a string, any character is candidate for a string literal
        if l.isString {
            if l.CurrentChar != "\"" {
                // Add the current character to the chars array
                l.currentTokenString += l.CurrentChar
            } else {
                // Sinc this is already a string, and we also got a closing double quote, we can add the string to the tokens array
                l.Tokens = append(l.Tokens, Token{Value: l.currentTokenString, Type: "string"})
                l.currentTokenString = ""
                l.isString = false
            }
            l.Position++
            continue
        }

        if l.CurrentChar == "\"" {
            err := l.HandleToken()
            if err != nil {
                println(err)
                return
            }
            l.currentTokenString = ""
            l.isString = true
            l.Position++
            continue
        }
        

        // if the current character is a space, we need to save current token and reset the chars array
        if l.CurrentChar == " " || l.CurrentChar == ";" {
            l.HandleToken()
            l.currentTokenString = ""
            l.Position++
            continue
        }

        // check if current character is a single valued tokens like =, >, <, etc
        if l.isSingleValuedToken() {
            err := l.HandleToken() // handles the current token so that we can process the next token
            if err != nil {
                println(err)
                return
            }
            l.Position++
            l.currentTokenString = l.CurrentChar
            if l.Position < len(l.chars) && l.handleMultiValuedToken(l.CurrentChar) {
                l.Position++
                l.currentTokenString = l.CurrentChar
                continue;
            }
            l.handleSingleValuedToken()
            l.currentTokenString = ""
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

    l.HandleToken()
    

}

func isStringANumber(s string) bool {
    _, err := strconv.ParseFloat(s, 64)
    return err == nil
}

func (l *Lexer) GetTokens() []Token {
    return l.Tokens
}

func (l *Lexer) isSingleValuedToken() bool {
    return l.CurrentChar == "=" || l.CurrentChar == ">" || l.CurrentChar == "<" || l.CurrentChar == "!" || l.CurrentChar == "+" || l.CurrentChar == "-" || l.CurrentChar == "*" || l.CurrentChar == "/" || l.CurrentChar == "%" || l.CurrentChar == "(" || l.CurrentChar == ")" || l.CurrentChar == "," || l.CurrentChar == "." || l.CurrentChar == ";" || l.CurrentChar == ":" || l.CurrentChar == "?"
}

func (l *Lexer) handleMultiValuedToken(lastToken string) bool {
    combinedToken := lastToken + l.chars[l.Position]
    matches := combinedToken == "!=" || combinedToken == "<=" || combinedToken == ">=" || combinedToken == "&&" || combinedToken == "||"
    if !matches {
        return false
    }
    l.Tokens = append(l.Tokens, Token{Value: combinedToken, Type: OPERATORS[combinedToken]})
    return true
}

func (l *Lexer) handleSingleValuedToken() {
    l.Tokens = append(l.Tokens, Token{Value: l.currentTokenString, Type: OPERATORS[l.currentTokenString]})
}


func (l *Lexer) HandleToken() error {
    if l.currentTokenString == "" {
        return nil
    }

    if l.isNumber {
        l.Tokens = append(l.Tokens, Token{Value: l.currentTokenString, Type: "NUMBER"})
        l.isNumber = false
        return nil
    }

    if l.isString {
        // If the string is not closed, we need to throw an error
        return errors.New("String started but not closed.")
    }

    if _, ok := TOKEN_TYPES[strings.ToUpper(l.currentTokenString)]; ok {
        l.Tokens = append(l.Tokens, Token{Value: l.currentTokenString, Type: TOKEN_TYPES[strings.ToUpper((l.currentTokenString))]})
        return nil
    }

    l.Tokens = append(l.Tokens, Token{Value: l.currentTokenString, Type: "IDENTIFIER"})

    return nil 
}
