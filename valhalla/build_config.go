package valhalla

import "encoding/json"

type Config struct {
	Json map[string]interface{}
}

func DefaultConfig() *Config {
	var config map[string]interface{}
	err := json.Unmarshal([]byte(defaultConfigString), &config)
	if err != nil {
		panic(err)
	}
	return &Config{Json: config}
}

func (config *Config) String() string {
	marshal, err := json.Marshal(config.Json)
	if err != nil {
		return err.Error()
	}
	return string(marshal)
}

func (config *Config) SetTileDirPath(path string) {
	mjolnir := config.Json["mjolnir"].(map[string]interface{})
	mjolnir["tile_dir"] = path
}

func (config *Config) SetTileExtractPath(path string) {
	mjolnir := config.Json["mjolnir"].(map[string]interface{})
	mjolnir["tile_extract"] = path
}

func (config *Config) SetLoggingVerbosity(verbose bool) {
	mjolnir := config.Json["mjolnir"].(map[string]interface{})
	logging := mjolnir["logging"].(map[string]interface{})
	logging["type"] = verbose
}
