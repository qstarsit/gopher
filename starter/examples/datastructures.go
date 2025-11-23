package examples

// Arrays are fixed length collections of elements of a single type.
// https://gobyexample.com/arrays
func arraySample() {
	var a [3]int                        // zero‑initialized: a = [0 0 0]
	b := [...]string{"go", "is", "fun"} // length inferred (3)
	c := a                              // c gets a copy of a, they are now independent objects
	c[0] = 42                           // does not change a
	_ = b
	_ = c
	_ = a // avoid unused warnings
}

// A slice is a lightweight descriptor (pointer, length, capacity) referencing an underlying array.
// Slices can have dynamic size, slicing does not copy elements, instead both slices share the same memory.
// https://gobyexample.com/slices
func sliceSample() {
	base := []int{1, 2, 3, 4, 5} // The underlying Array
	sub := base[1:3]             // elements 2,3; len=2 cap depends on base capacity from index 1
	sub[0] = 99                  // changes base[1]

	// Detach via copy (new "backing" array):
	detached := append([]int(nil), sub...)
	sub[1] = 77 // does not affect detached now
	_ = detached
	_ = sub
	_ = base
}

// Maps are hash tables: dynamic size, reference type
// Lookups return value and a boolean (presence). Deleting a missing key is safe.
func mapSample() {
	scores := map[string]int{"alice": 3, "bob": 5}
	scores["carol"] = 7     // insert
	scores["alice"]++       // update existing
	v, ok := scores["dave"] // ok is false; v is zero (0)
	delete(scores, "bob")   // remove key
	// Iteration order is deliberately randomized; do not rely on ordering.
	_ = v
	_ = ok
	_ = scores
}

// https://gobyexample.com/structs
type User struct {
	Name string
	Age  int
	Tags []string
}

func structSample() {
	var u User // zero value: Name="" Age=0 Tags=nil
	admin := User{Name: "Admin", Age: 42, Tags: []string{"root", "sys"}}

	p := &admin                       // reference p to admin for mutation of the original object
	p.Tags = append(p.Tags, "active") // admin.Tags is now ["root", "sys", "active"]
	p.Age++                           // admin.Age is now 43

	/* The profile struct embeds User, meaning it includes all fields of User directly.
	This is the "Composition over Inheritance" principle in Go.
	https://www.digitalocean.com/community/tutorials/composition-vs-inheritance
	*/
	type Profile struct {
		User
		Bio string
	}
	prof := Profile{User: User{Name: "Carol", Age: 29}, Bio: "Gopher"}

	_ = u
	_ = admin
	_ = prof
}
