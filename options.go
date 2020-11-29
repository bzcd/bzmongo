package bzmongo

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type ClientOptions interface {
	ClientOptions() *options.ClientOptions
}

type Options struct {
	URI             string        `yaml:"uri" json:"uri" toml:"uri"`
	MaxPoolSize     uint64        `yaml:"max_pool_size" json:"max_pool_size" toml:"max_pool_size"`
	MaxConnIdleTime time.Duration `yaml:"max_conn_idletime" json:"max_conn_idletime" toml:"max_conn_idletime"`
	LocalThreshold  time.Duration `yaml:"local_threshold" json:"local_threshold" toml:"local_threshold"`
}

func (o *Options) ClientOptions() *options.ClientOptions {
	return options.Client().ApplyURI(o.URI).
		SetLocalThreshold(o.LocalThreshold).
		SetMaxConnIdleTime(o.MaxConnIdleTime).
		SetMaxPoolSize(o.MaxPoolSize)
}
