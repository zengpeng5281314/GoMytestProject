package main

import (
	"fmt"
	"sync"
)

//五个工人
const noGoroutine = 5

//10个任务
const noTask = 10

var wg = sync.WaitGroup{}

func main() {
	//创建一个缓存池
	task := make(chan int, noTask)

	//启动数量为noGoroutine 的 goroutine
	for i := 1; i <= noGoroutine; i++ {
		wg.Add(1)
		go taskProcess(task, i)
	}

	//向通道里面放入任务
	for i := 1; i <= noTask; i++ {
		task <- i
	}
	//关闭通道(通道关闭后，还能继续读取没有读取的数据)
	close(task)
	//当wg.wait 和 close 互换位置的时候，就会出现死锁
	wg.Wait()
}

func taskProcess(tasks chan int, workerNo int) {
	defer wg.Done()

	for t := range tasks {
		fmt.Printf("Worker %d is procesing Task no:%d \n", workerNo, t)
	}

	fmt.Printf("Worker %d got off work \n", workerNo)
}
