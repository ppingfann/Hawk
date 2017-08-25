package interact_redis

import (
	"fmt"
	"connect"
	"k8s/resource/info"
)

type InteractRedis struct {

}

func (* InteractRedis) WritePodRedis(info *info.PodInfo) {
	_, err := connect.ConnRedi.Do("SET",info.Name,info.Ip)

	if err != nil {
		fmt.Printf("ERROR", err)
	}
}

func (* InteractRedis) WriteSvcRedis(info *info.SvcInfo) {
	_, err := connect.ConnRedi.Do("SET",info.Name,info.Selflink)

	if err != nil {
		fmt.Printf("ERROR", err)
	}
}
