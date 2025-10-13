//go:build dev

package config

var Config = config{
	DB: DBConfig{
		DBAddr: `127.0.0.1:3306`,
		DBUser: `root`,
		DBPassword: `root`,
	},
	RedisConfig: RedisConfig{
		RedisAddr: `127.0.0.1:6379`,
		RedisUser: ``,
		RedisPassword: `redis`,
	},
}