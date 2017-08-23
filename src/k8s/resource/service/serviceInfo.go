package service

import(
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiv1 "k8s.io/api/core/v1"
	"connect"

)

//get Service information

type Service struct {
	name string
	selflink string
}

func (s Service)GetServiceInfo() {

	serviceClient := connect.Clientset.Services(apiv1.NamespaceAll)

	list_2, err := serviceClient.List(metav1.ListOptions{})

	if err != nil {
		panic(err)
	}

	for _, d := range list_2.Items {
		s.name = d.GetName()
		s.selflink = d.SelfLink

		_, err_1 := connect.ConnRedi.Do("SET",s.name,s.selflink)

		if err_1 != nil {
			fmt.Printf("ERROR", err_1)
		}

	}
}
