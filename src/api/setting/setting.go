package setting

import "os"

// settngs
var (
	AppName    = "Test API Sima-land"
	AppVersion = "0.1"

	APIStaticDir  = "../../static"
	APIListenAddr = "127.0.0.1:9080"

	DBUser     = "root"
	DBPassword = ""
	DBHost     = ""
	DBName     = "test"
)

// OpenEnv - устанавливает настройки окружения
func OpenEnv() {
	var list = []string{
		"APIStaticDir", "APIListenAddr",
		"DBUser", "DBPassword", "DBHost", "DBName",
	}

	for _, key := range list {
		val := os.Getenv(key)
		if val != "" {
			readParam(key, val)
		}
	}
}

func readParam(key string, val interface{}) {
	switch key {
	case "APIStaticDir":
		APIStaticDir = val.(string)
	case "APIListenAddr":
		APIListenAddr = val.(string)
	case "DBUser":
		DBUser = val.(string)
	case "DBPassword":
		DBPassword = val.(string)
	case "DBHost":
		DBHost = val.(string)
	case "DBName":
		DBName = val.(string)
	}
}
