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

func CheckStatus(status int) (err error) {
	switch status {
	case StatusSuccess:
		err = nil
	case StatusNoSuchDriver:
		err = errors.New("StatusNoSuchDriver")

	case StatusNoSuchElement:
		err = errors.New("StatusNoSuchElement")

	case StatusNoSuchFrame:
		err = errors.New("StatusNoSuchFrame")

	case StatusUnknownCommand:
		err = errors.New("StatusUnknownCommand")

	case StatusStaleElementReference:
		err = errors.New("StatusStaleElementReference")

	case StatusElementNotVisible:
		err = errors.New("StatusElementNotVisible")

	case StatusInvalidElementState:
		err = errors.New("StatusInvalidElementState")

	case StatusUnknownError:
		err = errors.New("StatusUnknownError")

	case StatusElementIsNotSelectable:
		err = errors.New("StatusElementIsNotSelectable")

	case StatusJavaScriptError:
		err = errors.New("StatusJavaScriptError")

	case StatusXPathLookupError:
		err = errors.New("StatusXPathLookupError")

	case StatusTimeout:
		err = errors.New("StatusTimeout")

	case StatusNoSuchWindow:
		err = errors.New("StatusNoSuchWindow")

	case StatusInvalidCookieDomain:
		err = errors.New("StatusInvalidCookieDomain")

	case StatusUnableToSetCookie:
		err = errors.New("StatusUnableToSetCookie")

	case StatusUnexpectedAlertOpen:
		err = errors.New("StatusUnexpectedAlertOpen")

	case StatusNoAlertOpenError:
		err = errors.New("StatusNoAlertOpenError")

	case StatusScriptTimeout:
		err = errors.New("StatusScriptTimeout")

	case StatusInvalidElementCoordinates:
		err = errors.New("StatusInvalidElementCoordinates")

	case StatusIMEEngineActivationFailed:
		err = errors.New("StatusIMEEngineActivationFailed")

	case StatusInvalidSelector:
		err = errors.New("StatusInvalidSelector")

	case StatusSessionNotCreatedException:
		err = errors.New("StatusSessionNotCreatedException")

	case StatusMoveTargetOutOfBounds:
		err = errors.New("StatusMoveTargetOutOfBounds")
	}
	return
}
