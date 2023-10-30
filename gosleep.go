package main

import (
	"fmt"

	"github.com/lihan3238/gosleep/SleepSort"
)

func main() {

	fmt.Println("Go to sleep\n请输入数组大小：\n")
	var n int
	fmt.Scanln(&n)
	fmt.Println("请输入数组：\n")
	//用户输入数组

	unsort := []int{}
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		unsort = append(unsort, num)
	}
	fmt.Println("排序结果：\n")
	//用户输入数组
	for _, num := range SleepSort.GoSleep(unsort) {
		fmt.Println(num)
	}
}
