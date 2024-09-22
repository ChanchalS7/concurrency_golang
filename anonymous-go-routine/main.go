package main

import (
	"fmt"
	"time"
)

func main(){

	go func(){
		fmt.Println("In anonymous method......")
	}()
	time.Sleep(1*time.Second)
}
//so it works as similar to simple go-routine