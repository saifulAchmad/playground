package main

import "fmt"

//fungsi square sama seperti sebelumnya
//yang membedakan adalah fungsi ini
//menggunakan blank return
func main() {
	result1, result2 := square(4, 5)
	fmt.Printf("%d dan %d\n", result1, result2)
	fmt.Println(square(9, 8))
}

// TODO: answer here
func square (n,n1 int)(result1,result2 int){
	result1 = n*n
	result2 = n1*n1
   return 
}
