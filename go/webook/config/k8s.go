//go:build k8s

package config

var Config = config{
	DB: DBConfig{
		DBAddr: `127.0.0.1:3316`,
		DBUser: `root`,
		DBPassword: `root`,
	},
	RedisConfig: RedisConfig{
		RedisAddr: `127.0.0.1:6380`,
		RedisUser: ``,
		RedisPassword: `redis`,
	},
}