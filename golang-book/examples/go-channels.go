package main
import(
	"strconv"
	"fmt"
	"time"
)
var pizzaNum = 0
var pizzaName = ""
func makeDough(stringChan chan string){
	pizzaNum++
	pizzaName= "Pizza #"+strconv.Itoa(pizzaNum)
	fmt.Println("Make Dough and send for Sauce")
	stringChan <- pizzaName
	time.Sleep(time.Millisecond*10)

}
func addSauce(stringChan chan string){
	pizza := <- stringChan
	fmt.Println("Add sauce and send",pizza," for toppings")
	stringChan <- pizzaName
	time.Sleep(time.Millisecond*10)


}
func addTopings(stringChan chan string){
	pizza := <- stringChan
	fmt.Println("Add topings to ",pizza," and ship")
	time.Sleep(time.Millisecond*10)
}