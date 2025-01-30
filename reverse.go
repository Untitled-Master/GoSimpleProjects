package main 

import (
	"fmt"
)

func reverse(a int) int{
	var lastdigit int
	var reverse int 
	for a != 0 {
		lastdigit = a % 10
		reverse = reverse*10 + lastdigit
		a = a/10
	}
	return reverse
}

func main(){
	var a int 
	fmt.Println("Enter a number: ")
	fmt.Scan(&a)
	fmt.Printf("The reverse is %v", reverse(a))
}
