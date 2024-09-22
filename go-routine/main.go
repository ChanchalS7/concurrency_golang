package main
import (
	"fmt"
	"time"
)
func calculateSquare(i int){
	//In this we added time which is going to sleep for one second then print the result, this is going to help us understand about go routines. But more on this later.
	time.Sleep(1*time.Second)
	fmt.Println(i*i)
}

func main(){

	//start of the programm return the current local time when we start program
	start:=time.Now()

	for i := 1; i <= 10000; i++ {
		go calculateSquare(i)
	}

	//calculating the time elapsed while starting and calculate the time while executing the function
	elapsed:=time.Since(start)
	time.Sleep(2*time.Second)
	fmt.Println("Function took:",elapsed)
	//its just executed but did not see any of the square being print over the screen and this is because the main function exited over here without even waiting for all the go routines to be finished.

	//Ideally over here main function should wait for some time before exited so all the go routines can be execute so for this here be add timer in main function which will wait for one or two second before exiting the main function. So I added timer in line number -23

	//fist run without line number 23 and then add for better understanding.
	//Now more better and reliable way to execute the go-routine is waitGroup which we will discussing about it in next program
}