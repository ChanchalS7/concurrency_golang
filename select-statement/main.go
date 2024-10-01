package main

import (
	"fmt"
	"time"
)

func main() {
	//1.create two channels

	ch1 := make(chan string)
	ch2 := make(chan string)

	//2. create two go-routine.
	go goOne(ch1)
	go goTwo(ch2)

	//5. let's use select statement
	select {
		//receiving statement from channel 1
	case val1 := <-ch1:
		fmt.Println(val1)
		//receving value from channel 2
	case val2 := <-ch2:
		fmt.Println(val2)	
	}
	//6. we have two go-routines so we want main function to be wait for execution both, we are going to add either timer or waitGroup
	time.Sleep(1*time.Second)

}

// 3.
func goTwo(ch2 chan string) {
	ch2 <- "Channel-2"
}

// 4.
func goOne(ch1 chan string) {
	ch1 <- "Channel-1"
}

//7. run the program - 
//8. We see that second case statement was executed in this case, hence we got the value channel2 because it was the value that we sent on channel 2.
// Now, over here the output cannot be determined. Because either of the go-routine, goOne or goTwo might execute first. Now, let's say even if they execute at the same time. The select statement might choose them randomly. So output over here is non-deterministic.
