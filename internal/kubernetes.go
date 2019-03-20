package internal

import (
	log "github.com/sirupsen/logrus"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Kubernetes struct {
	Cfg *rest.Config
}

func NewKubernetes(cfg Config) (Kubernetes, error) {
	var k8s Kubernetes

	kcfg, err := kubeConfig(cfg)
	if err != nil {
		return k8s, err
	}
	k8s.Cfg = kcfg

	return k8s, nil
}

func (k Kubernetes) ClusterInfo() {

}

func (k Kubernetes) CRDs() ([]apiextensionsv1beta1.CustomResourceDefinition, error) {
	cl, err := apiextensionsclientset.NewForConfig(k.Cfg)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	crds, err := cl.ApiextensionsV1beta1().CustomResourceDefinitions().List(metav1.ListOptions{})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return crds.Items, nil
}

func (k Kubernetes) APIResources() {

}

func kubeConfig(cfg Config) (*rest.Config, error) {
	if cfg.RunningInPod {
		return kubeConfigPod(cfg)
	}
	return kubeConfigClient(cfg)
}

func kubeConfigClient(cfg Config) (*rest.Config, error) {
	return clientcmd.BuildConfigFromFlags("", cfg.KubeConfigPath)
}

func kubeConfigPod(cfg Config) (*rest.Config, error) {
	return nil, nil
}
