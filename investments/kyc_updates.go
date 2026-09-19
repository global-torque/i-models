package investments

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/webdevelop-pro/i-models/profiles"
)

// InvestmentQueryExecutor binds the update to the caller's transaction.
// The caller owns commit/rollback and appends any business facts before commit.
type InvestmentQueryExecutor interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
}

// MarkLegallyConfirmedForProfile moves confirmed investments to
// legally_confirmed when KYC/accreditation requirements are satisfied.
func MarkLegallyConfirmedForProfile(ctx context.Context, repo InvestmentQueryExecutor, profile *profiles.Profile) ([]InvestmentInvestment, error) {
	if profile == nil {
		return nil, fmt.Errorf("profile is nil")
	}

	sql, args := markLegallyConfirmedQuery(profile)

	rows, err := repo.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("update investments for profile %d: %w", profile.ID, err)
	}
	defer rows.Close()

	investments := []InvestmentInvestment{}
	for rows.Next() {
		var investment InvestmentInvestment
		if err := rows.Scan(&investment.ID, &investment.Status); err != nil {
			return investments, fmt.Errorf("scan investment for profile %d: %w", profile.ID, err)
		}
		investments = append(investments, investment)
	}
	if err := rows.Err(); err != nil {
		return investments, fmt.Errorf("iterate investments for profile %d: %w", profile.ID, err)
	}

	return investments, nil
}

func markLegallyConfirmedQuery(profile *profiles.Profile) (string, []any) {
	if profile.Type == profiles.Individual {
		return `
			UPDATE investment_investments as ii SET status=$1
			FROM investment_profiles AS ip
			WHERE ii.profile_id=ip.id
			AND ip.id=$2
			AND ii.user_id=$3 and ii.status=$4
			AND ip.type in ('individual', 'sdira', 'solo401k')
			RETURNING ii.id,ii.status`, []any{
				InvestmentTLegallyConfirmed, profile.ID, profile.UserID, InvestmentTConfirmed,
			}
	}

	return `
		UPDATE investment_investments SET status=$1 WHERE profile_id=$2 and status=$3
		RETURNING id,status`, []any{
			InvestmentTLegallyConfirmed, profile.ID, InvestmentTConfirmed,
		}
}
