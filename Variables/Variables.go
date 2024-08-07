package Variables

import (
	"fmt"
)

func main() {
	var a = "initial" //var keyword declaration creates a variable of a particular type
	fmt.Println(a)

	var b, c int = 1, 2 //we can declare and initialize a set of variables in a single declaration
	fmt.Println(b, c)

	var boolVal, floatVal, stringVal = true, 2.3, "four" //here we can omit the type and declare and initialize multiple variables of different types
	fmt.Println(boolVal, floatVal, stringVal)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple" //if the type is omitted it becomes determined by initializer expression
	fmt.Println(f)

	//if the expression is omitted, the initial value  os zero value for the type, which 0 is for numbers, false for booleans, "" for strings, and nil interfaces and reference types(slice, pointer, map, channel, function)
	// the zero value for an aggregate type like an array or a struct has the zero value of all its elements or fields

	//var f, err = os.Open(name) <- a set of variables can be initialized by calling a function that returns multiple values

	//form name := expression <- the type of name is deteremined by the type of expression

	pi := 3.14 //short var expression
	fmt.Println(pi)

	//a var declaration tends to be reserved for local variables that need an explicit type that differs from that of the initializer expression, or for when
	//the variable will be assigned a value letter and its initial value is unimportant

	number1, number2 := 0, 1 // we can also short declare and initialize multiple variables
	fmt.Println(number1, number2)

}
