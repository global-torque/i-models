package evmredemptionfulfillmentgrants

import "github.com/global-torque/go-common/orm/v2/pgtype"

// RedemptionFulfillmentGrant is the persisted narrow authorization used by
// automatic redemption fulfillment. Scope and provider evidence are retained
// as structured values so an operation can bind to one exact grant version.
type RedemptionFulfillmentGrant struct {
	ID                         int64              `db:"id" json:"id"`
	OfferID                    int                `db:"offer_id" json:"offer_id"`
	VaultContractID            int                `db:"vault_contract_id" json:"vault_contract_id"`
	Chain                      string             `db:"chain" json:"chain"`
	FundWalletAddress          string             `db:"fund_wallet_address" json:"fund_wallet_address"`
	ExecutorAddress            string             `db:"executor_address" json:"executor_address"`
	FulfillSelector            string             `db:"fulfill_selector" json:"fulfill_selector"`
	Scope                      any                `db:"scope" json:"scope"`
	ScopeHash                  string             `db:"scope_hash" json:"scope_hash"`
	Version                    int64              `db:"version" json:"version"`
	AuthorizationEpoch         int64              `db:"authorization_epoch" json:"authorization_epoch"`
	AuthorizationExpiresAt     pgtype.Timestamptz `db:"authorization_expires_at" json:"authorization_expires_at,omitempty"`
	Status                     StatusT            `db:"status" json:"status"`
	AuthorizedAt               pgtype.Timestamptz `db:"authorized_at" json:"authorized_at"`
	RevokedAt                  pgtype.Timestamptz `db:"revoked_at" json:"revoked_at,omitempty"`
	ProviderName               string             `db:"provider_name" json:"provider_name"`
	ProviderGrantID            string             `db:"provider_grant_id" json:"provider_grant_id"`
	ProviderPolicyID           string             `db:"provider_policy_id" json:"provider_policy_id"`
	ProviderActivityID         string             `db:"provider_activity_id" json:"provider_activity_id"`
	ProviderEvidence           any                `db:"provider_evidence" json:"provider_evidence"`
	ProviderVerifiedAt         pgtype.Timestamptz `db:"provider_verified_at" json:"provider_verified_at,omitempty"`
	ProviderRevocationActivity string             `db:"provider_revocation_activity_id" json:"provider_revocation_activity_id"`
	ProviderRevokedAt          pgtype.Timestamptz `db:"provider_revoked_at" json:"provider_revoked_at,omitempty"`
	SupersededAt               pgtype.Timestamptz `db:"superseded_at" json:"superseded_at,omitempty"`
	CreatedAt                  pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt                  pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
}

func (model RedemptionFulfillmentGrant) ToJSON() map[string]any {
	result := map[string]any{}
	for _, field := range model.Fields() {
		result[field] = model.value(field)
	}
	return result
}

func (model RedemptionFulfillmentGrant) value(field string) any {
	switch field {
	case "id":
		return model.ID
	case "offer_id":
		return model.OfferID
	case "vault_contract_id":
		return model.VaultContractID
	case "chain":
		return model.Chain
	case "fund_wallet_address":
		return model.FundWalletAddress
	case "executor_address":
		return model.ExecutorAddress
	case "fulfill_selector":
		return model.FulfillSelector
	case "scope":
		return model.Scope
	case "scope_hash":
		return model.ScopeHash
	case "version":
		return model.Version
	case "authorization_epoch":
		return model.AuthorizationEpoch
	case "authorization_expires_at":
		return model.AuthorizationExpiresAt
	case "status":
		return model.Status
	case "authorized_at":
		return model.AuthorizedAt
	case "revoked_at":
		return model.RevokedAt
	case "provider_name":
		return model.ProviderName
	case "provider_grant_id":
		return model.ProviderGrantID
	case "provider_policy_id":
		return model.ProviderPolicyID
	case "provider_activity_id":
		return model.ProviderActivityID
	case "provider_evidence":
		return model.ProviderEvidence
	case "provider_verified_at":
		return model.ProviderVerifiedAt
	case "provider_revocation_activity_id":
		return model.ProviderRevocationActivity
	case "provider_revoked_at":
		return model.ProviderRevokedAt
	case "superseded_at":
		return model.SupersededAt
	case "created_at":
		return model.CreatedAt
	case "updated_at":
		return model.UpdatedAt
	default:
		return nil
	}
}

func (model RedemptionFulfillmentGrant) Fields() []string {
	return []string{"id", "offer_id", "vault_contract_id", "chain", "fund_wallet_address", "executor_address", "fulfill_selector", "scope", "scope_hash", "version", "authorization_epoch", "authorization_expires_at", "status", "authorized_at", "revoked_at", "provider_name", "provider_grant_id", "provider_policy_id", "provider_activity_id", "provider_evidence", "provider_verified_at", "provider_revocation_activity_id", "provider_revoked_at", "superseded_at", "created_at", "updated_at"}
}

func (model RedemptionFulfillmentGrant) Table() string { return "evm_redemption_fulfillment_grants" }
func (model RedemptionFulfillmentGrant) GetID() any    { return model.ID }
func (model *RedemptionFulfillmentGrant) SetID(id any) { model.ID = id.(int64) }
