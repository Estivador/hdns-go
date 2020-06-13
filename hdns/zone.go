package hdns

import (
	"time"
)

// Zone represents a zone in the Hetzner DNS
type Zone struct {
	ID              string
	Created         time.Time
	Modified        time.Time
	LegacyDNSHost   string
	LegacyNs        []string
	Name            string
	Ns              []string
	Owner           string
	Paused          bool
	Permission      string
	Project         string
	Registrar       string
	Status          string
	TTL             uint64
	Verified        time.Time
	RecordsCount    uint64
	IsSecondaryDNS  bool
	TxtVerification TxtVerification
}
