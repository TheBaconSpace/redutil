package heartbeat

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreation(t *testing.T) {
	hb := New(&HeartbeatParam{
		Keyspace: "keyspace",
		Name:     "name",
		Interval: time.Duration(1 * time.Second),
	})
	defer hb.Close()

	assert.IsType(t, simpleHeartbeat{}, hb)
	assert.Equal(t, "keyspace", hb.Keyspace())
	assert.Equal(t, "name", hb.Name())
}

func TestClosing(t *testing.T) {
	t.Skip()
}

func TestUpdatesOnInterval(t *testing.T) {
	hb := New(&HeartbeatParam{
		Keyspace: "foo",
		Name:     "bar",
		Interval: time.Duration(500 * time.Millisecond),
	})
}
