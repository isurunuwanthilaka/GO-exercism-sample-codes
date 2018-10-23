package main

import (
	"fmt"
	"math"
	"time"
)

func f(s string) {
	fmt.Println(s)
}

func main() {
	nan := math.NaN()
	fmt.Print(nan)
	var num int
	channel := make(chan string, 1)
	go goroute(channel)
	go goroute1(&num)
	for {
		time.Sleep(3 * time.Second)
		fmt.Print(<-channel)
		fmt.Printf("Num val in main channel : %v", num)
	}
}
func goroute1(num *int) {
	count := 0
	for {
		time.Sleep(1 * time.Second)
		count++
		fmt.Println(count)
		*num = count
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
