package main

import (
	"fmt"
	"strconv"
)

// func main() {
// 	var a *int // 在类型上代表这个a是一个int变量的地址
// 	var b int = 10
// 	a = &b         // &表示取这个int变量的地址 然后赋值给a
// 	fmt.Println(a) // 打印地址变量
// 	add(a)
// 	fmt.Println(b)  // 打印int值
// 	fmt.Println(&b) // 打印int变量的地址
// 	fmt.Println(*a) // 打印地址中存的int值

// 	var c = new(int)
// 	fmt.Println(c)

// 	s := "abc"
// 	bb := []byte(s)
// 	s2 := string(bb)
// 	s3 := &s

// 	fmt.Println(&s)
// 	fmt.Println(&b)
// 	fmt.Println(&s2)

// 	fmt.Println(s3)
// 	fmt.Println(*s3)

// 	ar := [3]int{1, 2, 3}
// 	fmt.Println(ar)
// 	fmt.Println(&ar[0])
// }

// func add(a *int) {
// 	*a += 1
// }
// 将十进制数字转化为二进制字符串
func convertToBin(num int) string {
	s := ""

	if num == 0 {
		return "0"
	}

	// num /= 2 每次循环的时候 都将num除以2  再把结果赋值给 num
	for ; num > 0; num /= 2 {
		lsb := num % 2
		// strconv.Itoa() 将数字强制性转化为字符串
		s = strconv.Itoa(lsb) + s
	}
	return s
}

func main() {
	fmt.Println(
		convertToBin(2),
		convertToBin(19),
		convertToBin(15),
		convertToBin(0),
		convertToBin(8),
		convertToBin(-7),
	)
}
