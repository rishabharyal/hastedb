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
    TableDDLStatement TableDDLStatement
}

type TableStatement struct {
    Type string
    Table string
    Columns []string
    Values []string
    Where string
}

type TableDDLStatement struct {
    Type string
    Table string
    Columns []string
    Constraints []string
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
            panic("unsupported_query")
    }

    if err != nil {
        return Statement{}, err
    }

    return result, nil

}

func (sa *SemanticAnalyzer) analyzeCreate() (Statement, error) {
    var result Statement
    result.Type = "CREATE"
    TableStatement := TableDDLStatement{Type: "CREATE"}
    i := 1
    if sa.tokens[i].Type != "TABLE" {
        return result, errors.New("unknown_entity_expected_table")
    }

    i++

    if !(sa.tokens[i].Type == "IDENTIFIER" || sa.tokens[i].Type == "STRING") {
        return result, errors.New("table_name_unspecified")
    }

    TableStatement.Table = sa.tokens[i].Value
    i++

    if sa.tokens[i].Type != "LEFTPARENTHESIS" {
        return result, errors.New("unknown_entity_expected_left_parenthesis")
    }

    i++

    // Here we need to handle the columns, their types and other traits
    for i < len(sa.tokens) {
        if sa.tokens[i].Type == "RIGHTPARENTHESIS" {
            break
        }
        if !(sa.tokens[i].Type == "IDENTIFIER" || sa.tokens[i].Type == "STRING") {
            return result, errors.New("unknown_column_name")
        }

        column := sa.tokens[i].Value
        i++


    }

    return result, nil

}

func (sa *SemanticAnalyzer) analyzeGet() (Statement, error) {
    var result Statement
    result.Type = "GET"
    kvs := KVStatement{Type: "GET"}
    if len(sa.tokens) != 2 {
        return result, errors.New("invalid_get_statement")
    }
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
    if len(sa.tokens) < 3 {
        return result, errors.New("invalid_set_statement")
    }
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
