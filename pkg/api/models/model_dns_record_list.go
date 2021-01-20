/*
 * ArvanCloud CDN Services
 * API version: 4.0.0
 */
package models

type DnsRecordList struct {
	Data  []DnsRecord             `json:"data,omitempty"`
	Links *PaginatedResponseLinks `json:"links,omitempty"`
	Meta  *PaginatedResponseMeta  `json:"meta,omitempty"`
}
