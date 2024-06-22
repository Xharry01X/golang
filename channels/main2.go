package main


import "fmt";

func main()  {
	ch :=make(chan string)

	fmt.Println("Working on string")

	go addingMoreString(ch)

	ch <- "Harshit"
	fmt.Println("Go lang channels are very good")
}

func addingMoreString(ch chan string){
	fmt.Println("My name is", <- ch)
}