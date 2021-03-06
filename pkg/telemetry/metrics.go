package telemetry

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

// InitMeter ...
func InitMetrics() {

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		_ = http.ListenAndServe(":2222", nil)
	}()

}

//@todo too early to start using otel metrics. Review this again.
// InitMeter ...
//func InitMetricsOtel() {
//	exporter, err := prometheus.InstallNewPipeline(
//		prometheus.Config{
//			DefaultHistogramBoundaries: []float64{8000.00, 10000.00, 13000.00, 16000.00},
//		},
//	)
//	if err != nil {
//		log.Fatalf("Failed to init prometheus exporter: %v", err)
//	}
//
//	http.HandleFunc("/metrics", exporter.ServeHTTP)
//
//	go func() {
//		// runtime metrics instrumentation
//		if err = runtime.Start(runtime.WithMinimumReadMemStatsInterval(time.Second)); err != nil {
//			log.Fatalf("Failed to init runtime instrumentation: %v", err)
//		}
//
//		if err = host.Start(); err != nil {
//			log.Fatalf("Failed to init host instrumentation: %v", err)
//		}
//
//		_ = http.ListenAndServe(":2222", nil)
//
//		fmt.Println("Metric server running on :2222")
//	}()
//}
