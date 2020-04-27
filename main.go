package main

import (
	"fmt"
)

func main() {
	var a *int // 在类型上代表这个a是一个int变量的地址
	var b int = 10
	a = &b         // &表示取这个int变量的地址 然后赋值给a
	fmt.Println(a) // 打印地址变量
	add(a)
	fmt.Println(b)  // 打印int值
	fmt.Println(&b) // 打印int变量的地址
	fmt.Println(*a) // 打印地址中存的int值

	var c = new(int)
	fmt.Println(c)

	s := "abc"
	bb := []byte(s)
	s2 := string(bb)
	s3 := &s
				  
	fmt.Println(&s)
	fmt.Println(&b)
	fmt.Println(&s2)

	fmt.Println(s3)
	fmt.Println(*s3)

	ar := [3]int{1, 2, 3}
	fmt.Println(ar)
	fmt.Println(&ar[0])
}

func add(a *int) {
	*a += 1
}
