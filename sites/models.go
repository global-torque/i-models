// Package sites contains the persistence model for Django sites.
package sites

import "github.com/global-torque/go-common/db/v2"

const TableName = "django_site"

// DjangoSite is an object representing the django_site table.
type DjangoSite struct {
	ID     int    `db:"id" json:"id" yaml:"id"`
	Domain string `db:"domain" json:"domain" yaml:"domain"`
	Name   string `db:"name" json:"name" yaml:"name"`

	db db.Repository `db:"-" json:"-" yaml:"-"`
}

// Fields returns every persisted column required to hydrate a Site.
func (site DjangoSite) Fields() []string {
	return []string{"id", "domain", "name"}
}

// Table returns the canonical Django Site table name.
func (site DjangoSite) Table() string {
	return TableName
}

// GetID returns the Site primary key.
func (site DjangoSite) GetID() any {
	return site.ID
}

// SetID assigns the Site primary key.
func (site *DjangoSite) SetID(id any) {
	site.ID = id.(int)
}

// SetDB binds the Site to a repository.
func (site *DjangoSite) SetDB(repo db.Repository) {
	site.db = repo
}
