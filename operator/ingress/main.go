package main

import (
	"ingress/pkg"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		inClusterConfig, err := rest.InClusterConfig()
		if err != nil {
			log.Fatalln("can't find config")
		}
		config = inClusterConfig
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln("can't create client")
	}

	factory := informers.NewSharedInformerFactory(clientSet, 0)
	serviceInformer := factory.Core().V1().Services()
	ingessInformrer := factory.Networking().V1().Ingresses()
	controller := pkg.NewController(clientSet, serviceInformer, ingessInformrer)

	stopCh := make(chan struct{})
	factory.Start(stopCh)
	factory.WaitForCacheSync(stopCh)

	controller.Run(stopCh)
}
