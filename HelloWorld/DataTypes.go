package main 

import "fmt"

// func add(x float64, y float64) float64{
// 	return x+y
// }

// func add(x,y float64)float64{
// 	return x+y
// }

func add(x,y float32)float32{
	return x+y
}

func multiple (a,b string)(string,string){
	return a,b
}

// const x int = 5
func main(){
	// var num1 float64 = 5.6
	// var num2 float64 = 9.5

	// var num1,num2 float64 = 5.6,9.5
	// num1,num2 := 5.6,9.5	//float64 type would be assigned by default
	/////if we declare a variable and dont use it program wouldn't execute
	/////and will through an error variable not used
	// fmt.Println("The sum of ", num1 , " and ", num2, "is " , add(num1,num2))
	/*this will give error beacause arguments require float 32 but num1 and num2 
	are assigned float 64 by default*/
	// w1,w2 := "Hey","There"
	// fmt.Println(multiple(w1,w2))

	var a int = 62
	var b float64 = float64(a) //type casting
	fmt.Println(b)
	x:=a //x will be type int
	fmt.Println(x)

}