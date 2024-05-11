package main

import(
	"fmt"
)

func main()  {

	var num int
	fmt.Println("ENter 10 numbers")
   for i := 0; i < 10; i++ {
	fmt.Printf("Number %d:",i+1)
	fmt.Scanln(&num)
	fmt.Println("You entered",num)
   }
}