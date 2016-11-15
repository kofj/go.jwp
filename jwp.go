package jwp

func New(addr string) (j *Session, err error) {
	j = &Session{
		Server: addr,
	}
	err = j.Session(&ApiDesiredCapabilities{
		DesiredCapabilities: DesiredCapabilities{
			Proxy: Proxy{
				ProxyType: "dirct",
			},
		},
	})
	return
}
