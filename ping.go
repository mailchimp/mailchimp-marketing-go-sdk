// Code generated from our API definition. DO NOT EDIT.

package api

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/mailchimp/mailchimp-marketing-go-sdk/internal"
	big "math/big"
)

// API health status.
var (
	listPingResponseFieldHealthStatus = big.NewInt(1 << 0)
)

type ListPingResponse struct {
	// This will return a constant string value if the request is successful. Ex. "Everything's Chimpy!"
	HealthStatus *string `json:"health_status,omitempty" url:"health_status,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListPingResponse) GetHealthStatus() *string {
	if l == nil {
		return nil
	}
	return l.HealthStatus
}

func (l *ListPingResponse) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l *ListPingResponse) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetHealthStatus sets the HealthStatus field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListPingResponse) SetHealthStatus(healthStatus *string) {
	l.HealthStatus = healthStatus
	l.require(listPingResponseFieldHealthStatus)
}

func (l *ListPingResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListPingResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListPingResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListPingResponse) MarshalJSON() ([]byte, error) {
	type embed ListPingResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	explicitMarshaler := internal.HandleExplicitFields(marshaler, l.explicitFields)
	return json.Marshal(explicitMarshaler)
}

func (l *ListPingResponse) String() string {
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
