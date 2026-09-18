package evmvaultliquidityquarantines

import "github.com/global-torque/go-common/orm/v2/pgtype"

// VaultLiquidityQuarantine is operational safety evidence for a finalized
// shortfall. It never introduces a redemption business status.
type VaultLiquidityQuarantine struct {
	ID                   int64              `db:"id" json:"id"`
	OfferID              int                `db:"offer_id" json:"offer_id"`
	VaultContractID      int                `db:"vault_contract_id" json:"vault_contract_id"`
	ShortfallAssetsRaw   string             `db:"shortfall_assets_raw" json:"shortfall_assets_raw"`
	State                string             `db:"state" json:"state"`
	ReasonCode           string             `db:"reason_code" json:"reason_code"`
	RedemptionID         *int64             `db:"redemption_id" json:"redemption_id,omitempty"`
	ControllerAddress    *string            `db:"controller_address" json:"controller_address,omitempty"`
	AffectedEffectID     *int64             `db:"affected_effect_id" json:"affected_effect_id,omitempty"`
	FinalizedBlockNumber int64              `db:"finalized_block_number" json:"finalized_block_number"`
	FinalizedBlockHash   string             `db:"finalized_block_hash" json:"finalized_block_hash"`
	DetectedAt           pgtype.Timestamptz `db:"detected_at" json:"detected_at"`
	RecoveryEvidence     any                `db:"recovery_evidence" json:"recovery_evidence"`
	ClearedAt            pgtype.Timestamptz `db:"cleared_at" json:"cleared_at,omitempty"`
	CreatedAt            pgtype.Timestamptz `db:"created_at" json:"created_at"`
}

func (model VaultLiquidityQuarantine) ToJSON() map[string]any {
	return map[string]any{
		"id": model.ID, "offer_id": model.OfferID, "vault_contract_id": model.VaultContractID,
		"shortfall_assets_raw": model.ShortfallAssetsRaw, "state": model.State, "reason_code": model.ReasonCode, "redemption_id": model.RedemptionID,
		"controller_address": model.ControllerAddress, "affected_effect_id": model.AffectedEffectID,
		"finalized_block_number": model.FinalizedBlockNumber, "finalized_block_hash": model.FinalizedBlockHash,
		"detected_at": model.DetectedAt, "recovery_evidence": model.RecoveryEvidence,
		"cleared_at": model.ClearedAt, "created_at": model.CreatedAt,
	}
}
func (model VaultLiquidityQuarantine) Fields() []string {
	return []string{"id", "offer_id", "vault_contract_id", "shortfall_assets_raw", "state", "reason_code", "redemption_id", "controller_address", "affected_effect_id", "finalized_block_number", "finalized_block_hash", "detected_at", "recovery_evidence", "cleared_at", "created_at"}
}
func (model VaultLiquidityQuarantine) Table() string { return "evm_vault_liquidity_quarantines" }
func (model VaultLiquidityQuarantine) GetID() any    { return model.ID }
func (model *VaultLiquidityQuarantine) SetID(id any) { model.ID = id.(int64) }
