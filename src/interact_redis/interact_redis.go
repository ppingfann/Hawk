package interact_redis

import (
	"fmt"
	"k8s/resource/pod"
	"connect"
	"k8s/resource/service"
)

type InteractRedis struct {

}

func (* InteractRedis) WritePodRedis(info *pod.PodInfo) {
	_, err := connect.ConnRedi.Do("SET",info.Name,info.Ip)

	if err != nil {
		fmt.Printf("ERROR", err)
	}
}

func (* InteractRedis) WriteSvcRedis(info *service.SvcInfo) {
	_, err := connect.ConnRedi.Do("SET",info.Name,info.Selflink)

	if err != nil {
		fmt.Printf("ERROR", err)
	}
}
