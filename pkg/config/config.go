package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

var (
	dbConfig     DBConfig
	jwtConfig    JWTConfig
	jaegerConfig JaegerConfig
)

type JaegerConfig struct {
	Enabled      bool    `envconfig:"JAEGER_ENABLED" default:"false"`
	ServiceName  string  `envconfig:"JAEGER_SERVICE_NAME" default:"red-package-api"`
	Endpoint     string  `envconfig:"JAEGER_ENDPOINT" default:"127.0.0.1:6382"`
	SamplerType  string  `envconfig:"JAEGER_SAMPLER_TYPE" default:"const"`
	SamplerParam float64 `envconfig:"JAEGER_SAMPLER_PARAM" default:"1"`
}

type DBConfig struct {
	DBName     string `envconfig:"DBNAME"`
	DBURL      string `envconfig:"DBURL"`
	DBPort     string `envconfig:"DBPORT"`
	DBUserName string `envconfig:"DBUSERNAME"`
	DBPassword string `envconfig:"DBPASSWORD"`
}

type JWTConfig struct {
	APISecret         string `envconfig:"API_SECRET"`
	TokenHourLifeSpan string `envconfig:"TOKEN_HOUR_LIFESPAN"`
}

func SetConfig() {
	configs := []interface{}{
		&dbConfig,
		&jwtConfig,
		&jaegerConfig,
	}
	for _, instance := range configs {
		err := envconfig.Process("", instance)
		if err != nil {
			log.Fatalf("unable to init config: %v, err: %v", instance, err)
		}
	}
}

func DbConfig() DBConfig {
	return dbConfig
}

func JwtConfig() JWTConfig {
	return jwtConfig
}

func TracingConfig() JaegerConfig {
	return jaegerConfig
}
