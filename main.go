package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var totalReqCount = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "req_total",
	Help: "Count of total requests of this server",
})

var oneHandlerCount = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "oneHandlerCount",
}, []string{"status", "path"})

func main() {

	setTimeWeight()

	prometheus.MustRegister(totalReqCount, oneHandlerCount)

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		oneHandlerCount.WithLabelValues("200", r.URL.String()).Inc()
		totalReqCount.Add(1)
		mainHandler(w, r)
	})

	log.Println("start :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
