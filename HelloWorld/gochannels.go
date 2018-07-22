package main

import "fmt"

func foo(c chan int, someValue int){
	c <- someValue*5
}

func main(){
	someVal := make(chan int)
	go foo (someVal,5)
	go foo (someVal,3)

//	v1 := <- someVal
//	v2 := <- someVal
	v1,v2 := <- someVal, <- someVal
	fmt.Println(v1 , v2)
}