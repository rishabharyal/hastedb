package core

import "errors"

type SemanticAnalyzer struct {
    tokens []Token
}

func NewSemanticAnalyzer(tokens []Token) *SemanticAnalyzer {
    return &SemanticAnalyzer{tokens: tokens}
}

type Statement struct {
    Type string
    TableStatement TableStatement
    KVStatement KVStatement
}

type TableStatement struct {
    Type string
    Table string
    Columns []string
    Values []string
    Where string
}

type KVStatement struct {
    Type string
    Key string
    Value string
}

func (sa *SemanticAnalyzer) Analyze() (Statement, error) {
    firstToken := sa.tokens[0]
    var result Statement
    var err error
    switch firstToken.Type {
        case "SELECT":
            result, err = sa.analyzeSelect()
        case "INSERT":
            result, err = sa.analyzeInsert()
        case "UPDATE":
            result, err = sa.analyzeUpdate()
        case "DELETE":
            result, err = sa.analyzeDelete()
        case "CREATE":
            result, err = sa.analyzeCreate()
        case "DROP":
            result, err = sa.analyzeDrop()
        case "ALTER":
            result, err = sa.analyzeAlter()
        case "SET":
            result, err = sa.analyzeSet()
        case "GET":
            result, err = sa.analyzeGet()
        default:
            panic("Invalid SQL statement")
    }

    if err != nil {
        return Statement{}, err
    }

    return result, nil

}

func (sa *SemanticAnalyzer) analyzeCreate() (Statement, error) {
    var result Statement
    result.Type = "CREATE"
    return result, nil
}

func (sa *SemanticAnalyzer) analyzeGet() (Statement, error) {
    var result Statement
    result.Type = "GET"
    kvs := KVStatement{Type: "GET"}
    i := 1
    for i < len(sa.tokens) {
        if sa.tokens[i].Type == "IDENTIFIER" {
            kvs.Key = sa.tokens[i].Value
            result.KVStatement = kvs
            return result, nil
        }
        return result, errors.New("unknown_key_at_key_position")
    }
    return result, errors.New("unknown_key_at_key_position")
}

func (sa *SemanticAnalyzer) analyzeSet() (Statement, error) {
    var result Statement
    result.Type = "SET"
    kvs := KVStatement{Type: "SET"}
    i := 1
    for i < len(sa.tokens) {
        if sa.tokens[i].Type == "IDENTIFIER" {
            kvs.Key = sa.tokens[i].Value
            if sa.tokens[i+1].Type == "ASSIGN" {
                kvs.Value = sa.tokens[i+2].Value
                if sa.tokens[i+2].Type == "STRING" {
                    kvs.Value = sa.tokens[i+2].Value
                    result.KVStatement = kvs
                    return result, nil
                }
                // return error
                println(sa.tokens[i+2].Type)
                return result, errors.New("unknown_value_at_value_position")
            }
            return result, errors.New("unknown_value_at_assign_position")
        }
        return result, errors.New("unknown_key_at_key_position")
    }
    return result, errors.New("unknown_key_at_key_position")
}


func (sa *SemanticAnalyzer) analyzeDrop() (Statement, error) {
    var result Statement
    result.Type = "DROP"
    return result, nil
}

func (sa *SemanticAnalyzer) analyzeAlter() (Statement, error) {
    var result Statement
    result.Type = "ALTER"
    return result, nil
}

func (sa *SemanticAnalyzer) analyzeSelect() (Statement, error) {
    var result Statement
    result.Type = "SELECT"
    return result, nil
}

func (sa *SemanticAnalyzer) analyzeInsert() (Statement, error) {
    var result Statement
    result.Type = "INSERT"
    return result, nil
}

func (sa *SemanticAnalyzer) analyzeUpdate() (Statement, error) {
    var result Statement
    result.Type = "UPDATE"
    return result, nil
}

func (sa *SemanticAnalyzer) analyzeDelete() (Statement, error) {
    var result Statement
    result.Type = "DELETE"
    return result, nil
}
