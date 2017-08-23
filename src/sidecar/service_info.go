package sidecar

import(
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiv1 "k8s.io/api/core/v1"

)

func Service_info() {

	serviceClient := Clientset.Services(apiv1.NamespaceAll)

	list_2, err := serviceClient.List(metav1.ListOptions{})

	if err != nil {
		panic(err)
	}

	for _, d := range list_2.Items {
		_, err_1 := Conn_redi.Do("SET", d.GetName(), d.SelfLink)

		if err_1 != nil {
			fmt.Printf("ERROR", err_1)
		}

	}
}
