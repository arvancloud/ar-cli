/*
 * ArvanCloud CDN Services
 * API version: 4.0.0
 */
package models

type SrvRecord struct {
	Target   string `json:"target,omitempty"`
	Port     int32  `json:"port,omitempty"`
	Weight   int32  `json:"weight,omitempty"`
	Priority int32  `json:"priority,omitempty"`
}
