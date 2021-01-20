/*
 * ArvanCloud CDN Services
 * API version: 4.0.0
 */
package responses

type InlineResponse422 struct {
	Status  bool   `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	// List of parameters and related errors
	Errors map[string][]string `json:"errors,omitempty"`
}
