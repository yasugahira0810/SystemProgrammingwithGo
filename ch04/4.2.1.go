package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start sub()")
	done := make(chan bool)
	go func() {
		fmt.Println("sub() is running")
		time.Sleep(time.Second)
		fmt.Println("sub() is finished")
		done <- true
	}()
	<-done
	fmt.Println("Hey hey")
	time.Sleep(2 * time.Second)
}
