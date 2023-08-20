package bootstrap

type Application struct {
	Env *Env
	Db  *Database
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Db = NewDynamoDbDatabase(app.Env)
	return *app
}
