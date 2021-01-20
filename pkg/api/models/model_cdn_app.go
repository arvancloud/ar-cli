/*
 * ArvanCloud CDN Services
 * API version: 4.0.0
 */
package models

import (
	"time"
)

type CdnApp struct {
	Id               string `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	Slug             string `json:"slug,omitempty"`
	ShortDescription string `json:"short_description,omitempty"`
	Description      string `json:"description,omitempty"`
	// absolute link to logo image
	Logo        string       `json:"logo,omitempty"`
	Vendor      string       `json:"vendor,omitempty"`
	InstallJson *InstallJson `json:"install_json,omitempty"`
	Status      string       `json:"status,omitempty"`
	CreatedAt   time.Time    `json:"created_at,omitempty"`
	UpdatedAt   time.Time    `json:"updated_at,omitempty"`
}
