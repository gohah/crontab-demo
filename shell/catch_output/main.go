package main

import (
	"fmt"
	"os/exec"
)

//捕获子程序输出
func main() {
	var (
		cmd    *exec.Cmd
		err    error
		output []byte
	)
	//生成cmd
	cmd = exec.Command("/bin/bash", "-c", "echo hello,world")

	//执行了命令，捕获了子程序输出
	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return
	}

	//打印子程序输出
	fmt.Printf("%s", output)

}
