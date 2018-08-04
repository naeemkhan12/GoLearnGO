package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	cat := Cat{Animal{"cat", 4}, "MEW !!"}
	dog := Dog{Animal{"Dog", 9}, "BARK !!"}
	cat.MakeNoise(cat.voice)
	dog.MakeNoise(dog.voice)
	vehicle := new(Vehicle)
	vehicle.List[0] = Car{"BEEP"}
	vehicle.List[1] = Bike{"RING"}
	for i := range vehicle.List {
		PressHorn(vehicle.List[i])
	}

}

type Animal struct {
	Name     string
	Strength int
}
type Cat struct {
	Animal
	voice string
}
type Dog struct {
	Animal
	voice string
}

func (a Animal) MakeNoise(voice string) {
	for i := 0; i < a.Strength; i++ {
		fmt.Println(a.Name, " says ", voice)
	}
}
