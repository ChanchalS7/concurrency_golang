package main

import (
	"fmt"
	"sync"
)

// how go-routine leak can occure in program - never terminate and alway occupy the memory it has reserverd, we will see this in program
func main() {

	//1.create wait group
	var wg sync.WaitGroup
	//2. add waitgroup
	wg.Add(2)
	//3. called method as leak and pass wg to it. after 4step go added.
	go leak(&wg)
	//4. called wait method on our waitgroup, that is block the main function

}

// 5. create leak method
func leak(s *sync.WaitGroup) {
	//6. create channel inside this leak method.
	ch := make(chan int)

	//9. create anonymous go-routine
	go func() {
		//10. store value inside val variable
		val := <-ch
		//print the statement
		fmt.Println("Received", val)
		s.Done()

	}()

	//7. print statement
	fmt.Println("Exiting leak method")
	//8done method
	s.Done()
}
