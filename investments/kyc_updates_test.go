package investments

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/webdevelop-pro/i-models/profiles"
)

func TestMarkLegallyConfirmedQueryIsolatesIndividualProfile(t *testing.T) {
	t.Parallel()

	profile := &profiles.Profile{ID: 42, UserID: 7, Type: profiles.Individual}
	query, args := markLegallyConfirmedQuery(profile)

	for _, predicate := range []string{
		"ii.profile_id=ip.id",
		"ip.id=$2",
		"ii.user_id=$3",
		"ii.status=$4",
	} {
		if !strings.Contains(query, predicate) {
			t.Fatalf("query does not contain multi-profile isolation predicate %q: %s", predicate, query)
		}
	}

	wantArgs := []any{InvestmentTLegallyConfirmed, 42, 7, InvestmentTConfirmed}
	if !reflect.DeepEqual(args, wantArgs) {
		t.Fatalf("unexpected query arguments: got %#v, want %#v", args, wantArgs)
	}
}

func TestMarkLegallyConfirmedQueryUsesExactNonIndividualProfile(t *testing.T) {
	t.Parallel()

	profile := &profiles.Profile{ID: 51, UserID: 9, Type: profiles.Entity}
	query, args := markLegallyConfirmedQuery(profile)

	if !strings.Contains(query, "profile_id=$2") {
		t.Fatalf("query does not isolate the non-individual profile: %s", query)
	}
	wantArgs := []any{InvestmentTLegallyConfirmed, 51, InvestmentTConfirmed}
	if !reflect.DeepEqual(args, wantArgs) {
		t.Fatalf("unexpected query arguments: got %#v, want %#v", args, wantArgs)
	}
}

// The returned changed rows and their caller-owned transaction are the contract
// needed by both KYC and Escrow when atomically appending their own facts.
func TestMarkLegallyConfirmedUsesCallerTransaction(t *testing.T) {
	dsn := os.Getenv("BUSINESS_EVENTS_TEST_DSN")
	if dsn == "" {
		t.Skip("set BUSINESS_EVENTS_TEST_DSN to a disposable PostgreSQL database")
	}
	must := func(err error) {
		t.Helper()
		if err != nil {
			t.Fatal(err)
		}
	}
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dsn)
	must(err)
	t.Cleanup(func() { must(conn.Close(ctx)) })
	tx, err := conn.Begin(ctx)
	must(err)
	t.Cleanup(func() {
		err := tx.Rollback(ctx)
		if !errors.Is(err, pgx.ErrTxClosed) {
			must(err)
		}
	})
	// Temporary minimal query fixtures avoid changing real source tables or their business triggers.
	_, err = tx.Exec(ctx, `CREATE TEMP TABLE investment_profiles(id integer,user_id integer,type text);
 CREATE TEMP TABLE investment_investments(id integer,profile_id integer,user_id integer,status text);
 INSERT INTO investment_profiles VALUES (42,7,'individual'),(43,7,'individual');
 INSERT INTO investment_investments VALUES (100,42,7,'confirmed'),(101,43,7,'confirmed'),(102,42,8,'confirmed');`)
	must(err)
	changed, err := MarkLegallyConfirmedForProfile(ctx, tx, &profiles.Profile{ID: 42, UserID: 7, Type: profiles.Individual})
	must(err)
	if len(changed) != 1 {
		t.Fatalf("changed rows: %v", changed)
	}
	if changed[0].ID != 100 {
		t.Fatalf("changed ID: %d", changed[0].ID)
	}
	if changed[0].Status != InvestmentTLegallyConfirmed {
		t.Fatalf("status: %s", changed[0].Status)
	}
	replay, err := MarkLegallyConfirmedForProfile(ctx, tx, &profiles.Profile{ID: 42, UserID: 7, Type: profiles.Individual})
	must(err)
	if len(replay) != 0 {
		t.Fatalf("unexpected replay: %v", replay)
	}
	must(tx.Rollback(ctx))
	var fixtureExists bool
	must(conn.QueryRow(ctx, "SELECT to_regclass('pg_temp.investment_investments') IS NOT NULL").Scan(&fixtureExists))
	if fixtureExists {
		t.Fatal("rolled-back fixture survived")
	}
}
