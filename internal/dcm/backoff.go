package dcm

import "time"

// ParseRetryAfter parses the Retry-After header value.
// Returns the duration to wait and whether parsing succeeded.
func ParseRetryAfter(value string, now time.Time) (time.Duration, bool) {
	return -1, false
}
