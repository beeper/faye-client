package message_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/beeper/faye-client/pkg/message"
)

type wrapper struct {
	Error *message.Error `json:"error"`
}

func TestError_Decode(t *testing.T) {
	testCases := []struct {
		testCase string
		expected *message.Error
	}{
		{
			testCase: "401::No client ID",
			expected: &message.Error{401, nil, "No client ID"},
		},
		{
			testCase: "402:xj3sjdsjdsjad:Unknown Client ID",
			expected: &message.Error{402, []string{"xj3sjdsjdsjad"}, "Unknown Client ID"},
		},
		{
			testCase: "403:xj3sjdsjdsjad,/foo/bar:Subscription denied",
			expected: &message.Error{403, []string{"xj3sjdsjdsjad", "/foo/bar"}, "Subscription denied"},
		},
		{
			testCase: "404:/foo/bar:Unknown Channel",
			expected: &message.Error{404, []string{"/foo/bar"}, "Unknown Channel"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testCase, func(t *testing.T) {
			var w wrapper
			testCase := fmt.Sprintf(`{"error": "%s"}`, tc.testCase)
			err := json.Unmarshal([]byte(testCase), &w)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, w.Error)
		})
	}
}

func TestError_Encode(t *testing.T) {
	testCases := []struct {
		testCase *message.Error
		expected string
	}{
		{
			testCase: &message.Error{401, nil, "No client ID"},
			expected: "401::No client ID",
		},
		{
			testCase: &message.Error{402, []string{"xj3sjdsjdsjad"}, "Unknown Client ID"},
			expected: "402:xj3sjdsjdsjad:Unknown Client ID",
		},
		{
			testCase: &message.Error{403, []string{"xj3sjdsjdsjad", "/foo/bar"}, "Subscription denied"},
			expected: "403:xj3sjdsjdsjad,/foo/bar:Subscription denied",
		},
		{
			testCase: &message.Error{404, []string{"/foo/bar"}, "Unknown Channel"},
			expected: "404:/foo/bar:Unknown Channel",
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprint(tc.testCase), func(t *testing.T) {
			w := wrapper{tc.testCase}
			b, err := json.Marshal(w)
			assert.NoError(t, err)
			assert.Equal(t, fmt.Sprintf(`{"error":"%s"}`, tc.expected), string(b))
		})
	}
}
