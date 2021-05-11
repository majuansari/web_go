package app

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web/config"
	g "web/grpc"
	"web/pkg/cache"
	"web/pkg/db"
	appError "web/pkg/error"
	appHttp "web/pkg/http"
	"web/pkg/telemetry"
)

//App struct
type App struct {
	Echo       *echo.Echo
	DB         *gorm.DB
	Cache      cache.Manager
	Grpc       map[string]*grpc.ClientConn
	HttpClient *http.Client
	//EnvConfig *config.EnvConfig
}

//NewApp func
func NewApp(cfg *config.EnvConfig) (*App, func()) {
	//Initialise tracer
	tracerFlush := telemetry.InitTracer(cfg.Tracer)
	//Initialise metrics
	telemetry.InitMetrics()

	dbCon, dbClose, err := db.NewDBConnection(cfg.DB)
	if err != nil {
		log.Fatalf("%v", err)
	}

	cache, err := cache.NewCacheManager(cfg.Cache)
	if err != nil {
		log.Fatalf("Can't connect to cache %v", err)
	}

	httpClient := appHttp.InitHttpClient()

	userGrpcCon, userGrpcConClose := g.InitGrpC("127.0.0.1:50051", grpc.WithInsecure(), grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()))

	var grpcCons = map[string]*grpc.ClientConn{
		"UserServer": userGrpcCon,
	}

	cleanup := func() {
		log.Print("clean up func called")
		dbClose()
		tracerFlush()
		userGrpcConClose()
	}

	return &App{
		Echo:       echo.New(),
		DB:         dbCon,
		Cache:      cache,
		HttpClient: httpClient,
		Grpc:       grpcCons,
	}, cleanup
}

//NewMockServer func
func NewMockServer(cfg *config.EnvConfig) *App {
	return &App{
		Echo: echo.New(),
	}
}

// Start func
func (app *App) Start(port string, cleanUp func()) error {

	e := app.Echo

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	// Start the service listening for requests.
	go func() {
		e.Logger.Infof("main : Service listening on %s", port)
		if err := e.Start(":" + port); err != nil {
			e.Logger.Errorf("App stopped %v", err)
			serverErrors <- err
		}
	}()

	//@todo enable debugger only on debug mode
	// serving pprof on different port
	go func() {
		srv := http.Server{
			Addr: ":8001", //@todo load from config or flag
			//Handler: http.DefaultServeMux, // DefaultServeMux is served by default if no handler is provided
		}
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the debug app")
		}
	}()

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// =========================================================================
	// Shutdown

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		e.Logger.Error(err, "error: listening and serving")

	case <-shutdown:
		e.Logger.Info("main : Start shutdown")
		e.Logger.Info("Starting tear down")
		cleanUp()
		// Give outstanding requests a deadline for completion.
		const timeout = 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		// Asking listener to shutdown and load shed.
		err := e.Shutdown(ctx)
		if err != nil {
			e.Logger.Info("main : Graceful shutdown did not complete in %v : %v", timeout, err)
			err = e.Close()
		}

		if err != nil {
			e.Logger.Errorf("main : could not stop app gracefully : %v", err)
		}
	}
	return nil
}

// ConfigureLogger func
func (app *App) ConfigureLogger() {
	app.Echo.Logger.SetLevel(log.DEBUG)
	app.Echo.Logger.SetHeader("${time_rfc3339} ${level}")
}

type HTTPErrorHandler func(error, echo.Context)

// ConfigureErrorHandler func
func (app *App) ConfigureErrorHandler() {
	app.Echo.HTTPErrorHandler = appError.CustomHTTPErrorHandler
}

// Get Grpc Con func
func (app *App) GetGrpcCon(key string) *grpc.ClientConn {
	return app.Grpc[key]
}
