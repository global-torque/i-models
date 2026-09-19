package evmredemptiongrantcommands

import "testing"

func TestRedemptionGrantCommandIncludesSagaColumns(t *testing.T) {
	model := RedemptionGrantCommand{}
	want := map[string]bool{
		"prepare_command_uuid": true,
		"submission_response":  true,
		"updated_at":           true,
		"retry_at":             true,
	}
	for _, field := range model.Fields() {
		delete(want, field)
	}
	if len(want) != 0 {
		t.Fatalf("missing saga fields: %v", want)
	}
	json := model.ToJSON()
	for field := range map[string]struct{}{
		"prepare_command_uuid": {},
		"submission_response":  {},
		"updated_at":           {},
		"retry_at":             {},
	} {
		if _, ok := json[field]; !ok {
			t.Fatalf("ToJSON omitted saga field %q", field)
		}
	}
}
