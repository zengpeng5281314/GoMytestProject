package main

import "fmt"

func main() {
	fmt.Println("hello word!")

	var a int = 4
	var b int32
	var c float32
	var ptr *int

	/* 运算符实例 */
	fmt.Printf("第一行 - a 变量类型为 = %T\n", a)
	fmt.Printf("第二行 - b 变量类型为 = %T \n", b)
	fmt.Printf("第三行 - c 变量类型为 = %T \n", c)
	//fmt.Printf("第四行 - b 变量类型为 = %T \n", ptr)

	/* & 和 * 运算符实例 */
	ptr = &a
	fmt.Printf("第一行 - a 变量类型为 = %d\n", a)
	fmt.Printf("第四行 - b 变量类型为 = %d \n", *ptr)

	number1 := make([]int, 4, 4)
	for i := 0; i < 4; i++ {
		number1[i] = i
	}
	fmt.Printf("第四行 - b 变量类型为 = %d \n", number1)
}
