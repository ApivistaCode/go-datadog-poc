package main

import (
	"github.com/DataDog/datadog-go/statsd"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"log"
	"math/rand"
	"time"
)

func main() {
	statsdd, err := statsd.New("127.0.0.1:8125")
	if err != nil {
		log.Fatalf("fatal level", err)
	}

	tracer.Start(tracer.WithAgentAddr("datadog-agent:8126"), tracer.WithAnalytics(true))
	defer tracer.Stop()

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
