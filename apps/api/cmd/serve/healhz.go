package serve

import "time"

var startTime = time.Now()

type HealthCheckResonse struct {
	Message   string `json:"message"`
	Uptime    uint64 `json:"uptime"`
	Timestamp uint64 `json:"timestamp"`
}

func getHealthCheckResonse() *HealthCheckResonse {
	elapsedTime := time.Since(startTime)
	uptime := uint64(elapsedTime.Seconds())
	timestamp := uint64(time.Now().UnixNano() / int64(time.Millisecond))

	return &HealthCheckResonse{Message: "Not dead yet", Uptime: uptime, Timestamp: timestamp}
}
