package logs

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/global-torque/go-common/context/v2/keys"
	"github.com/global-torque/go-common/logger/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type rowFunc func(dest ...any) error

func (fn rowFunc) Scan(dest ...any) error {
	return fn(dest...)
}

type logRepository struct {
	queryRowCalls int
	queryRowArgs  [][]any
	row           pgx.Row
	rows          []pgx.Row
	execArgs      []any
	execTag       pgconn.CommandTag
	execErr       error
}

func (*logRepository) Query(context.Context, string, ...any) (pgx.Rows, error) {
	return nil, errors.New("unexpected Query")
}

func (repo *logRepository) Exec(_ context.Context, _ string, args ...any) (pgconn.CommandTag, error) {
	repo.execArgs = args
	return repo.execTag, repo.execErr
}

func (repo *logRepository) QueryRow(_ context.Context, _ string, args ...any) pgx.Row {
	repo.queryRowCalls++
	repo.queryRowArgs = append(repo.queryRowArgs, args)
	if len(repo.rows) > 0 {
		row := repo.rows[0]
		repo.rows = repo.rows[1:]
		return row
	}
	return repo.row
}

func rowWithID(id int) pgx.Row {
	return rowFunc(func(dest ...any) error {
		if len(dest) != 1 {
			return errors.New("expected one scan destination")
		}
		value, ok := dest[0].(*int)
		if !ok {
			return errors.New("expected an integer scan destination")
		}
		*value = id
		return nil
	})
}

func requestContext(requestID, messageID string) context.Context {
	return keys.SetCtxValues(context.Background(), map[keys.ContextKey]any{
		keys.RequestID: requestID,
		keys.MSGID:     messageID,
	})
}

func testLogRequest(repo *logRepository, ctx context.Context, service ServicesT) (*http.Request, *LogLog) {
	request := httptest.NewRequest(http.MethodPost, "https://example.test/rpc", strings.NewReader(`{"request":true}`)).WithContext(ctx)
	model := &LogLog{}
	log := logger.NewLogger(ctx, "logs-test", "disabled", io.Discard)
	model.LogRequest(log, repo, "42", service, "profiles", ObjectProfile)(request)
	return request, model
}

func TestLogHttpRequestPersistsV2RequestAndMessageIDs(t *testing.T) {
	ctx := requestContext("request-123", "message-456")
	repo := &logRepository{rows: []pgx.Row{rowWithID(11), rowWithID(27)}}
	request := httptest.NewRequest(http.MethodPost, "https://example.test/rpc", nil).WithContext(ctx)

	model, err := LogHttpRequest(
		ctx,
		repo,
		LogTypeTIncoming,
		"42",
		ServicesTAlchemy,
		"profiles",
		ObjectProfile,
		request,
		[]byte(`{"request":true}`),
	)
	if err != nil {
		t.Fatalf("LogHttpRequest() error = %v", err)
	}
	if model.ID != 27 {
		t.Fatalf("log ID = %d, want 27", model.ID)
	}
	if len(repo.queryRowArgs) != 2 {
		t.Fatalf("QueryRow calls = %d, want 2", len(repo.queryRowArgs))
	}

	args := repo.queryRowArgs[1]
	if got := args[6]; got != "request-123" {
		t.Fatalf("request_id argument = %v, want request-123", got)
	}
	if got := args[9]; got != "message-456" {
		t.Fatalf("msg_id argument = %v, want message-456", got)
	}
}

func TestLogRequestUsesV2RequestIDForHeaders(t *testing.T) {
	tests := []struct {
		name                   string
		requestID              string
		messageID              string
		wantRequestID          string
		wantIdempotency        string
		wantPersistedRequestID string
		wantPersistedMessageID string
	}{
		{
			name:                   "request ID",
			requestID:              "request-123",
			messageID:              "message-456",
			wantRequestID:          "request-123",
			wantIdempotency:        "request-123",
			wantPersistedRequestID: "request-123",
			wantPersistedMessageID: "message-456",
		},
		{
			name:                   "missing request ID fallback",
			wantRequestID:          "00000000-0000-0000-0000-000000000042",
			wantIdempotency:        "00000000-0000-0000-0000-000000000042",
			wantPersistedRequestID: "",
			wantPersistedMessageID: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := requestContext(test.requestID, test.messageID)
			repo := &logRepository{rows: []pgx.Row{rowWithID(11), rowWithID(27)}}
			request, model := testLogRequest(repo, ctx, ServicesTDwolla)

			if model.ID != 27 {
				t.Fatalf("log ID = %d, want 27", model.ID)
			}
			if got := request.Header.Get("X-Request-Id"); got != test.wantRequestID {
				t.Fatalf("X-Request-Id = %q, want %q", got, test.wantRequestID)
			}
			if got := request.Header.Get("Idempotency-Key"); got != test.wantIdempotency {
				t.Fatalf("Idempotency-Key = %q, want %q", got, test.wantIdempotency)
			}
			args := repo.queryRowArgs[1]
			if got := args[6]; got != test.wantPersistedRequestID {
				t.Fatalf("persisted request_id = %v, want %q", got, test.wantPersistedRequestID)
			}
			if got := args[9]; got != test.wantPersistedMessageID {
				t.Fatalf("persisted msg_id = %v, want %q", got, test.wantPersistedMessageID)
			}
		})
	}
}

func TestLogResponseUpdatesCreatedLog(t *testing.T) {
	ctx := requestContext("request-123", "message-456")
	repo := &logRepository{
		rows:    []pgx.Row{rowWithID(11), rowWithID(27)},
		execTag: pgconn.NewCommandTag("UPDATE 1"),
	}
	request, model := testLogRequest(repo, ctx, ServicesTAlchemy)

	response := &http.Response{
		StatusCode: 201,
		Header: http.Header{
			"X-Response-Id": {"response-789"},
		},
		Body:    io.NopCloser(strings.NewReader(`{"created":true}`)),
		Request: request,
	}
	model.LogResponse(logger.NewLogger(ctx, "logs-test", "disabled", io.Discard), repo, ServicesTAlchemy)(response)

	if len(repo.execArgs) != 4 {
		t.Fatalf("Exec arguments = %d, want 4", len(repo.execArgs))
	}
	if got := repo.execArgs[0]; got != model.ID {
		t.Fatalf("updated log ID = %v, want %d", got, model.ID)
	}
	if got := repo.execArgs[1]; got != 201 {
		t.Fatalf("status code = %v, want 201", got)
	}
	if got := repo.execArgs[2].(http.Header).Get("X-Response-Id"); got != "response-789" {
		t.Fatalf("response header = %q, want response-789", got)
	}
	if got := string(repo.execArgs[3].([]byte)); got != `{"created":true}` {
		t.Fatalf("response body = %q, want {\"created\":true}", got)
	}
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
