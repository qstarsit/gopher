package examples

import "fmt"

// If statements don't need parentheses, but braces are required.
// Variables declared in the condition are scoped to the if/else block.
// https://gobyexample.com/if-else
func ifSample() {
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else if x == 5 {
		fmt.Println("x equals 5")
	} else {
		fmt.Println("x is less than 5")
	}

	// Short declaration in condition; y is only available in the scope of if/else
	if y := x * 2; y > 15 {
		fmt.Println("y is greater than 15:", y)
	} else {
		fmt.Println("y is not greater than 15:", y)
	}

	// y is no longer accessible
}

// For is Go's only looping construct (no while)
// https://gobyexample.com/for
func forSample() {

	// Classic for loop with init; condition; post
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// "While"-like loop: Only a condition
	j := 0
	for j < 3 {
		fmt.Println("j:", j)
		j++
	}

	// Infinite loop
	k := 0
	for {
		k++
		if k > 2 {
			break // exit
		}
		fmt.Println("k:", k)
	}

	// Continue skips to next iteration
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue // skip even numbers
		}
		fmt.Println("odd:", i)
	}
}

// Range iterates over elements in various data structures.
// https://gobyexample.com/range
func rangeSample() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums { // index ignored with _
		sum += num
	}
	fmt.Println("sum:", sum)

	// Range on a slice gives the index and value
	for i, num := range nums {
		fmt.Printf("index %d: value %d\n", i, num)
	}

	// Range on a map gives the key and value
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// Range on a string gives the index and rune (Unicode code point)
	for i, c := range "go" {
		fmt.Printf("index %d: rune %c\n", i, c)
	}
}

// Switch statements compare a value against multiple cases.
// Cases don't fall through by default (unlike C/Java); use fallthrough keyword if needed.
// https://gobyexample.com/switch
func switchSample() {
	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4, 5: // multiple values in one case
		fmt.Println("Thursday or Friday")
	default:
		fmt.Println("Weekend")
	}
}

// Defer schedules a function call to run after the surrounding function completes.
// Multiple defers execute in LIFO order (last in, first out).
// https://gobyexample.com/defer
func deferSample() {
	defer fmt.Println("third")  // executed last
	defer fmt.Println("second") // executed second
	fmt.Println("first")        // executed first
}
