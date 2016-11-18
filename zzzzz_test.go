package jwp

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	err := client.Delete()
	require.NoError(t, err)

	list, err2 := client.ListSessions()
	require.NoError(t, err2)
	assert.Equal(t, []SessionItem{}, list.Value)
}

func TestCleanup(t *testing.T) {
	t.Log("clear resources")
	cancel()
	cmd.Process.Kill()
}
