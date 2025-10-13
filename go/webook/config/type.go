package config

type config struct {
	DBConfig DBConfig
	RedisConfig RedisConfig
}

type DBConfig struct {
	DBAddr string
	DBUser string
	DBPassword string
}

type RedisConfig struct {
	RedisAddr string
	RedisUser string
	RedisPassword string
}