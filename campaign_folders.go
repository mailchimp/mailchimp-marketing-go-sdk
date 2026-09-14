// Code generated from our API definition. DO NOT EDIT.

package api

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/mailchimp/mailchimp-marketing-go-sdk/internal"
	big "math/big"
)

var (
	createCampaignFoldersRequestFieldName = big.NewInt(1 << 0)
)

type CreateCampaignFoldersRequest struct {
	// Name to associate with the folder.
	Name string `json:"name" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CreateCampaignFoldersRequest) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateCampaignFoldersRequest) SetName(name string) {
	c.Name = name
	c.require(createCampaignFoldersRequestFieldName)
}

func (c *CreateCampaignFoldersRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateCampaignFoldersRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*c = CreateCampaignFoldersRequest(body)
	return nil
}

func (c *CreateCampaignFoldersRequest) MarshalJSON() ([]byte, error) {
	type embed CreateCampaignFoldersRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

var (
	deleteCampaignFoldersRequestFieldFolderID = big.NewInt(1 << 0)
)

type DeleteCampaignFoldersRequest struct {
	// The unique id for the campaign folder.
	FolderID string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (d *DeleteCampaignFoldersRequest) require(field *big.Int) {
	if d.explicitFields == nil {
		d.explicitFields = big.NewInt(0)
	}
	d.explicitFields.Or(d.explicitFields, field)
}

// SetFolderID sets the FolderID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (d *DeleteCampaignFoldersRequest) SetFolderID(folderID string) {
	d.FolderID = folderID
	d.require(deleteCampaignFoldersRequestFieldFolderID)
}

var (
	getCampaignFoldersRequestFieldFolderID      = big.NewInt(1 << 0)
	getCampaignFoldersRequestFieldFields        = big.NewInt(1 << 1)
	getCampaignFoldersRequestFieldExcludeFields = big.NewInt(1 << 2)
)

type GetCampaignFoldersRequest struct {
	// The unique id for the campaign folder.
	FolderID string `json:"-" url:"-"`
	// A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
	Fields []*string `json:"-" url:"fields,omitempty"`
	// A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
	ExcludeFields []*string `json:"-" url:"exclude_fields,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetCampaignFoldersRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetFolderID sets the FolderID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetCampaignFoldersRequest) SetFolderID(folderID string) {
	g.FolderID = folderID
	g.require(getCampaignFoldersRequestFieldFolderID)
}

// SetFields sets the Fields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetCampaignFoldersRequest) SetFields(fields []*string) {
	g.Fields = fields
	g.require(getCampaignFoldersRequestFieldFields)
}

// SetExcludeFields sets the ExcludeFields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetCampaignFoldersRequest) SetExcludeFields(excludeFields []*string) {
	g.ExcludeFields = excludeFields
	g.require(getCampaignFoldersRequestFieldExcludeFields)
}

var (
	listCampaignFoldersRequestFieldFields        = big.NewInt(1 << 0)
	listCampaignFoldersRequestFieldExcludeFields = big.NewInt(1 << 1)
	listCampaignFoldersRequestFieldCount         = big.NewInt(1 << 2)
	listCampaignFoldersRequestFieldOffset        = big.NewInt(1 << 3)
)

type ListCampaignFoldersRequest struct {
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

func (l *ListCampaignFoldersRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetFields sets the Fields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListCampaignFoldersRequest) SetFields(fields []*string) {
	l.Fields = fields
	l.require(listCampaignFoldersRequestFieldFields)
}

// SetExcludeFields sets the ExcludeFields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListCampaignFoldersRequest) SetExcludeFields(excludeFields []*string) {
	l.ExcludeFields = excludeFields
	l.require(listCampaignFoldersRequestFieldExcludeFields)
}

// SetCount sets the Count field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListCampaignFoldersRequest) SetCount(count *int) {
	l.Count = count
	l.require(listCampaignFoldersRequestFieldCount)
}

// SetOffset sets the Offset field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListCampaignFoldersRequest) SetOffset(offset *int) {
	l.Offset = offset
	l.require(listCampaignFoldersRequestFieldOffset)
}

// A list of campaign folders
var (
	campaignFoldersFieldLinks      = big.NewInt(1 << 0)
	campaignFoldersFieldFolders    = big.NewInt(1 << 1)
	campaignFoldersFieldTotalItems = big.NewInt(1 << 2)
)

type CampaignFolders struct {
	// A list of link types and descriptions for the API schema documents.
	Links []*CampaignFoldersLinksItem `json:"_links,omitempty" url:"_links,omitempty"`
	// An array of objects representing campaign folders.
	Folders []*CampaignFoldersFoldersItem `json:"folders,omitempty" url:"folders,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int `json:"total_items,omitempty" url:"total_items,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CampaignFolders) GetLinks() []*CampaignFoldersLinksItem {
	if c == nil {
		return nil
	}
	return c.Links
}

func (c *CampaignFolders) GetFolders() []*CampaignFoldersFoldersItem {
	if c == nil {
		return nil
	}
	return c.Folders
}

func (c *CampaignFolders) GetTotalItems() *int {
	if c == nil {
		return nil
	}
	return c.TotalItems
}

func (c *CampaignFolders) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *CampaignFolders) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetLinks sets the Links field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CampaignFolders) SetLinks(links []*CampaignFoldersLinksItem) {
	c.Links = links
	c.require(campaignFoldersFieldLinks)
}

// SetFolders sets the Folders field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CampaignFolders) SetFolders(folders []*CampaignFoldersFoldersItem) {
	c.Folders = folders
	c.require(campaignFoldersFieldFolders)
}

// SetTotalItems sets the TotalItems field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CampaignFolders) SetTotalItems(totalItems *int) {
	c.TotalItems = totalItems
	c.require(campaignFoldersFieldTotalItems)
}

func (c *CampaignFolders) UnmarshalJSON(data []byte) error {
	type unmarshaler CampaignFolders
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CampaignFolders(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CampaignFolders) MarshalJSON() ([]byte, error) {
	type embed CampaignFolders
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *CampaignFolders) String() string {
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

// A folder used to organize campaigns.
var (
	campaignFoldersFoldersItemFieldLinks = big.NewInt(1 << 0)
	campaignFoldersFoldersItemFieldCount = big.NewInt(1 << 1)
	campaignFoldersFoldersItemFieldID    = big.NewInt(1 << 2)
	campaignFoldersFoldersItemFieldName  = big.NewInt(1 << 3)
)

type CampaignFoldersFoldersItem struct {
	// A list of link types and descriptions for the API schema documents.
	Links []*CampaignFoldersFoldersItemLinksItem `json:"_links,omitempty" url:"_links,omitempty"`
	// The number of campaigns in the folder.
	Count *int `json:"count,omitempty" url:"count,omitempty"`
	// A string that uniquely identifies this campaign folder.
	ID *string `json:"id,omitempty" url:"id,omitempty"`
	// The name of the folder.
	Name *string `json:"name,omitempty" url:"name,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CampaignFoldersFoldersItem) GetLinks() []*CampaignFoldersFoldersItemLinksItem {
	if c == nil {
		return nil
	}
	return c.Links
}

func (c *CampaignFoldersFoldersItem) GetCount() *int {
	if c == nil {
		return nil
	}
	return c.Count
}

func (c *CampaignFoldersFoldersItem) GetID() *string {
	if c == nil {
		return nil
	}
	return c.ID
}

func (c *CampaignFoldersFoldersItem) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CampaignFoldersFoldersItem) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *CampaignFoldersFoldersItem) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetLinks sets the Links field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CampaignFoldersFoldersItem) SetLinks(links []*CampaignFoldersFoldersItemLinksItem) {
	c.Links = links
	c.require(campaignFoldersFoldersItemFieldLinks)
}

// SetCount sets the Count field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CampaignFoldersFoldersItem) SetCount(count *int) {
	c.Count = count
	c.require(campaignFoldersFoldersItemFieldCount)
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CampaignFoldersFoldersItem) SetID(id *string) {
	c.ID = id
	c.require(campaignFoldersFoldersItemFieldID)
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CampaignFoldersFoldersItem) SetName(name *string) {
	c.Name = name
	c.require(campaignFoldersFoldersItemFieldName)
}

func (c *CampaignFoldersFoldersItem) UnmarshalJSON(data []byte) error {
	type unmarshaler CampaignFoldersFoldersItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CampaignFoldersFoldersItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CampaignFoldersFoldersItem) MarshalJSON() ([]byte, error) {
	type embed CampaignFoldersFoldersItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *CampaignFoldersFoldersItem) String() string {
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

// This object represents a link from the resource where it is found to another resource or action that may be performed.
var (
	campaignFoldersFoldersItemLinksItemFieldHref         = big.NewInt(1 << 0)
	campaignFoldersFoldersItemLinksItemFieldMethod       = big.NewInt(1 << 1)
	campaignFoldersFoldersItemLinksItemFieldRel          = big.NewInt(1 << 2)
	campaignFoldersFoldersItemLinksItemFieldSchema       = big.NewInt(1 << 3)
	campaignFoldersFoldersItemLinksItemFieldTargetSchema = big.NewInt(1 << 4)
)

type CampaignFoldersFoldersItemLinksItem struct {
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href *string `json:"href,omitempty" url:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method *CampaignFoldersFoldersItemLinksItemMethod `json:"method,omitempty" url:"method,omitempty"`
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

func (c *CampaignFoldersFoldersItemLinksItem) GetHref() *string {
	if c == nil {
		return nil
	}
	return c.Href
}

func (c *CampaignFoldersFoldersItemLinksItem) GetMethod() *CampaignFoldersFoldersItemLinksItemMethod {
	if c == nil {
		return nil
	}
	return c.Method
}

func (c *CampaignFoldersFoldersItemLinksItem) GetRel() *string {
	if c == nil {
		return nil
	}
	return c.Rel
}

func (c *CampaignFoldersFoldersItemLinksItem) GetSchema() *string {
	if c == nil {
		return nil
	}
	return c.Schema
}

func (c *CampaignFoldersFoldersItemLinksItem) GetTargetSchema() *string {
	if c == nil {
		return nil
	}
	return c.TargetSchema
}

func (c *CampaignFoldersFoldersItemLinksItem) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *CampaignFoldersFoldersItemLinksItem) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetHref sets the Href field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CampaignFoldersFoldersItemLinksItem) SetHref(href *string) {
	c.Href = href
	c.require(campaignFoldersFoldersItemLinksItemFieldHref)
}

// SetMethod sets the Method field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CampaignFoldersFoldersItemLinksItem) SetMethod(method *CampaignFoldersFoldersItemLinksItemMethod) {
	c.Method = method
	c.require(campaignFoldersFoldersItemLinksItemFieldMethod)
}

// SetRel sets the Rel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CampaignFoldersFoldersItemLinksItem) SetRel(rel *string) {
	c.Rel = rel
	c.require(campaignFoldersFoldersItemLinksItemFieldRel)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CampaignFoldersFoldersItemLinksItem) SetSchema(schema *string) {
	c.Schema = schema
	c.require(campaignFoldersFoldersItemLinksItemFieldSchema)
}

// SetTargetSchema sets the TargetSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CampaignFoldersFoldersItemLinksItem) SetTargetSchema(targetSchema *string) {
	c.TargetSchema = targetSchema
	c.require(campaignFoldersFoldersItemLinksItemFieldTargetSchema)
}

func (c *CampaignFoldersFoldersItemLinksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler CampaignFoldersFoldersItemLinksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CampaignFoldersFoldersItemLinksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CampaignFoldersFoldersItemLinksItem) MarshalJSON() ([]byte, error) {
	type embed CampaignFoldersFoldersItemLinksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *CampaignFoldersFoldersItemLinksItem) String() string {
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

// The HTTP method that should be used when accessing the URL defined in 'href'.
type CampaignFoldersFoldersItemLinksItemMethod string

const (
	CampaignFoldersFoldersItemLinksItemMethodGet     CampaignFoldersFoldersItemLinksItemMethod = "GET"
	CampaignFoldersFoldersItemLinksItemMethodPost    CampaignFoldersFoldersItemLinksItemMethod = "POST"
	CampaignFoldersFoldersItemLinksItemMethodPut     CampaignFoldersFoldersItemLinksItemMethod = "PUT"
	CampaignFoldersFoldersItemLinksItemMethodPatch   CampaignFoldersFoldersItemLinksItemMethod = "PATCH"
	CampaignFoldersFoldersItemLinksItemMethodDelete  CampaignFoldersFoldersItemLinksItemMethod = "DELETE"
	CampaignFoldersFoldersItemLinksItemMethodOptions CampaignFoldersFoldersItemLinksItemMethod = "OPTIONS"
	CampaignFoldersFoldersItemLinksItemMethodHead    CampaignFoldersFoldersItemLinksItemMethod = "HEAD"
)

func NewCampaignFoldersFoldersItemLinksItemMethodFromString(s string) (CampaignFoldersFoldersItemLinksItemMethod, error) {
	switch s {
	case "GET":
		return CampaignFoldersFoldersItemLinksItemMethodGet, nil
	case "POST":
		return CampaignFoldersFoldersItemLinksItemMethodPost, nil
	case "PUT":
		return CampaignFoldersFoldersItemLinksItemMethodPut, nil
	case "PATCH":
		return CampaignFoldersFoldersItemLinksItemMethodPatch, nil
	case "DELETE":
		return CampaignFoldersFoldersItemLinksItemMethodDelete, nil
	case "OPTIONS":
		return CampaignFoldersFoldersItemLinksItemMethodOptions, nil
	case "HEAD":
		return CampaignFoldersFoldersItemLinksItemMethodHead, nil
	}
	var t CampaignFoldersFoldersItemLinksItemMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c CampaignFoldersFoldersItemLinksItemMethod) Ptr() *CampaignFoldersFoldersItemLinksItemMethod {
	return &c
}

// This object represents a link from the resource where it is found to another resource or action that may be performed.
var (
	campaignFoldersLinksItemFieldHref         = big.NewInt(1 << 0)
	campaignFoldersLinksItemFieldMethod       = big.NewInt(1 << 1)
	campaignFoldersLinksItemFieldRel          = big.NewInt(1 << 2)
	campaignFoldersLinksItemFieldSchema       = big.NewInt(1 << 3)
	campaignFoldersLinksItemFieldTargetSchema = big.NewInt(1 << 4)
)

type CampaignFoldersLinksItem struct {
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href *string `json:"href,omitempty" url:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method *CampaignFoldersLinksItemMethod `json:"method,omitempty" url:"method,omitempty"`
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

func (c *CampaignFoldersLinksItem) GetHref() *string {
	if c == nil {
		return nil
	}
	return c.Href
}

func (c *CampaignFoldersLinksItem) GetMethod() *CampaignFoldersLinksItemMethod {
	if c == nil {
		return nil
	}
	return c.Method
}

func (c *CampaignFoldersLinksItem) GetRel() *string {
	if c == nil {
		return nil
	}
	return c.Rel
}

func (c *CampaignFoldersLinksItem) GetSchema() *string {
	if c == nil {
		return nil
	}
	return c.Schema
}

func (c *CampaignFoldersLinksItem) GetTargetSchema() *string {
	if c == nil {
		return nil
	}
	return c.TargetSchema
}

func (c *CampaignFoldersLinksItem) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *CampaignFoldersLinksItem) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetHref sets the Href field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CampaignFoldersLinksItem) SetHref(href *string) {
	c.Href = href
	c.require(campaignFoldersLinksItemFieldHref)
}

// SetMethod sets the Method field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CampaignFoldersLinksItem) SetMethod(method *CampaignFoldersLinksItemMethod) {
	c.Method = method
	c.require(campaignFoldersLinksItemFieldMethod)
}

// SetRel sets the Rel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CampaignFoldersLinksItem) SetRel(rel *string) {
	c.Rel = rel
	c.require(campaignFoldersLinksItemFieldRel)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CampaignFoldersLinksItem) SetSchema(schema *string) {
	c.Schema = schema
	c.require(campaignFoldersLinksItemFieldSchema)
}

// SetTargetSchema sets the TargetSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CampaignFoldersLinksItem) SetTargetSchema(targetSchema *string) {
	c.TargetSchema = targetSchema
	c.require(campaignFoldersLinksItemFieldTargetSchema)
}

func (c *CampaignFoldersLinksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler CampaignFoldersLinksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CampaignFoldersLinksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CampaignFoldersLinksItem) MarshalJSON() ([]byte, error) {
	type embed CampaignFoldersLinksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *CampaignFoldersLinksItem) String() string {
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

// The HTTP method that should be used when accessing the URL defined in 'href'.
type CampaignFoldersLinksItemMethod string

const (
	CampaignFoldersLinksItemMethodGet     CampaignFoldersLinksItemMethod = "GET"
	CampaignFoldersLinksItemMethodPost    CampaignFoldersLinksItemMethod = "POST"
	CampaignFoldersLinksItemMethodPut     CampaignFoldersLinksItemMethod = "PUT"
	CampaignFoldersLinksItemMethodPatch   CampaignFoldersLinksItemMethod = "PATCH"
	CampaignFoldersLinksItemMethodDelete  CampaignFoldersLinksItemMethod = "DELETE"
	CampaignFoldersLinksItemMethodOptions CampaignFoldersLinksItemMethod = "OPTIONS"
	CampaignFoldersLinksItemMethodHead    CampaignFoldersLinksItemMethod = "HEAD"
)

func NewCampaignFoldersLinksItemMethodFromString(s string) (CampaignFoldersLinksItemMethod, error) {
	switch s {
	case "GET":
		return CampaignFoldersLinksItemMethodGet, nil
	case "POST":
		return CampaignFoldersLinksItemMethodPost, nil
	case "PUT":
		return CampaignFoldersLinksItemMethodPut, nil
	case "PATCH":
		return CampaignFoldersLinksItemMethodPatch, nil
	case "DELETE":
		return CampaignFoldersLinksItemMethodDelete, nil
	case "OPTIONS":
		return CampaignFoldersLinksItemMethodOptions, nil
	case "HEAD":
		return CampaignFoldersLinksItemMethodHead, nil
	}
	var t CampaignFoldersLinksItemMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c CampaignFoldersLinksItemMethod) Ptr() *CampaignFoldersLinksItemMethod {
	return &c
}

// A folder used to organize campaigns.
var (
	getCampaignFoldersResponseFieldLinks = big.NewInt(1 << 0)
	getCampaignFoldersResponseFieldCount = big.NewInt(1 << 1)
	getCampaignFoldersResponseFieldID    = big.NewInt(1 << 2)
	getCampaignFoldersResponseFieldName  = big.NewInt(1 << 3)
)

type GetCampaignFoldersResponse struct {
	// A list of link types and descriptions for the API schema documents.
	Links []*GetCampaignFoldersResponseLinksItem `json:"_links,omitempty" url:"_links,omitempty"`
	// The number of campaigns in the folder.
	Count *int `json:"count,omitempty" url:"count,omitempty"`
	// A string that uniquely identifies this campaign folder.
	ID *string `json:"id,omitempty" url:"id,omitempty"`
	// The name of the folder.
	Name *string `json:"name,omitempty" url:"name,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (g *GetCampaignFoldersResponse) GetLinks() []*GetCampaignFoldersResponseLinksItem {
	if g == nil {
		return nil
	}
	return g.Links
}

func (g *GetCampaignFoldersResponse) GetCount() *int {
	if g == nil {
		return nil
	}
	return g.Count
}

func (g *GetCampaignFoldersResponse) GetID() *string {
	if g == nil {
		return nil
	}
	return g.ID
}

func (g *GetCampaignFoldersResponse) GetName() *string {
	if g == nil {
		return nil
	}
	return g.Name
}

func (g *GetCampaignFoldersResponse) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g *GetCampaignFoldersResponse) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetLinks sets the Links field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetCampaignFoldersResponse) SetLinks(links []*GetCampaignFoldersResponseLinksItem) {
	g.Links = links
	g.require(getCampaignFoldersResponseFieldLinks)
}

// SetCount sets the Count field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetCampaignFoldersResponse) SetCount(count *int) {
	g.Count = count
	g.require(getCampaignFoldersResponseFieldCount)
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetCampaignFoldersResponse) SetID(id *string) {
	g.ID = id
	g.require(getCampaignFoldersResponseFieldID)
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetCampaignFoldersResponse) SetName(name *string) {
	g.Name = name
	g.require(getCampaignFoldersResponseFieldName)
}

func (g *GetCampaignFoldersResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetCampaignFoldersResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetCampaignFoldersResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetCampaignFoldersResponse) MarshalJSON() ([]byte, error) {
	type embed GetCampaignFoldersResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, g.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (g *GetCampaignFoldersResponse) String() string {
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
	getCampaignFoldersResponseLinksItemFieldHref         = big.NewInt(1 << 0)
	getCampaignFoldersResponseLinksItemFieldMethod       = big.NewInt(1 << 1)
	getCampaignFoldersResponseLinksItemFieldRel          = big.NewInt(1 << 2)
	getCampaignFoldersResponseLinksItemFieldSchema       = big.NewInt(1 << 3)
	getCampaignFoldersResponseLinksItemFieldTargetSchema = big.NewInt(1 << 4)
)

type GetCampaignFoldersResponseLinksItem struct {
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href *string `json:"href,omitempty" url:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method *GetCampaignFoldersResponseLinksItemMethod `json:"method,omitempty" url:"method,omitempty"`
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

func (g *GetCampaignFoldersResponseLinksItem) GetHref() *string {
	if g == nil {
		return nil
	}
	return g.Href
}

func (g *GetCampaignFoldersResponseLinksItem) GetMethod() *GetCampaignFoldersResponseLinksItemMethod {
	if g == nil {
		return nil
	}
	return g.Method
}

func (g *GetCampaignFoldersResponseLinksItem) GetRel() *string {
	if g == nil {
		return nil
	}
	return g.Rel
}

func (g *GetCampaignFoldersResponseLinksItem) GetSchema() *string {
	if g == nil {
		return nil
	}
	return g.Schema
}

func (g *GetCampaignFoldersResponseLinksItem) GetTargetSchema() *string {
	if g == nil {
		return nil
	}
	return g.TargetSchema
}

func (g *GetCampaignFoldersResponseLinksItem) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g *GetCampaignFoldersResponseLinksItem) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetHref sets the Href field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetCampaignFoldersResponseLinksItem) SetHref(href *string) {
	g.Href = href
	g.require(getCampaignFoldersResponseLinksItemFieldHref)
}

// SetMethod sets the Method field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetCampaignFoldersResponseLinksItem) SetMethod(method *GetCampaignFoldersResponseLinksItemMethod) {
	g.Method = method
	g.require(getCampaignFoldersResponseLinksItemFieldMethod)
}

// SetRel sets the Rel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetCampaignFoldersResponseLinksItem) SetRel(rel *string) {
	g.Rel = rel
	g.require(getCampaignFoldersResponseLinksItemFieldRel)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetCampaignFoldersResponseLinksItem) SetSchema(schema *string) {
	g.Schema = schema
	g.require(getCampaignFoldersResponseLinksItemFieldSchema)
}

// SetTargetSchema sets the TargetSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetCampaignFoldersResponseLinksItem) SetTargetSchema(targetSchema *string) {
	g.TargetSchema = targetSchema
	g.require(getCampaignFoldersResponseLinksItemFieldTargetSchema)
}

func (g *GetCampaignFoldersResponseLinksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler GetCampaignFoldersResponseLinksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetCampaignFoldersResponseLinksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetCampaignFoldersResponseLinksItem) MarshalJSON() ([]byte, error) {
	type embed GetCampaignFoldersResponseLinksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, g.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (g *GetCampaignFoldersResponseLinksItem) String() string {
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
type GetCampaignFoldersResponseLinksItemMethod string

const (
	GetCampaignFoldersResponseLinksItemMethodGet     GetCampaignFoldersResponseLinksItemMethod = "GET"
	GetCampaignFoldersResponseLinksItemMethodPost    GetCampaignFoldersResponseLinksItemMethod = "POST"
	GetCampaignFoldersResponseLinksItemMethodPut     GetCampaignFoldersResponseLinksItemMethod = "PUT"
	GetCampaignFoldersResponseLinksItemMethodPatch   GetCampaignFoldersResponseLinksItemMethod = "PATCH"
	GetCampaignFoldersResponseLinksItemMethodDelete  GetCampaignFoldersResponseLinksItemMethod = "DELETE"
	GetCampaignFoldersResponseLinksItemMethodOptions GetCampaignFoldersResponseLinksItemMethod = "OPTIONS"
	GetCampaignFoldersResponseLinksItemMethodHead    GetCampaignFoldersResponseLinksItemMethod = "HEAD"
)

func NewGetCampaignFoldersResponseLinksItemMethodFromString(s string) (GetCampaignFoldersResponseLinksItemMethod, error) {
	switch s {
	case "GET":
		return GetCampaignFoldersResponseLinksItemMethodGet, nil
	case "POST":
		return GetCampaignFoldersResponseLinksItemMethodPost, nil
	case "PUT":
		return GetCampaignFoldersResponseLinksItemMethodPut, nil
	case "PATCH":
		return GetCampaignFoldersResponseLinksItemMethodPatch, nil
	case "DELETE":
		return GetCampaignFoldersResponseLinksItemMethodDelete, nil
	case "OPTIONS":
		return GetCampaignFoldersResponseLinksItemMethodOptions, nil
	case "HEAD":
		return GetCampaignFoldersResponseLinksItemMethodHead, nil
	}
	var t GetCampaignFoldersResponseLinksItemMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (g GetCampaignFoldersResponseLinksItemMethod) Ptr() *GetCampaignFoldersResponseLinksItemMethod {
	return &g
}

// A folder used to organize campaigns.
var (
	updateCampaignFoldersResponseFieldLinks = big.NewInt(1 << 0)
	updateCampaignFoldersResponseFieldCount = big.NewInt(1 << 1)
	updateCampaignFoldersResponseFieldID    = big.NewInt(1 << 2)
	updateCampaignFoldersResponseFieldName  = big.NewInt(1 << 3)
)

type UpdateCampaignFoldersResponse struct {
	// A list of link types and descriptions for the API schema documents.
	Links []*UpdateCampaignFoldersResponseLinksItem `json:"_links,omitempty" url:"_links,omitempty"`
	// The number of campaigns in the folder.
	Count *int `json:"count,omitempty" url:"count,omitempty"`
	// A string that uniquely identifies this campaign folder.
	ID *string `json:"id,omitempty" url:"id,omitempty"`
	// The name of the folder.
	Name *string `json:"name,omitempty" url:"name,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (u *UpdateCampaignFoldersResponse) GetLinks() []*UpdateCampaignFoldersResponseLinksItem {
	if u == nil {
		return nil
	}
	return u.Links
}

func (u *UpdateCampaignFoldersResponse) GetCount() *int {
	if u == nil {
		return nil
	}
	return u.Count
}

func (u *UpdateCampaignFoldersResponse) GetID() *string {
	if u == nil {
		return nil
	}
	return u.ID
}

func (u *UpdateCampaignFoldersResponse) GetName() *string {
	if u == nil {
		return nil
	}
	return u.Name
}

func (u *UpdateCampaignFoldersResponse) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u *UpdateCampaignFoldersResponse) require(field *big.Int) {
	if u.explicitFields == nil {
		u.explicitFields = big.NewInt(0)
	}
	u.explicitFields.Or(u.explicitFields, field)
}

// SetLinks sets the Links field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (u *UpdateCampaignFoldersResponse) SetLinks(links []*UpdateCampaignFoldersResponseLinksItem) {
	u.Links = links
	u.require(updateCampaignFoldersResponseFieldLinks)
}

// SetCount sets the Count field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (u *UpdateCampaignFoldersResponse) SetCount(count *int) {
	u.Count = count
	u.require(updateCampaignFoldersResponseFieldCount)
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (u *UpdateCampaignFoldersResponse) SetID(id *string) {
	u.ID = id
	u.require(updateCampaignFoldersResponseFieldID)
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (u *UpdateCampaignFoldersResponse) SetName(name *string) {
	u.Name = name
	u.require(updateCampaignFoldersResponseFieldName)
}

func (u *UpdateCampaignFoldersResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateCampaignFoldersResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateCampaignFoldersResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	u.rawJSON = json.RawMessage(data)
	return nil
}

func (u *UpdateCampaignFoldersResponse) MarshalJSON() ([]byte, error) {
	type embed UpdateCampaignFoldersResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, u.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (u *UpdateCampaignFoldersResponse) String() string {
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

// This object represents a link from the resource where it is found to another resource or action that may be performed.
var (
	updateCampaignFoldersResponseLinksItemFieldHref         = big.NewInt(1 << 0)
	updateCampaignFoldersResponseLinksItemFieldMethod       = big.NewInt(1 << 1)
	updateCampaignFoldersResponseLinksItemFieldRel          = big.NewInt(1 << 2)
	updateCampaignFoldersResponseLinksItemFieldSchema       = big.NewInt(1 << 3)
	updateCampaignFoldersResponseLinksItemFieldTargetSchema = big.NewInt(1 << 4)
)

type UpdateCampaignFoldersResponseLinksItem struct {
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href *string `json:"href,omitempty" url:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method *UpdateCampaignFoldersResponseLinksItemMethod `json:"method,omitempty" url:"method,omitempty"`
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

func (u *UpdateCampaignFoldersResponseLinksItem) GetHref() *string {
	if u == nil {
		return nil
	}
	return u.Href
}

func (u *UpdateCampaignFoldersResponseLinksItem) GetMethod() *UpdateCampaignFoldersResponseLinksItemMethod {
	if u == nil {
		return nil
	}
	return u.Method
}

func (u *UpdateCampaignFoldersResponseLinksItem) GetRel() *string {
	if u == nil {
		return nil
	}
	return u.Rel
}

func (u *UpdateCampaignFoldersResponseLinksItem) GetSchema() *string {
	if u == nil {
		return nil
	}
	return u.Schema
}

func (u *UpdateCampaignFoldersResponseLinksItem) GetTargetSchema() *string {
	if u == nil {
		return nil
	}
	return u.TargetSchema
}

func (u *UpdateCampaignFoldersResponseLinksItem) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u *UpdateCampaignFoldersResponseLinksItem) require(field *big.Int) {
	if u.explicitFields == nil {
		u.explicitFields = big.NewInt(0)
	}
	u.explicitFields.Or(u.explicitFields, field)
}

// SetHref sets the Href field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (u *UpdateCampaignFoldersResponseLinksItem) SetHref(href *string) {
	u.Href = href
	u.require(updateCampaignFoldersResponseLinksItemFieldHref)
}

// SetMethod sets the Method field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (u *UpdateCampaignFoldersResponseLinksItem) SetMethod(method *UpdateCampaignFoldersResponseLinksItemMethod) {
	u.Method = method
	u.require(updateCampaignFoldersResponseLinksItemFieldMethod)
}

// SetRel sets the Rel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (u *UpdateCampaignFoldersResponseLinksItem) SetRel(rel *string) {
	u.Rel = rel
	u.require(updateCampaignFoldersResponseLinksItemFieldRel)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (u *UpdateCampaignFoldersResponseLinksItem) SetSchema(schema *string) {
	u.Schema = schema
	u.require(updateCampaignFoldersResponseLinksItemFieldSchema)
}

// SetTargetSchema sets the TargetSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (u *UpdateCampaignFoldersResponseLinksItem) SetTargetSchema(targetSchema *string) {
	u.TargetSchema = targetSchema
	u.require(updateCampaignFoldersResponseLinksItemFieldTargetSchema)
}

func (u *UpdateCampaignFoldersResponseLinksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateCampaignFoldersResponseLinksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateCampaignFoldersResponseLinksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	u.rawJSON = json.RawMessage(data)
	return nil
}

func (u *UpdateCampaignFoldersResponseLinksItem) MarshalJSON() ([]byte, error) {
	type embed UpdateCampaignFoldersResponseLinksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, u.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (u *UpdateCampaignFoldersResponseLinksItem) String() string {
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

// The HTTP method that should be used when accessing the URL defined in 'href'.
type UpdateCampaignFoldersResponseLinksItemMethod string

const (
	UpdateCampaignFoldersResponseLinksItemMethodGet     UpdateCampaignFoldersResponseLinksItemMethod = "GET"
	UpdateCampaignFoldersResponseLinksItemMethodPost    UpdateCampaignFoldersResponseLinksItemMethod = "POST"
	UpdateCampaignFoldersResponseLinksItemMethodPut     UpdateCampaignFoldersResponseLinksItemMethod = "PUT"
	UpdateCampaignFoldersResponseLinksItemMethodPatch   UpdateCampaignFoldersResponseLinksItemMethod = "PATCH"
	UpdateCampaignFoldersResponseLinksItemMethodDelete  UpdateCampaignFoldersResponseLinksItemMethod = "DELETE"
	UpdateCampaignFoldersResponseLinksItemMethodOptions UpdateCampaignFoldersResponseLinksItemMethod = "OPTIONS"
	UpdateCampaignFoldersResponseLinksItemMethodHead    UpdateCampaignFoldersResponseLinksItemMethod = "HEAD"
)

func NewUpdateCampaignFoldersResponseLinksItemMethodFromString(s string) (UpdateCampaignFoldersResponseLinksItemMethod, error) {
	switch s {
	case "GET":
		return UpdateCampaignFoldersResponseLinksItemMethodGet, nil
	case "POST":
		return UpdateCampaignFoldersResponseLinksItemMethodPost, nil
	case "PUT":
		return UpdateCampaignFoldersResponseLinksItemMethodPut, nil
	case "PATCH":
		return UpdateCampaignFoldersResponseLinksItemMethodPatch, nil
	case "DELETE":
		return UpdateCampaignFoldersResponseLinksItemMethodDelete, nil
	case "OPTIONS":
		return UpdateCampaignFoldersResponseLinksItemMethodOptions, nil
	case "HEAD":
		return UpdateCampaignFoldersResponseLinksItemMethodHead, nil
	}
	var t UpdateCampaignFoldersResponseLinksItemMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (u UpdateCampaignFoldersResponseLinksItemMethod) Ptr() *UpdateCampaignFoldersResponseLinksItemMethod {
	return &u
}

var (
	updateCampaignFoldersRequestFieldFolderID = big.NewInt(1 << 0)
	updateCampaignFoldersRequestFieldName     = big.NewInt(1 << 1)
)

type UpdateCampaignFoldersRequest struct {
	// The unique id for the campaign folder.
	FolderID string `json:"-" url:"-"`
	// Name to associate with the folder.
	Name string `json:"name" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (u *UpdateCampaignFoldersRequest) require(field *big.Int) {
	if u.explicitFields == nil {
		u.explicitFields = big.NewInt(0)
	}
	u.explicitFields.Or(u.explicitFields, field)
}

// SetFolderID sets the FolderID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (u *UpdateCampaignFoldersRequest) SetFolderID(folderID string) {
	u.FolderID = folderID
	u.require(updateCampaignFoldersRequestFieldFolderID)
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (u *UpdateCampaignFoldersRequest) SetName(name string) {
	u.Name = name
	u.require(updateCampaignFoldersRequestFieldName)
}

func (u *UpdateCampaignFoldersRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateCampaignFoldersRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*u = UpdateCampaignFoldersRequest(body)
	return nil
}

func (u *UpdateCampaignFoldersRequest) MarshalJSON() ([]byte, error) {
	type embed UpdateCampaignFoldersRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, u.explicitFields)
	return json.Marshal(explicitMarshaler)
}
