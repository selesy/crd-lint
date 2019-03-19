package main

import (
	//"k8s.io/apiextensions-apiserver/pkg/apiserver/validation"
	//"k8s.io/apimachinery/pkg/util/validation"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
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
	cfg, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Error(err)
		panic(err.Error)
	}

	// cl.BatchV1().Jobs("eio-swe")
	// _ = &batchv1.Job{
	// 	ObjectMeta: metav1.ObjectMeta{
	// 		Name:      "aleks-oneshot",
	// 		Namespace: "eio-swe",
	// 	},
	// 	Spec: batchv1.JobSpec{
	// 		Template: apiv1.PodTemplateSpec{
	// 			Spec: apiv1.PodSpec{
	// 				Containers: []apiv1.Container{
	// 					{
	// 						Name:  "aleks-oneshot",
	// 						Image: "aleks-generator",
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }

	// log.Info("Got here!")

	// pods, err := cl.CoreV1().Pods("").List(metav1.ListOptions{})
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	// for _, pod := range pods.Items {
	// 	log.Info("Pod name: ", pod.Name)
	// }

	extcs, err := apiextensionsclientset.NewForConfig(cfg)
	if err != nil {
		log.Error(err)
		panic(err.Error)
	}

	crds, err := extcs.ApiextensionsV1beta1().CustomResourceDefinitions().List(metav1.ListOptions{})
	if err != nil {
		log.Error(err)
		panic(err.Error)
	}

	//var specs = map[string]
	for _, crd := range crds.Items {
		log.Info(crd.ObjectMeta.Name)
		log.Info("     ", crd.Spec.Version)
		log.Info("          ", crd.Spec)
		v := crd.Spec.Validation
		if v != nil {
			log.Info("          ", v.OpenAPIV3Schema)

		}
	}

	data, err := ioutil.ReadFile("directory-service-helm-release.yaml")
	if err != nil {
		log.Error(err)
		panic(err.Error)
	}

	var crd apiextensionsv1beta1.CustomResourceDefinition
	err = yaml.Unmarshal(data, &crd)
	if err != nil {
		log.Error(err)
		panic(err.Error)
	}

	log.Info(crd)

	//TODO: errorList := apiextensionsvalidation.ValidateCustomResourceDefinition(crd)
	//log.Info(errorList)

	// cl.Get()
	// crd := extensions.CustomResourceDefinition{}

	// validation.ValidateCustomResourceDefinition(&crd)

	// Examples for error handling:
	// - Use helper functions like e.g. errors.IsNotFound()
	// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
	// namespace := "default"
	// pod := "example-xxxxx"
	// _, err = cl.CoreV1().Pods(namespace).Get(pod, metav1.GetOptions{})
	// if errors.IsNotFound(err) {
	// 	fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
	// } else if statusError, isStatus := err.(*errors.StatusError); isStatus {
	// 	fmt.Printf("Error getting pod %s in namespace %s: %v\n",
	// 		pod, namespace, statusError.ErrStatus.Message)
	// } else if err != nil {
	// 	panic(err.Error())
	// } else {
	// 	fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
	// }
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func crds() {

}
