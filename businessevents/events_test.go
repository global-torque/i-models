package businessevents

import (
	"encoding/json"
	"os"
	"testing"
)

func TestOfferWireContract(t *testing.T) {
	body, err := os.ReadFile("testdata/offer-published.json")
	if err != nil {
		t.Fatal(err)
	}
	event, err := DecodeV1(body)
	if err != nil {
		t.Fatal(err)
	}
	site, err := event.OfferSiteID()
	if err != nil || site != 9007199254740993 {
		t.Fatalf("site: %d, %v", site, err)
	}
	if string(event.Data["min_investment"]) != "1.000000000000000001" {
		t.Fatal("decimal precision lost")
	}
	for _, action := range []string{OfferPublished, OfferUnpublished, OfferContentEdited} {
		t.Run(action, func(t *testing.T) {
			event.Action = action
			event.Data["site_id"] = json.RawMessage("null")
			encoded, err := json.Marshal(event)
			if err != nil {
				t.Fatal(err)
			}
			decoded, err := DecodeV1(encoded)
			if err != nil {
				t.Fatal(err)
			}
			if _, err := decoded.OfferSiteID(); err == nil {
				t.Fatal("null site must fail routing, not decoding")
			}
		})
	}
	for _, name := range []string{"missing changed fields", "nested legacy shape", "excluded text", "wrong identity"} {
		t.Run(name, func(t *testing.T) {
			event, err := DecodeV1(body)
			if err != nil {
				t.Fatal(err)
			}
			switch name {
			case "missing changed fields":
				event.ChangedFields = nil
			case "nested legacy shape":
				event.Data = map[string]json.RawMessage{"values": json.RawMessage(`{}`)}
			case "excluded text":
				event.Data["description"] = json.RawMessage(`"private text"`)
			case "wrong identity":
				event.ObjectRef = "43"
			}
			encoded, err := json.Marshal(event)
			if err != nil {
				t.Fatal(err)
			}
			if _, err := DecodeV1(encoded); err == nil {
				t.Fatal("accepted invalid trigger contract")
			}
		})
	}
}

func TestRetainedApplicationEnvelope(t *testing.T) {
	for _, body := range []string{
		`{"id":"6d0aaf23-d5ea-4ed5-b020-60fb9ba72155","version":1,"action":"offer.content-edited.v1","sender":"offer-api","object_name":"offer","object_ref":"42","object_id":42,"occurred_at":"2026-11-05T10:00:00Z","data":{"changed_fields":["title"],"values":{}}}`,
		`{"id":"6d0aaf23-d5ea-4ed5-b020-60fb9ba72155","version":1,"action":"other.created.v1","sender":"other-api","object_name":"other","object_ref":"external-key","occurred_at":"2026-11-05T10:00:00Z","data":{}}`,
	} {
		if _, err := DecodeV1([]byte(body)); err != nil {
			t.Fatal(err)
		}
	}
}
