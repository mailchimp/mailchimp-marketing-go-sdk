// Code generated from our API definition. DO NOT EDIT.

package api

import (
	big "math/big"
)

var (
	createListSurveyActionCreateEmailSurveysRequestFieldListID   = big.NewInt(1 << 0)
	createListSurveyActionCreateEmailSurveysRequestFieldSurveyID = big.NewInt(1 << 1)
)

type CreateListSurveyActionCreateEmailSurveysRequest struct {
	// The unique ID for the list.
	ListID string `json:"-" url:"-"`
	// The ID of the survey.
	SurveyID string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CreateListSurveyActionCreateEmailSurveysRequest) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetListID sets the ListID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateListSurveyActionCreateEmailSurveysRequest) SetListID(listID string) {
	c.ListID = listID
	c.require(createListSurveyActionCreateEmailSurveysRequestFieldListID)
}

// SetSurveyID sets the SurveyID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateListSurveyActionCreateEmailSurveysRequest) SetSurveyID(surveyID string) {
	c.SurveyID = surveyID
	c.require(createListSurveyActionCreateEmailSurveysRequestFieldSurveyID)
}

var (
	createListSurveyActionPublishSurveysRequestFieldListID   = big.NewInt(1 << 0)
	createListSurveyActionPublishSurveysRequestFieldSurveyID = big.NewInt(1 << 1)
)

type CreateListSurveyActionPublishSurveysRequest struct {
	// The unique ID for the list.
	ListID string `json:"-" url:"-"`
	// The ID of the survey.
	SurveyID string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CreateListSurveyActionPublishSurveysRequest) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetListID sets the ListID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateListSurveyActionPublishSurveysRequest) SetListID(listID string) {
	c.ListID = listID
	c.require(createListSurveyActionPublishSurveysRequestFieldListID)
}

// SetSurveyID sets the SurveyID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateListSurveyActionPublishSurveysRequest) SetSurveyID(surveyID string) {
	c.SurveyID = surveyID
	c.require(createListSurveyActionPublishSurveysRequestFieldSurveyID)
}

var (
	createListSurveyActionUnpublishSurveysRequestFieldListID   = big.NewInt(1 << 0)
	createListSurveyActionUnpublishSurveysRequestFieldSurveyID = big.NewInt(1 << 1)
)

type CreateListSurveyActionUnpublishSurveysRequest struct {
	// The unique ID for the list.
	ListID string `json:"-" url:"-"`
	// The ID of the survey.
	SurveyID string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CreateListSurveyActionUnpublishSurveysRequest) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetListID sets the ListID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateListSurveyActionUnpublishSurveysRequest) SetListID(listID string) {
	c.ListID = listID
	c.require(createListSurveyActionUnpublishSurveysRequestFieldListID)
}

// SetSurveyID sets the SurveyID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateListSurveyActionUnpublishSurveysRequest) SetSurveyID(surveyID string) {
	c.SurveyID = surveyID
	c.require(createListSurveyActionUnpublishSurveysRequestFieldSurveyID)
}
