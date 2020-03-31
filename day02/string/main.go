package main

import (
	"fmt"
	"strings"
)

//字符串
func main() {
	//转义
	//path := "D:\GO\src\github.com"
	path := "\"D:\\GO\\src\\github.com\""
	fmt.Printf(path)

	s := "I'm ok!"
	fmt.Printf(s)

	//按原有格式输出

	//多行字符串
	s2 := `
哈哈
			窗前明月光
			疑是地上霜

			`
	fmt.Printf(s2)

	path1 := `D:\GO\src\github.com`
	fmt.Printf(path1)

	//字符串常用操作
	fmt.Println(len(s2)) //打印字符串长度
	//字符串拼接
	name := "我 "
	world := "真帅"
	ss := name + world
	fmt.Println(ss)
	fmt.Printf("%s%s", name, world)
	ss1 := fmt.Sprintf("%s%s", name, world) //拼接并返回
	fmt.Println(ss1)
	//分割
	s3 := "haha,haha"
	ret := strings.Split(s3, ",")
	fmt.Println(ret)
	//包含
	fmt.Println(strings.Contains(s3, "haha"))
	//前缀
	fmt.Println(strings.HasPrefix(s3, "h"))
	//后缀
	fmt.Println(strings.HasSuffix(s3, "h"))
	//子串出现的位置
	fmt.Println(strings.Index(s3, "hah"))     //0  子串第一次出现的位置
	fmt.Println(strings.LastIndex(s3, "hah")) //5  子串最后一次出现的位置
	//拼接
	fmt.Println(strings.Join(ret, "+"))

	//字符串修改
	sss2 := "白萝卜"
	sss3 := []rune(sss2)
	sss3[0] = '红'
	fmt.Println(string(sss3))

}
