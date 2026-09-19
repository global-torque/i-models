package evmwalletoperationeffects

import "github.com/global-torque/go-common/orm/v2/pgtype"

// WalletOperationEffect is one canonical native transfer, ERC-20 transfer, or
// contract event observed in a transaction receipt.
type WalletOperationEffect struct {
	ID                       int64              `db:"id" json:"id"`
	OperationID              int                `db:"operation_id" json:"operation_id"`
	ChainAccountID           *int               `db:"chain_account_id" json:"chain_account_id,omitempty"`
	Chain                    string             `db:"chain" json:"chain"`
	EffectKind               EffectKindT        `db:"effect_kind" json:"effect_kind"`
	EffectIndex              int64              `db:"effect_index" json:"effect_index"`
	WalletAddress            string             `db:"wallet_address" json:"wallet_address"`
	Direction                *EffectDirectionT  `db:"direction" json:"direction,omitempty"`
	TokenAddress             *string            `db:"token_address" json:"token_address,omitempty"`
	TokenDecimals            *int               `db:"token_decimals" json:"token_decimals,omitempty"`
	AmountRaw                *string            `db:"amount_raw" json:"amount_raw,omitempty"`
	CounterpartyAddress      string             `db:"counterparty_address" json:"counterparty_address"`
	ContractAddress          *string            `db:"contract_address" json:"contract_address,omitempty"`
	EventSignature           *string            `db:"event_signature" json:"event_signature,omitempty"`
	EventTopics              any                `db:"event_topics" json:"event_topics,omitempty"`
	EventData                *string            `db:"event_data" json:"event_data,omitempty"`
	EventAssetsRaw           *string            `db:"event_assets_raw" json:"event_assets_raw,omitempty"`
	EventSharesRaw           *string            `db:"event_shares_raw" json:"event_shares_raw,omitempty"`
	ClaimBusinessAllocations any                `db:"claim_business_allocations" json:"claim_business_allocations,omitempty"`
	SubjectAddress           string             `db:"subject_address" json:"subject_address"`
	ReceiptGeneration        int                `db:"receipt_generation" json:"receipt_generation"`
	Canonical                bool               `db:"canonical" json:"canonical"`
	InvalidatedAt            pgtype.Timestamptz `db:"invalidated_at" json:"invalidated_at"`
	CreatedAt                pgtype.Timestamptz `db:"created_at" json:"created_at"`
}

func (model WalletOperationEffect) ToJSON() map[string]any {
	return map[string]any{
		"id":                         model.ID,
		"operation_id":               model.OperationID,
		"chain_account_id":           model.ChainAccountID,
		"chain":                      model.Chain,
		"effect_kind":                model.EffectKind,
		"effect_index":               model.EffectIndex,
		"wallet_address":             model.WalletAddress,
		"direction":                  model.Direction,
		"token_address":              model.TokenAddress,
		"token_decimals":             model.TokenDecimals,
		"amount_raw":                 model.AmountRaw,
		"counterparty_address":       model.CounterpartyAddress,
		"contract_address":           model.ContractAddress,
		"event_signature":            model.EventSignature,
		"event_topics":               model.EventTopics,
		"event_data":                 model.EventData,
		"event_assets_raw":           model.EventAssetsRaw,
		"event_shares_raw":           model.EventSharesRaw,
		"claim_business_allocations": model.ClaimBusinessAllocations,
		"subject_address":            model.SubjectAddress,
		"receipt_generation":         model.ReceiptGeneration,
		"canonical":                  model.Canonical,
		"invalidated_at":             model.InvalidatedAt,
		"created_at":                 model.CreatedAt,
	}
}

func (model WalletOperationEffect) Fields() []string {
	return []string{
		"id", "operation_id", "chain_account_id", "chain", "effect_kind",
		"effect_index", "wallet_address", "direction", "token_address",
		"token_decimals", "amount_raw", "counterparty_address", "contract_address",
		"event_signature", "event_topics", "event_data", "event_assets_raw",
		"event_shares_raw", "claim_business_allocations", "subject_address", "receipt_generation",
		"canonical", "invalidated_at", "created_at",
	}
}

func (model WalletOperationEffect) Table() string { return "evm_wallet_operation_effects" }
func (model WalletOperationEffect) GetID() any    { return model.ID }
func (model *WalletOperationEffect) SetID(id any) { model.ID = id.(int64) }
