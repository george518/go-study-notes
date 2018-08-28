/************************************************************
** @Description: user
** @Author: george hao
** @Date:   2018-08-28 10:34
** @Last Modified by:  george hao
** @Last Modified time: 2018-08-28 10:34
*************************************************************/
package controller

type User struct{}

func (U *User) Create(request string, reply *string) error {
	*reply = " create ok " + request
	return nil
}

func (U *User) Get(request string, reply *string) error {
	*reply = "user info :" + request
	return nil
}
