package parser

type RequestQuery struct {
    splits []string
    qd QueryData
}

type QueryData struct {
    Query string
    Bindings []string
}

func NewRequestQuery(splits []string) *RequestQuery {
    return &RequestQuery{splits: splits}
}

func (r *RequestQuery) Parse() (error) {
    return nil
}

func (r *RequestQuery) GetParsedData() interface{} {
    return r.qd
}
