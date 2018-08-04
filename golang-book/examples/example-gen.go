package main

import "fmt"
import "time"

func main() {
	// go arrays
	array := [5]string{"Banana", "Apple", "Mango", "Grapes", "Watermelon"}
	fmt.Println(array[1])
	for i, value := range array {
		fmt.Println(i, value)
	}
	// go slices
	slices := []string{"Banana", "Apple", "Mango", "Grapes", "Watermelon"}
	for i, value := range slices {
		fmt.Println(i, value)
	}
	// maps

	maps := make(map[string]int)
	maps["Banana"] = 4
	maps["Apple"] = 2
	maps["Mango"] = 23
	fmt.Println(maps["Mango"])
	//defer to execute function at the end of program

	// pointers
	pointer := 0
	pointFunc(&pointer)
	fmt.Printf("Pointer %d \n", pointer)
	// structs and interfaces
	rect := Rectangle{20, 32}
	circ := Circle{12}
	fmt.Println("Area of a Rectangle: ",getArea(rect))
	fmt.Println("Area of a Circle: " ,getArea(circ))
	// go routines
	for i:=0; i<10; i++{
		go count(i)
	}
	time.Sleep(time.Millisecond*1100)
	stringChan := make(chan string)
	for i:=0; i<10; i++{
		go makeDough(stringChan)
		go addSauce(stringChan)
		go addTopings(stringChan)
		time.Sleep(time.Millisecond*1100)

	}
	time.Sleep(time.Millisecond*1100)
}

//function with return types
func add(numberList []float64) float64 {
	sum := 0.0
	for _, val := range numberList {
		sum += val
	}
	return sum
}
func multiArgs(args ...int) int {
	sum := 0
	for _, val := range args {
		sum += val
	}
	return sum
}
func multiReturn(numberList []float64) (largest float64, smallest float64) {
	// your function logic to get smallest and biggest
	return 0.0, 0.0

}
func pointFunc(x *int) {
	*x = 2
}

