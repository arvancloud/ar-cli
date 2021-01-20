/*
 * ArvanCloud CDN Services
 * API version: 4.0.0
 */
package models

type Dnssec struct {
	Enabled bool   `json:"enabled,omitempty"`
	Ds      string `json:"ds,omitempty"`
}
