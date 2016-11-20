package jwp

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestCurrentWindowSize(t *testing.T) {
	size, err := client.GetCurrentWindowSize()
	require.NoError(t, err)
	assert.Equal(t, 300, size.Height)
	assert.Equal(t, 400, size.Width)
}
