package clock

import "time"

func Since(c Clock, t time.Time) time.Duration {
	if c == nil {
		return time.Since(t)
	}
	return c.Now().Sub(t)
}
