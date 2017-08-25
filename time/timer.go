package main

import (
	//"fmt"
	//"time"
	"fmt"
)

func main() {
	ch := make(chan int)
for i:=0 ; i<10 ; i++{
	select {
	case ch <- 1:
	case ch <- 0:
	}
	go print(ch)
	}

}

func print(ch chan int) {
	for i := range ch{
	fmt.Println(i)
	}
}
