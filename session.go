package jwp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Session struct {
	Server       string
	SessionID    string
	Capabilities ApiCapabilities
}

func (s *Session) Session(dc *ApiDesiredCapabilities) (err error) {
	var (
		req  *http.Request
		res  *http.Response
		body []byte
		buf  = new(bytes.Buffer)
		v    ApiSession
	)

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

	json.Unmarshal(body, &v)
	s.SessionID = v.SessionId
	s.Capabilities = v.Value
	return
}

func (s *Session) ListSessions() (sess *ApiSessions, err error) {
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

	json.Unmarshal(body, &sess)
	return
}

func (s *Session) GetCapabilities() (caps *ApiCapabilities, err error) {
	var (
		req  *http.Request
		res  *http.Response
		body []byte
		sess ApiSession
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

	json.Unmarshal(body, &sess)
	s.Capabilities = sess.Value
	caps = &s.Capabilities

	return
}
