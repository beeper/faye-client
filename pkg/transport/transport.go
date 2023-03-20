package transport

import "encoding/json"

type Transport interface {
	Send(json.Marshaler) (json.Unmarshaler, error)
}
