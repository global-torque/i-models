package offers

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type OfferT string

// Enum values for OfferT
const (
	OfferTNew                  OfferT = "new"
	OfferTDraft                OfferT = "draft"
	OfferTLegalReview          OfferT = "legal_review"
	OfferTLegalDeclined        OfferT = "legal_declined"
	OfferTLegalAccepted        OfferT = "legal_accepted"
	OfferTPublished            OfferT = "published"
	OfferTLegalClosed          OfferT = "legal_closed"
	OfferTClosedSuccessfully   OfferT = "closed_successfully"
	OfferTClosedUnsuccessfully OfferT = "closed_unsuccessfully"

	OfferAppLabel  = "offer"
	OfferModelName = "offer"

	TableName = "offer_offers"
)

func AllOfferT() []OfferT {
	return []OfferT{
		OfferTNew,
		OfferTDraft,
		OfferTLegalReview,
		OfferTLegalDeclined,
		OfferTLegalAccepted,
		OfferTPublished,
		OfferTLegalClosed,
		OfferTClosedSuccessfully,
		OfferTClosedUnsuccessfully,
	}
}

func (e OfferT) IsValid() error {
	switch e {
	case OfferTNew, OfferTDraft, OfferTLegalReview, OfferTLegalDeclined, OfferTLegalAccepted, OfferTPublished, OfferTLegalClosed, OfferTClosedSuccessfully, OfferTClosedUnsuccessfully:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e OfferT) String() string {
	return string(e)
}

func (e *OfferT) Scan(value any) error {
	switch val := value.(type) {
	case nil:
		*e = ""
	case string:
		*e = OfferT(val)
	case []byte:
		*e = OfferT(val)
	default:
		return fmt.Errorf("cannot scan %T into OfferT", value)
	}

	return nil
}

// NullOfferT is a nullable OfferT enum type. It supports SQL and JSON serialization.
type NullOfferT struct {
	Val   OfferT
	Valid bool
}

// NullOfferTFrom creates a new OfferT that will never be blank.
func NullOfferTFrom(v OfferT) NullOfferT {
	return NewNullOfferT(v, true)
}

// NullOfferTFromPtr creates a new NullOfferT that be null if s is nil.
func NullOfferTFromPtr(v *OfferT) NullOfferT {
	if v == nil {
		return NewNullOfferT("", false)
	}
	return NewNullOfferT(*v, true)
}

// NewNullOfferT creates a new NullOfferT
func NewNullOfferT(v OfferT, valid bool) NullOfferT {
	return NullOfferT{
		Val:   v,
		Valid: valid,
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *NullOfferT) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		e.Val = ""
		e.Valid = false
		return nil
	}

	if err := json.Unmarshal(data, &e.Val); err != nil {
		return err
	}

	e.Valid = true
	return nil
}

// MarshalJSON implements json.Marshaler.
func (e NullOfferT) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Val)
}

// MarshalText implements encoding.TextMarshaler.
func (e NullOfferT) MarshalText() ([]byte, error) {
	if !e.Valid {
		return []byte{}, nil
	}
	return []byte(e.Val), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (e *NullOfferT) UnmarshalText(text []byte) error {
	if text == nil || len(text) == 0 {
		e.Valid = false
		return nil
	}

	e.Val = OfferT(text)
	e.Valid = true
	return nil
}

// SetValid changes this NullOfferT value and also sets it to be non-null.
func (e *NullOfferT) SetValid(v OfferT) {
	e.Val = v
	e.Valid = true
}

// Ptr returns a pointer to this NullOfferT value, or a nil pointer if this NullOfferT is null.
func (e NullOfferT) Ptr() *OfferT {
	if !e.Valid {
		return nil
	}
	return &e.Val
}

// IsZero returns true for null types.
func (e NullOfferT) IsZero() bool {
	return !e.Valid
}

// Value implements the driver Valuer interface.
func (e NullOfferT) Value() (any, error) {
	if !e.Valid {
		return nil, nil
	}
	return string(e.Val), nil
}

func (e *NullOfferT) Scan(value any) error {
	if value == nil {
		e.Val = ""
		e.Valid = false

		return nil
	}

	var offerType OfferT
	if err := offerType.Scan(value); err != nil {
		return err
	}

	e.Val = offerType
	e.Valid = true

	return nil
}

type OfferSecurityTypeT string

// Enum values for OfferSecurityTypeT
const (
	OfferSecurityTypeTEquity           OfferSecurityTypeT = "equity"
	OfferSecurityTypeTPreferredEquity  OfferSecurityTypeT = "preferred-equity"
	OfferSecurityTypeTDebt             OfferSecurityTypeT = "debt"
	OfferSecurityTypeTConvertibleDebt  OfferSecurityTypeT = "convertible-debt"
	OfferSecurityTypeTEquityWarrants   OfferSecurityTypeT = "equity-warrants"
	OfferSecurityTypeTConvertibleBonds OfferSecurityTypeT = "convertible-bonds"
	OfferSecurityTypeTPreferenceShares OfferSecurityTypeT = "preference-shares"
	OfferSecurityTypeTConvertibleNote  OfferSecurityTypeT = "convertible-note"
)

func AllOfferSecurityTypeT() []OfferSecurityTypeT {
	return []OfferSecurityTypeT{
		OfferSecurityTypeTEquity,
		OfferSecurityTypeTPreferredEquity,
		OfferSecurityTypeTDebt,
		OfferSecurityTypeTConvertibleDebt,
		OfferSecurityTypeTEquityWarrants,
		OfferSecurityTypeTConvertibleBonds,
		OfferSecurityTypeTPreferenceShares,
		OfferSecurityTypeTConvertibleNote,
	}
}

func (e OfferSecurityTypeT) IsValid() error {
	switch e {
	case OfferSecurityTypeTEquity, OfferSecurityTypeTPreferredEquity, OfferSecurityTypeTDebt, OfferSecurityTypeTConvertibleDebt, OfferSecurityTypeTEquityWarrants, OfferSecurityTypeTConvertibleBonds, OfferSecurityTypeTPreferenceShares, OfferSecurityTypeTConvertibleNote:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e OfferSecurityTypeT) String() string {
	return string(e)
} // OfferOffer is an object representing the database table.

func (e *OfferSecurityTypeT) Scan(value any) error {
	switch val := value.(type) {
	case nil:
		*e = ""
	case string:
		*e = OfferSecurityTypeT(val)
	case []byte:
		*e = OfferSecurityTypeT(val)
	default:
		return fmt.Errorf("cannot scan %T into OfferSecurityTypeT", value)
	}

	return nil
}

type OfferRegTypeT string
type TokenizationEngineT string
type TokenizationModelT string
type FundStructureT string

const (
	FundStructureOpenEnded   FundStructureT = "open_ended"
	FundStructureClosedEnded FundStructureT = "closed_ended"
)

func AllFundStructureT() []FundStructureT {
	return []FundStructureT{
		FundStructureOpenEnded,
		FundStructureClosedEnded,
	}
}

func (e FundStructureT) IsValid() error {
	switch e {
	case FundStructureOpenEnded, FundStructureClosedEnded:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e FundStructureT) String() string {
	return string(e)
}

const (
	TokenizationEngineNone    TokenizationEngineT = "none_not_tokenized"
	TokenizationEngineERC20   TokenizationEngineT = "ERC-20"
	TokenizationEngineERC721  TokenizationEngineT = "ERC-721"
	TokenizationEngineERC3643 TokenizationEngineT = "ERC-3643"
	TokenizationEngineERC7943 TokenizationEngineT = "ERC-7943"
)

const (
	TokenizationModelNone                           TokenizationModelT = "none_not_tokenized"
	TokenizationModelIssuerSponsoredOnchainRegister TokenizationModelT = "issuer_sponsored_onchain_register"
	TokenizationModelIssuerAuthorizedOffchain       TokenizationModelT = "issuer_authorized_offchain_register_transfer_instruction"
	TokenizationModelThirdPartyCustodialEntitlement TokenizationModelT = "third_party_custodial_entitlement"
	TokenizationModelThirdPartyLinkedSecurity       TokenizationModelT = "third_party_linked_security"
	TokenizationModelThirdPartySyntheticSwap        TokenizationModelT = "third_party_synthetic_security_based_swap"
	TokenizationModelPlatformEntitlementOrReceipt   TokenizationModelT = "platform_entitlement_or_receipt"
	TokenizationModelUnknown                        TokenizationModelT = "unknown"
)

func AllTokenizationEngineT() []TokenizationEngineT {
	return []TokenizationEngineT{
		TokenizationEngineNone,
		TokenizationEngineERC20,
		TokenizationEngineERC721,
		TokenizationEngineERC3643,
		TokenizationEngineERC7943,
	}
}

func AllTokenizationModelT() []TokenizationModelT {
	return []TokenizationModelT{
		TokenizationModelNone,
		TokenizationModelIssuerSponsoredOnchainRegister,
		TokenizationModelIssuerAuthorizedOffchain,
		TokenizationModelThirdPartyCustodialEntitlement,
		TokenizationModelThirdPartyLinkedSecurity,
		TokenizationModelThirdPartySyntheticSwap,
		TokenizationModelPlatformEntitlementOrReceipt,
		TokenizationModelUnknown,
	}
}

const (
	OfferRegTypeTUSRegAPlus              OfferRegTypeT = "US_REG_A_PLUS"
	OfferRegTypeTUSRegD506B              OfferRegTypeT = "US_REG_D_506B"
	OfferRegTypeTUSRegD506C              OfferRegTypeT = "US_REG_D_506C"
	OfferRegTypeTUSRegCF                 OfferRegTypeT = "US_REG_CF"
	OfferRegTypeTUSRegS                  OfferRegTypeT = "US_REG_S"
	OfferRegTypeTEUECSPCrowdfunding      OfferRegTypeT = "EU_ECSP_CROWDFUNDING"
	OfferRegTypeTEUMiCAOtherCryptoAsset  OfferRegTypeT = "EU_MICA_OTHER_CRYPTO_ASSET"
	OfferRegTypeTLUProspectusExempt      OfferRegTypeT = "LU_PROSPECTUS_EXEMPT"
	OfferRegTypeTDEEWPGCryptoSecurity    OfferRegTypeT = "DE_EWPG_CRYPTO_SECURITY"
	OfferRegTypeTCHFinSAPrivatePlacement OfferRegTypeT = "CH_FINSA_PRIVATE_PLACEMENT"
	OfferRegTypeTKYPrivateFund           OfferRegTypeT = "KY_PRIVATE_FUND"
	OfferRegTypeTKYVASPTokenIssuance     OfferRegTypeT = "KY_VASP_TOKEN_ISSUANCE"
	OfferRegTypeTKYTokenisedMutualFund   OfferRegTypeT = "KY_TOKENISED_MUTUAL_FUND"
	OfferRegTypeTKYSIBASecurityToken     OfferRegTypeT = "KY_SIBA_SECURITY_TOKEN"
)

func AllOfferRegTypeT() []OfferRegTypeT {
	return []OfferRegTypeT{
		OfferRegTypeTUSRegAPlus,
		OfferRegTypeTUSRegD506B,
		OfferRegTypeTUSRegD506C,
		OfferRegTypeTUSRegCF,
		OfferRegTypeTUSRegS,
		OfferRegTypeTEUECSPCrowdfunding,
		OfferRegTypeTEUMiCAOtherCryptoAsset,
		OfferRegTypeTLUProspectusExempt,
		OfferRegTypeTDEEWPGCryptoSecurity,
		OfferRegTypeTCHFinSAPrivatePlacement,
		OfferRegTypeTKYPrivateFund,
		OfferRegTypeTKYVASPTokenIssuance,
		OfferRegTypeTKYTokenisedMutualFund,
		OfferRegTypeTKYSIBASecurityToken,
	}
}

func (e OfferRegTypeT) IsValid() error {
	switch e {
	case OfferRegTypeTUSRegAPlus, OfferRegTypeTUSRegD506B, OfferRegTypeTUSRegD506C, OfferRegTypeTUSRegCF, OfferRegTypeTUSRegS, OfferRegTypeTEUECSPCrowdfunding, OfferRegTypeTEUMiCAOtherCryptoAsset, OfferRegTypeTLUProspectusExempt, OfferRegTypeTDEEWPGCryptoSecurity, OfferRegTypeTCHFinSAPrivatePlacement, OfferRegTypeTKYPrivateFund, OfferRegTypeTKYVASPTokenIssuance, OfferRegTypeTKYTokenisedMutualFund, OfferRegTypeTKYSIBASecurityToken:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

type OfferRegTypesT []OfferRegTypeT
type OfferJurisdictionsT []OfferJurisdictionT
type OfferingModesT []OfferingModeT

func scanOfferStringSlice(value any) ([]string, error) {
	var raw []byte
	switch val := value.(type) {
	case nil:
		return nil, nil
	case []byte:
		raw = val
	case string:
		raw = []byte(val)
	default:
		return nil, fmt.Errorf("cannot scan %T into JSON string slice", value)
	}

	var values []string
	if err := json.Unmarshal(raw, &values); err != nil {
		return nil, fmt.Errorf("decode JSON string slice: %w", err)
	}
	return values, nil
}

func (values *OfferRegTypesT) Scan(value any) error {
	strings, err := scanOfferStringSlice(value)
	if err != nil {
		return err
	}
	*values = make(OfferRegTypesT, len(strings))
	for index, value := range strings {
		(*values)[index] = OfferRegTypeT(value)
	}
	return nil
}

func (values OfferRegTypesT) Value() (driver.Value, error) {
	encoded, err := json.Marshal(values)
	if err != nil {
		return nil, fmt.Errorf("encode offer regulatory types: %w", err)
	}
	return encoded, nil
}

func (values *OfferJurisdictionsT) Scan(value any) error {
	strings, err := scanOfferStringSlice(value)
	if err != nil {
		return err
	}
	*values = make(OfferJurisdictionsT, len(strings))
	for index, value := range strings {
		(*values)[index] = OfferJurisdictionT(value)
	}
	return nil
}

func (values OfferJurisdictionsT) Value() (driver.Value, error) {
	encoded, err := json.Marshal(values)
	if err != nil {
		return nil, fmt.Errorf("encode offer jurisdictions: %w", err)
	}
	return encoded, nil
}

func (values *OfferingModesT) Scan(value any) error {
	strings, err := scanOfferStringSlice(value)
	if err != nil {
		return err
	}
	*values = make(OfferingModesT, len(strings))
	for index, value := range strings {
		(*values)[index] = OfferingModeT(value)
	}
	return nil
}

func (values OfferingModesT) Value() (driver.Value, error) {
	encoded, err := json.Marshal(values)
	if err != nil {
		return nil, fmt.Errorf("encode offering modes: %w", err)
	}
	return encoded, nil
}

type OfferJurisdictionT string

const (
	OfferJurisdictionTUS OfferJurisdictionT = "US"
	OfferJurisdictionTEU OfferJurisdictionT = "EU"
	OfferJurisdictionTLU OfferJurisdictionT = "LU"
	OfferJurisdictionTDE OfferJurisdictionT = "DE"
	OfferJurisdictionTCH OfferJurisdictionT = "CH"
	OfferJurisdictionTKY OfferJurisdictionT = "KY"
)

func AllOfferJurisdictionT() []OfferJurisdictionT {
	return []OfferJurisdictionT{
		OfferJurisdictionTUS,
		OfferJurisdictionTEU,
		OfferJurisdictionTLU,
		OfferJurisdictionTDE,
		OfferJurisdictionTCH,
		OfferJurisdictionTKY,
	}
}

func (e OfferJurisdictionT) IsValid() error {
	switch e {
	case OfferJurisdictionTUS, OfferJurisdictionTEU, OfferJurisdictionTLU, OfferJurisdictionTDE, OfferJurisdictionTCH, OfferJurisdictionTKY:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

type InstrumentClassificationT string

const (
	InstrumentClassificationTEquity             InstrumentClassificationT = "equity"
	InstrumentClassificationTDebt               InstrumentClassificationT = "debt"
	InstrumentClassificationTConvertibleNote    InstrumentClassificationT = "convertible_note"
	InstrumentClassificationTFundInterest       InstrumentClassificationT = "fund_interest"
	InstrumentClassificationTSecuritisationNote InstrumentClassificationT = "securitisation_note"
	InstrumentClassificationTAssetToken         InstrumentClassificationT = "asset_token"
	InstrumentClassificationTPaymentToken       InstrumentClassificationT = "payment_token"
	InstrumentClassificationTUtilityToken       InstrumentClassificationT = "utility_token"
	InstrumentClassificationTEMoneyToken        InstrumentClassificationT = "e_money_token"
	InstrumentClassificationTOtherCryptoAsset   InstrumentClassificationT = "other_crypto_asset"
	InstrumentClassificationTCryptoSecurity     InstrumentClassificationT = "crypto_security"
	InstrumentClassificationTDerivative         InstrumentClassificationT = "derivative"
)

func AllInstrumentClassificationT() []InstrumentClassificationT {
	return []InstrumentClassificationT{
		InstrumentClassificationTEquity,
		InstrumentClassificationTDebt,
		InstrumentClassificationTConvertibleNote,
		InstrumentClassificationTFundInterest,
		InstrumentClassificationTSecuritisationNote,
		InstrumentClassificationTAssetToken,
		InstrumentClassificationTPaymentToken,
		InstrumentClassificationTUtilityToken,
		InstrumentClassificationTEMoneyToken,
		InstrumentClassificationTOtherCryptoAsset,
		InstrumentClassificationTCryptoSecurity,
		InstrumentClassificationTDerivative,
	}
}

func (e InstrumentClassificationT) IsValid() error {
	switch e {
	case InstrumentClassificationTEquity, InstrumentClassificationTDebt, InstrumentClassificationTConvertibleNote, InstrumentClassificationTFundInterest, InstrumentClassificationTSecuritisationNote, InstrumentClassificationTAssetToken, InstrumentClassificationTPaymentToken, InstrumentClassificationTUtilityToken, InstrumentClassificationTEMoneyToken, InstrumentClassificationTOtherCryptoAsset, InstrumentClassificationTCryptoSecurity, InstrumentClassificationTDerivative:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

type OfferingModeT string

const (
	OfferingModeTPublicOffer            OfferingModeT = "public_offer"
	OfferingModeTPrivatePlacement       OfferingModeT = "private_placement"
	OfferingModeTProfessionalOnly       OfferingModeT = "professional_only"
	OfferingModeTQualifiedInvestorsOnly OfferingModeT = "qualified_investors_only"
	OfferingModeTCrowdfunding           OfferingModeT = "crowdfunding"
	OfferingModeTOffshore               OfferingModeT = "offshore"
	OfferingModeTRegistered             OfferingModeT = "registered"
)

func AllOfferingModeT() []OfferingModeT {
	return []OfferingModeT{
		OfferingModeTPublicOffer,
		OfferingModeTPrivatePlacement,
		OfferingModeTProfessionalOnly,
		OfferingModeTQualifiedInvestorsOnly,
		OfferingModeTCrowdfunding,
		OfferingModeTOffshore,
		OfferingModeTRegistered,
	}
}

func (e OfferingModeT) IsValid() error {
	switch e {
	case OfferingModeTPublicOffer, OfferingModeTPrivatePlacement, OfferingModeTProfessionalOnly, OfferingModeTQualifiedInvestorsOnly, OfferingModeTCrowdfunding, OfferingModeTOffshore, OfferingModeTRegistered:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

type CommentRelatedT string

// Enum values for CommentRelatedT
const (
	CommentRelatedTAdviser    CommentRelatedT = "adviser"
	CommentRelatedTEmployee   CommentRelatedT = "employee"
	CommentRelatedTAffiliated CommentRelatedT = "affiliated"
	CommentRelatedTInvestor   CommentRelatedT = "investor"
	CommentRelatedTNone       CommentRelatedT = "none"
)

func AllCommentRelatedT() []CommentRelatedT {
	return []CommentRelatedT{
		CommentRelatedTAdviser,
		CommentRelatedTEmployee,
		CommentRelatedTAffiliated,
		CommentRelatedTInvestor,
		CommentRelatedTNone,
	}
}

func (e CommentRelatedT) IsValid() error {
	switch e {
	case CommentRelatedTAdviser, CommentRelatedTEmployee, CommentRelatedTAffiliated, CommentRelatedTInvestor, CommentRelatedTNone:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e CommentRelatedT) String() string {
	return string(e)
}

// NullCommentRelatedT is a nullable CommentRelatedT enum type. It supports SQL and JSON serialization.
type NullCommentRelatedT struct {
	Val   CommentRelatedT
	Valid bool
}

// NullCommentRelatedTFrom creates a new CommentRelatedT that will never be blank.
func NullCommentRelatedTFrom(v CommentRelatedT) NullCommentRelatedT {
	return NewNullCommentRelatedT(v, true)
}

// NullCommentRelatedTFromPtr creates a new NullCommentRelatedT that be null if s is nil.
func NullCommentRelatedTFromPtr(v *CommentRelatedT) NullCommentRelatedT {
	if v == nil {
		return NewNullCommentRelatedT("", false)
	}
	return NewNullCommentRelatedT(*v, true)
}

// NewNullCommentRelatedT creates a new NullCommentRelatedT
func NewNullCommentRelatedT(v CommentRelatedT, valid bool) NullCommentRelatedT {
	return NullCommentRelatedT{
		Val:   v,
		Valid: valid,
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *NullCommentRelatedT) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &e.Val); err != nil {
		return err
	}

	e.Valid = true
	return nil
}

// MarshalJSON implements json.Marshaler.
func (e NullCommentRelatedT) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Val)
}

// MarshalText implements encoding.TextMarshaler.
func (e NullCommentRelatedT) MarshalText() ([]byte, error) {
	if !e.Valid {
		return []byte{}, nil
	}
	return []byte(e.Val), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (e *NullCommentRelatedT) UnmarshalText(text []byte) error {
	if text == nil || len(text) == 0 {
		e.Valid = false
		return nil
	}

	e.Val = CommentRelatedT(text)
	e.Valid = true
	return nil
}

// SetValid changes this NullCommentRelatedT value and also sets it to be non-null.
func (e *NullCommentRelatedT) SetValid(v CommentRelatedT) {
	e.Val = v
	e.Valid = true
}

// Ptr returns a pointer to this NullCommentRelatedT value, or a nil pointer if this NullCommentRelatedT is null.
func (e NullCommentRelatedT) Ptr() *CommentRelatedT {
	if !e.Valid {
		return nil
	}
	return &e.Val
}

// IsZero returns true for null types.
func (e NullCommentRelatedT) IsZero() bool {
	return !e.Valid
}

// Value implements the driver Valuer interface.
func (e NullCommentRelatedT) Value() (any, error) {
	if !e.Valid {
		return nil, nil
	}
	return string(e.Val), nil
}

type CommentStatusT string

// Enum values for CommentStatusT
const (
	CommentStatusTNew       CommentStatusT = "new"
	CommentStatusTPublished CommentStatusT = "published"
	CommentStatusTDeleted   CommentStatusT = "deleted"
)

func AllCommentStatusT() []CommentStatusT {
	return []CommentStatusT{
		CommentStatusTNew,
		CommentStatusTPublished,
		CommentStatusTDeleted,
	}
}

func (e CommentStatusT) IsValid() error {
	switch e {
	case CommentStatusTNew, CommentStatusTPublished, CommentStatusTDeleted:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e CommentStatusT) String() string {
	return string(e)
}

// NullCommentStatusT is a nullable CommentStatusT enum type. It supports SQL and JSON serialization.
type NullCommentStatusT struct {
	Val   CommentStatusT
	Valid bool
}

// NullCommentStatusTFrom creates a new CommentStatusT that will never be blank.
func NullCommentStatusTFrom(v CommentStatusT) NullCommentStatusT {
	return NewNullCommentStatusT(v, true)
}

// NullCommentStatusTFromPtr creates a new NullCommentStatusT that be null if s is nil.
func NullCommentStatusTFromPtr(v *CommentStatusT) NullCommentStatusT {
	if v == nil {
		return NewNullCommentStatusT("", false)
	}
	return NewNullCommentStatusT(*v, true)
}

// NewNullCommentStatusT creates a new NullCommentStatusT
func NewNullCommentStatusT(v CommentStatusT, valid bool) NullCommentStatusT {
	return NullCommentStatusT{
		Val:   v,
		Valid: valid,
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *NullCommentStatusT) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &e.Val); err != nil {
		return err
	}

	e.Valid = true
	return nil
}

// MarshalJSON implements json.Marshaler.
func (e NullCommentStatusT) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Val)
}

// MarshalText implements encoding.TextMarshaler.
func (e NullCommentStatusT) MarshalText() ([]byte, error) {
	if !e.Valid {
		return []byte{}, nil
	}
	return []byte(e.Val), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (e *NullCommentStatusT) UnmarshalText(text []byte) error {
	if text == nil || len(text) == 0 {
		e.Valid = false
		return nil
	}

	e.Val = CommentStatusT(text)
	e.Valid = true
	return nil
}

// SetValid changes this NullCommentStatusT value and also sets it to be non-null.
func (e *NullCommentStatusT) SetValid(v CommentStatusT) {
	e.Val = v
	e.Valid = true
}

// Ptr returns a pointer to this NullCommentStatusT value, or a nil pointer if this NullCommentStatusT is null.
func (e NullCommentStatusT) Ptr() *CommentStatusT {
	if !e.Valid {
		return nil
	}
	return &e.Val
}

// IsZero returns true for null types.
func (e NullCommentStatusT) IsZero() bool {
	return !e.Valid
}

// Value implements the driver Valuer interface.
func (e NullCommentStatusT) Value() (any, error) {
	if !e.Valid {
		return nil, nil
	}
	return string(e.Val), nil
}
