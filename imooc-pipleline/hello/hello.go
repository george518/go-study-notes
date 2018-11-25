/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2018-01-21 22:05:59
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-01-21 23:27:47
***********************************************/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello world%s", request.FormValue("name"))
	})

	http.ListenAndServe(":8888", nil)
}
