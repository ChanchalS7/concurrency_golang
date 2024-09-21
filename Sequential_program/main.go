package main

import (
	"fmt"
	"time"
)

func calculateSquare(i int) {
	//In this we added time which is going to sleep for one second then print the result, this is going to help us understand about go routines. But more on this later.

	time.Sleep(1*time.Second)
	fmt.Println(i*i)
}
func main() {

	//state of the programm return the current local time when we start program
	start:=time.Now()

	for i := 1; i <= 10000; i++ {
		calculateSquare(i)
	}

	//calculating the time elapsed while starting and calculate the time while executing the function
	elapsed:=time.Since(start)
	fmt.Println("Function took:",elapsed)

	//It will take 10000 second to get all the output
}