/************************************************************
** @Description: 这里演示命令执行exec包的基本用法
** @Author: haodaquan
** @Date:   2018-11-18 19:51
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-11-18 19:51
*************************************************************/
package main

import (
	"fmt"
	"os/exec"
)

func main() {

	var cmd *exec.Cmd
	var err error
	var output []byte

	//执行
	cmd = exec.Command("/bin/bash", "-c", "echo hello")
	err = cmd.Run()
	fmt.Println(err)

	//有输出的执行
	cmd = exec.Command("/bin/bash", "-c", "echo george")
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(output))
}
