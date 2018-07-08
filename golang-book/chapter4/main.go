 package main
 import "fmt"
 func main() {
	 /*Go variables*/
	 // with type declaration
	 var x string;
	 var y string="John";
	 //without type declaration
	 z :="Glad to See ya !!";
	 x="Hello ";
	 // constants
	 const constant string="I am a constant"
	 // declaring multiple variables
	 var (
		 a=12
		 b=23
		 c="Stringing"
	 )
	 fmt.Println(x,y,z);
	 fmt.Println(constant);
	 fmt.Println(a,b,c)

 }
