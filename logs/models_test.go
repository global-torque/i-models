package logs

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/webdevelop-pro/go-common/logger"
)

type rowFunc func(dest ...any) error

func (fn rowFunc) Scan(dest ...any) error {
	return fn(dest...)
}

type logRepository struct {
	queryRowCalls int
	row           pgx.Row
}

func (*logRepository) Query(context.Context, string, ...any) (pgx.Rows, error) {
	return nil, errors.New("unexpected Query")
}

func (*logRepository) Exec(context.Context, string, ...any) (pgconn.CommandTag, error) {
	return pgconn.CommandTag{}, errors.New("unexpected Exec")
}

func (repo *logRepository) QueryRow(context.Context, string, ...any) pgx.Row {
	repo.queryRowCalls++
	return repo.row
}

func TestLogRequestHandlesCanceledContext(t *testing.T) {
	tests := []struct {
		name           string
		cancelBefore   bool
		row            pgx.Row
		wantQueryCalls int
	}{
		{
			name:           "already canceled request skips persistence",
			cancelBefore:   true,
			wantQueryCalls: 0,
		},
		{
			name: "content type lookup canceled after request starts",
			row: rowFunc(func(...any) error {
				return context.Canceled
			}),
			wantQueryCalls: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			if test.cancelBefore {
				cancel()
			} else {
				defer cancel()
			}

			repo := &logRepository{row: test.row}
			model := &LogLog{}
			log := logger.NewLogger(ctx, "logs-test", "disabled", io.Discard)
			req := httptest.NewRequest(http.MethodPost, "https://example.test/rpc", nil).WithContext(ctx)

			model.LogRequest(log, repo, "", ServicesTAlchemy, "profiles", ObjectType("profile"))(req)

			if repo.queryRowCalls != test.wantQueryCalls {
				t.Fatalf("QueryRow calls = %d, want %d", repo.queryRowCalls, test.wantQueryCalls)
			}
			if model.ID != 0 {
				t.Fatalf("log ID = %d, want zero after cancellation", model.ID)
			}
		})
	}
}
