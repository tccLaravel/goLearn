package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) ReflectCallFunc() {
	fmt.Println("Allen.Wu ReflectCallFunc")
}
func (u User) ReflectCallFuncHasArgs(name string, age int) {
	fmt.Println("ReflectCallFuncHasArgs name: ", name, ", age:", age, "and origal User.Name:", u.Name)
}

func (u User) ReflectCallFuncNoArgs() {
	fmt.Println("ReflectCallFuncNoArgs")
}

func main() {
	var num float64 = 1.2345
	fmt.Println("old value of pointer:", num)

	// 通过reflect.ValueOf获取num中的reflect.Value，注意，参数必须是指针才能修改其值
	pointer := reflect.ValueOf(&num)
	fmt.Println("pointer is: ",pointer)
	newValue := pointer.Elem()
	fmt.Println("new value is: ",newValue)
	fmt.Println("type of pointer:", newValue.Type())
	fmt.Println("settability of pointer:", newValue.CanSet())
	// 重新赋值
	newValue.SetFloat(77)
	fmt.Println("new value of pointer:", num) //注意这里 的参数是  num

	user := User{1, "Allen.Wu", 25}

	DoFiledAndMethod(user)

	ReflectCallFuncValue(user)
}

// 通过接口来获取任意参数，然后一一揭晓
func DoFiledAndMethod(input interface{}) {
	fmt.Println("----------DoFiledAndMethod start --------")
	getType := reflect.TypeOf(input)
	fmt.Println("get Type is :", getType)

	fmt.Println("get Type name is :", getType.Name())

	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is:", getValue)

	// 获取方法字段
	// 1. 先获取 interface 的 reflect.Type，然后通过NumField进行遍历
	// 2. 再通过 reflect.Type 的 Field 获取其 Field
	// 3. 最后通过Field的Interface()得到对应的value
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 获取方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}

	fmt.Println("----------DoFiledAndMethod end--------")
}

func ReflectCallFuncValue(user User)  {
	fmt.Println("----------ReflectCallFuncValue start--------")
	// 1. 要通过反射来调用起对应的方法，必须要先通过 reflect.ValueOf(interface) 来获取到 reflect.Value，
	// 得到“反射类型对象”后才能做下一步处理
	getValue := reflect.ValueOf(user)

	// 一定要指定参数为正确的方法名
	// 2. 先看看带有参数的调用方法
	methodValue := getValue.MethodByName("ReflectCallFuncHasArgs")
	fmt.Println("methodValue is :",methodValue)
	args := []reflect.Value{reflect.ValueOf("wudebao"), reflect.ValueOf(30)}
	methodValue.Call(args)

	// 一定要指定参数为正确的方法名
	// 3. 再看看无参数的调用方法
	methodValue = getValue.MethodByName("ReflectCallFuncNoArgs")
	args = make([]reflect.Value, 0)
	methodValue.Call(args)
	fmt.Println("----------ReflectCallFuncValue end--------")
}