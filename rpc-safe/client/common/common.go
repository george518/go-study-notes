/************************************************************
** @Description: main
** @Author: george hao
** @Date:   2018-08-28 10:05
** @Last Modified by:  george hao
** @Last Modified time: 2018-08-28 10:05
*************************************************************/
package common

import (
	"net/rpc"
)

//包路径前缀，抽象路径,必须与服务端一致
const UserServiceName = "user"
const ProductServiceName = "product"

//客户端规则
type RPCClient struct {
	*rpc.Client
}

//var _ RpcInterface = (*RPCClient)(nil)

//可以建立service
func DialService(network, address string) (*RPCClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &RPCClient{Client: c}, nil
}

func (p *RPCClient) UserGet(request string, reply *string) error {
	return p.Client.Call(UserServiceName+".Get", request, reply)
}

func (p *RPCClient) UserCreate(request string, reply *string) error {
	return p.Client.Call(UserServiceName+".Create", request, reply)
}

func (p *RPCClient) ProductCreate(request string, reply *string) error {
	return p.Client.Call(ProductServiceName+".Create", request, reply)
}

func (p *RPCClient) ProductGet(request string, reply *string) error {
	return p.Client.Call(ProductServiceName+".Get", request, reply)
}
