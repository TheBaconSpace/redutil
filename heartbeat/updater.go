package heartbeat

import "time"

type Updater func(h Heartbeat, now time.Time) error

var (
	SetTimeUpdater Updater = func(h Heartbeat, now time.Time) error {
		cnx := h.Pool().Get()
		defer cnx.Close()

		if _, err := cnx.Do("HSET", h.Keyspace(), h.Name(),
			now.String()); err != nil {

			return err
		}

		return nil
	}
)
