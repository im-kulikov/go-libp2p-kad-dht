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
)

// Metrics
var (
	GetValueMsgReceivedCount     = stats.Int64("libp2p.io/dht/kad/get_value_msg_received_count", "Total number of GetValue messages received", stats.UnitDimensionless)
	PutValueMsgReceivedCount     = stats.Int64("libp2p.io/dht/kad/put_value_msg_received_count", "Total number of PutValue messages received", stats.UnitDimensionless)
	FindNodeMsgReceivedCount     = stats.Int64("libp2p.io/dht/kad/find_node_msg_received_count", "Total number of FindNode messages received", stats.UnitDimensionless)
	AddProviderMsgReceivedCount  = stats.Int64("libp2p.io/dht/kad/add_provider_msg_received_count", "Total number of AddProvider messages received", stats.UnitDimensionless)
	GetProvidersMsgReceivedCount = stats.Int64("libp2p.io/dht/kad/get_providers_msg_received_count", "Total number of GetProviders messages received", stats.UnitDimensionless)
	PingMsgReceivedCount         = stats.Int64("libp2p.io/dht/kad/ping_msg_received_count", "Total number of Ping messages received", stats.UnitDimensionless)
)

// Views
var (
	GetValueMsgReceivedCountView = &view.View{
		Name:        "libp2p.io/dht/kad/get_value_msg_received_count",
		Measure:     GetValueMsgReceivedCount,
		TagKeys:     []tag.Key{},
		Aggregation: DefaultMessageCountDistribution,
	}
	PutValueMsgReceivedCountView = &view.View{
		Name:        "libp2p.io/dht/kad/put_value_msg_received_count",
		Measure:     PutValueMsgReceivedCount,
		TagKeys:     []tag.Key{},
		Aggregation: DefaultMessageCountDistribution,
	}
	FindNodeMsgReceivedCountView = &view.View{
		Name:        "libp2p.io/dht/kad/find_node_msg_received_count",
		Measure:     FindNodeMsgReceivedCount,
		TagKeys:     []tag.Key{},
		Aggregation: DefaultMessageCountDistribution,
	}
	AddProviderMsgReceivedCountView = &view.View{
		Name:        "libp2p.io/dht/kad/add_provider_msg_received_count",
		Measure:     AddProviderMsgReceivedCount,
		TagKeys:     []tag.Key{},
		Aggregation: DefaultMessageCountDistribution,
	}
	GetProvidersMsgReceivedCountView = &view.View{
		Name:        "libp2p.io/dht/kad/get_providers_msg_received_count",
		Measure:     GetProvidersMsgReceivedCount,
		TagKeys:     []tag.Key{},
		Aggregation: DefaultMessageCountDistribution,
	}
	PingMsgReceivedCountView = &view.View{
		Name:        "libp2p.io/dht/kad/ping_msg_received_count",
		Measure:     PingMsgReceivedCount,
		TagKeys:     []tag.Key{},
		Aggregation: DefaultMessageCountDistribution,
	}
)
