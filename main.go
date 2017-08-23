package main

import (
	"k8s/resource/pod"
	"k8s.io/api/core/v1"
	"interact_redis"
)

func main() {
	pc := new(pod.PodCtroller)
	podinfo := pc.GetPodInfo(v1.NamespaceAll,"myweb-1527347769-r7nzw")
	ir := new(interact_redis.InteractRedis)
	ir.WriteRedis(podinfo)
}
