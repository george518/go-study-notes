/************************************************************
** @Description: base
** @Author: george hao
** @Date:   2018-08-28 10:33
** @Last Modified by:  george hao
** @Last Modified time: 2018-08-28 10:33
*************************************************************/
package base

import (
	"net/rpc"
)

//包路径前缀，抽象路径
const UserServiceName = "user"
const ProductServiceName = "product"

//创建rpcServer的接口,一个查询一个创建
type ServiceInterface = interface {
	Get(request string, reply *string) error
	Create(request string, reply *string) error
}

//注册服务
func RegisterService(model string, svc ServiceInterface) error {
	return rpc.RegisterName(model, svc)
}
