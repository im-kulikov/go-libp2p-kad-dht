package metrics

import (
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

var (
	defaultBytesDistribution        = view.Distribution(1024, 2048, 4096, 16384, 65536, 262144, 1048576, 4194304, 16777216, 67108864, 268435456, 1073741824, 4294967296)
	defaultMillisecondsDistribution = view.Distribution(0.01, 0.05, 0.1, 0.3, 0.6, 0.8, 1, 2, 3, 4, 5, 6, 8, 10, 13, 16, 20, 25, 30, 40, 50, 65, 80, 100, 130, 160, 200, 250, 300, 400, 500, 650, 800, 1000, 2000, 5000, 10000, 20000, 50000, 100000)
	defaultMessageCountDistribution = view.Distribution(1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536)
)

// Keys
var (
	// TODO: figure out some appropriate keys
	// KeyProtocol, _     = tag.NewKey("libp2p_protocol")
	KeyMessageType, _ = tag.NewKey("message_type")
	KeyPeerID, _      = tag.NewKey("peer_id")
	KeyRemoteID, _    = tag.NewKey("remotepeer_id")
	// KeyInstanceID identifies a dht instance by the pointer address.
	// Useful for differentiating between different dhts that have the same peer id.
	KeyInstanceID, _ = tag.NewKey("instance_id")
)

var (
	MRpcReceivedMessages = stats.Int64("libp2p.io/dht/kad/rpc_received_messages", "Total number of messages received per RPC", stats.UnitDimensionless)
	MRpcReceivedBytes    = stats.Int64("libp2p.io/dht/kad/rpc_received_bytes", "Total received bytes per RPC", stats.UnitBytes)
	MRpcLatencyMs        = stats.Float64("libp2p.io/dht/kad/rpc_latency_ms", "Latency per RPC", stats.UnitMilliseconds)
	MRpcSentMessages     = stats.Int64("libp2p.io/dht/kad/rpc_sent_messages", "Total number of messages sent per RPC", stats.UnitDimensionless)
	MRpcSentBytes        = stats.Int64("libp2p.io/dht/kad/rpc_sent_bytes", "Total sent bytes per RPC", stats.UnitBytes)
)

var (
	RpcReceivedMessagesView = &view.View{
		Measure:     MRpcReceivedMessages,
		TagKeys:     []tag.Key{KeyMessageType, KeyPeerID, KeyRemoteID, KeyInstanceID},
		Aggregation: defaultMessageCountDistribution,
	}
	RpcReceivedBytesView = &view.View{
		Measure:     MRpcReceivedBytes,
		TagKeys:     []tag.Key{KeyMessageType, KeyPeerID, KeyRemoteID, KeyInstanceID},
		Aggregation: defaultBytesDistribution,
	}
	RpcLatencyMsView = &view.View{
		Measure:     MRpcLatencyMs,
		TagKeys:     []tag.Key{KeyMessageType, KeyPeerID, KeyRemoteID, KeyInstanceID},
		Aggregation: defaultMillisecondsDistribution,
	}
	RpcSentMessagesView = &view.View{
		Measure:     MRpcSentMessages,
		TagKeys:     []tag.Key{KeyMessageType, KeyPeerID, KeyRemoteID, KeyInstanceID},
		Aggregation: defaultMessageCountDistribution,
	}
	RpcSentBytesView = &view.View{
		Measure:     MRpcSentBytes,
		TagKeys:     []tag.Key{KeyMessageType, KeyPeerID, KeyRemoteID, KeyInstanceID},
		Aggregation: defaultBytesDistribution,
	}
)
