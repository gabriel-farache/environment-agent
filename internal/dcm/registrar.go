package dcm

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/dcm-project/environment-agent/api/v1alpha1"
)

var (
	ErrNonRetryable  = errors.New("non-retryable DCM error")
	ErrNotRegistered = errors.New("agent not registered with DCM")
)

// ServiceTypeLister returns the set of currently advertisable service types
// (backed by SPs in Ready or Unhealthy state — NOT Unavailable).
type ServiceTypeLister interface {
	AdvertisableServiceTypes() []string
}

// ConsumerLagProvider returns the current consumer lag for heartbeat payloads.
type ConsumerLagProvider interface {
	ConsumerLag() int64
}

// ResourceCapacityProvider optionally returns resource availability for registration.
// Returns nil when not available (REQ-DCM-030 is SHOULD, not MUST).
type ResourceCapacityProvider interface {
	ResourceCapacity() *v1alpha1.ResourceCapacity
}

// RegistrarConfig holds the configuration for DCM registration.
type RegistrarConfig struct {
	AgentName         string
	Environment       string
	Cost              string
	TopicName         string
	RegistrationURL   string
	InitialBackoff    time.Duration
	MaxBackoff        time.Duration
	HeartbeatInterval time.Duration
}

// Registrar handles DCM registration and heartbeat lifecycle.
// Unlike k8s SP's Registrar (which exits after first successful registration),
// this registrar stays alive indefinitely for periodic heartbeats and
// service-type update notifications (REQ-DCM-120, REQ-DCM-140).
type Registrar struct {
	done chan struct{}
}

// NewRegistrar creates a Registrar. Returns error if config is invalid
// (e.g., unparseable RegistrationURL). Mirrors k8s SP constructor pattern.
func NewRegistrar(
	cfg RegistrarConfig,
	lister ServiceTypeLister,
	lag ConsumerLagProvider,
	resources ResourceCapacityProvider,
	logger *slog.Logger,
) (*Registrar, error) {
	return &Registrar{done: make(chan struct{})}, nil
}

// Start begins the async registration + heartbeat loop. Non-blocking, idempotent (sync.Once).
func (r *Registrar) Start(_ context.Context) {}

// Done returns a channel closed when the registrar exits.
func (r *Registrar) Done() <-chan struct{} {
	return r.done
}

// NotifyServiceTypeChange signals that the advertisable service types may have changed.
func (r *Registrar) NotifyServiceTypeChange() {}

// AgentID returns the DCM-assigned agent ID, or ("", false) if not yet registered.
func (r *Registrar) AgentID() (string, bool) {
	return "", false
}
