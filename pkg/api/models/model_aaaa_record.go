/*
 * ArvanCloud CDN Services
 * API version: 4.0.0
 */
package models

type AaaaRecord struct {
	Ip     string `json:"ip,omitempty"`
	Port   int32  `json:"port,omitempty"`
	Weight int32  `json:"weight,omitempty"`
	// ISO 3166 alpha-2 country code
	Country string `json:"country,omitempty"`
}
