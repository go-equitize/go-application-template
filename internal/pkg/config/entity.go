package config

type HttpConfig struct {
	BindAddress string
	Mode        string
	Prefix      string
}

type CommonConfig struct {
	LogLevel int8

	// HTTP Client

	HttpClientRetryCount              int `default:"3"`
	HttpClientRetryWaitTimeSeconds    int `default:"5"`
	HttpClientRetryMaxWaitTimeSeconds int `default:"30"`
}

type RedisConfig struct {
	RedisAddresses string
	Password       string
	MasterName     string
}

type AuthConfig struct {
	BaseURL      string
	BaseAdminURL string
}
