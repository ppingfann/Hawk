package pod

import (
	"connect"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodCtroller struct {

}

func (* PodCtroller)GetPodInfo(namespace string,name string) *PodInfo {
	podsClient := connect.Clientset.Pods(namespace)

	list, err := podsClient.List(metav1.ListOptions{})

	if err != nil {
		panic(err)
	}

	p := new(PodInfo)

	for _, d := range list.Items {
		if name == d.GetName(){
			p.Name = d.GetName()
			p.Ip = d.Status.PodIP

		}
	}
	return p
}
