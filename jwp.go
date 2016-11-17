package jwp

func Dial(addr string) *Client {
	return &Client{
		Server: addr,
	}
}
