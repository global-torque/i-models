package investments

import (
	"github.com/global-torque/go-common/orm/v2/pgtype"
)

// InvestmentInvestment is an object representing the database table.
type InvestmentInvestment struct {
	ID                              int                  `db:"id" json:"id" yaml:"id"`
	UserID                          *int                 `db:"user_id" json:"user_id,omitempty" yaml:"user_id,omitempty"`
	CreatedBy                       int                  `db:"created_by" json:"created_by" yaml:"created_by"`
	OfferID                         *int                 `db:"offer_id" json:"offer_id,omitempty" yaml:"offer_id,omitempty"`
	ProfileID                       *int                 `db:"profile_id" json:"profile_id,omitempty" yaml:"profile_id,omitempty"`
	SiteID                          *int                 `db:"site_id" json:"site_id,omitempty" yaml:"site_id,omitempty"`
	Amount                          *string              `db:"amount" json:"amount" yaml:"amount"`
	PricePerShare                   string               `db:"price_per_share" json:"price_per_share" yaml:"price_per_share"`
	NumberOfShares                  *string              `db:"number_of_shares" json:"number_of_shares" yaml:"number_of_shares"`
	PaymentType                     PaymentT             `db:"payment_type" json:"payment_type" yaml:"payment_type"`
	EscrowType                      EscrowT              `db:"escrow_type" json:"escrow_type" yaml:"escrow_type"`
	FundingType                     FundingT             `db:"funding_type" json:"funding_type" yaml:"funding_type"`
	FundingStatus                   FundingS             `db:"funding_status" json:"funding_status" yaml:"funding_status"`
	Status                          InvestmentT          `db:"status" json:"status" yaml:"status"`
	PrevStatus                      InvestmentT          `db:"prev_status" json:"prev_status" yaml:"prev_status"`
	Step                            InvestmentStepT      `db:"step" json:"step" yaml:"step"`
	Commission                      string               `db:"commission" json:"commission" yaml:"commission"`
	CancelationReason               string               `db:"cancelation_reason" json:"cancelation_reason" yaml:"cancelation_reason"`
	EntityID                        *string              `db:"entity_id" json:"entity_id,omitempty" yaml:"entity_id,omitempty"`
	TransactionRef                  *string              `db:"transaction_ref" json:"transaction_ref,omitempty" yaml:"transaction_ref,omitempty"`
	EscrowData                      map[string]any       `db:"escrow_data" json:"escrow_data" yaml:"escrow_data"`
	SignatureData                   map[string]any       `db:"signature_data" json:"signature_data" yaml:"signature_data"`
	PaymentData                     map[string]any       `db:"payment_data" json:"payment_data" yaml:"payment_data"`
	CanceledAt                      pgtype.Timestamptz   `db:"canceled_at" json:"canceled_at" yaml:"canceled_at"`
	SubmitedAt                      pgtype.Timestamptz   `db:"submited_at" json:"submited_at" yaml:"submited_at"`
	ClosedAt                        pgtype.Timestamptz   `db:"closed_at" json:"closed_at" yaml:"closed_at"`
	CreatedAt                       pgtype.Timestamptz   `db:"created_at" json:"created_at" yaml:"created_at"`
	UpdatedAt                       pgtype.Timestamptz   `db:"updated_at" json:"updated_at" yaml:"updated_at"`
	VaultContractID                 *int                 `db:"vault_contract_id" json:"vault_contract_id,omitempty" yaml:"vault_contract_id,omitempty"`
	RequestControllerChainAccountID *int                 `db:"request_controller_chain_account_id" json:"request_controller_chain_account_id,omitempty" yaml:"request_controller_chain_account_id,omitempty"`
	RequestControllerAddress        *string              `db:"request_controller_address" json:"request_controller_address,omitempty" yaml:"request_controller_address,omitempty"`
	VaultRequestOrigin              VaultRequestOriginT  `db:"vault_request_origin" json:"vault_request_origin" yaml:"vault_request_origin"`
	VaultRequestEffectID            *int64               `db:"vault_request_effect_id" json:"vault_request_effect_id,omitempty" yaml:"vault_request_effect_id,omitempty"`
	CustodyChainAccountID           *int                 `db:"custody_chain_account_id" json:"custody_chain_account_id,omitempty" yaml:"custody_chain_account_id,omitempty"`
	CustodyFundedAt                 pgtype.Timestamptz   `db:"custody_funded_at" json:"custody_funded_at,omitempty" yaml:"custody_funded_at,omitempty"`
	CustodyRefundLockedAt           pgtype.Timestamptz   `db:"custody_refund_locked_at" json:"custody_refund_locked_at,omitempty" yaml:"custody_refund_locked_at,omitempty"`
	CustodyReleasedAt               pgtype.Timestamptz   `db:"custody_released_at" json:"custody_released_at,omitempty" yaml:"custody_released_at,omitempty"`
	CustodyRefundedAt               pgtype.Timestamptz   `db:"custody_refunded_at" json:"custody_refunded_at,omitempty" yaml:"custody_refunded_at,omitempty"`
	CustodyTransitionVersion        int64                `db:"custody_transition_version" json:"custody_transition_version" yaml:"custody_transition_version"`
	DepositPriceSource              *DepositPriceSourceT `db:"deposit_price_source" json:"deposit_price_source,omitempty" yaml:"deposit_price_source,omitempty"`
	DepositPricingStatus            VaultPricingStatusT  `db:"deposit_pricing_status" json:"deposit_pricing_status" yaml:"deposit_pricing_status"`
	DepositPriceUSDCRaw             *string              `db:"deposit_price_usdc_raw" json:"deposit_price_usdc_raw,omitempty" yaml:"deposit_price_usdc_raw,omitempty"`
	DepositNAVRecordID              *int64               `db:"deposit_nav_record_id" json:"deposit_nav_record_id,omitempty" yaml:"deposit_nav_record_id,omitempty"`
	DepositNAVVersion               *int64               `db:"deposit_nav_version" json:"deposit_nav_version,omitempty" yaml:"deposit_nav_version,omitempty"`
	DepositPricedAt                 pgtype.Timestamptz   `db:"deposit_priced_at" json:"deposit_priced_at,omitempty" yaml:"deposit_priced_at,omitempty"`
	AssetAmountRaw                  *string              `db:"asset_amount_raw" json:"asset_amount_raw,omitempty" yaml:"asset_amount_raw,omitempty"`
	ShareAmountRaw                  *string              `db:"share_amount_raw" json:"share_amount_raw,omitempty" yaml:"share_amount_raw,omitempty"`
	PendingAssetsRaw                string               `db:"pending_assets_raw" json:"pending_assets_raw" yaml:"pending_assets_raw"`
	ClaimableAssetsRaw              string               `db:"claimable_assets_raw" json:"claimable_assets_raw" yaml:"claimable_assets_raw"`
	ClaimableSharesRaw              string               `db:"claimable_shares_raw" json:"claimable_shares_raw" yaml:"claimable_shares_raw"`
	ClaimedAssetsRaw                string               `db:"claimed_assets_raw" json:"claimed_assets_raw" yaml:"claimed_assets_raw"`
	ClaimedSharesRaw                string               `db:"claimed_shares_raw" json:"claimed_shares_raw" yaml:"claimed_shares_raw"`
	VaultTransitionVersion          int64                `db:"vault_transition_version" json:"vault_transition_version" yaml:"vault_transition_version"`
	VaultRequestLockedAt            pgtype.Timestamptz   `db:"vault_request_locked_at" json:"vault_request_locked_at,omitempty" yaml:"vault_request_locked_at,omitempty"`
	VaultRequestedAt                pgtype.Timestamptz   `db:"vault_requested_at" json:"vault_requested_at,omitempty" yaml:"vault_requested_at,omitempty"`
	VaultClaimableAt                pgtype.Timestamptz   `db:"vault_claimable_at" json:"vault_claimable_at,omitempty" yaml:"vault_claimable_at,omitempty"`
	VaultClaimedAt                  pgtype.Timestamptz   `db:"vault_claimed_at" json:"vault_claimed_at,omitempty" yaml:"vault_claimed_at,omitempty"`
}

// IsFundingTypeWire temporary fix for https://github.com/Joker/jade/issues/59
func (model InvestmentInvestment) IsFundingTypeWire() bool {
	return model.FundingType == "wire"
}

func (model InvestmentInvestment) ToJSON() map[string]any {
	return map[string]any{
		"id":                                  model.ID,
		"user_id":                             model.UserID,
		"created_by":                          model.CreatedBy,
		"offer_id":                            model.OfferID,
		"profile_id":                          model.ProfileID,
		"site_id":                             model.SiteID,
		"amount":                              model.Amount,
		"price_per_share":                     model.PricePerShare,
		"number_of_shares":                    model.NumberOfShares,
		"payment_type":                        model.PaymentType,
		"escrow_type":                         model.EscrowType,
		"funding_type":                        model.FundingType,
		"funding_status":                      model.FundingStatus,
		"status":                              model.Status,
		"prev_status":                         model.PrevStatus,
		"step":                                model.Step,
		"commission":                          model.Commission,
		"cancelation_reason":                  model.CancelationReason,
		"entity_id":                           model.EntityID,
		"transaction_ref":                     model.TransactionRef,
		"escrow_data":                         model.EscrowData,
		"signature_data":                      model.SignatureData,
		"payment_data":                        model.PaymentData,
		"canceled_at":                         model.CanceledAt,
		"submited_at":                         model.SubmitedAt,
		"closed_at":                           model.ClosedAt,
		"created_at":                          model.CreatedAt,
		"updated_at":                          model.UpdatedAt,
		"vault_contract_id":                   model.VaultContractID,
		"request_controller_chain_account_id": model.RequestControllerChainAccountID,
		"request_controller_address":          model.RequestControllerAddress,
		"vault_request_origin":                model.VaultRequestOrigin,
		"vault_request_effect_id":             model.VaultRequestEffectID,
		"custody_chain_account_id":            model.CustodyChainAccountID,
		"custody_funded_at":                   model.CustodyFundedAt,
		"custody_refund_locked_at":            model.CustodyRefundLockedAt,
		"custody_released_at":                 model.CustodyReleasedAt,
		"custody_refunded_at":                 model.CustodyRefundedAt,
		"custody_transition_version":          model.CustodyTransitionVersion,
		"deposit_price_source":                model.DepositPriceSource,
		"deposit_pricing_status":              model.DepositPricingStatus,
		"deposit_price_usdc_raw":              model.DepositPriceUSDCRaw,
		"deposit_nav_record_id":               model.DepositNAVRecordID,
		"deposit_nav_version":                 model.DepositNAVVersion,
		"deposit_priced_at":                   model.DepositPricedAt,
		"asset_amount_raw":                    model.AssetAmountRaw,
		"share_amount_raw":                    model.ShareAmountRaw,
		"pending_assets_raw":                  model.PendingAssetsRaw,
		"claimable_assets_raw":                model.ClaimableAssetsRaw,
		"claimable_shares_raw":                model.ClaimableSharesRaw,
		"claimed_assets_raw":                  model.ClaimedAssetsRaw,
		"claimed_shares_raw":                  model.ClaimedSharesRaw,
		"vault_transition_version":            model.VaultTransitionVersion,
		"vault_request_locked_at":             model.VaultRequestLockedAt,
		"vault_requested_at":                  model.VaultRequestedAt,
		"vault_claimable_at":                  model.VaultClaimableAt,
		"vault_claimed_at":                    model.VaultClaimedAt,
	}
}

func (model InvestmentInvestment) Fields() []string {
	return []string{
		"id", "user_id", "created_by", "offer_id", "profile_id", "site_id", "amount",
		"price_per_share", "number_of_shares", "payment_type", "escrow_type",
		"funding_type", "funding_status", "status", "prev_status", "step",
		"commission", "cancelation_reason", "entity_id", "transaction_ref",
		"escrow_data", "signature_data", "payment_data", "canceled_at",
		"submited_at", "closed_at", "created_at", "updated_at",
		"vault_contract_id", "request_controller_chain_account_id",
		"request_controller_address", "vault_request_origin", "vault_request_effect_id",
		"custody_chain_account_id", "custody_funded_at", "custody_refund_locked_at",
		"custody_released_at", "custody_refunded_at", "custody_transition_version",
		"deposit_price_source", "deposit_pricing_status", "deposit_price_usdc_raw",
		"deposit_nav_record_id", "deposit_nav_version", "deposit_priced_at",
		"asset_amount_raw", "share_amount_raw", "pending_assets_raw",
		"claimable_assets_raw", "claimable_shares_raw", "claimed_assets_raw",
		"claimed_shares_raw", "vault_transition_version", "vault_request_locked_at",
		"vault_requested_at", "vault_claimable_at", "vault_claimed_at",
	}
}

func (model InvestmentInvestment) Table() string {
	return "investment_investments"
}

func (model InvestmentInvestment) GetID() any {
	return model.ID
}

func (model *InvestmentInvestment) SetID(id any) {
	model.ID = id.(int)
}

// InvestmentRedemption is an exact-value projection of an asynchronous ERC-7540 redemption.
type InvestmentRedemption struct {
	ID                              int64               `db:"id" json:"id"`
	OfferID                         int                 `db:"offer_id" json:"offer_id"`
	ProfileID                       int                 `db:"profile_id" json:"profile_id"`
	InvestmentID                    *int                `db:"investment_id" json:"investment_id,omitempty"`
	VaultContractID                 *int                `db:"vault_contract_id" json:"vault_contract_id,omitempty"`
	RequestControllerChainAccountID *int                `db:"request_controller_chain_account_id" json:"request_controller_chain_account_id,omitempty"`
	RequestControllerAddress        *string             `db:"request_controller_address" json:"request_controller_address,omitempty"`
	VaultRequestOrigin              VaultRequestOriginT `db:"vault_request_origin" json:"vault_request_origin"`
	VaultRequestEffectID            *int64              `db:"vault_request_effect_id" json:"vault_request_effect_id,omitempty"`
	Status                          RedemptionStatusT   `db:"status" json:"status"`
	IdempotencyKey                  string              `db:"idempotency_key" json:"idempotency_key"`
	EstimatedNAVRecordID            *int64              `db:"estimated_nav_record_id" json:"estimated_nav_record_id,omitempty"`
	EstimatedNAVUSDCRaw             *string             `db:"estimated_nav_usdc_raw" json:"estimated_nav_usdc_raw,omitempty"`
	EstimatedNAVVersion             *int64              `db:"estimated_nav_version" json:"estimated_nav_version,omitempty"`
	EstimatedNAVValuationBlock      *int64              `db:"estimated_nav_valuation_block_number" json:"estimated_nav_valuation_block_number,omitempty"`
	EstimatedNAVValuationAsOf       pgtype.Timestamptz  `db:"estimated_nav_valuation_as_of" json:"estimated_nav_valuation_as_of,omitempty"`
	EstimatedAssetAmountRaw         *string             `db:"estimated_asset_amount_raw" json:"estimated_asset_amount_raw,omitempty"`
	EstimatedAt                     pgtype.Timestamptz  `db:"estimated_at" json:"estimated_at,omitempty"`
	DealingCutoffBlockNumber        *int64              `db:"dealing_cutoff_block_number" json:"dealing_cutoff_block_number,omitempty"`
	DealingCutoffAt                 pgtype.Timestamptz  `db:"dealing_cutoff_at" json:"dealing_cutoff_at,omitempty"`
	NAVRecordID                     *int64              `db:"nav_record_id" json:"nav_record_id,omitempty"`
	NAVUSDCRaw                      *string             `db:"nav_usdc_raw" json:"nav_usdc_raw,omitempty"`
	NAVVersion                      *int64              `db:"nav_version" json:"nav_version,omitempty"`
	NAVValuationBlockNumber         *int64              `db:"nav_valuation_block_number" json:"nav_valuation_block_number,omitempty"`
	NAVValuationAsOf                pgtype.Timestamptz  `db:"nav_valuation_as_of" json:"nav_valuation_as_of,omitempty"`
	DealingPriceUSDCRaw             *string             `db:"dealing_price_usdc_raw" json:"dealing_price_usdc_raw,omitempty"`
	PricingSource                   *string             `db:"pricing_source" json:"pricing_source,omitempty"`
	PricedByUserID                  *int                `db:"priced_by_user_id" json:"priced_by_user_id,omitempty"`
	PricedAt                        pgtype.Timestamptz  `db:"priced_at" json:"priced_at,omitempty"`
	ApprovedAt                      pgtype.Timestamptz  `db:"approved_at" json:"approved_at,omitempty"`
	ApprovedByUserID                *int                `db:"approved_by_user_id" json:"approved_by_user_id,omitempty"`
	DeniedAt                        pgtype.Timestamptz  `db:"denied_at" json:"denied_at,omitempty"`
	DeniedByUserID                  *int                `db:"denied_by_user_id" json:"denied_by_user_id,omitempty"`
	DenialReason                    *string             `db:"denial_reason" json:"denial_reason,omitempty"`
	DecisionIdempotencyKey          *string             `db:"decision_idempotency_key" json:"decision_idempotency_key,omitempty"`
	DecisionPayloadHash             *string             `db:"decision_payload_hash" json:"decision_payload_hash,omitempty"`
	ApprovalPreflightEvidenceID     *int64              `db:"approval_preflight_evidence_id" json:"approval_preflight_evidence_id,omitempty"`
	RequestEvidenceID               *int64              `db:"request_evidence_id" json:"request_evidence_id,omitempty"`
	ProtocolState                   string              `db:"protocol_state" json:"protocol_state"`
	LiquidityShortfallRaw           string              `db:"liquidity_shortfall_raw" json:"liquidity_shortfall_raw"`
	AssetAmountRaw                  *string             `db:"asset_amount_raw" json:"asset_amount_raw,omitempty"`
	ShareAmountRaw                  string              `db:"share_amount_raw" json:"share_amount_raw"`
	PendingSharesRaw                string              `db:"pending_shares_raw" json:"pending_shares_raw"`
	ClaimableAssetsRaw              string              `db:"claimable_assets_raw" json:"claimable_assets_raw"`
	ClaimableSharesRaw              string              `db:"claimable_shares_raw" json:"claimable_shares_raw"`
	ClaimedAssetsRaw                string              `db:"claimed_assets_raw" json:"claimed_assets_raw"`
	ClaimedSharesRaw                string              `db:"claimed_shares_raw" json:"claimed_shares_raw"`
	TransitionVersion               int64               `db:"transition_version" json:"transition_version"`
	RequestLockedAt                 pgtype.Timestamptz  `db:"request_locked_at" json:"request_locked_at,omitempty"`
	RequestedAt                     pgtype.Timestamptz  `db:"requested_at" json:"requested_at,omitempty"`
	ClaimableAt                     pgtype.Timestamptz  `db:"claimable_at" json:"claimable_at,omitempty"`
	ClaimedAt                       pgtype.Timestamptz  `db:"claimed_at" json:"claimed_at,omitempty"`
	CancelledAt                     pgtype.Timestamptz  `db:"cancelled_at" json:"cancelled_at,omitempty"`
	CreatedAt                       pgtype.Timestamptz  `db:"created_at" json:"created_at"`
	UpdatedAt                       pgtype.Timestamptz  `db:"updated_at" json:"updated_at"`
}

func (model InvestmentRedemption) ToJSON() map[string]any {
	return map[string]any{
		"id":                                   model.ID,
		"offer_id":                             model.OfferID,
		"profile_id":                           model.ProfileID,
		"investment_id":                        model.InvestmentID,
		"vault_contract_id":                    model.VaultContractID,
		"request_controller_chain_account_id":  model.RequestControllerChainAccountID,
		"request_controller_address":           model.RequestControllerAddress,
		"vault_request_origin":                 model.VaultRequestOrigin,
		"vault_request_effect_id":              model.VaultRequestEffectID,
		"status":                               model.Status,
		"idempotency_key":                      model.IdempotencyKey,
		"estimated_nav_record_id":              model.EstimatedNAVRecordID,
		"estimated_nav_usdc_raw":               model.EstimatedNAVUSDCRaw,
		"estimated_nav_version":                model.EstimatedNAVVersion,
		"estimated_nav_valuation_block_number": model.EstimatedNAVValuationBlock,
		"estimated_nav_valuation_as_of":        model.EstimatedNAVValuationAsOf,
		"estimated_asset_amount_raw":           model.EstimatedAssetAmountRaw,
		"estimated_at":                         model.EstimatedAt,
		"dealing_cutoff_block_number":          model.DealingCutoffBlockNumber,
		"dealing_cutoff_at":                    model.DealingCutoffAt,
		"nav_record_id":                        model.NAVRecordID,
		"nav_usdc_raw":                         model.NAVUSDCRaw,
		"nav_version":                          model.NAVVersion,
		"nav_valuation_block_number":           model.NAVValuationBlockNumber,
		"nav_valuation_as_of":                  model.NAVValuationAsOf,
		"dealing_price_usdc_raw":               model.DealingPriceUSDCRaw,
		"pricing_source":                       model.PricingSource,
		"priced_by_user_id":                    model.PricedByUserID,
		"priced_at":                            model.PricedAt,
		"approved_at":                          model.ApprovedAt,
		"approved_by_user_id":                  model.ApprovedByUserID,
		"denied_at":                            model.DeniedAt,
		"denied_by_user_id":                    model.DeniedByUserID,
		"denial_reason":                        model.DenialReason,
		"decision_idempotency_key":             model.DecisionIdempotencyKey,
		"decision_payload_hash":                model.DecisionPayloadHash,
		"approval_preflight_evidence_id":       model.ApprovalPreflightEvidenceID,
		"request_evidence_id":                  model.RequestEvidenceID,
		"protocol_state":                       model.ProtocolState,
		"liquidity_shortfall_raw":              model.LiquidityShortfallRaw,
		"asset_amount_raw":                     model.AssetAmountRaw,
		"share_amount_raw":                     model.ShareAmountRaw,
		"pending_shares_raw":                   model.PendingSharesRaw,
		"claimable_assets_raw":                 model.ClaimableAssetsRaw,
		"claimable_shares_raw":                 model.ClaimableSharesRaw,
		"claimed_assets_raw":                   model.ClaimedAssetsRaw,
		"claimed_shares_raw":                   model.ClaimedSharesRaw,
		"transition_version":                   model.TransitionVersion,
		"request_locked_at":                    model.RequestLockedAt,
		"requested_at":                         model.RequestedAt,
		"claimable_at":                         model.ClaimableAt,
		"claimed_at":                           model.ClaimedAt,
		"cancelled_at":                         model.CancelledAt,
		"created_at":                           model.CreatedAt,
		"updated_at":                           model.UpdatedAt,
	}
}

func (model InvestmentRedemption) Fields() []string {
	return []string{
		"id", "offer_id", "profile_id", "investment_id", "vault_contract_id",
		"request_controller_chain_account_id", "request_controller_address",
		"vault_request_origin", "vault_request_effect_id", "status", "idempotency_key",
		"estimated_nav_record_id", "estimated_nav_usdc_raw",
		"estimated_nav_version", "estimated_nav_valuation_block_number",
		"estimated_nav_valuation_as_of", "estimated_asset_amount_raw", "estimated_at",
		"dealing_cutoff_block_number", "dealing_cutoff_at", "nav_record_id",
		"nav_usdc_raw", "nav_version", "nav_valuation_block_number",
		"nav_valuation_as_of", "liquidity_shortfall_raw",
		"dealing_price_usdc_raw", "pricing_source", "priced_by_user_id",
		"priced_at", "approved_at", "approved_by_user_id", "denied_at", "denied_by_user_id",
		"denial_reason", "decision_idempotency_key", "decision_payload_hash",
		"approval_preflight_evidence_id",
		"request_evidence_id", "protocol_state",
		"asset_amount_raw", "share_amount_raw", "pending_shares_raw",
		"claimable_assets_raw", "claimable_shares_raw", "claimed_assets_raw",
		"claimed_shares_raw", "transition_version", "request_locked_at",
		"requested_at", "claimable_at", "claimed_at", "cancelled_at", "created_at",
		"updated_at",
	}
}

func (model InvestmentRedemption) Table() string { return "investment_redemptions" }
func (model InvestmentRedemption) GetID() any    { return model.ID }
func (model *InvestmentRedemption) SetID(id any) { model.ID = id.(int64) }

// InvestmentProfile is an object representing the database table.
type InvestmentProfile struct {
	ID                  int                `db:"id" json:"id" yaml:"id"`
	UserID              *int               `db:"user_id" json:"user_id,omitempty" yaml:"user_id,omitempty"`
	Type                ProfileT           `db:"type" json:"type" yaml:"type"`
	Data                any                `db:"data" json:"data" yaml:"data"`
	KycID               *string            `db:"kyc_id" json:"kyc_id,omitempty" yaml:"kyc_id,omitempty"`
	KycStatus           KycT               `db:"kyc_status" json:"kyc_status" yaml:"kyc_status"`
	KycData             any                `db:"kyc_data" json:"kyc_data" yaml:"kyc_data"`
	AccreditationID     *string            `db:"accreditation_id" json:"accreditation_id,omitempty" yaml:"accreditation_id,omitempty"`
	AccreditationStatus AccreditationT     `db:"accreditation_status" json:"accreditation_status" yaml:"accreditation_status"`
	AccreditationData   any                `db:"accreditation_data" json:"accreditation_data" yaml:"accreditation_data"`
	EscrowID            string             `db:"escrow_id" json:"escrow_id" yaml:"escrow_id"`
	KycAt               pgtype.Timestamptz `db:"kyc_at" json:"kyc_at,omitempty" yaml:"kyc_at,omitempty"`
	AccreditationAt     pgtype.Timestamptz `db:"accreditation_at" json:"accreditation_at,omitempty" yaml:"accreditation_at,omitempty"`
	CreatedAt           pgtype.Timestamptz `db:"created_at" json:"created_at,omitempty" yaml:"created_at,omitempty"`
	UpdatedAt           pgtype.Timestamptz `db:"updated_at" json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	WalletID            int                `db:"-" json:"wallet_id,omitempty" yaml:"wallet_id,omitempty"`
}

func (model InvestmentProfile) ToJSON() map[string]any {
	return map[string]any{
		"id":                   model.ID,
		"user_id":              model.UserID,
		"type":                 model.Type,
		"data":                 model.Data,
		"kyc_id":               model.KycID,
		"kyc_status":           model.KycStatus,
		"kyc_data":             model.KycData,
		"accreditation_id":     model.AccreditationID,
		"accreditation_status": model.AccreditationStatus,
		"accreditation_data":   model.AccreditationData,
		"escrow_id":            model.EscrowID,
		"kyc_at":               model.KycAt,
		"accreditation_at":     model.AccreditationAt,
		"created_at":           model.CreatedAt,
		"updated_at":           model.UpdatedAt,
		"wallet_id":            model.WalletID,
	}
}

func (model InvestmentProfile) Fields() []string {
	return []string{
		"id",
		"user_id",
		"type",
		"data",
		"kyc_id",
		"kyc_status",
		"kyc_data",
		"accreditation_id",
		"accreditation_status",
		"accreditation_data",
		"escrow_id",
		"created_at",
		"updated_at",
	}
}

func (model InvestmentProfile) Table() string {
	return "investment_profiles"
}

func (model InvestmentProfile) GetID() any {
	return model.ID
}

func (model *InvestmentProfile) SetID(id any) {
	model.ID = id.(int)
}
