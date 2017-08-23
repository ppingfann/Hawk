package connect

import (
	"flag"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"
)

func ConnectKube() *kubernetes.Clientset{//connect to cluster which use the config blow
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
	return Clientset
}

	var Clientset *kubernetes.Clientset = ConnectKube()
