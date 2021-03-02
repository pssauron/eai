package context

type Connection struct {
	Alias        string
	Parameter    map[string]interface{}
	ComponentKey string
}

func NewConnect(name, key string, param map[string]interface{}) *Connection {
	return &Connection{
		Alias:        name,
		Parameter:    param,
		ComponentKey: key,
	}
}
