package main

import (
	"flag"

	"github.com/renatafborges/stress-test-golang/internal/processor"
)

var (
	url         *string
	requests    *int
	concurrency *int64
)

func init() {
	url = flag.String("url", "http://google.com", "url to make request")
	requests = flag.Int("requests", 10, "number of requests to send to url")
	concurrency = flag.Int64("concurrency", 10, "number of concurrent executions")
	flag.Parse()
}

func main() {
	processor.MakeStressTest(*url, *requests, *concurrency)
}
