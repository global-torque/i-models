package evmredemptionapprovalpreflightevidence

import "github.com/global-torque/go-common/orm/v2/pgtype"

// RedemptionApprovalPreflightEvidence is an immutable finalized checkpoint
// snapshot used to fence one manager approval transaction.
type RedemptionApprovalPreflightEvidence struct {
	ID                           int64              `db:"id" json:"id"`
	RedemptionID                 int64              `db:"redemption_id" json:"redemption_id"`
	IdempotencyKey               string             `db:"idempotency_key" json:"idempotency_key"`
	PayloadHash                  string             `db:"payload_hash" json:"payload_hash"`
	GrantID                      int64              `db:"grant_id" json:"grant_id"`
	GrantVersion                 int64              `db:"grant_version" json:"grant_version"`
	GrantScope                   any                `db:"grant_scope" json:"grant_scope"`
	GrantScopeHash               string             `db:"grant_scope_hash" json:"grant_scope_hash"`
	VaultContractID              int                `db:"vault_contract_id" json:"vault_contract_id"`
	RequestControllerAddress     string             `db:"request_controller_address" json:"request_controller_address"`
	FinalizedBlockNumber         int64              `db:"finalized_block_number" json:"finalized_block_number"`
	FinalizedBlockHash           string             `db:"finalized_block_hash" json:"finalized_block_hash"`
	ScannerFinalityWatermark     int64              `db:"scanner_finality_watermark" json:"scanner_finality_watermark"`
	ScannerFinalityBlockHash     *string            `db:"scanner_finality_block_hash" json:"scanner_finality_block_hash,omitempty"`
	InvestorShareBalanceRaw      string             `db:"investor_share_balance_raw" json:"investor_share_balance_raw"`
	FinalizedFreeAssetsRaw       string             `db:"finalized_free_assets_raw" json:"finalized_free_assets_raw"`
	UnrepresentedReservationsRaw string             `db:"unrepresented_reservations_raw" json:"unrepresented_reservations_raw"`
	ActiveSetDigest              string             `db:"active_set_digest" json:"active_set_digest"`
	ExpectedPricePerShareRaw     string             `db:"expected_price_per_share_raw" json:"expected_price_per_share_raw"`
	ExpectedPayoutAssetsRaw      string             `db:"expected_payout_assets_raw" json:"expected_payout_assets_raw"`
	ObservedAt                   pgtype.Timestamptz `db:"observed_at" json:"observed_at"`
	ExpiresAt                    pgtype.Timestamptz `db:"expires_at" json:"expires_at"`
	CreatedAt                    pgtype.Timestamptz `db:"created_at" json:"created_at"`
}

func (model RedemptionApprovalPreflightEvidence) ToJSON() map[string]any {
	return map[string]any{
		"id": model.ID, "redemption_id": model.RedemptionID, "idempotency_key": model.IdempotencyKey,
		"payload_hash": model.PayloadHash, "grant_id": model.GrantID, "grant_version": model.GrantVersion,
		"grant_scope": model.GrantScope, "grant_scope_hash": model.GrantScopeHash,
		"vault_contract_id": model.VaultContractID, "request_controller_address": model.RequestControllerAddress,
		"finalized_block_number": model.FinalizedBlockNumber, "finalized_block_hash": model.FinalizedBlockHash,
		"scanner_finality_watermark": model.ScannerFinalityWatermark, "scanner_finality_block_hash": model.ScannerFinalityBlockHash,
		"investor_share_balance_raw": model.InvestorShareBalanceRaw, "finalized_free_assets_raw": model.FinalizedFreeAssetsRaw,
		"unrepresented_reservations_raw": model.UnrepresentedReservationsRaw, "active_set_digest": model.ActiveSetDigest,
		"expected_price_per_share_raw": model.ExpectedPricePerShareRaw, "expected_payout_assets_raw": model.ExpectedPayoutAssetsRaw,
		"observed_at": model.ObservedAt, "expires_at": model.ExpiresAt, "created_at": model.CreatedAt,
	}
}

func (model RedemptionApprovalPreflightEvidence) Fields() []string {
	return []string{"id", "redemption_id", "idempotency_key", "payload_hash", "grant_id", "grant_version", "grant_scope", "grant_scope_hash", "vault_contract_id", "request_controller_address", "finalized_block_number", "finalized_block_hash", "scanner_finality_watermark", "scanner_finality_block_hash", "investor_share_balance_raw", "finalized_free_assets_raw", "unrepresented_reservations_raw", "active_set_digest", "expected_price_per_share_raw", "expected_payout_assets_raw", "observed_at", "expires_at", "created_at"}
}
func (model RedemptionApprovalPreflightEvidence) Table() string {
	return "evm_redemption_approval_preflight_evidence"
}
func (model RedemptionApprovalPreflightEvidence) GetID() any    { return model.ID }
func (model *RedemptionApprovalPreflightEvidence) SetID(id any) { model.ID = id.(int64) }
