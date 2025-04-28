package config

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
)

func Load() (*Config, error) {
	cfg := new(Config)

	workDir, err := os.Getwd()

	if err != nil {
		return nil, fmt.Errorf("failed to get working dir: %w", err)
	}

	configDir := fmt.Sprintf("%s/configs/", workDir)

	// Read .env file from working directory
	err = cleanenv.ReadConfig(workDir+"/"+envFilename, cfg)
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to read from %s: %w", envFilename, err)
	}

	// Read base config
	err = cleanenv.ReadConfig(configDir+baseCfgFilename, cfg)
	if err != nil && !isEOFerr(err) {
		return nil, fmt.Errorf("failed to read from %s: %w", baseCfgFilename, err)
	}

	if !availableModes.Contains(cfg.AppMode) {
		cfg.AppMode = ProdMode
	}

	// Read environment specific config file
	modeFilename, ok := cfgFileMapper[cfg.AppMode]
	if ok {
		err = cleanenv.ReadConfig(configDir+modeFilename, cfg)
		if err != nil && !isEOFerr(err) {
			return nil, fmt.Errorf("failed to read from %s: %w", modeFilename, err)
		}
	}

	err = cfg.validate()
	if err != nil {
		return nil, err
	}

	return cfg, nil

}

func isEOFerr(err error) bool {
	return strings.HasSuffix(err.Error(), io.EOF.Error())
}
