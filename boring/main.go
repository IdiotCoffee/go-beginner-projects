package main

import (
	"fmt"
	"math/rand"
	"time"
)

// here, boring() is a fan-out pattern. We put 2 go-routines in this fan-out.
// Function fanIn will then take the 2 responses from the 2 go-routines, and combine them into 1 response
// This will be returned as a channel, and in func main(), I am printing the values from that channel.
func main() {
	c := fanIn(boring("Ishaan"), boring("Shreeraj"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring, I'm leaving.")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond) // wait for a random time before adding the next number.
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1 // receive from input1 and add into c
		}
	}()
	go func() {
		for {
			c <- <-input2 // receive from input2 and add into c
		}
	}()
	return c
}
