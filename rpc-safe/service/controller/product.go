/************************************************************
** @Description: controller
** @Author: george hao
** @Date:   2018-08-28 16:34
** @Last Modified by:  george hao
** @Last Modified time: 2018-08-28 16:34
*************************************************************/
package controller

type Product struct{}

func (p *Product) Create(request string, reply *string) error {
	*reply = " create product ok " + request
	return nil
}

func (p *Product) Get(request string, reply *string) error {
	*reply = "get product detail : " + request
	return nil
}
