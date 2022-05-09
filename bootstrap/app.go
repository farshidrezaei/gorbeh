package appBootstrap

import (
	libraryDotEnv "farshidrezaei/vms_mongo_log/library/dotenv"
)

func Bootstrap() {
	libraryDotEnv.LoadEnvironmentVariables()
}
