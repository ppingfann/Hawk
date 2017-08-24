package main

import (
	"k8s/resource/pod"
	"k8s.io/api/core/v1"
	"interact_redis"
	"k8s/resource/service"
)

func main() {
	pc := new(pod.PodCtroller)
	podinfo := pc.GetPodInfo(v1.NamespaceAll,"myweb-1527347769-r7nzw")
	ir := new(interact_redis.InteractRedis)
	ir.WritePodRedis(podinfo)

	sc := new(service.SvcController)
	svcinfo := sc.GetSvcInfo(v1.NamespaceAll,"mysql")
	ir1 := new(interact_redis.InteractRedis)
	ir1.WriteSvcRedis(svcinfo)
}
