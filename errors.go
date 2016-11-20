package jwp

import (
	"errors"
	"fmt"
	"net/http"
)

func CheckHTTPStatus(res *http.Response, body []byte) (err error) {
	if res.StatusCode >= 400 && res.StatusCode < 600 {
		fmt.Println("\t...")
		err = errors.New("Status Code: " + res.Status + ", Body: " + string(body))
	}
	return
}
