package evmwalletoperations

import (
	"github.com/global-torque/go-common/db/v2"
	"github.com/global-torque/go-common/orm/v2/pgtype"
)

// WalletOperation is the canonical unified transaction record.
//
// TxHash is pgtype.Text so platform pre-insert rows can carry SQL NULL
// and avoid the partial unique index `WHERE tx_hash IS NOT NULL`.
// BlockTimestamp is pgtype.Timestamptz so it can be NULL until the
// webhook resolves it from the block.
type WalletOperation struct {
	ID                                   int                `db:"id" json:"id"`
	UserID                               *int               `db:"user_id" json:"user_id,omitempty"`
	ProfileID                            *int               `db:"profile_id" json:"profile_id,omitempty"`
	InvestmentID                         *int               `db:"investment_id" json:"investment_id"`
	TokenID                              *int               `db:"token_id" json:"token_id"`
	WalletAddress                        string             `db:"wallet_address" json:"wallet_address"`
	Chain                                string             `db:"chain" json:"chain"`
	Source                               OperationSourceT   `db:"source" json:"source"`
	Type                                 OperationTypeT     `db:"type" json:"type"`
	Status                               OperationStatusT   `db:"status" json:"status"`
	Amount                               string             `db:"amount" json:"amount"`
	AmountRaw                            string             `db:"amount_raw" json:"amount_raw"`
	PriceUSD                             string             `db:"price_usd" json:"price_usd"`
	AmountUSD                            string             `db:"amount_usd" json:"amount_usd"`
	TokenTicker                          string             `db:"token_ticker" json:"token_ticker"`
	TokenAddress                         string             `db:"token_address" json:"token_address"`
	TokenName                            string             `db:"token_name" json:"token_name"`
	TokenLogo                            string             `db:"token_logo" json:"token_logo"`
	TokenDecimals                        int                `db:"token_decimals" json:"token_decimals"`
	CounterpartyAddress                  string             `db:"counterparty_address" json:"counterparty_address"`
	TxHash                               pgtype.Text        `db:"tx_hash" json:"tx_hash"`
	ExternalID                           string             `db:"external_id" json:"external_id"`
	IdempotencyKey                       string             `db:"idempotency_key" json:"idempotency_key"`
	ProviderEventID                      string             `db:"provider_event_id" json:"provider_event_id"`
	FailureReason                        string             `db:"failure_reason" json:"failure_reason"`
	BlockNumber                          int64              `db:"block_number" json:"block_number"`
	BlockTimestamp                       pgtype.Timestamptz `db:"block_timestamp" json:"block_timestamp"`
	ConfirmationCount                    int                `db:"confirmation_count" json:"confirmation_count"`
	ConfirmationTarget                   int                `db:"confirmation_target" json:"confirmation_target"`
	ReorgCount                           int                `db:"reorg_count" json:"reorg_count"`
	LastSeenBlock                        int64              `db:"last_seen_block" json:"last_seen_block"`
	RemovedAt                            pgtype.Timestamptz `db:"removed_at" json:"removed_at"`
	ProviderName                         string             `db:"provider_name" json:"provider_name"`
	WalletSessionID                      *int               `db:"wallet_session_id" json:"-"`
	WalletSessionReservationConsumed     bool               `db:"wallet_session_reservation_consumed" json:"-"`
	TurnkeyActivityID                    string             `db:"turnkey_activity_id" json:"-"`
	TurnkeyOrgID                         string             `db:"turnkey_org_id" json:"-"`
	TurnkeySubOrgID                      string             `db:"turnkey_sub_org_id" json:"-"`
	SubmissionStatus                     string             `db:"submission_status" json:"-"`
	ExternalTxHash                       string             `db:"external_tx_hash" json:"external_tx_hash,omitempty"`
	TurnkeyCredentialRole                string             `db:"turnkey_credential_role" json:"-"`
	TurnkeyCredentialFingerprint         string             `db:"turnkey_credential_fingerprint" json:"-"`
	TurnkeyCredentialVersion             int64              `db:"turnkey_credential_version" json:"-"`
	TurnkeyRequestTimestampMS            int64              `db:"turnkey_request_timestamp_ms" json:"-"`
	TurnkeyRequestBodyHash               string             `db:"turnkey_request_body_hash" json:"-"`
	ProviderPreparedCallRaw              string             `db:"provider_prepared_call_raw" json:"-"`
	ReconciliationAttemptCount           int                `db:"reconciliation_attempt_count" json:"-"`
	ReconciliationNextRetryAt            pgtype.Timestamptz `db:"reconciliation_next_retry_at" json:"-"`
	ReconciliationLastErrorCode          string             `db:"reconciliation_last_error_code" json:"-"`
	ComplianceStatus                     string             `db:"compliance_status" json:"-"`
	ComplianceCheckedAt                  pgtype.Timestamptz `db:"compliance_checked_at" json:"-"`
	ComplianceDecisionVersion            int                `db:"compliance_decision_version" json:"-"`
	CompliancePolicySHA256               string             `db:"compliance_policy_sha256" json:"-"`
	ExchangeRWAOperationID               *int               `db:"exchange_rwa_operation_id" json:"-"`
	PreparedTxRaw                        string             `db:"prepared_tx_raw" json:"-"`
	PreparedTxNonce                      *int64             `db:"prepared_tx_nonce" json:"-"`
	TransactionPreparedAt                pgtype.Timestamptz `db:"transaction_prepared_at" json:"-"`
	TransactionBroadcastAt               pgtype.Timestamptz `db:"transaction_broadcast_at" json:"-"`
	TransactionBroadcastAttemptCount     int                `db:"transaction_broadcast_attempt_count" json:"-"`
	TransactionBroadcastLastError        string             `db:"transaction_broadcast_last_error" json:"-"`
	ExchangePayoutAttemptNumber          int                `db:"exchange_payout_attempt_number" json:"-"`
	ExchangePayoutCompensationStatus     string             `db:"exchange_payout_compensation_status" json:"-"`
	ExchangePayoutCompensationRequiredAt pgtype.Timestamptz `db:"exchange_payout_compensation_required_at" json:"-"`
	ExchangeRedemptionApprovalReference  string             `db:"exchange_redemption_approval_reference" json:"-"`
	ExchangeRedemptionApprovedAt         pgtype.Timestamptz `db:"exchange_redemption_approved_at" json:"-"`
	InvestmentRedemptionID               *int64             `db:"investment_redemption_id" json:"investment_redemption_id,omitempty"`
	RedemptionFulfillmentGrantID         *int64             `db:"redemption_fulfillment_grant_id" json:"redemption_fulfillment_grant_id,omitempty"`
	RedemptionFulfillmentGrantVersion    *int64             `db:"redemption_fulfillment_grant_version" json:"redemption_fulfillment_grant_version,omitempty"`
	RedemptionFulfillmentGrantScopeHash  *string            `db:"redemption_fulfillment_grant_scope_hash" json:"redemption_fulfillment_grant_scope_hash,omitempty"`
	RedemptionFulfillmentGrantEpoch      *int64             `db:"redemption_fulfillment_grant_authorization_epoch" json:"redemption_fulfillment_grant_authorization_epoch,omitempty"`
	RedemptionFulfillmentGrantExpiresAt  pgtype.Timestamptz `db:"redemption_fulfillment_grant_expires_at" json:"redemption_fulfillment_grant_expires_at,omitempty"`
	RedemptionFulfillmentEvidenceDigest  *string            `db:"redemption_fulfillment_evidence_digest" json:"redemption_fulfillment_evidence_digest,omitempty"`
	RedemptionFulfillmentArmedAt         pgtype.Timestamptz `db:"redemption_fulfillment_armed_at" json:"redemption_fulfillment_armed_at,omitempty"`
	FundNAVRecordID                      *int64             `db:"fund_nav_record_id" json:"fund_nav_record_id,omitempty"`
	TransactionFromAddress               *string            `db:"transaction_from_address" json:"transaction_from_address,omitempty"`
	TransactionToAddress                 *string            `db:"transaction_to_address" json:"transaction_to_address,omitempty"`
	CallTargetAddress                    *string            `db:"call_target_address" json:"call_target_address,omitempty"`
	ContractID                           *int               `db:"contract_id" json:"contract_id,omitempty"`
	ContractFunctionSelector             *string            `db:"contract_function_selector" json:"contract_function_selector,omitempty"`
	CallData                             *string            `db:"call_data" json:"-"`
	TransactionCallData                  *string            `db:"transaction_call_data" json:"-"`
	ReceiptBlockHash                     *string            `db:"receipt_block_hash" json:"receipt_block_hash,omitempty"`
	CreatedAt                            pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt                            pgtype.Timestamptz `db:"updated_at" json:"updated_at"`

	db db.Repository `db:"-" json:"-"`
}

func (model WalletOperation) Fields() []string {
	return []string{
		"id", "user_id", "profile_id", "investment_id", "token_id", "wallet_address",
		"chain", "source", "type", "status", "amount", "amount_raw", "price_usd",
		"amount_usd", "token_ticker", "token_address", "token_name", "token_logo",
		"token_decimals", "counterparty_address", "tx_hash", "external_id",
		"idempotency_key", "provider_event_id", "failure_reason", "block_number",
		"block_timestamp", "confirmation_count", "confirmation_target", "reorg_count",
		"last_seen_block", "removed_at", "created_at", "updated_at",
		"investment_redemption_id", "fund_nav_record_id", "transaction_from_address",
		"redemption_fulfillment_grant_id", "redemption_fulfillment_grant_version",
		"redemption_fulfillment_grant_scope_hash", "redemption_fulfillment_grant_authorization_epoch",
		"redemption_fulfillment_grant_expires_at", "redemption_fulfillment_evidence_digest",
		"redemption_fulfillment_armed_at",
		"transaction_to_address", "call_target_address", "contract_id",
		"contract_function_selector", "call_data", "transaction_call_data", "receipt_block_hash",
		"provider_prepared_call_raw", "wallet_session_reservation_consumed",
	}
}

func (model WalletOperation) Table() string {
	return TableName
}

func (model WalletOperation) ToJSON() map[string]any {
	res := map[string]any{}
	for _, field := range model.Fields() {
		if field == "provider_prepared_call_raw" ||
			field == "wallet_session_reservation_consumed" {
			continue
		}
		res[field] = model.GetValueByTag(field)
	}
	return res
}

func (model WalletOperation) GetID() any {
	return model.ID
}

func (model *WalletOperation) SetID(id any) {
	model.ID = id.(int)
}

func (model *WalletOperation) SetDB(db db.Repository) {
	model.db = db
}

func (model WalletOperation) GetValueByTag(name string) any {
	switch name {
	case "id":
		return model.ID
	case "user_id":
		return model.UserID
	case "profile_id":
		return model.ProfileID
	case "investment_id":
		return model.InvestmentID
	case "token_id":
		return model.TokenID
	case "wallet_address":
		return model.WalletAddress
	case "chain":
		return model.Chain
	case "source":
		return model.Source
	case "type":
		return model.Type
	case "status":
		return model.Status
	case "amount":
		return model.Amount
	case "amount_raw":
		return model.AmountRaw
	case "price_usd":
		return model.PriceUSD
	case "amount_usd":
		return model.AmountUSD
	case "token_ticker":
		return model.TokenTicker
	case "token_address":
		return model.TokenAddress
	case "token_name":
		return model.TokenName
	case "token_logo":
		return model.TokenLogo
	case "token_decimals":
		return model.TokenDecimals
	case "counterparty_address":
		return model.CounterpartyAddress
	case "tx_hash":
		return model.TxHash
	case "external_id":
		return model.ExternalID
	case "idempotency_key":
		return model.IdempotencyKey
	case "provider_event_id":
		return model.ProviderEventID
	case "failure_reason":
		return model.FailureReason
	case "block_number":
		return model.BlockNumber
	case "block_timestamp":
		return model.BlockTimestamp
	case "confirmation_count":
		return model.ConfirmationCount
	case "confirmation_target":
		return model.ConfirmationTarget
	case "reorg_count":
		return model.ReorgCount
	case "last_seen_block":
		return model.LastSeenBlock
	case "removed_at":
		return model.RemovedAt
	case "investment_redemption_id":
		return model.InvestmentRedemptionID
	case "redemption_fulfillment_grant_id":
		return model.RedemptionFulfillmentGrantID
	case "redemption_fulfillment_grant_version":
		return model.RedemptionFulfillmentGrantVersion
	case "redemption_fulfillment_grant_scope_hash":
		return model.RedemptionFulfillmentGrantScopeHash
	case "redemption_fulfillment_grant_authorization_epoch":
		return model.RedemptionFulfillmentGrantEpoch
	case "redemption_fulfillment_grant_expires_at":
		return model.RedemptionFulfillmentGrantExpiresAt
	case "redemption_fulfillment_evidence_digest":
		return model.RedemptionFulfillmentEvidenceDigest
	case "redemption_fulfillment_armed_at":
		return model.RedemptionFulfillmentArmedAt
	case "fund_nav_record_id":
		return model.FundNAVRecordID
	case "transaction_from_address":
		return model.TransactionFromAddress
	case "transaction_to_address":
		return model.TransactionToAddress
	case "call_target_address":
		return model.CallTargetAddress
	case "contract_id":
		return model.ContractID
	case "contract_function_selector":
		return model.ContractFunctionSelector
	case "call_data":
		return model.CallData
	case "transaction_call_data":
		return model.TransactionCallData
	case "receipt_block_hash":
		return model.ReceiptBlockHash
	case "provider_prepared_call_raw":
		return model.ProviderPreparedCallRaw
	case "wallet_session_reservation_consumed":
		return model.WalletSessionReservationConsumed
	case "created_at":
		return model.CreatedAt
	case "updated_at":
		return model.UpdatedAt
	default:
		return nil
	}
}
