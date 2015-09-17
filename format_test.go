package gophercloud

import (
	"bytes"
	th "github.com/rackspace/gophercloud/testhelper"
	"os/exec"
	"testing"
)

func TestCodeFormattedCorrectly(t *testing.T) {
	cmd := exec.Command("gofmt", "-l", ".")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	th.CheckNoErr(t, err)
	th.CheckEquals(t, "", out.String())
}
