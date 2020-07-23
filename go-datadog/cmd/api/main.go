package main

import (
	"github.com/DataDog/datadog-go/statsd"
	"log"
	"math/rand"
	"time"
)

func main() {
	statsdd, err := statsd.New("127.0.0.1:8125")
	if err != nil {
		log.Fatalf("fatal level", err)
	}



	var i float64

	for {

		i += 1

		statsdd.Incr("example_metric.increment", []string{"environment:dev"}, 1)
		statsdd.Decr("example_metric.decrement", []string{"environment:dev"}, 1)
		statsdd.Count("example_metric.count", 2, []string{"environment:dev"}, 1)
		statsdd.Gauge("example_metric.gauge", i, []string{"environment:dev"}, 2)
		statsdd.Histogram("example_metric.histogram", float64(rand.Intn(20)), []string{"environment:dev"}, 1)

		time.Sleep(10 * time.Second)
	}

}
