package main2

import (
	"fmt"
	"time"
)

func main2() {
	messages := make(chan string, 2)

	go f("defer", messages)
	go f("defer2", messages)
	//time.Sleep(time.Duration(10) * time.Second)
	msg := <-messages
	fmt.Println("UN: " + msg)
	//time.Sleep(time.Duration(3) * time.Second)
	msg2 := <-messages
	fmt.Println("DOS: " + msg2)
}

func f(from string, mess chan string) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Duration(i) * time.Second)
		fmt.Println(from, ":", i)
	}
	mess <- "resp: " + from
}
