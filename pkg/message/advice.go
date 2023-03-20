package message

type ReconnectAdvice string

const (
	ReconnectAdviceNone      ReconnectAdvice = "none"
	ReconnectAdviceRetry     ReconnectAdvice = "retry"
	ReconnectAdviceHandshake ReconnectAdvice = "handshake"
)

type Advice struct {
	Reconnect ReconnectAdvice `json:"reconnect,omitempty"`
	Interval  *int            `json:"interval,omitempty"`
	Timeout   *int            `json:"timeout,omitempty"`
}
