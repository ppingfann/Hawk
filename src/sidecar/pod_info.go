package sidecar

import(
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//acquire pod information
func Pod_info(){
	podsClient := Clientset.Pods(apiv1.NamespaceAll)

	list, err := podsClient.List(metav1.ListOptions{
	})

	if err != nil {
		panic(err)
	}

	for _, d := range list.Items {
		_, err_1 := Conn_redi.Do("SET", d.GetName(), d.Status.PodIP)

		if err_1 != nil {
			fmt.Printf("ERROR", err_1)
		}

	}
}
