package metrics

import (
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"log"
)

var (
	// Create a metrics registry.
	reg = prometheus.NewRegistry()

	// Create some standard server metrics.
	GrpcMetrics = grpc_prometheus.NewServerMetrics()

	//responseLatency = prometheus.NewHistogramVec(
	//	prometheus.HistogramOpts{
	//		Name:    "response_latency",
	//		Help:    "Response calculate latency (histogram)",
	//		Buckets: []float64{.0001, .0005, .005, .015, .1},
	//	},
	//	[]string{"step"},
	//)
)

//func ResponseLatency(step string, duration float64) {
//	log.Println("step", step, "duration", duration)
//	responseLatency.WithLabelValues(step).Observe(duration)
//}

func init() {
	log.Println("Must register done")
	reg.MustRegister(GrpcMetrics) //, responseLatency)
}
