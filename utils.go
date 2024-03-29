package delaywheel

import (
	"time"
)

// truncate returns the result of rounding x toward zero to a multiple of m.
// If m <= 0, Truncate returns x unchanged.
func truncate(x, m int64) int64 {
	if m <= 0 {
		return x
	}
	return x - x%m
}

// timeToMs returns an integer number, which represents t in milliseconds.
func timeToMs(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// msToTime returns the UTC time corresponding to the given Unix time,
// t milliseconds since January 1, 1970 UTC.
func msToTime(t int64) time.Time {
	return time.Unix(0, t*int64(time.Millisecond)).UTC()
}

func refreshTimer(t *time.Timer, now int64, expireTime int64) (isRefresh bool) {
	// Ensure the timer is fully reset.
	if !t.Stop() {
		select {
		case <-t.C:
		default:
		}
	}

	deltaMs := expireTime - now
	if deltaMs <= 0 {
		return false
	}
	t.Reset(time.Duration(deltaMs) * time.Millisecond)
	return true
}
