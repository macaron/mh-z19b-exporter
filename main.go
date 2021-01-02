package main

import (
	"flag"
	"fmt"
	mhz19b "github.com/macaron/go-mh-z19b"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

var ppmGauge = promauto.NewGauge(prometheus.GaugeOpts{
	Namespace: "mhz19b",
	Name:      "co2_ppm",
	Help:      "Current ppm",
})

type Config struct {
	Path     string
	Port     string
	Device   string
	Interval int64
}

func recordMetrics(device string, interval int64) {
	go func() {
		for {
			ppm, err := mhz19b.Read(device)
			if err != nil {
				ppm = 0
			}

			ppmGauge.Set(float64(ppm))

			log.Println("Collection metrics")
			time.Sleep(time.Duration(interval) * time.Second)
		}
	}()
}

func main() {
	log.Println("starting mh-z19b-exporter")

	c := Config{
		Path:     "/metrics",
		Port:     "8080",
		Device:   "/dev/serial0",
		Interval: 60,
	}
	path := flag.String("path", c.Path, "Path for metrics")
	port := flag.String("port", c.Port, "Address for this exporter run")
	dev := flag.String("dev", c.Device, "MH-Z19B device")
	interval := flag.Int64("interval", c.Interval, "The frequency in seconds in which to gather data")
	flag.Parse()

	recordMetrics(*dev, *interval)

	http.Handle(*path, promhttp.Handler())
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *port), nil))
}
