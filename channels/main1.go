package main

import "fmt";

func main()  {
	ch :=make(chan int)
	
	fmt.Println("Hello from main")
	
	go multiPlyChannel(ch)
	ch <-10 //assigning ch to 10
	
	fmt.Println("BYE from main")
}

func multiPlyChannel(ch chan int){
fmt.Println(100 * <-ch) //100 * 10
}