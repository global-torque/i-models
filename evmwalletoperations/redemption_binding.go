package evmwalletoperations

import "github.com/global-torque/go-common/orm/v2/pgtype"

// RedemptionOperationBinding is the grant/evidence fence carried by an
// evm_wallet_operations row. It is a value object rather than a child table;
// operation history remains owned by WalletOperation.
type RedemptionOperationBinding struct {
	GrantID                *int64
	GrantVersion           *int64
	GrantScopeHash         *string
	AuthorizationEpoch     *int64
	AuthorizationExpiresAt pgtype.Timestamptz
	EvidenceDigest         *string
	ArmedAt                pgtype.Timestamptz
}
