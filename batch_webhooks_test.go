// Code generated from our API definition. DO NOT EDIT.

package api

import (
	json "encoding/json"
	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
	testing "testing"
)

func TestSettersCreateBatchWebhooksRequest(t *testing.T) {
	t.Run("SetEnabled", func(t *testing.T) {
		obj := &CreateBatchWebhooksRequest{}
		var fernTestValueEnabled *bool
		obj.SetEnabled(fernTestValueEnabled)
		assert.Equal(t, fernTestValueEnabled, obj.Enabled)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetURL", func(t *testing.T) {
		obj := &CreateBatchWebhooksRequest{}
		var fernTestValueURL string
		obj.SetURL(fernTestValueURL)
		assert.Equal(t, fernTestValueURL, obj.URL)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitCreateBatchWebhooksRequest(t *testing.T) {
	t.Run("SetEnabled_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateBatchWebhooksRequest{}
		var fernTestValueEnabled *bool

		// Act
		obj.SetEnabled(fernTestValueEnabled)

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

	t.Run("SetURL_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateBatchWebhooksRequest{}
		var fernTestValueURL string

		// Act
		obj.SetURL(fernTestValueURL)

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

func TestSettersDeleteBatchWebhooksRequest(t *testing.T) {
	t.Run("SetBatchWebhookID", func(t *testing.T) {
		obj := &DeleteBatchWebhooksRequest{}
		var fernTestValueBatchWebhookID string
		obj.SetBatchWebhookID(fernTestValueBatchWebhookID)
		assert.Equal(t, fernTestValueBatchWebhookID, obj.BatchWebhookID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitDeleteBatchWebhooksRequest(t *testing.T) {
	t.Run("SetBatchWebhookID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &DeleteBatchWebhooksRequest{}
		var fernTestValueBatchWebhookID string

		// Act
		obj.SetBatchWebhookID(fernTestValueBatchWebhookID)

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

func TestSettersGetBatchWebhooksRequest(t *testing.T) {
	t.Run("SetBatchWebhookID", func(t *testing.T) {
		obj := &GetBatchWebhooksRequest{}
		var fernTestValueBatchWebhookID string
		obj.SetBatchWebhookID(fernTestValueBatchWebhookID)
		assert.Equal(t, fernTestValueBatchWebhookID, obj.BatchWebhookID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetFields", func(t *testing.T) {
		obj := &GetBatchWebhooksRequest{}
		var fernTestValueFields []*string
		obj.SetFields(fernTestValueFields)
		assert.Equal(t, fernTestValueFields, obj.Fields)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetExcludeFields", func(t *testing.T) {
		obj := &GetBatchWebhooksRequest{}
		var fernTestValueExcludeFields []*string
		obj.SetExcludeFields(fernTestValueExcludeFields)
		assert.Equal(t, fernTestValueExcludeFields, obj.ExcludeFields)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitGetBatchWebhooksRequest(t *testing.T) {
	t.Run("SetBatchWebhookID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetBatchWebhooksRequest{}
		var fernTestValueBatchWebhookID string

		// Act
		obj.SetBatchWebhookID(fernTestValueBatchWebhookID)

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

	t.Run("SetFields_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetBatchWebhooksRequest{}
		var fernTestValueFields []*string

		// Act
		obj.SetFields(fernTestValueFields)

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

	t.Run("SetExcludeFields_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetBatchWebhooksRequest{}
		var fernTestValueExcludeFields []*string

		// Act
		obj.SetExcludeFields(fernTestValueExcludeFields)

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

func TestSettersListBatchWebhooksRequest(t *testing.T) {
	t.Run("SetFields", func(t *testing.T) {
		obj := &ListBatchWebhooksRequest{}
		var fernTestValueFields []*string
		obj.SetFields(fernTestValueFields)
		assert.Equal(t, fernTestValueFields, obj.Fields)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetExcludeFields", func(t *testing.T) {
		obj := &ListBatchWebhooksRequest{}
		var fernTestValueExcludeFields []*string
		obj.SetExcludeFields(fernTestValueExcludeFields)
		assert.Equal(t, fernTestValueExcludeFields, obj.ExcludeFields)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCount", func(t *testing.T) {
		obj := &ListBatchWebhooksRequest{}
		var fernTestValueCount *int
		obj.SetCount(fernTestValueCount)
		assert.Equal(t, fernTestValueCount, obj.Count)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetOffset", func(t *testing.T) {
		obj := &ListBatchWebhooksRequest{}
		var fernTestValueOffset *int
		obj.SetOffset(fernTestValueOffset)
		assert.Equal(t, fernTestValueOffset, obj.Offset)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitListBatchWebhooksRequest(t *testing.T) {
	t.Run("SetFields_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListBatchWebhooksRequest{}
		var fernTestValueFields []*string

		// Act
		obj.SetFields(fernTestValueFields)

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

	t.Run("SetExcludeFields_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListBatchWebhooksRequest{}
		var fernTestValueExcludeFields []*string

		// Act
		obj.SetExcludeFields(fernTestValueExcludeFields)

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

	t.Run("SetCount_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListBatchWebhooksRequest{}
		var fernTestValueCount *int

		// Act
		obj.SetCount(fernTestValueCount)

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

	t.Run("SetOffset_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListBatchWebhooksRequest{}
		var fernTestValueOffset *int

		// Act
		obj.SetOffset(fernTestValueOffset)

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

func TestSettersBatchWebhook(t *testing.T) {
	t.Run("SetLinks", func(t *testing.T) {
		obj := &BatchWebhook{}
		var fernTestValueLinks [][]*BatchWebhookLinksItemItem
		obj.SetLinks(fernTestValueLinks)
		assert.Equal(t, fernTestValueLinks, obj.Links)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetEnabled", func(t *testing.T) {
		obj := &BatchWebhook{}
		var fernTestValueEnabled *bool
		obj.SetEnabled(fernTestValueEnabled)
		assert.Equal(t, fernTestValueEnabled, obj.Enabled)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetID", func(t *testing.T) {
		obj := &BatchWebhook{}
		var fernTestValueID *string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSigningEnabled", func(t *testing.T) {
		obj := &BatchWebhook{}
		var fernTestValueSigningEnabled *bool
		obj.SetSigningEnabled(fernTestValueSigningEnabled)
		assert.Equal(t, fernTestValueSigningEnabled, obj.SigningEnabled)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetURL", func(t *testing.T) {
		obj := &BatchWebhook{}
		var fernTestValueURL *string
		obj.SetURL(fernTestValueURL)
		assert.Equal(t, fernTestValueURL, obj.URL)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersBatchWebhook(t *testing.T) {
	t.Run("GetLinks", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhook{}
		var expected [][]*BatchWebhookLinksItemItem
		obj.Links = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetLinks(), "getter should return the property value")
	})

	t.Run("GetLinks_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhook{}
		obj.Links = nil

		// Act & Assert
		assert.Nil(t, obj.GetLinks(), "getter should return nil when property is nil")
	})

	t.Run("GetLinks_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *BatchWebhook
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetLinks() // Should return zero value
	})

	t.Run("GetEnabled", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhook{}
		var expected *bool
		obj.Enabled = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEnabled(), "getter should return the property value")
	})

	t.Run("GetEnabled_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhook{}
		obj.Enabled = nil

		// Act & Assert
		assert.Nil(t, obj.GetEnabled(), "getter should return nil when property is nil")
	})

	t.Run("GetEnabled_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *BatchWebhook
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEnabled() // Should return zero value
	})

	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhook{}
		var expected *string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhook{}
		obj.ID = nil

		// Act & Assert
		assert.Nil(t, obj.GetID(), "getter should return nil when property is nil")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *BatchWebhook
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetID() // Should return zero value
	})

	t.Run("GetSigningEnabled", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhook{}
		var expected *bool
		obj.SigningEnabled = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSigningEnabled(), "getter should return the property value")
	})

	t.Run("GetSigningEnabled_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhook{}
		obj.SigningEnabled = nil

		// Act & Assert
		assert.Nil(t, obj.GetSigningEnabled(), "getter should return nil when property is nil")
	})

	t.Run("GetSigningEnabled_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *BatchWebhook
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSigningEnabled() // Should return zero value
	})

	t.Run("GetURL", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhook{}
		var expected *string
		obj.URL = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetURL(), "getter should return the property value")
	})

	t.Run("GetURL_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhook{}
		obj.URL = nil

		// Act & Assert
		assert.Nil(t, obj.GetURL(), "getter should return nil when property is nil")
	})

	t.Run("GetURL_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *BatchWebhook
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetURL() // Should return zero value
	})

}

func TestSettersMarkExplicitBatchWebhook(t *testing.T) {
	t.Run("SetLinks_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhook{}
		var fernTestValueLinks [][]*BatchWebhookLinksItemItem

		// Act
		obj.SetLinks(fernTestValueLinks)

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

	t.Run("SetEnabled_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhook{}
		var fernTestValueEnabled *bool

		// Act
		obj.SetEnabled(fernTestValueEnabled)

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

	t.Run("SetID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhook{}
		var fernTestValueID *string

		// Act
		obj.SetID(fernTestValueID)

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

	t.Run("SetSigningEnabled_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhook{}
		var fernTestValueSigningEnabled *bool

		// Act
		obj.SetSigningEnabled(fernTestValueSigningEnabled)

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

	t.Run("SetURL_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhook{}
		var fernTestValueURL *string

		// Act
		obj.SetURL(fernTestValueURL)

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

func TestSettersBatchWebhookLinksItemItem(t *testing.T) {
	t.Run("SetHref", func(t *testing.T) {
		obj := &BatchWebhookLinksItemItem{}
		var fernTestValueHref *string
		obj.SetHref(fernTestValueHref)
		assert.Equal(t, fernTestValueHref, obj.Href)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMethod", func(t *testing.T) {
		obj := &BatchWebhookLinksItemItem{}
		var fernTestValueMethod *BatchWebhookLinksItemItemMethod
		obj.SetMethod(fernTestValueMethod)
		assert.Equal(t, fernTestValueMethod, obj.Method)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRel", func(t *testing.T) {
		obj := &BatchWebhookLinksItemItem{}
		var fernTestValueRel *string
		obj.SetRel(fernTestValueRel)
		assert.Equal(t, fernTestValueRel, obj.Rel)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSchema", func(t *testing.T) {
		obj := &BatchWebhookLinksItemItem{}
		var fernTestValueSchema *string
		obj.SetSchema(fernTestValueSchema)
		assert.Equal(t, fernTestValueSchema, obj.Schema)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTargetSchema", func(t *testing.T) {
		obj := &BatchWebhookLinksItemItem{}
		var fernTestValueTargetSchema *string
		obj.SetTargetSchema(fernTestValueTargetSchema)
		assert.Equal(t, fernTestValueTargetSchema, obj.TargetSchema)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersBatchWebhookLinksItemItem(t *testing.T) {
	t.Run("GetHref", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhookLinksItemItem{}
		var expected *string
		obj.Href = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetHref(), "getter should return the property value")
	})

	t.Run("GetHref_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhookLinksItemItem{}
		obj.Href = nil

		// Act & Assert
		assert.Nil(t, obj.GetHref(), "getter should return nil when property is nil")
	})

	t.Run("GetHref_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *BatchWebhookLinksItemItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetHref() // Should return zero value
	})

	t.Run("GetMethod", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhookLinksItemItem{}
		var expected *BatchWebhookLinksItemItemMethod
		obj.Method = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMethod(), "getter should return the property value")
	})

	t.Run("GetMethod_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhookLinksItemItem{}
		obj.Method = nil

		// Act & Assert
		assert.Nil(t, obj.GetMethod(), "getter should return nil when property is nil")
	})

	t.Run("GetMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *BatchWebhookLinksItemItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetMethod() // Should return zero value
	})

	t.Run("GetRel", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhookLinksItemItem{}
		var expected *string
		obj.Rel = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetRel(), "getter should return the property value")
	})

	t.Run("GetRel_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhookLinksItemItem{}
		obj.Rel = nil

		// Act & Assert
		assert.Nil(t, obj.GetRel(), "getter should return nil when property is nil")
	})

	t.Run("GetRel_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *BatchWebhookLinksItemItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetRel() // Should return zero value
	})

	t.Run("GetSchema", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhookLinksItemItem{}
		var expected *string
		obj.Schema = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSchema(), "getter should return the property value")
	})

	t.Run("GetSchema_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhookLinksItemItem{}
		obj.Schema = nil

		// Act & Assert
		assert.Nil(t, obj.GetSchema(), "getter should return nil when property is nil")
	})

	t.Run("GetSchema_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *BatchWebhookLinksItemItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSchema() // Should return zero value
	})

	t.Run("GetTargetSchema", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhookLinksItemItem{}
		var expected *string
		obj.TargetSchema = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTargetSchema(), "getter should return the property value")
	})

	t.Run("GetTargetSchema_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhookLinksItemItem{}
		obj.TargetSchema = nil

		// Act & Assert
		assert.Nil(t, obj.GetTargetSchema(), "getter should return nil when property is nil")
	})

	t.Run("GetTargetSchema_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *BatchWebhookLinksItemItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTargetSchema() // Should return zero value
	})

}

func TestSettersMarkExplicitBatchWebhookLinksItemItem(t *testing.T) {
	t.Run("SetHref_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhookLinksItemItem{}
		var fernTestValueHref *string

		// Act
		obj.SetHref(fernTestValueHref)

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

	t.Run("SetMethod_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhookLinksItemItem{}
		var fernTestValueMethod *BatchWebhookLinksItemItemMethod

		// Act
		obj.SetMethod(fernTestValueMethod)

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

	t.Run("SetRel_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhookLinksItemItem{}
		var fernTestValueRel *string

		// Act
		obj.SetRel(fernTestValueRel)

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

	t.Run("SetSchema_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhookLinksItemItem{}
		var fernTestValueSchema *string

		// Act
		obj.SetSchema(fernTestValueSchema)

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

	t.Run("SetTargetSchema_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhookLinksItemItem{}
		var fernTestValueTargetSchema *string

		// Act
		obj.SetTargetSchema(fernTestValueTargetSchema)

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

func TestSettersCreateBatchWebhooksResponse(t *testing.T) {
	t.Run("SetLinks", func(t *testing.T) {
		obj := &CreateBatchWebhooksResponse{}
		var fernTestValueLinks [][]*BatchWebhookLinksItemItem
		obj.SetLinks(fernTestValueLinks)
		assert.Equal(t, fernTestValueLinks, obj.Links)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetEnabled", func(t *testing.T) {
		obj := &CreateBatchWebhooksResponse{}
		var fernTestValueEnabled *bool
		obj.SetEnabled(fernTestValueEnabled)
		assert.Equal(t, fernTestValueEnabled, obj.Enabled)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetID", func(t *testing.T) {
		obj := &CreateBatchWebhooksResponse{}
		var fernTestValueID *string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSigningEnabled", func(t *testing.T) {
		obj := &CreateBatchWebhooksResponse{}
		var fernTestValueSigningEnabled *bool
		obj.SetSigningEnabled(fernTestValueSigningEnabled)
		assert.Equal(t, fernTestValueSigningEnabled, obj.SigningEnabled)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetURL", func(t *testing.T) {
		obj := &CreateBatchWebhooksResponse{}
		var fernTestValueURL *string
		obj.SetURL(fernTestValueURL)
		assert.Equal(t, fernTestValueURL, obj.URL)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSigningSecret", func(t *testing.T) {
		obj := &CreateBatchWebhooksResponse{}
		var fernTestValueSigningSecret *string
		obj.SetSigningSecret(fernTestValueSigningSecret)
		assert.Equal(t, fernTestValueSigningSecret, obj.SigningSecret)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersCreateBatchWebhooksResponse(t *testing.T) {
	t.Run("GetLinks", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateBatchWebhooksResponse{}
		var expected [][]*BatchWebhookLinksItemItem
		obj.Links = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetLinks(), "getter should return the property value")
	})

	t.Run("GetLinks_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateBatchWebhooksResponse{}
		obj.Links = nil

		// Act & Assert
		assert.Nil(t, obj.GetLinks(), "getter should return nil when property is nil")
	})

	t.Run("GetLinks_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateBatchWebhooksResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetLinks() // Should return zero value
	})

	t.Run("GetEnabled", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateBatchWebhooksResponse{}
		var expected *bool
		obj.Enabled = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEnabled(), "getter should return the property value")
	})

	t.Run("GetEnabled_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateBatchWebhooksResponse{}
		obj.Enabled = nil

		// Act & Assert
		assert.Nil(t, obj.GetEnabled(), "getter should return nil when property is nil")
	})

	t.Run("GetEnabled_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateBatchWebhooksResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEnabled() // Should return zero value
	})

	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateBatchWebhooksResponse{}
		var expected *string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateBatchWebhooksResponse{}
		obj.ID = nil

		// Act & Assert
		assert.Nil(t, obj.GetID(), "getter should return nil when property is nil")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateBatchWebhooksResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetID() // Should return zero value
	})

	t.Run("GetSigningEnabled", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateBatchWebhooksResponse{}
		var expected *bool
		obj.SigningEnabled = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSigningEnabled(), "getter should return the property value")
	})

	t.Run("GetSigningEnabled_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateBatchWebhooksResponse{}
		obj.SigningEnabled = nil

		// Act & Assert
		assert.Nil(t, obj.GetSigningEnabled(), "getter should return nil when property is nil")
	})

	t.Run("GetSigningEnabled_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateBatchWebhooksResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSigningEnabled() // Should return zero value
	})

	t.Run("GetURL", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateBatchWebhooksResponse{}
		var expected *string
		obj.URL = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetURL(), "getter should return the property value")
	})

	t.Run("GetURL_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateBatchWebhooksResponse{}
		obj.URL = nil

		// Act & Assert
		assert.Nil(t, obj.GetURL(), "getter should return nil when property is nil")
	})

	t.Run("GetURL_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateBatchWebhooksResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetURL() // Should return zero value
	})

	t.Run("GetSigningSecret", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateBatchWebhooksResponse{}
		var expected *string
		obj.SigningSecret = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSigningSecret(), "getter should return the property value")
	})

	t.Run("GetSigningSecret_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateBatchWebhooksResponse{}
		obj.SigningSecret = nil

		// Act & Assert
		assert.Nil(t, obj.GetSigningSecret(), "getter should return nil when property is nil")
	})

	t.Run("GetSigningSecret_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateBatchWebhooksResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSigningSecret() // Should return zero value
	})

}

func TestSettersMarkExplicitCreateBatchWebhooksResponse(t *testing.T) {
	t.Run("SetLinks_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateBatchWebhooksResponse{}
		var fernTestValueLinks [][]*BatchWebhookLinksItemItem

		// Act
		obj.SetLinks(fernTestValueLinks)

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

	t.Run("SetEnabled_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateBatchWebhooksResponse{}
		var fernTestValueEnabled *bool

		// Act
		obj.SetEnabled(fernTestValueEnabled)

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

	t.Run("SetID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateBatchWebhooksResponse{}
		var fernTestValueID *string

		// Act
		obj.SetID(fernTestValueID)

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

	t.Run("SetSigningEnabled_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateBatchWebhooksResponse{}
		var fernTestValueSigningEnabled *bool

		// Act
		obj.SetSigningEnabled(fernTestValueSigningEnabled)

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

	t.Run("SetURL_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateBatchWebhooksResponse{}
		var fernTestValueURL *string

		// Act
		obj.SetURL(fernTestValueURL)

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

	t.Run("SetSigningSecret_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateBatchWebhooksResponse{}
		var fernTestValueSigningSecret *string

		// Act
		obj.SetSigningSecret(fernTestValueSigningSecret)

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

func TestSettersListBatchWebhooksResponse(t *testing.T) {
	t.Run("SetLinks", func(t *testing.T) {
		obj := &ListBatchWebhooksResponse{}
		var fernTestValueLinks []*ListBatchWebhooksResponseLinksItem
		obj.SetLinks(fernTestValueLinks)
		assert.Equal(t, fernTestValueLinks, obj.Links)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTotalItems", func(t *testing.T) {
		obj := &ListBatchWebhooksResponse{}
		var fernTestValueTotalItems *int
		obj.SetTotalItems(fernTestValueTotalItems)
		assert.Equal(t, fernTestValueTotalItems, obj.TotalItems)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetWebhooks", func(t *testing.T) {
		obj := &ListBatchWebhooksResponse{}
		var fernTestValueWebhooks []*BatchWebhook
		obj.SetWebhooks(fernTestValueWebhooks)
		assert.Equal(t, fernTestValueWebhooks, obj.Webhooks)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersListBatchWebhooksResponse(t *testing.T) {
	t.Run("GetLinks", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListBatchWebhooksResponse{}
		var expected []*ListBatchWebhooksResponseLinksItem
		obj.Links = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetLinks(), "getter should return the property value")
	})

	t.Run("GetLinks_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListBatchWebhooksResponse{}
		obj.Links = nil

		// Act & Assert
		assert.Nil(t, obj.GetLinks(), "getter should return nil when property is nil")
	})

	t.Run("GetLinks_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListBatchWebhooksResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetLinks() // Should return zero value
	})

	t.Run("GetTotalItems", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListBatchWebhooksResponse{}
		var expected *int
		obj.TotalItems = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTotalItems(), "getter should return the property value")
	})

	t.Run("GetTotalItems_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListBatchWebhooksResponse{}
		obj.TotalItems = nil

		// Act & Assert
		assert.Nil(t, obj.GetTotalItems(), "getter should return nil when property is nil")
	})

	t.Run("GetTotalItems_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListBatchWebhooksResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTotalItems() // Should return zero value
	})

	t.Run("GetWebhooks", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListBatchWebhooksResponse{}
		var expected []*BatchWebhook
		obj.Webhooks = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetWebhooks(), "getter should return the property value")
	})

	t.Run("GetWebhooks_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListBatchWebhooksResponse{}
		obj.Webhooks = nil

		// Act & Assert
		assert.Nil(t, obj.GetWebhooks(), "getter should return nil when property is nil")
	})

	t.Run("GetWebhooks_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListBatchWebhooksResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetWebhooks() // Should return zero value
	})

}

func TestSettersMarkExplicitListBatchWebhooksResponse(t *testing.T) {
	t.Run("SetLinks_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListBatchWebhooksResponse{}
		var fernTestValueLinks []*ListBatchWebhooksResponseLinksItem

		// Act
		obj.SetLinks(fernTestValueLinks)

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
		obj := &ListBatchWebhooksResponse{}
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

	t.Run("SetWebhooks_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListBatchWebhooksResponse{}
		var fernTestValueWebhooks []*BatchWebhook

		// Act
		obj.SetWebhooks(fernTestValueWebhooks)

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

func TestSettersListBatchWebhooksResponseLinksItem(t *testing.T) {
	t.Run("SetHref", func(t *testing.T) {
		obj := &ListBatchWebhooksResponseLinksItem{}
		var fernTestValueHref *string
		obj.SetHref(fernTestValueHref)
		assert.Equal(t, fernTestValueHref, obj.Href)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMethod", func(t *testing.T) {
		obj := &ListBatchWebhooksResponseLinksItem{}
		var fernTestValueMethod *ListBatchWebhooksResponseLinksItemMethod
		obj.SetMethod(fernTestValueMethod)
		assert.Equal(t, fernTestValueMethod, obj.Method)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRel", func(t *testing.T) {
		obj := &ListBatchWebhooksResponseLinksItem{}
		var fernTestValueRel *string
		obj.SetRel(fernTestValueRel)
		assert.Equal(t, fernTestValueRel, obj.Rel)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSchema", func(t *testing.T) {
		obj := &ListBatchWebhooksResponseLinksItem{}
		var fernTestValueSchema *string
		obj.SetSchema(fernTestValueSchema)
		assert.Equal(t, fernTestValueSchema, obj.Schema)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTargetSchema", func(t *testing.T) {
		obj := &ListBatchWebhooksResponseLinksItem{}
		var fernTestValueTargetSchema *string
		obj.SetTargetSchema(fernTestValueTargetSchema)
		assert.Equal(t, fernTestValueTargetSchema, obj.TargetSchema)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersListBatchWebhooksResponseLinksItem(t *testing.T) {
	t.Run("GetHref", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListBatchWebhooksResponseLinksItem{}
		var expected *string
		obj.Href = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetHref(), "getter should return the property value")
	})

	t.Run("GetHref_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListBatchWebhooksResponseLinksItem{}
		obj.Href = nil

		// Act & Assert
		assert.Nil(t, obj.GetHref(), "getter should return nil when property is nil")
	})

	t.Run("GetHref_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListBatchWebhooksResponseLinksItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetHref() // Should return zero value
	})

	t.Run("GetMethod", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListBatchWebhooksResponseLinksItem{}
		var expected *ListBatchWebhooksResponseLinksItemMethod
		obj.Method = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMethod(), "getter should return the property value")
	})

	t.Run("GetMethod_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListBatchWebhooksResponseLinksItem{}
		obj.Method = nil

		// Act & Assert
		assert.Nil(t, obj.GetMethod(), "getter should return nil when property is nil")
	})

	t.Run("GetMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListBatchWebhooksResponseLinksItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetMethod() // Should return zero value
	})

	t.Run("GetRel", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListBatchWebhooksResponseLinksItem{}
		var expected *string
		obj.Rel = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetRel(), "getter should return the property value")
	})

	t.Run("GetRel_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListBatchWebhooksResponseLinksItem{}
		obj.Rel = nil

		// Act & Assert
		assert.Nil(t, obj.GetRel(), "getter should return nil when property is nil")
	})

	t.Run("GetRel_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListBatchWebhooksResponseLinksItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetRel() // Should return zero value
	})

	t.Run("GetSchema", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListBatchWebhooksResponseLinksItem{}
		var expected *string
		obj.Schema = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSchema(), "getter should return the property value")
	})

	t.Run("GetSchema_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListBatchWebhooksResponseLinksItem{}
		obj.Schema = nil

		// Act & Assert
		assert.Nil(t, obj.GetSchema(), "getter should return nil when property is nil")
	})

	t.Run("GetSchema_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListBatchWebhooksResponseLinksItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSchema() // Should return zero value
	})

	t.Run("GetTargetSchema", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListBatchWebhooksResponseLinksItem{}
		var expected *string
		obj.TargetSchema = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTargetSchema(), "getter should return the property value")
	})

	t.Run("GetTargetSchema_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListBatchWebhooksResponseLinksItem{}
		obj.TargetSchema = nil

		// Act & Assert
		assert.Nil(t, obj.GetTargetSchema(), "getter should return nil when property is nil")
	})

	t.Run("GetTargetSchema_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListBatchWebhooksResponseLinksItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTargetSchema() // Should return zero value
	})

}

func TestSettersMarkExplicitListBatchWebhooksResponseLinksItem(t *testing.T) {
	t.Run("SetHref_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListBatchWebhooksResponseLinksItem{}
		var fernTestValueHref *string

		// Act
		obj.SetHref(fernTestValueHref)

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

	t.Run("SetMethod_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListBatchWebhooksResponseLinksItem{}
		var fernTestValueMethod *ListBatchWebhooksResponseLinksItemMethod

		// Act
		obj.SetMethod(fernTestValueMethod)

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

	t.Run("SetRel_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListBatchWebhooksResponseLinksItem{}
		var fernTestValueRel *string

		// Act
		obj.SetRel(fernTestValueRel)

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

	t.Run("SetSchema_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListBatchWebhooksResponseLinksItem{}
		var fernTestValueSchema *string

		// Act
		obj.SetSchema(fernTestValueSchema)

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

	t.Run("SetTargetSchema_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListBatchWebhooksResponseLinksItem{}
		var fernTestValueTargetSchema *string

		// Act
		obj.SetTargetSchema(fernTestValueTargetSchema)

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

func TestSettersUpdateBatchWebhooksRequest(t *testing.T) {
	t.Run("SetBatchWebhookID", func(t *testing.T) {
		obj := &UpdateBatchWebhooksRequest{}
		var fernTestValueBatchWebhookID string
		obj.SetBatchWebhookID(fernTestValueBatchWebhookID)
		assert.Equal(t, fernTestValueBatchWebhookID, obj.BatchWebhookID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetEnabled", func(t *testing.T) {
		obj := &UpdateBatchWebhooksRequest{}
		var fernTestValueEnabled *bool
		obj.SetEnabled(fernTestValueEnabled)
		assert.Equal(t, fernTestValueEnabled, obj.Enabled)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetURL", func(t *testing.T) {
		obj := &UpdateBatchWebhooksRequest{}
		var fernTestValueURL *string
		obj.SetURL(fernTestValueURL)
		assert.Equal(t, fernTestValueURL, obj.URL)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitUpdateBatchWebhooksRequest(t *testing.T) {
	t.Run("SetBatchWebhookID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &UpdateBatchWebhooksRequest{}
		var fernTestValueBatchWebhookID string

		// Act
		obj.SetBatchWebhookID(fernTestValueBatchWebhookID)

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

	t.Run("SetEnabled_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &UpdateBatchWebhooksRequest{}
		var fernTestValueEnabled *bool

		// Act
		obj.SetEnabled(fernTestValueEnabled)

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

	t.Run("SetURL_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &UpdateBatchWebhooksRequest{}
		var fernTestValueURL *string

		// Act
		obj.SetURL(fernTestValueURL)

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

func TestJSONMarshalingBatchWebhook(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhook{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled BatchWebhook
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj BatchWebhook
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj BatchWebhook
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingBatchWebhookLinksItemItem(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &BatchWebhookLinksItemItem{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled BatchWebhookLinksItemItem
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj BatchWebhookLinksItemItem
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj BatchWebhookLinksItemItem
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingCreateBatchWebhooksResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateBatchWebhooksResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled CreateBatchWebhooksResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj CreateBatchWebhooksResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj CreateBatchWebhooksResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingListBatchWebhooksResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListBatchWebhooksResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ListBatchWebhooksResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ListBatchWebhooksResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ListBatchWebhooksResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingListBatchWebhooksResponseLinksItem(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListBatchWebhooksResponseLinksItem{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ListBatchWebhooksResponseLinksItem
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ListBatchWebhooksResponseLinksItem
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ListBatchWebhooksResponseLinksItem
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestStringBatchWebhook(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &BatchWebhook{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *BatchWebhook
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringBatchWebhookLinksItemItem(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &BatchWebhookLinksItemItem{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *BatchWebhookLinksItemItem
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringCreateBatchWebhooksResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &CreateBatchWebhooksResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateBatchWebhooksResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringListBatchWebhooksResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ListBatchWebhooksResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListBatchWebhooksResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringListBatchWebhooksResponseLinksItem(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ListBatchWebhooksResponseLinksItem{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListBatchWebhooksResponseLinksItem
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestEnumBatchWebhookLinksItemItemMethod(t *testing.T) {
	t.Run("NewFromString_GET", func(t *testing.T) {
		t.Parallel()
		val, err := NewBatchWebhookLinksItemItemMethodFromString("GET")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, BatchWebhookLinksItemItemMethod("GET"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_POST", func(t *testing.T) {
		t.Parallel()
		val, err := NewBatchWebhookLinksItemItemMethodFromString("POST")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, BatchWebhookLinksItemItemMethod("POST"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_PUT", func(t *testing.T) {
		t.Parallel()
		val, err := NewBatchWebhookLinksItemItemMethodFromString("PUT")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, BatchWebhookLinksItemItemMethod("PUT"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_PATCH", func(t *testing.T) {
		t.Parallel()
		val, err := NewBatchWebhookLinksItemItemMethodFromString("PATCH")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, BatchWebhookLinksItemItemMethod("PATCH"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_DELETE", func(t *testing.T) {
		t.Parallel()
		val, err := NewBatchWebhookLinksItemItemMethodFromString("DELETE")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, BatchWebhookLinksItemItemMethod("DELETE"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_OPTIONS", func(t *testing.T) {
		t.Parallel()
		val, err := NewBatchWebhookLinksItemItemMethodFromString("OPTIONS")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, BatchWebhookLinksItemItemMethod("OPTIONS"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_HEAD", func(t *testing.T) {
		t.Parallel()
		val, err := NewBatchWebhookLinksItemItemMethodFromString("HEAD")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, BatchWebhookLinksItemItemMethod("HEAD"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewBatchWebhookLinksItemItemMethodFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewBatchWebhookLinksItemItemMethodFromString("GET")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumListBatchWebhooksResponseLinksItemMethod(t *testing.T) {
	t.Run("NewFromString_GET", func(t *testing.T) {
		t.Parallel()
		val, err := NewListBatchWebhooksResponseLinksItemMethodFromString("GET")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListBatchWebhooksResponseLinksItemMethod("GET"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_POST", func(t *testing.T) {
		t.Parallel()
		val, err := NewListBatchWebhooksResponseLinksItemMethodFromString("POST")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListBatchWebhooksResponseLinksItemMethod("POST"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_PUT", func(t *testing.T) {
		t.Parallel()
		val, err := NewListBatchWebhooksResponseLinksItemMethodFromString("PUT")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListBatchWebhooksResponseLinksItemMethod("PUT"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_PATCH", func(t *testing.T) {
		t.Parallel()
		val, err := NewListBatchWebhooksResponseLinksItemMethodFromString("PATCH")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListBatchWebhooksResponseLinksItemMethod("PATCH"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_DELETE", func(t *testing.T) {
		t.Parallel()
		val, err := NewListBatchWebhooksResponseLinksItemMethodFromString("DELETE")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListBatchWebhooksResponseLinksItemMethod("DELETE"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_OPTIONS", func(t *testing.T) {
		t.Parallel()
		val, err := NewListBatchWebhooksResponseLinksItemMethodFromString("OPTIONS")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListBatchWebhooksResponseLinksItemMethod("OPTIONS"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_HEAD", func(t *testing.T) {
		t.Parallel()
		val, err := NewListBatchWebhooksResponseLinksItemMethodFromString("HEAD")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListBatchWebhooksResponseLinksItemMethod("HEAD"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewListBatchWebhooksResponseLinksItemMethodFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewListBatchWebhooksResponseLinksItemMethodFromString("GET")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestExtraPropertiesBatchWebhook(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &BatchWebhook{}
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
		var obj *BatchWebhook
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesBatchWebhookLinksItemItem(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &BatchWebhookLinksItemItem{}
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
		var obj *BatchWebhookLinksItemItem
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesCreateBatchWebhooksResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &CreateBatchWebhooksResponse{}
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
		var obj *CreateBatchWebhooksResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesListBatchWebhooksResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ListBatchWebhooksResponse{}
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
		var obj *ListBatchWebhooksResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesListBatchWebhooksResponseLinksItem(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ListBatchWebhooksResponseLinksItem{}
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
		var obj *ListBatchWebhooksResponseLinksItem
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}
