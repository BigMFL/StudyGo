package main

import "fmt"

//fmt占位符总结

func main() {
	//int
	var n = 100
	fmt.Printf("%T\n", n) //输出类型
	fmt.Printf("%v\n", n) //输出变量的值
	fmt.Printf("%b\n", n) //输出二进制
	fmt.Printf("%d\n", n) //输出十进制
	fmt.Printf("%o\n", n) //输出八进制
	fmt.Printf("%x\n", n) //输出十六进制
	//字符串
	var s = "hahaha"
	fmt.Printf("%s\n", s)  //hahaha
	fmt.Printf("%v\n", s)  //hahaha
	fmt.Printf("%#v\n", s) //"hahaha"

}
