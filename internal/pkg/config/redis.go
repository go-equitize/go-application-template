package config

// RedisConfig contains all configuration of redis
type RedisConfig struct {
	RedisAddresses string
	Password       string
	MasterName     string
}
