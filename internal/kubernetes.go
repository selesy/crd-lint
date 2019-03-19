package internal

type Kubernetes struct {
	Cfg    Config
	CRDMap CRDMap
}

func NewKubernetes(cfg Config) (Kubernetes, error) {
	var k Kubernetes
	return k, nil
}

func (k Kubernetes) ClusterInfo() {

}

func (k Kubernetes) CRDs() []string {
	return nil
}

func (k Kubernetes) CRD(name string) {

}

// func kubeConfigPath() (string, error) {
// 	var kubeconfig string
// 	if home := homeDir(); home != "" {
// 		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
// 	} else {
// 		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
// 	}
// 	return kubeconfig, nil
// }
