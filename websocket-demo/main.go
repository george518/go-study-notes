/************************************************************
** @Description: PPGo_Websocket
** @Author: haodaquan
** @Date:   2018-08-25 10:00
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-08-25 10:00
*************************************************************/
package main

import (
	"net/http"

	"github.com/george518/go-study-notes/websocket-demo/service"
)

func main() {
	http.HandleFunc("/ws", service.WsHandler)
	http.ListenAndServe("0.0.0.0:7777", nil)
}
