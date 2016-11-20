package jwp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (c *Client) GetCurrentPageURL() (url string, err error) {
	var (
		req  *http.Request
		res  *http.Response
		body []byte
		resp ApiURL
	)

	req, err = http.NewRequest("GET", c.Server+"/session/"+c.SessionID+"/url", nil)
	if err != nil {
		return
	}

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	if err = CheckHTTPStatus(res, body); err != nil {
		return
	}

	json.Unmarshal(body, &resp)
	if err = CheckStatus(resp.Status); err != nil {
		return
	}

	url = resp.Value

	return
}
