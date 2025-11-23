package examples

import "fmt"

func pointerExample() {
	var age int = 30
	fmt.Println("Age: ", age)

	/* We declare ptr as a type of pointer to an int (*int)
	We get the pointer of age variable by referencing (&) it */

	var ptr *int = &age
	fmt.Println("Pointer to age variable: ", ptr)
	fmt.Println("Value of pointer: ", *ptr)

	/* ptr variable is a pointer to age variable
	By using the asterisk (*) on ptr we dereference it and can set it's value
	As both ptr and age are pointing to the same bit of memory, age is now changed also. */
	*ptr = 35

	fmt.Println("Modified the age", age)

	modifyAge(&age)
	fmt.Println("Age after calling modifyAge:", age)
}

func modifyAge(p *int) {
	*p = 40
}
