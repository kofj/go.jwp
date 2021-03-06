package jwp

import (
	"bytes"
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

func (c *Client) Open(url string) (err error) {
	var (
		req  *http.Request
		res  *http.Response
		body []byte
		buf  = new(bytes.Buffer)
		resp ApiMeta
	)

	err = json.NewEncoder(buf).Encode(PageURL{url})
	if err != nil {
		return
	}

	req, err = http.NewRequest("POST", c.Server+"/session/"+c.SessionID+"/url", buf)
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

// POST /session/:sessionId/forward
// Navigate forwards in the browser history, if possible.
func (c *Client) Forward() (err error) { return }

// POST /session/:sessionId/back
// Navigate backwards in the browser history, if possible.
func (c *Client) Back() (err error) { return }

// POST /session/:sessionId/refresh
// Refresh the current page.
func (c *Client) Refresh() (err error) { return }

// POST /session/:sessionId/execute
/* Inject a snippet of JavaScript into the page for execution
 * in the context of the currently selected frame. The executed script
 * is assumed to be synchronous and the result of evaluating the script
 * is returned to the client.
 *
 * JSON Parameters:
 * script - {string} The script to execute.
 * args - {Array.<*>} The script arguments.
 */
func (c *Client) Execute(es *ExecuteScript) (result string, err error) {
	var (
		req  *http.Request
		res  *http.Response
		body []byte
		buf  = new(bytes.Buffer)
		resp ApiExecuteScriptResult
	)

	err = json.NewEncoder(buf).Encode(*es)
	if err != nil {
		return
	}

	req, err = http.NewRequest("POST", c.Server+"/session/"+c.SessionID+"/execute", buf)
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
	result = resp.Value

	return
}
