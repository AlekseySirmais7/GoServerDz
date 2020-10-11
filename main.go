package main

import (
	"./handlers"
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

	handlers.SetTimeWeight()

	prometheus.MustRegister(totalReqCount, oneHandlerCount)

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		oneHandlerCount.WithLabelValues("200", r.URL.String()).Inc()
		totalReqCount.Add(1)
		handlers.MainHandler(w, r)
	})

	log.Println("start :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
