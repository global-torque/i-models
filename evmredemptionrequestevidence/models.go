package evmredemptionrequestevidence

import "github.com/global-torque/go-common/orm/v2/pgtype"

// RedemptionRequestEvidence is a checkpoint read captured before request
// calldata is signed. It becomes authoritative only when the operation,
// calldata hash, and cancellation fence are committed together.
type RedemptionRequestEvidence struct {
	ID                              int64              `db:"id" json:"id"`
	RedemptionID                    int64              `db:"redemption_id" json:"redemption_id"`
	VaultContractID                 int                `db:"vault_contract_id" json:"vault_contract_id"`
	RequestControllerAddress        string             `db:"request_controller_address" json:"request_controller_address"`
	RequestOperationID              *int64             `db:"request_operation_id" json:"request_operation_id,omitempty"`
	RequestCalldataHash             *string            `db:"request_calldata_hash" json:"request_calldata_hash,omitempty"`
	CancellationFence               *int64             `db:"cancellation_fence" json:"cancellation_fence,omitempty"`
	ObservedBlockNumber             int64              `db:"observed_block_number" json:"observed_block_number"`
	ObservedBlockHash               string             `db:"observed_block_hash" json:"observed_block_hash"`
	ObservedShareBalanceRaw         string             `db:"observed_share_balance_raw" json:"observed_share_balance_raw"`
	ReservedShareAmountRaw          string             `db:"reserved_share_amount_raw" json:"reserved_share_amount_raw"`
	RemainingReservedShareAmountRaw string             `db:"remaining_reserved_share_amount_raw" json:"remaining_reserved_share_amount_raw"`
	ReservedAssetAmountRaw          string             `db:"reserved_asset_amount_raw" json:"reserved_asset_amount_raw"`
	RemainingReservedAssetAmountRaw string             `db:"remaining_reserved_asset_amount_raw" json:"remaining_reserved_asset_amount_raw"`
	ActiveSetDigest                 string             `db:"active_set_digest" json:"active_set_digest"`
	Provisional                     bool               `db:"provisional" json:"provisional"`
	ExpiresAt                       pgtype.Timestamptz `db:"expires_at" json:"expires_at"`
	CreatedAt                       pgtype.Timestamptz `db:"created_at" json:"created_at"`
}

func (model RedemptionRequestEvidence) ToJSON() map[string]any {
	return map[string]any{
		"id": model.ID, "redemption_id": model.RedemptionID, "vault_contract_id": model.VaultContractID,
		"request_controller_address": model.RequestControllerAddress, "request_operation_id": model.RequestOperationID,
		"request_calldata_hash": model.RequestCalldataHash, "cancellation_fence": model.CancellationFence,
		"observed_block_number": model.ObservedBlockNumber, "observed_block_hash": model.ObservedBlockHash,
		"observed_share_balance_raw": model.ObservedShareBalanceRaw, "reserved_share_amount_raw": model.ReservedShareAmountRaw,
		"remaining_reserved_share_amount_raw": model.RemainingReservedShareAmountRaw, "reserved_asset_amount_raw": model.ReservedAssetAmountRaw,
		"remaining_reserved_asset_amount_raw": model.RemainingReservedAssetAmountRaw, "active_set_digest": model.ActiveSetDigest,
		"provisional": model.Provisional, "expires_at": model.ExpiresAt, "created_at": model.CreatedAt,
	}
}

func (model RedemptionRequestEvidence) Fields() []string {
	return []string{"id", "redemption_id", "vault_contract_id", "request_controller_address", "request_operation_id", "request_calldata_hash", "cancellation_fence", "observed_block_number", "observed_block_hash", "observed_share_balance_raw", "reserved_share_amount_raw", "remaining_reserved_share_amount_raw", "reserved_asset_amount_raw", "remaining_reserved_asset_amount_raw", "active_set_digest", "provisional", "expires_at", "created_at"}
}
func (model RedemptionRequestEvidence) Table() string { return "evm_redemption_request_evidence" }
func (model RedemptionRequestEvidence) GetID() any    { return model.ID }
func (model *RedemptionRequestEvidence) SetID(id any) { model.ID = id.(int64) }
