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
	createAccountExportsRequestFieldIncludeStages  = big.NewInt(1 << 0)
	createAccountExportsRequestFieldSinceTimestamp = big.NewInt(1 << 1)
)

type CreateAccountExportsRequest struct {
	// The stages of an account export to include.
	IncludeStages []CreateAccountExportsRequestIncludeStagesItem `json:"include_stages" url:"-"`
	// An ISO 8601 date that will limit the export to only records created after a given time. For instance, the reports stage will contain any campaign sent after the given timestamp. Audiences, however, are excluded from this limit.
	SinceTimestamp *time.Time `json:"since_timestamp,omitempty" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CreateAccountExportsRequest) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetIncludeStages sets the IncludeStages field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAccountExportsRequest) SetIncludeStages(includeStages []CreateAccountExportsRequestIncludeStagesItem) {
	c.IncludeStages = includeStages
	c.require(createAccountExportsRequestFieldIncludeStages)
}

// SetSinceTimestamp sets the SinceTimestamp field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAccountExportsRequest) SetSinceTimestamp(sinceTimestamp *time.Time) {
	c.SinceTimestamp = sinceTimestamp
	c.require(createAccountExportsRequestFieldSinceTimestamp)
}

func (c *CreateAccountExportsRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateAccountExportsRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*c = CreateAccountExportsRequest(body)
	return nil
}

func (c *CreateAccountExportsRequest) MarshalJSON() ([]byte, error) {
	type embed CreateAccountExportsRequest
	var marshaler = struct {
		embed
		SinceTimestamp *internal.DateTime `json:"since_timestamp,omitempty"`
	}{
		embed:          embed(*c),
		SinceTimestamp: internal.NewOptionalDateTime(c.SinceTimestamp),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

var (
	getAccountExportsRequestFieldExportID      = big.NewInt(1 << 0)
	getAccountExportsRequestFieldFields        = big.NewInt(1 << 1)
	getAccountExportsRequestFieldExcludeFields = big.NewInt(1 << 2)
)

type GetAccountExportsRequest struct {
	// The unique id for the account export.
	ExportID string `json:"-" url:"-"`
	// A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
	Fields []*string `json:"-" url:"fields,omitempty"`
	// A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
	ExcludeFields []*string `json:"-" url:"exclude_fields,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetAccountExportsRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetExportID sets the ExportID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAccountExportsRequest) SetExportID(exportID string) {
	g.ExportID = exportID
	g.require(getAccountExportsRequestFieldExportID)
}

// SetFields sets the Fields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAccountExportsRequest) SetFields(fields []*string) {
	g.Fields = fields
	g.require(getAccountExportsRequestFieldFields)
}

// SetExcludeFields sets the ExcludeFields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAccountExportsRequest) SetExcludeFields(excludeFields []*string) {
	g.ExcludeFields = excludeFields
	g.require(getAccountExportsRequestFieldExcludeFields)
}

var (
	listAccountExportsRequestFieldFields        = big.NewInt(1 << 0)
	listAccountExportsRequestFieldExcludeFields = big.NewInt(1 << 1)
	listAccountExportsRequestFieldCount         = big.NewInt(1 << 2)
	listAccountExportsRequestFieldOffset        = big.NewInt(1 << 3)
)

type ListAccountExportsRequest struct {
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

func (l *ListAccountExportsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetFields sets the Fields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAccountExportsRequest) SetFields(fields []*string) {
	l.Fields = fields
	l.require(listAccountExportsRequestFieldFields)
}

// SetExcludeFields sets the ExcludeFields field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAccountExportsRequest) SetExcludeFields(excludeFields []*string) {
	l.ExcludeFields = excludeFields
	l.require(listAccountExportsRequestFieldExcludeFields)
}

// SetCount sets the Count field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAccountExportsRequest) SetCount(count *int) {
	l.Count = count
	l.require(listAccountExportsRequestFieldCount)
}

// SetOffset sets the Offset field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAccountExportsRequest) SetOffset(offset *int) {
	l.Offset = offset
	l.require(listAccountExportsRequestFieldOffset)
}

type CreateAccountExportsRequestIncludeStagesItem string

const (
	CreateAccountExportsRequestIncludeStagesItemAudiences    CreateAccountExportsRequestIncludeStagesItem = "audiences"
	CreateAccountExportsRequestIncludeStagesItemCampaigns    CreateAccountExportsRequestIncludeStagesItem = "campaigns"
	CreateAccountExportsRequestIncludeStagesItemEvents       CreateAccountExportsRequestIncludeStagesItem = "events"
	CreateAccountExportsRequestIncludeStagesItemGalleryFiles CreateAccountExportsRequestIncludeStagesItem = "gallery_files"
	CreateAccountExportsRequestIncludeStagesItemReports      CreateAccountExportsRequestIncludeStagesItem = "reports"
	CreateAccountExportsRequestIncludeStagesItemTemplates    CreateAccountExportsRequestIncludeStagesItem = "templates"
)

func NewCreateAccountExportsRequestIncludeStagesItemFromString(s string) (CreateAccountExportsRequestIncludeStagesItem, error) {
	switch s {
	case "audiences":
		return CreateAccountExportsRequestIncludeStagesItemAudiences, nil
	case "campaigns":
		return CreateAccountExportsRequestIncludeStagesItemCampaigns, nil
	case "events":
		return CreateAccountExportsRequestIncludeStagesItemEvents, nil
	case "gallery_files":
		return CreateAccountExportsRequestIncludeStagesItemGalleryFiles, nil
	case "reports":
		return CreateAccountExportsRequestIncludeStagesItemReports, nil
	case "templates":
		return CreateAccountExportsRequestIncludeStagesItemTemplates, nil
	}
	var t CreateAccountExportsRequestIncludeStagesItem
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c CreateAccountExportsRequestIncludeStagesItem) Ptr() *CreateAccountExportsRequestIncludeStagesItem {
	return &c
}

// An account export.
var (
	createAccountExportsResponseFieldLinks       = big.NewInt(1 << 0)
	createAccountExportsResponseFieldDownloadURL = big.NewInt(1 << 1)
	createAccountExportsResponseFieldExportID    = big.NewInt(1 << 2)
	createAccountExportsResponseFieldFinished    = big.NewInt(1 << 3)
	createAccountExportsResponseFieldSizeInBytes = big.NewInt(1 << 4)
	createAccountExportsResponseFieldStarted     = big.NewInt(1 << 5)
)

type CreateAccountExportsResponse struct {
	// A list of link types and descriptions for the API schema documents.
	Links []*CreateAccountExportsResponseLinksItem `json:"_links,omitempty" url:"_links,omitempty"`
	// If the export is finished, the download URL for an export. URLs are only valid for 90 days after the export completes.
	DownloadURL *string `json:"download_url,omitempty" url:"download_url,omitempty"`
	// The ID for the export.
	ExportID *int `json:"export_id,omitempty" url:"export_id,omitempty"`
	// If finished, the finish time for the export.
	Finished *time.Time `json:"finished,omitempty" url:"finished,omitempty"`
	// The size of the uncompressed export in bytes.
	SizeInBytes *int `json:"size_in_bytes,omitempty" url:"size_in_bytes,omitempty"`
	// Start time for the export.
	Started *time.Time `json:"started,omitempty" url:"started,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateAccountExportsResponse) GetLinks() []*CreateAccountExportsResponseLinksItem {
	if c == nil {
		return nil
	}
	return c.Links
}

func (c *CreateAccountExportsResponse) GetDownloadURL() *string {
	if c == nil {
		return nil
	}
	return c.DownloadURL
}

func (c *CreateAccountExportsResponse) GetExportID() *int {
	if c == nil {
		return nil
	}
	return c.ExportID
}

func (c *CreateAccountExportsResponse) GetFinished() *time.Time {
	if c == nil {
		return nil
	}
	return c.Finished
}

func (c *CreateAccountExportsResponse) GetSizeInBytes() *int {
	if c == nil {
		return nil
	}
	return c.SizeInBytes
}

func (c *CreateAccountExportsResponse) GetStarted() *time.Time {
	if c == nil {
		return nil
	}
	return c.Started
}

func (c *CreateAccountExportsResponse) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *CreateAccountExportsResponse) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetLinks sets the Links field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAccountExportsResponse) SetLinks(links []*CreateAccountExportsResponseLinksItem) {
	c.Links = links
	c.require(createAccountExportsResponseFieldLinks)
}

// SetDownloadURL sets the DownloadURL field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAccountExportsResponse) SetDownloadURL(downloadURL *string) {
	c.DownloadURL = downloadURL
	c.require(createAccountExportsResponseFieldDownloadURL)
}

// SetExportID sets the ExportID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAccountExportsResponse) SetExportID(exportID *int) {
	c.ExportID = exportID
	c.require(createAccountExportsResponseFieldExportID)
}

// SetFinished sets the Finished field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAccountExportsResponse) SetFinished(finished *time.Time) {
	c.Finished = finished
	c.require(createAccountExportsResponseFieldFinished)
}

// SetSizeInBytes sets the SizeInBytes field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAccountExportsResponse) SetSizeInBytes(sizeInBytes *int) {
	c.SizeInBytes = sizeInBytes
	c.require(createAccountExportsResponseFieldSizeInBytes)
}

// SetStarted sets the Started field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAccountExportsResponse) SetStarted(started *time.Time) {
	c.Started = started
	c.require(createAccountExportsResponseFieldStarted)
}

func (c *CreateAccountExportsResponse) UnmarshalJSON(data []byte) error {
	type embed CreateAccountExportsResponse
	var unmarshaler = struct {
		embed
		Finished *internal.DateTime `json:"finished,omitempty"`
		Started  *internal.DateTime `json:"started,omitempty"`
	}{
		embed: embed(*c),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*c = CreateAccountExportsResponse(unmarshaler.embed)
	c.Finished = unmarshaler.Finished.TimePtr()
	c.Started = unmarshaler.Started.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateAccountExportsResponse) MarshalJSON() ([]byte, error) {
	type embed CreateAccountExportsResponse
	var marshaler = struct {
		embed
		Finished *internal.DateTime `json:"finished,omitempty"`
		Started  *internal.DateTime `json:"started,omitempty"`
	}{
		embed:    embed(*c),
		Finished: internal.NewOptionalDateTime(c.Finished),
		Started:  internal.NewOptionalDateTime(c.Started),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *CreateAccountExportsResponse) String() string {
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
	createAccountExportsResponseLinksItemFieldHref         = big.NewInt(1 << 0)
	createAccountExportsResponseLinksItemFieldMethod       = big.NewInt(1 << 1)
	createAccountExportsResponseLinksItemFieldRel          = big.NewInt(1 << 2)
	createAccountExportsResponseLinksItemFieldSchema       = big.NewInt(1 << 3)
	createAccountExportsResponseLinksItemFieldTargetSchema = big.NewInt(1 << 4)
)

type CreateAccountExportsResponseLinksItem struct {
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href *string `json:"href,omitempty" url:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method *CreateAccountExportsResponseLinksItemMethod `json:"method,omitempty" url:"method,omitempty"`
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

func (c *CreateAccountExportsResponseLinksItem) GetHref() *string {
	if c == nil {
		return nil
	}
	return c.Href
}

func (c *CreateAccountExportsResponseLinksItem) GetMethod() *CreateAccountExportsResponseLinksItemMethod {
	if c == nil {
		return nil
	}
	return c.Method
}

func (c *CreateAccountExportsResponseLinksItem) GetRel() *string {
	if c == nil {
		return nil
	}
	return c.Rel
}

func (c *CreateAccountExportsResponseLinksItem) GetSchema() *string {
	if c == nil {
		return nil
	}
	return c.Schema
}

func (c *CreateAccountExportsResponseLinksItem) GetTargetSchema() *string {
	if c == nil {
		return nil
	}
	return c.TargetSchema
}

func (c *CreateAccountExportsResponseLinksItem) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *CreateAccountExportsResponseLinksItem) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetHref sets the Href field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAccountExportsResponseLinksItem) SetHref(href *string) {
	c.Href = href
	c.require(createAccountExportsResponseLinksItemFieldHref)
}

// SetMethod sets the Method field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAccountExportsResponseLinksItem) SetMethod(method *CreateAccountExportsResponseLinksItemMethod) {
	c.Method = method
	c.require(createAccountExportsResponseLinksItemFieldMethod)
}

// SetRel sets the Rel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAccountExportsResponseLinksItem) SetRel(rel *string) {
	c.Rel = rel
	c.require(createAccountExportsResponseLinksItemFieldRel)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAccountExportsResponseLinksItem) SetSchema(schema *string) {
	c.Schema = schema
	c.require(createAccountExportsResponseLinksItemFieldSchema)
}

// SetTargetSchema sets the TargetSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateAccountExportsResponseLinksItem) SetTargetSchema(targetSchema *string) {
	c.TargetSchema = targetSchema
	c.require(createAccountExportsResponseLinksItemFieldTargetSchema)
}

func (c *CreateAccountExportsResponseLinksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateAccountExportsResponseLinksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateAccountExportsResponseLinksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateAccountExportsResponseLinksItem) MarshalJSON() ([]byte, error) {
	type embed CreateAccountExportsResponseLinksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *CreateAccountExportsResponseLinksItem) String() string {
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
type CreateAccountExportsResponseLinksItemMethod string

const (
	CreateAccountExportsResponseLinksItemMethodGet     CreateAccountExportsResponseLinksItemMethod = "GET"
	CreateAccountExportsResponseLinksItemMethodPost    CreateAccountExportsResponseLinksItemMethod = "POST"
	CreateAccountExportsResponseLinksItemMethodPut     CreateAccountExportsResponseLinksItemMethod = "PUT"
	CreateAccountExportsResponseLinksItemMethodPatch   CreateAccountExportsResponseLinksItemMethod = "PATCH"
	CreateAccountExportsResponseLinksItemMethodDelete  CreateAccountExportsResponseLinksItemMethod = "DELETE"
	CreateAccountExportsResponseLinksItemMethodOptions CreateAccountExportsResponseLinksItemMethod = "OPTIONS"
	CreateAccountExportsResponseLinksItemMethodHead    CreateAccountExportsResponseLinksItemMethod = "HEAD"
)

func NewCreateAccountExportsResponseLinksItemMethodFromString(s string) (CreateAccountExportsResponseLinksItemMethod, error) {
	switch s {
	case "GET":
		return CreateAccountExportsResponseLinksItemMethodGet, nil
	case "POST":
		return CreateAccountExportsResponseLinksItemMethodPost, nil
	case "PUT":
		return CreateAccountExportsResponseLinksItemMethodPut, nil
	case "PATCH":
		return CreateAccountExportsResponseLinksItemMethodPatch, nil
	case "DELETE":
		return CreateAccountExportsResponseLinksItemMethodDelete, nil
	case "OPTIONS":
		return CreateAccountExportsResponseLinksItemMethodOptions, nil
	case "HEAD":
		return CreateAccountExportsResponseLinksItemMethodHead, nil
	}
	var t CreateAccountExportsResponseLinksItemMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c CreateAccountExportsResponseLinksItemMethod) Ptr() *CreateAccountExportsResponseLinksItemMethod {
	return &c
}

// An account export.
var (
	getAccountExportsResponseFieldLinks       = big.NewInt(1 << 0)
	getAccountExportsResponseFieldDownloadURL = big.NewInt(1 << 1)
	getAccountExportsResponseFieldExportID    = big.NewInt(1 << 2)
	getAccountExportsResponseFieldFinished    = big.NewInt(1 << 3)
	getAccountExportsResponseFieldSizeInBytes = big.NewInt(1 << 4)
	getAccountExportsResponseFieldStarted     = big.NewInt(1 << 5)
)

type GetAccountExportsResponse struct {
	// A list of link types and descriptions for the API schema documents.
	Links []*GetAccountExportsResponseLinksItem `json:"_links,omitempty" url:"_links,omitempty"`
	// If the export is finished, the download URL for an export. URLs are only valid for 90 days after the export completes.
	DownloadURL *string `json:"download_url,omitempty" url:"download_url,omitempty"`
	// The ID for the export.
	ExportID *int `json:"export_id,omitempty" url:"export_id,omitempty"`
	// If finished, the finish time for the export.
	Finished *time.Time `json:"finished,omitempty" url:"finished,omitempty"`
	// The size of the uncompressed export in bytes.
	SizeInBytes *int `json:"size_in_bytes,omitempty" url:"size_in_bytes,omitempty"`
	// Start time for the export.
	Started *time.Time `json:"started,omitempty" url:"started,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (g *GetAccountExportsResponse) GetLinks() []*GetAccountExportsResponseLinksItem {
	if g == nil {
		return nil
	}
	return g.Links
}

func (g *GetAccountExportsResponse) GetDownloadURL() *string {
	if g == nil {
		return nil
	}
	return g.DownloadURL
}

func (g *GetAccountExportsResponse) GetExportID() *int {
	if g == nil {
		return nil
	}
	return g.ExportID
}

func (g *GetAccountExportsResponse) GetFinished() *time.Time {
	if g == nil {
		return nil
	}
	return g.Finished
}

func (g *GetAccountExportsResponse) GetSizeInBytes() *int {
	if g == nil {
		return nil
	}
	return g.SizeInBytes
}

func (g *GetAccountExportsResponse) GetStarted() *time.Time {
	if g == nil {
		return nil
	}
	return g.Started
}

func (g *GetAccountExportsResponse) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g *GetAccountExportsResponse) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetLinks sets the Links field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAccountExportsResponse) SetLinks(links []*GetAccountExportsResponseLinksItem) {
	g.Links = links
	g.require(getAccountExportsResponseFieldLinks)
}

// SetDownloadURL sets the DownloadURL field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAccountExportsResponse) SetDownloadURL(downloadURL *string) {
	g.DownloadURL = downloadURL
	g.require(getAccountExportsResponseFieldDownloadURL)
}

// SetExportID sets the ExportID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAccountExportsResponse) SetExportID(exportID *int) {
	g.ExportID = exportID
	g.require(getAccountExportsResponseFieldExportID)
}

// SetFinished sets the Finished field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAccountExportsResponse) SetFinished(finished *time.Time) {
	g.Finished = finished
	g.require(getAccountExportsResponseFieldFinished)
}

// SetSizeInBytes sets the SizeInBytes field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAccountExportsResponse) SetSizeInBytes(sizeInBytes *int) {
	g.SizeInBytes = sizeInBytes
	g.require(getAccountExportsResponseFieldSizeInBytes)
}

// SetStarted sets the Started field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAccountExportsResponse) SetStarted(started *time.Time) {
	g.Started = started
	g.require(getAccountExportsResponseFieldStarted)
}

func (g *GetAccountExportsResponse) UnmarshalJSON(data []byte) error {
	type embed GetAccountExportsResponse
	var unmarshaler = struct {
		embed
		Finished *internal.DateTime `json:"finished,omitempty"`
		Started  *internal.DateTime `json:"started,omitempty"`
	}{
		embed: embed(*g),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*g = GetAccountExportsResponse(unmarshaler.embed)
	g.Finished = unmarshaler.Finished.TimePtr()
	g.Started = unmarshaler.Started.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetAccountExportsResponse) MarshalJSON() ([]byte, error) {
	type embed GetAccountExportsResponse
	var marshaler = struct {
		embed
		Finished *internal.DateTime `json:"finished,omitempty"`
		Started  *internal.DateTime `json:"started,omitempty"`
	}{
		embed:    embed(*g),
		Finished: internal.NewOptionalDateTime(g.Finished),
		Started:  internal.NewOptionalDateTime(g.Started),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, g.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (g *GetAccountExportsResponse) String() string {
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
	getAccountExportsResponseLinksItemFieldHref         = big.NewInt(1 << 0)
	getAccountExportsResponseLinksItemFieldMethod       = big.NewInt(1 << 1)
	getAccountExportsResponseLinksItemFieldRel          = big.NewInt(1 << 2)
	getAccountExportsResponseLinksItemFieldSchema       = big.NewInt(1 << 3)
	getAccountExportsResponseLinksItemFieldTargetSchema = big.NewInt(1 << 4)
)

type GetAccountExportsResponseLinksItem struct {
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href *string `json:"href,omitempty" url:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method *GetAccountExportsResponseLinksItemMethod `json:"method,omitempty" url:"method,omitempty"`
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

func (g *GetAccountExportsResponseLinksItem) GetHref() *string {
	if g == nil {
		return nil
	}
	return g.Href
}

func (g *GetAccountExportsResponseLinksItem) GetMethod() *GetAccountExportsResponseLinksItemMethod {
	if g == nil {
		return nil
	}
	return g.Method
}

func (g *GetAccountExportsResponseLinksItem) GetRel() *string {
	if g == nil {
		return nil
	}
	return g.Rel
}

func (g *GetAccountExportsResponseLinksItem) GetSchema() *string {
	if g == nil {
		return nil
	}
	return g.Schema
}

func (g *GetAccountExportsResponseLinksItem) GetTargetSchema() *string {
	if g == nil {
		return nil
	}
	return g.TargetSchema
}

func (g *GetAccountExportsResponseLinksItem) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g *GetAccountExportsResponseLinksItem) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetHref sets the Href field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAccountExportsResponseLinksItem) SetHref(href *string) {
	g.Href = href
	g.require(getAccountExportsResponseLinksItemFieldHref)
}

// SetMethod sets the Method field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAccountExportsResponseLinksItem) SetMethod(method *GetAccountExportsResponseLinksItemMethod) {
	g.Method = method
	g.require(getAccountExportsResponseLinksItemFieldMethod)
}

// SetRel sets the Rel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAccountExportsResponseLinksItem) SetRel(rel *string) {
	g.Rel = rel
	g.require(getAccountExportsResponseLinksItemFieldRel)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAccountExportsResponseLinksItem) SetSchema(schema *string) {
	g.Schema = schema
	g.require(getAccountExportsResponseLinksItemFieldSchema)
}

// SetTargetSchema sets the TargetSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAccountExportsResponseLinksItem) SetTargetSchema(targetSchema *string) {
	g.TargetSchema = targetSchema
	g.require(getAccountExportsResponseLinksItemFieldTargetSchema)
}

func (g *GetAccountExportsResponseLinksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler GetAccountExportsResponseLinksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetAccountExportsResponseLinksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetAccountExportsResponseLinksItem) MarshalJSON() ([]byte, error) {
	type embed GetAccountExportsResponseLinksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, g.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (g *GetAccountExportsResponseLinksItem) String() string {
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
type GetAccountExportsResponseLinksItemMethod string

const (
	GetAccountExportsResponseLinksItemMethodGet     GetAccountExportsResponseLinksItemMethod = "GET"
	GetAccountExportsResponseLinksItemMethodPost    GetAccountExportsResponseLinksItemMethod = "POST"
	GetAccountExportsResponseLinksItemMethodPut     GetAccountExportsResponseLinksItemMethod = "PUT"
	GetAccountExportsResponseLinksItemMethodPatch   GetAccountExportsResponseLinksItemMethod = "PATCH"
	GetAccountExportsResponseLinksItemMethodDelete  GetAccountExportsResponseLinksItemMethod = "DELETE"
	GetAccountExportsResponseLinksItemMethodOptions GetAccountExportsResponseLinksItemMethod = "OPTIONS"
	GetAccountExportsResponseLinksItemMethodHead    GetAccountExportsResponseLinksItemMethod = "HEAD"
)

func NewGetAccountExportsResponseLinksItemMethodFromString(s string) (GetAccountExportsResponseLinksItemMethod, error) {
	switch s {
	case "GET":
		return GetAccountExportsResponseLinksItemMethodGet, nil
	case "POST":
		return GetAccountExportsResponseLinksItemMethodPost, nil
	case "PUT":
		return GetAccountExportsResponseLinksItemMethodPut, nil
	case "PATCH":
		return GetAccountExportsResponseLinksItemMethodPatch, nil
	case "DELETE":
		return GetAccountExportsResponseLinksItemMethodDelete, nil
	case "OPTIONS":
		return GetAccountExportsResponseLinksItemMethodOptions, nil
	case "HEAD":
		return GetAccountExportsResponseLinksItemMethodHead, nil
	}
	var t GetAccountExportsResponseLinksItemMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (g GetAccountExportsResponseLinksItemMethod) Ptr() *GetAccountExportsResponseLinksItemMethod {
	return &g
}

// An array of objects, each representing an account export.
var (
	listAccountExportsResponseFieldLinks      = big.NewInt(1 << 0)
	listAccountExportsResponseFieldExports    = big.NewInt(1 << 1)
	listAccountExportsResponseFieldTotalItems = big.NewInt(1 << 2)
)

type ListAccountExportsResponse struct {
	// A list of link types and descriptions for the API schema documents.
	Links []*ListAccountExportsResponseLinksItem `json:"_links,omitempty" url:"_links,omitempty"`
	// An array of objects, each representing an account export.
	Exports []*ListAccountExportsResponseExportsItem `json:"exports,omitempty" url:"exports,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int `json:"total_items,omitempty" url:"total_items,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListAccountExportsResponse) GetLinks() []*ListAccountExportsResponseLinksItem {
	if l == nil {
		return nil
	}
	return l.Links
}

func (l *ListAccountExportsResponse) GetExports() []*ListAccountExportsResponseExportsItem {
	if l == nil {
		return nil
	}
	return l.Exports
}

func (l *ListAccountExportsResponse) GetTotalItems() *int {
	if l == nil {
		return nil
	}
	return l.TotalItems
}

func (l *ListAccountExportsResponse) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListAccountExportsResponse) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetLinks sets the Links field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAccountExportsResponse) SetLinks(links []*ListAccountExportsResponseLinksItem) {
	l.Links = links
	l.require(listAccountExportsResponseFieldLinks)
}

// SetExports sets the Exports field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAccountExportsResponse) SetExports(exports []*ListAccountExportsResponseExportsItem) {
	l.Exports = exports
	l.require(listAccountExportsResponseFieldExports)
}

// SetTotalItems sets the TotalItems field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAccountExportsResponse) SetTotalItems(totalItems *int) {
	l.TotalItems = totalItems
	l.require(listAccountExportsResponseFieldTotalItems)
}

func (l *ListAccountExportsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListAccountExportsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListAccountExportsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListAccountExportsResponse) MarshalJSON() ([]byte, error) {
	type embed ListAccountExportsResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListAccountExportsResponse) String() string {
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

// An account export.
var (
	listAccountExportsResponseExportsItemFieldLinks       = big.NewInt(1 << 0)
	listAccountExportsResponseExportsItemFieldDownloadURL = big.NewInt(1 << 1)
	listAccountExportsResponseExportsItemFieldExportID    = big.NewInt(1 << 2)
	listAccountExportsResponseExportsItemFieldFinished    = big.NewInt(1 << 3)
	listAccountExportsResponseExportsItemFieldSizeInBytes = big.NewInt(1 << 4)
	listAccountExportsResponseExportsItemFieldStarted     = big.NewInt(1 << 5)
)

type ListAccountExportsResponseExportsItem struct {
	// A list of link types and descriptions for the API schema documents.
	Links []*ListAccountExportsResponseExportsItemLinksItem `json:"_links,omitempty" url:"_links,omitempty"`
	// If the export is finished, the download URL for an export. URLs are only valid for 90 days after the export completes.
	DownloadURL *string `json:"download_url,omitempty" url:"download_url,omitempty"`
	// The ID for the export.
	ExportID *int `json:"export_id,omitempty" url:"export_id,omitempty"`
	// If finished, the finish time for the export.
	Finished *time.Time `json:"finished,omitempty" url:"finished,omitempty"`
	// The size of the uncompressed export in bytes.
	SizeInBytes *int `json:"size_in_bytes,omitempty" url:"size_in_bytes,omitempty"`
	// Start time for the export.
	Started *time.Time `json:"started,omitempty" url:"started,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListAccountExportsResponseExportsItem) GetLinks() []*ListAccountExportsResponseExportsItemLinksItem {
	if l == nil {
		return nil
	}
	return l.Links
}

func (l *ListAccountExportsResponseExportsItem) GetDownloadURL() *string {
	if l == nil {
		return nil
	}
	return l.DownloadURL
}

func (l *ListAccountExportsResponseExportsItem) GetExportID() *int {
	if l == nil {
		return nil
	}
	return l.ExportID
}

func (l *ListAccountExportsResponseExportsItem) GetFinished() *time.Time {
	if l == nil {
		return nil
	}
	return l.Finished
}

func (l *ListAccountExportsResponseExportsItem) GetSizeInBytes() *int {
	if l == nil {
		return nil
	}
	return l.SizeInBytes
}

func (l *ListAccountExportsResponseExportsItem) GetStarted() *time.Time {
	if l == nil {
		return nil
	}
	return l.Started
}

func (l *ListAccountExportsResponseExportsItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListAccountExportsResponseExportsItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetLinks sets the Links field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAccountExportsResponseExportsItem) SetLinks(links []*ListAccountExportsResponseExportsItemLinksItem) {
	l.Links = links
	l.require(listAccountExportsResponseExportsItemFieldLinks)
}

// SetDownloadURL sets the DownloadURL field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAccountExportsResponseExportsItem) SetDownloadURL(downloadURL *string) {
	l.DownloadURL = downloadURL
	l.require(listAccountExportsResponseExportsItemFieldDownloadURL)
}

// SetExportID sets the ExportID field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAccountExportsResponseExportsItem) SetExportID(exportID *int) {
	l.ExportID = exportID
	l.require(listAccountExportsResponseExportsItemFieldExportID)
}

// SetFinished sets the Finished field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAccountExportsResponseExportsItem) SetFinished(finished *time.Time) {
	l.Finished = finished
	l.require(listAccountExportsResponseExportsItemFieldFinished)
}

// SetSizeInBytes sets the SizeInBytes field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAccountExportsResponseExportsItem) SetSizeInBytes(sizeInBytes *int) {
	l.SizeInBytes = sizeInBytes
	l.require(listAccountExportsResponseExportsItemFieldSizeInBytes)
}

// SetStarted sets the Started field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAccountExportsResponseExportsItem) SetStarted(started *time.Time) {
	l.Started = started
	l.require(listAccountExportsResponseExportsItemFieldStarted)
}

func (l *ListAccountExportsResponseExportsItem) UnmarshalJSON(data []byte) error {
	type embed ListAccountExportsResponseExportsItem
	var unmarshaler = struct {
		embed
		Finished *internal.DateTime `json:"finished,omitempty"`
		Started  *internal.DateTime `json:"started,omitempty"`
	}{
		embed: embed(*l),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*l = ListAccountExportsResponseExportsItem(unmarshaler.embed)
	l.Finished = unmarshaler.Finished.TimePtr()
	l.Started = unmarshaler.Started.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListAccountExportsResponseExportsItem) MarshalJSON() ([]byte, error) {
	type embed ListAccountExportsResponseExportsItem
	var marshaler = struct {
		embed
		Finished *internal.DateTime `json:"finished,omitempty"`
		Started  *internal.DateTime `json:"started,omitempty"`
	}{
		embed:    embed(*l),
		Finished: internal.NewOptionalDateTime(l.Finished),
		Started:  internal.NewOptionalDateTime(l.Started),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListAccountExportsResponseExportsItem) String() string {
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
	listAccountExportsResponseExportsItemLinksItemFieldHref         = big.NewInt(1 << 0)
	listAccountExportsResponseExportsItemLinksItemFieldMethod       = big.NewInt(1 << 1)
	listAccountExportsResponseExportsItemLinksItemFieldRel          = big.NewInt(1 << 2)
	listAccountExportsResponseExportsItemLinksItemFieldSchema       = big.NewInt(1 << 3)
	listAccountExportsResponseExportsItemLinksItemFieldTargetSchema = big.NewInt(1 << 4)
)

type ListAccountExportsResponseExportsItemLinksItem struct {
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href *string `json:"href,omitempty" url:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method *ListAccountExportsResponseExportsItemLinksItemMethod `json:"method,omitempty" url:"method,omitempty"`
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

func (l *ListAccountExportsResponseExportsItemLinksItem) GetHref() *string {
	if l == nil {
		return nil
	}
	return l.Href
}

func (l *ListAccountExportsResponseExportsItemLinksItem) GetMethod() *ListAccountExportsResponseExportsItemLinksItemMethod {
	if l == nil {
		return nil
	}
	return l.Method
}

func (l *ListAccountExportsResponseExportsItemLinksItem) GetRel() *string {
	if l == nil {
		return nil
	}
	return l.Rel
}

func (l *ListAccountExportsResponseExportsItemLinksItem) GetSchema() *string {
	if l == nil {
		return nil
	}
	return l.Schema
}

func (l *ListAccountExportsResponseExportsItemLinksItem) GetTargetSchema() *string {
	if l == nil {
		return nil
	}
	return l.TargetSchema
}

func (l *ListAccountExportsResponseExportsItemLinksItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListAccountExportsResponseExportsItemLinksItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetHref sets the Href field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAccountExportsResponseExportsItemLinksItem) SetHref(href *string) {
	l.Href = href
	l.require(listAccountExportsResponseExportsItemLinksItemFieldHref)
}

// SetMethod sets the Method field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAccountExportsResponseExportsItemLinksItem) SetMethod(method *ListAccountExportsResponseExportsItemLinksItemMethod) {
	l.Method = method
	l.require(listAccountExportsResponseExportsItemLinksItemFieldMethod)
}

// SetRel sets the Rel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAccountExportsResponseExportsItemLinksItem) SetRel(rel *string) {
	l.Rel = rel
	l.require(listAccountExportsResponseExportsItemLinksItemFieldRel)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAccountExportsResponseExportsItemLinksItem) SetSchema(schema *string) {
	l.Schema = schema
	l.require(listAccountExportsResponseExportsItemLinksItemFieldSchema)
}

// SetTargetSchema sets the TargetSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAccountExportsResponseExportsItemLinksItem) SetTargetSchema(targetSchema *string) {
	l.TargetSchema = targetSchema
	l.require(listAccountExportsResponseExportsItemLinksItemFieldTargetSchema)
}

func (l *ListAccountExportsResponseExportsItemLinksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler ListAccountExportsResponseExportsItemLinksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListAccountExportsResponseExportsItemLinksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListAccountExportsResponseExportsItemLinksItem) MarshalJSON() ([]byte, error) {
	type embed ListAccountExportsResponseExportsItemLinksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListAccountExportsResponseExportsItemLinksItem) String() string {
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
type ListAccountExportsResponseExportsItemLinksItemMethod string

const (
	ListAccountExportsResponseExportsItemLinksItemMethodGet     ListAccountExportsResponseExportsItemLinksItemMethod = "GET"
	ListAccountExportsResponseExportsItemLinksItemMethodPost    ListAccountExportsResponseExportsItemLinksItemMethod = "POST"
	ListAccountExportsResponseExportsItemLinksItemMethodPut     ListAccountExportsResponseExportsItemLinksItemMethod = "PUT"
	ListAccountExportsResponseExportsItemLinksItemMethodPatch   ListAccountExportsResponseExportsItemLinksItemMethod = "PATCH"
	ListAccountExportsResponseExportsItemLinksItemMethodDelete  ListAccountExportsResponseExportsItemLinksItemMethod = "DELETE"
	ListAccountExportsResponseExportsItemLinksItemMethodOptions ListAccountExportsResponseExportsItemLinksItemMethod = "OPTIONS"
	ListAccountExportsResponseExportsItemLinksItemMethodHead    ListAccountExportsResponseExportsItemLinksItemMethod = "HEAD"
)

func NewListAccountExportsResponseExportsItemLinksItemMethodFromString(s string) (ListAccountExportsResponseExportsItemLinksItemMethod, error) {
	switch s {
	case "GET":
		return ListAccountExportsResponseExportsItemLinksItemMethodGet, nil
	case "POST":
		return ListAccountExportsResponseExportsItemLinksItemMethodPost, nil
	case "PUT":
		return ListAccountExportsResponseExportsItemLinksItemMethodPut, nil
	case "PATCH":
		return ListAccountExportsResponseExportsItemLinksItemMethodPatch, nil
	case "DELETE":
		return ListAccountExportsResponseExportsItemLinksItemMethodDelete, nil
	case "OPTIONS":
		return ListAccountExportsResponseExportsItemLinksItemMethodOptions, nil
	case "HEAD":
		return ListAccountExportsResponseExportsItemLinksItemMethodHead, nil
	}
	var t ListAccountExportsResponseExportsItemLinksItemMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListAccountExportsResponseExportsItemLinksItemMethod) Ptr() *ListAccountExportsResponseExportsItemLinksItemMethod {
	return &l
}

// This object represents a link from the resource where it is found to another resource or action that may be performed.
var (
	listAccountExportsResponseLinksItemFieldHref         = big.NewInt(1 << 0)
	listAccountExportsResponseLinksItemFieldMethod       = big.NewInt(1 << 1)
	listAccountExportsResponseLinksItemFieldRel          = big.NewInt(1 << 2)
	listAccountExportsResponseLinksItemFieldSchema       = big.NewInt(1 << 3)
	listAccountExportsResponseLinksItemFieldTargetSchema = big.NewInt(1 << 4)
)

type ListAccountExportsResponseLinksItem struct {
	// This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.
	Href *string `json:"href,omitempty" url:"href,omitempty"`
	// The HTTP method that should be used when accessing the URL defined in 'href'.
	Method *ListAccountExportsResponseLinksItemMethod `json:"method,omitempty" url:"method,omitempty"`
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

func (l *ListAccountExportsResponseLinksItem) GetHref() *string {
	if l == nil {
		return nil
	}
	return l.Href
}

func (l *ListAccountExportsResponseLinksItem) GetMethod() *ListAccountExportsResponseLinksItemMethod {
	if l == nil {
		return nil
	}
	return l.Method
}

func (l *ListAccountExportsResponseLinksItem) GetRel() *string {
	if l == nil {
		return nil
	}
	return l.Rel
}

func (l *ListAccountExportsResponseLinksItem) GetSchema() *string {
	if l == nil {
		return nil
	}
	return l.Schema
}

func (l *ListAccountExportsResponseLinksItem) GetTargetSchema() *string {
	if l == nil {
		return nil
	}
	return l.TargetSchema
}

func (l *ListAccountExportsResponseLinksItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListAccountExportsResponseLinksItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetHref sets the Href field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAccountExportsResponseLinksItem) SetHref(href *string) {
	l.Href = href
	l.require(listAccountExportsResponseLinksItemFieldHref)
}

// SetMethod sets the Method field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAccountExportsResponseLinksItem) SetMethod(method *ListAccountExportsResponseLinksItemMethod) {
	l.Method = method
	l.require(listAccountExportsResponseLinksItemFieldMethod)
}

// SetRel sets the Rel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAccountExportsResponseLinksItem) SetRel(rel *string) {
	l.Rel = rel
	l.require(listAccountExportsResponseLinksItemFieldRel)
}

// SetSchema sets the Schema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAccountExportsResponseLinksItem) SetSchema(schema *string) {
	l.Schema = schema
	l.require(listAccountExportsResponseLinksItemFieldSchema)
}

// SetTargetSchema sets the TargetSchema field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAccountExportsResponseLinksItem) SetTargetSchema(targetSchema *string) {
	l.TargetSchema = targetSchema
	l.require(listAccountExportsResponseLinksItemFieldTargetSchema)
}

func (l *ListAccountExportsResponseLinksItem) UnmarshalJSON(data []byte) error {
	type unmarshaler ListAccountExportsResponseLinksItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListAccountExportsResponseLinksItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListAccountExportsResponseLinksItem) MarshalJSON() ([]byte, error) {
	type embed ListAccountExportsResponseLinksItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListAccountExportsResponseLinksItem) String() string {
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
type ListAccountExportsResponseLinksItemMethod string

const (
	ListAccountExportsResponseLinksItemMethodGet     ListAccountExportsResponseLinksItemMethod = "GET"
	ListAccountExportsResponseLinksItemMethodPost    ListAccountExportsResponseLinksItemMethod = "POST"
	ListAccountExportsResponseLinksItemMethodPut     ListAccountExportsResponseLinksItemMethod = "PUT"
	ListAccountExportsResponseLinksItemMethodPatch   ListAccountExportsResponseLinksItemMethod = "PATCH"
	ListAccountExportsResponseLinksItemMethodDelete  ListAccountExportsResponseLinksItemMethod = "DELETE"
	ListAccountExportsResponseLinksItemMethodOptions ListAccountExportsResponseLinksItemMethod = "OPTIONS"
	ListAccountExportsResponseLinksItemMethodHead    ListAccountExportsResponseLinksItemMethod = "HEAD"
)

func NewListAccountExportsResponseLinksItemMethodFromString(s string) (ListAccountExportsResponseLinksItemMethod, error) {
	switch s {
	case "GET":
		return ListAccountExportsResponseLinksItemMethodGet, nil
	case "POST":
		return ListAccountExportsResponseLinksItemMethodPost, nil
	case "PUT":
		return ListAccountExportsResponseLinksItemMethodPut, nil
	case "PATCH":
		return ListAccountExportsResponseLinksItemMethodPatch, nil
	case "DELETE":
		return ListAccountExportsResponseLinksItemMethodDelete, nil
	case "OPTIONS":
		return ListAccountExportsResponseLinksItemMethodOptions, nil
	case "HEAD":
		return ListAccountExportsResponseLinksItemMethodHead, nil
	}
	var t ListAccountExportsResponseLinksItemMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListAccountExportsResponseLinksItemMethod) Ptr() *ListAccountExportsResponseLinksItemMethod {
	return &l
}
