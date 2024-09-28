package main

import (
	"fmt"
	"time"
)

//we create two go routine to understand the channel reading writing


func main(){
	//1. declare channel of string type i.e. can only transfer data of type string
	ch:=make(chan string)

	//7. call or create go-routine
	go sell(ch)
	go buy(ch)
	//Let us add a timer of two seconds so our main function does not exit.
	time.Sleep(2*time.Second)
	// we can add waitgroup here as well.
}

//2. create two method sell() and buy()

//send data to channel
//channel is just like a slice or map, is sent as a reference implicitly  - so we don't need to use ampersand(&) or the asterisk(*) operator over here as parameter
func sell(ch chan string){
	//3. send data
	ch <- "Furniture"
	//4. print out a statement to the channel
	fmt.Println("Sent data to the  channel")

}
//receive data from the channel
func buy(ch chan string){
//5. 
fmt.Println("Waiting for data")
//6. received data
val := <-ch
fmt.Println("Received data-", val)
}


