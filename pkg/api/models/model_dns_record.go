/*
 * ArvanCloud CDN Services
 * API version: 4.0.0
 */
package models

import (
	"time"
)

type DnsRecord struct {
	Id   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
	// Value of the DNS record; depends on the record type
	Value         *OneOfDnsRecordValue   `json:"value,omitempty"`
	Ttl           int32                  `json:"ttl,omitempty"`
	Cloud         bool                   `json:"cloud,omitempty"`
	UpstreamHttps string                 `json:"upstream_https,omitempty"`
	IpFilterMode  *DnsRecordIpFilterMode `json:"ip_filter_mode,omitempty"`
	// This flag is deprecated in favor of is_protected flag
	CanDelete bool `json:"can_delete,omitempty"`
	// Protected records cannot be modified or deleted by user.
	IsProtected bool      `json:"is_protected,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
