// Code generated from our API definition. DO NOT EDIT.

package api

import (
	json "encoding/json"
	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
	testing "testing"
	time "time"
)

func TestSettersCreateAudienceContactRequest(t *testing.T) {
	t.Run("SetAudienceID", func(t *testing.T) {
		obj := &CreateAudienceContactRequest{}
		var fernTestValueAudienceID string
		obj.SetAudienceID(fernTestValueAudienceID)
		assert.Equal(t, fernTestValueAudienceID, obj.AudienceID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMergeFieldValidationMode", func(t *testing.T) {
		obj := &CreateAudienceContactRequest{}
		var fernTestValueMergeFieldValidationMode *CreateAudienceContactRequestMergeFieldValidationMode
		obj.SetMergeFieldValidationMode(fernTestValueMergeFieldValidationMode)
		assert.Equal(t, fernTestValueMergeFieldValidationMode, obj.MergeFieldValidationMode)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetDataMode", func(t *testing.T) {
		obj := &CreateAudienceContactRequest{}
		var fernTestValueDataMode *CreateAudienceContactRequestDataMode
		obj.SetDataMode(fernTestValueDataMode)
		assert.Equal(t, fernTestValueDataMode, obj.DataMode)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetEmailChannel", func(t *testing.T) {
		obj := &CreateAudienceContactRequest{}
		var fernTestValueEmailChannel *CreateAudienceContactRequestEmailChannel
		obj.SetEmailChannel(fernTestValueEmailChannel)
		assert.Equal(t, fernTestValueEmailChannel, obj.EmailChannel)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetLanguage", func(t *testing.T) {
		obj := &CreateAudienceContactRequest{}
		var fernTestValueLanguage *string
		obj.SetLanguage(fernTestValueLanguage)
		assert.Equal(t, fernTestValueLanguage, obj.Language)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMergeFields", func(t *testing.T) {
		obj := &CreateAudienceContactRequest{}
		var fernTestValueMergeFields map[string]*CreateAudienceContactRequestMergeFieldsValue
		obj.SetMergeFields(fernTestValueMergeFields)
		assert.Equal(t, fernTestValueMergeFields, obj.MergeFields)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSmsChannel", func(t *testing.T) {
		obj := &CreateAudienceContactRequest{}
		var fernTestValueSmsChannel *CreateAudienceContactRequestSmsChannel
		obj.SetSmsChannel(fernTestValueSmsChannel)
		assert.Equal(t, fernTestValueSmsChannel, obj.SmsChannel)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTags", func(t *testing.T) {
		obj := &CreateAudienceContactRequest{}
		var fernTestValueTags []*CreateAudienceContactRequestTagsItem
		obj.SetTags(fernTestValueTags)
		assert.Equal(t, fernTestValueTags, obj.Tags)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetUpdateExisting", func(t *testing.T) {
		obj := &CreateAudienceContactRequest{}
		var fernTestValueUpdateExisting *bool
		obj.SetUpdateExisting(fernTestValueUpdateExisting)
		assert.Equal(t, fernTestValueUpdateExisting, obj.UpdateExisting)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitCreateAudienceContactRequest(t *testing.T) {
	t.Run("SetAudienceID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequest{}
		var fernTestValueAudienceID string

		// Act
		obj.SetAudienceID(fernTestValueAudienceID)

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

	t.Run("SetMergeFieldValidationMode_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequest{}
		var fernTestValueMergeFieldValidationMode *CreateAudienceContactRequestMergeFieldValidationMode

		// Act
		obj.SetMergeFieldValidationMode(fernTestValueMergeFieldValidationMode)

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

	t.Run("SetDataMode_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequest{}
		var fernTestValueDataMode *CreateAudienceContactRequestDataMode

		// Act
		obj.SetDataMode(fernTestValueDataMode)

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

	t.Run("SetEmailChannel_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequest{}
		var fernTestValueEmailChannel *CreateAudienceContactRequestEmailChannel

		// Act
		obj.SetEmailChannel(fernTestValueEmailChannel)

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

	t.Run("SetLanguage_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequest{}
		var fernTestValueLanguage *string

		// Act
		obj.SetLanguage(fernTestValueLanguage)

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

	t.Run("SetMergeFields_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequest{}
		var fernTestValueMergeFields map[string]*CreateAudienceContactRequestMergeFieldsValue

		// Act
		obj.SetMergeFields(fernTestValueMergeFields)

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

	t.Run("SetSmsChannel_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequest{}
		var fernTestValueSmsChannel *CreateAudienceContactRequestSmsChannel

		// Act
		obj.SetSmsChannel(fernTestValueSmsChannel)

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

	t.Run("SetTags_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequest{}
		var fernTestValueTags []*CreateAudienceContactRequestTagsItem

		// Act
		obj.SetTags(fernTestValueTags)

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

	t.Run("SetUpdateExisting_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequest{}
		var fernTestValueUpdateExisting *bool

		// Act
		obj.SetUpdateExisting(fernTestValueUpdateExisting)

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

func TestSettersGetAudienceContactRequest(t *testing.T) {
	t.Run("SetAudienceID", func(t *testing.T) {
		obj := &GetAudienceContactRequest{}
		var fernTestValueAudienceID string
		obj.SetAudienceID(fernTestValueAudienceID)
		assert.Equal(t, fernTestValueAudienceID, obj.AudienceID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetContactID", func(t *testing.T) {
		obj := &GetAudienceContactRequest{}
		var fernTestValueContactID string
		obj.SetContactID(fernTestValueContactID)
		assert.Equal(t, fernTestValueContactID, obj.ContactID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetFields", func(t *testing.T) {
		obj := &GetAudienceContactRequest{}
		var fernTestValueFields []*string
		obj.SetFields(fernTestValueFields)
		assert.Equal(t, fernTestValueFields, obj.Fields)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetExcludeFields", func(t *testing.T) {
		obj := &GetAudienceContactRequest{}
		var fernTestValueExcludeFields []*string
		obj.SetExcludeFields(fernTestValueExcludeFields)
		assert.Equal(t, fernTestValueExcludeFields, obj.ExcludeFields)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitGetAudienceContactRequest(t *testing.T) {
	t.Run("SetAudienceID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetAudienceContactRequest{}
		var fernTestValueAudienceID string

		// Act
		obj.SetAudienceID(fernTestValueAudienceID)

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

	t.Run("SetContactID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetAudienceContactRequest{}
		var fernTestValueContactID string

		// Act
		obj.SetContactID(fernTestValueContactID)

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
		obj := &GetAudienceContactRequest{}
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
		obj := &GetAudienceContactRequest{}
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

func TestSettersGetAudienceContactListRequest(t *testing.T) {
	t.Run("SetAudienceID", func(t *testing.T) {
		obj := &GetAudienceContactListRequest{}
		var fernTestValueAudienceID string
		obj.SetAudienceID(fernTestValueAudienceID)
		assert.Equal(t, fernTestValueAudienceID, obj.AudienceID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetFields", func(t *testing.T) {
		obj := &GetAudienceContactListRequest{}
		var fernTestValueFields []*string
		obj.SetFields(fernTestValueFields)
		assert.Equal(t, fernTestValueFields, obj.Fields)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetExcludeFields", func(t *testing.T) {
		obj := &GetAudienceContactListRequest{}
		var fernTestValueExcludeFields []*string
		obj.SetExcludeFields(fernTestValueExcludeFields)
		assert.Equal(t, fernTestValueExcludeFields, obj.ExcludeFields)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCount", func(t *testing.T) {
		obj := &GetAudienceContactListRequest{}
		var fernTestValueCount *int
		obj.SetCount(fernTestValueCount)
		assert.Equal(t, fernTestValueCount, obj.Count)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCursor", func(t *testing.T) {
		obj := &GetAudienceContactListRequest{}
		var fernTestValueCursor *string
		obj.SetCursor(fernTestValueCursor)
		assert.Equal(t, fernTestValueCursor, obj.Cursor)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCreatedBefore", func(t *testing.T) {
		obj := &GetAudienceContactListRequest{}
		var fernTestValueCreatedBefore *time.Time
		obj.SetCreatedBefore(fernTestValueCreatedBefore)
		assert.Equal(t, fernTestValueCreatedBefore, obj.CreatedBefore)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCreatedSince", func(t *testing.T) {
		obj := &GetAudienceContactListRequest{}
		var fernTestValueCreatedSince *time.Time
		obj.SetCreatedSince(fernTestValueCreatedSince)
		assert.Equal(t, fernTestValueCreatedSince, obj.CreatedSince)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetUpdatedBefore", func(t *testing.T) {
		obj := &GetAudienceContactListRequest{}
		var fernTestValueUpdatedBefore *time.Time
		obj.SetUpdatedBefore(fernTestValueUpdatedBefore)
		assert.Equal(t, fernTestValueUpdatedBefore, obj.UpdatedBefore)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetUpdatedSince", func(t *testing.T) {
		obj := &GetAudienceContactListRequest{}
		var fernTestValueUpdatedSince *time.Time
		obj.SetUpdatedSince(fernTestValueUpdatedSince)
		assert.Equal(t, fernTestValueUpdatedSince, obj.UpdatedSince)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSortField", func(t *testing.T) {
		obj := &GetAudienceContactListRequest{}
		var fernTestValueSortField *GetAudienceContactListRequestSortField
		obj.SetSortField(fernTestValueSortField)
		assert.Equal(t, fernTestValueSortField, obj.SortField)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSortDir", func(t *testing.T) {
		obj := &GetAudienceContactListRequest{}
		var fernTestValueSortDir *GetAudienceContactListRequestSortDir
		obj.SetSortDir(fernTestValueSortDir)
		assert.Equal(t, fernTestValueSortDir, obj.SortDir)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitGetAudienceContactListRequest(t *testing.T) {
	t.Run("SetAudienceID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetAudienceContactListRequest{}
		var fernTestValueAudienceID string

		// Act
		obj.SetAudienceID(fernTestValueAudienceID)

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
		obj := &GetAudienceContactListRequest{}
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
		obj := &GetAudienceContactListRequest{}
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
		obj := &GetAudienceContactListRequest{}
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

	t.Run("SetCursor_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetAudienceContactListRequest{}
		var fernTestValueCursor *string

		// Act
		obj.SetCursor(fernTestValueCursor)

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

	t.Run("SetCreatedBefore_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetAudienceContactListRequest{}
		var fernTestValueCreatedBefore *time.Time

		// Act
		obj.SetCreatedBefore(fernTestValueCreatedBefore)

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

	t.Run("SetCreatedSince_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetAudienceContactListRequest{}
		var fernTestValueCreatedSince *time.Time

		// Act
		obj.SetCreatedSince(fernTestValueCreatedSince)

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

	t.Run("SetUpdatedBefore_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetAudienceContactListRequest{}
		var fernTestValueUpdatedBefore *time.Time

		// Act
		obj.SetUpdatedBefore(fernTestValueUpdatedBefore)

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

	t.Run("SetUpdatedSince_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetAudienceContactListRequest{}
		var fernTestValueUpdatedSince *time.Time

		// Act
		obj.SetUpdatedSince(fernTestValueUpdatedSince)

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

	t.Run("SetSortField_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetAudienceContactListRequest{}
		var fernTestValueSortField *GetAudienceContactListRequestSortField

		// Act
		obj.SetSortField(fernTestValueSortField)

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

	t.Run("SetSortDir_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetAudienceContactListRequest{}
		var fernTestValueSortDir *GetAudienceContactListRequestSortDir

		// Act
		obj.SetSortDir(fernTestValueSortDir)

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

func TestSettersPatchAudienceContactRequest(t *testing.T) {
	t.Run("SetAudienceID", func(t *testing.T) {
		obj := &PatchAudienceContactRequest{}
		var fernTestValueAudienceID string
		obj.SetAudienceID(fernTestValueAudienceID)
		assert.Equal(t, fernTestValueAudienceID, obj.AudienceID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetContactID", func(t *testing.T) {
		obj := &PatchAudienceContactRequest{}
		var fernTestValueContactID string
		obj.SetContactID(fernTestValueContactID)
		assert.Equal(t, fernTestValueContactID, obj.ContactID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMergeFieldValidationMode", func(t *testing.T) {
		obj := &PatchAudienceContactRequest{}
		var fernTestValueMergeFieldValidationMode *PatchAudienceContactRequestMergeFieldValidationMode
		obj.SetMergeFieldValidationMode(fernTestValueMergeFieldValidationMode)
		assert.Equal(t, fernTestValueMergeFieldValidationMode, obj.MergeFieldValidationMode)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetDataMode", func(t *testing.T) {
		obj := &PatchAudienceContactRequest{}
		var fernTestValueDataMode *PatchAudienceContactRequestDataMode
		obj.SetDataMode(fernTestValueDataMode)
		assert.Equal(t, fernTestValueDataMode, obj.DataMode)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetEmailChannel", func(t *testing.T) {
		obj := &PatchAudienceContactRequest{}
		var fernTestValueEmailChannel *PatchAudienceContactRequestEmailChannel
		obj.SetEmailChannel(fernTestValueEmailChannel)
		assert.Equal(t, fernTestValueEmailChannel, obj.EmailChannel)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetLanguage", func(t *testing.T) {
		obj := &PatchAudienceContactRequest{}
		var fernTestValueLanguage *string
		obj.SetLanguage(fernTestValueLanguage)
		assert.Equal(t, fernTestValueLanguage, obj.Language)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMergeFields", func(t *testing.T) {
		obj := &PatchAudienceContactRequest{}
		var fernTestValueMergeFields map[string]*PatchAudienceContactRequestMergeFieldsValue
		obj.SetMergeFields(fernTestValueMergeFields)
		assert.Equal(t, fernTestValueMergeFields, obj.MergeFields)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSmsChannel", func(t *testing.T) {
		obj := &PatchAudienceContactRequest{}
		var fernTestValueSmsChannel *PatchAudienceContactRequestSmsChannel
		obj.SetSmsChannel(fernTestValueSmsChannel)
		assert.Equal(t, fernTestValueSmsChannel, obj.SmsChannel)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTags", func(t *testing.T) {
		obj := &PatchAudienceContactRequest{}
		var fernTestValueTags []*PatchAudienceContactRequestTagsItem
		obj.SetTags(fernTestValueTags)
		assert.Equal(t, fernTestValueTags, obj.Tags)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitPatchAudienceContactRequest(t *testing.T) {
	t.Run("SetAudienceID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequest{}
		var fernTestValueAudienceID string

		// Act
		obj.SetAudienceID(fernTestValueAudienceID)

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

	t.Run("SetContactID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequest{}
		var fernTestValueContactID string

		// Act
		obj.SetContactID(fernTestValueContactID)

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

	t.Run("SetMergeFieldValidationMode_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequest{}
		var fernTestValueMergeFieldValidationMode *PatchAudienceContactRequestMergeFieldValidationMode

		// Act
		obj.SetMergeFieldValidationMode(fernTestValueMergeFieldValidationMode)

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

	t.Run("SetDataMode_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequest{}
		var fernTestValueDataMode *PatchAudienceContactRequestDataMode

		// Act
		obj.SetDataMode(fernTestValueDataMode)

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

	t.Run("SetEmailChannel_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequest{}
		var fernTestValueEmailChannel *PatchAudienceContactRequestEmailChannel

		// Act
		obj.SetEmailChannel(fernTestValueEmailChannel)

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

	t.Run("SetLanguage_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequest{}
		var fernTestValueLanguage *string

		// Act
		obj.SetLanguage(fernTestValueLanguage)

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

	t.Run("SetMergeFields_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequest{}
		var fernTestValueMergeFields map[string]*PatchAudienceContactRequestMergeFieldsValue

		// Act
		obj.SetMergeFields(fernTestValueMergeFields)

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

	t.Run("SetSmsChannel_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequest{}
		var fernTestValueSmsChannel *PatchAudienceContactRequestSmsChannel

		// Act
		obj.SetSmsChannel(fernTestValueSmsChannel)

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

	t.Run("SetTags_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequest{}
		var fernTestValueTags []*PatchAudienceContactRequestTagsItem

		// Act
		obj.SetTags(fernTestValueTags)

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

func TestSettersPostAudiencesContactsActionsArchiveRequest(t *testing.T) {
	t.Run("SetAudienceID", func(t *testing.T) {
		obj := &PostAudiencesContactsActionsArchiveRequest{}
		var fernTestValueAudienceID string
		obj.SetAudienceID(fernTestValueAudienceID)
		assert.Equal(t, fernTestValueAudienceID, obj.AudienceID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetContactID", func(t *testing.T) {
		obj := &PostAudiencesContactsActionsArchiveRequest{}
		var fernTestValueContactID string
		obj.SetContactID(fernTestValueContactID)
		assert.Equal(t, fernTestValueContactID, obj.ContactID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitPostAudiencesContactsActionsArchiveRequest(t *testing.T) {
	t.Run("SetAudienceID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PostAudiencesContactsActionsArchiveRequest{}
		var fernTestValueAudienceID string

		// Act
		obj.SetAudienceID(fernTestValueAudienceID)

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

	t.Run("SetContactID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PostAudiencesContactsActionsArchiveRequest{}
		var fernTestValueContactID string

		// Act
		obj.SetContactID(fernTestValueContactID)

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

func TestSettersPostAudiencesContactsActionsForgetRequest(t *testing.T) {
	t.Run("SetAudienceID", func(t *testing.T) {
		obj := &PostAudiencesContactsActionsForgetRequest{}
		var fernTestValueAudienceID string
		obj.SetAudienceID(fernTestValueAudienceID)
		assert.Equal(t, fernTestValueAudienceID, obj.AudienceID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetContactID", func(t *testing.T) {
		obj := &PostAudiencesContactsActionsForgetRequest{}
		var fernTestValueContactID string
		obj.SetContactID(fernTestValueContactID)
		assert.Equal(t, fernTestValueContactID, obj.ContactID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitPostAudiencesContactsActionsForgetRequest(t *testing.T) {
	t.Run("SetAudienceID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PostAudiencesContactsActionsForgetRequest{}
		var fernTestValueAudienceID string

		// Act
		obj.SetAudienceID(fernTestValueAudienceID)

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

	t.Run("SetContactID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PostAudiencesContactsActionsForgetRequest{}
		var fernTestValueContactID string

		// Act
		obj.SetContactID(fernTestValueContactID)

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

func TestSettersAudiencesContact(t *testing.T) {
	t.Run("SetAudienceID", func(t *testing.T) {
		obj := &AudiencesContact{}
		var fernTestValueAudienceID *string
		obj.SetAudienceID(fernTestValueAudienceID)
		assert.Equal(t, fernTestValueAudienceID, obj.AudienceID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCreatedAt", func(t *testing.T) {
		obj := &AudiencesContact{}
		var fernTestValueCreatedAt *time.Time
		obj.SetCreatedAt(fernTestValueCreatedAt)
		assert.Equal(t, fernTestValueCreatedAt, obj.CreatedAt)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetEmailChannel", func(t *testing.T) {
		obj := &AudiencesContact{}
		var fernTestValueEmailChannel *AudiencesContactEmailChannel
		obj.SetEmailChannel(fernTestValueEmailChannel)
		assert.Equal(t, fernTestValueEmailChannel, obj.EmailChannel)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetID", func(t *testing.T) {
		obj := &AudiencesContact{}
		var fernTestValueID *string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetLanguage", func(t *testing.T) {
		obj := &AudiencesContact{}
		var fernTestValueLanguage *AudiencesContactLanguage
		obj.SetLanguage(fernTestValueLanguage)
		assert.Equal(t, fernTestValueLanguage, obj.Language)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetLastUpdatedAt", func(t *testing.T) {
		obj := &AudiencesContact{}
		var fernTestValueLastUpdatedAt *time.Time
		obj.SetLastUpdatedAt(fernTestValueLastUpdatedAt)
		assert.Equal(t, fernTestValueLastUpdatedAt, obj.LastUpdatedAt)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMergeFields", func(t *testing.T) {
		obj := &AudiencesContact{}
		var fernTestValueMergeFields map[string]*AudiencesContactMergeFieldsValue
		obj.SetMergeFields(fernTestValueMergeFields)
		assert.Equal(t, fernTestValueMergeFields, obj.MergeFields)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSmsChannel", func(t *testing.T) {
		obj := &AudiencesContact{}
		var fernTestValueSmsChannel *AudiencesContactSmsChannel
		obj.SetSmsChannel(fernTestValueSmsChannel)
		assert.Equal(t, fernTestValueSmsChannel, obj.SmsChannel)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSource", func(t *testing.T) {
		obj := &AudiencesContact{}
		var fernTestValueSource *AudiencesContactSource
		obj.SetSource(fernTestValueSource)
		assert.Equal(t, fernTestValueSource, obj.Source)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStatus", func(t *testing.T) {
		obj := &AudiencesContact{}
		var fernTestValueStatus *AudiencesContactStatus
		obj.SetStatus(fernTestValueStatus)
		assert.Equal(t, fernTestValueStatus, obj.Status)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTags", func(t *testing.T) {
		obj := &AudiencesContact{}
		var fernTestValueTags []string
		obj.SetTags(fernTestValueTags)
		assert.Equal(t, fernTestValueTags, obj.Tags)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersAudiencesContact(t *testing.T) {
	t.Run("GetAudienceID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		var expected *string
		obj.AudienceID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetAudienceID(), "getter should return the property value")
	})

	t.Run("GetAudienceID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		obj.AudienceID = nil

		// Act & Assert
		assert.Nil(t, obj.GetAudienceID(), "getter should return nil when property is nil")
	})

	t.Run("GetAudienceID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContact
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetAudienceID() // Should return zero value
	})

	t.Run("GetCreatedAt", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		var expected *time.Time
		obj.CreatedAt = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCreatedAt(), "getter should return the property value")
	})

	t.Run("GetCreatedAt_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		obj.CreatedAt = nil

		// Act & Assert
		assert.Nil(t, obj.GetCreatedAt(), "getter should return nil when property is nil")
	})

	t.Run("GetCreatedAt_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContact
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCreatedAt() // Should return zero value
	})

	t.Run("GetEmailChannel", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		var expected *AudiencesContactEmailChannel
		obj.EmailChannel = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEmailChannel(), "getter should return the property value")
	})

	t.Run("GetEmailChannel_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		obj.EmailChannel = nil

		// Act & Assert
		assert.Nil(t, obj.GetEmailChannel(), "getter should return nil when property is nil")
	})

	t.Run("GetEmailChannel_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContact
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEmailChannel() // Should return zero value
	})

	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		var expected *string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		obj.ID = nil

		// Act & Assert
		assert.Nil(t, obj.GetID(), "getter should return nil when property is nil")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContact
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetID() // Should return zero value
	})

	t.Run("GetLanguage", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		var expected *AudiencesContactLanguage
		obj.Language = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetLanguage(), "getter should return the property value")
	})

	t.Run("GetLanguage_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		obj.Language = nil

		// Act & Assert
		assert.Nil(t, obj.GetLanguage(), "getter should return nil when property is nil")
	})

	t.Run("GetLanguage_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContact
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetLanguage() // Should return zero value
	})

	t.Run("GetLastUpdatedAt", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		var expected *time.Time
		obj.LastUpdatedAt = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetLastUpdatedAt(), "getter should return the property value")
	})

	t.Run("GetLastUpdatedAt_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		obj.LastUpdatedAt = nil

		// Act & Assert
		assert.Nil(t, obj.GetLastUpdatedAt(), "getter should return nil when property is nil")
	})

	t.Run("GetLastUpdatedAt_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContact
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetLastUpdatedAt() // Should return zero value
	})

	t.Run("GetMergeFields", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		var expected map[string]*AudiencesContactMergeFieldsValue
		obj.MergeFields = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMergeFields(), "getter should return the property value")
	})

	t.Run("GetMergeFields_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		obj.MergeFields = nil

		// Act & Assert
		assert.Nil(t, obj.GetMergeFields(), "getter should return nil when property is nil")
	})

	t.Run("GetMergeFields_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContact
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetMergeFields() // Should return zero value
	})

	t.Run("GetSmsChannel", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		var expected *AudiencesContactSmsChannel
		obj.SmsChannel = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSmsChannel(), "getter should return the property value")
	})

	t.Run("GetSmsChannel_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		obj.SmsChannel = nil

		// Act & Assert
		assert.Nil(t, obj.GetSmsChannel(), "getter should return nil when property is nil")
	})

	t.Run("GetSmsChannel_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContact
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSmsChannel() // Should return zero value
	})

	t.Run("GetSource", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		var expected *AudiencesContactSource
		obj.Source = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSource(), "getter should return the property value")
	})

	t.Run("GetSource_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		obj.Source = nil

		// Act & Assert
		assert.Nil(t, obj.GetSource(), "getter should return nil when property is nil")
	})

	t.Run("GetSource_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContact
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSource() // Should return zero value
	})

	t.Run("GetStatus", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		var expected *AudiencesContactStatus
		obj.Status = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStatus(), "getter should return the property value")
	})

	t.Run("GetStatus_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		obj.Status = nil

		// Act & Assert
		assert.Nil(t, obj.GetStatus(), "getter should return nil when property is nil")
	})

	t.Run("GetStatus_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContact
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStatus() // Should return zero value
	})

	t.Run("GetTags", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		var expected []string
		obj.Tags = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTags(), "getter should return the property value")
	})

	t.Run("GetTags_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		obj.Tags = nil

		// Act & Assert
		assert.Nil(t, obj.GetTags(), "getter should return nil when property is nil")
	})

	t.Run("GetTags_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContact
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTags() // Should return zero value
	})

}

func TestSettersMarkExplicitAudiencesContact(t *testing.T) {
	t.Run("SetAudienceID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		var fernTestValueAudienceID *string

		// Act
		obj.SetAudienceID(fernTestValueAudienceID)

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

	t.Run("SetCreatedAt_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		var fernTestValueCreatedAt *time.Time

		// Act
		obj.SetCreatedAt(fernTestValueCreatedAt)

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

	t.Run("SetEmailChannel_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		var fernTestValueEmailChannel *AudiencesContactEmailChannel

		// Act
		obj.SetEmailChannel(fernTestValueEmailChannel)

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
		obj := &AudiencesContact{}
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

	t.Run("SetLanguage_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		var fernTestValueLanguage *AudiencesContactLanguage

		// Act
		obj.SetLanguage(fernTestValueLanguage)

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

	t.Run("SetLastUpdatedAt_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		var fernTestValueLastUpdatedAt *time.Time

		// Act
		obj.SetLastUpdatedAt(fernTestValueLastUpdatedAt)

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

	t.Run("SetMergeFields_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		var fernTestValueMergeFields map[string]*AudiencesContactMergeFieldsValue

		// Act
		obj.SetMergeFields(fernTestValueMergeFields)

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

	t.Run("SetSmsChannel_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		var fernTestValueSmsChannel *AudiencesContactSmsChannel

		// Act
		obj.SetSmsChannel(fernTestValueSmsChannel)

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

	t.Run("SetSource_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		var fernTestValueSource *AudiencesContactSource

		// Act
		obj.SetSource(fernTestValueSource)

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
		obj := &AudiencesContact{}
		var fernTestValueStatus *AudiencesContactStatus

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

	t.Run("SetTags_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}
		var fernTestValueTags []string

		// Act
		obj.SetTags(fernTestValueTags)

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

func TestSettersAudiencesContactEmailChannel(t *testing.T) {
	t.Run("SetEffectiveSubscriptionStatus", func(t *testing.T) {
		obj := &AudiencesContactEmailChannel{}
		var fernTestValueEffectiveSubscriptionStatus *AudiencesContactEmailChannelEffectiveSubscriptionStatus
		obj.SetEffectiveSubscriptionStatus(fernTestValueEffectiveSubscriptionStatus)
		assert.Equal(t, fernTestValueEffectiveSubscriptionStatus, obj.EffectiveSubscriptionStatus)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetEmail", func(t *testing.T) {
		obj := &AudiencesContactEmailChannel{}
		var fernTestValueEmail *string
		obj.SetEmail(fernTestValueEmail)
		assert.Equal(t, fernTestValueEmail, obj.Email)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetHashedEmail", func(t *testing.T) {
		obj := &AudiencesContactEmailChannel{}
		var fernTestValueHashedEmail *string
		obj.SetHashedEmail(fernTestValueHashedEmail)
		assert.Equal(t, fernTestValueHashedEmail, obj.HashedEmail)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMarketingConsent", func(t *testing.T) {
		obj := &AudiencesContactEmailChannel{}
		var fernTestValueMarketingConsent *AudiencesContactEmailChannelMarketingConsent
		obj.SetMarketingConsent(fernTestValueMarketingConsent)
		assert.Equal(t, fernTestValueMarketingConsent, obj.MarketingConsent)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSource", func(t *testing.T) {
		obj := &AudiencesContactEmailChannel{}
		var fernTestValueSource *AudiencesContactEmailChannelSource
		obj.SetSource(fernTestValueSource)
		assert.Equal(t, fernTestValueSource, obj.Source)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersAudiencesContactEmailChannel(t *testing.T) {
	t.Run("GetEffectiveSubscriptionStatus", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannel{}
		var expected *AudiencesContactEmailChannelEffectiveSubscriptionStatus
		obj.EffectiveSubscriptionStatus = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEffectiveSubscriptionStatus(), "getter should return the property value")
	})

	t.Run("GetEffectiveSubscriptionStatus_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannel{}
		obj.EffectiveSubscriptionStatus = nil

		// Act & Assert
		assert.Nil(t, obj.GetEffectiveSubscriptionStatus(), "getter should return nil when property is nil")
	})

	t.Run("GetEffectiveSubscriptionStatus_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactEmailChannel
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEffectiveSubscriptionStatus() // Should return zero value
	})

	t.Run("GetEmail", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannel{}
		var expected *string
		obj.Email = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEmail(), "getter should return the property value")
	})

	t.Run("GetEmail_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannel{}
		obj.Email = nil

		// Act & Assert
		assert.Nil(t, obj.GetEmail(), "getter should return nil when property is nil")
	})

	t.Run("GetEmail_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactEmailChannel
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEmail() // Should return zero value
	})

	t.Run("GetHashedEmail", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannel{}
		var expected *string
		obj.HashedEmail = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetHashedEmail(), "getter should return the property value")
	})

	t.Run("GetHashedEmail_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannel{}
		obj.HashedEmail = nil

		// Act & Assert
		assert.Nil(t, obj.GetHashedEmail(), "getter should return nil when property is nil")
	})

	t.Run("GetHashedEmail_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactEmailChannel
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetHashedEmail() // Should return zero value
	})

	t.Run("GetMarketingConsent", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannel{}
		var expected *AudiencesContactEmailChannelMarketingConsent
		obj.MarketingConsent = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMarketingConsent(), "getter should return the property value")
	})

	t.Run("GetMarketingConsent_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannel{}
		obj.MarketingConsent = nil

		// Act & Assert
		assert.Nil(t, obj.GetMarketingConsent(), "getter should return nil when property is nil")
	})

	t.Run("GetMarketingConsent_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactEmailChannel
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetMarketingConsent() // Should return zero value
	})

	t.Run("GetSource", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannel{}
		var expected *AudiencesContactEmailChannelSource
		obj.Source = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSource(), "getter should return the property value")
	})

	t.Run("GetSource_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannel{}
		obj.Source = nil

		// Act & Assert
		assert.Nil(t, obj.GetSource(), "getter should return nil when property is nil")
	})

	t.Run("GetSource_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactEmailChannel
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSource() // Should return zero value
	})

}

func TestSettersMarkExplicitAudiencesContactEmailChannel(t *testing.T) {
	t.Run("SetEffectiveSubscriptionStatus_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannel{}
		var fernTestValueEffectiveSubscriptionStatus *AudiencesContactEmailChannelEffectiveSubscriptionStatus

		// Act
		obj.SetEffectiveSubscriptionStatus(fernTestValueEffectiveSubscriptionStatus)

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

	t.Run("SetEmail_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannel{}
		var fernTestValueEmail *string

		// Act
		obj.SetEmail(fernTestValueEmail)

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

	t.Run("SetHashedEmail_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannel{}
		var fernTestValueHashedEmail *string

		// Act
		obj.SetHashedEmail(fernTestValueHashedEmail)

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

	t.Run("SetMarketingConsent_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannel{}
		var fernTestValueMarketingConsent *AudiencesContactEmailChannelMarketingConsent

		// Act
		obj.SetMarketingConsent(fernTestValueMarketingConsent)

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

	t.Run("SetSource_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannel{}
		var fernTestValueSource *AudiencesContactEmailChannelSource

		// Act
		obj.SetSource(fernTestValueSource)

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

func TestSettersAudiencesContactEmailChannelEffectiveSubscriptionStatus(t *testing.T) {
	t.Run("SetValue", func(t *testing.T) {
		obj := &AudiencesContactEmailChannelEffectiveSubscriptionStatus{}
		var fernTestValueValue *AudiencesContactEmailChannelEffectiveSubscriptionStatusValue
		obj.SetValue(fernTestValueValue)
		assert.Equal(t, fernTestValueValue, obj.Value)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersAudiencesContactEmailChannelEffectiveSubscriptionStatus(t *testing.T) {
	t.Run("GetValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannelEffectiveSubscriptionStatus{}
		var expected *AudiencesContactEmailChannelEffectiveSubscriptionStatusValue
		obj.Value = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetValue(), "getter should return the property value")
	})

	t.Run("GetValue_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannelEffectiveSubscriptionStatus{}
		obj.Value = nil

		// Act & Assert
		assert.Nil(t, obj.GetValue(), "getter should return nil when property is nil")
	})

	t.Run("GetValue_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactEmailChannelEffectiveSubscriptionStatus
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetValue() // Should return zero value
	})

}

func TestSettersMarkExplicitAudiencesContactEmailChannelEffectiveSubscriptionStatus(t *testing.T) {
	t.Run("SetValue_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannelEffectiveSubscriptionStatus{}
		var fernTestValueValue *AudiencesContactEmailChannelEffectiveSubscriptionStatusValue

		// Act
		obj.SetValue(fernTestValueValue)

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

func TestSettersAudiencesContactEmailChannelMarketingConsent(t *testing.T) {
	t.Run("SetSource", func(t *testing.T) {
		obj := &AudiencesContactEmailChannelMarketingConsent{}
		var fernTestValueSource *AudiencesContactEmailChannelMarketingConsentSource
		obj.SetSource(fernTestValueSource)
		assert.Equal(t, fernTestValueSource, obj.Source)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStatus", func(t *testing.T) {
		obj := &AudiencesContactEmailChannelMarketingConsent{}
		var fernTestValueStatus *AudiencesContactEmailChannelMarketingConsentStatus
		obj.SetStatus(fernTestValueStatus)
		assert.Equal(t, fernTestValueStatus, obj.Status)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCapturedAt", func(t *testing.T) {
		obj := &AudiencesContactEmailChannelMarketingConsent{}
		var fernTestValueCapturedAt *time.Time
		obj.SetCapturedAt(fernTestValueCapturedAt)
		assert.Equal(t, fernTestValueCapturedAt, obj.CapturedAt)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersAudiencesContactEmailChannelMarketingConsent(t *testing.T) {
	t.Run("GetSource", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannelMarketingConsent{}
		var expected *AudiencesContactEmailChannelMarketingConsentSource
		obj.Source = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSource(), "getter should return the property value")
	})

	t.Run("GetSource_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannelMarketingConsent{}
		obj.Source = nil

		// Act & Assert
		assert.Nil(t, obj.GetSource(), "getter should return nil when property is nil")
	})

	t.Run("GetSource_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactEmailChannelMarketingConsent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSource() // Should return zero value
	})

	t.Run("GetStatus", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannelMarketingConsent{}
		var expected *AudiencesContactEmailChannelMarketingConsentStatus
		obj.Status = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStatus(), "getter should return the property value")
	})

	t.Run("GetStatus_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannelMarketingConsent{}
		obj.Status = nil

		// Act & Assert
		assert.Nil(t, obj.GetStatus(), "getter should return nil when property is nil")
	})

	t.Run("GetStatus_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactEmailChannelMarketingConsent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStatus() // Should return zero value
	})

	t.Run("GetCapturedAt", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannelMarketingConsent{}
		var expected *time.Time
		obj.CapturedAt = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCapturedAt(), "getter should return the property value")
	})

	t.Run("GetCapturedAt_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannelMarketingConsent{}
		obj.CapturedAt = nil

		// Act & Assert
		assert.Nil(t, obj.GetCapturedAt(), "getter should return nil when property is nil")
	})

	t.Run("GetCapturedAt_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactEmailChannelMarketingConsent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCapturedAt() // Should return zero value
	})

}

func TestSettersMarkExplicitAudiencesContactEmailChannelMarketingConsent(t *testing.T) {
	t.Run("SetSource_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannelMarketingConsent{}
		var fernTestValueSource *AudiencesContactEmailChannelMarketingConsentSource

		// Act
		obj.SetSource(fernTestValueSource)

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
		obj := &AudiencesContactEmailChannelMarketingConsent{}
		var fernTestValueStatus *AudiencesContactEmailChannelMarketingConsentStatus

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

	t.Run("SetCapturedAt_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannelMarketingConsent{}
		var fernTestValueCapturedAt *time.Time

		// Act
		obj.SetCapturedAt(fernTestValueCapturedAt)

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

func TestSettersAudiencesContactEmailChannelMarketingConsentSource(t *testing.T) {
	t.Run("SetName", func(t *testing.T) {
		obj := &AudiencesContactEmailChannelMarketingConsentSource{}
		var fernTestValueName *string
		obj.SetName(fernTestValueName)
		assert.Equal(t, fernTestValueName, obj.Name)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersAudiencesContactEmailChannelMarketingConsentSource(t *testing.T) {
	t.Run("GetName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannelMarketingConsentSource{}
		var expected *string
		obj.Name = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetName(), "getter should return the property value")
	})

	t.Run("GetName_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannelMarketingConsentSource{}
		obj.Name = nil

		// Act & Assert
		assert.Nil(t, obj.GetName(), "getter should return nil when property is nil")
	})

	t.Run("GetName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactEmailChannelMarketingConsentSource
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetName() // Should return zero value
	})

}

func TestSettersMarkExplicitAudiencesContactEmailChannelMarketingConsentSource(t *testing.T) {
	t.Run("SetName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannelMarketingConsentSource{}
		var fernTestValueName *string

		// Act
		obj.SetName(fernTestValueName)

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

func TestSettersAudiencesContactEmailChannelSource(t *testing.T) {
	t.Run("SetName", func(t *testing.T) {
		obj := &AudiencesContactEmailChannelSource{}
		var fernTestValueName *string
		obj.SetName(fernTestValueName)
		assert.Equal(t, fernTestValueName, obj.Name)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersAudiencesContactEmailChannelSource(t *testing.T) {
	t.Run("GetName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannelSource{}
		var expected *string
		obj.Name = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetName(), "getter should return the property value")
	})

	t.Run("GetName_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannelSource{}
		obj.Name = nil

		// Act & Assert
		assert.Nil(t, obj.GetName(), "getter should return nil when property is nil")
	})

	t.Run("GetName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactEmailChannelSource
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetName() // Should return zero value
	})

}

func TestSettersMarkExplicitAudiencesContactEmailChannelSource(t *testing.T) {
	t.Run("SetName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannelSource{}
		var fernTestValueName *string

		// Act
		obj.SetName(fernTestValueName)

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

func TestGettersAudiencesContactMergeFieldsValue(t *testing.T) {
	t.Run("GetAudiencesContactMergeFieldsValueAddr1", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactMergeFieldsValue{}
		var expected *AudiencesContactMergeFieldsValueAddr1
		obj.AudiencesContactMergeFieldsValueAddr1 = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetAudiencesContactMergeFieldsValueAddr1(), "getter should return the property value")
	})

	t.Run("GetAudiencesContactMergeFieldsValueAddr1_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactMergeFieldsValue{}
		obj.AudiencesContactMergeFieldsValueAddr1 = nil

		// Act & Assert
		assert.Nil(t, obj.GetAudiencesContactMergeFieldsValueAddr1(), "getter should return nil when property is nil")
	})

	t.Run("GetAudiencesContactMergeFieldsValueAddr1_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactMergeFieldsValue
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetAudiencesContactMergeFieldsValueAddr1() // Should return zero value
	})

	t.Run("GetString", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactMergeFieldsValue{}
		var expected string
		obj.String = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetString(), "getter should return the property value")
	})

	t.Run("GetString_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactMergeFieldsValue
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetString() // Should return zero value
	})

	t.Run("GetDouble", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactMergeFieldsValue{}
		var expected float64
		obj.Double = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetDouble(), "getter should return the property value")
	})

	t.Run("GetDouble_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactMergeFieldsValue
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetDouble() // Should return zero value
	})

}

func TestSettersAudiencesContactMergeFieldsValueAddr1(t *testing.T) {
	t.Run("SetAddr1", func(t *testing.T) {
		obj := &AudiencesContactMergeFieldsValueAddr1{}
		var fernTestValueAddr1 string
		obj.SetAddr1(fernTestValueAddr1)
		assert.Equal(t, fernTestValueAddr1, obj.Addr1)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetAddr2", func(t *testing.T) {
		obj := &AudiencesContactMergeFieldsValueAddr1{}
		var fernTestValueAddr2 *string
		obj.SetAddr2(fernTestValueAddr2)
		assert.Equal(t, fernTestValueAddr2, obj.Addr2)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCity", func(t *testing.T) {
		obj := &AudiencesContactMergeFieldsValueAddr1{}
		var fernTestValueCity string
		obj.SetCity(fernTestValueCity)
		assert.Equal(t, fernTestValueCity, obj.City)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetState", func(t *testing.T) {
		obj := &AudiencesContactMergeFieldsValueAddr1{}
		var fernTestValueState string
		obj.SetState(fernTestValueState)
		assert.Equal(t, fernTestValueState, obj.State)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetZip", func(t *testing.T) {
		obj := &AudiencesContactMergeFieldsValueAddr1{}
		var fernTestValueZip string
		obj.SetZip(fernTestValueZip)
		assert.Equal(t, fernTestValueZip, obj.Zip)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCountry", func(t *testing.T) {
		obj := &AudiencesContactMergeFieldsValueAddr1{}
		var fernTestValueCountry *string
		obj.SetCountry(fernTestValueCountry)
		assert.Equal(t, fernTestValueCountry, obj.Country)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersAudiencesContactMergeFieldsValueAddr1(t *testing.T) {
	t.Run("GetAddr1", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactMergeFieldsValueAddr1{}
		var expected string
		obj.Addr1 = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetAddr1(), "getter should return the property value")
	})

	t.Run("GetAddr1_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactMergeFieldsValueAddr1
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetAddr1() // Should return zero value
	})

	t.Run("GetAddr2", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactMergeFieldsValueAddr1{}
		var expected *string
		obj.Addr2 = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetAddr2(), "getter should return the property value")
	})

	t.Run("GetAddr2_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactMergeFieldsValueAddr1{}
		obj.Addr2 = nil

		// Act & Assert
		assert.Nil(t, obj.GetAddr2(), "getter should return nil when property is nil")
	})

	t.Run("GetAddr2_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactMergeFieldsValueAddr1
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetAddr2() // Should return zero value
	})

	t.Run("GetCity", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactMergeFieldsValueAddr1{}
		var expected string
		obj.City = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCity(), "getter should return the property value")
	})

	t.Run("GetCity_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactMergeFieldsValueAddr1
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCity() // Should return zero value
	})

	t.Run("GetState", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactMergeFieldsValueAddr1{}
		var expected string
		obj.State = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetState(), "getter should return the property value")
	})

	t.Run("GetState_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactMergeFieldsValueAddr1
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetState() // Should return zero value
	})

	t.Run("GetZip", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactMergeFieldsValueAddr1{}
		var expected string
		obj.Zip = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetZip(), "getter should return the property value")
	})

	t.Run("GetZip_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactMergeFieldsValueAddr1
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetZip() // Should return zero value
	})

	t.Run("GetCountry", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactMergeFieldsValueAddr1{}
		var expected *string
		obj.Country = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCountry(), "getter should return the property value")
	})

	t.Run("GetCountry_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactMergeFieldsValueAddr1{}
		obj.Country = nil

		// Act & Assert
		assert.Nil(t, obj.GetCountry(), "getter should return nil when property is nil")
	})

	t.Run("GetCountry_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactMergeFieldsValueAddr1
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCountry() // Should return zero value
	})

}

func TestSettersMarkExplicitAudiencesContactMergeFieldsValueAddr1(t *testing.T) {
	t.Run("SetAddr1_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactMergeFieldsValueAddr1{}
		var fernTestValueAddr1 string

		// Act
		obj.SetAddr1(fernTestValueAddr1)

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

	t.Run("SetAddr2_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactMergeFieldsValueAddr1{}
		var fernTestValueAddr2 *string

		// Act
		obj.SetAddr2(fernTestValueAddr2)

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

	t.Run("SetCity_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactMergeFieldsValueAddr1{}
		var fernTestValueCity string

		// Act
		obj.SetCity(fernTestValueCity)

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

	t.Run("SetState_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactMergeFieldsValueAddr1{}
		var fernTestValueState string

		// Act
		obj.SetState(fernTestValueState)

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

	t.Run("SetZip_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactMergeFieldsValueAddr1{}
		var fernTestValueZip string

		// Act
		obj.SetZip(fernTestValueZip)

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

	t.Run("SetCountry_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactMergeFieldsValueAddr1{}
		var fernTestValueCountry *string

		// Act
		obj.SetCountry(fernTestValueCountry)

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

func TestSettersAudiencesContactSmsChannel(t *testing.T) {
	t.Run("SetEffectiveSubscriptionStatus", func(t *testing.T) {
		obj := &AudiencesContactSmsChannel{}
		var fernTestValueEffectiveSubscriptionStatus *AudiencesContactSmsChannelEffectiveSubscriptionStatus
		obj.SetEffectiveSubscriptionStatus(fernTestValueEffectiveSubscriptionStatus)
		assert.Equal(t, fernTestValueEffectiveSubscriptionStatus, obj.EffectiveSubscriptionStatus)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMarketingConsent", func(t *testing.T) {
		obj := &AudiencesContactSmsChannel{}
		var fernTestValueMarketingConsent *AudiencesContactSmsChannelMarketingConsent
		obj.SetMarketingConsent(fernTestValueMarketingConsent)
		assert.Equal(t, fernTestValueMarketingConsent, obj.MarketingConsent)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSmsPhone", func(t *testing.T) {
		obj := &AudiencesContactSmsChannel{}
		var fernTestValueSmsPhone *string
		obj.SetSmsPhone(fernTestValueSmsPhone)
		assert.Equal(t, fernTestValueSmsPhone, obj.SmsPhone)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSource", func(t *testing.T) {
		obj := &AudiencesContactSmsChannel{}
		var fernTestValueSource *AudiencesContactSmsChannelSource
		obj.SetSource(fernTestValueSource)
		assert.Equal(t, fernTestValueSource, obj.Source)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetHashedSmsPhone", func(t *testing.T) {
		obj := &AudiencesContactSmsChannel{}
		var fernTestValueHashedSmsPhone *string
		obj.SetHashedSmsPhone(fernTestValueHashedSmsPhone)
		assert.Equal(t, fernTestValueHashedSmsPhone, obj.HashedSmsPhone)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersAudiencesContactSmsChannel(t *testing.T) {
	t.Run("GetEffectiveSubscriptionStatus", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannel{}
		var expected *AudiencesContactSmsChannelEffectiveSubscriptionStatus
		obj.EffectiveSubscriptionStatus = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEffectiveSubscriptionStatus(), "getter should return the property value")
	})

	t.Run("GetEffectiveSubscriptionStatus_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannel{}
		obj.EffectiveSubscriptionStatus = nil

		// Act & Assert
		assert.Nil(t, obj.GetEffectiveSubscriptionStatus(), "getter should return nil when property is nil")
	})

	t.Run("GetEffectiveSubscriptionStatus_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactSmsChannel
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEffectiveSubscriptionStatus() // Should return zero value
	})

	t.Run("GetMarketingConsent", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannel{}
		var expected *AudiencesContactSmsChannelMarketingConsent
		obj.MarketingConsent = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMarketingConsent(), "getter should return the property value")
	})

	t.Run("GetMarketingConsent_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannel{}
		obj.MarketingConsent = nil

		// Act & Assert
		assert.Nil(t, obj.GetMarketingConsent(), "getter should return nil when property is nil")
	})

	t.Run("GetMarketingConsent_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactSmsChannel
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetMarketingConsent() // Should return zero value
	})

	t.Run("GetSmsPhone", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannel{}
		var expected *string
		obj.SmsPhone = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSmsPhone(), "getter should return the property value")
	})

	t.Run("GetSmsPhone_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannel{}
		obj.SmsPhone = nil

		// Act & Assert
		assert.Nil(t, obj.GetSmsPhone(), "getter should return nil when property is nil")
	})

	t.Run("GetSmsPhone_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactSmsChannel
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSmsPhone() // Should return zero value
	})

	t.Run("GetSource", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannel{}
		var expected *AudiencesContactSmsChannelSource
		obj.Source = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSource(), "getter should return the property value")
	})

	t.Run("GetSource_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannel{}
		obj.Source = nil

		// Act & Assert
		assert.Nil(t, obj.GetSource(), "getter should return nil when property is nil")
	})

	t.Run("GetSource_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactSmsChannel
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSource() // Should return zero value
	})

	t.Run("GetHashedSmsPhone", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannel{}
		var expected *string
		obj.HashedSmsPhone = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetHashedSmsPhone(), "getter should return the property value")
	})

	t.Run("GetHashedSmsPhone_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannel{}
		obj.HashedSmsPhone = nil

		// Act & Assert
		assert.Nil(t, obj.GetHashedSmsPhone(), "getter should return nil when property is nil")
	})

	t.Run("GetHashedSmsPhone_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactSmsChannel
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetHashedSmsPhone() // Should return zero value
	})

}

func TestSettersMarkExplicitAudiencesContactSmsChannel(t *testing.T) {
	t.Run("SetEffectiveSubscriptionStatus_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannel{}
		var fernTestValueEffectiveSubscriptionStatus *AudiencesContactSmsChannelEffectiveSubscriptionStatus

		// Act
		obj.SetEffectiveSubscriptionStatus(fernTestValueEffectiveSubscriptionStatus)

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

	t.Run("SetMarketingConsent_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannel{}
		var fernTestValueMarketingConsent *AudiencesContactSmsChannelMarketingConsent

		// Act
		obj.SetMarketingConsent(fernTestValueMarketingConsent)

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

	t.Run("SetSmsPhone_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannel{}
		var fernTestValueSmsPhone *string

		// Act
		obj.SetSmsPhone(fernTestValueSmsPhone)

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

	t.Run("SetSource_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannel{}
		var fernTestValueSource *AudiencesContactSmsChannelSource

		// Act
		obj.SetSource(fernTestValueSource)

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

	t.Run("SetHashedSmsPhone_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannel{}
		var fernTestValueHashedSmsPhone *string

		// Act
		obj.SetHashedSmsPhone(fernTestValueHashedSmsPhone)

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

func TestSettersAudiencesContactSmsChannelEffectiveSubscriptionStatus(t *testing.T) {
	t.Run("SetValue", func(t *testing.T) {
		obj := &AudiencesContactSmsChannelEffectiveSubscriptionStatus{}
		var fernTestValueValue *AudiencesContactSmsChannelEffectiveSubscriptionStatusValue
		obj.SetValue(fernTestValueValue)
		assert.Equal(t, fernTestValueValue, obj.Value)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersAudiencesContactSmsChannelEffectiveSubscriptionStatus(t *testing.T) {
	t.Run("GetValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannelEffectiveSubscriptionStatus{}
		var expected *AudiencesContactSmsChannelEffectiveSubscriptionStatusValue
		obj.Value = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetValue(), "getter should return the property value")
	})

	t.Run("GetValue_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannelEffectiveSubscriptionStatus{}
		obj.Value = nil

		// Act & Assert
		assert.Nil(t, obj.GetValue(), "getter should return nil when property is nil")
	})

	t.Run("GetValue_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactSmsChannelEffectiveSubscriptionStatus
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetValue() // Should return zero value
	})

}

func TestSettersMarkExplicitAudiencesContactSmsChannelEffectiveSubscriptionStatus(t *testing.T) {
	t.Run("SetValue_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannelEffectiveSubscriptionStatus{}
		var fernTestValueValue *AudiencesContactSmsChannelEffectiveSubscriptionStatusValue

		// Act
		obj.SetValue(fernTestValueValue)

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

func TestSettersAudiencesContactSmsChannelMarketingConsent(t *testing.T) {
	t.Run("SetSource", func(t *testing.T) {
		obj := &AudiencesContactSmsChannelMarketingConsent{}
		var fernTestValueSource *AudiencesContactSmsChannelMarketingConsentSource
		obj.SetSource(fernTestValueSource)
		assert.Equal(t, fernTestValueSource, obj.Source)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStatus", func(t *testing.T) {
		obj := &AudiencesContactSmsChannelMarketingConsent{}
		var fernTestValueStatus *AudiencesContactSmsChannelMarketingConsentStatus
		obj.SetStatus(fernTestValueStatus)
		assert.Equal(t, fernTestValueStatus, obj.Status)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCapturedAt", func(t *testing.T) {
		obj := &AudiencesContactSmsChannelMarketingConsent{}
		var fernTestValueCapturedAt *time.Time
		obj.SetCapturedAt(fernTestValueCapturedAt)
		assert.Equal(t, fernTestValueCapturedAt, obj.CapturedAt)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersAudiencesContactSmsChannelMarketingConsent(t *testing.T) {
	t.Run("GetSource", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannelMarketingConsent{}
		var expected *AudiencesContactSmsChannelMarketingConsentSource
		obj.Source = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSource(), "getter should return the property value")
	})

	t.Run("GetSource_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannelMarketingConsent{}
		obj.Source = nil

		// Act & Assert
		assert.Nil(t, obj.GetSource(), "getter should return nil when property is nil")
	})

	t.Run("GetSource_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactSmsChannelMarketingConsent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSource() // Should return zero value
	})

	t.Run("GetStatus", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannelMarketingConsent{}
		var expected *AudiencesContactSmsChannelMarketingConsentStatus
		obj.Status = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStatus(), "getter should return the property value")
	})

	t.Run("GetStatus_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannelMarketingConsent{}
		obj.Status = nil

		// Act & Assert
		assert.Nil(t, obj.GetStatus(), "getter should return nil when property is nil")
	})

	t.Run("GetStatus_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactSmsChannelMarketingConsent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStatus() // Should return zero value
	})

	t.Run("GetCapturedAt", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannelMarketingConsent{}
		var expected *time.Time
		obj.CapturedAt = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCapturedAt(), "getter should return the property value")
	})

	t.Run("GetCapturedAt_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannelMarketingConsent{}
		obj.CapturedAt = nil

		// Act & Assert
		assert.Nil(t, obj.GetCapturedAt(), "getter should return nil when property is nil")
	})

	t.Run("GetCapturedAt_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactSmsChannelMarketingConsent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCapturedAt() // Should return zero value
	})

}

func TestSettersMarkExplicitAudiencesContactSmsChannelMarketingConsent(t *testing.T) {
	t.Run("SetSource_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannelMarketingConsent{}
		var fernTestValueSource *AudiencesContactSmsChannelMarketingConsentSource

		// Act
		obj.SetSource(fernTestValueSource)

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
		obj := &AudiencesContactSmsChannelMarketingConsent{}
		var fernTestValueStatus *AudiencesContactSmsChannelMarketingConsentStatus

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

	t.Run("SetCapturedAt_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannelMarketingConsent{}
		var fernTestValueCapturedAt *time.Time

		// Act
		obj.SetCapturedAt(fernTestValueCapturedAt)

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

func TestSettersAudiencesContactSmsChannelMarketingConsentSource(t *testing.T) {
	t.Run("SetName", func(t *testing.T) {
		obj := &AudiencesContactSmsChannelMarketingConsentSource{}
		var fernTestValueName *string
		obj.SetName(fernTestValueName)
		assert.Equal(t, fernTestValueName, obj.Name)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersAudiencesContactSmsChannelMarketingConsentSource(t *testing.T) {
	t.Run("GetName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannelMarketingConsentSource{}
		var expected *string
		obj.Name = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetName(), "getter should return the property value")
	})

	t.Run("GetName_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannelMarketingConsentSource{}
		obj.Name = nil

		// Act & Assert
		assert.Nil(t, obj.GetName(), "getter should return nil when property is nil")
	})

	t.Run("GetName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactSmsChannelMarketingConsentSource
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetName() // Should return zero value
	})

}

func TestSettersMarkExplicitAudiencesContactSmsChannelMarketingConsentSource(t *testing.T) {
	t.Run("SetName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannelMarketingConsentSource{}
		var fernTestValueName *string

		// Act
		obj.SetName(fernTestValueName)

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

func TestSettersAudiencesContactSmsChannelSource(t *testing.T) {
	t.Run("SetName", func(t *testing.T) {
		obj := &AudiencesContactSmsChannelSource{}
		var fernTestValueName *string
		obj.SetName(fernTestValueName)
		assert.Equal(t, fernTestValueName, obj.Name)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersAudiencesContactSmsChannelSource(t *testing.T) {
	t.Run("GetName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannelSource{}
		var expected *string
		obj.Name = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetName(), "getter should return the property value")
	})

	t.Run("GetName_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannelSource{}
		obj.Name = nil

		// Act & Assert
		assert.Nil(t, obj.GetName(), "getter should return nil when property is nil")
	})

	t.Run("GetName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactSmsChannelSource
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetName() // Should return zero value
	})

}

func TestSettersMarkExplicitAudiencesContactSmsChannelSource(t *testing.T) {
	t.Run("SetName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannelSource{}
		var fernTestValueName *string

		// Act
		obj.SetName(fernTestValueName)

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

func TestSettersAudiencesContactSource(t *testing.T) {
	t.Run("SetName", func(t *testing.T) {
		obj := &AudiencesContactSource{}
		var fernTestValueName *string
		obj.SetName(fernTestValueName)
		assert.Equal(t, fernTestValueName, obj.Name)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersAudiencesContactSource(t *testing.T) {
	t.Run("GetName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSource{}
		var expected *string
		obj.Name = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetName(), "getter should return the property value")
	})

	t.Run("GetName_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSource{}
		obj.Name = nil

		// Act & Assert
		assert.Nil(t, obj.GetName(), "getter should return nil when property is nil")
	})

	t.Run("GetName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactSource
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetName() // Should return zero value
	})

}

func TestSettersMarkExplicitAudiencesContactSource(t *testing.T) {
	t.Run("SetName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSource{}
		var fernTestValueName *string

		// Act
		obj.SetName(fernTestValueName)

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

func TestSettersCreateAudienceContactRequestEmailChannel(t *testing.T) {
	t.Run("SetEmail", func(t *testing.T) {
		obj := &CreateAudienceContactRequestEmailChannel{}
		var fernTestValueEmail *string
		obj.SetEmail(fernTestValueEmail)
		assert.Equal(t, fernTestValueEmail, obj.Email)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMarketingConsent", func(t *testing.T) {
		obj := &CreateAudienceContactRequestEmailChannel{}
		var fernTestValueMarketingConsent *CreateAudienceContactRequestEmailChannelMarketingConsent
		obj.SetMarketingConsent(fernTestValueMarketingConsent)
		assert.Equal(t, fernTestValueMarketingConsent, obj.MarketingConsent)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersCreateAudienceContactRequestEmailChannel(t *testing.T) {
	t.Run("GetEmail", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestEmailChannel{}
		var expected *string
		obj.Email = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEmail(), "getter should return the property value")
	})

	t.Run("GetEmail_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestEmailChannel{}
		obj.Email = nil

		// Act & Assert
		assert.Nil(t, obj.GetEmail(), "getter should return nil when property is nil")
	})

	t.Run("GetEmail_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateAudienceContactRequestEmailChannel
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEmail() // Should return zero value
	})

	t.Run("GetMarketingConsent", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestEmailChannel{}
		var expected *CreateAudienceContactRequestEmailChannelMarketingConsent
		obj.MarketingConsent = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMarketingConsent(), "getter should return the property value")
	})

	t.Run("GetMarketingConsent_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestEmailChannel{}
		obj.MarketingConsent = nil

		// Act & Assert
		assert.Nil(t, obj.GetMarketingConsent(), "getter should return nil when property is nil")
	})

	t.Run("GetMarketingConsent_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateAudienceContactRequestEmailChannel
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetMarketingConsent() // Should return zero value
	})

}

func TestSettersMarkExplicitCreateAudienceContactRequestEmailChannel(t *testing.T) {
	t.Run("SetEmail_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestEmailChannel{}
		var fernTestValueEmail *string

		// Act
		obj.SetEmail(fernTestValueEmail)

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

	t.Run("SetMarketingConsent_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestEmailChannel{}
		var fernTestValueMarketingConsent *CreateAudienceContactRequestEmailChannelMarketingConsent

		// Act
		obj.SetMarketingConsent(fernTestValueMarketingConsent)

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

func TestSettersCreateAudienceContactRequestEmailChannelMarketingConsent(t *testing.T) {
	t.Run("SetStatus", func(t *testing.T) {
		obj := &CreateAudienceContactRequestEmailChannelMarketingConsent{}
		var fernTestValueStatus *CreateAudienceContactRequestEmailChannelMarketingConsentStatus
		obj.SetStatus(fernTestValueStatus)
		assert.Equal(t, fernTestValueStatus, obj.Status)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersCreateAudienceContactRequestEmailChannelMarketingConsent(t *testing.T) {
	t.Run("GetStatus", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestEmailChannelMarketingConsent{}
		var expected *CreateAudienceContactRequestEmailChannelMarketingConsentStatus
		obj.Status = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStatus(), "getter should return the property value")
	})

	t.Run("GetStatus_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestEmailChannelMarketingConsent{}
		obj.Status = nil

		// Act & Assert
		assert.Nil(t, obj.GetStatus(), "getter should return nil when property is nil")
	})

	t.Run("GetStatus_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateAudienceContactRequestEmailChannelMarketingConsent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStatus() // Should return zero value
	})

}

func TestSettersMarkExplicitCreateAudienceContactRequestEmailChannelMarketingConsent(t *testing.T) {
	t.Run("SetStatus_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestEmailChannelMarketingConsent{}
		var fernTestValueStatus *CreateAudienceContactRequestEmailChannelMarketingConsentStatus

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

}

func TestGettersCreateAudienceContactRequestMergeFieldsValue(t *testing.T) {
	t.Run("GetCreateAudienceContactRequestMergeFieldsValueAddr1", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestMergeFieldsValue{}
		var expected *CreateAudienceContactRequestMergeFieldsValueAddr1
		obj.CreateAudienceContactRequestMergeFieldsValueAddr1 = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCreateAudienceContactRequestMergeFieldsValueAddr1(), "getter should return the property value")
	})

	t.Run("GetCreateAudienceContactRequestMergeFieldsValueAddr1_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestMergeFieldsValue{}
		obj.CreateAudienceContactRequestMergeFieldsValueAddr1 = nil

		// Act & Assert
		assert.Nil(t, obj.GetCreateAudienceContactRequestMergeFieldsValueAddr1(), "getter should return nil when property is nil")
	})

	t.Run("GetCreateAudienceContactRequestMergeFieldsValueAddr1_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateAudienceContactRequestMergeFieldsValue
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCreateAudienceContactRequestMergeFieldsValueAddr1() // Should return zero value
	})

	t.Run("GetString", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestMergeFieldsValue{}
		var expected string
		obj.String = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetString(), "getter should return the property value")
	})

	t.Run("GetString_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateAudienceContactRequestMergeFieldsValue
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetString() // Should return zero value
	})

	t.Run("GetDouble", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestMergeFieldsValue{}
		var expected float64
		obj.Double = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetDouble(), "getter should return the property value")
	})

	t.Run("GetDouble_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateAudienceContactRequestMergeFieldsValue
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetDouble() // Should return zero value
	})

}

func TestSettersCreateAudienceContactRequestMergeFieldsValueAddr1(t *testing.T) {
	t.Run("SetAddr1", func(t *testing.T) {
		obj := &CreateAudienceContactRequestMergeFieldsValueAddr1{}
		var fernTestValueAddr1 string
		obj.SetAddr1(fernTestValueAddr1)
		assert.Equal(t, fernTestValueAddr1, obj.Addr1)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetAddr2", func(t *testing.T) {
		obj := &CreateAudienceContactRequestMergeFieldsValueAddr1{}
		var fernTestValueAddr2 *string
		obj.SetAddr2(fernTestValueAddr2)
		assert.Equal(t, fernTestValueAddr2, obj.Addr2)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCity", func(t *testing.T) {
		obj := &CreateAudienceContactRequestMergeFieldsValueAddr1{}
		var fernTestValueCity string
		obj.SetCity(fernTestValueCity)
		assert.Equal(t, fernTestValueCity, obj.City)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetState", func(t *testing.T) {
		obj := &CreateAudienceContactRequestMergeFieldsValueAddr1{}
		var fernTestValueState string
		obj.SetState(fernTestValueState)
		assert.Equal(t, fernTestValueState, obj.State)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetZip", func(t *testing.T) {
		obj := &CreateAudienceContactRequestMergeFieldsValueAddr1{}
		var fernTestValueZip string
		obj.SetZip(fernTestValueZip)
		assert.Equal(t, fernTestValueZip, obj.Zip)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCountry", func(t *testing.T) {
		obj := &CreateAudienceContactRequestMergeFieldsValueAddr1{}
		var fernTestValueCountry *string
		obj.SetCountry(fernTestValueCountry)
		assert.Equal(t, fernTestValueCountry, obj.Country)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersCreateAudienceContactRequestMergeFieldsValueAddr1(t *testing.T) {
	t.Run("GetAddr1", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestMergeFieldsValueAddr1{}
		var expected string
		obj.Addr1 = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetAddr1(), "getter should return the property value")
	})

	t.Run("GetAddr1_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateAudienceContactRequestMergeFieldsValueAddr1
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetAddr1() // Should return zero value
	})

	t.Run("GetAddr2", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestMergeFieldsValueAddr1{}
		var expected *string
		obj.Addr2 = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetAddr2(), "getter should return the property value")
	})

	t.Run("GetAddr2_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestMergeFieldsValueAddr1{}
		obj.Addr2 = nil

		// Act & Assert
		assert.Nil(t, obj.GetAddr2(), "getter should return nil when property is nil")
	})

	t.Run("GetAddr2_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateAudienceContactRequestMergeFieldsValueAddr1
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetAddr2() // Should return zero value
	})

	t.Run("GetCity", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestMergeFieldsValueAddr1{}
		var expected string
		obj.City = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCity(), "getter should return the property value")
	})

	t.Run("GetCity_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateAudienceContactRequestMergeFieldsValueAddr1
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCity() // Should return zero value
	})

	t.Run("GetState", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestMergeFieldsValueAddr1{}
		var expected string
		obj.State = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetState(), "getter should return the property value")
	})

	t.Run("GetState_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateAudienceContactRequestMergeFieldsValueAddr1
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetState() // Should return zero value
	})

	t.Run("GetZip", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestMergeFieldsValueAddr1{}
		var expected string
		obj.Zip = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetZip(), "getter should return the property value")
	})

	t.Run("GetZip_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateAudienceContactRequestMergeFieldsValueAddr1
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetZip() // Should return zero value
	})

	t.Run("GetCountry", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestMergeFieldsValueAddr1{}
		var expected *string
		obj.Country = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCountry(), "getter should return the property value")
	})

	t.Run("GetCountry_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestMergeFieldsValueAddr1{}
		obj.Country = nil

		// Act & Assert
		assert.Nil(t, obj.GetCountry(), "getter should return nil when property is nil")
	})

	t.Run("GetCountry_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateAudienceContactRequestMergeFieldsValueAddr1
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCountry() // Should return zero value
	})

}

func TestSettersMarkExplicitCreateAudienceContactRequestMergeFieldsValueAddr1(t *testing.T) {
	t.Run("SetAddr1_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestMergeFieldsValueAddr1{}
		var fernTestValueAddr1 string

		// Act
		obj.SetAddr1(fernTestValueAddr1)

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

	t.Run("SetAddr2_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestMergeFieldsValueAddr1{}
		var fernTestValueAddr2 *string

		// Act
		obj.SetAddr2(fernTestValueAddr2)

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

	t.Run("SetCity_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestMergeFieldsValueAddr1{}
		var fernTestValueCity string

		// Act
		obj.SetCity(fernTestValueCity)

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

	t.Run("SetState_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestMergeFieldsValueAddr1{}
		var fernTestValueState string

		// Act
		obj.SetState(fernTestValueState)

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

	t.Run("SetZip_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestMergeFieldsValueAddr1{}
		var fernTestValueZip string

		// Act
		obj.SetZip(fernTestValueZip)

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

	t.Run("SetCountry_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestMergeFieldsValueAddr1{}
		var fernTestValueCountry *string

		// Act
		obj.SetCountry(fernTestValueCountry)

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

func TestSettersCreateAudienceContactRequestSmsChannel(t *testing.T) {
	t.Run("SetMarketingConsent", func(t *testing.T) {
		obj := &CreateAudienceContactRequestSmsChannel{}
		var fernTestValueMarketingConsent *CreateAudienceContactRequestSmsChannelMarketingConsent
		obj.SetMarketingConsent(fernTestValueMarketingConsent)
		assert.Equal(t, fernTestValueMarketingConsent, obj.MarketingConsent)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSmsPhone", func(t *testing.T) {
		obj := &CreateAudienceContactRequestSmsChannel{}
		var fernTestValueSmsPhone *string
		obj.SetSmsPhone(fernTestValueSmsPhone)
		assert.Equal(t, fernTestValueSmsPhone, obj.SmsPhone)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersCreateAudienceContactRequestSmsChannel(t *testing.T) {
	t.Run("GetMarketingConsent", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestSmsChannel{}
		var expected *CreateAudienceContactRequestSmsChannelMarketingConsent
		obj.MarketingConsent = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMarketingConsent(), "getter should return the property value")
	})

	t.Run("GetMarketingConsent_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestSmsChannel{}
		obj.MarketingConsent = nil

		// Act & Assert
		assert.Nil(t, obj.GetMarketingConsent(), "getter should return nil when property is nil")
	})

	t.Run("GetMarketingConsent_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateAudienceContactRequestSmsChannel
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetMarketingConsent() // Should return zero value
	})

	t.Run("GetSmsPhone", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestSmsChannel{}
		var expected *string
		obj.SmsPhone = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSmsPhone(), "getter should return the property value")
	})

	t.Run("GetSmsPhone_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestSmsChannel{}
		obj.SmsPhone = nil

		// Act & Assert
		assert.Nil(t, obj.GetSmsPhone(), "getter should return nil when property is nil")
	})

	t.Run("GetSmsPhone_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateAudienceContactRequestSmsChannel
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSmsPhone() // Should return zero value
	})

}

func TestSettersMarkExplicitCreateAudienceContactRequestSmsChannel(t *testing.T) {
	t.Run("SetMarketingConsent_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestSmsChannel{}
		var fernTestValueMarketingConsent *CreateAudienceContactRequestSmsChannelMarketingConsent

		// Act
		obj.SetMarketingConsent(fernTestValueMarketingConsent)

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

	t.Run("SetSmsPhone_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestSmsChannel{}
		var fernTestValueSmsPhone *string

		// Act
		obj.SetSmsPhone(fernTestValueSmsPhone)

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

func TestSettersCreateAudienceContactRequestSmsChannelMarketingConsent(t *testing.T) {
	t.Run("SetSource", func(t *testing.T) {
		obj := &CreateAudienceContactRequestSmsChannelMarketingConsent{}
		var fernTestValueSource *CreateAudienceContactRequestSmsChannelMarketingConsentSource
		obj.SetSource(fernTestValueSource)
		assert.Equal(t, fernTestValueSource, obj.Source)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStatus", func(t *testing.T) {
		obj := &CreateAudienceContactRequestSmsChannelMarketingConsent{}
		var fernTestValueStatus *CreateAudienceContactRequestSmsChannelMarketingConsentStatus
		obj.SetStatus(fernTestValueStatus)
		assert.Equal(t, fernTestValueStatus, obj.Status)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCapturedAt", func(t *testing.T) {
		obj := &CreateAudienceContactRequestSmsChannelMarketingConsent{}
		var fernTestValueCapturedAt *time.Time
		obj.SetCapturedAt(fernTestValueCapturedAt)
		assert.Equal(t, fernTestValueCapturedAt, obj.CapturedAt)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersCreateAudienceContactRequestSmsChannelMarketingConsent(t *testing.T) {
	t.Run("GetSource", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestSmsChannelMarketingConsent{}
		var expected *CreateAudienceContactRequestSmsChannelMarketingConsentSource
		obj.Source = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSource(), "getter should return the property value")
	})

	t.Run("GetSource_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestSmsChannelMarketingConsent{}
		obj.Source = nil

		// Act & Assert
		assert.Nil(t, obj.GetSource(), "getter should return nil when property is nil")
	})

	t.Run("GetSource_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateAudienceContactRequestSmsChannelMarketingConsent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSource() // Should return zero value
	})

	t.Run("GetStatus", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestSmsChannelMarketingConsent{}
		var expected *CreateAudienceContactRequestSmsChannelMarketingConsentStatus
		obj.Status = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStatus(), "getter should return the property value")
	})

	t.Run("GetStatus_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestSmsChannelMarketingConsent{}
		obj.Status = nil

		// Act & Assert
		assert.Nil(t, obj.GetStatus(), "getter should return nil when property is nil")
	})

	t.Run("GetStatus_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateAudienceContactRequestSmsChannelMarketingConsent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStatus() // Should return zero value
	})

	t.Run("GetCapturedAt", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestSmsChannelMarketingConsent{}
		var expected *time.Time
		obj.CapturedAt = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCapturedAt(), "getter should return the property value")
	})

	t.Run("GetCapturedAt_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestSmsChannelMarketingConsent{}
		obj.CapturedAt = nil

		// Act & Assert
		assert.Nil(t, obj.GetCapturedAt(), "getter should return nil when property is nil")
	})

	t.Run("GetCapturedAt_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateAudienceContactRequestSmsChannelMarketingConsent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCapturedAt() // Should return zero value
	})

}

func TestSettersMarkExplicitCreateAudienceContactRequestSmsChannelMarketingConsent(t *testing.T) {
	t.Run("SetSource_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestSmsChannelMarketingConsent{}
		var fernTestValueSource *CreateAudienceContactRequestSmsChannelMarketingConsentSource

		// Act
		obj.SetSource(fernTestValueSource)

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
		obj := &CreateAudienceContactRequestSmsChannelMarketingConsent{}
		var fernTestValueStatus *CreateAudienceContactRequestSmsChannelMarketingConsentStatus

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

	t.Run("SetCapturedAt_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestSmsChannelMarketingConsent{}
		var fernTestValueCapturedAt *time.Time

		// Act
		obj.SetCapturedAt(fernTestValueCapturedAt)

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

func TestSettersCreateAudienceContactRequestSmsChannelMarketingConsentSource(t *testing.T) {
	t.Run("SetName", func(t *testing.T) {
		obj := &CreateAudienceContactRequestSmsChannelMarketingConsentSource{}
		var fernTestValueName *string
		obj.SetName(fernTestValueName)
		assert.Equal(t, fernTestValueName, obj.Name)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersCreateAudienceContactRequestSmsChannelMarketingConsentSource(t *testing.T) {
	t.Run("GetName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestSmsChannelMarketingConsentSource{}
		var expected *string
		obj.Name = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetName(), "getter should return the property value")
	})

	t.Run("GetName_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestSmsChannelMarketingConsentSource{}
		obj.Name = nil

		// Act & Assert
		assert.Nil(t, obj.GetName(), "getter should return nil when property is nil")
	})

	t.Run("GetName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateAudienceContactRequestSmsChannelMarketingConsentSource
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetName() // Should return zero value
	})

}

func TestSettersMarkExplicitCreateAudienceContactRequestSmsChannelMarketingConsentSource(t *testing.T) {
	t.Run("SetName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestSmsChannelMarketingConsentSource{}
		var fernTestValueName *string

		// Act
		obj.SetName(fernTestValueName)

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

func TestGettersCreateAudienceContactRequestTagsItem(t *testing.T) {
	t.Run("GetString", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestTagsItem{}
		var expected string
		obj.String = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetString(), "getter should return the property value")
	})

	t.Run("GetString_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateAudienceContactRequestTagsItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetString() // Should return zero value
	})

	t.Run("GetCreateAudienceContactRequestTagsItemName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestTagsItem{}
		var expected *CreateAudienceContactRequestTagsItemName
		obj.CreateAudienceContactRequestTagsItemName = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCreateAudienceContactRequestTagsItemName(), "getter should return the property value")
	})

	t.Run("GetCreateAudienceContactRequestTagsItemName_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestTagsItem{}
		obj.CreateAudienceContactRequestTagsItemName = nil

		// Act & Assert
		assert.Nil(t, obj.GetCreateAudienceContactRequestTagsItemName(), "getter should return nil when property is nil")
	})

	t.Run("GetCreateAudienceContactRequestTagsItemName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateAudienceContactRequestTagsItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCreateAudienceContactRequestTagsItemName() // Should return zero value
	})

}

func TestSettersCreateAudienceContactRequestTagsItemName(t *testing.T) {
	t.Run("SetName", func(t *testing.T) {
		obj := &CreateAudienceContactRequestTagsItemName{}
		var fernTestValueName string
		obj.SetName(fernTestValueName)
		assert.Equal(t, fernTestValueName, obj.Name)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStatus", func(t *testing.T) {
		obj := &CreateAudienceContactRequestTagsItemName{}
		var fernTestValueStatus CreateAudienceContactRequestTagsItemNameStatus
		obj.SetStatus(fernTestValueStatus)
		assert.Equal(t, fernTestValueStatus, obj.Status)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersCreateAudienceContactRequestTagsItemName(t *testing.T) {
	t.Run("GetName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestTagsItemName{}
		var expected string
		obj.Name = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetName(), "getter should return the property value")
	})

	t.Run("GetName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateAudienceContactRequestTagsItemName
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetName() // Should return zero value
	})

	t.Run("GetStatus", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestTagsItemName{}
		var expected CreateAudienceContactRequestTagsItemNameStatus
		obj.Status = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStatus(), "getter should return the property value")
	})

	t.Run("GetStatus_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateAudienceContactRequestTagsItemName
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStatus() // Should return zero value
	})

}

func TestSettersMarkExplicitCreateAudienceContactRequestTagsItemName(t *testing.T) {
	t.Run("SetName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestTagsItemName{}
		var fernTestValueName string

		// Act
		obj.SetName(fernTestValueName)

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
		obj := &CreateAudienceContactRequestTagsItemName{}
		var fernTestValueStatus CreateAudienceContactRequestTagsItemNameStatus

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

}

func TestSettersGetAudienceContactListResponse(t *testing.T) {
	t.Run("SetContacts", func(t *testing.T) {
		obj := &GetAudienceContactListResponse{}
		var fernTestValueContacts []*AudiencesContact
		obj.SetContacts(fernTestValueContacts)
		assert.Equal(t, fernTestValueContacts, obj.Contacts)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetNextCursor", func(t *testing.T) {
		obj := &GetAudienceContactListResponse{}
		var fernTestValueNextCursor *string
		obj.SetNextCursor(fernTestValueNextCursor)
		assert.Equal(t, fernTestValueNextCursor, obj.NextCursor)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetLinks", func(t *testing.T) {
		obj := &GetAudienceContactListResponse{}
		var fernTestValueLinks []*GetAudienceContactListResponseLinksItem
		obj.SetLinks(fernTestValueLinks)
		assert.Equal(t, fernTestValueLinks, obj.Links)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersGetAudienceContactListResponse(t *testing.T) {
	t.Run("GetContacts", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetAudienceContactListResponse{}
		var expected []*AudiencesContact
		obj.Contacts = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetContacts(), "getter should return the property value")
	})

	t.Run("GetContacts_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetAudienceContactListResponse{}
		obj.Contacts = nil

		// Act & Assert
		assert.Nil(t, obj.GetContacts(), "getter should return nil when property is nil")
	})

	t.Run("GetContacts_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetAudienceContactListResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetContacts() // Should return zero value
	})

	t.Run("GetNextCursor", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetAudienceContactListResponse{}
		var expected *string
		obj.NextCursor = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetNextCursor(), "getter should return the property value")
	})

	t.Run("GetNextCursor_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetAudienceContactListResponse{}
		obj.NextCursor = nil

		// Act & Assert
		assert.Nil(t, obj.GetNextCursor(), "getter should return nil when property is nil")
	})

	t.Run("GetNextCursor_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetAudienceContactListResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetNextCursor() // Should return zero value
	})

	t.Run("GetLinks", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetAudienceContactListResponse{}
		var expected []*GetAudienceContactListResponseLinksItem
		obj.Links = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetLinks(), "getter should return the property value")
	})

	t.Run("GetLinks_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetAudienceContactListResponse{}
		obj.Links = nil

		// Act & Assert
		assert.Nil(t, obj.GetLinks(), "getter should return nil when property is nil")
	})

	t.Run("GetLinks_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetAudienceContactListResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetLinks() // Should return zero value
	})

}

func TestSettersMarkExplicitGetAudienceContactListResponse(t *testing.T) {
	t.Run("SetContacts_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetAudienceContactListResponse{}
		var fernTestValueContacts []*AudiencesContact

		// Act
		obj.SetContacts(fernTestValueContacts)

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

	t.Run("SetNextCursor_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetAudienceContactListResponse{}
		var fernTestValueNextCursor *string

		// Act
		obj.SetNextCursor(fernTestValueNextCursor)

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

	t.Run("SetLinks_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetAudienceContactListResponse{}
		var fernTestValueLinks []*GetAudienceContactListResponseLinksItem

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

}

func TestSettersGetAudienceContactListResponseLinksItem(t *testing.T) {
	t.Run("SetHref", func(t *testing.T) {
		obj := &GetAudienceContactListResponseLinksItem{}
		var fernTestValueHref *string
		obj.SetHref(fernTestValueHref)
		assert.Equal(t, fernTestValueHref, obj.Href)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMethod", func(t *testing.T) {
		obj := &GetAudienceContactListResponseLinksItem{}
		var fernTestValueMethod *GetAudienceContactListResponseLinksItemMethod
		obj.SetMethod(fernTestValueMethod)
		assert.Equal(t, fernTestValueMethod, obj.Method)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRel", func(t *testing.T) {
		obj := &GetAudienceContactListResponseLinksItem{}
		var fernTestValueRel *string
		obj.SetRel(fernTestValueRel)
		assert.Equal(t, fernTestValueRel, obj.Rel)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSchema", func(t *testing.T) {
		obj := &GetAudienceContactListResponseLinksItem{}
		var fernTestValueSchema *string
		obj.SetSchema(fernTestValueSchema)
		assert.Equal(t, fernTestValueSchema, obj.Schema)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTargetSchema", func(t *testing.T) {
		obj := &GetAudienceContactListResponseLinksItem{}
		var fernTestValueTargetSchema *string
		obj.SetTargetSchema(fernTestValueTargetSchema)
		assert.Equal(t, fernTestValueTargetSchema, obj.TargetSchema)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersGetAudienceContactListResponseLinksItem(t *testing.T) {
	t.Run("GetHref", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetAudienceContactListResponseLinksItem{}
		var expected *string
		obj.Href = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetHref(), "getter should return the property value")
	})

	t.Run("GetHref_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetAudienceContactListResponseLinksItem{}
		obj.Href = nil

		// Act & Assert
		assert.Nil(t, obj.GetHref(), "getter should return nil when property is nil")
	})

	t.Run("GetHref_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetAudienceContactListResponseLinksItem
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
		obj := &GetAudienceContactListResponseLinksItem{}
		var expected *GetAudienceContactListResponseLinksItemMethod
		obj.Method = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMethod(), "getter should return the property value")
	})

	t.Run("GetMethod_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetAudienceContactListResponseLinksItem{}
		obj.Method = nil

		// Act & Assert
		assert.Nil(t, obj.GetMethod(), "getter should return nil when property is nil")
	})

	t.Run("GetMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetAudienceContactListResponseLinksItem
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
		obj := &GetAudienceContactListResponseLinksItem{}
		var expected *string
		obj.Rel = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetRel(), "getter should return the property value")
	})

	t.Run("GetRel_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetAudienceContactListResponseLinksItem{}
		obj.Rel = nil

		// Act & Assert
		assert.Nil(t, obj.GetRel(), "getter should return nil when property is nil")
	})

	t.Run("GetRel_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetAudienceContactListResponseLinksItem
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
		obj := &GetAudienceContactListResponseLinksItem{}
		var expected *string
		obj.Schema = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSchema(), "getter should return the property value")
	})

	t.Run("GetSchema_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetAudienceContactListResponseLinksItem{}
		obj.Schema = nil

		// Act & Assert
		assert.Nil(t, obj.GetSchema(), "getter should return nil when property is nil")
	})

	t.Run("GetSchema_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetAudienceContactListResponseLinksItem
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
		obj := &GetAudienceContactListResponseLinksItem{}
		var expected *string
		obj.TargetSchema = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTargetSchema(), "getter should return the property value")
	})

	t.Run("GetTargetSchema_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetAudienceContactListResponseLinksItem{}
		obj.TargetSchema = nil

		// Act & Assert
		assert.Nil(t, obj.GetTargetSchema(), "getter should return nil when property is nil")
	})

	t.Run("GetTargetSchema_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetAudienceContactListResponseLinksItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTargetSchema() // Should return zero value
	})

}

func TestSettersMarkExplicitGetAudienceContactListResponseLinksItem(t *testing.T) {
	t.Run("SetHref_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetAudienceContactListResponseLinksItem{}
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
		obj := &GetAudienceContactListResponseLinksItem{}
		var fernTestValueMethod *GetAudienceContactListResponseLinksItemMethod

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
		obj := &GetAudienceContactListResponseLinksItem{}
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
		obj := &GetAudienceContactListResponseLinksItem{}
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
		obj := &GetAudienceContactListResponseLinksItem{}
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

func TestSettersPatchAudienceContactRequestEmailChannel(t *testing.T) {
	t.Run("SetEmail", func(t *testing.T) {
		obj := &PatchAudienceContactRequestEmailChannel{}
		var fernTestValueEmail *string
		obj.SetEmail(fernTestValueEmail)
		assert.Equal(t, fernTestValueEmail, obj.Email)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMarketingConsent", func(t *testing.T) {
		obj := &PatchAudienceContactRequestEmailChannel{}
		var fernTestValueMarketingConsent *PatchAudienceContactRequestEmailChannelMarketingConsent
		obj.SetMarketingConsent(fernTestValueMarketingConsent)
		assert.Equal(t, fernTestValueMarketingConsent, obj.MarketingConsent)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersPatchAudienceContactRequestEmailChannel(t *testing.T) {
	t.Run("GetEmail", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestEmailChannel{}
		var expected *string
		obj.Email = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEmail(), "getter should return the property value")
	})

	t.Run("GetEmail_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestEmailChannel{}
		obj.Email = nil

		// Act & Assert
		assert.Nil(t, obj.GetEmail(), "getter should return nil when property is nil")
	})

	t.Run("GetEmail_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestEmailChannel
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEmail() // Should return zero value
	})

	t.Run("GetMarketingConsent", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestEmailChannel{}
		var expected *PatchAudienceContactRequestEmailChannelMarketingConsent
		obj.MarketingConsent = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMarketingConsent(), "getter should return the property value")
	})

	t.Run("GetMarketingConsent_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestEmailChannel{}
		obj.MarketingConsent = nil

		// Act & Assert
		assert.Nil(t, obj.GetMarketingConsent(), "getter should return nil when property is nil")
	})

	t.Run("GetMarketingConsent_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestEmailChannel
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetMarketingConsent() // Should return zero value
	})

}

func TestSettersMarkExplicitPatchAudienceContactRequestEmailChannel(t *testing.T) {
	t.Run("SetEmail_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestEmailChannel{}
		var fernTestValueEmail *string

		// Act
		obj.SetEmail(fernTestValueEmail)

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

	t.Run("SetMarketingConsent_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestEmailChannel{}
		var fernTestValueMarketingConsent *PatchAudienceContactRequestEmailChannelMarketingConsent

		// Act
		obj.SetMarketingConsent(fernTestValueMarketingConsent)

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

func TestSettersPatchAudienceContactRequestEmailChannelMarketingConsent(t *testing.T) {
	t.Run("SetSource", func(t *testing.T) {
		obj := &PatchAudienceContactRequestEmailChannelMarketingConsent{}
		var fernTestValueSource *PatchAudienceContactRequestEmailChannelMarketingConsentSource
		obj.SetSource(fernTestValueSource)
		assert.Equal(t, fernTestValueSource, obj.Source)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStatus", func(t *testing.T) {
		obj := &PatchAudienceContactRequestEmailChannelMarketingConsent{}
		var fernTestValueStatus *PatchAudienceContactRequestEmailChannelMarketingConsentStatus
		obj.SetStatus(fernTestValueStatus)
		assert.Equal(t, fernTestValueStatus, obj.Status)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCapturedAt", func(t *testing.T) {
		obj := &PatchAudienceContactRequestEmailChannelMarketingConsent{}
		var fernTestValueCapturedAt *time.Time
		obj.SetCapturedAt(fernTestValueCapturedAt)
		assert.Equal(t, fernTestValueCapturedAt, obj.CapturedAt)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersPatchAudienceContactRequestEmailChannelMarketingConsent(t *testing.T) {
	t.Run("GetSource", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestEmailChannelMarketingConsent{}
		var expected *PatchAudienceContactRequestEmailChannelMarketingConsentSource
		obj.Source = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSource(), "getter should return the property value")
	})

	t.Run("GetSource_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestEmailChannelMarketingConsent{}
		obj.Source = nil

		// Act & Assert
		assert.Nil(t, obj.GetSource(), "getter should return nil when property is nil")
	})

	t.Run("GetSource_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestEmailChannelMarketingConsent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSource() // Should return zero value
	})

	t.Run("GetStatus", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestEmailChannelMarketingConsent{}
		var expected *PatchAudienceContactRequestEmailChannelMarketingConsentStatus
		obj.Status = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStatus(), "getter should return the property value")
	})

	t.Run("GetStatus_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestEmailChannelMarketingConsent{}
		obj.Status = nil

		// Act & Assert
		assert.Nil(t, obj.GetStatus(), "getter should return nil when property is nil")
	})

	t.Run("GetStatus_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestEmailChannelMarketingConsent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStatus() // Should return zero value
	})

	t.Run("GetCapturedAt", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestEmailChannelMarketingConsent{}
		var expected *time.Time
		obj.CapturedAt = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCapturedAt(), "getter should return the property value")
	})

	t.Run("GetCapturedAt_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestEmailChannelMarketingConsent{}
		obj.CapturedAt = nil

		// Act & Assert
		assert.Nil(t, obj.GetCapturedAt(), "getter should return nil when property is nil")
	})

	t.Run("GetCapturedAt_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestEmailChannelMarketingConsent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCapturedAt() // Should return zero value
	})

}

func TestSettersMarkExplicitPatchAudienceContactRequestEmailChannelMarketingConsent(t *testing.T) {
	t.Run("SetSource_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestEmailChannelMarketingConsent{}
		var fernTestValueSource *PatchAudienceContactRequestEmailChannelMarketingConsentSource

		// Act
		obj.SetSource(fernTestValueSource)

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
		obj := &PatchAudienceContactRequestEmailChannelMarketingConsent{}
		var fernTestValueStatus *PatchAudienceContactRequestEmailChannelMarketingConsentStatus

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

	t.Run("SetCapturedAt_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestEmailChannelMarketingConsent{}
		var fernTestValueCapturedAt *time.Time

		// Act
		obj.SetCapturedAt(fernTestValueCapturedAt)

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

func TestSettersPatchAudienceContactRequestEmailChannelMarketingConsentSource(t *testing.T) {
	t.Run("SetName", func(t *testing.T) {
		obj := &PatchAudienceContactRequestEmailChannelMarketingConsentSource{}
		var fernTestValueName *string
		obj.SetName(fernTestValueName)
		assert.Equal(t, fernTestValueName, obj.Name)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersPatchAudienceContactRequestEmailChannelMarketingConsentSource(t *testing.T) {
	t.Run("GetName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestEmailChannelMarketingConsentSource{}
		var expected *string
		obj.Name = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetName(), "getter should return the property value")
	})

	t.Run("GetName_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestEmailChannelMarketingConsentSource{}
		obj.Name = nil

		// Act & Assert
		assert.Nil(t, obj.GetName(), "getter should return nil when property is nil")
	})

	t.Run("GetName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestEmailChannelMarketingConsentSource
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetName() // Should return zero value
	})

}

func TestSettersMarkExplicitPatchAudienceContactRequestEmailChannelMarketingConsentSource(t *testing.T) {
	t.Run("SetName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestEmailChannelMarketingConsentSource{}
		var fernTestValueName *string

		// Act
		obj.SetName(fernTestValueName)

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

func TestGettersPatchAudienceContactRequestMergeFieldsValue(t *testing.T) {
	t.Run("GetPatchAudienceContactRequestMergeFieldsValueAddr1", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestMergeFieldsValue{}
		var expected *PatchAudienceContactRequestMergeFieldsValueAddr1
		obj.PatchAudienceContactRequestMergeFieldsValueAddr1 = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetPatchAudienceContactRequestMergeFieldsValueAddr1(), "getter should return the property value")
	})

	t.Run("GetPatchAudienceContactRequestMergeFieldsValueAddr1_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestMergeFieldsValue{}
		obj.PatchAudienceContactRequestMergeFieldsValueAddr1 = nil

		// Act & Assert
		assert.Nil(t, obj.GetPatchAudienceContactRequestMergeFieldsValueAddr1(), "getter should return nil when property is nil")
	})

	t.Run("GetPatchAudienceContactRequestMergeFieldsValueAddr1_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestMergeFieldsValue
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetPatchAudienceContactRequestMergeFieldsValueAddr1() // Should return zero value
	})

	t.Run("GetString", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestMergeFieldsValue{}
		var expected string
		obj.String = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetString(), "getter should return the property value")
	})

	t.Run("GetString_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestMergeFieldsValue
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetString() // Should return zero value
	})

	t.Run("GetDouble", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestMergeFieldsValue{}
		var expected float64
		obj.Double = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetDouble(), "getter should return the property value")
	})

	t.Run("GetDouble_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestMergeFieldsValue
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetDouble() // Should return zero value
	})

}

func TestSettersPatchAudienceContactRequestMergeFieldsValueAddr1(t *testing.T) {
	t.Run("SetAddr1", func(t *testing.T) {
		obj := &PatchAudienceContactRequestMergeFieldsValueAddr1{}
		var fernTestValueAddr1 string
		obj.SetAddr1(fernTestValueAddr1)
		assert.Equal(t, fernTestValueAddr1, obj.Addr1)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetAddr2", func(t *testing.T) {
		obj := &PatchAudienceContactRequestMergeFieldsValueAddr1{}
		var fernTestValueAddr2 *string
		obj.SetAddr2(fernTestValueAddr2)
		assert.Equal(t, fernTestValueAddr2, obj.Addr2)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCity", func(t *testing.T) {
		obj := &PatchAudienceContactRequestMergeFieldsValueAddr1{}
		var fernTestValueCity string
		obj.SetCity(fernTestValueCity)
		assert.Equal(t, fernTestValueCity, obj.City)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetState", func(t *testing.T) {
		obj := &PatchAudienceContactRequestMergeFieldsValueAddr1{}
		var fernTestValueState string
		obj.SetState(fernTestValueState)
		assert.Equal(t, fernTestValueState, obj.State)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetZip", func(t *testing.T) {
		obj := &PatchAudienceContactRequestMergeFieldsValueAddr1{}
		var fernTestValueZip string
		obj.SetZip(fernTestValueZip)
		assert.Equal(t, fernTestValueZip, obj.Zip)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCountry", func(t *testing.T) {
		obj := &PatchAudienceContactRequestMergeFieldsValueAddr1{}
		var fernTestValueCountry *string
		obj.SetCountry(fernTestValueCountry)
		assert.Equal(t, fernTestValueCountry, obj.Country)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersPatchAudienceContactRequestMergeFieldsValueAddr1(t *testing.T) {
	t.Run("GetAddr1", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestMergeFieldsValueAddr1{}
		var expected string
		obj.Addr1 = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetAddr1(), "getter should return the property value")
	})

	t.Run("GetAddr1_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestMergeFieldsValueAddr1
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetAddr1() // Should return zero value
	})

	t.Run("GetAddr2", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestMergeFieldsValueAddr1{}
		var expected *string
		obj.Addr2 = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetAddr2(), "getter should return the property value")
	})

	t.Run("GetAddr2_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestMergeFieldsValueAddr1{}
		obj.Addr2 = nil

		// Act & Assert
		assert.Nil(t, obj.GetAddr2(), "getter should return nil when property is nil")
	})

	t.Run("GetAddr2_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestMergeFieldsValueAddr1
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetAddr2() // Should return zero value
	})

	t.Run("GetCity", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestMergeFieldsValueAddr1{}
		var expected string
		obj.City = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCity(), "getter should return the property value")
	})

	t.Run("GetCity_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestMergeFieldsValueAddr1
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCity() // Should return zero value
	})

	t.Run("GetState", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestMergeFieldsValueAddr1{}
		var expected string
		obj.State = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetState(), "getter should return the property value")
	})

	t.Run("GetState_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestMergeFieldsValueAddr1
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetState() // Should return zero value
	})

	t.Run("GetZip", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestMergeFieldsValueAddr1{}
		var expected string
		obj.Zip = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetZip(), "getter should return the property value")
	})

	t.Run("GetZip_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestMergeFieldsValueAddr1
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetZip() // Should return zero value
	})

	t.Run("GetCountry", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestMergeFieldsValueAddr1{}
		var expected *string
		obj.Country = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCountry(), "getter should return the property value")
	})

	t.Run("GetCountry_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestMergeFieldsValueAddr1{}
		obj.Country = nil

		// Act & Assert
		assert.Nil(t, obj.GetCountry(), "getter should return nil when property is nil")
	})

	t.Run("GetCountry_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestMergeFieldsValueAddr1
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCountry() // Should return zero value
	})

}

func TestSettersMarkExplicitPatchAudienceContactRequestMergeFieldsValueAddr1(t *testing.T) {
	t.Run("SetAddr1_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestMergeFieldsValueAddr1{}
		var fernTestValueAddr1 string

		// Act
		obj.SetAddr1(fernTestValueAddr1)

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

	t.Run("SetAddr2_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestMergeFieldsValueAddr1{}
		var fernTestValueAddr2 *string

		// Act
		obj.SetAddr2(fernTestValueAddr2)

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

	t.Run("SetCity_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestMergeFieldsValueAddr1{}
		var fernTestValueCity string

		// Act
		obj.SetCity(fernTestValueCity)

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

	t.Run("SetState_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestMergeFieldsValueAddr1{}
		var fernTestValueState string

		// Act
		obj.SetState(fernTestValueState)

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

	t.Run("SetZip_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestMergeFieldsValueAddr1{}
		var fernTestValueZip string

		// Act
		obj.SetZip(fernTestValueZip)

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

	t.Run("SetCountry_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestMergeFieldsValueAddr1{}
		var fernTestValueCountry *string

		// Act
		obj.SetCountry(fernTestValueCountry)

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

func TestSettersPatchAudienceContactRequestSmsChannel(t *testing.T) {
	t.Run("SetMarketingConsent", func(t *testing.T) {
		obj := &PatchAudienceContactRequestSmsChannel{}
		var fernTestValueMarketingConsent *PatchAudienceContactRequestSmsChannelMarketingConsent
		obj.SetMarketingConsent(fernTestValueMarketingConsent)
		assert.Equal(t, fernTestValueMarketingConsent, obj.MarketingConsent)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSmsPhone", func(t *testing.T) {
		obj := &PatchAudienceContactRequestSmsChannel{}
		var fernTestValueSmsPhone *string
		obj.SetSmsPhone(fernTestValueSmsPhone)
		assert.Equal(t, fernTestValueSmsPhone, obj.SmsPhone)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersPatchAudienceContactRequestSmsChannel(t *testing.T) {
	t.Run("GetMarketingConsent", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestSmsChannel{}
		var expected *PatchAudienceContactRequestSmsChannelMarketingConsent
		obj.MarketingConsent = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMarketingConsent(), "getter should return the property value")
	})

	t.Run("GetMarketingConsent_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestSmsChannel{}
		obj.MarketingConsent = nil

		// Act & Assert
		assert.Nil(t, obj.GetMarketingConsent(), "getter should return nil when property is nil")
	})

	t.Run("GetMarketingConsent_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestSmsChannel
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetMarketingConsent() // Should return zero value
	})

	t.Run("GetSmsPhone", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestSmsChannel{}
		var expected *string
		obj.SmsPhone = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSmsPhone(), "getter should return the property value")
	})

	t.Run("GetSmsPhone_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestSmsChannel{}
		obj.SmsPhone = nil

		// Act & Assert
		assert.Nil(t, obj.GetSmsPhone(), "getter should return nil when property is nil")
	})

	t.Run("GetSmsPhone_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestSmsChannel
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSmsPhone() // Should return zero value
	})

}

func TestSettersMarkExplicitPatchAudienceContactRequestSmsChannel(t *testing.T) {
	t.Run("SetMarketingConsent_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestSmsChannel{}
		var fernTestValueMarketingConsent *PatchAudienceContactRequestSmsChannelMarketingConsent

		// Act
		obj.SetMarketingConsent(fernTestValueMarketingConsent)

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

	t.Run("SetSmsPhone_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestSmsChannel{}
		var fernTestValueSmsPhone *string

		// Act
		obj.SetSmsPhone(fernTestValueSmsPhone)

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

func TestSettersPatchAudienceContactRequestSmsChannelMarketingConsent(t *testing.T) {
	t.Run("SetSource", func(t *testing.T) {
		obj := &PatchAudienceContactRequestSmsChannelMarketingConsent{}
		var fernTestValueSource *PatchAudienceContactRequestSmsChannelMarketingConsentSource
		obj.SetSource(fernTestValueSource)
		assert.Equal(t, fernTestValueSource, obj.Source)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStatus", func(t *testing.T) {
		obj := &PatchAudienceContactRequestSmsChannelMarketingConsent{}
		var fernTestValueStatus *PatchAudienceContactRequestSmsChannelMarketingConsentStatus
		obj.SetStatus(fernTestValueStatus)
		assert.Equal(t, fernTestValueStatus, obj.Status)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCapturedAt", func(t *testing.T) {
		obj := &PatchAudienceContactRequestSmsChannelMarketingConsent{}
		var fernTestValueCapturedAt *time.Time
		obj.SetCapturedAt(fernTestValueCapturedAt)
		assert.Equal(t, fernTestValueCapturedAt, obj.CapturedAt)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersPatchAudienceContactRequestSmsChannelMarketingConsent(t *testing.T) {
	t.Run("GetSource", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestSmsChannelMarketingConsent{}
		var expected *PatchAudienceContactRequestSmsChannelMarketingConsentSource
		obj.Source = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSource(), "getter should return the property value")
	})

	t.Run("GetSource_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestSmsChannelMarketingConsent{}
		obj.Source = nil

		// Act & Assert
		assert.Nil(t, obj.GetSource(), "getter should return nil when property is nil")
	})

	t.Run("GetSource_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestSmsChannelMarketingConsent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSource() // Should return zero value
	})

	t.Run("GetStatus", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestSmsChannelMarketingConsent{}
		var expected *PatchAudienceContactRequestSmsChannelMarketingConsentStatus
		obj.Status = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStatus(), "getter should return the property value")
	})

	t.Run("GetStatus_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestSmsChannelMarketingConsent{}
		obj.Status = nil

		// Act & Assert
		assert.Nil(t, obj.GetStatus(), "getter should return nil when property is nil")
	})

	t.Run("GetStatus_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestSmsChannelMarketingConsent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStatus() // Should return zero value
	})

	t.Run("GetCapturedAt", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestSmsChannelMarketingConsent{}
		var expected *time.Time
		obj.CapturedAt = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCapturedAt(), "getter should return the property value")
	})

	t.Run("GetCapturedAt_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestSmsChannelMarketingConsent{}
		obj.CapturedAt = nil

		// Act & Assert
		assert.Nil(t, obj.GetCapturedAt(), "getter should return nil when property is nil")
	})

	t.Run("GetCapturedAt_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestSmsChannelMarketingConsent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCapturedAt() // Should return zero value
	})

}

func TestSettersMarkExplicitPatchAudienceContactRequestSmsChannelMarketingConsent(t *testing.T) {
	t.Run("SetSource_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestSmsChannelMarketingConsent{}
		var fernTestValueSource *PatchAudienceContactRequestSmsChannelMarketingConsentSource

		// Act
		obj.SetSource(fernTestValueSource)

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
		obj := &PatchAudienceContactRequestSmsChannelMarketingConsent{}
		var fernTestValueStatus *PatchAudienceContactRequestSmsChannelMarketingConsentStatus

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

	t.Run("SetCapturedAt_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestSmsChannelMarketingConsent{}
		var fernTestValueCapturedAt *time.Time

		// Act
		obj.SetCapturedAt(fernTestValueCapturedAt)

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

func TestSettersPatchAudienceContactRequestSmsChannelMarketingConsentSource(t *testing.T) {
	t.Run("SetName", func(t *testing.T) {
		obj := &PatchAudienceContactRequestSmsChannelMarketingConsentSource{}
		var fernTestValueName *string
		obj.SetName(fernTestValueName)
		assert.Equal(t, fernTestValueName, obj.Name)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersPatchAudienceContactRequestSmsChannelMarketingConsentSource(t *testing.T) {
	t.Run("GetName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestSmsChannelMarketingConsentSource{}
		var expected *string
		obj.Name = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetName(), "getter should return the property value")
	})

	t.Run("GetName_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestSmsChannelMarketingConsentSource{}
		obj.Name = nil

		// Act & Assert
		assert.Nil(t, obj.GetName(), "getter should return nil when property is nil")
	})

	t.Run("GetName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestSmsChannelMarketingConsentSource
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetName() // Should return zero value
	})

}

func TestSettersMarkExplicitPatchAudienceContactRequestSmsChannelMarketingConsentSource(t *testing.T) {
	t.Run("SetName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestSmsChannelMarketingConsentSource{}
		var fernTestValueName *string

		// Act
		obj.SetName(fernTestValueName)

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

func TestGettersPatchAudienceContactRequestTagsItem(t *testing.T) {
	t.Run("GetString", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestTagsItem{}
		var expected string
		obj.String = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetString(), "getter should return the property value")
	})

	t.Run("GetString_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestTagsItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetString() // Should return zero value
	})

	t.Run("GetPatchAudienceContactRequestTagsItemName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestTagsItem{}
		var expected *PatchAudienceContactRequestTagsItemName
		obj.PatchAudienceContactRequestTagsItemName = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetPatchAudienceContactRequestTagsItemName(), "getter should return the property value")
	})

	t.Run("GetPatchAudienceContactRequestTagsItemName_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestTagsItem{}
		obj.PatchAudienceContactRequestTagsItemName = nil

		// Act & Assert
		assert.Nil(t, obj.GetPatchAudienceContactRequestTagsItemName(), "getter should return nil when property is nil")
	})

	t.Run("GetPatchAudienceContactRequestTagsItemName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestTagsItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetPatchAudienceContactRequestTagsItemName() // Should return zero value
	})

}

func TestSettersPatchAudienceContactRequestTagsItemName(t *testing.T) {
	t.Run("SetName", func(t *testing.T) {
		obj := &PatchAudienceContactRequestTagsItemName{}
		var fernTestValueName string
		obj.SetName(fernTestValueName)
		assert.Equal(t, fernTestValueName, obj.Name)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStatus", func(t *testing.T) {
		obj := &PatchAudienceContactRequestTagsItemName{}
		var fernTestValueStatus PatchAudienceContactRequestTagsItemNameStatus
		obj.SetStatus(fernTestValueStatus)
		assert.Equal(t, fernTestValueStatus, obj.Status)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersPatchAudienceContactRequestTagsItemName(t *testing.T) {
	t.Run("GetName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestTagsItemName{}
		var expected string
		obj.Name = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetName(), "getter should return the property value")
	})

	t.Run("GetName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestTagsItemName
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetName() // Should return zero value
	})

	t.Run("GetStatus", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestTagsItemName{}
		var expected PatchAudienceContactRequestTagsItemNameStatus
		obj.Status = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStatus(), "getter should return the property value")
	})

	t.Run("GetStatus_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestTagsItemName
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStatus() // Should return zero value
	})

}

func TestSettersMarkExplicitPatchAudienceContactRequestTagsItemName(t *testing.T) {
	t.Run("SetName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestTagsItemName{}
		var fernTestValueName string

		// Act
		obj.SetName(fernTestValueName)

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
		obj := &PatchAudienceContactRequestTagsItemName{}
		var fernTestValueStatus PatchAudienceContactRequestTagsItemNameStatus

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

}

func TestJSONMarshalingAudiencesContact(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContact{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled AudiencesContact
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj AudiencesContact
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj AudiencesContact
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingAudiencesContactEmailChannel(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannel{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled AudiencesContactEmailChannel
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj AudiencesContactEmailChannel
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj AudiencesContactEmailChannel
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingAudiencesContactEmailChannelEffectiveSubscriptionStatus(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannelEffectiveSubscriptionStatus{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled AudiencesContactEmailChannelEffectiveSubscriptionStatus
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj AudiencesContactEmailChannelEffectiveSubscriptionStatus
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj AudiencesContactEmailChannelEffectiveSubscriptionStatus
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingAudiencesContactEmailChannelMarketingConsent(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannelMarketingConsent{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled AudiencesContactEmailChannelMarketingConsent
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj AudiencesContactEmailChannelMarketingConsent
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj AudiencesContactEmailChannelMarketingConsent
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingAudiencesContactEmailChannelMarketingConsentSource(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannelMarketingConsentSource{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled AudiencesContactEmailChannelMarketingConsentSource
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj AudiencesContactEmailChannelMarketingConsentSource
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj AudiencesContactEmailChannelMarketingConsentSource
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingAudiencesContactEmailChannelSource(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactEmailChannelSource{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled AudiencesContactEmailChannelSource
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj AudiencesContactEmailChannelSource
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj AudiencesContactEmailChannelSource
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingAudiencesContactMergeFieldsValueAddr1(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactMergeFieldsValueAddr1{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled AudiencesContactMergeFieldsValueAddr1
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj AudiencesContactMergeFieldsValueAddr1
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj AudiencesContactMergeFieldsValueAddr1
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingAudiencesContactSmsChannel(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannel{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled AudiencesContactSmsChannel
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj AudiencesContactSmsChannel
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj AudiencesContactSmsChannel
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingAudiencesContactSmsChannelEffectiveSubscriptionStatus(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannelEffectiveSubscriptionStatus{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled AudiencesContactSmsChannelEffectiveSubscriptionStatus
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj AudiencesContactSmsChannelEffectiveSubscriptionStatus
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj AudiencesContactSmsChannelEffectiveSubscriptionStatus
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingAudiencesContactSmsChannelMarketingConsent(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannelMarketingConsent{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled AudiencesContactSmsChannelMarketingConsent
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj AudiencesContactSmsChannelMarketingConsent
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj AudiencesContactSmsChannelMarketingConsent
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingAudiencesContactSmsChannelMarketingConsentSource(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannelMarketingConsentSource{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled AudiencesContactSmsChannelMarketingConsentSource
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj AudiencesContactSmsChannelMarketingConsentSource
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj AudiencesContactSmsChannelMarketingConsentSource
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingAudiencesContactSmsChannelSource(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSmsChannelSource{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled AudiencesContactSmsChannelSource
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj AudiencesContactSmsChannelSource
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj AudiencesContactSmsChannelSource
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingAudiencesContactSource(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &AudiencesContactSource{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled AudiencesContactSource
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj AudiencesContactSource
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj AudiencesContactSource
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingCreateAudienceContactRequestEmailChannel(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestEmailChannel{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled CreateAudienceContactRequestEmailChannel
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj CreateAudienceContactRequestEmailChannel
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj CreateAudienceContactRequestEmailChannel
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingCreateAudienceContactRequestEmailChannelMarketingConsent(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestEmailChannelMarketingConsent{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled CreateAudienceContactRequestEmailChannelMarketingConsent
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj CreateAudienceContactRequestEmailChannelMarketingConsent
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj CreateAudienceContactRequestEmailChannelMarketingConsent
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingCreateAudienceContactRequestMergeFieldsValueAddr1(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestMergeFieldsValueAddr1{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled CreateAudienceContactRequestMergeFieldsValueAddr1
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj CreateAudienceContactRequestMergeFieldsValueAddr1
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj CreateAudienceContactRequestMergeFieldsValueAddr1
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingCreateAudienceContactRequestSmsChannel(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestSmsChannel{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled CreateAudienceContactRequestSmsChannel
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj CreateAudienceContactRequestSmsChannel
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj CreateAudienceContactRequestSmsChannel
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingCreateAudienceContactRequestSmsChannelMarketingConsent(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestSmsChannelMarketingConsent{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled CreateAudienceContactRequestSmsChannelMarketingConsent
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj CreateAudienceContactRequestSmsChannelMarketingConsent
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj CreateAudienceContactRequestSmsChannelMarketingConsent
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingCreateAudienceContactRequestSmsChannelMarketingConsentSource(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestSmsChannelMarketingConsentSource{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled CreateAudienceContactRequestSmsChannelMarketingConsentSource
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj CreateAudienceContactRequestSmsChannelMarketingConsentSource
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj CreateAudienceContactRequestSmsChannelMarketingConsentSource
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingCreateAudienceContactRequestTagsItemName(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateAudienceContactRequestTagsItemName{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled CreateAudienceContactRequestTagsItemName
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj CreateAudienceContactRequestTagsItemName
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj CreateAudienceContactRequestTagsItemName
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingGetAudienceContactListResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetAudienceContactListResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled GetAudienceContactListResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj GetAudienceContactListResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj GetAudienceContactListResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingGetAudienceContactListResponseLinksItem(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetAudienceContactListResponseLinksItem{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled GetAudienceContactListResponseLinksItem
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj GetAudienceContactListResponseLinksItem
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj GetAudienceContactListResponseLinksItem
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingPatchAudienceContactRequestEmailChannel(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestEmailChannel{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled PatchAudienceContactRequestEmailChannel
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj PatchAudienceContactRequestEmailChannel
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj PatchAudienceContactRequestEmailChannel
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingPatchAudienceContactRequestEmailChannelMarketingConsent(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestEmailChannelMarketingConsent{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled PatchAudienceContactRequestEmailChannelMarketingConsent
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj PatchAudienceContactRequestEmailChannelMarketingConsent
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj PatchAudienceContactRequestEmailChannelMarketingConsent
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingPatchAudienceContactRequestEmailChannelMarketingConsentSource(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestEmailChannelMarketingConsentSource{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled PatchAudienceContactRequestEmailChannelMarketingConsentSource
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj PatchAudienceContactRequestEmailChannelMarketingConsentSource
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj PatchAudienceContactRequestEmailChannelMarketingConsentSource
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingPatchAudienceContactRequestMergeFieldsValueAddr1(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestMergeFieldsValueAddr1{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled PatchAudienceContactRequestMergeFieldsValueAddr1
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj PatchAudienceContactRequestMergeFieldsValueAddr1
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj PatchAudienceContactRequestMergeFieldsValueAddr1
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingPatchAudienceContactRequestSmsChannel(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestSmsChannel{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled PatchAudienceContactRequestSmsChannel
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj PatchAudienceContactRequestSmsChannel
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj PatchAudienceContactRequestSmsChannel
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingPatchAudienceContactRequestSmsChannelMarketingConsent(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestSmsChannelMarketingConsent{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled PatchAudienceContactRequestSmsChannelMarketingConsent
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj PatchAudienceContactRequestSmsChannelMarketingConsent
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj PatchAudienceContactRequestSmsChannelMarketingConsent
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingPatchAudienceContactRequestSmsChannelMarketingConsentSource(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestSmsChannelMarketingConsentSource{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled PatchAudienceContactRequestSmsChannelMarketingConsentSource
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj PatchAudienceContactRequestSmsChannelMarketingConsentSource
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj PatchAudienceContactRequestSmsChannelMarketingConsentSource
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingPatchAudienceContactRequestTagsItemName(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &PatchAudienceContactRequestTagsItemName{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled PatchAudienceContactRequestTagsItemName
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj PatchAudienceContactRequestTagsItemName
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj PatchAudienceContactRequestTagsItemName
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestStringAudiencesContact(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &AudiencesContact{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContact
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringAudiencesContactEmailChannel(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &AudiencesContactEmailChannel{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactEmailChannel
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringAudiencesContactEmailChannelEffectiveSubscriptionStatus(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &AudiencesContactEmailChannelEffectiveSubscriptionStatus{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactEmailChannelEffectiveSubscriptionStatus
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringAudiencesContactEmailChannelMarketingConsent(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &AudiencesContactEmailChannelMarketingConsent{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactEmailChannelMarketingConsent
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringAudiencesContactEmailChannelMarketingConsentSource(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &AudiencesContactEmailChannelMarketingConsentSource{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactEmailChannelMarketingConsentSource
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringAudiencesContactEmailChannelSource(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &AudiencesContactEmailChannelSource{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactEmailChannelSource
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringAudiencesContactMergeFieldsValueAddr1(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &AudiencesContactMergeFieldsValueAddr1{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactMergeFieldsValueAddr1
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringAudiencesContactSmsChannel(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &AudiencesContactSmsChannel{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactSmsChannel
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringAudiencesContactSmsChannelEffectiveSubscriptionStatus(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &AudiencesContactSmsChannelEffectiveSubscriptionStatus{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactSmsChannelEffectiveSubscriptionStatus
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringAudiencesContactSmsChannelMarketingConsent(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &AudiencesContactSmsChannelMarketingConsent{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactSmsChannelMarketingConsent
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringAudiencesContactSmsChannelMarketingConsentSource(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &AudiencesContactSmsChannelMarketingConsentSource{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactSmsChannelMarketingConsentSource
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringAudiencesContactSmsChannelSource(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &AudiencesContactSmsChannelSource{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactSmsChannelSource
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringAudiencesContactSource(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &AudiencesContactSource{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *AudiencesContactSource
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringCreateAudienceContactRequestEmailChannel(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &CreateAudienceContactRequestEmailChannel{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateAudienceContactRequestEmailChannel
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringCreateAudienceContactRequestEmailChannelMarketingConsent(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &CreateAudienceContactRequestEmailChannelMarketingConsent{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateAudienceContactRequestEmailChannelMarketingConsent
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringCreateAudienceContactRequestMergeFieldsValueAddr1(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &CreateAudienceContactRequestMergeFieldsValueAddr1{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateAudienceContactRequestMergeFieldsValueAddr1
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringCreateAudienceContactRequestSmsChannel(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &CreateAudienceContactRequestSmsChannel{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateAudienceContactRequestSmsChannel
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringCreateAudienceContactRequestSmsChannelMarketingConsent(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &CreateAudienceContactRequestSmsChannelMarketingConsent{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateAudienceContactRequestSmsChannelMarketingConsent
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringCreateAudienceContactRequestSmsChannelMarketingConsentSource(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &CreateAudienceContactRequestSmsChannelMarketingConsentSource{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateAudienceContactRequestSmsChannelMarketingConsentSource
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringCreateAudienceContactRequestTagsItemName(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &CreateAudienceContactRequestTagsItemName{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateAudienceContactRequestTagsItemName
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringGetAudienceContactListResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &GetAudienceContactListResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetAudienceContactListResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringGetAudienceContactListResponseLinksItem(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &GetAudienceContactListResponseLinksItem{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *GetAudienceContactListResponseLinksItem
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringPatchAudienceContactRequestEmailChannel(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &PatchAudienceContactRequestEmailChannel{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestEmailChannel
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringPatchAudienceContactRequestEmailChannelMarketingConsent(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &PatchAudienceContactRequestEmailChannelMarketingConsent{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestEmailChannelMarketingConsent
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringPatchAudienceContactRequestEmailChannelMarketingConsentSource(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &PatchAudienceContactRequestEmailChannelMarketingConsentSource{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestEmailChannelMarketingConsentSource
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringPatchAudienceContactRequestMergeFieldsValueAddr1(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &PatchAudienceContactRequestMergeFieldsValueAddr1{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestMergeFieldsValueAddr1
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringPatchAudienceContactRequestSmsChannel(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &PatchAudienceContactRequestSmsChannel{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestSmsChannel
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringPatchAudienceContactRequestSmsChannelMarketingConsent(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &PatchAudienceContactRequestSmsChannelMarketingConsent{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestSmsChannelMarketingConsent
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringPatchAudienceContactRequestSmsChannelMarketingConsentSource(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &PatchAudienceContactRequestSmsChannelMarketingConsentSource{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestSmsChannelMarketingConsentSource
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringPatchAudienceContactRequestTagsItemName(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &PatchAudienceContactRequestTagsItemName{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *PatchAudienceContactRequestTagsItemName
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestEnumAudiencesContactEmailChannelEffectiveSubscriptionStatusValue(t *testing.T) {
	t.Run("NewFromString_subscribed", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactEmailChannelEffectiveSubscriptionStatusValueFromString("subscribed")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactEmailChannelEffectiveSubscriptionStatusValue("subscribed"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_unsubscribed", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactEmailChannelEffectiveSubscriptionStatusValueFromString("unsubscribed")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactEmailChannelEffectiveSubscriptionStatusValue("unsubscribed"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_nonsubscribed", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactEmailChannelEffectiveSubscriptionStatusValueFromString("nonsubscribed")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactEmailChannelEffectiveSubscriptionStatusValue("nonsubscribed"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_pending", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactEmailChannelEffectiveSubscriptionStatusValueFromString("pending")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactEmailChannelEffectiveSubscriptionStatusValue("pending"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewAudiencesContactEmailChannelEffectiveSubscriptionStatusValueFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewAudiencesContactEmailChannelEffectiveSubscriptionStatusValueFromString("subscribed")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumAudiencesContactEmailChannelMarketingConsentStatus(t *testing.T) {
	t.Run("NewFromString_consented", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactEmailChannelMarketingConsentStatusFromString("consented")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactEmailChannelMarketingConsentStatus("consented"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_denied", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactEmailChannelMarketingConsentStatusFromString("denied")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactEmailChannelMarketingConsentStatus("denied"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_confirmed", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactEmailChannelMarketingConsentStatusFromString("confirmed")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactEmailChannelMarketingConsentStatus("confirmed"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_unknown", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactEmailChannelMarketingConsentStatusFromString("unknown")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactEmailChannelMarketingConsentStatus("unknown"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewAudiencesContactEmailChannelMarketingConsentStatusFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewAudiencesContactEmailChannelMarketingConsentStatusFromString("consented")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumAudiencesContactLanguage(t *testing.T) {
	t.Run("NewFromString_empty_string", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage(""), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_en", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("en")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("en"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_ar", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("ar")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("ar"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_af", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("af")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("af"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_be", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("be")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("be"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_bg", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("bg")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("bg"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_ca", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("ca")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("ca"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_zh", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("zh")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("zh"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_zh_CN", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("zh_CN")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("zh_CN"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_hr", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("hr")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("hr"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_cs", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("cs")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("cs"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_da", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("da")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("da"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_nl", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("nl")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("nl"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_et", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("et")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("et"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_fa", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("fa")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("fa"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_fi", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("fi")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("fi"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_fr", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("fr")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("fr"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_fr_CA", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("fr_CA")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("fr_CA"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_de", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("de")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("de"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_el", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("el")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("el"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_he", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("he")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("he"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_hi", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("hi")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("hi"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_hu", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("hu")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("hu"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_is", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("is")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("is"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_id", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("id")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("id"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_ga", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("ga")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("ga"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_it", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("it")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("it"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_ja", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("ja")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("ja"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_km", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("km")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("km"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_ko", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("ko")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("ko"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_lv", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("lv")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("lv"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_lt", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("lt")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("lt"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_mt", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("mt")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("mt"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_ms", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("ms")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("ms"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_mk", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("mk")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("mk"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_no", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("no")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("no"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_pl", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("pl")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("pl"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_pt", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("pt")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("pt"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_pt_PT", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("pt_PT")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("pt_PT"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_ro", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("ro")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("ro"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_ru", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("ru")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("ru"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_sr", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("sr")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("sr"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_sk", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("sk")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("sk"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_sl", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("sl")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("sl"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_es", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("es")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("es"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_es_ES", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("es_ES")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("es_ES"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_sw", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("sw")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("sw"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_sv", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("sv")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("sv"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_ta", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("ta")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("ta"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_th", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("th")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("th"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_tr", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("tr")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("tr"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_uk", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("uk")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("uk"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_vi", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactLanguageFromString("vi")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactLanguage("vi"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewAudiencesContactLanguageFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewAudiencesContactLanguageFromString("")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumAudiencesContactSmsChannelEffectiveSubscriptionStatusValue(t *testing.T) {
	t.Run("NewFromString_subscribed", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactSmsChannelEffectiveSubscriptionStatusValueFromString("subscribed")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactSmsChannelEffectiveSubscriptionStatusValue("subscribed"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_unsubscribed", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactSmsChannelEffectiveSubscriptionStatusValueFromString("unsubscribed")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactSmsChannelEffectiveSubscriptionStatusValue("unsubscribed"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_nonsubscribed", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactSmsChannelEffectiveSubscriptionStatusValueFromString("nonsubscribed")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactSmsChannelEffectiveSubscriptionStatusValue("nonsubscribed"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_pending", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactSmsChannelEffectiveSubscriptionStatusValueFromString("pending")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactSmsChannelEffectiveSubscriptionStatusValue("pending"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewAudiencesContactSmsChannelEffectiveSubscriptionStatusValueFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewAudiencesContactSmsChannelEffectiveSubscriptionStatusValueFromString("subscribed")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumAudiencesContactSmsChannelMarketingConsentStatus(t *testing.T) {
	t.Run("NewFromString_consented", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactSmsChannelMarketingConsentStatusFromString("consented")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactSmsChannelMarketingConsentStatus("consented"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_confirmed", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactSmsChannelMarketingConsentStatusFromString("confirmed")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactSmsChannelMarketingConsentStatus("confirmed"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_denied", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactSmsChannelMarketingConsentStatusFromString("denied")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactSmsChannelMarketingConsentStatus("denied"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_unknown", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactSmsChannelMarketingConsentStatusFromString("unknown")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactSmsChannelMarketingConsentStatus("unknown"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewAudiencesContactSmsChannelMarketingConsentStatusFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewAudiencesContactSmsChannelMarketingConsentStatusFromString("consented")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumAudiencesContactStatus(t *testing.T) {
	t.Run("NewFromString_active", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactStatusFromString("active")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactStatus("active"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_archived", func(t *testing.T) {
		t.Parallel()
		val, err := NewAudiencesContactStatusFromString("archived")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, AudiencesContactStatus("archived"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewAudiencesContactStatusFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewAudiencesContactStatusFromString("active")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumCreateAudienceContactRequestDataMode(t *testing.T) {
	t.Run("NewFromString_historical", func(t *testing.T) {
		t.Parallel()
		val, err := NewCreateAudienceContactRequestDataModeFromString("historical")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, CreateAudienceContactRequestDataMode("historical"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_live", func(t *testing.T) {
		t.Parallel()
		val, err := NewCreateAudienceContactRequestDataModeFromString("live")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, CreateAudienceContactRequestDataMode("live"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewCreateAudienceContactRequestDataModeFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewCreateAudienceContactRequestDataModeFromString("historical")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumCreateAudienceContactRequestEmailChannelMarketingConsentStatus(t *testing.T) {
	t.Run("NewFromString_confirmed", func(t *testing.T) {
		t.Parallel()
		val, err := NewCreateAudienceContactRequestEmailChannelMarketingConsentStatusFromString("confirmed")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, CreateAudienceContactRequestEmailChannelMarketingConsentStatus("confirmed"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_consented", func(t *testing.T) {
		t.Parallel()
		val, err := NewCreateAudienceContactRequestEmailChannelMarketingConsentStatusFromString("consented")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, CreateAudienceContactRequestEmailChannelMarketingConsentStatus("consented"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_denied", func(t *testing.T) {
		t.Parallel()
		val, err := NewCreateAudienceContactRequestEmailChannelMarketingConsentStatusFromString("denied")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, CreateAudienceContactRequestEmailChannelMarketingConsentStatus("denied"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_unknown", func(t *testing.T) {
		t.Parallel()
		val, err := NewCreateAudienceContactRequestEmailChannelMarketingConsentStatusFromString("unknown")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, CreateAudienceContactRequestEmailChannelMarketingConsentStatus("unknown"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewCreateAudienceContactRequestEmailChannelMarketingConsentStatusFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewCreateAudienceContactRequestEmailChannelMarketingConsentStatusFromString("confirmed")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumCreateAudienceContactRequestMergeFieldValidationMode(t *testing.T) {
	t.Run("NewFromString_ignore_required_checks", func(t *testing.T) {
		t.Parallel()
		val, err := NewCreateAudienceContactRequestMergeFieldValidationModeFromString("ignore_required_checks")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, CreateAudienceContactRequestMergeFieldValidationMode("ignore_required_checks"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_strict", func(t *testing.T) {
		t.Parallel()
		val, err := NewCreateAudienceContactRequestMergeFieldValidationModeFromString("strict")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, CreateAudienceContactRequestMergeFieldValidationMode("strict"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewCreateAudienceContactRequestMergeFieldValidationModeFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewCreateAudienceContactRequestMergeFieldValidationModeFromString("ignore_required_checks")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumCreateAudienceContactRequestSmsChannelMarketingConsentStatus(t *testing.T) {
	t.Run("NewFromString_consented", func(t *testing.T) {
		t.Parallel()
		val, err := NewCreateAudienceContactRequestSmsChannelMarketingConsentStatusFromString("consented")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, CreateAudienceContactRequestSmsChannelMarketingConsentStatus("consented"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_confirmed", func(t *testing.T) {
		t.Parallel()
		val, err := NewCreateAudienceContactRequestSmsChannelMarketingConsentStatusFromString("confirmed")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, CreateAudienceContactRequestSmsChannelMarketingConsentStatus("confirmed"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_unknown", func(t *testing.T) {
		t.Parallel()
		val, err := NewCreateAudienceContactRequestSmsChannelMarketingConsentStatusFromString("unknown")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, CreateAudienceContactRequestSmsChannelMarketingConsentStatus("unknown"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewCreateAudienceContactRequestSmsChannelMarketingConsentStatusFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewCreateAudienceContactRequestSmsChannelMarketingConsentStatusFromString("consented")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumCreateAudienceContactRequestTagsItemNameStatus(t *testing.T) {
	t.Run("NewFromString_active", func(t *testing.T) {
		t.Parallel()
		val, err := NewCreateAudienceContactRequestTagsItemNameStatusFromString("active")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, CreateAudienceContactRequestTagsItemNameStatus("active"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_inactive", func(t *testing.T) {
		t.Parallel()
		val, err := NewCreateAudienceContactRequestTagsItemNameStatusFromString("inactive")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, CreateAudienceContactRequestTagsItemNameStatus("inactive"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewCreateAudienceContactRequestTagsItemNameStatusFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewCreateAudienceContactRequestTagsItemNameStatusFromString("active")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumGetAudienceContactListRequestSortDir(t *testing.T) {
	t.Run("NewFromString_ASC", func(t *testing.T) {
		t.Parallel()
		val, err := NewGetAudienceContactListRequestSortDirFromString("ASC")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, GetAudienceContactListRequestSortDir("ASC"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_DESC", func(t *testing.T) {
		t.Parallel()
		val, err := NewGetAudienceContactListRequestSortDirFromString("DESC")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, GetAudienceContactListRequestSortDir("DESC"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewGetAudienceContactListRequestSortDirFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewGetAudienceContactListRequestSortDirFromString("ASC")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumGetAudienceContactListRequestSortField(t *testing.T) {
	t.Run("NewFromString_created_at", func(t *testing.T) {
		t.Parallel()
		val, err := NewGetAudienceContactListRequestSortFieldFromString("created_at")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, GetAudienceContactListRequestSortField("created_at"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_updated_at", func(t *testing.T) {
		t.Parallel()
		val, err := NewGetAudienceContactListRequestSortFieldFromString("updated_at")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, GetAudienceContactListRequestSortField("updated_at"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewGetAudienceContactListRequestSortFieldFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewGetAudienceContactListRequestSortFieldFromString("created_at")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumGetAudienceContactListResponseLinksItemMethod(t *testing.T) {
	t.Run("NewFromString_GET", func(t *testing.T) {
		t.Parallel()
		val, err := NewGetAudienceContactListResponseLinksItemMethodFromString("GET")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, GetAudienceContactListResponseLinksItemMethod("GET"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_POST", func(t *testing.T) {
		t.Parallel()
		val, err := NewGetAudienceContactListResponseLinksItemMethodFromString("POST")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, GetAudienceContactListResponseLinksItemMethod("POST"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_PUT", func(t *testing.T) {
		t.Parallel()
		val, err := NewGetAudienceContactListResponseLinksItemMethodFromString("PUT")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, GetAudienceContactListResponseLinksItemMethod("PUT"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_PATCH", func(t *testing.T) {
		t.Parallel()
		val, err := NewGetAudienceContactListResponseLinksItemMethodFromString("PATCH")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, GetAudienceContactListResponseLinksItemMethod("PATCH"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_DELETE", func(t *testing.T) {
		t.Parallel()
		val, err := NewGetAudienceContactListResponseLinksItemMethodFromString("DELETE")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, GetAudienceContactListResponseLinksItemMethod("DELETE"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_OPTIONS", func(t *testing.T) {
		t.Parallel()
		val, err := NewGetAudienceContactListResponseLinksItemMethodFromString("OPTIONS")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, GetAudienceContactListResponseLinksItemMethod("OPTIONS"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_HEAD", func(t *testing.T) {
		t.Parallel()
		val, err := NewGetAudienceContactListResponseLinksItemMethodFromString("HEAD")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, GetAudienceContactListResponseLinksItemMethod("HEAD"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewGetAudienceContactListResponseLinksItemMethodFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewGetAudienceContactListResponseLinksItemMethodFromString("GET")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumPatchAudienceContactRequestDataMode(t *testing.T) {
	t.Run("NewFromString_historical", func(t *testing.T) {
		t.Parallel()
		val, err := NewPatchAudienceContactRequestDataModeFromString("historical")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, PatchAudienceContactRequestDataMode("historical"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_live", func(t *testing.T) {
		t.Parallel()
		val, err := NewPatchAudienceContactRequestDataModeFromString("live")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, PatchAudienceContactRequestDataMode("live"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewPatchAudienceContactRequestDataModeFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewPatchAudienceContactRequestDataModeFromString("historical")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumPatchAudienceContactRequestEmailChannelMarketingConsentStatus(t *testing.T) {
	t.Run("NewFromString_consented", func(t *testing.T) {
		t.Parallel()
		val, err := NewPatchAudienceContactRequestEmailChannelMarketingConsentStatusFromString("consented")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, PatchAudienceContactRequestEmailChannelMarketingConsentStatus("consented"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_denied", func(t *testing.T) {
		t.Parallel()
		val, err := NewPatchAudienceContactRequestEmailChannelMarketingConsentStatusFromString("denied")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, PatchAudienceContactRequestEmailChannelMarketingConsentStatus("denied"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_confirmed", func(t *testing.T) {
		t.Parallel()
		val, err := NewPatchAudienceContactRequestEmailChannelMarketingConsentStatusFromString("confirmed")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, PatchAudienceContactRequestEmailChannelMarketingConsentStatus("confirmed"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_unknown", func(t *testing.T) {
		t.Parallel()
		val, err := NewPatchAudienceContactRequestEmailChannelMarketingConsentStatusFromString("unknown")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, PatchAudienceContactRequestEmailChannelMarketingConsentStatus("unknown"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewPatchAudienceContactRequestEmailChannelMarketingConsentStatusFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewPatchAudienceContactRequestEmailChannelMarketingConsentStatusFromString("consented")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumPatchAudienceContactRequestMergeFieldValidationMode(t *testing.T) {
	t.Run("NewFromString_ignore_required_checks", func(t *testing.T) {
		t.Parallel()
		val, err := NewPatchAudienceContactRequestMergeFieldValidationModeFromString("ignore_required_checks")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, PatchAudienceContactRequestMergeFieldValidationMode("ignore_required_checks"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_strict", func(t *testing.T) {
		t.Parallel()
		val, err := NewPatchAudienceContactRequestMergeFieldValidationModeFromString("strict")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, PatchAudienceContactRequestMergeFieldValidationMode("strict"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewPatchAudienceContactRequestMergeFieldValidationModeFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewPatchAudienceContactRequestMergeFieldValidationModeFromString("ignore_required_checks")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumPatchAudienceContactRequestSmsChannelMarketingConsentStatus(t *testing.T) {
	t.Run("NewFromString_consented", func(t *testing.T) {
		t.Parallel()
		val, err := NewPatchAudienceContactRequestSmsChannelMarketingConsentStatusFromString("consented")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, PatchAudienceContactRequestSmsChannelMarketingConsentStatus("consented"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_confirmed", func(t *testing.T) {
		t.Parallel()
		val, err := NewPatchAudienceContactRequestSmsChannelMarketingConsentStatusFromString("confirmed")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, PatchAudienceContactRequestSmsChannelMarketingConsentStatus("confirmed"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_denied", func(t *testing.T) {
		t.Parallel()
		val, err := NewPatchAudienceContactRequestSmsChannelMarketingConsentStatusFromString("denied")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, PatchAudienceContactRequestSmsChannelMarketingConsentStatus("denied"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_unknown", func(t *testing.T) {
		t.Parallel()
		val, err := NewPatchAudienceContactRequestSmsChannelMarketingConsentStatusFromString("unknown")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, PatchAudienceContactRequestSmsChannelMarketingConsentStatus("unknown"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewPatchAudienceContactRequestSmsChannelMarketingConsentStatusFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewPatchAudienceContactRequestSmsChannelMarketingConsentStatusFromString("consented")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumPatchAudienceContactRequestTagsItemNameStatus(t *testing.T) {
	t.Run("NewFromString_active", func(t *testing.T) {
		t.Parallel()
		val, err := NewPatchAudienceContactRequestTagsItemNameStatusFromString("active")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, PatchAudienceContactRequestTagsItemNameStatus("active"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_inactive", func(t *testing.T) {
		t.Parallel()
		val, err := NewPatchAudienceContactRequestTagsItemNameStatusFromString("inactive")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, PatchAudienceContactRequestTagsItemNameStatus("inactive"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewPatchAudienceContactRequestTagsItemNameStatusFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewPatchAudienceContactRequestTagsItemNameStatusFromString("active")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestExtraPropertiesAudiencesContact(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &AudiencesContact{}
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
		var obj *AudiencesContact
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesAudiencesContactEmailChannel(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &AudiencesContactEmailChannel{}
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
		var obj *AudiencesContactEmailChannel
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesAudiencesContactEmailChannelEffectiveSubscriptionStatus(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &AudiencesContactEmailChannelEffectiveSubscriptionStatus{}
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
		var obj *AudiencesContactEmailChannelEffectiveSubscriptionStatus
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesAudiencesContactEmailChannelMarketingConsent(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &AudiencesContactEmailChannelMarketingConsent{}
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
		var obj *AudiencesContactEmailChannelMarketingConsent
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesAudiencesContactEmailChannelMarketingConsentSource(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &AudiencesContactEmailChannelMarketingConsentSource{}
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
		var obj *AudiencesContactEmailChannelMarketingConsentSource
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesAudiencesContactEmailChannelSource(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &AudiencesContactEmailChannelSource{}
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
		var obj *AudiencesContactEmailChannelSource
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesAudiencesContactMergeFieldsValueAddr1(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &AudiencesContactMergeFieldsValueAddr1{}
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
		var obj *AudiencesContactMergeFieldsValueAddr1
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesAudiencesContactSmsChannel(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &AudiencesContactSmsChannel{}
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
		var obj *AudiencesContactSmsChannel
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesAudiencesContactSmsChannelEffectiveSubscriptionStatus(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &AudiencesContactSmsChannelEffectiveSubscriptionStatus{}
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
		var obj *AudiencesContactSmsChannelEffectiveSubscriptionStatus
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesAudiencesContactSmsChannelMarketingConsent(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &AudiencesContactSmsChannelMarketingConsent{}
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
		var obj *AudiencesContactSmsChannelMarketingConsent
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesAudiencesContactSmsChannelMarketingConsentSource(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &AudiencesContactSmsChannelMarketingConsentSource{}
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
		var obj *AudiencesContactSmsChannelMarketingConsentSource
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesAudiencesContactSmsChannelSource(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &AudiencesContactSmsChannelSource{}
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
		var obj *AudiencesContactSmsChannelSource
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesAudiencesContactSource(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &AudiencesContactSource{}
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
		var obj *AudiencesContactSource
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesCreateAudienceContactRequestEmailChannel(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &CreateAudienceContactRequestEmailChannel{}
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
		var obj *CreateAudienceContactRequestEmailChannel
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesCreateAudienceContactRequestEmailChannelMarketingConsent(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &CreateAudienceContactRequestEmailChannelMarketingConsent{}
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
		var obj *CreateAudienceContactRequestEmailChannelMarketingConsent
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesCreateAudienceContactRequestMergeFieldsValueAddr1(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &CreateAudienceContactRequestMergeFieldsValueAddr1{}
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
		var obj *CreateAudienceContactRequestMergeFieldsValueAddr1
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesCreateAudienceContactRequestSmsChannel(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &CreateAudienceContactRequestSmsChannel{}
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
		var obj *CreateAudienceContactRequestSmsChannel
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesCreateAudienceContactRequestSmsChannelMarketingConsent(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &CreateAudienceContactRequestSmsChannelMarketingConsent{}
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
		var obj *CreateAudienceContactRequestSmsChannelMarketingConsent
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesCreateAudienceContactRequestSmsChannelMarketingConsentSource(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &CreateAudienceContactRequestSmsChannelMarketingConsentSource{}
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
		var obj *CreateAudienceContactRequestSmsChannelMarketingConsentSource
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesCreateAudienceContactRequestTagsItemName(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &CreateAudienceContactRequestTagsItemName{}
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
		var obj *CreateAudienceContactRequestTagsItemName
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesGetAudienceContactListResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &GetAudienceContactListResponse{}
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
		var obj *GetAudienceContactListResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesGetAudienceContactListResponseLinksItem(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &GetAudienceContactListResponseLinksItem{}
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
		var obj *GetAudienceContactListResponseLinksItem
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesPatchAudienceContactRequestEmailChannel(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &PatchAudienceContactRequestEmailChannel{}
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
		var obj *PatchAudienceContactRequestEmailChannel
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesPatchAudienceContactRequestEmailChannelMarketingConsent(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &PatchAudienceContactRequestEmailChannelMarketingConsent{}
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
		var obj *PatchAudienceContactRequestEmailChannelMarketingConsent
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesPatchAudienceContactRequestEmailChannelMarketingConsentSource(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &PatchAudienceContactRequestEmailChannelMarketingConsentSource{}
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
		var obj *PatchAudienceContactRequestEmailChannelMarketingConsentSource
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesPatchAudienceContactRequestMergeFieldsValueAddr1(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &PatchAudienceContactRequestMergeFieldsValueAddr1{}
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
		var obj *PatchAudienceContactRequestMergeFieldsValueAddr1
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesPatchAudienceContactRequestSmsChannel(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &PatchAudienceContactRequestSmsChannel{}
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
		var obj *PatchAudienceContactRequestSmsChannel
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesPatchAudienceContactRequestSmsChannelMarketingConsent(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &PatchAudienceContactRequestSmsChannelMarketingConsent{}
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
		var obj *PatchAudienceContactRequestSmsChannelMarketingConsent
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesPatchAudienceContactRequestSmsChannelMarketingConsentSource(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &PatchAudienceContactRequestSmsChannelMarketingConsentSource{}
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
		var obj *PatchAudienceContactRequestSmsChannelMarketingConsentSource
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesPatchAudienceContactRequestTagsItemName(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &PatchAudienceContactRequestTagsItemName{}
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
		var obj *PatchAudienceContactRequestTagsItemName
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}
