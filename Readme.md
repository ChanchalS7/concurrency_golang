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


### Go runtime scheduler
So, what happens when you start a Go program ?
: The Go program will launch a operating system threads equal to the number of logical cpus available to it. Now, these threads are OS threads, and they are completely managed by the kernel or the operating system. From creating to blocking to scheduling them on CPU course, the entire responsibility is off the operating system.

- We can find out the number of logical processors using the` runtime.Numcpus` method from the runtime package.

- Logical cores = number of physical cores * number of threads that can run on each core(hardware threads)

Let's see how we can calculate the logical processors in our system.
-We know that Go-routines considered as a lightweight application-level thread that has a separate independent execution.
- The go runtime has its own scheduler that will multiplex the go-routines on the os level threads in go runtime.
- It schedules an arbitrary number of go-routines onto an arbitrary number of OS threads which is also called as `m:n multiplexing` .

Let us understand the go-runtime scheduler - 
As we know our OS scheduler manages the OS thread each logical cores in our system within go runtime each of this thread will have one queue associated with it. It is called the LRQ or local run queue.
- It consist of all the go-routine that will be executed in the context of that thread.
- The go-runtime scheduler will be doing the scheduling and the context switching of the go-routines belonging to a particular LRQ.
- Also we have one more queue called the global run queue or the GRQ. It contains all the go-routines that haven't moved to any LRQ or any OS thread.
- The Go scheduler will assign a go-routine form this queue to the local run queue of any operating system thread and well, that was the high level overview of how Go Scheduler works and multiplexes the go-routines on the operating system threads.
[page-22]
#### Cooperative Scheduler
- Golang scheduler is a `cooperative scheduler` it means that there is no time-based preemption that is happening form OS.
- It's a style of scheduling in which the OS never interrupts a running process to initiate a context switch from one process to another.
- In face processes must `voluntarily` yield control periodically or when logically blocked on a resource.
- Of course, there are some specific check points where go-routine can yield its execution to other go-routine. These are called `context-switches`.
For eg. The runtime calls the scheduler on function calls to decide whether a new go-routine needs to be scheduled. 
-So basically when a go-routine make any function call, in that case scheduler will be called and context switch might happen meaning a new go-routine might be scheduled. Well, it's also possible that the existing go-routine continues its execution and no context switching happens.

Some example of context switching:
-Functions call
- Garbage collection
- Network calls
- Channel operations
- On using go keyword
(It depends on the scheduler to do a context switch or not)

#### Go-routines vs Threads
- Go-routines are cheaper, they are only a few KILOBYTES in stack and stack can grow and shrink according to the needs of the application.
- Whereas for a thread, the stack size has to be specified and is fixed. OS threads generally start with ONE MEGABYTE.
Since go-routines are cheap you can launch hundreds and thousands of go-routines while you can only launch a few thousand threads. 

- Go-routines are `multiplexed to a fewer number of OS threads
`. . There might be on threads of program with thousands of go-routines.
- The scheduling of go-routine is done by go runtime and hence it is quite faster. Where as in case of threads, the scheduling of threads done by OS runtime. Hence, the context switching time of go-routines is much master than the context switching time of threads.
- Go-routines communicate using channels.
- Channels, by design prevent race conditions from happening when accessing share memory using go-routines.
- Channels can be thought as a pipe using which go-routines communicate.



### WaitGroups
- In the last or above section , we see the problem that - The problem with go-routines was the `main go-routine terminating` before the go-routines completed or even began their execution.

To wait for multiple go-routines to finish, we have much elegant solution that `timeout` method. And these are the` waitgroups`.

- A wait group is a `synchronized primitive` that allows multiple go-routines to wait for each other.
- This package acts like a counter that block execution in a structured way until its internal counter becomes 0.

[syntax] 👍```
    import "sync"
    var wg sync.WaitGroup




#### Wait groups - Methods
   ` wg.Add(int)`
   - This indicated the number of available go-routines to wait for. The integer in the functions parameters acts like a counter.

  `  wg.Wait()`
   - This methods blocks the execution of code until the internal counter reduces to value=0.

  ` wg.Done()`
   -This method decreases the internal count parameter in the add method by one.

   Let us understand it by example:
  [page33]
  -Assuming we have three go-routines running and we decide to use the WaitGroup. Initially the counter would be zero then we add three go-routines to wait group using the add method. The internal counter goes to three.
  - We also have wait method that will block the execution of code until this internal counter reduces to a zero value.
  - Let's say G1 finishes its execution and calls done function. This would reduces the counter by 1.
  - Now, when G2 finished its execution and calls the done function. The counter again get reduced by one, and similarly for G3.
  Note-all of this would e happening concurrently and not link this sequentially, so when counter goes to zero. The wait function finally unblock before execution.
  Now go to WaitGroup directory for understanding it through example.



  ### Channels 
  - In Go channels are a means through which different go-routines communicate. - It is basically a programming construct, That allows us to move data between different parts of our code often from different go-routines.
  - Along with go-routines channel makes concurrent programming convenient, fun and lowers the difficulty of concurrent programming.

  `Do not communicate by sharing memory; instead share memory by communicating.`
    `-Rob Pike`

 - So the traditional threading models which are commonly used when we write, Java or C++, requires the programmer to communicate between threads using shared memory.

 - Communicate by sharing memory - `Threads and mutexes`.
 Typically, shared data structures are protected by locks and threads contained over those logs to access the data. Hence communicating with each other by sharing memory in form of data.
 - While Go's concurrency primitives which go-routines and channels provide us and elegant and distinct means of structuring concurrent software. Instead of explicitly using locks to mediate access to shared data.
 - Go encourages the use of `channels` to pass reference or memory to data between go-routines.
 - Share memory by communicating - `Go-routines and channels`
 -This approach ensues only one Go-routines has access to the data at a given time. Hence, sharing memory by communicating.

 ##### Channels
 - The communication in channel is `bidirectional by default`, meaning that you can `send and receive` values from the `same channel`.
 - By default, channels send and receive until the other side is ready.
 - This allows go-routines to synchronize without explicit locks or condition variable.

 - Now each channel can hold a data only  of a particular data type. Let's say it could only string or integer type and also go uses special keyword which declaring channel called `chan`
  [Syntax]
  `var c chan string`
  
  -Alternatively we can use the make function to declare and initialize the channel.
`  c:= make(chan string)`
- Channel type can be bi-directional or single-directional but for our course, we would be limiting the discussion to the bi-directional channel.
-Now, a channel is meant for communicating between go-routines.
- Hence, it has many operations such as [channel-operations]
        - Sending a value
        - Receiving a value
        - Closing a channel
        - Querying Buffer of a channel
        - Querying length of a channel

   In Details understand the above operations.     

   ``Channel operations: Sending a value``
   `ch <- v` 
   - `<-` is a channel send operator
   - The operator is used to send a value to the channel.
   - v must be a value which assignable to the element type of channel `ch`.


   ``Channel operations: Receiving a value``
   `val:=<-ch`
   - `<-` is a channel receiving operator.
   - This is used to receive a value from a channel.
   - val is variable in which the read data from the channel will be stored.


    ``Channel operations: Closing a channel``
    `close(ch)`
    -`close()` is a built-in function.
    - The argument of a close function call must be a channel value.

    ``Channel operations: Querying buffer of a channel``
    `cap(ch)`
    -`cap()` is a built-in function.
    - returns an integer denoting the buffer of the specified channel.


    ``Channel operations: Querying length of a channel``
    `len(ch)`
    -`len()` is a built-in function.
    - returns an integer denoting the length of the specified channel.