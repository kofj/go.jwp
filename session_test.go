package jwp

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	id string
)

func TestListSessions(t *testing.T) {
	list, err := client.ListSessions()
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.IsType(t, &ApiSessions{}, list)
	assert.IsType(t, []SessionItem{}, list.Value)
	assert.IsType(t, ApiCapabilities{}, list.Value[0].Capabilities)
	id = list.Value[0].Id
}

func TestGetCapabilities(t *testing.T) {
	capabilities, err := client.GetCapabilities()
	require.NoError(t, err)
	assert.Equal(t, "phantomjs", capabilities.BrowserName)
}
