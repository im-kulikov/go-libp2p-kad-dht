package dht

import (
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

// Copied from https://github.com/census-instrumentation/opencensus-go/blob/master/plugin/ocgrpc/stats_common.go
var (
	DefaultBytesDistribution        = view.Distribution(1024, 2048, 4096, 16384, 65536, 262144, 1048576, 4194304, 16777216, 67108864, 268435456, 1073741824, 4294967296)
	DefaultMillisecondsDistribution = view.Distribution(0.01, 0.05, 0.1, 0.3, 0.6, 0.8, 1, 2, 3, 4, 5, 6, 8, 10, 13, 16, 20, 25, 30, 40, 50, 65, 80, 100, 130, 160, 200, 250, 300, 400, 500, 650, 800, 1000, 2000, 5000, 10000, 20000, 50000, 100000)
	DefaultMessageCountDistribution = view.Distribution(1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536)
)

// Keys
var (
	// TODO: figure out some appropriate keys
	// TODO: add a key for each query type to be able to tag every query made
	// KeyProtocol, _     = tag.NewKey("libp2p_protocol")
	// KeyPeerID, _       = tag.NewKey("libp2p_peer_id")
	// KeyRemotePeerID, _ = tag.NewKey("libp2p_remote_peer_id")
	KeyMessageType, _ = tag.NewKey("message_type")
)

// Metrics
var (
	ReceivedMessagesPerRPC = stats.Int64("libp2p.io/dht/kad/received_messages_per_rpc", "Total number of messages received per RPC", stats.UnitDimensionless)
	ReceivedBytesPerRPC    = stats.Int64("libp2p.io/dht/kad/received_bytes_per_rpc", "Total received bytes per RPC", stats.UnitBytes)
	LatencyPerRPC          = stats.Float64("libp2p.io/dht/kad/latency_per_rpc", "Latency per RPC", stats.UnitMilliseconds)
)

// Views
var (
	ReceivedMessagesPerRPCView = &view.View{
		Name:        "libp2p.io/dht/kad/received_messages_per_rpc",
		Measure:     ReceivedMessagesPerRPC,
		TagKeys:     []tag.Key{KeyMessageType},
		Aggregation: DefaultMessageCountDistribution,
	}

	ReceivedBytesPerRPCView = &view.View{
		Name:        "libp2p.io/dht/kad/received_bytes_per_rpc",
		Measure:     ReceivedBytesPerRPC,
		TagKeys:     []tag.Key{KeyMessageType},
		Aggregation: DefaultBytesDistribution,
	}

	LatencyPerRPCView = &view.View{
		Name:        "libp2p.io/dht/kad/latency_per_rpc",
		Measure:     LatencyPerRPC,
		TagKeys:     []tag.Key{KeyMessageType},
		Aggregation: DefaultMillisecondsDistribution,
	}
)
