package parser

import (
	"errors"
	"strings"
)

type RequestQuery struct {
    splits []string
    qd QueryData
}

type QueryData struct {
    Query string
    Bindings []string
    Token string
    Type string
}

func NewRequestQuery(splits []string) *RequestQuery {
    return &RequestQuery{splits: splits}
}

func (r *RequestQuery) Parse() (error) {
    if len(r.splits) < 5 {
        return errors.New("Invalid number of arguments")
    }

    if r.splits[2] == "" {
        return errors.New("Token cannot be empty")
    }

    if r.splits[4] == "" {
        return errors.New("Type cannot be empty")
    }

    if r.splits[6] == "" {
        return errors.New("Query cannot be empty")
    }

    token := strings.Split(r.splits[2], ":");
    if len(token) != 2 {
        return errors.New("Invalid token format")
    }

    if strings.TrimSpace(token[0]) != "TOKEN" {
        return errors.New("Invalid token format")
    }

    tokenString := strings.TrimSpace(token[1])
    if tokenString == "" {
        return errors.New("Token cannot be empty")
    }

    r.qd.Token = tokenString

    queryType := strings.Split(r.splits[4], ":");
    if len(queryType) != 2 {
        return errors.New("Invalid query type format")
    }

    if strings.TrimSpace(queryType[0]) != "TYPE" {
        return errors.New("Invalid query type format")
    }

    queryTypeString := strings.TrimSpace(queryType[1])
    if queryTypeString == "" {
        return errors.New("Query type cannot be empty")
    }

    r.qd.Type = queryTypeString

    r.qd.Query = r.splits[6]

    if len(r.splits) > 7 {
        for i := 5; i < len(r.splits); i++ {
            r.qd.Bindings = append(r.qd.Bindings, r.splits[i])
        }
    }

    return nil
}

func (r *RequestQuery) GetParsedData() interface{} {
    return r.qd
}
