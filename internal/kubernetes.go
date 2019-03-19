package internal

import (
	"k8s.io/client-go/tools/clientcmd"
)

type Kubernetes struct {
	Cfg    clientcmd.ClientConfig
	CRDMap CRDMap
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

func (k Kubernetes) APIResources() {

}

func kubeConfig(cfg Config) (clientcmd.ClientConfig, error) {
	if cfg.RunningInPod {
		return kubeConfigPod(cfg)
	}
	return kubeConfigClient(cfg)
}

func kubeConfigClient(cfg Config) (clientcmd.ClientConfig, error) {
	//return clientcmd.BuildConfigFromFlags("", &cfg.KubeConfigPath)
	return nil, nil
}

func kubeConfigPod(cfg Config) (clientcmd.ClientConfig, error) {
	return nil, nil
}
