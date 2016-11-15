package jwp

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewSession(t *testing.T) {
	session, err := New("http://127.0.0.1:8000")
	require.NoError(t, err)
	assert.NotNil(t, session.Capabilities)
	assert.NotEmpty(t, session.SessionID)
}
