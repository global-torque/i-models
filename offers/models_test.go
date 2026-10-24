package offers

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestOfferRegulatoryArraysPreserveOrderThroughJSONAndDatabaseScanning(t *testing.T) {
	t.Parallel()

	model := OfferOffer{
		RegType:           OfferRegTypesT{OfferRegTypeTUSRegS, OfferRegTypeTUSRegD506B},
		OfferJurisdiction: OfferJurisdictionsT{OfferJurisdictionTCH, OfferJurisdictionTUS},
		OfferingMode:      OfferingModesT{OfferingModeTOffshore, OfferingModeTPrivatePlacement},
	}

	encoded, err := json.Marshal(model)
	if err != nil {
		t.Fatalf("marshal offer: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(encoded, &payload); err != nil {
		t.Fatalf("unmarshal offer JSON: %v", err)
	}
	if got, want := payload["reg_type"], []any{"US_REG_S", "US_REG_D_506B"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("reg_type order changed: got %#v, want %#v", got, want)
	}

	var regTypes OfferRegTypesT
	if err := regTypes.Scan([]byte(`["US_REG_S","US_REG_D_506B"]`)); err != nil {
		t.Fatalf("scan reg_type: %v", err)
	}
	if !reflect.DeepEqual(regTypes, model.RegType) {
		t.Fatalf("scanned reg_type changed: got %#v, want %#v", regTypes, model.RegType)
	}
	value, err := model.OfferJurisdiction.Value()
	if err != nil {
		t.Fatalf("value offer_jurisdiction: %v", err)
	}
	if got, want := string(value.([]byte)), `["CH","US"]`; got != want {
		t.Fatalf("offer_jurisdiction order changed: got %q, want %q", got, want)
	}
}

func TestOfferShareFieldsUseExactJSONStrings(t *testing.T) {
	t.Parallel()

	model := OfferOffer{
		MinInvestment:    "0.000000000000000001",
		TotalShares:      "99999999999999999999.999999999999999999",
		PricePerShare:    "12345678901234567890.123456789012345678",
		SubscribedShares: "98765432109876543210.123456789012345678",
		ConfirmedShares:  "87654321098765432109.012345678901234567",
	}

	encoded, err := json.Marshal(model)
	if err != nil {
		t.Fatalf("marshal offer: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(encoded, &payload); err != nil {
		t.Fatalf("unmarshal offer JSON: %v", err)
	}

	want := map[string]string{
		"min_investment":    model.MinInvestment,
		"total_shares":      model.TotalShares,
		"price_per_share":   model.PricePerShare,
		"subscribed_shares": model.SubscribedShares,
		"confirmed_shares":  model.ConfirmedShares,
	}
	for field, expected := range want {
		if got := payload[field]; got != expected {
			t.Errorf("%s must be an exact JSON string: got %#v, want %q", field, got, expected)
		}
	}

	modelPayload := model.ToJSON()
	for field, expected := range want {
		if got := modelPayload[field]; got != expected {
			t.Errorf("ToJSON %s must retain its exact string: got %#v, want %q", field, got, expected)
		}
	}
}

func TestOfferMinInvestmentRejectsJSONNumber(t *testing.T) {
	t.Parallel()

	var model OfferOffer
	if err := json.Unmarshal([]byte(`{"min_investment":0.25}`), &model); err == nil {
		t.Fatal("min_investment must reject a JSON numeric token")
	}
}

func TestOfferFieldsIncludeVaultDeploymentSelectors(t *testing.T) {
	t.Parallel()

	fields := make(map[string]struct{})
	for _, field := range (OfferOffer{}).Fields() {
		fields[field] = struct{}{}
	}

	for _, field := range []string{"tokenization_engine", "fund_structure"} {
		if _, ok := fields[field]; !ok {
			t.Fatalf("OfferOffer.Fields must select %q for Vault deployment decisions", field)
		}
	}
}

func TestOfferModelExcludesDroppedTokenizationFramework(t *testing.T) {
	t.Parallel()

	if _, ok := reflect.TypeOf(OfferOffer{}).FieldByName("TokenizationFramework"); ok {
		t.Fatal("OfferOffer must not expose the dropped tokenization_framework column")
	}

	for _, field := range (OfferOffer{}).Fields() {
		if field == "tokenization_framework" {
			t.Fatal("OfferOffer.Fields must not select the dropped tokenization_framework column")
		}
	}
}
