/*
 * ArvanCloud CDN Services
 * API version: 4.0.0
 */
package models

type DnsRecordResponse struct {
	Message string     `json:"message,omitempty"`
	Data    *DnsRecord `json:"data,omitempty"`
}
