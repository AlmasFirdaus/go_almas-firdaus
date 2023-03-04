package main

import "fmt"

type Car struct {
	fuel float64
}

func (car *Car) jarak() {
	distance := car.fuel / 1.5
	fmt.Println("Bensin mobil terisi sebanyak", car.fuel,"L, perkiraan jarak yang dapat ditempuh yaitu:", distance,"mil")
}

func main(){
	var fuelCharge float64
	fmt.Print("isi bahan bakar : ")
	fmt.Scan(&fuelCharge)
	car1 := Car{
		fuel: fuelCharge,
	}
	car1.jarak()
}