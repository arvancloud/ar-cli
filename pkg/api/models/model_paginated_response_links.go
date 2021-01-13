/*
 * ArvanCloud CDN Services
 * API version: 4.0.0
 */
package models

type PaginatedResponseLinks struct {
	First string `json:"first"`
	Last  string `json:"last,omitempty"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
}
