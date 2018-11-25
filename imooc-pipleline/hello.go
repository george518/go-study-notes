/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2018-01-21 22:07:42
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-01-21 22:08:12
***********************************************/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	len := len(arr)
	for i := 0; i < len; i++ {
		fmt.Println(arr[i])
	}

	for i, v := range arr {
		fmt.Println(i, v)
	}

	//取全部
	slice1 := arr[:]
	fmt.Println(slice1)
	//取下标2到4，前闭后开区间
	slice2 := arr[2:4]
	fmt.Println(slice2)
	//取下标2之后的全部区间
	slice3 := arr[2:]
	fmt.Println(slice3)
	//取到第四个数组，0-3
	slice4 := arr[:4]
	fmt.Println(slice4)
	//对切片还可以取切片，slice2中的下标2到4，前闭后开区间
	slice5 := slice2[2:4]
	fmt.Println(slice5)

	var s []int
	s1 := []int{1, 2, 3}
	s2 := make([]int, 10, 32)

	fmt.Println(s, s1, s2)

	s = append(s, 1)
	s1 = append(s1, 4)
	s2 = append(s2, 11)

	fmt.Println(s, s1, s2)

	ss1 := []int{1, 2, 3}
	ss2 := make([]int, 8, 16)

	copy(ss2, ss1)
	fmt.Println(ss2)

	ss3 := []int{1, 2, 3, 4, 5, 6}
	ss4 := append(ss3[:2], ss3[3:]...)
	fmt.Println(ss4)

}
