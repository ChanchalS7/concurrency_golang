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

/*
- we can see that we go error exiting leak method and deadlock.
- we created channel at step 6 and then we created go-routine at step 9.
- and inside the anonymous go-routine, where it is blocking the anonymous go-routine.
- It's waiting for a value to be send on this channel.
- Now while this go-routine is waiting, the leak function returns.
- Now After the leak function return there is no other part of the program that can send a signal over the channel,
- Because it was a local variable that was created inside the leak method.
- This leaves anounymous go-routine over here waiting  indefinitely.
- So make sure whenever you start go-routine it will terminate eventually else it will cause similar memory leaks in your program.


*/
