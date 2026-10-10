package users

import (
	"context"
	"errors"
	"testing"

	"github.com/global-torque/go-common/logger/v2"
	"github.com/global-torque/go-common/orm/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type mergeRepository struct {
	commandTag pgconn.CommandTag
	err        error
	sql        string
	args       []any
}

func (r *mergeRepository) Query(context.Context, string, ...any) (pgx.Rows, error) {
	return nil, errors.New("unexpected Query")
}

func (r *mergeRepository) QueryRow(context.Context, string, ...any) pgx.Row {
	return nil
}

func (r *mergeRepository) Exec(_ context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	r.sql = sql
	r.args = args
	return r.commandTag, r.err
}

func (r *mergeRepository) Lg() logger.Logger { return logger.Logger{} }

func TestMergeDataByID(t *testing.T) {
	tests := []struct {
		name       string
		commandTag pgconn.CommandTag
		execErr    error
		wantErr    error
	}{
		{name: "merged", commandTag: pgconn.NewCommandTag("UPDATE 1")},
		{name: "missing", commandTag: pgconn.NewCommandTag("UPDATE 0"), wantErr: orm.ErrNoRowsAffected},
		{name: "database error", execErr: errors.New("database failed"), wantErr: errors.New("database failed")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mergeRepository{commandTag: tt.commandTag, err: tt.execErr}
			err := MergeDataByID(context.Background(), repo, 17, map[string]any{"issuer": "nc-1"})
			if tt.wantErr == nil && err != nil {
				t.Fatalf("MergeDataByID returned error: %v", err)
			}
			if tt.wantErr != nil && !errors.Is(err, tt.wantErr) && (tt.execErr == nil || !errors.Is(err, tt.execErr)) {
				t.Fatalf("expected %v, got %v", tt.wantErr, err)
			}
		})
	}
}
