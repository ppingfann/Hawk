package main

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiv1 "k8s.io/api/core/v1"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"flag"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"
)

func main() {
	//connect to cluster which use the config blow
	kubeconfig := flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	flag.Parse()

	if *kubeconfig == "" {
		panic("-kubeconfig not specified")
	}

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	Clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		panic(err)
	}
	//connect to redis
	conn,err := redis.Dial("tcp","127.0.0.1:6379")

	//acquire pod information
	podsClient := Clientset.Pods(apiv1.NamespaceAll)

	list, err := podsClient.List(metav1.ListOptions{
	})

	if err != nil {
		panic(err)
	}

	for _, d := range list.Items {
		_ ,err_1 := conn.Do("SET",d.GetName(),d.Status.PodIP)

		if err_1 != nil{
			fmt.Printf("ERROR",err_1)
		}

	}

	serviceClient := Clientset.Services(apiv1.NamespaceAll)

	list_2, err := serviceClient.List(metav1.ListOptions{})

	if err != nil {
		panic(err)
	}

	for _, d := range list_2.Items {
		_ ,err_1 := conn.Do("SET",d.GetName(),d.SelfLink)

		if err_1 != nil{
			fmt.Printf("ERROR",err_1)
		}

	}

}
