// Code generated from our API definition. DO NOT EDIT.

package api

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/mailchimp/mailchimp-marketing-go-sdk/internal"
	big "math/big"
)

var (
	listSearchMembersRequestFieldFields        = big.NewInt(1 << 0)
	listSearchMembersRequestFieldExcludeFields = big.NewInt(1 << 1)
	listSearchMembersRequestFieldQuery         = big.NewInt(1 << 2)
	listSearchMembersRequestFieldListID        = big.NewInt(1 << 3)
)

type ListSearchMembersRequest struct {
	// A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
	Fields []*string `json:"-" url:"fields,omitempty"`
	// A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
	ExcludeFields []*string `json:"-" url:"exclude_fields,omitempty"`
	// The search query used to filter results. Query should be a valid email, or a string representing a contact's first or last name.
	Query string `json:"-" url:"query"`
	// The unique id for the list.
	ListID *string `json:"-" url:"list_id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListSearchMembersRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetFields sets the Fields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSearchMembersRequest) SetFields(fields []*string) {
	l.Fields = fields
	l.require(listSearchMembersRequestFieldFields)
}

// SetExcludeFields sets the ExcludeFields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSearchMembersRequest) SetExcludeFields(excludeFields []*string) {
	l.ExcludeFields = excludeFields
	l.require(listSearchMembersRequestFieldExcludeFields)
}

// SetQuery sets the Query field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSearchMembersRequest) SetQuery(query string) {
	l.Query = query
	l.require(listSearchMembersRequestFieldQuery)
}

// SetListID sets the ListID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSearchMembersRequest) SetListID(listID *string) {
	l.ListID = listID
	l.require(listSearchMembersRequestFieldListID)
}

// Members found for given search term
var (
	listSearchMembersResponseFieldLinks        = big.NewInt(1 << 0)
	listSearchMembersResponseFieldExactMatches = big.NewInt(1 << 1)
	listSearchMembersResponseFieldFullSearch   = big.NewInt(1 << 2)
)

type ListSearchMembersResponse struct {
	// A list of link types and descriptions for the API schema documents.
	Links []*ListSearchMembersResponseLinksItem `json:"_links,omitempty" url:"_links,omitempty"`
	// Exact matches of the provided search query.
	ExactMatches *ListSearchMembersResponseExactMatches `json:"exact_matches,omitempty" url:"exact_matches,omitempty"`
	// Partial matches of the provided search query.
	FullSearch *ListSearchMembersResponseFullSearch `json:"full_search,omitempty" url:"full_search,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListSearchMembersResponse) GetLinks() []*ListSearchMembersResponseLinksItem {
	if l == nil {
		return nil
	}
	return l.Links
}

func (l *ListSearchMembersResponse) GetExactMatches() *ListSearchMembersResponseExactMatches {
	if l == nil {
		return nil
	}
	return l.ExactMatches
}

func (l *ListSearchMembersResponse) GetFullSearch() *ListSearchMembersResponseFullSearch {
	if l == nil {
		return nil
	}
	return l.FullSearch
}

func (l *ListSearchMembersResponse) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListSearchMembersResponse) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetLinks sets the Links field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSearchMembersResponse) SetLinks(links []*ListSearchMembersResponseLinksItem) {
	l.Links = links
	l.require(listSearchMembersResponseFieldLinks)
}

// SetExactMatches sets the ExactMatches field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSearchMembersResponse) SetExactMatches(exactMatches *ListSearchMembersResponseExactMatches) {
	l.ExactMatches = exactMatches
	l.require(listSearchMembersResponseFieldExactMatches)
}

// SetFullSearch sets the FullSearch field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSearchMembersResponse) SetFullSearch(fullSearch *ListSearchMembersResponseFullSearch) {
	l.FullSearch = fullSearch
	l.require(listSearchMembersResponseFieldFullSearch)
}

func (l *ListSearchMembersResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListSearchMembersResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListSearchMembersResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListSearchMembersResponse) MarshalJSON() ([]byte, error) {
	type embed ListSearchMembersResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListSearchMembersResponse) String() string {
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

// Exact matches of the provided search query.
var (
	listSearchMembersResponseExactMatchesFieldMembers    = big.NewInt(1 << 0)
	listSearchMembersResponseExactMatchesFieldTotalItems = big.NewInt(1 << 1)
)

type ListSearchMembersResponseExactMatches struct {
	// An array of objects, each representing a specific list member.
	Members []*ListMembers `json:"members,omitempty" url:"members,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int `json:"total_items,omitempty" url:"total_items,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListSearchMembersResponseExactMatches) GetMembers() []*ListMembers {
	if l == nil {
		return nil
	}
	return l.Members
}

func (l *ListSearchMembersResponseExactMatches) GetTotalItems() *int {
	if l == nil {
		return nil
	}
	return l.TotalItems
}

func (l *ListSearchMembersResponseExactMatches) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListSearchMembersResponseExactMatches) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetMembers sets the Members field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSearchMembersResponseExactMatches) SetMembers(members []*ListMembers) {
	l.Members = members
	l.require(listSearchMembersResponseExactMatchesFieldMembers)
}

// SetTotalItems sets the TotalItems field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSearchMembersResponseExactMatches) SetTotalItems(totalItems *int) {
	l.TotalItems = totalItems
	l.require(listSearchMembersResponseExactMatchesFieldTotalItems)
}

func (l *ListSearchMembersResponseExactMatches) UnmarshalJSON(data []byte) error {
	type unmarshaler ListSearchMembersResponseExactMatches
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListSearchMembersResponseExactMatches(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListSearchMembersResponseExactMatches) MarshalJSON() ([]byte, error) {
	type embed ListSearchMembersResponseExactMatches
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListSearchMembersResponseExactMatches) String() string {
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

// Partial matches of the provided search query.
var (
	listSearchMembersResponseFullSearchFieldMembers    = big.NewInt(1 << 0)
	listSearchMembersResponseFullSearchFieldTotalItems = big.NewInt(1 << 1)
)

type ListSearchMembersResponseFullSearch struct {
	// An array of objects, each representing a specific list member.
	Members []*ListMembers `json:"members,omitempty" url:"members,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int `json:"total_items,omitempty" url:"total_items,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListSearchMembersResponseFullSearch) GetMembers() []*ListMembers {
	if l == nil {
		return nil
	}
	return l.Members
}

func (l *ListSearchMembersResponseFullSearch) GetTotalItems() *int {
	if l == nil {
		return nil
	}
	return l.TotalItems
}

func (l *ListSearchMembersResponseFullSearch) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListSearchMembersResponseFullSearch) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetMembers sets the Members field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSearchMembersResponseFullSearch) SetMembers(members []*ListMembers) {
	l.Members = members
	l.require(listSearchMembersResponseFullSearchFieldMembers)
}

// SetTotalItems sets the TotalItems field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSearchMembersResponseFullSearch) SetTotalItems(totalItems *int) {
	l.TotalItems = totalItems
	l.require(listSearchMembersResponseFullSearchFieldTotalItems)
}

func (l *ListSearchMembersResponseFullSearch) UnmarshalJSON(data []byte) error {
	type unmarshaler ListSearchMembersResponseFullSearch
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListSearchMembersResponseFullSearch(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListSearchMembersResponseFullSearch) MarshalJSON() ([]byte, error) {
	type embed ListSearchMembersResponseFullSearch
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListSearchMembersResponseFullSearch) String() string {
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
	listSearchMembersResponseLinksItemFieldHref         = big.NewInt(1 << 0)
	listSearchMembersResponseLinksItemFieldMethod       = big.NewInt(1 << 1)
	listSearchMembersResponseLinksItemFieldRel          = big.NewInt(1 << 2)
	listSearchMembersResponseLinksItemFieldSchema       = big.NewInt(1 << 3)
	listSearchMembersResponseLinksItemFieldTargetSchema = big.NewInt(1 << 4)
)

type ListSearchMembersResponseLinksItem struct {
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href *string `json:"href,omitempty" url:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method *ListSearchMembersResponseLinksItemMethod `json:"method,omitempty" url:"method,omitempty"`
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

func (l *ListSearchMembersResponseLinksItem) GetHref() *string {
	if l == nil {
		return nil
	}
	return l.Href
}

func (l *ListSearchMembersResponseLinksItem) GetMethod() *ListSearchMembersResponseLinksItemMethod {
	if l == nil {
		return nil
	}
	return l.Method
}

func (l *ListSearchMembersResponseLinksItem) GetRel() *string {
	if l == nil {
		return nil
	}
	return l.Rel
}

func (l *ListSearchMembersResponseLinksItem) GetSchema() *string {
	if l == nil {
		return nil
	}
	return l.Schema
}

func (l *ListSearchMembersResponseLinksItem) GetTargetSchema() *string {
	if l == nil {
		return nil
	}
	return l.TargetSchema
}

func (l *ListSearchMembersResponseLinksItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListSearchMembersResponseLinksItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetHref sets the Href field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSearchMembersResponseLinksItem) SetHref(href *string) {
	l.Href = href
	l.require(listSearchMembersResponseLinksItemFieldHref)
}

// SetMethod sets the Method field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSearchMembersResponseLinksItem) SetMethod(method *ListSearchMembersResponseLinksItemMethod) {
	l.Method = method
	l.require(listSearchMembersResponseLinksItemFieldMethod)
}

// SetRel sets the Rel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSearchMembersResponseLinksItem) SetRel(rel *string) {
	l.Rel = rel
	l.require(listSearchMembersResponseLinksItemFieldRel)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSearchMembersResponseLinksItem) SetSchema(schema *string) {
	l.Schema = schema
	l.require(listSearchMembersResponseLinksItemFieldSchema)
}

// SetTargetSchema sets the TargetSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListSearchMembersResponseLinksItem) SetTargetSchema(targetSchema *string) {
	l.TargetSchema = targetSchema
	l.require(listSearchMembersResponseLinksItemFieldTargetSchema)
}

func (l *ListSearchMembersResponseLinksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler ListSearchMembersResponseLinksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListSearchMembersResponseLinksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListSearchMembersResponseLinksItem) MarshalJSON() ([]byte, error) {
	type embed ListSearchMembersResponseLinksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListSearchMembersResponseLinksItem) String() string {
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
type ListSearchMembersResponseLinksItemMethod string

const (
	ListSearchMembersResponseLinksItemMethodGet     ListSearchMembersResponseLinksItemMethod = "GET"
	ListSearchMembersResponseLinksItemMethodPost    ListSearchMembersResponseLinksItemMethod = "POST"
	ListSearchMembersResponseLinksItemMethodPut     ListSearchMembersResponseLinksItemMethod = "PUT"
	ListSearchMembersResponseLinksItemMethodPatch   ListSearchMembersResponseLinksItemMethod = "PATCH"
	ListSearchMembersResponseLinksItemMethodDelete  ListSearchMembersResponseLinksItemMethod = "DELETE"
	ListSearchMembersResponseLinksItemMethodOptions ListSearchMembersResponseLinksItemMethod = "OPTIONS"
	ListSearchMembersResponseLinksItemMethodHead    ListSearchMembersResponseLinksItemMethod = "HEAD"
)

func NewListSearchMembersResponseLinksItemMethodFromString(s string) (ListSearchMembersResponseLinksItemMethod, error) {
	switch s {
	case "GET":
		return ListSearchMembersResponseLinksItemMethodGet, nil
	case "POST":
		return ListSearchMembersResponseLinksItemMethodPost, nil
	case "PUT":
		return ListSearchMembersResponseLinksItemMethodPut, nil
	case "PATCH":
		return ListSearchMembersResponseLinksItemMethodPatch, nil
	case "DELETE":
		return ListSearchMembersResponseLinksItemMethodDelete, nil
	case "OPTIONS":
		return ListSearchMembersResponseLinksItemMethodOptions, nil
	case "HEAD":
		return ListSearchMembersResponseLinksItemMethodHead, nil
	}
	var t ListSearchMembersResponseLinksItemMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListSearchMembersResponseLinksItemMethod) Ptr() *ListSearchMembersResponseLinksItemMethod {
	return &l
}
