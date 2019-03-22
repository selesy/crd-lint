package internal

import (
	"errors"
	"io"
	"os"

	"k8s.io/apimachinery/pkg/util/yaml"

	log "github.com/sirupsen/logrus"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

type CRDMap map[string]apiextensionsv1beta1.CustomResourceDefinition

func NewCRDMap(cfg Config, k8s Kubernetes) (CRDMap, error) {
	log.Trace("-> NewCRDMap(Config, Kubernetes)")

	crds, err := loadCRDs(cfg, k8s)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	c := make(CRDMap)
	for _, crd := range crds {
		c[crd.ObjectMeta.Name] = crd
	}

	if len(c) > 0 {
		log.Info("Custom Resource Definitions (CRDs)")
		for k, v := range c {
			if v.Spec.Validation == nil {
				log.Warn("  ", k, " - ", v.Spec.Version, " (No validation provided)")
				continue
			}
			log.Info("  ", k, " - ", v.Spec.Version)
		}
	}

	log.Trace("NewCRDMap(Config, Kubernetes) ->")
	return c, nil
}

func loadCRDs(cfg Config, k8s Kubernetes) ([]apiextensionsv1beta1.CustomResourceDefinition, error) {
	if cfg.Offline() {
		return loadCRDsFromPath(cfg)
	}
	return loadCRDsFromKubernetes(k8s)
}

func loadCRDsFromKubernetes(k8s Kubernetes) ([]apiextensionsv1beta1.CustomResourceDefinition, error) {
	return k8s.CRDs()
}

func loadCRDsFromPath(cfg Config) ([]apiextensionsv1beta1.CustomResourceDefinition, error) {
	fi, err := os.Stat(cfg.CRDPath)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if !fi.IsDir() {
		err = errors.New("")
		log.Error(err)
		return nil, err
	}

	dir, err := os.Open(cfg.CRDPath)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	fis, err := dir.Readdir(0)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var c []apiextensionsv1beta1.CustomResourceDefinition
	log.Info("Files:")
	for _, fi := range fis {
		fn := cfg.CRDPath + fi.Name()
		log.Info("  ", fn)
		if fi.IsDir() {
			continue
		}
		f, err := os.Open(fn)
		if err != nil {
			log.Error(err)
			continue
		}
		defer f.Close()
		var crd apiextensionsv1beta1.CustomResourceDefinition
		err = yaml.NewYAMLOrJSONDecoder(f, 100).Decode(&crd)
		if err != nil {
			if err != io.EOF {
				log.Error(err)
			}
			continue
		}
		if crd.Kind == "CustomResourceDefinition" {
			c = append(c, crd)
		}
	}

	return c, nil
}
