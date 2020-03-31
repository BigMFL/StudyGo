package main

import "fmt"

//append() 为切片增加元素

func main() {
	s1 := []string{"beijing", "shanghai"}
	fmt.Printf("s1=%v  len(s1)=%d  cap(s1)=%d\n", s1, len(s1), cap(s1)) //s1=[beijing shanghai]  len(s1)=2  cap(s1)=2
	s1 = append(s1, "guangzhou")                                        //调用append函数必须用原变量接收返回值
	fmt.Printf("s1=%v  len(s1)=%d  cap(s1)=%d\n", s1, len(s1), cap(s1)) //s1=[beijing shanghai guangzhou]  len(s1)=3  cap(s1)=4 扩容后cap不足，扩大cap（底层数组更换）
	s1 = append(s1, "hangzhou", "chengdu")
	fmt.Printf("s1=%v  len(s1)=%d  cap(s1)=%d\n", s1, len(s1), cap(s1)) //s1=[beijing shanghai guangzhou hangzhou chengdu]  len(s1)=5  cap(s1)=8
	ss := []string{"xian", "jinan", "suzhou"}
	s1 = append(s1, ss...)                                              //...表示拆开
	fmt.Printf("s1=%v  len(s1)=%d  cap(s1)=%d\n", s1, len(s1), cap(s1)) //s1=[beijing shanghai guangzhou hangzhou chengdu xian jinan suzhou]  len(s1)=8  cap(s1)=8
}
