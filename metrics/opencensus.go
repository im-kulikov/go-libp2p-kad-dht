package metrics

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

var (
	MRpcReceivedMessages = stats.Int64("libp2p.io/dht/kad/rpc_received_messages", "Total number of messages received per RPC", stats.UnitDimensionless)
	MRpcReceivedBytes    = stats.Int64("libp2p.io/dht/kad/rpc_received_bytes", "Total received bytes per RPC", stats.UnitBytes)
	MRpcLatencyMs        = stats.Float64("libp2p.io/dht/kad/rpc_latency_ms", "Latency per RPC", stats.UnitMilliseconds)
	RpcReceivedMessages  = &view.View{
		Measure:     MRpcReceivedMessages,
		TagKeys:     []tag.Key{KeyMessageType},
		Aggregation: view.Count(),
	}
	RpcReceivedBytes = &view.View{
		Measure:     MRpcReceivedBytes,
		TagKeys:     []tag.Key{KeyMessageType},
		Aggregation: view.Distribution(0, 100, 1024),
	}
	RpcLatencyMs = &view.View{
		Measure:     MRpcLatencyMs,
		TagKeys:     []tag.Key{KeyMessageType},
		Aggregation: view.Distribution(0, 1, 5, 20, 50, 100, 150, 250, 750, 3000),
	}
)

// Maybe if the views are self describing, they should just all be in a big slice or map.
func AllViews() []*view.View {
	return append([]*view.View(nil), RpcLatencyMs, RpcReceivedBytes, RpcReceivedMessages)
}
