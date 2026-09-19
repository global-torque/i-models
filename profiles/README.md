# `profiles`

This package models account/investor profiles stored in
`investment_profiles`. `Profile` is the shared identity/KYC projection;
`ProfileData`, `KYC`, and `Accreditation` describe structured JSON data and
history. Profile-type, KYC, and accreditation constants are used throughout
API validation and event handling.

The package also owns three atomic JSONB/domain mutations:

- `AppendKYCByID` updates KYC state, prepends an audit item, and optionally
  merges profile data;
- `MergeDataByID` merges an object into the `data` JSONB column and requires
  exactly one affected row;
- `UpdateRelatedKYCByProfile` propagates KYC state to related IRA or child
  profiles and returns the affected IDs.

`kyc-api` is the primary mutation consumer. `evm-api`, `wallet-api`,
`payment-api`, `escrow-api`, and `email-worker` use the model/constants for
wallet eligibility, settlement, provider integration, and notifications.

The similarly named `investments.InvestmentProfile` is a separate service
projection of the same domain/table boundary. Choose the projection whose
`Fields()` match the query being performed.

See [holder flows](holder-flows.md) for the cross-repository KYC-pass event
and token-holder boundary.

See the [root guide](../README.md) for projection and transaction conventions.
