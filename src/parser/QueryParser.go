package parser

type QueryParser struct {
    queryString string
}

func NewQueryParser(queryString string) *QueryParser {
    return &QueryParser{
        queryString: queryString,
    }
}

func (q *QueryParser) Parse() error {
    return nil
}
