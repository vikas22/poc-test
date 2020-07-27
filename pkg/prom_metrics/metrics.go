package prom_metrics

import (
  "fmt"
  "github.com/prometheus/client_golang/prometheus"
  "github.com/prometheus/client_golang/prometheus/promauto"
  "time"
)

const (
	Namespace = "bench"
)

var (
	opsProcessed = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: Namespace,
		Name:      "db_operation",
		Help:      "A request counter for db operation",
	}, []string{"operation", "success"})

	dbRequestDurationHistogram = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: Namespace,
		Name:      "db_duration_histogram",
		Help:      "DB request duration histogram",
		Buckets:   prometheus.LinearBuckets(5, 10, 2000), // bucket starting with 5ms, 10 ms width till 20_000 ms
	}, []string{"operation", "success"})

	dbRequestDurationSummary = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace:  Namespace,
			Name:       "db_duration_summary",
			Help:       "DB request duration summary",
			Objectives: map[float64]float64{0.5: 0.5, 0.9: 0.9, 0.99: 0.99},
		}, []string{"operation", "success"})
)

func IncOperation(operation string, success bool) {
	successStr := "false"
	if success {
		successStr = "true"
	}
	opsProcessed.WithLabelValues(operation, successStr).Inc()
}

func DbRequestDuration(operation string, success bool, t0 time.Time) {
	successStr := "false"
	if success {
		successStr = "true"
	}
	duration := float64(time.Now().Sub(t0).Milliseconds())

	fmt.Println(operation, " = " , duration)
	dbRequestDurationSummary.WithLabelValues(operation, successStr).Observe(duration)
	dbRequestDurationHistogram.WithLabelValues(operation, successStr).Observe(duration)

}
