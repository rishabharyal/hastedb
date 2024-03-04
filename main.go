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

func (r *Response) Render(){
    // Write the byte back
    println(r.Message)
}

func main() {

    requestString := 
    `QUERY

TOKEN: d6t243fbhmwtdg86y3wp77qrj

TYPE: CREATE

CREATE TABLE users (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255)
);

0 => 1
1 = "rishabh@gmail.com
`

    request_parser := parser.NewRequest(requestString);

    connection_type, err := request_parser.Parse();
    if err != nil {
        response := NewResponse("ERROR", err.Error(), "PARSE")
        response.Render()
        return
    }

    switch connection_type {
        case "CONNECT":
            conn_parser := request_parser.GetParser();
            if connection_data, ok := conn_parser.GetParsedData().(parser.ConnectionData); ok {
                println(connection_data.User)
                println(connection_data.Pass)
                // we will handle the connection here...
            }
        case "DISCONNECT":
            conn_parser := request_parser.GetParser();
            if connection_data, ok := conn_parser.GetParsedData().(parser.DisconnectData); ok {
                println(connection_data.Token)
            }
        case "QUERY":
            conn_parser := request_parser.GetParser();
            if connection_data, ok := conn_parser.GetParsedData().(parser.QueryData); ok {
                queryParser := parser.NewQueryParser(connection_data.Query);
                queryParser.Parse();
            }
    }
}
