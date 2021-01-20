/*
 * ArvanCloud CDN Services
 * API version: 4.0.0
 */
package models

type SuccessResponse struct {
	Message string       `json:"message,omitempty"`
	Data    *interface{} `json:"data,omitempty"`
}
