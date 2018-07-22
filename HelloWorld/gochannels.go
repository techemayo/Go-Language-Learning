package main

import ("fmt"
		"sync")

var wg sync.WaitGroup

func foo(c chan int, someValue int){
	defer wg.Done()
	c <- someValue*5
}

func main(){
	someVal := make(chan int,10)
	for i := 0 ; i <10 ; i++{
		wg.Add(1)
		go foo(someVal,i)
	}

	wg.Wait()
	close(someVal)

	for item := range someVal{
		fmt.Println(item)
	}

	// go foo (someVal,5)
	// go foo (someVal,3)


//	v1 := <- someVal
//	v2 := <- someVal
	// v1,v2 := <- someVal, <- someVal
	// fmt.Println(v1 , v2)
}