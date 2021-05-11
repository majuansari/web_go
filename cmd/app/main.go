package app

//
//import (
//	"web/app"
//	"web/config"
//	"web/routes"
//)
//
//func main() {
//	cfg := config.NewConfig()
//
//	app, cleanUp := app.NewApp(cfg)
//	defer cleanUp()
//
//	app.ConfigureLogger()
//	app.ConfigureErrorHandler()
//
//	//@todo review if we need to pass cfg just telemetry
//	routes.RegisterRoutes(app, cfg)
//
//	app.Start("8000")
//}
