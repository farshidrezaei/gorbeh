package config

import (
	libraryConfig "farshidrezaei/vms_mongo_log/library/config"
	"farshidrezaei/vms_mongo_log/library/dotenv"
)

func Database(key string, defaultValue any) any {

	configs := map[string]any{
		"connections": map[string]any{
			"mysql": map[string]any{
				"host":     libraryDotEnv.Env("MYSQL_HOST", "127.0.0.1"),
				"port":     libraryDotEnv.Env("MYSQL_PORT", "3306"),
				"database": libraryDotEnv.Env("MYSQL_DATABASE", "mikey"),
				"username": libraryDotEnv.Env("MYSQL_USERNAME", "root"),
				"password": libraryDotEnv.Env("MYSQL_PASSWORD", ""),
			},
			"redis": map[string]any{
				"host":     libraryDotEnv.Env("REDIS_HOST", "127.0.0.1"),
				"port":     libraryDotEnv.Env("REDIS_PORT", "6379"),
				"password": libraryDotEnv.Env("REDIS_PASSWORD", ""),
				"db":       libraryDotEnv.Env("REDIS_DB", "0"),
			},
		},
		"time_zone": "UTC",
		"env":       "local",
	}
	return libraryConfig.Get(configs, key, defaultValue)

}
