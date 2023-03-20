package message

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type Error struct {
	ErrorCode    int
	ErrorArgs    []string
	ErrorMessage string
}

func (e *Error) UnmarshalJSON(b []byte) (err error) {
	b = bytes.Trim(b, `"`)
	parts := bytes.Split(b, []byte{':'})
	if len(parts) != 3 {
		return fmt.Errorf("invalid error format: %s", b)
	}
	e.ErrorCode, err = strconv.Atoi(string(parts[0]))
	if err != nil {
		return err
	}
	args := string(parts[1])
	if len(args) > 0 {
		e.ErrorArgs = strings.Split(args, ",")
	}
	e.ErrorMessage = string(parts[2])
	return nil
}

func (e *Error) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, e.String())), nil
}

func (e *Error) String() string {
	return fmt.Sprintf("%d:%s:%s", e.ErrorCode, strings.Join(e.ErrorArgs, ","), e.ErrorMessage)
}

func (e *Error) Error() string {
	return e.String()
}
