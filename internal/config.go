package internal

import (
	"flag"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

const (
	KubeConfigPathUsage         string = "absolute path to the kubeconfig file"
	KubeConfigPathOptionalUsage string = "(optional) " + KubeConfigPathUsage
	CRDPathUsage                string = "absolute path to a directory containing the CRD paths"
	RunningInPodUsage           string = "true configures crd-lint to run in a Kubernetes pod"
)

type Config struct {
	KubeConfigPath string
	CRDPath        string
	RunningInPod   bool
	LogLevel       log.Level
	LogFormatter   log.Formatter
	ManifestPaths  []string
}

func NewConfig() Config {
	var cfg Config
	if home := homeDir(); home != "" {
		flag.StringVar(&cfg.KubeConfigPath, "kubeconfig", filepath.Join(home, ".kube", "config"), KubeConfigPathUsage)
	} else {
		flag.StringVar(&cfg.KubeConfigPath, "kubeconfig", "", KubeConfigPathOptionalUsage)
	}
	flag.StringVar(&cfg.CRDPath, "crds", "", CRDPathUsage)
	flag.BoolVar(&cfg.RunningInPod, "pod", false, RunningInPodUsage)
	flag.Parse()

	log.Info("Configuration: ")
	log.Info("  Offline mode: ", cfg.Offline())
	log.Info("  CRD source: ", cfg.crdSource())
	log.Info("  Running in pod: ", cfg.RunningInPod)

	mans := flag.Args()
	if len(mans) < 1 {
		log.Error("No manifests to validate/lint")
		os.Exit(1)
	}
	log.Info("Manifests:")
	for _, man := range mans {
		log.Info("  ", man)
	}

	return cfg
}

func (c Config) Offline() bool {
	return c.CRDPath != ""
}

func (c Config) crdSource() string {
	if c.Offline() {
		return c.CRDPath
	}
	return "Kubernetes cluster"
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" { // Linux and Mac OSX
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
