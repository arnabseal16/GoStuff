package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
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

	kubeConfigPath := *kubeconfig

	config, err := getClusterConfig(kubeConfigPath)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config) //Builds the clientset from config obj
	if err != nil {
		panic(err)
	}
	name := "haproxy-ingress"
	labelSelector := metav1.LabelSelector{MatchLabels: map[string]string{"app": name}}
	options := metav1.ListOptions{
		LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
		//FieldSelector: resources.limits.cpu=100
	}
	deploymentsClient := clientset.AppsV1().Deployments("atmos-system")

	fmt.Println(deploymentsClient.List(context.TODO(), options))
	// result, getErr := deploymentsClient.Get(context.TODO(), "", metav1.GetOptions{})
	// if getErr != nil {
	// 	panic(fmt.Errorf("failed to get latest version of deployment: %v", getErr))
	// }

	// fmt.Println(deploymentsClient)
	// fmt.Println(result)
}

func getClusterConfig(kubeConfigPath string) (*rest.Config, error) {
	if kubeConfigPath != "" {
		//  when not running in cluster
		return clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	}
	return rest.InClusterConfig()
}
