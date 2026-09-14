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
	listChimpChatterActivityFeedRequestFieldCount  = big.NewInt(1 << 0)
	listChimpChatterActivityFeedRequestFieldOffset = big.NewInt(1 << 1)
)

type ListChimpChatterActivityFeedRequest struct {
	// The number of records to return. Default value is 10. Maximum value is 1000
	Count *int `json:"-" url:"count,omitempty"`
	// Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this is the number of records from a collection to skip. Default value is 0.
	Offset *int `json:"-" url:"offset,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListChimpChatterActivityFeedRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetCount sets the Count field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListChimpChatterActivityFeedRequest) SetCount(count *int) {
	l.Count = count
	l.require(listChimpChatterActivityFeedRequestFieldCount)
}

// SetOffset sets the Offset field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListChimpChatterActivityFeedRequest) SetOffset(offset *int) {
	l.Offset = offset
	l.require(listChimpChatterActivityFeedRequestFieldOffset)
}

// This object represents a link from the resource where it is found to another resource or action that may be performed.
var (
	listActivityFeedResponseItemFieldHref         = big.NewInt(1 << 0)
	listActivityFeedResponseItemFieldMethod       = big.NewInt(1 << 1)
	listActivityFeedResponseItemFieldRel          = big.NewInt(1 << 2)
	listActivityFeedResponseItemFieldSchema       = big.NewInt(1 << 3)
	listActivityFeedResponseItemFieldTargetSchema = big.NewInt(1 << 4)
)

type ListActivityFeedResponseItem struct {
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href *string `json:"href,omitempty" url:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method *ListActivityFeedResponseItemMethod `json:"method,omitempty" url:"method,omitempty"`
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

func (l *ListActivityFeedResponseItem) GetHref() *string {
	if l == nil {
		return nil
	}
	return l.Href
}

func (l *ListActivityFeedResponseItem) GetMethod() *ListActivityFeedResponseItemMethod {
	if l == nil {
		return nil
	}
	return l.Method
}

func (l *ListActivityFeedResponseItem) GetRel() *string {
	if l == nil {
		return nil
	}
	return l.Rel
}

func (l *ListActivityFeedResponseItem) GetSchema() *string {
	if l == nil {
		return nil
	}
	return l.Schema
}

func (l *ListActivityFeedResponseItem) GetTargetSchema() *string {
	if l == nil {
		return nil
	}
	return l.TargetSchema
}

func (l *ListActivityFeedResponseItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListActivityFeedResponseItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetHref sets the Href field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListActivityFeedResponseItem) SetHref(href *string) {
	l.Href = href
	l.require(listActivityFeedResponseItemFieldHref)
}

// SetMethod sets the Method field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListActivityFeedResponseItem) SetMethod(method *ListActivityFeedResponseItemMethod) {
	l.Method = method
	l.require(listActivityFeedResponseItemFieldMethod)
}

// SetRel sets the Rel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListActivityFeedResponseItem) SetRel(rel *string) {
	l.Rel = rel
	l.require(listActivityFeedResponseItemFieldRel)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListActivityFeedResponseItem) SetSchema(schema *string) {
	l.Schema = schema
	l.require(listActivityFeedResponseItemFieldSchema)
}

// SetTargetSchema sets the TargetSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListActivityFeedResponseItem) SetTargetSchema(targetSchema *string) {
	l.TargetSchema = targetSchema
	l.require(listActivityFeedResponseItemFieldTargetSchema)
}

func (l *ListActivityFeedResponseItem) UnmarshalJSON(data []byte) error {
	type unmarshaler ListActivityFeedResponseItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListActivityFeedResponseItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListActivityFeedResponseItem) MarshalJSON() ([]byte, error) {
	type embed ListActivityFeedResponseItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListActivityFeedResponseItem) String() string {
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
type ListActivityFeedResponseItemMethod string

const (
	ListActivityFeedResponseItemMethodGet     ListActivityFeedResponseItemMethod = "GET"
	ListActivityFeedResponseItemMethodPost    ListActivityFeedResponseItemMethod = "POST"
	ListActivityFeedResponseItemMethodPut     ListActivityFeedResponseItemMethod = "PUT"
	ListActivityFeedResponseItemMethodPatch   ListActivityFeedResponseItemMethod = "PATCH"
	ListActivityFeedResponseItemMethodDelete  ListActivityFeedResponseItemMethod = "DELETE"
	ListActivityFeedResponseItemMethodOptions ListActivityFeedResponseItemMethod = "OPTIONS"
	ListActivityFeedResponseItemMethodHead    ListActivityFeedResponseItemMethod = "HEAD"
)

func NewListActivityFeedResponseItemMethodFromString(s string) (ListActivityFeedResponseItemMethod, error) {
	switch s {
	case "GET":
		return ListActivityFeedResponseItemMethodGet, nil
	case "POST":
		return ListActivityFeedResponseItemMethodPost, nil
	case "PUT":
		return ListActivityFeedResponseItemMethodPut, nil
	case "PATCH":
		return ListActivityFeedResponseItemMethodPatch, nil
	case "DELETE":
		return ListActivityFeedResponseItemMethodDelete, nil
	case "OPTIONS":
		return ListActivityFeedResponseItemMethodOptions, nil
	case "HEAD":
		return ListActivityFeedResponseItemMethodHead, nil
	}
	var t ListActivityFeedResponseItemMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListActivityFeedResponseItemMethod) Ptr() *ListActivityFeedResponseItemMethod {
	return &l
}

// An array of Chimp Chatter messages. There's a maximum of 200 messages present for an account.
var (
	listChimpChatterActivityFeedResponseFieldLinks        = big.NewInt(1 << 0)
	listChimpChatterActivityFeedResponseFieldChimpChatter = big.NewInt(1 << 1)
	listChimpChatterActivityFeedResponseFieldTotalItems   = big.NewInt(1 << 2)
)

type ListChimpChatterActivityFeedResponse struct {
	// A list of link types and descriptions for the API schema documents.
	Links []*ListChimpChatterActivityFeedResponseLinksItem `json:"_links,omitempty" url:"_links,omitempty"`
	// An array of Chimp Chatter messages. There's a maximum of 200 messages present for an account.
	ChimpChatter []*ListChimpChatterActivityFeedResponseChimpChatterItem `json:"chimp_chatter,omitempty" url:"chimp_chatter,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int `json:"total_items,omitempty" url:"total_items,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListChimpChatterActivityFeedResponse) GetLinks() []*ListChimpChatterActivityFeedResponseLinksItem {
	if l == nil {
		return nil
	}
	return l.Links
}

func (l *ListChimpChatterActivityFeedResponse) GetChimpChatter() []*ListChimpChatterActivityFeedResponseChimpChatterItem {
	if l == nil {
		return nil
	}
	return l.ChimpChatter
}

func (l *ListChimpChatterActivityFeedResponse) GetTotalItems() *int {
	if l == nil {
		return nil
	}
	return l.TotalItems
}

func (l *ListChimpChatterActivityFeedResponse) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListChimpChatterActivityFeedResponse) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetLinks sets the Links field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListChimpChatterActivityFeedResponse) SetLinks(links []*ListChimpChatterActivityFeedResponseLinksItem) {
	l.Links = links
	l.require(listChimpChatterActivityFeedResponseFieldLinks)
}

// SetChimpChatter sets the ChimpChatter field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListChimpChatterActivityFeedResponse) SetChimpChatter(chimpChatter []*ListChimpChatterActivityFeedResponseChimpChatterItem) {
	l.ChimpChatter = chimpChatter
	l.require(listChimpChatterActivityFeedResponseFieldChimpChatter)
}

// SetTotalItems sets the TotalItems field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListChimpChatterActivityFeedResponse) SetTotalItems(totalItems *int) {
	l.TotalItems = totalItems
	l.require(listChimpChatterActivityFeedResponseFieldTotalItems)
}

func (l *ListChimpChatterActivityFeedResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListChimpChatterActivityFeedResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListChimpChatterActivityFeedResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListChimpChatterActivityFeedResponse) MarshalJSON() ([]byte, error) {
	type embed ListChimpChatterActivityFeedResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListChimpChatterActivityFeedResponse) String() string {
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

// A Chimp Chatter message
var (
	listChimpChatterActivityFeedResponseChimpChatterItemFieldCampaignID = big.NewInt(1 << 0)
	listChimpChatterActivityFeedResponseChimpChatterItemFieldListID     = big.NewInt(1 << 1)
	listChimpChatterActivityFeedResponseChimpChatterItemFieldMessage    = big.NewInt(1 << 2)
	listChimpChatterActivityFeedResponseChimpChatterItemFieldTitle      = big.NewInt(1 << 3)
	listChimpChatterActivityFeedResponseChimpChatterItemFieldType       = big.NewInt(1 << 4)
	listChimpChatterActivityFeedResponseChimpChatterItemFieldUpdateTime = big.NewInt(1 << 5)
	listChimpChatterActivityFeedResponseChimpChatterItemFieldURL        = big.NewInt(1 << 6)
)

type ListChimpChatterActivityFeedResponseChimpChatterItem struct {
	// If it exists, campaign ID for the associated campaign
	CampaignID *string `json:"campaign_id,omitempty" url:"campaign_id,omitempty"`
	// If it exists, list ID for the associated list
	ListID  *string `json:"list_id,omitempty" url:"list_id,omitempty"`
	Message *string `json:"message,omitempty" url:"message,omitempty"`
	Title   *string `json:"title,omitempty" url:"title,omitempty"`
	// The type of activity
	Type *ListChimpChatterActivityFeedResponseChimpChatterItemType `json:"type,omitempty" url:"type,omitempty"`
	// The date and time this activity was updated.
	UpdateTime *time.Time `json:"update_time,omitempty" url:"update_time,omitempty"`
	// URL to a report that includes this activity
	URL *string `json:"url,omitempty" url:"url,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListChimpChatterActivityFeedResponseChimpChatterItem) GetCampaignID() *string {
	if l == nil {
		return nil
	}
	return l.CampaignID
}

func (l *ListChimpChatterActivityFeedResponseChimpChatterItem) GetListID() *string {
	if l == nil {
		return nil
	}
	return l.ListID
}

func (l *ListChimpChatterActivityFeedResponseChimpChatterItem) GetMessage() *string {
	if l == nil {
		return nil
	}
	return l.Message
}

func (l *ListChimpChatterActivityFeedResponseChimpChatterItem) GetTitle() *string {
	if l == nil {
		return nil
	}
	return l.Title
}

func (l *ListChimpChatterActivityFeedResponseChimpChatterItem) GetType() *ListChimpChatterActivityFeedResponseChimpChatterItemType {
	if l == nil {
		return nil
	}
	return l.Type
}

func (l *ListChimpChatterActivityFeedResponseChimpChatterItem) GetUpdateTime() *time.Time {
	if l == nil {
		return nil
	}
	return l.UpdateTime
}

func (l *ListChimpChatterActivityFeedResponseChimpChatterItem) GetURL() *string {
	if l == nil {
		return nil
	}
	return l.URL
}

func (l *ListChimpChatterActivityFeedResponseChimpChatterItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListChimpChatterActivityFeedResponseChimpChatterItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetCampaignID sets the CampaignID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListChimpChatterActivityFeedResponseChimpChatterItem) SetCampaignID(campaignID *string) {
	l.CampaignID = campaignID
	l.require(listChimpChatterActivityFeedResponseChimpChatterItemFieldCampaignID)
}

// SetListID sets the ListID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListChimpChatterActivityFeedResponseChimpChatterItem) SetListID(listID *string) {
	l.ListID = listID
	l.require(listChimpChatterActivityFeedResponseChimpChatterItemFieldListID)
}

// SetMessage sets the Message field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListChimpChatterActivityFeedResponseChimpChatterItem) SetMessage(message *string) {
	l.Message = message
	l.require(listChimpChatterActivityFeedResponseChimpChatterItemFieldMessage)
}

// SetTitle sets the Title field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListChimpChatterActivityFeedResponseChimpChatterItem) SetTitle(title *string) {
	l.Title = title
	l.require(listChimpChatterActivityFeedResponseChimpChatterItemFieldTitle)
}

// SetType sets the Type field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListChimpChatterActivityFeedResponseChimpChatterItem) SetType(type_ *ListChimpChatterActivityFeedResponseChimpChatterItemType) {
	l.Type = type_
	l.require(listChimpChatterActivityFeedResponseChimpChatterItemFieldType)
}

// SetUpdateTime sets the UpdateTime field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListChimpChatterActivityFeedResponseChimpChatterItem) SetUpdateTime(updateTime *time.Time) {
	l.UpdateTime = updateTime
	l.require(listChimpChatterActivityFeedResponseChimpChatterItemFieldUpdateTime)
}

// SetURL sets the URL field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListChimpChatterActivityFeedResponseChimpChatterItem) SetURL(url *string) {
	l.URL = url
	l.require(listChimpChatterActivityFeedResponseChimpChatterItemFieldURL)
}

func (l *ListChimpChatterActivityFeedResponseChimpChatterItem) UnmarshalJSON(data []byte) error {
	type embed ListChimpChatterActivityFeedResponseChimpChatterItem
	var unmarshaler = struct {
		embed
		UpdateTime *internal.DateTime `json:"update_time,omitempty"`
	}{
		embed: embed(*l),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*l = ListChimpChatterActivityFeedResponseChimpChatterItem(unmarshaler.embed)
	l.UpdateTime = unmarshaler.UpdateTime.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListChimpChatterActivityFeedResponseChimpChatterItem) MarshalJSON() ([]byte, error) {
	type embed ListChimpChatterActivityFeedResponseChimpChatterItem
	var marshaler = struct {
		embed
		UpdateTime *internal.DateTime `json:"update_time,omitempty"`
	}{
		embed:      embed(*l),
		UpdateTime: internal.NewOptionalDateTime(l.UpdateTime),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListChimpChatterActivityFeedResponseChimpChatterItem) String() string {
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

// The type of activity
type ListChimpChatterActivityFeedResponseChimpChatterItemType string

const (
	ListChimpChatterActivityFeedResponseChimpChatterItemTypeListsNewSubscriber       ListChimpChatterActivityFeedResponseChimpChatterItemType = "lists:new-subscriber"
	ListChimpChatterActivityFeedResponseChimpChatterItemTypeListsUnsubscribes        ListChimpChatterActivityFeedResponseChimpChatterItemType = "lists:unsubscribes"
	ListChimpChatterActivityFeedResponseChimpChatterItemTypeListsProfileUpdates      ListChimpChatterActivityFeedResponseChimpChatterItemType = "lists:profile-updates"
	ListChimpChatterActivityFeedResponseChimpChatterItemTypeCampaignsFacebookLikes   ListChimpChatterActivityFeedResponseChimpChatterItemType = "campaigns:facebook-likes"
	ListChimpChatterActivityFeedResponseChimpChatterItemTypeCampaignsForwardToFriend ListChimpChatterActivityFeedResponseChimpChatterItemType = "campaigns:forward-to-friend"
	ListChimpChatterActivityFeedResponseChimpChatterItemTypeListsImports             ListChimpChatterActivityFeedResponseChimpChatterItemType = "lists:imports"
)

func NewListChimpChatterActivityFeedResponseChimpChatterItemTypeFromString(s string) (ListChimpChatterActivityFeedResponseChimpChatterItemType, error) {
	switch s {
	case "lists:new-subscriber":
		return ListChimpChatterActivityFeedResponseChimpChatterItemTypeListsNewSubscriber, nil
	case "lists:unsubscribes":
		return ListChimpChatterActivityFeedResponseChimpChatterItemTypeListsUnsubscribes, nil
	case "lists:profile-updates":
		return ListChimpChatterActivityFeedResponseChimpChatterItemTypeListsProfileUpdates, nil
	case "campaigns:facebook-likes":
		return ListChimpChatterActivityFeedResponseChimpChatterItemTypeCampaignsFacebookLikes, nil
	case "campaigns:forward-to-friend":
		return ListChimpChatterActivityFeedResponseChimpChatterItemTypeCampaignsForwardToFriend, nil
	case "lists:imports":
		return ListChimpChatterActivityFeedResponseChimpChatterItemTypeListsImports, nil
	}
	var t ListChimpChatterActivityFeedResponseChimpChatterItemType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListChimpChatterActivityFeedResponseChimpChatterItemType) Ptr() *ListChimpChatterActivityFeedResponseChimpChatterItemType {
	return &l
}

// This object represents a link from the resource where it is found to another resource or action that may be performed.
var (
	listChimpChatterActivityFeedResponseLinksItemFieldHref         = big.NewInt(1 << 0)
	listChimpChatterActivityFeedResponseLinksItemFieldMethod       = big.NewInt(1 << 1)
	listChimpChatterActivityFeedResponseLinksItemFieldRel          = big.NewInt(1 << 2)
	listChimpChatterActivityFeedResponseLinksItemFieldSchema       = big.NewInt(1 << 3)
	listChimpChatterActivityFeedResponseLinksItemFieldTargetSchema = big.NewInt(1 << 4)
)

type ListChimpChatterActivityFeedResponseLinksItem struct {
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href *string `json:"href,omitempty" url:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method *ListChimpChatterActivityFeedResponseLinksItemMethod `json:"method,omitempty" url:"method,omitempty"`
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

func (l *ListChimpChatterActivityFeedResponseLinksItem) GetHref() *string {
	if l == nil {
		return nil
	}
	return l.Href
}

func (l *ListChimpChatterActivityFeedResponseLinksItem) GetMethod() *ListChimpChatterActivityFeedResponseLinksItemMethod {
	if l == nil {
		return nil
	}
	return l.Method
}

func (l *ListChimpChatterActivityFeedResponseLinksItem) GetRel() *string {
	if l == nil {
		return nil
	}
	return l.Rel
}

func (l *ListChimpChatterActivityFeedResponseLinksItem) GetSchema() *string {
	if l == nil {
		return nil
	}
	return l.Schema
}

func (l *ListChimpChatterActivityFeedResponseLinksItem) GetTargetSchema() *string {
	if l == nil {
		return nil
	}
	return l.TargetSchema
}

func (l *ListChimpChatterActivityFeedResponseLinksItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListChimpChatterActivityFeedResponseLinksItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetHref sets the Href field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListChimpChatterActivityFeedResponseLinksItem) SetHref(href *string) {
	l.Href = href
	l.require(listChimpChatterActivityFeedResponseLinksItemFieldHref)
}

// SetMethod sets the Method field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListChimpChatterActivityFeedResponseLinksItem) SetMethod(method *ListChimpChatterActivityFeedResponseLinksItemMethod) {
	l.Method = method
	l.require(listChimpChatterActivityFeedResponseLinksItemFieldMethod)
}

// SetRel sets the Rel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListChimpChatterActivityFeedResponseLinksItem) SetRel(rel *string) {
	l.Rel = rel
	l.require(listChimpChatterActivityFeedResponseLinksItemFieldRel)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListChimpChatterActivityFeedResponseLinksItem) SetSchema(schema *string) {
	l.Schema = schema
	l.require(listChimpChatterActivityFeedResponseLinksItemFieldSchema)
}

// SetTargetSchema sets the TargetSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListChimpChatterActivityFeedResponseLinksItem) SetTargetSchema(targetSchema *string) {
	l.TargetSchema = targetSchema
	l.require(listChimpChatterActivityFeedResponseLinksItemFieldTargetSchema)
}

func (l *ListChimpChatterActivityFeedResponseLinksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler ListChimpChatterActivityFeedResponseLinksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListChimpChatterActivityFeedResponseLinksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListChimpChatterActivityFeedResponseLinksItem) MarshalJSON() ([]byte, error) {
	type embed ListChimpChatterActivityFeedResponseLinksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListChimpChatterActivityFeedResponseLinksItem) String() string {
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
type ListChimpChatterActivityFeedResponseLinksItemMethod string

const (
	ListChimpChatterActivityFeedResponseLinksItemMethodGet     ListChimpChatterActivityFeedResponseLinksItemMethod = "GET"
	ListChimpChatterActivityFeedResponseLinksItemMethodPost    ListChimpChatterActivityFeedResponseLinksItemMethod = "POST"
	ListChimpChatterActivityFeedResponseLinksItemMethodPut     ListChimpChatterActivityFeedResponseLinksItemMethod = "PUT"
	ListChimpChatterActivityFeedResponseLinksItemMethodPatch   ListChimpChatterActivityFeedResponseLinksItemMethod = "PATCH"
	ListChimpChatterActivityFeedResponseLinksItemMethodDelete  ListChimpChatterActivityFeedResponseLinksItemMethod = "DELETE"
	ListChimpChatterActivityFeedResponseLinksItemMethodOptions ListChimpChatterActivityFeedResponseLinksItemMethod = "OPTIONS"
	ListChimpChatterActivityFeedResponseLinksItemMethodHead    ListChimpChatterActivityFeedResponseLinksItemMethod = "HEAD"
)

func NewListChimpChatterActivityFeedResponseLinksItemMethodFromString(s string) (ListChimpChatterActivityFeedResponseLinksItemMethod, error) {
	switch s {
	case "GET":
		return ListChimpChatterActivityFeedResponseLinksItemMethodGet, nil
	case "POST":
		return ListChimpChatterActivityFeedResponseLinksItemMethodPost, nil
	case "PUT":
		return ListChimpChatterActivityFeedResponseLinksItemMethodPut, nil
	case "PATCH":
		return ListChimpChatterActivityFeedResponseLinksItemMethodPatch, nil
	case "DELETE":
		return ListChimpChatterActivityFeedResponseLinksItemMethodDelete, nil
	case "OPTIONS":
		return ListChimpChatterActivityFeedResponseLinksItemMethodOptions, nil
	case "HEAD":
		return ListChimpChatterActivityFeedResponseLinksItemMethodHead, nil
	}
	var t ListChimpChatterActivityFeedResponseLinksItemMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListChimpChatterActivityFeedResponseLinksItemMethod) Ptr() *ListChimpChatterActivityFeedResponseLinksItemMethod {
	return &l
}
