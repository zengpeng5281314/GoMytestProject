package main

import (
	"fmt"
	"math/rand"
)

var done = false
var Mess = make(map[int]bool)

//随机生成一些数字写入通道中，然后另一个通道读取该通道的值，如果通道中的值已经存在了， 就关闭生成通道，将前面所有的值都相加起来
func main() {
	//随机生成50个数字
	for i := 0; i < 50; i++ {
		i := getRandomNum(50, 10)
		fmt.Printf("随机生成的数字为：%d \n", i)
	}

	//定义一个通道A
	A := make(chan int)
	//定义一个通道B
	B := make(chan int)
	go sendNum(50, 0, A)
	go readRum(A, B)
	fmt.Printf("总数为：%d \n", sumRum(B))
}

//生成随机函数
func getRandomNum(max int, min int) int {
	return rand.Intn(max-min) + min
}

func sendNum(max, min int, A chan<- int) {
	for {
		if done {
			close(A)
			return
		}
		A <- getRandomNum(max, min)
	}
}

func readRum(A chan int, B chan int) {
	for i := range A {
		fmt.Printf("num is :%d \n", i)
		_, ok := Mess[i]
		if ok {
			fmt.Printf(" duplication num is :%d  \n", i)
			done = true
		} else {
			Mess[i] = true
			B <- i
		}
	}
	close(B)
}

func sumRum(B chan int) int {
	var sum int
	for i := range B {
		sum += i
	}
	fmt.Printf("The Sum is :%d \n", sum)
	return sum
}
