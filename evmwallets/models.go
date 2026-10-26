package evmwallets

import (
	"context"

	"github.com/global-torque/go-common/db/v2"
	"github.com/global-torque/go-common/logger/v2"
	"github.com/global-torque/go-common/orm/v2"
	"github.com/global-torque/go-common/orm/v2/pgtype"
	"github.com/pkg/errors"
)

// Wallet is an object representing the database table.
type Wallet struct {
	ID            int    `db:"id" json:"id" yaml:"id"`
	ContentTypeID int    `db:"content_type_id" json:"content_type_id" yaml:"content_type_id"`
	UserID        *int   `db:"user_id" json:"user_id,omitempty" yaml:"user_id,omitempty"`
	ObjectID      string `db:"object_id" json:"object_id" yaml:"object_id"`

	PublicKey  string `db:"public_key" json:"public_key" yaml:"public_key"`
	PrivateKey string `db:"-" json:"-" yaml:"-"`

	Balance                                    float64            `db:"balance" json:"balance" yaml:"balance"`
	IncBalance                                 float64            `db:"inc_balance" json:"inc_balance" yaml:"inc_balance"`
	OutBalance                                 float64            `db:"out_balance" json:"out_balance" yaml:"out_balance"`
	Status                                     WalletStatusT      `db:"status" json:"status" yaml:"status"`
	CreatedAt                                  pgtype.Timestamptz `db:"created_at" json:"created_at" yaml:"created_at"`
	UpdatedAt                                  pgtype.Timestamptz `db:"updated_at" json:"updated_at" yaml:"updated_at"`
	ProviderName                               string             `db:"provider_name" json:"provider_name" yaml:"provider_name"`
	ProviderUserID                             string             `db:"provider_user_id" json:"-" yaml:"-"`
	MFANotifiedAt                              pgtype.Timestamptz `db:"mfa_notified_at" json:"-" yaml:"-"`
	TurnkeyOrgID                               string             `db:"turnkey_org_id" json:"-" yaml:"-"`
	TurnkeySubOrgID                            string             `db:"turnkey_sub_org_id" json:"-" yaml:"-"`
	TurnkeyUserID                              string             `db:"turnkey_user_id" json:"-" yaml:"-"`
	TurnkeyWalletID                            string             `db:"turnkey_wallet_id" json:"-" yaml:"-"`
	TurnkeyDelegatedUserID                     string             `db:"turnkey_delegated_user_id" json:"-" yaml:"-"`
	TurnkeyDelegatedAPIKeyPublicKey            string             `db:"turnkey_delegated_api_key_public_key" json:"-" yaml:"-"`
	TurnkeyAccountKey                          string             `db:"turnkey_account_key" json:"-" yaml:"-"`
	TurnkeyDelegatedCredentialFingerprint      string             `db:"turnkey_delegated_credential_fingerprint" json:"-" yaml:"-"`
	TurnkeyDelegatedCredentialVersion          int64              `db:"turnkey_delegated_credential_version" json:"-" yaml:"-"`
	TurnkeyDelegatedNonRootVerifiedAt          pgtype.Timestamptz `db:"turnkey_delegated_non_root_verified_at" json:"-" yaml:"-"`
	TurnkeyDelegatedNonRootVerificationVersion int                `db:"turnkey_delegated_non_root_verification_version" json:"-" yaml:"-"`
	TurnkeyExecutionAllowlisted                bool               `db:"turnkey_execution_allowlisted" json:"-" yaml:"-"`
	TurnkeyExecutionAllowlistedAt              pgtype.Timestamptz `db:"turnkey_execution_allowlisted_at" json:"-" yaml:"-"`
	TurnkeyExecutionAllowlistVersion           int                `db:"turnkey_execution_allowlist_version" json:"-" yaml:"-"`
	TurnkeyExecutionAllowlistApprovalID        string             `db:"turnkey_execution_allowlist_approval_id" json:"-" yaml:"-"`
	TurnkeyExecutionInventoryAuditSHA256       string             `db:"turnkey_execution_inventory_audit_sha256" json:"-" yaml:"-"`
	TurnkeyExecutionLifecycleEvidenceSHA256    string             `db:"turnkey_execution_lifecycle_evidence_sha256" json:"-" yaml:"-"`

	updatedFields []string       `db:"-" json:"-"`
	fns           map[string]any `db:"-" json:"-"`
	db            db.Repository  `db:"-" json:"-"`
}

func New(db db.Repository) *Wallet {
	return &Wallet{
		db:  db,
		fns: map[string]any{},
	}
}

func (model Wallet) GetField(name string) any {
	switch name {
	case "ID":
		return model.ID
	case "UserID":
		return model.UserID
	case "PublicKey":
		return model.PublicKey
	case "PrivateKey":
		return model.PrivateKey
	case "Balance":
		return model.Balance
	case "IncBalance":
		return model.IncBalance
	case "OutBalance":
		return model.OutBalance
	case "Status":
		return model.Status
	case "CreatedAt":
		return model.CreatedAt
	case "UpdatedAt":
		return model.UpdatedAt
	}
	return nil
}

func (model Wallet) GetValueByTag(name string) any {
	switch name {
	case "id":
		return model.ID
	case "user_id":
		return model.UserID
	case "public_key":
		return model.PublicKey
	case "private_key":
		return model.PrivateKey
	case "balance":
		return model.Balance
	case "inc_balance":
		return model.IncBalance
	case "out_balance":
		return model.OutBalance
	case "status":
		return model.Status
	case "created_at":
		return model.CreatedAt
	case "updated_at":
		return model.UpdatedAt
	}
	return nil
}

func (model Wallet) ToJSON() map[string]any {
	res := map[string]any{}
	fields := model.Fields()
	for _, key := range fields {
		if key == "private_key" {
			continue
		}
		res[key] = model.GetValueByTag(key)
	}
	return res
}

func (model Wallet) Fields() []string {
	return []string{
		"id", "user_id", "public_key", "balance", "inc_balance",
		"out_balance", "status", "created_at", "updated_at",
	}
}

func (model Wallet) Table() string {
	return TableName
}

func (model Wallet) GetID() any {
	return model.ID
}

func (model *Wallet) SetID(id any) {
	model.ID = id.(int)
}

func (model *Wallet) SetDB(db db.Repository) {
	model.db = db
}

func (model Wallet) Save(ctx context.Context) error {
	if model.ID == 0 {
		err := errors.Errorf("%s: Wallet %d", orm.ErrEmptyID, model.ID)
		logger.FromCtx(ctx, pkgName).Error().Stack().Err(err).Msg(orm.ErrEmptyID.Error())
		return err
	}

	updates := map[string]any{}
	for _, field := range model.updatedFields {
		updates[field] = model.GetValueByTag(field)
	}
	updated, err := orm.Update[Wallet](
		ctx,
		model.db,
		map[string]any{
			"id": model.ID,
		},
		updates,
	)
	if err != nil {
		err = errors.Wrapf(err, "cannot update %d", model.ID)
		return err
	}
	if updated == false {
		err := errors.Errorf("%s: Wallet %d", orm.ErrNoRowsAffected, model.ID)
		logger.FromCtx(ctx, pkgName).Error().Stack().Err(err).Msg(orm.ErrNoRowsAffected.Error())
		return err
	}
	return nil
}
