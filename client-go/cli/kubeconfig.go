package cli

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func KubeCLI() kubernetes.Clientset {
	var kubeConfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeConfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeConfig = flag.String("kubeconfig", "~/.kube/config", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		panic("kubeConfig error")
	}
	cli, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic("Clientset error")
	}
	return *cli
}
