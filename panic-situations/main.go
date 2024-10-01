package main

import "fmt"

func main(){
	//1. create channel
	ch := make(chan int, 10)
	//2. send value to the channel
	ch <- 10
	ch <- 11
	//3. catch and print value with boolean
	val, ok := <-ch 
	fmt.Println(val,ok)
	//4. close the channel.
	close(ch)
	//5. this will arise the panic situation
	ch<-12
	//6. Because we must always send values before closing the channel, A channel should only be closed if, you are sure that no other values has to be sent to that channel.
	//7. Another panic situation can closing an already closed channel, so if we try to close this channel again it will create panic situations.
	//8. closing again
	close(ch)

}