/*
 * ArvanCloud CDN Services
 * API version: 4.0.0
 */
package models

type MxRecord struct {
	Host     string `json:"host,omitempty"`
	Priority int32  `json:"priority,omitempty"`
}
