package main

import "fmt"

/*
*
定义浮点型数据
*/
//func main() {
//	var floatValue1 float32
//	floatValue1 = 10
//	floatValue2 := 10.0 // 如果不加小数点，floatValue2 会被推导为整型而不是浮点型
//	floatValue3 := 1.1e-10
//	fmt.Println("floatValue1:", floatValue1)
//	fmt.Println("floatValue2:", floatValue2)
//	fmt.Println("floatValue3:", floatValue3)
//}

/*
浮点型精度
*/
//func main() {
//	floatValue4 := 0.1
//	floatValue5 := 0.7
//	floatValue6 := floatValue4 + floatValue5
//	fmt.Println("floatValue6:", floatValue6)
//	// 输出 floatValue6: 0.7999999999999999
//}

/*
浮点数比较
*/
//func main() {
//	floatValue1 := 10
//	floatValue2 := 10.0 // 如果不加小数点，floatValue2 会被推导为整型而不是浮点型
//	p := 0.00001
//	// 判断 floatValue1 与 floatValue2 是否相等
//	if math.Dim(float64(floatValue1), floatValue2) < p {
//		fmt.Println("floatValue1 和 floatValue2 相等")
//	}
//}

/*
复数类型
*/
func main() {
	var complexValue complex64
	complexValue = 1.10 + 10i   // 由两个 float32 实数构成的复数类型
	complexValue2 := 1.10 + 10i // 和浮点型一样，默认自动推导的实数类型是float64，所以 complexValue2 是 complex128 类型
	complexValue3 := complex(1.10, 10)
	var complexValue4 complex64 = complex(1, 2)
	var complexValue5 complex128 = complex(11.11, 22.22)
	fmt.Println("complexValue:", complexValue)
	fmt.Println("complexValue2:", complexValue2)
	fmt.Println("complexValue3:", complexValue3)
	fmt.Println("complexValue4:", complexValue4)
	fmt.Println("complexValue5:", complexValue5)
	fmt.Printf("complexValue5 实部: %v, 虚部:%v", real(complexValue5), imag(complexValue5))
}
