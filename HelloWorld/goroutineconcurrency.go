package main

import("time"
		"fmt")

func say(s string){
	for i:=0;i<3;i++{
	fmt.Println(s)
	time.Sleep(time.Millisecond*100)
	}
}
func main() {
	go say("Hey") //go keyword makes a small thread that will not execute if the program ends before thread completes
					// this go thread will be used to implement concurrency
	say("There")
}