package evmredemptionapprovalpreflightentries

// RedemptionApprovalPreflightEntry is one checkpoint-bound active obligation.
// Effect and receipt metadata are retained so coverage is independently
// revalidated instead of inferred from a mutable projection.
type RedemptionApprovalPreflightEntry struct {
	EvidenceID                   int64     `db:"evidence_id" json:"evidence_id"`
	RedemptionID                 int64     `db:"redemption_id" json:"redemption_id"`
	PayoutAssetsRaw              string    `db:"payout_assets_raw" json:"payout_assets_raw"`
	ClaimedAssetsRaw             string    `db:"claimed_assets_raw" json:"claimed_assets_raw"`
	RemainingAssetsRaw           string    `db:"remaining_assets_raw" json:"remaining_assets_raw"`
	FulfillmentEffectID          *int64    `db:"fulfillment_effect_id" json:"fulfillment_effect_id,omitempty"`
	RequestEffectID              *int64    `db:"request_effect_id" json:"request_effect_id,omitempty"`
	RequestEffectSignature       *string   `db:"request_effect_signature" json:"request_effect_signature,omitempty"`
	RequestReceiptGeneration     *int      `db:"request_receipt_generation" json:"request_receipt_generation,omitempty"`
	RequestBlockNumber           *int64    `db:"request_block_number" json:"request_block_number,omitempty"`
	RequestBlockHash             *string   `db:"request_block_hash" json:"request_block_hash,omitempty"`
	FulfillmentEffectSignature   *string   `db:"fulfillment_effect_signature" json:"fulfillment_effect_signature,omitempty"`
	FulfillmentReceiptGeneration *int      `db:"fulfillment_receipt_generation" json:"fulfillment_receipt_generation,omitempty"`
	FulfillmentBlockNumber       *int64    `db:"fulfillment_block_number" json:"fulfillment_block_number,omitempty"`
	FulfillmentBlockHash         *string   `db:"fulfillment_block_hash" json:"fulfillment_block_hash,omitempty"`
	CheckpointBlockNumber        *int64    `db:"checkpoint_block_number" json:"checkpoint_block_number,omitempty"`
	CheckpointBlockHash          *string   `db:"checkpoint_block_hash" json:"checkpoint_block_hash,omitempty"`
	CheckpointFact               any       `db:"checkpoint_fact" json:"checkpoint_fact"`
	Coverage                     CoverageT `db:"coverage" json:"coverage"`
}

func (model RedemptionApprovalPreflightEntry) ToJSON() map[string]any {
	return map[string]any{
		"evidence_id": model.EvidenceID, "redemption_id": model.RedemptionID,
		"payout_assets_raw": model.PayoutAssetsRaw, "claimed_assets_raw": model.ClaimedAssetsRaw,
		"remaining_assets_raw": model.RemainingAssetsRaw, "fulfillment_effect_id": model.FulfillmentEffectID,
		"request_effect_id": model.RequestEffectID, "request_effect_signature": model.RequestEffectSignature,
		"request_receipt_generation": model.RequestReceiptGeneration, "request_block_number": model.RequestBlockNumber,
		"request_block_hash": model.RequestBlockHash, "fulfillment_effect_signature": model.FulfillmentEffectSignature,
		"fulfillment_receipt_generation": model.FulfillmentReceiptGeneration, "fulfillment_block_number": model.FulfillmentBlockNumber,
		"fulfillment_block_hash": model.FulfillmentBlockHash, "checkpoint_block_number": model.CheckpointBlockNumber,
		"checkpoint_block_hash": model.CheckpointBlockHash, "checkpoint_fact": model.CheckpointFact, "coverage": model.Coverage,
	}
}

func (model RedemptionApprovalPreflightEntry) Fields() []string {
	return []string{"evidence_id", "redemption_id", "payout_assets_raw", "claimed_assets_raw", "remaining_assets_raw", "fulfillment_effect_id", "request_effect_id", "request_effect_signature", "request_receipt_generation", "request_block_number", "request_block_hash", "fulfillment_effect_signature", "fulfillment_receipt_generation", "fulfillment_block_number", "fulfillment_block_hash", "checkpoint_block_number", "checkpoint_block_hash", "checkpoint_fact", "coverage"}
}
func (model RedemptionApprovalPreflightEntry) Table() string {
	return "evm_redemption_approval_preflight_entries"
}
func (model RedemptionApprovalPreflightEntry) GetID() any    { return model.RedemptionID }
func (model *RedemptionApprovalPreflightEntry) SetID(id any) { model.RedemptionID = id.(int64) }
