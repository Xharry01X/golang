package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello, world!")
}
 
 func sayMyName(){
	fmt.Println("My name is Harshit")
 }

func main() {
	go sayHello() // This runs sayHello() concurrently.
	go sayMyName()
	time.Sleep(time.Second) // Give goroutine time to complete before exiting main.
}
