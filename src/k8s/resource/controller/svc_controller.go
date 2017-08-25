package controller

import (
	"connect"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"interact_redis"
	"k8s/resource/info"
)

type SvcController struct {

}

func (*SvcController)GetSvcInfo(namespace string) {
	serviceClient := connect.Clientset.Services(namespace)

	list, err := serviceClient.List(metav1.ListOptions{})

	if err != nil {
		panic(err)
	}

	svcinfo := new(info.SvcInfo)

	for _, d := range list.Items {
			svcinfo.Name = d.GetName()
			svcinfo.Selflink = d.SelfLink
			ir1 := new(interact_redis.InteractRedis)
			ir1.WriteSvcRedis(svcinfo)
		}
}


