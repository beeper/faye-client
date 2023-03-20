package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/beeper/faye-client/pkg/message"
)

type LongPollingTransport struct {
	url    *url.URL
	client *http.Client

	idCounter int64
}

func NewLongPollingTransport(url *url.URL) *LongPollingTransport {
	return &LongPollingTransport{
		url:    url,
		client: &http.Client{},
	}
}

func (t *LongPollingTransport) Send(ctx context.Context, msg *message.Message) ([]*message.Message, error) {
	t.idCounter++
	msg.ID = strconv.FormatInt(t.idCounter, 32)

	marshaled, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	resp, err := t.client.Post(t.url.String(), "application/json", bytes.NewBuffer(marshaled))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("handshake failed with status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if buf[0] == '[' {
		var messages []*message.Message
		err = json.Unmarshal(buf, &messages)
		if err != nil {
			return nil, err
		}
		return messages, nil
	} else {
		var msg message.Message
		err = json.Unmarshal(buf, &msg)
		if err != nil {
			return nil, err
		}
		return []*message.Message{&msg}, nil
	}
}
