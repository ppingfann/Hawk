package main

import (
	"k8s.io/api/core/v1"
	"k8s/resource/controller"
	"time"
	"fmt"
)

func main() {
	ch := make(chan int)
	go execute(ch)
	for i:=0; i<10;i++  {
		fmt.Println("hello")
	}
	select {
	case <- ch:

	}

}

func execute(ch chan int)  {
	for  {
			pc := new(controller.PodCtroller)
			sc := new(controller.SvcController)
			pc.GetPodInfo(v1.NamespaceAll)
			sc.GetSvcInfo(v1.NamespaceAll)

			time.Sleep(5 * time.Second)
		}
	ch <- 1
}
