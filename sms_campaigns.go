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
	createSmsCampaignsRequestFieldName             = big.NewInt(1 << 0)
	createSmsCampaignsRequestFieldListID           = big.NewInt(1 << 1)
	createSmsCampaignsRequestFieldFolderID         = big.NewInt(1 << 2)
	createSmsCampaignsRequestFieldSegments         = big.NewInt(1 << 3)
	createSmsCampaignsRequestFieldExcludedSegments = big.NewInt(1 << 4)
)

type CreateSmsCampaignsRequest struct {
	// The name of the campaign.
	Name string `json:"name" url:"-"`
	// The numeric ID of the list to send the campaign to.
	ListID *int `json:"list_id,omitempty" url:"-"`
	// The ID of the folder to place this campaign in.
	FolderID *string `json:"folder_id,omitempty" url:"-"`
	// The segment IDs to target for this campaign.
	Segments []int `json:"segments,omitempty" url:"-"`
	// The segment IDs to exclude from this campaign.
	ExcludedSegments []int `json:"excluded_segments,omitempty" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CreateSmsCampaignsRequest) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateSmsCampaignsRequest) SetName(name string) {
	c.Name = name
	c.require(createSmsCampaignsRequestFieldName)
}

// SetListID sets the ListID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateSmsCampaignsRequest) SetListID(listID *int) {
	c.ListID = listID
	c.require(createSmsCampaignsRequestFieldListID)
}

// SetFolderID sets the FolderID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateSmsCampaignsRequest) SetFolderID(folderID *string) {
	c.FolderID = folderID
	c.require(createSmsCampaignsRequestFieldFolderID)
}

// SetSegments sets the Segments field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateSmsCampaignsRequest) SetSegments(segments []int) {
	c.Segments = segments
	c.require(createSmsCampaignsRequestFieldSegments)
}

// SetExcludedSegments sets the ExcludedSegments field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateSmsCampaignsRequest) SetExcludedSegments(excludedSegments []int) {
	c.ExcludedSegments = excludedSegments
	c.require(createSmsCampaignsRequestFieldExcludedSegments)
}

func (c *CreateSmsCampaignsRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateSmsCampaignsRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*c = CreateSmsCampaignsRequest(body)
	return nil
}

func (c *CreateSmsCampaignsRequest) MarshalJSON() ([]byte, error) {
	type embed CreateSmsCampaignsRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

var (
	createActionCancelSendSmsCampaignsRequestFieldSmsCampaignID = big.NewInt(1 << 0)
)

type CreateActionCancelSendSmsCampaignsRequest struct {
	// The unique id for the SMS campaign.
	SmsCampaignID string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CreateActionCancelSendSmsCampaignsRequest) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetSmsCampaignID sets the SmsCampaignID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateActionCancelSendSmsCampaignsRequest) SetSmsCampaignID(smsCampaignID string) {
	c.SmsCampaignID = smsCampaignID
	c.require(createActionCancelSendSmsCampaignsRequestFieldSmsCampaignID)
}

var (
	createActionScheduleSmsCampaignsRequestFieldSmsCampaignID = big.NewInt(1 << 0)
	createActionScheduleSmsCampaignsRequestFieldScheduleTime  = big.NewInt(1 << 1)
)

type CreateActionScheduleSmsCampaignsRequest struct {
	// The unique id for the SMS campaign.
	SmsCampaignID string `json:"-" url:"-"`
	// The UTC date and time to schedule the campaign.
	ScheduleTime time.Time `json:"schedule_time" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CreateActionScheduleSmsCampaignsRequest) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetSmsCampaignID sets the SmsCampaignID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateActionScheduleSmsCampaignsRequest) SetSmsCampaignID(smsCampaignID string) {
	c.SmsCampaignID = smsCampaignID
	c.require(createActionScheduleSmsCampaignsRequestFieldSmsCampaignID)
}

// SetScheduleTime sets the ScheduleTime field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateActionScheduleSmsCampaignsRequest) SetScheduleTime(scheduleTime time.Time) {
	c.ScheduleTime = scheduleTime
	c.require(createActionScheduleSmsCampaignsRequestFieldScheduleTime)
}

func (c *CreateActionScheduleSmsCampaignsRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateActionScheduleSmsCampaignsRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*c = CreateActionScheduleSmsCampaignsRequest(body)
	return nil
}

func (c *CreateActionScheduleSmsCampaignsRequest) MarshalJSON() ([]byte, error) {
	type embed CreateActionScheduleSmsCampaignsRequest
	var marshaler = struct {
		embed
		ScheduleTime *internal.DateTime `json:"schedule_time"`
	}{
		embed:        embed(*c),
		ScheduleTime: internal.NewDateTime(c.ScheduleTime),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

var (
	createActionSendSmsCampaignsRequestFieldSmsCampaignID = big.NewInt(1 << 0)
)

type CreateActionSendSmsCampaignsRequest struct {
	// The unique id for the SMS campaign.
	SmsCampaignID string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CreateActionSendSmsCampaignsRequest) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetSmsCampaignID sets the SmsCampaignID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateActionSendSmsCampaignsRequest) SetSmsCampaignID(smsCampaignID string) {
	c.SmsCampaignID = smsCampaignID
	c.require(createActionSendSmsCampaignsRequestFieldSmsCampaignID)
}

var (
	deleteSmsCampaignsRequestFieldSmsCampaignID = big.NewInt(1 << 0)
)

type DeleteSmsCampaignsRequest struct {
	// The unique id for the SMS campaign.
	SmsCampaignID string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (d *DeleteSmsCampaignsRequest) require(field *big.Int) {
	if d.explicitFields == nil {
		d.explicitFields = big.NewInt(0)
	}
	d.explicitFields.Or(d.explicitFields, field)
}

// SetSmsCampaignID sets the SmsCampaignID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (d *DeleteSmsCampaignsRequest) SetSmsCampaignID(smsCampaignID string) {
	d.SmsCampaignID = smsCampaignID
	d.require(deleteSmsCampaignsRequestFieldSmsCampaignID)
}

var (
	getSmsCampaignsRequestFieldSmsCampaignID = big.NewInt(1 << 0)
	getSmsCampaignsRequestFieldFields        = big.NewInt(1 << 1)
	getSmsCampaignsRequestFieldExcludeFields = big.NewInt(1 << 2)
)

type GetSmsCampaignsRequest struct {
	// The unique id for the SMS campaign.
	SmsCampaignID string `json:"-" url:"-"`
	// A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
	Fields []*string `json:"-" url:"fields,omitempty"`
	// A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
	ExcludeFields []*string `json:"-" url:"exclude_fields,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetSmsCampaignsRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetSmsCampaignID sets the SmsCampaignID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSmsCampaignsRequest) SetSmsCampaignID(smsCampaignID string) {
	g.SmsCampaignID = smsCampaignID
	g.require(getSmsCampaignsRequestFieldSmsCampaignID)
}

// SetFields sets the Fields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSmsCampaignsRequest) SetFields(fields []*string) {
	g.Fields = fields
	g.require(getSmsCampaignsRequestFieldFields)
}

// SetExcludeFields sets the ExcludeFields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSmsCampaignsRequest) SetExcludeFields(excludeFields []*string) {
	g.ExcludeFields = excludeFields
	g.require(getSmsCampaignsRequestFieldExcludeFields)
}

var (
	getContentSmsCampaignsRequestFieldSmsCampaignID = big.NewInt(1 << 0)
	getContentSmsCampaignsRequestFieldFields        = big.NewInt(1 << 1)
	getContentSmsCampaignsRequestFieldExcludeFields = big.NewInt(1 << 2)
)

type GetContentSmsCampaignsRequest struct {
	// The unique id for the SMS campaign.
	SmsCampaignID string `json:"-" url:"-"`
	// A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
	Fields []*string `json:"-" url:"fields,omitempty"`
	// A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
	ExcludeFields []*string `json:"-" url:"exclude_fields,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetContentSmsCampaignsRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetSmsCampaignID sets the SmsCampaignID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetContentSmsCampaignsRequest) SetSmsCampaignID(smsCampaignID string) {
	g.SmsCampaignID = smsCampaignID
	g.require(getContentSmsCampaignsRequestFieldSmsCampaignID)
}

// SetFields sets the Fields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetContentSmsCampaignsRequest) SetFields(fields []*string) {
	g.Fields = fields
	g.require(getContentSmsCampaignsRequestFieldFields)
}

// SetExcludeFields sets the ExcludeFields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetContentSmsCampaignsRequest) SetExcludeFields(excludeFields []*string) {
	g.ExcludeFields = excludeFields
	g.require(getContentSmsCampaignsRequestFieldExcludeFields)
}

var (
	listSmsCampaignsRequestFieldFields        = big.NewInt(1 << 0)
	listSmsCampaignsRequestFieldExcludeFields = big.NewInt(1 << 1)
	listSmsCampaignsRequestFieldCount         = big.NewInt(1 << 2)
	listSmsCampaignsRequestFieldOffset        = big.NewInt(1 << 3)
)

type ListSmsCampaignsRequest struct {
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

func (l *ListSmsCampaignsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetFields sets the Fields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSmsCampaignsRequest) SetFields(fields []*string) {
	l.Fields = fields
	l.require(listSmsCampaignsRequestFieldFields)
}

// SetExcludeFields sets the ExcludeFields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSmsCampaignsRequest) SetExcludeFields(excludeFields []*string) {
	l.ExcludeFields = excludeFields
	l.require(listSmsCampaignsRequestFieldExcludeFields)
}

// SetCount sets the Count field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSmsCampaignsRequest) SetCount(count *int) {
	l.Count = count
	l.require(listSmsCampaignsRequestFieldCount)
}

// SetOffset sets the Offset field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSmsCampaignsRequest) SetOffset(offset *int) {
	l.Offset = offset
	l.require(listSmsCampaignsRequestFieldOffset)
}

// A single SMS campaign.
var (
	smsCampaignFieldID               = big.NewInt(1 << 0)
	smsCampaignFieldWebID            = big.NewInt(1 << 1)
	smsCampaignFieldName             = big.NewInt(1 << 2)
	smsCampaignFieldStatus           = big.NewInt(1 << 3)
	smsCampaignFieldChannel          = big.NewInt(1 << 4)
	smsCampaignFieldListID           = big.NewInt(1 << 5)
	smsCampaignFieldRecipientCount   = big.NewInt(1 << 6)
	smsCampaignFieldCreateTime       = big.NewInt(1 << 7)
	smsCampaignFieldSendTime         = big.NewInt(1 << 8)
	smsCampaignFieldUpdatedAt        = big.NewInt(1 << 9)
	smsCampaignFieldExpireTime       = big.NewInt(1 << 10)
	smsCampaignFieldIsSendNow        = big.NewInt(1 << 11)
	smsCampaignFieldFolderID         = big.NewInt(1 << 12)
	smsCampaignFieldSegments         = big.NewInt(1 << 13)
	smsCampaignFieldExcludedSegments = big.NewInt(1 << 14)
	smsCampaignFieldLinks            = big.NewInt(1 << 15)
)

type SmsCampaign struct {
	// A string that uniquely identifies this campaign.
	ID *string `json:"id,omitempty" url:"id,omitempty"`
	// The ID used in the Mailchimp web application.
	WebID *string `json:"web_id,omitempty" url:"web_id,omitempty"`
	// The name of the campaign.
	Name *string `json:"name,omitempty" url:"name,omitempty"`
	// The current status of the campaign.
	Status *string `json:"status,omitempty" url:"status,omitempty"`
	// The channel for this campaign (sms or whatsapp).
	Channel *string `json:"channel,omitempty" url:"channel,omitempty"`
	// The numeric ID of the list associated with this campaign.
	ListID *int `json:"list_id,omitempty" url:"list_id,omitempty"`
	// The number of recipients for this campaign.
	RecipientCount *int `json:"recipient_count,omitempty" url:"recipient_count,omitempty"`
	// The date and time the campaign was created.
	CreateTime *time.Time `json:"create_time,omitempty" url:"create_time,omitempty"`
	// The date and time the campaign is scheduled to send.
	SendTime *time.Time `json:"send_time,omitempty" url:"send_time,omitempty"`
	// The date and time the campaign was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty" url:"updated_at,omitempty"`
	// The date and time the campaign will stop sending in ISO 8601 format.
	ExpireTime *time.Time `json:"expire_time,omitempty" url:"expire_time,omitempty"`
	// Whether the campaign is configured to send immediately.
	IsSendNow *bool `json:"is_send_now,omitempty" url:"is_send_now,omitempty"`
	// The ID of the folder this campaign is in.
	FolderID *string `json:"folder_id,omitempty" url:"folder_id,omitempty"`
	// The segment IDs used to target recipients for this campaign.
	Segments []int `json:"segments,omitempty" url:"segments,omitempty"`
	// The segment IDs excluded from receiving this campaign.
	ExcludedSegments []int `json:"excluded_segments,omitempty" url:"excluded_segments,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []*SmsCampaignLinksItem `json:"_links,omitempty" url:"_links,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *SmsCampaign) GetID() *string {
	if s == nil {
		return nil
	}
	return s.ID
}

func (s *SmsCampaign) GetWebID() *string {
	if s == nil {
		return nil
	}
	return s.WebID
}

func (s *SmsCampaign) GetName() *string {
	if s == nil {
		return nil
	}
	return s.Name
}

func (s *SmsCampaign) GetStatus() *string {
	if s == nil {
		return nil
	}
	return s.Status
}

func (s *SmsCampaign) GetChannel() *string {
	if s == nil {
		return nil
	}
	return s.Channel
}

func (s *SmsCampaign) GetListID() *int {
	if s == nil {
		return nil
	}
	return s.ListID
}

func (s *SmsCampaign) GetRecipientCount() *int {
	if s == nil {
		return nil
	}
	return s.RecipientCount
}

func (s *SmsCampaign) GetCreateTime() *time.Time {
	if s == nil {
		return nil
	}
	return s.CreateTime
}

func (s *SmsCampaign) GetSendTime() *time.Time {
	if s == nil {
		return nil
	}
	return s.SendTime
}

func (s *SmsCampaign) GetUpdatedAt() *time.Time {
	if s == nil {
		return nil
	}
	return s.UpdatedAt
}

func (s *SmsCampaign) GetExpireTime() *time.Time {
	if s == nil {
		return nil
	}
	return s.ExpireTime
}

func (s *SmsCampaign) GetIsSendNow() *bool {
	if s == nil {
		return nil
	}
	return s.IsSendNow
}

func (s *SmsCampaign) GetFolderID() *string {
	if s == nil {
		return nil
	}
	return s.FolderID
}

func (s *SmsCampaign) GetSegments() []int {
	if s == nil {
		return nil
	}
	return s.Segments
}

func (s *SmsCampaign) GetExcludedSegments() []int {
	if s == nil {
		return nil
	}
	return s.ExcludedSegments
}

func (s *SmsCampaign) GetLinks() []*SmsCampaignLinksItem {
	if s == nil {
		return nil
	}
	return s.Links
}

func (s *SmsCampaign) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s *SmsCampaign) require(field *big.Int) {
	if s.explicitFields == nil {
		s.explicitFields = big.NewInt(0)
	}
	s.explicitFields.Or(s.explicitFields, field)
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaign) SetID(id *string) {
	s.ID = id
	s.require(smsCampaignFieldID)
}

// SetWebID sets the WebID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaign) SetWebID(webID *string) {
	s.WebID = webID
	s.require(smsCampaignFieldWebID)
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaign) SetName(name *string) {
	s.Name = name
	s.require(smsCampaignFieldName)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaign) SetStatus(status *string) {
	s.Status = status
	s.require(smsCampaignFieldStatus)
}

// SetChannel sets the Channel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaign) SetChannel(channel *string) {
	s.Channel = channel
	s.require(smsCampaignFieldChannel)
}

// SetListID sets the ListID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaign) SetListID(listID *int) {
	s.ListID = listID
	s.require(smsCampaignFieldListID)
}

// SetRecipientCount sets the RecipientCount field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaign) SetRecipientCount(recipientCount *int) {
	s.RecipientCount = recipientCount
	s.require(smsCampaignFieldRecipientCount)
}

// SetCreateTime sets the CreateTime field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaign) SetCreateTime(createTime *time.Time) {
	s.CreateTime = createTime
	s.require(smsCampaignFieldCreateTime)
}

// SetSendTime sets the SendTime field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaign) SetSendTime(sendTime *time.Time) {
	s.SendTime = sendTime
	s.require(smsCampaignFieldSendTime)
}

// SetUpdatedAt sets the UpdatedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaign) SetUpdatedAt(updatedAt *time.Time) {
	s.UpdatedAt = updatedAt
	s.require(smsCampaignFieldUpdatedAt)
}

// SetExpireTime sets the ExpireTime field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaign) SetExpireTime(expireTime *time.Time) {
	s.ExpireTime = expireTime
	s.require(smsCampaignFieldExpireTime)
}

// SetIsSendNow sets the IsSendNow field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaign) SetIsSendNow(isSendNow *bool) {
	s.IsSendNow = isSendNow
	s.require(smsCampaignFieldIsSendNow)
}

// SetFolderID sets the FolderID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaign) SetFolderID(folderID *string) {
	s.FolderID = folderID
	s.require(smsCampaignFieldFolderID)
}

// SetSegments sets the Segments field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaign) SetSegments(segments []int) {
	s.Segments = segments
	s.require(smsCampaignFieldSegments)
}

// SetExcludedSegments sets the ExcludedSegments field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaign) SetExcludedSegments(excludedSegments []int) {
	s.ExcludedSegments = excludedSegments
	s.require(smsCampaignFieldExcludedSegments)
}

// SetLinks sets the Links field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaign) SetLinks(links []*SmsCampaignLinksItem) {
	s.Links = links
	s.require(smsCampaignFieldLinks)
}

func (s *SmsCampaign) UnmarshalJSON(data []byte) error {
	type embed SmsCampaign
	var unmarshaler = struct {
		embed
		CreateTime *internal.DateTime `json:"create_time,omitempty"`
		SendTime   *internal.DateTime `json:"send_time,omitempty"`
		UpdatedAt  *internal.DateTime `json:"updated_at,omitempty"`
		ExpireTime *internal.DateTime `json:"expire_time,omitempty"`
	}{
		embed: embed(*s),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*s = SmsCampaign(unmarshaler.embed)
	s.CreateTime = unmarshaler.CreateTime.TimePtr()
	s.SendTime = unmarshaler.SendTime.TimePtr()
	s.UpdatedAt = unmarshaler.UpdatedAt.TimePtr()
	s.ExpireTime = unmarshaler.ExpireTime.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *SmsCampaign) MarshalJSON() ([]byte, error) {
	type embed SmsCampaign
	var marshaler = struct {
		embed
		CreateTime *internal.DateTime `json:"create_time,omitempty"`
		SendTime   *internal.DateTime `json:"send_time,omitempty"`
		UpdatedAt  *internal.DateTime `json:"updated_at,omitempty"`
		ExpireTime *internal.DateTime `json:"expire_time,omitempty"`
	}{
		embed:      embed(*s),
		CreateTime: internal.NewOptionalDateTime(s.CreateTime),
		SendTime:   internal.NewOptionalDateTime(s.SendTime),
		UpdatedAt:  internal.NewOptionalDateTime(s.UpdatedAt),
		ExpireTime: internal.NewOptionalDateTime(s.ExpireTime),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, s.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (s *SmsCampaign) String() string {
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

// The content of an SMS campaign.
var (
	smsCampaignContentFieldMessageBody       = big.NewInt(1 << 0)
	smsCampaignContentFieldEstimatedSegments = big.NewInt(1 << 1)
	smsCampaignContentFieldMergeFields       = big.NewInt(1 << 2)
	smsCampaignContentFieldMedia             = big.NewInt(1 << 3)
	smsCampaignContentFieldSource            = big.NewInt(1 << 4)
	smsCampaignContentFieldProperties        = big.NewInt(1 << 5)
	smsCampaignContentFieldLinks             = big.NewInt(1 << 6)
)

type SmsCampaignContent struct {
	// The SMS message body.
	MessageBody *string `json:"message_body,omitempty" url:"message_body,omitempty"`
	// The estimated number of message segments this content will use.
	EstimatedSegments *int `json:"estimated_segments,omitempty" url:"estimated_segments,omitempty"`
	// The merge fields used in the message body.
	MergeFields []string `json:"merge_fields,omitempty" url:"merge_fields,omitempty"`
	// Attached images or files.
	Media []*SmsCampaignContentMediaItem `json:"media,omitempty" url:"media,omitempty"`
	// The source that created or imported this content.
	Source *SmsCampaignContentSource `json:"source,omitempty" url:"source,omitempty"`
	// Additional content properties.
	Properties *SmsCampaignContentProperties `json:"properties,omitempty" url:"properties,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []*SmsCampaignContentLinksItem `json:"_links,omitempty" url:"_links,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *SmsCampaignContent) GetMessageBody() *string {
	if s == nil {
		return nil
	}
	return s.MessageBody
}

func (s *SmsCampaignContent) GetEstimatedSegments() *int {
	if s == nil {
		return nil
	}
	return s.EstimatedSegments
}

func (s *SmsCampaignContent) GetMergeFields() []string {
	if s == nil {
		return nil
	}
	return s.MergeFields
}

func (s *SmsCampaignContent) GetMedia() []*SmsCampaignContentMediaItem {
	if s == nil {
		return nil
	}
	return s.Media
}

func (s *SmsCampaignContent) GetSource() *SmsCampaignContentSource {
	if s == nil {
		return nil
	}
	return s.Source
}

func (s *SmsCampaignContent) GetProperties() *SmsCampaignContentProperties {
	if s == nil {
		return nil
	}
	return s.Properties
}

func (s *SmsCampaignContent) GetLinks() []*SmsCampaignContentLinksItem {
	if s == nil {
		return nil
	}
	return s.Links
}

func (s *SmsCampaignContent) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s *SmsCampaignContent) require(field *big.Int) {
	if s.explicitFields == nil {
		s.explicitFields = big.NewInt(0)
	}
	s.explicitFields.Or(s.explicitFields, field)
}

// SetMessageBody sets the MessageBody field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaignContent) SetMessageBody(messageBody *string) {
	s.MessageBody = messageBody
	s.require(smsCampaignContentFieldMessageBody)
}

// SetEstimatedSegments sets the EstimatedSegments field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaignContent) SetEstimatedSegments(estimatedSegments *int) {
	s.EstimatedSegments = estimatedSegments
	s.require(smsCampaignContentFieldEstimatedSegments)
}

// SetMergeFields sets the MergeFields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaignContent) SetMergeFields(mergeFields []string) {
	s.MergeFields = mergeFields
	s.require(smsCampaignContentFieldMergeFields)
}

// SetMedia sets the Media field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaignContent) SetMedia(media []*SmsCampaignContentMediaItem) {
	s.Media = media
	s.require(smsCampaignContentFieldMedia)
}

// SetSource sets the Source field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaignContent) SetSource(source *SmsCampaignContentSource) {
	s.Source = source
	s.require(smsCampaignContentFieldSource)
}

// SetProperties sets the Properties field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaignContent) SetProperties(properties *SmsCampaignContentProperties) {
	s.Properties = properties
	s.require(smsCampaignContentFieldProperties)
}

// SetLinks sets the Links field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaignContent) SetLinks(links []*SmsCampaignContentLinksItem) {
	s.Links = links
	s.require(smsCampaignContentFieldLinks)
}

func (s *SmsCampaignContent) UnmarshalJSON(data []byte) error {
	type unmarshaler SmsCampaignContent
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SmsCampaignContent(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *SmsCampaignContent) MarshalJSON() ([]byte, error) {
	type embed SmsCampaignContent
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, s.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (s *SmsCampaignContent) String() string {
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

// This object represents a link from the resource where it is found to another resource or action that may be performed.
var (
	smsCampaignContentLinksItemFieldHref         = big.NewInt(1 << 0)
	smsCampaignContentLinksItemFieldMethod       = big.NewInt(1 << 1)
	smsCampaignContentLinksItemFieldRel          = big.NewInt(1 << 2)
	smsCampaignContentLinksItemFieldSchema       = big.NewInt(1 << 3)
	smsCampaignContentLinksItemFieldTargetSchema = big.NewInt(1 << 4)
)

type SmsCampaignContentLinksItem struct {
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href *string `json:"href,omitempty" url:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method *SmsCampaignContentLinksItemMethod `json:"method,omitempty" url:"method,omitempty"`
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

func (s *SmsCampaignContentLinksItem) GetHref() *string {
	if s == nil {
		return nil
	}
	return s.Href
}

func (s *SmsCampaignContentLinksItem) GetMethod() *SmsCampaignContentLinksItemMethod {
	if s == nil {
		return nil
	}
	return s.Method
}

func (s *SmsCampaignContentLinksItem) GetRel() *string {
	if s == nil {
		return nil
	}
	return s.Rel
}

func (s *SmsCampaignContentLinksItem) GetSchema() *string {
	if s == nil {
		return nil
	}
	return s.Schema
}

func (s *SmsCampaignContentLinksItem) GetTargetSchema() *string {
	if s == nil {
		return nil
	}
	return s.TargetSchema
}

func (s *SmsCampaignContentLinksItem) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s *SmsCampaignContentLinksItem) require(field *big.Int) {
	if s.explicitFields == nil {
		s.explicitFields = big.NewInt(0)
	}
	s.explicitFields.Or(s.explicitFields, field)
}

// SetHref sets the Href field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaignContentLinksItem) SetHref(href *string) {
	s.Href = href
	s.require(smsCampaignContentLinksItemFieldHref)
}

// SetMethod sets the Method field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaignContentLinksItem) SetMethod(method *SmsCampaignContentLinksItemMethod) {
	s.Method = method
	s.require(smsCampaignContentLinksItemFieldMethod)
}

// SetRel sets the Rel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaignContentLinksItem) SetRel(rel *string) {
	s.Rel = rel
	s.require(smsCampaignContentLinksItemFieldRel)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaignContentLinksItem) SetSchema(schema *string) {
	s.Schema = schema
	s.require(smsCampaignContentLinksItemFieldSchema)
}

// SetTargetSchema sets the TargetSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaignContentLinksItem) SetTargetSchema(targetSchema *string) {
	s.TargetSchema = targetSchema
	s.require(smsCampaignContentLinksItemFieldTargetSchema)
}

func (s *SmsCampaignContentLinksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler SmsCampaignContentLinksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SmsCampaignContentLinksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *SmsCampaignContentLinksItem) MarshalJSON() ([]byte, error) {
	type embed SmsCampaignContentLinksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, s.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (s *SmsCampaignContentLinksItem) String() string {
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

// The HTTP method that should be used when accessing the URL defined in 'href'.
type SmsCampaignContentLinksItemMethod string

const (
	SmsCampaignContentLinksItemMethodGet     SmsCampaignContentLinksItemMethod = "GET"
	SmsCampaignContentLinksItemMethodPost    SmsCampaignContentLinksItemMethod = "POST"
	SmsCampaignContentLinksItemMethodPut     SmsCampaignContentLinksItemMethod = "PUT"
	SmsCampaignContentLinksItemMethodPatch   SmsCampaignContentLinksItemMethod = "PATCH"
	SmsCampaignContentLinksItemMethodDelete  SmsCampaignContentLinksItemMethod = "DELETE"
	SmsCampaignContentLinksItemMethodOptions SmsCampaignContentLinksItemMethod = "OPTIONS"
	SmsCampaignContentLinksItemMethodHead    SmsCampaignContentLinksItemMethod = "HEAD"
)

func NewSmsCampaignContentLinksItemMethodFromString(s string) (SmsCampaignContentLinksItemMethod, error) {
	switch s {
	case "GET":
		return SmsCampaignContentLinksItemMethodGet, nil
	case "POST":
		return SmsCampaignContentLinksItemMethodPost, nil
	case "PUT":
		return SmsCampaignContentLinksItemMethodPut, nil
	case "PATCH":
		return SmsCampaignContentLinksItemMethodPatch, nil
	case "DELETE":
		return SmsCampaignContentLinksItemMethodDelete, nil
	case "OPTIONS":
		return SmsCampaignContentLinksItemMethodOptions, nil
	case "HEAD":
		return SmsCampaignContentLinksItemMethodHead, nil
	}
	var t SmsCampaignContentLinksItemMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (s SmsCampaignContentLinksItemMethod) Ptr() *SmsCampaignContentLinksItemMethod {
	return &s
}

var (
	smsCampaignContentMediaItemFieldURL = big.NewInt(1 << 0)
)

type SmsCampaignContentMediaItem struct {
	// The URL of the media file.
	URL *string `json:"url,omitempty" url:"url,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *SmsCampaignContentMediaItem) GetURL() *string {
	if s == nil {
		return nil
	}
	return s.URL
}

func (s *SmsCampaignContentMediaItem) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s *SmsCampaignContentMediaItem) require(field *big.Int) {
	if s.explicitFields == nil {
		s.explicitFields = big.NewInt(0)
	}
	s.explicitFields.Or(s.explicitFields, field)
}

// SetURL sets the URL field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaignContentMediaItem) SetURL(url *string) {
	s.URL = url
	s.require(smsCampaignContentMediaItemFieldURL)
}

func (s *SmsCampaignContentMediaItem) UnmarshalJSON(data []byte) error {
	type unmarshaler SmsCampaignContentMediaItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SmsCampaignContentMediaItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *SmsCampaignContentMediaItem) MarshalJSON() ([]byte, error) {
	type embed SmsCampaignContentMediaItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, s.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (s *SmsCampaignContentMediaItem) String() string {
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

// Additional content properties.
var (
	smsCampaignContentPropertiesFieldContentType           = big.NewInt(1 << 0)
	smsCampaignContentPropertiesFieldSender                = big.NewInt(1 << 1)
	smsCampaignContentPropertiesFieldOptoutMessageLanguage = big.NewInt(1 << 2)
)

type SmsCampaignContentProperties struct {
	// The content type of the message.
	ContentType *string `json:"content_type,omitempty" url:"content_type,omitempty"`
	// The sender identifier for the message.
	Sender *string `json:"sender,omitempty" url:"sender,omitempty"`
	// The language of the opt-out message.
	OptoutMessageLanguage *string `json:"optout_message_language,omitempty" url:"optout_message_language,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *SmsCampaignContentProperties) GetContentType() *string {
	if s == nil {
		return nil
	}
	return s.ContentType
}

func (s *SmsCampaignContentProperties) GetSender() *string {
	if s == nil {
		return nil
	}
	return s.Sender
}

func (s *SmsCampaignContentProperties) GetOptoutMessageLanguage() *string {
	if s == nil {
		return nil
	}
	return s.OptoutMessageLanguage
}

func (s *SmsCampaignContentProperties) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s *SmsCampaignContentProperties) require(field *big.Int) {
	if s.explicitFields == nil {
		s.explicitFields = big.NewInt(0)
	}
	s.explicitFields.Or(s.explicitFields, field)
}

// SetContentType sets the ContentType field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaignContentProperties) SetContentType(contentType *string) {
	s.ContentType = contentType
	s.require(smsCampaignContentPropertiesFieldContentType)
}

// SetSender sets the Sender field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaignContentProperties) SetSender(sender *string) {
	s.Sender = sender
	s.require(smsCampaignContentPropertiesFieldSender)
}

// SetOptoutMessageLanguage sets the OptoutMessageLanguage field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaignContentProperties) SetOptoutMessageLanguage(optoutMessageLanguage *string) {
	s.OptoutMessageLanguage = optoutMessageLanguage
	s.require(smsCampaignContentPropertiesFieldOptoutMessageLanguage)
}

func (s *SmsCampaignContentProperties) UnmarshalJSON(data []byte) error {
	type unmarshaler SmsCampaignContentProperties
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SmsCampaignContentProperties(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *SmsCampaignContentProperties) MarshalJSON() ([]byte, error) {
	type embed SmsCampaignContentProperties
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, s.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (s *SmsCampaignContentProperties) String() string {
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

// The source that created or imported this content.
var (
	smsCampaignContentSourceFieldType = big.NewInt(1 << 0)
	smsCampaignContentSourceFieldID   = big.NewInt(1 << 1)
)

type SmsCampaignContentSource struct {
	// The type of source.
	Type *string `json:"type,omitempty" url:"type,omitempty"`
	// The ID of the source.
	ID *string `json:"id,omitempty" url:"id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *SmsCampaignContentSource) GetType() *string {
	if s == nil {
		return nil
	}
	return s.Type
}

func (s *SmsCampaignContentSource) GetID() *string {
	if s == nil {
		return nil
	}
	return s.ID
}

func (s *SmsCampaignContentSource) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s *SmsCampaignContentSource) require(field *big.Int) {
	if s.explicitFields == nil {
		s.explicitFields = big.NewInt(0)
	}
	s.explicitFields.Or(s.explicitFields, field)
}

// SetType sets the Type field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaignContentSource) SetType(type_ *string) {
	s.Type = type_
	s.require(smsCampaignContentSourceFieldType)
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaignContentSource) SetID(id *string) {
	s.ID = id
	s.require(smsCampaignContentSourceFieldID)
}

func (s *SmsCampaignContentSource) UnmarshalJSON(data []byte) error {
	type unmarshaler SmsCampaignContentSource
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SmsCampaignContentSource(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *SmsCampaignContentSource) MarshalJSON() ([]byte, error) {
	type embed SmsCampaignContentSource
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, s.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (s *SmsCampaignContentSource) String() string {
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

// This object represents a link from the resource where it is found to another resource or action that may be performed.
var (
	smsCampaignLinksItemFieldHref         = big.NewInt(1 << 0)
	smsCampaignLinksItemFieldMethod       = big.NewInt(1 << 1)
	smsCampaignLinksItemFieldRel          = big.NewInt(1 << 2)
	smsCampaignLinksItemFieldSchema       = big.NewInt(1 << 3)
	smsCampaignLinksItemFieldTargetSchema = big.NewInt(1 << 4)
)

type SmsCampaignLinksItem struct {
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href *string `json:"href,omitempty" url:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method *SmsCampaignLinksItemMethod `json:"method,omitempty" url:"method,omitempty"`
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

func (s *SmsCampaignLinksItem) GetHref() *string {
	if s == nil {
		return nil
	}
	return s.Href
}

func (s *SmsCampaignLinksItem) GetMethod() *SmsCampaignLinksItemMethod {
	if s == nil {
		return nil
	}
	return s.Method
}

func (s *SmsCampaignLinksItem) GetRel() *string {
	if s == nil {
		return nil
	}
	return s.Rel
}

func (s *SmsCampaignLinksItem) GetSchema() *string {
	if s == nil {
		return nil
	}
	return s.Schema
}

func (s *SmsCampaignLinksItem) GetTargetSchema() *string {
	if s == nil {
		return nil
	}
	return s.TargetSchema
}

func (s *SmsCampaignLinksItem) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s *SmsCampaignLinksItem) require(field *big.Int) {
	if s.explicitFields == nil {
		s.explicitFields = big.NewInt(0)
	}
	s.explicitFields.Or(s.explicitFields, field)
}

// SetHref sets the Href field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaignLinksItem) SetHref(href *string) {
	s.Href = href
	s.require(smsCampaignLinksItemFieldHref)
}

// SetMethod sets the Method field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaignLinksItem) SetMethod(method *SmsCampaignLinksItemMethod) {
	s.Method = method
	s.require(smsCampaignLinksItemFieldMethod)
}

// SetRel sets the Rel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaignLinksItem) SetRel(rel *string) {
	s.Rel = rel
	s.require(smsCampaignLinksItemFieldRel)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaignLinksItem) SetSchema(schema *string) {
	s.Schema = schema
	s.require(smsCampaignLinksItemFieldSchema)
}

// SetTargetSchema sets the TargetSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SmsCampaignLinksItem) SetTargetSchema(targetSchema *string) {
	s.TargetSchema = targetSchema
	s.require(smsCampaignLinksItemFieldTargetSchema)
}

func (s *SmsCampaignLinksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler SmsCampaignLinksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SmsCampaignLinksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *SmsCampaignLinksItem) MarshalJSON() ([]byte, error) {
	type embed SmsCampaignLinksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, s.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (s *SmsCampaignLinksItem) String() string {
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

// The HTTP method that should be used when accessing the URL defined in 'href'.
type SmsCampaignLinksItemMethod string

const (
	SmsCampaignLinksItemMethodGet     SmsCampaignLinksItemMethod = "GET"
	SmsCampaignLinksItemMethodPost    SmsCampaignLinksItemMethod = "POST"
	SmsCampaignLinksItemMethodPut     SmsCampaignLinksItemMethod = "PUT"
	SmsCampaignLinksItemMethodPatch   SmsCampaignLinksItemMethod = "PATCH"
	SmsCampaignLinksItemMethodDelete  SmsCampaignLinksItemMethod = "DELETE"
	SmsCampaignLinksItemMethodOptions SmsCampaignLinksItemMethod = "OPTIONS"
	SmsCampaignLinksItemMethodHead    SmsCampaignLinksItemMethod = "HEAD"
)

func NewSmsCampaignLinksItemMethodFromString(s string) (SmsCampaignLinksItemMethod, error) {
	switch s {
	case "GET":
		return SmsCampaignLinksItemMethodGet, nil
	case "POST":
		return SmsCampaignLinksItemMethodPost, nil
	case "PUT":
		return SmsCampaignLinksItemMethodPut, nil
	case "PATCH":
		return SmsCampaignLinksItemMethodPatch, nil
	case "DELETE":
		return SmsCampaignLinksItemMethodDelete, nil
	case "OPTIONS":
		return SmsCampaignLinksItemMethodOptions, nil
	case "HEAD":
		return SmsCampaignLinksItemMethodHead, nil
	}
	var t SmsCampaignLinksItemMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (s SmsCampaignLinksItemMethod) Ptr() *SmsCampaignLinksItemMethod {
	return &s
}

// A collection of SMS campaigns.
var (
	listSmsCampaignsResponseFieldSmsCampaigns = big.NewInt(1 << 0)
	listSmsCampaignsResponseFieldTotalItems   = big.NewInt(1 << 1)
	listSmsCampaignsResponseFieldLinks        = big.NewInt(1 << 2)
)

type ListSmsCampaignsResponse struct {
	// An array of SMS campaigns.
	SmsCampaigns []*SmsCampaign `json:"sms_campaigns,omitempty" url:"sms_campaigns,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int `json:"total_items,omitempty" url:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []*ListSmsCampaignsResponseLinksItem `json:"_links,omitempty" url:"_links,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListSmsCampaignsResponse) GetSmsCampaigns() []*SmsCampaign {
	if l == nil {
		return nil
	}
	return l.SmsCampaigns
}

func (l *ListSmsCampaignsResponse) GetTotalItems() *int {
	if l == nil {
		return nil
	}
	return l.TotalItems
}

func (l *ListSmsCampaignsResponse) GetLinks() []*ListSmsCampaignsResponseLinksItem {
	if l == nil {
		return nil
	}
	return l.Links
}

func (l *ListSmsCampaignsResponse) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListSmsCampaignsResponse) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetSmsCampaigns sets the SmsCampaigns field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSmsCampaignsResponse) SetSmsCampaigns(smsCampaigns []*SmsCampaign) {
	l.SmsCampaigns = smsCampaigns
	l.require(listSmsCampaignsResponseFieldSmsCampaigns)
}

// SetTotalItems sets the TotalItems field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSmsCampaignsResponse) SetTotalItems(totalItems *int) {
	l.TotalItems = totalItems
	l.require(listSmsCampaignsResponseFieldTotalItems)
}

// SetLinks sets the Links field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSmsCampaignsResponse) SetLinks(links []*ListSmsCampaignsResponseLinksItem) {
	l.Links = links
	l.require(listSmsCampaignsResponseFieldLinks)
}

func (l *ListSmsCampaignsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListSmsCampaignsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListSmsCampaignsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListSmsCampaignsResponse) MarshalJSON() ([]byte, error) {
	type embed ListSmsCampaignsResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListSmsCampaignsResponse) String() string {
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
	listSmsCampaignsResponseLinksItemFieldHref         = big.NewInt(1 << 0)
	listSmsCampaignsResponseLinksItemFieldMethod       = big.NewInt(1 << 1)
	listSmsCampaignsResponseLinksItemFieldRel          = big.NewInt(1 << 2)
	listSmsCampaignsResponseLinksItemFieldSchema       = big.NewInt(1 << 3)
	listSmsCampaignsResponseLinksItemFieldTargetSchema = big.NewInt(1 << 4)
)

type ListSmsCampaignsResponseLinksItem struct {
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href *string `json:"href,omitempty" url:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method *ListSmsCampaignsResponseLinksItemMethod `json:"method,omitempty" url:"method,omitempty"`
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

func (l *ListSmsCampaignsResponseLinksItem) GetHref() *string {
	if l == nil {
		return nil
	}
	return l.Href
}

func (l *ListSmsCampaignsResponseLinksItem) GetMethod() *ListSmsCampaignsResponseLinksItemMethod {
	if l == nil {
		return nil
	}
	return l.Method
}

func (l *ListSmsCampaignsResponseLinksItem) GetRel() *string {
	if l == nil {
		return nil
	}
	return l.Rel
}

func (l *ListSmsCampaignsResponseLinksItem) GetSchema() *string {
	if l == nil {
		return nil
	}
	return l.Schema
}

func (l *ListSmsCampaignsResponseLinksItem) GetTargetSchema() *string {
	if l == nil {
		return nil
	}
	return l.TargetSchema
}

func (l *ListSmsCampaignsResponseLinksItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListSmsCampaignsResponseLinksItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetHref sets the Href field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSmsCampaignsResponseLinksItem) SetHref(href *string) {
	l.Href = href
	l.require(listSmsCampaignsResponseLinksItemFieldHref)
}

// SetMethod sets the Method field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSmsCampaignsResponseLinksItem) SetMethod(method *ListSmsCampaignsResponseLinksItemMethod) {
	l.Method = method
	l.require(listSmsCampaignsResponseLinksItemFieldMethod)
}

// SetRel sets the Rel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSmsCampaignsResponseLinksItem) SetRel(rel *string) {
	l.Rel = rel
	l.require(listSmsCampaignsResponseLinksItemFieldRel)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSmsCampaignsResponseLinksItem) SetSchema(schema *string) {
	l.Schema = schema
	l.require(listSmsCampaignsResponseLinksItemFieldSchema)
}

// SetTargetSchema sets the TargetSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSmsCampaignsResponseLinksItem) SetTargetSchema(targetSchema *string) {
	l.TargetSchema = targetSchema
	l.require(listSmsCampaignsResponseLinksItemFieldTargetSchema)
}

func (l *ListSmsCampaignsResponseLinksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler ListSmsCampaignsResponseLinksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListSmsCampaignsResponseLinksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListSmsCampaignsResponseLinksItem) MarshalJSON() ([]byte, error) {
	type embed ListSmsCampaignsResponseLinksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListSmsCampaignsResponseLinksItem) String() string {
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
type ListSmsCampaignsResponseLinksItemMethod string

const (
	ListSmsCampaignsResponseLinksItemMethodGet     ListSmsCampaignsResponseLinksItemMethod = "GET"
	ListSmsCampaignsResponseLinksItemMethodPost    ListSmsCampaignsResponseLinksItemMethod = "POST"
	ListSmsCampaignsResponseLinksItemMethodPut     ListSmsCampaignsResponseLinksItemMethod = "PUT"
	ListSmsCampaignsResponseLinksItemMethodPatch   ListSmsCampaignsResponseLinksItemMethod = "PATCH"
	ListSmsCampaignsResponseLinksItemMethodDelete  ListSmsCampaignsResponseLinksItemMethod = "DELETE"
	ListSmsCampaignsResponseLinksItemMethodOptions ListSmsCampaignsResponseLinksItemMethod = "OPTIONS"
	ListSmsCampaignsResponseLinksItemMethodHead    ListSmsCampaignsResponseLinksItemMethod = "HEAD"
)

func NewListSmsCampaignsResponseLinksItemMethodFromString(s string) (ListSmsCampaignsResponseLinksItemMethod, error) {
	switch s {
	case "GET":
		return ListSmsCampaignsResponseLinksItemMethodGet, nil
	case "POST":
		return ListSmsCampaignsResponseLinksItemMethodPost, nil
	case "PUT":
		return ListSmsCampaignsResponseLinksItemMethodPut, nil
	case "PATCH":
		return ListSmsCampaignsResponseLinksItemMethodPatch, nil
	case "DELETE":
		return ListSmsCampaignsResponseLinksItemMethodDelete, nil
	case "OPTIONS":
		return ListSmsCampaignsResponseLinksItemMethodOptions, nil
	case "HEAD":
		return ListSmsCampaignsResponseLinksItemMethodHead, nil
	}
	var t ListSmsCampaignsResponseLinksItemMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListSmsCampaignsResponseLinksItemMethod) Ptr() *ListSmsCampaignsResponseLinksItemMethod {
	return &l
}

var (
	upsertContentSmsCampaignsRequestMediaItemFieldURL = big.NewInt(1 << 0)
)

type UpsertContentSmsCampaignsRequestMediaItem struct {
	// The URL of the media file.
	URL *string `json:"url,omitempty" url:"url,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (u *UpsertContentSmsCampaignsRequestMediaItem) GetURL() *string {
	if u == nil {
		return nil
	}
	return u.URL
}

func (u *UpsertContentSmsCampaignsRequestMediaItem) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u *UpsertContentSmsCampaignsRequestMediaItem) require(field *big.Int) {
	if u.explicitFields == nil {
		u.explicitFields = big.NewInt(0)
	}
	u.explicitFields.Or(u.explicitFields, field)
}

// SetURL sets the URL field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (u *UpsertContentSmsCampaignsRequestMediaItem) SetURL(url *string) {
	u.URL = url
	u.require(upsertContentSmsCampaignsRequestMediaItemFieldURL)
}

func (u *UpsertContentSmsCampaignsRequestMediaItem) UnmarshalJSON(data []byte) error {
	type unmarshaler UpsertContentSmsCampaignsRequestMediaItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpsertContentSmsCampaignsRequestMediaItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	u.rawJSON = json.RawMessage(data)
	return nil
}

func (u *UpsertContentSmsCampaignsRequestMediaItem) MarshalJSON() ([]byte, error) {
	type embed UpsertContentSmsCampaignsRequestMediaItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, u.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (u *UpsertContentSmsCampaignsRequestMediaItem) String() string {
	if u == nil {
		return "<nil>"
	}
	if len(u.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(u.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

var (
	updateSmsCampaignsRequestFieldSmsCampaignID    = big.NewInt(1 << 0)
	updateSmsCampaignsRequestFieldName             = big.NewInt(1 << 1)
	updateSmsCampaignsRequestFieldFolderID         = big.NewInt(1 << 2)
	updateSmsCampaignsRequestFieldSegments         = big.NewInt(1 << 3)
	updateSmsCampaignsRequestFieldExcludedSegments = big.NewInt(1 << 4)
)

type UpdateSmsCampaignsRequest struct {
	// The unique id for the SMS campaign.
	SmsCampaignID string `json:"-" url:"-"`
	// The name of the campaign.
	Name *string `json:"name,omitempty" url:"-"`
	// The ID of the folder to place this campaign in.
	FolderID *string `json:"folder_id,omitempty" url:"-"`
	// The segment IDs to target for this campaign.
	Segments []int `json:"segments,omitempty" url:"-"`
	// The segment IDs to exclude from this campaign.
	ExcludedSegments []int `json:"excluded_segments,omitempty" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (u *UpdateSmsCampaignsRequest) require(field *big.Int) {
	if u.explicitFields == nil {
		u.explicitFields = big.NewInt(0)
	}
	u.explicitFields.Or(u.explicitFields, field)
}

// SetSmsCampaignID sets the SmsCampaignID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (u *UpdateSmsCampaignsRequest) SetSmsCampaignID(smsCampaignID string) {
	u.SmsCampaignID = smsCampaignID
	u.require(updateSmsCampaignsRequestFieldSmsCampaignID)
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (u *UpdateSmsCampaignsRequest) SetName(name *string) {
	u.Name = name
	u.require(updateSmsCampaignsRequestFieldName)
}

// SetFolderID sets the FolderID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (u *UpdateSmsCampaignsRequest) SetFolderID(folderID *string) {
	u.FolderID = folderID
	u.require(updateSmsCampaignsRequestFieldFolderID)
}

// SetSegments sets the Segments field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (u *UpdateSmsCampaignsRequest) SetSegments(segments []int) {
	u.Segments = segments
	u.require(updateSmsCampaignsRequestFieldSegments)
}

// SetExcludedSegments sets the ExcludedSegments field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (u *UpdateSmsCampaignsRequest) SetExcludedSegments(excludedSegments []int) {
	u.ExcludedSegments = excludedSegments
	u.require(updateSmsCampaignsRequestFieldExcludedSegments)
}

func (u *UpdateSmsCampaignsRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateSmsCampaignsRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*u = UpdateSmsCampaignsRequest(body)
	return nil
}

func (u *UpdateSmsCampaignsRequest) MarshalJSON() ([]byte, error) {
	type embed UpdateSmsCampaignsRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, u.explicitFields)
	return json.Marshal(explicitMarshaler)
}

var (
	upsertContentSmsCampaignsRequestFieldSmsCampaignID = big.NewInt(1 << 0)
	upsertContentSmsCampaignsRequestFieldMessageBody   = big.NewInt(1 << 1)
	upsertContentSmsCampaignsRequestFieldMedia         = big.NewInt(1 << 2)
)

type UpsertContentSmsCampaignsRequest struct {
	// The unique id for the SMS campaign.
	SmsCampaignID string `json:"-" url:"-"`
	// The SMS message body.
	MessageBody string `json:"message_body" url:"-"`
	// Attached images or files. Limited to one item. Omitting this field or sending an empty array removes any existing media; to keep the current media while updating other fields, re-send the media array.
	Media []*UpsertContentSmsCampaignsRequestMediaItem `json:"media,omitempty" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (u *UpsertContentSmsCampaignsRequest) require(field *big.Int) {
	if u.explicitFields == nil {
		u.explicitFields = big.NewInt(0)
	}
	u.explicitFields.Or(u.explicitFields, field)
}

// SetSmsCampaignID sets the SmsCampaignID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (u *UpsertContentSmsCampaignsRequest) SetSmsCampaignID(smsCampaignID string) {
	u.SmsCampaignID = smsCampaignID
	u.require(upsertContentSmsCampaignsRequestFieldSmsCampaignID)
}

// SetMessageBody sets the MessageBody field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (u *UpsertContentSmsCampaignsRequest) SetMessageBody(messageBody string) {
	u.MessageBody = messageBody
	u.require(upsertContentSmsCampaignsRequestFieldMessageBody)
}

// SetMedia sets the Media field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (u *UpsertContentSmsCampaignsRequest) SetMedia(media []*UpsertContentSmsCampaignsRequestMediaItem) {
	u.Media = media
	u.require(upsertContentSmsCampaignsRequestFieldMedia)
}

func (u *UpsertContentSmsCampaignsRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler UpsertContentSmsCampaignsRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*u = UpsertContentSmsCampaignsRequest(body)
	return nil
}

func (u *UpsertContentSmsCampaignsRequest) MarshalJSON() ([]byte, error) {
	type embed UpsertContentSmsCampaignsRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, u.explicitFields)
	return json.Marshal(explicitMarshaler)
}
