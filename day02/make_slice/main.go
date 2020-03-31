package main

import "fmt"

// make()函数创造切片(可以指定长度和容量)
// make([]T,len,cap)
func main() {
	s1 := make([]int, 5) //cap省略的时候默认等于len
	s2 := make([]int, 5, 10)
	fmt.Printf("s1=%v  len(s1)=%d  cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v  len(s2)=%d  cap(s2)=%d\n", s2, len(s2), cap(s2))

	//判断切片是否为空，应该判断len是否为0

	//切片的赋值
	s3 := []int{1, 3, 5}
	s4 := s3
	fmt.Println(s3, s4)
	s3[0] = 1000 //切片是引用类型，s3和s4指向的是同一个底层数组
	fmt.Println(s3, s4)

	//切片遍历
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}

	for i, v := range s3 {
		fmt.Println(i, v)
	}

}
