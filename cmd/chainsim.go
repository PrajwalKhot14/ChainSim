package main

import (
	"github.com/PrajwalKhot14/ChainSim/core"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {
	registry := core.NewRegistry()

	node := &core.Node{
		ID:        uuid.New().String(),
		Product:   core.PriceFeed,
		Status:    core.NodeStatusRunning,
		StartTime: time.Now(),
		SLIMetrics: core.SLIMetrics{
			Uptime:    99.9,
			LatencyMs: 120,
			ErrorRate: 0.01,
		},
	}

	registry.Register(node)

	for _, n := range registry.List() {
		fmt.Printf("Node %s (%s) - Uptime: %.2f%%\n", n.ID, n.Product, n.SLIMetrics.Uptime)
	}
}
