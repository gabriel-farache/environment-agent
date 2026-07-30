// Package backoff provides utilities for calculating exponential backoff durations with jitter.
package backoff

import "time"

// CalculateBackoff returns the deterministic component: min(initial × 2^attempt, max).
func CalculateBackoff(initial, max time.Duration, attempt int) time.Duration { //nolint:revive // stub
	return -1
}

// ApplyJitter applies full jitter: uniform random in [0, calculated].
// randFn allows deterministic testing; production callers pass math/rand.Float64.
func ApplyJitter(calculated time.Duration, randFn func() float64) time.Duration { //nolint:revive // stub
	return -1
}
