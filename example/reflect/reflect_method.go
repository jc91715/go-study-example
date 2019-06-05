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

	user := &User{Name: "hello name0"}

	obj := reflect.ValueOf(user)

	//调用带参数的方法
	args := []reflect.Value{reflect.ValueOf("hello set Name1")}
	obj.MethodByName("SetName").Call(args) //输出设置成功：hello set Name1

	//调用非参数的方法
	args = make([]reflect.Value, 0)
	obj.MethodByName("GetName").Call(args) //获取成功：hello set Name1

}
