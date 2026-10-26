package offers

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestOfferRegulatoryArraysPreserveOrderThroughJSONAndDatabaseScanning(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		model      OfferOffer
		jsonField  string
		jsonValues []any
		scanJSON   string
		scan       func([]byte) (any, error)
		value      func() (any, error)
		want       any
		wantValue  string
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
			scanJSON:   `["US_REG_S","US_REG_D_506B"]`,
			scan: func(raw []byte) (any, error) {
				var values OfferRegTypesT
				if err := values.Scan(raw); err != nil {
					return nil, err
				}
				return values, nil
			},
			value: func() (any, error) {
				return OfferRegTypesT{
					OfferRegTypeTUSRegS,
					OfferRegTypeTUSRegD506B,
				}.Value()
			},
			want:      OfferRegTypesT{OfferRegTypeTUSRegS, OfferRegTypeTUSRegD506B},
			wantValue: `["US_REG_S","US_REG_D_506B"]`,
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
			scanJSON:   `["CH","US"]`,
			scan: func(raw []byte) (any, error) {
				var values OfferJurisdictionsT
				if err := values.Scan(raw); err != nil {
					return nil, err
				}
				return values, nil
			},
			value: func() (any, error) {
				return OfferJurisdictionsT{
					OfferJurisdictionTCH,
					OfferJurisdictionTUS,
				}.Value()
			},
			want:      OfferJurisdictionsT{OfferJurisdictionTCH, OfferJurisdictionTUS},
			wantValue: `["CH","US"]`,
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
			scanJSON:   `["offshore","private_placement"]`,
			scan: func(raw []byte) (any, error) {
				var values OfferingModesT
				if err := values.Scan(raw); err != nil {
					return nil, err
				}
				return values, nil
			},
			value: func() (any, error) {
				return OfferingModesT{
					OfferingModeTOffshore,
					OfferingModeTPrivatePlacement,
				}.Value()
			},
			want:      OfferingModesT{OfferingModeTOffshore, OfferingModeTPrivatePlacement},
			wantValue: `["offshore","private_placement"]`,
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

			scanned, err := test.scan([]byte(test.scanJSON))
			if err != nil {
				t.Fatalf("scan %s: %v", test.name, err)
			}
			if !reflect.DeepEqual(scanned, test.want) {
				t.Fatalf("scan order changed: got %#v, want %#v", scanned, test.want)
			}

			value, err := test.value()
			if err != nil {
				t.Fatalf("value %s: %v", test.name, err)
			}
			if got := string(value.([]byte)); got != test.wantValue {
				t.Fatalf("database value order changed: got %q, want %q", got, test.wantValue)
			}
		})
	}
}

func TestOfferRegTypesScanValidation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   any
		want    []string
		wantErr bool
	}{
		{
			name:  "unknown strings preserve order",
			input: []byte(`["future_route","legacy value"]`),
			want:  []string{"future_route", "legacy value"},
		},
		{name: "SQL nil", input: nil, wantErr: true},
		{name: "JSON null", input: []byte(`null`), wantErr: true},
		{name: "empty array", input: []byte(`[]`), wantErr: true},
		{name: "blank string", input: []byte(`[""]`), wantErr: true},
		{name: "whitespace string", input: []byte(`[" \t"]`), wantErr: true},
		{name: "duplicate string", input: []byte(`["future_route","future_route"]`), wantErr: true},
		{name: "scalar", input: []byte(`"future_route"`), wantErr: true},
		{name: "object", input: []byte(`{"value":"future_route"}`), wantErr: true},
		{name: "non-string element", input: []byte(`[42]`), wantErr: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var values OfferRegTypesT
			err := values.Scan(test.input)
			if test.wantErr {
				if err == nil {
					t.Fatal("OfferRegTypesT.Scan unexpectedly succeeded")
				}
				return
			}
			if err != nil {
				t.Fatalf("OfferRegTypesT.Scan failed: %v", err)
			}
			got := make([]string, len(values))
			for index, value := range values {
				got[index] = string(value)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("OfferRegTypesT.Scan changed values: got %#v, want %#v", got, test.want)
			}
		})
	}
}

func TestOfferRegulatoryArrayScansRejectMalformedValues(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		scan func(any) error
	}{
		{
			name: "offer jurisdiction",
			scan: func(value any) error {
				var values OfferJurisdictionsT
				return values.Scan(value)
			},
		},
		{
			name: "offering mode",
			scan: func(value any) error {
				var values OfferingModesT
				return values.Scan(value)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := test.scan([]byte(`[42]`)); err == nil {
				t.Fatal("malformed regulatory array unexpectedly scanned")
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
