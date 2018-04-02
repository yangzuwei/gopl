package main

import (
	"fmt"
	"reflect"
	"stack"
	"time"
)

func badCall() {
	panic("bad end")
}

func gc() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()

	fmt.Println("calling test \r\n")
	badCall()

}

func say(words string) {
	fmt.Println(words)
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	time.Sleep(1 * time.Second)
	c <- sum
	c <- 99
}

func fibo(ch, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {

	go gc()
	go say("hello i am a gorotine")

	fmt.Println("after calling test \r\n")

	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x := <-c
	y := <-c
	z := <-c
	q := <-c
	fmt.Println(x, y, z, q)

	ch := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		quit <- 0
	}()

	fibo(ch, quit)

	var funStack stack.Stack

	funStack.Push(func() int {
		return 2 * 2
	})

	funStack.Push(func() int {
		return 3 * 3
	})

	for funStack.Len() > 0 {
		f, _ := funStack.Pop()
		fmt.Println(reflect.TypeOf(f))
		//通过类型断言将其转换为可使用的func后再调用
		fmt.Println(f.(func() int)())
	}
}
