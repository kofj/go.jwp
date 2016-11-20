package jwp

type ApiMeta struct {
	SessionId string `json:"sessionId,omitempty"`
	Status    int    `json:"status,omitempty"`
}

type ApiServerStatusBuild struct {
	Version  string `json:"version,omitempty"`
	Revision string `json:"revision,omitempty"`
	Time     string `json:"time,omitempty"`
}

type ApiServerStatusOS struct {
	Arch    string `json:"arch,omitempty"`
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

type ApiServerStatus struct {
	ApiMeta `json:",inline"`
	Value   struct {
		ApiServerStatusBuild `json:"build,omitempty"`
		ApiServerStatusOS    `json:"os,omitempty"`
	} `json:"value"`
}

type ApiCapabilities struct {
	BrowserName              string `json:"browserName"`
	Version                  string `json:"version"`
	DriverName               string `json:"driverName"`
	DriverVersion            string `json:"driverVersion"`
	Platform                 string `json:"platform"`
	JavascriptEnabled        bool   `json:"javascriptEnabled"`
	TakesScreenshot          bool   `json:"takesScreenshot"`
	HandlesAlerts            bool   `json:"handlesAlerts"`
	DatabaseEnabled          bool   `json:"databaseEnabled"`
	LocationContextEnabled   bool   `json:"locationContextEnabled"`
	ApplicationCacheEnabled  bool   `json:"applicationCacheEnabled"`
	BrowserConnectionEnabled bool   `json:"browserConnectionEnabled"`
	CSSSelectorsEnabled      bool   `json:"cssSelectorsEnabled"`
	WebStorageEnabled        bool   `json:"webStorageEnabled"`
	Rotatable                bool   `json:"rotatable"`
	AcceptSslCerts           bool   `json:"acceptSslCerts"`
	NativeEvents             bool   `json:"nativeEvents"`
	Proxy                    `json:"proxy"`

	// special
	UserAgent string `json:"phantomjs.page.settings.userAgent,omitempty"`
}
type Proxy struct {
	/* direct - A direct connection - no proxy in use,
	 * manual - Manual proxy settings configured, e.g. setting a proxy for HTTP, a proxy for FTP, etc,
	 * pac - Proxy autoconfiguration from a URL, autodetect - Proxy autodetection, probably with WPAD,
	 * system - Use system settings
	 */
	ProxyType string `json:"proxyType,omitempty"`

	// Required if proxyType == pac, Ignored otherwise) Specifies the URL to be used for proxy autoconfiguration
	ProxyAutoconfigUrl string `json:"proxyAutoconfigUrl,omitempty"`

	/* (Optional, Ignored if proxyType != manual) Specifies the proxies to be used for FTP, HTTP, HTTPS and SOCKS requests respectively.
	 * Behaviour is undefined if a request is made, where the proxy for the particular protocol is undefined, if proxyType is manual.
	 * Expected format example: hostname.com:1234
	 */
	FtpProxy   string `json:"ftpProxy,omitempty"`
	HttpProxy  string `json:"httpProxy,omitempty"`
	SslProxy   string `json:"sslProxy,omitempty"`
	SocksProxy string `json:"socksProxy,omitempty"`

	SocksUsername string `json:"socksUsername,omitempty"`
	SocksPassword string `json:"socksPassword,omitempty"`

	// Specifies proxy bypass addresses. Format is driver specific.
	NoProxy string `json:"noProxy,omitempty"`
}

type ApiSession struct {
	ApiMeta `json:",inline"`
	Value   ApiCapabilities `json:"value"`
}

type DesiredCapabilities struct {
	UserAgent       string `json:"phantomjs.page.settings.userAgent,omitempty"`
	TakesScreenshot bool   `json:"takesScreenshot,omitempty"`
	Proxy           Proxy  `json:"proxy,omitempty"`
}

type ApiDesiredCapabilities struct {
	DesiredCapabilities `json:"desiredCapabilities,omitempty"`
}

type SessionItem struct {
	Id           string          `json:"id"`
	Capabilities ApiCapabilities `json:"capabilities"`
}

type ApiSessions struct {
	ApiMeta `json:",inline"`
	Value   []SessionItem `json:"value"`
}

type WindowSize struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

type ApiWindowSize struct {
	ApiMeta `json:",inline"`
	Value   WindowSize `json:"value"`
}

type ApiURL struct {
	ApiMeta `json:",inline"`
	Value   string `json:"value"`
}
