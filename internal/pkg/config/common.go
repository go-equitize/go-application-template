package config

// CommonConfig contains all common configuration of application
type CommonConfig struct {
	LogLevel                          int `default:"0"`
	HttpClientRetryCount              int `default:"3"`
	HttpClientRetryWaitTimeSeconds    int `default:"5"`
	HttpClientRetryMaxWaitTimeSeconds int `default:"30"`
}
