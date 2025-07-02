package core

import "time"

type NodeStatus string

const (
	NodeStatusRunning NodeStatus = "RUNNING"
	NodeStatusStopped NodeStatus = "STOPPED"
)

type ChainlinkProduct string

const (
	PriceFeed ChainlinkProduct = "PRICE_FEED"
	VRF       ChainlinkProduct = "VRF"
	CCIP      ChainlinkProduct = "CCIP"
)

type Node struct {
	ID         string
	Product    ChainlinkProduct
	Status     NodeStatus
	StartTime  time.Time
	SLIMetrics SLIMetrics
}

type SLIMetrics struct {
	Uptime    float64
	LatencyMs float64
	ErrorRate float64
}
