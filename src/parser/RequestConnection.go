package parser

import (
  "errors"
  "strings"
)

type ConnectionData struct {
  User string
  Pass string
}

type RequestConnection struct {
  splits []string
  connectionString string
  cd ConnectionData
}

func NewRequestConnection(connectionString string) *RequestConnection {
  return &RequestConnection{connectionString: connectionString}
}

func (r *RequestConnection) GetParsedData() interface{} {
  return r.cd
}

func (r *RequestConnection) Parse() (error) {
  r.splits = strings.Split(r.connectionString, "\n")
  if len(r.splits) != 4 {
    return errors.New("Invalid connection request. Please ask the driver maintainer to handle this.")
  }
  // Check if the user and pass are present
  if !strings.Contains(r.splits[2], "USER") || !strings.Contains(r.splits[3], "PASS") {
    return errors.New("Invalid connection request. Please ask the driver maintainer to handle this.")
  }

  // Get the user and pass
  user := strings.Split(r.splits[2], ":")
  pass := strings.Split(r.splits[3], ":")


  // Check if the user and pass are present
  if len(user) != 2 || len(pass) != 2 {
    return errors.New("Invalid connection request. Please ask the driver maintainer to handle this.")
  }

  // Check if the user and pass are not empty
  if user[1] == "" || pass[1] == "" {
    return errors.New("Invalid connection request. Please ask the driver maintainer to handle this.")
  }

  // Set the user and pass after trimming the spaces
  r.cd.User = strings.TrimSpace(user[1])
  r.cd.Pass = strings.TrimSpace(pass[1])

  return nil

}
