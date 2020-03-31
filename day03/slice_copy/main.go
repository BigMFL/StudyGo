package main

import (
	"fmt"
	"sort"
)

// copy()

func main() {
	s1 := []int{1, 2, 3}
	s2 := s1
	s3 := make([]int, 3, 3)
	copy(s3, s1) // copy 将底层数组的值复制到另一个切片的底层数组
	fmt.Println(s1, s2, s3)
	s1[0] = 100
	fmt.Println(s1, s2, s3)

	//删除切片中的元素
	a1 := [...]int{1, 3, 5}
	s11 := a1[:]
	s11 = append(s11[:1], s11[2:]...)                                         //修改了底层数组
	fmt.Printf("s11=%v  len(s11)=%d  cap(s11)=%d\n", s11, len(s11), cap(s11)) //s11=[1 5]  len(s11)=2  cap(s11)=3
	fmt.Println(a1)                                                           // [1,5,5]

	sli := make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		sli = append(sli, i)
	}
	fmt.Println(sli) //[0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]

	var arr = [...]int{1, 3, 5, 2, 4, 6, 9, 7}
	sort.Ints(arr[:]) //对切片进行排序
	fmt.Println(arr)

}
