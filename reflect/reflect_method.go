package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
}

func (u *User) GetName() string {
	fmt.Println("获取成功：" + u.Name)
	return u.Name
}
func (u *User) SetName(name string) {
	u.Name = name
	fmt.Println("设置成功：" + u.Name)

}

func main() {

	user := &User{Name: "hello world"}

	obj := reflect.ValueOf(user)

	setName := obj.MethodByName("SetName")
	args := []reflect.Value{reflect.ValueOf("hello set Name")}
	setName.Call(args)

	getName := obj.MethodByName("GetName")
	args = make([]reflect.Value, 0)
	getName.Call(args)

}
