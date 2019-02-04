package main

import (
	"fmt"
	"time"
)
var num int = 100
func say(s string){
	for i := 0; i < num; i++ {
		fmt.Println(s)
		time.Sleep(100 * time.Millisecond)
	}
}

func main(){
	go say("hi")
	say("hello")
}
