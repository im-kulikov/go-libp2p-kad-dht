package dht

import (
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

// Keys
var (
	// TODO: figure out some appropriate keys
	// KeyProtocol, _     = tag.NewKey("libp2p_protocol")
	KeyMessageType, _ = tag.NewKey("message_type")
)

type NodeMetrics struct {
	rpcReceivedMessages *stats.Int64Measure
	rpcReceivedBytes    *stats.Int64Measure
	rpcLatencyMs        *stats.Float64Measure
	RpcReceivedMessages *view.View
	RpcReceivedBytes    *view.View
	RpcLatencyMs        *view.View
}

func newNodeMetrics() (nm NodeMetrics) {
	nm.rpcReceivedMessages = stats.Int64("libp2p.io/dht/kad/rpc_received_messages", "Total number of messages received per RPC", stats.UnitDimensionless)
	nm.rpcReceivedBytes = stats.Int64("libp2p.io/dht/kad/rpc_received_bytes", "Total received bytes per RPC", stats.UnitBytes)
	nm.rpcLatencyMs = stats.Float64("libp2p.io/dht/kad/rpc_latency_ms", "Latency per RPC", stats.UnitMilliseconds)
	nm.RpcReceivedMessages = &view.View{
		Measure:     nm.rpcReceivedMessages,
		TagKeys:     []tag.Key{KeyMessageType},
		Aggregation: view.Count(),
	}
	nm.RpcReceivedBytes = &view.View{
		Measure:     nm.rpcReceivedBytes,
		TagKeys:     []tag.Key{KeyMessageType},
		Aggregation: view.Distribution(0, 100, 1024),
	}
	nm.RpcLatencyMs = &view.View{
		Measure:     nm.rpcLatencyMs,
		TagKeys:     []tag.Key{KeyMessageType},
		Aggregation: view.Distribution(0, 1, 5, 20, 50, 100, 150, 250, 750, 3000),
	}
	return
}

// Maybe if the views are self describing, they should just all be in a big slice or map.
func (nm NodeMetrics) AllViews() []*view.View {
	return append([]*view.View(nil), nm.RpcLatencyMs, nm.RpcReceivedBytes, nm.RpcReceivedMessages)
}
