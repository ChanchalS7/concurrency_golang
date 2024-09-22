package main

import (
	"fmt"
	"time"
)

//so in this program we understand how the go-routine execute independent go-routine as there is no child and parent relationship between them
func main(){

	//start method in main()
	go start()
	//we add time to make sure that main  function doesn't exit before all go-routine exited.
	time.Sleep(1*time.Second)
}

//first method start
func start(){

	//called process method inside start()
	go process()
	fmt.Println("In start")

}
//second method process
func process(){
fmt.Println("In process")
}
