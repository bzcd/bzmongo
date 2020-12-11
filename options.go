package bzmongo

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type ClientOptions interface {
	ClientOptions() *options.ClientOptions
}

type Options struct {
	URI                    string        `yaml:"uri" json:"uri" toml:"uri"`
	ConnectTimeout         time.Duration `yaml:"connect_timeout" json:"connect_timeout" toml:"connect_timeout"`
	HeartbeatInterval      time.Duration `yaml:"heartbeat_interval" json:"heartbeat_interval" toml:"heartbeat_interval"`
	LocalThreshold         time.Duration `yaml:"local_threshold" json:"local_threshold" toml:"local_threshold"`
	MaxConnIdleTime        time.Duration `yaml:"max_conn_idletime" json:"max_conn_idletime" toml:"max_conn_idletime"`
	MaxPoolSize            uint64        `yaml:"max_pool_size" json:"max_pool_size" toml:"max_pool_size"`
	MinPoolSize            uint64        `yaml:"min_pool_size" json:"min_pool_size" toml:"min_pool_size"`
	ServerSelectionTimeout time.Duration `yaml:"server_selection_timeout" json:"server_selection_timeout" toml:"server_selection_timeout"`
	SocketTimeout          time.Duration `yaml:"socket_timeout" json:"socket_timeout" toml:"socket_timeout"`
}

func (o *Options) ClientOptions() *options.ClientOptions {
	opt := options.Client().ApplyURI(o.URI).
		SetLocalThreshold(o.LocalThreshold).
		SetMaxConnIdleTime(o.MaxConnIdleTime).
		SetMaxPoolSize(o.MaxPoolSize).
		SetConnectTimeout(o.ConnectTimeout).
		SetServerSelectionTimeout(o.ServerSelectionTimeout).
		SetSocketTimeout(o.SocketTimeout)

	if o.HeartbeatInterval > 0 {
		opt.SetHeartbeatInterval(o.HeartbeatInterval)
	}

	return opt
}
