package config

import "os"

const (
	StatusPageIOApiBase = "https://api.statuspage.io/v1"
)

var (
	ApiRoot string
	Env     string
	Port    string
	Debug   string

	StatusPageIOApiKey   string
	StatusPageIOMetricID string
	StatusPageIOPageID   string
	StatusPageIOEnable   string

	MongoURL          string
	MongoDatabaseName string

	PrivateKey []byte

	NewRelicLicense string
	NewRelicAppName string

	MemcacheURL      string
	MemcacheUsername string
	MemcachePassword string
)

func init() {
	ApiRoot = env("API_ROOT", false)
	Env = env("ENV", false)
	Port = env("PORT", false)
	Debug = env("DEBUG", true)

	StatusPageIOEnable = env("STATUSPAGEIO_ENABLE", false)
	StatusPageIOApiKey = env("STATUSPAGEIO_API_KEY", StatusPageIOEnable != "true")
	StatusPageIOMetricID = env("STATUSPAGEIO_METRIC_ID", StatusPageIOEnable != "true")
	StatusPageIOPageID = env("STATUSPAGEIO_PAGE_ID", StatusPageIOEnable != "true")

	MongoURL = env("MONGO_URL", false)
	MongoDatabaseName = env("MONGO_DATABASE_NAME", false)

	PrivateKey = []byte(env("PRIVATE_KEY", false))

	NewRelicLicense = env("NEW_RELIC_LICENSE", true)
	NewRelicAppName = env("NEW_RELIC_APP_NAME", true)

	MemcacheURL = env("MEMCACHED_URL", false)
	MemcacheUsername = env("MEMCACHED_USERNAME", true)
	MemcachePassword = env("MEMCACHED_PASSWORD", true)
}

func env(s string, optional bool) string {
	v := os.Getenv(s)
	if v == "" && !optional {
		panic("VAR " + s + " empty, run:\n\nexport " + s + "='SOME VALUE'\n\n")
	}

	return v
}
