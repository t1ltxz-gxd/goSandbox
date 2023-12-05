package tasks

import (
	"fmt"
	"reflect"
)

func Fourteenth() {
	var var1 interface{} = 42
	var var2 interface{} = "Hello, world!"
	var var3 interface{} = true
	var var4 interface{} = make(chan int)

	fmt.Println("Тип переменной var1:", getType(var1))
	fmt.Println("Тип переменной var2:", getType(var2))
	fmt.Println("Тип переменной var3:", getType(var3))
	fmt.Println("Тип переменной var4:", getType(var4))
}

func getType(v interface{}) string {
	return reflect.TypeOf(v).String()
}
