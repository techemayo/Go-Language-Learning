package main

import "fmt"

func main(){

	x:=15 //x = 15 with type int
	fmt.Println(x) // prints the value of x which is 15
	a:=&x ////assign the memmory address of x 
		 ////to a memmory address are hexadecimals
	fmt.Println(a) //print the value of a which is the memmory address of x
	fmt.Println(*a) ////print the value stored on the 
					//// memmory address stored in the variable a
	*a = 5 //assigns the value of 5 to the memmory address stored in variable a
	fmt.Println(x) ////prints the value assigned to the 
				   ////memmory address store in variable a,
				   ////"a" is a pointer of variable x
	*a = *a**a  //// assign the value to the value of memmory address stored in variable a the value of the value stored at the memmory address of a times the value of the value stored at the memmory address   
	fmt.Println(x) //prints the value of x
	fmt.Println(*a) // print the value store at memmory address store in a
}