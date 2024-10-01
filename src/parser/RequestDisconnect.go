package parser

import (
  "errors"
  "strings"
)

type RequestDisconnect struct {
  splits []string
  disconnectString string
  dd DisconnectData
}

type DisconnectData struct {
  Token string
}

func NewRequestDisconnect(disconnectString string) *RequestDisconnect {
  return &RequestDisconnect{disconnectString: disconnectString}
}

func (r *RequestDisconnect) Parse() (error) {
  r.splits = strings.Split(r.disconnectString, "\n")
  if len(r.splits) != 3 {
    return errors.New("Invalid number of arguments")
  }

  if r.splits[2] == "" {
    return errors.New("Token cannot be empty")
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

  r.dd.Token = tokenString

  return nil
}

func (r *RequestDisconnect) GetParsedData() interface{} {
  return r.dd
}
