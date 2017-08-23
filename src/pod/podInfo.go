package pod

import(
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"connect"
)

//acquire pod information
type PodInfo struct{
	name string
	ip string
}

func (p *PodInfo) GetPodInfo() {
	podsClient := connect.Clientset.Pods(apiv1.NamespaceAll)

	list, err := podsClient.List(metav1.ListOptions{})

	if err != nil {
		panic(err)
	}

	for _, d := range list.Items {

		p.name = d.GetName()
		p.ip = d.Status.PodIP

		_, err := connect.ConnRedi.Do("SET",p.name,p.ip)

		if err != nil {
			fmt.Printf("ERROR", err)
		}

	}
}
