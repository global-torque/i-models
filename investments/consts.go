package investments

import "github.com/pkg/errors"

type PaymentT string

// Enum values for PaymentT
const (
	PaymentTNone         PaymentT = "none"
	PaymentTDwolla       PaymentT = "dwolla"
	PaymentTNorthCapital PaymentT = "north_capital"
	PaymentTPrimeTrust   PaymentT = "prime_trust"
	PaymentTApexGroup    PaymentT = "apex_group"
	PaymentTEVM          PaymentT = "evm"
	PaymentTSolana       PaymentT = "solana"

	AppLabel  = "investment"
	ModelName = "investment"
)

func AllPaymentT() []PaymentT {
	return []PaymentT{
		PaymentTNone,
		PaymentTDwolla,
		PaymentTNorthCapital,
		PaymentTPrimeTrust,
		PaymentTApexGroup,
		PaymentTEVM,
		PaymentTSolana,
	}
}

func (e PaymentT) IsValid() error {
	switch e {
	case PaymentTNone, PaymentTDwolla, PaymentTNorthCapital, PaymentTPrimeTrust, PaymentTApexGroup, PaymentTEVM, PaymentTSolana:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e PaymentT) String() string {
	return string(e)
}

type EscrowT string

// Enum values for EscrowT
const (
	EscrowTNone         EscrowT = "none"
	EscrowTNorthCapital EscrowT = "north_capital"
	EscrowTPrimeTrust   EscrowT = "prime_trust"
	EscrowTApexGroup    EscrowT = "apex_group"
)

func AllEscrowT() []EscrowT {
	return []EscrowT{
		EscrowTNone,
		EscrowTNorthCapital,
		EscrowTPrimeTrust,
		EscrowTApexGroup,
	}
}

func (e EscrowT) IsValid() error {
	switch e {
	case EscrowTNone, EscrowTNorthCapital, EscrowTPrimeTrust, EscrowTApexGroup:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e EscrowT) String() string {
	return string(e)
}

type FundingT string

// Enum values for FundingT
const (
	FundingTNone         FundingT = "none"
	FundingTWire         FundingT = "wire"
	FundingTAch          FundingT = "ach"
	FundingTWallet       FundingT = "wallet"
	FundingTCryptoWallet FundingT = "crypto_wallet"
)

func AllFundingT() []FundingT {
	return []FundingT{
		FundingTNone,
		FundingTWire,
		FundingTAch,
		FundingTWallet,
		FundingTCryptoWallet,
	}
}

func (e FundingT) IsValid() error {
	switch e {
	case FundingTNone, FundingTWire, FundingTAch, FundingTWallet, FundingTCryptoWallet:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e FundingT) String() string {
	return string(e)
}

type FundingS string

// Enum values for FundingS
const (
	FundingSCreationError   FundingS = "creation_error"
	FundingSNew             FundingS = "new"
	FundingSInitialize      FundingS = "initialize"
	FundingSInProgress      FundingS = "in_progress"
	FundingSReceived        FundingS = "received"
	FundingSSettled         FundingS = "settled"
	FundingSFailed          FundingS = "failed"
	FundingSSentBackPending FundingS = "sent_back_pending"
	FundingSSentBackSettled FundingS = "sent_back_settled"
	FundingSCanceled        FundingS = "canceled"
)

func AllFundingS() []FundingS {
	return []FundingS{
		FundingSCreationError,
		FundingSNew,
		FundingSInitialize,
		FundingSInProgress,
		FundingSReceived,
		FundingSSettled,
		FundingSFailed,
		FundingSSentBackPending,
		FundingSSentBackSettled,
		FundingSCanceled,
	}
}

func (e FundingS) IsValid() error {
	switch e {
	case FundingSCreationError, FundingSNew, FundingSInitialize, FundingSInProgress, FundingSReceived, FundingSSettled, FundingSFailed, FundingSSentBackPending, FundingSSentBackSettled, FundingSCanceled:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e FundingS) String() string {
	return string(e)
}

type InvestmentT string

// Enum values for InvestmentT
const (
	InvestmentTNew                       InvestmentT = "new"
	InvestmentTConfirmed                 InvestmentT = "confirmed"
	InvestmentTLegallyConfirmed          InvestmentT = "legally_confirmed"
	InvestmentTClosedSuccessfully        InvestmentT = "closed_successfully"
	InvestmentTCancelledDuringInvestment InvestmentT = "cancelled_during_investment"
	InvestmentTCancelledAfterInvestment  InvestmentT = "cancelled_after_investment"
	InvestmentTCancelledByManager        InvestmentT = "cancelled_by_manager"
	InvestmentTSold                      InvestmentT = "sold"
	InvestmentTExited                    InvestmentT = "exited"
	InvestmentTSystemError               InvestmentT = "system_error"
)

func AllInvestmentT() []InvestmentT {
	return []InvestmentT{
		InvestmentTNew,
		InvestmentTConfirmed,
		InvestmentTLegallyConfirmed,
		InvestmentTClosedSuccessfully,
		InvestmentTCancelledDuringInvestment,
		InvestmentTCancelledAfterInvestment,
		InvestmentTCancelledByManager,
		InvestmentTSold,
		InvestmentTExited,
		InvestmentTSystemError,
	}
}

func (e InvestmentT) IsValid() error {
	switch e {
	case InvestmentTNew, InvestmentTConfirmed, InvestmentTLegallyConfirmed, InvestmentTClosedSuccessfully, InvestmentTCancelledDuringInvestment, InvestmentTCancelledAfterInvestment, InvestmentTCancelledByManager, InvestmentTSold, InvestmentTExited, InvestmentTSystemError:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e InvestmentT) String() string {
	return string(e)
}

type InvestmentStepT string

// Enum values for InvestmentStepT
const (
	InvestmentStepTNew           InvestmentStepT = "new"
	InvestmentStepTAmount        InvestmentStepT = "amount"
	InvestmentStepTOwnership     InvestmentStepT = "ownership"
	InvestmentStepTSignature     InvestmentStepT = "signature"
	InvestmentStepTFunding       InvestmentStepT = "funding"
	InvestmentStepTAccreditation InvestmentStepT = "accreditation"
	InvestmentStepTConfirmation  InvestmentStepT = "confirmation"
	InvestmentStepTReview        InvestmentStepT = "review"
)

func AllInvestmentStepT() []InvestmentStepT {
	return []InvestmentStepT{
		InvestmentStepTNew,
		InvestmentStepTAmount,
		InvestmentStepTOwnership,
		InvestmentStepTSignature,
		InvestmentStepTFunding,
		InvestmentStepTAccreditation,
		InvestmentStepTConfirmation,
		InvestmentStepTReview,
	}
}

func (e InvestmentStepT) IsValid() error {
	switch e {
	case InvestmentStepTNew, InvestmentStepTAmount, InvestmentStepTOwnership, InvestmentStepTSignature, InvestmentStepTFunding, InvestmentStepTAccreditation, InvestmentStepTConfirmation, InvestmentStepTReview:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e InvestmentStepT) String() string {
	return string(e)
}

type VaultRequestOriginT string

const (
	VaultRequestOriginApplication VaultRequestOriginT = "application"
	VaultRequestOriginChain       VaultRequestOriginT = "chain"
)

func AllVaultRequestOriginT() []VaultRequestOriginT {
	return []VaultRequestOriginT{VaultRequestOriginApplication, VaultRequestOriginChain}
}

func (e VaultRequestOriginT) IsValid() error {
	switch e {
	case VaultRequestOriginApplication, VaultRequestOriginChain:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e VaultRequestOriginT) String() string {
	return string(e)
}

type VaultPricingStatusT string

const (
	VaultPricingStatusAwaitingDealingNAV VaultPricingStatusT = "awaiting_dealing_nav"
	VaultPricingStatusPriced             VaultPricingStatusT = "priced"
)

func AllVaultPricingStatusT() []VaultPricingStatusT {
	return []VaultPricingStatusT{VaultPricingStatusAwaitingDealingNAV, VaultPricingStatusPriced}
}

func (e VaultPricingStatusT) IsValid() error {
	switch e {
	case VaultPricingStatusAwaitingDealingNAV, VaultPricingStatusPriced:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e VaultPricingStatusT) String() string {
	return string(e)
}

type DepositPriceSourceT string

const (
	DepositPriceSourceOfferDeck       DepositPriceSourceT = "offer_deck"
	DepositPriceSourceFinalizedNAV    DepositPriceSourceT = "finalized_nav"
	DepositPriceSourceChainForwardNAV DepositPriceSourceT = "chain_forward_nav"
	DepositPriceSourceManagerDealing  DepositPriceSourceT = "manager_dealing_price"
)

func AllDepositPriceSourceT() []DepositPriceSourceT {
	return []DepositPriceSourceT{
		DepositPriceSourceOfferDeck,
		DepositPriceSourceFinalizedNAV,
		DepositPriceSourceChainForwardNAV,
		DepositPriceSourceManagerDealing,
	}
}

func (e DepositPriceSourceT) IsValid() error {
	switch e {
	case DepositPriceSourceOfferDeck, DepositPriceSourceFinalizedNAV,
		DepositPriceSourceChainForwardNAV, DepositPriceSourceManagerDealing:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e DepositPriceSourceT) String() string {
	return string(e)
}

type RedemptionStatusT string

const (
	RedemptionStatusPending   RedemptionStatusT = "pending"
	RedemptionStatusApproved  RedemptionStatusT = "approved"
	RedemptionStatusDenied    RedemptionStatusT = "denied"
	RedemptionStatusCancelled RedemptionStatusT = "cancelled"
	RedemptionStatusCompleted RedemptionStatusT = "completed"
)

func AllRedemptionStatusT() []RedemptionStatusT {
	return []RedemptionStatusT{
		RedemptionStatusPending,
		RedemptionStatusApproved,
		RedemptionStatusDenied,
		RedemptionStatusCancelled,
		RedemptionStatusCompleted,
	}
}

func (e RedemptionStatusT) IsValid() error {
	switch e {
	case RedemptionStatusPending, RedemptionStatusApproved, RedemptionStatusDenied,
		RedemptionStatusCancelled, RedemptionStatusCompleted:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e RedemptionStatusT) String() string {
	return string(e)
}

type ProfileT string

// Enum values for ProfileT
const (
	ProfileTIndividual ProfileT = "individual"
	ProfileTEntity     ProfileT = "entity"
	ProfileTTrust      ProfileT = "trust"
	ProfileTSdira      ProfileT = "sdira"
	ProfileTSolo401k   ProfileT = "solo401k"
)

func AllProfileT() []ProfileT {
	return []ProfileT{
		ProfileTIndividual,
		ProfileTEntity,
		ProfileTTrust,
		ProfileTSdira,
		ProfileTSolo401k,
	}
}

func (e ProfileT) IsValid() error {
	switch e {
	case ProfileTIndividual, ProfileTEntity, ProfileTTrust, ProfileTSdira, ProfileTSolo401k:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e ProfileT) String() string {
	return string(e)
}

type KycT string

// Enum values for KycT
const (
	KycTNew      KycT = "new"
	KycTPending  KycT = "pending"
	KycTApproved KycT = "approved"
	KycTDeclined KycT = "declined"
)

func AllKycT() []KycT {
	return []KycT{
		KycTNew,
		KycTPending,
		KycTApproved,
		KycTDeclined,
	}
}

func (e KycT) IsValid() error {
	switch e {
	case KycTNew, KycTPending, KycTApproved, KycTDeclined:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e KycT) String() string {
	return string(e)
}

type AccreditationT string

// Enum values for AccreditationT
const (
	AccreditationTNew              AccreditationT = "new"
	AccreditationTPending          AccreditationT = "pending"
	AccreditationTDocumentRequired AccreditationT = "document_required"
	AccreditationTDocumentPending  AccreditationT = "document_pending"
	AccreditationTInfoRequired     AccreditationT = "info_required"
	AccreditationTInfoPending      AccreditationT = "info_pending"
	AccreditationTExpired          AccreditationT = "expired"
	AccreditationTApproved         AccreditationT = "approved"
	AccreditationTDeclined         AccreditationT = "declined"
)

func AllAccreditationT() []AccreditationT {
	return []AccreditationT{
		AccreditationTNew,
		AccreditationTPending,
		AccreditationTDocumentRequired,
		AccreditationTDocumentPending,
		AccreditationTInfoRequired,
		AccreditationTInfoPending,
		AccreditationTExpired,
		AccreditationTApproved,
		AccreditationTDeclined,
	}
}

func (e AccreditationT) IsValid() error {
	switch e {
	case AccreditationTNew, AccreditationTPending, AccreditationTDocumentRequired, AccreditationTDocumentPending, AccreditationTInfoRequired, AccreditationTInfoPending, AccreditationTExpired, AccreditationTApproved, AccreditationTDeclined:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e AccreditationT) String() string {
	return string(e)
}
