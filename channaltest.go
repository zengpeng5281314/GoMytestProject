package main

import "fmt"

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	//a := <-c
	//fmt.Printf("aaa %d:", a)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

	d := make(chan int, 5)
	d <- 1
	d <- 2
	fmt.Println(<-d)
	fmt.Println(<-d)
	d <- 3
	d <- 4
	d <- 3
	d <- 4
	d <- 4
	// 获取这两个数据
	fmt.Println(<-d)
	fmt.Println(<-d)
	fmt.Println(<-d)
	fmt.Println(<-d)
	fmt.Println(<-d)

	cc := make(chan int, 10)
	go fibonacci(cap(cc), cc)
	for i := range cc {
		fmt.Println(i)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for i := range s {
		sum += s[i]
	}
	c <- sum
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

//1 x=0 y=1    x=1 y=1
//2 x=1 y=1    x=1 y=2
//3 x=1 y=2    x=2 y=3
//4 x=2 y=3    x=3 y=5
//5 x=3 y=5    x=5 y=8
//0, 1, 1,
