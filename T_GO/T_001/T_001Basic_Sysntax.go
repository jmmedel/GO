package main
// i get error if i use math/rand  very bad
import "fmt"
import "math/rand"
import "math"

// you can have the same main another file 
// you get an error 

func foo (){

	fmt.Println("The Square root of 4 is ",math.Sqrt(4))
}

func main(){
	foo()
	fmt.Println("A nmber from 1 to 100 ", rand.Intn(100))
}



