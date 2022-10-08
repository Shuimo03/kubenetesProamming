package resource

import (
	"client-go/cli"
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetPod() []string {
	kubeCli := cli.KubeCLI()
	res := make([]string, 0)
	podList, err := kubeCli.CoreV1().Pods("kube-system").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}
	for _, pod := range podList.Items {
		res = append(res, pod.Name)
	}
	return res
}
