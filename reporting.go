// Code generated from our API definition. DO NOT EDIT.

package api

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/mailchimp/mailchimp-marketing-go-sdk/internal"
	big "math/big"
	time "time"
)

var (
	getFacebookAdReportingRequestFieldOutreachID    = big.NewInt(1 << 0)
	getFacebookAdReportingRequestFieldFields        = big.NewInt(1 << 1)
	getFacebookAdReportingRequestFieldExcludeFields = big.NewInt(1 << 2)
)

type GetFacebookAdReportingRequest struct {
	// The outreach id.
	OutreachID string `json:"-" url:"-"`
	// A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
	Fields []*string `json:"-" url:"fields,omitempty"`
	// A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
	ExcludeFields []*string `json:"-" url:"exclude_fields,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetFacebookAdReportingRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetOutreachID sets the OutreachID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetFacebookAdReportingRequest) SetOutreachID(outreachID string) {
	g.OutreachID = outreachID
	g.require(getFacebookAdReportingRequestFieldOutreachID)
}

// SetFields sets the Fields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetFacebookAdReportingRequest) SetFields(fields []*string) {
	g.Fields = fields
	g.require(getFacebookAdReportingRequestFieldFields)
}

// SetExcludeFields sets the ExcludeFields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetFacebookAdReportingRequest) SetExcludeFields(excludeFields []*string) {
	g.ExcludeFields = excludeFields
	g.require(getFacebookAdReportingRequestFieldExcludeFields)
}

var (
	getLandingPageReportingRequestFieldOutreachID    = big.NewInt(1 << 0)
	getLandingPageReportingRequestFieldFields        = big.NewInt(1 << 1)
	getLandingPageReportingRequestFieldExcludeFields = big.NewInt(1 << 2)
)

type GetLandingPageReportingRequest struct {
	// The outreach id.
	OutreachID string `json:"-" url:"-"`
	// A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
	Fields []*string `json:"-" url:"fields,omitempty"`
	// A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
	ExcludeFields []*string `json:"-" url:"exclude_fields,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetLandingPageReportingRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetOutreachID sets the OutreachID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetLandingPageReportingRequest) SetOutreachID(outreachID string) {
	g.OutreachID = outreachID
	g.require(getLandingPageReportingRequestFieldOutreachID)
}

// SetFields sets the Fields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetLandingPageReportingRequest) SetFields(fields []*string) {
	g.Fields = fields
	g.require(getLandingPageReportingRequestFieldFields)
}

// SetExcludeFields sets the ExcludeFields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetLandingPageReportingRequest) SetExcludeFields(excludeFields []*string) {
	g.ExcludeFields = excludeFields
	g.require(getLandingPageReportingRequestFieldExcludeFields)
}

var (
	getSurveyReportingRequestFieldSurveyID      = big.NewInt(1 << 0)
	getSurveyReportingRequestFieldFields        = big.NewInt(1 << 1)
	getSurveyReportingRequestFieldExcludeFields = big.NewInt(1 << 2)
)

type GetSurveyReportingRequest struct {
	// The ID of the survey.
	SurveyID string `json:"-" url:"-"`
	// A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
	Fields []*string `json:"-" url:"fields,omitempty"`
	// A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
	ExcludeFields []*string `json:"-" url:"exclude_fields,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetSurveyReportingRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetSurveyID sets the SurveyID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyReportingRequest) SetSurveyID(surveyID string) {
	g.SurveyID = surveyID
	g.require(getSurveyReportingRequestFieldSurveyID)
}

// SetFields sets the Fields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyReportingRequest) SetFields(fields []*string) {
	g.Fields = fields
	g.require(getSurveyReportingRequestFieldFields)
}

// SetExcludeFields sets the ExcludeFields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyReportingRequest) SetExcludeFields(excludeFields []*string) {
	g.ExcludeFields = excludeFields
	g.require(getSurveyReportingRequestFieldExcludeFields)
}

var (
	getSurveyQuestionReportingRequestFieldSurveyID      = big.NewInt(1 << 0)
	getSurveyQuestionReportingRequestFieldQuestionID    = big.NewInt(1 << 1)
	getSurveyQuestionReportingRequestFieldFields        = big.NewInt(1 << 2)
	getSurveyQuestionReportingRequestFieldExcludeFields = big.NewInt(1 << 3)
)

type GetSurveyQuestionReportingRequest struct {
	// The ID of the survey.
	SurveyID string `json:"-" url:"-"`
	// The ID of the survey question.
	QuestionID string `json:"-" url:"-"`
	// A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
	Fields []*string `json:"-" url:"fields,omitempty"`
	// A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
	ExcludeFields []*string `json:"-" url:"exclude_fields,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetSurveyQuestionReportingRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetSurveyID sets the SurveyID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyQuestionReportingRequest) SetSurveyID(surveyID string) {
	g.SurveyID = surveyID
	g.require(getSurveyQuestionReportingRequestFieldSurveyID)
}

// SetQuestionID sets the QuestionID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyQuestionReportingRequest) SetQuestionID(questionID string) {
	g.QuestionID = questionID
	g.require(getSurveyQuestionReportingRequestFieldQuestionID)
}

// SetFields sets the Fields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyQuestionReportingRequest) SetFields(fields []*string) {
	g.Fields = fields
	g.require(getSurveyQuestionReportingRequestFieldFields)
}

// SetExcludeFields sets the ExcludeFields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyQuestionReportingRequest) SetExcludeFields(excludeFields []*string) {
	g.ExcludeFields = excludeFields
	g.require(getSurveyQuestionReportingRequestFieldExcludeFields)
}

var (
	getSurveyResponsReportingRequestFieldSurveyID   = big.NewInt(1 << 0)
	getSurveyResponsReportingRequestFieldResponseID = big.NewInt(1 << 1)
)

type GetSurveyResponsReportingRequest struct {
	// The ID of the survey.
	SurveyID string `json:"-" url:"-"`
	// The ID of the survey response.
	ResponseID string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetSurveyResponsReportingRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetSurveyID sets the SurveyID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyResponsReportingRequest) SetSurveyID(surveyID string) {
	g.SurveyID = surveyID
	g.require(getSurveyResponsReportingRequestFieldSurveyID)
}

// SetResponseID sets the ResponseID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyResponsReportingRequest) SetResponseID(responseID string) {
	g.ResponseID = responseID
	g.require(getSurveyResponsReportingRequestFieldResponseID)
}

var (
	listFacebookAdEcommerceProductActivityReportingRequestFieldOutreachID    = big.NewInt(1 << 0)
	listFacebookAdEcommerceProductActivityReportingRequestFieldFields        = big.NewInt(1 << 1)
	listFacebookAdEcommerceProductActivityReportingRequestFieldExcludeFields = big.NewInt(1 << 2)
	listFacebookAdEcommerceProductActivityReportingRequestFieldCount         = big.NewInt(1 << 3)
	listFacebookAdEcommerceProductActivityReportingRequestFieldOffset        = big.NewInt(1 << 4)
	listFacebookAdEcommerceProductActivityReportingRequestFieldSortField     = big.NewInt(1 << 5)
)

type ListFacebookAdEcommerceProductActivityReportingRequest struct {
	// The outreach id.
	OutreachID string `json:"-" url:"-"`
	// A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
	Fields []*string `json:"-" url:"fields,omitempty"`
	// A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
	ExcludeFields []*string `json:"-" url:"exclude_fields,omitempty"`
	// The number of records to return. Default value is 10. Maximum value is 1000
	Count *int `json:"-" url:"count,omitempty"`
	// Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
	Offset *int `json:"-" url:"offset,omitempty"`
	// Returns files sorted by the specified field.
	SortField *ListFacebookAdEcommerceProductActivityReportingRequestSortField `json:"-" url:"sort_field,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListFacebookAdEcommerceProductActivityReportingRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetOutreachID sets the OutreachID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdEcommerceProductActivityReportingRequest) SetOutreachID(outreachID string) {
	l.OutreachID = outreachID
	l.require(listFacebookAdEcommerceProductActivityReportingRequestFieldOutreachID)
}

// SetFields sets the Fields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdEcommerceProductActivityReportingRequest) SetFields(fields []*string) {
	l.Fields = fields
	l.require(listFacebookAdEcommerceProductActivityReportingRequestFieldFields)
}

// SetExcludeFields sets the ExcludeFields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdEcommerceProductActivityReportingRequest) SetExcludeFields(excludeFields []*string) {
	l.ExcludeFields = excludeFields
	l.require(listFacebookAdEcommerceProductActivityReportingRequestFieldExcludeFields)
}

// SetCount sets the Count field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdEcommerceProductActivityReportingRequest) SetCount(count *int) {
	l.Count = count
	l.require(listFacebookAdEcommerceProductActivityReportingRequestFieldCount)
}

// SetOffset sets the Offset field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdEcommerceProductActivityReportingRequest) SetOffset(offset *int) {
	l.Offset = offset
	l.require(listFacebookAdEcommerceProductActivityReportingRequestFieldOffset)
}

// SetSortField sets the SortField field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdEcommerceProductActivityReportingRequest) SetSortField(sortField *ListFacebookAdEcommerceProductActivityReportingRequestSortField) {
	l.SortField = sortField
	l.require(listFacebookAdEcommerceProductActivityReportingRequestFieldSortField)
}

var (
	listFacebookAdsReportingRequestFieldFields        = big.NewInt(1 << 0)
	listFacebookAdsReportingRequestFieldExcludeFields = big.NewInt(1 << 1)
	listFacebookAdsReportingRequestFieldCount         = big.NewInt(1 << 2)
	listFacebookAdsReportingRequestFieldOffset        = big.NewInt(1 << 3)
	listFacebookAdsReportingRequestFieldSortField     = big.NewInt(1 << 4)
	listFacebookAdsReportingRequestFieldSortDir       = big.NewInt(1 << 5)
)

type ListFacebookAdsReportingRequest struct {
	// A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
	Fields []*string `json:"-" url:"fields,omitempty"`
	// A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
	ExcludeFields []*string `json:"-" url:"exclude_fields,omitempty"`
	// The number of records to return. Default value is 10. Maximum value is 1000
	Count *int `json:"-" url:"count,omitempty"`
	// Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
	Offset *int `json:"-" url:"offset,omitempty"`
	// Returns files sorted by the specified field.
	SortField *ListFacebookAdsReportingRequestSortField `json:"-" url:"sort_field,omitempty"`
	// Determines the order direction for sorted results.
	SortDir *ListFacebookAdsReportingRequestSortDir `json:"-" url:"sort_dir,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListFacebookAdsReportingRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetFields sets the Fields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdsReportingRequest) SetFields(fields []*string) {
	l.Fields = fields
	l.require(listFacebookAdsReportingRequestFieldFields)
}

// SetExcludeFields sets the ExcludeFields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdsReportingRequest) SetExcludeFields(excludeFields []*string) {
	l.ExcludeFields = excludeFields
	l.require(listFacebookAdsReportingRequestFieldExcludeFields)
}

// SetCount sets the Count field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdsReportingRequest) SetCount(count *int) {
	l.Count = count
	l.require(listFacebookAdsReportingRequestFieldCount)
}

// SetOffset sets the Offset field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdsReportingRequest) SetOffset(offset *int) {
	l.Offset = offset
	l.require(listFacebookAdsReportingRequestFieldOffset)
}

// SetSortField sets the SortField field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdsReportingRequest) SetSortField(sortField *ListFacebookAdsReportingRequestSortField) {
	l.SortField = sortField
	l.require(listFacebookAdsReportingRequestFieldSortField)
}

// SetSortDir sets the SortDir field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdsReportingRequest) SetSortDir(sortDir *ListFacebookAdsReportingRequestSortDir) {
	l.SortDir = sortDir
	l.require(listFacebookAdsReportingRequestFieldSortDir)
}

var (
	listLandingPagesReportingRequestFieldFields        = big.NewInt(1 << 0)
	listLandingPagesReportingRequestFieldExcludeFields = big.NewInt(1 << 1)
	listLandingPagesReportingRequestFieldCount         = big.NewInt(1 << 2)
	listLandingPagesReportingRequestFieldOffset        = big.NewInt(1 << 3)
)

type ListLandingPagesReportingRequest struct {
	// A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
	Fields []*string `json:"-" url:"fields,omitempty"`
	// A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
	ExcludeFields []*string `json:"-" url:"exclude_fields,omitempty"`
	// The number of records to return. Default value is 10. Maximum value is 1000
	Count *int `json:"-" url:"count,omitempty"`
	// Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
	Offset *int `json:"-" url:"offset,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListLandingPagesReportingRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetFields sets the Fields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListLandingPagesReportingRequest) SetFields(fields []*string) {
	l.Fields = fields
	l.require(listLandingPagesReportingRequestFieldFields)
}

// SetExcludeFields sets the ExcludeFields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListLandingPagesReportingRequest) SetExcludeFields(excludeFields []*string) {
	l.ExcludeFields = excludeFields
	l.require(listLandingPagesReportingRequestFieldExcludeFields)
}

// SetCount sets the Count field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListLandingPagesReportingRequest) SetCount(count *int) {
	l.Count = count
	l.require(listLandingPagesReportingRequestFieldCount)
}

// SetOffset sets the Offset field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListLandingPagesReportingRequest) SetOffset(offset *int) {
	l.Offset = offset
	l.require(listLandingPagesReportingRequestFieldOffset)
}

var (
	listSurveyQuestionAnswersReportingRequestFieldSurveyID                = big.NewInt(1 << 0)
	listSurveyQuestionAnswersReportingRequestFieldQuestionID              = big.NewInt(1 << 1)
	listSurveyQuestionAnswersReportingRequestFieldFields                  = big.NewInt(1 << 2)
	listSurveyQuestionAnswersReportingRequestFieldExcludeFields           = big.NewInt(1 << 3)
	listSurveyQuestionAnswersReportingRequestFieldRespondentFamiliarityIs = big.NewInt(1 << 4)
)

type ListSurveyQuestionAnswersReportingRequest struct {
	// The ID of the survey.
	SurveyID string `json:"-" url:"-"`
	// The ID of the survey question.
	QuestionID string `json:"-" url:"-"`
	// A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
	Fields []*string `json:"-" url:"fields,omitempty"`
	// A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
	ExcludeFields []*string `json:"-" url:"exclude_fields,omitempty"`
	// Filter survey responses by familiarity of the respondents.
	RespondentFamiliarityIs *ListSurveyQuestionAnswersReportingRequestRespondentFamiliarityIs `json:"-" url:"respondent_familiarity_is,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListSurveyQuestionAnswersReportingRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetSurveyID sets the SurveyID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionAnswersReportingRequest) SetSurveyID(surveyID string) {
	l.SurveyID = surveyID
	l.require(listSurveyQuestionAnswersReportingRequestFieldSurveyID)
}

// SetQuestionID sets the QuestionID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionAnswersReportingRequest) SetQuestionID(questionID string) {
	l.QuestionID = questionID
	l.require(listSurveyQuestionAnswersReportingRequestFieldQuestionID)
}

// SetFields sets the Fields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionAnswersReportingRequest) SetFields(fields []*string) {
	l.Fields = fields
	l.require(listSurveyQuestionAnswersReportingRequestFieldFields)
}

// SetExcludeFields sets the ExcludeFields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionAnswersReportingRequest) SetExcludeFields(excludeFields []*string) {
	l.ExcludeFields = excludeFields
	l.require(listSurveyQuestionAnswersReportingRequestFieldExcludeFields)
}

// SetRespondentFamiliarityIs sets the RespondentFamiliarityIs field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionAnswersReportingRequest) SetRespondentFamiliarityIs(respondentFamiliarityIs *ListSurveyQuestionAnswersReportingRequestRespondentFamiliarityIs) {
	l.RespondentFamiliarityIs = respondentFamiliarityIs
	l.require(listSurveyQuestionAnswersReportingRequestFieldRespondentFamiliarityIs)
}

var (
	listSurveyQuestionsReportingRequestFieldSurveyID      = big.NewInt(1 << 0)
	listSurveyQuestionsReportingRequestFieldFields        = big.NewInt(1 << 1)
	listSurveyQuestionsReportingRequestFieldExcludeFields = big.NewInt(1 << 2)
)

type ListSurveyQuestionsReportingRequest struct {
	// The ID of the survey.
	SurveyID string `json:"-" url:"-"`
	// A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
	Fields []*string `json:"-" url:"fields,omitempty"`
	// A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
	ExcludeFields []*string `json:"-" url:"exclude_fields,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListSurveyQuestionsReportingRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetSurveyID sets the SurveyID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionsReportingRequest) SetSurveyID(surveyID string) {
	l.SurveyID = surveyID
	l.require(listSurveyQuestionsReportingRequestFieldSurveyID)
}

// SetFields sets the Fields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionsReportingRequest) SetFields(fields []*string) {
	l.Fields = fields
	l.require(listSurveyQuestionsReportingRequestFieldFields)
}

// SetExcludeFields sets the ExcludeFields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionsReportingRequest) SetExcludeFields(excludeFields []*string) {
	l.ExcludeFields = excludeFields
	l.require(listSurveyQuestionsReportingRequestFieldExcludeFields)
}

var (
	listSurveyResponsesReportingRequestFieldSurveyID                = big.NewInt(1 << 0)
	listSurveyResponsesReportingRequestFieldFields                  = big.NewInt(1 << 1)
	listSurveyResponsesReportingRequestFieldExcludeFields           = big.NewInt(1 << 2)
	listSurveyResponsesReportingRequestFieldAnsweredQuestion        = big.NewInt(1 << 3)
	listSurveyResponsesReportingRequestFieldChoseAnswer             = big.NewInt(1 << 4)
	listSurveyResponsesReportingRequestFieldRespondentFamiliarityIs = big.NewInt(1 << 5)
)

type ListSurveyResponsesReportingRequest struct {
	// The ID of the survey.
	SurveyID string `json:"-" url:"-"`
	// A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
	Fields []*string `json:"-" url:"fields,omitempty"`
	// A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
	ExcludeFields []*string `json:"-" url:"exclude_fields,omitempty"`
	// The ID of the question that was answered.
	AnsweredQuestion *int `json:"-" url:"answered_question,omitempty"`
	// The ID of the option chosen to filter responses on.
	ChoseAnswer *string `json:"-" url:"chose_answer,omitempty"`
	// Filter survey responses by familiarity of the respondents.
	RespondentFamiliarityIs *ListSurveyResponsesReportingRequestRespondentFamiliarityIs `json:"-" url:"respondent_familiarity_is,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListSurveyResponsesReportingRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetSurveyID sets the SurveyID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyResponsesReportingRequest) SetSurveyID(surveyID string) {
	l.SurveyID = surveyID
	l.require(listSurveyResponsesReportingRequestFieldSurveyID)
}

// SetFields sets the Fields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyResponsesReportingRequest) SetFields(fields []*string) {
	l.Fields = fields
	l.require(listSurveyResponsesReportingRequestFieldFields)
}

// SetExcludeFields sets the ExcludeFields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyResponsesReportingRequest) SetExcludeFields(excludeFields []*string) {
	l.ExcludeFields = excludeFields
	l.require(listSurveyResponsesReportingRequestFieldExcludeFields)
}

// SetAnsweredQuestion sets the AnsweredQuestion field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyResponsesReportingRequest) SetAnsweredQuestion(answeredQuestion *int) {
	l.AnsweredQuestion = answeredQuestion
	l.require(listSurveyResponsesReportingRequestFieldAnsweredQuestion)
}

// SetChoseAnswer sets the ChoseAnswer field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyResponsesReportingRequest) SetChoseAnswer(choseAnswer *string) {
	l.ChoseAnswer = choseAnswer
	l.require(listSurveyResponsesReportingRequestFieldChoseAnswer)
}

// SetRespondentFamiliarityIs sets the RespondentFamiliarityIs field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyResponsesReportingRequest) SetRespondentFamiliarityIs(respondentFamiliarityIs *ListSurveyResponsesReportingRequestRespondentFamiliarityIs) {
	l.RespondentFamiliarityIs = respondentFamiliarityIs
	l.require(listSurveyResponsesReportingRequestFieldRespondentFamiliarityIs)
}

var (
	listSurveysReportingRequestFieldFields        = big.NewInt(1 << 0)
	listSurveysReportingRequestFieldExcludeFields = big.NewInt(1 << 1)
	listSurveysReportingRequestFieldCount         = big.NewInt(1 << 2)
	listSurveysReportingRequestFieldOffset        = big.NewInt(1 << 3)
)

type ListSurveysReportingRequest struct {
	// A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
	Fields []*string `json:"-" url:"fields,omitempty"`
	// A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
	ExcludeFields []*string `json:"-" url:"exclude_fields,omitempty"`
	// The number of records to return. Default value is 10. Maximum value is 1000
	Count *int `json:"-" url:"count,omitempty"`
	// Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
	Offset *int `json:"-" url:"offset,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListSurveysReportingRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetFields sets the Fields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveysReportingRequest) SetFields(fields []*string) {
	l.Fields = fields
	l.require(listSurveysReportingRequestFieldFields)
}

// SetExcludeFields sets the ExcludeFields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveysReportingRequest) SetExcludeFields(excludeFields []*string) {
	l.ExcludeFields = excludeFields
	l.require(listSurveysReportingRequestFieldExcludeFields)
}

// SetCount sets the Count field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveysReportingRequest) SetCount(count *int) {
	l.Count = count
	l.require(listSurveysReportingRequestFieldCount)
}

// SetOffset sets the Offset field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveysReportingRequest) SetOffset(offset *int) {
	l.Offset = offset
	l.require(listSurveysReportingRequestFieldOffset)
}

// A summary of an individual landing page's settings and content.
var (
	landingPageReportFieldLinks          = big.NewInt(1 << 0)
	landingPageReportFieldClicks         = big.NewInt(1 << 1)
	landingPageReportFieldConversionRate = big.NewInt(1 << 2)
	landingPageReportFieldEcommerce      = big.NewInt(1 << 3)
	landingPageReportFieldID             = big.NewInt(1 << 4)
	landingPageReportFieldListID         = big.NewInt(1 << 5)
	landingPageReportFieldListName       = big.NewInt(1 << 6)
	landingPageReportFieldName           = big.NewInt(1 << 7)
	landingPageReportFieldPublishedAt    = big.NewInt(1 << 8)
	landingPageReportFieldSignupTags     = big.NewInt(1 << 9)
	landingPageReportFieldStatus         = big.NewInt(1 << 10)
	landingPageReportFieldSubscribes     = big.NewInt(1 << 11)
	landingPageReportFieldTimeseries     = big.NewInt(1 << 12)
	landingPageReportFieldTitle          = big.NewInt(1 << 13)
	landingPageReportFieldUniqueVisits   = big.NewInt(1 << 14)
	landingPageReportFieldUnpublishedAt  = big.NewInt(1 << 15)
	landingPageReportFieldURL            = big.NewInt(1 << 16)
	landingPageReportFieldVisits         = big.NewInt(1 << 17)
	landingPageReportFieldWebID          = big.NewInt(1 << 18)
)

type LandingPageReport struct {
	// A list of link types and descriptions for the API schema documents.
	Links []*LandingPageReportLinksItem `json:"_links,omitempty" url:"_links,omitempty"`
	// The number of clicks to this landing pages.
	Clicks *int `json:"clicks,omitempty" url:"clicks,omitempty"`
	// The percentage of people who visited your landing page and were added to your list.
	ConversionRate *float64                    `json:"conversion_rate,omitempty" url:"conversion_rate,omitempty"`
	Ecommerce      *LandingPageReportEcommerce `json:"ecommerce,omitempty" url:"ecommerce,omitempty"`
	// A string that uniquely identifies this landing page.
	ID *string `json:"id,omitempty" url:"id,omitempty"`
	// The list id connected to this landing page.
	ListID *string `json:"list_id,omitempty" url:"list_id,omitempty"`
	// List Name
	ListName *string `json:"list_name,omitempty" url:"list_name,omitempty"`
	// The name of this landing page the user will see.
	Name *string `json:"name,omitempty" url:"name,omitempty"`
	// The time this landing page was published.
	PublishedAt *time.Time `json:"published_at,omitempty" url:"published_at,omitempty"`
	// A list of tags associated to the landing page.
	SignupTags []*LandingPageReportSignupTagsItem `json:"signup_tags,omitempty" url:"signup_tags,omitempty"`
	// The status of the landing page.
	Status *string `json:"status,omitempty" url:"status,omitempty"`
	// The number of subscribes to this landing pages.
	Subscribes *int                         `json:"subscribes,omitempty" url:"subscribes,omitempty"`
	Timeseries *LandingPageReportTimeseries `json:"timeseries,omitempty" url:"timeseries,omitempty"`
	// The name of the landing page the user's customers will see.
	Title *string `json:"title,omitempty" url:"title,omitempty"`
	// The number of unique visits to this landing pages.
	UniqueVisits *int `json:"unique_visits,omitempty" url:"unique_visits,omitempty"`
	// The time this landing page was unpublished.
	UnpublishedAt *time.Time `json:"unpublished_at,omitempty" url:"unpublished_at,omitempty"`
	// The landing page url.
	URL *string `json:"url,omitempty" url:"url,omitempty"`
	// The number of visits to this landing pages.
	Visits *int `json:"visits,omitempty" url:"visits,omitempty"`
	// The ID used in the Mailchimp web application.
	WebID *int `json:"web_id,omitempty" url:"web_id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *LandingPageReport) GetLinks() []*LandingPageReportLinksItem {
	if l == nil {
		return nil
	}
	return l.Links
}

func (l *LandingPageReport) GetClicks() *int {
	if l == nil {
		return nil
	}
	return l.Clicks
}

func (l *LandingPageReport) GetConversionRate() *float64 {
	if l == nil {
		return nil
	}
	return l.ConversionRate
}

func (l *LandingPageReport) GetEcommerce() *LandingPageReportEcommerce {
	if l == nil {
		return nil
	}
	return l.Ecommerce
}

func (l *LandingPageReport) GetID() *string {
	if l == nil {
		return nil
	}
	return l.ID
}

func (l *LandingPageReport) GetListID() *string {
	if l == nil {
		return nil
	}
	return l.ListID
}

func (l *LandingPageReport) GetListName() *string {
	if l == nil {
		return nil
	}
	return l.ListName
}

func (l *LandingPageReport) GetName() *string {
	if l == nil {
		return nil
	}
	return l.Name
}

func (l *LandingPageReport) GetPublishedAt() *time.Time {
	if l == nil {
		return nil
	}
	return l.PublishedAt
}

func (l *LandingPageReport) GetSignupTags() []*LandingPageReportSignupTagsItem {
	if l == nil {
		return nil
	}
	return l.SignupTags
}

func (l *LandingPageReport) GetStatus() *string {
	if l == nil {
		return nil
	}
	return l.Status
}

func (l *LandingPageReport) GetSubscribes() *int {
	if l == nil {
		return nil
	}
	return l.Subscribes
}

func (l *LandingPageReport) GetTimeseries() *LandingPageReportTimeseries {
	if l == nil {
		return nil
	}
	return l.Timeseries
}

func (l *LandingPageReport) GetTitle() *string {
	if l == nil {
		return nil
	}
	return l.Title
}

func (l *LandingPageReport) GetUniqueVisits() *int {
	if l == nil {
		return nil
	}
	return l.UniqueVisits
}

func (l *LandingPageReport) GetUnpublishedAt() *time.Time {
	if l == nil {
		return nil
	}
	return l.UnpublishedAt
}

func (l *LandingPageReport) GetURL() *string {
	if l == nil {
		return nil
	}
	return l.URL
}

func (l *LandingPageReport) GetVisits() *int {
	if l == nil {
		return nil
	}
	return l.Visits
}

func (l *LandingPageReport) GetWebID() *int {
	if l == nil {
		return nil
	}
	return l.WebID
}

func (l *LandingPageReport) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *LandingPageReport) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetLinks sets the Links field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReport) SetLinks(links []*LandingPageReportLinksItem) {
	l.Links = links
	l.require(landingPageReportFieldLinks)
}

// SetClicks sets the Clicks field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReport) SetClicks(clicks *int) {
	l.Clicks = clicks
	l.require(landingPageReportFieldClicks)
}

// SetConversionRate sets the ConversionRate field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReport) SetConversionRate(conversionRate *float64) {
	l.ConversionRate = conversionRate
	l.require(landingPageReportFieldConversionRate)
}

// SetEcommerce sets the Ecommerce field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReport) SetEcommerce(ecommerce *LandingPageReportEcommerce) {
	l.Ecommerce = ecommerce
	l.require(landingPageReportFieldEcommerce)
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReport) SetID(id *string) {
	l.ID = id
	l.require(landingPageReportFieldID)
}

// SetListID sets the ListID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReport) SetListID(listID *string) {
	l.ListID = listID
	l.require(landingPageReportFieldListID)
}

// SetListName sets the ListName field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReport) SetListName(listName *string) {
	l.ListName = listName
	l.require(landingPageReportFieldListName)
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReport) SetName(name *string) {
	l.Name = name
	l.require(landingPageReportFieldName)
}

// SetPublishedAt sets the PublishedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReport) SetPublishedAt(publishedAt *time.Time) {
	l.PublishedAt = publishedAt
	l.require(landingPageReportFieldPublishedAt)
}

// SetSignupTags sets the SignupTags field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReport) SetSignupTags(signupTags []*LandingPageReportSignupTagsItem) {
	l.SignupTags = signupTags
	l.require(landingPageReportFieldSignupTags)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReport) SetStatus(status *string) {
	l.Status = status
	l.require(landingPageReportFieldStatus)
}

// SetSubscribes sets the Subscribes field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReport) SetSubscribes(subscribes *int) {
	l.Subscribes = subscribes
	l.require(landingPageReportFieldSubscribes)
}

// SetTimeseries sets the Timeseries field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReport) SetTimeseries(timeseries *LandingPageReportTimeseries) {
	l.Timeseries = timeseries
	l.require(landingPageReportFieldTimeseries)
}

// SetTitle sets the Title field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReport) SetTitle(title *string) {
	l.Title = title
	l.require(landingPageReportFieldTitle)
}

// SetUniqueVisits sets the UniqueVisits field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReport) SetUniqueVisits(uniqueVisits *int) {
	l.UniqueVisits = uniqueVisits
	l.require(landingPageReportFieldUniqueVisits)
}

// SetUnpublishedAt sets the UnpublishedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReport) SetUnpublishedAt(unpublishedAt *time.Time) {
	l.UnpublishedAt = unpublishedAt
	l.require(landingPageReportFieldUnpublishedAt)
}

// SetURL sets the URL field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReport) SetURL(url *string) {
	l.URL = url
	l.require(landingPageReportFieldURL)
}

// SetVisits sets the Visits field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReport) SetVisits(visits *int) {
	l.Visits = visits
	l.require(landingPageReportFieldVisits)
}

// SetWebID sets the WebID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReport) SetWebID(webID *int) {
	l.WebID = webID
	l.require(landingPageReportFieldWebID)
}

func (l *LandingPageReport) UnmarshalJSON(data []byte) error {
	type embed LandingPageReport
	var unmarshaler = struct {
		embed
		PublishedAt   *internal.DateTime `json:"published_at,omitempty"`
		UnpublishedAt *internal.DateTime `json:"unpublished_at,omitempty"`
	}{
		embed: embed(*l),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*l = LandingPageReport(unmarshaler.embed)
	l.PublishedAt = unmarshaler.PublishedAt.TimePtr()
	l.UnpublishedAt = unmarshaler.UnpublishedAt.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *LandingPageReport) MarshalJSON() ([]byte, error) {
	type embed LandingPageReport
	var marshaler = struct {
		embed
		PublishedAt   *internal.DateTime `json:"published_at,omitempty"`
		UnpublishedAt *internal.DateTime `json:"unpublished_at,omitempty"`
	}{
		embed:         embed(*l),
		PublishedAt:   internal.NewOptionalDateTime(l.PublishedAt),
		UnpublishedAt: internal.NewOptionalDateTime(l.UnpublishedAt),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *LandingPageReport) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

var (
	landingPageReportEcommerceFieldAverageOrderRevenue = big.NewInt(1 << 0)
	landingPageReportEcommerceFieldCurrencyCode        = big.NewInt(1 << 1)
	landingPageReportEcommerceFieldTotalOrders         = big.NewInt(1 << 2)
	landingPageReportEcommerceFieldTotalRevenue        = big.NewInt(1 << 3)
)

type LandingPageReportEcommerce struct {
	// The average order revenue of this landing page.
	AverageOrderRevenue *float64 `json:"average_order_revenue,omitempty" url:"average_order_revenue,omitempty"`
	// The user's currency code.
	CurrencyCode *string `json:"currency_code,omitempty" url:"currency_code,omitempty"`
	// The total number of orders associated with this landing page.
	TotalOrders *int `json:"total_orders,omitempty" url:"total_orders,omitempty"`
	// The total revenue of this landing page.
	TotalRevenue *float64 `json:"total_revenue,omitempty" url:"total_revenue,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *LandingPageReportEcommerce) GetAverageOrderRevenue() *float64 {
	if l == nil {
		return nil
	}
	return l.AverageOrderRevenue
}

func (l *LandingPageReportEcommerce) GetCurrencyCode() *string {
	if l == nil {
		return nil
	}
	return l.CurrencyCode
}

func (l *LandingPageReportEcommerce) GetTotalOrders() *int {
	if l == nil {
		return nil
	}
	return l.TotalOrders
}

func (l *LandingPageReportEcommerce) GetTotalRevenue() *float64 {
	if l == nil {
		return nil
	}
	return l.TotalRevenue
}

func (l *LandingPageReportEcommerce) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *LandingPageReportEcommerce) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetAverageOrderRevenue sets the AverageOrderRevenue field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportEcommerce) SetAverageOrderRevenue(averageOrderRevenue *float64) {
	l.AverageOrderRevenue = averageOrderRevenue
	l.require(landingPageReportEcommerceFieldAverageOrderRevenue)
}

// SetCurrencyCode sets the CurrencyCode field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportEcommerce) SetCurrencyCode(currencyCode *string) {
	l.CurrencyCode = currencyCode
	l.require(landingPageReportEcommerceFieldCurrencyCode)
}

// SetTotalOrders sets the TotalOrders field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportEcommerce) SetTotalOrders(totalOrders *int) {
	l.TotalOrders = totalOrders
	l.require(landingPageReportEcommerceFieldTotalOrders)
}

// SetTotalRevenue sets the TotalRevenue field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportEcommerce) SetTotalRevenue(totalRevenue *float64) {
	l.TotalRevenue = totalRevenue
	l.require(landingPageReportEcommerceFieldTotalRevenue)
}

func (l *LandingPageReportEcommerce) UnmarshalJSON(data []byte) error {
	type unmarshaler LandingPageReportEcommerce
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LandingPageReportEcommerce(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *LandingPageReportEcommerce) MarshalJSON() ([]byte, error) {
	type embed LandingPageReportEcommerce
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *LandingPageReportEcommerce) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// This object represents a link from the resource where it is found to another resource or action that may be performed.
var (
	landingPageReportLinksItemFieldHref         = big.NewInt(1 << 0)
	landingPageReportLinksItemFieldMethod       = big.NewInt(1 << 1)
	landingPageReportLinksItemFieldRel          = big.NewInt(1 << 2)
	landingPageReportLinksItemFieldSchema       = big.NewInt(1 << 3)
	landingPageReportLinksItemFieldTargetSchema = big.NewInt(1 << 4)
)

type LandingPageReportLinksItem struct {
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href *string `json:"href,omitempty" url:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method *LandingPageReportLinksItemMethod `json:"method,omitempty" url:"method,omitempty"`
	// As with an HTML 'rel' attribute, this describes the type of link.
	Rel *string `json:"rel,omitempty" url:"rel,omitempty"`
	// For HTTP methods that can receive bodies (POST and PUT), this is a URL representing the schema that the body should conform to.
	Schema *string `json:"schema,omitempty" url:"schema,omitempty"`
	// For GETs, this is a URL representing the schema that the response should conform to.
	TargetSchema *string `json:"targetSchema,omitempty" url:"targetSchema,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *LandingPageReportLinksItem) GetHref() *string {
	if l == nil {
		return nil
	}
	return l.Href
}

func (l *LandingPageReportLinksItem) GetMethod() *LandingPageReportLinksItemMethod {
	if l == nil {
		return nil
	}
	return l.Method
}

func (l *LandingPageReportLinksItem) GetRel() *string {
	if l == nil {
		return nil
	}
	return l.Rel
}

func (l *LandingPageReportLinksItem) GetSchema() *string {
	if l == nil {
		return nil
	}
	return l.Schema
}

func (l *LandingPageReportLinksItem) GetTargetSchema() *string {
	if l == nil {
		return nil
	}
	return l.TargetSchema
}

func (l *LandingPageReportLinksItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *LandingPageReportLinksItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetHref sets the Href field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportLinksItem) SetHref(href *string) {
	l.Href = href
	l.require(landingPageReportLinksItemFieldHref)
}

// SetMethod sets the Method field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportLinksItem) SetMethod(method *LandingPageReportLinksItemMethod) {
	l.Method = method
	l.require(landingPageReportLinksItemFieldMethod)
}

// SetRel sets the Rel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportLinksItem) SetRel(rel *string) {
	l.Rel = rel
	l.require(landingPageReportLinksItemFieldRel)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportLinksItem) SetSchema(schema *string) {
	l.Schema = schema
	l.require(landingPageReportLinksItemFieldSchema)
}

// SetTargetSchema sets the TargetSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportLinksItem) SetTargetSchema(targetSchema *string) {
	l.TargetSchema = targetSchema
	l.require(landingPageReportLinksItemFieldTargetSchema)
}

func (l *LandingPageReportLinksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler LandingPageReportLinksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LandingPageReportLinksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *LandingPageReportLinksItem) MarshalJSON() ([]byte, error) {
	type embed LandingPageReportLinksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *LandingPageReportLinksItem) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// The HTTP method that should be used when accessing the URL defined in 'href'.
type LandingPageReportLinksItemMethod string

const (
	LandingPageReportLinksItemMethodGet     LandingPageReportLinksItemMethod = "GET"
	LandingPageReportLinksItemMethodPost    LandingPageReportLinksItemMethod = "POST"
	LandingPageReportLinksItemMethodPut     LandingPageReportLinksItemMethod = "PUT"
	LandingPageReportLinksItemMethodPatch   LandingPageReportLinksItemMethod = "PATCH"
	LandingPageReportLinksItemMethodDelete  LandingPageReportLinksItemMethod = "DELETE"
	LandingPageReportLinksItemMethodOptions LandingPageReportLinksItemMethod = "OPTIONS"
	LandingPageReportLinksItemMethodHead    LandingPageReportLinksItemMethod = "HEAD"
)

func NewLandingPageReportLinksItemMethodFromString(s string) (LandingPageReportLinksItemMethod, error) {
	switch s {
	case "GET":
		return LandingPageReportLinksItemMethodGet, nil
	case "POST":
		return LandingPageReportLinksItemMethodPost, nil
	case "PUT":
		return LandingPageReportLinksItemMethodPut, nil
	case "PATCH":
		return LandingPageReportLinksItemMethodPatch, nil
	case "DELETE":
		return LandingPageReportLinksItemMethodDelete, nil
	case "OPTIONS":
		return LandingPageReportLinksItemMethodOptions, nil
	case "HEAD":
		return LandingPageReportLinksItemMethodHead, nil
	}
	var t LandingPageReportLinksItemMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l LandingPageReportLinksItemMethod) Ptr() *LandingPageReportLinksItemMethod {
	return &l
}

var (
	landingPageReportSignupTagsItemFieldTagID   = big.NewInt(1 << 0)
	landingPageReportSignupTagsItemFieldTagName = big.NewInt(1 << 1)
)

type LandingPageReportSignupTagsItem struct {
	// The unique id for the tag.
	TagID *int `json:"tag_id,omitempty" url:"tag_id,omitempty"`
	// The name of the tag.
	TagName *string `json:"tag_name,omitempty" url:"tag_name,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *LandingPageReportSignupTagsItem) GetTagID() *int {
	if l == nil {
		return nil
	}
	return l.TagID
}

func (l *LandingPageReportSignupTagsItem) GetTagName() *string {
	if l == nil {
		return nil
	}
	return l.TagName
}

func (l *LandingPageReportSignupTagsItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *LandingPageReportSignupTagsItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetTagID sets the TagID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportSignupTagsItem) SetTagID(tagID *int) {
	l.TagID = tagID
	l.require(landingPageReportSignupTagsItemFieldTagID)
}

// SetTagName sets the TagName field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportSignupTagsItem) SetTagName(tagName *string) {
	l.TagName = tagName
	l.require(landingPageReportSignupTagsItemFieldTagName)
}

func (l *LandingPageReportSignupTagsItem) UnmarshalJSON(data []byte) error {
	type unmarshaler LandingPageReportSignupTagsItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LandingPageReportSignupTagsItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *LandingPageReportSignupTagsItem) MarshalJSON() ([]byte, error) {
	type embed LandingPageReportSignupTagsItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *LandingPageReportSignupTagsItem) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

var (
	landingPageReportTimeseriesFieldDailyStats  = big.NewInt(1 << 0)
	landingPageReportTimeseriesFieldWeeklyStats = big.NewInt(1 << 1)
)

type LandingPageReportTimeseries struct {
	// The clicks and visits data from the last seven days.
	DailyStats *LandingPageReportTimeseriesDailyStats `json:"daily_stats,omitempty" url:"daily_stats,omitempty"`
	// The clicks and visits data from the last five weeks.
	WeeklyStats *LandingPageReportTimeseriesWeeklyStats `json:"weekly_stats,omitempty" url:"weekly_stats,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *LandingPageReportTimeseries) GetDailyStats() *LandingPageReportTimeseriesDailyStats {
	if l == nil {
		return nil
	}
	return l.DailyStats
}

func (l *LandingPageReportTimeseries) GetWeeklyStats() *LandingPageReportTimeseriesWeeklyStats {
	if l == nil {
		return nil
	}
	return l.WeeklyStats
}

func (l *LandingPageReportTimeseries) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *LandingPageReportTimeseries) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetDailyStats sets the DailyStats field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportTimeseries) SetDailyStats(dailyStats *LandingPageReportTimeseriesDailyStats) {
	l.DailyStats = dailyStats
	l.require(landingPageReportTimeseriesFieldDailyStats)
}

// SetWeeklyStats sets the WeeklyStats field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportTimeseries) SetWeeklyStats(weeklyStats *LandingPageReportTimeseriesWeeklyStats) {
	l.WeeklyStats = weeklyStats
	l.require(landingPageReportTimeseriesFieldWeeklyStats)
}

func (l *LandingPageReportTimeseries) UnmarshalJSON(data []byte) error {
	type unmarshaler LandingPageReportTimeseries
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LandingPageReportTimeseries(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *LandingPageReportTimeseries) MarshalJSON() ([]byte, error) {
	type embed LandingPageReportTimeseries
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *LandingPageReportTimeseries) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// The clicks and visits data from the last seven days.
var (
	landingPageReportTimeseriesDailyStatsFieldClicks       = big.NewInt(1 << 0)
	landingPageReportTimeseriesDailyStatsFieldUniqueVisits = big.NewInt(1 << 1)
	landingPageReportTimeseriesDailyStatsFieldVisits       = big.NewInt(1 << 2)
)

type LandingPageReportTimeseriesDailyStats struct {
	Clicks       []*LandingPageReportTimeseriesDailyStatsClicksItem       `json:"clicks,omitempty" url:"clicks,omitempty"`
	UniqueVisits []*LandingPageReportTimeseriesDailyStatsUniqueVisitsItem `json:"unique_visits,omitempty" url:"unique_visits,omitempty"`
	Visits       []*LandingPageReportTimeseriesDailyStatsVisitsItem       `json:"visits,omitempty" url:"visits,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *LandingPageReportTimeseriesDailyStats) GetClicks() []*LandingPageReportTimeseriesDailyStatsClicksItem {
	if l == nil {
		return nil
	}
	return l.Clicks
}

func (l *LandingPageReportTimeseriesDailyStats) GetUniqueVisits() []*LandingPageReportTimeseriesDailyStatsUniqueVisitsItem {
	if l == nil {
		return nil
	}
	return l.UniqueVisits
}

func (l *LandingPageReportTimeseriesDailyStats) GetVisits() []*LandingPageReportTimeseriesDailyStatsVisitsItem {
	if l == nil {
		return nil
	}
	return l.Visits
}

func (l *LandingPageReportTimeseriesDailyStats) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *LandingPageReportTimeseriesDailyStats) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetClicks sets the Clicks field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportTimeseriesDailyStats) SetClicks(clicks []*LandingPageReportTimeseriesDailyStatsClicksItem) {
	l.Clicks = clicks
	l.require(landingPageReportTimeseriesDailyStatsFieldClicks)
}

// SetUniqueVisits sets the UniqueVisits field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportTimeseriesDailyStats) SetUniqueVisits(uniqueVisits []*LandingPageReportTimeseriesDailyStatsUniqueVisitsItem) {
	l.UniqueVisits = uniqueVisits
	l.require(landingPageReportTimeseriesDailyStatsFieldUniqueVisits)
}

// SetVisits sets the Visits field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportTimeseriesDailyStats) SetVisits(visits []*LandingPageReportTimeseriesDailyStatsVisitsItem) {
	l.Visits = visits
	l.require(landingPageReportTimeseriesDailyStatsFieldVisits)
}

func (l *LandingPageReportTimeseriesDailyStats) UnmarshalJSON(data []byte) error {
	type unmarshaler LandingPageReportTimeseriesDailyStats
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LandingPageReportTimeseriesDailyStats(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *LandingPageReportTimeseriesDailyStats) MarshalJSON() ([]byte, error) {
	type embed LandingPageReportTimeseriesDailyStats
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *LandingPageReportTimeseriesDailyStats) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

var (
	landingPageReportTimeseriesDailyStatsClicksItemFieldDate = big.NewInt(1 << 0)
	landingPageReportTimeseriesDailyStatsClicksItemFieldVal  = big.NewInt(1 << 1)
)

type LandingPageReportTimeseriesDailyStatsClicksItem struct {
	Date *string `json:"date,omitempty" url:"date,omitempty"`
	Val  *int    `json:"val,omitempty" url:"val,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *LandingPageReportTimeseriesDailyStatsClicksItem) GetDate() *string {
	if l == nil {
		return nil
	}
	return l.Date
}

func (l *LandingPageReportTimeseriesDailyStatsClicksItem) GetVal() *int {
	if l == nil {
		return nil
	}
	return l.Val
}

func (l *LandingPageReportTimeseriesDailyStatsClicksItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *LandingPageReportTimeseriesDailyStatsClicksItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetDate sets the Date field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportTimeseriesDailyStatsClicksItem) SetDate(date *string) {
	l.Date = date
	l.require(landingPageReportTimeseriesDailyStatsClicksItemFieldDate)
}

// SetVal sets the Val field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportTimeseriesDailyStatsClicksItem) SetVal(val *int) {
	l.Val = val
	l.require(landingPageReportTimeseriesDailyStatsClicksItemFieldVal)
}

func (l *LandingPageReportTimeseriesDailyStatsClicksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler LandingPageReportTimeseriesDailyStatsClicksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LandingPageReportTimeseriesDailyStatsClicksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *LandingPageReportTimeseriesDailyStatsClicksItem) MarshalJSON() ([]byte, error) {
	type embed LandingPageReportTimeseriesDailyStatsClicksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *LandingPageReportTimeseriesDailyStatsClicksItem) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

var (
	landingPageReportTimeseriesDailyStatsUniqueVisitsItemFieldDate = big.NewInt(1 << 0)
	landingPageReportTimeseriesDailyStatsUniqueVisitsItemFieldVal  = big.NewInt(1 << 1)
)

type LandingPageReportTimeseriesDailyStatsUniqueVisitsItem struct {
	Date *string `json:"date,omitempty" url:"date,omitempty"`
	Val  *int    `json:"val,omitempty" url:"val,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *LandingPageReportTimeseriesDailyStatsUniqueVisitsItem) GetDate() *string {
	if l == nil {
		return nil
	}
	return l.Date
}

func (l *LandingPageReportTimeseriesDailyStatsUniqueVisitsItem) GetVal() *int {
	if l == nil {
		return nil
	}
	return l.Val
}

func (l *LandingPageReportTimeseriesDailyStatsUniqueVisitsItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *LandingPageReportTimeseriesDailyStatsUniqueVisitsItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetDate sets the Date field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportTimeseriesDailyStatsUniqueVisitsItem) SetDate(date *string) {
	l.Date = date
	l.require(landingPageReportTimeseriesDailyStatsUniqueVisitsItemFieldDate)
}

// SetVal sets the Val field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportTimeseriesDailyStatsUniqueVisitsItem) SetVal(val *int) {
	l.Val = val
	l.require(landingPageReportTimeseriesDailyStatsUniqueVisitsItemFieldVal)
}

func (l *LandingPageReportTimeseriesDailyStatsUniqueVisitsItem) UnmarshalJSON(data []byte) error {
	type unmarshaler LandingPageReportTimeseriesDailyStatsUniqueVisitsItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LandingPageReportTimeseriesDailyStatsUniqueVisitsItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *LandingPageReportTimeseriesDailyStatsUniqueVisitsItem) MarshalJSON() ([]byte, error) {
	type embed LandingPageReportTimeseriesDailyStatsUniqueVisitsItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *LandingPageReportTimeseriesDailyStatsUniqueVisitsItem) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

var (
	landingPageReportTimeseriesDailyStatsVisitsItemFieldDate = big.NewInt(1 << 0)
	landingPageReportTimeseriesDailyStatsVisitsItemFieldVal  = big.NewInt(1 << 1)
)

type LandingPageReportTimeseriesDailyStatsVisitsItem struct {
	Date *string `json:"date,omitempty" url:"date,omitempty"`
	Val  *int    `json:"val,omitempty" url:"val,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *LandingPageReportTimeseriesDailyStatsVisitsItem) GetDate() *string {
	if l == nil {
		return nil
	}
	return l.Date
}

func (l *LandingPageReportTimeseriesDailyStatsVisitsItem) GetVal() *int {
	if l == nil {
		return nil
	}
	return l.Val
}

func (l *LandingPageReportTimeseriesDailyStatsVisitsItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *LandingPageReportTimeseriesDailyStatsVisitsItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetDate sets the Date field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportTimeseriesDailyStatsVisitsItem) SetDate(date *string) {
	l.Date = date
	l.require(landingPageReportTimeseriesDailyStatsVisitsItemFieldDate)
}

// SetVal sets the Val field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportTimeseriesDailyStatsVisitsItem) SetVal(val *int) {
	l.Val = val
	l.require(landingPageReportTimeseriesDailyStatsVisitsItemFieldVal)
}

func (l *LandingPageReportTimeseriesDailyStatsVisitsItem) UnmarshalJSON(data []byte) error {
	type unmarshaler LandingPageReportTimeseriesDailyStatsVisitsItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LandingPageReportTimeseriesDailyStatsVisitsItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *LandingPageReportTimeseriesDailyStatsVisitsItem) MarshalJSON() ([]byte, error) {
	type embed LandingPageReportTimeseriesDailyStatsVisitsItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *LandingPageReportTimeseriesDailyStatsVisitsItem) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// The clicks and visits data from the last five weeks.
var (
	landingPageReportTimeseriesWeeklyStatsFieldClicks       = big.NewInt(1 << 0)
	landingPageReportTimeseriesWeeklyStatsFieldUniqueVisits = big.NewInt(1 << 1)
	landingPageReportTimeseriesWeeklyStatsFieldVisits       = big.NewInt(1 << 2)
)

type LandingPageReportTimeseriesWeeklyStats struct {
	// The total number of clicks in a week.
	Clicks       []*LandingPageReportTimeseriesWeeklyStatsClicksItem       `json:"clicks,omitempty" url:"clicks,omitempty"`
	UniqueVisits []*LandingPageReportTimeseriesWeeklyStatsUniqueVisitsItem `json:"unique_visits,omitempty" url:"unique_visits,omitempty"`
	// The total number of visits in a week.
	Visits []*LandingPageReportTimeseriesWeeklyStatsVisitsItem `json:"visits,omitempty" url:"visits,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *LandingPageReportTimeseriesWeeklyStats) GetClicks() []*LandingPageReportTimeseriesWeeklyStatsClicksItem {
	if l == nil {
		return nil
	}
	return l.Clicks
}

func (l *LandingPageReportTimeseriesWeeklyStats) GetUniqueVisits() []*LandingPageReportTimeseriesWeeklyStatsUniqueVisitsItem {
	if l == nil {
		return nil
	}
	return l.UniqueVisits
}

func (l *LandingPageReportTimeseriesWeeklyStats) GetVisits() []*LandingPageReportTimeseriesWeeklyStatsVisitsItem {
	if l == nil {
		return nil
	}
	return l.Visits
}

func (l *LandingPageReportTimeseriesWeeklyStats) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *LandingPageReportTimeseriesWeeklyStats) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetClicks sets the Clicks field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportTimeseriesWeeklyStats) SetClicks(clicks []*LandingPageReportTimeseriesWeeklyStatsClicksItem) {
	l.Clicks = clicks
	l.require(landingPageReportTimeseriesWeeklyStatsFieldClicks)
}

// SetUniqueVisits sets the UniqueVisits field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportTimeseriesWeeklyStats) SetUniqueVisits(uniqueVisits []*LandingPageReportTimeseriesWeeklyStatsUniqueVisitsItem) {
	l.UniqueVisits = uniqueVisits
	l.require(landingPageReportTimeseriesWeeklyStatsFieldUniqueVisits)
}

// SetVisits sets the Visits field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportTimeseriesWeeklyStats) SetVisits(visits []*LandingPageReportTimeseriesWeeklyStatsVisitsItem) {
	l.Visits = visits
	l.require(landingPageReportTimeseriesWeeklyStatsFieldVisits)
}

func (l *LandingPageReportTimeseriesWeeklyStats) UnmarshalJSON(data []byte) error {
	type unmarshaler LandingPageReportTimeseriesWeeklyStats
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LandingPageReportTimeseriesWeeklyStats(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *LandingPageReportTimeseriesWeeklyStats) MarshalJSON() ([]byte, error) {
	type embed LandingPageReportTimeseriesWeeklyStats
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *LandingPageReportTimeseriesWeeklyStats) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

var (
	landingPageReportTimeseriesWeeklyStatsClicksItemFieldDate = big.NewInt(1 << 0)
	landingPageReportTimeseriesWeeklyStatsClicksItemFieldVal  = big.NewInt(1 << 1)
)

type LandingPageReportTimeseriesWeeklyStatsClicksItem struct {
	Date *string `json:"date,omitempty" url:"date,omitempty"`
	Val  *int    `json:"val,omitempty" url:"val,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *LandingPageReportTimeseriesWeeklyStatsClicksItem) GetDate() *string {
	if l == nil {
		return nil
	}
	return l.Date
}

func (l *LandingPageReportTimeseriesWeeklyStatsClicksItem) GetVal() *int {
	if l == nil {
		return nil
	}
	return l.Val
}

func (l *LandingPageReportTimeseriesWeeklyStatsClicksItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *LandingPageReportTimeseriesWeeklyStatsClicksItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetDate sets the Date field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportTimeseriesWeeklyStatsClicksItem) SetDate(date *string) {
	l.Date = date
	l.require(landingPageReportTimeseriesWeeklyStatsClicksItemFieldDate)
}

// SetVal sets the Val field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportTimeseriesWeeklyStatsClicksItem) SetVal(val *int) {
	l.Val = val
	l.require(landingPageReportTimeseriesWeeklyStatsClicksItemFieldVal)
}

func (l *LandingPageReportTimeseriesWeeklyStatsClicksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler LandingPageReportTimeseriesWeeklyStatsClicksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LandingPageReportTimeseriesWeeklyStatsClicksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *LandingPageReportTimeseriesWeeklyStatsClicksItem) MarshalJSON() ([]byte, error) {
	type embed LandingPageReportTimeseriesWeeklyStatsClicksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *LandingPageReportTimeseriesWeeklyStatsClicksItem) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

var (
	landingPageReportTimeseriesWeeklyStatsUniqueVisitsItemFieldDate = big.NewInt(1 << 0)
	landingPageReportTimeseriesWeeklyStatsUniqueVisitsItemFieldVal  = big.NewInt(1 << 1)
)

type LandingPageReportTimeseriesWeeklyStatsUniqueVisitsItem struct {
	Date *string `json:"date,omitempty" url:"date,omitempty"`
	Val  *int    `json:"val,omitempty" url:"val,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *LandingPageReportTimeseriesWeeklyStatsUniqueVisitsItem) GetDate() *string {
	if l == nil {
		return nil
	}
	return l.Date
}

func (l *LandingPageReportTimeseriesWeeklyStatsUniqueVisitsItem) GetVal() *int {
	if l == nil {
		return nil
	}
	return l.Val
}

func (l *LandingPageReportTimeseriesWeeklyStatsUniqueVisitsItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *LandingPageReportTimeseriesWeeklyStatsUniqueVisitsItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetDate sets the Date field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportTimeseriesWeeklyStatsUniqueVisitsItem) SetDate(date *string) {
	l.Date = date
	l.require(landingPageReportTimeseriesWeeklyStatsUniqueVisitsItemFieldDate)
}

// SetVal sets the Val field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportTimeseriesWeeklyStatsUniqueVisitsItem) SetVal(val *int) {
	l.Val = val
	l.require(landingPageReportTimeseriesWeeklyStatsUniqueVisitsItemFieldVal)
}

func (l *LandingPageReportTimeseriesWeeklyStatsUniqueVisitsItem) UnmarshalJSON(data []byte) error {
	type unmarshaler LandingPageReportTimeseriesWeeklyStatsUniqueVisitsItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LandingPageReportTimeseriesWeeklyStatsUniqueVisitsItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *LandingPageReportTimeseriesWeeklyStatsUniqueVisitsItem) MarshalJSON() ([]byte, error) {
	type embed LandingPageReportTimeseriesWeeklyStatsUniqueVisitsItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *LandingPageReportTimeseriesWeeklyStatsUniqueVisitsItem) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

var (
	landingPageReportTimeseriesWeeklyStatsVisitsItemFieldDate = big.NewInt(1 << 0)
	landingPageReportTimeseriesWeeklyStatsVisitsItemFieldVal  = big.NewInt(1 << 1)
)

type LandingPageReportTimeseriesWeeklyStatsVisitsItem struct {
	Date *string `json:"date,omitempty" url:"date,omitempty"`
	Val  *int    `json:"val,omitempty" url:"val,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *LandingPageReportTimeseriesWeeklyStatsVisitsItem) GetDate() *string {
	if l == nil {
		return nil
	}
	return l.Date
}

func (l *LandingPageReportTimeseriesWeeklyStatsVisitsItem) GetVal() *int {
	if l == nil {
		return nil
	}
	return l.Val
}

func (l *LandingPageReportTimeseriesWeeklyStatsVisitsItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *LandingPageReportTimeseriesWeeklyStatsVisitsItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetDate sets the Date field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportTimeseriesWeeklyStatsVisitsItem) SetDate(date *string) {
	l.Date = date
	l.require(landingPageReportTimeseriesWeeklyStatsVisitsItemFieldDate)
}

// SetVal sets the Val field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *LandingPageReportTimeseriesWeeklyStatsVisitsItem) SetVal(val *int) {
	l.Val = val
	l.require(landingPageReportTimeseriesWeeklyStatsVisitsItemFieldVal)
}

func (l *LandingPageReportTimeseriesWeeklyStatsVisitsItem) UnmarshalJSON(data []byte) error {
	type unmarshaler LandingPageReportTimeseriesWeeklyStatsVisitsItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LandingPageReportTimeseriesWeeklyStatsVisitsItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *LandingPageReportTimeseriesWeeklyStatsVisitsItem) MarshalJSON() ([]byte, error) {
	type embed LandingPageReportTimeseriesWeeklyStatsVisitsItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *LandingPageReportTimeseriesWeeklyStatsVisitsItem) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

var (
	reportingFacebookAdFieldEmailSourceName       = big.NewInt(1 << 0)
	reportingFacebookAdFieldEndTime               = big.NewInt(1 << 1)
	reportingFacebookAdFieldNeedsAttention        = big.NewInt(1 << 2)
	reportingFacebookAdFieldPausedAt              = big.NewInt(1 << 3)
	reportingFacebookAdFieldThumbnail             = big.NewInt(1 << 4)
	reportingFacebookAdFieldWasCanceledByFacebook = big.NewInt(1 << 5)
	reportingFacebookAdFieldAudience              = big.NewInt(1 << 6)
	reportingFacebookAdFieldAudienceActivity      = big.NewInt(1 << 7)
	reportingFacebookAdFieldBudget                = big.NewInt(1 << 8)
	reportingFacebookAdFieldChannel               = big.NewInt(1 << 9)
	reportingFacebookAdFieldReportSummary         = big.NewInt(1 << 10)
	reportingFacebookAdFieldLinks                 = big.NewInt(1 << 11)
	reportingFacebookAdFieldCanceledAt            = big.NewInt(1 << 12)
	reportingFacebookAdFieldCreateTime            = big.NewInt(1 << 13)
	reportingFacebookAdFieldHasSegment            = big.NewInt(1 << 14)
	reportingFacebookAdFieldID                    = big.NewInt(1 << 15)
	reportingFacebookAdFieldName                  = big.NewInt(1 << 16)
	reportingFacebookAdFieldPublishedTime         = big.NewInt(1 << 17)
	reportingFacebookAdFieldRecipients            = big.NewInt(1 << 18)
	reportingFacebookAdFieldShowReport            = big.NewInt(1 << 19)
	reportingFacebookAdFieldStartTime             = big.NewInt(1 << 20)
	reportingFacebookAdFieldStatus                = big.NewInt(1 << 21)
	reportingFacebookAdFieldType                  = big.NewInt(1 << 22)
	reportingFacebookAdFieldUpdatedAt             = big.NewInt(1 << 23)
	reportingFacebookAdFieldWebID                 = big.NewInt(1 << 24)
)

type ReportingFacebookAd struct {
	EmailSourceName *string `json:"email_source_name,omitempty" url:"email_source_name,omitempty"`
	// The date and time the ad was ended in ISO 8601 format.
	EndTime *time.Time `json:"end_time,omitempty" url:"end_time,omitempty"`
	// If the ad has a problem and needs attention.
	NeedsAttention *bool `json:"needs_attention,omitempty" url:"needs_attention,omitempty"`
	// The date and time the ad was paused in ISO 8601 format.
	PausedAt *time.Time `json:"paused_at,omitempty" url:"paused_at,omitempty"`
	// The URL of the thumbnail for this outreach.
	Thumbnail             *string `json:"thumbnail,omitempty" url:"thumbnail,omitempty"`
	WasCanceledByFacebook *bool   `json:"was_canceled_by_facebook,omitempty" url:"was_canceled_by_facebook,omitempty"`
	// Audience settings
	Audience         *ReportingFacebookAdAudience         `json:"audience,omitempty" url:"audience,omitempty"`
	AudienceActivity *ReportingFacebookAdAudienceActivity `json:"audience_activity,omitempty" url:"audience_activity,omitempty"`
	Budget           *ReportingFacebookAdBudget           `json:"budget,omitempty" url:"budget,omitempty"`
	// Channel settings
	Channel *ReportingFacebookAdChannel `json:"channel,omitempty" url:"channel,omitempty"`
	// Report summary of facebook ad
	ReportSummary *ReportingFacebookAdReportSummary `json:"report_summary,omitempty" url:"report_summary,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []*ReportingFacebookAdLinksItem `json:"_links,omitempty" url:"_links,omitempty"`
	// The date and time the outreach was canceled in ISO 8601 format.
	CanceledAt *time.Time `json:"canceled_at,omitempty" url:"canceled_at,omitempty"`
	// The date and time the outreach was created in ISO 8601 format.
	CreateTime *time.Time `json:"create_time,omitempty" url:"create_time,omitempty"`
	// If this outreach targets a segment of your audience.
	HasSegment *bool `json:"has_segment,omitempty" url:"has_segment,omitempty"`
	// Unique ID of an Outreach.
	ID *string `json:"id,omitempty" url:"id,omitempty"`
	// Title or name of an Outreach.
	Name *string `json:"name,omitempty" url:"name,omitempty"`
	// The date and time the outreach was (or will be) published in ISO 8601 format.
	PublishedTime *time.Time `json:"published_time,omitempty" url:"published_time,omitempty"`
	// High level audience information for who the outreach targets.
	Recipients *FacebookAdRecipients `json:"recipients,omitempty" url:"recipients,omitempty"`
	// Outreach report availability. Note: This property is hotly debated in what it _should_ convey. See [MCP-1371](https://jira.mailchimp.com/browse/MCP-1371) for more context.
	ShowReport *bool `json:"show_report,omitempty" url:"show_report,omitempty"`
	// The date and time the outreach was started in ISO 8601 format.
	StartTime *time.Time `json:"start_time,omitempty" url:"start_time,omitempty"`
	// The status of this outreach.
	Status *FacebookAdStatus `json:"status,omitempty" url:"status,omitempty"`
	// The type of outreach this object is.
	Type *FacebookAdType `json:"type,omitempty" url:"type,omitempty"`
	// The date and time the outreach was last updated in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty" url:"updated_at,omitempty"`
	// The ID used in the Mailchimp web application. For example, for a `regular` outreach, you can view this campaign in your Mailchimp account at `https://{dc}.admin.mailchimp.com/campaigns/show/?id={web_id}`.
	WebID *int `json:"web_id,omitempty" url:"web_id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *ReportingFacebookAd) GetEmailSourceName() *string {
	if r == nil {
		return nil
	}
	return r.EmailSourceName
}

func (r *ReportingFacebookAd) GetEndTime() *time.Time {
	if r == nil {
		return nil
	}
	return r.EndTime
}

func (r *ReportingFacebookAd) GetNeedsAttention() *bool {
	if r == nil {
		return nil
	}
	return r.NeedsAttention
}

func (r *ReportingFacebookAd) GetPausedAt() *time.Time {
	if r == nil {
		return nil
	}
	return r.PausedAt
}

func (r *ReportingFacebookAd) GetThumbnail() *string {
	if r == nil {
		return nil
	}
	return r.Thumbnail
}

func (r *ReportingFacebookAd) GetWasCanceledByFacebook() *bool {
	if r == nil {
		return nil
	}
	return r.WasCanceledByFacebook
}

func (r *ReportingFacebookAd) GetAudience() *ReportingFacebookAdAudience {
	if r == nil {
		return nil
	}
	return r.Audience
}

func (r *ReportingFacebookAd) GetAudienceActivity() *ReportingFacebookAdAudienceActivity {
	if r == nil {
		return nil
	}
	return r.AudienceActivity
}

func (r *ReportingFacebookAd) GetBudget() *ReportingFacebookAdBudget {
	if r == nil {
		return nil
	}
	return r.Budget
}

func (r *ReportingFacebookAd) GetChannel() *ReportingFacebookAdChannel {
	if r == nil {
		return nil
	}
	return r.Channel
}

func (r *ReportingFacebookAd) GetReportSummary() *ReportingFacebookAdReportSummary {
	if r == nil {
		return nil
	}
	return r.ReportSummary
}

func (r *ReportingFacebookAd) GetLinks() []*ReportingFacebookAdLinksItem {
	if r == nil {
		return nil
	}
	return r.Links
}

func (r *ReportingFacebookAd) GetCanceledAt() *time.Time {
	if r == nil {
		return nil
	}
	return r.CanceledAt
}

func (r *ReportingFacebookAd) GetCreateTime() *time.Time {
	if r == nil {
		return nil
	}
	return r.CreateTime
}

func (r *ReportingFacebookAd) GetHasSegment() *bool {
	if r == nil {
		return nil
	}
	return r.HasSegment
}

func (r *ReportingFacebookAd) GetID() *string {
	if r == nil {
		return nil
	}
	return r.ID
}

func (r *ReportingFacebookAd) GetName() *string {
	if r == nil {
		return nil
	}
	return r.Name
}

func (r *ReportingFacebookAd) GetPublishedTime() *time.Time {
	if r == nil {
		return nil
	}
	return r.PublishedTime
}

func (r *ReportingFacebookAd) GetRecipients() *FacebookAdRecipients {
	if r == nil {
		return nil
	}
	return r.Recipients
}

func (r *ReportingFacebookAd) GetShowReport() *bool {
	if r == nil {
		return nil
	}
	return r.ShowReport
}

func (r *ReportingFacebookAd) GetStartTime() *time.Time {
	if r == nil {
		return nil
	}
	return r.StartTime
}

func (r *ReportingFacebookAd) GetStatus() *FacebookAdStatus {
	if r == nil {
		return nil
	}
	return r.Status
}

func (r *ReportingFacebookAd) GetType() *FacebookAdType {
	if r == nil {
		return nil
	}
	return r.Type
}

func (r *ReportingFacebookAd) GetUpdatedAt() *time.Time {
	if r == nil {
		return nil
	}
	return r.UpdatedAt
}

func (r *ReportingFacebookAd) GetWebID() *int {
	if r == nil {
		return nil
	}
	return r.WebID
}

func (r *ReportingFacebookAd) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r *ReportingFacebookAd) require(field *big.Int) {
	if r.explicitFields == nil {
		r.explicitFields = big.NewInt(0)
	}
	r.explicitFields.Or(r.explicitFields, field)
}

// SetEmailSourceName sets the EmailSourceName field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAd) SetEmailSourceName(emailSourceName *string) {
	r.EmailSourceName = emailSourceName
	r.require(reportingFacebookAdFieldEmailSourceName)
}

// SetEndTime sets the EndTime field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAd) SetEndTime(endTime *time.Time) {
	r.EndTime = endTime
	r.require(reportingFacebookAdFieldEndTime)
}

// SetNeedsAttention sets the NeedsAttention field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAd) SetNeedsAttention(needsAttention *bool) {
	r.NeedsAttention = needsAttention
	r.require(reportingFacebookAdFieldNeedsAttention)
}

// SetPausedAt sets the PausedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAd) SetPausedAt(pausedAt *time.Time) {
	r.PausedAt = pausedAt
	r.require(reportingFacebookAdFieldPausedAt)
}

// SetThumbnail sets the Thumbnail field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAd) SetThumbnail(thumbnail *string) {
	r.Thumbnail = thumbnail
	r.require(reportingFacebookAdFieldThumbnail)
}

// SetWasCanceledByFacebook sets the WasCanceledByFacebook field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAd) SetWasCanceledByFacebook(wasCanceledByFacebook *bool) {
	r.WasCanceledByFacebook = wasCanceledByFacebook
	r.require(reportingFacebookAdFieldWasCanceledByFacebook)
}

// SetAudience sets the Audience field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAd) SetAudience(audience *ReportingFacebookAdAudience) {
	r.Audience = audience
	r.require(reportingFacebookAdFieldAudience)
}

// SetAudienceActivity sets the AudienceActivity field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAd) SetAudienceActivity(audienceActivity *ReportingFacebookAdAudienceActivity) {
	r.AudienceActivity = audienceActivity
	r.require(reportingFacebookAdFieldAudienceActivity)
}

// SetBudget sets the Budget field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAd) SetBudget(budget *ReportingFacebookAdBudget) {
	r.Budget = budget
	r.require(reportingFacebookAdFieldBudget)
}

// SetChannel sets the Channel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAd) SetChannel(channel *ReportingFacebookAdChannel) {
	r.Channel = channel
	r.require(reportingFacebookAdFieldChannel)
}

// SetReportSummary sets the ReportSummary field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAd) SetReportSummary(reportSummary *ReportingFacebookAdReportSummary) {
	r.ReportSummary = reportSummary
	r.require(reportingFacebookAdFieldReportSummary)
}

// SetLinks sets the Links field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAd) SetLinks(links []*ReportingFacebookAdLinksItem) {
	r.Links = links
	r.require(reportingFacebookAdFieldLinks)
}

// SetCanceledAt sets the CanceledAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAd) SetCanceledAt(canceledAt *time.Time) {
	r.CanceledAt = canceledAt
	r.require(reportingFacebookAdFieldCanceledAt)
}

// SetCreateTime sets the CreateTime field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAd) SetCreateTime(createTime *time.Time) {
	r.CreateTime = createTime
	r.require(reportingFacebookAdFieldCreateTime)
}

// SetHasSegment sets the HasSegment field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAd) SetHasSegment(hasSegment *bool) {
	r.HasSegment = hasSegment
	r.require(reportingFacebookAdFieldHasSegment)
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAd) SetID(id *string) {
	r.ID = id
	r.require(reportingFacebookAdFieldID)
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAd) SetName(name *string) {
	r.Name = name
	r.require(reportingFacebookAdFieldName)
}

// SetPublishedTime sets the PublishedTime field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAd) SetPublishedTime(publishedTime *time.Time) {
	r.PublishedTime = publishedTime
	r.require(reportingFacebookAdFieldPublishedTime)
}

// SetRecipients sets the Recipients field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAd) SetRecipients(recipients *FacebookAdRecipients) {
	r.Recipients = recipients
	r.require(reportingFacebookAdFieldRecipients)
}

// SetShowReport sets the ShowReport field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAd) SetShowReport(showReport *bool) {
	r.ShowReport = showReport
	r.require(reportingFacebookAdFieldShowReport)
}

// SetStartTime sets the StartTime field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAd) SetStartTime(startTime *time.Time) {
	r.StartTime = startTime
	r.require(reportingFacebookAdFieldStartTime)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAd) SetStatus(status *FacebookAdStatus) {
	r.Status = status
	r.require(reportingFacebookAdFieldStatus)
}

// SetType sets the Type field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAd) SetType(type_ *FacebookAdType) {
	r.Type = type_
	r.require(reportingFacebookAdFieldType)
}

// SetUpdatedAt sets the UpdatedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAd) SetUpdatedAt(updatedAt *time.Time) {
	r.UpdatedAt = updatedAt
	r.require(reportingFacebookAdFieldUpdatedAt)
}

// SetWebID sets the WebID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAd) SetWebID(webID *int) {
	r.WebID = webID
	r.require(reportingFacebookAdFieldWebID)
}

func (r *ReportingFacebookAd) UnmarshalJSON(data []byte) error {
	type embed ReportingFacebookAd
	var unmarshaler = struct {
		embed
		EndTime       *internal.DateTime `json:"end_time,omitempty"`
		PausedAt      *internal.DateTime `json:"paused_at,omitempty"`
		CanceledAt    *internal.DateTime `json:"canceled_at,omitempty"`
		CreateTime    *internal.DateTime `json:"create_time,omitempty"`
		PublishedTime *internal.DateTime `json:"published_time,omitempty"`
		StartTime     *internal.DateTime `json:"start_time,omitempty"`
		UpdatedAt     *internal.DateTime `json:"updated_at,omitempty"`
	}{
		embed: embed(*r),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*r = ReportingFacebookAd(unmarshaler.embed)
	r.EndTime = unmarshaler.EndTime.TimePtr()
	r.PausedAt = unmarshaler.PausedAt.TimePtr()
	r.CanceledAt = unmarshaler.CanceledAt.TimePtr()
	r.CreateTime = unmarshaler.CreateTime.TimePtr()
	r.PublishedTime = unmarshaler.PublishedTime.TimePtr()
	r.StartTime = unmarshaler.StartTime.TimePtr()
	r.UpdatedAt = unmarshaler.UpdatedAt.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *ReportingFacebookAd) MarshalJSON() ([]byte, error) {
	type embed ReportingFacebookAd
	var marshaler = struct {
		embed
		EndTime       *internal.DateTime `json:"end_time,omitempty"`
		PausedAt      *internal.DateTime `json:"paused_at,omitempty"`
		CanceledAt    *internal.DateTime `json:"canceled_at,omitempty"`
		CreateTime    *internal.DateTime `json:"create_time,omitempty"`
		PublishedTime *internal.DateTime `json:"published_time,omitempty"`
		StartTime     *internal.DateTime `json:"start_time,omitempty"`
		UpdatedAt     *internal.DateTime `json:"updated_at,omitempty"`
	}{
		embed:         embed(*r),
		EndTime:       internal.NewOptionalDateTime(r.EndTime),
		PausedAt:      internal.NewOptionalDateTime(r.PausedAt),
		CanceledAt:    internal.NewOptionalDateTime(r.CanceledAt),
		CreateTime:    internal.NewOptionalDateTime(r.CreateTime),
		PublishedTime: internal.NewOptionalDateTime(r.PublishedTime),
		StartTime:     internal.NewOptionalDateTime(r.StartTime),
		UpdatedAt:     internal.NewOptionalDateTime(r.UpdatedAt),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, r.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (r *ReportingFacebookAd) String() string {
	if r == nil {
		return "<nil>"
	}
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

// Audience settings
var (
	reportingFacebookAdAudienceFieldEmailSource           = big.NewInt(1 << 0)
	reportingFacebookAdAudienceFieldIncludeSourceInTarget = big.NewInt(1 << 1)
	reportingFacebookAdAudienceFieldLookalikeCountryCode  = big.NewInt(1 << 2)
	reportingFacebookAdAudienceFieldSourceType            = big.NewInt(1 << 3)
	reportingFacebookAdAudienceFieldTargetingSpecs        = big.NewInt(1 << 4)
	reportingFacebookAdAudienceFieldType                  = big.NewInt(1 << 5)
)

type ReportingFacebookAdAudience struct {
	EmailSource *ReportingFacebookAdAudienceEmailSource `json:"email_source,omitempty" url:"email_source,omitempty"`
	// To include list contacts as part of audience
	IncludeSourceInTarget *bool `json:"include_source_in_target,omitempty" url:"include_source_in_target,omitempty"`
	// To find similar audience in given country
	LookalikeCountryCode *string `json:"lookalike_country_code,omitempty" url:"lookalike_country_code,omitempty"`
	// List or Facebook based audience
	SourceType     *ReportingFacebookAdAudienceSourceType     `json:"source_type,omitempty" url:"source_type,omitempty"`
	TargetingSpecs *ReportingFacebookAdAudienceTargetingSpecs `json:"targeting_specs,omitempty" url:"targeting_specs,omitempty"`
	// Type of the audience
	Type *ReportingFacebookAdAudienceType `json:"type,omitempty" url:"type,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *ReportingFacebookAdAudience) GetEmailSource() *ReportingFacebookAdAudienceEmailSource {
	if r == nil {
		return nil
	}
	return r.EmailSource
}

func (r *ReportingFacebookAdAudience) GetIncludeSourceInTarget() *bool {
	if r == nil {
		return nil
	}
	return r.IncludeSourceInTarget
}

func (r *ReportingFacebookAdAudience) GetLookalikeCountryCode() *string {
	if r == nil {
		return nil
	}
	return r.LookalikeCountryCode
}

func (r *ReportingFacebookAdAudience) GetSourceType() *ReportingFacebookAdAudienceSourceType {
	if r == nil {
		return nil
	}
	return r.SourceType
}

func (r *ReportingFacebookAdAudience) GetTargetingSpecs() *ReportingFacebookAdAudienceTargetingSpecs {
	if r == nil {
		return nil
	}
	return r.TargetingSpecs
}

func (r *ReportingFacebookAdAudience) GetType() *ReportingFacebookAdAudienceType {
	if r == nil {
		return nil
	}
	return r.Type
}

func (r *ReportingFacebookAdAudience) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r *ReportingFacebookAdAudience) require(field *big.Int) {
	if r.explicitFields == nil {
		r.explicitFields = big.NewInt(0)
	}
	r.explicitFields.Or(r.explicitFields, field)
}

// SetEmailSource sets the EmailSource field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdAudience) SetEmailSource(emailSource *ReportingFacebookAdAudienceEmailSource) {
	r.EmailSource = emailSource
	r.require(reportingFacebookAdAudienceFieldEmailSource)
}

// SetIncludeSourceInTarget sets the IncludeSourceInTarget field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdAudience) SetIncludeSourceInTarget(includeSourceInTarget *bool) {
	r.IncludeSourceInTarget = includeSourceInTarget
	r.require(reportingFacebookAdAudienceFieldIncludeSourceInTarget)
}

// SetLookalikeCountryCode sets the LookalikeCountryCode field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdAudience) SetLookalikeCountryCode(lookalikeCountryCode *string) {
	r.LookalikeCountryCode = lookalikeCountryCode
	r.require(reportingFacebookAdAudienceFieldLookalikeCountryCode)
}

// SetSourceType sets the SourceType field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdAudience) SetSourceType(sourceType *ReportingFacebookAdAudienceSourceType) {
	r.SourceType = sourceType
	r.require(reportingFacebookAdAudienceFieldSourceType)
}

// SetTargetingSpecs sets the TargetingSpecs field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdAudience) SetTargetingSpecs(targetingSpecs *ReportingFacebookAdAudienceTargetingSpecs) {
	r.TargetingSpecs = targetingSpecs
	r.require(reportingFacebookAdAudienceFieldTargetingSpecs)
}

// SetType sets the Type field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdAudience) SetType(type_ *ReportingFacebookAdAudienceType) {
	r.Type = type_
	r.require(reportingFacebookAdAudienceFieldType)
}

func (r *ReportingFacebookAdAudience) UnmarshalJSON(data []byte) error {
	type unmarshaler ReportingFacebookAdAudience
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ReportingFacebookAdAudience(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *ReportingFacebookAdAudience) MarshalJSON() ([]byte, error) {
	type embed ReportingFacebookAdAudience
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, r.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (r *ReportingFacebookAdAudience) String() string {
	if r == nil {
		return "<nil>"
	}
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

var (
	reportingFacebookAdAudienceActivityFieldClicks      = big.NewInt(1 << 0)
	reportingFacebookAdAudienceActivityFieldImpressions = big.NewInt(1 << 1)
	reportingFacebookAdAudienceActivityFieldRevenue     = big.NewInt(1 << 2)
)

type ReportingFacebookAdAudienceActivity struct {
	Clicks      []*ReportingFacebookAdAudienceActivityClicksItem      `json:"clicks,omitempty" url:"clicks,omitempty"`
	Impressions []*ReportingFacebookAdAudienceActivityImpressionsItem `json:"impressions,omitempty" url:"impressions,omitempty"`
	Revenue     []*ReportingFacebookAdAudienceActivityRevenueItem     `json:"revenue,omitempty" url:"revenue,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *ReportingFacebookAdAudienceActivity) GetClicks() []*ReportingFacebookAdAudienceActivityClicksItem {
	if r == nil {
		return nil
	}
	return r.Clicks
}

func (r *ReportingFacebookAdAudienceActivity) GetImpressions() []*ReportingFacebookAdAudienceActivityImpressionsItem {
	if r == nil {
		return nil
	}
	return r.Impressions
}

func (r *ReportingFacebookAdAudienceActivity) GetRevenue() []*ReportingFacebookAdAudienceActivityRevenueItem {
	if r == nil {
		return nil
	}
	return r.Revenue
}

func (r *ReportingFacebookAdAudienceActivity) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r *ReportingFacebookAdAudienceActivity) require(field *big.Int) {
	if r.explicitFields == nil {
		r.explicitFields = big.NewInt(0)
	}
	r.explicitFields.Or(r.explicitFields, field)
}

// SetClicks sets the Clicks field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdAudienceActivity) SetClicks(clicks []*ReportingFacebookAdAudienceActivityClicksItem) {
	r.Clicks = clicks
	r.require(reportingFacebookAdAudienceActivityFieldClicks)
}

// SetImpressions sets the Impressions field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdAudienceActivity) SetImpressions(impressions []*ReportingFacebookAdAudienceActivityImpressionsItem) {
	r.Impressions = impressions
	r.require(reportingFacebookAdAudienceActivityFieldImpressions)
}

// SetRevenue sets the Revenue field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdAudienceActivity) SetRevenue(revenue []*ReportingFacebookAdAudienceActivityRevenueItem) {
	r.Revenue = revenue
	r.require(reportingFacebookAdAudienceActivityFieldRevenue)
}

func (r *ReportingFacebookAdAudienceActivity) UnmarshalJSON(data []byte) error {
	type unmarshaler ReportingFacebookAdAudienceActivity
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ReportingFacebookAdAudienceActivity(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *ReportingFacebookAdAudienceActivity) MarshalJSON() ([]byte, error) {
	type embed ReportingFacebookAdAudienceActivity
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, r.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (r *ReportingFacebookAdAudienceActivity) String() string {
	if r == nil {
		return "<nil>"
	}
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

var (
	reportingFacebookAdAudienceActivityClicksItemFieldClicks = big.NewInt(1 << 0)
	reportingFacebookAdAudienceActivityClicksItemFieldDate   = big.NewInt(1 << 1)
)

type ReportingFacebookAdAudienceActivityClicksItem struct {
	Clicks *int    `json:"clicks,omitempty" url:"clicks,omitempty"`
	Date   *string `json:"date,omitempty" url:"date,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *ReportingFacebookAdAudienceActivityClicksItem) GetClicks() *int {
	if r == nil {
		return nil
	}
	return r.Clicks
}

func (r *ReportingFacebookAdAudienceActivityClicksItem) GetDate() *string {
	if r == nil {
		return nil
	}
	return r.Date
}

func (r *ReportingFacebookAdAudienceActivityClicksItem) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r *ReportingFacebookAdAudienceActivityClicksItem) require(field *big.Int) {
	if r.explicitFields == nil {
		r.explicitFields = big.NewInt(0)
	}
	r.explicitFields.Or(r.explicitFields, field)
}

// SetClicks sets the Clicks field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdAudienceActivityClicksItem) SetClicks(clicks *int) {
	r.Clicks = clicks
	r.require(reportingFacebookAdAudienceActivityClicksItemFieldClicks)
}

// SetDate sets the Date field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdAudienceActivityClicksItem) SetDate(date *string) {
	r.Date = date
	r.require(reportingFacebookAdAudienceActivityClicksItemFieldDate)
}

func (r *ReportingFacebookAdAudienceActivityClicksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler ReportingFacebookAdAudienceActivityClicksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ReportingFacebookAdAudienceActivityClicksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *ReportingFacebookAdAudienceActivityClicksItem) MarshalJSON() ([]byte, error) {
	type embed ReportingFacebookAdAudienceActivityClicksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, r.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (r *ReportingFacebookAdAudienceActivityClicksItem) String() string {
	if r == nil {
		return "<nil>"
	}
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

var (
	reportingFacebookAdAudienceActivityImpressionsItemFieldDate        = big.NewInt(1 << 0)
	reportingFacebookAdAudienceActivityImpressionsItemFieldImpressions = big.NewInt(1 << 1)
)

type ReportingFacebookAdAudienceActivityImpressionsItem struct {
	Date        *string `json:"date,omitempty" url:"date,omitempty"`
	Impressions *int    `json:"impressions,omitempty" url:"impressions,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *ReportingFacebookAdAudienceActivityImpressionsItem) GetDate() *string {
	if r == nil {
		return nil
	}
	return r.Date
}

func (r *ReportingFacebookAdAudienceActivityImpressionsItem) GetImpressions() *int {
	if r == nil {
		return nil
	}
	return r.Impressions
}

func (r *ReportingFacebookAdAudienceActivityImpressionsItem) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r *ReportingFacebookAdAudienceActivityImpressionsItem) require(field *big.Int) {
	if r.explicitFields == nil {
		r.explicitFields = big.NewInt(0)
	}
	r.explicitFields.Or(r.explicitFields, field)
}

// SetDate sets the Date field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdAudienceActivityImpressionsItem) SetDate(date *string) {
	r.Date = date
	r.require(reportingFacebookAdAudienceActivityImpressionsItemFieldDate)
}

// SetImpressions sets the Impressions field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdAudienceActivityImpressionsItem) SetImpressions(impressions *int) {
	r.Impressions = impressions
	r.require(reportingFacebookAdAudienceActivityImpressionsItemFieldImpressions)
}

func (r *ReportingFacebookAdAudienceActivityImpressionsItem) UnmarshalJSON(data []byte) error {
	type unmarshaler ReportingFacebookAdAudienceActivityImpressionsItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ReportingFacebookAdAudienceActivityImpressionsItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *ReportingFacebookAdAudienceActivityImpressionsItem) MarshalJSON() ([]byte, error) {
	type embed ReportingFacebookAdAudienceActivityImpressionsItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, r.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (r *ReportingFacebookAdAudienceActivityImpressionsItem) String() string {
	if r == nil {
		return "<nil>"
	}
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

var (
	reportingFacebookAdAudienceActivityRevenueItemFieldDate    = big.NewInt(1 << 0)
	reportingFacebookAdAudienceActivityRevenueItemFieldRevenue = big.NewInt(1 << 1)
)

type ReportingFacebookAdAudienceActivityRevenueItem struct {
	Date    *string  `json:"date,omitempty" url:"date,omitempty"`
	Revenue *float64 `json:"revenue,omitempty" url:"revenue,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *ReportingFacebookAdAudienceActivityRevenueItem) GetDate() *string {
	if r == nil {
		return nil
	}
	return r.Date
}

func (r *ReportingFacebookAdAudienceActivityRevenueItem) GetRevenue() *float64 {
	if r == nil {
		return nil
	}
	return r.Revenue
}

func (r *ReportingFacebookAdAudienceActivityRevenueItem) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r *ReportingFacebookAdAudienceActivityRevenueItem) require(field *big.Int) {
	if r.explicitFields == nil {
		r.explicitFields = big.NewInt(0)
	}
	r.explicitFields.Or(r.explicitFields, field)
}

// SetDate sets the Date field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdAudienceActivityRevenueItem) SetDate(date *string) {
	r.Date = date
	r.require(reportingFacebookAdAudienceActivityRevenueItemFieldDate)
}

// SetRevenue sets the Revenue field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdAudienceActivityRevenueItem) SetRevenue(revenue *float64) {
	r.Revenue = revenue
	r.require(reportingFacebookAdAudienceActivityRevenueItemFieldRevenue)
}

func (r *ReportingFacebookAdAudienceActivityRevenueItem) UnmarshalJSON(data []byte) error {
	type unmarshaler ReportingFacebookAdAudienceActivityRevenueItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ReportingFacebookAdAudienceActivityRevenueItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *ReportingFacebookAdAudienceActivityRevenueItem) MarshalJSON() ([]byte, error) {
	type embed ReportingFacebookAdAudienceActivityRevenueItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, r.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (r *ReportingFacebookAdAudienceActivityRevenueItem) String() string {
	if r == nil {
		return "<nil>"
	}
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

var (
	reportingFacebookAdAudienceEmailSourceFieldIsSegment   = big.NewInt(1 << 0)
	reportingFacebookAdAudienceEmailSourceFieldListName    = big.NewInt(1 << 1)
	reportingFacebookAdAudienceEmailSourceFieldName        = big.NewInt(1 << 2)
	reportingFacebookAdAudienceEmailSourceFieldSegmentType = big.NewInt(1 << 3)
	reportingFacebookAdAudienceEmailSourceFieldType        = big.NewInt(1 << 4)
)

type ReportingFacebookAdAudienceEmailSource struct {
	// Is the source reference a segment
	IsSegment *bool `json:"is_segment,omitempty" url:"is_segment,omitempty"`
	// Associated list name to the source
	ListName *string `json:"list_name,omitempty" url:"list_name,omitempty"`
	// Email source name
	Name *string `json:"name,omitempty" url:"name,omitempty"`
	// Segment type if this source is tied to a segment
	SegmentType *string `json:"segment_type,omitempty" url:"segment_type,omitempty"`
	// Type of the email source
	Type *string `json:"type,omitempty" url:"type,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *ReportingFacebookAdAudienceEmailSource) GetIsSegment() *bool {
	if r == nil {
		return nil
	}
	return r.IsSegment
}

func (r *ReportingFacebookAdAudienceEmailSource) GetListName() *string {
	if r == nil {
		return nil
	}
	return r.ListName
}

func (r *ReportingFacebookAdAudienceEmailSource) GetName() *string {
	if r == nil {
		return nil
	}
	return r.Name
}

func (r *ReportingFacebookAdAudienceEmailSource) GetSegmentType() *string {
	if r == nil {
		return nil
	}
	return r.SegmentType
}

func (r *ReportingFacebookAdAudienceEmailSource) GetType() *string {
	if r == nil {
		return nil
	}
	return r.Type
}

func (r *ReportingFacebookAdAudienceEmailSource) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r *ReportingFacebookAdAudienceEmailSource) require(field *big.Int) {
	if r.explicitFields == nil {
		r.explicitFields = big.NewInt(0)
	}
	r.explicitFields.Or(r.explicitFields, field)
}

// SetIsSegment sets the IsSegment field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdAudienceEmailSource) SetIsSegment(isSegment *bool) {
	r.IsSegment = isSegment
	r.require(reportingFacebookAdAudienceEmailSourceFieldIsSegment)
}

// SetListName sets the ListName field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdAudienceEmailSource) SetListName(listName *string) {
	r.ListName = listName
	r.require(reportingFacebookAdAudienceEmailSourceFieldListName)
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdAudienceEmailSource) SetName(name *string) {
	r.Name = name
	r.require(reportingFacebookAdAudienceEmailSourceFieldName)
}

// SetSegmentType sets the SegmentType field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdAudienceEmailSource) SetSegmentType(segmentType *string) {
	r.SegmentType = segmentType
	r.require(reportingFacebookAdAudienceEmailSourceFieldSegmentType)
}

// SetType sets the Type field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdAudienceEmailSource) SetType(type_ *string) {
	r.Type = type_
	r.require(reportingFacebookAdAudienceEmailSourceFieldType)
}

func (r *ReportingFacebookAdAudienceEmailSource) UnmarshalJSON(data []byte) error {
	type unmarshaler ReportingFacebookAdAudienceEmailSource
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ReportingFacebookAdAudienceEmailSource(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *ReportingFacebookAdAudienceEmailSource) MarshalJSON() ([]byte, error) {
	type embed ReportingFacebookAdAudienceEmailSource
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, r.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (r *ReportingFacebookAdAudienceEmailSource) String() string {
	if r == nil {
		return "<nil>"
	}
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

// List or Facebook based audience
type ReportingFacebookAdAudienceSourceType string

const (
	ReportingFacebookAdAudienceSourceTypeFacebook ReportingFacebookAdAudienceSourceType = "facebook"
	ReportingFacebookAdAudienceSourceTypeList     ReportingFacebookAdAudienceSourceType = "list"
)

func NewReportingFacebookAdAudienceSourceTypeFromString(s string) (ReportingFacebookAdAudienceSourceType, error) {
	switch s {
	case "facebook":
		return ReportingFacebookAdAudienceSourceTypeFacebook, nil
	case "list":
		return ReportingFacebookAdAudienceSourceTypeList, nil
	}
	var t ReportingFacebookAdAudienceSourceType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (r ReportingFacebookAdAudienceSourceType) Ptr() *ReportingFacebookAdAudienceSourceType {
	return &r
}

var (
	reportingFacebookAdAudienceTargetingSpecsFieldGender    = big.NewInt(1 << 0)
	reportingFacebookAdAudienceTargetingSpecsFieldInterests = big.NewInt(1 << 1)
	reportingFacebookAdAudienceTargetingSpecsFieldLocations = big.NewInt(1 << 2)
	reportingFacebookAdAudienceTargetingSpecsFieldMaxAge    = big.NewInt(1 << 3)
	reportingFacebookAdAudienceTargetingSpecsFieldMinAge    = big.NewInt(1 << 4)
)

type ReportingFacebookAdAudienceTargetingSpecs struct {
	Gender    *int                                                      `json:"gender,omitempty" url:"gender,omitempty"`
	Interests []*ReportingFacebookAdAudienceTargetingSpecsInterestsItem `json:"interests,omitempty" url:"interests,omitempty"`
	Locations *ReportingFacebookAdAudienceTargetingSpecsLocations       `json:"locations,omitempty" url:"locations,omitempty"`
	MaxAge    *int                                                      `json:"max_age,omitempty" url:"max_age,omitempty"`
	MinAge    *int                                                      `json:"min_age,omitempty" url:"min_age,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *ReportingFacebookAdAudienceTargetingSpecs) GetGender() *int {
	if r == nil {
		return nil
	}
	return r.Gender
}

func (r *ReportingFacebookAdAudienceTargetingSpecs) GetInterests() []*ReportingFacebookAdAudienceTargetingSpecsInterestsItem {
	if r == nil {
		return nil
	}
	return r.Interests
}

func (r *ReportingFacebookAdAudienceTargetingSpecs) GetLocations() *ReportingFacebookAdAudienceTargetingSpecsLocations {
	if r == nil {
		return nil
	}
	return r.Locations
}

func (r *ReportingFacebookAdAudienceTargetingSpecs) GetMaxAge() *int {
	if r == nil {
		return nil
	}
	return r.MaxAge
}

func (r *ReportingFacebookAdAudienceTargetingSpecs) GetMinAge() *int {
	if r == nil {
		return nil
	}
	return r.MinAge
}

func (r *ReportingFacebookAdAudienceTargetingSpecs) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r *ReportingFacebookAdAudienceTargetingSpecs) require(field *big.Int) {
	if r.explicitFields == nil {
		r.explicitFields = big.NewInt(0)
	}
	r.explicitFields.Or(r.explicitFields, field)
}

// SetGender sets the Gender field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdAudienceTargetingSpecs) SetGender(gender *int) {
	r.Gender = gender
	r.require(reportingFacebookAdAudienceTargetingSpecsFieldGender)
}

// SetInterests sets the Interests field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdAudienceTargetingSpecs) SetInterests(interests []*ReportingFacebookAdAudienceTargetingSpecsInterestsItem) {
	r.Interests = interests
	r.require(reportingFacebookAdAudienceTargetingSpecsFieldInterests)
}

// SetLocations sets the Locations field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdAudienceTargetingSpecs) SetLocations(locations *ReportingFacebookAdAudienceTargetingSpecsLocations) {
	r.Locations = locations
	r.require(reportingFacebookAdAudienceTargetingSpecsFieldLocations)
}

// SetMaxAge sets the MaxAge field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdAudienceTargetingSpecs) SetMaxAge(maxAge *int) {
	r.MaxAge = maxAge
	r.require(reportingFacebookAdAudienceTargetingSpecsFieldMaxAge)
}

// SetMinAge sets the MinAge field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdAudienceTargetingSpecs) SetMinAge(minAge *int) {
	r.MinAge = minAge
	r.require(reportingFacebookAdAudienceTargetingSpecsFieldMinAge)
}

func (r *ReportingFacebookAdAudienceTargetingSpecs) UnmarshalJSON(data []byte) error {
	type unmarshaler ReportingFacebookAdAudienceTargetingSpecs
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ReportingFacebookAdAudienceTargetingSpecs(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *ReportingFacebookAdAudienceTargetingSpecs) MarshalJSON() ([]byte, error) {
	type embed ReportingFacebookAdAudienceTargetingSpecs
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, r.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (r *ReportingFacebookAdAudienceTargetingSpecs) String() string {
	if r == nil {
		return "<nil>"
	}
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

var (
	reportingFacebookAdAudienceTargetingSpecsInterestsItemFieldName = big.NewInt(1 << 0)
)

type ReportingFacebookAdAudienceTargetingSpecsInterestsItem struct {
	Name *string `json:"name,omitempty" url:"name,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *ReportingFacebookAdAudienceTargetingSpecsInterestsItem) GetName() *string {
	if r == nil {
		return nil
	}
	return r.Name
}

func (r *ReportingFacebookAdAudienceTargetingSpecsInterestsItem) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r *ReportingFacebookAdAudienceTargetingSpecsInterestsItem) require(field *big.Int) {
	if r.explicitFields == nil {
		r.explicitFields = big.NewInt(0)
	}
	r.explicitFields.Or(r.explicitFields, field)
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdAudienceTargetingSpecsInterestsItem) SetName(name *string) {
	r.Name = name
	r.require(reportingFacebookAdAudienceTargetingSpecsInterestsItemFieldName)
}

func (r *ReportingFacebookAdAudienceTargetingSpecsInterestsItem) UnmarshalJSON(data []byte) error {
	type unmarshaler ReportingFacebookAdAudienceTargetingSpecsInterestsItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ReportingFacebookAdAudienceTargetingSpecsInterestsItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *ReportingFacebookAdAudienceTargetingSpecsInterestsItem) MarshalJSON() ([]byte, error) {
	type embed ReportingFacebookAdAudienceTargetingSpecsInterestsItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, r.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (r *ReportingFacebookAdAudienceTargetingSpecsInterestsItem) String() string {
	if r == nil {
		return "<nil>"
	}
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

var (
	reportingFacebookAdAudienceTargetingSpecsLocationsFieldCities    = big.NewInt(1 << 0)
	reportingFacebookAdAudienceTargetingSpecsLocationsFieldCountries = big.NewInt(1 << 1)
	reportingFacebookAdAudienceTargetingSpecsLocationsFieldRegions   = big.NewInt(1 << 2)
	reportingFacebookAdAudienceTargetingSpecsLocationsFieldZips      = big.NewInt(1 << 3)
)

type ReportingFacebookAdAudienceTargetingSpecsLocations struct {
	Cities    []string `json:"cities,omitempty" url:"cities,omitempty"`
	Countries []string `json:"countries,omitempty" url:"countries,omitempty"`
	Regions   []string `json:"regions,omitempty" url:"regions,omitempty"`
	Zips      []string `json:"zips,omitempty" url:"zips,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *ReportingFacebookAdAudienceTargetingSpecsLocations) GetCities() []string {
	if r == nil {
		return nil
	}
	return r.Cities
}

func (r *ReportingFacebookAdAudienceTargetingSpecsLocations) GetCountries() []string {
	if r == nil {
		return nil
	}
	return r.Countries
}

func (r *ReportingFacebookAdAudienceTargetingSpecsLocations) GetRegions() []string {
	if r == nil {
		return nil
	}
	return r.Regions
}

func (r *ReportingFacebookAdAudienceTargetingSpecsLocations) GetZips() []string {
	if r == nil {
		return nil
	}
	return r.Zips
}

func (r *ReportingFacebookAdAudienceTargetingSpecsLocations) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r *ReportingFacebookAdAudienceTargetingSpecsLocations) require(field *big.Int) {
	if r.explicitFields == nil {
		r.explicitFields = big.NewInt(0)
	}
	r.explicitFields.Or(r.explicitFields, field)
}

// SetCities sets the Cities field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdAudienceTargetingSpecsLocations) SetCities(cities []string) {
	r.Cities = cities
	r.require(reportingFacebookAdAudienceTargetingSpecsLocationsFieldCities)
}

// SetCountries sets the Countries field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdAudienceTargetingSpecsLocations) SetCountries(countries []string) {
	r.Countries = countries
	r.require(reportingFacebookAdAudienceTargetingSpecsLocationsFieldCountries)
}

// SetRegions sets the Regions field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdAudienceTargetingSpecsLocations) SetRegions(regions []string) {
	r.Regions = regions
	r.require(reportingFacebookAdAudienceTargetingSpecsLocationsFieldRegions)
}

// SetZips sets the Zips field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdAudienceTargetingSpecsLocations) SetZips(zips []string) {
	r.Zips = zips
	r.require(reportingFacebookAdAudienceTargetingSpecsLocationsFieldZips)
}

func (r *ReportingFacebookAdAudienceTargetingSpecsLocations) UnmarshalJSON(data []byte) error {
	type unmarshaler ReportingFacebookAdAudienceTargetingSpecsLocations
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ReportingFacebookAdAudienceTargetingSpecsLocations(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *ReportingFacebookAdAudienceTargetingSpecsLocations) MarshalJSON() ([]byte, error) {
	type embed ReportingFacebookAdAudienceTargetingSpecsLocations
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, r.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (r *ReportingFacebookAdAudienceTargetingSpecsLocations) String() string {
	if r == nil {
		return "<nil>"
	}
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

// Type of the audience
type ReportingFacebookAdAudienceType string

const (
	ReportingFacebookAdAudienceTypeCustomAudience        ReportingFacebookAdAudienceType = "Custom Audience"
	ReportingFacebookAdAudienceTypeLookalikeAudience     ReportingFacebookAdAudienceType = "Lookalike Audience"
	ReportingFacebookAdAudienceTypeInterestBasedAudience ReportingFacebookAdAudienceType = "Interest-based Audience"
)

func NewReportingFacebookAdAudienceTypeFromString(s string) (ReportingFacebookAdAudienceType, error) {
	switch s {
	case "Custom Audience":
		return ReportingFacebookAdAudienceTypeCustomAudience, nil
	case "Lookalike Audience":
		return ReportingFacebookAdAudienceTypeLookalikeAudience, nil
	case "Interest-based Audience":
		return ReportingFacebookAdAudienceTypeInterestBasedAudience, nil
	}
	var t ReportingFacebookAdAudienceType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (r ReportingFacebookAdAudienceType) Ptr() *ReportingFacebookAdAudienceType {
	return &r
}

var (
	reportingFacebookAdBudgetFieldCurrencyCode = big.NewInt(1 << 0)
	reportingFacebookAdBudgetFieldDuration     = big.NewInt(1 << 1)
	reportingFacebookAdBudgetFieldTotalAmount  = big.NewInt(1 << 2)
)

type ReportingFacebookAdBudget struct {
	// Currency code
	CurrencyCode *string `json:"currency_code,omitempty" url:"currency_code,omitempty"`
	// Duration of the ad in seconds
	Duration *int `json:"duration,omitempty" url:"duration,omitempty"`
	// Total budget of the ad
	TotalAmount *float64 `json:"total_amount,omitempty" url:"total_amount,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *ReportingFacebookAdBudget) GetCurrencyCode() *string {
	if r == nil {
		return nil
	}
	return r.CurrencyCode
}

func (r *ReportingFacebookAdBudget) GetDuration() *int {
	if r == nil {
		return nil
	}
	return r.Duration
}

func (r *ReportingFacebookAdBudget) GetTotalAmount() *float64 {
	if r == nil {
		return nil
	}
	return r.TotalAmount
}

func (r *ReportingFacebookAdBudget) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r *ReportingFacebookAdBudget) require(field *big.Int) {
	if r.explicitFields == nil {
		r.explicitFields = big.NewInt(0)
	}
	r.explicitFields.Or(r.explicitFields, field)
}

// SetCurrencyCode sets the CurrencyCode field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdBudget) SetCurrencyCode(currencyCode *string) {
	r.CurrencyCode = currencyCode
	r.require(reportingFacebookAdBudgetFieldCurrencyCode)
}

// SetDuration sets the Duration field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdBudget) SetDuration(duration *int) {
	r.Duration = duration
	r.require(reportingFacebookAdBudgetFieldDuration)
}

// SetTotalAmount sets the TotalAmount field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdBudget) SetTotalAmount(totalAmount *float64) {
	r.TotalAmount = totalAmount
	r.require(reportingFacebookAdBudgetFieldTotalAmount)
}

func (r *ReportingFacebookAdBudget) UnmarshalJSON(data []byte) error {
	type unmarshaler ReportingFacebookAdBudget
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ReportingFacebookAdBudget(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *ReportingFacebookAdBudget) MarshalJSON() ([]byte, error) {
	type embed ReportingFacebookAdBudget
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, r.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (r *ReportingFacebookAdBudget) String() string {
	if r == nil {
		return "<nil>"
	}
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

// Channel settings
var (
	reportingFacebookAdChannelFieldFbPlacementAudience = big.NewInt(1 << 0)
	reportingFacebookAdChannelFieldFbPlacementFeed     = big.NewInt(1 << 1)
	reportingFacebookAdChannelFieldIgPlacementFeed     = big.NewInt(1 << 2)
)

type ReportingFacebookAdChannel struct {
	// Is this for facebook audience
	FbPlacementAudience *bool `json:"fb_placement_audience,omitempty" url:"fb_placement_audience,omitempty"`
	// Is this for facebook feed
	FbPlacementFeed *bool `json:"fb_placement_feed,omitempty" url:"fb_placement_feed,omitempty"`
	// Is this for instagram feed
	IgPlacementFeed *bool `json:"ig_placement_feed,omitempty" url:"ig_placement_feed,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *ReportingFacebookAdChannel) GetFbPlacementAudience() *bool {
	if r == nil {
		return nil
	}
	return r.FbPlacementAudience
}

func (r *ReportingFacebookAdChannel) GetFbPlacementFeed() *bool {
	if r == nil {
		return nil
	}
	return r.FbPlacementFeed
}

func (r *ReportingFacebookAdChannel) GetIgPlacementFeed() *bool {
	if r == nil {
		return nil
	}
	return r.IgPlacementFeed
}

func (r *ReportingFacebookAdChannel) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r *ReportingFacebookAdChannel) require(field *big.Int) {
	if r.explicitFields == nil {
		r.explicitFields = big.NewInt(0)
	}
	r.explicitFields.Or(r.explicitFields, field)
}

// SetFbPlacementAudience sets the FbPlacementAudience field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdChannel) SetFbPlacementAudience(fbPlacementAudience *bool) {
	r.FbPlacementAudience = fbPlacementAudience
	r.require(reportingFacebookAdChannelFieldFbPlacementAudience)
}

// SetFbPlacementFeed sets the FbPlacementFeed field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdChannel) SetFbPlacementFeed(fbPlacementFeed *bool) {
	r.FbPlacementFeed = fbPlacementFeed
	r.require(reportingFacebookAdChannelFieldFbPlacementFeed)
}

// SetIgPlacementFeed sets the IgPlacementFeed field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdChannel) SetIgPlacementFeed(igPlacementFeed *bool) {
	r.IgPlacementFeed = igPlacementFeed
	r.require(reportingFacebookAdChannelFieldIgPlacementFeed)
}

func (r *ReportingFacebookAdChannel) UnmarshalJSON(data []byte) error {
	type unmarshaler ReportingFacebookAdChannel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ReportingFacebookAdChannel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *ReportingFacebookAdChannel) MarshalJSON() ([]byte, error) {
	type embed ReportingFacebookAdChannel
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, r.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (r *ReportingFacebookAdChannel) String() string {
	if r == nil {
		return "<nil>"
	}
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

// This object represents a link from the resource where it is found to another resource or action that may be performed.
var (
	reportingFacebookAdLinksItemFieldHref         = big.NewInt(1 << 0)
	reportingFacebookAdLinksItemFieldMethod       = big.NewInt(1 << 1)
	reportingFacebookAdLinksItemFieldRel          = big.NewInt(1 << 2)
	reportingFacebookAdLinksItemFieldSchema       = big.NewInt(1 << 3)
	reportingFacebookAdLinksItemFieldTargetSchema = big.NewInt(1 << 4)
)

type ReportingFacebookAdLinksItem struct {
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href *string `json:"href,omitempty" url:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method *ReportingFacebookAdLinksItemMethod `json:"method,omitempty" url:"method,omitempty"`
	// As with an HTML 'rel' attribute, this describes the type of link.
	Rel *string `json:"rel,omitempty" url:"rel,omitempty"`
	// For HTTP methods that can receive bodies (POST and PUT), this is a URL representing the schema that the body should conform to.
	Schema *string `json:"schema,omitempty" url:"schema,omitempty"`
	// For GETs, this is a URL representing the schema that the response should conform to.
	TargetSchema *string `json:"targetSchema,omitempty" url:"targetSchema,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *ReportingFacebookAdLinksItem) GetHref() *string {
	if r == nil {
		return nil
	}
	return r.Href
}

func (r *ReportingFacebookAdLinksItem) GetMethod() *ReportingFacebookAdLinksItemMethod {
	if r == nil {
		return nil
	}
	return r.Method
}

func (r *ReportingFacebookAdLinksItem) GetRel() *string {
	if r == nil {
		return nil
	}
	return r.Rel
}

func (r *ReportingFacebookAdLinksItem) GetSchema() *string {
	if r == nil {
		return nil
	}
	return r.Schema
}

func (r *ReportingFacebookAdLinksItem) GetTargetSchema() *string {
	if r == nil {
		return nil
	}
	return r.TargetSchema
}

func (r *ReportingFacebookAdLinksItem) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r *ReportingFacebookAdLinksItem) require(field *big.Int) {
	if r.explicitFields == nil {
		r.explicitFields = big.NewInt(0)
	}
	r.explicitFields.Or(r.explicitFields, field)
}

// SetHref sets the Href field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdLinksItem) SetHref(href *string) {
	r.Href = href
	r.require(reportingFacebookAdLinksItemFieldHref)
}

// SetMethod sets the Method field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdLinksItem) SetMethod(method *ReportingFacebookAdLinksItemMethod) {
	r.Method = method
	r.require(reportingFacebookAdLinksItemFieldMethod)
}

// SetRel sets the Rel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdLinksItem) SetRel(rel *string) {
	r.Rel = rel
	r.require(reportingFacebookAdLinksItemFieldRel)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdLinksItem) SetSchema(schema *string) {
	r.Schema = schema
	r.require(reportingFacebookAdLinksItemFieldSchema)
}

// SetTargetSchema sets the TargetSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdLinksItem) SetTargetSchema(targetSchema *string) {
	r.TargetSchema = targetSchema
	r.require(reportingFacebookAdLinksItemFieldTargetSchema)
}

func (r *ReportingFacebookAdLinksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler ReportingFacebookAdLinksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ReportingFacebookAdLinksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *ReportingFacebookAdLinksItem) MarshalJSON() ([]byte, error) {
	type embed ReportingFacebookAdLinksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, r.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (r *ReportingFacebookAdLinksItem) String() string {
	if r == nil {
		return "<nil>"
	}
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

// The HTTP method that should be used when accessing the URL defined in 'href'.
type ReportingFacebookAdLinksItemMethod string

const (
	ReportingFacebookAdLinksItemMethodGet     ReportingFacebookAdLinksItemMethod = "GET"
	ReportingFacebookAdLinksItemMethodPost    ReportingFacebookAdLinksItemMethod = "POST"
	ReportingFacebookAdLinksItemMethodPut     ReportingFacebookAdLinksItemMethod = "PUT"
	ReportingFacebookAdLinksItemMethodPatch   ReportingFacebookAdLinksItemMethod = "PATCH"
	ReportingFacebookAdLinksItemMethodDelete  ReportingFacebookAdLinksItemMethod = "DELETE"
	ReportingFacebookAdLinksItemMethodOptions ReportingFacebookAdLinksItemMethod = "OPTIONS"
	ReportingFacebookAdLinksItemMethodHead    ReportingFacebookAdLinksItemMethod = "HEAD"
)

func NewReportingFacebookAdLinksItemMethodFromString(s string) (ReportingFacebookAdLinksItemMethod, error) {
	switch s {
	case "GET":
		return ReportingFacebookAdLinksItemMethodGet, nil
	case "POST":
		return ReportingFacebookAdLinksItemMethodPost, nil
	case "PUT":
		return ReportingFacebookAdLinksItemMethodPut, nil
	case "PATCH":
		return ReportingFacebookAdLinksItemMethodPatch, nil
	case "DELETE":
		return ReportingFacebookAdLinksItemMethodDelete, nil
	case "OPTIONS":
		return ReportingFacebookAdLinksItemMethodOptions, nil
	case "HEAD":
		return ReportingFacebookAdLinksItemMethodHead, nil
	}
	var t ReportingFacebookAdLinksItemMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (r ReportingFacebookAdLinksItemMethod) Ptr() *ReportingFacebookAdLinksItemMethod {
	return &r
}

// Report summary of facebook ad
var (
	reportingFacebookAdReportSummaryFieldAverageDailyBudget    = big.NewInt(1 << 0)
	reportingFacebookAdReportSummaryFieldAverageOrderAmount    = big.NewInt(1 << 1)
	reportingFacebookAdReportSummaryFieldClickRate             = big.NewInt(1 << 2)
	reportingFacebookAdReportSummaryFieldClicks                = big.NewInt(1 << 3)
	reportingFacebookAdReportSummaryFieldComments              = big.NewInt(1 << 4)
	reportingFacebookAdReportSummaryFieldCostPerClick          = big.NewInt(1 << 5)
	reportingFacebookAdReportSummaryFieldEcommerce             = big.NewInt(1 << 6)
	reportingFacebookAdReportSummaryFieldExtendedAt            = big.NewInt(1 << 7)
	reportingFacebookAdReportSummaryFieldFirstTimeBuyers       = big.NewInt(1 << 8)
	reportingFacebookAdReportSummaryFieldHasExtendedAdDuration = big.NewInt(1 << 9)
	reportingFacebookAdReportSummaryFieldImpressions           = big.NewInt(1 << 10)
	reportingFacebookAdReportSummaryFieldLikes                 = big.NewInt(1 << 11)
	reportingFacebookAdReportSummaryFieldReach                 = big.NewInt(1 << 12)
	reportingFacebookAdReportSummaryFieldReturnOnInvestment    = big.NewInt(1 << 13)
	reportingFacebookAdReportSummaryFieldShares                = big.NewInt(1 << 14)
	reportingFacebookAdReportSummaryFieldTotalOrders           = big.NewInt(1 << 15)
	reportingFacebookAdReportSummaryFieldTotalProductsSold     = big.NewInt(1 << 16)
	reportingFacebookAdReportSummaryFieldUniqueClicks          = big.NewInt(1 << 17)
)

type ReportingFacebookAdReportSummary struct {
	AverageDailyBudget    *ReportingFacebookAdReportSummaryAverageDailyBudget `json:"average_daily_budget,omitempty" url:"average_daily_budget,omitempty"`
	AverageOrderAmount    *ReportingFacebookAdReportSummaryAverageOrderAmount `json:"average_order_amount,omitempty" url:"average_order_amount,omitempty"`
	ClickRate             *float64                                            `json:"click_rate,omitempty" url:"click_rate,omitempty"`
	Clicks                *int                                                `json:"clicks,omitempty" url:"clicks,omitempty"`
	Comments              *int                                                `json:"comments,omitempty" url:"comments,omitempty"`
	CostPerClick          *ReportingFacebookAdReportSummaryCostPerClick       `json:"cost_per_click,omitempty" url:"cost_per_click,omitempty"`
	Ecommerce             *ReportingFacebookAdReportSummaryEcommerce          `json:"ecommerce,omitempty" url:"ecommerce,omitempty"`
	ExtendedAt            *ReportingFacebookAdReportSummaryExtendedAt         `json:"extended_at,omitempty" url:"extended_at,omitempty"`
	FirstTimeBuyers       *int                                                `json:"first_time_buyers,omitempty" url:"first_time_buyers,omitempty"`
	HasExtendedAdDuration *bool                                               `json:"has_extended_ad_duration,omitempty" url:"has_extended_ad_duration,omitempty"`
	Impressions           *int                                                `json:"impressions,omitempty" url:"impressions,omitempty"`
	Likes                 *int                                                `json:"likes,omitempty" url:"likes,omitempty"`
	Reach                 *int                                                `json:"reach,omitempty" url:"reach,omitempty"`
	ReturnOnInvestment    *float64                                            `json:"return_on_investment,omitempty" url:"return_on_investment,omitempty"`
	Shares                *int                                                `json:"shares,omitempty" url:"shares,omitempty"`
	TotalOrders           *int                                                `json:"total_orders,omitempty" url:"total_orders,omitempty"`
	TotalProductsSold     *int                                                `json:"total_products_sold,omitempty" url:"total_products_sold,omitempty"`
	UniqueClicks          *int                                                `json:"unique_clicks,omitempty" url:"unique_clicks,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *ReportingFacebookAdReportSummary) GetAverageDailyBudget() *ReportingFacebookAdReportSummaryAverageDailyBudget {
	if r == nil {
		return nil
	}
	return r.AverageDailyBudget
}

func (r *ReportingFacebookAdReportSummary) GetAverageOrderAmount() *ReportingFacebookAdReportSummaryAverageOrderAmount {
	if r == nil {
		return nil
	}
	return r.AverageOrderAmount
}

func (r *ReportingFacebookAdReportSummary) GetClickRate() *float64 {
	if r == nil {
		return nil
	}
	return r.ClickRate
}

func (r *ReportingFacebookAdReportSummary) GetClicks() *int {
	if r == nil {
		return nil
	}
	return r.Clicks
}

func (r *ReportingFacebookAdReportSummary) GetComments() *int {
	if r == nil {
		return nil
	}
	return r.Comments
}

func (r *ReportingFacebookAdReportSummary) GetCostPerClick() *ReportingFacebookAdReportSummaryCostPerClick {
	if r == nil {
		return nil
	}
	return r.CostPerClick
}

func (r *ReportingFacebookAdReportSummary) GetEcommerce() *ReportingFacebookAdReportSummaryEcommerce {
	if r == nil {
		return nil
	}
	return r.Ecommerce
}

func (r *ReportingFacebookAdReportSummary) GetExtendedAt() *ReportingFacebookAdReportSummaryExtendedAt {
	if r == nil {
		return nil
	}
	return r.ExtendedAt
}

func (r *ReportingFacebookAdReportSummary) GetFirstTimeBuyers() *int {
	if r == nil {
		return nil
	}
	return r.FirstTimeBuyers
}

func (r *ReportingFacebookAdReportSummary) GetHasExtendedAdDuration() *bool {
	if r == nil {
		return nil
	}
	return r.HasExtendedAdDuration
}

func (r *ReportingFacebookAdReportSummary) GetImpressions() *int {
	if r == nil {
		return nil
	}
	return r.Impressions
}

func (r *ReportingFacebookAdReportSummary) GetLikes() *int {
	if r == nil {
		return nil
	}
	return r.Likes
}

func (r *ReportingFacebookAdReportSummary) GetReach() *int {
	if r == nil {
		return nil
	}
	return r.Reach
}

func (r *ReportingFacebookAdReportSummary) GetReturnOnInvestment() *float64 {
	if r == nil {
		return nil
	}
	return r.ReturnOnInvestment
}

func (r *ReportingFacebookAdReportSummary) GetShares() *int {
	if r == nil {
		return nil
	}
	return r.Shares
}

func (r *ReportingFacebookAdReportSummary) GetTotalOrders() *int {
	if r == nil {
		return nil
	}
	return r.TotalOrders
}

func (r *ReportingFacebookAdReportSummary) GetTotalProductsSold() *int {
	if r == nil {
		return nil
	}
	return r.TotalProductsSold
}

func (r *ReportingFacebookAdReportSummary) GetUniqueClicks() *int {
	if r == nil {
		return nil
	}
	return r.UniqueClicks
}

func (r *ReportingFacebookAdReportSummary) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r *ReportingFacebookAdReportSummary) require(field *big.Int) {
	if r.explicitFields == nil {
		r.explicitFields = big.NewInt(0)
	}
	r.explicitFields.Or(r.explicitFields, field)
}

// SetAverageDailyBudget sets the AverageDailyBudget field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdReportSummary) SetAverageDailyBudget(averageDailyBudget *ReportingFacebookAdReportSummaryAverageDailyBudget) {
	r.AverageDailyBudget = averageDailyBudget
	r.require(reportingFacebookAdReportSummaryFieldAverageDailyBudget)
}

// SetAverageOrderAmount sets the AverageOrderAmount field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdReportSummary) SetAverageOrderAmount(averageOrderAmount *ReportingFacebookAdReportSummaryAverageOrderAmount) {
	r.AverageOrderAmount = averageOrderAmount
	r.require(reportingFacebookAdReportSummaryFieldAverageOrderAmount)
}

// SetClickRate sets the ClickRate field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdReportSummary) SetClickRate(clickRate *float64) {
	r.ClickRate = clickRate
	r.require(reportingFacebookAdReportSummaryFieldClickRate)
}

// SetClicks sets the Clicks field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdReportSummary) SetClicks(clicks *int) {
	r.Clicks = clicks
	r.require(reportingFacebookAdReportSummaryFieldClicks)
}

// SetComments sets the Comments field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdReportSummary) SetComments(comments *int) {
	r.Comments = comments
	r.require(reportingFacebookAdReportSummaryFieldComments)
}

// SetCostPerClick sets the CostPerClick field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdReportSummary) SetCostPerClick(costPerClick *ReportingFacebookAdReportSummaryCostPerClick) {
	r.CostPerClick = costPerClick
	r.require(reportingFacebookAdReportSummaryFieldCostPerClick)
}

// SetEcommerce sets the Ecommerce field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdReportSummary) SetEcommerce(ecommerce *ReportingFacebookAdReportSummaryEcommerce) {
	r.Ecommerce = ecommerce
	r.require(reportingFacebookAdReportSummaryFieldEcommerce)
}

// SetExtendedAt sets the ExtendedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdReportSummary) SetExtendedAt(extendedAt *ReportingFacebookAdReportSummaryExtendedAt) {
	r.ExtendedAt = extendedAt
	r.require(reportingFacebookAdReportSummaryFieldExtendedAt)
}

// SetFirstTimeBuyers sets the FirstTimeBuyers field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdReportSummary) SetFirstTimeBuyers(firstTimeBuyers *int) {
	r.FirstTimeBuyers = firstTimeBuyers
	r.require(reportingFacebookAdReportSummaryFieldFirstTimeBuyers)
}

// SetHasExtendedAdDuration sets the HasExtendedAdDuration field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdReportSummary) SetHasExtendedAdDuration(hasExtendedAdDuration *bool) {
	r.HasExtendedAdDuration = hasExtendedAdDuration
	r.require(reportingFacebookAdReportSummaryFieldHasExtendedAdDuration)
}

// SetImpressions sets the Impressions field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdReportSummary) SetImpressions(impressions *int) {
	r.Impressions = impressions
	r.require(reportingFacebookAdReportSummaryFieldImpressions)
}

// SetLikes sets the Likes field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdReportSummary) SetLikes(likes *int) {
	r.Likes = likes
	r.require(reportingFacebookAdReportSummaryFieldLikes)
}

// SetReach sets the Reach field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdReportSummary) SetReach(reach *int) {
	r.Reach = reach
	r.require(reportingFacebookAdReportSummaryFieldReach)
}

// SetReturnOnInvestment sets the ReturnOnInvestment field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdReportSummary) SetReturnOnInvestment(returnOnInvestment *float64) {
	r.ReturnOnInvestment = returnOnInvestment
	r.require(reportingFacebookAdReportSummaryFieldReturnOnInvestment)
}

// SetShares sets the Shares field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdReportSummary) SetShares(shares *int) {
	r.Shares = shares
	r.require(reportingFacebookAdReportSummaryFieldShares)
}

// SetTotalOrders sets the TotalOrders field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdReportSummary) SetTotalOrders(totalOrders *int) {
	r.TotalOrders = totalOrders
	r.require(reportingFacebookAdReportSummaryFieldTotalOrders)
}

// SetTotalProductsSold sets the TotalProductsSold field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdReportSummary) SetTotalProductsSold(totalProductsSold *int) {
	r.TotalProductsSold = totalProductsSold
	r.require(reportingFacebookAdReportSummaryFieldTotalProductsSold)
}

// SetUniqueClicks sets the UniqueClicks field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdReportSummary) SetUniqueClicks(uniqueClicks *int) {
	r.UniqueClicks = uniqueClicks
	r.require(reportingFacebookAdReportSummaryFieldUniqueClicks)
}

func (r *ReportingFacebookAdReportSummary) UnmarshalJSON(data []byte) error {
	type unmarshaler ReportingFacebookAdReportSummary
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ReportingFacebookAdReportSummary(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *ReportingFacebookAdReportSummary) MarshalJSON() ([]byte, error) {
	type embed ReportingFacebookAdReportSummary
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, r.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (r *ReportingFacebookAdReportSummary) String() string {
	if r == nil {
		return "<nil>"
	}
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

var (
	reportingFacebookAdReportSummaryAverageDailyBudgetFieldAmount       = big.NewInt(1 << 0)
	reportingFacebookAdReportSummaryAverageDailyBudgetFieldCurrencyCode = big.NewInt(1 << 1)
)

type ReportingFacebookAdReportSummaryAverageDailyBudget struct {
	Amount       *float64 `json:"amount,omitempty" url:"amount,omitempty"`
	CurrencyCode *string  `json:"currency_code,omitempty" url:"currency_code,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *ReportingFacebookAdReportSummaryAverageDailyBudget) GetAmount() *float64 {
	if r == nil {
		return nil
	}
	return r.Amount
}

func (r *ReportingFacebookAdReportSummaryAverageDailyBudget) GetCurrencyCode() *string {
	if r == nil {
		return nil
	}
	return r.CurrencyCode
}

func (r *ReportingFacebookAdReportSummaryAverageDailyBudget) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r *ReportingFacebookAdReportSummaryAverageDailyBudget) require(field *big.Int) {
	if r.explicitFields == nil {
		r.explicitFields = big.NewInt(0)
	}
	r.explicitFields.Or(r.explicitFields, field)
}

// SetAmount sets the Amount field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdReportSummaryAverageDailyBudget) SetAmount(amount *float64) {
	r.Amount = amount
	r.require(reportingFacebookAdReportSummaryAverageDailyBudgetFieldAmount)
}

// SetCurrencyCode sets the CurrencyCode field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdReportSummaryAverageDailyBudget) SetCurrencyCode(currencyCode *string) {
	r.CurrencyCode = currencyCode
	r.require(reportingFacebookAdReportSummaryAverageDailyBudgetFieldCurrencyCode)
}

func (r *ReportingFacebookAdReportSummaryAverageDailyBudget) UnmarshalJSON(data []byte) error {
	type unmarshaler ReportingFacebookAdReportSummaryAverageDailyBudget
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ReportingFacebookAdReportSummaryAverageDailyBudget(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *ReportingFacebookAdReportSummaryAverageDailyBudget) MarshalJSON() ([]byte, error) {
	type embed ReportingFacebookAdReportSummaryAverageDailyBudget
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, r.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (r *ReportingFacebookAdReportSummaryAverageDailyBudget) String() string {
	if r == nil {
		return "<nil>"
	}
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

var (
	reportingFacebookAdReportSummaryAverageOrderAmountFieldAmount       = big.NewInt(1 << 0)
	reportingFacebookAdReportSummaryAverageOrderAmountFieldCurrencyCode = big.NewInt(1 << 1)
)

type ReportingFacebookAdReportSummaryAverageOrderAmount struct {
	Amount       *float64 `json:"amount,omitempty" url:"amount,omitempty"`
	CurrencyCode *string  `json:"currency_code,omitempty" url:"currency_code,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *ReportingFacebookAdReportSummaryAverageOrderAmount) GetAmount() *float64 {
	if r == nil {
		return nil
	}
	return r.Amount
}

func (r *ReportingFacebookAdReportSummaryAverageOrderAmount) GetCurrencyCode() *string {
	if r == nil {
		return nil
	}
	return r.CurrencyCode
}

func (r *ReportingFacebookAdReportSummaryAverageOrderAmount) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r *ReportingFacebookAdReportSummaryAverageOrderAmount) require(field *big.Int) {
	if r.explicitFields == nil {
		r.explicitFields = big.NewInt(0)
	}
	r.explicitFields.Or(r.explicitFields, field)
}

// SetAmount sets the Amount field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdReportSummaryAverageOrderAmount) SetAmount(amount *float64) {
	r.Amount = amount
	r.require(reportingFacebookAdReportSummaryAverageOrderAmountFieldAmount)
}

// SetCurrencyCode sets the CurrencyCode field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdReportSummaryAverageOrderAmount) SetCurrencyCode(currencyCode *string) {
	r.CurrencyCode = currencyCode
	r.require(reportingFacebookAdReportSummaryAverageOrderAmountFieldCurrencyCode)
}

func (r *ReportingFacebookAdReportSummaryAverageOrderAmount) UnmarshalJSON(data []byte) error {
	type unmarshaler ReportingFacebookAdReportSummaryAverageOrderAmount
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ReportingFacebookAdReportSummaryAverageOrderAmount(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *ReportingFacebookAdReportSummaryAverageOrderAmount) MarshalJSON() ([]byte, error) {
	type embed ReportingFacebookAdReportSummaryAverageOrderAmount
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, r.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (r *ReportingFacebookAdReportSummaryAverageOrderAmount) String() string {
	if r == nil {
		return "<nil>"
	}
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

var (
	reportingFacebookAdReportSummaryCostPerClickFieldAmount       = big.NewInt(1 << 0)
	reportingFacebookAdReportSummaryCostPerClickFieldCurrencyCode = big.NewInt(1 << 1)
)

type ReportingFacebookAdReportSummaryCostPerClick struct {
	Amount       *float64 `json:"amount,omitempty" url:"amount,omitempty"`
	CurrencyCode *string  `json:"currency_code,omitempty" url:"currency_code,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *ReportingFacebookAdReportSummaryCostPerClick) GetAmount() *float64 {
	if r == nil {
		return nil
	}
	return r.Amount
}

func (r *ReportingFacebookAdReportSummaryCostPerClick) GetCurrencyCode() *string {
	if r == nil {
		return nil
	}
	return r.CurrencyCode
}

func (r *ReportingFacebookAdReportSummaryCostPerClick) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r *ReportingFacebookAdReportSummaryCostPerClick) require(field *big.Int) {
	if r.explicitFields == nil {
		r.explicitFields = big.NewInt(0)
	}
	r.explicitFields.Or(r.explicitFields, field)
}

// SetAmount sets the Amount field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdReportSummaryCostPerClick) SetAmount(amount *float64) {
	r.Amount = amount
	r.require(reportingFacebookAdReportSummaryCostPerClickFieldAmount)
}

// SetCurrencyCode sets the CurrencyCode field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdReportSummaryCostPerClick) SetCurrencyCode(currencyCode *string) {
	r.CurrencyCode = currencyCode
	r.require(reportingFacebookAdReportSummaryCostPerClickFieldCurrencyCode)
}

func (r *ReportingFacebookAdReportSummaryCostPerClick) UnmarshalJSON(data []byte) error {
	type unmarshaler ReportingFacebookAdReportSummaryCostPerClick
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ReportingFacebookAdReportSummaryCostPerClick(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *ReportingFacebookAdReportSummaryCostPerClick) MarshalJSON() ([]byte, error) {
	type embed ReportingFacebookAdReportSummaryCostPerClick
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, r.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (r *ReportingFacebookAdReportSummaryCostPerClick) String() string {
	if r == nil {
		return "<nil>"
	}
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

var (
	reportingFacebookAdReportSummaryEcommerceFieldCurrencyCode = big.NewInt(1 << 0)
	reportingFacebookAdReportSummaryEcommerceFieldTotalRevenue = big.NewInt(1 << 1)
)

type ReportingFacebookAdReportSummaryEcommerce struct {
	CurrencyCode *string  `json:"currency_code,omitempty" url:"currency_code,omitempty"`
	TotalRevenue *float64 `json:"total_revenue,omitempty" url:"total_revenue,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *ReportingFacebookAdReportSummaryEcommerce) GetCurrencyCode() *string {
	if r == nil {
		return nil
	}
	return r.CurrencyCode
}

func (r *ReportingFacebookAdReportSummaryEcommerce) GetTotalRevenue() *float64 {
	if r == nil {
		return nil
	}
	return r.TotalRevenue
}

func (r *ReportingFacebookAdReportSummaryEcommerce) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r *ReportingFacebookAdReportSummaryEcommerce) require(field *big.Int) {
	if r.explicitFields == nil {
		r.explicitFields = big.NewInt(0)
	}
	r.explicitFields.Or(r.explicitFields, field)
}

// SetCurrencyCode sets the CurrencyCode field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdReportSummaryEcommerce) SetCurrencyCode(currencyCode *string) {
	r.CurrencyCode = currencyCode
	r.require(reportingFacebookAdReportSummaryEcommerceFieldCurrencyCode)
}

// SetTotalRevenue sets the TotalRevenue field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdReportSummaryEcommerce) SetTotalRevenue(totalRevenue *float64) {
	r.TotalRevenue = totalRevenue
	r.require(reportingFacebookAdReportSummaryEcommerceFieldTotalRevenue)
}

func (r *ReportingFacebookAdReportSummaryEcommerce) UnmarshalJSON(data []byte) error {
	type unmarshaler ReportingFacebookAdReportSummaryEcommerce
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ReportingFacebookAdReportSummaryEcommerce(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *ReportingFacebookAdReportSummaryEcommerce) MarshalJSON() ([]byte, error) {
	type embed ReportingFacebookAdReportSummaryEcommerce
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, r.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (r *ReportingFacebookAdReportSummaryEcommerce) String() string {
	if r == nil {
		return "<nil>"
	}
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

var (
	reportingFacebookAdReportSummaryExtendedAtFieldDatetime = big.NewInt(1 << 0)
	reportingFacebookAdReportSummaryExtendedAtFieldTimezone = big.NewInt(1 << 1)
)

type ReportingFacebookAdReportSummaryExtendedAt struct {
	Datetime *string `json:"datetime,omitempty" url:"datetime,omitempty"`
	Timezone *string `json:"timezone,omitempty" url:"timezone,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *ReportingFacebookAdReportSummaryExtendedAt) GetDatetime() *string {
	if r == nil {
		return nil
	}
	return r.Datetime
}

func (r *ReportingFacebookAdReportSummaryExtendedAt) GetTimezone() *string {
	if r == nil {
		return nil
	}
	return r.Timezone
}

func (r *ReportingFacebookAdReportSummaryExtendedAt) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r *ReportingFacebookAdReportSummaryExtendedAt) require(field *big.Int) {
	if r.explicitFields == nil {
		r.explicitFields = big.NewInt(0)
	}
	r.explicitFields.Or(r.explicitFields, field)
}

// SetDatetime sets the Datetime field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdReportSummaryExtendedAt) SetDatetime(datetime *string) {
	r.Datetime = datetime
	r.require(reportingFacebookAdReportSummaryExtendedAtFieldDatetime)
}

// SetTimezone sets the Timezone field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ReportingFacebookAdReportSummaryExtendedAt) SetTimezone(timezone *string) {
	r.Timezone = timezone
	r.require(reportingFacebookAdReportSummaryExtendedAtFieldTimezone)
}

func (r *ReportingFacebookAdReportSummaryExtendedAt) UnmarshalJSON(data []byte) error {
	type unmarshaler ReportingFacebookAdReportSummaryExtendedAt
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ReportingFacebookAdReportSummaryExtendedAt(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *ReportingFacebookAdReportSummaryExtendedAt) MarshalJSON() ([]byte, error) {
	type embed ReportingFacebookAdReportSummaryExtendedAt
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, r.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (r *ReportingFacebookAdReportSummaryExtendedAt) String() string {
	if r == nil {
		return "<nil>"
	}
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

// The details of a survey question's report.
var (
	surveyQuestionReportFieldAverageRating            = big.NewInt(1 << 0)
	surveyQuestionReportFieldContactCounts            = big.NewInt(1 << 1)
	surveyQuestionReportFieldHasOther                 = big.NewInt(1 << 2)
	surveyQuestionReportFieldID                       = big.NewInt(1 << 3)
	surveyQuestionReportFieldIsRequired               = big.NewInt(1 << 4)
	surveyQuestionReportFieldMergeField               = big.NewInt(1 << 5)
	surveyQuestionReportFieldOptions                  = big.NewInt(1 << 6)
	surveyQuestionReportFieldOtherLabel               = big.NewInt(1 << 7)
	surveyQuestionReportFieldPlaceholderLabel         = big.NewInt(1 << 8)
	surveyQuestionReportFieldQuery                    = big.NewInt(1 << 9)
	surveyQuestionReportFieldRangeHighLabel           = big.NewInt(1 << 10)
	surveyQuestionReportFieldRangeLowLabel            = big.NewInt(1 << 11)
	surveyQuestionReportFieldSubscribeCheckboxEnabled = big.NewInt(1 << 12)
	surveyQuestionReportFieldSubscribeCheckboxLabel   = big.NewInt(1 << 13)
	surveyQuestionReportFieldSurveyID                 = big.NewInt(1 << 14)
	surveyQuestionReportFieldTotalResponses           = big.NewInt(1 << 15)
	surveyQuestionReportFieldType                     = big.NewInt(1 << 16)
)

type SurveyQuestionReport struct {
	// The average rating for this range question.
	AverageRating *float64 `json:"average_rating,omitempty" url:"average_rating,omitempty"`
	// For email question types, how many are new, known, or unknown contacts.
	ContactCounts *SurveyQuestionReportContactCounts `json:"contact_counts,omitempty" url:"contact_counts,omitempty"`
	// Whether this survey question has an 'other' option.
	HasOther *bool `json:"has_other,omitempty" url:"has_other,omitempty"`
	// The ID of the survey question.
	ID *string `json:"id,omitempty" url:"id,omitempty"`
	// Whether this survey question is required to answer.
	IsRequired *bool `json:"is_required,omitempty" url:"is_required,omitempty"`
	// A [merge field](https://mailchimp.com/developer/marketing/docs/merge-fields/) for an audience.
	MergeField *SurveyQuestionReportMergeField `json:"merge_field,omitempty" url:"merge_field,omitempty"`
	// The answer choices for this question.
	Options []*SurveyQuestionReportOptionsItem `json:"options,omitempty" url:"options,omitempty"`
	// Label used for the 'other' option of this survey question.
	OtherLabel *string `json:"other_label,omitempty" url:"other_label,omitempty"`
	// Placeholder text for this survey question's answer box.
	PlaceholderLabel *string `json:"placeholder_label,omitempty" url:"placeholder_label,omitempty"`
	// The query of the survey question.
	Query *string `json:"query,omitempty" url:"query,omitempty"`
	// Label for the high end of the range.
	RangeHighLabel *string `json:"range_high_label,omitempty" url:"range_high_label,omitempty"`
	// Label for the low end of the range.
	RangeLowLabel *string `json:"range_low_label,omitempty" url:"range_low_label,omitempty"`
	// Whether the subscribe checkbox is shown for this email question.
	SubscribeCheckboxEnabled *bool `json:"subscribe_checkbox_enabled,omitempty" url:"subscribe_checkbox_enabled,omitempty"`
	// Label used for the subscribe checkbox for this email question.
	SubscribeCheckboxLabel *string `json:"subscribe_checkbox_label,omitempty" url:"subscribe_checkbox_label,omitempty"`
	// The unique ID of the survey.
	SurveyID *string `json:"survey_id,omitempty" url:"survey_id,omitempty"`
	// The total number of responses to this question.
	TotalResponses *int `json:"total_responses,omitempty" url:"total_responses,omitempty"`
	// The response type of the survey question.
	Type *SurveyQuestionReportType `json:"type,omitempty" url:"type,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *SurveyQuestionReport) GetAverageRating() *float64 {
	if s == nil {
		return nil
	}
	return s.AverageRating
}

func (s *SurveyQuestionReport) GetContactCounts() *SurveyQuestionReportContactCounts {
	if s == nil {
		return nil
	}
	return s.ContactCounts
}

func (s *SurveyQuestionReport) GetHasOther() *bool {
	if s == nil {
		return nil
	}
	return s.HasOther
}

func (s *SurveyQuestionReport) GetID() *string {
	if s == nil {
		return nil
	}
	return s.ID
}

func (s *SurveyQuestionReport) GetIsRequired() *bool {
	if s == nil {
		return nil
	}
	return s.IsRequired
}

func (s *SurveyQuestionReport) GetMergeField() *SurveyQuestionReportMergeField {
	if s == nil {
		return nil
	}
	return s.MergeField
}

func (s *SurveyQuestionReport) GetOptions() []*SurveyQuestionReportOptionsItem {
	if s == nil {
		return nil
	}
	return s.Options
}

func (s *SurveyQuestionReport) GetOtherLabel() *string {
	if s == nil {
		return nil
	}
	return s.OtherLabel
}

func (s *SurveyQuestionReport) GetPlaceholderLabel() *string {
	if s == nil {
		return nil
	}
	return s.PlaceholderLabel
}

func (s *SurveyQuestionReport) GetQuery() *string {
	if s == nil {
		return nil
	}
	return s.Query
}

func (s *SurveyQuestionReport) GetRangeHighLabel() *string {
	if s == nil {
		return nil
	}
	return s.RangeHighLabel
}

func (s *SurveyQuestionReport) GetRangeLowLabel() *string {
	if s == nil {
		return nil
	}
	return s.RangeLowLabel
}

func (s *SurveyQuestionReport) GetSubscribeCheckboxEnabled() *bool {
	if s == nil {
		return nil
	}
	return s.SubscribeCheckboxEnabled
}

func (s *SurveyQuestionReport) GetSubscribeCheckboxLabel() *string {
	if s == nil {
		return nil
	}
	return s.SubscribeCheckboxLabel
}

func (s *SurveyQuestionReport) GetSurveyID() *string {
	if s == nil {
		return nil
	}
	return s.SurveyID
}

func (s *SurveyQuestionReport) GetTotalResponses() *int {
	if s == nil {
		return nil
	}
	return s.TotalResponses
}

func (s *SurveyQuestionReport) GetType() *SurveyQuestionReportType {
	if s == nil {
		return nil
	}
	return s.Type
}

func (s *SurveyQuestionReport) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s *SurveyQuestionReport) require(field *big.Int) {
	if s.explicitFields == nil {
		s.explicitFields = big.NewInt(0)
	}
	s.explicitFields.Or(s.explicitFields, field)
}

// SetAverageRating sets the AverageRating field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SurveyQuestionReport) SetAverageRating(averageRating *float64) {
	s.AverageRating = averageRating
	s.require(surveyQuestionReportFieldAverageRating)
}

// SetContactCounts sets the ContactCounts field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SurveyQuestionReport) SetContactCounts(contactCounts *SurveyQuestionReportContactCounts) {
	s.ContactCounts = contactCounts
	s.require(surveyQuestionReportFieldContactCounts)
}

// SetHasOther sets the HasOther field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SurveyQuestionReport) SetHasOther(hasOther *bool) {
	s.HasOther = hasOther
	s.require(surveyQuestionReportFieldHasOther)
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SurveyQuestionReport) SetID(id *string) {
	s.ID = id
	s.require(surveyQuestionReportFieldID)
}

// SetIsRequired sets the IsRequired field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SurveyQuestionReport) SetIsRequired(isRequired *bool) {
	s.IsRequired = isRequired
	s.require(surveyQuestionReportFieldIsRequired)
}

// SetMergeField sets the MergeField field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SurveyQuestionReport) SetMergeField(mergeField *SurveyQuestionReportMergeField) {
	s.MergeField = mergeField
	s.require(surveyQuestionReportFieldMergeField)
}

// SetOptions sets the Options field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SurveyQuestionReport) SetOptions(options []*SurveyQuestionReportOptionsItem) {
	s.Options = options
	s.require(surveyQuestionReportFieldOptions)
}

// SetOtherLabel sets the OtherLabel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SurveyQuestionReport) SetOtherLabel(otherLabel *string) {
	s.OtherLabel = otherLabel
	s.require(surveyQuestionReportFieldOtherLabel)
}

// SetPlaceholderLabel sets the PlaceholderLabel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SurveyQuestionReport) SetPlaceholderLabel(placeholderLabel *string) {
	s.PlaceholderLabel = placeholderLabel
	s.require(surveyQuestionReportFieldPlaceholderLabel)
}

// SetQuery sets the Query field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SurveyQuestionReport) SetQuery(query *string) {
	s.Query = query
	s.require(surveyQuestionReportFieldQuery)
}

// SetRangeHighLabel sets the RangeHighLabel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SurveyQuestionReport) SetRangeHighLabel(rangeHighLabel *string) {
	s.RangeHighLabel = rangeHighLabel
	s.require(surveyQuestionReportFieldRangeHighLabel)
}

// SetRangeLowLabel sets the RangeLowLabel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SurveyQuestionReport) SetRangeLowLabel(rangeLowLabel *string) {
	s.RangeLowLabel = rangeLowLabel
	s.require(surveyQuestionReportFieldRangeLowLabel)
}

// SetSubscribeCheckboxEnabled sets the SubscribeCheckboxEnabled field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SurveyQuestionReport) SetSubscribeCheckboxEnabled(subscribeCheckboxEnabled *bool) {
	s.SubscribeCheckboxEnabled = subscribeCheckboxEnabled
	s.require(surveyQuestionReportFieldSubscribeCheckboxEnabled)
}

// SetSubscribeCheckboxLabel sets the SubscribeCheckboxLabel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SurveyQuestionReport) SetSubscribeCheckboxLabel(subscribeCheckboxLabel *string) {
	s.SubscribeCheckboxLabel = subscribeCheckboxLabel
	s.require(surveyQuestionReportFieldSubscribeCheckboxLabel)
}

// SetSurveyID sets the SurveyID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SurveyQuestionReport) SetSurveyID(surveyID *string) {
	s.SurveyID = surveyID
	s.require(surveyQuestionReportFieldSurveyID)
}

// SetTotalResponses sets the TotalResponses field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SurveyQuestionReport) SetTotalResponses(totalResponses *int) {
	s.TotalResponses = totalResponses
	s.require(surveyQuestionReportFieldTotalResponses)
}

// SetType sets the Type field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SurveyQuestionReport) SetType(type_ *SurveyQuestionReportType) {
	s.Type = type_
	s.require(surveyQuestionReportFieldType)
}

func (s *SurveyQuestionReport) UnmarshalJSON(data []byte) error {
	type unmarshaler SurveyQuestionReport
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SurveyQuestionReport(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *SurveyQuestionReport) MarshalJSON() ([]byte, error) {
	type embed SurveyQuestionReport
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, s.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (s *SurveyQuestionReport) String() string {
	if s == nil {
		return "<nil>"
	}
	if len(s.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(s.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

// For email question types, how many are new, known, or unknown contacts.
var (
	surveyQuestionReportContactCountsFieldKnown   = big.NewInt(1 << 0)
	surveyQuestionReportContactCountsFieldNew     = big.NewInt(1 << 1)
	surveyQuestionReportContactCountsFieldUnknown = big.NewInt(1 << 2)
)

type SurveyQuestionReportContactCounts struct {
	// The number of known contacts that responded to this survey.
	Known *int `json:"known,omitempty" url:"known,omitempty"`
	// The number of new contacts that responded to this survey.
	New *int `json:"new,omitempty" url:"new,omitempty"`
	// The number of unknown contacts that responded to this survey.
	Unknown *int `json:"unknown,omitempty" url:"unknown,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *SurveyQuestionReportContactCounts) GetKnown() *int {
	if s == nil {
		return nil
	}
	return s.Known
}

func (s *SurveyQuestionReportContactCounts) GetNew() *int {
	if s == nil {
		return nil
	}
	return s.New
}

func (s *SurveyQuestionReportContactCounts) GetUnknown() *int {
	if s == nil {
		return nil
	}
	return s.Unknown
}

func (s *SurveyQuestionReportContactCounts) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s *SurveyQuestionReportContactCounts) require(field *big.Int) {
	if s.explicitFields == nil {
		s.explicitFields = big.NewInt(0)
	}
	s.explicitFields.Or(s.explicitFields, field)
}

// SetKnown sets the Known field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SurveyQuestionReportContactCounts) SetKnown(known *int) {
	s.Known = known
	s.require(surveyQuestionReportContactCountsFieldKnown)
}

// SetNew sets the New field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SurveyQuestionReportContactCounts) SetNew(new_ *int) {
	s.New = new_
	s.require(surveyQuestionReportContactCountsFieldNew)
}

// SetUnknown sets the Unknown field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SurveyQuestionReportContactCounts) SetUnknown(unknown *int) {
	s.Unknown = unknown
	s.require(surveyQuestionReportContactCountsFieldUnknown)
}

func (s *SurveyQuestionReportContactCounts) UnmarshalJSON(data []byte) error {
	type unmarshaler SurveyQuestionReportContactCounts
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SurveyQuestionReportContactCounts(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *SurveyQuestionReportContactCounts) MarshalJSON() ([]byte, error) {
	type embed SurveyQuestionReportContactCounts
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, s.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (s *SurveyQuestionReportContactCounts) String() string {
	if s == nil {
		return "<nil>"
	}
	if len(s.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(s.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

// A [merge field](https://mailchimp.com/developer/marketing/docs/merge-fields/) for an audience.
var (
	surveyQuestionReportMergeFieldFieldID    = big.NewInt(1 << 0)
	surveyQuestionReportMergeFieldFieldLabel = big.NewInt(1 << 1)
	surveyQuestionReportMergeFieldFieldType  = big.NewInt(1 << 2)
)

type SurveyQuestionReportMergeField struct {
	// An unchanging id for the merge field.
	ID *int `json:"id,omitempty" url:"id,omitempty"`
	// The [label](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for the merge field.
	Label *string `json:"label,omitempty" url:"label,omitempty"`
	// The [type](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for the merge field.
	Type *SurveyQuestionReportMergeFieldType `json:"type,omitempty" url:"type,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *SurveyQuestionReportMergeField) GetID() *int {
	if s == nil {
		return nil
	}
	return s.ID
}

func (s *SurveyQuestionReportMergeField) GetLabel() *string {
	if s == nil {
		return nil
	}
	return s.Label
}

func (s *SurveyQuestionReportMergeField) GetType() *SurveyQuestionReportMergeFieldType {
	if s == nil {
		return nil
	}
	return s.Type
}

func (s *SurveyQuestionReportMergeField) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s *SurveyQuestionReportMergeField) require(field *big.Int) {
	if s.explicitFields == nil {
		s.explicitFields = big.NewInt(0)
	}
	s.explicitFields.Or(s.explicitFields, field)
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SurveyQuestionReportMergeField) SetID(id *int) {
	s.ID = id
	s.require(surveyQuestionReportMergeFieldFieldID)
}

// SetLabel sets the Label field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SurveyQuestionReportMergeField) SetLabel(label *string) {
	s.Label = label
	s.require(surveyQuestionReportMergeFieldFieldLabel)
}

// SetType sets the Type field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SurveyQuestionReportMergeField) SetType(type_ *SurveyQuestionReportMergeFieldType) {
	s.Type = type_
	s.require(surveyQuestionReportMergeFieldFieldType)
}

func (s *SurveyQuestionReportMergeField) UnmarshalJSON(data []byte) error {
	type unmarshaler SurveyQuestionReportMergeField
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SurveyQuestionReportMergeField(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *SurveyQuestionReportMergeField) MarshalJSON() ([]byte, error) {
	type embed SurveyQuestionReportMergeField
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, s.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (s *SurveyQuestionReportMergeField) String() string {
	if s == nil {
		return "<nil>"
	}
	if len(s.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(s.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

// The [type](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for the merge field.
type SurveyQuestionReportMergeFieldType string

const (
	SurveyQuestionReportMergeFieldTypeText     SurveyQuestionReportMergeFieldType = "text"
	SurveyQuestionReportMergeFieldTypeNumber   SurveyQuestionReportMergeFieldType = "number"
	SurveyQuestionReportMergeFieldTypeAddress  SurveyQuestionReportMergeFieldType = "address"
	SurveyQuestionReportMergeFieldTypePhone    SurveyQuestionReportMergeFieldType = "phone"
	SurveyQuestionReportMergeFieldTypeDate     SurveyQuestionReportMergeFieldType = "date"
	SurveyQuestionReportMergeFieldTypeURL      SurveyQuestionReportMergeFieldType = "url"
	SurveyQuestionReportMergeFieldTypeImageurl SurveyQuestionReportMergeFieldType = "imageurl"
	SurveyQuestionReportMergeFieldTypeRadio    SurveyQuestionReportMergeFieldType = "radio"
	SurveyQuestionReportMergeFieldTypeDropdown SurveyQuestionReportMergeFieldType = "dropdown"
	SurveyQuestionReportMergeFieldTypeBirthday SurveyQuestionReportMergeFieldType = "birthday"
	SurveyQuestionReportMergeFieldTypeZip      SurveyQuestionReportMergeFieldType = "zip"
)

func NewSurveyQuestionReportMergeFieldTypeFromString(s string) (SurveyQuestionReportMergeFieldType, error) {
	switch s {
	case "text":
		return SurveyQuestionReportMergeFieldTypeText, nil
	case "number":
		return SurveyQuestionReportMergeFieldTypeNumber, nil
	case "address":
		return SurveyQuestionReportMergeFieldTypeAddress, nil
	case "phone":
		return SurveyQuestionReportMergeFieldTypePhone, nil
	case "date":
		return SurveyQuestionReportMergeFieldTypeDate, nil
	case "url":
		return SurveyQuestionReportMergeFieldTypeURL, nil
	case "imageurl":
		return SurveyQuestionReportMergeFieldTypeImageurl, nil
	case "radio":
		return SurveyQuestionReportMergeFieldTypeRadio, nil
	case "dropdown":
		return SurveyQuestionReportMergeFieldTypeDropdown, nil
	case "birthday":
		return SurveyQuestionReportMergeFieldTypeBirthday, nil
	case "zip":
		return SurveyQuestionReportMergeFieldTypeZip, nil
	}
	var t SurveyQuestionReportMergeFieldType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (s SurveyQuestionReportMergeFieldType) Ptr() *SurveyQuestionReportMergeFieldType {
	return &s
}

var (
	surveyQuestionReportOptionsItemFieldCount = big.NewInt(1 << 0)
	surveyQuestionReportOptionsItemFieldID    = big.NewInt(1 << 1)
	surveyQuestionReportOptionsItemFieldLabel = big.NewInt(1 << 2)
)

type SurveyQuestionReportOptionsItem struct {
	// The count of responses that selected this survey question option.
	Count *int `json:"count,omitempty" url:"count,omitempty"`
	// The ID for this survey question option.
	ID *string `json:"id,omitempty" url:"id,omitempty"`
	// The label for this survey question option.
	Label *string `json:"label,omitempty" url:"label,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *SurveyQuestionReportOptionsItem) GetCount() *int {
	if s == nil {
		return nil
	}
	return s.Count
}

func (s *SurveyQuestionReportOptionsItem) GetID() *string {
	if s == nil {
		return nil
	}
	return s.ID
}

func (s *SurveyQuestionReportOptionsItem) GetLabel() *string {
	if s == nil {
		return nil
	}
	return s.Label
}

func (s *SurveyQuestionReportOptionsItem) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s *SurveyQuestionReportOptionsItem) require(field *big.Int) {
	if s.explicitFields == nil {
		s.explicitFields = big.NewInt(0)
	}
	s.explicitFields.Or(s.explicitFields, field)
}

// SetCount sets the Count field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SurveyQuestionReportOptionsItem) SetCount(count *int) {
	s.Count = count
	s.require(surveyQuestionReportOptionsItemFieldCount)
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SurveyQuestionReportOptionsItem) SetID(id *string) {
	s.ID = id
	s.require(surveyQuestionReportOptionsItemFieldID)
}

// SetLabel sets the Label field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SurveyQuestionReportOptionsItem) SetLabel(label *string) {
	s.Label = label
	s.require(surveyQuestionReportOptionsItemFieldLabel)
}

func (s *SurveyQuestionReportOptionsItem) UnmarshalJSON(data []byte) error {
	type unmarshaler SurveyQuestionReportOptionsItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SurveyQuestionReportOptionsItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *SurveyQuestionReportOptionsItem) MarshalJSON() ([]byte, error) {
	type embed SurveyQuestionReportOptionsItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, s.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (s *SurveyQuestionReportOptionsItem) String() string {
	if s == nil {
		return "<nil>"
	}
	if len(s.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(s.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

// The response type of the survey question.
type SurveyQuestionReportType string

const (
	SurveyQuestionReportTypePickOne            SurveyQuestionReportType = "pickOne"
	SurveyQuestionReportTypePickMany           SurveyQuestionReportType = "pickMany"
	SurveyQuestionReportTypeRange              SurveyQuestionReportType = "range"
	SurveyQuestionReportTypeText               SurveyQuestionReportType = "text"
	SurveyQuestionReportTypeEmail              SurveyQuestionReportType = "email"
	SurveyQuestionReportTypeContactInformation SurveyQuestionReportType = "contactInformation"
	SurveyQuestionReportTypeDropdown           SurveyQuestionReportType = "dropdown"
)

func NewSurveyQuestionReportTypeFromString(s string) (SurveyQuestionReportType, error) {
	switch s {
	case "pickOne":
		return SurveyQuestionReportTypePickOne, nil
	case "pickMany":
		return SurveyQuestionReportTypePickMany, nil
	case "range":
		return SurveyQuestionReportTypeRange, nil
	case "text":
		return SurveyQuestionReportTypeText, nil
	case "email":
		return SurveyQuestionReportTypeEmail, nil
	case "contactInformation":
		return SurveyQuestionReportTypeContactInformation, nil
	case "dropdown":
		return SurveyQuestionReportTypeDropdown, nil
	}
	var t SurveyQuestionReportType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (s SurveyQuestionReportType) Ptr() *SurveyQuestionReportType {
	return &s
}

// The report for a survey.
var (
	getSurveyReportingResponseFieldCreatedAt      = big.NewInt(1 << 0)
	getSurveyReportingResponseFieldID             = big.NewInt(1 << 1)
	getSurveyReportingResponseFieldListID         = big.NewInt(1 << 2)
	getSurveyReportingResponseFieldListName       = big.NewInt(1 << 3)
	getSurveyReportingResponseFieldPublishedAt    = big.NewInt(1 << 4)
	getSurveyReportingResponseFieldStatus         = big.NewInt(1 << 5)
	getSurveyReportingResponseFieldTitle          = big.NewInt(1 << 6)
	getSurveyReportingResponseFieldTotalResponses = big.NewInt(1 << 7)
	getSurveyReportingResponseFieldUpdatedAt      = big.NewInt(1 << 8)
	getSurveyReportingResponseFieldURL            = big.NewInt(1 << 9)
	getSurveyReportingResponseFieldWebID          = big.NewInt(1 << 10)
)

type GetSurveyReportingResponse struct {
	// The date and time the survey was created in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty" url:"created_at,omitempty"`
	// A string that uniquely identifies this survey.
	ID *string `json:"id,omitempty" url:"id,omitempty"`
	// The ID of the list connected to this survey.
	ListID *string `json:"list_id,omitempty" url:"list_id,omitempty"`
	// The name of the list connected to this survey.
	ListName *string `json:"list_name,omitempty" url:"list_name,omitempty"`
	// The date and time the survey was published in ISO 8601 format.
	PublishedAt *time.Time `json:"published_at,omitempty" url:"published_at,omitempty"`
	// The survey's status.
	Status *GetSurveyReportingResponseStatus `json:"status,omitempty" url:"status,omitempty"`
	// The title of the survey.
	Title *string `json:"title,omitempty" url:"title,omitempty"`
	// The total number of responses to this survey.
	TotalResponses *int `json:"total_responses,omitempty" url:"total_responses,omitempty"`
	// The date and time the survey was last updated in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty" url:"updated_at,omitempty"`
	// The URL for the survey.
	URL *string `json:"url,omitempty" url:"url,omitempty"`
	// The ID used in the Mailchimp web application. View this survey report in your Mailchimp account at `https://{dc}.admin.mailchimp.com/lists/surveys/results?survey_id={web_id}`.
	WebID *int `json:"web_id,omitempty" url:"web_id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (g *GetSurveyReportingResponse) GetCreatedAt() *time.Time {
	if g == nil {
		return nil
	}
	return g.CreatedAt
}

func (g *GetSurveyReportingResponse) GetID() *string {
	if g == nil {
		return nil
	}
	return g.ID
}

func (g *GetSurveyReportingResponse) GetListID() *string {
	if g == nil {
		return nil
	}
	return g.ListID
}

func (g *GetSurveyReportingResponse) GetListName() *string {
	if g == nil {
		return nil
	}
	return g.ListName
}

func (g *GetSurveyReportingResponse) GetPublishedAt() *time.Time {
	if g == nil {
		return nil
	}
	return g.PublishedAt
}

func (g *GetSurveyReportingResponse) GetStatus() *GetSurveyReportingResponseStatus {
	if g == nil {
		return nil
	}
	return g.Status
}

func (g *GetSurveyReportingResponse) GetTitle() *string {
	if g == nil {
		return nil
	}
	return g.Title
}

func (g *GetSurveyReportingResponse) GetTotalResponses() *int {
	if g == nil {
		return nil
	}
	return g.TotalResponses
}

func (g *GetSurveyReportingResponse) GetUpdatedAt() *time.Time {
	if g == nil {
		return nil
	}
	return g.UpdatedAt
}

func (g *GetSurveyReportingResponse) GetURL() *string {
	if g == nil {
		return nil
	}
	return g.URL
}

func (g *GetSurveyReportingResponse) GetWebID() *int {
	if g == nil {
		return nil
	}
	return g.WebID
}

func (g *GetSurveyReportingResponse) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g *GetSurveyReportingResponse) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetCreatedAt sets the CreatedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyReportingResponse) SetCreatedAt(createdAt *time.Time) {
	g.CreatedAt = createdAt
	g.require(getSurveyReportingResponseFieldCreatedAt)
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyReportingResponse) SetID(id *string) {
	g.ID = id
	g.require(getSurveyReportingResponseFieldID)
}

// SetListID sets the ListID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyReportingResponse) SetListID(listID *string) {
	g.ListID = listID
	g.require(getSurveyReportingResponseFieldListID)
}

// SetListName sets the ListName field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyReportingResponse) SetListName(listName *string) {
	g.ListName = listName
	g.require(getSurveyReportingResponseFieldListName)
}

// SetPublishedAt sets the PublishedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyReportingResponse) SetPublishedAt(publishedAt *time.Time) {
	g.PublishedAt = publishedAt
	g.require(getSurveyReportingResponseFieldPublishedAt)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyReportingResponse) SetStatus(status *GetSurveyReportingResponseStatus) {
	g.Status = status
	g.require(getSurveyReportingResponseFieldStatus)
}

// SetTitle sets the Title field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyReportingResponse) SetTitle(title *string) {
	g.Title = title
	g.require(getSurveyReportingResponseFieldTitle)
}

// SetTotalResponses sets the TotalResponses field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyReportingResponse) SetTotalResponses(totalResponses *int) {
	g.TotalResponses = totalResponses
	g.require(getSurveyReportingResponseFieldTotalResponses)
}

// SetUpdatedAt sets the UpdatedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyReportingResponse) SetUpdatedAt(updatedAt *time.Time) {
	g.UpdatedAt = updatedAt
	g.require(getSurveyReportingResponseFieldUpdatedAt)
}

// SetURL sets the URL field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyReportingResponse) SetURL(url *string) {
	g.URL = url
	g.require(getSurveyReportingResponseFieldURL)
}

// SetWebID sets the WebID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyReportingResponse) SetWebID(webID *int) {
	g.WebID = webID
	g.require(getSurveyReportingResponseFieldWebID)
}

func (g *GetSurveyReportingResponse) UnmarshalJSON(data []byte) error {
	type embed GetSurveyReportingResponse
	var unmarshaler = struct {
		embed
		CreatedAt   *internal.DateTime `json:"created_at,omitempty"`
		PublishedAt *internal.DateTime `json:"published_at,omitempty"`
		UpdatedAt   *internal.DateTime `json:"updated_at,omitempty"`
	}{
		embed: embed(*g),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*g = GetSurveyReportingResponse(unmarshaler.embed)
	g.CreatedAt = unmarshaler.CreatedAt.TimePtr()
	g.PublishedAt = unmarshaler.PublishedAt.TimePtr()
	g.UpdatedAt = unmarshaler.UpdatedAt.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetSurveyReportingResponse) MarshalJSON() ([]byte, error) {
	type embed GetSurveyReportingResponse
	var marshaler = struct {
		embed
		CreatedAt   *internal.DateTime `json:"created_at,omitempty"`
		PublishedAt *internal.DateTime `json:"published_at,omitempty"`
		UpdatedAt   *internal.DateTime `json:"updated_at,omitempty"`
	}{
		embed:       embed(*g),
		CreatedAt:   internal.NewOptionalDateTime(g.CreatedAt),
		PublishedAt: internal.NewOptionalDateTime(g.PublishedAt),
		UpdatedAt:   internal.NewOptionalDateTime(g.UpdatedAt),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, g.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (g *GetSurveyReportingResponse) String() string {
	if g == nil {
		return "<nil>"
	}
	if len(g.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(g.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

// The survey's status.
type GetSurveyReportingResponseStatus string

const (
	GetSurveyReportingResponseStatusPublished   GetSurveyReportingResponseStatus = "published"
	GetSurveyReportingResponseStatusUnpublished GetSurveyReportingResponseStatus = "unpublished"
)

func NewGetSurveyReportingResponseStatusFromString(s string) (GetSurveyReportingResponseStatus, error) {
	switch s {
	case "published":
		return GetSurveyReportingResponseStatusPublished, nil
	case "unpublished":
		return GetSurveyReportingResponseStatusUnpublished, nil
	}
	var t GetSurveyReportingResponseStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (g GetSurveyReportingResponseStatus) Ptr() *GetSurveyReportingResponseStatus {
	return &g
}

// A single survey response.
var (
	getSurveyResponsReportingResponseFieldContact      = big.NewInt(1 << 0)
	getSurveyResponsReportingResponseFieldIsNewContact = big.NewInt(1 << 1)
	getSurveyResponsReportingResponseFieldResponseID   = big.NewInt(1 << 2)
	getSurveyResponsReportingResponseFieldResults      = big.NewInt(1 << 3)
	getSurveyResponsReportingResponseFieldSubmittedAt  = big.NewInt(1 << 4)
)

type GetSurveyResponsReportingResponse struct {
	// Information about the contact.
	Contact *GetSurveyResponsReportingResponseContact `json:"contact,omitempty" url:"contact,omitempty"`
	// If this contact was added to the Mailchimp audience via this survey.
	IsNewContact *bool `json:"is_new_contact,omitempty" url:"is_new_contact,omitempty"`
	// The ID for the survey response.
	ResponseID *string `json:"response_id,omitempty" url:"response_id,omitempty"`
	// The survey questions and the answers to those questions.
	Results []*GetSurveyResponsReportingResponseResultsItem `json:"results,omitempty" url:"results,omitempty"`
	// The date and time when the survey response was submitted in ISO 8601 format.
	SubmittedAt *time.Time `json:"submitted_at,omitempty" url:"submitted_at,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (g *GetSurveyResponsReportingResponse) GetContact() *GetSurveyResponsReportingResponseContact {
	if g == nil {
		return nil
	}
	return g.Contact
}

func (g *GetSurveyResponsReportingResponse) GetIsNewContact() *bool {
	if g == nil {
		return nil
	}
	return g.IsNewContact
}

func (g *GetSurveyResponsReportingResponse) GetResponseID() *string {
	if g == nil {
		return nil
	}
	return g.ResponseID
}

func (g *GetSurveyResponsReportingResponse) GetResults() []*GetSurveyResponsReportingResponseResultsItem {
	if g == nil {
		return nil
	}
	return g.Results
}

func (g *GetSurveyResponsReportingResponse) GetSubmittedAt() *time.Time {
	if g == nil {
		return nil
	}
	return g.SubmittedAt
}

func (g *GetSurveyResponsReportingResponse) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g *GetSurveyResponsReportingResponse) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetContact sets the Contact field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyResponsReportingResponse) SetContact(contact *GetSurveyResponsReportingResponseContact) {
	g.Contact = contact
	g.require(getSurveyResponsReportingResponseFieldContact)
}

// SetIsNewContact sets the IsNewContact field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyResponsReportingResponse) SetIsNewContact(isNewContact *bool) {
	g.IsNewContact = isNewContact
	g.require(getSurveyResponsReportingResponseFieldIsNewContact)
}

// SetResponseID sets the ResponseID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyResponsReportingResponse) SetResponseID(responseID *string) {
	g.ResponseID = responseID
	g.require(getSurveyResponsReportingResponseFieldResponseID)
}

// SetResults sets the Results field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyResponsReportingResponse) SetResults(results []*GetSurveyResponsReportingResponseResultsItem) {
	g.Results = results
	g.require(getSurveyResponsReportingResponseFieldResults)
}

// SetSubmittedAt sets the SubmittedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyResponsReportingResponse) SetSubmittedAt(submittedAt *time.Time) {
	g.SubmittedAt = submittedAt
	g.require(getSurveyResponsReportingResponseFieldSubmittedAt)
}

func (g *GetSurveyResponsReportingResponse) UnmarshalJSON(data []byte) error {
	type embed GetSurveyResponsReportingResponse
	var unmarshaler = struct {
		embed
		SubmittedAt *internal.DateTime `json:"submitted_at,omitempty"`
	}{
		embed: embed(*g),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*g = GetSurveyResponsReportingResponse(unmarshaler.embed)
	g.SubmittedAt = unmarshaler.SubmittedAt.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetSurveyResponsReportingResponse) MarshalJSON() ([]byte, error) {
	type embed GetSurveyResponsReportingResponse
	var marshaler = struct {
		embed
		SubmittedAt *internal.DateTime `json:"submitted_at,omitempty"`
	}{
		embed:       embed(*g),
		SubmittedAt: internal.NewOptionalDateTime(g.SubmittedAt),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, g.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (g *GetSurveyResponsReportingResponse) String() string {
	if g == nil {
		return "<nil>"
	}
	if len(g.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(g.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

// Information about the contact.
var (
	getSurveyResponsReportingResponseContactFieldAvatarURL                   = big.NewInt(1 << 0)
	getSurveyResponsReportingResponseContactFieldConsentsToOneToOneMessaging = big.NewInt(1 << 1)
	getSurveyResponsReportingResponseContactFieldContactID                   = big.NewInt(1 << 2)
	getSurveyResponsReportingResponseContactFieldEmail                       = big.NewInt(1 << 3)
	getSurveyResponsReportingResponseContactFieldEmailID                     = big.NewInt(1 << 4)
	getSurveyResponsReportingResponseContactFieldFullName                    = big.NewInt(1 << 5)
	getSurveyResponsReportingResponseContactFieldPhone                       = big.NewInt(1 << 6)
	getSurveyResponsReportingResponseContactFieldStatus                      = big.NewInt(1 << 7)
)

type GetSurveyResponsReportingResponseContact struct {
	// URL for the contact's avatar or profile image.
	AvatarURL *string `json:"avatar_url,omitempty" url:"avatar_url,omitempty"`
	// Indicates whether a contact consents to 1:1 messaging.
	ConsentsToOneToOneMessaging *bool `json:"consents_to_one_to_one_messaging,omitempty" url:"consents_to_one_to_one_messaging,omitempty"`
	// The ID of this contact.
	ContactID *string `json:"contact_id,omitempty" url:"contact_id,omitempty"`
	// The contact's email address.
	Email *string `json:"email,omitempty" url:"email,omitempty"`
	// The MD5 hash of the lowercase version of the list member's email address.
	EmailID *string `json:"email_id,omitempty" url:"email_id,omitempty"`
	// The contact's full name.
	FullName *string `json:"full_name,omitempty" url:"full_name,omitempty"`
	// The contact's sms phone number.
	Phone *string `json:"phone,omitempty" url:"phone,omitempty"`
	// The contact's current status.
	Status *GetSurveyResponsReportingResponseContactStatus `json:"status,omitempty" url:"status,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (g *GetSurveyResponsReportingResponseContact) GetAvatarURL() *string {
	if g == nil {
		return nil
	}
	return g.AvatarURL
}

func (g *GetSurveyResponsReportingResponseContact) GetConsentsToOneToOneMessaging() *bool {
	if g == nil {
		return nil
	}
	return g.ConsentsToOneToOneMessaging
}

func (g *GetSurveyResponsReportingResponseContact) GetContactID() *string {
	if g == nil {
		return nil
	}
	return g.ContactID
}

func (g *GetSurveyResponsReportingResponseContact) GetEmail() *string {
	if g == nil {
		return nil
	}
	return g.Email
}

func (g *GetSurveyResponsReportingResponseContact) GetEmailID() *string {
	if g == nil {
		return nil
	}
	return g.EmailID
}

func (g *GetSurveyResponsReportingResponseContact) GetFullName() *string {
	if g == nil {
		return nil
	}
	return g.FullName
}

func (g *GetSurveyResponsReportingResponseContact) GetPhone() *string {
	if g == nil {
		return nil
	}
	return g.Phone
}

func (g *GetSurveyResponsReportingResponseContact) GetStatus() *GetSurveyResponsReportingResponseContactStatus {
	if g == nil {
		return nil
	}
	return g.Status
}

func (g *GetSurveyResponsReportingResponseContact) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g *GetSurveyResponsReportingResponseContact) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetAvatarURL sets the AvatarURL field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyResponsReportingResponseContact) SetAvatarURL(avatarURL *string) {
	g.AvatarURL = avatarURL
	g.require(getSurveyResponsReportingResponseContactFieldAvatarURL)
}

// SetConsentsToOneToOneMessaging sets the ConsentsToOneToOneMessaging field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyResponsReportingResponseContact) SetConsentsToOneToOneMessaging(consentsToOneToOneMessaging *bool) {
	g.ConsentsToOneToOneMessaging = consentsToOneToOneMessaging
	g.require(getSurveyResponsReportingResponseContactFieldConsentsToOneToOneMessaging)
}

// SetContactID sets the ContactID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyResponsReportingResponseContact) SetContactID(contactID *string) {
	g.ContactID = contactID
	g.require(getSurveyResponsReportingResponseContactFieldContactID)
}

// SetEmail sets the Email field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyResponsReportingResponseContact) SetEmail(email *string) {
	g.Email = email
	g.require(getSurveyResponsReportingResponseContactFieldEmail)
}

// SetEmailID sets the EmailID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyResponsReportingResponseContact) SetEmailID(emailID *string) {
	g.EmailID = emailID
	g.require(getSurveyResponsReportingResponseContactFieldEmailID)
}

// SetFullName sets the FullName field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyResponsReportingResponseContact) SetFullName(fullName *string) {
	g.FullName = fullName
	g.require(getSurveyResponsReportingResponseContactFieldFullName)
}

// SetPhone sets the Phone field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyResponsReportingResponseContact) SetPhone(phone *string) {
	g.Phone = phone
	g.require(getSurveyResponsReportingResponseContactFieldPhone)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyResponsReportingResponseContact) SetStatus(status *GetSurveyResponsReportingResponseContactStatus) {
	g.Status = status
	g.require(getSurveyResponsReportingResponseContactFieldStatus)
}

func (g *GetSurveyResponsReportingResponseContact) UnmarshalJSON(data []byte) error {
	type unmarshaler GetSurveyResponsReportingResponseContact
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetSurveyResponsReportingResponseContact(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetSurveyResponsReportingResponseContact) MarshalJSON() ([]byte, error) {
	type embed GetSurveyResponsReportingResponseContact
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, g.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (g *GetSurveyResponsReportingResponseContact) String() string {
	if g == nil {
		return "<nil>"
	}
	if len(g.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(g.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

// The contact's current status.
type GetSurveyResponsReportingResponseContactStatus string

const (
	GetSurveyResponsReportingResponseContactStatusSubscribed    GetSurveyResponsReportingResponseContactStatus = "Subscribed"
	GetSurveyResponsReportingResponseContactStatusUnsubscribed  GetSurveyResponsReportingResponseContactStatus = "Unsubscribed"
	GetSurveyResponsReportingResponseContactStatusNonSubscribed GetSurveyResponsReportingResponseContactStatus = "Non-Subscribed"
	GetSurveyResponsReportingResponseContactStatusCleaned       GetSurveyResponsReportingResponseContactStatus = "Cleaned"
	GetSurveyResponsReportingResponseContactStatusArchived      GetSurveyResponsReportingResponseContactStatus = "Archived"
)

func NewGetSurveyResponsReportingResponseContactStatusFromString(s string) (GetSurveyResponsReportingResponseContactStatus, error) {
	switch s {
	case "Subscribed":
		return GetSurveyResponsReportingResponseContactStatusSubscribed, nil
	case "Unsubscribed":
		return GetSurveyResponsReportingResponseContactStatusUnsubscribed, nil
	case "Non-Subscribed":
		return GetSurveyResponsReportingResponseContactStatusNonSubscribed, nil
	case "Cleaned":
		return GetSurveyResponsReportingResponseContactStatusCleaned, nil
	case "Archived":
		return GetSurveyResponsReportingResponseContactStatusArchived, nil
	}
	var t GetSurveyResponsReportingResponseContactStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (g GetSurveyResponsReportingResponseContactStatus) Ptr() *GetSurveyResponsReportingResponseContactStatus {
	return &g
}

// A single question and the response to that question.
var (
	getSurveyResponsReportingResponseResultsItemFieldAnswer       = big.NewInt(1 << 0)
	getSurveyResponsReportingResponseResultsItemFieldQuery        = big.NewInt(1 << 1)
	getSurveyResponsReportingResponseResultsItemFieldQuestionID   = big.NewInt(1 << 2)
	getSurveyResponsReportingResponseResultsItemFieldQuestionType = big.NewInt(1 << 3)
)

type GetSurveyResponsReportingResponseResultsItem struct {
	// The answer to this survey question.
	Answer *string `json:"answer,omitempty" url:"answer,omitempty"`
	// The survey question.
	Query *string `json:"query,omitempty" url:"query,omitempty"`
	// The unique ID for this question.
	QuestionID *string `json:"question_id,omitempty" url:"question_id,omitempty"`
	// The type of question this is.
	QuestionType *GetSurveyResponsReportingResponseResultsItemQuestionType `json:"question_type,omitempty" url:"question_type,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (g *GetSurveyResponsReportingResponseResultsItem) GetAnswer() *string {
	if g == nil {
		return nil
	}
	return g.Answer
}

func (g *GetSurveyResponsReportingResponseResultsItem) GetQuery() *string {
	if g == nil {
		return nil
	}
	return g.Query
}

func (g *GetSurveyResponsReportingResponseResultsItem) GetQuestionID() *string {
	if g == nil {
		return nil
	}
	return g.QuestionID
}

func (g *GetSurveyResponsReportingResponseResultsItem) GetQuestionType() *GetSurveyResponsReportingResponseResultsItemQuestionType {
	if g == nil {
		return nil
	}
	return g.QuestionType
}

func (g *GetSurveyResponsReportingResponseResultsItem) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g *GetSurveyResponsReportingResponseResultsItem) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetAnswer sets the Answer field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyResponsReportingResponseResultsItem) SetAnswer(answer *string) {
	g.Answer = answer
	g.require(getSurveyResponsReportingResponseResultsItemFieldAnswer)
}

// SetQuery sets the Query field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyResponsReportingResponseResultsItem) SetQuery(query *string) {
	g.Query = query
	g.require(getSurveyResponsReportingResponseResultsItemFieldQuery)
}

// SetQuestionID sets the QuestionID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyResponsReportingResponseResultsItem) SetQuestionID(questionID *string) {
	g.QuestionID = questionID
	g.require(getSurveyResponsReportingResponseResultsItemFieldQuestionID)
}

// SetQuestionType sets the QuestionType field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSurveyResponsReportingResponseResultsItem) SetQuestionType(questionType *GetSurveyResponsReportingResponseResultsItemQuestionType) {
	g.QuestionType = questionType
	g.require(getSurveyResponsReportingResponseResultsItemFieldQuestionType)
}

func (g *GetSurveyResponsReportingResponseResultsItem) UnmarshalJSON(data []byte) error {
	type unmarshaler GetSurveyResponsReportingResponseResultsItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetSurveyResponsReportingResponseResultsItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetSurveyResponsReportingResponseResultsItem) MarshalJSON() ([]byte, error) {
	type embed GetSurveyResponsReportingResponseResultsItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, g.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (g *GetSurveyResponsReportingResponseResultsItem) String() string {
	if g == nil {
		return "<nil>"
	}
	if len(g.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(g.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

// The type of question this is.
type GetSurveyResponsReportingResponseResultsItemQuestionType string

const (
	GetSurveyResponsReportingResponseResultsItemQuestionTypePickOne            GetSurveyResponsReportingResponseResultsItemQuestionType = "pickOne"
	GetSurveyResponsReportingResponseResultsItemQuestionTypePickMany           GetSurveyResponsReportingResponseResultsItemQuestionType = "pickMany"
	GetSurveyResponsReportingResponseResultsItemQuestionTypeRange              GetSurveyResponsReportingResponseResultsItemQuestionType = "range"
	GetSurveyResponsReportingResponseResultsItemQuestionTypeText               GetSurveyResponsReportingResponseResultsItemQuestionType = "text"
	GetSurveyResponsReportingResponseResultsItemQuestionTypeEmail              GetSurveyResponsReportingResponseResultsItemQuestionType = "email"
	GetSurveyResponsReportingResponseResultsItemQuestionTypeContactInformation GetSurveyResponsReportingResponseResultsItemQuestionType = "contactInformation"
	GetSurveyResponsReportingResponseResultsItemQuestionTypeDropdown           GetSurveyResponsReportingResponseResultsItemQuestionType = "dropdown"
)

func NewGetSurveyResponsReportingResponseResultsItemQuestionTypeFromString(s string) (GetSurveyResponsReportingResponseResultsItemQuestionType, error) {
	switch s {
	case "pickOne":
		return GetSurveyResponsReportingResponseResultsItemQuestionTypePickOne, nil
	case "pickMany":
		return GetSurveyResponsReportingResponseResultsItemQuestionTypePickMany, nil
	case "range":
		return GetSurveyResponsReportingResponseResultsItemQuestionTypeRange, nil
	case "text":
		return GetSurveyResponsReportingResponseResultsItemQuestionTypeText, nil
	case "email":
		return GetSurveyResponsReportingResponseResultsItemQuestionTypeEmail, nil
	case "contactInformation":
		return GetSurveyResponsReportingResponseResultsItemQuestionTypeContactInformation, nil
	case "dropdown":
		return GetSurveyResponsReportingResponseResultsItemQuestionTypeDropdown, nil
	}
	var t GetSurveyResponsReportingResponseResultsItemQuestionType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (g GetSurveyResponsReportingResponseResultsItemQuestionType) Ptr() *GetSurveyResponsReportingResponseResultsItemQuestionType {
	return &g
}

type ListFacebookAdEcommerceProductActivityReportingRequestSortField string

const (
	ListFacebookAdEcommerceProductActivityReportingRequestSortFieldTitle          ListFacebookAdEcommerceProductActivityReportingRequestSortField = "title"
	ListFacebookAdEcommerceProductActivityReportingRequestSortFieldTotalRevenue   ListFacebookAdEcommerceProductActivityReportingRequestSortField = "total_revenue"
	ListFacebookAdEcommerceProductActivityReportingRequestSortFieldTotalPurchased ListFacebookAdEcommerceProductActivityReportingRequestSortField = "total_purchased"
)

func NewListFacebookAdEcommerceProductActivityReportingRequestSortFieldFromString(s string) (ListFacebookAdEcommerceProductActivityReportingRequestSortField, error) {
	switch s {
	case "title":
		return ListFacebookAdEcommerceProductActivityReportingRequestSortFieldTitle, nil
	case "total_revenue":
		return ListFacebookAdEcommerceProductActivityReportingRequestSortFieldTotalRevenue, nil
	case "total_purchased":
		return ListFacebookAdEcommerceProductActivityReportingRequestSortFieldTotalPurchased, nil
	}
	var t ListFacebookAdEcommerceProductActivityReportingRequestSortField
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListFacebookAdEcommerceProductActivityReportingRequestSortField) Ptr() *ListFacebookAdEcommerceProductActivityReportingRequestSortField {
	return &l
}

// A collection of ecommerce products.
var (
	listFacebookAdEcommerceProductActivityReportingResponseFieldLinks      = big.NewInt(1 << 0)
	listFacebookAdEcommerceProductActivityReportingResponseFieldProducts   = big.NewInt(1 << 1)
	listFacebookAdEcommerceProductActivityReportingResponseFieldTotalItems = big.NewInt(1 << 2)
)

type ListFacebookAdEcommerceProductActivityReportingResponse struct {
	// A list of link types and descriptions for the API schema documents.
	Links    []*ListFacebookAdEcommerceProductActivityReportingResponseLinksItem    `json:"_links,omitempty" url:"_links,omitempty"`
	Products []*ListFacebookAdEcommerceProductActivityReportingResponseProductsItem `json:"products,omitempty" url:"products,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int `json:"total_items,omitempty" url:"total_items,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponse) GetLinks() []*ListFacebookAdEcommerceProductActivityReportingResponseLinksItem {
	if l == nil {
		return nil
	}
	return l.Links
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponse) GetProducts() []*ListFacebookAdEcommerceProductActivityReportingResponseProductsItem {
	if l == nil {
		return nil
	}
	return l.Products
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponse) GetTotalItems() *int {
	if l == nil {
		return nil
	}
	return l.TotalItems
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponse) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponse) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetLinks sets the Links field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdEcommerceProductActivityReportingResponse) SetLinks(links []*ListFacebookAdEcommerceProductActivityReportingResponseLinksItem) {
	l.Links = links
	l.require(listFacebookAdEcommerceProductActivityReportingResponseFieldLinks)
}

// SetProducts sets the Products field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdEcommerceProductActivityReportingResponse) SetProducts(products []*ListFacebookAdEcommerceProductActivityReportingResponseProductsItem) {
	l.Products = products
	l.require(listFacebookAdEcommerceProductActivityReportingResponseFieldProducts)
}

// SetTotalItems sets the TotalItems field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdEcommerceProductActivityReportingResponse) SetTotalItems(totalItems *int) {
	l.TotalItems = totalItems
	l.require(listFacebookAdEcommerceProductActivityReportingResponseFieldTotalItems)
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListFacebookAdEcommerceProductActivityReportingResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListFacebookAdEcommerceProductActivityReportingResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponse) MarshalJSON() ([]byte, error) {
	type embed ListFacebookAdEcommerceProductActivityReportingResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponse) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// This object represents a link from the resource where it is found to another resource or action that may be performed.
var (
	listFacebookAdEcommerceProductActivityReportingResponseLinksItemFieldHref         = big.NewInt(1 << 0)
	listFacebookAdEcommerceProductActivityReportingResponseLinksItemFieldMethod       = big.NewInt(1 << 1)
	listFacebookAdEcommerceProductActivityReportingResponseLinksItemFieldRel          = big.NewInt(1 << 2)
	listFacebookAdEcommerceProductActivityReportingResponseLinksItemFieldSchema       = big.NewInt(1 << 3)
	listFacebookAdEcommerceProductActivityReportingResponseLinksItemFieldTargetSchema = big.NewInt(1 << 4)
)

type ListFacebookAdEcommerceProductActivityReportingResponseLinksItem struct {
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href *string `json:"href,omitempty" url:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method *ListFacebookAdEcommerceProductActivityReportingResponseLinksItemMethod `json:"method,omitempty" url:"method,omitempty"`
	// As with an HTML 'rel' attribute, this describes the type of link.
	Rel *string `json:"rel,omitempty" url:"rel,omitempty"`
	// For HTTP methods that can receive bodies (POST and PUT), this is a URL representing the schema that the body should conform to.
	Schema *string `json:"schema,omitempty" url:"schema,omitempty"`
	// For GETs, this is a URL representing the schema that the response should conform to.
	TargetSchema *string `json:"targetSchema,omitempty" url:"targetSchema,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponseLinksItem) GetHref() *string {
	if l == nil {
		return nil
	}
	return l.Href
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponseLinksItem) GetMethod() *ListFacebookAdEcommerceProductActivityReportingResponseLinksItemMethod {
	if l == nil {
		return nil
	}
	return l.Method
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponseLinksItem) GetRel() *string {
	if l == nil {
		return nil
	}
	return l.Rel
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponseLinksItem) GetSchema() *string {
	if l == nil {
		return nil
	}
	return l.Schema
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponseLinksItem) GetTargetSchema() *string {
	if l == nil {
		return nil
	}
	return l.TargetSchema
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponseLinksItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponseLinksItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetHref sets the Href field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdEcommerceProductActivityReportingResponseLinksItem) SetHref(href *string) {
	l.Href = href
	l.require(listFacebookAdEcommerceProductActivityReportingResponseLinksItemFieldHref)
}

// SetMethod sets the Method field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdEcommerceProductActivityReportingResponseLinksItem) SetMethod(method *ListFacebookAdEcommerceProductActivityReportingResponseLinksItemMethod) {
	l.Method = method
	l.require(listFacebookAdEcommerceProductActivityReportingResponseLinksItemFieldMethod)
}

// SetRel sets the Rel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdEcommerceProductActivityReportingResponseLinksItem) SetRel(rel *string) {
	l.Rel = rel
	l.require(listFacebookAdEcommerceProductActivityReportingResponseLinksItemFieldRel)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdEcommerceProductActivityReportingResponseLinksItem) SetSchema(schema *string) {
	l.Schema = schema
	l.require(listFacebookAdEcommerceProductActivityReportingResponseLinksItemFieldSchema)
}

// SetTargetSchema sets the TargetSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdEcommerceProductActivityReportingResponseLinksItem) SetTargetSchema(targetSchema *string) {
	l.TargetSchema = targetSchema
	l.require(listFacebookAdEcommerceProductActivityReportingResponseLinksItemFieldTargetSchema)
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponseLinksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler ListFacebookAdEcommerceProductActivityReportingResponseLinksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListFacebookAdEcommerceProductActivityReportingResponseLinksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponseLinksItem) MarshalJSON() ([]byte, error) {
	type embed ListFacebookAdEcommerceProductActivityReportingResponseLinksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponseLinksItem) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// The HTTP method that should be used when accessing the URL defined in 'href'.
type ListFacebookAdEcommerceProductActivityReportingResponseLinksItemMethod string

const (
	ListFacebookAdEcommerceProductActivityReportingResponseLinksItemMethodGet     ListFacebookAdEcommerceProductActivityReportingResponseLinksItemMethod = "GET"
	ListFacebookAdEcommerceProductActivityReportingResponseLinksItemMethodPost    ListFacebookAdEcommerceProductActivityReportingResponseLinksItemMethod = "POST"
	ListFacebookAdEcommerceProductActivityReportingResponseLinksItemMethodPut     ListFacebookAdEcommerceProductActivityReportingResponseLinksItemMethod = "PUT"
	ListFacebookAdEcommerceProductActivityReportingResponseLinksItemMethodPatch   ListFacebookAdEcommerceProductActivityReportingResponseLinksItemMethod = "PATCH"
	ListFacebookAdEcommerceProductActivityReportingResponseLinksItemMethodDelete  ListFacebookAdEcommerceProductActivityReportingResponseLinksItemMethod = "DELETE"
	ListFacebookAdEcommerceProductActivityReportingResponseLinksItemMethodOptions ListFacebookAdEcommerceProductActivityReportingResponseLinksItemMethod = "OPTIONS"
	ListFacebookAdEcommerceProductActivityReportingResponseLinksItemMethodHead    ListFacebookAdEcommerceProductActivityReportingResponseLinksItemMethod = "HEAD"
)

func NewListFacebookAdEcommerceProductActivityReportingResponseLinksItemMethodFromString(s string) (ListFacebookAdEcommerceProductActivityReportingResponseLinksItemMethod, error) {
	switch s {
	case "GET":
		return ListFacebookAdEcommerceProductActivityReportingResponseLinksItemMethodGet, nil
	case "POST":
		return ListFacebookAdEcommerceProductActivityReportingResponseLinksItemMethodPost, nil
	case "PUT":
		return ListFacebookAdEcommerceProductActivityReportingResponseLinksItemMethodPut, nil
	case "PATCH":
		return ListFacebookAdEcommerceProductActivityReportingResponseLinksItemMethodPatch, nil
	case "DELETE":
		return ListFacebookAdEcommerceProductActivityReportingResponseLinksItemMethodDelete, nil
	case "OPTIONS":
		return ListFacebookAdEcommerceProductActivityReportingResponseLinksItemMethodOptions, nil
	case "HEAD":
		return ListFacebookAdEcommerceProductActivityReportingResponseLinksItemMethodHead, nil
	}
	var t ListFacebookAdEcommerceProductActivityReportingResponseLinksItemMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListFacebookAdEcommerceProductActivityReportingResponseLinksItemMethod) Ptr() *ListFacebookAdEcommerceProductActivityReportingResponseLinksItemMethod {
	return &l
}

var (
	listFacebookAdEcommerceProductActivityReportingResponseProductsItemFieldCurrencyCode            = big.NewInt(1 << 0)
	listFacebookAdEcommerceProductActivityReportingResponseProductsItemFieldImageURL                = big.NewInt(1 << 1)
	listFacebookAdEcommerceProductActivityReportingResponseProductsItemFieldRecommendationPurchased = big.NewInt(1 << 2)
	listFacebookAdEcommerceProductActivityReportingResponseProductsItemFieldRecommendationTotal     = big.NewInt(1 << 3)
	listFacebookAdEcommerceProductActivityReportingResponseProductsItemFieldSku                     = big.NewInt(1 << 4)
	listFacebookAdEcommerceProductActivityReportingResponseProductsItemFieldTitle                   = big.NewInt(1 << 5)
	listFacebookAdEcommerceProductActivityReportingResponseProductsItemFieldTotalPurchased          = big.NewInt(1 << 6)
	listFacebookAdEcommerceProductActivityReportingResponseProductsItemFieldTotalRevenue            = big.NewInt(1 << 7)
)

type ListFacebookAdEcommerceProductActivityReportingResponseProductsItem struct {
	CurrencyCode            *string  `json:"currency_code,omitempty" url:"currency_code,omitempty"`
	ImageURL                *string  `json:"image_url,omitempty" url:"image_url,omitempty"`
	RecommendationPurchased *int     `json:"recommendation_purchased,omitempty" url:"recommendation_purchased,omitempty"`
	RecommendationTotal     *int     `json:"recommendation_total,omitempty" url:"recommendation_total,omitempty"`
	Sku                     *string  `json:"sku,omitempty" url:"sku,omitempty"`
	Title                   *string  `json:"title,omitempty" url:"title,omitempty"`
	TotalPurchased          *float64 `json:"total_purchased,omitempty" url:"total_purchased,omitempty"`
	TotalRevenue            *float64 `json:"total_revenue,omitempty" url:"total_revenue,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponseProductsItem) GetCurrencyCode() *string {
	if l == nil {
		return nil
	}
	return l.CurrencyCode
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponseProductsItem) GetImageURL() *string {
	if l == nil {
		return nil
	}
	return l.ImageURL
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponseProductsItem) GetRecommendationPurchased() *int {
	if l == nil {
		return nil
	}
	return l.RecommendationPurchased
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponseProductsItem) GetRecommendationTotal() *int {
	if l == nil {
		return nil
	}
	return l.RecommendationTotal
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponseProductsItem) GetSku() *string {
	if l == nil {
		return nil
	}
	return l.Sku
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponseProductsItem) GetTitle() *string {
	if l == nil {
		return nil
	}
	return l.Title
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponseProductsItem) GetTotalPurchased() *float64 {
	if l == nil {
		return nil
	}
	return l.TotalPurchased
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponseProductsItem) GetTotalRevenue() *float64 {
	if l == nil {
		return nil
	}
	return l.TotalRevenue
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponseProductsItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponseProductsItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetCurrencyCode sets the CurrencyCode field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdEcommerceProductActivityReportingResponseProductsItem) SetCurrencyCode(currencyCode *string) {
	l.CurrencyCode = currencyCode
	l.require(listFacebookAdEcommerceProductActivityReportingResponseProductsItemFieldCurrencyCode)
}

// SetImageURL sets the ImageURL field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdEcommerceProductActivityReportingResponseProductsItem) SetImageURL(imageURL *string) {
	l.ImageURL = imageURL
	l.require(listFacebookAdEcommerceProductActivityReportingResponseProductsItemFieldImageURL)
}

// SetRecommendationPurchased sets the RecommendationPurchased field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdEcommerceProductActivityReportingResponseProductsItem) SetRecommendationPurchased(recommendationPurchased *int) {
	l.RecommendationPurchased = recommendationPurchased
	l.require(listFacebookAdEcommerceProductActivityReportingResponseProductsItemFieldRecommendationPurchased)
}

// SetRecommendationTotal sets the RecommendationTotal field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdEcommerceProductActivityReportingResponseProductsItem) SetRecommendationTotal(recommendationTotal *int) {
	l.RecommendationTotal = recommendationTotal
	l.require(listFacebookAdEcommerceProductActivityReportingResponseProductsItemFieldRecommendationTotal)
}

// SetSku sets the Sku field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdEcommerceProductActivityReportingResponseProductsItem) SetSku(sku *string) {
	l.Sku = sku
	l.require(listFacebookAdEcommerceProductActivityReportingResponseProductsItemFieldSku)
}

// SetTitle sets the Title field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdEcommerceProductActivityReportingResponseProductsItem) SetTitle(title *string) {
	l.Title = title
	l.require(listFacebookAdEcommerceProductActivityReportingResponseProductsItemFieldTitle)
}

// SetTotalPurchased sets the TotalPurchased field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdEcommerceProductActivityReportingResponseProductsItem) SetTotalPurchased(totalPurchased *float64) {
	l.TotalPurchased = totalPurchased
	l.require(listFacebookAdEcommerceProductActivityReportingResponseProductsItemFieldTotalPurchased)
}

// SetTotalRevenue sets the TotalRevenue field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdEcommerceProductActivityReportingResponseProductsItem) SetTotalRevenue(totalRevenue *float64) {
	l.TotalRevenue = totalRevenue
	l.require(listFacebookAdEcommerceProductActivityReportingResponseProductsItemFieldTotalRevenue)
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponseProductsItem) UnmarshalJSON(data []byte) error {
	type unmarshaler ListFacebookAdEcommerceProductActivityReportingResponseProductsItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListFacebookAdEcommerceProductActivityReportingResponseProductsItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponseProductsItem) MarshalJSON() ([]byte, error) {
	type embed ListFacebookAdEcommerceProductActivityReportingResponseProductsItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListFacebookAdEcommerceProductActivityReportingResponseProductsItem) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

type ListFacebookAdsReportingRequestSortDir string

const (
	ListFacebookAdsReportingRequestSortDirAsc  ListFacebookAdsReportingRequestSortDir = "ASC"
	ListFacebookAdsReportingRequestSortDirDesc ListFacebookAdsReportingRequestSortDir = "DESC"
)

func NewListFacebookAdsReportingRequestSortDirFromString(s string) (ListFacebookAdsReportingRequestSortDir, error) {
	switch s {
	case "ASC":
		return ListFacebookAdsReportingRequestSortDirAsc, nil
	case "DESC":
		return ListFacebookAdsReportingRequestSortDirDesc, nil
	}
	var t ListFacebookAdsReportingRequestSortDir
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListFacebookAdsReportingRequestSortDir) Ptr() *ListFacebookAdsReportingRequestSortDir {
	return &l
}

type ListFacebookAdsReportingRequestSortField string

const (
	ListFacebookAdsReportingRequestSortFieldCreatedAt ListFacebookAdsReportingRequestSortField = "created_at"
	ListFacebookAdsReportingRequestSortFieldUpdatedAt ListFacebookAdsReportingRequestSortField = "updated_at"
	ListFacebookAdsReportingRequestSortFieldEndTime   ListFacebookAdsReportingRequestSortField = "end_time"
)

func NewListFacebookAdsReportingRequestSortFieldFromString(s string) (ListFacebookAdsReportingRequestSortField, error) {
	switch s {
	case "created_at":
		return ListFacebookAdsReportingRequestSortFieldCreatedAt, nil
	case "updated_at":
		return ListFacebookAdsReportingRequestSortFieldUpdatedAt, nil
	case "end_time":
		return ListFacebookAdsReportingRequestSortFieldEndTime, nil
	}
	var t ListFacebookAdsReportingRequestSortField
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListFacebookAdsReportingRequestSortField) Ptr() *ListFacebookAdsReportingRequestSortField {
	return &l
}

// A collection of Facebook ads.
var (
	listFacebookAdsReportingResponseFieldLinks       = big.NewInt(1 << 0)
	listFacebookAdsReportingResponseFieldFacebookAds = big.NewInt(1 << 1)
	listFacebookAdsReportingResponseFieldTotalItems  = big.NewInt(1 << 2)
)

type ListFacebookAdsReportingResponse struct {
	// A list of link types and descriptions for the API schema documents.
	Links       []*ListFacebookAdsReportingResponseLinksItem `json:"_links,omitempty" url:"_links,omitempty"`
	FacebookAds []*ReportingFacebookAd                       `json:"facebook_ads,omitempty" url:"facebook_ads,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int `json:"total_items,omitempty" url:"total_items,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListFacebookAdsReportingResponse) GetLinks() []*ListFacebookAdsReportingResponseLinksItem {
	if l == nil {
		return nil
	}
	return l.Links
}

func (l *ListFacebookAdsReportingResponse) GetFacebookAds() []*ReportingFacebookAd {
	if l == nil {
		return nil
	}
	return l.FacebookAds
}

func (l *ListFacebookAdsReportingResponse) GetTotalItems() *int {
	if l == nil {
		return nil
	}
	return l.TotalItems
}

func (l *ListFacebookAdsReportingResponse) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListFacebookAdsReportingResponse) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetLinks sets the Links field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdsReportingResponse) SetLinks(links []*ListFacebookAdsReportingResponseLinksItem) {
	l.Links = links
	l.require(listFacebookAdsReportingResponseFieldLinks)
}

// SetFacebookAds sets the FacebookAds field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdsReportingResponse) SetFacebookAds(facebookAds []*ReportingFacebookAd) {
	l.FacebookAds = facebookAds
	l.require(listFacebookAdsReportingResponseFieldFacebookAds)
}

// SetTotalItems sets the TotalItems field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdsReportingResponse) SetTotalItems(totalItems *int) {
	l.TotalItems = totalItems
	l.require(listFacebookAdsReportingResponseFieldTotalItems)
}

func (l *ListFacebookAdsReportingResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListFacebookAdsReportingResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListFacebookAdsReportingResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListFacebookAdsReportingResponse) MarshalJSON() ([]byte, error) {
	type embed ListFacebookAdsReportingResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListFacebookAdsReportingResponse) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// This object represents a link from the resource where it is found to another resource or action that may be performed.
var (
	listFacebookAdsReportingResponseLinksItemFieldHref         = big.NewInt(1 << 0)
	listFacebookAdsReportingResponseLinksItemFieldMethod       = big.NewInt(1 << 1)
	listFacebookAdsReportingResponseLinksItemFieldRel          = big.NewInt(1 << 2)
	listFacebookAdsReportingResponseLinksItemFieldSchema       = big.NewInt(1 << 3)
	listFacebookAdsReportingResponseLinksItemFieldTargetSchema = big.NewInt(1 << 4)
)

type ListFacebookAdsReportingResponseLinksItem struct {
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href *string `json:"href,omitempty" url:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method *ListFacebookAdsReportingResponseLinksItemMethod `json:"method,omitempty" url:"method,omitempty"`
	// As with an HTML 'rel' attribute, this describes the type of link.
	Rel *string `json:"rel,omitempty" url:"rel,omitempty"`
	// For HTTP methods that can receive bodies (POST and PUT), this is a URL representing the schema that the body should conform to.
	Schema *string `json:"schema,omitempty" url:"schema,omitempty"`
	// For GETs, this is a URL representing the schema that the response should conform to.
	TargetSchema *string `json:"targetSchema,omitempty" url:"targetSchema,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListFacebookAdsReportingResponseLinksItem) GetHref() *string {
	if l == nil {
		return nil
	}
	return l.Href
}

func (l *ListFacebookAdsReportingResponseLinksItem) GetMethod() *ListFacebookAdsReportingResponseLinksItemMethod {
	if l == nil {
		return nil
	}
	return l.Method
}

func (l *ListFacebookAdsReportingResponseLinksItem) GetRel() *string {
	if l == nil {
		return nil
	}
	return l.Rel
}

func (l *ListFacebookAdsReportingResponseLinksItem) GetSchema() *string {
	if l == nil {
		return nil
	}
	return l.Schema
}

func (l *ListFacebookAdsReportingResponseLinksItem) GetTargetSchema() *string {
	if l == nil {
		return nil
	}
	return l.TargetSchema
}

func (l *ListFacebookAdsReportingResponseLinksItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListFacebookAdsReportingResponseLinksItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetHref sets the Href field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdsReportingResponseLinksItem) SetHref(href *string) {
	l.Href = href
	l.require(listFacebookAdsReportingResponseLinksItemFieldHref)
}

// SetMethod sets the Method field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdsReportingResponseLinksItem) SetMethod(method *ListFacebookAdsReportingResponseLinksItemMethod) {
	l.Method = method
	l.require(listFacebookAdsReportingResponseLinksItemFieldMethod)
}

// SetRel sets the Rel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdsReportingResponseLinksItem) SetRel(rel *string) {
	l.Rel = rel
	l.require(listFacebookAdsReportingResponseLinksItemFieldRel)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdsReportingResponseLinksItem) SetSchema(schema *string) {
	l.Schema = schema
	l.require(listFacebookAdsReportingResponseLinksItemFieldSchema)
}

// SetTargetSchema sets the TargetSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFacebookAdsReportingResponseLinksItem) SetTargetSchema(targetSchema *string) {
	l.TargetSchema = targetSchema
	l.require(listFacebookAdsReportingResponseLinksItemFieldTargetSchema)
}

func (l *ListFacebookAdsReportingResponseLinksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler ListFacebookAdsReportingResponseLinksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListFacebookAdsReportingResponseLinksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListFacebookAdsReportingResponseLinksItem) MarshalJSON() ([]byte, error) {
	type embed ListFacebookAdsReportingResponseLinksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListFacebookAdsReportingResponseLinksItem) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// The HTTP method that should be used when accessing the URL defined in 'href'.
type ListFacebookAdsReportingResponseLinksItemMethod string

const (
	ListFacebookAdsReportingResponseLinksItemMethodGet     ListFacebookAdsReportingResponseLinksItemMethod = "GET"
	ListFacebookAdsReportingResponseLinksItemMethodPost    ListFacebookAdsReportingResponseLinksItemMethod = "POST"
	ListFacebookAdsReportingResponseLinksItemMethodPut     ListFacebookAdsReportingResponseLinksItemMethod = "PUT"
	ListFacebookAdsReportingResponseLinksItemMethodPatch   ListFacebookAdsReportingResponseLinksItemMethod = "PATCH"
	ListFacebookAdsReportingResponseLinksItemMethodDelete  ListFacebookAdsReportingResponseLinksItemMethod = "DELETE"
	ListFacebookAdsReportingResponseLinksItemMethodOptions ListFacebookAdsReportingResponseLinksItemMethod = "OPTIONS"
	ListFacebookAdsReportingResponseLinksItemMethodHead    ListFacebookAdsReportingResponseLinksItemMethod = "HEAD"
)

func NewListFacebookAdsReportingResponseLinksItemMethodFromString(s string) (ListFacebookAdsReportingResponseLinksItemMethod, error) {
	switch s {
	case "GET":
		return ListFacebookAdsReportingResponseLinksItemMethodGet, nil
	case "POST":
		return ListFacebookAdsReportingResponseLinksItemMethodPost, nil
	case "PUT":
		return ListFacebookAdsReportingResponseLinksItemMethodPut, nil
	case "PATCH":
		return ListFacebookAdsReportingResponseLinksItemMethodPatch, nil
	case "DELETE":
		return ListFacebookAdsReportingResponseLinksItemMethodDelete, nil
	case "OPTIONS":
		return ListFacebookAdsReportingResponseLinksItemMethodOptions, nil
	case "HEAD":
		return ListFacebookAdsReportingResponseLinksItemMethodHead, nil
	}
	var t ListFacebookAdsReportingResponseLinksItemMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListFacebookAdsReportingResponseLinksItemMethod) Ptr() *ListFacebookAdsReportingResponseLinksItemMethod {
	return &l
}

// A collection of landing pages.
var (
	listLandingPagesReportingResponseFieldLinks        = big.NewInt(1 << 0)
	listLandingPagesReportingResponseFieldLandingPages = big.NewInt(1 << 1)
	listLandingPagesReportingResponseFieldTotalItems   = big.NewInt(1 << 2)
)

type ListLandingPagesReportingResponse struct {
	// A list of link types and descriptions for the API schema documents.
	Links        []*ListLandingPagesReportingResponseLinksItem `json:"_links,omitempty" url:"_links,omitempty"`
	LandingPages []*LandingPageReport                          `json:"landing_pages,omitempty" url:"landing_pages,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int `json:"total_items,omitempty" url:"total_items,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListLandingPagesReportingResponse) GetLinks() []*ListLandingPagesReportingResponseLinksItem {
	if l == nil {
		return nil
	}
	return l.Links
}

func (l *ListLandingPagesReportingResponse) GetLandingPages() []*LandingPageReport {
	if l == nil {
		return nil
	}
	return l.LandingPages
}

func (l *ListLandingPagesReportingResponse) GetTotalItems() *int {
	if l == nil {
		return nil
	}
	return l.TotalItems
}

func (l *ListLandingPagesReportingResponse) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListLandingPagesReportingResponse) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetLinks sets the Links field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListLandingPagesReportingResponse) SetLinks(links []*ListLandingPagesReportingResponseLinksItem) {
	l.Links = links
	l.require(listLandingPagesReportingResponseFieldLinks)
}

// SetLandingPages sets the LandingPages field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListLandingPagesReportingResponse) SetLandingPages(landingPages []*LandingPageReport) {
	l.LandingPages = landingPages
	l.require(listLandingPagesReportingResponseFieldLandingPages)
}

// SetTotalItems sets the TotalItems field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListLandingPagesReportingResponse) SetTotalItems(totalItems *int) {
	l.TotalItems = totalItems
	l.require(listLandingPagesReportingResponseFieldTotalItems)
}

func (l *ListLandingPagesReportingResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListLandingPagesReportingResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListLandingPagesReportingResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListLandingPagesReportingResponse) MarshalJSON() ([]byte, error) {
	type embed ListLandingPagesReportingResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListLandingPagesReportingResponse) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// This object represents a link from the resource where it is found to another resource or action that may be performed.
var (
	listLandingPagesReportingResponseLinksItemFieldHref         = big.NewInt(1 << 0)
	listLandingPagesReportingResponseLinksItemFieldMethod       = big.NewInt(1 << 1)
	listLandingPagesReportingResponseLinksItemFieldRel          = big.NewInt(1 << 2)
	listLandingPagesReportingResponseLinksItemFieldSchema       = big.NewInt(1 << 3)
	listLandingPagesReportingResponseLinksItemFieldTargetSchema = big.NewInt(1 << 4)
)

type ListLandingPagesReportingResponseLinksItem struct {
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href *string `json:"href,omitempty" url:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method *ListLandingPagesReportingResponseLinksItemMethod `json:"method,omitempty" url:"method,omitempty"`
	// As with an HTML 'rel' attribute, this describes the type of link.
	Rel *string `json:"rel,omitempty" url:"rel,omitempty"`
	// For HTTP methods that can receive bodies (POST and PUT), this is a URL representing the schema that the body should conform to.
	Schema *string `json:"schema,omitempty" url:"schema,omitempty"`
	// For GETs, this is a URL representing the schema that the response should conform to.
	TargetSchema *string `json:"targetSchema,omitempty" url:"targetSchema,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListLandingPagesReportingResponseLinksItem) GetHref() *string {
	if l == nil {
		return nil
	}
	return l.Href
}

func (l *ListLandingPagesReportingResponseLinksItem) GetMethod() *ListLandingPagesReportingResponseLinksItemMethod {
	if l == nil {
		return nil
	}
	return l.Method
}

func (l *ListLandingPagesReportingResponseLinksItem) GetRel() *string {
	if l == nil {
		return nil
	}
	return l.Rel
}

func (l *ListLandingPagesReportingResponseLinksItem) GetSchema() *string {
	if l == nil {
		return nil
	}
	return l.Schema
}

func (l *ListLandingPagesReportingResponseLinksItem) GetTargetSchema() *string {
	if l == nil {
		return nil
	}
	return l.TargetSchema
}

func (l *ListLandingPagesReportingResponseLinksItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListLandingPagesReportingResponseLinksItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetHref sets the Href field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListLandingPagesReportingResponseLinksItem) SetHref(href *string) {
	l.Href = href
	l.require(listLandingPagesReportingResponseLinksItemFieldHref)
}

// SetMethod sets the Method field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListLandingPagesReportingResponseLinksItem) SetMethod(method *ListLandingPagesReportingResponseLinksItemMethod) {
	l.Method = method
	l.require(listLandingPagesReportingResponseLinksItemFieldMethod)
}

// SetRel sets the Rel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListLandingPagesReportingResponseLinksItem) SetRel(rel *string) {
	l.Rel = rel
	l.require(listLandingPagesReportingResponseLinksItemFieldRel)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListLandingPagesReportingResponseLinksItem) SetSchema(schema *string) {
	l.Schema = schema
	l.require(listLandingPagesReportingResponseLinksItemFieldSchema)
}

// SetTargetSchema sets the TargetSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListLandingPagesReportingResponseLinksItem) SetTargetSchema(targetSchema *string) {
	l.TargetSchema = targetSchema
	l.require(listLandingPagesReportingResponseLinksItemFieldTargetSchema)
}

func (l *ListLandingPagesReportingResponseLinksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler ListLandingPagesReportingResponseLinksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListLandingPagesReportingResponseLinksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListLandingPagesReportingResponseLinksItem) MarshalJSON() ([]byte, error) {
	type embed ListLandingPagesReportingResponseLinksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListLandingPagesReportingResponseLinksItem) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// The HTTP method that should be used when accessing the URL defined in 'href'.
type ListLandingPagesReportingResponseLinksItemMethod string

const (
	ListLandingPagesReportingResponseLinksItemMethodGet     ListLandingPagesReportingResponseLinksItemMethod = "GET"
	ListLandingPagesReportingResponseLinksItemMethodPost    ListLandingPagesReportingResponseLinksItemMethod = "POST"
	ListLandingPagesReportingResponseLinksItemMethodPut     ListLandingPagesReportingResponseLinksItemMethod = "PUT"
	ListLandingPagesReportingResponseLinksItemMethodPatch   ListLandingPagesReportingResponseLinksItemMethod = "PATCH"
	ListLandingPagesReportingResponseLinksItemMethodDelete  ListLandingPagesReportingResponseLinksItemMethod = "DELETE"
	ListLandingPagesReportingResponseLinksItemMethodOptions ListLandingPagesReportingResponseLinksItemMethod = "OPTIONS"
	ListLandingPagesReportingResponseLinksItemMethodHead    ListLandingPagesReportingResponseLinksItemMethod = "HEAD"
)

func NewListLandingPagesReportingResponseLinksItemMethodFromString(s string) (ListLandingPagesReportingResponseLinksItemMethod, error) {
	switch s {
	case "GET":
		return ListLandingPagesReportingResponseLinksItemMethodGet, nil
	case "POST":
		return ListLandingPagesReportingResponseLinksItemMethodPost, nil
	case "PUT":
		return ListLandingPagesReportingResponseLinksItemMethodPut, nil
	case "PATCH":
		return ListLandingPagesReportingResponseLinksItemMethodPatch, nil
	case "DELETE":
		return ListLandingPagesReportingResponseLinksItemMethodDelete, nil
	case "OPTIONS":
		return ListLandingPagesReportingResponseLinksItemMethodOptions, nil
	case "HEAD":
		return ListLandingPagesReportingResponseLinksItemMethodHead, nil
	}
	var t ListLandingPagesReportingResponseLinksItemMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListLandingPagesReportingResponseLinksItemMethod) Ptr() *ListLandingPagesReportingResponseLinksItemMethod {
	return &l
}

// This object represents a link from the resource where it is found to another resource or action that may be performed.
var (
	listReportingResponseItemFieldHref         = big.NewInt(1 << 0)
	listReportingResponseItemFieldMethod       = big.NewInt(1 << 1)
	listReportingResponseItemFieldRel          = big.NewInt(1 << 2)
	listReportingResponseItemFieldSchema       = big.NewInt(1 << 3)
	listReportingResponseItemFieldTargetSchema = big.NewInt(1 << 4)
)

type ListReportingResponseItem struct {
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href *string `json:"href,omitempty" url:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method *ListReportingResponseItemMethod `json:"method,omitempty" url:"method,omitempty"`
	// As with an HTML 'rel' attribute, this describes the type of link.
	Rel *string `json:"rel,omitempty" url:"rel,omitempty"`
	// For HTTP methods that can receive bodies (POST and PUT), this is a URL representing the schema that the body should conform to.
	Schema *string `json:"schema,omitempty" url:"schema,omitempty"`
	// For GETs, this is a URL representing the schema that the response should conform to.
	TargetSchema *string `json:"targetSchema,omitempty" url:"targetSchema,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListReportingResponseItem) GetHref() *string {
	if l == nil {
		return nil
	}
	return l.Href
}

func (l *ListReportingResponseItem) GetMethod() *ListReportingResponseItemMethod {
	if l == nil {
		return nil
	}
	return l.Method
}

func (l *ListReportingResponseItem) GetRel() *string {
	if l == nil {
		return nil
	}
	return l.Rel
}

func (l *ListReportingResponseItem) GetSchema() *string {
	if l == nil {
		return nil
	}
	return l.Schema
}

func (l *ListReportingResponseItem) GetTargetSchema() *string {
	if l == nil {
		return nil
	}
	return l.TargetSchema
}

func (l *ListReportingResponseItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListReportingResponseItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetHref sets the Href field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListReportingResponseItem) SetHref(href *string) {
	l.Href = href
	l.require(listReportingResponseItemFieldHref)
}

// SetMethod sets the Method field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListReportingResponseItem) SetMethod(method *ListReportingResponseItemMethod) {
	l.Method = method
	l.require(listReportingResponseItemFieldMethod)
}

// SetRel sets the Rel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListReportingResponseItem) SetRel(rel *string) {
	l.Rel = rel
	l.require(listReportingResponseItemFieldRel)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListReportingResponseItem) SetSchema(schema *string) {
	l.Schema = schema
	l.require(listReportingResponseItemFieldSchema)
}

// SetTargetSchema sets the TargetSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListReportingResponseItem) SetTargetSchema(targetSchema *string) {
	l.TargetSchema = targetSchema
	l.require(listReportingResponseItemFieldTargetSchema)
}

func (l *ListReportingResponseItem) UnmarshalJSON(data []byte) error {
	type unmarshaler ListReportingResponseItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListReportingResponseItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListReportingResponseItem) MarshalJSON() ([]byte, error) {
	type embed ListReportingResponseItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListReportingResponseItem) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// The HTTP method that should be used when accessing the URL defined in 'href'.
type ListReportingResponseItemMethod string

const (
	ListReportingResponseItemMethodGet     ListReportingResponseItemMethod = "GET"
	ListReportingResponseItemMethodPost    ListReportingResponseItemMethod = "POST"
	ListReportingResponseItemMethodPut     ListReportingResponseItemMethod = "PUT"
	ListReportingResponseItemMethodPatch   ListReportingResponseItemMethod = "PATCH"
	ListReportingResponseItemMethodDelete  ListReportingResponseItemMethod = "DELETE"
	ListReportingResponseItemMethodOptions ListReportingResponseItemMethod = "OPTIONS"
	ListReportingResponseItemMethodHead    ListReportingResponseItemMethod = "HEAD"
)

func NewListReportingResponseItemMethodFromString(s string) (ListReportingResponseItemMethod, error) {
	switch s {
	case "GET":
		return ListReportingResponseItemMethodGet, nil
	case "POST":
		return ListReportingResponseItemMethodPost, nil
	case "PUT":
		return ListReportingResponseItemMethodPut, nil
	case "PATCH":
		return ListReportingResponseItemMethodPatch, nil
	case "DELETE":
		return ListReportingResponseItemMethodDelete, nil
	case "OPTIONS":
		return ListReportingResponseItemMethodOptions, nil
	case "HEAD":
		return ListReportingResponseItemMethodHead, nil
	}
	var t ListReportingResponseItemMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListReportingResponseItemMethod) Ptr() *ListReportingResponseItemMethod {
	return &l
}

type ListSurveyQuestionAnswersReportingRequestRespondentFamiliarityIs string

const (
	ListSurveyQuestionAnswersReportingRequestRespondentFamiliarityIsNew     ListSurveyQuestionAnswersReportingRequestRespondentFamiliarityIs = "new"
	ListSurveyQuestionAnswersReportingRequestRespondentFamiliarityIsKnown   ListSurveyQuestionAnswersReportingRequestRespondentFamiliarityIs = "known"
	ListSurveyQuestionAnswersReportingRequestRespondentFamiliarityIsUnknown ListSurveyQuestionAnswersReportingRequestRespondentFamiliarityIs = "unknown"
)

func NewListSurveyQuestionAnswersReportingRequestRespondentFamiliarityIsFromString(s string) (ListSurveyQuestionAnswersReportingRequestRespondentFamiliarityIs, error) {
	switch s {
	case "new":
		return ListSurveyQuestionAnswersReportingRequestRespondentFamiliarityIsNew, nil
	case "known":
		return ListSurveyQuestionAnswersReportingRequestRespondentFamiliarityIsKnown, nil
	case "unknown":
		return ListSurveyQuestionAnswersReportingRequestRespondentFamiliarityIsUnknown, nil
	}
	var t ListSurveyQuestionAnswersReportingRequestRespondentFamiliarityIs
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListSurveyQuestionAnswersReportingRequestRespondentFamiliarityIs) Ptr() *ListSurveyQuestionAnswersReportingRequestRespondentFamiliarityIs {
	return &l
}

var (
	listSurveyQuestionAnswersReportingResponseFieldLinks      = big.NewInt(1 << 0)
	listSurveyQuestionAnswersReportingResponseFieldAnswers    = big.NewInt(1 << 1)
	listSurveyQuestionAnswersReportingResponseFieldTotalItems = big.NewInt(1 << 2)
)

type ListSurveyQuestionAnswersReportingResponse struct {
	// A list of link types and descriptions for the API schema documents.
	Links []*ListSurveyQuestionAnswersReportingResponseLinksItem `json:"_links,omitempty" url:"_links,omitempty"`
	// An array of answers for a question on the survey.
	Answers []*ListSurveyQuestionAnswersReportingResponseAnswersItem `json:"answers,omitempty" url:"answers,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int `json:"total_items,omitempty" url:"total_items,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListSurveyQuestionAnswersReportingResponse) GetLinks() []*ListSurveyQuestionAnswersReportingResponseLinksItem {
	if l == nil {
		return nil
	}
	return l.Links
}

func (l *ListSurveyQuestionAnswersReportingResponse) GetAnswers() []*ListSurveyQuestionAnswersReportingResponseAnswersItem {
	if l == nil {
		return nil
	}
	return l.Answers
}

func (l *ListSurveyQuestionAnswersReportingResponse) GetTotalItems() *int {
	if l == nil {
		return nil
	}
	return l.TotalItems
}

func (l *ListSurveyQuestionAnswersReportingResponse) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListSurveyQuestionAnswersReportingResponse) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetLinks sets the Links field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionAnswersReportingResponse) SetLinks(links []*ListSurveyQuestionAnswersReportingResponseLinksItem) {
	l.Links = links
	l.require(listSurveyQuestionAnswersReportingResponseFieldLinks)
}

// SetAnswers sets the Answers field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionAnswersReportingResponse) SetAnswers(answers []*ListSurveyQuestionAnswersReportingResponseAnswersItem) {
	l.Answers = answers
	l.require(listSurveyQuestionAnswersReportingResponseFieldAnswers)
}

// SetTotalItems sets the TotalItems field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionAnswersReportingResponse) SetTotalItems(totalItems *int) {
	l.TotalItems = totalItems
	l.require(listSurveyQuestionAnswersReportingResponseFieldTotalItems)
}

func (l *ListSurveyQuestionAnswersReportingResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListSurveyQuestionAnswersReportingResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListSurveyQuestionAnswersReportingResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListSurveyQuestionAnswersReportingResponse) MarshalJSON() ([]byte, error) {
	type embed ListSurveyQuestionAnswersReportingResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListSurveyQuestionAnswersReportingResponse) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// The details of a survey question's answer.
var (
	listSurveyQuestionAnswersReportingResponseAnswersItemFieldContact      = big.NewInt(1 << 0)
	listSurveyQuestionAnswersReportingResponseAnswersItemFieldID           = big.NewInt(1 << 1)
	listSurveyQuestionAnswersReportingResponseAnswersItemFieldIsNewContact = big.NewInt(1 << 2)
	listSurveyQuestionAnswersReportingResponseAnswersItemFieldResponseID   = big.NewInt(1 << 3)
	listSurveyQuestionAnswersReportingResponseAnswersItemFieldSubmittedAt  = big.NewInt(1 << 4)
	listSurveyQuestionAnswersReportingResponseAnswersItemFieldValue        = big.NewInt(1 << 5)
)

type ListSurveyQuestionAnswersReportingResponseAnswersItem struct {
	// Information about the contact.
	Contact *ListSurveyQuestionAnswersReportingResponseAnswersItemContact `json:"contact,omitempty" url:"contact,omitempty"`
	// The ID of the answer.
	ID *string `json:"id,omitempty" url:"id,omitempty"`
	// If this contact was added to the Mailchimp audience via this survey.
	IsNewContact *bool `json:"is_new_contact,omitempty" url:"is_new_contact,omitempty"`
	// The ID of the survey response.
	ResponseID *string `json:"response_id,omitempty" url:"response_id,omitempty"`
	// The date and time when the survey response was submitted in ISO 8601 format.
	SubmittedAt *time.Time `json:"submitted_at,omitempty" url:"submitted_at,omitempty"`
	// The raw text answer.
	Value *string `json:"value,omitempty" url:"value,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListSurveyQuestionAnswersReportingResponseAnswersItem) GetContact() *ListSurveyQuestionAnswersReportingResponseAnswersItemContact {
	if l == nil {
		return nil
	}
	return l.Contact
}

func (l *ListSurveyQuestionAnswersReportingResponseAnswersItem) GetID() *string {
	if l == nil {
		return nil
	}
	return l.ID
}

func (l *ListSurveyQuestionAnswersReportingResponseAnswersItem) GetIsNewContact() *bool {
	if l == nil {
		return nil
	}
	return l.IsNewContact
}

func (l *ListSurveyQuestionAnswersReportingResponseAnswersItem) GetResponseID() *string {
	if l == nil {
		return nil
	}
	return l.ResponseID
}

func (l *ListSurveyQuestionAnswersReportingResponseAnswersItem) GetSubmittedAt() *time.Time {
	if l == nil {
		return nil
	}
	return l.SubmittedAt
}

func (l *ListSurveyQuestionAnswersReportingResponseAnswersItem) GetValue() *string {
	if l == nil {
		return nil
	}
	return l.Value
}

func (l *ListSurveyQuestionAnswersReportingResponseAnswersItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListSurveyQuestionAnswersReportingResponseAnswersItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetContact sets the Contact field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionAnswersReportingResponseAnswersItem) SetContact(contact *ListSurveyQuestionAnswersReportingResponseAnswersItemContact) {
	l.Contact = contact
	l.require(listSurveyQuestionAnswersReportingResponseAnswersItemFieldContact)
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionAnswersReportingResponseAnswersItem) SetID(id *string) {
	l.ID = id
	l.require(listSurveyQuestionAnswersReportingResponseAnswersItemFieldID)
}

// SetIsNewContact sets the IsNewContact field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionAnswersReportingResponseAnswersItem) SetIsNewContact(isNewContact *bool) {
	l.IsNewContact = isNewContact
	l.require(listSurveyQuestionAnswersReportingResponseAnswersItemFieldIsNewContact)
}

// SetResponseID sets the ResponseID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionAnswersReportingResponseAnswersItem) SetResponseID(responseID *string) {
	l.ResponseID = responseID
	l.require(listSurveyQuestionAnswersReportingResponseAnswersItemFieldResponseID)
}

// SetSubmittedAt sets the SubmittedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionAnswersReportingResponseAnswersItem) SetSubmittedAt(submittedAt *time.Time) {
	l.SubmittedAt = submittedAt
	l.require(listSurveyQuestionAnswersReportingResponseAnswersItemFieldSubmittedAt)
}

// SetValue sets the Value field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionAnswersReportingResponseAnswersItem) SetValue(value *string) {
	l.Value = value
	l.require(listSurveyQuestionAnswersReportingResponseAnswersItemFieldValue)
}

func (l *ListSurveyQuestionAnswersReportingResponseAnswersItem) UnmarshalJSON(data []byte) error {
	type embed ListSurveyQuestionAnswersReportingResponseAnswersItem
	var unmarshaler = struct {
		embed
		SubmittedAt *internal.DateTime `json:"submitted_at,omitempty"`
	}{
		embed: embed(*l),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*l = ListSurveyQuestionAnswersReportingResponseAnswersItem(unmarshaler.embed)
	l.SubmittedAt = unmarshaler.SubmittedAt.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListSurveyQuestionAnswersReportingResponseAnswersItem) MarshalJSON() ([]byte, error) {
	type embed ListSurveyQuestionAnswersReportingResponseAnswersItem
	var marshaler = struct {
		embed
		SubmittedAt *internal.DateTime `json:"submitted_at,omitempty"`
	}{
		embed:       embed(*l),
		SubmittedAt: internal.NewOptionalDateTime(l.SubmittedAt),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListSurveyQuestionAnswersReportingResponseAnswersItem) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// Information about the contact.
var (
	listSurveyQuestionAnswersReportingResponseAnswersItemContactFieldAvatarURL                   = big.NewInt(1 << 0)
	listSurveyQuestionAnswersReportingResponseAnswersItemContactFieldConsentsToOneToOneMessaging = big.NewInt(1 << 1)
	listSurveyQuestionAnswersReportingResponseAnswersItemContactFieldContactID                   = big.NewInt(1 << 2)
	listSurveyQuestionAnswersReportingResponseAnswersItemContactFieldEmail                       = big.NewInt(1 << 3)
	listSurveyQuestionAnswersReportingResponseAnswersItemContactFieldEmailID                     = big.NewInt(1 << 4)
	listSurveyQuestionAnswersReportingResponseAnswersItemContactFieldFullName                    = big.NewInt(1 << 5)
	listSurveyQuestionAnswersReportingResponseAnswersItemContactFieldPhone                       = big.NewInt(1 << 6)
	listSurveyQuestionAnswersReportingResponseAnswersItemContactFieldStatus                      = big.NewInt(1 << 7)
)

type ListSurveyQuestionAnswersReportingResponseAnswersItemContact struct {
	// URL for the contact's avatar or profile image.
	AvatarURL *string `json:"avatar_url,omitempty" url:"avatar_url,omitempty"`
	// Indicates whether a contact consents to 1:1 messaging.
	ConsentsToOneToOneMessaging *bool `json:"consents_to_one_to_one_messaging,omitempty" url:"consents_to_one_to_one_messaging,omitempty"`
	// The ID of this contact.
	ContactID *string `json:"contact_id,omitempty" url:"contact_id,omitempty"`
	// The contact's email address.
	Email *string `json:"email,omitempty" url:"email,omitempty"`
	// The MD5 hash of the lowercase version of the list member's email address.
	EmailID *string `json:"email_id,omitempty" url:"email_id,omitempty"`
	// The contact's full name.
	FullName *string `json:"full_name,omitempty" url:"full_name,omitempty"`
	// The contact's sms phone number.
	Phone *string `json:"phone,omitempty" url:"phone,omitempty"`
	// The contact's current status.
	Status *ListSurveyQuestionAnswersReportingResponseAnswersItemContactStatus `json:"status,omitempty" url:"status,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListSurveyQuestionAnswersReportingResponseAnswersItemContact) GetAvatarURL() *string {
	if l == nil {
		return nil
	}
	return l.AvatarURL
}

func (l *ListSurveyQuestionAnswersReportingResponseAnswersItemContact) GetConsentsToOneToOneMessaging() *bool {
	if l == nil {
		return nil
	}
	return l.ConsentsToOneToOneMessaging
}

func (l *ListSurveyQuestionAnswersReportingResponseAnswersItemContact) GetContactID() *string {
	if l == nil {
		return nil
	}
	return l.ContactID
}

func (l *ListSurveyQuestionAnswersReportingResponseAnswersItemContact) GetEmail() *string {
	if l == nil {
		return nil
	}
	return l.Email
}

func (l *ListSurveyQuestionAnswersReportingResponseAnswersItemContact) GetEmailID() *string {
	if l == nil {
		return nil
	}
	return l.EmailID
}

func (l *ListSurveyQuestionAnswersReportingResponseAnswersItemContact) GetFullName() *string {
	if l == nil {
		return nil
	}
	return l.FullName
}

func (l *ListSurveyQuestionAnswersReportingResponseAnswersItemContact) GetPhone() *string {
	if l == nil {
		return nil
	}
	return l.Phone
}

func (l *ListSurveyQuestionAnswersReportingResponseAnswersItemContact) GetStatus() *ListSurveyQuestionAnswersReportingResponseAnswersItemContactStatus {
	if l == nil {
		return nil
	}
	return l.Status
}

func (l *ListSurveyQuestionAnswersReportingResponseAnswersItemContact) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListSurveyQuestionAnswersReportingResponseAnswersItemContact) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetAvatarURL sets the AvatarURL field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionAnswersReportingResponseAnswersItemContact) SetAvatarURL(avatarURL *string) {
	l.AvatarURL = avatarURL
	l.require(listSurveyQuestionAnswersReportingResponseAnswersItemContactFieldAvatarURL)
}

// SetConsentsToOneToOneMessaging sets the ConsentsToOneToOneMessaging field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionAnswersReportingResponseAnswersItemContact) SetConsentsToOneToOneMessaging(consentsToOneToOneMessaging *bool) {
	l.ConsentsToOneToOneMessaging = consentsToOneToOneMessaging
	l.require(listSurveyQuestionAnswersReportingResponseAnswersItemContactFieldConsentsToOneToOneMessaging)
}

// SetContactID sets the ContactID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionAnswersReportingResponseAnswersItemContact) SetContactID(contactID *string) {
	l.ContactID = contactID
	l.require(listSurveyQuestionAnswersReportingResponseAnswersItemContactFieldContactID)
}

// SetEmail sets the Email field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionAnswersReportingResponseAnswersItemContact) SetEmail(email *string) {
	l.Email = email
	l.require(listSurveyQuestionAnswersReportingResponseAnswersItemContactFieldEmail)
}

// SetEmailID sets the EmailID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionAnswersReportingResponseAnswersItemContact) SetEmailID(emailID *string) {
	l.EmailID = emailID
	l.require(listSurveyQuestionAnswersReportingResponseAnswersItemContactFieldEmailID)
}

// SetFullName sets the FullName field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionAnswersReportingResponseAnswersItemContact) SetFullName(fullName *string) {
	l.FullName = fullName
	l.require(listSurveyQuestionAnswersReportingResponseAnswersItemContactFieldFullName)
}

// SetPhone sets the Phone field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionAnswersReportingResponseAnswersItemContact) SetPhone(phone *string) {
	l.Phone = phone
	l.require(listSurveyQuestionAnswersReportingResponseAnswersItemContactFieldPhone)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionAnswersReportingResponseAnswersItemContact) SetStatus(status *ListSurveyQuestionAnswersReportingResponseAnswersItemContactStatus) {
	l.Status = status
	l.require(listSurveyQuestionAnswersReportingResponseAnswersItemContactFieldStatus)
}

func (l *ListSurveyQuestionAnswersReportingResponseAnswersItemContact) UnmarshalJSON(data []byte) error {
	type unmarshaler ListSurveyQuestionAnswersReportingResponseAnswersItemContact
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListSurveyQuestionAnswersReportingResponseAnswersItemContact(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListSurveyQuestionAnswersReportingResponseAnswersItemContact) MarshalJSON() ([]byte, error) {
	type embed ListSurveyQuestionAnswersReportingResponseAnswersItemContact
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListSurveyQuestionAnswersReportingResponseAnswersItemContact) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// The contact's current status.
type ListSurveyQuestionAnswersReportingResponseAnswersItemContactStatus string

const (
	ListSurveyQuestionAnswersReportingResponseAnswersItemContactStatusSubscribed    ListSurveyQuestionAnswersReportingResponseAnswersItemContactStatus = "Subscribed"
	ListSurveyQuestionAnswersReportingResponseAnswersItemContactStatusUnsubscribed  ListSurveyQuestionAnswersReportingResponseAnswersItemContactStatus = "Unsubscribed"
	ListSurveyQuestionAnswersReportingResponseAnswersItemContactStatusNonSubscribed ListSurveyQuestionAnswersReportingResponseAnswersItemContactStatus = "Non-Subscribed"
	ListSurveyQuestionAnswersReportingResponseAnswersItemContactStatusCleaned       ListSurveyQuestionAnswersReportingResponseAnswersItemContactStatus = "Cleaned"
	ListSurveyQuestionAnswersReportingResponseAnswersItemContactStatusArchived      ListSurveyQuestionAnswersReportingResponseAnswersItemContactStatus = "Archived"
)

func NewListSurveyQuestionAnswersReportingResponseAnswersItemContactStatusFromString(s string) (ListSurveyQuestionAnswersReportingResponseAnswersItemContactStatus, error) {
	switch s {
	case "Subscribed":
		return ListSurveyQuestionAnswersReportingResponseAnswersItemContactStatusSubscribed, nil
	case "Unsubscribed":
		return ListSurveyQuestionAnswersReportingResponseAnswersItemContactStatusUnsubscribed, nil
	case "Non-Subscribed":
		return ListSurveyQuestionAnswersReportingResponseAnswersItemContactStatusNonSubscribed, nil
	case "Cleaned":
		return ListSurveyQuestionAnswersReportingResponseAnswersItemContactStatusCleaned, nil
	case "Archived":
		return ListSurveyQuestionAnswersReportingResponseAnswersItemContactStatusArchived, nil
	}
	var t ListSurveyQuestionAnswersReportingResponseAnswersItemContactStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListSurveyQuestionAnswersReportingResponseAnswersItemContactStatus) Ptr() *ListSurveyQuestionAnswersReportingResponseAnswersItemContactStatus {
	return &l
}

// This object represents a link from the resource where it is found to another resource or action that may be performed.
var (
	listSurveyQuestionAnswersReportingResponseLinksItemFieldHref         = big.NewInt(1 << 0)
	listSurveyQuestionAnswersReportingResponseLinksItemFieldMethod       = big.NewInt(1 << 1)
	listSurveyQuestionAnswersReportingResponseLinksItemFieldRel          = big.NewInt(1 << 2)
	listSurveyQuestionAnswersReportingResponseLinksItemFieldSchema       = big.NewInt(1 << 3)
	listSurveyQuestionAnswersReportingResponseLinksItemFieldTargetSchema = big.NewInt(1 << 4)
)

type ListSurveyQuestionAnswersReportingResponseLinksItem struct {
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href *string `json:"href,omitempty" url:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method *ListSurveyQuestionAnswersReportingResponseLinksItemMethod `json:"method,omitempty" url:"method,omitempty"`
	// As with an HTML 'rel' attribute, this describes the type of link.
	Rel *string `json:"rel,omitempty" url:"rel,omitempty"`
	// For HTTP methods that can receive bodies (POST and PUT), this is a URL representing the schema that the body should conform to.
	Schema *string `json:"schema,omitempty" url:"schema,omitempty"`
	// For GETs, this is a URL representing the schema that the response should conform to.
	TargetSchema *string `json:"targetSchema,omitempty" url:"targetSchema,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListSurveyQuestionAnswersReportingResponseLinksItem) GetHref() *string {
	if l == nil {
		return nil
	}
	return l.Href
}

func (l *ListSurveyQuestionAnswersReportingResponseLinksItem) GetMethod() *ListSurveyQuestionAnswersReportingResponseLinksItemMethod {
	if l == nil {
		return nil
	}
	return l.Method
}

func (l *ListSurveyQuestionAnswersReportingResponseLinksItem) GetRel() *string {
	if l == nil {
		return nil
	}
	return l.Rel
}

func (l *ListSurveyQuestionAnswersReportingResponseLinksItem) GetSchema() *string {
	if l == nil {
		return nil
	}
	return l.Schema
}

func (l *ListSurveyQuestionAnswersReportingResponseLinksItem) GetTargetSchema() *string {
	if l == nil {
		return nil
	}
	return l.TargetSchema
}

func (l *ListSurveyQuestionAnswersReportingResponseLinksItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListSurveyQuestionAnswersReportingResponseLinksItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetHref sets the Href field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionAnswersReportingResponseLinksItem) SetHref(href *string) {
	l.Href = href
	l.require(listSurveyQuestionAnswersReportingResponseLinksItemFieldHref)
}

// SetMethod sets the Method field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionAnswersReportingResponseLinksItem) SetMethod(method *ListSurveyQuestionAnswersReportingResponseLinksItemMethod) {
	l.Method = method
	l.require(listSurveyQuestionAnswersReportingResponseLinksItemFieldMethod)
}

// SetRel sets the Rel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionAnswersReportingResponseLinksItem) SetRel(rel *string) {
	l.Rel = rel
	l.require(listSurveyQuestionAnswersReportingResponseLinksItemFieldRel)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionAnswersReportingResponseLinksItem) SetSchema(schema *string) {
	l.Schema = schema
	l.require(listSurveyQuestionAnswersReportingResponseLinksItemFieldSchema)
}

// SetTargetSchema sets the TargetSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionAnswersReportingResponseLinksItem) SetTargetSchema(targetSchema *string) {
	l.TargetSchema = targetSchema
	l.require(listSurveyQuestionAnswersReportingResponseLinksItemFieldTargetSchema)
}

func (l *ListSurveyQuestionAnswersReportingResponseLinksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler ListSurveyQuestionAnswersReportingResponseLinksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListSurveyQuestionAnswersReportingResponseLinksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListSurveyQuestionAnswersReportingResponseLinksItem) MarshalJSON() ([]byte, error) {
	type embed ListSurveyQuestionAnswersReportingResponseLinksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListSurveyQuestionAnswersReportingResponseLinksItem) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// The HTTP method that should be used when accessing the URL defined in 'href'.
type ListSurveyQuestionAnswersReportingResponseLinksItemMethod string

const (
	ListSurveyQuestionAnswersReportingResponseLinksItemMethodGet     ListSurveyQuestionAnswersReportingResponseLinksItemMethod = "GET"
	ListSurveyQuestionAnswersReportingResponseLinksItemMethodPost    ListSurveyQuestionAnswersReportingResponseLinksItemMethod = "POST"
	ListSurveyQuestionAnswersReportingResponseLinksItemMethodPut     ListSurveyQuestionAnswersReportingResponseLinksItemMethod = "PUT"
	ListSurveyQuestionAnswersReportingResponseLinksItemMethodPatch   ListSurveyQuestionAnswersReportingResponseLinksItemMethod = "PATCH"
	ListSurveyQuestionAnswersReportingResponseLinksItemMethodDelete  ListSurveyQuestionAnswersReportingResponseLinksItemMethod = "DELETE"
	ListSurveyQuestionAnswersReportingResponseLinksItemMethodOptions ListSurveyQuestionAnswersReportingResponseLinksItemMethod = "OPTIONS"
	ListSurveyQuestionAnswersReportingResponseLinksItemMethodHead    ListSurveyQuestionAnswersReportingResponseLinksItemMethod = "HEAD"
)

func NewListSurveyQuestionAnswersReportingResponseLinksItemMethodFromString(s string) (ListSurveyQuestionAnswersReportingResponseLinksItemMethod, error) {
	switch s {
	case "GET":
		return ListSurveyQuestionAnswersReportingResponseLinksItemMethodGet, nil
	case "POST":
		return ListSurveyQuestionAnswersReportingResponseLinksItemMethodPost, nil
	case "PUT":
		return ListSurveyQuestionAnswersReportingResponseLinksItemMethodPut, nil
	case "PATCH":
		return ListSurveyQuestionAnswersReportingResponseLinksItemMethodPatch, nil
	case "DELETE":
		return ListSurveyQuestionAnswersReportingResponseLinksItemMethodDelete, nil
	case "OPTIONS":
		return ListSurveyQuestionAnswersReportingResponseLinksItemMethodOptions, nil
	case "HEAD":
		return ListSurveyQuestionAnswersReportingResponseLinksItemMethodHead, nil
	}
	var t ListSurveyQuestionAnswersReportingResponseLinksItemMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListSurveyQuestionAnswersReportingResponseLinksItemMethod) Ptr() *ListSurveyQuestionAnswersReportingResponseLinksItemMethod {
	return &l
}

var (
	listSurveyQuestionsReportingResponseFieldLinks      = big.NewInt(1 << 0)
	listSurveyQuestionsReportingResponseFieldQuestions  = big.NewInt(1 << 1)
	listSurveyQuestionsReportingResponseFieldTotalItems = big.NewInt(1 << 2)
)

type ListSurveyQuestionsReportingResponse struct {
	// A list of link types and descriptions for the API schema documents.
	Links []*ListSurveyQuestionsReportingResponseLinksItem `json:"_links,omitempty" url:"_links,omitempty"`
	// An array of reports for each question on the survey.
	Questions []*SurveyQuestionReport `json:"questions,omitempty" url:"questions,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int `json:"total_items,omitempty" url:"total_items,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListSurveyQuestionsReportingResponse) GetLinks() []*ListSurveyQuestionsReportingResponseLinksItem {
	if l == nil {
		return nil
	}
	return l.Links
}

func (l *ListSurveyQuestionsReportingResponse) GetQuestions() []*SurveyQuestionReport {
	if l == nil {
		return nil
	}
	return l.Questions
}

func (l *ListSurveyQuestionsReportingResponse) GetTotalItems() *int {
	if l == nil {
		return nil
	}
	return l.TotalItems
}

func (l *ListSurveyQuestionsReportingResponse) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListSurveyQuestionsReportingResponse) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetLinks sets the Links field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionsReportingResponse) SetLinks(links []*ListSurveyQuestionsReportingResponseLinksItem) {
	l.Links = links
	l.require(listSurveyQuestionsReportingResponseFieldLinks)
}

// SetQuestions sets the Questions field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionsReportingResponse) SetQuestions(questions []*SurveyQuestionReport) {
	l.Questions = questions
	l.require(listSurveyQuestionsReportingResponseFieldQuestions)
}

// SetTotalItems sets the TotalItems field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionsReportingResponse) SetTotalItems(totalItems *int) {
	l.TotalItems = totalItems
	l.require(listSurveyQuestionsReportingResponseFieldTotalItems)
}

func (l *ListSurveyQuestionsReportingResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListSurveyQuestionsReportingResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListSurveyQuestionsReportingResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListSurveyQuestionsReportingResponse) MarshalJSON() ([]byte, error) {
	type embed ListSurveyQuestionsReportingResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListSurveyQuestionsReportingResponse) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// This object represents a link from the resource where it is found to another resource or action that may be performed.
var (
	listSurveyQuestionsReportingResponseLinksItemFieldHref         = big.NewInt(1 << 0)
	listSurveyQuestionsReportingResponseLinksItemFieldMethod       = big.NewInt(1 << 1)
	listSurveyQuestionsReportingResponseLinksItemFieldRel          = big.NewInt(1 << 2)
	listSurveyQuestionsReportingResponseLinksItemFieldSchema       = big.NewInt(1 << 3)
	listSurveyQuestionsReportingResponseLinksItemFieldTargetSchema = big.NewInt(1 << 4)
)

type ListSurveyQuestionsReportingResponseLinksItem struct {
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href *string `json:"href,omitempty" url:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method *ListSurveyQuestionsReportingResponseLinksItemMethod `json:"method,omitempty" url:"method,omitempty"`
	// As with an HTML 'rel' attribute, this describes the type of link.
	Rel *string `json:"rel,omitempty" url:"rel,omitempty"`
	// For HTTP methods that can receive bodies (POST and PUT), this is a URL representing the schema that the body should conform to.
	Schema *string `json:"schema,omitempty" url:"schema,omitempty"`
	// For GETs, this is a URL representing the schema that the response should conform to.
	TargetSchema *string `json:"targetSchema,omitempty" url:"targetSchema,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListSurveyQuestionsReportingResponseLinksItem) GetHref() *string {
	if l == nil {
		return nil
	}
	return l.Href
}

func (l *ListSurveyQuestionsReportingResponseLinksItem) GetMethod() *ListSurveyQuestionsReportingResponseLinksItemMethod {
	if l == nil {
		return nil
	}
	return l.Method
}

func (l *ListSurveyQuestionsReportingResponseLinksItem) GetRel() *string {
	if l == nil {
		return nil
	}
	return l.Rel
}

func (l *ListSurveyQuestionsReportingResponseLinksItem) GetSchema() *string {
	if l == nil {
		return nil
	}
	return l.Schema
}

func (l *ListSurveyQuestionsReportingResponseLinksItem) GetTargetSchema() *string {
	if l == nil {
		return nil
	}
	return l.TargetSchema
}

func (l *ListSurveyQuestionsReportingResponseLinksItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListSurveyQuestionsReportingResponseLinksItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetHref sets the Href field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionsReportingResponseLinksItem) SetHref(href *string) {
	l.Href = href
	l.require(listSurveyQuestionsReportingResponseLinksItemFieldHref)
}

// SetMethod sets the Method field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionsReportingResponseLinksItem) SetMethod(method *ListSurveyQuestionsReportingResponseLinksItemMethod) {
	l.Method = method
	l.require(listSurveyQuestionsReportingResponseLinksItemFieldMethod)
}

// SetRel sets the Rel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionsReportingResponseLinksItem) SetRel(rel *string) {
	l.Rel = rel
	l.require(listSurveyQuestionsReportingResponseLinksItemFieldRel)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionsReportingResponseLinksItem) SetSchema(schema *string) {
	l.Schema = schema
	l.require(listSurveyQuestionsReportingResponseLinksItemFieldSchema)
}

// SetTargetSchema sets the TargetSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyQuestionsReportingResponseLinksItem) SetTargetSchema(targetSchema *string) {
	l.TargetSchema = targetSchema
	l.require(listSurveyQuestionsReportingResponseLinksItemFieldTargetSchema)
}

func (l *ListSurveyQuestionsReportingResponseLinksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler ListSurveyQuestionsReportingResponseLinksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListSurveyQuestionsReportingResponseLinksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListSurveyQuestionsReportingResponseLinksItem) MarshalJSON() ([]byte, error) {
	type embed ListSurveyQuestionsReportingResponseLinksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListSurveyQuestionsReportingResponseLinksItem) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// The HTTP method that should be used when accessing the URL defined in 'href'.
type ListSurveyQuestionsReportingResponseLinksItemMethod string

const (
	ListSurveyQuestionsReportingResponseLinksItemMethodGet     ListSurveyQuestionsReportingResponseLinksItemMethod = "GET"
	ListSurveyQuestionsReportingResponseLinksItemMethodPost    ListSurveyQuestionsReportingResponseLinksItemMethod = "POST"
	ListSurveyQuestionsReportingResponseLinksItemMethodPut     ListSurveyQuestionsReportingResponseLinksItemMethod = "PUT"
	ListSurveyQuestionsReportingResponseLinksItemMethodPatch   ListSurveyQuestionsReportingResponseLinksItemMethod = "PATCH"
	ListSurveyQuestionsReportingResponseLinksItemMethodDelete  ListSurveyQuestionsReportingResponseLinksItemMethod = "DELETE"
	ListSurveyQuestionsReportingResponseLinksItemMethodOptions ListSurveyQuestionsReportingResponseLinksItemMethod = "OPTIONS"
	ListSurveyQuestionsReportingResponseLinksItemMethodHead    ListSurveyQuestionsReportingResponseLinksItemMethod = "HEAD"
)

func NewListSurveyQuestionsReportingResponseLinksItemMethodFromString(s string) (ListSurveyQuestionsReportingResponseLinksItemMethod, error) {
	switch s {
	case "GET":
		return ListSurveyQuestionsReportingResponseLinksItemMethodGet, nil
	case "POST":
		return ListSurveyQuestionsReportingResponseLinksItemMethodPost, nil
	case "PUT":
		return ListSurveyQuestionsReportingResponseLinksItemMethodPut, nil
	case "PATCH":
		return ListSurveyQuestionsReportingResponseLinksItemMethodPatch, nil
	case "DELETE":
		return ListSurveyQuestionsReportingResponseLinksItemMethodDelete, nil
	case "OPTIONS":
		return ListSurveyQuestionsReportingResponseLinksItemMethodOptions, nil
	case "HEAD":
		return ListSurveyQuestionsReportingResponseLinksItemMethodHead, nil
	}
	var t ListSurveyQuestionsReportingResponseLinksItemMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListSurveyQuestionsReportingResponseLinksItemMethod) Ptr() *ListSurveyQuestionsReportingResponseLinksItemMethod {
	return &l
}

type ListSurveyResponsesReportingRequestRespondentFamiliarityIs string

const (
	ListSurveyResponsesReportingRequestRespondentFamiliarityIsNew     ListSurveyResponsesReportingRequestRespondentFamiliarityIs = "new"
	ListSurveyResponsesReportingRequestRespondentFamiliarityIsKnown   ListSurveyResponsesReportingRequestRespondentFamiliarityIs = "known"
	ListSurveyResponsesReportingRequestRespondentFamiliarityIsUnknown ListSurveyResponsesReportingRequestRespondentFamiliarityIs = "unknown"
)

func NewListSurveyResponsesReportingRequestRespondentFamiliarityIsFromString(s string) (ListSurveyResponsesReportingRequestRespondentFamiliarityIs, error) {
	switch s {
	case "new":
		return ListSurveyResponsesReportingRequestRespondentFamiliarityIsNew, nil
	case "known":
		return ListSurveyResponsesReportingRequestRespondentFamiliarityIsKnown, nil
	case "unknown":
		return ListSurveyResponsesReportingRequestRespondentFamiliarityIsUnknown, nil
	}
	var t ListSurveyResponsesReportingRequestRespondentFamiliarityIs
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListSurveyResponsesReportingRequestRespondentFamiliarityIs) Ptr() *ListSurveyResponsesReportingRequestRespondentFamiliarityIs {
	return &l
}

var (
	listSurveyResponsesReportingResponseFieldLinks      = big.NewInt(1 << 0)
	listSurveyResponsesReportingResponseFieldResponses  = big.NewInt(1 << 1)
	listSurveyResponsesReportingResponseFieldTotalItems = big.NewInt(1 << 2)
)

type ListSurveyResponsesReportingResponse struct {
	// A list of link types and descriptions for the API schema documents.
	Links []*ListSurveyResponsesReportingResponseLinksItem `json:"_links,omitempty" url:"_links,omitempty"`
	// An array of responses to a survey.
	Responses []*ListSurveyResponsesReportingResponseResponsesItem `json:"responses,omitempty" url:"responses,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int `json:"total_items,omitempty" url:"total_items,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListSurveyResponsesReportingResponse) GetLinks() []*ListSurveyResponsesReportingResponseLinksItem {
	if l == nil {
		return nil
	}
	return l.Links
}

func (l *ListSurveyResponsesReportingResponse) GetResponses() []*ListSurveyResponsesReportingResponseResponsesItem {
	if l == nil {
		return nil
	}
	return l.Responses
}

func (l *ListSurveyResponsesReportingResponse) GetTotalItems() *int {
	if l == nil {
		return nil
	}
	return l.TotalItems
}

func (l *ListSurveyResponsesReportingResponse) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListSurveyResponsesReportingResponse) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetLinks sets the Links field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyResponsesReportingResponse) SetLinks(links []*ListSurveyResponsesReportingResponseLinksItem) {
	l.Links = links
	l.require(listSurveyResponsesReportingResponseFieldLinks)
}

// SetResponses sets the Responses field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyResponsesReportingResponse) SetResponses(responses []*ListSurveyResponsesReportingResponseResponsesItem) {
	l.Responses = responses
	l.require(listSurveyResponsesReportingResponseFieldResponses)
}

// SetTotalItems sets the TotalItems field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyResponsesReportingResponse) SetTotalItems(totalItems *int) {
	l.TotalItems = totalItems
	l.require(listSurveyResponsesReportingResponseFieldTotalItems)
}

func (l *ListSurveyResponsesReportingResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListSurveyResponsesReportingResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListSurveyResponsesReportingResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListSurveyResponsesReportingResponse) MarshalJSON() ([]byte, error) {
	type embed ListSurveyResponsesReportingResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListSurveyResponsesReportingResponse) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// This object represents a link from the resource where it is found to another resource or action that may be performed.
var (
	listSurveyResponsesReportingResponseLinksItemFieldHref         = big.NewInt(1 << 0)
	listSurveyResponsesReportingResponseLinksItemFieldMethod       = big.NewInt(1 << 1)
	listSurveyResponsesReportingResponseLinksItemFieldRel          = big.NewInt(1 << 2)
	listSurveyResponsesReportingResponseLinksItemFieldSchema       = big.NewInt(1 << 3)
	listSurveyResponsesReportingResponseLinksItemFieldTargetSchema = big.NewInt(1 << 4)
)

type ListSurveyResponsesReportingResponseLinksItem struct {
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href *string `json:"href,omitempty" url:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method *ListSurveyResponsesReportingResponseLinksItemMethod `json:"method,omitempty" url:"method,omitempty"`
	// As with an HTML 'rel' attribute, this describes the type of link.
	Rel *string `json:"rel,omitempty" url:"rel,omitempty"`
	// For HTTP methods that can receive bodies (POST and PUT), this is a URL representing the schema that the body should conform to.
	Schema *string `json:"schema,omitempty" url:"schema,omitempty"`
	// For GETs, this is a URL representing the schema that the response should conform to.
	TargetSchema *string `json:"targetSchema,omitempty" url:"targetSchema,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListSurveyResponsesReportingResponseLinksItem) GetHref() *string {
	if l == nil {
		return nil
	}
	return l.Href
}

func (l *ListSurveyResponsesReportingResponseLinksItem) GetMethod() *ListSurveyResponsesReportingResponseLinksItemMethod {
	if l == nil {
		return nil
	}
	return l.Method
}

func (l *ListSurveyResponsesReportingResponseLinksItem) GetRel() *string {
	if l == nil {
		return nil
	}
	return l.Rel
}

func (l *ListSurveyResponsesReportingResponseLinksItem) GetSchema() *string {
	if l == nil {
		return nil
	}
	return l.Schema
}

func (l *ListSurveyResponsesReportingResponseLinksItem) GetTargetSchema() *string {
	if l == nil {
		return nil
	}
	return l.TargetSchema
}

func (l *ListSurveyResponsesReportingResponseLinksItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListSurveyResponsesReportingResponseLinksItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetHref sets the Href field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyResponsesReportingResponseLinksItem) SetHref(href *string) {
	l.Href = href
	l.require(listSurveyResponsesReportingResponseLinksItemFieldHref)
}

// SetMethod sets the Method field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyResponsesReportingResponseLinksItem) SetMethod(method *ListSurveyResponsesReportingResponseLinksItemMethod) {
	l.Method = method
	l.require(listSurveyResponsesReportingResponseLinksItemFieldMethod)
}

// SetRel sets the Rel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyResponsesReportingResponseLinksItem) SetRel(rel *string) {
	l.Rel = rel
	l.require(listSurveyResponsesReportingResponseLinksItemFieldRel)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyResponsesReportingResponseLinksItem) SetSchema(schema *string) {
	l.Schema = schema
	l.require(listSurveyResponsesReportingResponseLinksItemFieldSchema)
}

// SetTargetSchema sets the TargetSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyResponsesReportingResponseLinksItem) SetTargetSchema(targetSchema *string) {
	l.TargetSchema = targetSchema
	l.require(listSurveyResponsesReportingResponseLinksItemFieldTargetSchema)
}

func (l *ListSurveyResponsesReportingResponseLinksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler ListSurveyResponsesReportingResponseLinksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListSurveyResponsesReportingResponseLinksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListSurveyResponsesReportingResponseLinksItem) MarshalJSON() ([]byte, error) {
	type embed ListSurveyResponsesReportingResponseLinksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListSurveyResponsesReportingResponseLinksItem) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// The HTTP method that should be used when accessing the URL defined in 'href'.
type ListSurveyResponsesReportingResponseLinksItemMethod string

const (
	ListSurveyResponsesReportingResponseLinksItemMethodGet     ListSurveyResponsesReportingResponseLinksItemMethod = "GET"
	ListSurveyResponsesReportingResponseLinksItemMethodPost    ListSurveyResponsesReportingResponseLinksItemMethod = "POST"
	ListSurveyResponsesReportingResponseLinksItemMethodPut     ListSurveyResponsesReportingResponseLinksItemMethod = "PUT"
	ListSurveyResponsesReportingResponseLinksItemMethodPatch   ListSurveyResponsesReportingResponseLinksItemMethod = "PATCH"
	ListSurveyResponsesReportingResponseLinksItemMethodDelete  ListSurveyResponsesReportingResponseLinksItemMethod = "DELETE"
	ListSurveyResponsesReportingResponseLinksItemMethodOptions ListSurveyResponsesReportingResponseLinksItemMethod = "OPTIONS"
	ListSurveyResponsesReportingResponseLinksItemMethodHead    ListSurveyResponsesReportingResponseLinksItemMethod = "HEAD"
)

func NewListSurveyResponsesReportingResponseLinksItemMethodFromString(s string) (ListSurveyResponsesReportingResponseLinksItemMethod, error) {
	switch s {
	case "GET":
		return ListSurveyResponsesReportingResponseLinksItemMethodGet, nil
	case "POST":
		return ListSurveyResponsesReportingResponseLinksItemMethodPost, nil
	case "PUT":
		return ListSurveyResponsesReportingResponseLinksItemMethodPut, nil
	case "PATCH":
		return ListSurveyResponsesReportingResponseLinksItemMethodPatch, nil
	case "DELETE":
		return ListSurveyResponsesReportingResponseLinksItemMethodDelete, nil
	case "OPTIONS":
		return ListSurveyResponsesReportingResponseLinksItemMethodOptions, nil
	case "HEAD":
		return ListSurveyResponsesReportingResponseLinksItemMethodHead, nil
	}
	var t ListSurveyResponsesReportingResponseLinksItemMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListSurveyResponsesReportingResponseLinksItemMethod) Ptr() *ListSurveyResponsesReportingResponseLinksItemMethod {
	return &l
}

// Survey respondent details.
var (
	listSurveyResponsesReportingResponseResponsesItemFieldContact      = big.NewInt(1 << 0)
	listSurveyResponsesReportingResponseResponsesItemFieldIsNewContact = big.NewInt(1 << 1)
	listSurveyResponsesReportingResponseResponsesItemFieldResponseID   = big.NewInt(1 << 2)
	listSurveyResponsesReportingResponseResponsesItemFieldSubmittedAt  = big.NewInt(1 << 3)
)

type ListSurveyResponsesReportingResponseResponsesItem struct {
	// Information about the contact.
	Contact *ListSurveyResponsesReportingResponseResponsesItemContact `json:"contact,omitempty" url:"contact,omitempty"`
	// If this contact was added to the Mailchimp audience via this survey.
	IsNewContact *bool `json:"is_new_contact,omitempty" url:"is_new_contact,omitempty"`
	// The ID for the survey response.
	ResponseID *string `json:"response_id,omitempty" url:"response_id,omitempty"`
	// The date and time when the survey response was submitted in ISO 8601 format.
	SubmittedAt *time.Time `json:"submitted_at,omitempty" url:"submitted_at,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListSurveyResponsesReportingResponseResponsesItem) GetContact() *ListSurveyResponsesReportingResponseResponsesItemContact {
	if l == nil {
		return nil
	}
	return l.Contact
}

func (l *ListSurveyResponsesReportingResponseResponsesItem) GetIsNewContact() *bool {
	if l == nil {
		return nil
	}
	return l.IsNewContact
}

func (l *ListSurveyResponsesReportingResponseResponsesItem) GetResponseID() *string {
	if l == nil {
		return nil
	}
	return l.ResponseID
}

func (l *ListSurveyResponsesReportingResponseResponsesItem) GetSubmittedAt() *time.Time {
	if l == nil {
		return nil
	}
	return l.SubmittedAt
}

func (l *ListSurveyResponsesReportingResponseResponsesItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListSurveyResponsesReportingResponseResponsesItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetContact sets the Contact field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyResponsesReportingResponseResponsesItem) SetContact(contact *ListSurveyResponsesReportingResponseResponsesItemContact) {
	l.Contact = contact
	l.require(listSurveyResponsesReportingResponseResponsesItemFieldContact)
}

// SetIsNewContact sets the IsNewContact field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyResponsesReportingResponseResponsesItem) SetIsNewContact(isNewContact *bool) {
	l.IsNewContact = isNewContact
	l.require(listSurveyResponsesReportingResponseResponsesItemFieldIsNewContact)
}

// SetResponseID sets the ResponseID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyResponsesReportingResponseResponsesItem) SetResponseID(responseID *string) {
	l.ResponseID = responseID
	l.require(listSurveyResponsesReportingResponseResponsesItemFieldResponseID)
}

// SetSubmittedAt sets the SubmittedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyResponsesReportingResponseResponsesItem) SetSubmittedAt(submittedAt *time.Time) {
	l.SubmittedAt = submittedAt
	l.require(listSurveyResponsesReportingResponseResponsesItemFieldSubmittedAt)
}

func (l *ListSurveyResponsesReportingResponseResponsesItem) UnmarshalJSON(data []byte) error {
	type embed ListSurveyResponsesReportingResponseResponsesItem
	var unmarshaler = struct {
		embed
		SubmittedAt *internal.DateTime `json:"submitted_at,omitempty"`
	}{
		embed: embed(*l),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*l = ListSurveyResponsesReportingResponseResponsesItem(unmarshaler.embed)
	l.SubmittedAt = unmarshaler.SubmittedAt.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListSurveyResponsesReportingResponseResponsesItem) MarshalJSON() ([]byte, error) {
	type embed ListSurveyResponsesReportingResponseResponsesItem
	var marshaler = struct {
		embed
		SubmittedAt *internal.DateTime `json:"submitted_at,omitempty"`
	}{
		embed:       embed(*l),
		SubmittedAt: internal.NewOptionalDateTime(l.SubmittedAt),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListSurveyResponsesReportingResponseResponsesItem) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// Information about the contact.
var (
	listSurveyResponsesReportingResponseResponsesItemContactFieldAvatarURL                   = big.NewInt(1 << 0)
	listSurveyResponsesReportingResponseResponsesItemContactFieldConsentsToOneToOneMessaging = big.NewInt(1 << 1)
	listSurveyResponsesReportingResponseResponsesItemContactFieldContactID                   = big.NewInt(1 << 2)
	listSurveyResponsesReportingResponseResponsesItemContactFieldEmail                       = big.NewInt(1 << 3)
	listSurveyResponsesReportingResponseResponsesItemContactFieldEmailID                     = big.NewInt(1 << 4)
	listSurveyResponsesReportingResponseResponsesItemContactFieldFullName                    = big.NewInt(1 << 5)
	listSurveyResponsesReportingResponseResponsesItemContactFieldPhone                       = big.NewInt(1 << 6)
	listSurveyResponsesReportingResponseResponsesItemContactFieldStatus                      = big.NewInt(1 << 7)
)

type ListSurveyResponsesReportingResponseResponsesItemContact struct {
	// URL for the contact's avatar or profile image.
	AvatarURL *string `json:"avatar_url,omitempty" url:"avatar_url,omitempty"`
	// Indicates whether a contact consents to 1:1 messaging.
	ConsentsToOneToOneMessaging *bool `json:"consents_to_one_to_one_messaging,omitempty" url:"consents_to_one_to_one_messaging,omitempty"`
	// The ID of this contact.
	ContactID *string `json:"contact_id,omitempty" url:"contact_id,omitempty"`
	// The contact's email address.
	Email *string `json:"email,omitempty" url:"email,omitempty"`
	// The MD5 hash of the lowercase version of the list member's email address.
	EmailID *string `json:"email_id,omitempty" url:"email_id,omitempty"`
	// The contact's full name.
	FullName *string `json:"full_name,omitempty" url:"full_name,omitempty"`
	// The contact's sms phone number.
	Phone *string `json:"phone,omitempty" url:"phone,omitempty"`
	// The contact's current status.
	Status *ListSurveyResponsesReportingResponseResponsesItemContactStatus `json:"status,omitempty" url:"status,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListSurveyResponsesReportingResponseResponsesItemContact) GetAvatarURL() *string {
	if l == nil {
		return nil
	}
	return l.AvatarURL
}

func (l *ListSurveyResponsesReportingResponseResponsesItemContact) GetConsentsToOneToOneMessaging() *bool {
	if l == nil {
		return nil
	}
	return l.ConsentsToOneToOneMessaging
}

func (l *ListSurveyResponsesReportingResponseResponsesItemContact) GetContactID() *string {
	if l == nil {
		return nil
	}
	return l.ContactID
}

func (l *ListSurveyResponsesReportingResponseResponsesItemContact) GetEmail() *string {
	if l == nil {
		return nil
	}
	return l.Email
}

func (l *ListSurveyResponsesReportingResponseResponsesItemContact) GetEmailID() *string {
	if l == nil {
		return nil
	}
	return l.EmailID
}

func (l *ListSurveyResponsesReportingResponseResponsesItemContact) GetFullName() *string {
	if l == nil {
		return nil
	}
	return l.FullName
}

func (l *ListSurveyResponsesReportingResponseResponsesItemContact) GetPhone() *string {
	if l == nil {
		return nil
	}
	return l.Phone
}

func (l *ListSurveyResponsesReportingResponseResponsesItemContact) GetStatus() *ListSurveyResponsesReportingResponseResponsesItemContactStatus {
	if l == nil {
		return nil
	}
	return l.Status
}

func (l *ListSurveyResponsesReportingResponseResponsesItemContact) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListSurveyResponsesReportingResponseResponsesItemContact) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetAvatarURL sets the AvatarURL field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyResponsesReportingResponseResponsesItemContact) SetAvatarURL(avatarURL *string) {
	l.AvatarURL = avatarURL
	l.require(listSurveyResponsesReportingResponseResponsesItemContactFieldAvatarURL)
}

// SetConsentsToOneToOneMessaging sets the ConsentsToOneToOneMessaging field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyResponsesReportingResponseResponsesItemContact) SetConsentsToOneToOneMessaging(consentsToOneToOneMessaging *bool) {
	l.ConsentsToOneToOneMessaging = consentsToOneToOneMessaging
	l.require(listSurveyResponsesReportingResponseResponsesItemContactFieldConsentsToOneToOneMessaging)
}

// SetContactID sets the ContactID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyResponsesReportingResponseResponsesItemContact) SetContactID(contactID *string) {
	l.ContactID = contactID
	l.require(listSurveyResponsesReportingResponseResponsesItemContactFieldContactID)
}

// SetEmail sets the Email field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyResponsesReportingResponseResponsesItemContact) SetEmail(email *string) {
	l.Email = email
	l.require(listSurveyResponsesReportingResponseResponsesItemContactFieldEmail)
}

// SetEmailID sets the EmailID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyResponsesReportingResponseResponsesItemContact) SetEmailID(emailID *string) {
	l.EmailID = emailID
	l.require(listSurveyResponsesReportingResponseResponsesItemContactFieldEmailID)
}

// SetFullName sets the FullName field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyResponsesReportingResponseResponsesItemContact) SetFullName(fullName *string) {
	l.FullName = fullName
	l.require(listSurveyResponsesReportingResponseResponsesItemContactFieldFullName)
}

// SetPhone sets the Phone field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyResponsesReportingResponseResponsesItemContact) SetPhone(phone *string) {
	l.Phone = phone
	l.require(listSurveyResponsesReportingResponseResponsesItemContactFieldPhone)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveyResponsesReportingResponseResponsesItemContact) SetStatus(status *ListSurveyResponsesReportingResponseResponsesItemContactStatus) {
	l.Status = status
	l.require(listSurveyResponsesReportingResponseResponsesItemContactFieldStatus)
}

func (l *ListSurveyResponsesReportingResponseResponsesItemContact) UnmarshalJSON(data []byte) error {
	type unmarshaler ListSurveyResponsesReportingResponseResponsesItemContact
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListSurveyResponsesReportingResponseResponsesItemContact(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListSurveyResponsesReportingResponseResponsesItemContact) MarshalJSON() ([]byte, error) {
	type embed ListSurveyResponsesReportingResponseResponsesItemContact
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListSurveyResponsesReportingResponseResponsesItemContact) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// The contact's current status.
type ListSurveyResponsesReportingResponseResponsesItemContactStatus string

const (
	ListSurveyResponsesReportingResponseResponsesItemContactStatusSubscribed    ListSurveyResponsesReportingResponseResponsesItemContactStatus = "Subscribed"
	ListSurveyResponsesReportingResponseResponsesItemContactStatusUnsubscribed  ListSurveyResponsesReportingResponseResponsesItemContactStatus = "Unsubscribed"
	ListSurveyResponsesReportingResponseResponsesItemContactStatusNonSubscribed ListSurveyResponsesReportingResponseResponsesItemContactStatus = "Non-Subscribed"
	ListSurveyResponsesReportingResponseResponsesItemContactStatusCleaned       ListSurveyResponsesReportingResponseResponsesItemContactStatus = "Cleaned"
	ListSurveyResponsesReportingResponseResponsesItemContactStatusArchived      ListSurveyResponsesReportingResponseResponsesItemContactStatus = "Archived"
)

func NewListSurveyResponsesReportingResponseResponsesItemContactStatusFromString(s string) (ListSurveyResponsesReportingResponseResponsesItemContactStatus, error) {
	switch s {
	case "Subscribed":
		return ListSurveyResponsesReportingResponseResponsesItemContactStatusSubscribed, nil
	case "Unsubscribed":
		return ListSurveyResponsesReportingResponseResponsesItemContactStatusUnsubscribed, nil
	case "Non-Subscribed":
		return ListSurveyResponsesReportingResponseResponsesItemContactStatusNonSubscribed, nil
	case "Cleaned":
		return ListSurveyResponsesReportingResponseResponsesItemContactStatusCleaned, nil
	case "Archived":
		return ListSurveyResponsesReportingResponseResponsesItemContactStatusArchived, nil
	}
	var t ListSurveyResponsesReportingResponseResponsesItemContactStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListSurveyResponsesReportingResponseResponsesItemContactStatus) Ptr() *ListSurveyResponsesReportingResponseResponsesItemContactStatus {
	return &l
}

var (
	listSurveysReportingResponseFieldLinks      = big.NewInt(1 << 0)
	listSurveysReportingResponseFieldSurveys    = big.NewInt(1 << 1)
	listSurveysReportingResponseFieldTotalItems = big.NewInt(1 << 2)
)

type ListSurveysReportingResponse struct {
	// A list of link types and descriptions for the API schema documents.
	Links []*ListSurveysReportingResponseLinksItem `json:"_links,omitempty" url:"_links,omitempty"`
	// The surveys that have reports available.
	Surveys []*ListSurveysReportingResponseSurveysItem `json:"surveys,omitempty" url:"surveys,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int `json:"total_items,omitempty" url:"total_items,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListSurveysReportingResponse) GetLinks() []*ListSurveysReportingResponseLinksItem {
	if l == nil {
		return nil
	}
	return l.Links
}

func (l *ListSurveysReportingResponse) GetSurveys() []*ListSurveysReportingResponseSurveysItem {
	if l == nil {
		return nil
	}
	return l.Surveys
}

func (l *ListSurveysReportingResponse) GetTotalItems() *int {
	if l == nil {
		return nil
	}
	return l.TotalItems
}

func (l *ListSurveysReportingResponse) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListSurveysReportingResponse) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetLinks sets the Links field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveysReportingResponse) SetLinks(links []*ListSurveysReportingResponseLinksItem) {
	l.Links = links
	l.require(listSurveysReportingResponseFieldLinks)
}

// SetSurveys sets the Surveys field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveysReportingResponse) SetSurveys(surveys []*ListSurveysReportingResponseSurveysItem) {
	l.Surveys = surveys
	l.require(listSurveysReportingResponseFieldSurveys)
}

// SetTotalItems sets the TotalItems field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveysReportingResponse) SetTotalItems(totalItems *int) {
	l.TotalItems = totalItems
	l.require(listSurveysReportingResponseFieldTotalItems)
}

func (l *ListSurveysReportingResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListSurveysReportingResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListSurveysReportingResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListSurveysReportingResponse) MarshalJSON() ([]byte, error) {
	type embed ListSurveysReportingResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListSurveysReportingResponse) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// This object represents a link from the resource where it is found to another resource or action that may be performed.
var (
	listSurveysReportingResponseLinksItemFieldHref         = big.NewInt(1 << 0)
	listSurveysReportingResponseLinksItemFieldMethod       = big.NewInt(1 << 1)
	listSurveysReportingResponseLinksItemFieldRel          = big.NewInt(1 << 2)
	listSurveysReportingResponseLinksItemFieldSchema       = big.NewInt(1 << 3)
	listSurveysReportingResponseLinksItemFieldTargetSchema = big.NewInt(1 << 4)
)

type ListSurveysReportingResponseLinksItem struct {
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href *string `json:"href,omitempty" url:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method *ListSurveysReportingResponseLinksItemMethod `json:"method,omitempty" url:"method,omitempty"`
	// As with an HTML 'rel' attribute, this describes the type of link.
	Rel *string `json:"rel,omitempty" url:"rel,omitempty"`
	// For HTTP methods that can receive bodies (POST and PUT), this is a URL representing the schema that the body should conform to.
	Schema *string `json:"schema,omitempty" url:"schema,omitempty"`
	// For GETs, this is a URL representing the schema that the response should conform to.
	TargetSchema *string `json:"targetSchema,omitempty" url:"targetSchema,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListSurveysReportingResponseLinksItem) GetHref() *string {
	if l == nil {
		return nil
	}
	return l.Href
}

func (l *ListSurveysReportingResponseLinksItem) GetMethod() *ListSurveysReportingResponseLinksItemMethod {
	if l == nil {
		return nil
	}
	return l.Method
}

func (l *ListSurveysReportingResponseLinksItem) GetRel() *string {
	if l == nil {
		return nil
	}
	return l.Rel
}

func (l *ListSurveysReportingResponseLinksItem) GetSchema() *string {
	if l == nil {
		return nil
	}
	return l.Schema
}

func (l *ListSurveysReportingResponseLinksItem) GetTargetSchema() *string {
	if l == nil {
		return nil
	}
	return l.TargetSchema
}

func (l *ListSurveysReportingResponseLinksItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListSurveysReportingResponseLinksItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetHref sets the Href field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveysReportingResponseLinksItem) SetHref(href *string) {
	l.Href = href
	l.require(listSurveysReportingResponseLinksItemFieldHref)
}

// SetMethod sets the Method field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveysReportingResponseLinksItem) SetMethod(method *ListSurveysReportingResponseLinksItemMethod) {
	l.Method = method
	l.require(listSurveysReportingResponseLinksItemFieldMethod)
}

// SetRel sets the Rel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveysReportingResponseLinksItem) SetRel(rel *string) {
	l.Rel = rel
	l.require(listSurveysReportingResponseLinksItemFieldRel)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveysReportingResponseLinksItem) SetSchema(schema *string) {
	l.Schema = schema
	l.require(listSurveysReportingResponseLinksItemFieldSchema)
}

// SetTargetSchema sets the TargetSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveysReportingResponseLinksItem) SetTargetSchema(targetSchema *string) {
	l.TargetSchema = targetSchema
	l.require(listSurveysReportingResponseLinksItemFieldTargetSchema)
}

func (l *ListSurveysReportingResponseLinksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler ListSurveysReportingResponseLinksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListSurveysReportingResponseLinksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListSurveysReportingResponseLinksItem) MarshalJSON() ([]byte, error) {
	type embed ListSurveysReportingResponseLinksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListSurveysReportingResponseLinksItem) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// The HTTP method that should be used when accessing the URL defined in 'href'.
type ListSurveysReportingResponseLinksItemMethod string

const (
	ListSurveysReportingResponseLinksItemMethodGet     ListSurveysReportingResponseLinksItemMethod = "GET"
	ListSurveysReportingResponseLinksItemMethodPost    ListSurveysReportingResponseLinksItemMethod = "POST"
	ListSurveysReportingResponseLinksItemMethodPut     ListSurveysReportingResponseLinksItemMethod = "PUT"
	ListSurveysReportingResponseLinksItemMethodPatch   ListSurveysReportingResponseLinksItemMethod = "PATCH"
	ListSurveysReportingResponseLinksItemMethodDelete  ListSurveysReportingResponseLinksItemMethod = "DELETE"
	ListSurveysReportingResponseLinksItemMethodOptions ListSurveysReportingResponseLinksItemMethod = "OPTIONS"
	ListSurveysReportingResponseLinksItemMethodHead    ListSurveysReportingResponseLinksItemMethod = "HEAD"
)

func NewListSurveysReportingResponseLinksItemMethodFromString(s string) (ListSurveysReportingResponseLinksItemMethod, error) {
	switch s {
	case "GET":
		return ListSurveysReportingResponseLinksItemMethodGet, nil
	case "POST":
		return ListSurveysReportingResponseLinksItemMethodPost, nil
	case "PUT":
		return ListSurveysReportingResponseLinksItemMethodPut, nil
	case "PATCH":
		return ListSurveysReportingResponseLinksItemMethodPatch, nil
	case "DELETE":
		return ListSurveysReportingResponseLinksItemMethodDelete, nil
	case "OPTIONS":
		return ListSurveysReportingResponseLinksItemMethodOptions, nil
	case "HEAD":
		return ListSurveysReportingResponseLinksItemMethodHead, nil
	}
	var t ListSurveysReportingResponseLinksItemMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListSurveysReportingResponseLinksItemMethod) Ptr() *ListSurveysReportingResponseLinksItemMethod {
	return &l
}

// The report for a survey.
var (
	listSurveysReportingResponseSurveysItemFieldCreatedAt      = big.NewInt(1 << 0)
	listSurveysReportingResponseSurveysItemFieldID             = big.NewInt(1 << 1)
	listSurveysReportingResponseSurveysItemFieldListID         = big.NewInt(1 << 2)
	listSurveysReportingResponseSurveysItemFieldListName       = big.NewInt(1 << 3)
	listSurveysReportingResponseSurveysItemFieldPublishedAt    = big.NewInt(1 << 4)
	listSurveysReportingResponseSurveysItemFieldStatus         = big.NewInt(1 << 5)
	listSurveysReportingResponseSurveysItemFieldTitle          = big.NewInt(1 << 6)
	listSurveysReportingResponseSurveysItemFieldTotalResponses = big.NewInt(1 << 7)
	listSurveysReportingResponseSurveysItemFieldUpdatedAt      = big.NewInt(1 << 8)
	listSurveysReportingResponseSurveysItemFieldURL            = big.NewInt(1 << 9)
	listSurveysReportingResponseSurveysItemFieldWebID          = big.NewInt(1 << 10)
)

type ListSurveysReportingResponseSurveysItem struct {
	// The date and time the survey was created in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty" url:"created_at,omitempty"`
	// A string that uniquely identifies this survey.
	ID *string `json:"id,omitempty" url:"id,omitempty"`
	// The ID of the list connected to this survey.
	ListID *string `json:"list_id,omitempty" url:"list_id,omitempty"`
	// The name of the list connected to this survey.
	ListName *string `json:"list_name,omitempty" url:"list_name,omitempty"`
	// The date and time the survey was published in ISO 8601 format.
	PublishedAt *time.Time `json:"published_at,omitempty" url:"published_at,omitempty"`
	// The survey's status.
	Status *ListSurveysReportingResponseSurveysItemStatus `json:"status,omitempty" url:"status,omitempty"`
	// The title of the survey.
	Title *string `json:"title,omitempty" url:"title,omitempty"`
	// The total number of responses to this survey.
	TotalResponses *int `json:"total_responses,omitempty" url:"total_responses,omitempty"`
	// The date and time the survey was last updated in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty" url:"updated_at,omitempty"`
	// The URL for the survey.
	URL *string `json:"url,omitempty" url:"url,omitempty"`
	// The ID used in the Mailchimp web application. View this survey report in your Mailchimp account at `https://{dc}.admin.mailchimp.com/lists/surveys/results?survey_id={web_id}`.
	WebID *int `json:"web_id,omitempty" url:"web_id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListSurveysReportingResponseSurveysItem) GetCreatedAt() *time.Time {
	if l == nil {
		return nil
	}
	return l.CreatedAt
}

func (l *ListSurveysReportingResponseSurveysItem) GetID() *string {
	if l == nil {
		return nil
	}
	return l.ID
}

func (l *ListSurveysReportingResponseSurveysItem) GetListID() *string {
	if l == nil {
		return nil
	}
	return l.ListID
}

func (l *ListSurveysReportingResponseSurveysItem) GetListName() *string {
	if l == nil {
		return nil
	}
	return l.ListName
}

func (l *ListSurveysReportingResponseSurveysItem) GetPublishedAt() *time.Time {
	if l == nil {
		return nil
	}
	return l.PublishedAt
}

func (l *ListSurveysReportingResponseSurveysItem) GetStatus() *ListSurveysReportingResponseSurveysItemStatus {
	if l == nil {
		return nil
	}
	return l.Status
}

func (l *ListSurveysReportingResponseSurveysItem) GetTitle() *string {
	if l == nil {
		return nil
	}
	return l.Title
}

func (l *ListSurveysReportingResponseSurveysItem) GetTotalResponses() *int {
	if l == nil {
		return nil
	}
	return l.TotalResponses
}

func (l *ListSurveysReportingResponseSurveysItem) GetUpdatedAt() *time.Time {
	if l == nil {
		return nil
	}
	return l.UpdatedAt
}

func (l *ListSurveysReportingResponseSurveysItem) GetURL() *string {
	if l == nil {
		return nil
	}
	return l.URL
}

func (l *ListSurveysReportingResponseSurveysItem) GetWebID() *int {
	if l == nil {
		return nil
	}
	return l.WebID
}

func (l *ListSurveysReportingResponseSurveysItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListSurveysReportingResponseSurveysItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetCreatedAt sets the CreatedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveysReportingResponseSurveysItem) SetCreatedAt(createdAt *time.Time) {
	l.CreatedAt = createdAt
	l.require(listSurveysReportingResponseSurveysItemFieldCreatedAt)
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveysReportingResponseSurveysItem) SetID(id *string) {
	l.ID = id
	l.require(listSurveysReportingResponseSurveysItemFieldID)
}

// SetListID sets the ListID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveysReportingResponseSurveysItem) SetListID(listID *string) {
	l.ListID = listID
	l.require(listSurveysReportingResponseSurveysItemFieldListID)
}

// SetListName sets the ListName field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveysReportingResponseSurveysItem) SetListName(listName *string) {
	l.ListName = listName
	l.require(listSurveysReportingResponseSurveysItemFieldListName)
}

// SetPublishedAt sets the PublishedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveysReportingResponseSurveysItem) SetPublishedAt(publishedAt *time.Time) {
	l.PublishedAt = publishedAt
	l.require(listSurveysReportingResponseSurveysItemFieldPublishedAt)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveysReportingResponseSurveysItem) SetStatus(status *ListSurveysReportingResponseSurveysItemStatus) {
	l.Status = status
	l.require(listSurveysReportingResponseSurveysItemFieldStatus)
}

// SetTitle sets the Title field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveysReportingResponseSurveysItem) SetTitle(title *string) {
	l.Title = title
	l.require(listSurveysReportingResponseSurveysItemFieldTitle)
}

// SetTotalResponses sets the TotalResponses field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveysReportingResponseSurveysItem) SetTotalResponses(totalResponses *int) {
	l.TotalResponses = totalResponses
	l.require(listSurveysReportingResponseSurveysItemFieldTotalResponses)
}

// SetUpdatedAt sets the UpdatedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveysReportingResponseSurveysItem) SetUpdatedAt(updatedAt *time.Time) {
	l.UpdatedAt = updatedAt
	l.require(listSurveysReportingResponseSurveysItemFieldUpdatedAt)
}

// SetURL sets the URL field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveysReportingResponseSurveysItem) SetURL(url *string) {
	l.URL = url
	l.require(listSurveysReportingResponseSurveysItemFieldURL)
}

// SetWebID sets the WebID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSurveysReportingResponseSurveysItem) SetWebID(webID *int) {
	l.WebID = webID
	l.require(listSurveysReportingResponseSurveysItemFieldWebID)
}

func (l *ListSurveysReportingResponseSurveysItem) UnmarshalJSON(data []byte) error {
	type embed ListSurveysReportingResponseSurveysItem
	var unmarshaler = struct {
		embed
		CreatedAt   *internal.DateTime `json:"created_at,omitempty"`
		PublishedAt *internal.DateTime `json:"published_at,omitempty"`
		UpdatedAt   *internal.DateTime `json:"updated_at,omitempty"`
	}{
		embed: embed(*l),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*l = ListSurveysReportingResponseSurveysItem(unmarshaler.embed)
	l.CreatedAt = unmarshaler.CreatedAt.TimePtr()
	l.PublishedAt = unmarshaler.PublishedAt.TimePtr()
	l.UpdatedAt = unmarshaler.UpdatedAt.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListSurveysReportingResponseSurveysItem) MarshalJSON() ([]byte, error) {
	type embed ListSurveysReportingResponseSurveysItem
	var marshaler = struct {
		embed
		CreatedAt   *internal.DateTime `json:"created_at,omitempty"`
		PublishedAt *internal.DateTime `json:"published_at,omitempty"`
		UpdatedAt   *internal.DateTime `json:"updated_at,omitempty"`
	}{
		embed:       embed(*l),
		CreatedAt:   internal.NewOptionalDateTime(l.CreatedAt),
		PublishedAt: internal.NewOptionalDateTime(l.PublishedAt),
		UpdatedAt:   internal.NewOptionalDateTime(l.UpdatedAt),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListSurveysReportingResponseSurveysItem) String() string {
	if l == nil {
		return "<nil>"
	}
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// The survey's status.
type ListSurveysReportingResponseSurveysItemStatus string

const (
	ListSurveysReportingResponseSurveysItemStatusPublished   ListSurveysReportingResponseSurveysItemStatus = "published"
	ListSurveysReportingResponseSurveysItemStatusUnpublished ListSurveysReportingResponseSurveysItemStatus = "unpublished"
)

func NewListSurveysReportingResponseSurveysItemStatusFromString(s string) (ListSurveysReportingResponseSurveysItemStatus, error) {
	switch s {
	case "published":
		return ListSurveysReportingResponseSurveysItemStatusPublished, nil
	case "unpublished":
		return ListSurveysReportingResponseSurveysItemStatusUnpublished, nil
	}
	var t ListSurveysReportingResponseSurveysItemStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListSurveysReportingResponseSurveysItemStatus) Ptr() *ListSurveysReportingResponseSurveysItemStatus {
	return &l
}
