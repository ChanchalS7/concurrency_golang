package main

import "fmt"

func main() {
	//1.create a buffered channel
	ch := make(chan int, 10)

	//2. sent some value from the channel
	ch <- 10
	ch <- 11
	ch <- 12
	//3
	val, ok := <-ch
	//4
	fmt.Println(val,ok)
	//5.
	close(ch)

	val,ok= <-ch
	fmt.Println(val,ok)
	val,ok = <- ch
	fmt.Println(val,ok)

	// so it will print the value along with true, until the value available in channel, even after closing the channel.


}