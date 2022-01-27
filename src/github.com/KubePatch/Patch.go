package services

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const typeAnnotation string = "original-service-type"

// Update changes the service type
func Update(namespace, name, serviceType string, client kubernetes.Interface) error {
	service, err := client.CoreV1().Services(namespace).Get(name, metav1.GetOptions{})

	if err != nil {
		return fmt.Errorf("Failed to get service in namespace \"%s\": %s", namespace, err)
	}

	if _, ok := service.ObjectMeta.Annotations[typeAnnotation]; !ok {
		service.ObjectMeta.Annotations[typeAnnotation] = string(service.Spec.Type)
	}
	service.Spec.Type = v1.ServiceType(serviceType)

	if serviceType != "NodePort" && serviceType != "LoadBalancer" {
		for i := range service.Spec.Ports {
			service.Spec.Ports[i].NodePort = 0
		}
	}

	_, err = client.CoreV1().Services(namespace).Update(service)

	if err != nil {
		return fmt.Errorf("failed to update service: %s", err)
	}

	log.Success("Updated service \"%s\" in namespace \"%s\" to %s", name, namespace, serviceType)
	return nil
}

// Revert service type (only applies if tectl has previously changes service type)
func Revert(namespace, name string, client kubernetes.Interface) error {
	service, err := client.CoreV1().Services(namespace).Get(name, metav1.GetOptions{})

	if err != nil {
		return fmt.Errorf("Failed to get service in namespace \"%s\": %s", namespace, err)
	}

	if serviceType, ok := service.ObjectMeta.Annotations[typeAnnotation]; ok {
		service.Spec.Type = v1.ServiceType(serviceType)

		if serviceType != "NodePort" && serviceType != "LoadBalancer" {
			for i := range service.Spec.Ports {
				service.Spec.Ports[i].NodePort = 0
			}
		}

		delete(service.ObjectMeta.Annotations, typeAnnotation)
		_, err = client.CoreV1().Services(namespace).Update(service)

		if err != nil {
			return fmt.Errorf("failed to update service: %s", err)
		}

		log.Success("Reverted service \"%s\" in namespace \"%s\" to %s", name, namespace, serviceType)
		return nil
	}

	return fmt.Errorf("unable to find required annotation")
}
