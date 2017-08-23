package interact_redis

import (
	"fmt"
	"k8s/resource/pod"
	"connect"
)

type InteractRedis struct {

}
func (* InteractRedis) WriteRedis(info *pod.PodInfo) {
	_, err := connect.ConnRedi.Do("SET",info.Name,info.Ip)

	if err != nil {
		fmt.Printf("ERROR", err)
	}
}