package controller

import (
	"connect"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"interact_redis"
	"k8s/resource/info"
)

type PodCtroller struct {

}

func (* PodCtroller)GetPodInfo(namespace string) {
	podsClient := connect.Clientset.Pods(namespace)

	list, err := podsClient.List(metav1.ListOptions{})

	if err != nil {
		panic(err)
	}

	p := new(info.PodInfo)

	for _, d := range list.Items {
			p.Name = d.GetName()
			p.Ip = d.Status.PodIP
			ir := new(interact_redis.InteractRedis)
			ir.WritePodRedis(p)
	}
}
