package jwp

import (
	"context"
	"log"
	"net"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	ctx    context.Context
	cancel context.CancelFunc
	cmd    *exec.Cmd

	port   string
	client *Client
)

func init() {
	// start PhantomJS GhostServer
	_, _, port = randPort()
	ctx, cancel = context.WithCancel(context.Background())
	cmd = exec.CommandContext(ctx, "phantomjs", "--webdriver=0:"+port, " &")

	// cmd.Stdout = os.Stderr
	err := cmd.Start()
	if err != nil {
		log.Fatalf("[init]start PhantomJS GhostServer error: %s\n", err.Error())
	}
	log.Printf("listen at 0: %s\t PID: %d\n", port, cmd.Process.Pid)
	time.Sleep(5e9)
}

func randPort() (hostport, host, port string) {
	l, err := net.Listen("tcp", "0:0")
	defer l.Close()
	if err != nil {
		log.Fatalf("New socket error: %s\n", err.Error())
	}
	hostport = l.Addr().String()
	host, port, _ = net.SplitHostPort(hostport)
	return
}

func TestDial(t *testing.T) {
	client = Dial("http://127.0.0.1:" + port)
	require.NotNil(t, client)
}

func TestNewSession(t *testing.T) {
	err := client.NewSession(&ApiDesiredCapabilities{})
	require.NoError(t, err)
	assert.NotNil(t, client.Capabilities)
	assert.NotEmpty(t, client.SessionID)
}
