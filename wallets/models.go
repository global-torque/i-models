package wallets

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/global-torque/go-common/db/v2"
	"github.com/global-torque/go-common/logger/v2"
	"github.com/global-torque/go-common/orm/v2"
	"github.com/global-torque/go-common/orm/v2/pgtype"
	"github.com/pkg/errors"
	"github.com/webdevelop-pro/i-models/historylogs"
	"github.com/webdevelop-pro/i-models/logs"
	"github.com/webdevelop-pro/i-models/notifications"
)

const Table = "wallet_wallets"
const pkgName = "models/wallets"

// Wallet is an object representing the database table.
type Wallet struct {
	ID     int  `db:"id" json:"id" yaml:"id"`
	UserID *int `db:"user_id" json:"user_id,omitempty" yaml:"user_id,omitempty"`

	EntityID        *string            `db:"entity_id" json:"entity_id,omitempty"`
	EntityBalanceID *string            `db:"entity_balance_id" json:"entity_balance_id,omitempty"`
	Balance         float64            `db:"balance" json:"balance" yaml:"balance"`
	IncBalance      float64            `db:"inc_balance" json:"inc_balance" yaml:"inc_balance"`
	OutBalance      float64            `db:"out_balance" json:"out_balance" yaml:"out_balance"`
	Status          WalletStatusT      `db:"status" json:"status" yaml:"status"`
	CreatedAt       pgtype.Timestamptz `db:"created_at" json:"created_at" yaml:"created_at"`
	UpdatedAt       pgtype.Timestamptz `db:"updated_at" json:"updated_at" yaml:"updated_at"`

	ObjectID      string `db:"object_id" json:"object_id" yaml:"object_id"`
	ContentTypeID int    `db:"content_type_id" json:"content_type_id" yaml:"content_type_id"`

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

func (model Wallet) Fields() []string {
	return []string{
		"id",
		"user_id",
		"entity_id",
		"entity_balance_id",
		"balance",
		"inc_balance",
		"out_balance",
		"status",
		"created_at",
		"updated_at",
		"object_id",
		"content_type_id",
	}
}

func (model Wallet) GetField(name string) any {
	switch name {
	case "ID":
		return model.ID
	case "UserID":
		return model.UserID
	case "EntityID":
		return model.EntityID
	case "EntityBalanceID":
		return model.EntityBalanceID
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
	case "ObjectID":
		return model.ObjectID
	case "ContentTypeID":
		return model.ContentTypeID
	}
	return nil
}

func (model Wallet) GetValueByTag(field string) any {
	switch field {
	case "id":
		return model.ID
	case "user_id":
		return model.UserID
	case "entity_id":
		return model.EntityID
	case "entity_balance_id":
		return model.EntityBalanceID
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
	case "object_id":
		return model.ObjectID
	case "content_type_id":
		return model.ContentTypeID
	}
	return nil
}

func (model *Wallet) SetOutBalance(val float64) {
	model.OutBalance = val
	model.updatedFields = append(model.updatedFields, "out_balance")
}

func (model *Wallet) SetIncBalance(val float64) {
	model.IncBalance = val
	model.updatedFields = append(model.updatedFields, "inc_balance")
}

func (model *Wallet) SetIncBalanceFN(val float64) {
	model.IncBalance += val
	model.fns["fn_inc_balance"] = sq.Expr("inc_balance + ?", val)
	model.updatedFields = append(model.updatedFields, "fn_inc_balance")
}

func (model *Wallet) SetOutBalanceFN(val float64) {
	model.IncBalance += val
	model.fns["fn_out_balance"] = sq.Expr("out_balance + ?", val)
	model.updatedFields = append(model.updatedFields, "fn_out_balance")
}

func (model *Wallet) SetBalanceFN(val float64) {
	model.IncBalance += val
	model.fns["fn_balance"] = sq.Expr("balance + ?", val)
	model.updatedFields = append(model.updatedFields, "fn_balance")
}

func (model *Wallet) SetBalance(val float64) {
	model.Balance = val
	model.updatedFields = append(model.updatedFields, "balance")
}

func (model Wallet) Save(ctx context.Context) error {
	if model.ID == 0 {
		err := errors.Errorf("%s: wallet %d", orm.ErrEmptyID, model.ID)
		logger.FromCtx(ctx, pkgName).Error().Stack().Err(err).Msg(orm.ErrEmptyID.Error())
		return err
	}

	updates := map[string]any{}
	for _, field := range model.updatedFields {
		if field[0:3] == "fn_" {
			updates[field[3:]] = model.fns[field]
		} else {
			updates[field] = model.GetValueByTag(field)
		}
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
		err := errors.Errorf("%s: wallet %d", orm.ErrNoRowsAffected, model.ID)
		logger.FromCtx(ctx, pkgName).Error().Stack().Err(err).Msg(orm.ErrNoRowsAffected.Error())
		return err
	}
	if model.UserID != nil {
		model.DefaultPostUpdate(ctx, *model.UserID, updates)
	}
	return nil
}

func Get(ctx context.Context, db db.Repository, where map[string]any) (*Wallet, error) {
	model, err := orm.RetrieveOne[Wallet](
		ctx,
		db,
		sq.Eq(where),
	)
	if err != nil {
		err = errors.Wrapf(err, "cannot get model")
		return nil, err
	}
	model.db = db
	model.fns = map[string]any{}
	return model, nil
}

func GetByID(ctx context.Context, db db.Repository, id int) (*Wallet, error) {
	model, err := orm.RetrieveOne[Wallet](
		ctx,
		db,
		sq.Eq{
			"id": id,
		},
	)
	if err != nil {
		err = errors.Wrapf(err, "cannot get model")
		return nil, err
	}
	model.db = db
	model.fns = map[string]any{}
	return model, nil
}

func (model Wallet) ToMap() map[string]any {
	res := map[string]any{}
	fields := model.Fields()
	for _, key := range fields {
		res[key] = model.GetValueByTag(key)
	}
	return res
}

func (model Wallet) Table() string {
	return Table
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

func (model *Wallet) NotificationUserUpdate(ctx context.Context, userID int, data map[string]any) error {
	if userID != 0 {
		// what user see for wallet: balance, inc/out balance, status
		watchFields := []string{
			"balance",
			"inc_balance",
			"out_balance",
			"status",
		}
		notifData := map[string]any{}
		for _, field := range watchFields {
			val, ok := data[field]
			if ok {
				notifData[field] = val
			}
		}

		if len(notifData) > 0 {
			_, err := orm.Create[notifications.Notification](
				ctx,
				model.db,
				map[string]any{
					"user_id": userID,
					"content": "",
					"data":    notifData,
				},
			)
			return err
		}
	}

	return nil
}

func (model *Wallet) HistoryLogUpdate(ctx context.Context, userID int, data map[string]any) error {
	// cannot save anything without userID
	if userID != 0 {
		upFields := ""
		for key, _ := range data {
			upFields = fmt.Sprintf("%s%s,", upFields, key)
		}
		if len(upFields) > 0 {
			upFields = upFields[0 : len(upFields)-1]
		}

		ContentID, err := logs.GetContentID(ctx, model.db, AppLabel, ModelName)
		if err != nil {
			return err
		}

		data = map[string]any{
			"action_flag":     historylogs.ActionChange,
			"change_message":  fmt.Sprintf(`{"changed":{"fields": "%s"}}`, upFields),
			"object_repr":     fmt.Sprintf("%s: %d", ModelName, model.ID),
			"action_time":     time.Now(),
			"user_id":         userID,
			"content_type_id": ContentID,
			"object_id":       fmt.Sprintf("%d", model.ID),
		}

		_, err = orm.Create[historylogs.HistoryLog](
			ctx,
			model.db,
			data,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (model *Wallet) DefaultPostUpdate(ctx context.Context, userID int, updatedData map[string]any) error {
	// model.NotificationUserUpdate(ctx, userID, updatedData)
	model.HistoryLogUpdate(ctx, userID, updatedData)
	return nil
}
