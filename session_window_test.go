package jwp

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestGetCurrentWindowSize(t *testing.T) {
	size, err := client.GetCurrentWindowSize()
	require.NoError(t, err)
	assert.Equal(t, 300, size.Height)
	assert.Equal(t, 400, size.Width)
}

func TestSetCurrentWindowSize(t *testing.T) {
	size := &WindowSize{Height: 800, Width: 600}
	err := client.SetCurrentWindowSize(size)
	require.NoError(t, err)
	size, err = client.GetCurrentWindowSize()
	require.NoError(t, err)
	assert.Equal(t, 800, size.Height)
	assert.Equal(t, 600, size.Width)
}
