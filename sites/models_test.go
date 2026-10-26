package sites

import (
	"reflect"
	"testing"
)

func TestDjangoSitePersistenceContract(t *testing.T) {
	t.Parallel()

	site := DjangoSite{}
	site.SetID(17)

	if site.ID != 17 {
		t.Fatalf("site ID = %d, want 17", site.ID)
	}
	if site.Table() != TableName {
		t.Fatalf("table = %q, want %q", site.Table(), TableName)
	}
	wantFields := []string{"id", "domain", "name"}
	if !reflect.DeepEqual(site.Fields(), wantFields) {
		t.Fatalf("fields = %#v, want %#v", site.Fields(), wantFields)
	}
}
