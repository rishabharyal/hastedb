package parser

import (
	"errors"
	"strings"
)

type Parser interface {
    Parse() error
    GetParsedData() interface{}
}

type Request struct {
    query string
    parser Parser
}

func NewRequest(query string) *Request {
    return &Request{
        query: query,
    }
}

func (r *Request) Parse() (string, error) {
    split := strings.Split(r.query, "\n");
    REQ_TYPE := split[0]

    var parser Parser
    
    switch REQ_TYPE {
        case "CONNECT":
            parser = NewRequestConnection(split)
        case "DISCONNECT":
            parser = NewRequestDisconnect(split)
        case "QUERY":
            parser = NewRequestQuery(split)
        default:
            return "", errors.New("Invalid request type")
    }

    error := parser.Parse()
    if error != nil {
        return "", error
    }

    r.parser = parser

    return REQ_TYPE, error
}

func (r *Request) GetParser() Parser {
    return r.parser
}
