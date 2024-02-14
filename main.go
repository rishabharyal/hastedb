package main

import (
	"hastedb/src/parser"
)

type Response struct {
    Status string
    Message string
    Action string
}

func NewResponse(status string, message string, action string) *Response {
    return &Response{
        Status: status,
        Message: message,
        Action: action,
    }
}

func (r *Response) Render() (string, error) {
    return r.Status, nil
}

func main() {

    requestString := 
    `CONNECT

    USER: admin
    PASS: password`

    request_parser := parser.NewRequest(requestString);

    connection_data, err := request_parser.Parse();
    if err != nil {
        response := NewResponse("ERROR", err.Error(), "PARSE")
        response.Render()
        return
    }

    println(connection_data)


    




}
