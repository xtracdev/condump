package main

import (
	"time"
	"fmt"
)

const (
	MetricsDumpInterval = 5 * time.Second
)

func main() {
	c := time.Tick(MetricsDumpInterval)
	for range c {
		//Signal self to dump metrics to stdout
		dumpHealth()
	}
}

func dumpHealth() {
	fmt.Printf("yeah\n")
}