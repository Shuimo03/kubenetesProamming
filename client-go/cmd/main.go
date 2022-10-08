package main

import (
	"client-go/resource"
	"fmt"
)

func main() {
	pod := resource.GetPod()
	fmt.Println(pod)
}
