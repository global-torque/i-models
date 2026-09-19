package evmredemptiongrantcommands

import (
	"github.com/global-torque/go-common/orm/v2/pgtype"
	"github.com/google/uuid"
)

// RedemptionGrantCommand is the durable prepare/confirm/revoke saga record.
type RedemptionGrantCommand struct {
	ID                 int64              `db:"id" json:"id"`
	CommandUUID        uuid.UUID          `db:"command_uuid" json:"command_uuid"`
	OfferID            int                `db:"offer_id" json:"offer_id"`
	CommandKind        KindT              `db:"command_kind" json:"command_kind"`
	IdempotencyKey     string             `db:"idempotency_key" json:"idempotency_key"`
	PayloadHash        string             `db:"payload_hash" json:"payload_hash"`
	GrantID            *int64             `db:"grant_id" json:"grant_id,omitempty"`
	ChallengeID        uuid.UUID          `db:"challenge_id" json:"challenge_id"`
	PrepareCommandUUID *uuid.UUID         `db:"prepare_command_uuid" json:"prepare_command_uuid,omitempty"`
	OperationID        *int64             `db:"operation_id" json:"operation_id,omitempty"`
	SignablePayload    any                `db:"signable_payload" json:"signable_payload"`
	State              StateT             `db:"state" json:"state"`
	Attempts           int                `db:"attempts" json:"attempts"`
	LeaseOwner         *string            `db:"lease_owner" json:"lease_owner,omitempty"`
	LeaseDeadline      pgtype.Timestamptz `db:"lease_deadline" json:"lease_deadline,omitempty"`
	ChallengeExpiresAt pgtype.Timestamptz `db:"challenge_expires_at" json:"challenge_expires_at,omitempty"`
	AuthorizationEpoch int64              `db:"authorization_epoch" json:"authorization_epoch"`
	ProviderPlanHash   *string            `db:"provider_plan_hash" json:"provider_plan_hash,omitempty"`
	CreatedAt          pgtype.Timestamptz `db:"created_at" json:"created_at"`
	ExpiresAt          pgtype.Timestamptz `db:"expires_at" json:"expires_at,omitempty"`
	FailureCode        *string            `db:"failure_code" json:"failure_code,omitempty"`
	ConsumedAt         pgtype.Timestamptz `db:"consumed_at" json:"consumed_at,omitempty"`
	SubmissionResponse any                `db:"submission_response" json:"submission_response"`
	UpdatedAt          pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
	RetryAt            pgtype.Timestamptz `db:"retry_at" json:"retry_at,omitempty"`
}

func (model RedemptionGrantCommand) ToJSON() map[string]any {
	return map[string]any{
		"id": model.ID, "command_uuid": model.CommandUUID, "offer_id": model.OfferID, "command_kind": model.CommandKind,
		"idempotency_key": model.IdempotencyKey, "payload_hash": model.PayloadHash,
		"grant_id": model.GrantID, "challenge_id": model.ChallengeID, "prepare_command_uuid": model.PrepareCommandUUID,
		"operation_id": model.OperationID, "signable_payload": model.SignablePayload,
		"state": model.State, "attempts": model.Attempts, "lease_owner": model.LeaseOwner,
		"lease_deadline": model.LeaseDeadline, "challenge_expires_at": model.ChallengeExpiresAt,
		"authorization_epoch": model.AuthorizationEpoch, "provider_plan_hash": model.ProviderPlanHash,
		"created_at": model.CreatedAt, "expires_at": model.ExpiresAt, "failure_code": model.FailureCode, "consumed_at": model.ConsumedAt,
		"submission_response": model.SubmissionResponse, "updated_at": model.UpdatedAt, "retry_at": model.RetryAt,
	}
}

func (model RedemptionGrantCommand) Fields() []string {
	return []string{"id", "command_uuid", "offer_id", "command_kind", "idempotency_key", "payload_hash", "grant_id", "challenge_id", "prepare_command_uuid", "operation_id", "signable_payload", "state", "attempts", "lease_owner", "lease_deadline", "challenge_expires_at", "authorization_epoch", "provider_plan_hash", "created_at", "expires_at", "failure_code", "consumed_at", "submission_response", "updated_at", "retry_at"}
}
func (model RedemptionGrantCommand) Table() string { return "evm_redemption_grant_commands" }
func (model RedemptionGrantCommand) GetID() any    { return model.ID }
func (model *RedemptionGrantCommand) SetID(id any) { model.ID = id.(int64) }
