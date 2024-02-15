package parser

type RequestDisconnect struct {
    splits []string
    dd DisconnectData
}

type DisconnectData struct {
    Token string
}

func NewRequestDisconnect(splits []string) *RequestDisconnect {
    return &RequestDisconnect{splits: splits}
}

func (r *RequestDisconnect) Parse() (error) {
    return nil
}

func (r *RequestDisconnect) GetParsedData() interface{} {
    return r.dd
}
