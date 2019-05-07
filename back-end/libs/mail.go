package libs

import (
	"fmt"
	"os/exec"
)

func SendMail(to, sub, content string) error {
	exe := fmt.Sprintf("echo %s|mail -s %s %s", content, sub, to)
	cmd := exec.Command("/bin/sh", "-c", exe)
	err := cmd.Run()
	return err
}
