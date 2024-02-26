package core

type SemanticAnalyzer struct {
    tokens []Token
}

func NewSemanticAnalyzer(tokens []Token) *SemanticAnalyzer {
    return &SemanticAnalyzer{tokens: tokens}
}

type Statement struct {
    Type string
    Table string
    Columns []string
    Values []string
    Where string
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
    return result, nil
}

func (sa *SemanticAnalyzer) analyzeSet() (Statement, error) {
    var result Statement
    result.Type = "SET"
    return result, nil
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
