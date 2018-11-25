package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	err    error
	output []byte
}

//执行1个cmd,让它在一个协程里面去执行，让它执行2秒
//1秒的时候我们杀死它
func main() {
	var (
		cmd        *exec.Cmd
		ctx        context.Context
		cancelFunc context.CancelFunc
		resultChan chan *result
		res        *result
	)

	resultChan = make(chan *result)

	ctx, cancelFunc = context.WithCancel(context.TODO())

	go func() {
		var (
			err    error
			output []byte
		)

		//执行带上下文的命令
		cmd = exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 3;echo hello,world")

		//捕获输出
		output, err = cmd.CombinedOutput()

		//协程之间通信用chan
		resultChan <- &result{
			err:    err,
			output: output,
		}

	}()

	//继续往下走
	time.Sleep(1 * time.Second)

	//取消上下文
	cancelFunc()

	res = <-resultChan

	fmt.Println(string(res.output))
}
