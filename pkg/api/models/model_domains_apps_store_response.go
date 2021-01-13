/*
 * ArvanCloud CDN Services
 * API version: 4.0.0
 */
package models

type DomainsAppsStoreResponse struct {
	Message string        `json:"message,omitempty"`
	Data    *DomainCdnApp `json:"data,omitempty"`
}
