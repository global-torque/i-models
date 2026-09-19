package investments

import (
	"context"
	"errors"
	"strings"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var errQueryStopped = errors.New("query stopped after SQL capture")

type persistenceRepository struct {
	querySQL  string
	queryArgs []any
	execSQL   string
	execArgs  []any
	command   pgconn.CommandTag
}

func (repo *persistenceRepository) Query(
	_ context.Context,
	sql string,
	args ...any,
) (pgx.Rows, error) {
	repo.querySQL = sql
	repo.queryArgs = args
	return nil, errQueryStopped
}

func (repo *persistenceRepository) QueryRow(
	context.Context,
	string,
	...any,
) pgx.Row {
	panic("unexpected QueryRow")
}

func (repo *persistenceRepository) Exec(
	_ context.Context,
	sql string,
	args ...any,
) (pgconn.CommandTag, error) {
	repo.execSQL = sql
	repo.execArgs = args
	return repo.command, nil
}

func TestRetrieveBusinessRowsForUpdate(t *testing.T) {
	tests := []struct {
		name  string
		call  func(context.Context, *persistenceRepository) error
		table string
	}{
		{
			name: "investment",
			call: func(ctx context.Context, repo *persistenceRepository) error {
				_, err := RetrieveInvestmentForUpdate(ctx, repo, 17)
				return err
			},
			table: "investment_investments",
		},
		{
			name: "redemption",
			call: func(ctx context.Context, repo *persistenceRepository) error {
				_, err := RetrieveRedemptionForUpdate(ctx, repo, 23)
				return err
			},
			table: "investment_redemptions",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := &persistenceRepository{}
			err := test.call(context.Background(), repo)
			if !errors.Is(err, errQueryStopped) {
				t.Fatalf("expected captured query error, got %v", err)
			}
			if !strings.Contains(repo.querySQL, "FROM "+test.table) ||
				!strings.HasSuffix(repo.querySQL, "FOR UPDATE") {
				t.Fatalf("query does not hold the shared business-row lock: %s", repo.querySQL)
			}
			if len(repo.queryArgs) != 1 {
				t.Fatalf("expected one business ID argument, got %#v", repo.queryArgs)
			}
		})
	}
}

func TestConditionalBusinessUpdatesRequireStatusAndUnlockedPath(t *testing.T) {
	tests := []struct {
		name       string
		command    string
		call       func(context.Context, *persistenceRepository) (bool, error)
		table      string
		lockColumn string
	}{
		{
			name:    "investment changed",
			command: "UPDATE 1",
			call: func(ctx context.Context, repo *persistenceRepository) (bool, error) {
				return UpdateUnlockedInvestment(
					ctx,
					repo,
					17,
					InvestmentT("legally_confirmed"),
					map[string]any{"vault_request_locked_at": sq.Expr("clock_timestamp()")},
					sq.Expr("custody_funded_at IS NOT NULL"),
				)
			},
			table:      "investment_investments",
			lockColumn: "vault_request_locked_at IS NULL",
		},
		{
			name:    "redemption conflict",
			command: "UPDATE 0",
			call: func(ctx context.Context, repo *persistenceRepository) (bool, error) {
				return UpdateUnlockedRedemption(
					ctx,
					repo,
					23,
					RedemptionStatusPending,
					map[string]any{"updated_at": sq.Expr("clock_timestamp()")},
					sq.Expr("cancelled_at IS NULL"),
				)
			},
			table:      "investment_redemptions",
			lockColumn: "request_locked_at IS NULL",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := &persistenceRepository{command: pgconn.NewCommandTag(test.command)}
			changed, err := test.call(context.Background(), repo)
			if err != nil {
				t.Fatalf("conditional update returned error: %v", err)
			}
			wantChanged := test.command == "UPDATE 1"
			if changed != wantChanged {
				t.Fatalf("changed=%v, want %v", changed, wantChanged)
			}
			for _, fragment := range []string{
				"UPDATE " + test.table,
				"id =",
				"status =",
				test.lockColumn,
			} {
				if !strings.Contains(repo.execSQL, fragment) {
					t.Fatalf("conditional SQL is missing %q: %s", fragment, repo.execSQL)
				}
			}
			if strings.Contains(repo.execSQL, "business_transition_version") {
				t.Fatalf("conditional SQL uses forbidden business version: %s", repo.execSQL)
			}
		})
	}
}
