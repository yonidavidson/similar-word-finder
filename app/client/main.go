package main

import (
	"encoding/json"
	"fmt"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

const RATE_PER_SECOND = 100
const DURATION_SECONDS = 4

func main() {

	rate := uint64(RATE_PER_SECOND) // per second
	duration := DURATION_SECONDS * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "http://web:8000/api/v1/similar?word=apple",
	})
	attacker := vegeta.NewAttacker()
	fmt.Printf("duration(s):%d, rate(s):%d \n\n", DURATION_SECONDS, RATE_PER_SECOND)
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration) {
		metrics.Add(res)
	}
	metrics.Close()
	json, _ := json.Marshal(metrics)
	fmt.Printf("metrics:\n %s\n", string(json))
}
