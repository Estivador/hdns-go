package schema

import "time"

// Zone defines the schema of DNS zone information which may be included
// in responses.
type Zone struct {
	ID              string          `json:"id"`
	Created         time.Time       `json:"created"`
	Modified        time.Time       `json:"modified"`
	LegacyDnsHost   string          `json:"legacy_dns_host"`
	LegacyNs        []string        `json:"legacy_ns"`
	Name            string          `json:"name"`
	Ns              []string        `json:"ns"`
	Owner           string          `json:"owner"`
	Paused          bool            `json:"paused"`
	Permission      string          `json:"permission"`
	Project         string          `json:"project"`
	Registrar       string          `json:"registrar"`
	Status          string          `json:"status"`
	Ttl             uint64          `json:"ttl"`
	Verified        time.Time       `json:"verified"`
	RecordsCount    uint64          `json:"records_count"`
	IsSecondaryDns  bool            `json:"is_sedondary_dns"`
	TxtVerification TxtVerification `json:"txt_verification"`
}

// ZoneResponse defines the schema of a response containing zone information.
type ZoneResponse struct {
	Zone Zone `json:"zone"`
}
