package heartbeat

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

type Heartbeat interface {
	Pool() *redis.Pool
	Keyspace() string
	Name() string

	Errors() <-chan error
	Close()
}

type HeartbeatParam struct {
	Pool *redis.Pool

	Keyspace string
	Name     string

	Updater  Updater
	Interval time.Duration
}
