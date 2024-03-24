package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "C:\\Users\\ariji\\.kube\\config", "Location of KubeConfig - ")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		//handle errors
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		//handle error
	}
	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		//handle error
	}
	fmt.Println("pods in default namespace are - ")
	for _, pod := range pods.Items {
		fmt.Printf("%s", pod.Name)
	}

	deployments, err := clientset.AppsV1().Deployments("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		//handle error
	}
	fmt.Println("")
	fmt.Println("deployments in default namespaces are - ")
	for _, deployment := range deployments.Items {
		fmt.Printf("%s", deployment.Name)
	}
}
