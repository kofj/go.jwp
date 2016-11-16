package jwp

import "testing"

func TestCleanup(t *testing.T) {
	t.Log("clear resources")
	cancel()
	cmd.Process.Kill()
}
