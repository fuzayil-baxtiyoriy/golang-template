package config

import "golang-template/pkg/ds"

const (
	DevMode  = "DEVELOPMENT"
	ProdMode = "PRODUCTION"
)

const (
	baseCfgFilename = "base.yaml"
	envFilename     = ".env"
)

var cfgFileMapper = map[string]string{
	DevMode:  "dev.yaml",
	ProdMode: "prod.yaml",
}

var availableModes = ds.NewSet(DevMode, ProdMode)
