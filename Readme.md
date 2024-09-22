## Sequential vs Cocurrency -
        Sequential programming commands or instructions are executed one after another it means next commands or instructions can not be executed until the previous command or instructions are not completely executed.
 ## page-3
        For Eg: Calculate the sum of two number,
        Instructions are:
         - ask first number
          -second number -
           calculate sum 
           - return sum
 ## page-4          
      Sequential programming are  often used for simple program where commands or instruction are not  dependent on each other.


But CPU is super fast, so keeping occupied it for single task is not a good idea. So here Multitask come in the picture, it is the ability of the operating system to execute more than one task simultaneously in cpu. 
It happens in cpu by using switching jobs using small time quantum.
When we talk about multitasking we generally refer to single core processor but it switches one program to another so quickly so it gives us the appearance of executing all of the program at the same time.
 ## page-5
### Now then what the concurrency is 
 - Concurrency is the notion of multiple things happening at the same time. It is the potential for multiple processes to be in progress at the same time. We can also say concurrency can achieve through multitasking.

 - Now concurrency and multi core cpu would look something like this there is two core and two task each core is doing both task by switching among them over time. This system is also knows as multi processing system. The CPU core are added for increase computing power.

 ## page-7

 However we should not mistake  concurrency with parallelism : 

 - Parallelism refers to using multiple processing elements simultaneously for solving any problem. For eg:

 ## page-8

Simple examples of concurrent processing can be any interactive program such as
- use Interactive program (Like Text Editor in file typing and saving at same time, in this case this looks like all operation look like they are happening simultaneously at the same time but that's not case )
In contrast for Parallel processing
- distributed data processing ( it involved large scale data processing and it uses parallel processor but programmer see the entire system as a single database)


## Go - routines 
        It is the core concept in go concurrency model.
        Go routine can be thought of as a light weight thread that has a separate independent execution.
        Can execute concurrently with other go-routines.
        All go-routines are managed concurrently by Go runtime schedulers.


        Syntax : to start a go routine add special keyword called go before a function or method call.
        ```
            go calculate() 
            go func(){}()
        ```
    - That function or method now will be executed within go routine. I think we need to note one point here, it is not the function or method which determine the go routine if we call that function or method with go keyword then that function or method said to be executing in go-routine .
    - Over here we have a method calculate() which is executing separate go routine.

#### Now in the directory called  for leveraging sequential programming you can try out for better understanding of go-routine the example before going down -----------

#### So now let see how we can do the same thing using go-routine so lets go to the go-routine directory where you can find proper explanation of it.

#### Lets understand about waitGroup which is more reliable way to execute go-routine.

### Main go-routine
    - Main function in the main package is the main go-routine.
    - All go-routines are started form the main go-routine, these go-routine can start multiple other go-routine and so on. 
    - But the Main go-routine represent the main program. Once it exit that means entire program exits.
    - Go-routine do not have parents and children, when we starts a go-routine, it just executes alongside all other go-routine. Each go-routine exits only when its function returns. 
    The only exception to that is that all go-routine exits when the main go-routine, the one that runs  the main function exit

#### So let's understand the main go routine in the directory called main-go-routine 
So here in program you find this:
-So sometime `In Process` printed later on and some time first.
- So basically output is non deterministic.
-So basically when `In Process` printed later on shows go-routine does not have parent-child relationship and they exist as a independent execution. -Because `In Process` was printed after the start method printed its statement and exited. This means process and start executing completely independent of each other, this also the reason why our output is not deterministic.


### Anonymous go-routine
- In Golang, anonymous function are those functions that don't have any name. Simply put, anonymous functions don't use any variables as a name when they are declared.
- Anonymous functions in golang can also be called using go-routine.
Syntax:
```
go func(){


}(args...)
```

So basically no difference in behavior when calling simple go-routine or anonymous method.

