package main

import "fmt"

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {

	back1, back2:= 0, 1
	return func() int {
		temp := back1
		// 重新赋值
		back1, back2 = back2, back1 + back2
		//返回temp
		return temp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 5; i++ {
		fmt.Println(f())
	}
}