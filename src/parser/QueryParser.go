package parser

import "hastedb/src/parser/core"

type QueryParser struct {
    queryString string
}

func NewQueryParser(queryString string) *QueryParser {
    return &QueryParser{
        queryString: queryString,
    }
}

func (q *QueryParser) Parse() error {
    lexer := core.NewLexer(q.queryString)
    lexer.Tokenize()
    tokens := lexer.GetTokens()
    for _, token := range tokens {
        println(token)
    }
    return nil
}
