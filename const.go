package jwp

// UA
const (
	Chrome54         = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.59 Safari/537.36"
	AppleiPhoneiOS91 = `Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1`
)

const (
	BlankScreen = "iVBORw0KGgoAAAANSUhEUgAAAZAAAAEsCAYAAADtt+XCAAAACXBIWXMAAAsTAAALEwEAmpwYAAAB6ElEQVR4nO3BMQEAAADCoPVPbQ0PoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHg3VJUAAfkd0+cAAAAASUVORK5CYII="
)

// status code
const (
	StatusSuccess                    = 0
	StatusNoSuchDriver               = 6
	StatusNoSuchElement              = 7
	StatusNoSuchFrame                = 8
	StatusUnknownCommand             = 9
	StatusStaleElementReference      = 10
	StatusElementNotVisible          = 11
	StatusInvalidElementState        = 12
	StatusUnknownError               = 13
	StatusElementIsNotSelectable     = 15
	StatusJavaScriptError            = 17
	StatusXPathLookupError           = 19
	StatusTimeout                    = 21
	StatusNoSuchWindow               = 23
	StatusInvalidCookieDomain        = 24
	StatusUnableToSetCookie          = 25
	StatusUnexpectedAlertOpen        = 26
	StatusNoAlertOpenError           = 27
	StatusScriptTimeout              = 28
	StatusInvalidElementCoordinates  = 29
	StatueIMENotAvailable            = 30
	StatusIMEEngineActivationFailed  = 31
	StatusInvalidSelector            = 32
	StatusSessionNotCreatedException = 33
	StatusMoveTargetOutOfBounds      = 34
)
