package main

import (
	"fmt"
	"time"
)

func main() {
	//1.create channel
	ch1 := make(chan int)
	//2. create and call go-routine with method
	go sendValues(ch1)
	//3. use select statement
	select {
	//4.received value called
	case msg := <-ch1:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("select time out")
	}

}

// 5. define method
func sendValues(ch1 chan int) {
	//6.add time - make go routint to wait for more than one second
	time.Sleep(3 * time.Second)
	ch1 <- 10
}

// if we exuted right after steps 5, it print 10,
// if we put some sleep time as a step 6 so it prints the select time out.
// 7.we were waiting for the receive operation to complete on the channel.
// Over here is the second select case, we had our time.After method which was waiting for one second.
//Basically, the select statement will wait for at leas one second.
//After that the time.After case statement is going to be executed.
//So, it is an operation that gets unblocked after some time. We can use it to create timeouts using select statements. Where might have to wait for API call or long I/O calles.
