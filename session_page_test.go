package jwp

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestGetCurrentPageURL(t *testing.T) {
	url, err := client.GetCurrentPageURL()
	require.NoError(t, err)
	assert.Equal(t, "about:blank", url)
}

func TestOpen(t *testing.T) {
	err := client.Open("https://www.baidu.com/")
	require.NoError(t, err)
	url, _ := client.GetCurrentPageURL()
	require.Equal(t, "https://www.baidu.com/", url)
}

func TestExecute(t *testing.T) {
	ret, err := client.Execute(&ExecuteScript{
		Script: "return document.getElementsByTagName('title')[0].text",
		Args:   []map[string]interface{}{{"prefix": "title"}},
	})
	require.NoError(t, err)
	assert.Equal(t, "百度一下，你就知道", ret)
}
