package main

import (
	"os"

	"github.com/selesy/crd-lint/internal"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Trace("-> main()")
	log.Info("----- crd-lint -----")

	cfg := internal.NewConfig()
	log.Debug("Config: ", cfg)

	crds, err := loadCRDs(cfg)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	if len(crds) <= 1 {
		log.Warn("No CRDs were loaded ... no validation will be done")
	}
	if len(crds) > 0 {
		log.Info("Custom Resource Definitions (CRDs)")
		for key := range crds {
			log.Info("  ", key)
		}
	}

	log.Trace("main() ->")
}

func loadCRDs(cfg internal.Config) (internal.CRDMap, error) {
	if cfg.Offline() {
		return loadCRDsFromPath(cfg)
	}
	return loadCRDsFromKubernetes(cfg)
}

func loadCRDsFromKubernetes(cfg internal.Config) (internal.CRDMap, error) {
	k8s, err := internal.NewKubernetes(cfg)
	if err != nil {
		return nil, err
	}
	return k8s.CRDMap, nil
}

func loadCRDsFromPath(cfg internal.Config) (internal.CRDMap, error) {
	return make(internal.CRDMap), nil
}
