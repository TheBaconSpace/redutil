package heartbeat

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

func New(param *HeartbeatParam) Heartbeat {
	if param.Updater == nil {
		param.Updater = SetTimeUpdater
	}

	h := simpleHeartbeat{
		Updater: param.Updater,
		pool:    param.Pool,

		Interval: param.Interval,
		keyspace: param.Keyspace,
		name:     param.Name,

		closer: make(chan interface{}),
		errors: make(chan error),
	}
	go h.heartbeat()

	return h
}

type simpleHeartbeat struct {
	Updater Updater

	Interval time.Duration
	keyspace string
	name     string

	pool   *redis.Pool
	closer chan interface{}
	errors chan error
}

func (h simpleHeartbeat) Pool() *redis.Pool    { return h.pool }
func (h simpleHeartbeat) Keyspace() string     { return h.keyspace }
func (h simpleHeartbeat) Name() string         { return h.name }
func (h simpleHeartbeat) Close()               { h.closer <- struct{}{} }
func (h simpleHeartbeat) Errors() <-chan error { return h.errors }

func (h simpleHeartbeat) heartbeat() {
	tick := time.NewTicker(h.Interval)
	defer tick.Stop()

	for {
		now := time.Now()

		select {
		case <-tick.C:
			if err := h.Updater(h, now); err != nil {
				h.errors <- err
			}
		case <-h.closer:
			return
		}
	}
}
