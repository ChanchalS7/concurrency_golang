package main

import (
	"fmt"
	"sync"
)

//buffered channels
func main(){
	//2. create waitgroup
	var wg sync.WaitGroup
	//3.add waitgroup of counter 2
	wg.Add(2)
//1.create the channels
ch := make(chan int, 3)

//5. create a method that is going to send the values to our channel
go sell(ch,&wg) //17. going to call sell in go-routine as well.

//4.call wg function here until all go routine return
wg.Wait()

}
//6. create method sell
func sell(ch chan int, s *sync.WaitGroup) {
	//7.we are going to send some value to the channels as buffer size is size 3 so we are going to send three values
	ch <- 10
	ch <- 11
	ch <- 12
	//22. go buy(ch,s)
	//19. exceed buffer limit
	ch <-13
	
	//10. Let's create one more method buy that would receive values from the channel
	go buy(ch, s) //16. Now we want buy method to be executed in another go routine so we are going to make it go routine here
	//8. Print statement at the end of the method.
	fmt.Println("Sent all data to the channel")
	//9. Since the process is done we are going to call the done method on our WaitGroup
	s.Done()

}

//11. In this method we are going to receive values from the channel
func buy(ch chan int, s *sync.WaitGroup){
	//12.
	fmt.Println("Waiting for data")
	//13.received value, print
	fmt.Println("Received data", <-ch)
	//14.
	s.Done()

}
//15. Now we want buy method to be executed in another go-routine.

//18.Over here the buy method received teh value 1- and it got printed the output as well. You also noticed here unlike unbuffered channel , a buffered channel over here did'nt block the sell go-routine, this is also because we did'nt exceed the buffer limit. Let's try to send one more value to the channel.

//20. Now it prints the output like : fatal error: all goroutines are asleep - deadlock! . It is because the buy method on line number 34, inside the sell method blocking sell method and even the buy method could not be called over here. Hence it's creating a deadlock situation.
//21. So, for buffered channel the send operation is going to block the go-routine, if it exceeds the buffer limit. Now if call this go-routine before the blocking line which is the fourth channel, so before it. Its not going to give us any error. Because the go-routine that's going to receive the value from the channel has been called before this blocking line, which was earlier blocking our sell method earlier.
//23. Now, talking about a concept that a receive on a channel is blocked, when the channel is empty.
//24. So let's make some change here in the in the above code and not send anything, from this channel.
//25. We again see the error that all go-routines are asleep and there a deadlock situation. Over here, since the channel is empty, the  buy method is going to be blocked. To resolve this error we can simply send some data to the channel and we'll see that there's no error at all. So, if I send some value , we can see there is no blocking over the channel and buy method receive the value one from the channel itself.