/************************************************************
** @Description: 获取错误和输出
** @Author: haodaquan
** @Date:   2018-11-18 21:08
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-11-18 21:08
*************************************************************/
package main

import (
	"os/exec"

	"bytes"
	"fmt"
)

func main() {
	var cmd *exec.Cmd
	var err error
	bufOut := new(bytes.Buffer)
	bufErr := new(bytes.Buffer)
	cmd = exec.Command(
		"/bin/bash", "-c", "echo 2;sleep 2;echo george")
	cmd.Stdout = bufOut
	cmd.Stderr = bufErr
	cmd.Run()

	err = cmd.Process.Kill()
	fmt.Println(err)
	fmt.Println(bufOut)
	fmt.Println("--------")
	fmt.Println(bufErr)

}
