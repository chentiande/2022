package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var cpuTemp prometheus.Gauge

func init() {
	//prometheus.MustRegister(cpuTemp)
}
func main() {

	cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "chentiande",
		Help: "Current temperature of the CPU.",
	})
	go settemp()
	prometheus.MustRegister(cpuTemp)

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func settemp() {
	for {
		time.Sleep(time.Second * 3)
		rand.Seed(time.Now().Unix())
		cpuTemp.Set(rand.Float64())
	}

}
