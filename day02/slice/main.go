package main

import "fmt"

//切片 slice 是一个拥有相同类型元素的可变长度的序列，是基于数组类型的一层封装，支持自动扩容
//切片是引用类型，它的内部结构包含  地址、长度、容量
//切片一般用于快速地操作一块数据集合

func main() {
	//声明
	//var name []T
	var s []int     //定义一个存放int类型元素的切片
	var s1 []string //定义一个存放string类型元素的切片
	fmt.Println(s, s1)
	fmt.Println(s == nil)  //true
	fmt.Println(s1 == nil) //true

	//1 初始化
	s = []int{1, 2, 3}
	s1 = []string{"haha", "hehe", "heihei"}
	fmt.Println(s, s1)
	fmt.Println(s == nil)  //false
	fmt.Println(s1 == nil) //false

	fmt.Printf("len(s):%d  cap(s):%d\n", len(s), cap(s))
	fmt.Printf("len(s1):%d  cap(s1):%d\n", len(s1), cap(s1))

	//2 由数组得到切片
	a1 := [...]int{1, 2, 3, 4, 5}
	s2 := a1[1:4] //[2,3,4]基于数组得到切片，左闭右开
	s3 := a1[:4]  //[0,4]
	s4 := a1[3:]  //[3,len(a1)]
	s5 := a1[:]   //[0,len(a1)]
	fmt.Println(s2, s3, s4, s5)
	//切片的长度是切片包含元素的数量
	//切片的容量是底层数组中从切片第一个元素到数组最后一个元素的长度
	fmt.Printf("len(s2):%d  cap(s2):%d\n", len(s2), cap(s2)) //3 4 cap(s2) = len(a1) - 1
	fmt.Printf("len(s3):%d  cap(s3):%d\n", len(s3), cap(s3)) //4 5 cap(s3) = len(a1) - 0
	fmt.Printf("len(s4):%d  cap(s4):%d\n", len(s4), cap(s4)) //2 2 cap(s4) = len(a1) - 3
	fmt.Printf("len(s5):%d  cap(s5):%d\n", len(s5), cap(s5)) //5 5 cap(s5) = len(a1) - 0
	//3 切片到切片
	s6 := s3[2:]
	fmt.Println(s3)                                          //[1 2 3 4]
	fmt.Println(s6)                                          //[3 4]
	fmt.Printf("len(s6):%d  cap(s6):%d\n", len(s6), cap(s6)) //2 3 //cap(s6) = cap(s3) - 2

}
