// Code generated from our API definition. DO NOT EDIT.

package api

import (
	json "encoding/json"
	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
	testing "testing"
)

func TestSettersListPingResponse(t *testing.T) {
	t.Run("SetHealthStatus", func(t *testing.T) {
		obj := &ListPingResponse{}
		var fernTestValueHealthStatus *string
		obj.SetHealthStatus(fernTestValueHealthStatus)
		assert.Equal(t, fernTestValueHealthStatus, obj.HealthStatus)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersListPingResponse(t *testing.T) {
	t.Run("GetHealthStatus", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListPingResponse{}
		var expected *string
		obj.HealthStatus = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetHealthStatus(), "getter should return the property value")
	})

	t.Run("GetHealthStatus_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListPingResponse{}
		obj.HealthStatus = nil

		// Act & Assert
		assert.Nil(t, obj.GetHealthStatus(), "getter should return nil when property is nil")
	})

	t.Run("GetHealthStatus_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListPingResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetHealthStatus() // Should return zero value
	})

}

func TestSettersMarkExplicitListPingResponse(t *testing.T) {
	t.Run("SetHealthStatus_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListPingResponse{}
		var fernTestValueHealthStatus *string

		// Act
		obj.SetHealthStatus(fernTestValueHealthStatus)

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

func TestJSONMarshalingListPingResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListPingResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ListPingResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ListPingResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ListPingResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestStringListPingResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ListPingResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListPingResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestExtraPropertiesListPingResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ListPingResponse{}
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
		var obj *ListPingResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}
