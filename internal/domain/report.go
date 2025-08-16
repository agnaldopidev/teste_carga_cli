package domain

import "time"

type Report struct {
	URL           string
	Duration      time.Duration
	TotalRequests int
	StatusCount   map[int]int
}
