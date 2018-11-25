/**********************************************
** @Des: 内部排序
** @Author: haodaquan
** @Date:   2018-01-21 23:59:31
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-01-22 00:01:23
***********************************************/

package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 3, 5, 2, 9, 8, 11, 7}
	sort.Ints(a)

	fmt.Println(a)
}
