package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func(){
		c <- 42
		close(c)
	}()
	
	fmt.Println(<-c)
	
	c1 := make(chan int,1)
	c1 <- 32
	fmt.Println(<-c1)

}