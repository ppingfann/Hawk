package service

import (
	"connect"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SvcController struct {

}

func (*SvcController)GetSvcInfo(Namespace string,Name string) *SvcInfo{
	serviceClient := connect.Clientset.Services(Namespace)

	list, err := serviceClient.List(metav1.ListOptions{})

	if err != nil {
		panic(err)
	}

	svcinfo := new(SvcInfo)

	for _, d := range list.Items {
		if Name == d.GetName() {
			svcinfo.Name = d.GetName()
			svcinfo.Selflink = d.SelfLink
		}
	}
	return svcinfo
}


