package main

import (
	"fmt"
	"math"
)

var (
	name string
	age  int
	ok   bool
)

var s1 string = "hah"
var s3 = "heh"

const pi = 3.1415926

//批量声明常量时，如果某一行没有赋值，默认跟上一行一致
const (
	n1 = 100 //100
	n2       //100
	n3       //100
)

const (
	n8  = iota //0
	n9         //1
	n10        //2
	n11        //3
)

const (
	n4 = iota //0
	n5        //1
	n6        //2
	n7        //3
)

const (
	n12 = iota //0
	n13 = 100  //100
	n14 = iota //2
	n15        //3
)

const (
	m1, m2 = iota + 1, iota //1  0
	m3, m4                  //2  1
)

const (
	d1, d2 = iota + 1, iota + 2 //d1 = 1,d2 = 2
	d3, d4 = iota + 1, iota + 2 //d3 = 2,d4 = 3
)

//字节单位（注意，大写开头的常量为export常量，需要加注释）
const (
	_  = iota
	KB = 1 << (10 * iota) //1024
	MB = 1 << (10 * iota) //1024 * 1024
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {

	s2 := "heh" //只能在函数内使用的赋值语句

	name = "Li Xiang"
	age = 18
	ok = false
	fmt.Println("Hello world!")
	fmt.Printf("name:%s\n", name)
	fmt.Println(s2)

	var (
		x int
		y string
	)

	x, _ = foo()
	_, y = foo()

	fmt.Println("x=", x)
	fmt.Println("y=", y)

	fmt.Println(m1, m2, m3, m4)

	//进制转换
	var i1 = 101
	fmt.Printf("%d\n", i1) //输出十进制
	fmt.Printf("%b\n", i1) //输出二进制
	fmt.Printf("%o\n", i1) //输出八进制
	fmt.Printf("%x\n", i1) //输出十六进制
	var i2 = 077
	fmt.Printf("%T\n", i2) //变量类型输出

	fmt.Println(math.MaxFloat32)

}

func foo() (int, string) {
	return 10, "haha"
}
