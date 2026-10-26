// Package organizationbranding contains organization branding persistence models.
package organizationbranding

import (
	"github.com/global-torque/go-common/db/v2"
	"github.com/global-torque/go-common/orm/v2/pgtype"
)

const TableName = "organization_branding"

// OrganizationBranding is an object representing the organization_branding table.
type OrganizationBranding struct {
	SiteID      int                `db:"site_id" json:"site_id" yaml:"site_id"`
	Draft       map[string]any     `db:"draft" json:"draft" yaml:"draft"`
	Published   map[string]any     `db:"published" json:"published" yaml:"published"`
	PublishedAt pgtype.Timestamptz `db:"published_at" json:"published_at" yaml:"published_at"`
	UpdatedAt   pgtype.Timestamptz `db:"updated_at" json:"updated_at" yaml:"updated_at"`

	db db.Repository `db:"-" json:"-" yaml:"-"`
}

// Fields returns every persisted column required to hydrate organization branding.
func (branding OrganizationBranding) Fields() []string {
	return []string{"site_id", "draft", "published", "published_at", "updated_at"}
}

// Table returns the canonical organization branding table name.
func (branding OrganizationBranding) Table() string {
	return TableName
}

// GetID returns the branding primary key.
func (branding OrganizationBranding) GetID() any {
	return branding.SiteID
}

// SetID assigns the branding primary key.
func (branding *OrganizationBranding) SetID(id any) {
	branding.SiteID = id.(int)
}

// SetDB binds the branding to a repository.
func (branding *OrganizationBranding) SetDB(repo db.Repository) {
	branding.db = repo
}
