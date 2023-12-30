package main

import (
	//Internal 
	"internal/config"
	
	//STD
	"context"
	"flag"
	"fmt"
	"path/filepath"
	// "time"
	// "os"
	"log"
	"github.com/joho/godotenv"

	//K8s
	"k8s.io/client-go/rest"
	// "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

func init () {
	err := godotenv.Load(".env")
	if (err != nil) {
		log.Println("WARNING Failed to load from .env file")
	}
}

func main () {
	
	var err error
	conf := config.New()

	var config *rest.Config
	if !conf.InCluster {
		var kubeconfig *string
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()

		// Set Config variable
		config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
	
		if err != nil {
			panic(err.Error())
		}
	
	} else if conf.InCluster {
		config, err = rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
	}


	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	certCM, err := clientset.CoreV1().ConfigMaps("kube-system").Get(context.TODO(),"kube-root-ca.crt", metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(certCM.Data["ca.crt"])
}