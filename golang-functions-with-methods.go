package main

import (
	"fmt"
	"time"
)

type profilableFunc func()

func (f profilableFunc) run() {
	start := time.Now().Unix()
	f()
	end := time.Now().Unix()
	fmt.Println("Ellapsed (seconds):", end-start)
}

func sleep1() {
	fmt.Println("Sleeping for 1 second...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done!")
}

func main() {
	profilable := profilableFunc(sleep1)

	// You can call the custom methods
	profilable.run()

	fmt.Println("")

	// You can use it as a regular function as well
	profilable()
}

// Output:
//
// Sleeping for 1 second...
// Done!
// Ellapsed (seconds): 1
//
// Sleeping for 1 second...
// Done!
