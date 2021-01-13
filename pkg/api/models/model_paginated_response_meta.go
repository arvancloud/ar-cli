/*
 * ArvanCloud CDN Services
 * API version: 4.0.0
 */
package models

type PaginatedResponseMeta struct {
	CurrentPage int32  `json:"current_page,omitempty"`
	From        int32  `json:"from,omitempty"`
	LastPage    int32  `json:"last_page,omitempty"`
	Path        string `json:"path,omitempty"`
	PerPage     int32  `json:"per_page,omitempty"`
	To          int32  `json:"to,omitempty"`
	Total       int32  `json:"total,omitempty"`
}
