package main

import "fmt"

							// return value ???????
func add(x float64,y float64)float64{
	return x+y
}
// short hand
func multiple(a,b string) (string,string){
	return a,b
}
func main(){
	// short hand
	var num1,num2 float64 = 5.6, 9.5
	
	fmt.Println(add(num1,num2))
	// short hand 
	w1, w2 := "Hey","There"

	fmt.Println(multiple(w1,w2))

	/*
	multiple line comments 

	*/
}