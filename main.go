package main

import (
	"pod"
	"service"
)

func main() {
	pInfo := new(pod.PodInfo)
	pInfo.GetPodInfo()
	sInfo := &service.Service{}
	sInfo.GetServiceInfo()
}
