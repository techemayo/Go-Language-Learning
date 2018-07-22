package main

import("time"
		"fmt"
		"sync")

var wg sync.WaitGroup // Wait group allows us to group go threads to wait for
func say(s string){
	for i:=0;i<3;i++{
	fmt.Println(s)
	time.Sleep(time.Millisecond*100)
	}
	wg.Done() //Done() tell the object that wait group has done executed
}
func main() {
	wg.Add(1) // add will make the thread add in wait group of sync
	go say("Hey") //go keyword makes a small thread that will not execute if the program ends before thread completes
	wg.Add(1)				// this go thread will be used to implement concurrency
	go say("There") //if we concurrent everything with go thread nothing will block program from ending and nothing will be completed
	//time.Sleep(time.Second) //1 second sleep to provide time fot go thrads to execute
	wg.Wait() //wait() will wait to execute untill sync is done 
	wg.Add(1)
	go say("Hi")
	wg.Wait()


}