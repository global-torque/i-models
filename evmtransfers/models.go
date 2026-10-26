package evmtransfers

import (
	"context"

	"github.com/global-torque/go-common/db/v2"
	"github.com/global-torque/go-common/logger/v2"
	"github.com/global-torque/go-common/orm/v2"
	"github.com/global-torque/go-common/orm/v2/pgtype"
	"github.com/pkg/errors"
)

// Transfer is an object representing the database table.
type Transfer struct {
	ID             int  `db:"id" json:"id" yaml:"id"`
	UserID         *int `db:"user_id" json:"user_id,omitempty" yaml:"user_id,omitempty"`
	TokenID        *int `db:"token_id" json:"token_id" yaml:"token_id"`
	DestWalletID   *int `db:"dest_wallet_id" json:"dest_wallet_id" yaml:"dest_wallet_id"`
	SourceWalletID *int `db:"source_wallet_id" json:"source_wallet_id" yaml:"source_wallet_id"`
	InvestmentID   *int `db:"investment_id" json:"investment_id" yaml:"investment_id"`

	Type          string  `db:"type" json:"type" yaml:"type"`
	Amount        string  `db:"amount" json:"amount" yaml:"amount"`
	Network       string  `db:"network" json:"network" yaml:"network"`
	Status        *string `db:"status" json:"status,omitempty" yaml:"status,omitempty"`
	TransactionTX string  `db:"transaction_tx" json:"transaction_tx" yaml:"transaction_tx"`

	CreatedAt pgtype.Timestamptz `db:"created_at" json:"created_at" yaml:"created_at"`
	UpdatedAt pgtype.Timestamptz `db:"updated_at" json:"updated_at" yaml:"updated_at"`

	updatedFields []string       `db:"-" json:"-"`
	fns           map[string]any `db:"-" json:"-"`
	db            db.Repository  `db:"-" json:"-"`
}

func New(db db.Repository) *Transfer {
	return &Transfer{
		db:  db,
		fns: map[string]any{},
	}
}

func (model Transfer) GetField(name string) any {
	switch name {
	case "ID":
		return model.ID
	case "UserID":
		return model.UserID
	case "DestWalletID":
		return model.DestWalletID
	case "SourceWalletID":
		return model.SourceWalletID
	case "InvestmentID":
		return model.InvestmentID
	case "TransactionTX":
		return model.TransactionTX
	case "Status":
		return model.Status
	case "CreatedAt":
		return model.CreatedAt
	case "UpdatedAt":
		return model.UpdatedAt
	}
	return nil
}

func (model Transfer) GetValueByTag(name string) any {
	switch name {
	case "id":
		return model.ID
	case "user_id":
		return model.UserID
	case "token_id":
		return model.TokenID
	case "dest_wallet_id":
		return model.DestWalletID
	case "source_wallet_id":
		return model.SourceWalletID
	case "investment_id":
		return model.InvestmentID
	case "type":
		return model.Type
	case "amount":
		return model.Amount
	case "network":
		return model.Network
	case "transaction_tx":
		return model.TransactionTX
	case "status":
		return model.Status
	case "created_at":
		return model.CreatedAt
	case "updated_at":
		return model.UpdatedAt
	}
	return nil
}

func (model Transfer) ToJSON() map[string]any {
	res := map[string]any{}
	fields := model.Fields()
	for _, key := range fields {
		res[key] = model.GetValueByTag(key)
	}
	return res
}

func (model Transfer) Fields() []string {
	return []string{
		"id", "user_id", "token_id", "dest_wallet_id", "source_wallet_id",
		"investment_id", "type", "amount", "network", "status", "transaction_tx",
		"created_at", "updated_at",
	}
}

func (model Transfer) Table() string {
	return TableName
}

func (model Transfer) GetID() any {
	return model.ID
}

func (model *Transfer) SetID(id any) {
	model.ID = id.(int)
}

func (model *Transfer) SetDB(db db.Repository) {
	model.db = db
}

func (model Transfer) Save(ctx context.Context) error {
	if model.ID == 0 {
		err := errors.Errorf("%s: Transfer %d", orm.ErrEmptyID, model.ID)
		logger.FromCtx(ctx, pkgName).Error().Stack().Err(err).Msg(orm.ErrEmptyID.Error())
		return err
	}

	updates := map[string]any{}
	for _, field := range model.updatedFields {
		updates[field] = model.GetValueByTag(field)
	}
	updated, err := orm.Update[Transfer](
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
		err := errors.Errorf("%s: Transfer %d", orm.ErrNoRowsAffected, model.ID)
		logger.FromCtx(ctx, pkgName).Error().Stack().Err(err).Msg(orm.ErrNoRowsAffected.Error())
		return err
	}
	return nil
}
