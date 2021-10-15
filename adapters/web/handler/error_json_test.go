package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_json_error(t *testing.T) {
	msg := "Hello Json"
	result := jsonError(msg)
	require.Equal(t, []byte(`{"Message":"Hello Json"}`), result)
}
