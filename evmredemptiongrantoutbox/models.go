package evmredemptiongrantoutbox

import (
	"github.com/global-torque/go-common/orm/v2/pgtype"
	"github.com/google/uuid"
)

// GrantOutbox is the transactional provider-work item consumed by the
// existing EVM reconciliation worker.
type GrantOutbox struct {
	ID                 int64              `db:"id" json:"id"`
	CommandUUID        uuid.UUID          `db:"command_uuid" json:"command_uuid"`
	Action             string             `db:"action" json:"action"`
	Payload            any                `db:"payload" json:"payload"`
	PayloadHash        string             `db:"payload_hash" json:"payload_hash"`
	State              string             `db:"state" json:"state"`
	Attempts           int                `db:"attempts" json:"attempts"`
	LeaseOwner         *string            `db:"lease_owner" json:"lease_owner,omitempty"`
	LeaseDeadline      pgtype.Timestamptz `db:"lease_deadline" json:"lease_deadline,omitempty"`
	NextAttemptAt      pgtype.Timestamptz `db:"next_attempt_at" json:"next_attempt_at"`
	ProviderActivityID *string            `db:"provider_activity_id" json:"provider_activity_id,omitempty"`
	ProviderEvidence   any                `db:"provider_evidence" json:"provider_evidence"`
	FailureCode        *string            `db:"failure_code" json:"failure_code,omitempty"`
	CreatedAt          pgtype.Timestamptz `db:"created_at" json:"created_at"`
	CompletedAt        pgtype.Timestamptz `db:"completed_at" json:"completed_at,omitempty"`
}

func (model GrantOutbox) ToJSON() map[string]any {
	return map[string]any{
		"id": model.ID, "command_uuid": model.CommandUUID, "action": model.Action,
		"payload": model.Payload, "payload_hash": model.PayloadHash, "state": model.State,
		"attempts": model.Attempts, "lease_owner": model.LeaseOwner, "lease_deadline": model.LeaseDeadline,
		"next_attempt_at": model.NextAttemptAt, "provider_activity_id": model.ProviderActivityID,
		"provider_evidence": model.ProviderEvidence, "failure_code": model.FailureCode,
		"created_at": model.CreatedAt, "completed_at": model.CompletedAt,
	}
}
func (model GrantOutbox) Fields() []string {
	return []string{"id", "command_uuid", "action", "payload", "payload_hash", "state", "attempts", "lease_owner", "lease_deadline", "next_attempt_at", "provider_activity_id", "provider_evidence", "failure_code", "created_at", "completed_at"}
}
func (model GrantOutbox) Table() string { return "evm_redemption_grant_outbox" }
func (model GrantOutbox) GetID() any    { return model.ID }
func (model *GrantOutbox) SetID(id any) { model.ID = id.(int64) }
