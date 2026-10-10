package evmcontracts

import (
	"context"

	"github.com/global-torque/go-common/db/v2"
	"github.com/global-torque/go-common/logger/v2"
	"github.com/global-torque/go-common/orm/v2"
	"github.com/global-torque/go-common/orm/v2/pgtype"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// Wallet is an object representing the database table.
type Contract struct {
	ID      int  `db:"id" json:"id" yaml:"id"`
	UserID  *int `db:"user_id" json:"user_id,omitempty" yaml:"user_id,omitempty"`
	OfferID *int `db:"offer_id" json:"offer_id,omitempty" yaml:"offer_id,omitempty"`

	Name                         string             `db:"name" json:"name" yaml:"name"`
	Status                       StatusT            `db:"status" json:"status" yaml:"status"`
	Symbol                       string             `db:"symbol" json:"symbol" yaml:"symbol"`
	Address                      string             `db:"address" json:"address" yaml:"address"`
	TransactionTX                string             `db:"transaction_tx" json:"transaction_tx" yaml:"transaction_tx"`
	DeploymentOperationID        *uuid.UUID         `db:"deployment_operation_id" json:"deployment_operation_id,omitempty" yaml:"deployment_operation_id,omitempty"`
	DeploymentLeg                *DeploymentLegT    `db:"deployment_leg" json:"deployment_leg,omitempty" yaml:"deployment_leg,omitempty"`
	Chain                        *string            `db:"chain" json:"chain,omitempty" yaml:"chain,omitempty"`
	FulfillmentChainAccountID    *int               `db:"fulfillment_chain_account_id" json:"fulfillment_chain_account_id,omitempty" yaml:"fulfillment_chain_account_id,omitempty"`
	FulfillmentControllerAddress *string            `db:"fulfillment_controller_address" json:"fulfillment_controller_address,omitempty" yaml:"fulfillment_controller_address,omitempty"`
	AssetAddress                 *string            `db:"asset_address" json:"asset_address,omitempty" yaml:"asset_address,omitempty"`
	AssetDecimals                *int               `db:"asset_decimals" json:"asset_decimals,omitempty" yaml:"asset_decimals,omitempty"`
	ShareDecimals                *int               `db:"share_decimals" json:"share_decimals,omitempty" yaml:"share_decimals,omitempty"`
	VaultStandard                *string            `db:"vault_standard" json:"vault_standard,omitempty" yaml:"vault_standard,omitempty"`
	TokenizationEngine           *string            `db:"tokenization_engine" json:"tokenization_engine,omitempty" yaml:"tokenization_engine,omitempty"`
	EventsScannedThroughBlock    *int64             `db:"events_scanned_through_block_number" json:"events_scanned_through_block_number,omitempty" yaml:"events_scanned_through_block_number,omitempty"`
	EventsScannedThroughHash     *string            `db:"events_scanned_through_block_hash" json:"events_scanned_through_block_hash,omitempty" yaml:"events_scanned_through_block_hash,omitempty"`
	RedemptionEnabled            bool               `db:"redemption_enabled" json:"redemption_enabled" yaml:"redemption_enabled"`
	RedemptionChain              string             `db:"redemption_chain" json:"redemption_chain" yaml:"redemption_chain"`
	RedemptionApprovedAt         pgtype.Timestamptz `db:"redemption_approved_at" json:"redemption_approved_at,omitempty" yaml:"redemption_approved_at,omitempty"`
	RedemptionApprovalReference  string             `db:"redemption_approval_reference" json:"redemption_approval_reference,omitempty" yaml:"redemption_approval_reference,omitempty"`

	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at" yaml:"created_at"`
	UpdatedAt pgtype.Timestamptz `db:"updated_at" json:"updated_at" yaml:"updated_at"`

	updatedFields []string       `db:"-" json:"-"`
	fns           map[string]any `db:"-" json:"-"`
	db            db.Repository  `db:"-" json:"-"`
}

func New(db db.Repository) *Contract {
	return &Contract{
		db:  db,
		fns: map[string]any{},
	}
}

func (model Contract) GetField(name string) any {
	switch name {
	case "ID":
		return model.ID
	case "UserID":
		return model.UserID
	case "OfferID":
		return model.OfferID
	case "TransactionTX":
		return model.TransactionTX
	case "Status":
		return model.Status
	case "DeploymentOperationID":
		return model.DeploymentOperationID
	case "Chain":
		return model.Chain
	case "CreatedAt":
		return model.CreatedAt
	case "UpdatedAt":
		return model.UpdatedAt
	}
	return nil
}

func (model Contract) GetValueByTag(name string) any {
	switch name {
	case "id":
		return model.ID
	case "user_id":
		return model.UserID
	case "offer_id":
		return model.OfferID
	case "name":
		return model.Name
	case "symbol":
		return model.Symbol
	case "address":
		return model.Address
	case "transaction_tx":
		return model.TransactionTX
	case "status":
		return model.Status
	case "deployment_operation_id":
		return model.DeploymentOperationID
	case "deployment_leg":
		return model.DeploymentLeg
	case "chain":
		return model.Chain
	case "fulfillment_chain_account_id":
		return model.FulfillmentChainAccountID
	case "fulfillment_controller_address":
		return model.FulfillmentControllerAddress
	case "asset_address":
		return model.AssetAddress
	case "asset_decimals":
		return model.AssetDecimals
	case "share_decimals":
		return model.ShareDecimals
	case "vault_standard":
		return model.VaultStandard
	case "tokenization_engine":
		return model.TokenizationEngine
	case "events_scanned_through_block_number":
		return model.EventsScannedThroughBlock
	case "events_scanned_through_block_hash":
		return model.EventsScannedThroughHash
	case "redemption_enabled":
		return model.RedemptionEnabled
	case "redemption_chain":
		return model.RedemptionChain
	case "redemption_approved_at":
		return model.RedemptionApprovedAt
	case "redemption_approval_reference":
		return model.RedemptionApprovalReference
	case "created_at":
		return model.CreatedAt
	case "updated_at":
		return model.UpdatedAt
	}
	return nil
}

func (model Contract) ToJSON() map[string]any {
	res := map[string]any{}
	fields := model.Fields()
	for _, key := range fields {
		res[key] = model.GetValueByTag(key)
	}
	return res
}

func (model Contract) Fields() []string {
	return []string{
		"id", "user_id", "offer_id", "name", "status", "symbol", "address",
		"transaction_tx", "deployment_operation_id", "deployment_leg", "chain",
		"fulfillment_chain_account_id", "fulfillment_controller_address",
		"asset_address", "asset_decimals", "share_decimals", "vault_standard",
		"tokenization_engine", "events_scanned_through_block_number",
		"events_scanned_through_block_hash", "redemption_enabled", "redemption_chain",
		"redemption_approved_at", "redemption_approval_reference", "created_at", "updated_at",
	}
}

func (model Contract) Table() string {
	return TableName
}

func (model Contract) GetID() any {
	return model.ID
}

func (model *Contract) SetID(id any) {
	model.ID = id.(int)
}

func (model *Contract) SetDB(db db.Repository) {
	model.db = db
}

func (model Contract) Save(ctx context.Context) error {
	if model.ID == 0 {
		err := errors.Errorf("%s: Contract %d", orm.ErrEmptyID, model.ID)
		logger.FromCtx(ctx, pkgName).Error().Stack().Err(err).Msg(orm.ErrEmptyID.Error())
		return err
	}

	updates := map[string]any{}
	for _, field := range model.updatedFields {
		updates[field] = model.GetValueByTag(field)
	}
	updated, err := orm.Update[Contract](
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
		err := errors.Errorf("%s: Contract %d", orm.ErrNoRowsAffected, model.ID)
		logger.FromCtx(ctx, pkgName).Error().Stack().Err(err).Msg(orm.ErrNoRowsAffected.Error())
		return err
	}
	return nil
}
