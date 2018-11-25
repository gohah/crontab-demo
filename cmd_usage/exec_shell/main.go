package main

import (
	"fmt"
	"os/exec"
)

//执行shell命令
func main() {
	var (
		cmd *exec.Cmd
		err error
	)
	cmd = exec.Command("/bin/bash", "-c", "echo hello,world")
	err = cmd.Run()
	fmt.Println(err)
}
