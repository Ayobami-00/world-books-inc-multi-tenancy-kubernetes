package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv            string `mapstructure:"ENVIRONMENT"`
	ServerPort        string `mapstructure:"SERVER_PORT"`
	ContextTimeout    int    `mapstructure:"CONTEXT_TIMEOUT"`
	BackendServiceUrl string `mapstructure:"BACKEND_SERVICE_URL"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
