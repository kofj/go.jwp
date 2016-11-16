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
	list, err := sess.ListSessions()
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.IsType(t, &ApiSessions{}, list)
	assert.IsType(t, []SessionItem{}, list.Value)
	assert.IsType(t, ApiCapabilities{}, list.Value[0].Capabilities)
	id = list.Value[0].Id
}
