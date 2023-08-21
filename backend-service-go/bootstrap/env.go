package bootstrap

type Env struct {
	AppEnv            string `mapstructure:"ENVIRONMENT"`
	ServerPort        string `mapstructure:"SERVER_PORT"`
	ContextTimeout    int    `mapstructure:"CONTEXT_TIMEOUT"`
	BackendServiceUrl string `mapstructure:"BACKEND_SERVICE_URL"`
}

func NewEnv() *Env {
	env := Env{
		AppEnv:            "PRODUCTION",
		ServerPort:        ":8080",
		ContextTimeout:    10,
		BackendServiceUrl: "0.0.0.0:8080",
	}

	// viper.SetConfigFile(".env")

	// err := viper.ReadInConfig()
	// if err != nil {
	// 	log.Fatal("Can't find the file .env : ", err)
	// }

	// err = viper.Unmarshal(&env)
	// if err != nil {
	// 	log.Fatal("Environment can't be loaded: ", err)
	// }

	// if env.AppEnv == "development" {
	// 	log.Println("The App is running in development env")
	// }

	return &env
}
