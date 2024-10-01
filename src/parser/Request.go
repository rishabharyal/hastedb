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
  split := strings.SplitN(r.query, "\n", 1)
  REQ_TYPE := split[0]

  var parser Parser

  switch REQ_TYPE {
  case "CONNECT":
    parser = NewRequestConnection(split[1])
  case "DISCONNECT":
    parser = NewRequestDisconnect(split[1])
  case "QUERY":
    parser = NewRequestQuery(split[1])
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
