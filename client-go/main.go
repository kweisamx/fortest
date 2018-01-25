/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	//"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func main() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		//pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
		//if err != nil {
		//	panic(err.Error())
		//}
		//fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		// Examples for error handling:
		// - Use helper functions like e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		namespace := "default"
		//pod := "stress-gin-866bc678bd-kqdt6"
		//p, _ := clientset.CoreV1().Pods(namespace).Get(pod, metav1.GetOptions{})
		p, _ := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
		rc := "stress-gin"
		for _, e := range p.Items {
			fmt.Printf(e.Spec.NodeName + ": ")
			fmt.Println(e.ObjectMeta.Name)
			//fmt.Println(e.Status.ContainerStatuses[0].Name)
		}
		//r, _ := clientset.ExtensionsV1beta1().Deployments(namespace).List(metav1.ListOptions{})
		r, _ := clientset.ExtensionsV1beta1().Deployments(namespace).GetScale(rc, metav1.GetOptions{})
		//fmt.Println("rc num is " , *(r.Items[0].Spec.Replicas))
		fmt.Println("rc num is", r.Spec.Replicas)
		r.Spec.Replicas = 1
		_, err := clientset.ExtensionsV1beta1().Deployments(namespace).UpdateScale(rc, r)

		fmt.Println("rc num is", r.Spec.Replicas)

		fmt.Println("err", err)

		/*
			if errors.IsNotFound(err) {
				fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
			} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
				fmt.Printf("Error getting pod %s in namespace %s: %v\n",
					pod, namespace, statusError.ErrStatus.Message)
			} else if err != nil {
				panic(err.Error())
			} else {
				fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
			}
		*/
		time.Sleep(2 * time.Second)
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
