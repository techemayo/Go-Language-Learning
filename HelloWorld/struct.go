package main // main package is being used

import "fmt"
//fmt library imported

type car struct { //structure car an structure is a composite datatype that is used to store the data that have different datatype than traditional datatypes usaully stores during handling of composite data
	gas_pedal uint16 //unsigned integer 16 can store values from 0 to 	65535 	
	break_pedal uint16	
	steering_wheel int16 //int16 can store values from -32k to +32k
	top_speed_kmh float64
}

func main() {
	a_car := car{gas_pedal:22431,
				 break_pedal:0,
				 steering_wheel:12561,
				 top_speed_kmh:255.0}

//	a_car := car{22431,0,12561,255.0} //hardcoded or shorthand way it is a legit way to cereate new car from car structure but the chaces of misshandling are higher i.e: chances of human error

	fmt.Println(a_car.gas_pedal)
}