package db

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Dev-Siri/gdn/constants"
	"github.com/Dev-Siri/gdn/models"
	"github.com/Dev-Siri/gdn/utils"
)

var CDNConfig models.Config

func SetupConfig() error {
	if _, err := os.Stat(constants.ConfigFile); os.IsNotExist(err) {
		return fmt.Errorf("a " + constants.ConfigFile + " config file is required with the `origin_server` to start the cdn")
	}

	fileContent, err := os.ReadFile(constants.ConfigFile)

	if err != nil {
		return fmt.Errorf("failed to read config file")
	}

	var config models.Config

	if err := json.Unmarshal(fileContent, &config); err != nil {
		return fmt.Errorf("failed to parse config file")
	}

	if config.OriginServer == "" {
		return fmt.Errorf("the `origin_server` property is required")
	}

	if !utils.IsValidURL(config.OriginServer) {
		return fmt.Errorf("the given origin server property is not valid")
	}

	if config.CacheDir == "" {
		config.CacheDir = ".cache"
	}

	CDNConfig = config

	return nil
}
