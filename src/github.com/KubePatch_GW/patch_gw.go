package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {

	//fetch the kubeconfig
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("041922983057-host1", *kubeconfig) //Build config obj from kubeconfig
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config) //Builds the clientset from config obj
	if err != nil {
		panic(err)
	}

	options := metav1.ListOptions{
		LabelSelector: "app=<APPNAME>",
	}
	deploymentsClient, err := clientset.AppsV1().Deployments("contorller-systems").List(context.TODO(), options)
	if err != nil {
		panic(err)
	}

	// result, getErr := deploymentsClient.Get(context.TODO(), "", metav1.GetOptions{})
	// if getErr != nil {
	// 	panic(fmt.Errorf("failed to get latest version of deployment: %v", getErr))
	// }

	fmt.Println(deploymentsClient)

}
