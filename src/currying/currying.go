package main

/**
* 柯里化是函数式编程中的一个术语
* 柯里化 (Currying) 是把多个参数的函数转变为只接受一个参数并返回接收剩余参数的函数的过程，这个过程持续多次直到收集到所有所需参数。
* 在 FP 语言中函数（而不是类）被作为参数进行传递，Currying 常常用于转化一个函数的接口以便于其他代码调用。
* 函数的接口就是它的参数，于是 Currying 通常用于减少函数参数的数量
 */

import (
	"fmt"
)

//这里我们来生产一个汽车 指定车名称和速度
func carFactory(name string) func(int) {
	return func(speed int) {
		fmt.Println("This is a", name, "and it's speed is", speed)
	}
}

func main() {
	//生产一个 Jeep 速度为 130
	carFactory("Jeep")(130)
}
