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

func (d *DnsRecord) ValueString() string {
	switch d.Type {
	case "a":
		return d.Value.ARecord.Ip
	case "aaaa":
		return d.Value.AaaaRecord.Ip
	case "ns":
		return d.Value.NsRecord.Host
	case "txt":
		return d.Value.TxtRecord.Text
	case "cname":
		return d.Value.CnameRecord.Host
	case "mx":
		return d.Value.MxRecord.Host
	case "srv":
		return d.Value.SrvRecord.Target
	case "spf":
		return d.Value.SpfRecord.Text
	case "dkim":
		return d.Value.DkimRecord.Text
	case "aname":
		return d.Value.AnameRecord.Location
	case "ptr":
		return d.Value.PtrRecord.Domain
	default:
		return ""
	}
}
