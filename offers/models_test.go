package offers

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestOfferRegulatoryArraysPreserveOrderThroughJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		model      OfferOffer
		jsonField  string
		jsonValues []any
	}{
		{
			name: "reg_type",
			model: OfferOffer{
				RegType: OfferRegTypesT{
					OfferRegTypeTUSRegS,
					OfferRegTypeTUSRegD506B,
				},
			},
			jsonField:  "reg_type",
			jsonValues: []any{"US_REG_S", "US_REG_D_506B"},
		},
		{
			name: "offer_jurisdiction",
			model: OfferOffer{
				OfferJurisdiction: OfferJurisdictionsT{
					OfferJurisdictionTCH,
					OfferJurisdictionTUS,
				},
			},
			jsonField:  "offer_jurisdiction",
			jsonValues: []any{"CH", "US"},
		},
		{
			name: "offering_mode",
			model: OfferOffer{
				OfferingMode: OfferingModesT{
					OfferingModeTOffshore,
					OfferingModeTPrivatePlacement,
				},
			},
			jsonField:  "offering_mode",
			jsonValues: []any{"offshore", "private_placement"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			encoded, err := json.Marshal(test.model)
			if err != nil {
				t.Fatalf("marshal offer: %v", err)
			}
			var payload map[string]any
			if err := json.Unmarshal(encoded, &payload); err != nil {
				t.Fatalf("unmarshal offer JSON: %v", err)
			}
			if got := payload[test.jsonField]; !reflect.DeepEqual(got, test.jsonValues) {
				t.Fatalf("JSON order changed: got %#v, want %#v", got, test.jsonValues)
			}
		})
	}
}

type offerRegulatoryArrayAdapter struct {
	name  string
	field string
	scan  func(any) ([]string, error)
	value func([]string) (any, error)
}

func offerRegulatoryStrings[T ~string](values []T) []string {
	strings := make([]string, len(values))
	for index, value := range values {
		strings[index] = string(value)
	}
	return strings
}

func offerRegulatoryValues[T ~string](values []string) []T {
	if values == nil {
		return nil
	}

	result := make([]T, len(values))
	for index, value := range values {
		result[index] = T(value)
	}
	return result
}

func offerRegulatoryArrayAdapters() []offerRegulatoryArrayAdapter {
	return []offerRegulatoryArrayAdapter{
		{
			name: "reg_type", field: "offer regulatory types",
			scan: func(value any) ([]string, error) {
				var values OfferRegTypesT
				if err := values.Scan(value); err != nil {
					return nil, err
				}
				return offerRegulatoryStrings(values), nil
			},
			value: func(values []string) (any, error) {
				return OfferRegTypesT(offerRegulatoryValues[OfferRegTypeT](values)).Value()
			},
		},
		{
			name: "offer_jurisdiction", field: "offer jurisdictions",
			scan: func(value any) ([]string, error) {
				var values OfferJurisdictionsT
				if err := values.Scan(value); err != nil {
					return nil, err
				}
				return offerRegulatoryStrings(values), nil
			},
			value: func(values []string) (any, error) {
				return OfferJurisdictionsT(offerRegulatoryValues[OfferJurisdictionT](values)).Value()
			},
		},
		{
			name: "offering_mode", field: "offering modes",
			scan: func(value any) ([]string, error) {
				var values OfferingModesT
				if err := values.Scan(value); err != nil {
					return nil, err
				}
				return offerRegulatoryStrings(values), nil
			},
			value: func(values []string) (any, error) {
				return OfferingModesT(offerRegulatoryValues[OfferingModeT](values)).Value()
			},
		},
	}
}

func TestOfferRegulatoryArraysValidateStructure(t *testing.T) {
	t.Parallel()

	validValues := []string{" future_route ", "legacy value"}
	validJSON := `[" future_route ","legacy value"]`
	scanInvalid := []struct {
		name  string
		input any
	}{
		{name: "SQL nil", input: nil},
		{name: "unsupported SQL type", input: 42},
		{name: "invalid JSON", input: []byte(`not-json`)},
		{name: "JSON null", input: []byte(`null`)},
		{name: "empty array", input: []byte(`[]`)},
		{name: "blank string", input: []byte(`[""]`)},
		{name: "whitespace string", input: []byte(`[" \t"]`)},
		{name: "duplicate string", input: []byte(`["future_route","future_route"]`)},
		{name: "scalar", input: []byte(`"future_route"`)},
		{name: "object", input: []byte(`{"value":"future_route"}`)},
		{name: "non-string element", input: []byte(`[42]`)},
	}
	valueInvalid := []struct {
		name  string
		input []string
	}{
		{name: "nil", input: nil},
		{name: "empty array", input: []string{}},
		{name: "blank string", input: []string{""}},
		{name: "whitespace string", input: []string{" \t"}},
		{name: "duplicate string", input: []string{"future_route", "future_route"}},
	}

	for _, adapter := range offerRegulatoryArrayAdapters() {
		adapter := adapter
		t.Run(adapter.name, func(t *testing.T) {
			t.Run("valid ordered unknown strings", func(t *testing.T) {
				scanned, err := adapter.scan([]byte(validJSON))
				if err != nil {
					t.Fatalf("scan: %v", err)
				}
				if !reflect.DeepEqual(scanned, validValues) {
					t.Fatalf("scan changed values: got %#v, want %#v", scanned, validValues)
				}

				encoded, err := adapter.value(validValues)
				if err != nil {
					t.Fatalf("value: %v", err)
				}
				got, ok := encoded.([]byte)
				if !ok {
					t.Fatalf("value returned %T, want []byte", encoded)
				}
				if string(got) != validJSON {
					t.Fatalf("value changed order or text: got %q, want %q", got, validJSON)
				}
			})

			for _, test := range scanInvalid {
				test := test
				t.Run("scan/"+test.name, func(t *testing.T) {
					if _, err := adapter.scan(test.input); err == nil {
						t.Fatal("scan unexpectedly succeeded")
					}
				})
			}

			for _, test := range valueInvalid {
				test := test
				t.Run("value/"+test.name, func(t *testing.T) {
					_, err := adapter.value(test.input)
					if err == nil {
						t.Fatal("value unexpectedly succeeded")
					}
					if !strings.Contains(err.Error(), adapter.field) {
						t.Fatalf("value error lost field context: %v", err)
					}
				})
			}
		})
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
