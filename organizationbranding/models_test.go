package organizationbranding

import (
	"reflect"
	"testing"
)

func TestOrganizationBrandingPersistenceContract(t *testing.T) {
	t.Parallel()

	branding := OrganizationBranding{}
	branding.SetID(17)

	if branding.SiteID != 17 {
		t.Fatalf("site ID = %d, want 17", branding.SiteID)
	}
	if branding.Table() != TableName {
		t.Fatalf("table = %q, want %q", branding.Table(), TableName)
	}
	wantFields := []string{"site_id", "draft", "published", "published_at", "updated_at"}
	if !reflect.DeepEqual(branding.Fields(), wantFields) {
		t.Fatalf("fields = %#v, want %#v", branding.Fields(), wantFields)
	}
}
