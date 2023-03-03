package main

import "fmt"

/*
字符串声明和初始化
*/
//func main() {
//	var str string
//	str = "Hello World"
//	str2 := "你好，中国"
//	fmt.Println("str:", str)
//	fmt.Println("str2:", str2)
//}

/*
格式化输出
*/
//func main() {
//	var str string
//	str = "Hello World"
//	ch := str[0]
//	fmt.Printf("The length of %s is %d \n", str, len(str))
//	fmt.Printf("The first character of %s is %c. \n", str, ch)
//}

/*
多行字符串
*/
//func main() {
//	results := `Search results for "Golang":
//	- Go
//	- Golang
//	Golang Programming
//	`
//	fmt.Printf("%s", results)
//}

/*
字符串拼接
*/
//func main() {
//	str := "你好"
//	str += "中国"
//	fmt.Println("str:", str)
//}

/*
字符串切片
*/
//func main() {
//	str := "hello, world"
//	str1 := str[:5]  // 获取索引5（不含）之前的子串
//	str2 := str[7:]  // 获取索引7（含）之后的子串
//	str3 := str[0:5] // 获取从索引0（含）到索引5（不含）之间的子串
//	fmt.Println("str1:", str1)
//	fmt.Println("str2:", str2)
//	fmt.Println("str3:", str3)
//}

/*
字符串遍历
*/
//func main() {
//	str := "Hello, 世界"
//	n := len(str)
//	for i := 0; i < n; i++ {
//		ch := str[i] // 依据下标取字符串中的字符，ch 类型为 byte
//		fmt.Println(i, ch)
//	}
//}

/*
以Unicode字符遍历
*/
//func main() {
//	str := "Hello, 世界"
//	for i, ch := range str {
//		fmt.Println(i, ch) // ch 的类型为 rune
//	}
//}

/*
将Unicode编码转为可打印字符
*/
func main() {
	str := "Hello, 世界"
	for i, ch := range str {
		fmt.Println(i, string(ch))
	}
}
