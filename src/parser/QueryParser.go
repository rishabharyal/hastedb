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
    semanticAnalyzer := core.NewSemanticAnalyzer(tokens)
    statement, err := semanticAnalyzer.Analyze()
    if err != nil {
        println(err.Error())
        return err
    }
    println(statement.Type)
    return nil
}
