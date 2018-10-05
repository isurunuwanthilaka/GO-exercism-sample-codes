package main

import (
	"fmt"
	"time"
)

func f(s string) {
	fmt.Println(s)
}

func main() {
	channel := make(chan string, 1)
	go goroute(channel)
	for {
		time.Sleep(3 * time.Second)
		fmt.Print(<-channel)
	}
}

func goroute(channel chan string) {
	c := 0
	for {
		time.Sleep(1 * time.Second)
		c++
		fmt.Printf("In Go Route %v\n", c)
		if len(channel) != 0 {
			<-channel
			fmt.Print("Unloading\n")
		}
		channel <- fmt.Sprintf("In Channel %v\n", c)
	}
}
