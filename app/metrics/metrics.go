package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	ResponseLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "response_latency",
			Help:    "Response calculate latency (histogram)",
			Buckets: []float64{.0001, .0005, .005, .015, .1},
		},
		[]string{"step"},
	)
)

func init() {
	prometheus.MustRegister(ResponseLatency)
}

func ObserveLatency(step string, duration float64) {
	ResponseLatency.WithLabelValues(step).Observe(duration)
}
