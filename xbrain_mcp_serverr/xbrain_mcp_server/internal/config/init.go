package config

import "os"

func InitConfig() {
	// Set the run environment
	RunEnvironment = os.Getenv("RUN_ENV")
	if RunEnvironment == "" {
		RunEnvironment = "dev" // Default to dev if not set
	}

	// Set the XnoApiUrl based on the environment
	if RunEnvironment == "prod" {
		XnoApiUrl = "https://api-v2.xno.vn"
		XnoV1ApiUrl = "https://api.xno.vn"
	} else {
		XnoApiUrl = "https://dev-api-v2.xno.vn"
		XnoV1ApiUrl = "https://api.xno.vn"
	}
}
