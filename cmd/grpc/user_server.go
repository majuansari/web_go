package grpc

import (
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"web/app"
	"web/cmd"
	"web/config"
	"web/grpc/interceptor"
	userpb "web/grpc/proto/user/go"
	"web/grpc/servers/user"
	"web/pkg/telemetry"
)

var (
	// Create a metrics registry.
	promRegistry = prometheus.NewRegistry()

	// Create some standard server metrics.
	grpcMetrics = grpc_prometheus.NewServerMetrics()
)

// helloCmd represents the serve command
var GrpcStart = &cobra.Command{
	Use:   "start-user-grpc-server",
	Short: "Start user grpc server",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

var addr, promAddr int

func init() {
	promRegistry.MustRegister(grpcMetrics) //, prometheus.NewGoCollector(),prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))
	cmd.RootCmd.AddCommand(GrpcStart)
	GrpcStart.Flags().IntVarP(&addr, "grpc_port", "g", 50051, "Grpc Port")
	GrpcStart.Flags().IntVarP(&promAddr, "prom_port", "p", 9092, "Prometheus Port")

}
func start() {
	cfg := config.NewEnvConfig()

	app, _ := app.NewApp(cfg)

	//Initialise tracer
	flush := telemetry.InitTracer(cfg.Tracer)
	defer flush()
	con, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", addr))
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			otelgrpc.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			interceptor.ClientAuthUnaryInterceptor(),
		)),
	)
	// Initialize all metrics.
	grpcMetrics.InitializeMetrics(grpcServer)

	promHandler := promhttp.HandlerFor(
		promRegistry,
		promhttp.HandlerOpts{
			// Opt into OpenMetrics to support exemplars.
			EnableOpenMetrics: true,
		},
	)
	//http.Handle("/metrics", promHandler)
	// Create a HTTP app for prometheus.
	httpServer := &http.Server{Handler: promHandler, Addr: fmt.Sprintf("0.0.0.0:%d", promAddr)}
	// Start your http app for prometheus.
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal("Unable to start a http app.")
		}
	}()

	userpb.RegisterUserDetailsServiceServer(grpcServer, &user.Server{app}) //&app is a grpc handler receiver which implements a method
	// Start your gRPC app.
	log.Fatal(grpcServer.Serve(con))
}
