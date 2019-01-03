/************************************************************
** @Description: interfaces
** @Author: george hao
** @Date:   2018-12-12 14:40
** @Last Modified by:  george hao
** @Last Modified time: 2018-12-12 14:40
*************************************************************/
package main

import "fmt"

func main() {
	var animal animal
	c := &cat{"mi"}
	d := &dog{"wang"}

	animal = c
	animal.eat()
	animal.sound()

	animal = d
	animal.eat()
	animal.sound()
}

//动物animal cat dog

type animal interface {
	eat()
	sound()
}

type cat struct {
	name string
}

type dog struct {
	name string
}

func (c *cat) eat() {
	fmt.Println("cat eats fish")
}

func (c *cat) sound() {
	fmt.Println("cat mi")
}

func (d *dog) eat() {
	fmt.Println("dog eats bone")
}

func (d *dog) sound() {
	fmt.Println("dog wang")
}
