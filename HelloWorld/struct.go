package main // main package is being used

import "fmt"
//fmt library imported


const usinxteenbitmax float64 = 65535	
const kmh_multiple float64 = 1.60934 //1 mile is equals to 1.60934 miles

//we have two types of 	methods 1. Value Reciever : Methods that just perform clculation on values
							////2. Pointer Recievers : Methods that change or modify a value or something in a struct or somewhere else 


type car struct { //structure car an structure is a composite datatype that is used to store the data that have different datatype than traditional datatypes usaully stores during handling of composite data
	gas_pedal uint16 //unsigned integer 16 can store values from 0 to 	65535 	
	break_pedal uint16	
	steering_wheel int16 //int16 can store values from -32k to +32k
	top_speed_kmh float64
}

func (c *car) new_top_speed(newspeed float64){ //pointer function here it is used to set new speed in car top speed
	c.top_speed_kmh = newspeed 					//*car pointer makes it pointer function	
}												//permanently updates top speed for the structure

func (c car) kmh() float64{  //value receiver function/method
	// (c car) associate car type to this function and c is used to refrence car struc 
	 //this function calculate the top speed in kilometer per hour
	//c.top_speed_kmh = 500 //temporarily set top speed to 500 and do not store in structure
	return float64(c.gas_pedal)*(c.top_speed_kmh / usinxteenbitmax)
}
// func (c *car) kmh() float64{  //pointer function
// 	c.top_speed_kmh = 500 //permanently set top speed to 500 and do not store in structure
// 	return float64(c.gas_pedal)*(c.top_speed_kmh / usinxteenbitmax)
// }

func (c *car) mph() float64 { //value reciever function
	c.top_speed_kmh=500 //permanantly ipdate value of top speed for the car
	return float64(c.gas_pedal)*(c.top_speed_kmh/usinxteenbitmax/kmh_multiple)
}

//we use pointer function where we need performace gains
//value receiver functions are expensive

func newer_car_speed(c car ,speed float64) car{
	c.top_speed_kmh=speed
	return c
}

func main() {
	a_car := car{//gas_pedal:22431,
				 gas_pedal:65000,
				 break_pedal:0,
				 steering_wheel:12561,
				 top_speed_kmh:255.0}

//	a_car := car{22431,0,12561,255.0} //hardcoded or shorthand way it is a legit way to cereate new car from car structure but the chaces of misshandling are higher i.e: chances of human error

	//fmt.Println(a_car.gas_pedal)

	fmt.Println( a_car.kmh() )
	fmt.Println( a_car.mph() )
	// a_car.new_top_speed(500)
	a_car = newer_car_speed(a_car, 1100)
	fmt.Println( a_car.kmh() )
	fmt.Println( a_car.mph() )

}
