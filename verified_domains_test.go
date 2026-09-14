// Code generated from our API definition. DO NOT EDIT.

package api

import (
	json "encoding/json"
	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
	testing "testing"
	time "time"
)

func TestSettersCreateVerifiedDomainsRequest(t *testing.T) {
	t.Run("SetVerificationEmail", func(t *testing.T) {
		obj := &CreateVerifiedDomainsRequest{}
		var fernTestValueVerificationEmail string
		obj.SetVerificationEmail(fernTestValueVerificationEmail)
		assert.Equal(t, fernTestValueVerificationEmail, obj.VerificationEmail)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitCreateVerifiedDomainsRequest(t *testing.T) {
	t.Run("SetVerificationEmail_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateVerifiedDomainsRequest{}
		var fernTestValueVerificationEmail string

		// Act
		obj.SetVerificationEmail(fernTestValueVerificationEmail)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersCreateActionVerifyVerifiedDomainsRequest(t *testing.T) {
	t.Run("SetDomainName", func(t *testing.T) {
		obj := &CreateActionVerifyVerifiedDomainsRequest{}
		var fernTestValueDomainName string
		obj.SetDomainName(fernTestValueDomainName)
		assert.Equal(t, fernTestValueDomainName, obj.DomainName)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCode", func(t *testing.T) {
		obj := &CreateActionVerifyVerifiedDomainsRequest{}
		var fernTestValueCode string
		obj.SetCode(fernTestValueCode)
		assert.Equal(t, fernTestValueCode, obj.Code)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitCreateActionVerifyVerifiedDomainsRequest(t *testing.T) {
	t.Run("SetDomainName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateActionVerifyVerifiedDomainsRequest{}
		var fernTestValueDomainName string

		// Act
		obj.SetDomainName(fernTestValueDomainName)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetCode_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateActionVerifyVerifiedDomainsRequest{}
		var fernTestValueCode string

		// Act
		obj.SetCode(fernTestValueCode)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersDeleteVerifiedDomainsRequest(t *testing.T) {
	t.Run("SetDomainName", func(t *testing.T) {
		obj := &DeleteVerifiedDomainsRequest{}
		var fernTestValueDomainName string
		obj.SetDomainName(fernTestValueDomainName)
		assert.Equal(t, fernTestValueDomainName, obj.DomainName)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitDeleteVerifiedDomainsRequest(t *testing.T) {
	t.Run("SetDomainName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &DeleteVerifiedDomainsRequest{}
		var fernTestValueDomainName string

		// Act
		obj.SetDomainName(fernTestValueDomainName)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersGetVerifiedDomainsRequest(t *testing.T) {
	t.Run("SetDomainName", func(t *testing.T) {
		obj := &GetVerifiedDomainsRequest{}
		var fernTestValueDomainName string
		obj.SetDomainName(fernTestValueDomainName)
		assert.Equal(t, fernTestValueDomainName, obj.DomainName)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitGetVerifiedDomainsRequest(t *testing.T) {
	t.Run("SetDomainName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetVerifiedDomainsRequest{}
		var fernTestValueDomainName string

		// Act
		obj.SetDomainName(fernTestValueDomainName)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersCreateActionVerifyVerifiedDomainsResponse(t *testing.T) {
	t.Run("SetAuthenticated", func(t *testing.T) {
		obj := &CreateActionVerifyVerifiedDomainsResponse{}
		var fernTestValueAuthenticated *bool
		obj.SetAuthenticated(fernTestValueAuthenticated)
		assert.Equal(t, fernTestValueAuthenticated, obj.Authenticated)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetDomain", func(t *testing.T) {
		obj := &CreateActionVerifyVerifiedDomainsResponse{}
		var fernTestValueDomain *string
		obj.SetDomain(fernTestValueDomain)
		assert.Equal(t, fernTestValueDomain, obj.Domain)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetIsFreeEmailProvider", func(t *testing.T) {
		obj := &CreateActionVerifyVerifiedDomainsResponse{}
		var fernTestValueIsFreeEmailProvider *bool
		obj.SetIsFreeEmailProvider(fernTestValueIsFreeEmailProvider)
		assert.Equal(t, fernTestValueIsFreeEmailProvider, obj.IsFreeEmailProvider)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStatus", func(t *testing.T) {
		obj := &CreateActionVerifyVerifiedDomainsResponse{}
		var fernTestValueStatus *CreateActionVerifyVerifiedDomainsResponseStatus
		obj.SetStatus(fernTestValueStatus)
		assert.Equal(t, fernTestValueStatus, obj.Status)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetVerificationEmail", func(t *testing.T) {
		obj := &CreateActionVerifyVerifiedDomainsResponse{}
		var fernTestValueVerificationEmail *string
		obj.SetVerificationEmail(fernTestValueVerificationEmail)
		assert.Equal(t, fernTestValueVerificationEmail, obj.VerificationEmail)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetVerificationSent", func(t *testing.T) {
		obj := &CreateActionVerifyVerifiedDomainsResponse{}
		var fernTestValueVerificationSent *time.Time
		obj.SetVerificationSent(fernTestValueVerificationSent)
		assert.Equal(t, fernTestValueVerificationSent, obj.VerificationSent)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetVerified", func(t *testing.T) {
		obj := &CreateActionVerifyVerifiedDomainsResponse{}
		var fernTestValueVerified *bool
		obj.SetVerified(fernTestValueVerified)
		assert.Equal(t, fernTestValueVerified, obj.Verified)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersCreateActionVerifyVerifiedDomainsResponse(t *testing.T) {
	t.Run("GetAuthenticated", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateActionVerifyVerifiedDomainsResponse{}
		var expected *bool
		obj.Authenticated = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetAuthenticated(), "getter should return the property value")
	})

	t.Run("GetAuthenticated_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateActionVerifyVerifiedDomainsResponse{}
		obj.Authenticated = nil

		// Act & Assert
		assert.Nil(t, obj.GetAuthenticated(), "getter should return nil when property is nil")
	})

	t.Run("GetAuthenticated_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateActionVerifyVerifiedDomainsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetAuthenticated() // Should return zero value
	})

	t.Run("GetDomain", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateActionVerifyVerifiedDomainsResponse{}
		var expected *string
		obj.Domain = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetDomain(), "getter should return the property value")
	})

	t.Run("GetDomain_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateActionVerifyVerifiedDomainsResponse{}
		obj.Domain = nil

		// Act & Assert
		assert.Nil(t, obj.GetDomain(), "getter should return nil when property is nil")
	})

	t.Run("GetDomain_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateActionVerifyVerifiedDomainsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetDomain() // Should return zero value
	})

	t.Run("GetIsFreeEmailProvider", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateActionVerifyVerifiedDomainsResponse{}
		var expected *bool
		obj.IsFreeEmailProvider = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetIsFreeEmailProvider(), "getter should return the property value")
	})

	t.Run("GetIsFreeEmailProvider_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateActionVerifyVerifiedDomainsResponse{}
		obj.IsFreeEmailProvider = nil

		// Act & Assert
		assert.Nil(t, obj.GetIsFreeEmailProvider(), "getter should return nil when property is nil")
	})

	t.Run("GetIsFreeEmailProvider_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateActionVerifyVerifiedDomainsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetIsFreeEmailProvider() // Should return zero value
	})

	t.Run("GetStatus", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateActionVerifyVerifiedDomainsResponse{}
		var expected *CreateActionVerifyVerifiedDomainsResponseStatus
		obj.Status = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStatus(), "getter should return the property value")
	})

	t.Run("GetStatus_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateActionVerifyVerifiedDomainsResponse{}
		obj.Status = nil

		// Act & Assert
		assert.Nil(t, obj.GetStatus(), "getter should return nil when property is nil")
	})

	t.Run("GetStatus_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateActionVerifyVerifiedDomainsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStatus() // Should return zero value
	})

	t.Run("GetVerificationEmail", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateActionVerifyVerifiedDomainsResponse{}
		var expected *string
		obj.VerificationEmail = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetVerificationEmail(), "getter should return the property value")
	})

	t.Run("GetVerificationEmail_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateActionVerifyVerifiedDomainsResponse{}
		obj.VerificationEmail = nil

		// Act & Assert
		assert.Nil(t, obj.GetVerificationEmail(), "getter should return nil when property is nil")
	})

	t.Run("GetVerificationEmail_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateActionVerifyVerifiedDomainsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetVerificationEmail() // Should return zero value
	})

	t.Run("GetVerificationSent", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateActionVerifyVerifiedDomainsResponse{}
		var expected *time.Time
		obj.VerificationSent = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetVerificationSent(), "getter should return the property value")
	})

	t.Run("GetVerificationSent_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateActionVerifyVerifiedDomainsResponse{}
		obj.VerificationSent = nil

		// Act & Assert
		assert.Nil(t, obj.GetVerificationSent(), "getter should return nil when property is nil")
	})

	t.Run("GetVerificationSent_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateActionVerifyVerifiedDomainsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetVerificationSent() // Should return zero value
	})

	t.Run("GetVerified", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateActionVerifyVerifiedDomainsResponse{}
		var expected *bool
		obj.Verified = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetVerified(), "getter should return the property value")
	})

	t.Run("GetVerified_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateActionVerifyVerifiedDomainsResponse{}
		obj.Verified = nil

		// Act & Assert
		assert.Nil(t, obj.GetVerified(), "getter should return nil when property is nil")
	})

	t.Run("GetVerified_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateActionVerifyVerifiedDomainsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetVerified() // Should return zero value
	})

}

func TestSettersMarkExplicitCreateActionVerifyVerifiedDomainsResponse(t *testing.T) {
	t.Run("SetAuthenticated_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateActionVerifyVerifiedDomainsResponse{}
		var fernTestValueAuthenticated *bool

		// Act
		obj.SetAuthenticated(fernTestValueAuthenticated)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetDomain_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateActionVerifyVerifiedDomainsResponse{}
		var fernTestValueDomain *string

		// Act
		obj.SetDomain(fernTestValueDomain)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetIsFreeEmailProvider_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateActionVerifyVerifiedDomainsResponse{}
		var fernTestValueIsFreeEmailProvider *bool

		// Act
		obj.SetIsFreeEmailProvider(fernTestValueIsFreeEmailProvider)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetStatus_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateActionVerifyVerifiedDomainsResponse{}
		var fernTestValueStatus *CreateActionVerifyVerifiedDomainsResponseStatus

		// Act
		obj.SetStatus(fernTestValueStatus)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetVerificationEmail_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateActionVerifyVerifiedDomainsResponse{}
		var fernTestValueVerificationEmail *string

		// Act
		obj.SetVerificationEmail(fernTestValueVerificationEmail)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetVerificationSent_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateActionVerifyVerifiedDomainsResponse{}
		var fernTestValueVerificationSent *time.Time

		// Act
		obj.SetVerificationSent(fernTestValueVerificationSent)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetVerified_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateActionVerifyVerifiedDomainsResponse{}
		var fernTestValueVerified *bool

		// Act
		obj.SetVerified(fernTestValueVerified)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersCreateVerifiedDomainsResponse(t *testing.T) {
	t.Run("SetAuthenticated", func(t *testing.T) {
		obj := &CreateVerifiedDomainsResponse{}
		var fernTestValueAuthenticated *bool
		obj.SetAuthenticated(fernTestValueAuthenticated)
		assert.Equal(t, fernTestValueAuthenticated, obj.Authenticated)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetDomain", func(t *testing.T) {
		obj := &CreateVerifiedDomainsResponse{}
		var fernTestValueDomain *string
		obj.SetDomain(fernTestValueDomain)
		assert.Equal(t, fernTestValueDomain, obj.Domain)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetIsFreeEmailProvider", func(t *testing.T) {
		obj := &CreateVerifiedDomainsResponse{}
		var fernTestValueIsFreeEmailProvider *bool
		obj.SetIsFreeEmailProvider(fernTestValueIsFreeEmailProvider)
		assert.Equal(t, fernTestValueIsFreeEmailProvider, obj.IsFreeEmailProvider)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStatus", func(t *testing.T) {
		obj := &CreateVerifiedDomainsResponse{}
		var fernTestValueStatus *CreateVerifiedDomainsResponseStatus
		obj.SetStatus(fernTestValueStatus)
		assert.Equal(t, fernTestValueStatus, obj.Status)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetVerificationEmail", func(t *testing.T) {
		obj := &CreateVerifiedDomainsResponse{}
		var fernTestValueVerificationEmail *string
		obj.SetVerificationEmail(fernTestValueVerificationEmail)
		assert.Equal(t, fernTestValueVerificationEmail, obj.VerificationEmail)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetVerificationSent", func(t *testing.T) {
		obj := &CreateVerifiedDomainsResponse{}
		var fernTestValueVerificationSent *time.Time
		obj.SetVerificationSent(fernTestValueVerificationSent)
		assert.Equal(t, fernTestValueVerificationSent, obj.VerificationSent)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetVerified", func(t *testing.T) {
		obj := &CreateVerifiedDomainsResponse{}
		var fernTestValueVerified *bool
		obj.SetVerified(fernTestValueVerified)
		assert.Equal(t, fernTestValueVerified, obj.Verified)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersCreateVerifiedDomainsResponse(t *testing.T) {
	t.Run("GetAuthenticated", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateVerifiedDomainsResponse{}
		var expected *bool
		obj.Authenticated = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetAuthenticated(), "getter should return the property value")
	})

	t.Run("GetAuthenticated_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateVerifiedDomainsResponse{}
		obj.Authenticated = nil

		// Act & Assert
		assert.Nil(t, obj.GetAuthenticated(), "getter should return nil when property is nil")
	})

	t.Run("GetAuthenticated_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateVerifiedDomainsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetAuthenticated() // Should return zero value
	})

	t.Run("GetDomain", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateVerifiedDomainsResponse{}
		var expected *string
		obj.Domain = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetDomain(), "getter should return the property value")
	})

	t.Run("GetDomain_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateVerifiedDomainsResponse{}
		obj.Domain = nil

		// Act & Assert
		assert.Nil(t, obj.GetDomain(), "getter should return nil when property is nil")
	})

	t.Run("GetDomain_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateVerifiedDomainsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetDomain() // Should return zero value
	})

	t.Run("GetIsFreeEmailProvider", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateVerifiedDomainsResponse{}
		var expected *bool
		obj.IsFreeEmailProvider = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetIsFreeEmailProvider(), "getter should return the property value")
	})

	t.Run("GetIsFreeEmailProvider_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateVerifiedDomainsResponse{}
		obj.IsFreeEmailProvider = nil

		// Act & Assert
		assert.Nil(t, obj.GetIsFreeEmailProvider(), "getter should return nil when property is nil")
	})

	t.Run("GetIsFreeEmailProvider_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateVerifiedDomainsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetIsFreeEmailProvider() // Should return zero value
	})

	t.Run("GetStatus", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateVerifiedDomainsResponse{}
		var expected *CreateVerifiedDomainsResponseStatus
		obj.Status = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStatus(), "getter should return the property value")
	})

	t.Run("GetStatus_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateVerifiedDomainsResponse{}
		obj.Status = nil

		// Act & Assert
		assert.Nil(t, obj.GetStatus(), "getter should return nil when property is nil")
	})

	t.Run("GetStatus_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateVerifiedDomainsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStatus() // Should return zero value
	})

	t.Run("GetVerificationEmail", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateVerifiedDomainsResponse{}
		var expected *string
		obj.VerificationEmail = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetVerificationEmail(), "getter should return the property value")
	})

	t.Run("GetVerificationEmail_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateVerifiedDomainsResponse{}
		obj.VerificationEmail = nil

		// Act & Assert
		assert.Nil(t, obj.GetVerificationEmail(), "getter should return nil when property is nil")
	})

	t.Run("GetVerificationEmail_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateVerifiedDomainsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetVerificationEmail() // Should return zero value
	})

	t.Run("GetVerificationSent", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateVerifiedDomainsResponse{}
		var expected *time.Time
		obj.VerificationSent = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetVerificationSent(), "getter should return the property value")
	})

	t.Run("GetVerificationSent_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateVerifiedDomainsResponse{}
		obj.VerificationSent = nil

		// Act & Assert
		assert.Nil(t, obj.GetVerificationSent(), "getter should return nil when property is nil")
	})

	t.Run("GetVerificationSent_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateVerifiedDomainsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetVerificationSent() // Should return zero value
	})

	t.Run("GetVerified", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateVerifiedDomainsResponse{}
		var expected *bool
		obj.Verified = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetVerified(), "getter should return the property value")
	})

	t.Run("GetVerified_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateVerifiedDomainsResponse{}
		obj.Verified = nil

		// Act & Assert
		assert.Nil(t, obj.GetVerified(), "getter should return nil when property is nil")
	})

	t.Run("GetVerified_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateVerifiedDomainsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetVerified() // Should return zero value
	})

}

func TestSettersMarkExplicitCreateVerifiedDomainsResponse(t *testing.T) {
	t.Run("SetAuthenticated_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateVerifiedDomainsResponse{}
		var fernTestValueAuthenticated *bool

		// Act
		obj.SetAuthenticated(fernTestValueAuthenticated)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetDomain_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateVerifiedDomainsResponse{}
		var fernTestValueDomain *string

		// Act
		obj.SetDomain(fernTestValueDomain)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetIsFreeEmailProvider_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateVerifiedDomainsResponse{}
		var fernTestValueIsFreeEmailProvider *bool

		// Act
		obj.SetIsFreeEmailProvider(fernTestValueIsFreeEmailProvider)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetStatus_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateVerifiedDomainsResponse{}
		var fernTestValueStatus *CreateVerifiedDomainsResponseStatus

		// Act
		obj.SetStatus(fernTestValueStatus)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetVerificationEmail_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateVerifiedDomainsResponse{}
		var fernTestValueVerificationEmail *string

		// Act
		obj.SetVerificationEmail(fernTestValueVerificationEmail)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetVerificationSent_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateVerifiedDomainsResponse{}
		var fernTestValueVerificationSent *time.Time

		// Act
		obj.SetVerificationSent(fernTestValueVerificationSent)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetVerified_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateVerifiedDomainsResponse{}
		var fernTestValueVerified *bool

		// Act
		obj.SetVerified(fernTestValueVerified)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersGetVerifiedDomainsResponse(t *testing.T) {
	t.Run("SetAuthenticated", func(t *testing.T) {
		obj := &GetVerifiedDomainsResponse{}
		var fernTestValueAuthenticated *bool
		obj.SetAuthenticated(fernTestValueAuthenticated)
		assert.Equal(t, fernTestValueAuthenticated, obj.Authenticated)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetDomain", func(t *testing.T) {
		obj := &GetVerifiedDomainsResponse{}
		var fernTestValueDomain *string
		obj.SetDomain(fernTestValueDomain)
		assert.Equal(t, fernTestValueDomain, obj.Domain)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetIsFreeEmailProvider", func(t *testing.T) {
		obj := &GetVerifiedDomainsResponse{}
		var fernTestValueIsFreeEmailProvider *bool
		obj.SetIsFreeEmailProvider(fernTestValueIsFreeEmailProvider)
		assert.Equal(t, fernTestValueIsFreeEmailProvider, obj.IsFreeEmailProvider)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStatus", func(t *testing.T) {
		obj := &GetVerifiedDomainsResponse{}
		var fernTestValueStatus *GetVerifiedDomainsResponseStatus
		obj.SetStatus(fernTestValueStatus)
		assert.Equal(t, fernTestValueStatus, obj.Status)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetVerificationEmail", func(t *testing.T) {
		obj := &GetVerifiedDomainsResponse{}
		var fernTestValueVerificationEmail *string
		obj.SetVerificationEmail(fernTestValueVerificationEmail)
		assert.Equal(t, fernTestValueVerificationEmail, obj.VerificationEmail)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetVerificationSent", func(t *testing.T) {
		obj := &GetVerifiedDomainsResponse{}
		var fernTestValueVerificationSent *time.Time
		obj.SetVerificationSent(fernTestValueVerificationSent)
		assert.Equal(t, fernTestValueVerificationSent, obj.VerificationSent)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetVerified", func(t *testing.T) {
		obj := &GetVerifiedDomainsResponse{}
		var fernTestValueVerified *bool
		obj.SetVerified(fernTestValueVerified)
		assert.Equal(t, fernTestValueVerified, obj.Verified)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersGetVerifiedDomainsResponse(t *testing.T) {
	t.Run("GetAuthenticated", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetVerifiedDomainsResponse{}
		var expected *bool
		obj.Authenticated = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetAuthenticated(), "getter should return the property value")
	})

	t.Run("GetAuthenticated_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetVerifiedDomainsResponse{}
		obj.Authenticated = nil

		// Act & Assert
		assert.Nil(t, obj.GetAuthenticated(), "getter should return nil when property is nil")
	})

	t.Run("GetAuthenticated_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetVerifiedDomainsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetAuthenticated() // Should return zero value
	})

	t.Run("GetDomain", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetVerifiedDomainsResponse{}
		var expected *string
		obj.Domain = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetDomain(), "getter should return the property value")
	})

	t.Run("GetDomain_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetVerifiedDomainsResponse{}
		obj.Domain = nil

		// Act & Assert
		assert.Nil(t, obj.GetDomain(), "getter should return nil when property is nil")
	})

	t.Run("GetDomain_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetVerifiedDomainsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetDomain() // Should return zero value
	})

	t.Run("GetIsFreeEmailProvider", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetVerifiedDomainsResponse{}
		var expected *bool
		obj.IsFreeEmailProvider = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetIsFreeEmailProvider(), "getter should return the property value")
	})

	t.Run("GetIsFreeEmailProvider_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetVerifiedDomainsResponse{}
		obj.IsFreeEmailProvider = nil

		// Act & Assert
		assert.Nil(t, obj.GetIsFreeEmailProvider(), "getter should return nil when property is nil")
	})

	t.Run("GetIsFreeEmailProvider_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetVerifiedDomainsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetIsFreeEmailProvider() // Should return zero value
	})

	t.Run("GetStatus", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetVerifiedDomainsResponse{}
		var expected *GetVerifiedDomainsResponseStatus
		obj.Status = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStatus(), "getter should return the property value")
	})

	t.Run("GetStatus_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetVerifiedDomainsResponse{}
		obj.Status = nil

		// Act & Assert
		assert.Nil(t, obj.GetStatus(), "getter should return nil when property is nil")
	})

	t.Run("GetStatus_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetVerifiedDomainsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStatus() // Should return zero value
	})

	t.Run("GetVerificationEmail", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetVerifiedDomainsResponse{}
		var expected *string
		obj.VerificationEmail = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetVerificationEmail(), "getter should return the property value")
	})

	t.Run("GetVerificationEmail_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetVerifiedDomainsResponse{}
		obj.VerificationEmail = nil

		// Act & Assert
		assert.Nil(t, obj.GetVerificationEmail(), "getter should return nil when property is nil")
	})

	t.Run("GetVerificationEmail_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetVerifiedDomainsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetVerificationEmail() // Should return zero value
	})

	t.Run("GetVerificationSent", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetVerifiedDomainsResponse{}
		var expected *time.Time
		obj.VerificationSent = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetVerificationSent(), "getter should return the property value")
	})

	t.Run("GetVerificationSent_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetVerifiedDomainsResponse{}
		obj.VerificationSent = nil

		// Act & Assert
		assert.Nil(t, obj.GetVerificationSent(), "getter should return nil when property is nil")
	})

	t.Run("GetVerificationSent_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetVerifiedDomainsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetVerificationSent() // Should return zero value
	})

	t.Run("GetVerified", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetVerifiedDomainsResponse{}
		var expected *bool
		obj.Verified = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetVerified(), "getter should return the property value")
	})

	t.Run("GetVerified_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetVerifiedDomainsResponse{}
		obj.Verified = nil

		// Act & Assert
		assert.Nil(t, obj.GetVerified(), "getter should return nil when property is nil")
	})

	t.Run("GetVerified_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetVerifiedDomainsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetVerified() // Should return zero value
	})

}

func TestSettersMarkExplicitGetVerifiedDomainsResponse(t *testing.T) {
	t.Run("SetAuthenticated_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetVerifiedDomainsResponse{}
		var fernTestValueAuthenticated *bool

		// Act
		obj.SetAuthenticated(fernTestValueAuthenticated)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetDomain_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetVerifiedDomainsResponse{}
		var fernTestValueDomain *string

		// Act
		obj.SetDomain(fernTestValueDomain)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetIsFreeEmailProvider_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetVerifiedDomainsResponse{}
		var fernTestValueIsFreeEmailProvider *bool

		// Act
		obj.SetIsFreeEmailProvider(fernTestValueIsFreeEmailProvider)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetStatus_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetVerifiedDomainsResponse{}
		var fernTestValueStatus *GetVerifiedDomainsResponseStatus

		// Act
		obj.SetStatus(fernTestValueStatus)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetVerificationEmail_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetVerifiedDomainsResponse{}
		var fernTestValueVerificationEmail *string

		// Act
		obj.SetVerificationEmail(fernTestValueVerificationEmail)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetVerificationSent_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetVerifiedDomainsResponse{}
		var fernTestValueVerificationSent *time.Time

		// Act
		obj.SetVerificationSent(fernTestValueVerificationSent)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetVerified_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetVerifiedDomainsResponse{}
		var fernTestValueVerified *bool

		// Act
		obj.SetVerified(fernTestValueVerified)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersListVerifiedDomainsResponse(t *testing.T) {
	t.Run("SetDomains", func(t *testing.T) {
		obj := &ListVerifiedDomainsResponse{}
		var fernTestValueDomains []*ListVerifiedDomainsResponseDomainsItem
		obj.SetDomains(fernTestValueDomains)
		assert.Equal(t, fernTestValueDomains, obj.Domains)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTotalItems", func(t *testing.T) {
		obj := &ListVerifiedDomainsResponse{}
		var fernTestValueTotalItems *int
		obj.SetTotalItems(fernTestValueTotalItems)
		assert.Equal(t, fernTestValueTotalItems, obj.TotalItems)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersListVerifiedDomainsResponse(t *testing.T) {
	t.Run("GetDomains", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListVerifiedDomainsResponse{}
		var expected []*ListVerifiedDomainsResponseDomainsItem
		obj.Domains = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetDomains(), "getter should return the property value")
	})

	t.Run("GetDomains_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListVerifiedDomainsResponse{}
		obj.Domains = nil

		// Act & Assert
		assert.Nil(t, obj.GetDomains(), "getter should return nil when property is nil")
	})

	t.Run("GetDomains_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListVerifiedDomainsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetDomains() // Should return zero value
	})

	t.Run("GetTotalItems", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListVerifiedDomainsResponse{}
		var expected *int
		obj.TotalItems = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTotalItems(), "getter should return the property value")
	})

	t.Run("GetTotalItems_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListVerifiedDomainsResponse{}
		obj.TotalItems = nil

		// Act & Assert
		assert.Nil(t, obj.GetTotalItems(), "getter should return nil when property is nil")
	})

	t.Run("GetTotalItems_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListVerifiedDomainsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTotalItems() // Should return zero value
	})

}

func TestSettersMarkExplicitListVerifiedDomainsResponse(t *testing.T) {
	t.Run("SetDomains_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListVerifiedDomainsResponse{}
		var fernTestValueDomains []*ListVerifiedDomainsResponseDomainsItem

		// Act
		obj.SetDomains(fernTestValueDomains)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetTotalItems_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListVerifiedDomainsResponse{}
		var fernTestValueTotalItems *int

		// Act
		obj.SetTotalItems(fernTestValueTotalItems)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestSettersListVerifiedDomainsResponseDomainsItem(t *testing.T) {
	t.Run("SetAuthenticated", func(t *testing.T) {
		obj := &ListVerifiedDomainsResponseDomainsItem{}
		var fernTestValueAuthenticated *bool
		obj.SetAuthenticated(fernTestValueAuthenticated)
		assert.Equal(t, fernTestValueAuthenticated, obj.Authenticated)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetDomain", func(t *testing.T) {
		obj := &ListVerifiedDomainsResponseDomainsItem{}
		var fernTestValueDomain *string
		obj.SetDomain(fernTestValueDomain)
		assert.Equal(t, fernTestValueDomain, obj.Domain)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetIsFreeEmailProvider", func(t *testing.T) {
		obj := &ListVerifiedDomainsResponseDomainsItem{}
		var fernTestValueIsFreeEmailProvider *bool
		obj.SetIsFreeEmailProvider(fernTestValueIsFreeEmailProvider)
		assert.Equal(t, fernTestValueIsFreeEmailProvider, obj.IsFreeEmailProvider)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStatus", func(t *testing.T) {
		obj := &ListVerifiedDomainsResponseDomainsItem{}
		var fernTestValueStatus *ListVerifiedDomainsResponseDomainsItemStatus
		obj.SetStatus(fernTestValueStatus)
		assert.Equal(t, fernTestValueStatus, obj.Status)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetVerificationEmail", func(t *testing.T) {
		obj := &ListVerifiedDomainsResponseDomainsItem{}
		var fernTestValueVerificationEmail *string
		obj.SetVerificationEmail(fernTestValueVerificationEmail)
		assert.Equal(t, fernTestValueVerificationEmail, obj.VerificationEmail)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetVerificationSent", func(t *testing.T) {
		obj := &ListVerifiedDomainsResponseDomainsItem{}
		var fernTestValueVerificationSent *time.Time
		obj.SetVerificationSent(fernTestValueVerificationSent)
		assert.Equal(t, fernTestValueVerificationSent, obj.VerificationSent)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetVerified", func(t *testing.T) {
		obj := &ListVerifiedDomainsResponseDomainsItem{}
		var fernTestValueVerified *bool
		obj.SetVerified(fernTestValueVerified)
		assert.Equal(t, fernTestValueVerified, obj.Verified)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersListVerifiedDomainsResponseDomainsItem(t *testing.T) {
	t.Run("GetAuthenticated", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListVerifiedDomainsResponseDomainsItem{}
		var expected *bool
		obj.Authenticated = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetAuthenticated(), "getter should return the property value")
	})

	t.Run("GetAuthenticated_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListVerifiedDomainsResponseDomainsItem{}
		obj.Authenticated = nil

		// Act & Assert
		assert.Nil(t, obj.GetAuthenticated(), "getter should return nil when property is nil")
	})

	t.Run("GetAuthenticated_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListVerifiedDomainsResponseDomainsItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetAuthenticated() // Should return zero value
	})

	t.Run("GetDomain", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListVerifiedDomainsResponseDomainsItem{}
		var expected *string
		obj.Domain = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetDomain(), "getter should return the property value")
	})

	t.Run("GetDomain_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListVerifiedDomainsResponseDomainsItem{}
		obj.Domain = nil

		// Act & Assert
		assert.Nil(t, obj.GetDomain(), "getter should return nil when property is nil")
	})

	t.Run("GetDomain_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListVerifiedDomainsResponseDomainsItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetDomain() // Should return zero value
	})

	t.Run("GetIsFreeEmailProvider", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListVerifiedDomainsResponseDomainsItem{}
		var expected *bool
		obj.IsFreeEmailProvider = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetIsFreeEmailProvider(), "getter should return the property value")
	})

	t.Run("GetIsFreeEmailProvider_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListVerifiedDomainsResponseDomainsItem{}
		obj.IsFreeEmailProvider = nil

		// Act & Assert
		assert.Nil(t, obj.GetIsFreeEmailProvider(), "getter should return nil when property is nil")
	})

	t.Run("GetIsFreeEmailProvider_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListVerifiedDomainsResponseDomainsItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetIsFreeEmailProvider() // Should return zero value
	})

	t.Run("GetStatus", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListVerifiedDomainsResponseDomainsItem{}
		var expected *ListVerifiedDomainsResponseDomainsItemStatus
		obj.Status = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStatus(), "getter should return the property value")
	})

	t.Run("GetStatus_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListVerifiedDomainsResponseDomainsItem{}
		obj.Status = nil

		// Act & Assert
		assert.Nil(t, obj.GetStatus(), "getter should return nil when property is nil")
	})

	t.Run("GetStatus_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListVerifiedDomainsResponseDomainsItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStatus() // Should return zero value
	})

	t.Run("GetVerificationEmail", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListVerifiedDomainsResponseDomainsItem{}
		var expected *string
		obj.VerificationEmail = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetVerificationEmail(), "getter should return the property value")
	})

	t.Run("GetVerificationEmail_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListVerifiedDomainsResponseDomainsItem{}
		obj.VerificationEmail = nil

		// Act & Assert
		assert.Nil(t, obj.GetVerificationEmail(), "getter should return nil when property is nil")
	})

	t.Run("GetVerificationEmail_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListVerifiedDomainsResponseDomainsItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetVerificationEmail() // Should return zero value
	})

	t.Run("GetVerificationSent", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListVerifiedDomainsResponseDomainsItem{}
		var expected *time.Time
		obj.VerificationSent = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetVerificationSent(), "getter should return the property value")
	})

	t.Run("GetVerificationSent_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListVerifiedDomainsResponseDomainsItem{}
		obj.VerificationSent = nil

		// Act & Assert
		assert.Nil(t, obj.GetVerificationSent(), "getter should return nil when property is nil")
	})

	t.Run("GetVerificationSent_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListVerifiedDomainsResponseDomainsItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetVerificationSent() // Should return zero value
	})

	t.Run("GetVerified", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListVerifiedDomainsResponseDomainsItem{}
		var expected *bool
		obj.Verified = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetVerified(), "getter should return the property value")
	})

	t.Run("GetVerified_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListVerifiedDomainsResponseDomainsItem{}
		obj.Verified = nil

		// Act & Assert
		assert.Nil(t, obj.GetVerified(), "getter should return nil when property is nil")
	})

	t.Run("GetVerified_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListVerifiedDomainsResponseDomainsItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetVerified() // Should return zero value
	})

}

func TestSettersMarkExplicitListVerifiedDomainsResponseDomainsItem(t *testing.T) {
	t.Run("SetAuthenticated_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListVerifiedDomainsResponseDomainsItem{}
		var fernTestValueAuthenticated *bool

		// Act
		obj.SetAuthenticated(fernTestValueAuthenticated)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetDomain_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListVerifiedDomainsResponseDomainsItem{}
		var fernTestValueDomain *string

		// Act
		obj.SetDomain(fernTestValueDomain)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetIsFreeEmailProvider_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListVerifiedDomainsResponseDomainsItem{}
		var fernTestValueIsFreeEmailProvider *bool

		// Act
		obj.SetIsFreeEmailProvider(fernTestValueIsFreeEmailProvider)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetStatus_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListVerifiedDomainsResponseDomainsItem{}
		var fernTestValueStatus *ListVerifiedDomainsResponseDomainsItemStatus

		// Act
		obj.SetStatus(fernTestValueStatus)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetVerificationEmail_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListVerifiedDomainsResponseDomainsItem{}
		var fernTestValueVerificationEmail *string

		// Act
		obj.SetVerificationEmail(fernTestValueVerificationEmail)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetVerificationSent_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListVerifiedDomainsResponseDomainsItem{}
		var fernTestValueVerificationSent *time.Time

		// Act
		obj.SetVerificationSent(fernTestValueVerificationSent)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetVerified_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListVerifiedDomainsResponseDomainsItem{}
		var fernTestValueVerified *bool

		// Act
		obj.SetVerified(fernTestValueVerified)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

}

func TestJSONMarshalingCreateActionVerifyVerifiedDomainsResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateActionVerifyVerifiedDomainsResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled CreateActionVerifyVerifiedDomainsResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj CreateActionVerifyVerifiedDomainsResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj CreateActionVerifyVerifiedDomainsResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingCreateVerifiedDomainsResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateVerifiedDomainsResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled CreateVerifiedDomainsResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj CreateVerifiedDomainsResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj CreateVerifiedDomainsResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingGetVerifiedDomainsResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetVerifiedDomainsResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled GetVerifiedDomainsResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj GetVerifiedDomainsResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj GetVerifiedDomainsResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingListVerifiedDomainsResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListVerifiedDomainsResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ListVerifiedDomainsResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ListVerifiedDomainsResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ListVerifiedDomainsResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingListVerifiedDomainsResponseDomainsItem(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListVerifiedDomainsResponseDomainsItem{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ListVerifiedDomainsResponseDomainsItem
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ListVerifiedDomainsResponseDomainsItem
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ListVerifiedDomainsResponseDomainsItem
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestStringCreateActionVerifyVerifiedDomainsResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &CreateActionVerifyVerifiedDomainsResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateActionVerifyVerifiedDomainsResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringCreateVerifiedDomainsResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &CreateVerifiedDomainsResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateVerifiedDomainsResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringGetVerifiedDomainsResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &GetVerifiedDomainsResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetVerifiedDomainsResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringListVerifiedDomainsResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ListVerifiedDomainsResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListVerifiedDomainsResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringListVerifiedDomainsResponseDomainsItem(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ListVerifiedDomainsResponseDomainsItem{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListVerifiedDomainsResponseDomainsItem
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestEnumCreateActionVerifyVerifiedDomainsResponseStatus(t *testing.T) {
	t.Run("NewFromString_VERIFICATION_IN_PROGRESS", func(t *testing.T) {
		t.Parallel()
		val, err := NewCreateActionVerifyVerifiedDomainsResponseStatusFromString("VERIFICATION_IN_PROGRESS")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, CreateActionVerifyVerifiedDomainsResponseStatus("VERIFICATION_IN_PROGRESS"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_VERIFIED", func(t *testing.T) {
		t.Parallel()
		val, err := NewCreateActionVerifyVerifiedDomainsResponseStatusFromString("VERIFIED")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, CreateActionVerifyVerifiedDomainsResponseStatus("VERIFIED"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_EXPIRED", func(t *testing.T) {
		t.Parallel()
		val, err := NewCreateActionVerifyVerifiedDomainsResponseStatusFromString("EXPIRED")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, CreateActionVerifyVerifiedDomainsResponseStatus("EXPIRED"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_ERROR", func(t *testing.T) {
		t.Parallel()
		val, err := NewCreateActionVerifyVerifiedDomainsResponseStatusFromString("ERROR")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, CreateActionVerifyVerifiedDomainsResponseStatus("ERROR"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_AUTHENTICATION_IN_PROGRESS", func(t *testing.T) {
		t.Parallel()
		val, err := NewCreateActionVerifyVerifiedDomainsResponseStatusFromString("AUTHENTICATION_IN_PROGRESS")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, CreateActionVerifyVerifiedDomainsResponseStatus("AUTHENTICATION_IN_PROGRESS"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_AUTHENTICATION_ERROR", func(t *testing.T) {
		t.Parallel()
		val, err := NewCreateActionVerifyVerifiedDomainsResponseStatusFromString("AUTHENTICATION_ERROR")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, CreateActionVerifyVerifiedDomainsResponseStatus("AUTHENTICATION_ERROR"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_AUTHENTICATED", func(t *testing.T) {
		t.Parallel()
		val, err := NewCreateActionVerifyVerifiedDomainsResponseStatusFromString("AUTHENTICATED")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, CreateActionVerifyVerifiedDomainsResponseStatus("AUTHENTICATED"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewCreateActionVerifyVerifiedDomainsResponseStatusFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewCreateActionVerifyVerifiedDomainsResponseStatusFromString("VERIFICATION_IN_PROGRESS")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumCreateVerifiedDomainsResponseStatus(t *testing.T) {
	t.Run("NewFromString_VERIFICATION_IN_PROGRESS", func(t *testing.T) {
		t.Parallel()
		val, err := NewCreateVerifiedDomainsResponseStatusFromString("VERIFICATION_IN_PROGRESS")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, CreateVerifiedDomainsResponseStatus("VERIFICATION_IN_PROGRESS"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_VERIFIED", func(t *testing.T) {
		t.Parallel()
		val, err := NewCreateVerifiedDomainsResponseStatusFromString("VERIFIED")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, CreateVerifiedDomainsResponseStatus("VERIFIED"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_EXPIRED", func(t *testing.T) {
		t.Parallel()
		val, err := NewCreateVerifiedDomainsResponseStatusFromString("EXPIRED")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, CreateVerifiedDomainsResponseStatus("EXPIRED"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_ERROR", func(t *testing.T) {
		t.Parallel()
		val, err := NewCreateVerifiedDomainsResponseStatusFromString("ERROR")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, CreateVerifiedDomainsResponseStatus("ERROR"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_AUTHENTICATION_IN_PROGRESS", func(t *testing.T) {
		t.Parallel()
		val, err := NewCreateVerifiedDomainsResponseStatusFromString("AUTHENTICATION_IN_PROGRESS")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, CreateVerifiedDomainsResponseStatus("AUTHENTICATION_IN_PROGRESS"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_AUTHENTICATION_ERROR", func(t *testing.T) {
		t.Parallel()
		val, err := NewCreateVerifiedDomainsResponseStatusFromString("AUTHENTICATION_ERROR")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, CreateVerifiedDomainsResponseStatus("AUTHENTICATION_ERROR"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_AUTHENTICATED", func(t *testing.T) {
		t.Parallel()
		val, err := NewCreateVerifiedDomainsResponseStatusFromString("AUTHENTICATED")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, CreateVerifiedDomainsResponseStatus("AUTHENTICATED"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewCreateVerifiedDomainsResponseStatusFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewCreateVerifiedDomainsResponseStatusFromString("VERIFICATION_IN_PROGRESS")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumGetVerifiedDomainsResponseStatus(t *testing.T) {
	t.Run("NewFromString_VERIFICATION_IN_PROGRESS", func(t *testing.T) {
		t.Parallel()
		val, err := NewGetVerifiedDomainsResponseStatusFromString("VERIFICATION_IN_PROGRESS")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, GetVerifiedDomainsResponseStatus("VERIFICATION_IN_PROGRESS"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_VERIFIED", func(t *testing.T) {
		t.Parallel()
		val, err := NewGetVerifiedDomainsResponseStatusFromString("VERIFIED")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, GetVerifiedDomainsResponseStatus("VERIFIED"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_EXPIRED", func(t *testing.T) {
		t.Parallel()
		val, err := NewGetVerifiedDomainsResponseStatusFromString("EXPIRED")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, GetVerifiedDomainsResponseStatus("EXPIRED"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_ERROR", func(t *testing.T) {
		t.Parallel()
		val, err := NewGetVerifiedDomainsResponseStatusFromString("ERROR")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, GetVerifiedDomainsResponseStatus("ERROR"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_AUTHENTICATION_IN_PROGRESS", func(t *testing.T) {
		t.Parallel()
		val, err := NewGetVerifiedDomainsResponseStatusFromString("AUTHENTICATION_IN_PROGRESS")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, GetVerifiedDomainsResponseStatus("AUTHENTICATION_IN_PROGRESS"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_AUTHENTICATION_ERROR", func(t *testing.T) {
		t.Parallel()
		val, err := NewGetVerifiedDomainsResponseStatusFromString("AUTHENTICATION_ERROR")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, GetVerifiedDomainsResponseStatus("AUTHENTICATION_ERROR"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_AUTHENTICATED", func(t *testing.T) {
		t.Parallel()
		val, err := NewGetVerifiedDomainsResponseStatusFromString("AUTHENTICATED")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, GetVerifiedDomainsResponseStatus("AUTHENTICATED"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewGetVerifiedDomainsResponseStatusFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewGetVerifiedDomainsResponseStatusFromString("VERIFICATION_IN_PROGRESS")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumListVerifiedDomainsResponseDomainsItemStatus(t *testing.T) {
	t.Run("NewFromString_VERIFICATION_IN_PROGRESS", func(t *testing.T) {
		t.Parallel()
		val, err := NewListVerifiedDomainsResponseDomainsItemStatusFromString("VERIFICATION_IN_PROGRESS")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListVerifiedDomainsResponseDomainsItemStatus("VERIFICATION_IN_PROGRESS"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_VERIFIED", func(t *testing.T) {
		t.Parallel()
		val, err := NewListVerifiedDomainsResponseDomainsItemStatusFromString("VERIFIED")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListVerifiedDomainsResponseDomainsItemStatus("VERIFIED"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_EXPIRED", func(t *testing.T) {
		t.Parallel()
		val, err := NewListVerifiedDomainsResponseDomainsItemStatusFromString("EXPIRED")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListVerifiedDomainsResponseDomainsItemStatus("EXPIRED"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_ERROR", func(t *testing.T) {
		t.Parallel()
		val, err := NewListVerifiedDomainsResponseDomainsItemStatusFromString("ERROR")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListVerifiedDomainsResponseDomainsItemStatus("ERROR"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_AUTHENTICATION_IN_PROGRESS", func(t *testing.T) {
		t.Parallel()
		val, err := NewListVerifiedDomainsResponseDomainsItemStatusFromString("AUTHENTICATION_IN_PROGRESS")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListVerifiedDomainsResponseDomainsItemStatus("AUTHENTICATION_IN_PROGRESS"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_AUTHENTICATION_ERROR", func(t *testing.T) {
		t.Parallel()
		val, err := NewListVerifiedDomainsResponseDomainsItemStatusFromString("AUTHENTICATION_ERROR")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListVerifiedDomainsResponseDomainsItemStatus("AUTHENTICATION_ERROR"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_AUTHENTICATED", func(t *testing.T) {
		t.Parallel()
		val, err := NewListVerifiedDomainsResponseDomainsItemStatusFromString("AUTHENTICATED")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListVerifiedDomainsResponseDomainsItemStatus("AUTHENTICATED"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewListVerifiedDomainsResponseDomainsItemStatusFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewListVerifiedDomainsResponseDomainsItemStatusFromString("VERIFICATION_IN_PROGRESS")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestExtraPropertiesCreateActionVerifyVerifiedDomainsResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &CreateActionVerifyVerifiedDomainsResponse{}
		// Should not panic when calling GetExtraProperties()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("GetExtraProperties() panicked: %v", r)
			}
		}()
		extraProps := obj.GetExtraProperties()
		// Result can be nil or an empty/non-empty map
		_ = extraProps
	})

	t.Run("GetExtraProperties_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateActionVerifyVerifiedDomainsResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesCreateVerifiedDomainsResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &CreateVerifiedDomainsResponse{}
		// Should not panic when calling GetExtraProperties()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("GetExtraProperties() panicked: %v", r)
			}
		}()
		extraProps := obj.GetExtraProperties()
		// Result can be nil or an empty/non-empty map
		_ = extraProps
	})

	t.Run("GetExtraProperties_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateVerifiedDomainsResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesGetVerifiedDomainsResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &GetVerifiedDomainsResponse{}
		// Should not panic when calling GetExtraProperties()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("GetExtraProperties() panicked: %v", r)
			}
		}()
		extraProps := obj.GetExtraProperties()
		// Result can be nil or an empty/non-empty map
		_ = extraProps
	})

	t.Run("GetExtraProperties_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetVerifiedDomainsResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesListVerifiedDomainsResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ListVerifiedDomainsResponse{}
		// Should not panic when calling GetExtraProperties()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("GetExtraProperties() panicked: %v", r)
			}
		}()
		extraProps := obj.GetExtraProperties()
		// Result can be nil or an empty/non-empty map
		_ = extraProps
	})

	t.Run("GetExtraProperties_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListVerifiedDomainsResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesListVerifiedDomainsResponseDomainsItem(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ListVerifiedDomainsResponseDomainsItem{}
		// Should not panic when calling GetExtraProperties()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("GetExtraProperties() panicked: %v", r)
			}
		}()
		extraProps := obj.GetExtraProperties()
		// Result can be nil or an empty/non-empty map
		_ = extraProps
	})

	t.Run("GetExtraProperties_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListVerifiedDomainsResponseDomainsItem
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}
