package config

import (
	libraryConfig "farshidrezaei/vms_mongo_log/library/config"
	libraryDotEnv "farshidrezaei/vms_mongo_log/library/dotenv"
)

func App(key string, defaultValue any) any {

	configs := map[string]any{
		"name":      libraryDotEnv.Env("APP_NAME", "goravel"),
		"lang":      "apple",
		"time_zone": "UTC",
		"env":       "local",
	}

	return libraryConfig.Get(configs, key, defaultValue)
}
