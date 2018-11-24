/************************************************************
** @Description: go_distributed_demo
** @Author: haodaquan
** @Date:   2018-08-25 15:36
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-08-25 15:36
*************************************************************/
package main

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"encoding/json"
	"flag"
	"math/rand"
	"strings"
)

//节点id,ip,port
type NodeInfo struct {
	NodeId     int    `json:"nodeId"`
	NodeIpAddr string `json:"nodeIpAddr"`
	NodePort   string `json:"nodePort"`
}

//source:发起请求的节点
//dest:接收请求的节点
//message:发送的信息
type AddToClusterMessage struct {
	Source  NodeInfo `json:"source"`
	Dest    NodeInfo `json:"dest"`
	Message string   `json:"message"`
}

func (node *NodeInfo) String() string {
	return "NodeInfo:{nodeId:" + strconv.Itoa(node.NodeId) +
		",nodeIpAddr:" + node.NodeIpAddr +
		",nodePort:" + node.NodePort + "}"
}

func (req *AddToClusterMessage) String() string {
	return "AddToClusterMessage:{source:" + req.Source.String() +
		",dest:" + req.Dest.String() +
		",message:" + req.Message + "}"
}

//连接到集群
func connectToCluster(me, dest NodeInfo) bool {
	//tcp连接到dest节点
	conn, err := net.DialTimeout("tcp", dest.NodeIpAddr+":"+dest.NodePort, time.Duration(10)*time.Second)
	if err != nil {
		if _, ok := err.(net.Error); ok {
			fmt.Println("没有连接到集群", me.NodeId)
			return false
		}
	} else {

		fmt.Println("成功连接到集群")
		text := "大哥，小弟（" + strconv.Itoa(me.NodeId) + ")来了，加我"
		//初始化请求信息
		req := getAddToClusterMessage(me, dest, text)
		//发送请求
		json.NewEncoder(conn).Encode(&req)
		var resp AddToClusterMessage
		//接收响应
		json.NewDecoder(conn).Decode(&resp)
		fmt.Println("得到数据响应：", resp.String())

		return true

	}
	return false
}

//获取请求信息
func getAddToClusterMessage(source, dest NodeInfo, message string) AddToClusterMessage {
	return AddToClusterMessage{
		Source: NodeInfo{
			NodeId:     source.NodeId,
			NodeIpAddr: source.NodeIpAddr,
			NodePort:   source.NodePort,
		},
		Dest: NodeInfo{
			NodeId:     dest.NodeId,
			NodeIpAddr: dest.NodeIpAddr,
			NodePort:   dest.NodePort,
		},
		Message: message,
	}
}

//监听连接
func listenOnPort(me NodeInfo) {

	In, _ := net.Listen("tcp", ":"+me.NodePort)
	for {
		con, err := In.Accept()
		if err != nil {
			if _, ok := err.(net.Error); ok {
				fmt.Println("接收信息出错", me.NodeId)
			}
		} else {
			var req AddToClusterMessage
			//接收请求
			json.NewDecoder(con).Decode(&req)
			fmt.Println("接收到的信息是：", req.String())
			text := "兄弟，我收到你的消息了，欢迎"
			resp := getAddToClusterMessage(me, req.Source, text)
			//发送响应
			json.NewEncoder(con).Encode(&resp)
			con.Close()

		}
	}

}

func main() {
	fmt.Println("test start")
	makeMasterOnError := flag.Bool("makeMasterOnErr", false, "ip没有连接到集群，我们将此ip设置为初创节点")
	clusterip := flag.String("clusterip", "127.0.0.1:8080", "任何节点都可以连接本节点ip")
	myport := flag.String("myport", "8001", "正在运行这个节点，端口8001")
	flag.Parse()
	fmt.Println(*makeMasterOnError) //输出信息
	fmt.Println(*clusterip)
	fmt.Println(myport)

	rand.Seed(time.Now().UnixNano())
	myid := rand.Intn(9999999)
	fmt.Println(myid)
	myip, _ := net.InterfaceAddrs()
	fmt.Println("myip:", myip[0])
	//创建当前节点信息
	me := NodeInfo{NodeId: myid, NodeIpAddr: myip[0].String(), NodePort: *myport}
	fmt.Println("当前节点：", me)

	dest := NodeInfo{NodeId: -1, NodeIpAddr: strings.Split(*clusterip, ":")[0], NodePort: strings.Split(*clusterip, ":")[1]}
	fmt.Println("目标节点：", dest)
	ableToConnect := connectToCluster(me, dest)
	//监听其他节点加入集群，或者本节点设置为初始节点
	if ableToConnect || (!ableToConnect && *makeMasterOnError) {
		if *makeMasterOnError {
			fmt.Println("本节点是初始节点")
		}
		listenOnPort(me)
	} else {
		fmt.Println("退出系统，错误")
	}
}
