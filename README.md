# i-models

`i-models` is the shared Go schema-contract module used by Torque services. It
defines PostgreSQL-backed model structs, table and content-type names, enum
values, JSON projections, and a small generic CRUD facade. It does not own
database migrations and it is not a standalone service.

The module path is `github.com/webdevelop-pro/i-models`. The repository is
currently built with Go 1.25.8. The `v0.1.0` release contains the canonical
stablecoin-redemption status, approval/preflight, grant, operation-binding, and
quarantine contracts used by the coordinated service rollout.

## Where it sits

```text
PostgreSQL migrations
        │ define tables, columns and enum labels
        ▼
     i-models
        │ shared structs, constants and persistence helpers
        ├── evm-api: EVM projections, contract/offer constants, audit logs
        ├── wallet-api: wallets, transactions and funding sources
        ├── payment-api: investment and crypto-settlement projections
        ├── escrow-api: offer, investment-profile and user projections
        ├── kyc-api: profile/KYC mutation and audit records
        ├── esign-api: investment, offer, filer and notification contracts
        ├── email-worker: notification/email data and event qualification
        └── filer-api: Pub/Sub processing leases
```

Services are free to define narrower projections by embedding an `i-models`
type or by implementing the same `Fields`, `Table`, and `SetID` contract. Many
newer services query those projections with `go-common/orm`; older paths use
the compatibility helpers in [`models`](models/README.md).

## Basic use

```go
import (
    "github.com/webdevelop-pro/i-models/models"
    "github.com/webdevelop-pro/i-models/profiles"
)

profile, err := models.RetrieveOne[profiles.Profile](
    ctx,
    repo,
    map[string]any{"id": profileID},
)
```

Every type accepted by the generic helpers has a pointer form that implements:

```go
SetID(any)
Fields() []string
Table() string
```

Types with `SetDB` also receive the repository after `RetrieveOne`,
`RetrieveAll`, or `Create`. The supported operations are:

- `RetrieveOne` and `RetrieveAll` with equality filters expressed as
  `map[string]any`;
- `Create` with a column/value map;
- `Update`, `Exists`, and `Delete` with equality filters;
- deprecated `RetriveOne` and `RetriveAll` spellings for source compatibility.

For joins, ordering, row locks, compare-and-swap updates, or transaction-bound
work, use `go-common/orm/v2`, Squirrel, or repository SQL directly. The
`investments` and `pubsubactivities` packages expose dedicated helpers where a
shared concurrency protocol is part of the domain contract.

## Model conventions

- Database columns are identified by `db` tags. JSON/YAML tags describe API or
  event serialization and may intentionally hide fields.
- `Fields()` is the default SELECT projection. It is not guaranteed to contain
  every tagged field in the struct.
- `Table()` is the canonical table name used by the generic ORM.
- `ToJSON`/`ToMap` are explicit projections; inspect them before using their
  output at an external boundary.
- Enum types mirror PostgreSQL enum labels. Call `IsValid` where available at
  untrusted input boundaries; an enum's Go type alone does not validate a cast
  string.
- Nullable enums use generated-style wrappers such as `offers.NullOfferT`.
- [`pgtype.Timestamptz`](pgtype/README.md) is a compatibility wrapper for SQL,
  JSON, finite timestamps, and PostgreSQL infinity values.

## Package catalog

| Package | Responsibility | Verified direct consumers |
|---|---|---|
| [`distributions`](distributions/README.md) | Distribution, report, and filer-link rows | No sibling runtime import found |
| [`emails`](emails/README.md) | Outbound email records and provider interface | `email-worker` |
| [`evmallowancesequences`](evmallowancesequences/README.md) | ERC-20 allowance nonce/sequence rows | No direct import; `evm-api` uses the table via local repositories |
| [`evmchainaccounts`](evmchainaccounts/README.md) | Chain-specific EVM accounts | No direct import; `evm-api` has local projections for this table |
| [`evmcontracts`](evmcontracts/README.md) | Deployed contract records and shared table names | `evm-api` |
| [`evmtransfers`](evmtransfers/README.md) | Retired legacy EVM transfer rows | No current consumer |
| [`evmwalletbalances`](evmwalletbalances/README.md) | Materialized wallet asset balances | `evm-api` |
| [`evmwalletoperationeffects`](evmwalletoperationeffects/README.md) | Ledger effects produced by wallet operations | No direct import; `evm-api`/`payment-api` query the table |
| [`evmwalletoperations`](evmwalletoperations/README.md) | EVM operation lifecycle | `evm-api`, `payment-api`, `email-worker` |
| [`evmwallets`](evmwallets/README.md) | Embedded/on-chain wallet records | `evm-api` |
| [`filers`](filers/README.md) | Hierarchical filer metadata | `esign-api` |
| [`fundingsources`](fundingsources/README.md) | Linked bank/funding sources | `wallet-api`, `email-worker` |
| [`fundnavrecords`](fundnavrecords/README.md) | Fund NAV history | No direct import; `evm-api` queries the table |
| [`historylogs`](historylogs/README.md) | Django-compatible admin history | `kyc-api`; also used internally by wallet/transaction hooks |
| [`investments`](investments/README.md) | Investments, redemptions, and investor profiles | `evm-api`, `wallet-api`, `payment-api`, `escrow-api`, `kyc-api`, `esign-api`, `email-worker` |
| [`logs`](logs/README.md) | HTTP/provider audit logs | Most API services and `email-worker` |
| [`models`](models/README.md) | Generic CRUD compatibility facade | `wallet-api`, `email-worker`; model contract used broadly |
| [`notifications`](notifications/README.md) | User notification rows | `evm-api`, `esign-api`, `email-worker` |
| [`offers`](offers/README.md) | Offers, filer links, comments, and enums | `evm-api`, `wallet-api`, `payment-api`, `escrow-api`, `esign-api`, `email-worker` |
| [`pgtype`](pgtype/README.md) | Timestamp compatibility types | `wallet-api` directly; model fields throughout this module |
| [`profiles`](profiles/README.md) | Account profile/KYC data and JSONB mutation | `evm-api`, `wallet-api`, `payment-api`, `escrow-api`, `kyc-api`, `email-worker` |
| [`pubsubactivities`](pubsubactivities/README.md) | Pub/Sub deduplication and domain-reaction leases | All event-consuming sibling services |
| [`pubsublogs`](pubsublogs/README.md) | Legacy Pub/Sub log rows | No sibling runtime import found |
| [`schemaconformance`](schemaconformance/README.md) | Canonical projection and optional live-schema tests | Repository tests only |
| [`transactions`](transactions/README.md) | Legacy wallet transaction lifecycle | `wallet-api`, `email-worker` |
| [`userinvitations`](userinvitations/README.md) | Invitation roles, kinds, and states | `email-worker` |
| [`users`](users/README.md) | User identity rows and JSONB data | `evm-api`, `wallet-api`, `escrow-api`, `esign-api`, `kyc-api`, `email-worker` |
| [`wallets`](wallets/README.md) | Legacy custodial wallet balances | `wallet-api`, `email-worker` |

“Direct consumer” means an import was found in the checked sibling Go module;
it does not imply ownership of the corresponding table. A service may still
access a table through a local projection or raw SQL.

## Transactions and concurrency

Pass a `pgx.Tx` (or another compatible `db.Repository`) to helpers when several
writes must commit atomically. Do not read with the pool and write with the
transaction in the same unit of work.

Investment funding/redemption transitions use explicit `FOR UPDATE` helpers in
[`investments`](investments/README.md). Pub/Sub consumers should use the fenced
lease API in [`pubsubactivities`](pubsubactivities/README.md); a boolean dedupe
check alone cannot prevent stale workers from overwriting newer outcomes.

## Important compatibility notes

- The `wallets` and `transactions` packages retain stateful `Save` APIs for
  legacy consumers. New code should prefer explicit `models.Update` or
  transaction-scoped ORM calls.
- `evmwallets`, `evmcontracts`, and the retired `evmtransfers` types also expose
  `Save`, but currently have no public setters that populate their internal
  update map. Treat those methods as compatibility-only.
- `wallets.SetOutBalanceFN` and `wallets.SetBalanceFN` currently mutate
  `IncBalance`; callers should not use them until that legacy behavior is fixed.
- The documentation describes the checked source, not an assumed generated ORM
  API. In particular, `i-models/models` accepts equality maps only.

## Validation

Run the normal suite from this repository:

```sh
./make.sh test
go vet ./...
```

The default suite checks enum contracts, projections, persistence behavior, and
the in-memory parts of Pub/Sub coordination. Live PostgreSQL schema comparison
is opt-in because it requires a database configured with the application
schema:

```sh
I_MODELS_SCHEMA_TEST=1 go test ./schemaconformance
```

When changing a table, column, or PostgreSQL enum, coordinate the migration,
the relevant model and its `Fields()` projection, schema-conformance registry,
and every consuming service in the same release plan.

## Non-runtime tooling

[`experiments`](experiments/Readme.md) is a separate Go module containing
SQLBoiler template experiments. It is not imported by the runtime packages and
is not part of `go test ./...` from the root module.
