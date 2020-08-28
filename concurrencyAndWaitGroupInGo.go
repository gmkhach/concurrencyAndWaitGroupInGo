package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// Let's pritn out some facts about our machine
	fmt.Println("OS:\t\t", runtime.GOOS)
	fmt.Println("ARCH:\t\t", runtime.GOARCH)
	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())

	wg.Add(1)
	
	// The go keyword launches the function foo() into its own routine
	go foo()
	bar()
	
	// Printing out the number of Goroutines shows 2 routings now
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())

	// This command will hold the execution of func main at this point untill wg.Done() is called in foo(), indicating that the function execution is complete
	wg.Wait()
}

func foo() {
	for i := 0; i < 100; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 100; i++ {
		fmt.Println("bar:", i)
	}
}
