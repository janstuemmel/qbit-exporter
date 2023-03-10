package main

import (
	"log"
	"net/http"
	"os"
	"qbit-exporter/internal/collector"
	"qbit-exporter/internal/qbit"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	user := os.Getenv("QB_USER")
	pass := os.Getenv("QB_PASS")
	url := os.Getenv("QB_URL")

	qbit := qbit.NewClient(user, pass, url)
	col := collector.NewTorrentCollector(qbit)

	reg := prometheus.NewRegistry()
	reg.MustRegister(col)

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(":9897", nil))
}
