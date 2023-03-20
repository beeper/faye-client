package message

import "encoding/json"

type ConnectionType string

const (
	ConnectionTypeLongPolling ConnectionType = "long-polling"
	ConnectionTypeWebSocket   ConnectionType = "websocket"
)

type Message struct {
	VeryRaw json.RawMessage `json:"-"`

	Channel                  string           `json:"channel"`
	Version                  string           `json:"version"`
	MinimumVersion           string           `json:"minimumVersion,omitempty"`
	SupportedConnectionTypes []ConnectionType `json:"supportedConnectionTypes,omitempty"`
	ClientID                 string           `json:"clientId,omitempty"`
	Advice                   *Advice          `json:"advice,omitempty"`
	ConnectionType           ConnectionType   `json:"connectionType,omitempty"`
	ID                       string           `json:"id,omitempty"`
	Data                     map[string]any   `json:"data,omitempty"`
	Successful               bool             `json:"successful,omitempty"`
	Error                    *Error           `json:"error,omitempty"`
	Ext                      map[string]any   `json:"ext,omitempty"`
}
