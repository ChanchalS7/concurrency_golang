package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 1; i < 10; i++ {
		/*go func() {
			fmt.Println(i)
			wg.Done()
		}()
		*/
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("Done.")
}

/*
- our waitgrup over here,
- passed the parameter 10 to add method.
- because you basicallye want to wait for 10 go-routines to finish their execution.
- the loop is there which spins up 10 go-routines in total.
- so in each iteration, we are creating a go-routine. that particular go-routine takes the value i,
- print it and call the done method on the weight group.
- after all the statement called done and exit the main function.
- now let's run the program and see what happens.

Ideally output should look like : we are iterating from 0 to 10 . So, it should be : 0,1,2,3,4,.... till 10 , order of the output will be non-deterministic.

All go-routines are running, but not printing the desired or expecting result.
- So why doens't each go-routine take value of the loop variable with it. At the same time it was spawned ?
- Why all of them printed 11 value over here or all of them printed some different value than what we exepcted ?
- The point is that these go-routines do not start running immediately.
- We are talking about concurrency over here and not parallelism.
- On a perfectly parallel processor architecutre, the go-routines could indeed start running right away.
In realy life, however, the start usually gets delayed by the go-routine scheduler.
- The fact that there are not always enough CPU cores available, for running all of the go-routines in a separate system thread.
- So what happens is this loop is going to spin up 10 go-routines.
- The go-routines are ready to run,
- But maybe the main go-routine is still claiming the CPU times. So whenever the go-routine starts off, it takes up the value of
i  at that point of time.
- So when all of them printed 11,

*/

//Ideally you should not use the above way or closure to pass data between go-routine, so what is the proper way to do this is to pass the value :
