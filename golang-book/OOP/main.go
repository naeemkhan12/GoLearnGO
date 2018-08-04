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
