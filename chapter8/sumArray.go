package main

import "fmt"

func main()  {

var n int
fmt.Print("Enter the size")

fmt.Scan(&n)

arr :=make([]int,n)

fmt.Print("Enter the number")
for i := 0; i < n; i++ {
	fmt.Scan(&arr[i])
}


  total :=0

   for i := 0; i < n; i++ {
	total +=arr[i]
   }
   fmt.Print(total)
}
