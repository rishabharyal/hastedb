package connection

type Credentials struct {
    Username string
    Password string
}

type Connection struct {
    credentials Credentials
}

func NewConnection(credentials Credentials) *Connection {
    return &Connection{
        credentials: credentials,
    }
}

func (c *Connection) Connect() (Connection, error) {
    return *c, nil
}

func (c *Connection) Disconnect() (Connection, error) {
    return *c, nil
}
