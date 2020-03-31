package main

import "fmt"

//数组
//存放元素的容器
//必须制定存放元素的类型和容量（长度）
//数组的长度是数组类型的一部分
func main() {
	var a1 [3]bool //[false false false]
	var a2 [4]bool
	fmt.Printf("%T\n %T\n", a1, a2)

	//数组的初始化
	//如果不初始化，默认元素都是默认值
	fmt.Println(a1)
	//1
	a1 = [3]bool{true, true, true}
	//2 根据初始值自动推断数组长度
	a10 := [...]int{0, 1, 3, 4, 5, 6, 43, 2, 1}
	fmt.Println(a10)
	//3
	a3 := [5]int{1, 3}
	fmt.Println(a3)
	//4 根据索引
	a4 := [5]int{0: 1, 3: 3}
	fmt.Println(a4)

	//遍历
	citys := [...]string{"Beijing", "ShangHai", "Guangzhou"}
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}

	for i, v := range citys {
		fmt.Println(i, v)
	}

	for _, v := range citys {
		fmt.Println(v)
	}

	//多维数组[[1 2] [3 4] [5 6]]
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a11)

	//多维数组遍历
	for _, v := range a11 {
		for _, v1 := range v {
			fmt.Println(v1)
		}
	}

	//数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2)
}
