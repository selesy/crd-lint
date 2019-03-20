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

	var k8s internal.Kubernetes
	var err error
	if !cfg.Offline() {
		k8s, err = internal.NewKubernetes(cfg)
	}
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	crds, err := internal.NewCRDMap(cfg, k8s)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	if len(crds) <= 1 {
		log.Warn("No CRDs were loaded ... no validation will be done")
	}

	log.Trace("main() ->")
}
