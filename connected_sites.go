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
	createConnectedSitesRequestFieldDomain    = big.NewInt(1 << 0)
	createConnectedSitesRequestFieldForeignID = big.NewInt(1 << 1)
)

type CreateConnectedSitesRequest struct {
	// The connected site domain.
	Domain string `json:"domain" url:"-"`
	// The unique identifier for the site.
	ForeignID string `json:"foreign_id" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CreateConnectedSitesRequest) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetDomain sets the Domain field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateConnectedSitesRequest) SetDomain(domain string) {
	c.Domain = domain
	c.require(createConnectedSitesRequestFieldDomain)
}

// SetForeignID sets the ForeignID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateConnectedSitesRequest) SetForeignID(foreignID string) {
	c.ForeignID = foreignID
	c.require(createConnectedSitesRequestFieldForeignID)
}

func (c *CreateConnectedSitesRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateConnectedSitesRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*c = CreateConnectedSitesRequest(body)
	return nil
}

func (c *CreateConnectedSitesRequest) MarshalJSON() ([]byte, error) {
	type embed CreateConnectedSitesRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

var (
	createActionVerifyScriptInstallationConnectedSitesRequestFieldConnectedSiteID = big.NewInt(1 << 0)
)

type CreateActionVerifyScriptInstallationConnectedSitesRequest struct {
	// The unique identifier for the site.
	ConnectedSiteID string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CreateActionVerifyScriptInstallationConnectedSitesRequest) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetConnectedSiteID sets the ConnectedSiteID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateActionVerifyScriptInstallationConnectedSitesRequest) SetConnectedSiteID(connectedSiteID string) {
	c.ConnectedSiteID = connectedSiteID
	c.require(createActionVerifyScriptInstallationConnectedSitesRequestFieldConnectedSiteID)
}

var (
	deleteConnectedSitesRequestFieldConnectedSiteID = big.NewInt(1 << 0)
)

type DeleteConnectedSitesRequest struct {
	// The unique identifier for the site.
	ConnectedSiteID string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (d *DeleteConnectedSitesRequest) require(field *big.Int) {
	if d.explicitFields == nil {
		d.explicitFields = big.NewInt(0)
	}
	d.explicitFields.Or(d.explicitFields, field)
}

// SetConnectedSiteID sets the ConnectedSiteID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (d *DeleteConnectedSitesRequest) SetConnectedSiteID(connectedSiteID string) {
	d.ConnectedSiteID = connectedSiteID
	d.require(deleteConnectedSitesRequestFieldConnectedSiteID)
}

var (
	getConnectedSitesRequestFieldConnectedSiteID = big.NewInt(1 << 0)
	getConnectedSitesRequestFieldFields          = big.NewInt(1 << 1)
	getConnectedSitesRequestFieldExcludeFields   = big.NewInt(1 << 2)
)

type GetConnectedSitesRequest struct {
	// The unique identifier for the site.
	ConnectedSiteID string `json:"-" url:"-"`
	// A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
	Fields []*string `json:"-" url:"fields,omitempty"`
	// A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
	ExcludeFields []*string `json:"-" url:"exclude_fields,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetConnectedSitesRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetConnectedSiteID sets the ConnectedSiteID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetConnectedSitesRequest) SetConnectedSiteID(connectedSiteID string) {
	g.ConnectedSiteID = connectedSiteID
	g.require(getConnectedSitesRequestFieldConnectedSiteID)
}

// SetFields sets the Fields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetConnectedSitesRequest) SetFields(fields []*string) {
	g.Fields = fields
	g.require(getConnectedSitesRequestFieldFields)
}

// SetExcludeFields sets the ExcludeFields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetConnectedSitesRequest) SetExcludeFields(excludeFields []*string) {
	g.ExcludeFields = excludeFields
	g.require(getConnectedSitesRequestFieldExcludeFields)
}

var (
	listConnectedSitesRequestFieldFields        = big.NewInt(1 << 0)
	listConnectedSitesRequestFieldExcludeFields = big.NewInt(1 << 1)
	listConnectedSitesRequestFieldCount         = big.NewInt(1 << 2)
	listConnectedSitesRequestFieldOffset        = big.NewInt(1 << 3)
)

type ListConnectedSitesRequest struct {
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

func (l *ListConnectedSitesRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetFields sets the Fields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConnectedSitesRequest) SetFields(fields []*string) {
	l.Fields = fields
	l.require(listConnectedSitesRequestFieldFields)
}

// SetExcludeFields sets the ExcludeFields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConnectedSitesRequest) SetExcludeFields(excludeFields []*string) {
	l.ExcludeFields = excludeFields
	l.require(listConnectedSitesRequestFieldExcludeFields)
}

// SetCount sets the Count field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConnectedSitesRequest) SetCount(count *int) {
	l.Count = count
	l.require(listConnectedSitesRequestFieldCount)
}

// SetOffset sets the Offset field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConnectedSitesRequest) SetOffset(offset *int) {
	l.Offset = offset
	l.require(listConnectedSitesRequestFieldOffset)
}

// Information about a specific connected site.
var (
	connectedSiteFieldLinks      = big.NewInt(1 << 0)
	connectedSiteFieldCreatedAt  = big.NewInt(1 << 1)
	connectedSiteFieldDomain     = big.NewInt(1 << 2)
	connectedSiteFieldForeignID  = big.NewInt(1 << 3)
	connectedSiteFieldPlatform   = big.NewInt(1 << 4)
	connectedSiteFieldSiteScript = big.NewInt(1 << 5)
	connectedSiteFieldStoreID    = big.NewInt(1 << 6)
	connectedSiteFieldUpdatedAt  = big.NewInt(1 << 7)
)

type ConnectedSite struct {
	// A list of link types and descriptions for the API schema documents.
	Links []*ConnectedSiteLinksItem `json:"_links,omitempty" url:"_links,omitempty"`
	// The date and time the connected site was created in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty" url:"created_at,omitempty"`
	// The connected site domain.
	Domain *string `json:"domain,omitempty" url:"domain,omitempty"`
	// The unique identifier for the site.
	ForeignID *string `json:"foreign_id,omitempty" url:"foreign_id,omitempty"`
	// The platform of the connected site.
	Platform *string `json:"platform,omitempty" url:"platform,omitempty"`
	// The script used to connect your site with Mailchimp.
	SiteScript *ConnectedSiteSiteScript `json:"site_script,omitempty" url:"site_script,omitempty"`
	// The unique identifier for the ecommerce store that's associated with the connected site (if any). The store_id for a specific connected site can't change.
	StoreID *string `json:"store_id,omitempty" url:"store_id,omitempty"`
	// The date and time the connected site was last updated in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty" url:"updated_at,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *ConnectedSite) GetLinks() []*ConnectedSiteLinksItem {
	if c == nil {
		return nil
	}
	return c.Links
}

func (c *ConnectedSite) GetCreatedAt() *time.Time {
	if c == nil {
		return nil
	}
	return c.CreatedAt
}

func (c *ConnectedSite) GetDomain() *string {
	if c == nil {
		return nil
	}
	return c.Domain
}

func (c *ConnectedSite) GetForeignID() *string {
	if c == nil {
		return nil
	}
	return c.ForeignID
}

func (c *ConnectedSite) GetPlatform() *string {
	if c == nil {
		return nil
	}
	return c.Platform
}

func (c *ConnectedSite) GetSiteScript() *ConnectedSiteSiteScript {
	if c == nil {
		return nil
	}
	return c.SiteScript
}

func (c *ConnectedSite) GetStoreID() *string {
	if c == nil {
		return nil
	}
	return c.StoreID
}

func (c *ConnectedSite) GetUpdatedAt() *time.Time {
	if c == nil {
		return nil
	}
	return c.UpdatedAt
}

func (c *ConnectedSite) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *ConnectedSite) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetLinks sets the Links field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ConnectedSite) SetLinks(links []*ConnectedSiteLinksItem) {
	c.Links = links
	c.require(connectedSiteFieldLinks)
}

// SetCreatedAt sets the CreatedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ConnectedSite) SetCreatedAt(createdAt *time.Time) {
	c.CreatedAt = createdAt
	c.require(connectedSiteFieldCreatedAt)
}

// SetDomain sets the Domain field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ConnectedSite) SetDomain(domain *string) {
	c.Domain = domain
	c.require(connectedSiteFieldDomain)
}

// SetForeignID sets the ForeignID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ConnectedSite) SetForeignID(foreignID *string) {
	c.ForeignID = foreignID
	c.require(connectedSiteFieldForeignID)
}

// SetPlatform sets the Platform field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ConnectedSite) SetPlatform(platform *string) {
	c.Platform = platform
	c.require(connectedSiteFieldPlatform)
}

// SetSiteScript sets the SiteScript field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ConnectedSite) SetSiteScript(siteScript *ConnectedSiteSiteScript) {
	c.SiteScript = siteScript
	c.require(connectedSiteFieldSiteScript)
}

// SetStoreID sets the StoreID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ConnectedSite) SetStoreID(storeID *string) {
	c.StoreID = storeID
	c.require(connectedSiteFieldStoreID)
}

// SetUpdatedAt sets the UpdatedAt field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ConnectedSite) SetUpdatedAt(updatedAt *time.Time) {
	c.UpdatedAt = updatedAt
	c.require(connectedSiteFieldUpdatedAt)
}

func (c *ConnectedSite) UnmarshalJSON(data []byte) error {
	type embed ConnectedSite
	var unmarshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"created_at,omitempty"`
		UpdatedAt *internal.DateTime `json:"updated_at,omitempty"`
	}{
		embed: embed(*c),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*c = ConnectedSite(unmarshaler.embed)
	c.CreatedAt = unmarshaler.CreatedAt.TimePtr()
	c.UpdatedAt = unmarshaler.UpdatedAt.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *ConnectedSite) MarshalJSON() ([]byte, error) {
	type embed ConnectedSite
	var marshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"created_at,omitempty"`
		UpdatedAt *internal.DateTime `json:"updated_at,omitempty"`
	}{
		embed:     embed(*c),
		CreatedAt: internal.NewOptionalDateTime(c.CreatedAt),
		UpdatedAt: internal.NewOptionalDateTime(c.UpdatedAt),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *ConnectedSite) String() string {
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
	connectedSiteLinksItemFieldHref         = big.NewInt(1 << 0)
	connectedSiteLinksItemFieldMethod       = big.NewInt(1 << 1)
	connectedSiteLinksItemFieldRel          = big.NewInt(1 << 2)
	connectedSiteLinksItemFieldSchema       = big.NewInt(1 << 3)
	connectedSiteLinksItemFieldTargetSchema = big.NewInt(1 << 4)
)

type ConnectedSiteLinksItem struct {
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href *string `json:"href,omitempty" url:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method *ConnectedSiteLinksItemMethod `json:"method,omitempty" url:"method,omitempty"`
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

func (c *ConnectedSiteLinksItem) GetHref() *string {
	if c == nil {
		return nil
	}
	return c.Href
}

func (c *ConnectedSiteLinksItem) GetMethod() *ConnectedSiteLinksItemMethod {
	if c == nil {
		return nil
	}
	return c.Method
}

func (c *ConnectedSiteLinksItem) GetRel() *string {
	if c == nil {
		return nil
	}
	return c.Rel
}

func (c *ConnectedSiteLinksItem) GetSchema() *string {
	if c == nil {
		return nil
	}
	return c.Schema
}

func (c *ConnectedSiteLinksItem) GetTargetSchema() *string {
	if c == nil {
		return nil
	}
	return c.TargetSchema
}

func (c *ConnectedSiteLinksItem) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *ConnectedSiteLinksItem) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetHref sets the Href field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ConnectedSiteLinksItem) SetHref(href *string) {
	c.Href = href
	c.require(connectedSiteLinksItemFieldHref)
}

// SetMethod sets the Method field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ConnectedSiteLinksItem) SetMethod(method *ConnectedSiteLinksItemMethod) {
	c.Method = method
	c.require(connectedSiteLinksItemFieldMethod)
}

// SetRel sets the Rel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ConnectedSiteLinksItem) SetRel(rel *string) {
	c.Rel = rel
	c.require(connectedSiteLinksItemFieldRel)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ConnectedSiteLinksItem) SetSchema(schema *string) {
	c.Schema = schema
	c.require(connectedSiteLinksItemFieldSchema)
}

// SetTargetSchema sets the TargetSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ConnectedSiteLinksItem) SetTargetSchema(targetSchema *string) {
	c.TargetSchema = targetSchema
	c.require(connectedSiteLinksItemFieldTargetSchema)
}

func (c *ConnectedSiteLinksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler ConnectedSiteLinksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConnectedSiteLinksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *ConnectedSiteLinksItem) MarshalJSON() ([]byte, error) {
	type embed ConnectedSiteLinksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *ConnectedSiteLinksItem) String() string {
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
type ConnectedSiteLinksItemMethod string

const (
	ConnectedSiteLinksItemMethodGet     ConnectedSiteLinksItemMethod = "GET"
	ConnectedSiteLinksItemMethodPost    ConnectedSiteLinksItemMethod = "POST"
	ConnectedSiteLinksItemMethodPut     ConnectedSiteLinksItemMethod = "PUT"
	ConnectedSiteLinksItemMethodPatch   ConnectedSiteLinksItemMethod = "PATCH"
	ConnectedSiteLinksItemMethodDelete  ConnectedSiteLinksItemMethod = "DELETE"
	ConnectedSiteLinksItemMethodOptions ConnectedSiteLinksItemMethod = "OPTIONS"
	ConnectedSiteLinksItemMethodHead    ConnectedSiteLinksItemMethod = "HEAD"
)

func NewConnectedSiteLinksItemMethodFromString(s string) (ConnectedSiteLinksItemMethod, error) {
	switch s {
	case "GET":
		return ConnectedSiteLinksItemMethodGet, nil
	case "POST":
		return ConnectedSiteLinksItemMethodPost, nil
	case "PUT":
		return ConnectedSiteLinksItemMethodPut, nil
	case "PATCH":
		return ConnectedSiteLinksItemMethodPatch, nil
	case "DELETE":
		return ConnectedSiteLinksItemMethodDelete, nil
	case "OPTIONS":
		return ConnectedSiteLinksItemMethodOptions, nil
	case "HEAD":
		return ConnectedSiteLinksItemMethodHead, nil
	}
	var t ConnectedSiteLinksItemMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c ConnectedSiteLinksItemMethod) Ptr() *ConnectedSiteLinksItemMethod {
	return &c
}

// The script used to connect your site with Mailchimp.
var (
	connectedSiteSiteScriptFieldFragment = big.NewInt(1 << 0)
	connectedSiteSiteScriptFieldURL      = big.NewInt(1 << 1)
)

type ConnectedSiteSiteScript struct {
	// A pre-built script that you can copy-and-paste into your site to integrate it with Mailchimp.
	Fragment *string `json:"fragment,omitempty" url:"fragment,omitempty"`
	// The URL used for any integrations that offer built-in support for connected sites.
	URL *string `json:"url,omitempty" url:"url,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *ConnectedSiteSiteScript) GetFragment() *string {
	if c == nil {
		return nil
	}
	return c.Fragment
}

func (c *ConnectedSiteSiteScript) GetURL() *string {
	if c == nil {
		return nil
	}
	return c.URL
}

func (c *ConnectedSiteSiteScript) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *ConnectedSiteSiteScript) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetFragment sets the Fragment field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ConnectedSiteSiteScript) SetFragment(fragment *string) {
	c.Fragment = fragment
	c.require(connectedSiteSiteScriptFieldFragment)
}

// SetURL sets the URL field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *ConnectedSiteSiteScript) SetURL(url *string) {
	c.URL = url
	c.require(connectedSiteSiteScriptFieldURL)
}

func (c *ConnectedSiteSiteScript) UnmarshalJSON(data []byte) error {
	type unmarshaler ConnectedSiteSiteScript
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConnectedSiteSiteScript(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *ConnectedSiteSiteScript) MarshalJSON() ([]byte, error) {
	type embed ConnectedSiteSiteScript
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *ConnectedSiteSiteScript) String() string {
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

// A collection of connected sites in the account.
var (
	listConnectedSitesResponseFieldLinks      = big.NewInt(1 << 0)
	listConnectedSitesResponseFieldSites      = big.NewInt(1 << 1)
	listConnectedSitesResponseFieldTotalItems = big.NewInt(1 << 2)
)

type ListConnectedSitesResponse struct {
	// A list of link types and descriptions for the API schema documents.
	Links []*ListConnectedSitesResponseLinksItem `json:"_links,omitempty" url:"_links,omitempty"`
	// An array of objects, each representing a connected site.
	Sites []*ConnectedSite `json:"sites,omitempty" url:"sites,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int `json:"total_items,omitempty" url:"total_items,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListConnectedSitesResponse) GetLinks() []*ListConnectedSitesResponseLinksItem {
	if l == nil {
		return nil
	}
	return l.Links
}

func (l *ListConnectedSitesResponse) GetSites() []*ConnectedSite {
	if l == nil {
		return nil
	}
	return l.Sites
}

func (l *ListConnectedSitesResponse) GetTotalItems() *int {
	if l == nil {
		return nil
	}
	return l.TotalItems
}

func (l *ListConnectedSitesResponse) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListConnectedSitesResponse) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetLinks sets the Links field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConnectedSitesResponse) SetLinks(links []*ListConnectedSitesResponseLinksItem) {
	l.Links = links
	l.require(listConnectedSitesResponseFieldLinks)
}

// SetSites sets the Sites field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConnectedSitesResponse) SetSites(sites []*ConnectedSite) {
	l.Sites = sites
	l.require(listConnectedSitesResponseFieldSites)
}

// SetTotalItems sets the TotalItems field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConnectedSitesResponse) SetTotalItems(totalItems *int) {
	l.TotalItems = totalItems
	l.require(listConnectedSitesResponseFieldTotalItems)
}

func (l *ListConnectedSitesResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListConnectedSitesResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListConnectedSitesResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListConnectedSitesResponse) MarshalJSON() ([]byte, error) {
	type embed ListConnectedSitesResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListConnectedSitesResponse) String() string {
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
	listConnectedSitesResponseLinksItemFieldHref         = big.NewInt(1 << 0)
	listConnectedSitesResponseLinksItemFieldMethod       = big.NewInt(1 << 1)
	listConnectedSitesResponseLinksItemFieldRel          = big.NewInt(1 << 2)
	listConnectedSitesResponseLinksItemFieldSchema       = big.NewInt(1 << 3)
	listConnectedSitesResponseLinksItemFieldTargetSchema = big.NewInt(1 << 4)
)

type ListConnectedSitesResponseLinksItem struct {
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href *string `json:"href,omitempty" url:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method *ListConnectedSitesResponseLinksItemMethod `json:"method,omitempty" url:"method,omitempty"`
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

func (l *ListConnectedSitesResponseLinksItem) GetHref() *string {
	if l == nil {
		return nil
	}
	return l.Href
}

func (l *ListConnectedSitesResponseLinksItem) GetMethod() *ListConnectedSitesResponseLinksItemMethod {
	if l == nil {
		return nil
	}
	return l.Method
}

func (l *ListConnectedSitesResponseLinksItem) GetRel() *string {
	if l == nil {
		return nil
	}
	return l.Rel
}

func (l *ListConnectedSitesResponseLinksItem) GetSchema() *string {
	if l == nil {
		return nil
	}
	return l.Schema
}

func (l *ListConnectedSitesResponseLinksItem) GetTargetSchema() *string {
	if l == nil {
		return nil
	}
	return l.TargetSchema
}

func (l *ListConnectedSitesResponseLinksItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListConnectedSitesResponseLinksItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetHref sets the Href field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConnectedSitesResponseLinksItem) SetHref(href *string) {
	l.Href = href
	l.require(listConnectedSitesResponseLinksItemFieldHref)
}

// SetMethod sets the Method field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConnectedSitesResponseLinksItem) SetMethod(method *ListConnectedSitesResponseLinksItemMethod) {
	l.Method = method
	l.require(listConnectedSitesResponseLinksItemFieldMethod)
}

// SetRel sets the Rel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConnectedSitesResponseLinksItem) SetRel(rel *string) {
	l.Rel = rel
	l.require(listConnectedSitesResponseLinksItemFieldRel)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConnectedSitesResponseLinksItem) SetSchema(schema *string) {
	l.Schema = schema
	l.require(listConnectedSitesResponseLinksItemFieldSchema)
}

// SetTargetSchema sets the TargetSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConnectedSitesResponseLinksItem) SetTargetSchema(targetSchema *string) {
	l.TargetSchema = targetSchema
	l.require(listConnectedSitesResponseLinksItemFieldTargetSchema)
}

func (l *ListConnectedSitesResponseLinksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler ListConnectedSitesResponseLinksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListConnectedSitesResponseLinksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListConnectedSitesResponseLinksItem) MarshalJSON() ([]byte, error) {
	type embed ListConnectedSitesResponseLinksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListConnectedSitesResponseLinksItem) String() string {
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
type ListConnectedSitesResponseLinksItemMethod string

const (
	ListConnectedSitesResponseLinksItemMethodGet     ListConnectedSitesResponseLinksItemMethod = "GET"
	ListConnectedSitesResponseLinksItemMethodPost    ListConnectedSitesResponseLinksItemMethod = "POST"
	ListConnectedSitesResponseLinksItemMethodPut     ListConnectedSitesResponseLinksItemMethod = "PUT"
	ListConnectedSitesResponseLinksItemMethodPatch   ListConnectedSitesResponseLinksItemMethod = "PATCH"
	ListConnectedSitesResponseLinksItemMethodDelete  ListConnectedSitesResponseLinksItemMethod = "DELETE"
	ListConnectedSitesResponseLinksItemMethodOptions ListConnectedSitesResponseLinksItemMethod = "OPTIONS"
	ListConnectedSitesResponseLinksItemMethodHead    ListConnectedSitesResponseLinksItemMethod = "HEAD"
)

func NewListConnectedSitesResponseLinksItemMethodFromString(s string) (ListConnectedSitesResponseLinksItemMethod, error) {
	switch s {
	case "GET":
		return ListConnectedSitesResponseLinksItemMethodGet, nil
	case "POST":
		return ListConnectedSitesResponseLinksItemMethodPost, nil
	case "PUT":
		return ListConnectedSitesResponseLinksItemMethodPut, nil
	case "PATCH":
		return ListConnectedSitesResponseLinksItemMethodPatch, nil
	case "DELETE":
		return ListConnectedSitesResponseLinksItemMethodDelete, nil
	case "OPTIONS":
		return ListConnectedSitesResponseLinksItemMethodOptions, nil
	case "HEAD":
		return ListConnectedSitesResponseLinksItemMethodHead, nil
	}
	var t ListConnectedSitesResponseLinksItemMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListConnectedSitesResponseLinksItemMethod) Ptr() *ListConnectedSitesResponseLinksItemMethod {
	return &l
}
