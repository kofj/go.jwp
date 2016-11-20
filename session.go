package jwp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Server       string
	SessionID    string
	Capabilities ApiCapabilities
}

func (s *Client) NewSession(dc *ApiDesiredCapabilities) (err error) {
	var (
		req  *http.Request
		res  *http.Response
		body []byte
		buf  = new(bytes.Buffer)
		resp ApiSession
	)

	if dc.DesiredCapabilities.Proxy.ProxyType == "" {
		dc.DesiredCapabilities.Proxy.ProxyType = "dirct"
	}

	err = json.NewEncoder(buf).Encode(dc)
	if err != nil {
		return
	}

	req, err = http.NewRequest("POST", s.Server+"/session", buf)
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

	s.SessionID = resp.SessionId
	s.Capabilities = resp.Value
	return
}

func (s *Client) ListSessions() (resp *ApiSessions, err error) {
	var (
		req  *http.Request
		res  *http.Response
		body []byte
	)

	req, err = http.NewRequest("GET", s.Server+"/sessions", nil)
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

func (s *Client) GetCapabilities() (caps *ApiCapabilities, err error) {
	var (
		req  *http.Request
		res  *http.Response
		body []byte
		resp ApiSession
	)

	req, err = http.NewRequest("GET", s.Server+"/session/"+s.SessionID, nil)
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

	s.Capabilities = resp.Value
	caps = &s.Capabilities

	return
}

func (c *Client) Delete() (err error) {
	var (
		req  *http.Request
		res  *http.Response
		body []byte
		resp ApiMeta
	)

	req, err = http.NewRequest("DELETE", c.Server+"/session/"+c.SessionID, nil)
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

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}
	if err = CheckStatus(resp.Status); err != nil {
		return
	}
	return
}
