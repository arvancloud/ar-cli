/*
 * ArvanCloud CDN Services
 * API version: 4.0.0
 */
package models

type CustomPage struct {
	StatusCode int32  `json:"status_code,omitempty"`
	Type       string `json:"type,omitempty"`
	Url        string `json:"url,omitempty"`
}
