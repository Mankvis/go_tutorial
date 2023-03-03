package main

import "fmt"

/*
*
数据类型默认值
*/
//func main() {
//	var v1 int      // 整型
//	var v2 string   // 字符串
//	var v3 bool     // 布尔型
//	var v4 [10]int  // 数组，数组元素类型为整型
//	var v5 struct { // 结构体，成员变量 f 的类型为 64 位浮点型
//		f float64
//	}
//	var v6 *int            // 指针类型，指向整型
//	var v7 map[string]int  // map（字典），key 为字符串类型，value 为整型
//	var v8 func(a int) int // 函数类型，参数类型为整数，返回值类型为整数
//	fmt.Println("v1:", v1)
//	fmt.Println("v2:", v2)
//	fmt.Println("v3:", v3)
//	fmt.Println("v4:", v4)
//	fmt.Println("v5:", v5)
//	fmt.Println("v6:", v6)
//	fmt.Println("v7:", v7)
//	fmt.Println("v8:", v8)
//}

/**
变量初始化
*/
//func main() {
//	var v1 int = 10 // 方式一，常规的初始化操作
//	var v2 = 10     // 方式二，此时变量类型会被编译器自动推导出来
//	v3 := 10        // 方式三，可以省略 var，编译器可以自动推导出 v3 的类型
//	fmt.Println("v1:", v1)
//	fmt.Println("v2:", v2)
//	fmt.Println("v3:", v3)
//}

/**
变量多重赋值
*/
//func main() {
//	i := "i"
//	j := "j"
//	j, i = i, j
//	fmt.Println("i:", i)
//	fmt.Println("j:", j)
//}

/**
匿名变量
*/
//func main() {
//	userName, nickName := GetName()
//	fmt.Println("username:", userName, "nickName", nickName)
//	// 只获取nickName
//	_, nickName2 := GetName()
//	fmt.Println("nickName:", nickName2)
//}
//
//func GetName() (userName, nickName string) {
//	return "Tony", "stackTony"
//}

/*
常量定义
*/
/*func main() {
	const Pi = 3.14159265358979323846
	const zero = 0.0 // 无类型浮点型常量
	const (
		size int64 = 1024
		eof        = -1 // 无类型整型常量
	)
	const u, v float64 = 0, 3    // u = 0.0, v = 3.0，常量的多重赋值
	const a, b, c = 3, 4, "fool" // a = 3, b = 4, c = "foo", 无类型整型和字符串常量
	fmt.Println("Pi: ", Pi)
	fmt.Println("zero: ", zero)
	fmt.Println("size: ", size)
	fmt.Println("eof: ", eof)
	fmt.Println("u: ", u)
	fmt.Println("v: ", v)
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
}*/

/*
*
预定义常量
*/
//func main() {
//	const ( // iota 被重置为 0
//		c0 = iota // c0 = 0
//		c1 = iota // c1 = 1
//		c2 = iota // c2 = 2
//	)
//	const (
//		u = iota * 2 // u = 0
//		v = iota * 2 // u = 2
//		w = iota * 2 // w = 4
//	)
//	const x = iota // x = 0
//	const y = iota // y = 0
//	fmt.Println("c0:", c0)
//	fmt.Println("c1:", c1)
//	fmt.Println("c2:", c2)
//	fmt.Println("u:", u)
//	fmt.Println("v:", v)
//	fmt.Println("w:", w)
//	fmt.Println("x:", x)
//	fmt.Println("y:", y)
//}

func main() {
	const (
		Sunday = iota
		Monday
		TuesDay
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays
	)
	fmt.Println("Sunday:", Sunday)
	fmt.Println("Monday:", Monday)
	fmt.Println("TuesDay:", TuesDay)
	fmt.Println("Wednesday:", Wednesday)
	fmt.Println("Thursday:", Thursday)
	fmt.Println("Friday:", Friday)
	fmt.Println("Saturday:", Saturday)
	fmt.Println("numberOfDays:", numberOfDays)
}
