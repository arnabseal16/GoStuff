package services_gwcp

import (
	"flag"
	"fmt"
	"path/filepath"

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

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig) //Build config obj from kubeconfig
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config) //Builds the clientset from config obj
	if err != nil {
		panic(err)
	}

	fmt.Println(clientset)
	//deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

}
