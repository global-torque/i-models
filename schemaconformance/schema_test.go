package schemaconformance_test

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5"

	"github.com/webdevelop-pro/i-models/distributions"
	"github.com/webdevelop-pro/i-models/emails"
	"github.com/webdevelop-pro/i-models/evmallowancesequences"
	"github.com/webdevelop-pro/i-models/evmchainaccounts"
	"github.com/webdevelop-pro/i-models/evmcontracts"
	"github.com/webdevelop-pro/i-models/evmredemptionapprovalpreflightentries"
	"github.com/webdevelop-pro/i-models/evmredemptionapprovalpreflightevidence"
	"github.com/webdevelop-pro/i-models/evmredemptionfulfillmentgrants"
	"github.com/webdevelop-pro/i-models/evmredemptiongrantcommands"
	"github.com/webdevelop-pro/i-models/evmredemptiongrantoutbox"
	"github.com/webdevelop-pro/i-models/evmredemptionrequestevidence"
	"github.com/webdevelop-pro/i-models/evmvaultliquidityquarantines"
	"github.com/webdevelop-pro/i-models/evmwalletbalances"
	"github.com/webdevelop-pro/i-models/evmwalletoperationeffects"
	"github.com/webdevelop-pro/i-models/evmwalletoperations"
	"github.com/webdevelop-pro/i-models/evmwallets"
	"github.com/webdevelop-pro/i-models/filers"
	"github.com/webdevelop-pro/i-models/fundingsources"
	"github.com/webdevelop-pro/i-models/fundnavrecords"
	"github.com/webdevelop-pro/i-models/historylogs"
	"github.com/webdevelop-pro/i-models/investments"
	"github.com/webdevelop-pro/i-models/logs"
	"github.com/webdevelop-pro/i-models/notifications"
	"github.com/webdevelop-pro/i-models/offers"
	"github.com/webdevelop-pro/i-models/pubsublogs"
	"github.com/webdevelop-pro/i-models/transactions"
	"github.com/webdevelop-pro/i-models/userinvitations"
	"github.com/webdevelop-pro/i-models/users"
	"github.com/webdevelop-pro/i-models/wallets"
)

func TestStablecoinDigestFixtureIsSharedByteExact(t *testing.T) {
	contents, err := os.ReadFile("stablecoin_redemption_digest_fixtures.json")
	if err != nil {
		t.Fatal(err)
	}
	hash := fmt.Sprintf("%x", sha256.Sum256(contents))
	const expected = "cfbc75bdc7d46eb86c583c570f66f586ee5bbed9fc470ca0aad52cfde8079150"
	if hash != expected {
		t.Fatalf("shared canonical fixture hash = %s, want %s", hash, expected)
	}
	var fixture struct {
		InvalidRaw []struct {
			Name string `json:"name"`
			Raw  string `json:"raw"`
		} `json:"invalidRaw"`
	}
	if err := json.Unmarshal(contents, &fixture); err != nil {
		t.Fatal(err)
	}
	wantNames := []string{
		"duplicate-key", "negative-zero", "fraction", "exponent",
		"unsafe-integer", "lone-surrogate",
	}
	if len(fixture.InvalidRaw) != len(wantNames) {
		t.Fatalf("invalidRaw fixture entries=%d, want %d", len(fixture.InvalidRaw), len(wantNames))
	}
	for index, wantName := range wantNames {
		if fixture.InvalidRaw[index].Name != wantName || fixture.InvalidRaw[index].Raw == "" {
			t.Fatalf("invalidRaw[%d]=%+v, want non-empty %q", index, fixture.InvalidRaw[index], wantName)
		}
	}
}

type tableModel interface {
	Fields() []string
	Table() string
}

type modelSpec struct {
	name  string
	model tableModel
}

type databaseColumn struct {
	Name       string `db:"column_name"`
	DataType   string `db:"data_type"`
	UDTName    string `db:"udt_name"`
	IsNullable string `db:"is_nullable"`
}

// evmtransfers is intentionally absent: evm_wallet_operations replaced the
// retired evm_transfers table.
var persistedModels = []modelSpec{ //nolint:gochecknoglobals // This is the canonical model registry.
	{name: "DistributionDistribution", model: distributions.DistributionDistribution{}},
	{name: "DistributionDistributionFiler", model: distributions.DistributionDistributionFiler{}},
	{name: "DistributionDistributionReport", model: distributions.DistributionDistributionReport{}},
	{name: "EmailEmail", model: emails.EmailEmail{}},
	{name: "AllowanceSequence", model: evmallowancesequences.AllowanceSequence{}},
	{name: "EvmChainAccount", model: evmchainaccounts.EvmChainAccount{}},
	{name: "Contract", model: evmcontracts.Contract{}},
	{name: "RedemptionApprovalPreflightEntry", model: evmredemptionapprovalpreflightentries.RedemptionApprovalPreflightEntry{}},
	{name: "RedemptionApprovalPreflightEvidence", model: evmredemptionapprovalpreflightevidence.RedemptionApprovalPreflightEvidence{}},
	{name: "RedemptionFulfillmentGrant", model: evmredemptionfulfillmentgrants.RedemptionFulfillmentGrant{}},
	{name: "RedemptionGrantCommand", model: evmredemptiongrantcommands.RedemptionGrantCommand{}},
	{name: "RedemptionGrantOutbox", model: evmredemptiongrantoutbox.GrantOutbox{}},
	{name: "RedemptionRequestEvidence", model: evmredemptionrequestevidence.RedemptionRequestEvidence{}},
	{name: "VaultLiquidityQuarantine", model: evmvaultliquidityquarantines.VaultLiquidityQuarantine{}},
	{name: "WalletBalance", model: evmwalletbalances.WalletBalance{}},
	{name: "WalletOperationEffect", model: evmwalletoperationeffects.WalletOperationEffect{}},
	{name: "WalletOperation", model: evmwalletoperations.WalletOperation{}},
	{name: "EvmWallet", model: evmwallets.Wallet{}},
	{name: "FilerFiler", model: filers.FilerFiler{}},
	{name: "WalletFundingSource", model: fundingsources.WalletFundingSource{}},
	{name: "FundNAVRecord", model: fundnavrecords.FundNAVRecord{}},
	{name: "HistoryLog", model: historylogs.HistoryLog{}},
	{name: "InvestmentInvestment", model: investments.InvestmentInvestment{}},
	{name: "InvestmentRedemption", model: investments.InvestmentRedemption{}},
	{name: "InvestmentProfile", model: investments.InvestmentProfile{}},
	{name: "LogLog", model: logs.LogLog{}},
	{name: "Notification", model: notifications.Notification{}},
	{name: "OfferComment", model: offers.OfferComment{}},
	{name: "OfferOffer", model: offers.OfferOffer{}},
	{name: "PubsubLog", model: pubsublogs.PubsubLog{}},
	{name: "Transaction", model: transactions.Transaction{}},
	{name: "UserInvitation", model: userinvitations.UserInvitation{}},
	{name: "UserUser", model: users.UserUser{}},
	{name: "Wallet", model: wallets.Wallet{}},
}

var persistedEnums = map[string][]string{ //nolint:gochecknoglobals // PostgreSQL enum contract registry.
	"accreditation_t":                           enumStrings(investments.AllAccreditationT()),
	"comment_related_t":                         enumStrings(offers.AllCommentRelatedT()),
	"comment_status_t":                          enumStrings(offers.AllCommentStatusT()),
	"distribution_filer_t":                      enumStrings(distributions.AllDistributionFilerT()),
	"distribution_report_t":                     enumStrings(distributions.AllDistributionReportT()),
	"distribution_t":                            enumStrings(distributions.AllDistributionT()),
	"email_status_t":                            enumStrings(emails.AllEmailStatusT()),
	"escrow_t":                                  enumStrings(investments.AllEscrowT()),
	"evm_erc20_allowance_sequence_purpose_t":    enumStrings(evmallowancesequences.AllPurposeT()),
	"evm_status_t":                              enumStrings(evmwallets.AllWalletStatusT()),
	"evm_wallet_account_mode_t":                 enumStrings(evmchainaccounts.AllAccountModeT()),
	"evm_wallet_chain_account_status_t":         enumStrings(evmchainaccounts.AllStatusT()),
	"evm_wallet_chain_t":                        enumStrings(evmchainaccounts.AllChainT()),
	"evm_wallet_operation_source_t":             enumStrings(evmwalletoperations.AllOperationSourceT()),
	"evm_wallet_operation_status_t":             enumStrings(evmwalletoperations.AllOperationStatusT()),
	"evm_wallet_operation_type_t":               enumStrings(evmwalletoperations.AllOperationTypeT()),
	"evm_wallet_operation_effect_direction_t":   enumStrings(evmwalletoperationeffects.AllEffectDirectionT()),
	"evm_wallet_operation_effect_kind_t":        enumStrings(evmwalletoperationeffects.AllEffectKindT()),
	"evm_redemption_fulfillment_grant_status_t": enumStrings(evmredemptionfulfillmentgrants.AllStatusT()),
	"evm_redemption_grant_command_kind_t":       enumStrings(evmredemptiongrantcommands.AllKindT()),
	"evm_redemption_preflight_coverage_t":       enumStrings(evmredemptionapprovalpreflightentries.AllCoverageT()),
	"file_type":                                 enumStrings(filers.AllType()),
	"founding_source_t":                         enumStrings(fundingsources.AllFoundingSourceT()),
	"funding_s":                                 enumStrings(investments.AllFundingS()),
	"funding_t":                                 enumStrings(investments.AllFundingT()),
	"fund_structure_t":                          enumStrings(offers.AllFundStructureT()),
	"instrument_classification_t":               enumStrings(offers.AllInstrumentClassificationT()),
	"investment_step_t":                         enumStrings(investments.AllInvestmentStepT()),
	"investment_t":                              enumStrings(investments.AllInvestmentT()),
	"investment_deposit_price_source_t":         enumStrings(investments.AllDepositPriceSourceT()),
	"investment_redemption_status_t":            enumStrings(investments.AllRedemptionStatusT()),
	"investment_vault_pricing_status_t":         enumStrings(investments.AllVaultPricingStatusT()),
	"investment_vault_request_origin_t":         enumStrings(investments.AllVaultRequestOriginT()),
	"kyc_t":                                     enumStrings(investments.AllKycT()),
	"log_type_t":                                enumStrings(logs.AllLogTypeT()),
	"notification_status_t":                     enumStrings(notifications.AllNotificationStatusT()),
	"notification_type_t":                       enumStrings(notifications.AllNotificationTypeT()),
	"offer_jurisdiction_t":                      enumStrings(offers.AllOfferJurisdictionT()),
	"offer_reg_type_t":                          enumStrings(offers.AllOfferRegTypeT()),
	"offer_security_type_t":                     enumStrings(offers.AllOfferSecurityTypeT()),
	"offer_t":                                   enumStrings(offers.AllOfferT()),
	"offering_mode_t":                           enumStrings(offers.AllOfferingModeT()),
	"payment_t":                                 enumStrings(investments.AllPaymentT()),
	"profile_t":                                 enumStrings(investments.AllProfileT()),
	"services_t":                                enumStrings(logs.AllServicesT()),
	"tokenization_engine_t":                     enumStrings(offers.AllTokenizationEngineT()),
	"tokenization_model_t":                      enumStrings(offers.AllTokenizationModelT()),
	"transactions_status_t":                     enumStrings(transactions.AllTransactionsStatusT()),
	"transactions_type_t":                       enumStrings(transactions.AllTransactionsTypeT()),
	"user_user_invitation_role_t":               enumStrings(userinvitations.AllRoles()),
	"user_user_invitation_status_t":             enumStrings(userinvitations.AllStatuses()),
	"wallet_status_t":                           enumStrings(wallets.AllWalletStatusT()),
}

func TestPersistedModelTagsAndDefaults(t *testing.T) {
	for _, spec := range persistedModels {
		t.Run(spec.name, func(t *testing.T) {
			columns := taggedColumns(t, reflect.TypeOf(spec.model))
			if len(columns) == 0 {
				t.Fatal("model has no persisted db tags")
			}

			seen := make(map[string]struct{}, len(spec.model.Fields()))
			for _, field := range spec.model.Fields() {
				column := defaultColumn(field)
				if _, ok := columns[column]; !ok {
					t.Errorf("default field %q is not a canonical model column", field)
				}
				if _, ok := seen[column]; ok {
					t.Errorf("default field %q is duplicated", field)
				}
				seen[column] = struct{}{}
			}
		})
	}
}

func TestPersistedModelsMatchDatabaseColumns(t *testing.T) {
	if os.Getenv("I_MODELS_SCHEMA_TEST") != "1" {
		t.Skip("set I_MODELS_SCHEMA_TEST=1 to compare canonical models with PostgreSQL")
	}

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, schemaDatabaseURL())
	if err != nil {
		t.Fatalf("connect to schema database: %v", err)
	}
	t.Cleanup(func() {
		if err := conn.Close(ctx); err != nil {
			t.Errorf("close schema database: %v", err)
		}
	})

	for _, spec := range persistedModels {
		t.Run(spec.name, func(t *testing.T) {
			modelType := reflect.TypeOf(spec.model)
			modelColumns := taggedColumns(t, modelType)
			databaseColumns := databaseColumns(t, ctx, conn, spec.model.Table())
			databaseColumnSet := make(map[string]struct{}, len(databaseColumns))
			for name := range databaseColumns {
				databaseColumnSet[name] = struct{}{}
			}

			if missing := setDifference(databaseColumnSet, modelColumns); len(missing) > 0 {
				t.Errorf("model is missing database columns: %s", strings.Join(missing, ", "))
			}
			if extra := setDifference(modelColumns, databaseColumnSet); len(extra) > 0 {
				t.Errorf("model has columns absent from database: %s", strings.Join(extra, ", "))
			}

			validateColumnTypes(t, modelType, databaseColumns)
		})
	}
}

func TestPersistedEnumsMatchDatabase(t *testing.T) {
	if os.Getenv("I_MODELS_SCHEMA_TEST") != "1" {
		t.Skip("set I_MODELS_SCHEMA_TEST=1 to compare enum constants with PostgreSQL")
	}

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, schemaDatabaseURL())
	if err != nil {
		t.Fatalf("connect to schema database: %v", err)
	}
	t.Cleanup(func() {
		if err := conn.Close(ctx); err != nil {
			t.Errorf("close schema database: %v", err)
		}
	})

	for enumName, modelValues := range persistedEnums {
		t.Run(enumName, func(t *testing.T) {
			rows, err := conn.Query(ctx, `
				SELECT e.enumlabel
				FROM pg_type t
				JOIN pg_enum e ON e.enumtypid = t.oid
				WHERE t.typname = $1
				ORDER BY e.enumsortorder
			`, enumName)
			if err != nil {
				t.Fatalf("query enum %s: %v", enumName, err)
			}
			databaseValues, err := pgx.CollectRows(rows, pgx.RowTo[string])
			if err != nil {
				t.Fatalf("collect enum %s: %v", enumName, err)
			}

			modelSet := stringSet(modelValues)
			databaseSet := stringSet(databaseValues)
			if missing := setDifference(databaseSet, modelSet); len(missing) > 0 {
				t.Errorf("model enum is missing values: %s", strings.Join(missing, ", "))
			}
			if extra := setDifference(modelSet, databaseSet); len(extra) > 0 {
				t.Errorf("model enum has values absent from database: %s", strings.Join(extra, ", "))
			}
		})
	}
}

func taggedColumns(t *testing.T, typ reflect.Type) map[string]struct{} {
	t.Helper()

	columns := make(map[string]struct{}, typ.NumField())
	for i := range typ.NumField() {
		field := typ.Field(i)
		tag, ok := field.Tag.Lookup("db")
		if !ok {
			t.Errorf("field %s has no explicit db tag", field.Name)
			continue
		}

		column, _, _ := strings.Cut(tag, ",")
		if column == "-" {
			continue
		}
		if column == "" {
			t.Errorf("field %s has an empty db tag", field.Name)
			continue
		}
		if _, ok := columns[column]; ok {
			t.Errorf("db column %q is tagged more than once", column)
		}
		columns[column] = struct{}{}
	}

	return columns
}

func databaseColumns(t *testing.T, ctx context.Context, conn *pgx.Conn, table string) map[string]databaseColumn {
	t.Helper()

	rows, err := conn.Query(ctx, `
		SELECT column_name, data_type, udt_name, is_nullable
		FROM information_schema.columns
		WHERE table_schema = 'public' AND table_name = $1
	`, table)
	if err != nil {
		t.Fatalf("query columns for %s: %v", table, err)
	}

	columns, err := pgx.CollectRows(rows, pgx.RowToStructByName[databaseColumn])
	if err != nil {
		t.Fatalf("collect columns for %s: %v", table, err)
	}
	if len(columns) == 0 {
		t.Fatalf("table %s does not exist or has no columns", table)
	}

	result := make(map[string]databaseColumn, len(columns))
	for _, column := range columns {
		result[column.Name] = column
	}
	return result
}

func validateColumnTypes(t *testing.T, typ reflect.Type, databaseColumns map[string]databaseColumn) {
	t.Helper()

	for i := range typ.NumField() {
		field := typ.Field(i)
		columnName, _, _ := strings.Cut(field.Tag.Get("db"), ",")
		column, ok := databaseColumns[columnName]
		if !ok {
			continue
		}

		isPointerOrNullField := field.Type.Kind() == reflect.Pointer || strings.HasPrefix(field.Type.Name(), "Null")
		isStatusAwareField := field.Type.PkgPath() == "github.com/global-torque/go-common/orm/v2/pgtype" ||
			field.Type.Kind() == reflect.Interface
		if column.IsNullable == "NO" && isPointerOrNullField {
			t.Errorf("column %s is not nullable but field %s has nullable type %s", columnName, field.Name, field.Type)
		}
		if column.IsNullable == "YES" && !isPointerOrNullField && !isStatusAwareField {
			t.Errorf("column %s is nullable but field %s has non-nullable type %s", columnName, field.Name, field.Type)
		}
		if !isCompatibleGoType(field.Type, column) {
			t.Errorf(
				"column %s has PostgreSQL type %s/%s but field %s has type %s",
				columnName,
				column.DataType,
				column.UDTName,
				field.Name,
				field.Type,
			)
		}
	}
}

func isCompatibleGoType(typ reflect.Type, column databaseColumn) bool {
	for typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
	if typ.Kind() == reflect.Interface {
		return true
	}
	if strings.HasPrefix(typ.Name(), "Null") {
		return true
	}
	if typ.PkgPath() == "github.com/global-torque/go-common/orm/v2/pgtype" {
		return column.UDTName == "timestamptz" || column.UDTName == "varchar" || column.UDTName == "text"
	}

	switch column.UDTName {
	case "int2", "int4", "int8":
		return typ.Kind() >= reflect.Int && typ.Kind() <= reflect.Int64
	case "numeric":
		return typ.Kind() == reflect.Float32 || typ.Kind() == reflect.Float64 || typ.Kind() == reflect.String
	case "bool":
		return typ.Kind() == reflect.Bool
	case "json", "jsonb":
		return typ.Kind() == reflect.Map || typ.Kind() == reflect.Slice || typ.Kind() == reflect.Struct
	case "uuid":
		return typ.PkgPath() == "github.com/google/uuid" && typ.Name() == "UUID"
	case "timestamp", "timestamptz":
		return typ.PkgPath() == "time" && typ.Name() == "Time"
	default:
		return typ.Kind() == reflect.String
	}
}

func setDifference(left, right map[string]struct{}) []string {
	difference := make([]string, 0)
	for item := range left {
		if _, ok := right[item]; !ok {
			difference = append(difference, item)
		}
	}
	sort.Strings(difference)
	return difference
}

func enumStrings[T ~string](values []T) []string {
	result := make([]string, len(values))
	for i, value := range values {
		result[i] = string(value)
	}
	return result
}

func stringSet(values []string) map[string]struct{} {
	result := make(map[string]struct{}, len(values))
	for _, value := range values {
		result[value] = struct{}{}
	}
	return result
}

func defaultColumn(field string) string {
	parts := strings.Fields(field)
	if len(parts) >= 3 && strings.EqualFold(parts[len(parts)-2], "AS") {
		return parts[len(parts)-1]
	}
	return field
}

func schemaDatabaseURL() string {
	host := envOrDefault("DB_HOST", "127.0.0.1")
	port := envOrDefault("DB_PORT", "5432")
	user := envOrDefault("DB_USER", "postgres")
	password := envOrDefault("DB_PASSWORD", "postgres")
	database := envOrDefault("DB_DATABASE", "invest_platform")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, database)
}

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
