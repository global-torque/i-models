// Package businessevents defines the shared business journal wire contract.
package businessevents

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Database-authored offer actions and producer identity.
const (
	OfferPublished     = "offer.published.v1"
	OfferUnpublished   = "offer.unpublished.v1"
	OfferContentEdited = "offer.content-edited.v1"
	OfferTriggerSender = "postgres-offer-publication-trigger"
)

// OfferSnapshotColumns is the flat database snapshot allowlist, in schema order.
const OfferSnapshotColumns = `id user_id image_link_id legal_representative_id platform_representative_id site_id
data_room_group_id name legal_name slug title seo_title seo_description website state city
address1 address2 zip_code country entity_id ticker esign_id twitter facebook github
instagram telegram linkedin mastodon status security_type instrument_classification
offering_currency tokenization_engine tokenization_model fund_structure min_investment
valuation total_shares price_per_share subscribed_shares confirmed_shares start_at close_at data security_info
reg_type offer_jurisdiction offering_mode investor_eligibility marketed_jurisdictions
regulated_activities approved_at created_at updated_at`

// EventV1 is a business journal envelope. Raw data preserves numeric precision and SQL nulls.
// Retained application-authored events can omit ChangedFields and use their original data shape.
type EventV1 struct {
	ID            string                     `json:"id"`
	Version       int                        `json:"version"`
	Action        string                     `json:"action"`
	Sender        string                     `json:"sender"`
	ObjectName    string                     `json:"object_name"`
	ObjectRef     string                     `json:"object_ref"`
	ObjectID      int64                      `json:"object_id"`
	OccurredAt    time.Time                  `json:"occurred_at"`
	ChangedFields []string                   `json:"changed_fields,omitempty"`
	Data          map[string]json.RawMessage `json:"data"`
}

// DecodeV1 decodes a retained generic envelope or the strict database offer snapshot.
// A null site_id is valid storage data. OfferSiteID reports the routing failure separately.
func DecodeV1(body []byte) (EventV1, error) {
	var event EventV1
	if err := json.Unmarshal(body, &event); err != nil {
		return event, fmt.Errorf("decode business event: %w", err)
	}
	identity, identityErr := uuid.Parse(event.ID)
	if identityErr != nil || identity == uuid.Nil || identity.String() != event.ID || event.Version != 1 || event.Action == "" || event.Sender == "" ||
		event.ObjectName == "" || event.ObjectRef == "" || event.ObjectID < 0 || event.OccurredAt.IsZero() || event.Data == nil {
		return event, fmt.Errorf("invalid business event V1 envelope")
	}
	if event.Sender != OfferTriggerSender {
		return event, nil
	}
	switch event.Action {
	case OfferPublished, OfferUnpublished, OfferContentEdited:
	default:
		return event, fmt.Errorf("unsupported offer trigger action %q", event.Action)
	}
	if event.ObjectID <= 0 || event.ObjectName != "offer" || event.ObjectRef != strconv.FormatInt(event.ObjectID, 10) {
		return event, fmt.Errorf("inconsistent offer event identity")
	}
	if len(event.ChangedFields) == 0 || !sort.StringsAreSorted(event.ChangedFields) {
		return event, fmt.Errorf("offer changed_fields must be nonempty and sorted")
	}
	for i, field := range event.ChangedFields {
		if field == "" || (i > 0 && event.ChangedFields[i-1] == field) {
			return event, fmt.Errorf("invalid offer changed_fields")
		}
	}
	columns := strings.Fields(OfferSnapshotColumns)
	if len(event.Data) != len(columns) {
		return event, fmt.Errorf("invalid offer snapshot columns")
	}
	for _, column := range columns {
		if _, ok := event.Data[column]; !ok {
			return event, fmt.Errorf("missing offer snapshot column %q", column)
		}
	}
	var id int64
	if err := json.Unmarshal(event.Data["id"], &id); err != nil || id != event.ObjectID {
		return event, fmt.Errorf("inconsistent offer snapshot id")
	}
	return event, nil
}

// OfferSiteID returns the positive site routing key without converting JSON numbers to float64.
func (event EventV1) OfferSiteID() (int64, error) {
	var id *int64
	if err := json.Unmarshal(event.Data["site_id"], &id); err != nil {
		return 0, fmt.Errorf("invalid offer site_id: %w", err)
	}
	if id == nil || *id <= 0 {
		return 0, fmt.Errorf("offer site_id must be a positive integer")
	}
	return *id, nil
}
