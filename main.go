package main

import (
	"TestMetrck/log"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			log.Logger.Info("counter: ", opsProcessed)
			time.Sleep(2 * time.Second)
			error.Set(123)
		}
	}()

}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})

	error = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "myapp_signal_errors",
		Help: "сигналы ошибок в системе",
	})
)

func main() {

	log.InitLogger()

	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())

	log.Logger.Info("server started")

	go http.ListenAndServe(":2121", nil)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	log.Logger.Info("\nserver shutdown")
}
