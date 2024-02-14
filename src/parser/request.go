package parser

import (
	"errors"
	"strings"
)

type Request struct {
    query string
}

type ConnectionData struct {
    Username string
    Password string
}

func NewRequest(query string) *Request {
    return &Request{
        query: query,
    }
}

func (r *Request) Parse() (string, error) {
    split := strings.Split(r.query, "\n");
    REQ_TYPE := split[0]
    
    switch REQ_TYPE {
        case "CONNECT":
            // Check if index exists
            if len(split) != 4 {
                return "", errors.New("Invalid connection request. Please ask the driver maintainer to handle this.")
            }
            // Check if the user and pass are present
            if !strings.Contains(split[1], "USER") || !strings.Contains(split[2], "PASS") {
                return "", errors.New("Invalid connection request. Please ask the driver maintainer to handle this.")
            }

            // Get the user and pass
            user := strings.Split(split[1], ":")
            pass := strings.Split(split[2], ":")

            // Check if the user and pass are present
            if len(user) != 2 || len(pass) != 2 {
                return "", errors.New("Invalid connection request. Please ask the driver maintainer to handle this.")
            }

            // Check if the user and pass are not empty
            if user[1] == "" || pass[1] == "" {
                return "", errors.New("Invalid connection request. Please ask the driver maintainer to handle this.")
            }

            return REQ_TYPE, nil
        case "DISCONNECT":
        case "QUERY":
        default:
            return "", nil
    }
    return REQ_TYPE, nil
}
