/*
 * ArvanCloud CDN Services
 * API version: 4.0.0
 */
package models

import (
	"time"
)

type DomainCdnApp struct {
	Id            string      `json:"id,omitempty"`
	DomainId      string      `json:"domain_id,omitempty"`
	ApplicationId string      `json:"application_id,omitempty"`
	Active        bool        `json:"active,omitempty"`
	Options       *AppOptions `json:"options,omitempty"`
	CreatedAt     time.Time   `json:"created_at,omitempty"`
	UpdatedAt     time.Time   `json:"updated_at,omitempty"`
}
