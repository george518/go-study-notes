/************************************************************
** @Description: 可以杀死写成里的命令
** @Author: haodaquan
** @Date:   2018-11-18 20:03
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-11-18 20:03
*************************************************************/
package main

import (
	"fmt"
	"os/exec"
	"time"

	"golang.org/x/net/context"
)

type result struct {
	output []byte
	err    error
}

func main() {
	var cmd *exec.Cmd
	var resultChan chan *result
	var ctx context.Context
	var cancelFunc context.CancelFunc
	var res *result

	ctx, cancelFunc = context.WithCancel(context.TODO())
	resultChan = make(chan *result)
	go func() {
		var output []byte
		var err error
		var re *result
		cmd = exec.CommandContext(ctx,
			"/bin/bash", "-c", " echo hello; sleep 3;echo george")
		output, err = cmd.CombinedOutput()
		re = new(result)
		re.err = err
		re.output = output

		resultChan <- re
	}()

	time.Sleep(5 * time.Second)
	cancelFunc()
	res = <-resultChan

	fmt.Println(string(res.output), res.err)
}
