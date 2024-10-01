package main

import (
	"fmt"
	"sync"
)

//1. Simple demonstration how range in for loop works.
/*
func main(){

	//for range
	nums := [5]int{1,2,3,4,5}

	for i, item:=range nums {
		fmt.Println(i,"-",item)
	}
}
*/

// We can also use to iterate over values received from the channel.

func main(){
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int)
	go sell(ch, &wg)
	go buy(ch, &wg)
	wg.Wait()
}
//sending value
func sell(ch chan int,  wg *sync.WaitGroup){
	ch <-1 
	ch <- 2
	fmt.Println("Sent all data")
	close(ch)
	wg.Done()
}
//receiving value
func buy(ch chan int, wg *sync.WaitGroup){
	fmt.Println("Waiting for data")
	// add for range for receiving value
	for val := range ch {
		fmt.Println("Received:",val)
	}
	wg.Done()
}

//so the range keeps on iterating over each element as it is received from the queue, now what happens if we don't close the channel.
// so just comment out the line 36 and run program again.
// There is some error from the WaitGroup, which is fatal-error, all go-routines are asleep, and it's a deadlock. This happens because the for-range is never going to finish until the channel is closed. And hence it created a deadlock situation for us.
// so whenever you are using for-range while iterating over the channel for receiving value, make sure to close the channel.
// Now for unbuffered channel its almost the same.
//Now let's see how for-range works in unbuffered channel.

/*
func main(){
	ch := make(chan int, 5)
	ch <-100
	ch <- 200
	 close(ch)

	 for val := range ch {
		fmt.Println(val)
	 }
}
	 */