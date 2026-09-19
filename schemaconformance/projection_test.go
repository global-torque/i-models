package schemaconformance_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/global-torque/go-common/orm/v2"
	"github.com/global-torque/go-common/orm/v2/pgtype"
	"github.com/jackc/pgx/v5"
	"github.com/webdevelop-pro/go-common/logger"

	"github.com/webdevelop-pro/i-models/evmwalletoperationeffects"
	"github.com/webdevelop-pro/i-models/evmwalletoperations"
	"github.com/webdevelop-pro/i-models/fundnavrecords"
	"github.com/webdevelop-pro/i-models/investments"
	"github.com/webdevelop-pro/i-models/offers"
)

type transactionRepository struct {
	pgx.Tx
}

func (transactionRepository) Lg() logger.Logger { return logger.Logger{} }

type offerProjection struct {
	ID                    int                `db:"id"`
	Status                string             `db:"status"`
	Data                  map[string]any     `db:"data"`
	InvestorEligibility   []string           `db:"investor_eligibility"`
	MarketedJurisdictions []string           `db:"marketed_jurisdictions"`
	RiskDisclosures       string             `db:"risk_disclosures"`
	StartAt               pgtype.Timestamptz `db:"start_at"`
	MinInvestment         string             `db:"min_investment"`
	TotalShares           string             `db:"total_shares"`
	PricePerShare         string             `db:"price_per_share"`
	DataRoomGroupID       *int               `db:"data_room_group_id"`
	FundStructure         string             `db:"fund_structure"`
}

func TestOfferProjectionAgainstPostgreSQL(t *testing.T) {
	if os.Getenv("I_MODELS_SCHEMA_TEST") != "1" {
		t.Skip("set I_MODELS_SCHEMA_TEST=1 to run PostgreSQL projection parity")
	}

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, schemaDatabaseURL())
	if err != nil {
		t.Fatalf("connect to projection database: %v", err)
	}
	t.Cleanup(func() {
		if err := conn.Close(ctx); err != nil {
			t.Errorf("close projection database: %v", err)
		}
	})

	tx, err := conn.Begin(ctx)
	if err != nil {
		t.Fatalf("begin projection transaction: %v", err)
	}
	t.Cleanup(func() {
		if err := tx.Rollback(ctx); err != nil && err != pgx.ErrTxClosed {
			t.Errorf("rollback projection transaction: %v", err)
		}
	})
	repo := transactionRepository{Tx: tx}

	var userID int
	err = tx.QueryRow(ctx, `
		INSERT INTO user_users(id, password, email, first_name, last_name, identity_id)
		VALUES (-700001, '', 'projection@example.test', 'Projection', 'Test', 'projection-schema-test')
		RETURNING id
	`).Scan(&userID)
	if err != nil {
		t.Fatalf("insert projection user: %v", err)
	}

	var offerID int
	err = tx.QueryRow(ctx, `
		INSERT INTO offer_offers(
			user_id, name, slug, min_investment, status, data, investor_eligibility,
			marketed_jurisdictions, risk_disclosures, start_at, total_shares, price_per_share
		) VALUES (
			$1, $2, $3, $4, $5, $6::jsonb, $7::jsonb, $8::jsonb, $9, NULL, $10, $11
		) RETURNING id
	`,
		userID,
		"Projection offer",
		"projection-offer",
		"0.000000000000000001",
		"legal_review",
		`{"nc_stamping_text":"Stamped \"quoted\" ✓"}`,
		`["accredited","qualified_investor"]`,
		`["US","EU"]`,
		"Risk disclosure with Unicode ✓",
		125.25,
		2.5,
	).Scan(&offerID)
	if err != nil {
		t.Fatalf("insert projection offer: %v", err)
	}

	projection, err := orm.RetrieveOneAs[offers.OfferOffer, offerProjection](
		ctx,
		repo,
		sq.Eq{"id": offerID},
	)
	if err != nil {
		t.Fatalf("retrieve offer projection: %v", err)
	}
	if projection.Status != "legal_review" || projection.StartAt.Status != pgtype.Null {
		t.Fatalf("unexpected enum or nullable timestamp: %#v", projection)
	}
	if len(projection.InvestorEligibility) != 2 || len(projection.MarketedJurisdictions) != 2 {
		t.Fatalf("unexpected JSON arrays: %#v", projection)
	}
	if projection.Data["nc_stamping_text"] != "Stamped \"quoted\" ✓" {
		t.Fatalf("unexpected JSON data: %#v", projection.Data)
	}
	if projection.RiskDisclosures != "Risk disclosure with Unicode ✓" {
		t.Fatalf("unexpected risk disclosure: %q", projection.RiskDisclosures)
	}
	if projection.MinInvestment != "0.000000000000000001" ||
		projection.TotalShares != "125.250000000000000000" ||
		projection.PricePerShare != "2.500000000000000000" {
		t.Fatalf(
			"fractional offer values did not retain exact PostgreSQL decimals: minimum=%q shares=%q price=%q",
			projection.MinInvestment,
			projection.TotalShares,
			projection.PricePerShare,
		)
	}
	if projection.FundStructure != "open_ended" {
		t.Fatalf("unexpected default fund structure: %q", projection.FundStructure)
	}
	if projection.DataRoomGroupID != nil {
		t.Fatalf("legacy offer unexpectedly has data-room binding: %v", *projection.DataRoomGroupID)
	}

	empty, err := orm.RetrieveAllAs[offers.OfferOffer, offerProjection](
		ctx,
		repo,
		sq.Eq{"id": -1},
	)
	if err != nil {
		t.Fatalf("retrieve empty offer projections: %v", err)
	}
	if empty == nil || len(empty) != 0 {
		t.Fatalf("expected non-nil empty projections, got %#v", empty)
	}

	var profileID int
	err = tx.QueryRow(ctx, `
		INSERT INTO investment_profiles(id, user_id)
		VALUES (-700001, $1)
		RETURNING id
	`, userID).Scan(&profileID)
	if err != nil {
		t.Fatalf("insert projection profile: %v", err)
	}

	var investmentID int
	err = tx.QueryRow(ctx, `
		INSERT INTO investment_investments(
			id, user_id, created_by, offer_id, profile_id, amount, price_per_share,
			number_of_shares
		) VALUES (
			-700001, $1, $1, $2, $3, 123.456789, 2.5, 49.3827156
		)
		RETURNING id
	`, userID, offerID, profileID).Scan(&investmentID)
	if err != nil {
		t.Fatalf("insert projection investment: %v", err)
	}

	investment, err := orm.RetrieveOne[
		investments.InvestmentInvestment,
		*investments.InvestmentInvestment,
	](ctx, repo, sq.Eq{"id": investmentID})
	if err != nil {
		t.Fatalf("retrieve investment projection: %v", err)
	}
	if investment.CreatedBy != userID ||
		investment.Amount == nil ||
		*investment.Amount != "123.456789000000000000" ||
		investment.PricePerShare != "2.500000000000000000" ||
		investment.NumberOfShares == nil ||
		*investment.NumberOfShares != "49.382715600000000000" {
		t.Fatalf(
			"investment decimals were not projected exactly: amount=%v price=%q shares=%v",
			investment.Amount,
			investment.PricePerShare,
			investment.NumberOfShares,
		)
	}

	_, err = tx.Exec(ctx, `
		INSERT INTO investment_investments(
			id, user_id, created_by, offer_id, profile_id, amount, price_per_share,
			number_of_shares
		) VALUES (
			-700002, $1, $1, $2, $3, 123.456789, 2.5, NULL
		)
	`, userID, offerID, profileID)
	if err != nil {
		t.Fatalf("insert unpriced projection investment: %v", err)
	}
	unpricedInvestment, err := orm.RetrieveOne[
		investments.InvestmentInvestment,
		*investments.InvestmentInvestment,
	](ctx, repo, sq.Eq{"id": -700002})
	if err != nil {
		t.Fatalf("retrieve unpriced investment projection: %v", err)
	}
	if unpricedInvestment.NumberOfShares != nil {
		t.Fatalf("unpriced investment shares = %q, want nil", *unpricedInvestment.NumberOfShares)
	}

	const externalStrategyAssetsRaw = "12345678901234567890123456789012345678901234567890"
	var navID int64
	err = tx.QueryRow(ctx, `
		INSERT INTO fund_nav_records(
			id, offer_id, nav, recorded_at, external_strategy_assets_raw
		) VALUES (
			-700001, $1, 1.234567890123456789, clock_timestamp(), $2
		)
		RETURNING id
	`, offerID, externalStrategyAssetsRaw).Scan(&navID)
	if err != nil {
		t.Fatalf("insert projection NAV: %v", err)
	}

	nav, err := orm.RetrieveOne[
		fundnavrecords.FundNAVRecord,
		*fundnavrecords.FundNAVRecord,
	](ctx, repo, sq.Eq{"id": navID})
	if err != nil {
		t.Fatalf("retrieve NAV projection: %v", err)
	}
	if nav.NAV != "1.234567890123456789" ||
		nav.ExternalStrategyAssetsRaw != externalStrategyAssetsRaw {
		t.Fatalf(
			"NAV values were not projected exactly: nav=%q strategy=%q",
			nav.NAV,
			nav.ExternalStrategyAssetsRaw,
		)
	}

	const claimableSharesRaw = "9876543210987654321098765432109876543210"
	var redemptionID int64
	err = tx.QueryRow(ctx, `
		INSERT INTO investment_redemptions(
			id, offer_id, profile_id, investment_id, vault_request_origin,
			idempotency_key, share_amount_raw, claimable_assets_raw,
			claimable_shares_raw, claimable_at
		) VALUES (
			-700001, $1, $2, $3, 'chain',
			'projection-redemption', $4, 0, $4, clock_timestamp()
		)
		RETURNING id
	`, offerID, profileID, investmentID, claimableSharesRaw).Scan(&redemptionID)
	if err != nil {
		t.Fatalf("insert projection redemption: %v", err)
	}

	redemption, err := orm.RetrieveOne[
		investments.InvestmentRedemption,
		*investments.InvestmentRedemption,
	](ctx, repo, sq.Eq{"id": redemptionID})
	if err != nil {
		t.Fatalf("retrieve redemption projection: %v", err)
	}
	if redemption.ShareAmountRaw != claimableSharesRaw ||
		redemption.ClaimableSharesRaw != claimableSharesRaw ||
		redemption.ClaimableAssetsRaw != "0" {
		t.Fatalf(
			"redemption exact/share-derived Claimable projection mismatch: %#v",
			redemption,
		)
	}

	const (
		scopedCallSignature = "0xdcd240c0b26ed449aae1a951dc6b5047eda64e79ebeea2186944a090f1fd3c8f"
		scopedCallDigest    = "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
		executorTopic       = "0x0000000000000000000000001111111111111111111111111111111111111111"
		targetTopic         = "0x0000000000000000000000002222222222222222222222222222222222222222"
		executorAddress     = "0x1111111111111111111111111111111111111111"
		targetAddress       = "0x2222222222222222222222222222222222222222"
		zeroValueData       = "0x0000000000000000000000000000000000000000000000000000000000000000"
	)
	var operationID int
	err = tx.QueryRow(ctx, `
		INSERT INTO evm_wallet_operations(
			id, user_id, profile_id, chain, wallet_address, token_ticker,
			token_address, token_decimals, amount, amount_raw, type, status,
			source, submission_status, counterparty_address, reorg_count,
			idempotency_key, transaction_from_address
		) VALUES (
			-700001, $1, $2, 'ethereum-sepolia', $3, 'USDC',
			$4, 6, 0, '0', 'deposit', 'submitted',
			'chain', 'submitted', $4, 3, 'projection-scoped-call', $3
		)
		RETURNING id
	`, userID, profileID, executorAddress, targetAddress).Scan(&operationID)
	if err != nil {
		t.Fatalf("insert ScopedCallExecuted parent operation: %v", err)
	}

	operation, err := orm.RetrieveOne[
		evmwalletoperations.WalletOperation,
		*evmwalletoperations.WalletOperation,
	](ctx, repo, sq.Eq{"id": operationID})
	if err != nil {
		t.Fatalf("retrieve transaction sender projection: %v", err)
	}
	if operation.TransactionFromAddress == nil ||
		*operation.TransactionFromAddress != executorAddress {
		t.Fatalf(
			"transaction sender did not project exactly: %#v",
			operation.TransactionFromAddress,
		)
	}

	var effectID int64
	err = tx.QueryRow(ctx, `
		INSERT INTO evm_wallet_operation_effects(
			id, operation_id, chain, effect_kind, effect_index,
			wallet_address, contract_address, event_signature, event_topics,
			event_data, subject_address, receipt_generation, canonical
		) VALUES (
			-700001, $1, 'ethereum-sepolia', 'contract_event', 17,
			$2, $2, $3::text, jsonb_build_array(
				$3::text, $4::text, $5::text, $6::text
			),
			$7, $2, 3, true
		)
		RETURNING id
	`,
		operationID,
		executorAddress,
		scopedCallSignature,
		scopedCallDigest,
		executorTopic,
		targetTopic,
		zeroValueData,
	).Scan(&effectID)
	if err != nil {
		t.Fatalf("insert ScopedCallExecuted effect: %v", err)
	}

	effect, err := orm.RetrieveOne[
		evmwalletoperationeffects.WalletOperationEffect,
		*evmwalletoperationeffects.WalletOperationEffect,
	](ctx, repo, sq.Eq{"id": effectID})
	if err != nil {
		t.Fatalf("retrieve ScopedCallExecuted effect: %v", err)
	}
	topicsJSON, err := json.Marshal(effect.EventTopics)
	if err != nil {
		t.Fatalf("marshal projected ScopedCallExecuted topics: %v", err)
	}
	wantTopicsJSON := `["` + scopedCallSignature + `","` + scopedCallDigest +
		`","` + executorTopic + `","` + targetTopic + `"]`
	if effect.EffectIndex != 17 ||
		effect.ReceiptGeneration != 3 ||
		!effect.Canonical ||
		effect.EventSignature == nil ||
		*effect.EventSignature != scopedCallSignature ||
		effect.EventData == nil ||
		*effect.EventData != zeroValueData ||
		string(topicsJSON) != wantTopicsJSON {
		t.Fatalf(
			"ScopedCallExecuted evidence did not project exactly: effect=%#v topics=%s",
			effect,
			topicsJSON,
		)
	}
}
