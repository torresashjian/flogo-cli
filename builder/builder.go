package builder

import (
	"os/exec"
)

func DoGoGet(url string) error {
	cmd := exec.Command("go", "get", "-d", url)
	return cmd.Run()
}
