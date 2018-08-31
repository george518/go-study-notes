/************************************************************
** @Description: snowflake 分布式ID生成器
** @Author: george hao
** @Date:   2018-08-31 09:03
** @Last Modified by:  george hao
** @Last Modified time: 2018-08-31 09:03
*************************************************************/
package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/snowflake"
)

func main() {
	n, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 0; i < 3; i++ {
		id := n.Generate()
		fmt.Println("id ", id)
		fmt.Println(" node:", id.Node(), " step:", id.Step(), " time:", id.Time(), "\n")
	}
}
