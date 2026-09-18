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
	createAudienceContactRequestFieldAudienceID               = big.NewInt(1 << 0)
	createAudienceContactRequestFieldMergeFieldValidationMode = big.NewInt(1 << 1)
	createAudienceContactRequestFieldDataMode                 = big.NewInt(1 << 2)
	createAudienceContactRequestFieldEmailChannel             = big.NewInt(1 << 3)
	createAudienceContactRequestFieldLanguage                 = big.NewInt(1 << 4)
	createAudienceContactRequestFieldMergeFields              = big.NewInt(1 << 5)
	createAudienceContactRequestFieldSmsChannel               = big.NewInt(1 << 6)
	createAudienceContactRequestFieldTags                     = big.NewInt(1 << 7)
	createAudienceContactRequestFieldUpdateExisting           = big.NewInt(1 << 8)
)

type CreateAudienceContactRequest struct {
	// The unique ID for the audience.
	AudienceID string `json:"-" url:"-"`
	// Defines how merge field validation is handled. When set to `ignore_required_checks`, the API does not raise an error if required merge fields are missing from the request. When set to `strict`, the API enforces validation and returns an error if any required merge field is not provided. If this setting is omitted, `strict` is applied by default.
	MergeFieldValidationMode *CreateAudienceContactRequestMergeFieldValidationMode `json:"-" url:"merge_field_validation_mode,omitempty"`
	// Indicates the data processing mode. In `historical` mode, contact data changes do not trigger automations or webhooks. In `live mode`, such changes do trigger them.
	DataMode     *CreateAudienceContactRequestDataMode     `json:"-" url:"data_mode,omitempty"`
	EmailChannel *CreateAudienceContactRequestEmailChannel `json:"email_channel,omitempty" url:"-"`
	// The contact's detected language.
	Language *string `json:"language,omitempty" url:"-"`
	// A dictionary of merge fields where the keys are the merge tags. See the [Merge Fields documentation](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for more about the structure.
	MergeFields map[string]*CreateAudienceContactRequestMergeFieldsValue `json:"merge_fields,omitempty" url:"-"`
	SmsChannel  *CreateAudienceContactRequestSmsChannel                  `json:"sms_channel,omitempty" url:"-"`
	// An array of tags to add to the contact. Accepts tag name strings or objects with name and status. This operation is append-only; existing tags will be preserved, and only new tags from this array will be added.
	Tags []*CreateAudienceContactRequestTagsItem `json:"tags,omitempty" url:"-"`
	// If a contact already exists, update them instead of returning a conflict error. When `true` and a matching contact is found (by email or phone), the existing contact is updated with the provided channel data. Defaults to `false`.
	UpdateExisting *bool `json:"update_existing,omitempty" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CreateAudienceContactRequest) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetAudienceID sets the AudienceID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAudienceContactRequest) SetAudienceID(audienceID string) {
	c.AudienceID = audienceID
	c.require(createAudienceContactRequestFieldAudienceID)
}

// SetMergeFieldValidationMode sets the MergeFieldValidationMode field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAudienceContactRequest) SetMergeFieldValidationMode(mergeFieldValidationMode *CreateAudienceContactRequestMergeFieldValidationMode) {
	c.MergeFieldValidationMode = mergeFieldValidationMode
	c.require(createAudienceContactRequestFieldMergeFieldValidationMode)
}

// SetDataMode sets the DataMode field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAudienceContactRequest) SetDataMode(dataMode *CreateAudienceContactRequestDataMode) {
	c.DataMode = dataMode
	c.require(createAudienceContactRequestFieldDataMode)
}

// SetEmailChannel sets the EmailChannel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAudienceContactRequest) SetEmailChannel(emailChannel *CreateAudienceContactRequestEmailChannel) {
	c.EmailChannel = emailChannel
	c.require(createAudienceContactRequestFieldEmailChannel)
}

// SetLanguage sets the Language field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAudienceContactRequest) SetLanguage(language *string) {
	c.Language = language
	c.require(createAudienceContactRequestFieldLanguage)
}

// SetMergeFields sets the MergeFields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAudienceContactRequest) SetMergeFields(mergeFields map[string]*CreateAudienceContactRequestMergeFieldsValue) {
	c.MergeFields = mergeFields
	c.require(createAudienceContactRequestFieldMergeFields)
}

// SetSmsChannel sets the SmsChannel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAudienceContactRequest) SetSmsChannel(smsChannel *CreateAudienceContactRequestSmsChannel) {
	c.SmsChannel = smsChannel
	c.require(createAudienceContactRequestFieldSmsChannel)
}

// SetTags sets the Tags field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAudienceContactRequest) SetTags(tags []*CreateAudienceContactRequestTagsItem) {
	c.Tags = tags
	c.require(createAudienceContactRequestFieldTags)
}

// SetUpdateExisting sets the UpdateExisting field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAudienceContactRequest) SetUpdateExisting(updateExisting *bool) {
	c.UpdateExisting = updateExisting
	c.require(createAudienceContactRequestFieldUpdateExisting)
}

func (c *CreateAudienceContactRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateAudienceContactRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*c = CreateAudienceContactRequest(body)
	return nil
}

func (c *CreateAudienceContactRequest) MarshalJSON() ([]byte, error) {
	type embed CreateAudienceContactRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

var (
	getAudienceContactRequestFieldAudienceID    = big.NewInt(1 << 0)
	getAudienceContactRequestFieldContactID     = big.NewInt(1 << 1)
	getAudienceContactRequestFieldFields        = big.NewInt(1 << 2)
	getAudienceContactRequestFieldExcludeFields = big.NewInt(1 << 3)
)

type GetAudienceContactRequest struct {
	// The unique ID for the audience.
	AudienceID string `json:"-" url:"-"`
	// A unique identifier for the contact, which can be a Mailchimp contact ID or a channel hash. A channel hash must follow the format email:[md5_hash] (where the hash is the MD5 of the lowercased email address) or sms:[sha256_hash] (where the hash is the SHA256 of the E.164-formatted phone number).
	ContactID string `json:"-" url:"-"`
	// A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
	Fields []*string `json:"-" url:"fields,omitempty"`
	// A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
	ExcludeFields []*string `json:"-" url:"exclude_fields,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetAudienceContactRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetAudienceID sets the AudienceID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAudienceContactRequest) SetAudienceID(audienceID string) {
	g.AudienceID = audienceID
	g.require(getAudienceContactRequestFieldAudienceID)
}

// SetContactID sets the ContactID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAudienceContactRequest) SetContactID(contactID string) {
	g.ContactID = contactID
	g.require(getAudienceContactRequestFieldContactID)
}

// SetFields sets the Fields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAudienceContactRequest) SetFields(fields []*string) {
	g.Fields = fields
	g.require(getAudienceContactRequestFieldFields)
}

// SetExcludeFields sets the ExcludeFields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAudienceContactRequest) SetExcludeFields(excludeFields []*string) {
	g.ExcludeFields = excludeFields
	g.require(getAudienceContactRequestFieldExcludeFields)
}

var (
	getAudienceContactListRequestFieldAudienceID    = big.NewInt(1 << 0)
	getAudienceContactListRequestFieldFields        = big.NewInt(1 << 1)
	getAudienceContactListRequestFieldExcludeFields = big.NewInt(1 << 2)
	getAudienceContactListRequestFieldCount         = big.NewInt(1 << 3)
	getAudienceContactListRequestFieldCursor        = big.NewInt(1 << 4)
	getAudienceContactListRequestFieldCreatedBefore = big.NewInt(1 << 5)
	getAudienceContactListRequestFieldCreatedSince  = big.NewInt(1 << 6)
	getAudienceContactListRequestFieldUpdatedBefore = big.NewInt(1 << 7)
	getAudienceContactListRequestFieldUpdatedSince  = big.NewInt(1 << 8)
	getAudienceContactListRequestFieldSortField     = big.NewInt(1 << 9)
	getAudienceContactListRequestFieldSortDir       = big.NewInt(1 << 10)
)

type GetAudienceContactListRequest struct {
	// The unique ID for the audience.
	AudienceID string `json:"-" url:"-"`
	// A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
	Fields []*string `json:"-" url:"fields,omitempty"`
	// A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
	ExcludeFields []*string `json:"-" url:"exclude_fields,omitempty"`
	// The number of records to return. Default value is 10. Maximum value is 1000
	Count *int `json:"-" url:"count,omitempty"`
	// Paginate through a collection of records by setting the `cursor` parameter to a `next_cursor` attribute returned by a previous request. Default value fetches the first "page" of results.
	Cursor *string `json:"-" url:"cursor,omitempty"`
	// Restricts the response to contacts created at or before the specified time (inclusive). Uses ISO 8601 format: 2025-04-23T15:41:36+00:00.
	CreatedBefore *time.Time `json:"-" url:"created_before,omitempty"`
	// Restricts the response to contacts created after the specified time (exclusive). Uses ISO 8601 format: 2025-04-23T15:41:36+00:00.
	CreatedSince *time.Time `json:"-" url:"created_since,omitempty"`
	// Restricts the response to contacts updated at or before the specified time (inclusive). Uses ISO 8601 format: 2025-04-23T15:41:36+00:00.
	UpdatedBefore *time.Time `json:"-" url:"updated_before,omitempty"`
	// Restricts the response to contacts updated after the specified time (exclusive). Uses ISO 8601 format: 2025-04-23T15:41:36+00:00.
	UpdatedSince *time.Time `json:"-" url:"updated_since,omitempty"`
	// Specifies the field to sort the returned contacts by.
	SortField *GetAudienceContactListRequestSortField `json:"-" url:"sort_field,omitempty"`
	// Determines the order direction for sorted results.
	SortDir *GetAudienceContactListRequestSortDir `json:"-" url:"sort_dir,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetAudienceContactListRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetAudienceID sets the AudienceID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAudienceContactListRequest) SetAudienceID(audienceID string) {
	g.AudienceID = audienceID
	g.require(getAudienceContactListRequestFieldAudienceID)
}

// SetFields sets the Fields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAudienceContactListRequest) SetFields(fields []*string) {
	g.Fields = fields
	g.require(getAudienceContactListRequestFieldFields)
}

// SetExcludeFields sets the ExcludeFields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAudienceContactListRequest) SetExcludeFields(excludeFields []*string) {
	g.ExcludeFields = excludeFields
	g.require(getAudienceContactListRequestFieldExcludeFields)
}

// SetCount sets the Count field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAudienceContactListRequest) SetCount(count *int) {
	g.Count = count
	g.require(getAudienceContactListRequestFieldCount)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAudienceContactListRequest) SetCursor(cursor *string) {
	g.Cursor = cursor
	g.require(getAudienceContactListRequestFieldCursor)
}

// SetCreatedBefore sets the CreatedBefore field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAudienceContactListRequest) SetCreatedBefore(createdBefore *time.Time) {
	g.CreatedBefore = createdBefore
	g.require(getAudienceContactListRequestFieldCreatedBefore)
}

// SetCreatedSince sets the CreatedSince field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAudienceContactListRequest) SetCreatedSince(createdSince *time.Time) {
	g.CreatedSince = createdSince
	g.require(getAudienceContactListRequestFieldCreatedSince)
}

// SetUpdatedBefore sets the UpdatedBefore field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAudienceContactListRequest) SetUpdatedBefore(updatedBefore *time.Time) {
	g.UpdatedBefore = updatedBefore
	g.require(getAudienceContactListRequestFieldUpdatedBefore)
}

// SetUpdatedSince sets the UpdatedSince field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAudienceContactListRequest) SetUpdatedSince(updatedSince *time.Time) {
	g.UpdatedSince = updatedSince
	g.require(getAudienceContactListRequestFieldUpdatedSince)
}

// SetSortField sets the SortField field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAudienceContactListRequest) SetSortField(sortField *GetAudienceContactListRequestSortField) {
	g.SortField = sortField
	g.require(getAudienceContactListRequestFieldSortField)
}

// SetSortDir sets the SortDir field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAudienceContactListRequest) SetSortDir(sortDir *GetAudienceContactListRequestSortDir) {
	g.SortDir = sortDir
	g.require(getAudienceContactListRequestFieldSortDir)
}

var (
	patchAudienceContactRequestFieldAudienceID               = big.NewInt(1 << 0)
	patchAudienceContactRequestFieldContactID                = big.NewInt(1 << 1)
	patchAudienceContactRequestFieldMergeFieldValidationMode = big.NewInt(1 << 2)
	patchAudienceContactRequestFieldDataMode                 = big.NewInt(1 << 3)
	patchAudienceContactRequestFieldEmailChannel             = big.NewInt(1 << 4)
	patchAudienceContactRequestFieldLanguage                 = big.NewInt(1 << 5)
	patchAudienceContactRequestFieldMergeFields              = big.NewInt(1 << 6)
	patchAudienceContactRequestFieldSmsChannel               = big.NewInt(1 << 7)
	patchAudienceContactRequestFieldTags                     = big.NewInt(1 << 8)
)

type PatchAudienceContactRequest struct {
	// The unique ID for the audience.
	AudienceID string `json:"-" url:"-"`
	// The unique id for the contact.
	ContactID string `json:"-" url:"-"`
	// Defines how merge field validation is handled. When set to `ignore_required_checks`, the API does not raise an error if required merge fields are missing from the request. When set to `strict`, the API enforces validation and returns an error if any required merge field is not provided. If this setting is omitted, `strict` is applied by default.
	MergeFieldValidationMode *PatchAudienceContactRequestMergeFieldValidationMode `json:"-" url:"merge_field_validation_mode,omitempty"`
	// Indicates the data processing mode. In `historical` mode, contact data changes do not trigger automations or webhooks. In `live mode`, such changes do trigger them.
	DataMode     *PatchAudienceContactRequestDataMode     `json:"-" url:"data_mode,omitempty"`
	EmailChannel *PatchAudienceContactRequestEmailChannel `json:"email_channel,omitempty" url:"-"`
	// The contact's detected language.
	Language *string `json:"language,omitempty" url:"-"`
	// A dictionary of merge fields where the keys are the merge tags. See the [Merge Fields documentation](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for more about the structure.
	MergeFields map[string]*PatchAudienceContactRequestMergeFieldsValue `json:"merge_fields,omitempty" url:"-"`
	SmsChannel  *PatchAudienceContactRequestSmsChannel                  `json:"sms_channel,omitempty" url:"-"`
	// An array of tags to add to the contact. Accepts tag name strings or objects with name and status. This operation is append-only; existing tags will be preserved, and only new tags from this array will be added.
	Tags []*PatchAudienceContactRequestTagsItem `json:"tags,omitempty" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (p *PatchAudienceContactRequest) require(field *big.Int) {
	if p.explicitFields == nil {
		p.explicitFields = big.NewInt(0)
	}
	p.explicitFields.Or(p.explicitFields, field)
}

// SetAudienceID sets the AudienceID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PatchAudienceContactRequest) SetAudienceID(audienceID string) {
	p.AudienceID = audienceID
	p.require(patchAudienceContactRequestFieldAudienceID)
}

// SetContactID sets the ContactID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PatchAudienceContactRequest) SetContactID(contactID string) {
	p.ContactID = contactID
	p.require(patchAudienceContactRequestFieldContactID)
}

// SetMergeFieldValidationMode sets the MergeFieldValidationMode field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PatchAudienceContactRequest) SetMergeFieldValidationMode(mergeFieldValidationMode *PatchAudienceContactRequestMergeFieldValidationMode) {
	p.MergeFieldValidationMode = mergeFieldValidationMode
	p.require(patchAudienceContactRequestFieldMergeFieldValidationMode)
}

// SetDataMode sets the DataMode field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PatchAudienceContactRequest) SetDataMode(dataMode *PatchAudienceContactRequestDataMode) {
	p.DataMode = dataMode
	p.require(patchAudienceContactRequestFieldDataMode)
}

// SetEmailChannel sets the EmailChannel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PatchAudienceContactRequest) SetEmailChannel(emailChannel *PatchAudienceContactRequestEmailChannel) {
	p.EmailChannel = emailChannel
	p.require(patchAudienceContactRequestFieldEmailChannel)
}

// SetLanguage sets the Language field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PatchAudienceContactRequest) SetLanguage(language *string) {
	p.Language = language
	p.require(patchAudienceContactRequestFieldLanguage)
}

// SetMergeFields sets the MergeFields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PatchAudienceContactRequest) SetMergeFields(mergeFields map[string]*PatchAudienceContactRequestMergeFieldsValue) {
	p.MergeFields = mergeFields
	p.require(patchAudienceContactRequestFieldMergeFields)
}

// SetSmsChannel sets the SmsChannel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PatchAudienceContactRequest) SetSmsChannel(smsChannel *PatchAudienceContactRequestSmsChannel) {
	p.SmsChannel = smsChannel
	p.require(patchAudienceContactRequestFieldSmsChannel)
}

// SetTags sets the Tags field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PatchAudienceContactRequest) SetTags(tags []*PatchAudienceContactRequestTagsItem) {
	p.Tags = tags
	p.require(patchAudienceContactRequestFieldTags)
}

func (p *PatchAudienceContactRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler PatchAudienceContactRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*p = PatchAudienceContactRequest(body)
	return nil
}

func (p *PatchAudienceContactRequest) MarshalJSON() ([]byte, error) {
	type embed PatchAudienceContactRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, p.explicitFields)
	return json.Marshal(explicitMarshaler)
}

var (
	postAudiencesContactsActionsArchiveRequestFieldAudienceID = big.NewInt(1 << 0)
	postAudiencesContactsActionsArchiveRequestFieldContactID  = big.NewInt(1 << 1)
)

type PostAudiencesContactsActionsArchiveRequest struct {
	// The unique ID for the audience.
	AudienceID string `json:"-" url:"-"`
	// The unique id for the contact.
	ContactID string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (p *PostAudiencesContactsActionsArchiveRequest) require(field *big.Int) {
	if p.explicitFields == nil {
		p.explicitFields = big.NewInt(0)
	}
	p.explicitFields.Or(p.explicitFields, field)
}

// SetAudienceID sets the AudienceID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PostAudiencesContactsActionsArchiveRequest) SetAudienceID(audienceID string) {
	p.AudienceID = audienceID
	p.require(postAudiencesContactsActionsArchiveRequestFieldAudienceID)
}

// SetContactID sets the ContactID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PostAudiencesContactsActionsArchiveRequest) SetContactID(contactID string) {
	p.ContactID = contactID
	p.require(postAudiencesContactsActionsArchiveRequestFieldContactID)
}

var (
	postAudiencesContactsActionsForgetRequestFieldAudienceID = big.NewInt(1 << 0)
	postAudiencesContactsActionsForgetRequestFieldContactID  = big.NewInt(1 << 1)
)

type PostAudiencesContactsActionsForgetRequest struct {
	// The unique ID for the audience.
	AudienceID string `json:"-" url:"-"`
	// The unique id for the contact.
	ContactID string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (p *PostAudiencesContactsActionsForgetRequest) require(field *big.Int) {
	if p.explicitFields == nil {
		p.explicitFields = big.NewInt(0)
	}
	p.explicitFields.Or(p.explicitFields, field)
}

// SetAudienceID sets the AudienceID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PostAudiencesContactsActionsForgetRequest) SetAudienceID(audienceID string) {
	p.AudienceID = audienceID
	p.require(postAudiencesContactsActionsForgetRequestFieldAudienceID)
}

// SetContactID sets the ContactID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PostAudiencesContactsActionsForgetRequest) SetContactID(contactID string) {
	p.ContactID = contactID
	p.require(postAudiencesContactsActionsForgetRequestFieldContactID)
}

// An instance of a contact.
var (
	audiencesContactFieldAudienceID    = big.NewInt(1 << 0)
	audiencesContactFieldCreatedAt     = big.NewInt(1 << 1)
	audiencesContactFieldEmailChannel  = big.NewInt(1 << 2)
	audiencesContactFieldID            = big.NewInt(1 << 3)
	audiencesContactFieldLanguage      = big.NewInt(1 << 4)
	audiencesContactFieldLastUpdatedAt = big.NewInt(1 << 5)
	audiencesContactFieldMergeFields   = big.NewInt(1 << 6)
	audiencesContactFieldSmsChannel    = big.NewInt(1 << 7)
	audiencesContactFieldSource        = big.NewInt(1 << 8)
	audiencesContactFieldStatus        = big.NewInt(1 << 9)
	audiencesContactFieldTags          = big.NewInt(1 << 10)
)

type AudiencesContact struct {
	// The unique ID for the audience.
	AudienceID *string `json:"audience_id,omitempty" url:"audience_id,omitempty"`
	// The date that the contact was created.
	CreatedAt    *time.Time                    `json:"created_at,omitempty" url:"created_at,omitempty"`
	EmailChannel *AudiencesContactEmailChannel `json:"email_channel,omitempty" url:"email_channel,omitempty"`
	// The unique ID for the contact.
	ID *string `json:"id,omitempty" url:"id,omitempty"`
	// The contact's detected language. Empty string when no language has been detected or set.
	Language *AudiencesContactLanguage `json:"language,omitempty" url:"language,omitempty"`
	// The date that the contact was last updated.
	LastUpdatedAt *time.Time `json:"last_updated_at,omitempty" url:"last_updated_at,omitempty"`
	// A dictionary of merge fields where the keys are the merge tags. See the [Merge Fields documentation](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for more about the structure.
	MergeFields map[string]*AudiencesContactMergeFieldsValue `json:"merge_fields,omitempty" url:"merge_fields,omitempty"`
	SmsChannel  *AudiencesContactSmsChannel                  `json:"sms_channel,omitempty" url:"sms_channel,omitempty"`
	// The source from which the parent's entity was created.
	Source *AudiencesContactSource `json:"source,omitempty" url:"source,omitempty"`
	// The status of a contact.
	Status *AudiencesContactStatus `json:"status,omitempty" url:"status,omitempty"`
	// The tags assigned to this contact.
	Tags []string `json:"tags,omitempty" url:"tags,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *AudiencesContact) GetAudienceID() *string {
	if a == nil {
		return nil
	}
	return a.AudienceID
}

func (a *AudiencesContact) GetCreatedAt() *time.Time {
	if a == nil {
		return nil
	}
	return a.CreatedAt
}

func (a *AudiencesContact) GetEmailChannel() *AudiencesContactEmailChannel {
	if a == nil {
		return nil
	}
	return a.EmailChannel
}

func (a *AudiencesContact) GetID() *string {
	if a == nil {
		return nil
	}
	return a.ID
}

func (a *AudiencesContact) GetLanguage() *AudiencesContactLanguage {
	if a == nil {
		return nil
	}
	return a.Language
}

func (a *AudiencesContact) GetLastUpdatedAt() *time.Time {
	if a == nil {
		return nil
	}
	return a.LastUpdatedAt
}

func (a *AudiencesContact) GetMergeFields() map[string]*AudiencesContactMergeFieldsValue {
	if a == nil {
		return nil
	}
	return a.MergeFields
}

func (a *AudiencesContact) GetSmsChannel() *AudiencesContactSmsChannel {
	if a == nil {
		return nil
	}
	return a.SmsChannel
}

func (a *AudiencesContact) GetSource() *AudiencesContactSource {
	if a == nil {
		return nil
	}
	return a.Source
}

func (a *AudiencesContact) GetStatus() *AudiencesContactStatus {
	if a == nil {
		return nil
	}
	return a.Status
}

func (a *AudiencesContact) GetTags() []string {
	if a == nil {
		return nil
	}
	return a.Tags
}

func (a *AudiencesContact) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a *AudiencesContact) require(field *big.Int) {
	if a.explicitFields == nil {
		a.explicitFields = big.NewInt(0)
	}
	a.explicitFields.Or(a.explicitFields, field)
}

// SetAudienceID sets the AudienceID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContact) SetAudienceID(audienceID *string) {
	a.AudienceID = audienceID
	a.require(audiencesContactFieldAudienceID)
}

// SetCreatedAt sets the CreatedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContact) SetCreatedAt(createdAt *time.Time) {
	a.CreatedAt = createdAt
	a.require(audiencesContactFieldCreatedAt)
}

// SetEmailChannel sets the EmailChannel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContact) SetEmailChannel(emailChannel *AudiencesContactEmailChannel) {
	a.EmailChannel = emailChannel
	a.require(audiencesContactFieldEmailChannel)
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContact) SetID(id *string) {
	a.ID = id
	a.require(audiencesContactFieldID)
}

// SetLanguage sets the Language field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContact) SetLanguage(language *AudiencesContactLanguage) {
	a.Language = language
	a.require(audiencesContactFieldLanguage)
}

// SetLastUpdatedAt sets the LastUpdatedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContact) SetLastUpdatedAt(lastUpdatedAt *time.Time) {
	a.LastUpdatedAt = lastUpdatedAt
	a.require(audiencesContactFieldLastUpdatedAt)
}

// SetMergeFields sets the MergeFields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContact) SetMergeFields(mergeFields map[string]*AudiencesContactMergeFieldsValue) {
	a.MergeFields = mergeFields
	a.require(audiencesContactFieldMergeFields)
}

// SetSmsChannel sets the SmsChannel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContact) SetSmsChannel(smsChannel *AudiencesContactSmsChannel) {
	a.SmsChannel = smsChannel
	a.require(audiencesContactFieldSmsChannel)
}

// SetSource sets the Source field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContact) SetSource(source *AudiencesContactSource) {
	a.Source = source
	a.require(audiencesContactFieldSource)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContact) SetStatus(status *AudiencesContactStatus) {
	a.Status = status
	a.require(audiencesContactFieldStatus)
}

// SetTags sets the Tags field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContact) SetTags(tags []string) {
	a.Tags = tags
	a.require(audiencesContactFieldTags)
}

func (a *AudiencesContact) UnmarshalJSON(data []byte) error {
	type embed AudiencesContact
	var unmarshaler = struct {
		embed
		CreatedAt     *internal.DateTime `json:"created_at,omitempty"`
		LastUpdatedAt *internal.DateTime `json:"last_updated_at,omitempty"`
	}{
		embed: embed(*a),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*a = AudiencesContact(unmarshaler.embed)
	a.CreatedAt = unmarshaler.CreatedAt.TimePtr()
	a.LastUpdatedAt = unmarshaler.LastUpdatedAt.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *AudiencesContact) MarshalJSON() ([]byte, error) {
	type embed AudiencesContact
	var marshaler = struct {
		embed
		CreatedAt     *internal.DateTime `json:"created_at,omitempty"`
		LastUpdatedAt *internal.DateTime `json:"last_updated_at,omitempty"`
	}{
		embed:         embed(*a),
		CreatedAt:     internal.NewOptionalDateTime(a.CreatedAt),
		LastUpdatedAt: internal.NewOptionalDateTime(a.LastUpdatedAt),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, a.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (a *AudiencesContact) String() string {
	if a == nil {
		return "<nil>"
	}
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

var (
	audiencesContactEmailChannelFieldEffectiveSubscriptionStatus = big.NewInt(1 << 0)
	audiencesContactEmailChannelFieldEmail                       = big.NewInt(1 << 1)
	audiencesContactEmailChannelFieldHashedEmail                 = big.NewInt(1 << 2)
	audiencesContactEmailChannelFieldMarketingConsent            = big.NewInt(1 << 3)
	audiencesContactEmailChannelFieldSource                      = big.NewInt(1 << 4)
)

type AudiencesContactEmailChannel struct {
	// A computation performed by the Mailchimp platform, triggered whenever any of its inputs change. Some inputs are controlled by API users, while others are tracked internally by the platform. Computation is based on: audience opt-in configuration (single vs. double opt-in), marketing consent status, and deliverability status (an internal state for a contact, maintained by Mailchimp for a specific marketing channel instance). This new API field is distinct from how contacts are displayed in the UI. See the [Audiences (BETA) documentation](https://mailchimp.com/developer/marketing/docs/audiences-introduction) to learn about supported values.
	EffectiveSubscriptionStatus *AudiencesContactEmailChannelEffectiveSubscriptionStatus `json:"effective_subscription_status,omitempty" url:"effective_subscription_status,omitempty"`
	// Email address
	Email *string `json:"email,omitempty" url:"email,omitempty"`
	// MD5 hash of the email address
	HashedEmail *string `json:"hashed_email,omitempty" url:"hashed_email,omitempty"`
	// A contact's current consent status for email marketing communications. See the [Audiences (BETA) documentation](https://mailchimp.com/developer/marketing/docs/audiences-introduction) to learn about supported values.
	MarketingConsent *AudiencesContactEmailChannelMarketingConsent `json:"marketing_consent,omitempty" url:"marketing_consent,omitempty"`
	// The source from which the parent's entity was created.
	Source *AudiencesContactEmailChannelSource `json:"source,omitempty" url:"source,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *AudiencesContactEmailChannel) GetEffectiveSubscriptionStatus() *AudiencesContactEmailChannelEffectiveSubscriptionStatus {
	if a == nil {
		return nil
	}
	return a.EffectiveSubscriptionStatus
}

func (a *AudiencesContactEmailChannel) GetEmail() *string {
	if a == nil {
		return nil
	}
	return a.Email
}

func (a *AudiencesContactEmailChannel) GetHashedEmail() *string {
	if a == nil {
		return nil
	}
	return a.HashedEmail
}

func (a *AudiencesContactEmailChannel) GetMarketingConsent() *AudiencesContactEmailChannelMarketingConsent {
	if a == nil {
		return nil
	}
	return a.MarketingConsent
}

func (a *AudiencesContactEmailChannel) GetSource() *AudiencesContactEmailChannelSource {
	if a == nil {
		return nil
	}
	return a.Source
}

func (a *AudiencesContactEmailChannel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a *AudiencesContactEmailChannel) require(field *big.Int) {
	if a.explicitFields == nil {
		a.explicitFields = big.NewInt(0)
	}
	a.explicitFields.Or(a.explicitFields, field)
}

// SetEffectiveSubscriptionStatus sets the EffectiveSubscriptionStatus field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContactEmailChannel) SetEffectiveSubscriptionStatus(effectiveSubscriptionStatus *AudiencesContactEmailChannelEffectiveSubscriptionStatus) {
	a.EffectiveSubscriptionStatus = effectiveSubscriptionStatus
	a.require(audiencesContactEmailChannelFieldEffectiveSubscriptionStatus)
}

// SetEmail sets the Email field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContactEmailChannel) SetEmail(email *string) {
	a.Email = email
	a.require(audiencesContactEmailChannelFieldEmail)
}

// SetHashedEmail sets the HashedEmail field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContactEmailChannel) SetHashedEmail(hashedEmail *string) {
	a.HashedEmail = hashedEmail
	a.require(audiencesContactEmailChannelFieldHashedEmail)
}

// SetMarketingConsent sets the MarketingConsent field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContactEmailChannel) SetMarketingConsent(marketingConsent *AudiencesContactEmailChannelMarketingConsent) {
	a.MarketingConsent = marketingConsent
	a.require(audiencesContactEmailChannelFieldMarketingConsent)
}

// SetSource sets the Source field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContactEmailChannel) SetSource(source *AudiencesContactEmailChannelSource) {
	a.Source = source
	a.require(audiencesContactEmailChannelFieldSource)
}

func (a *AudiencesContactEmailChannel) UnmarshalJSON(data []byte) error {
	type unmarshaler AudiencesContactEmailChannel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AudiencesContactEmailChannel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *AudiencesContactEmailChannel) MarshalJSON() ([]byte, error) {
	type embed AudiencesContactEmailChannel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, a.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (a *AudiencesContactEmailChannel) String() string {
	if a == nil {
		return "<nil>"
	}
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

// A computation performed by the Mailchimp platform, triggered whenever any of its inputs change. Some inputs are controlled by API users, while others are tracked internally by the platform. Computation is based on: audience opt-in configuration (single vs. double opt-in), marketing consent status, and deliverability status (an internal state for a contact, maintained by Mailchimp for a specific marketing channel instance). This new API field is distinct from how contacts are displayed in the UI. See the [Audiences (BETA) documentation](https://mailchimp.com/developer/marketing/docs/audiences-introduction) to learn about supported values.
var (
	audiencesContactEmailChannelEffectiveSubscriptionStatusFieldValue = big.NewInt(1 << 0)
)

type AudiencesContactEmailChannelEffectiveSubscriptionStatus struct {
	Value *AudiencesContactEmailChannelEffectiveSubscriptionStatusValue `json:"value,omitempty" url:"value,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *AudiencesContactEmailChannelEffectiveSubscriptionStatus) GetValue() *AudiencesContactEmailChannelEffectiveSubscriptionStatusValue {
	if a == nil {
		return nil
	}
	return a.Value
}

func (a *AudiencesContactEmailChannelEffectiveSubscriptionStatus) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a *AudiencesContactEmailChannelEffectiveSubscriptionStatus) require(field *big.Int) {
	if a.explicitFields == nil {
		a.explicitFields = big.NewInt(0)
	}
	a.explicitFields.Or(a.explicitFields, field)
}

// SetValue sets the Value field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContactEmailChannelEffectiveSubscriptionStatus) SetValue(value *AudiencesContactEmailChannelEffectiveSubscriptionStatusValue) {
	a.Value = value
	a.require(audiencesContactEmailChannelEffectiveSubscriptionStatusFieldValue)
}

func (a *AudiencesContactEmailChannelEffectiveSubscriptionStatus) UnmarshalJSON(data []byte) error {
	type unmarshaler AudiencesContactEmailChannelEffectiveSubscriptionStatus
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AudiencesContactEmailChannelEffectiveSubscriptionStatus(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *AudiencesContactEmailChannelEffectiveSubscriptionStatus) MarshalJSON() ([]byte, error) {
	type embed AudiencesContactEmailChannelEffectiveSubscriptionStatus
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, a.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (a *AudiencesContactEmailChannelEffectiveSubscriptionStatus) String() string {
	if a == nil {
		return "<nil>"
	}
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type AudiencesContactEmailChannelEffectiveSubscriptionStatusValue string

const (
	AudiencesContactEmailChannelEffectiveSubscriptionStatusValueSubscribed    AudiencesContactEmailChannelEffectiveSubscriptionStatusValue = "subscribed"
	AudiencesContactEmailChannelEffectiveSubscriptionStatusValueUnsubscribed  AudiencesContactEmailChannelEffectiveSubscriptionStatusValue = "unsubscribed"
	AudiencesContactEmailChannelEffectiveSubscriptionStatusValueNonsubscribed AudiencesContactEmailChannelEffectiveSubscriptionStatusValue = "nonsubscribed"
	AudiencesContactEmailChannelEffectiveSubscriptionStatusValuePending       AudiencesContactEmailChannelEffectiveSubscriptionStatusValue = "pending"
)

func NewAudiencesContactEmailChannelEffectiveSubscriptionStatusValueFromString(s string) (AudiencesContactEmailChannelEffectiveSubscriptionStatusValue, error) {
	switch s {
	case "subscribed":
		return AudiencesContactEmailChannelEffectiveSubscriptionStatusValueSubscribed, nil
	case "unsubscribed":
		return AudiencesContactEmailChannelEffectiveSubscriptionStatusValueUnsubscribed, nil
	case "nonsubscribed":
		return AudiencesContactEmailChannelEffectiveSubscriptionStatusValueNonsubscribed, nil
	case "pending":
		return AudiencesContactEmailChannelEffectiveSubscriptionStatusValuePending, nil
	}
	var t AudiencesContactEmailChannelEffectiveSubscriptionStatusValue
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (a AudiencesContactEmailChannelEffectiveSubscriptionStatusValue) Ptr() *AudiencesContactEmailChannelEffectiveSubscriptionStatusValue {
	return &a
}

// A contact's current consent status for email marketing communications. See the [Audiences (BETA) documentation](https://mailchimp.com/developer/marketing/docs/audiences-introduction) to learn about supported values.
var (
	audiencesContactEmailChannelMarketingConsentFieldSource     = big.NewInt(1 << 0)
	audiencesContactEmailChannelMarketingConsentFieldStatus     = big.NewInt(1 << 1)
	audiencesContactEmailChannelMarketingConsentFieldCapturedAt = big.NewInt(1 << 2)
)

type AudiencesContactEmailChannelMarketingConsent struct {
	// The source from which the parent's entity was created.
	Source *AudiencesContactEmailChannelMarketingConsentSource `json:"source,omitempty" url:"source,omitempty"`
	Status *AudiencesContactEmailChannelMarketingConsentStatus `json:"status,omitempty" url:"status,omitempty"`
	// The ISO 8601 timestamp when the email marketing consent state was recorded; accepted and returned only when status is `confirmed` or `consented`; defaults to the current time if omitted; ignored if older than an existing stored timestamp (staleness guard).
	CapturedAt *time.Time `json:"captured_at,omitempty" url:"captured_at,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *AudiencesContactEmailChannelMarketingConsent) GetSource() *AudiencesContactEmailChannelMarketingConsentSource {
	if a == nil {
		return nil
	}
	return a.Source
}

func (a *AudiencesContactEmailChannelMarketingConsent) GetStatus() *AudiencesContactEmailChannelMarketingConsentStatus {
	if a == nil {
		return nil
	}
	return a.Status
}

func (a *AudiencesContactEmailChannelMarketingConsent) GetCapturedAt() *time.Time {
	if a == nil {
		return nil
	}
	return a.CapturedAt
}

func (a *AudiencesContactEmailChannelMarketingConsent) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a *AudiencesContactEmailChannelMarketingConsent) require(field *big.Int) {
	if a.explicitFields == nil {
		a.explicitFields = big.NewInt(0)
	}
	a.explicitFields.Or(a.explicitFields, field)
}

// SetSource sets the Source field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContactEmailChannelMarketingConsent) SetSource(source *AudiencesContactEmailChannelMarketingConsentSource) {
	a.Source = source
	a.require(audiencesContactEmailChannelMarketingConsentFieldSource)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContactEmailChannelMarketingConsent) SetStatus(status *AudiencesContactEmailChannelMarketingConsentStatus) {
	a.Status = status
	a.require(audiencesContactEmailChannelMarketingConsentFieldStatus)
}

// SetCapturedAt sets the CapturedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContactEmailChannelMarketingConsent) SetCapturedAt(capturedAt *time.Time) {
	a.CapturedAt = capturedAt
	a.require(audiencesContactEmailChannelMarketingConsentFieldCapturedAt)
}

func (a *AudiencesContactEmailChannelMarketingConsent) UnmarshalJSON(data []byte) error {
	type embed AudiencesContactEmailChannelMarketingConsent
	var unmarshaler = struct {
		embed
		CapturedAt *internal.DateTime `json:"captured_at,omitempty"`
	}{
		embed: embed(*a),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*a = AudiencesContactEmailChannelMarketingConsent(unmarshaler.embed)
	a.CapturedAt = unmarshaler.CapturedAt.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *AudiencesContactEmailChannelMarketingConsent) MarshalJSON() ([]byte, error) {
	type embed AudiencesContactEmailChannelMarketingConsent
	var marshaler = struct {
		embed
		CapturedAt *internal.DateTime `json:"captured_at,omitempty"`
	}{
		embed:      embed(*a),
		CapturedAt: internal.NewOptionalDateTime(a.CapturedAt),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, a.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (a *AudiencesContactEmailChannelMarketingConsent) String() string {
	if a == nil {
		return "<nil>"
	}
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

// The source from which the parent's entity was created.
var (
	audiencesContactEmailChannelMarketingConsentSourceFieldName = big.NewInt(1 << 0)
)

type AudiencesContactEmailChannelMarketingConsentSource struct {
	// The name of the entity's source
	Name *string `json:"name,omitempty" url:"name,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *AudiencesContactEmailChannelMarketingConsentSource) GetName() *string {
	if a == nil {
		return nil
	}
	return a.Name
}

func (a *AudiencesContactEmailChannelMarketingConsentSource) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a *AudiencesContactEmailChannelMarketingConsentSource) require(field *big.Int) {
	if a.explicitFields == nil {
		a.explicitFields = big.NewInt(0)
	}
	a.explicitFields.Or(a.explicitFields, field)
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContactEmailChannelMarketingConsentSource) SetName(name *string) {
	a.Name = name
	a.require(audiencesContactEmailChannelMarketingConsentSourceFieldName)
}

func (a *AudiencesContactEmailChannelMarketingConsentSource) UnmarshalJSON(data []byte) error {
	type unmarshaler AudiencesContactEmailChannelMarketingConsentSource
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AudiencesContactEmailChannelMarketingConsentSource(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *AudiencesContactEmailChannelMarketingConsentSource) MarshalJSON() ([]byte, error) {
	type embed AudiencesContactEmailChannelMarketingConsentSource
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, a.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (a *AudiencesContactEmailChannelMarketingConsentSource) String() string {
	if a == nil {
		return "<nil>"
	}
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type AudiencesContactEmailChannelMarketingConsentStatus string

const (
	AudiencesContactEmailChannelMarketingConsentStatusConsented AudiencesContactEmailChannelMarketingConsentStatus = "consented"
	AudiencesContactEmailChannelMarketingConsentStatusDenied    AudiencesContactEmailChannelMarketingConsentStatus = "denied"
	AudiencesContactEmailChannelMarketingConsentStatusConfirmed AudiencesContactEmailChannelMarketingConsentStatus = "confirmed"
	AudiencesContactEmailChannelMarketingConsentStatusUnknown   AudiencesContactEmailChannelMarketingConsentStatus = "unknown"
)

func NewAudiencesContactEmailChannelMarketingConsentStatusFromString(s string) (AudiencesContactEmailChannelMarketingConsentStatus, error) {
	switch s {
	case "consented":
		return AudiencesContactEmailChannelMarketingConsentStatusConsented, nil
	case "denied":
		return AudiencesContactEmailChannelMarketingConsentStatusDenied, nil
	case "confirmed":
		return AudiencesContactEmailChannelMarketingConsentStatusConfirmed, nil
	case "unknown":
		return AudiencesContactEmailChannelMarketingConsentStatusUnknown, nil
	}
	var t AudiencesContactEmailChannelMarketingConsentStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (a AudiencesContactEmailChannelMarketingConsentStatus) Ptr() *AudiencesContactEmailChannelMarketingConsentStatus {
	return &a
}

// The source from which the parent's entity was created.
var (
	audiencesContactEmailChannelSourceFieldName = big.NewInt(1 << 0)
)

type AudiencesContactEmailChannelSource struct {
	// The name of the entity's source
	Name *string `json:"name,omitempty" url:"name,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *AudiencesContactEmailChannelSource) GetName() *string {
	if a == nil {
		return nil
	}
	return a.Name
}

func (a *AudiencesContactEmailChannelSource) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a *AudiencesContactEmailChannelSource) require(field *big.Int) {
	if a.explicitFields == nil {
		a.explicitFields = big.NewInt(0)
	}
	a.explicitFields.Or(a.explicitFields, field)
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContactEmailChannelSource) SetName(name *string) {
	a.Name = name
	a.require(audiencesContactEmailChannelSourceFieldName)
}

func (a *AudiencesContactEmailChannelSource) UnmarshalJSON(data []byte) error {
	type unmarshaler AudiencesContactEmailChannelSource
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AudiencesContactEmailChannelSource(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *AudiencesContactEmailChannelSource) MarshalJSON() ([]byte, error) {
	type embed AudiencesContactEmailChannelSource
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, a.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (a *AudiencesContactEmailChannelSource) String() string {
	if a == nil {
		return "<nil>"
	}
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

// The contact's detected language. Empty string when no language has been detected or set.
type AudiencesContactLanguage string

const (
	AudiencesContactLanguageEmpty AudiencesContactLanguage = ""
	AudiencesContactLanguageEn    AudiencesContactLanguage = "en"
	AudiencesContactLanguageAr    AudiencesContactLanguage = "ar"
	AudiencesContactLanguageAf    AudiencesContactLanguage = "af"
	AudiencesContactLanguageBe    AudiencesContactLanguage = "be"
	AudiencesContactLanguageBg    AudiencesContactLanguage = "bg"
	AudiencesContactLanguageCa    AudiencesContactLanguage = "ca"
	AudiencesContactLanguageZh    AudiencesContactLanguage = "zh"
	AudiencesContactLanguageZhCn  AudiencesContactLanguage = "zh_CN"
	AudiencesContactLanguageHr    AudiencesContactLanguage = "hr"
	AudiencesContactLanguageCs    AudiencesContactLanguage = "cs"
	AudiencesContactLanguageDa    AudiencesContactLanguage = "da"
	AudiencesContactLanguageNl    AudiencesContactLanguage = "nl"
	AudiencesContactLanguageEt    AudiencesContactLanguage = "et"
	AudiencesContactLanguageFa    AudiencesContactLanguage = "fa"
	AudiencesContactLanguageFi    AudiencesContactLanguage = "fi"
	AudiencesContactLanguageFr    AudiencesContactLanguage = "fr"
	AudiencesContactLanguageFrCa  AudiencesContactLanguage = "fr_CA"
	AudiencesContactLanguageDe    AudiencesContactLanguage = "de"
	AudiencesContactLanguageEl    AudiencesContactLanguage = "el"
	AudiencesContactLanguageHe    AudiencesContactLanguage = "he"
	AudiencesContactLanguageHi    AudiencesContactLanguage = "hi"
	AudiencesContactLanguageHu    AudiencesContactLanguage = "hu"
	AudiencesContactLanguageIs    AudiencesContactLanguage = "is"
	AudiencesContactLanguageID    AudiencesContactLanguage = "id"
	AudiencesContactLanguageGa    AudiencesContactLanguage = "ga"
	AudiencesContactLanguageIt    AudiencesContactLanguage = "it"
	AudiencesContactLanguageJa    AudiencesContactLanguage = "ja"
	AudiencesContactLanguageKm    AudiencesContactLanguage = "km"
	AudiencesContactLanguageKo    AudiencesContactLanguage = "ko"
	AudiencesContactLanguageLv    AudiencesContactLanguage = "lv"
	AudiencesContactLanguageLt    AudiencesContactLanguage = "lt"
	AudiencesContactLanguageMt    AudiencesContactLanguage = "mt"
	AudiencesContactLanguageMs    AudiencesContactLanguage = "ms"
	AudiencesContactLanguageMk    AudiencesContactLanguage = "mk"
	AudiencesContactLanguageNo    AudiencesContactLanguage = "no"
	AudiencesContactLanguagePl    AudiencesContactLanguage = "pl"
	AudiencesContactLanguagePt    AudiencesContactLanguage = "pt"
	AudiencesContactLanguagePtPt  AudiencesContactLanguage = "pt_PT"
	AudiencesContactLanguageRo    AudiencesContactLanguage = "ro"
	AudiencesContactLanguageRu    AudiencesContactLanguage = "ru"
	AudiencesContactLanguageSr    AudiencesContactLanguage = "sr"
	AudiencesContactLanguageSk    AudiencesContactLanguage = "sk"
	AudiencesContactLanguageSl    AudiencesContactLanguage = "sl"
	AudiencesContactLanguageEs    AudiencesContactLanguage = "es"
	AudiencesContactLanguageEsEs  AudiencesContactLanguage = "es_ES"
	AudiencesContactLanguageSw    AudiencesContactLanguage = "sw"
	AudiencesContactLanguageSv    AudiencesContactLanguage = "sv"
	AudiencesContactLanguageTa    AudiencesContactLanguage = "ta"
	AudiencesContactLanguageTh    AudiencesContactLanguage = "th"
	AudiencesContactLanguageTr    AudiencesContactLanguage = "tr"
	AudiencesContactLanguageUk    AudiencesContactLanguage = "uk"
	AudiencesContactLanguageVi    AudiencesContactLanguage = "vi"
)

func NewAudiencesContactLanguageFromString(s string) (AudiencesContactLanguage, error) {
	switch s {
	case "":
		return AudiencesContactLanguageEmpty, nil
	case "en":
		return AudiencesContactLanguageEn, nil
	case "ar":
		return AudiencesContactLanguageAr, nil
	case "af":
		return AudiencesContactLanguageAf, nil
	case "be":
		return AudiencesContactLanguageBe, nil
	case "bg":
		return AudiencesContactLanguageBg, nil
	case "ca":
		return AudiencesContactLanguageCa, nil
	case "zh":
		return AudiencesContactLanguageZh, nil
	case "zh_CN":
		return AudiencesContactLanguageZhCn, nil
	case "hr":
		return AudiencesContactLanguageHr, nil
	case "cs":
		return AudiencesContactLanguageCs, nil
	case "da":
		return AudiencesContactLanguageDa, nil
	case "nl":
		return AudiencesContactLanguageNl, nil
	case "et":
		return AudiencesContactLanguageEt, nil
	case "fa":
		return AudiencesContactLanguageFa, nil
	case "fi":
		return AudiencesContactLanguageFi, nil
	case "fr":
		return AudiencesContactLanguageFr, nil
	case "fr_CA":
		return AudiencesContactLanguageFrCa, nil
	case "de":
		return AudiencesContactLanguageDe, nil
	case "el":
		return AudiencesContactLanguageEl, nil
	case "he":
		return AudiencesContactLanguageHe, nil
	case "hi":
		return AudiencesContactLanguageHi, nil
	case "hu":
		return AudiencesContactLanguageHu, nil
	case "is":
		return AudiencesContactLanguageIs, nil
	case "id":
		return AudiencesContactLanguageID, nil
	case "ga":
		return AudiencesContactLanguageGa, nil
	case "it":
		return AudiencesContactLanguageIt, nil
	case "ja":
		return AudiencesContactLanguageJa, nil
	case "km":
		return AudiencesContactLanguageKm, nil
	case "ko":
		return AudiencesContactLanguageKo, nil
	case "lv":
		return AudiencesContactLanguageLv, nil
	case "lt":
		return AudiencesContactLanguageLt, nil
	case "mt":
		return AudiencesContactLanguageMt, nil
	case "ms":
		return AudiencesContactLanguageMs, nil
	case "mk":
		return AudiencesContactLanguageMk, nil
	case "no":
		return AudiencesContactLanguageNo, nil
	case "pl":
		return AudiencesContactLanguagePl, nil
	case "pt":
		return AudiencesContactLanguagePt, nil
	case "pt_PT":
		return AudiencesContactLanguagePtPt, nil
	case "ro":
		return AudiencesContactLanguageRo, nil
	case "ru":
		return AudiencesContactLanguageRu, nil
	case "sr":
		return AudiencesContactLanguageSr, nil
	case "sk":
		return AudiencesContactLanguageSk, nil
	case "sl":
		return AudiencesContactLanguageSl, nil
	case "es":
		return AudiencesContactLanguageEs, nil
	case "es_ES":
		return AudiencesContactLanguageEsEs, nil
	case "sw":
		return AudiencesContactLanguageSw, nil
	case "sv":
		return AudiencesContactLanguageSv, nil
	case "ta":
		return AudiencesContactLanguageTa, nil
	case "th":
		return AudiencesContactLanguageTh, nil
	case "tr":
		return AudiencesContactLanguageTr, nil
	case "uk":
		return AudiencesContactLanguageUk, nil
	case "vi":
		return AudiencesContactLanguageVi, nil
	}
	var t AudiencesContactLanguage
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (a AudiencesContactLanguage) Ptr() *AudiencesContactLanguage {
	return &a
}

// This object's keys are merge tags (like FNAME). It's values are the values to be added to the merge field.
type AudiencesContactMergeFieldsValue struct {
	AudiencesContactMergeFieldsValueAddr1 *AudiencesContactMergeFieldsValueAddr1
	String                                string
	Double                                float64

	typ string
}

func (a *AudiencesContactMergeFieldsValue) GetAudiencesContactMergeFieldsValueAddr1() *AudiencesContactMergeFieldsValueAddr1 {
	if a == nil {
		return nil
	}
	return a.AudiencesContactMergeFieldsValueAddr1
}

func (a *AudiencesContactMergeFieldsValue) GetString() string {
	if a == nil {
		return ""
	}
	return a.String
}

func (a *AudiencesContactMergeFieldsValue) GetDouble() float64 {
	if a == nil {
		return 0
	}
	return a.Double
}

func (a *AudiencesContactMergeFieldsValue) UnmarshalJSON(data []byte) error {
	valueAudiencesContactMergeFieldsValueAddr1 := new(AudiencesContactMergeFieldsValueAddr1)
	if err := json.Unmarshal(data, &valueAudiencesContactMergeFieldsValueAddr1); err == nil {
		a.typ = "AudiencesContactMergeFieldsValueAddr1"
		a.AudiencesContactMergeFieldsValueAddr1 = valueAudiencesContactMergeFieldsValueAddr1
		return nil
	}
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		a.typ = "String"
		a.String = valueString
		return nil
	}
	var valueDouble float64
	if err := json.Unmarshal(data, &valueDouble); err == nil {
		a.typ = "Double"
		a.Double = valueDouble
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
}

func (a AudiencesContactMergeFieldsValue) MarshalJSON() ([]byte, error) {
	if a.typ == "AudiencesContactMergeFieldsValueAddr1" || a.AudiencesContactMergeFieldsValueAddr1 != nil {
		return json.Marshal(a.AudiencesContactMergeFieldsValueAddr1)
	}
	if a.typ == "String" || a.String != "" {
		return json.Marshal(a.String)
	}
	if a.typ == "Double" || a.Double != 0 {
		return json.Marshal(a.Double)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", a)
}

type AudiencesContactMergeFieldsValueVisitor interface {
	VisitAudiencesContactMergeFieldsValueAddr1(*AudiencesContactMergeFieldsValueAddr1) error
	VisitString(string) error
	VisitDouble(float64) error
}

func (a *AudiencesContactMergeFieldsValue) Accept(visitor AudiencesContactMergeFieldsValueVisitor) error {
	if a.typ == "AudiencesContactMergeFieldsValueAddr1" || a.AudiencesContactMergeFieldsValueAddr1 != nil {
		return visitor.VisitAudiencesContactMergeFieldsValueAddr1(a.AudiencesContactMergeFieldsValueAddr1)
	}
	if a.typ == "String" || a.String != "" {
		return visitor.VisitString(a.String)
	}
	if a.typ == "Double" || a.Double != 0 {
		return visitor.VisitDouble(a.Double)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", a)
}

var (
	audiencesContactMergeFieldsValueAddr1FieldAddr1   = big.NewInt(1 << 0)
	audiencesContactMergeFieldsValueAddr1FieldAddr2   = big.NewInt(1 << 1)
	audiencesContactMergeFieldsValueAddr1FieldCity    = big.NewInt(1 << 2)
	audiencesContactMergeFieldsValueAddr1FieldState   = big.NewInt(1 << 3)
	audiencesContactMergeFieldsValueAddr1FieldZip     = big.NewInt(1 << 4)
	audiencesContactMergeFieldsValueAddr1FieldCountry = big.NewInt(1 << 5)
)

type AudiencesContactMergeFieldsValueAddr1 struct {
	Addr1   string  `json:"addr1" url:"addr1"`
	Addr2   *string `json:"addr2,omitempty" url:"addr2,omitempty"`
	City    string  `json:"city" url:"city"`
	State   string  `json:"state" url:"state"`
	Zip     string  `json:"zip" url:"zip"`
	Country *string `json:"country,omitempty" url:"country,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *AudiencesContactMergeFieldsValueAddr1) GetAddr1() string {
	if a == nil {
		return ""
	}
	return a.Addr1
}

func (a *AudiencesContactMergeFieldsValueAddr1) GetAddr2() *string {
	if a == nil {
		return nil
	}
	return a.Addr2
}

func (a *AudiencesContactMergeFieldsValueAddr1) GetCity() string {
	if a == nil {
		return ""
	}
	return a.City
}

func (a *AudiencesContactMergeFieldsValueAddr1) GetState() string {
	if a == nil {
		return ""
	}
	return a.State
}

func (a *AudiencesContactMergeFieldsValueAddr1) GetZip() string {
	if a == nil {
		return ""
	}
	return a.Zip
}

func (a *AudiencesContactMergeFieldsValueAddr1) GetCountry() *string {
	if a == nil {
		return nil
	}
	return a.Country
}

func (a *AudiencesContactMergeFieldsValueAddr1) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a *AudiencesContactMergeFieldsValueAddr1) require(field *big.Int) {
	if a.explicitFields == nil {
		a.explicitFields = big.NewInt(0)
	}
	a.explicitFields.Or(a.explicitFields, field)
}

// SetAddr1 sets the Addr1 field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContactMergeFieldsValueAddr1) SetAddr1(addr1 string) {
	a.Addr1 = addr1
	a.require(audiencesContactMergeFieldsValueAddr1FieldAddr1)
}

// SetAddr2 sets the Addr2 field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContactMergeFieldsValueAddr1) SetAddr2(addr2 *string) {
	a.Addr2 = addr2
	a.require(audiencesContactMergeFieldsValueAddr1FieldAddr2)
}

// SetCity sets the City field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContactMergeFieldsValueAddr1) SetCity(city string) {
	a.City = city
	a.require(audiencesContactMergeFieldsValueAddr1FieldCity)
}

// SetState sets the State field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContactMergeFieldsValueAddr1) SetState(state string) {
	a.State = state
	a.require(audiencesContactMergeFieldsValueAddr1FieldState)
}

// SetZip sets the Zip field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContactMergeFieldsValueAddr1) SetZip(zip string) {
	a.Zip = zip
	a.require(audiencesContactMergeFieldsValueAddr1FieldZip)
}

// SetCountry sets the Country field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContactMergeFieldsValueAddr1) SetCountry(country *string) {
	a.Country = country
	a.require(audiencesContactMergeFieldsValueAddr1FieldCountry)
}

func (a *AudiencesContactMergeFieldsValueAddr1) UnmarshalJSON(data []byte) error {
	type unmarshaler AudiencesContactMergeFieldsValueAddr1
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AudiencesContactMergeFieldsValueAddr1(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *AudiencesContactMergeFieldsValueAddr1) MarshalJSON() ([]byte, error) {
	type embed AudiencesContactMergeFieldsValueAddr1
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, a.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (a *AudiencesContactMergeFieldsValueAddr1) String() string {
	if a == nil {
		return "<nil>"
	}
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

var (
	audiencesContactSmsChannelFieldEffectiveSubscriptionStatus = big.NewInt(1 << 0)
	audiencesContactSmsChannelFieldMarketingConsent            = big.NewInt(1 << 1)
	audiencesContactSmsChannelFieldSmsPhone                    = big.NewInt(1 << 2)
	audiencesContactSmsChannelFieldSource                      = big.NewInt(1 << 3)
	audiencesContactSmsChannelFieldHashedSmsPhone              = big.NewInt(1 << 4)
)

type AudiencesContactSmsChannel struct {
	// A computation performed by the Mailchimp platform, triggered whenever any of its inputs change. Some inputs are controlled by API users, while others are tracked internally by the platform. Computation is based on: audience opt-in configuration (single vs. double opt-in), marketing consent status, and deliverability status (an internal state for a contact, maintained by Mailchimp for a specific marketing channel instance). This new API field is distinct from how contacts are displayed in the UI. See the [Audiences (BETA) documentation](https://mailchimp.com/developer/marketing/docs/audiences-introduction) to learn about supported values.
	EffectiveSubscriptionStatus *AudiencesContactSmsChannelEffectiveSubscriptionStatus `json:"effective_subscription_status,omitempty" url:"effective_subscription_status,omitempty"`
	// A contact's current consent status for SMS marketing communications. See the [Audiences (BETA) documentation](https://mailchimp.com/developer/marketing/docs/audiences-introduction) to learn about supported values.
	MarketingConsent *AudiencesContactSmsChannelMarketingConsent `json:"marketing_consent,omitempty" url:"marketing_consent,omitempty"`
	// SMS Phone Number
	SmsPhone *string `json:"sms_phone,omitempty" url:"sms_phone,omitempty"`
	// The source from which the parent's entity was created.
	Source *AudiencesContactSmsChannelSource `json:"source,omitempty" url:"source,omitempty"`
	// SHA256 hash of the SMS phone number
	HashedSmsPhone *string `json:"hashed_sms_phone,omitempty" url:"hashed_sms_phone,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *AudiencesContactSmsChannel) GetEffectiveSubscriptionStatus() *AudiencesContactSmsChannelEffectiveSubscriptionStatus {
	if a == nil {
		return nil
	}
	return a.EffectiveSubscriptionStatus
}

func (a *AudiencesContactSmsChannel) GetMarketingConsent() *AudiencesContactSmsChannelMarketingConsent {
	if a == nil {
		return nil
	}
	return a.MarketingConsent
}

func (a *AudiencesContactSmsChannel) GetSmsPhone() *string {
	if a == nil {
		return nil
	}
	return a.SmsPhone
}

func (a *AudiencesContactSmsChannel) GetSource() *AudiencesContactSmsChannelSource {
	if a == nil {
		return nil
	}
	return a.Source
}

func (a *AudiencesContactSmsChannel) GetHashedSmsPhone() *string {
	if a == nil {
		return nil
	}
	return a.HashedSmsPhone
}

func (a *AudiencesContactSmsChannel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a *AudiencesContactSmsChannel) require(field *big.Int) {
	if a.explicitFields == nil {
		a.explicitFields = big.NewInt(0)
	}
	a.explicitFields.Or(a.explicitFields, field)
}

// SetEffectiveSubscriptionStatus sets the EffectiveSubscriptionStatus field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContactSmsChannel) SetEffectiveSubscriptionStatus(effectiveSubscriptionStatus *AudiencesContactSmsChannelEffectiveSubscriptionStatus) {
	a.EffectiveSubscriptionStatus = effectiveSubscriptionStatus
	a.require(audiencesContactSmsChannelFieldEffectiveSubscriptionStatus)
}

// SetMarketingConsent sets the MarketingConsent field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContactSmsChannel) SetMarketingConsent(marketingConsent *AudiencesContactSmsChannelMarketingConsent) {
	a.MarketingConsent = marketingConsent
	a.require(audiencesContactSmsChannelFieldMarketingConsent)
}

// SetSmsPhone sets the SmsPhone field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContactSmsChannel) SetSmsPhone(smsPhone *string) {
	a.SmsPhone = smsPhone
	a.require(audiencesContactSmsChannelFieldSmsPhone)
}

// SetSource sets the Source field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContactSmsChannel) SetSource(source *AudiencesContactSmsChannelSource) {
	a.Source = source
	a.require(audiencesContactSmsChannelFieldSource)
}

// SetHashedSmsPhone sets the HashedSmsPhone field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContactSmsChannel) SetHashedSmsPhone(hashedSmsPhone *string) {
	a.HashedSmsPhone = hashedSmsPhone
	a.require(audiencesContactSmsChannelFieldHashedSmsPhone)
}

func (a *AudiencesContactSmsChannel) UnmarshalJSON(data []byte) error {
	type unmarshaler AudiencesContactSmsChannel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AudiencesContactSmsChannel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *AudiencesContactSmsChannel) MarshalJSON() ([]byte, error) {
	type embed AudiencesContactSmsChannel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, a.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (a *AudiencesContactSmsChannel) String() string {
	if a == nil {
		return "<nil>"
	}
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

// A computation performed by the Mailchimp platform, triggered whenever any of its inputs change. Some inputs are controlled by API users, while others are tracked internally by the platform. Computation is based on: audience opt-in configuration (single vs. double opt-in), marketing consent status, and deliverability status (an internal state for a contact, maintained by Mailchimp for a specific marketing channel instance). This new API field is distinct from how contacts are displayed in the UI. See the [Audiences (BETA) documentation](https://mailchimp.com/developer/marketing/docs/audiences-introduction) to learn about supported values.
var (
	audiencesContactSmsChannelEffectiveSubscriptionStatusFieldValue = big.NewInt(1 << 0)
)

type AudiencesContactSmsChannelEffectiveSubscriptionStatus struct {
	Value *AudiencesContactSmsChannelEffectiveSubscriptionStatusValue `json:"value,omitempty" url:"value,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *AudiencesContactSmsChannelEffectiveSubscriptionStatus) GetValue() *AudiencesContactSmsChannelEffectiveSubscriptionStatusValue {
	if a == nil {
		return nil
	}
	return a.Value
}

func (a *AudiencesContactSmsChannelEffectiveSubscriptionStatus) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a *AudiencesContactSmsChannelEffectiveSubscriptionStatus) require(field *big.Int) {
	if a.explicitFields == nil {
		a.explicitFields = big.NewInt(0)
	}
	a.explicitFields.Or(a.explicitFields, field)
}

// SetValue sets the Value field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContactSmsChannelEffectiveSubscriptionStatus) SetValue(value *AudiencesContactSmsChannelEffectiveSubscriptionStatusValue) {
	a.Value = value
	a.require(audiencesContactSmsChannelEffectiveSubscriptionStatusFieldValue)
}

func (a *AudiencesContactSmsChannelEffectiveSubscriptionStatus) UnmarshalJSON(data []byte) error {
	type unmarshaler AudiencesContactSmsChannelEffectiveSubscriptionStatus
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AudiencesContactSmsChannelEffectiveSubscriptionStatus(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *AudiencesContactSmsChannelEffectiveSubscriptionStatus) MarshalJSON() ([]byte, error) {
	type embed AudiencesContactSmsChannelEffectiveSubscriptionStatus
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, a.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (a *AudiencesContactSmsChannelEffectiveSubscriptionStatus) String() string {
	if a == nil {
		return "<nil>"
	}
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type AudiencesContactSmsChannelEffectiveSubscriptionStatusValue string

const (
	AudiencesContactSmsChannelEffectiveSubscriptionStatusValueSubscribed    AudiencesContactSmsChannelEffectiveSubscriptionStatusValue = "subscribed"
	AudiencesContactSmsChannelEffectiveSubscriptionStatusValueUnsubscribed  AudiencesContactSmsChannelEffectiveSubscriptionStatusValue = "unsubscribed"
	AudiencesContactSmsChannelEffectiveSubscriptionStatusValueNonsubscribed AudiencesContactSmsChannelEffectiveSubscriptionStatusValue = "nonsubscribed"
	AudiencesContactSmsChannelEffectiveSubscriptionStatusValuePending       AudiencesContactSmsChannelEffectiveSubscriptionStatusValue = "pending"
)

func NewAudiencesContactSmsChannelEffectiveSubscriptionStatusValueFromString(s string) (AudiencesContactSmsChannelEffectiveSubscriptionStatusValue, error) {
	switch s {
	case "subscribed":
		return AudiencesContactSmsChannelEffectiveSubscriptionStatusValueSubscribed, nil
	case "unsubscribed":
		return AudiencesContactSmsChannelEffectiveSubscriptionStatusValueUnsubscribed, nil
	case "nonsubscribed":
		return AudiencesContactSmsChannelEffectiveSubscriptionStatusValueNonsubscribed, nil
	case "pending":
		return AudiencesContactSmsChannelEffectiveSubscriptionStatusValuePending, nil
	}
	var t AudiencesContactSmsChannelEffectiveSubscriptionStatusValue
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (a AudiencesContactSmsChannelEffectiveSubscriptionStatusValue) Ptr() *AudiencesContactSmsChannelEffectiveSubscriptionStatusValue {
	return &a
}

// A contact's current consent status for SMS marketing communications. See the [Audiences (BETA) documentation](https://mailchimp.com/developer/marketing/docs/audiences-introduction) to learn about supported values.
var (
	audiencesContactSmsChannelMarketingConsentFieldSource     = big.NewInt(1 << 0)
	audiencesContactSmsChannelMarketingConsentFieldStatus     = big.NewInt(1 << 1)
	audiencesContactSmsChannelMarketingConsentFieldCapturedAt = big.NewInt(1 << 2)
)

type AudiencesContactSmsChannelMarketingConsent struct {
	// The source from which the parent's entity was created.
	Source *AudiencesContactSmsChannelMarketingConsentSource `json:"source,omitempty" url:"source,omitempty"`
	// The contact's SMS marketing consent status. Use `confirmed` for double opt-in audiences, `consented` for single opt-in audiences. `denied` is accepted on PATCH/PUT only (not POST) and drives an API-initiated unsubscribe; it cannot be used when creating a new contact.
	Status *AudiencesContactSmsChannelMarketingConsentStatus `json:"status,omitempty" url:"status,omitempty"`
	// The timestamp when SMS marketing consent was captured (ISO 8601). Only accepted and returned when status is `confirmed`. The timestamp of the consent state change being recorded. Defaults to the current time if not provided. If the contact already has a consent timestamp on record that is equal to or newer than the supplied value, the supplied value is ignored (staleness guard); to update the consent timestamp supply a value strictly newer than the stored one.
	CapturedAt *time.Time `json:"captured_at,omitempty" url:"captured_at,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *AudiencesContactSmsChannelMarketingConsent) GetSource() *AudiencesContactSmsChannelMarketingConsentSource {
	if a == nil {
		return nil
	}
	return a.Source
}

func (a *AudiencesContactSmsChannelMarketingConsent) GetStatus() *AudiencesContactSmsChannelMarketingConsentStatus {
	if a == nil {
		return nil
	}
	return a.Status
}

func (a *AudiencesContactSmsChannelMarketingConsent) GetCapturedAt() *time.Time {
	if a == nil {
		return nil
	}
	return a.CapturedAt
}

func (a *AudiencesContactSmsChannelMarketingConsent) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a *AudiencesContactSmsChannelMarketingConsent) require(field *big.Int) {
	if a.explicitFields == nil {
		a.explicitFields = big.NewInt(0)
	}
	a.explicitFields.Or(a.explicitFields, field)
}

// SetSource sets the Source field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContactSmsChannelMarketingConsent) SetSource(source *AudiencesContactSmsChannelMarketingConsentSource) {
	a.Source = source
	a.require(audiencesContactSmsChannelMarketingConsentFieldSource)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContactSmsChannelMarketingConsent) SetStatus(status *AudiencesContactSmsChannelMarketingConsentStatus) {
	a.Status = status
	a.require(audiencesContactSmsChannelMarketingConsentFieldStatus)
}

// SetCapturedAt sets the CapturedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContactSmsChannelMarketingConsent) SetCapturedAt(capturedAt *time.Time) {
	a.CapturedAt = capturedAt
	a.require(audiencesContactSmsChannelMarketingConsentFieldCapturedAt)
}

func (a *AudiencesContactSmsChannelMarketingConsent) UnmarshalJSON(data []byte) error {
	type embed AudiencesContactSmsChannelMarketingConsent
	var unmarshaler = struct {
		embed
		CapturedAt *internal.DateTime `json:"captured_at,omitempty"`
	}{
		embed: embed(*a),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*a = AudiencesContactSmsChannelMarketingConsent(unmarshaler.embed)
	a.CapturedAt = unmarshaler.CapturedAt.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *AudiencesContactSmsChannelMarketingConsent) MarshalJSON() ([]byte, error) {
	type embed AudiencesContactSmsChannelMarketingConsent
	var marshaler = struct {
		embed
		CapturedAt *internal.DateTime `json:"captured_at,omitempty"`
	}{
		embed:      embed(*a),
		CapturedAt: internal.NewOptionalDateTime(a.CapturedAt),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, a.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (a *AudiencesContactSmsChannelMarketingConsent) String() string {
	if a == nil {
		return "<nil>"
	}
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

// The source from which the parent's entity was created.
var (
	audiencesContactSmsChannelMarketingConsentSourceFieldName = big.NewInt(1 << 0)
)

type AudiencesContactSmsChannelMarketingConsentSource struct {
	// The name of the entity's source
	Name *string `json:"name,omitempty" url:"name,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *AudiencesContactSmsChannelMarketingConsentSource) GetName() *string {
	if a == nil {
		return nil
	}
	return a.Name
}

func (a *AudiencesContactSmsChannelMarketingConsentSource) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a *AudiencesContactSmsChannelMarketingConsentSource) require(field *big.Int) {
	if a.explicitFields == nil {
		a.explicitFields = big.NewInt(0)
	}
	a.explicitFields.Or(a.explicitFields, field)
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContactSmsChannelMarketingConsentSource) SetName(name *string) {
	a.Name = name
	a.require(audiencesContactSmsChannelMarketingConsentSourceFieldName)
}

func (a *AudiencesContactSmsChannelMarketingConsentSource) UnmarshalJSON(data []byte) error {
	type unmarshaler AudiencesContactSmsChannelMarketingConsentSource
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AudiencesContactSmsChannelMarketingConsentSource(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *AudiencesContactSmsChannelMarketingConsentSource) MarshalJSON() ([]byte, error) {
	type embed AudiencesContactSmsChannelMarketingConsentSource
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, a.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (a *AudiencesContactSmsChannelMarketingConsentSource) String() string {
	if a == nil {
		return "<nil>"
	}
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

// The contact's SMS marketing consent status. Use `confirmed` for double opt-in audiences, `consented` for single opt-in audiences. `denied` is accepted on PATCH/PUT only (not POST) and drives an API-initiated unsubscribe; it cannot be used when creating a new contact.
type AudiencesContactSmsChannelMarketingConsentStatus string

const (
	AudiencesContactSmsChannelMarketingConsentStatusConsented AudiencesContactSmsChannelMarketingConsentStatus = "consented"
	AudiencesContactSmsChannelMarketingConsentStatusConfirmed AudiencesContactSmsChannelMarketingConsentStatus = "confirmed"
	AudiencesContactSmsChannelMarketingConsentStatusDenied    AudiencesContactSmsChannelMarketingConsentStatus = "denied"
	AudiencesContactSmsChannelMarketingConsentStatusUnknown   AudiencesContactSmsChannelMarketingConsentStatus = "unknown"
)

func NewAudiencesContactSmsChannelMarketingConsentStatusFromString(s string) (AudiencesContactSmsChannelMarketingConsentStatus, error) {
	switch s {
	case "consented":
		return AudiencesContactSmsChannelMarketingConsentStatusConsented, nil
	case "confirmed":
		return AudiencesContactSmsChannelMarketingConsentStatusConfirmed, nil
	case "denied":
		return AudiencesContactSmsChannelMarketingConsentStatusDenied, nil
	case "unknown":
		return AudiencesContactSmsChannelMarketingConsentStatusUnknown, nil
	}
	var t AudiencesContactSmsChannelMarketingConsentStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (a AudiencesContactSmsChannelMarketingConsentStatus) Ptr() *AudiencesContactSmsChannelMarketingConsentStatus {
	return &a
}

// The source from which the parent's entity was created.
var (
	audiencesContactSmsChannelSourceFieldName = big.NewInt(1 << 0)
)

type AudiencesContactSmsChannelSource struct {
	// The name of the entity's source
	Name *string `json:"name,omitempty" url:"name,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *AudiencesContactSmsChannelSource) GetName() *string {
	if a == nil {
		return nil
	}
	return a.Name
}

func (a *AudiencesContactSmsChannelSource) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a *AudiencesContactSmsChannelSource) require(field *big.Int) {
	if a.explicitFields == nil {
		a.explicitFields = big.NewInt(0)
	}
	a.explicitFields.Or(a.explicitFields, field)
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContactSmsChannelSource) SetName(name *string) {
	a.Name = name
	a.require(audiencesContactSmsChannelSourceFieldName)
}

func (a *AudiencesContactSmsChannelSource) UnmarshalJSON(data []byte) error {
	type unmarshaler AudiencesContactSmsChannelSource
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AudiencesContactSmsChannelSource(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *AudiencesContactSmsChannelSource) MarshalJSON() ([]byte, error) {
	type embed AudiencesContactSmsChannelSource
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, a.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (a *AudiencesContactSmsChannelSource) String() string {
	if a == nil {
		return "<nil>"
	}
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

// The source from which the parent's entity was created.
var (
	audiencesContactSourceFieldName = big.NewInt(1 << 0)
)

type AudiencesContactSource struct {
	// The name of the entity's source
	Name *string `json:"name,omitempty" url:"name,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *AudiencesContactSource) GetName() *string {
	if a == nil {
		return nil
	}
	return a.Name
}

func (a *AudiencesContactSource) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a *AudiencesContactSource) require(field *big.Int) {
	if a.explicitFields == nil {
		a.explicitFields = big.NewInt(0)
	}
	a.explicitFields.Or(a.explicitFields, field)
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (a *AudiencesContactSource) SetName(name *string) {
	a.Name = name
	a.require(audiencesContactSourceFieldName)
}

func (a *AudiencesContactSource) UnmarshalJSON(data []byte) error {
	type unmarshaler AudiencesContactSource
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AudiencesContactSource(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *AudiencesContactSource) MarshalJSON() ([]byte, error) {
	type embed AudiencesContactSource
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, a.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (a *AudiencesContactSource) String() string {
	if a == nil {
		return "<nil>"
	}
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

// The status of a contact.
type AudiencesContactStatus string

const (
	AudiencesContactStatusActive   AudiencesContactStatus = "active"
	AudiencesContactStatusArchived AudiencesContactStatus = "archived"
)

func NewAudiencesContactStatusFromString(s string) (AudiencesContactStatus, error) {
	switch s {
	case "active":
		return AudiencesContactStatusActive, nil
	case "archived":
		return AudiencesContactStatusArchived, nil
	}
	var t AudiencesContactStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (a AudiencesContactStatus) Ptr() *AudiencesContactStatus {
	return &a
}

type CreateAudienceContactRequestDataMode string

const (
	CreateAudienceContactRequestDataModeHistorical CreateAudienceContactRequestDataMode = "historical"
	CreateAudienceContactRequestDataModeLive       CreateAudienceContactRequestDataMode = "live"
)

func NewCreateAudienceContactRequestDataModeFromString(s string) (CreateAudienceContactRequestDataMode, error) {
	switch s {
	case "historical":
		return CreateAudienceContactRequestDataModeHistorical, nil
	case "live":
		return CreateAudienceContactRequestDataModeLive, nil
	}
	var t CreateAudienceContactRequestDataMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c CreateAudienceContactRequestDataMode) Ptr() *CreateAudienceContactRequestDataMode {
	return &c
}

var (
	createAudienceContactRequestEmailChannelFieldEmail            = big.NewInt(1 << 0)
	createAudienceContactRequestEmailChannelFieldMarketingConsent = big.NewInt(1 << 1)
)

type CreateAudienceContactRequestEmailChannel struct {
	// Email address
	Email *string `json:"email,omitempty" url:"email,omitempty"`
	// A contact's current consent status for email marketing communications. See the [Audiences (BETA) documentation](https://mailchimp.com/developer/marketing/docs/audiences-introduction) to learn about supported values.
	MarketingConsent *CreateAudienceContactRequestEmailChannelMarketingConsent `json:"marketing_consent,omitempty" url:"marketing_consent,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateAudienceContactRequestEmailChannel) GetEmail() *string {
	if c == nil {
		return nil
	}
	return c.Email
}

func (c *CreateAudienceContactRequestEmailChannel) GetMarketingConsent() *CreateAudienceContactRequestEmailChannelMarketingConsent {
	if c == nil {
		return nil
	}
	return c.MarketingConsent
}

func (c *CreateAudienceContactRequestEmailChannel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *CreateAudienceContactRequestEmailChannel) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetEmail sets the Email field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAudienceContactRequestEmailChannel) SetEmail(email *string) {
	c.Email = email
	c.require(createAudienceContactRequestEmailChannelFieldEmail)
}

// SetMarketingConsent sets the MarketingConsent field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAudienceContactRequestEmailChannel) SetMarketingConsent(marketingConsent *CreateAudienceContactRequestEmailChannelMarketingConsent) {
	c.MarketingConsent = marketingConsent
	c.require(createAudienceContactRequestEmailChannelFieldMarketingConsent)
}

func (c *CreateAudienceContactRequestEmailChannel) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateAudienceContactRequestEmailChannel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateAudienceContactRequestEmailChannel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateAudienceContactRequestEmailChannel) MarshalJSON() ([]byte, error) {
	type embed CreateAudienceContactRequestEmailChannel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *CreateAudienceContactRequestEmailChannel) String() string {
	if c == nil {
		return "<nil>"
	}
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

// A contact's current consent status for email marketing communications. See the [Audiences (BETA) documentation](https://mailchimp.com/developer/marketing/docs/audiences-introduction) to learn about supported values.
var (
	createAudienceContactRequestEmailChannelMarketingConsentFieldStatus = big.NewInt(1 << 0)
)

type CreateAudienceContactRequestEmailChannelMarketingConsent struct {
	// Status of a contacts Marketing Consent
	Status *CreateAudienceContactRequestEmailChannelMarketingConsentStatus `json:"status,omitempty" url:"status,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateAudienceContactRequestEmailChannelMarketingConsent) GetStatus() *CreateAudienceContactRequestEmailChannelMarketingConsentStatus {
	if c == nil {
		return nil
	}
	return c.Status
}

func (c *CreateAudienceContactRequestEmailChannelMarketingConsent) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *CreateAudienceContactRequestEmailChannelMarketingConsent) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAudienceContactRequestEmailChannelMarketingConsent) SetStatus(status *CreateAudienceContactRequestEmailChannelMarketingConsentStatus) {
	c.Status = status
	c.require(createAudienceContactRequestEmailChannelMarketingConsentFieldStatus)
}

func (c *CreateAudienceContactRequestEmailChannelMarketingConsent) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateAudienceContactRequestEmailChannelMarketingConsent
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateAudienceContactRequestEmailChannelMarketingConsent(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateAudienceContactRequestEmailChannelMarketingConsent) MarshalJSON() ([]byte, error) {
	type embed CreateAudienceContactRequestEmailChannelMarketingConsent
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *CreateAudienceContactRequestEmailChannelMarketingConsent) String() string {
	if c == nil {
		return "<nil>"
	}
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

// Status of a contacts Marketing Consent
type CreateAudienceContactRequestEmailChannelMarketingConsentStatus string

const (
	CreateAudienceContactRequestEmailChannelMarketingConsentStatusConfirmed CreateAudienceContactRequestEmailChannelMarketingConsentStatus = "confirmed"
	CreateAudienceContactRequestEmailChannelMarketingConsentStatusConsented CreateAudienceContactRequestEmailChannelMarketingConsentStatus = "consented"
	CreateAudienceContactRequestEmailChannelMarketingConsentStatusDenied    CreateAudienceContactRequestEmailChannelMarketingConsentStatus = "denied"
	CreateAudienceContactRequestEmailChannelMarketingConsentStatusUnknown   CreateAudienceContactRequestEmailChannelMarketingConsentStatus = "unknown"
)

func NewCreateAudienceContactRequestEmailChannelMarketingConsentStatusFromString(s string) (CreateAudienceContactRequestEmailChannelMarketingConsentStatus, error) {
	switch s {
	case "confirmed":
		return CreateAudienceContactRequestEmailChannelMarketingConsentStatusConfirmed, nil
	case "consented":
		return CreateAudienceContactRequestEmailChannelMarketingConsentStatusConsented, nil
	case "denied":
		return CreateAudienceContactRequestEmailChannelMarketingConsentStatusDenied, nil
	case "unknown":
		return CreateAudienceContactRequestEmailChannelMarketingConsentStatusUnknown, nil
	}
	var t CreateAudienceContactRequestEmailChannelMarketingConsentStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c CreateAudienceContactRequestEmailChannelMarketingConsentStatus) Ptr() *CreateAudienceContactRequestEmailChannelMarketingConsentStatus {
	return &c
}

type CreateAudienceContactRequestMergeFieldValidationMode string

const (
	CreateAudienceContactRequestMergeFieldValidationModeIgnoreRequiredChecks CreateAudienceContactRequestMergeFieldValidationMode = "ignore_required_checks"
	CreateAudienceContactRequestMergeFieldValidationModeStrict               CreateAudienceContactRequestMergeFieldValidationMode = "strict"
)

func NewCreateAudienceContactRequestMergeFieldValidationModeFromString(s string) (CreateAudienceContactRequestMergeFieldValidationMode, error) {
	switch s {
	case "ignore_required_checks":
		return CreateAudienceContactRequestMergeFieldValidationModeIgnoreRequiredChecks, nil
	case "strict":
		return CreateAudienceContactRequestMergeFieldValidationModeStrict, nil
	}
	var t CreateAudienceContactRequestMergeFieldValidationMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c CreateAudienceContactRequestMergeFieldValidationMode) Ptr() *CreateAudienceContactRequestMergeFieldValidationMode {
	return &c
}

// This object's keys are merge tags (like FNAME). It's values are the values to be added to the merge field.
type CreateAudienceContactRequestMergeFieldsValue struct {
	CreateAudienceContactRequestMergeFieldsValueAddr1 *CreateAudienceContactRequestMergeFieldsValueAddr1
	String                                            string
	Double                                            float64

	typ string
}

func (c *CreateAudienceContactRequestMergeFieldsValue) GetCreateAudienceContactRequestMergeFieldsValueAddr1() *CreateAudienceContactRequestMergeFieldsValueAddr1 {
	if c == nil {
		return nil
	}
	return c.CreateAudienceContactRequestMergeFieldsValueAddr1
}

func (c *CreateAudienceContactRequestMergeFieldsValue) GetString() string {
	if c == nil {
		return ""
	}
	return c.String
}

func (c *CreateAudienceContactRequestMergeFieldsValue) GetDouble() float64 {
	if c == nil {
		return 0
	}
	return c.Double
}

func (c *CreateAudienceContactRequestMergeFieldsValue) UnmarshalJSON(data []byte) error {
	valueCreateAudienceContactRequestMergeFieldsValueAddr1 := new(CreateAudienceContactRequestMergeFieldsValueAddr1)
	if err := json.Unmarshal(data, &valueCreateAudienceContactRequestMergeFieldsValueAddr1); err == nil {
		c.typ = "CreateAudienceContactRequestMergeFieldsValueAddr1"
		c.CreateAudienceContactRequestMergeFieldsValueAddr1 = valueCreateAudienceContactRequestMergeFieldsValueAddr1
		return nil
	}
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		c.typ = "String"
		c.String = valueString
		return nil
	}
	var valueDouble float64
	if err := json.Unmarshal(data, &valueDouble); err == nil {
		c.typ = "Double"
		c.Double = valueDouble
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, c)
}

func (c CreateAudienceContactRequestMergeFieldsValue) MarshalJSON() ([]byte, error) {
	if c.typ == "CreateAudienceContactRequestMergeFieldsValueAddr1" || c.CreateAudienceContactRequestMergeFieldsValueAddr1 != nil {
		return json.Marshal(c.CreateAudienceContactRequestMergeFieldsValueAddr1)
	}
	if c.typ == "String" || c.String != "" {
		return json.Marshal(c.String)
	}
	if c.typ == "Double" || c.Double != 0 {
		return json.Marshal(c.Double)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", c)
}

type CreateAudienceContactRequestMergeFieldsValueVisitor interface {
	VisitCreateAudienceContactRequestMergeFieldsValueAddr1(*CreateAudienceContactRequestMergeFieldsValueAddr1) error
	VisitString(string) error
	VisitDouble(float64) error
}

func (c *CreateAudienceContactRequestMergeFieldsValue) Accept(visitor CreateAudienceContactRequestMergeFieldsValueVisitor) error {
	if c.typ == "CreateAudienceContactRequestMergeFieldsValueAddr1" || c.CreateAudienceContactRequestMergeFieldsValueAddr1 != nil {
		return visitor.VisitCreateAudienceContactRequestMergeFieldsValueAddr1(c.CreateAudienceContactRequestMergeFieldsValueAddr1)
	}
	if c.typ == "String" || c.String != "" {
		return visitor.VisitString(c.String)
	}
	if c.typ == "Double" || c.Double != 0 {
		return visitor.VisitDouble(c.Double)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", c)
}

var (
	createAudienceContactRequestMergeFieldsValueAddr1FieldAddr1   = big.NewInt(1 << 0)
	createAudienceContactRequestMergeFieldsValueAddr1FieldAddr2   = big.NewInt(1 << 1)
	createAudienceContactRequestMergeFieldsValueAddr1FieldCity    = big.NewInt(1 << 2)
	createAudienceContactRequestMergeFieldsValueAddr1FieldState   = big.NewInt(1 << 3)
	createAudienceContactRequestMergeFieldsValueAddr1FieldZip     = big.NewInt(1 << 4)
	createAudienceContactRequestMergeFieldsValueAddr1FieldCountry = big.NewInt(1 << 5)
)

type CreateAudienceContactRequestMergeFieldsValueAddr1 struct {
	Addr1   string  `json:"addr1" url:"addr1"`
	Addr2   *string `json:"addr2,omitempty" url:"addr2,omitempty"`
	City    string  `json:"city" url:"city"`
	State   string  `json:"state" url:"state"`
	Zip     string  `json:"zip" url:"zip"`
	Country *string `json:"country,omitempty" url:"country,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateAudienceContactRequestMergeFieldsValueAddr1) GetAddr1() string {
	if c == nil {
		return ""
	}
	return c.Addr1
}

func (c *CreateAudienceContactRequestMergeFieldsValueAddr1) GetAddr2() *string {
	if c == nil {
		return nil
	}
	return c.Addr2
}

func (c *CreateAudienceContactRequestMergeFieldsValueAddr1) GetCity() string {
	if c == nil {
		return ""
	}
	return c.City
}

func (c *CreateAudienceContactRequestMergeFieldsValueAddr1) GetState() string {
	if c == nil {
		return ""
	}
	return c.State
}

func (c *CreateAudienceContactRequestMergeFieldsValueAddr1) GetZip() string {
	if c == nil {
		return ""
	}
	return c.Zip
}

func (c *CreateAudienceContactRequestMergeFieldsValueAddr1) GetCountry() *string {
	if c == nil {
		return nil
	}
	return c.Country
}

func (c *CreateAudienceContactRequestMergeFieldsValueAddr1) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *CreateAudienceContactRequestMergeFieldsValueAddr1) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetAddr1 sets the Addr1 field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAudienceContactRequestMergeFieldsValueAddr1) SetAddr1(addr1 string) {
	c.Addr1 = addr1
	c.require(createAudienceContactRequestMergeFieldsValueAddr1FieldAddr1)
}

// SetAddr2 sets the Addr2 field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAudienceContactRequestMergeFieldsValueAddr1) SetAddr2(addr2 *string) {
	c.Addr2 = addr2
	c.require(createAudienceContactRequestMergeFieldsValueAddr1FieldAddr2)
}

// SetCity sets the City field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAudienceContactRequestMergeFieldsValueAddr1) SetCity(city string) {
	c.City = city
	c.require(createAudienceContactRequestMergeFieldsValueAddr1FieldCity)
}

// SetState sets the State field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAudienceContactRequestMergeFieldsValueAddr1) SetState(state string) {
	c.State = state
	c.require(createAudienceContactRequestMergeFieldsValueAddr1FieldState)
}

// SetZip sets the Zip field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAudienceContactRequestMergeFieldsValueAddr1) SetZip(zip string) {
	c.Zip = zip
	c.require(createAudienceContactRequestMergeFieldsValueAddr1FieldZip)
}

// SetCountry sets the Country field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAudienceContactRequestMergeFieldsValueAddr1) SetCountry(country *string) {
	c.Country = country
	c.require(createAudienceContactRequestMergeFieldsValueAddr1FieldCountry)
}

func (c *CreateAudienceContactRequestMergeFieldsValueAddr1) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateAudienceContactRequestMergeFieldsValueAddr1
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateAudienceContactRequestMergeFieldsValueAddr1(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateAudienceContactRequestMergeFieldsValueAddr1) MarshalJSON() ([]byte, error) {
	type embed CreateAudienceContactRequestMergeFieldsValueAddr1
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *CreateAudienceContactRequestMergeFieldsValueAddr1) String() string {
	if c == nil {
		return "<nil>"
	}
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

var (
	createAudienceContactRequestSmsChannelFieldMarketingConsent = big.NewInt(1 << 0)
	createAudienceContactRequestSmsChannelFieldSmsPhone         = big.NewInt(1 << 1)
)

type CreateAudienceContactRequestSmsChannel struct {
	// A contact's current consent status for SMS marketing communications. See the [Audiences (BETA) documentation](https://mailchimp.com/developer/marketing/docs/audiences-introduction) to learn about supported values.
	MarketingConsent *CreateAudienceContactRequestSmsChannelMarketingConsent `json:"marketing_consent,omitempty" url:"marketing_consent,omitempty"`
	// SMS Phone Number
	SmsPhone *string `json:"sms_phone,omitempty" url:"sms_phone,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateAudienceContactRequestSmsChannel) GetMarketingConsent() *CreateAudienceContactRequestSmsChannelMarketingConsent {
	if c == nil {
		return nil
	}
	return c.MarketingConsent
}

func (c *CreateAudienceContactRequestSmsChannel) GetSmsPhone() *string {
	if c == nil {
		return nil
	}
	return c.SmsPhone
}

func (c *CreateAudienceContactRequestSmsChannel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *CreateAudienceContactRequestSmsChannel) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetMarketingConsent sets the MarketingConsent field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAudienceContactRequestSmsChannel) SetMarketingConsent(marketingConsent *CreateAudienceContactRequestSmsChannelMarketingConsent) {
	c.MarketingConsent = marketingConsent
	c.require(createAudienceContactRequestSmsChannelFieldMarketingConsent)
}

// SetSmsPhone sets the SmsPhone field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAudienceContactRequestSmsChannel) SetSmsPhone(smsPhone *string) {
	c.SmsPhone = smsPhone
	c.require(createAudienceContactRequestSmsChannelFieldSmsPhone)
}

func (c *CreateAudienceContactRequestSmsChannel) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateAudienceContactRequestSmsChannel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateAudienceContactRequestSmsChannel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateAudienceContactRequestSmsChannel) MarshalJSON() ([]byte, error) {
	type embed CreateAudienceContactRequestSmsChannel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *CreateAudienceContactRequestSmsChannel) String() string {
	if c == nil {
		return "<nil>"
	}
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

// A contact's current consent status for SMS marketing communications. See the [Audiences (BETA) documentation](https://mailchimp.com/developer/marketing/docs/audiences-introduction) to learn about supported values.
var (
	createAudienceContactRequestSmsChannelMarketingConsentFieldSource     = big.NewInt(1 << 0)
	createAudienceContactRequestSmsChannelMarketingConsentFieldStatus     = big.NewInt(1 << 1)
	createAudienceContactRequestSmsChannelMarketingConsentFieldCapturedAt = big.NewInt(1 << 2)
)

type CreateAudienceContactRequestSmsChannelMarketingConsent struct {
	// The source from which the parent's entity was created.
	Source *CreateAudienceContactRequestSmsChannelMarketingConsentSource `json:"source,omitempty" url:"source,omitempty"`
	// The contact's SMS marketing consent status. Use `confirmed` for double opt-in audiences, `consented` for single opt-in audiences.
	Status *CreateAudienceContactRequestSmsChannelMarketingConsentStatus `json:"status,omitempty" url:"status,omitempty"`
	// The timestamp when SMS marketing consent was captured (ISO 8601). Only accepted and returned when status is `confirmed`. The timestamp of the consent state change being recorded. Defaults to the current time if not provided. If the contact already has a consent timestamp on record that is equal to or newer than the supplied value, the supplied value is ignored (staleness guard); to update the consent timestamp supply a value strictly newer than the stored one.
	CapturedAt *time.Time `json:"captured_at,omitempty" url:"captured_at,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateAudienceContactRequestSmsChannelMarketingConsent) GetSource() *CreateAudienceContactRequestSmsChannelMarketingConsentSource {
	if c == nil {
		return nil
	}
	return c.Source
}

func (c *CreateAudienceContactRequestSmsChannelMarketingConsent) GetStatus() *CreateAudienceContactRequestSmsChannelMarketingConsentStatus {
	if c == nil {
		return nil
	}
	return c.Status
}

func (c *CreateAudienceContactRequestSmsChannelMarketingConsent) GetCapturedAt() *time.Time {
	if c == nil {
		return nil
	}
	return c.CapturedAt
}

func (c *CreateAudienceContactRequestSmsChannelMarketingConsent) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *CreateAudienceContactRequestSmsChannelMarketingConsent) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetSource sets the Source field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAudienceContactRequestSmsChannelMarketingConsent) SetSource(source *CreateAudienceContactRequestSmsChannelMarketingConsentSource) {
	c.Source = source
	c.require(createAudienceContactRequestSmsChannelMarketingConsentFieldSource)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAudienceContactRequestSmsChannelMarketingConsent) SetStatus(status *CreateAudienceContactRequestSmsChannelMarketingConsentStatus) {
	c.Status = status
	c.require(createAudienceContactRequestSmsChannelMarketingConsentFieldStatus)
}

// SetCapturedAt sets the CapturedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAudienceContactRequestSmsChannelMarketingConsent) SetCapturedAt(capturedAt *time.Time) {
	c.CapturedAt = capturedAt
	c.require(createAudienceContactRequestSmsChannelMarketingConsentFieldCapturedAt)
}

func (c *CreateAudienceContactRequestSmsChannelMarketingConsent) UnmarshalJSON(data []byte) error {
	type embed CreateAudienceContactRequestSmsChannelMarketingConsent
	var unmarshaler = struct {
		embed
		CapturedAt *internal.DateTime `json:"captured_at,omitempty"`
	}{
		embed: embed(*c),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*c = CreateAudienceContactRequestSmsChannelMarketingConsent(unmarshaler.embed)
	c.CapturedAt = unmarshaler.CapturedAt.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateAudienceContactRequestSmsChannelMarketingConsent) MarshalJSON() ([]byte, error) {
	type embed CreateAudienceContactRequestSmsChannelMarketingConsent
	var marshaler = struct {
		embed
		CapturedAt *internal.DateTime `json:"captured_at,omitempty"`
	}{
		embed:      embed(*c),
		CapturedAt: internal.NewOptionalDateTime(c.CapturedAt),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *CreateAudienceContactRequestSmsChannelMarketingConsent) String() string {
	if c == nil {
		return "<nil>"
	}
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

// The source from which the parent's entity was created.
var (
	createAudienceContactRequestSmsChannelMarketingConsentSourceFieldName = big.NewInt(1 << 0)
)

type CreateAudienceContactRequestSmsChannelMarketingConsentSource struct {
	// The name of the entity's source
	Name *string `json:"name,omitempty" url:"name,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateAudienceContactRequestSmsChannelMarketingConsentSource) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CreateAudienceContactRequestSmsChannelMarketingConsentSource) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *CreateAudienceContactRequestSmsChannelMarketingConsentSource) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAudienceContactRequestSmsChannelMarketingConsentSource) SetName(name *string) {
	c.Name = name
	c.require(createAudienceContactRequestSmsChannelMarketingConsentSourceFieldName)
}

func (c *CreateAudienceContactRequestSmsChannelMarketingConsentSource) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateAudienceContactRequestSmsChannelMarketingConsentSource
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateAudienceContactRequestSmsChannelMarketingConsentSource(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateAudienceContactRequestSmsChannelMarketingConsentSource) MarshalJSON() ([]byte, error) {
	type embed CreateAudienceContactRequestSmsChannelMarketingConsentSource
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *CreateAudienceContactRequestSmsChannelMarketingConsentSource) String() string {
	if c == nil {
		return "<nil>"
	}
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

// The contact's SMS marketing consent status. Use `confirmed` for double opt-in audiences, `consented` for single opt-in audiences.
type CreateAudienceContactRequestSmsChannelMarketingConsentStatus string

const (
	CreateAudienceContactRequestSmsChannelMarketingConsentStatusConsented CreateAudienceContactRequestSmsChannelMarketingConsentStatus = "consented"
	CreateAudienceContactRequestSmsChannelMarketingConsentStatusConfirmed CreateAudienceContactRequestSmsChannelMarketingConsentStatus = "confirmed"
	CreateAudienceContactRequestSmsChannelMarketingConsentStatusUnknown   CreateAudienceContactRequestSmsChannelMarketingConsentStatus = "unknown"
)

func NewCreateAudienceContactRequestSmsChannelMarketingConsentStatusFromString(s string) (CreateAudienceContactRequestSmsChannelMarketingConsentStatus, error) {
	switch s {
	case "consented":
		return CreateAudienceContactRequestSmsChannelMarketingConsentStatusConsented, nil
	case "confirmed":
		return CreateAudienceContactRequestSmsChannelMarketingConsentStatusConfirmed, nil
	case "unknown":
		return CreateAudienceContactRequestSmsChannelMarketingConsentStatusUnknown, nil
	}
	var t CreateAudienceContactRequestSmsChannelMarketingConsentStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c CreateAudienceContactRequestSmsChannelMarketingConsentStatus) Ptr() *CreateAudienceContactRequestSmsChannelMarketingConsentStatus {
	return &c
}

type CreateAudienceContactRequestTagsItem struct {
	String                                   string
	CreateAudienceContactRequestTagsItemName *CreateAudienceContactRequestTagsItemName

	typ string
}

func (c *CreateAudienceContactRequestTagsItem) GetString() string {
	if c == nil {
		return ""
	}
	return c.String
}

func (c *CreateAudienceContactRequestTagsItem) GetCreateAudienceContactRequestTagsItemName() *CreateAudienceContactRequestTagsItemName {
	if c == nil {
		return nil
	}
	return c.CreateAudienceContactRequestTagsItemName
}

func (c *CreateAudienceContactRequestTagsItem) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		c.typ = "String"
		c.String = valueString
		return nil
	}
	valueCreateAudienceContactRequestTagsItemName := new(CreateAudienceContactRequestTagsItemName)
	if err := json.Unmarshal(data, &valueCreateAudienceContactRequestTagsItemName); err == nil {
		c.typ = "CreateAudienceContactRequestTagsItemName"
		c.CreateAudienceContactRequestTagsItemName = valueCreateAudienceContactRequestTagsItemName
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, c)
}

func (c CreateAudienceContactRequestTagsItem) MarshalJSON() ([]byte, error) {
	if c.typ == "String" || c.String != "" {
		return json.Marshal(c.String)
	}
	if c.typ == "CreateAudienceContactRequestTagsItemName" || c.CreateAudienceContactRequestTagsItemName != nil {
		return json.Marshal(c.CreateAudienceContactRequestTagsItemName)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", c)
}

type CreateAudienceContactRequestTagsItemVisitor interface {
	VisitString(string) error
	VisitCreateAudienceContactRequestTagsItemName(*CreateAudienceContactRequestTagsItemName) error
}

func (c *CreateAudienceContactRequestTagsItem) Accept(visitor CreateAudienceContactRequestTagsItemVisitor) error {
	if c.typ == "String" || c.String != "" {
		return visitor.VisitString(c.String)
	}
	if c.typ == "CreateAudienceContactRequestTagsItemName" || c.CreateAudienceContactRequestTagsItemName != nil {
		return visitor.VisitCreateAudienceContactRequestTagsItemName(c.CreateAudienceContactRequestTagsItemName)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", c)
}

var (
	createAudienceContactRequestTagsItemNameFieldName   = big.NewInt(1 << 0)
	createAudienceContactRequestTagsItemNameFieldStatus = big.NewInt(1 << 1)
)

type CreateAudienceContactRequestTagsItemName struct {
	Name   string                                         `json:"name" url:"name"`
	Status CreateAudienceContactRequestTagsItemNameStatus `json:"status" url:"status"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateAudienceContactRequestTagsItemName) GetName() string {
	if c == nil {
		return ""
	}
	return c.Name
}

func (c *CreateAudienceContactRequestTagsItemName) GetStatus() CreateAudienceContactRequestTagsItemNameStatus {
	if c == nil {
		return ""
	}
	return c.Status
}

func (c *CreateAudienceContactRequestTagsItemName) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *CreateAudienceContactRequestTagsItemName) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAudienceContactRequestTagsItemName) SetName(name string) {
	c.Name = name
	c.require(createAudienceContactRequestTagsItemNameFieldName)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAudienceContactRequestTagsItemName) SetStatus(status CreateAudienceContactRequestTagsItemNameStatus) {
	c.Status = status
	c.require(createAudienceContactRequestTagsItemNameFieldStatus)
}

func (c *CreateAudienceContactRequestTagsItemName) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateAudienceContactRequestTagsItemName
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateAudienceContactRequestTagsItemName(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateAudienceContactRequestTagsItemName) MarshalJSON() ([]byte, error) {
	type embed CreateAudienceContactRequestTagsItemName
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *CreateAudienceContactRequestTagsItemName) String() string {
	if c == nil {
		return "<nil>"
	}
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type CreateAudienceContactRequestTagsItemNameStatus string

const (
	CreateAudienceContactRequestTagsItemNameStatusActive   CreateAudienceContactRequestTagsItemNameStatus = "active"
	CreateAudienceContactRequestTagsItemNameStatusInactive CreateAudienceContactRequestTagsItemNameStatus = "inactive"
)

func NewCreateAudienceContactRequestTagsItemNameStatusFromString(s string) (CreateAudienceContactRequestTagsItemNameStatus, error) {
	switch s {
	case "active":
		return CreateAudienceContactRequestTagsItemNameStatusActive, nil
	case "inactive":
		return CreateAudienceContactRequestTagsItemNameStatusInactive, nil
	}
	var t CreateAudienceContactRequestTagsItemNameStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c CreateAudienceContactRequestTagsItemNameStatus) Ptr() *CreateAudienceContactRequestTagsItemNameStatus {
	return &c
}

type GetAudienceContactListRequestSortDir string

const (
	GetAudienceContactListRequestSortDirAsc  GetAudienceContactListRequestSortDir = "ASC"
	GetAudienceContactListRequestSortDirDesc GetAudienceContactListRequestSortDir = "DESC"
)

func NewGetAudienceContactListRequestSortDirFromString(s string) (GetAudienceContactListRequestSortDir, error) {
	switch s {
	case "ASC":
		return GetAudienceContactListRequestSortDirAsc, nil
	case "DESC":
		return GetAudienceContactListRequestSortDirDesc, nil
	}
	var t GetAudienceContactListRequestSortDir
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (g GetAudienceContactListRequestSortDir) Ptr() *GetAudienceContactListRequestSortDir {
	return &g
}

type GetAudienceContactListRequestSortField string

const (
	GetAudienceContactListRequestSortFieldCreatedAt GetAudienceContactListRequestSortField = "created_at"
	GetAudienceContactListRequestSortFieldUpdatedAt GetAudienceContactListRequestSortField = "updated_at"
)

func NewGetAudienceContactListRequestSortFieldFromString(s string) (GetAudienceContactListRequestSortField, error) {
	switch s {
	case "created_at":
		return GetAudienceContactListRequestSortFieldCreatedAt, nil
	case "updated_at":
		return GetAudienceContactListRequestSortFieldUpdatedAt, nil
	}
	var t GetAudienceContactListRequestSortField
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (g GetAudienceContactListRequestSortField) Ptr() *GetAudienceContactListRequestSortField {
	return &g
}

// An array of objects, each representing a contact record.
var (
	getAudienceContactListResponseFieldContacts   = big.NewInt(1 << 0)
	getAudienceContactListResponseFieldNextCursor = big.NewInt(1 << 1)
	getAudienceContactListResponseFieldLinks      = big.NewInt(1 << 2)
)

type GetAudienceContactListResponse struct {
	// An array of objects, each representing a contact record.
	Contacts []*AudiencesContact `json:"contacts,omitempty" url:"contacts,omitempty"`
	// A cursor pointing to the last item on this page of the collection. Paginate through a collection of records by setting the `cursor` parameter on a subsequent request to this value.
	NextCursor *string `json:"next_cursor,omitempty" url:"next_cursor,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []*GetAudienceContactListResponseLinksItem `json:"_links,omitempty" url:"_links,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (g *GetAudienceContactListResponse) GetContacts() []*AudiencesContact {
	if g == nil {
		return nil
	}
	return g.Contacts
}

func (g *GetAudienceContactListResponse) GetNextCursor() *string {
	if g == nil {
		return nil
	}
	return g.NextCursor
}

func (g *GetAudienceContactListResponse) GetLinks() []*GetAudienceContactListResponseLinksItem {
	if g == nil {
		return nil
	}
	return g.Links
}

func (g *GetAudienceContactListResponse) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g *GetAudienceContactListResponse) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetContacts sets the Contacts field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAudienceContactListResponse) SetContacts(contacts []*AudiencesContact) {
	g.Contacts = contacts
	g.require(getAudienceContactListResponseFieldContacts)
}

// SetNextCursor sets the NextCursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAudienceContactListResponse) SetNextCursor(nextCursor *string) {
	g.NextCursor = nextCursor
	g.require(getAudienceContactListResponseFieldNextCursor)
}

// SetLinks sets the Links field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAudienceContactListResponse) SetLinks(links []*GetAudienceContactListResponseLinksItem) {
	g.Links = links
	g.require(getAudienceContactListResponseFieldLinks)
}

func (g *GetAudienceContactListResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetAudienceContactListResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetAudienceContactListResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetAudienceContactListResponse) MarshalJSON() ([]byte, error) {
	type embed GetAudienceContactListResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, g.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (g *GetAudienceContactListResponse) String() string {
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

// This object represents a link from the resource where it is found to another resource or action that may be performed.
var (
	getAudienceContactListResponseLinksItemFieldHref         = big.NewInt(1 << 0)
	getAudienceContactListResponseLinksItemFieldMethod       = big.NewInt(1 << 1)
	getAudienceContactListResponseLinksItemFieldRel          = big.NewInt(1 << 2)
	getAudienceContactListResponseLinksItemFieldSchema       = big.NewInt(1 << 3)
	getAudienceContactListResponseLinksItemFieldTargetSchema = big.NewInt(1 << 4)
)

type GetAudienceContactListResponseLinksItem struct {
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href *string `json:"href,omitempty" url:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method *GetAudienceContactListResponseLinksItemMethod `json:"method,omitempty" url:"method,omitempty"`
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

func (g *GetAudienceContactListResponseLinksItem) GetHref() *string {
	if g == nil {
		return nil
	}
	return g.Href
}

func (g *GetAudienceContactListResponseLinksItem) GetMethod() *GetAudienceContactListResponseLinksItemMethod {
	if g == nil {
		return nil
	}
	return g.Method
}

func (g *GetAudienceContactListResponseLinksItem) GetRel() *string {
	if g == nil {
		return nil
	}
	return g.Rel
}

func (g *GetAudienceContactListResponseLinksItem) GetSchema() *string {
	if g == nil {
		return nil
	}
	return g.Schema
}

func (g *GetAudienceContactListResponseLinksItem) GetTargetSchema() *string {
	if g == nil {
		return nil
	}
	return g.TargetSchema
}

func (g *GetAudienceContactListResponseLinksItem) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g *GetAudienceContactListResponseLinksItem) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetHref sets the Href field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAudienceContactListResponseLinksItem) SetHref(href *string) {
	g.Href = href
	g.require(getAudienceContactListResponseLinksItemFieldHref)
}

// SetMethod sets the Method field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAudienceContactListResponseLinksItem) SetMethod(method *GetAudienceContactListResponseLinksItemMethod) {
	g.Method = method
	g.require(getAudienceContactListResponseLinksItemFieldMethod)
}

// SetRel sets the Rel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAudienceContactListResponseLinksItem) SetRel(rel *string) {
	g.Rel = rel
	g.require(getAudienceContactListResponseLinksItemFieldRel)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAudienceContactListResponseLinksItem) SetSchema(schema *string) {
	g.Schema = schema
	g.require(getAudienceContactListResponseLinksItemFieldSchema)
}

// SetTargetSchema sets the TargetSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAudienceContactListResponseLinksItem) SetTargetSchema(targetSchema *string) {
	g.TargetSchema = targetSchema
	g.require(getAudienceContactListResponseLinksItemFieldTargetSchema)
}

func (g *GetAudienceContactListResponseLinksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler GetAudienceContactListResponseLinksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetAudienceContactListResponseLinksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetAudienceContactListResponseLinksItem) MarshalJSON() ([]byte, error) {
	type embed GetAudienceContactListResponseLinksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, g.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (g *GetAudienceContactListResponseLinksItem) String() string {
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

// The HTTP method that should be used when accessing the URL defined in 'href'.
type GetAudienceContactListResponseLinksItemMethod string

const (
	GetAudienceContactListResponseLinksItemMethodGet     GetAudienceContactListResponseLinksItemMethod = "GET"
	GetAudienceContactListResponseLinksItemMethodPost    GetAudienceContactListResponseLinksItemMethod = "POST"
	GetAudienceContactListResponseLinksItemMethodPut     GetAudienceContactListResponseLinksItemMethod = "PUT"
	GetAudienceContactListResponseLinksItemMethodPatch   GetAudienceContactListResponseLinksItemMethod = "PATCH"
	GetAudienceContactListResponseLinksItemMethodDelete  GetAudienceContactListResponseLinksItemMethod = "DELETE"
	GetAudienceContactListResponseLinksItemMethodOptions GetAudienceContactListResponseLinksItemMethod = "OPTIONS"
	GetAudienceContactListResponseLinksItemMethodHead    GetAudienceContactListResponseLinksItemMethod = "HEAD"
)

func NewGetAudienceContactListResponseLinksItemMethodFromString(s string) (GetAudienceContactListResponseLinksItemMethod, error) {
	switch s {
	case "GET":
		return GetAudienceContactListResponseLinksItemMethodGet, nil
	case "POST":
		return GetAudienceContactListResponseLinksItemMethodPost, nil
	case "PUT":
		return GetAudienceContactListResponseLinksItemMethodPut, nil
	case "PATCH":
		return GetAudienceContactListResponseLinksItemMethodPatch, nil
	case "DELETE":
		return GetAudienceContactListResponseLinksItemMethodDelete, nil
	case "OPTIONS":
		return GetAudienceContactListResponseLinksItemMethodOptions, nil
	case "HEAD":
		return GetAudienceContactListResponseLinksItemMethodHead, nil
	}
	var t GetAudienceContactListResponseLinksItemMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (g GetAudienceContactListResponseLinksItemMethod) Ptr() *GetAudienceContactListResponseLinksItemMethod {
	return &g
}

type PatchAudienceContactRequestDataMode string

const (
	PatchAudienceContactRequestDataModeHistorical PatchAudienceContactRequestDataMode = "historical"
	PatchAudienceContactRequestDataModeLive       PatchAudienceContactRequestDataMode = "live"
)

func NewPatchAudienceContactRequestDataModeFromString(s string) (PatchAudienceContactRequestDataMode, error) {
	switch s {
	case "historical":
		return PatchAudienceContactRequestDataModeHistorical, nil
	case "live":
		return PatchAudienceContactRequestDataModeLive, nil
	}
	var t PatchAudienceContactRequestDataMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (p PatchAudienceContactRequestDataMode) Ptr() *PatchAudienceContactRequestDataMode {
	return &p
}

var (
	patchAudienceContactRequestEmailChannelFieldEmail            = big.NewInt(1 << 0)
	patchAudienceContactRequestEmailChannelFieldMarketingConsent = big.NewInt(1 << 1)
)

type PatchAudienceContactRequestEmailChannel struct {
	// Email address
	Email *string `json:"email,omitempty" url:"email,omitempty"`
	// A contact's current consent status for email marketing communications. See the [Audiences (BETA) documentation](https://mailchimp.com/developer/marketing/docs/audiences-introduction) to learn about supported values.
	MarketingConsent *PatchAudienceContactRequestEmailChannelMarketingConsent `json:"marketing_consent,omitempty" url:"marketing_consent,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PatchAudienceContactRequestEmailChannel) GetEmail() *string {
	if p == nil {
		return nil
	}
	return p.Email
}

func (p *PatchAudienceContactRequestEmailChannel) GetMarketingConsent() *PatchAudienceContactRequestEmailChannelMarketingConsent {
	if p == nil {
		return nil
	}
	return p.MarketingConsent
}

func (p *PatchAudienceContactRequestEmailChannel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p *PatchAudienceContactRequestEmailChannel) require(field *big.Int) {
	if p.explicitFields == nil {
		p.explicitFields = big.NewInt(0)
	}
	p.explicitFields.Or(p.explicitFields, field)
}

// SetEmail sets the Email field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PatchAudienceContactRequestEmailChannel) SetEmail(email *string) {
	p.Email = email
	p.require(patchAudienceContactRequestEmailChannelFieldEmail)
}

// SetMarketingConsent sets the MarketingConsent field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PatchAudienceContactRequestEmailChannel) SetMarketingConsent(marketingConsent *PatchAudienceContactRequestEmailChannelMarketingConsent) {
	p.MarketingConsent = marketingConsent
	p.require(patchAudienceContactRequestEmailChannelFieldMarketingConsent)
}

func (p *PatchAudienceContactRequestEmailChannel) UnmarshalJSON(data []byte) error {
	type unmarshaler PatchAudienceContactRequestEmailChannel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PatchAudienceContactRequestEmailChannel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PatchAudienceContactRequestEmailChannel) MarshalJSON() ([]byte, error) {
	type embed PatchAudienceContactRequestEmailChannel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, p.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (p *PatchAudienceContactRequestEmailChannel) String() string {
	if p == nil {
		return "<nil>"
	}
	if len(p.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(p.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

// A contact's current consent status for email marketing communications. See the [Audiences (BETA) documentation](https://mailchimp.com/developer/marketing/docs/audiences-introduction) to learn about supported values.
var (
	patchAudienceContactRequestEmailChannelMarketingConsentFieldSource     = big.NewInt(1 << 0)
	patchAudienceContactRequestEmailChannelMarketingConsentFieldStatus     = big.NewInt(1 << 1)
	patchAudienceContactRequestEmailChannelMarketingConsentFieldCapturedAt = big.NewInt(1 << 2)
)

type PatchAudienceContactRequestEmailChannelMarketingConsent struct {
	// The source from which the parent's entity was created.
	Source *PatchAudienceContactRequestEmailChannelMarketingConsentSource `json:"source,omitempty" url:"source,omitempty"`
	Status *PatchAudienceContactRequestEmailChannelMarketingConsentStatus `json:"status,omitempty" url:"status,omitempty"`
	// The ISO 8601 timestamp when the email marketing consent state was recorded; accepted and returned only when status is `confirmed` or `consented`; defaults to the current time if omitted; ignored if older than an existing stored timestamp (staleness guard).
	CapturedAt *time.Time `json:"captured_at,omitempty" url:"captured_at,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PatchAudienceContactRequestEmailChannelMarketingConsent) GetSource() *PatchAudienceContactRequestEmailChannelMarketingConsentSource {
	if p == nil {
		return nil
	}
	return p.Source
}

func (p *PatchAudienceContactRequestEmailChannelMarketingConsent) GetStatus() *PatchAudienceContactRequestEmailChannelMarketingConsentStatus {
	if p == nil {
		return nil
	}
	return p.Status
}

func (p *PatchAudienceContactRequestEmailChannelMarketingConsent) GetCapturedAt() *time.Time {
	if p == nil {
		return nil
	}
	return p.CapturedAt
}

func (p *PatchAudienceContactRequestEmailChannelMarketingConsent) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p *PatchAudienceContactRequestEmailChannelMarketingConsent) require(field *big.Int) {
	if p.explicitFields == nil {
		p.explicitFields = big.NewInt(0)
	}
	p.explicitFields.Or(p.explicitFields, field)
}

// SetSource sets the Source field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PatchAudienceContactRequestEmailChannelMarketingConsent) SetSource(source *PatchAudienceContactRequestEmailChannelMarketingConsentSource) {
	p.Source = source
	p.require(patchAudienceContactRequestEmailChannelMarketingConsentFieldSource)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PatchAudienceContactRequestEmailChannelMarketingConsent) SetStatus(status *PatchAudienceContactRequestEmailChannelMarketingConsentStatus) {
	p.Status = status
	p.require(patchAudienceContactRequestEmailChannelMarketingConsentFieldStatus)
}

// SetCapturedAt sets the CapturedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PatchAudienceContactRequestEmailChannelMarketingConsent) SetCapturedAt(capturedAt *time.Time) {
	p.CapturedAt = capturedAt
	p.require(patchAudienceContactRequestEmailChannelMarketingConsentFieldCapturedAt)
}

func (p *PatchAudienceContactRequestEmailChannelMarketingConsent) UnmarshalJSON(data []byte) error {
	type embed PatchAudienceContactRequestEmailChannelMarketingConsent
	var unmarshaler = struct {
		embed
		CapturedAt *internal.DateTime `json:"captured_at,omitempty"`
	}{
		embed: embed(*p),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*p = PatchAudienceContactRequestEmailChannelMarketingConsent(unmarshaler.embed)
	p.CapturedAt = unmarshaler.CapturedAt.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PatchAudienceContactRequestEmailChannelMarketingConsent) MarshalJSON() ([]byte, error) {
	type embed PatchAudienceContactRequestEmailChannelMarketingConsent
	var marshaler = struct {
		embed
		CapturedAt *internal.DateTime `json:"captured_at,omitempty"`
	}{
		embed:      embed(*p),
		CapturedAt: internal.NewOptionalDateTime(p.CapturedAt),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, p.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (p *PatchAudienceContactRequestEmailChannelMarketingConsent) String() string {
	if p == nil {
		return "<nil>"
	}
	if len(p.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(p.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

// The source from which the parent's entity was created.
var (
	patchAudienceContactRequestEmailChannelMarketingConsentSourceFieldName = big.NewInt(1 << 0)
)

type PatchAudienceContactRequestEmailChannelMarketingConsentSource struct {
	// The name of the entity's source
	Name *string `json:"name,omitempty" url:"name,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PatchAudienceContactRequestEmailChannelMarketingConsentSource) GetName() *string {
	if p == nil {
		return nil
	}
	return p.Name
}

func (p *PatchAudienceContactRequestEmailChannelMarketingConsentSource) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p *PatchAudienceContactRequestEmailChannelMarketingConsentSource) require(field *big.Int) {
	if p.explicitFields == nil {
		p.explicitFields = big.NewInt(0)
	}
	p.explicitFields.Or(p.explicitFields, field)
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PatchAudienceContactRequestEmailChannelMarketingConsentSource) SetName(name *string) {
	p.Name = name
	p.require(patchAudienceContactRequestEmailChannelMarketingConsentSourceFieldName)
}

func (p *PatchAudienceContactRequestEmailChannelMarketingConsentSource) UnmarshalJSON(data []byte) error {
	type unmarshaler PatchAudienceContactRequestEmailChannelMarketingConsentSource
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PatchAudienceContactRequestEmailChannelMarketingConsentSource(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PatchAudienceContactRequestEmailChannelMarketingConsentSource) MarshalJSON() ([]byte, error) {
	type embed PatchAudienceContactRequestEmailChannelMarketingConsentSource
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, p.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (p *PatchAudienceContactRequestEmailChannelMarketingConsentSource) String() string {
	if p == nil {
		return "<nil>"
	}
	if len(p.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(p.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type PatchAudienceContactRequestEmailChannelMarketingConsentStatus string

const (
	PatchAudienceContactRequestEmailChannelMarketingConsentStatusConsented PatchAudienceContactRequestEmailChannelMarketingConsentStatus = "consented"
	PatchAudienceContactRequestEmailChannelMarketingConsentStatusDenied    PatchAudienceContactRequestEmailChannelMarketingConsentStatus = "denied"
	PatchAudienceContactRequestEmailChannelMarketingConsentStatusConfirmed PatchAudienceContactRequestEmailChannelMarketingConsentStatus = "confirmed"
	PatchAudienceContactRequestEmailChannelMarketingConsentStatusUnknown   PatchAudienceContactRequestEmailChannelMarketingConsentStatus = "unknown"
)

func NewPatchAudienceContactRequestEmailChannelMarketingConsentStatusFromString(s string) (PatchAudienceContactRequestEmailChannelMarketingConsentStatus, error) {
	switch s {
	case "consented":
		return PatchAudienceContactRequestEmailChannelMarketingConsentStatusConsented, nil
	case "denied":
		return PatchAudienceContactRequestEmailChannelMarketingConsentStatusDenied, nil
	case "confirmed":
		return PatchAudienceContactRequestEmailChannelMarketingConsentStatusConfirmed, nil
	case "unknown":
		return PatchAudienceContactRequestEmailChannelMarketingConsentStatusUnknown, nil
	}
	var t PatchAudienceContactRequestEmailChannelMarketingConsentStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (p PatchAudienceContactRequestEmailChannelMarketingConsentStatus) Ptr() *PatchAudienceContactRequestEmailChannelMarketingConsentStatus {
	return &p
}

type PatchAudienceContactRequestMergeFieldValidationMode string

const (
	PatchAudienceContactRequestMergeFieldValidationModeIgnoreRequiredChecks PatchAudienceContactRequestMergeFieldValidationMode = "ignore_required_checks"
	PatchAudienceContactRequestMergeFieldValidationModeStrict               PatchAudienceContactRequestMergeFieldValidationMode = "strict"
)

func NewPatchAudienceContactRequestMergeFieldValidationModeFromString(s string) (PatchAudienceContactRequestMergeFieldValidationMode, error) {
	switch s {
	case "ignore_required_checks":
		return PatchAudienceContactRequestMergeFieldValidationModeIgnoreRequiredChecks, nil
	case "strict":
		return PatchAudienceContactRequestMergeFieldValidationModeStrict, nil
	}
	var t PatchAudienceContactRequestMergeFieldValidationMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (p PatchAudienceContactRequestMergeFieldValidationMode) Ptr() *PatchAudienceContactRequestMergeFieldValidationMode {
	return &p
}

// This object's keys are merge tags (like FNAME). It's values are the values to be added to the merge field.
type PatchAudienceContactRequestMergeFieldsValue struct {
	PatchAudienceContactRequestMergeFieldsValueAddr1 *PatchAudienceContactRequestMergeFieldsValueAddr1
	String                                           string
	Double                                           float64

	typ string
}

func (p *PatchAudienceContactRequestMergeFieldsValue) GetPatchAudienceContactRequestMergeFieldsValueAddr1() *PatchAudienceContactRequestMergeFieldsValueAddr1 {
	if p == nil {
		return nil
	}
	return p.PatchAudienceContactRequestMergeFieldsValueAddr1
}

func (p *PatchAudienceContactRequestMergeFieldsValue) GetString() string {
	if p == nil {
		return ""
	}
	return p.String
}

func (p *PatchAudienceContactRequestMergeFieldsValue) GetDouble() float64 {
	if p == nil {
		return 0
	}
	return p.Double
}

func (p *PatchAudienceContactRequestMergeFieldsValue) UnmarshalJSON(data []byte) error {
	valuePatchAudienceContactRequestMergeFieldsValueAddr1 := new(PatchAudienceContactRequestMergeFieldsValueAddr1)
	if err := json.Unmarshal(data, &valuePatchAudienceContactRequestMergeFieldsValueAddr1); err == nil {
		p.typ = "PatchAudienceContactRequestMergeFieldsValueAddr1"
		p.PatchAudienceContactRequestMergeFieldsValueAddr1 = valuePatchAudienceContactRequestMergeFieldsValueAddr1
		return nil
	}
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		p.typ = "String"
		p.String = valueString
		return nil
	}
	var valueDouble float64
	if err := json.Unmarshal(data, &valueDouble); err == nil {
		p.typ = "Double"
		p.Double = valueDouble
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, p)
}

func (p PatchAudienceContactRequestMergeFieldsValue) MarshalJSON() ([]byte, error) {
	if p.typ == "PatchAudienceContactRequestMergeFieldsValueAddr1" || p.PatchAudienceContactRequestMergeFieldsValueAddr1 != nil {
		return json.Marshal(p.PatchAudienceContactRequestMergeFieldsValueAddr1)
	}
	if p.typ == "String" || p.String != "" {
		return json.Marshal(p.String)
	}
	if p.typ == "Double" || p.Double != 0 {
		return json.Marshal(p.Double)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", p)
}

type PatchAudienceContactRequestMergeFieldsValueVisitor interface {
	VisitPatchAudienceContactRequestMergeFieldsValueAddr1(*PatchAudienceContactRequestMergeFieldsValueAddr1) error
	VisitString(string) error
	VisitDouble(float64) error
}

func (p *PatchAudienceContactRequestMergeFieldsValue) Accept(visitor PatchAudienceContactRequestMergeFieldsValueVisitor) error {
	if p.typ == "PatchAudienceContactRequestMergeFieldsValueAddr1" || p.PatchAudienceContactRequestMergeFieldsValueAddr1 != nil {
		return visitor.VisitPatchAudienceContactRequestMergeFieldsValueAddr1(p.PatchAudienceContactRequestMergeFieldsValueAddr1)
	}
	if p.typ == "String" || p.String != "" {
		return visitor.VisitString(p.String)
	}
	if p.typ == "Double" || p.Double != 0 {
		return visitor.VisitDouble(p.Double)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", p)
}

var (
	patchAudienceContactRequestMergeFieldsValueAddr1FieldAddr1   = big.NewInt(1 << 0)
	patchAudienceContactRequestMergeFieldsValueAddr1FieldAddr2   = big.NewInt(1 << 1)
	patchAudienceContactRequestMergeFieldsValueAddr1FieldCity    = big.NewInt(1 << 2)
	patchAudienceContactRequestMergeFieldsValueAddr1FieldState   = big.NewInt(1 << 3)
	patchAudienceContactRequestMergeFieldsValueAddr1FieldZip     = big.NewInt(1 << 4)
	patchAudienceContactRequestMergeFieldsValueAddr1FieldCountry = big.NewInt(1 << 5)
)

type PatchAudienceContactRequestMergeFieldsValueAddr1 struct {
	Addr1   string  `json:"addr1" url:"addr1"`
	Addr2   *string `json:"addr2,omitempty" url:"addr2,omitempty"`
	City    string  `json:"city" url:"city"`
	State   string  `json:"state" url:"state"`
	Zip     string  `json:"zip" url:"zip"`
	Country *string `json:"country,omitempty" url:"country,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PatchAudienceContactRequestMergeFieldsValueAddr1) GetAddr1() string {
	if p == nil {
		return ""
	}
	return p.Addr1
}

func (p *PatchAudienceContactRequestMergeFieldsValueAddr1) GetAddr2() *string {
	if p == nil {
		return nil
	}
	return p.Addr2
}

func (p *PatchAudienceContactRequestMergeFieldsValueAddr1) GetCity() string {
	if p == nil {
		return ""
	}
	return p.City
}

func (p *PatchAudienceContactRequestMergeFieldsValueAddr1) GetState() string {
	if p == nil {
		return ""
	}
	return p.State
}

func (p *PatchAudienceContactRequestMergeFieldsValueAddr1) GetZip() string {
	if p == nil {
		return ""
	}
	return p.Zip
}

func (p *PatchAudienceContactRequestMergeFieldsValueAddr1) GetCountry() *string {
	if p == nil {
		return nil
	}
	return p.Country
}

func (p *PatchAudienceContactRequestMergeFieldsValueAddr1) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p *PatchAudienceContactRequestMergeFieldsValueAddr1) require(field *big.Int) {
	if p.explicitFields == nil {
		p.explicitFields = big.NewInt(0)
	}
	p.explicitFields.Or(p.explicitFields, field)
}

// SetAddr1 sets the Addr1 field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PatchAudienceContactRequestMergeFieldsValueAddr1) SetAddr1(addr1 string) {
	p.Addr1 = addr1
	p.require(patchAudienceContactRequestMergeFieldsValueAddr1FieldAddr1)
}

// SetAddr2 sets the Addr2 field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PatchAudienceContactRequestMergeFieldsValueAddr1) SetAddr2(addr2 *string) {
	p.Addr2 = addr2
	p.require(patchAudienceContactRequestMergeFieldsValueAddr1FieldAddr2)
}

// SetCity sets the City field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PatchAudienceContactRequestMergeFieldsValueAddr1) SetCity(city string) {
	p.City = city
	p.require(patchAudienceContactRequestMergeFieldsValueAddr1FieldCity)
}

// SetState sets the State field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PatchAudienceContactRequestMergeFieldsValueAddr1) SetState(state string) {
	p.State = state
	p.require(patchAudienceContactRequestMergeFieldsValueAddr1FieldState)
}

// SetZip sets the Zip field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PatchAudienceContactRequestMergeFieldsValueAddr1) SetZip(zip string) {
	p.Zip = zip
	p.require(patchAudienceContactRequestMergeFieldsValueAddr1FieldZip)
}

// SetCountry sets the Country field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PatchAudienceContactRequestMergeFieldsValueAddr1) SetCountry(country *string) {
	p.Country = country
	p.require(patchAudienceContactRequestMergeFieldsValueAddr1FieldCountry)
}

func (p *PatchAudienceContactRequestMergeFieldsValueAddr1) UnmarshalJSON(data []byte) error {
	type unmarshaler PatchAudienceContactRequestMergeFieldsValueAddr1
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PatchAudienceContactRequestMergeFieldsValueAddr1(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PatchAudienceContactRequestMergeFieldsValueAddr1) MarshalJSON() ([]byte, error) {
	type embed PatchAudienceContactRequestMergeFieldsValueAddr1
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, p.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (p *PatchAudienceContactRequestMergeFieldsValueAddr1) String() string {
	if p == nil {
		return "<nil>"
	}
	if len(p.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(p.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

var (
	patchAudienceContactRequestSmsChannelFieldMarketingConsent = big.NewInt(1 << 0)
	patchAudienceContactRequestSmsChannelFieldSmsPhone         = big.NewInt(1 << 1)
)

type PatchAudienceContactRequestSmsChannel struct {
	// A contact's current consent status for SMS marketing communications. See the [Audiences (BETA) documentation](https://mailchimp.com/developer/marketing/docs/audiences-introduction) to learn about supported values.
	MarketingConsent *PatchAudienceContactRequestSmsChannelMarketingConsent `json:"marketing_consent,omitempty" url:"marketing_consent,omitempty"`
	// SMS Phone Number
	SmsPhone *string `json:"sms_phone,omitempty" url:"sms_phone,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PatchAudienceContactRequestSmsChannel) GetMarketingConsent() *PatchAudienceContactRequestSmsChannelMarketingConsent {
	if p == nil {
		return nil
	}
	return p.MarketingConsent
}

func (p *PatchAudienceContactRequestSmsChannel) GetSmsPhone() *string {
	if p == nil {
		return nil
	}
	return p.SmsPhone
}

func (p *PatchAudienceContactRequestSmsChannel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p *PatchAudienceContactRequestSmsChannel) require(field *big.Int) {
	if p.explicitFields == nil {
		p.explicitFields = big.NewInt(0)
	}
	p.explicitFields.Or(p.explicitFields, field)
}

// SetMarketingConsent sets the MarketingConsent field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PatchAudienceContactRequestSmsChannel) SetMarketingConsent(marketingConsent *PatchAudienceContactRequestSmsChannelMarketingConsent) {
	p.MarketingConsent = marketingConsent
	p.require(patchAudienceContactRequestSmsChannelFieldMarketingConsent)
}

// SetSmsPhone sets the SmsPhone field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PatchAudienceContactRequestSmsChannel) SetSmsPhone(smsPhone *string) {
	p.SmsPhone = smsPhone
	p.require(patchAudienceContactRequestSmsChannelFieldSmsPhone)
}

func (p *PatchAudienceContactRequestSmsChannel) UnmarshalJSON(data []byte) error {
	type unmarshaler PatchAudienceContactRequestSmsChannel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PatchAudienceContactRequestSmsChannel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PatchAudienceContactRequestSmsChannel) MarshalJSON() ([]byte, error) {
	type embed PatchAudienceContactRequestSmsChannel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, p.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (p *PatchAudienceContactRequestSmsChannel) String() string {
	if p == nil {
		return "<nil>"
	}
	if len(p.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(p.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

// A contact's current consent status for SMS marketing communications. See the [Audiences (BETA) documentation](https://mailchimp.com/developer/marketing/docs/audiences-introduction) to learn about supported values.
var (
	patchAudienceContactRequestSmsChannelMarketingConsentFieldSource     = big.NewInt(1 << 0)
	patchAudienceContactRequestSmsChannelMarketingConsentFieldStatus     = big.NewInt(1 << 1)
	patchAudienceContactRequestSmsChannelMarketingConsentFieldCapturedAt = big.NewInt(1 << 2)
)

type PatchAudienceContactRequestSmsChannelMarketingConsent struct {
	// The source from which the parent's entity was created.
	Source *PatchAudienceContactRequestSmsChannelMarketingConsentSource `json:"source,omitempty" url:"source,omitempty"`
	// The contact's SMS marketing consent status. Use `confirmed` for double opt-in audiences, `consented` for single opt-in audiences. `denied` is accepted on PATCH/PUT only (not POST) and drives an API-initiated unsubscribe; it cannot be used when creating a new contact.
	Status *PatchAudienceContactRequestSmsChannelMarketingConsentStatus `json:"status,omitempty" url:"status,omitempty"`
	// The timestamp when SMS marketing consent was captured (ISO 8601). Only accepted and returned when status is `confirmed`. The timestamp of the consent state change being recorded. Defaults to the current time if not provided. If the contact already has a consent timestamp on record that is equal to or newer than the supplied value, the supplied value is ignored (staleness guard); to update the consent timestamp supply a value strictly newer than the stored one.
	CapturedAt *time.Time `json:"captured_at,omitempty" url:"captured_at,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PatchAudienceContactRequestSmsChannelMarketingConsent) GetSource() *PatchAudienceContactRequestSmsChannelMarketingConsentSource {
	if p == nil {
		return nil
	}
	return p.Source
}

func (p *PatchAudienceContactRequestSmsChannelMarketingConsent) GetStatus() *PatchAudienceContactRequestSmsChannelMarketingConsentStatus {
	if p == nil {
		return nil
	}
	return p.Status
}

func (p *PatchAudienceContactRequestSmsChannelMarketingConsent) GetCapturedAt() *time.Time {
	if p == nil {
		return nil
	}
	return p.CapturedAt
}

func (p *PatchAudienceContactRequestSmsChannelMarketingConsent) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p *PatchAudienceContactRequestSmsChannelMarketingConsent) require(field *big.Int) {
	if p.explicitFields == nil {
		p.explicitFields = big.NewInt(0)
	}
	p.explicitFields.Or(p.explicitFields, field)
}

// SetSource sets the Source field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PatchAudienceContactRequestSmsChannelMarketingConsent) SetSource(source *PatchAudienceContactRequestSmsChannelMarketingConsentSource) {
	p.Source = source
	p.require(patchAudienceContactRequestSmsChannelMarketingConsentFieldSource)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PatchAudienceContactRequestSmsChannelMarketingConsent) SetStatus(status *PatchAudienceContactRequestSmsChannelMarketingConsentStatus) {
	p.Status = status
	p.require(patchAudienceContactRequestSmsChannelMarketingConsentFieldStatus)
}

// SetCapturedAt sets the CapturedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PatchAudienceContactRequestSmsChannelMarketingConsent) SetCapturedAt(capturedAt *time.Time) {
	p.CapturedAt = capturedAt
	p.require(patchAudienceContactRequestSmsChannelMarketingConsentFieldCapturedAt)
}

func (p *PatchAudienceContactRequestSmsChannelMarketingConsent) UnmarshalJSON(data []byte) error {
	type embed PatchAudienceContactRequestSmsChannelMarketingConsent
	var unmarshaler = struct {
		embed
		CapturedAt *internal.DateTime `json:"captured_at,omitempty"`
	}{
		embed: embed(*p),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*p = PatchAudienceContactRequestSmsChannelMarketingConsent(unmarshaler.embed)
	p.CapturedAt = unmarshaler.CapturedAt.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PatchAudienceContactRequestSmsChannelMarketingConsent) MarshalJSON() ([]byte, error) {
	type embed PatchAudienceContactRequestSmsChannelMarketingConsent
	var marshaler = struct {
		embed
		CapturedAt *internal.DateTime `json:"captured_at,omitempty"`
	}{
		embed:      embed(*p),
		CapturedAt: internal.NewOptionalDateTime(p.CapturedAt),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, p.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (p *PatchAudienceContactRequestSmsChannelMarketingConsent) String() string {
	if p == nil {
		return "<nil>"
	}
	if len(p.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(p.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

// The source from which the parent's entity was created.
var (
	patchAudienceContactRequestSmsChannelMarketingConsentSourceFieldName = big.NewInt(1 << 0)
)

type PatchAudienceContactRequestSmsChannelMarketingConsentSource struct {
	// The name of the entity's source
	Name *string `json:"name,omitempty" url:"name,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PatchAudienceContactRequestSmsChannelMarketingConsentSource) GetName() *string {
	if p == nil {
		return nil
	}
	return p.Name
}

func (p *PatchAudienceContactRequestSmsChannelMarketingConsentSource) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p *PatchAudienceContactRequestSmsChannelMarketingConsentSource) require(field *big.Int) {
	if p.explicitFields == nil {
		p.explicitFields = big.NewInt(0)
	}
	p.explicitFields.Or(p.explicitFields, field)
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PatchAudienceContactRequestSmsChannelMarketingConsentSource) SetName(name *string) {
	p.Name = name
	p.require(patchAudienceContactRequestSmsChannelMarketingConsentSourceFieldName)
}

func (p *PatchAudienceContactRequestSmsChannelMarketingConsentSource) UnmarshalJSON(data []byte) error {
	type unmarshaler PatchAudienceContactRequestSmsChannelMarketingConsentSource
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PatchAudienceContactRequestSmsChannelMarketingConsentSource(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PatchAudienceContactRequestSmsChannelMarketingConsentSource) MarshalJSON() ([]byte, error) {
	type embed PatchAudienceContactRequestSmsChannelMarketingConsentSource
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, p.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (p *PatchAudienceContactRequestSmsChannelMarketingConsentSource) String() string {
	if p == nil {
		return "<nil>"
	}
	if len(p.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(p.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

// The contact's SMS marketing consent status. Use `confirmed` for double opt-in audiences, `consented` for single opt-in audiences. `denied` is accepted on PATCH/PUT only (not POST) and drives an API-initiated unsubscribe; it cannot be used when creating a new contact.
type PatchAudienceContactRequestSmsChannelMarketingConsentStatus string

const (
	PatchAudienceContactRequestSmsChannelMarketingConsentStatusConsented PatchAudienceContactRequestSmsChannelMarketingConsentStatus = "consented"
	PatchAudienceContactRequestSmsChannelMarketingConsentStatusConfirmed PatchAudienceContactRequestSmsChannelMarketingConsentStatus = "confirmed"
	PatchAudienceContactRequestSmsChannelMarketingConsentStatusDenied    PatchAudienceContactRequestSmsChannelMarketingConsentStatus = "denied"
	PatchAudienceContactRequestSmsChannelMarketingConsentStatusUnknown   PatchAudienceContactRequestSmsChannelMarketingConsentStatus = "unknown"
)

func NewPatchAudienceContactRequestSmsChannelMarketingConsentStatusFromString(s string) (PatchAudienceContactRequestSmsChannelMarketingConsentStatus, error) {
	switch s {
	case "consented":
		return PatchAudienceContactRequestSmsChannelMarketingConsentStatusConsented, nil
	case "confirmed":
		return PatchAudienceContactRequestSmsChannelMarketingConsentStatusConfirmed, nil
	case "denied":
		return PatchAudienceContactRequestSmsChannelMarketingConsentStatusDenied, nil
	case "unknown":
		return PatchAudienceContactRequestSmsChannelMarketingConsentStatusUnknown, nil
	}
	var t PatchAudienceContactRequestSmsChannelMarketingConsentStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (p PatchAudienceContactRequestSmsChannelMarketingConsentStatus) Ptr() *PatchAudienceContactRequestSmsChannelMarketingConsentStatus {
	return &p
}

type PatchAudienceContactRequestTagsItem struct {
	String                                  string
	PatchAudienceContactRequestTagsItemName *PatchAudienceContactRequestTagsItemName

	typ string
}

func (p *PatchAudienceContactRequestTagsItem) GetString() string {
	if p == nil {
		return ""
	}
	return p.String
}

func (p *PatchAudienceContactRequestTagsItem) GetPatchAudienceContactRequestTagsItemName() *PatchAudienceContactRequestTagsItemName {
	if p == nil {
		return nil
	}
	return p.PatchAudienceContactRequestTagsItemName
}

func (p *PatchAudienceContactRequestTagsItem) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		p.typ = "String"
		p.String = valueString
		return nil
	}
	valuePatchAudienceContactRequestTagsItemName := new(PatchAudienceContactRequestTagsItemName)
	if err := json.Unmarshal(data, &valuePatchAudienceContactRequestTagsItemName); err == nil {
		p.typ = "PatchAudienceContactRequestTagsItemName"
		p.PatchAudienceContactRequestTagsItemName = valuePatchAudienceContactRequestTagsItemName
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, p)
}

func (p PatchAudienceContactRequestTagsItem) MarshalJSON() ([]byte, error) {
	if p.typ == "String" || p.String != "" {
		return json.Marshal(p.String)
	}
	if p.typ == "PatchAudienceContactRequestTagsItemName" || p.PatchAudienceContactRequestTagsItemName != nil {
		return json.Marshal(p.PatchAudienceContactRequestTagsItemName)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", p)
}

type PatchAudienceContactRequestTagsItemVisitor interface {
	VisitString(string) error
	VisitPatchAudienceContactRequestTagsItemName(*PatchAudienceContactRequestTagsItemName) error
}

func (p *PatchAudienceContactRequestTagsItem) Accept(visitor PatchAudienceContactRequestTagsItemVisitor) error {
	if p.typ == "String" || p.String != "" {
		return visitor.VisitString(p.String)
	}
	if p.typ == "PatchAudienceContactRequestTagsItemName" || p.PatchAudienceContactRequestTagsItemName != nil {
		return visitor.VisitPatchAudienceContactRequestTagsItemName(p.PatchAudienceContactRequestTagsItemName)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", p)
}

var (
	patchAudienceContactRequestTagsItemNameFieldName   = big.NewInt(1 << 0)
	patchAudienceContactRequestTagsItemNameFieldStatus = big.NewInt(1 << 1)
)

type PatchAudienceContactRequestTagsItemName struct {
	Name   string                                        `json:"name" url:"name"`
	Status PatchAudienceContactRequestTagsItemNameStatus `json:"status" url:"status"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PatchAudienceContactRequestTagsItemName) GetName() string {
	if p == nil {
		return ""
	}
	return p.Name
}

func (p *PatchAudienceContactRequestTagsItemName) GetStatus() PatchAudienceContactRequestTagsItemNameStatus {
	if p == nil {
		return ""
	}
	return p.Status
}

func (p *PatchAudienceContactRequestTagsItemName) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p *PatchAudienceContactRequestTagsItemName) require(field *big.Int) {
	if p.explicitFields == nil {
		p.explicitFields = big.NewInt(0)
	}
	p.explicitFields.Or(p.explicitFields, field)
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PatchAudienceContactRequestTagsItemName) SetName(name string) {
	p.Name = name
	p.require(patchAudienceContactRequestTagsItemNameFieldName)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PatchAudienceContactRequestTagsItemName) SetStatus(status PatchAudienceContactRequestTagsItemNameStatus) {
	p.Status = status
	p.require(patchAudienceContactRequestTagsItemNameFieldStatus)
}

func (p *PatchAudienceContactRequestTagsItemName) UnmarshalJSON(data []byte) error {
	type unmarshaler PatchAudienceContactRequestTagsItemName
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PatchAudienceContactRequestTagsItemName(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PatchAudienceContactRequestTagsItemName) MarshalJSON() ([]byte, error) {
	type embed PatchAudienceContactRequestTagsItemName
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, p.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (p *PatchAudienceContactRequestTagsItemName) String() string {
	if p == nil {
		return "<nil>"
	}
	if len(p.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(p.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type PatchAudienceContactRequestTagsItemNameStatus string

const (
	PatchAudienceContactRequestTagsItemNameStatusActive   PatchAudienceContactRequestTagsItemNameStatus = "active"
	PatchAudienceContactRequestTagsItemNameStatusInactive PatchAudienceContactRequestTagsItemNameStatus = "inactive"
)

func NewPatchAudienceContactRequestTagsItemNameStatusFromString(s string) (PatchAudienceContactRequestTagsItemNameStatus, error) {
	switch s {
	case "active":
		return PatchAudienceContactRequestTagsItemNameStatusActive, nil
	case "inactive":
		return PatchAudienceContactRequestTagsItemNameStatusInactive, nil
	}
	var t PatchAudienceContactRequestTagsItemNameStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (p PatchAudienceContactRequestTagsItemNameStatus) Ptr() *PatchAudienceContactRequestTagsItemNameStatus {
	return &p
}
