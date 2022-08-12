package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubenetesproamming/cli"
)

func main() {
	kubernetesClient := cli.KubeCLI()
	pod, _ := kubernetesClient.CoreV1().Pods("kube-system").Get(context.TODO(), "pod", metav1.GetOptions{})
	fmt.Println(pod)
}
