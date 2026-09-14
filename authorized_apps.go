// Code generated from our API definition. DO NOT EDIT.

package api

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/mailchimp/mailchimp-marketing-go-sdk/internal"
	big "math/big"
)

var (
	getAuthorizedAppsRequestFieldAppID         = big.NewInt(1 << 0)
	getAuthorizedAppsRequestFieldFields        = big.NewInt(1 << 1)
	getAuthorizedAppsRequestFieldExcludeFields = big.NewInt(1 << 2)
)

type GetAuthorizedAppsRequest struct {
	// The unique id for the connected authorized application.
	AppID string `json:"-" url:"-"`
	// A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
	Fields []*string `json:"-" url:"fields,omitempty"`
	// A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
	ExcludeFields []*string `json:"-" url:"exclude_fields,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetAuthorizedAppsRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetAppID sets the AppID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAuthorizedAppsRequest) SetAppID(appID string) {
	g.AppID = appID
	g.require(getAuthorizedAppsRequestFieldAppID)
}

// SetFields sets the Fields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAuthorizedAppsRequest) SetFields(fields []*string) {
	g.Fields = fields
	g.require(getAuthorizedAppsRequestFieldFields)
}

// SetExcludeFields sets the ExcludeFields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAuthorizedAppsRequest) SetExcludeFields(excludeFields []*string) {
	g.ExcludeFields = excludeFields
	g.require(getAuthorizedAppsRequestFieldExcludeFields)
}

var (
	listAuthorizedAppsRequestFieldFields        = big.NewInt(1 << 0)
	listAuthorizedAppsRequestFieldExcludeFields = big.NewInt(1 << 1)
	listAuthorizedAppsRequestFieldCount         = big.NewInt(1 << 2)
	listAuthorizedAppsRequestFieldOffset        = big.NewInt(1 << 3)
)

type ListAuthorizedAppsRequest struct {
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

func (l *ListAuthorizedAppsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetFields sets the Fields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAuthorizedAppsRequest) SetFields(fields []*string) {
	l.Fields = fields
	l.require(listAuthorizedAppsRequestFieldFields)
}

// SetExcludeFields sets the ExcludeFields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAuthorizedAppsRequest) SetExcludeFields(excludeFields []*string) {
	l.ExcludeFields = excludeFields
	l.require(listAuthorizedAppsRequestFieldExcludeFields)
}

// SetCount sets the Count field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAuthorizedAppsRequest) SetCount(count *int) {
	l.Count = count
	l.require(listAuthorizedAppsRequestFieldCount)
}

// SetOffset sets the Offset field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAuthorizedAppsRequest) SetOffset(offset *int) {
	l.Offset = offset
	l.require(listAuthorizedAppsRequestFieldOffset)
}

// An authorized app.
var (
	getAuthorizedAppsResponseFieldLinks       = big.NewInt(1 << 0)
	getAuthorizedAppsResponseFieldDescription = big.NewInt(1 << 1)
	getAuthorizedAppsResponseFieldID          = big.NewInt(1 << 2)
	getAuthorizedAppsResponseFieldName        = big.NewInt(1 << 3)
	getAuthorizedAppsResponseFieldUsers       = big.NewInt(1 << 4)
)

type GetAuthorizedAppsResponse struct {
	// A list of link types and descriptions for the API schema documents.
	Links []*GetAuthorizedAppsResponseLinksItem `json:"_links,omitempty" url:"_links,omitempty"`
	// A short description of the application.
	Description *string `json:"description,omitempty" url:"description,omitempty"`
	// The ID for the application.
	ID *int `json:"id,omitempty" url:"id,omitempty"`
	// The name of the application.
	Name *string `json:"name,omitempty" url:"name,omitempty"`
	// An array of usernames for users who have linked the app.
	Users []string `json:"users,omitempty" url:"users,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (g *GetAuthorizedAppsResponse) GetLinks() []*GetAuthorizedAppsResponseLinksItem {
	if g == nil {
		return nil
	}
	return g.Links
}

func (g *GetAuthorizedAppsResponse) GetDescription() *string {
	if g == nil {
		return nil
	}
	return g.Description
}

func (g *GetAuthorizedAppsResponse) GetID() *int {
	if g == nil {
		return nil
	}
	return g.ID
}

func (g *GetAuthorizedAppsResponse) GetName() *string {
	if g == nil {
		return nil
	}
	return g.Name
}

func (g *GetAuthorizedAppsResponse) GetUsers() []string {
	if g == nil {
		return nil
	}
	return g.Users
}

func (g *GetAuthorizedAppsResponse) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g *GetAuthorizedAppsResponse) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetLinks sets the Links field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAuthorizedAppsResponse) SetLinks(links []*GetAuthorizedAppsResponseLinksItem) {
	g.Links = links
	g.require(getAuthorizedAppsResponseFieldLinks)
}

// SetDescription sets the Description field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAuthorizedAppsResponse) SetDescription(description *string) {
	g.Description = description
	g.require(getAuthorizedAppsResponseFieldDescription)
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAuthorizedAppsResponse) SetID(id *int) {
	g.ID = id
	g.require(getAuthorizedAppsResponseFieldID)
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAuthorizedAppsResponse) SetName(name *string) {
	g.Name = name
	g.require(getAuthorizedAppsResponseFieldName)
}

// SetUsers sets the Users field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAuthorizedAppsResponse) SetUsers(users []string) {
	g.Users = users
	g.require(getAuthorizedAppsResponseFieldUsers)
}

func (g *GetAuthorizedAppsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetAuthorizedAppsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetAuthorizedAppsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetAuthorizedAppsResponse) MarshalJSON() ([]byte, error) {
	type embed GetAuthorizedAppsResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, g.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (g *GetAuthorizedAppsResponse) String() string {
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
	getAuthorizedAppsResponseLinksItemFieldHref         = big.NewInt(1 << 0)
	getAuthorizedAppsResponseLinksItemFieldMethod       = big.NewInt(1 << 1)
	getAuthorizedAppsResponseLinksItemFieldRel          = big.NewInt(1 << 2)
	getAuthorizedAppsResponseLinksItemFieldSchema       = big.NewInt(1 << 3)
	getAuthorizedAppsResponseLinksItemFieldTargetSchema = big.NewInt(1 << 4)
)

type GetAuthorizedAppsResponseLinksItem struct {
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href *string `json:"href,omitempty" url:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method *GetAuthorizedAppsResponseLinksItemMethod `json:"method,omitempty" url:"method,omitempty"`
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

func (g *GetAuthorizedAppsResponseLinksItem) GetHref() *string {
	if g == nil {
		return nil
	}
	return g.Href
}

func (g *GetAuthorizedAppsResponseLinksItem) GetMethod() *GetAuthorizedAppsResponseLinksItemMethod {
	if g == nil {
		return nil
	}
	return g.Method
}

func (g *GetAuthorizedAppsResponseLinksItem) GetRel() *string {
	if g == nil {
		return nil
	}
	return g.Rel
}

func (g *GetAuthorizedAppsResponseLinksItem) GetSchema() *string {
	if g == nil {
		return nil
	}
	return g.Schema
}

func (g *GetAuthorizedAppsResponseLinksItem) GetTargetSchema() *string {
	if g == nil {
		return nil
	}
	return g.TargetSchema
}

func (g *GetAuthorizedAppsResponseLinksItem) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g *GetAuthorizedAppsResponseLinksItem) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetHref sets the Href field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAuthorizedAppsResponseLinksItem) SetHref(href *string) {
	g.Href = href
	g.require(getAuthorizedAppsResponseLinksItemFieldHref)
}

// SetMethod sets the Method field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAuthorizedAppsResponseLinksItem) SetMethod(method *GetAuthorizedAppsResponseLinksItemMethod) {
	g.Method = method
	g.require(getAuthorizedAppsResponseLinksItemFieldMethod)
}

// SetRel sets the Rel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAuthorizedAppsResponseLinksItem) SetRel(rel *string) {
	g.Rel = rel
	g.require(getAuthorizedAppsResponseLinksItemFieldRel)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAuthorizedAppsResponseLinksItem) SetSchema(schema *string) {
	g.Schema = schema
	g.require(getAuthorizedAppsResponseLinksItemFieldSchema)
}

// SetTargetSchema sets the TargetSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAuthorizedAppsResponseLinksItem) SetTargetSchema(targetSchema *string) {
	g.TargetSchema = targetSchema
	g.require(getAuthorizedAppsResponseLinksItemFieldTargetSchema)
}

func (g *GetAuthorizedAppsResponseLinksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler GetAuthorizedAppsResponseLinksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetAuthorizedAppsResponseLinksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetAuthorizedAppsResponseLinksItem) MarshalJSON() ([]byte, error) {
	type embed GetAuthorizedAppsResponseLinksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, g.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (g *GetAuthorizedAppsResponseLinksItem) String() string {
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
type GetAuthorizedAppsResponseLinksItemMethod string

const (
	GetAuthorizedAppsResponseLinksItemMethodGet     GetAuthorizedAppsResponseLinksItemMethod = "GET"
	GetAuthorizedAppsResponseLinksItemMethodPost    GetAuthorizedAppsResponseLinksItemMethod = "POST"
	GetAuthorizedAppsResponseLinksItemMethodPut     GetAuthorizedAppsResponseLinksItemMethod = "PUT"
	GetAuthorizedAppsResponseLinksItemMethodPatch   GetAuthorizedAppsResponseLinksItemMethod = "PATCH"
	GetAuthorizedAppsResponseLinksItemMethodDelete  GetAuthorizedAppsResponseLinksItemMethod = "DELETE"
	GetAuthorizedAppsResponseLinksItemMethodOptions GetAuthorizedAppsResponseLinksItemMethod = "OPTIONS"
	GetAuthorizedAppsResponseLinksItemMethodHead    GetAuthorizedAppsResponseLinksItemMethod = "HEAD"
)

func NewGetAuthorizedAppsResponseLinksItemMethodFromString(s string) (GetAuthorizedAppsResponseLinksItemMethod, error) {
	switch s {
	case "GET":
		return GetAuthorizedAppsResponseLinksItemMethodGet, nil
	case "POST":
		return GetAuthorizedAppsResponseLinksItemMethodPost, nil
	case "PUT":
		return GetAuthorizedAppsResponseLinksItemMethodPut, nil
	case "PATCH":
		return GetAuthorizedAppsResponseLinksItemMethodPatch, nil
	case "DELETE":
		return GetAuthorizedAppsResponseLinksItemMethodDelete, nil
	case "OPTIONS":
		return GetAuthorizedAppsResponseLinksItemMethodOptions, nil
	case "HEAD":
		return GetAuthorizedAppsResponseLinksItemMethodHead, nil
	}
	var t GetAuthorizedAppsResponseLinksItemMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (g GetAuthorizedAppsResponseLinksItemMethod) Ptr() *GetAuthorizedAppsResponseLinksItemMethod {
	return &g
}

// An array of objects, each representing an authorized application.
var (
	listAuthorizedAppsResponseFieldLinks      = big.NewInt(1 << 0)
	listAuthorizedAppsResponseFieldApps       = big.NewInt(1 << 1)
	listAuthorizedAppsResponseFieldTotalItems = big.NewInt(1 << 2)
)

type ListAuthorizedAppsResponse struct {
	// A list of link types and descriptions for the API schema documents.
	Links []*ListAuthorizedAppsResponseLinksItem `json:"_links,omitempty" url:"_links,omitempty"`
	// An array of objects, each representing an authorized application.
	Apps []*ListAuthorizedAppsResponseAppsItem `json:"apps,omitempty" url:"apps,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int `json:"total_items,omitempty" url:"total_items,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListAuthorizedAppsResponse) GetLinks() []*ListAuthorizedAppsResponseLinksItem {
	if l == nil {
		return nil
	}
	return l.Links
}

func (l *ListAuthorizedAppsResponse) GetApps() []*ListAuthorizedAppsResponseAppsItem {
	if l == nil {
		return nil
	}
	return l.Apps
}

func (l *ListAuthorizedAppsResponse) GetTotalItems() *int {
	if l == nil {
		return nil
	}
	return l.TotalItems
}

func (l *ListAuthorizedAppsResponse) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListAuthorizedAppsResponse) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetLinks sets the Links field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAuthorizedAppsResponse) SetLinks(links []*ListAuthorizedAppsResponseLinksItem) {
	l.Links = links
	l.require(listAuthorizedAppsResponseFieldLinks)
}

// SetApps sets the Apps field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAuthorizedAppsResponse) SetApps(apps []*ListAuthorizedAppsResponseAppsItem) {
	l.Apps = apps
	l.require(listAuthorizedAppsResponseFieldApps)
}

// SetTotalItems sets the TotalItems field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAuthorizedAppsResponse) SetTotalItems(totalItems *int) {
	l.TotalItems = totalItems
	l.require(listAuthorizedAppsResponseFieldTotalItems)
}

func (l *ListAuthorizedAppsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListAuthorizedAppsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListAuthorizedAppsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListAuthorizedAppsResponse) MarshalJSON() ([]byte, error) {
	type embed ListAuthorizedAppsResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListAuthorizedAppsResponse) String() string {
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

// An authorized app.
var (
	listAuthorizedAppsResponseAppsItemFieldLinks       = big.NewInt(1 << 0)
	listAuthorizedAppsResponseAppsItemFieldDescription = big.NewInt(1 << 1)
	listAuthorizedAppsResponseAppsItemFieldID          = big.NewInt(1 << 2)
	listAuthorizedAppsResponseAppsItemFieldName        = big.NewInt(1 << 3)
	listAuthorizedAppsResponseAppsItemFieldUsers       = big.NewInt(1 << 4)
)

type ListAuthorizedAppsResponseAppsItem struct {
	// A list of link types and descriptions for the API schema documents.
	Links []*ListAuthorizedAppsResponseAppsItemLinksItem `json:"_links,omitempty" url:"_links,omitempty"`
	// A short description of the application.
	Description *string `json:"description,omitempty" url:"description,omitempty"`
	// The ID for the application.
	ID *int `json:"id,omitempty" url:"id,omitempty"`
	// The name of the application.
	Name *string `json:"name,omitempty" url:"name,omitempty"`
	// An array of usernames for users who have linked the app.
	Users []string `json:"users,omitempty" url:"users,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListAuthorizedAppsResponseAppsItem) GetLinks() []*ListAuthorizedAppsResponseAppsItemLinksItem {
	if l == nil {
		return nil
	}
	return l.Links
}

func (l *ListAuthorizedAppsResponseAppsItem) GetDescription() *string {
	if l == nil {
		return nil
	}
	return l.Description
}

func (l *ListAuthorizedAppsResponseAppsItem) GetID() *int {
	if l == nil {
		return nil
	}
	return l.ID
}

func (l *ListAuthorizedAppsResponseAppsItem) GetName() *string {
	if l == nil {
		return nil
	}
	return l.Name
}

func (l *ListAuthorizedAppsResponseAppsItem) GetUsers() []string {
	if l == nil {
		return nil
	}
	return l.Users
}

func (l *ListAuthorizedAppsResponseAppsItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListAuthorizedAppsResponseAppsItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetLinks sets the Links field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAuthorizedAppsResponseAppsItem) SetLinks(links []*ListAuthorizedAppsResponseAppsItemLinksItem) {
	l.Links = links
	l.require(listAuthorizedAppsResponseAppsItemFieldLinks)
}

// SetDescription sets the Description field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAuthorizedAppsResponseAppsItem) SetDescription(description *string) {
	l.Description = description
	l.require(listAuthorizedAppsResponseAppsItemFieldDescription)
}

// SetID sets the ID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAuthorizedAppsResponseAppsItem) SetID(id *int) {
	l.ID = id
	l.require(listAuthorizedAppsResponseAppsItemFieldID)
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAuthorizedAppsResponseAppsItem) SetName(name *string) {
	l.Name = name
	l.require(listAuthorizedAppsResponseAppsItemFieldName)
}

// SetUsers sets the Users field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAuthorizedAppsResponseAppsItem) SetUsers(users []string) {
	l.Users = users
	l.require(listAuthorizedAppsResponseAppsItemFieldUsers)
}

func (l *ListAuthorizedAppsResponseAppsItem) UnmarshalJSON(data []byte) error {
	type unmarshaler ListAuthorizedAppsResponseAppsItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListAuthorizedAppsResponseAppsItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListAuthorizedAppsResponseAppsItem) MarshalJSON() ([]byte, error) {
	type embed ListAuthorizedAppsResponseAppsItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListAuthorizedAppsResponseAppsItem) String() string {
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
	listAuthorizedAppsResponseAppsItemLinksItemFieldHref         = big.NewInt(1 << 0)
	listAuthorizedAppsResponseAppsItemLinksItemFieldMethod       = big.NewInt(1 << 1)
	listAuthorizedAppsResponseAppsItemLinksItemFieldRel          = big.NewInt(1 << 2)
	listAuthorizedAppsResponseAppsItemLinksItemFieldSchema       = big.NewInt(1 << 3)
	listAuthorizedAppsResponseAppsItemLinksItemFieldTargetSchema = big.NewInt(1 << 4)
)

type ListAuthorizedAppsResponseAppsItemLinksItem struct {
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href *string `json:"href,omitempty" url:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method *ListAuthorizedAppsResponseAppsItemLinksItemMethod `json:"method,omitempty" url:"method,omitempty"`
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

func (l *ListAuthorizedAppsResponseAppsItemLinksItem) GetHref() *string {
	if l == nil {
		return nil
	}
	return l.Href
}

func (l *ListAuthorizedAppsResponseAppsItemLinksItem) GetMethod() *ListAuthorizedAppsResponseAppsItemLinksItemMethod {
	if l == nil {
		return nil
	}
	return l.Method
}

func (l *ListAuthorizedAppsResponseAppsItemLinksItem) GetRel() *string {
	if l == nil {
		return nil
	}
	return l.Rel
}

func (l *ListAuthorizedAppsResponseAppsItemLinksItem) GetSchema() *string {
	if l == nil {
		return nil
	}
	return l.Schema
}

func (l *ListAuthorizedAppsResponseAppsItemLinksItem) GetTargetSchema() *string {
	if l == nil {
		return nil
	}
	return l.TargetSchema
}

func (l *ListAuthorizedAppsResponseAppsItemLinksItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListAuthorizedAppsResponseAppsItemLinksItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetHref sets the Href field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAuthorizedAppsResponseAppsItemLinksItem) SetHref(href *string) {
	l.Href = href
	l.require(listAuthorizedAppsResponseAppsItemLinksItemFieldHref)
}

// SetMethod sets the Method field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAuthorizedAppsResponseAppsItemLinksItem) SetMethod(method *ListAuthorizedAppsResponseAppsItemLinksItemMethod) {
	l.Method = method
	l.require(listAuthorizedAppsResponseAppsItemLinksItemFieldMethod)
}

// SetRel sets the Rel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAuthorizedAppsResponseAppsItemLinksItem) SetRel(rel *string) {
	l.Rel = rel
	l.require(listAuthorizedAppsResponseAppsItemLinksItemFieldRel)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAuthorizedAppsResponseAppsItemLinksItem) SetSchema(schema *string) {
	l.Schema = schema
	l.require(listAuthorizedAppsResponseAppsItemLinksItemFieldSchema)
}

// SetTargetSchema sets the TargetSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAuthorizedAppsResponseAppsItemLinksItem) SetTargetSchema(targetSchema *string) {
	l.TargetSchema = targetSchema
	l.require(listAuthorizedAppsResponseAppsItemLinksItemFieldTargetSchema)
}

func (l *ListAuthorizedAppsResponseAppsItemLinksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler ListAuthorizedAppsResponseAppsItemLinksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListAuthorizedAppsResponseAppsItemLinksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListAuthorizedAppsResponseAppsItemLinksItem) MarshalJSON() ([]byte, error) {
	type embed ListAuthorizedAppsResponseAppsItemLinksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListAuthorizedAppsResponseAppsItemLinksItem) String() string {
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
type ListAuthorizedAppsResponseAppsItemLinksItemMethod string

const (
	ListAuthorizedAppsResponseAppsItemLinksItemMethodGet     ListAuthorizedAppsResponseAppsItemLinksItemMethod = "GET"
	ListAuthorizedAppsResponseAppsItemLinksItemMethodPost    ListAuthorizedAppsResponseAppsItemLinksItemMethod = "POST"
	ListAuthorizedAppsResponseAppsItemLinksItemMethodPut     ListAuthorizedAppsResponseAppsItemLinksItemMethod = "PUT"
	ListAuthorizedAppsResponseAppsItemLinksItemMethodPatch   ListAuthorizedAppsResponseAppsItemLinksItemMethod = "PATCH"
	ListAuthorizedAppsResponseAppsItemLinksItemMethodDelete  ListAuthorizedAppsResponseAppsItemLinksItemMethod = "DELETE"
	ListAuthorizedAppsResponseAppsItemLinksItemMethodOptions ListAuthorizedAppsResponseAppsItemLinksItemMethod = "OPTIONS"
	ListAuthorizedAppsResponseAppsItemLinksItemMethodHead    ListAuthorizedAppsResponseAppsItemLinksItemMethod = "HEAD"
)

func NewListAuthorizedAppsResponseAppsItemLinksItemMethodFromString(s string) (ListAuthorizedAppsResponseAppsItemLinksItemMethod, error) {
	switch s {
	case "GET":
		return ListAuthorizedAppsResponseAppsItemLinksItemMethodGet, nil
	case "POST":
		return ListAuthorizedAppsResponseAppsItemLinksItemMethodPost, nil
	case "PUT":
		return ListAuthorizedAppsResponseAppsItemLinksItemMethodPut, nil
	case "PATCH":
		return ListAuthorizedAppsResponseAppsItemLinksItemMethodPatch, nil
	case "DELETE":
		return ListAuthorizedAppsResponseAppsItemLinksItemMethodDelete, nil
	case "OPTIONS":
		return ListAuthorizedAppsResponseAppsItemLinksItemMethodOptions, nil
	case "HEAD":
		return ListAuthorizedAppsResponseAppsItemLinksItemMethodHead, nil
	}
	var t ListAuthorizedAppsResponseAppsItemLinksItemMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListAuthorizedAppsResponseAppsItemLinksItemMethod) Ptr() *ListAuthorizedAppsResponseAppsItemLinksItemMethod {
	return &l
}

// This object represents a link from the resource where it is found to another resource or action that may be performed.
var (
	listAuthorizedAppsResponseLinksItemFieldHref         = big.NewInt(1 << 0)
	listAuthorizedAppsResponseLinksItemFieldMethod       = big.NewInt(1 << 1)
	listAuthorizedAppsResponseLinksItemFieldRel          = big.NewInt(1 << 2)
	listAuthorizedAppsResponseLinksItemFieldSchema       = big.NewInt(1 << 3)
	listAuthorizedAppsResponseLinksItemFieldTargetSchema = big.NewInt(1 << 4)
)

type ListAuthorizedAppsResponseLinksItem struct {
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href *string `json:"href,omitempty" url:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method *ListAuthorizedAppsResponseLinksItemMethod `json:"method,omitempty" url:"method,omitempty"`
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

func (l *ListAuthorizedAppsResponseLinksItem) GetHref() *string {
	if l == nil {
		return nil
	}
	return l.Href
}

func (l *ListAuthorizedAppsResponseLinksItem) GetMethod() *ListAuthorizedAppsResponseLinksItemMethod {
	if l == nil {
		return nil
	}
	return l.Method
}

func (l *ListAuthorizedAppsResponseLinksItem) GetRel() *string {
	if l == nil {
		return nil
	}
	return l.Rel
}

func (l *ListAuthorizedAppsResponseLinksItem) GetSchema() *string {
	if l == nil {
		return nil
	}
	return l.Schema
}

func (l *ListAuthorizedAppsResponseLinksItem) GetTargetSchema() *string {
	if l == nil {
		return nil
	}
	return l.TargetSchema
}

func (l *ListAuthorizedAppsResponseLinksItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListAuthorizedAppsResponseLinksItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetHref sets the Href field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAuthorizedAppsResponseLinksItem) SetHref(href *string) {
	l.Href = href
	l.require(listAuthorizedAppsResponseLinksItemFieldHref)
}

// SetMethod sets the Method field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAuthorizedAppsResponseLinksItem) SetMethod(method *ListAuthorizedAppsResponseLinksItemMethod) {
	l.Method = method
	l.require(listAuthorizedAppsResponseLinksItemFieldMethod)
}

// SetRel sets the Rel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAuthorizedAppsResponseLinksItem) SetRel(rel *string) {
	l.Rel = rel
	l.require(listAuthorizedAppsResponseLinksItemFieldRel)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAuthorizedAppsResponseLinksItem) SetSchema(schema *string) {
	l.Schema = schema
	l.require(listAuthorizedAppsResponseLinksItemFieldSchema)
}

// SetTargetSchema sets the TargetSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAuthorizedAppsResponseLinksItem) SetTargetSchema(targetSchema *string) {
	l.TargetSchema = targetSchema
	l.require(listAuthorizedAppsResponseLinksItemFieldTargetSchema)
}

func (l *ListAuthorizedAppsResponseLinksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler ListAuthorizedAppsResponseLinksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListAuthorizedAppsResponseLinksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListAuthorizedAppsResponseLinksItem) MarshalJSON() ([]byte, error) {
	type embed ListAuthorizedAppsResponseLinksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListAuthorizedAppsResponseLinksItem) String() string {
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
type ListAuthorizedAppsResponseLinksItemMethod string

const (
	ListAuthorizedAppsResponseLinksItemMethodGet     ListAuthorizedAppsResponseLinksItemMethod = "GET"
	ListAuthorizedAppsResponseLinksItemMethodPost    ListAuthorizedAppsResponseLinksItemMethod = "POST"
	ListAuthorizedAppsResponseLinksItemMethodPut     ListAuthorizedAppsResponseLinksItemMethod = "PUT"
	ListAuthorizedAppsResponseLinksItemMethodPatch   ListAuthorizedAppsResponseLinksItemMethod = "PATCH"
	ListAuthorizedAppsResponseLinksItemMethodDelete  ListAuthorizedAppsResponseLinksItemMethod = "DELETE"
	ListAuthorizedAppsResponseLinksItemMethodOptions ListAuthorizedAppsResponseLinksItemMethod = "OPTIONS"
	ListAuthorizedAppsResponseLinksItemMethodHead    ListAuthorizedAppsResponseLinksItemMethod = "HEAD"
)

func NewListAuthorizedAppsResponseLinksItemMethodFromString(s string) (ListAuthorizedAppsResponseLinksItemMethod, error) {
	switch s {
	case "GET":
		return ListAuthorizedAppsResponseLinksItemMethodGet, nil
	case "POST":
		return ListAuthorizedAppsResponseLinksItemMethodPost, nil
	case "PUT":
		return ListAuthorizedAppsResponseLinksItemMethodPut, nil
	case "PATCH":
		return ListAuthorizedAppsResponseLinksItemMethodPatch, nil
	case "DELETE":
		return ListAuthorizedAppsResponseLinksItemMethodDelete, nil
	case "OPTIONS":
		return ListAuthorizedAppsResponseLinksItemMethodOptions, nil
	case "HEAD":
		return ListAuthorizedAppsResponseLinksItemMethodHead, nil
	}
	var t ListAuthorizedAppsResponseLinksItemMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListAuthorizedAppsResponseLinksItemMethod) Ptr() *ListAuthorizedAppsResponseLinksItemMethod {
	return &l
}
