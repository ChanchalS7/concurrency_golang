package main

import (
	"fmt"
	"sync"
	"time"
)
//wg added later in calculateSquare method later as a parameter

func calculateSquare(i int, wg *sync.WaitGroup) {
fmt.Println(i*i)
//in this method after done with processing let's call the done method
wg.Done() // it will be called on each of the go routine and , will keep decreasing the counter by one for each go-routine.
}
func main() {
	//now let's create waitgroups
	var wg sync.WaitGroup
	start := time.Now()
	wg.Add(10)
	for i := 0; i < 10; i++ {
		//go routine
		go calculateSquare(i, &wg)
	}
	//elapsed time
	elapsed := time.Since(start)
	//now we also want to block the execution of our main go-routine until all our go-routine execute hence we use wg.wait() method.
	wg.Wait()
	fmt.Println("Function took", elapsed)

	//it prints the output in not specific order because it is non-deterministic way
	//Because it can't determine which go-routine go in what order ?
	

}