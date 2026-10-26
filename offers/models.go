package offers

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/global-torque/go-common/db/v2"
	"github.com/global-torque/go-common/orm/v2"
	"github.com/global-torque/go-common/orm/v2/pgtype"
	"github.com/pkg/errors"
)

type OfferOffer struct {
	ID                       int                       `db:"id" json:"id" yaml:"id"`
	UserID                   int                       `db:"user_id" json:"user_id" yaml:"user_id"`
	Name                     string                    `db:"name" json:"name" yaml:"name"`
	Slug                     string                    `db:"slug" json:"slug" yaml:"slug"`
	MinInvestment            string                    `db:"min_investment" json:"min_investment" yaml:"min_investment"`
	Description              string                    `db:"description" json:"description" yaml:"description"`
	Title                    string                    `db:"title" json:"title" yaml:"title"`
	Highlights               string                    `db:"highlights" json:"highlights" yaml:"highlights"`
	Valuation                float64                   `db:"valuation" json:"valuation" yaml:"valuation"`
	TotalShares              string                    `db:"total_shares" json:"total_shares" yaml:"total_shares"`
	PricePerShare            string                    `db:"price_per_share" json:"price_per_share" yaml:"price_per_share"`
	Status                   NullOfferT                `db:"status" json:"status,omitempty" yaml:"status,omitempty"`
	SecurityType             OfferSecurityTypeT        `db:"security_type" json:"security_type" yaml:"security_type"`
	Notes                    string                    `db:"notes" json:"notes" yaml:"notes"`
	SubscribedShares         string                    `db:"subscribed_shares" json:"subscribed_shares" yaml:"subscribed_shares"`
	ConfirmedShares          string                    `db:"confirmed_shares" json:"confirmed_shares" yaml:"confirmed_shares"`
	AdditionalDetails        string                    `db:"additional_details" json:"additional_details" yaml:"additional_details"`
	SeoTitle                 string                    `db:"seo_title" json:"seo_title" yaml:"seo_title"`
	SeoDescription           string                    `db:"seo_description" json:"seo_description" yaml:"seo_description"`
	Website                  string                    `db:"website" json:"website" yaml:"website"`
	State                    string                    `db:"state" json:"state" yaml:"state"`
	City                     string                    `db:"city" json:"city" yaml:"city"`
	Address1                 string                    `db:"address1" json:"address1" yaml:"address1"`
	Address2                 string                    `db:"address2" json:"address2" yaml:"address2"`
	ZipCode                  string                    `db:"zip_code" json:"zip_code" yaml:"zip_code"`
	Country                  string                    `db:"country" json:"country" yaml:"country"`
	EntityID                 string                    `db:"entity_id" json:"entity_id" yaml:"entity_id"`
	Data                     any                       `db:"data" json:"data" yaml:"data"`
	SecurityInfo             any                       `db:"security_info" json:"security_info,omitempty" yaml:"security_info,omitempty"`
	Ticker                   string                    `db:"ticker" json:"ticker,omitempty" yaml:"ticker,omitempty"`
	StartAt                  pgtype.Timestamptz        `db:"start_at" json:"start_at,omitempty" yaml:"start_at,omitempty"`
	CloseAt                  pgtype.Timestamptz        `db:"close_at" json:"close_at,omitempty" yaml:"close_at,omitempty"`
	ApprovedAt               pgtype.Timestamptz        `db:"approved_at" json:"approved_at,omitempty" yaml:"approved_at,omitempty"`
	CreatedAt                pgtype.Timestamptz        `db:"created_at" json:"created_at" yaml:"created_at"`
	UpdatedAt                pgtype.Timestamptz        `db:"updated_at" json:"updated_at" yaml:"updated_at"`
	ImageLinkID              *int                      `db:"image_link_id" json:"image_link_id,omitempty" yaml:"image_link_id,omitempty"`
	LegalName                string                    `db:"legal_name" json:"legal_name" yaml:"legal_name"`
	Twitter                  string                    `db:"twitter" json:"twitter" yaml:"twitter"`
	Facebook                 string                    `db:"facebook" json:"facebook" yaml:"facebook"`
	Github                   string                    `db:"github" json:"github" yaml:"github"`
	Instagram                string                    `db:"instagram" json:"instagram" yaml:"instagram"`
	Telegram                 string                    `db:"telegram" json:"telegram" yaml:"telegram"`
	Linkedin                 string                    `db:"linkedin" json:"linkedin" yaml:"linkedin"`
	Mastodon                 string                    `db:"mastodon" json:"mastodon" yaml:"mastodon"`
	EsignID                  string                    `db:"esign_id" json:"esign_id" yaml:"esign_id"`
	RiskDisclosures          string                    `db:"risk_disclosures" json:"risk_disclosures,omitempty" yaml:"risk_disclosures,omitempty"`
	TokenizationEngine       TokenizationEngineT       `db:"tokenization_engine" json:"tokenization_engine" yaml:"tokenization_engine"`
	LegalRepresentativeID    *int                      `db:"legal_representative_id" json:"legal_representative_id,omitempty" yaml:"legal_representative_id,omitempty"`
	PlatformRepresentativeID *int                      `db:"platform_representative_id" json:"platform_representative_id,omitempty" yaml:"platform_representative_id,omitempty"`
	SiteID                   *int                      `db:"site_id" json:"site_id,omitempty" yaml:"site_id,omitempty"`
	DataRoomGroupID          *int                      `db:"data_room_group_id" json:"data_room_group_id,omitempty" yaml:"data_room_group_id,omitempty"`
	FundStructure            FundStructureT            `db:"fund_structure" json:"fund_structure" yaml:"fund_structure"`
	RegType                  OfferRegTypesT            `db:"reg_type" json:"reg_type" yaml:"reg_type"`
	OfferJurisdiction        OfferJurisdictionsT       `db:"offer_jurisdiction" json:"offer_jurisdiction" yaml:"offer_jurisdiction"`
	InstrumentClassification InstrumentClassificationT `db:"instrument_classification" json:"instrument_classification" yaml:"instrument_classification"`
	OfferingMode             OfferingModesT            `db:"offering_mode" json:"offering_mode" yaml:"offering_mode"`
	OfferingCurrency         string                    `db:"offering_currency" json:"offering_currency" yaml:"offering_currency"`
	InvestorEligibility      any                       `db:"investor_eligibility" json:"investor_eligibility,omitempty" yaml:"investor_eligibility,omitempty"`
	MarketedJurisdictions    any                       `db:"marketed_jurisdictions" json:"marketed_jurisdictions,omitempty" yaml:"marketed_jurisdictions,omitempty"`
	RegulatedActivities      any                       `db:"regulated_activities" json:"regulated_activities,omitempty" yaml:"regulated_activities,omitempty"`
	TokenizationModel        TokenizationModelT        `db:"tokenization_model" json:"tokenization_model" yaml:"tokenization_model"`

	updatedFields []string       `db:"-" json:"-"`
	fns           map[string]any `db:"-" json:"-"`
	db            db.Repository  `db:"-" json:"-"`
}

func New(db db.Repository) *OfferOffer {
	return &OfferOffer{
		db:  db,
		fns: map[string]any{},
	}
}

func (model OfferOffer) ToJSON() map[string]any {
	return map[string]any{
		"id":                         model.ID,
		"user_id":                    model.UserID,
		"name":                       model.Name,
		"slug":                       model.Slug,
		"min_investment":             model.MinInvestment,
		"description":                model.Description,
		"title":                      model.Title,
		"highlights":                 model.Highlights,
		"valuation":                  model.Valuation,
		"total_shares":               model.TotalShares,
		"price_per_share":            model.PricePerShare,
		"status":                     model.Status,
		"security_type":              model.SecurityType,
		"notes":                      model.Notes,
		"subscribed_shares":          model.SubscribedShares,
		"confirmed_shares":           model.ConfirmedShares,
		"additional_details":         model.AdditionalDetails,
		"seo_title":                  model.SeoTitle,
		"seo_description":            model.SeoDescription,
		"website":                    model.Website,
		"state":                      model.State,
		"city":                       model.City,
		"address1":                   model.Address1,
		"address2":                   model.Address2,
		"zip_code":                   model.ZipCode,
		"country":                    model.Country,
		"entity_id":                  model.EntityID,
		"data":                       model.Data,
		"security_info":              model.SecurityInfo,
		"ticker":                     model.Ticker,
		"start_at":                   model.StartAt,
		"close_at":                   model.CloseAt,
		"approved_at":                model.ApprovedAt,
		"created_at":                 model.CreatedAt,
		"updated_at":                 model.UpdatedAt,
		"image_link_id":              model.ImageLinkID,
		"legal_name":                 model.LegalName,
		"twitter":                    model.Twitter,
		"facebook":                   model.Facebook,
		"github":                     model.Github,
		"instagram":                  model.Instagram,
		"telegram":                   model.Telegram,
		"linkedin":                   model.Linkedin,
		"mastodon":                   model.Mastodon,
		"esign_id":                   model.EsignID,
		"risk_disclosures":           model.RiskDisclosures,
		"tokenization_engine":        model.TokenizationEngine,
		"legal_representative_id":    model.LegalRepresentativeID,
		"platform_representative_id": model.PlatformRepresentativeID,
		"site_id":                    model.SiteID,
		"data_room_group_id":         model.DataRoomGroupID,
		"fund_structure":             model.FundStructure,
		"reg_type":                   model.RegType,
		"offer_jurisdiction":         model.OfferJurisdiction,
		"instrument_classification":  model.InstrumentClassification,
		"offering_mode":              model.OfferingMode,
		"offering_currency":          model.OfferingCurrency,
		"investor_eligibility":       model.InvestorEligibility,
		"marketed_jurisdictions":     model.MarketedJurisdictions,
		"regulated_activities":       model.RegulatedActivities,
		"tokenization_model":         model.TokenizationModel,
	}
}

func (model OfferOffer) Fields() []string {
	return []string{
		"id",
		"user_id",
		"name",
		"slug",
		"min_investment",
		"description",
		"title",
		"highlights",
		"valuation",
		"total_shares",
		"price_per_share",
		"status",
		"security_type",
		"notes",
		"subscribed_shares",
		"confirmed_shares",
		"additional_details",
		"seo_title",
		"seo_description",
		"website",
		"state",
		"city",
		"address1",
		"address2",
		"zip_code",
		"country",
		"entity_id",
		"data",
		"security_info",
		"ticker",
		"start_at",
		"close_at",
		"approved_at",
		"created_at",
		"updated_at",
		"image_link_id",
		"legal_name",
		"twitter",
		"facebook",
		"github",
		"instagram",
		"telegram",
		"linkedin",
		"mastodon",
		"esign_id",
		"risk_disclosures",
		"tokenization_engine",
		"legal_representative_id",
		"platform_representative_id",
		"site_id",
		"data_room_group_id",
		"fund_structure",
		"reg_type",
		"offer_jurisdiction",
		"instrument_classification",
		"offering_mode",
		"offering_currency",
		"investor_eligibility",
		"marketed_jurisdictions",
		"regulated_activities",
		"tokenization_model",
	}
}

func (model OfferOffer) Table() string {
	return TableName
}

func (model OfferOffer) GetID() any {
	return model.ID
}

func (model *OfferOffer) SetID(id any) {
	model.ID = id.(int)
}

func (model *OfferOffer) SetDB(db db.Repository) {
	model.db = db
}

func Get(ctx context.Context, db db.Repository, where map[string]any) (*OfferOffer, error) {
	model, err := orm.RetrieveOne[OfferOffer](
		ctx,
		db,
		sq.Eq(where),
	)
	if err != nil {
		err = errors.Wrapf(err, "cannot get model")
		return nil, err
	}
	model.db = db
	return model, nil
}
