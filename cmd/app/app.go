package app

import (
	"github.com/spf13/viper"
	"web/app"
	"web/app/routes"
	"web/cmd"
	"web/config"

	"github.com/spf13/cobra"
)

// appCmd represents the app command
var appCmd = &cobra.Command{
	Use:   "start-web-app",
	Short: "Start Web Application",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

var (
	port    string
	dbgPort string
	debug   bool
)

func init() {
	cmd.RootCmd.AddCommand(appCmd)
	appCmd.Flags().StringVarP(&port, "port", "p", "8000", "Application port")
	appCmd.Flags().StringVarP(&dbgPort, "dbg_port", "d", "8001", "Application debug port")
	appCmd.Flags().BoolVar(&debug, "debug", true, "Enable/Disable Debug mode")

	viper.BindPFlag("debug", appCmd.Flags().Lookup("debug"))
}

func run() {
	cfg := config.NewEnvConfig()

	app, cleanUp := app.NewApp(cfg)
	app.ConfigureLogger()
	app.ConfigureErrorHandler()

	//@todo review if we need to pass cfg just telemetry
	routes.RegisterRoutes(app, cfg)
	//passing tear down to start method to do the tear down along with shut down
	app.Start("8000", cleanUp)
}
