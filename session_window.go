package jwp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (c *Client) GetCurrentWindowSize() (size *WindowSize, err error) {
	var (
		req  *http.Request
		res  *http.Response
		body []byte
		resp ApiWindowSize
	)

	req, err = http.NewRequest("GET", c.Server+"/session/"+c.SessionID+"/window/current/size", nil)
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

	size = &resp.Value

	return
}

func (c *Client) SetCurrentWindowSize(size *WindowSize) (err error) {
	var (
		req  *http.Request
		res  *http.Response
		body []byte
		buf  = new(bytes.Buffer)
		resp ApiWindowSize
	)

	if size.Height < 100 {
		size.Height = 100
	}
	if size.Width < 100 {
		size.Width = 100
	}

	err = json.NewEncoder(buf).Encode(*size)
	if err != nil {
		return
	}

	req, err = http.NewRequest("POST", c.Server+"/session/"+c.SessionID+"/window/current/size", buf)
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

	return
}
