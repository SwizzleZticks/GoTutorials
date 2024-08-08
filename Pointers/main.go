package main

import "fmt"

func main() {
	var age int
	age = 42
	fmt.Println(age)

	var pointer *int      //this is a variable that is pointing to an int
	fmt.Println(&age)     //this prints the memory address of the variable age
	pointer = &age        //here we are assigning ages memory address to pointer, which is of type (pointer to an int)
	fmt.Println(*pointer) //when you place thr "*" in front of a variable name now you are dereferencing a pointer
	*pointer = 21         //this will change the value of age from 42 to 21
	fmt.Println(age)      //this is 21

	//since pointer is using ages memory address, altering the value of pointer variable will also alter ages variable since they share the same address

	var x, y int                               //the zero value for a pointer is nil, pointers are comparable
	fmt.Println(&x == &x, &x == &y, &x == nil) // true, false, false

	v := 1
	incr(&v)              //side effect: v is now 2
	fmt.Println(incr(&v)) // "3" (and v is 3)

}

// Because a pointer contains the address of a variable, passing a pointer argument to a function
// makes it possible for the function to update the variable that was indirectly passed
func incr(p *int) int {
	*p++ //increments what p points to, does not change p
	return *p
}

// A pointer value is the address of a variable
//
//With a pointer we can read or update the value of a variable indirectly
//without using or even knowing the name of the variable, if it does indeed have a name

//if a variable is declared var x int the expression &x yields a pointer to an int variable
