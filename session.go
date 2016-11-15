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
