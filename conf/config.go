package conf

import "github.com/caarlos0/env/v6"

type AppConfig struct {
	Port string `env:"PORT" envDefault:"8080"`

	LogFormat string `env:"LOG_FORMAT" envDefault:"text"`
	DBHost    string `env:"DB_HOST" envDefault:"localhost"`
	DBPort    string `env:"DB_PORT" envDefault:"5432"`
	DBUser    string `env:"DB_USER" envDefault:"postgres"`
	DBPass    string `env:"DB_PASS" envDefault:"P@ssw0rd"`
	DBName    string `env:"DB_NAME" envDefault:"intern_db"`
	EnableDB  string `env:"ENABLE_DB" envDefault:"true"`

	SecretKey string `env:"SECRET_KEY"`
}

var config AppConfig

func SetEnv() {
	_ = env.Parse(&config)
}

func LoadEnv() AppConfig {
	return config
}
