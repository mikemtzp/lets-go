package main

import "fmt"

// INITIAL is run once at the beginning of the loop and can create variables within the scope of the loop.
// CONDITION is checked before each iteration. If the condition doesn't pass then the loop breaks.
// AFTER is run after each iteration.

// for INITIAL; CONDITION; AFTER{
// 	do something
// }

// Omit conditions
	// for INITIAL; ; AFTER {
		// do something forever
	// }

// While loops
	// for CONDITION {
		// do some stuff while CONDITION is true
	// }

func loops() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	} // Prints 0 through 9

	// While loops
	plantHeight := 1
	for plantHeight < 5 {
		fmt.Println("still growing! current height:", plantHeight)
		plantHeight++
	}
	fmt.Println("plant has grown to ", plantHeight, "inches")
	// still growing! current height: 1
	// still growing! current height: 2
	// still growing! current height: 3
	// still growing! current height: 4
	// plant has grown to 5 inches

	// The continue keyword stops the current iteration of a loop and continues to the next iteration.
	// continue is a powerful way to use the "guard clause" pattern within loops.
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println(i)
	}
	// 1
	// 3
	// 5
	// 7
	// 9

	// The break keyword stops the current iteration of a loop and exits the loop.
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	// 0
	// 1
	// 2
	// 3
	// 4
}