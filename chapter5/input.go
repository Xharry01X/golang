package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Enter a number")
	var x float64
	fmt.Scanf("%f",&x)

	output := x*2

	fmt.Println("Now multiplaction is ",output)
}