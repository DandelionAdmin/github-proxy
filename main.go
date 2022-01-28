package main

import (
	"fmt"
	"github.com/creack/pty"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func isWindows() bool {
	flag := runtime.GOOS == "windows"
	return flag
}

func execCommand(cmd string) {
	args := [2]string{
		"bash", "-c",
	}
	if isWindows() {
		args[0] = "cmd"
		args[1] = "/C"
	}
	c := exec.Command(args[0], args[1], cmd)
	f, err := pty.Start(c)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(os.Stdout, f)
	if err != nil {
		return
	}
}

func githubProxy() {
	var cmds = os.Args[1:]
	for i := 1; i < len(cmds); i++ {
		cmd := cmds[i]
		if strings.Contains(cmd, "https://github.com") {
			cmds[i] = strings.ReplaceAll(cmd, "https://github.com", "https://ghproxy.com/https://github.com")
		}

		if strings.Contains(cmd, "https://raw.githubusercontent.com") {
			cmds[i] = strings.ReplaceAll(cmd, "https://raw.githubusercontent.com", "https://ghproxy.com/https://raw.githubusercontent.com")
		}

		if strings.Contains(cmd, "https://gist.github.com") {
			cmds[i] = strings.ReplaceAll(cmd, "https://gist.github.com", "https://ghproxy.com/https://gist.github.com")
		}
	}

	cmd := strings.Join(cmds, " ")
	fmt.Printf("👉 : %v\n", cmd)
	execCommand(cmd)
}

func main() {
	githubProxy()
}
