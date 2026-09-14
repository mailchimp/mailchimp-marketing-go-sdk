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
	createVerifiedDomainsRequestFieldVerificationEmail = big.NewInt(1 << 0)
)

type CreateVerifiedDomainsRequest struct {
	// The e-mail address at the domain you want to verify. This will receive a two-factor challenge to be used in the verify action.
	VerificationEmail string `json:"verification_email" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CreateVerifiedDomainsRequest) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetVerificationEmail sets the VerificationEmail field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateVerifiedDomainsRequest) SetVerificationEmail(verificationEmail string) {
	c.VerificationEmail = verificationEmail
	c.require(createVerifiedDomainsRequestFieldVerificationEmail)
}

func (c *CreateVerifiedDomainsRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateVerifiedDomainsRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*c = CreateVerifiedDomainsRequest(body)
	return nil
}

func (c *CreateVerifiedDomainsRequest) MarshalJSON() ([]byte, error) {
	type embed CreateVerifiedDomainsRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

var (
	createActionVerifyVerifiedDomainsRequestFieldDomainName = big.NewInt(1 << 0)
	createActionVerifyVerifiedDomainsRequestFieldCode       = big.NewInt(1 << 1)
)

type CreateActionVerifyVerifiedDomainsRequest struct {
	// The domain name.
	DomainName string `json:"-" url:"-"`
	// The code that was sent to the email address provided when adding a new domain to verify.
	Code string `json:"code" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (c *CreateActionVerifyVerifiedDomainsRequest) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetDomainName sets the DomainName field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateActionVerifyVerifiedDomainsRequest) SetDomainName(domainName string) {
	c.DomainName = domainName
	c.require(createActionVerifyVerifiedDomainsRequestFieldDomainName)
}

// SetCode sets the Code field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateActionVerifyVerifiedDomainsRequest) SetCode(code string) {
	c.Code = code
	c.require(createActionVerifyVerifiedDomainsRequestFieldCode)
}

func (c *CreateActionVerifyVerifiedDomainsRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateActionVerifyVerifiedDomainsRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*c = CreateActionVerifyVerifiedDomainsRequest(body)
	return nil
}

func (c *CreateActionVerifyVerifiedDomainsRequest) MarshalJSON() ([]byte, error) {
	type embed CreateActionVerifyVerifiedDomainsRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

var (
	deleteVerifiedDomainsRequestFieldDomainName = big.NewInt(1 << 0)
)

type DeleteVerifiedDomainsRequest struct {
	// The domain name.
	DomainName string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (d *DeleteVerifiedDomainsRequest) require(field *big.Int) {
	if d.explicitFields == nil {
		d.explicitFields = big.NewInt(0)
	}
	d.explicitFields.Or(d.explicitFields, field)
}

// SetDomainName sets the DomainName field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (d *DeleteVerifiedDomainsRequest) SetDomainName(domainName string) {
	d.DomainName = domainName
	d.require(deleteVerifiedDomainsRequestFieldDomainName)
}

var (
	getVerifiedDomainsRequestFieldDomainName = big.NewInt(1 << 0)
)

type GetVerifiedDomainsRequest struct {
	// The domain name.
	DomainName string `json:"-" url:"-"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetVerifiedDomainsRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetDomainName sets the DomainName field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetVerifiedDomainsRequest) SetDomainName(domainName string) {
	g.DomainName = domainName
	g.require(getVerifiedDomainsRequestFieldDomainName)
}

// The verified domains currently on the account.
var (
	createActionVerifyVerifiedDomainsResponseFieldAuthenticated       = big.NewInt(1 << 0)
	createActionVerifyVerifiedDomainsResponseFieldDomain              = big.NewInt(1 << 1)
	createActionVerifyVerifiedDomainsResponseFieldIsFreeEmailProvider = big.NewInt(1 << 2)
	createActionVerifyVerifiedDomainsResponseFieldStatus              = big.NewInt(1 << 3)
	createActionVerifyVerifiedDomainsResponseFieldVerificationEmail   = big.NewInt(1 << 4)
	createActionVerifyVerifiedDomainsResponseFieldVerificationSent    = big.NewInt(1 << 5)
	createActionVerifyVerifiedDomainsResponseFieldVerified            = big.NewInt(1 << 6)
)

type CreateActionVerifyVerifiedDomainsResponse struct {
	// Whether domain authentication is enabled for this domain.
	Authenticated *bool `json:"authenticated,omitempty" url:"authenticated,omitempty"`
	// The name of this domain.
	Domain *string `json:"domain,omitempty" url:"domain,omitempty"`
	// Returns whether the domain used is a public / free email provider. See [Limitations of Free Email Addresses](https://mailchimp.com/help/limitations-of-free-email-addresses/) for more details.
	IsFreeEmailProvider *bool `json:"is_free_email_provider,omitempty" url:"is_free_email_provider,omitempty"`
	// The Domain's current status.
	Status *CreateActionVerifyVerifiedDomainsResponseStatus `json:"status,omitempty" url:"status,omitempty"`
	// The e-mail address receiving the two-factor challenge for this domain.
	VerificationEmail *string `json:"verification_email,omitempty" url:"verification_email,omitempty"`
	// The date/time that the two-factor challenge was sent to the verification email.
	VerificationSent *time.Time `json:"verification_sent,omitempty" url:"verification_sent,omitempty"`
	// Whether the domain has been verified for sending.
	Verified *bool `json:"verified,omitempty" url:"verified,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateActionVerifyVerifiedDomainsResponse) GetAuthenticated() *bool {
	if c == nil {
		return nil
	}
	return c.Authenticated
}

func (c *CreateActionVerifyVerifiedDomainsResponse) GetDomain() *string {
	if c == nil {
		return nil
	}
	return c.Domain
}

func (c *CreateActionVerifyVerifiedDomainsResponse) GetIsFreeEmailProvider() *bool {
	if c == nil {
		return nil
	}
	return c.IsFreeEmailProvider
}

func (c *CreateActionVerifyVerifiedDomainsResponse) GetStatus() *CreateActionVerifyVerifiedDomainsResponseStatus {
	if c == nil {
		return nil
	}
	return c.Status
}

func (c *CreateActionVerifyVerifiedDomainsResponse) GetVerificationEmail() *string {
	if c == nil {
		return nil
	}
	return c.VerificationEmail
}

func (c *CreateActionVerifyVerifiedDomainsResponse) GetVerificationSent() *time.Time {
	if c == nil {
		return nil
	}
	return c.VerificationSent
}

func (c *CreateActionVerifyVerifiedDomainsResponse) GetVerified() *bool {
	if c == nil {
		return nil
	}
	return c.Verified
}

func (c *CreateActionVerifyVerifiedDomainsResponse) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *CreateActionVerifyVerifiedDomainsResponse) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetAuthenticated sets the Authenticated field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateActionVerifyVerifiedDomainsResponse) SetAuthenticated(authenticated *bool) {
	c.Authenticated = authenticated
	c.require(createActionVerifyVerifiedDomainsResponseFieldAuthenticated)
}

// SetDomain sets the Domain field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateActionVerifyVerifiedDomainsResponse) SetDomain(domain *string) {
	c.Domain = domain
	c.require(createActionVerifyVerifiedDomainsResponseFieldDomain)
}

// SetIsFreeEmailProvider sets the IsFreeEmailProvider field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateActionVerifyVerifiedDomainsResponse) SetIsFreeEmailProvider(isFreeEmailProvider *bool) {
	c.IsFreeEmailProvider = isFreeEmailProvider
	c.require(createActionVerifyVerifiedDomainsResponseFieldIsFreeEmailProvider)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateActionVerifyVerifiedDomainsResponse) SetStatus(status *CreateActionVerifyVerifiedDomainsResponseStatus) {
	c.Status = status
	c.require(createActionVerifyVerifiedDomainsResponseFieldStatus)
}

// SetVerificationEmail sets the VerificationEmail field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateActionVerifyVerifiedDomainsResponse) SetVerificationEmail(verificationEmail *string) {
	c.VerificationEmail = verificationEmail
	c.require(createActionVerifyVerifiedDomainsResponseFieldVerificationEmail)
}

// SetVerificationSent sets the VerificationSent field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateActionVerifyVerifiedDomainsResponse) SetVerificationSent(verificationSent *time.Time) {
	c.VerificationSent = verificationSent
	c.require(createActionVerifyVerifiedDomainsResponseFieldVerificationSent)
}

// SetVerified sets the Verified field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateActionVerifyVerifiedDomainsResponse) SetVerified(verified *bool) {
	c.Verified = verified
	c.require(createActionVerifyVerifiedDomainsResponseFieldVerified)
}

func (c *CreateActionVerifyVerifiedDomainsResponse) UnmarshalJSON(data []byte) error {
	type embed CreateActionVerifyVerifiedDomainsResponse
	var unmarshaler = struct {
		embed
		VerificationSent *internal.DateTime `json:"verification_sent,omitempty"`
	}{
		embed: embed(*c),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*c = CreateActionVerifyVerifiedDomainsResponse(unmarshaler.embed)
	c.VerificationSent = unmarshaler.VerificationSent.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateActionVerifyVerifiedDomainsResponse) MarshalJSON() ([]byte, error) {
	type embed CreateActionVerifyVerifiedDomainsResponse
	var marshaler = struct {
		embed
		VerificationSent *internal.DateTime `json:"verification_sent,omitempty"`
	}{
		embed:            embed(*c),
		VerificationSent: internal.NewOptionalDateTime(c.VerificationSent),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *CreateActionVerifyVerifiedDomainsResponse) String() string {
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

// The Domain's current status.
type CreateActionVerifyVerifiedDomainsResponseStatus string

const (
	CreateActionVerifyVerifiedDomainsResponseStatusVerificationInProgress   CreateActionVerifyVerifiedDomainsResponseStatus = "VERIFICATION_IN_PROGRESS"
	CreateActionVerifyVerifiedDomainsResponseStatusVerified                 CreateActionVerifyVerifiedDomainsResponseStatus = "VERIFIED"
	CreateActionVerifyVerifiedDomainsResponseStatusExpired                  CreateActionVerifyVerifiedDomainsResponseStatus = "EXPIRED"
	CreateActionVerifyVerifiedDomainsResponseStatusError                    CreateActionVerifyVerifiedDomainsResponseStatus = "ERROR"
	CreateActionVerifyVerifiedDomainsResponseStatusAuthenticationInProgress CreateActionVerifyVerifiedDomainsResponseStatus = "AUTHENTICATION_IN_PROGRESS"
	CreateActionVerifyVerifiedDomainsResponseStatusAuthenticationError      CreateActionVerifyVerifiedDomainsResponseStatus = "AUTHENTICATION_ERROR"
	CreateActionVerifyVerifiedDomainsResponseStatusAuthenticated            CreateActionVerifyVerifiedDomainsResponseStatus = "AUTHENTICATED"
)

func NewCreateActionVerifyVerifiedDomainsResponseStatusFromString(s string) (CreateActionVerifyVerifiedDomainsResponseStatus, error) {
	switch s {
	case "VERIFICATION_IN_PROGRESS":
		return CreateActionVerifyVerifiedDomainsResponseStatusVerificationInProgress, nil
	case "VERIFIED":
		return CreateActionVerifyVerifiedDomainsResponseStatusVerified, nil
	case "EXPIRED":
		return CreateActionVerifyVerifiedDomainsResponseStatusExpired, nil
	case "ERROR":
		return CreateActionVerifyVerifiedDomainsResponseStatusError, nil
	case "AUTHENTICATION_IN_PROGRESS":
		return CreateActionVerifyVerifiedDomainsResponseStatusAuthenticationInProgress, nil
	case "AUTHENTICATION_ERROR":
		return CreateActionVerifyVerifiedDomainsResponseStatusAuthenticationError, nil
	case "AUTHENTICATED":
		return CreateActionVerifyVerifiedDomainsResponseStatusAuthenticated, nil
	}
	var t CreateActionVerifyVerifiedDomainsResponseStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c CreateActionVerifyVerifiedDomainsResponseStatus) Ptr() *CreateActionVerifyVerifiedDomainsResponseStatus {
	return &c
}

// The verified domains currently on the account.
var (
	createVerifiedDomainsResponseFieldAuthenticated       = big.NewInt(1 << 0)
	createVerifiedDomainsResponseFieldDomain              = big.NewInt(1 << 1)
	createVerifiedDomainsResponseFieldIsFreeEmailProvider = big.NewInt(1 << 2)
	createVerifiedDomainsResponseFieldStatus              = big.NewInt(1 << 3)
	createVerifiedDomainsResponseFieldVerificationEmail   = big.NewInt(1 << 4)
	createVerifiedDomainsResponseFieldVerificationSent    = big.NewInt(1 << 5)
	createVerifiedDomainsResponseFieldVerified            = big.NewInt(1 << 6)
)

type CreateVerifiedDomainsResponse struct {
	// Whether domain authentication is enabled for this domain.
	Authenticated *bool `json:"authenticated,omitempty" url:"authenticated,omitempty"`
	// The name of this domain.
	Domain *string `json:"domain,omitempty" url:"domain,omitempty"`
	// Returns whether the domain used is a public / free email provider. See [Limitations of Free Email Addresses](https://mailchimp.com/help/limitations-of-free-email-addresses/) for more details.
	IsFreeEmailProvider *bool `json:"is_free_email_provider,omitempty" url:"is_free_email_provider,omitempty"`
	// The Domain's current status.
	Status *CreateVerifiedDomainsResponseStatus `json:"status,omitempty" url:"status,omitempty"`
	// The e-mail address receiving the two-factor challenge for this domain.
	VerificationEmail *string `json:"verification_email,omitempty" url:"verification_email,omitempty"`
	// The date/time that the two-factor challenge was sent to the verification email.
	VerificationSent *time.Time `json:"verification_sent,omitempty" url:"verification_sent,omitempty"`
	// Whether the domain has been verified for sending.
	Verified *bool `json:"verified,omitempty" url:"verified,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateVerifiedDomainsResponse) GetAuthenticated() *bool {
	if c == nil {
		return nil
	}
	return c.Authenticated
}

func (c *CreateVerifiedDomainsResponse) GetDomain() *string {
	if c == nil {
		return nil
	}
	return c.Domain
}

func (c *CreateVerifiedDomainsResponse) GetIsFreeEmailProvider() *bool {
	if c == nil {
		return nil
	}
	return c.IsFreeEmailProvider
}

func (c *CreateVerifiedDomainsResponse) GetStatus() *CreateVerifiedDomainsResponseStatus {
	if c == nil {
		return nil
	}
	return c.Status
}

func (c *CreateVerifiedDomainsResponse) GetVerificationEmail() *string {
	if c == nil {
		return nil
	}
	return c.VerificationEmail
}

func (c *CreateVerifiedDomainsResponse) GetVerificationSent() *time.Time {
	if c == nil {
		return nil
	}
	return c.VerificationSent
}

func (c *CreateVerifiedDomainsResponse) GetVerified() *bool {
	if c == nil {
		return nil
	}
	return c.Verified
}

func (c *CreateVerifiedDomainsResponse) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c *CreateVerifiedDomainsResponse) require(field *big.Int) {
	if c.explicitFields == nil {
		c.explicitFields = big.NewInt(0)
	}
	c.explicitFields.Or(c.explicitFields, field)
}

// SetAuthenticated sets the Authenticated field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateVerifiedDomainsResponse) SetAuthenticated(authenticated *bool) {
	c.Authenticated = authenticated
	c.require(createVerifiedDomainsResponseFieldAuthenticated)
}

// SetDomain sets the Domain field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateVerifiedDomainsResponse) SetDomain(domain *string) {
	c.Domain = domain
	c.require(createVerifiedDomainsResponseFieldDomain)
}

// SetIsFreeEmailProvider sets the IsFreeEmailProvider field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateVerifiedDomainsResponse) SetIsFreeEmailProvider(isFreeEmailProvider *bool) {
	c.IsFreeEmailProvider = isFreeEmailProvider
	c.require(createVerifiedDomainsResponseFieldIsFreeEmailProvider)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateVerifiedDomainsResponse) SetStatus(status *CreateVerifiedDomainsResponseStatus) {
	c.Status = status
	c.require(createVerifiedDomainsResponseFieldStatus)
}

// SetVerificationEmail sets the VerificationEmail field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateVerifiedDomainsResponse) SetVerificationEmail(verificationEmail *string) {
	c.VerificationEmail = verificationEmail
	c.require(createVerifiedDomainsResponseFieldVerificationEmail)
}

// SetVerificationSent sets the VerificationSent field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateVerifiedDomainsResponse) SetVerificationSent(verificationSent *time.Time) {
	c.VerificationSent = verificationSent
	c.require(createVerifiedDomainsResponseFieldVerificationSent)
}

// SetVerified sets the Verified field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (c *CreateVerifiedDomainsResponse) SetVerified(verified *bool) {
	c.Verified = verified
	c.require(createVerifiedDomainsResponseFieldVerified)
}

func (c *CreateVerifiedDomainsResponse) UnmarshalJSON(data []byte) error {
	type embed CreateVerifiedDomainsResponse
	var unmarshaler = struct {
		embed
		VerificationSent *internal.DateTime `json:"verification_sent,omitempty"`
	}{
		embed: embed(*c),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*c = CreateVerifiedDomainsResponse(unmarshaler.embed)
	c.VerificationSent = unmarshaler.VerificationSent.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateVerifiedDomainsResponse) MarshalJSON() ([]byte, error) {
	type embed CreateVerifiedDomainsResponse
	var marshaler = struct {
		embed
		VerificationSent *internal.DateTime `json:"verification_sent,omitempty"`
	}{
		embed:            embed(*c),
		VerificationSent: internal.NewOptionalDateTime(c.VerificationSent),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, c.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (c *CreateVerifiedDomainsResponse) String() string {
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

// The Domain's current status.
type CreateVerifiedDomainsResponseStatus string

const (
	CreateVerifiedDomainsResponseStatusVerificationInProgress   CreateVerifiedDomainsResponseStatus = "VERIFICATION_IN_PROGRESS"
	CreateVerifiedDomainsResponseStatusVerified                 CreateVerifiedDomainsResponseStatus = "VERIFIED"
	CreateVerifiedDomainsResponseStatusExpired                  CreateVerifiedDomainsResponseStatus = "EXPIRED"
	CreateVerifiedDomainsResponseStatusError                    CreateVerifiedDomainsResponseStatus = "ERROR"
	CreateVerifiedDomainsResponseStatusAuthenticationInProgress CreateVerifiedDomainsResponseStatus = "AUTHENTICATION_IN_PROGRESS"
	CreateVerifiedDomainsResponseStatusAuthenticationError      CreateVerifiedDomainsResponseStatus = "AUTHENTICATION_ERROR"
	CreateVerifiedDomainsResponseStatusAuthenticated            CreateVerifiedDomainsResponseStatus = "AUTHENTICATED"
)

func NewCreateVerifiedDomainsResponseStatusFromString(s string) (CreateVerifiedDomainsResponseStatus, error) {
	switch s {
	case "VERIFICATION_IN_PROGRESS":
		return CreateVerifiedDomainsResponseStatusVerificationInProgress, nil
	case "VERIFIED":
		return CreateVerifiedDomainsResponseStatusVerified, nil
	case "EXPIRED":
		return CreateVerifiedDomainsResponseStatusExpired, nil
	case "ERROR":
		return CreateVerifiedDomainsResponseStatusError, nil
	case "AUTHENTICATION_IN_PROGRESS":
		return CreateVerifiedDomainsResponseStatusAuthenticationInProgress, nil
	case "AUTHENTICATION_ERROR":
		return CreateVerifiedDomainsResponseStatusAuthenticationError, nil
	case "AUTHENTICATED":
		return CreateVerifiedDomainsResponseStatusAuthenticated, nil
	}
	var t CreateVerifiedDomainsResponseStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c CreateVerifiedDomainsResponseStatus) Ptr() *CreateVerifiedDomainsResponseStatus {
	return &c
}

// The verified domains currently on the account.
var (
	getVerifiedDomainsResponseFieldAuthenticated       = big.NewInt(1 << 0)
	getVerifiedDomainsResponseFieldDomain              = big.NewInt(1 << 1)
	getVerifiedDomainsResponseFieldIsFreeEmailProvider = big.NewInt(1 << 2)
	getVerifiedDomainsResponseFieldStatus              = big.NewInt(1 << 3)
	getVerifiedDomainsResponseFieldVerificationEmail   = big.NewInt(1 << 4)
	getVerifiedDomainsResponseFieldVerificationSent    = big.NewInt(1 << 5)
	getVerifiedDomainsResponseFieldVerified            = big.NewInt(1 << 6)
)

type GetVerifiedDomainsResponse struct {
	// Whether domain authentication is enabled for this domain.
	Authenticated *bool `json:"authenticated,omitempty" url:"authenticated,omitempty"`
	// The name of this domain.
	Domain *string `json:"domain,omitempty" url:"domain,omitempty"`
	// Returns whether the domain used is a public / free email provider. See [Limitations of Free Email Addresses](https://mailchimp.com/help/limitations-of-free-email-addresses/) for more details.
	IsFreeEmailProvider *bool `json:"is_free_email_provider,omitempty" url:"is_free_email_provider,omitempty"`
	// The Domain's current status.
	Status *GetVerifiedDomainsResponseStatus `json:"status,omitempty" url:"status,omitempty"`
	// The e-mail address receiving the two-factor challenge for this domain.
	VerificationEmail *string `json:"verification_email,omitempty" url:"verification_email,omitempty"`
	// The date/time that the two-factor challenge was sent to the verification email.
	VerificationSent *time.Time `json:"verification_sent,omitempty" url:"verification_sent,omitempty"`
	// Whether the domain has been verified for sending.
	Verified *bool `json:"verified,omitempty" url:"verified,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (g *GetVerifiedDomainsResponse) GetAuthenticated() *bool {
	if g == nil {
		return nil
	}
	return g.Authenticated
}

func (g *GetVerifiedDomainsResponse) GetDomain() *string {
	if g == nil {
		return nil
	}
	return g.Domain
}

func (g *GetVerifiedDomainsResponse) GetIsFreeEmailProvider() *bool {
	if g == nil {
		return nil
	}
	return g.IsFreeEmailProvider
}

func (g *GetVerifiedDomainsResponse) GetStatus() *GetVerifiedDomainsResponseStatus {
	if g == nil {
		return nil
	}
	return g.Status
}

func (g *GetVerifiedDomainsResponse) GetVerificationEmail() *string {
	if g == nil {
		return nil
	}
	return g.VerificationEmail
}

func (g *GetVerifiedDomainsResponse) GetVerificationSent() *time.Time {
	if g == nil {
		return nil
	}
	return g.VerificationSent
}

func (g *GetVerifiedDomainsResponse) GetVerified() *bool {
	if g == nil {
		return nil
	}
	return g.Verified
}

func (g *GetVerifiedDomainsResponse) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g *GetVerifiedDomainsResponse) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetAuthenticated sets the Authenticated field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetVerifiedDomainsResponse) SetAuthenticated(authenticated *bool) {
	g.Authenticated = authenticated
	g.require(getVerifiedDomainsResponseFieldAuthenticated)
}

// SetDomain sets the Domain field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetVerifiedDomainsResponse) SetDomain(domain *string) {
	g.Domain = domain
	g.require(getVerifiedDomainsResponseFieldDomain)
}

// SetIsFreeEmailProvider sets the IsFreeEmailProvider field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetVerifiedDomainsResponse) SetIsFreeEmailProvider(isFreeEmailProvider *bool) {
	g.IsFreeEmailProvider = isFreeEmailProvider
	g.require(getVerifiedDomainsResponseFieldIsFreeEmailProvider)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetVerifiedDomainsResponse) SetStatus(status *GetVerifiedDomainsResponseStatus) {
	g.Status = status
	g.require(getVerifiedDomainsResponseFieldStatus)
}

// SetVerificationEmail sets the VerificationEmail field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetVerifiedDomainsResponse) SetVerificationEmail(verificationEmail *string) {
	g.VerificationEmail = verificationEmail
	g.require(getVerifiedDomainsResponseFieldVerificationEmail)
}

// SetVerificationSent sets the VerificationSent field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetVerifiedDomainsResponse) SetVerificationSent(verificationSent *time.Time) {
	g.VerificationSent = verificationSent
	g.require(getVerifiedDomainsResponseFieldVerificationSent)
}

// SetVerified sets the Verified field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetVerifiedDomainsResponse) SetVerified(verified *bool) {
	g.Verified = verified
	g.require(getVerifiedDomainsResponseFieldVerified)
}

func (g *GetVerifiedDomainsResponse) UnmarshalJSON(data []byte) error {
	type embed GetVerifiedDomainsResponse
	var unmarshaler = struct {
		embed
		VerificationSent *internal.DateTime `json:"verification_sent,omitempty"`
	}{
		embed: embed(*g),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*g = GetVerifiedDomainsResponse(unmarshaler.embed)
	g.VerificationSent = unmarshaler.VerificationSent.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetVerifiedDomainsResponse) MarshalJSON() ([]byte, error) {
	type embed GetVerifiedDomainsResponse
	var marshaler = struct {
		embed
		VerificationSent *internal.DateTime `json:"verification_sent,omitempty"`
	}{
		embed:            embed(*g),
		VerificationSent: internal.NewOptionalDateTime(g.VerificationSent),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, g.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (g *GetVerifiedDomainsResponse) String() string {
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

// The Domain's current status.
type GetVerifiedDomainsResponseStatus string

const (
	GetVerifiedDomainsResponseStatusVerificationInProgress   GetVerifiedDomainsResponseStatus = "VERIFICATION_IN_PROGRESS"
	GetVerifiedDomainsResponseStatusVerified                 GetVerifiedDomainsResponseStatus = "VERIFIED"
	GetVerifiedDomainsResponseStatusExpired                  GetVerifiedDomainsResponseStatus = "EXPIRED"
	GetVerifiedDomainsResponseStatusError                    GetVerifiedDomainsResponseStatus = "ERROR"
	GetVerifiedDomainsResponseStatusAuthenticationInProgress GetVerifiedDomainsResponseStatus = "AUTHENTICATION_IN_PROGRESS"
	GetVerifiedDomainsResponseStatusAuthenticationError      GetVerifiedDomainsResponseStatus = "AUTHENTICATION_ERROR"
	GetVerifiedDomainsResponseStatusAuthenticated            GetVerifiedDomainsResponseStatus = "AUTHENTICATED"
)

func NewGetVerifiedDomainsResponseStatusFromString(s string) (GetVerifiedDomainsResponseStatus, error) {
	switch s {
	case "VERIFICATION_IN_PROGRESS":
		return GetVerifiedDomainsResponseStatusVerificationInProgress, nil
	case "VERIFIED":
		return GetVerifiedDomainsResponseStatusVerified, nil
	case "EXPIRED":
		return GetVerifiedDomainsResponseStatusExpired, nil
	case "ERROR":
		return GetVerifiedDomainsResponseStatusError, nil
	case "AUTHENTICATION_IN_PROGRESS":
		return GetVerifiedDomainsResponseStatusAuthenticationInProgress, nil
	case "AUTHENTICATION_ERROR":
		return GetVerifiedDomainsResponseStatusAuthenticationError, nil
	case "AUTHENTICATED":
		return GetVerifiedDomainsResponseStatusAuthenticated, nil
	}
	var t GetVerifiedDomainsResponseStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (g GetVerifiedDomainsResponseStatus) Ptr() *GetVerifiedDomainsResponseStatus {
	return &g
}

// The verified domains currently on the account.
var (
	listVerifiedDomainsResponseFieldDomains    = big.NewInt(1 << 0)
	listVerifiedDomainsResponseFieldTotalItems = big.NewInt(1 << 1)
)

type ListVerifiedDomainsResponse struct {
	// The domains on the account
	Domains []*ListVerifiedDomainsResponseDomainsItem `json:"domains,omitempty" url:"domains,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int `json:"total_items,omitempty" url:"total_items,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListVerifiedDomainsResponse) GetDomains() []*ListVerifiedDomainsResponseDomainsItem {
	if l == nil {
		return nil
	}
	return l.Domains
}

func (l *ListVerifiedDomainsResponse) GetTotalItems() *int {
	if l == nil {
		return nil
	}
	return l.TotalItems
}

func (l *ListVerifiedDomainsResponse) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListVerifiedDomainsResponse) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetDomains sets the Domains field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListVerifiedDomainsResponse) SetDomains(domains []*ListVerifiedDomainsResponseDomainsItem) {
	l.Domains = domains
	l.require(listVerifiedDomainsResponseFieldDomains)
}

// SetTotalItems sets the TotalItems field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListVerifiedDomainsResponse) SetTotalItems(totalItems *int) {
	l.TotalItems = totalItems
	l.require(listVerifiedDomainsResponseFieldTotalItems)
}

func (l *ListVerifiedDomainsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListVerifiedDomainsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListVerifiedDomainsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListVerifiedDomainsResponse) MarshalJSON() ([]byte, error) {
	type embed ListVerifiedDomainsResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListVerifiedDomainsResponse) String() string {
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

// The verified domains currently on the account.
var (
	listVerifiedDomainsResponseDomainsItemFieldAuthenticated       = big.NewInt(1 << 0)
	listVerifiedDomainsResponseDomainsItemFieldDomain              = big.NewInt(1 << 1)
	listVerifiedDomainsResponseDomainsItemFieldIsFreeEmailProvider = big.NewInt(1 << 2)
	listVerifiedDomainsResponseDomainsItemFieldStatus              = big.NewInt(1 << 3)
	listVerifiedDomainsResponseDomainsItemFieldVerificationEmail   = big.NewInt(1 << 4)
	listVerifiedDomainsResponseDomainsItemFieldVerificationSent    = big.NewInt(1 << 5)
	listVerifiedDomainsResponseDomainsItemFieldVerified            = big.NewInt(1 << 6)
)

type ListVerifiedDomainsResponseDomainsItem struct {
	// Whether domain authentication is enabled for this domain.
	Authenticated *bool `json:"authenticated,omitempty" url:"authenticated,omitempty"`
	// The name of this domain.
	Domain *string `json:"domain,omitempty" url:"domain,omitempty"`
	// Returns whether the domain used is a public / free email provider. See [Limitations of Free Email Addresses](https://mailchimp.com/help/limitations-of-free-email-addresses/) for more details.
	IsFreeEmailProvider *bool `json:"is_free_email_provider,omitempty" url:"is_free_email_provider,omitempty"`
	// The Domain's current status.
	Status *ListVerifiedDomainsResponseDomainsItemStatus `json:"status,omitempty" url:"status,omitempty"`
	// The e-mail address receiving the two-factor challenge for this domain.
	VerificationEmail *string `json:"verification_email,omitempty" url:"verification_email,omitempty"`
	// The date/time that the two-factor challenge was sent to the verification email.
	VerificationSent *time.Time `json:"verification_sent,omitempty" url:"verification_sent,omitempty"`
	// Whether the domain has been verified for sending.
	Verified *bool `json:"verified,omitempty" url:"verified,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListVerifiedDomainsResponseDomainsItem) GetAuthenticated() *bool {
	if l == nil {
		return nil
	}
	return l.Authenticated
}

func (l *ListVerifiedDomainsResponseDomainsItem) GetDomain() *string {
	if l == nil {
		return nil
	}
	return l.Domain
}

func (l *ListVerifiedDomainsResponseDomainsItem) GetIsFreeEmailProvider() *bool {
	if l == nil {
		return nil
	}
	return l.IsFreeEmailProvider
}

func (l *ListVerifiedDomainsResponseDomainsItem) GetStatus() *ListVerifiedDomainsResponseDomainsItemStatus {
	if l == nil {
		return nil
	}
	return l.Status
}

func (l *ListVerifiedDomainsResponseDomainsItem) GetVerificationEmail() *string {
	if l == nil {
		return nil
	}
	return l.VerificationEmail
}

func (l *ListVerifiedDomainsResponseDomainsItem) GetVerificationSent() *time.Time {
	if l == nil {
		return nil
	}
	return l.VerificationSent
}

func (l *ListVerifiedDomainsResponseDomainsItem) GetVerified() *bool {
	if l == nil {
		return nil
	}
	return l.Verified
}

func (l *ListVerifiedDomainsResponseDomainsItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListVerifiedDomainsResponseDomainsItem) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetAuthenticated sets the Authenticated field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListVerifiedDomainsResponseDomainsItem) SetAuthenticated(authenticated *bool) {
	l.Authenticated = authenticated
	l.require(listVerifiedDomainsResponseDomainsItemFieldAuthenticated)
}

// SetDomain sets the Domain field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListVerifiedDomainsResponseDomainsItem) SetDomain(domain *string) {
	l.Domain = domain
	l.require(listVerifiedDomainsResponseDomainsItemFieldDomain)
}

// SetIsFreeEmailProvider sets the IsFreeEmailProvider field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListVerifiedDomainsResponseDomainsItem) SetIsFreeEmailProvider(isFreeEmailProvider *bool) {
	l.IsFreeEmailProvider = isFreeEmailProvider
	l.require(listVerifiedDomainsResponseDomainsItemFieldIsFreeEmailProvider)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListVerifiedDomainsResponseDomainsItem) SetStatus(status *ListVerifiedDomainsResponseDomainsItemStatus) {
	l.Status = status
	l.require(listVerifiedDomainsResponseDomainsItemFieldStatus)
}

// SetVerificationEmail sets the VerificationEmail field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListVerifiedDomainsResponseDomainsItem) SetVerificationEmail(verificationEmail *string) {
	l.VerificationEmail = verificationEmail
	l.require(listVerifiedDomainsResponseDomainsItemFieldVerificationEmail)
}

// SetVerificationSent sets the VerificationSent field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListVerifiedDomainsResponseDomainsItem) SetVerificationSent(verificationSent *time.Time) {
	l.VerificationSent = verificationSent
	l.require(listVerifiedDomainsResponseDomainsItemFieldVerificationSent)
}

// SetVerified sets the Verified field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListVerifiedDomainsResponseDomainsItem) SetVerified(verified *bool) {
	l.Verified = verified
	l.require(listVerifiedDomainsResponseDomainsItemFieldVerified)
}

func (l *ListVerifiedDomainsResponseDomainsItem) UnmarshalJSON(data []byte) error {
	type embed ListVerifiedDomainsResponseDomainsItem
	var unmarshaler = struct {
		embed
		VerificationSent *internal.DateTime `json:"verification_sent,omitempty"`
	}{
		embed: embed(*l),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*l = ListVerifiedDomainsResponseDomainsItem(unmarshaler.embed)
	l.VerificationSent = unmarshaler.VerificationSent.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListVerifiedDomainsResponseDomainsItem) MarshalJSON() ([]byte, error) {
	type embed ListVerifiedDomainsResponseDomainsItem
	var marshaler = struct {
		embed
		VerificationSent *internal.DateTime `json:"verification_sent,omitempty"`
	}{
		embed:            embed(*l),
		VerificationSent: internal.NewOptionalDateTime(l.VerificationSent),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListVerifiedDomainsResponseDomainsItem) String() string {
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

// The Domain's current status.
type ListVerifiedDomainsResponseDomainsItemStatus string

const (
	ListVerifiedDomainsResponseDomainsItemStatusVerificationInProgress   ListVerifiedDomainsResponseDomainsItemStatus = "VERIFICATION_IN_PROGRESS"
	ListVerifiedDomainsResponseDomainsItemStatusVerified                 ListVerifiedDomainsResponseDomainsItemStatus = "VERIFIED"
	ListVerifiedDomainsResponseDomainsItemStatusExpired                  ListVerifiedDomainsResponseDomainsItemStatus = "EXPIRED"
	ListVerifiedDomainsResponseDomainsItemStatusError                    ListVerifiedDomainsResponseDomainsItemStatus = "ERROR"
	ListVerifiedDomainsResponseDomainsItemStatusAuthenticationInProgress ListVerifiedDomainsResponseDomainsItemStatus = "AUTHENTICATION_IN_PROGRESS"
	ListVerifiedDomainsResponseDomainsItemStatusAuthenticationError      ListVerifiedDomainsResponseDomainsItemStatus = "AUTHENTICATION_ERROR"
	ListVerifiedDomainsResponseDomainsItemStatusAuthenticated            ListVerifiedDomainsResponseDomainsItemStatus = "AUTHENTICATED"
)

func NewListVerifiedDomainsResponseDomainsItemStatusFromString(s string) (ListVerifiedDomainsResponseDomainsItemStatus, error) {
	switch s {
	case "VERIFICATION_IN_PROGRESS":
		return ListVerifiedDomainsResponseDomainsItemStatusVerificationInProgress, nil
	case "VERIFIED":
		return ListVerifiedDomainsResponseDomainsItemStatusVerified, nil
	case "EXPIRED":
		return ListVerifiedDomainsResponseDomainsItemStatusExpired, nil
	case "ERROR":
		return ListVerifiedDomainsResponseDomainsItemStatusError, nil
	case "AUTHENTICATION_IN_PROGRESS":
		return ListVerifiedDomainsResponseDomainsItemStatusAuthenticationInProgress, nil
	case "AUTHENTICATION_ERROR":
		return ListVerifiedDomainsResponseDomainsItemStatusAuthenticationError, nil
	case "AUTHENTICATED":
		return ListVerifiedDomainsResponseDomainsItemStatusAuthenticated, nil
	}
	var t ListVerifiedDomainsResponseDomainsItemStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l ListVerifiedDomainsResponseDomainsItemStatus) Ptr() *ListVerifiedDomainsResponseDomainsItemStatus {
	return &l
}
