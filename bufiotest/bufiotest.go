package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please input your name:")
	//以\n为分隔符读取一段内容
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("has en error！%s", err)
	} else {
		//因为是以\n为分割的，所以需要去掉内容的最后一个\n
		fmt.Printf("your name is %s", input[:len(input)-1])
	}
}
