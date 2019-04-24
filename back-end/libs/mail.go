package libs

import (
	"fmt"
	"os/exec"
)

func SendMail(to, sub, content string) {
	exe := fmt.Sprintf("echo %s|mail -s %s %s", content, sub, to)
	cmd := exec.Command("/bin/sh", "-c", exe)
	cmd.Run()
}
