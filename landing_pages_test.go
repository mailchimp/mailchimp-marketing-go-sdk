// Code generated from our API definition. DO NOT EDIT.

package api

import (
	json "encoding/json"
	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
	testing "testing"
	time "time"
)

func TestSettersCreateLandingPagesRequest(t *testing.T) {
	t.Run("SetUseDefaultList", func(t *testing.T) {
		obj := &CreateLandingPagesRequest{}
		var fernTestValueUseDefaultList *bool
		obj.SetUseDefaultList(fernTestValueUseDefaultList)
		assert.Equal(t, fernTestValueUseDefaultList, obj.UseDefaultList)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetDescription", func(t *testing.T) {
		obj := &CreateLandingPagesRequest{}
		var fernTestValueDescription *string
		obj.SetDescription(fernTestValueDescription)
		assert.Equal(t, fernTestValueDescription, obj.Description)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetListID", func(t *testing.T) {
		obj := &CreateLandingPagesRequest{}
		var fernTestValueListID *string
		obj.SetListID(fernTestValueListID)
		assert.Equal(t, fernTestValueListID, obj.ListID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetName", func(t *testing.T) {
		obj := &CreateLandingPagesRequest{}
		var fernTestValueName *string
		obj.SetName(fernTestValueName)
		assert.Equal(t, fernTestValueName, obj.Name)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStoreID", func(t *testing.T) {
		obj := &CreateLandingPagesRequest{}
		var fernTestValueStoreID *string
		obj.SetStoreID(fernTestValueStoreID)
		assert.Equal(t, fernTestValueStoreID, obj.StoreID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTemplateID", func(t *testing.T) {
		obj := &CreateLandingPagesRequest{}
		var fernTestValueTemplateID *int
		obj.SetTemplateID(fernTestValueTemplateID)
		assert.Equal(t, fernTestValueTemplateID, obj.TemplateID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTitle", func(t *testing.T) {
		obj := &CreateLandingPagesRequest{}
		var fernTestValueTitle *string
		obj.SetTitle(fernTestValueTitle)
		assert.Equal(t, fernTestValueTitle, obj.Title)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTracking", func(t *testing.T) {
		obj := &CreateLandingPagesRequest{}
		var fernTestValueTracking *CreateLandingPagesRequestTracking
		obj.SetTracking(fernTestValueTracking)
		assert.Equal(t, fernTestValueTracking, obj.Tracking)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetType", func(t *testing.T) {
		obj := &CreateLandingPagesRequest{}
		var fernTestValueType *CreateLandingPagesRequestType
		obj.SetType(fernTestValueType)
		assert.Equal(t, fernTestValueType, obj.Type)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitCreateLandingPagesRequest(t *testing.T) {
	t.Run("SetUseDefaultList_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateLandingPagesRequest{}
		var fernTestValueUseDefaultList *bool

		// Act
		obj.SetUseDefaultList(fernTestValueUseDefaultList)

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

	t.Run("SetDescription_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateLandingPagesRequest{}
		var fernTestValueDescription *string

		// Act
		obj.SetDescription(fernTestValueDescription)

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

	t.Run("SetListID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateLandingPagesRequest{}
		var fernTestValueListID *string

		// Act
		obj.SetListID(fernTestValueListID)

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

	t.Run("SetName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateLandingPagesRequest{}
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

	t.Run("SetStoreID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateLandingPagesRequest{}
		var fernTestValueStoreID *string

		// Act
		obj.SetStoreID(fernTestValueStoreID)

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

	t.Run("SetTemplateID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateLandingPagesRequest{}
		var fernTestValueTemplateID *int

		// Act
		obj.SetTemplateID(fernTestValueTemplateID)

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

	t.Run("SetTitle_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateLandingPagesRequest{}
		var fernTestValueTitle *string

		// Act
		obj.SetTitle(fernTestValueTitle)

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

	t.Run("SetTracking_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateLandingPagesRequest{}
		var fernTestValueTracking *CreateLandingPagesRequestTracking

		// Act
		obj.SetTracking(fernTestValueTracking)

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

	t.Run("SetType_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateLandingPagesRequest{}
		var fernTestValueType *CreateLandingPagesRequestType

		// Act
		obj.SetType(fernTestValueType)

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

func TestSettersCreateActionPublishLandingPagesRequest(t *testing.T) {
	t.Run("SetPageID", func(t *testing.T) {
		obj := &CreateActionPublishLandingPagesRequest{}
		var fernTestValuePageID string
		obj.SetPageID(fernTestValuePageID)
		assert.Equal(t, fernTestValuePageID, obj.PageID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitCreateActionPublishLandingPagesRequest(t *testing.T) {
	t.Run("SetPageID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateActionPublishLandingPagesRequest{}
		var fernTestValuePageID string

		// Act
		obj.SetPageID(fernTestValuePageID)

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

func TestSettersCreateActionUnpublishLandingPagesRequest(t *testing.T) {
	t.Run("SetPageID", func(t *testing.T) {
		obj := &CreateActionUnpublishLandingPagesRequest{}
		var fernTestValuePageID string
		obj.SetPageID(fernTestValuePageID)
		assert.Equal(t, fernTestValuePageID, obj.PageID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitCreateActionUnpublishLandingPagesRequest(t *testing.T) {
	t.Run("SetPageID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateActionUnpublishLandingPagesRequest{}
		var fernTestValuePageID string

		// Act
		obj.SetPageID(fernTestValuePageID)

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

func TestSettersDeleteLandingPagesRequest(t *testing.T) {
	t.Run("SetPageID", func(t *testing.T) {
		obj := &DeleteLandingPagesRequest{}
		var fernTestValuePageID string
		obj.SetPageID(fernTestValuePageID)
		assert.Equal(t, fernTestValuePageID, obj.PageID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitDeleteLandingPagesRequest(t *testing.T) {
	t.Run("SetPageID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &DeleteLandingPagesRequest{}
		var fernTestValuePageID string

		// Act
		obj.SetPageID(fernTestValuePageID)

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

func TestSettersGetLandingPagesRequest(t *testing.T) {
	t.Run("SetPageID", func(t *testing.T) {
		obj := &GetLandingPagesRequest{}
		var fernTestValuePageID string
		obj.SetPageID(fernTestValuePageID)
		assert.Equal(t, fernTestValuePageID, obj.PageID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetFields", func(t *testing.T) {
		obj := &GetLandingPagesRequest{}
		var fernTestValueFields []*string
		obj.SetFields(fernTestValueFields)
		assert.Equal(t, fernTestValueFields, obj.Fields)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetExcludeFields", func(t *testing.T) {
		obj := &GetLandingPagesRequest{}
		var fernTestValueExcludeFields []*string
		obj.SetExcludeFields(fernTestValueExcludeFields)
		assert.Equal(t, fernTestValueExcludeFields, obj.ExcludeFields)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitGetLandingPagesRequest(t *testing.T) {
	t.Run("SetPageID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetLandingPagesRequest{}
		var fernTestValuePageID string

		// Act
		obj.SetPageID(fernTestValuePageID)

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
		obj := &GetLandingPagesRequest{}
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
		obj := &GetLandingPagesRequest{}
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

func TestSettersListLandingPagesRequest(t *testing.T) {
	t.Run("SetSortDir", func(t *testing.T) {
		obj := &ListLandingPagesRequest{}
		var fernTestValueSortDir *ListLandingPagesRequestSortDir
		obj.SetSortDir(fernTestValueSortDir)
		assert.Equal(t, fernTestValueSortDir, obj.SortDir)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSortField", func(t *testing.T) {
		obj := &ListLandingPagesRequest{}
		var fernTestValueSortField *ListLandingPagesRequestSortField
		obj.SetSortField(fernTestValueSortField)
		assert.Equal(t, fernTestValueSortField, obj.SortField)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetFields", func(t *testing.T) {
		obj := &ListLandingPagesRequest{}
		var fernTestValueFields []*string
		obj.SetFields(fernTestValueFields)
		assert.Equal(t, fernTestValueFields, obj.Fields)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetExcludeFields", func(t *testing.T) {
		obj := &ListLandingPagesRequest{}
		var fernTestValueExcludeFields []*string
		obj.SetExcludeFields(fernTestValueExcludeFields)
		assert.Equal(t, fernTestValueExcludeFields, obj.ExcludeFields)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCount", func(t *testing.T) {
		obj := &ListLandingPagesRequest{}
		var fernTestValueCount *int
		obj.SetCount(fernTestValueCount)
		assert.Equal(t, fernTestValueCount, obj.Count)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitListLandingPagesRequest(t *testing.T) {
	t.Run("SetSortDir_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListLandingPagesRequest{}
		var fernTestValueSortDir *ListLandingPagesRequestSortDir

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

	t.Run("SetSortField_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListLandingPagesRequest{}
		var fernTestValueSortField *ListLandingPagesRequestSortField

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

	t.Run("SetFields_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListLandingPagesRequest{}
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
		obj := &ListLandingPagesRequest{}
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
		obj := &ListLandingPagesRequest{}
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

}

func TestSettersListContentLandingPagesRequest(t *testing.T) {
	t.Run("SetPageID", func(t *testing.T) {
		obj := &ListContentLandingPagesRequest{}
		var fernTestValuePageID string
		obj.SetPageID(fernTestValuePageID)
		assert.Equal(t, fernTestValuePageID, obj.PageID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetFields", func(t *testing.T) {
		obj := &ListContentLandingPagesRequest{}
		var fernTestValueFields []*string
		obj.SetFields(fernTestValueFields)
		assert.Equal(t, fernTestValueFields, obj.Fields)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetExcludeFields", func(t *testing.T) {
		obj := &ListContentLandingPagesRequest{}
		var fernTestValueExcludeFields []*string
		obj.SetExcludeFields(fernTestValueExcludeFields)
		assert.Equal(t, fernTestValueExcludeFields, obj.ExcludeFields)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitListContentLandingPagesRequest(t *testing.T) {
	t.Run("SetPageID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListContentLandingPagesRequest{}
		var fernTestValuePageID string

		// Act
		obj.SetPageID(fernTestValuePageID)

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
		obj := &ListContentLandingPagesRequest{}
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
		obj := &ListContentLandingPagesRequest{}
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

func TestSettersLandingPage(t *testing.T) {
	t.Run("SetLinks", func(t *testing.T) {
		obj := &LandingPage{}
		var fernTestValueLinks []*LandingPageLinksItem
		obj.SetLinks(fernTestValueLinks)
		assert.Equal(t, fernTestValueLinks, obj.Links)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCreatedAt", func(t *testing.T) {
		obj := &LandingPage{}
		var fernTestValueCreatedAt *time.Time
		obj.SetCreatedAt(fernTestValueCreatedAt)
		assert.Equal(t, fernTestValueCreatedAt, obj.CreatedAt)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCreatedBySource", func(t *testing.T) {
		obj := &LandingPage{}
		var fernTestValueCreatedBySource *string
		obj.SetCreatedBySource(fernTestValueCreatedBySource)
		assert.Equal(t, fernTestValueCreatedBySource, obj.CreatedBySource)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetDescription", func(t *testing.T) {
		obj := &LandingPage{}
		var fernTestValueDescription *string
		obj.SetDescription(fernTestValueDescription)
		assert.Equal(t, fernTestValueDescription, obj.Description)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetID", func(t *testing.T) {
		obj := &LandingPage{}
		var fernTestValueID *string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetListID", func(t *testing.T) {
		obj := &LandingPage{}
		var fernTestValueListID *string
		obj.SetListID(fernTestValueListID)
		assert.Equal(t, fernTestValueListID, obj.ListID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetName", func(t *testing.T) {
		obj := &LandingPage{}
		var fernTestValueName *string
		obj.SetName(fernTestValueName)
		assert.Equal(t, fernTestValueName, obj.Name)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetPublishedAt", func(t *testing.T) {
		obj := &LandingPage{}
		var fernTestValuePublishedAt *time.Time
		obj.SetPublishedAt(fernTestValuePublishedAt)
		assert.Equal(t, fernTestValuePublishedAt, obj.PublishedAt)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStatus", func(t *testing.T) {
		obj := &LandingPage{}
		var fernTestValueStatus *LandingPageStatus
		obj.SetStatus(fernTestValueStatus)
		assert.Equal(t, fernTestValueStatus, obj.Status)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStoreID", func(t *testing.T) {
		obj := &LandingPage{}
		var fernTestValueStoreID *string
		obj.SetStoreID(fernTestValueStoreID)
		assert.Equal(t, fernTestValueStoreID, obj.StoreID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTemplateID", func(t *testing.T) {
		obj := &LandingPage{}
		var fernTestValueTemplateID *int
		obj.SetTemplateID(fernTestValueTemplateID)
		assert.Equal(t, fernTestValueTemplateID, obj.TemplateID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTitle", func(t *testing.T) {
		obj := &LandingPage{}
		var fernTestValueTitle *string
		obj.SetTitle(fernTestValueTitle)
		assert.Equal(t, fernTestValueTitle, obj.Title)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTracking", func(t *testing.T) {
		obj := &LandingPage{}
		var fernTestValueTracking *LandingPageTracking
		obj.SetTracking(fernTestValueTracking)
		assert.Equal(t, fernTestValueTracking, obj.Tracking)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetUnpublishedAt", func(t *testing.T) {
		obj := &LandingPage{}
		var fernTestValueUnpublishedAt *time.Time
		obj.SetUnpublishedAt(fernTestValueUnpublishedAt)
		assert.Equal(t, fernTestValueUnpublishedAt, obj.UnpublishedAt)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetUpdatedAt", func(t *testing.T) {
		obj := &LandingPage{}
		var fernTestValueUpdatedAt *time.Time
		obj.SetUpdatedAt(fernTestValueUpdatedAt)
		assert.Equal(t, fernTestValueUpdatedAt, obj.UpdatedAt)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetURL", func(t *testing.T) {
		obj := &LandingPage{}
		var fernTestValueURL *string
		obj.SetURL(fernTestValueURL)
		assert.Equal(t, fernTestValueURL, obj.URL)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetWebID", func(t *testing.T) {
		obj := &LandingPage{}
		var fernTestValueWebID *int
		obj.SetWebID(fernTestValueWebID)
		assert.Equal(t, fernTestValueWebID, obj.WebID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersLandingPage(t *testing.T) {
	t.Run("GetLinks", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		var expected []*LandingPageLinksItem
		obj.Links = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetLinks(), "getter should return the property value")
	})

	t.Run("GetLinks_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		obj.Links = nil

		// Act & Assert
		assert.Nil(t, obj.GetLinks(), "getter should return nil when property is nil")
	})

	t.Run("GetLinks_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *LandingPage
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetLinks() // Should return zero value
	})

	t.Run("GetCreatedAt", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		var expected *time.Time
		obj.CreatedAt = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCreatedAt(), "getter should return the property value")
	})

	t.Run("GetCreatedAt_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		obj.CreatedAt = nil

		// Act & Assert
		assert.Nil(t, obj.GetCreatedAt(), "getter should return nil when property is nil")
	})

	t.Run("GetCreatedAt_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *LandingPage
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCreatedAt() // Should return zero value
	})

	t.Run("GetCreatedBySource", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		var expected *string
		obj.CreatedBySource = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCreatedBySource(), "getter should return the property value")
	})

	t.Run("GetCreatedBySource_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		obj.CreatedBySource = nil

		// Act & Assert
		assert.Nil(t, obj.GetCreatedBySource(), "getter should return nil when property is nil")
	})

	t.Run("GetCreatedBySource_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *LandingPage
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCreatedBySource() // Should return zero value
	})

	t.Run("GetDescription", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		var expected *string
		obj.Description = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetDescription(), "getter should return the property value")
	})

	t.Run("GetDescription_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		obj.Description = nil

		// Act & Assert
		assert.Nil(t, obj.GetDescription(), "getter should return nil when property is nil")
	})

	t.Run("GetDescription_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *LandingPage
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetDescription() // Should return zero value
	})

	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		var expected *string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		obj.ID = nil

		// Act & Assert
		assert.Nil(t, obj.GetID(), "getter should return nil when property is nil")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *LandingPage
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetID() // Should return zero value
	})

	t.Run("GetListID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		var expected *string
		obj.ListID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetListID(), "getter should return the property value")
	})

	t.Run("GetListID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		obj.ListID = nil

		// Act & Assert
		assert.Nil(t, obj.GetListID(), "getter should return nil when property is nil")
	})

	t.Run("GetListID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *LandingPage
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetListID() // Should return zero value
	})

	t.Run("GetName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		var expected *string
		obj.Name = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetName(), "getter should return the property value")
	})

	t.Run("GetName_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		obj.Name = nil

		// Act & Assert
		assert.Nil(t, obj.GetName(), "getter should return nil when property is nil")
	})

	t.Run("GetName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *LandingPage
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetName() // Should return zero value
	})

	t.Run("GetPublishedAt", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		var expected *time.Time
		obj.PublishedAt = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetPublishedAt(), "getter should return the property value")
	})

	t.Run("GetPublishedAt_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		obj.PublishedAt = nil

		// Act & Assert
		assert.Nil(t, obj.GetPublishedAt(), "getter should return nil when property is nil")
	})

	t.Run("GetPublishedAt_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *LandingPage
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetPublishedAt() // Should return zero value
	})

	t.Run("GetStatus", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		var expected *LandingPageStatus
		obj.Status = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStatus(), "getter should return the property value")
	})

	t.Run("GetStatus_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		obj.Status = nil

		// Act & Assert
		assert.Nil(t, obj.GetStatus(), "getter should return nil when property is nil")
	})

	t.Run("GetStatus_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *LandingPage
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStatus() // Should return zero value
	})

	t.Run("GetStoreID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		var expected *string
		obj.StoreID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStoreID(), "getter should return the property value")
	})

	t.Run("GetStoreID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		obj.StoreID = nil

		// Act & Assert
		assert.Nil(t, obj.GetStoreID(), "getter should return nil when property is nil")
	})

	t.Run("GetStoreID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *LandingPage
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStoreID() // Should return zero value
	})

	t.Run("GetTemplateID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		var expected *int
		obj.TemplateID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTemplateID(), "getter should return the property value")
	})

	t.Run("GetTemplateID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		obj.TemplateID = nil

		// Act & Assert
		assert.Nil(t, obj.GetTemplateID(), "getter should return nil when property is nil")
	})

	t.Run("GetTemplateID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *LandingPage
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTemplateID() // Should return zero value
	})

	t.Run("GetTitle", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		var expected *string
		obj.Title = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTitle(), "getter should return the property value")
	})

	t.Run("GetTitle_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		obj.Title = nil

		// Act & Assert
		assert.Nil(t, obj.GetTitle(), "getter should return nil when property is nil")
	})

	t.Run("GetTitle_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *LandingPage
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTitle() // Should return zero value
	})

	t.Run("GetTracking", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		var expected *LandingPageTracking
		obj.Tracking = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTracking(), "getter should return the property value")
	})

	t.Run("GetTracking_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		obj.Tracking = nil

		// Act & Assert
		assert.Nil(t, obj.GetTracking(), "getter should return nil when property is nil")
	})

	t.Run("GetTracking_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *LandingPage
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTracking() // Should return zero value
	})

	t.Run("GetUnpublishedAt", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		var expected *time.Time
		obj.UnpublishedAt = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetUnpublishedAt(), "getter should return the property value")
	})

	t.Run("GetUnpublishedAt_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		obj.UnpublishedAt = nil

		// Act & Assert
		assert.Nil(t, obj.GetUnpublishedAt(), "getter should return nil when property is nil")
	})

	t.Run("GetUnpublishedAt_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *LandingPage
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetUnpublishedAt() // Should return zero value
	})

	t.Run("GetUpdatedAt", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		var expected *time.Time
		obj.UpdatedAt = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetUpdatedAt(), "getter should return the property value")
	})

	t.Run("GetUpdatedAt_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		obj.UpdatedAt = nil

		// Act & Assert
		assert.Nil(t, obj.GetUpdatedAt(), "getter should return nil when property is nil")
	})

	t.Run("GetUpdatedAt_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *LandingPage
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetUpdatedAt() // Should return zero value
	})

	t.Run("GetURL", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		var expected *string
		obj.URL = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetURL(), "getter should return the property value")
	})

	t.Run("GetURL_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		obj.URL = nil

		// Act & Assert
		assert.Nil(t, obj.GetURL(), "getter should return nil when property is nil")
	})

	t.Run("GetURL_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *LandingPage
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetURL() // Should return zero value
	})

	t.Run("GetWebID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		var expected *int
		obj.WebID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetWebID(), "getter should return the property value")
	})

	t.Run("GetWebID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		obj.WebID = nil

		// Act & Assert
		assert.Nil(t, obj.GetWebID(), "getter should return nil when property is nil")
	})

	t.Run("GetWebID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *LandingPage
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetWebID() // Should return zero value
	})

}

func TestSettersMarkExplicitLandingPage(t *testing.T) {
	t.Run("SetLinks_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		var fernTestValueLinks []*LandingPageLinksItem

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

	t.Run("SetCreatedAt_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
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

	t.Run("SetCreatedBySource_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		var fernTestValueCreatedBySource *string

		// Act
		obj.SetCreatedBySource(fernTestValueCreatedBySource)

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

	t.Run("SetDescription_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		var fernTestValueDescription *string

		// Act
		obj.SetDescription(fernTestValueDescription)

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
		obj := &LandingPage{}
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

	t.Run("SetListID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		var fernTestValueListID *string

		// Act
		obj.SetListID(fernTestValueListID)

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

	t.Run("SetName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
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

	t.Run("SetPublishedAt_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		var fernTestValuePublishedAt *time.Time

		// Act
		obj.SetPublishedAt(fernTestValuePublishedAt)

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
		obj := &LandingPage{}
		var fernTestValueStatus *LandingPageStatus

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

	t.Run("SetStoreID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		var fernTestValueStoreID *string

		// Act
		obj.SetStoreID(fernTestValueStoreID)

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

	t.Run("SetTemplateID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		var fernTestValueTemplateID *int

		// Act
		obj.SetTemplateID(fernTestValueTemplateID)

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

	t.Run("SetTitle_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		var fernTestValueTitle *string

		// Act
		obj.SetTitle(fernTestValueTitle)

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

	t.Run("SetTracking_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		var fernTestValueTracking *LandingPageTracking

		// Act
		obj.SetTracking(fernTestValueTracking)

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

	t.Run("SetUnpublishedAt_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		var fernTestValueUnpublishedAt *time.Time

		// Act
		obj.SetUnpublishedAt(fernTestValueUnpublishedAt)

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

	t.Run("SetUpdatedAt_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		var fernTestValueUpdatedAt *time.Time

		// Act
		obj.SetUpdatedAt(fernTestValueUpdatedAt)

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
		obj := &LandingPage{}
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

	t.Run("SetWebID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}
		var fernTestValueWebID *int

		// Act
		obj.SetWebID(fernTestValueWebID)

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

func TestSettersLandingPageLinksItem(t *testing.T) {
	t.Run("SetHref", func(t *testing.T) {
		obj := &LandingPageLinksItem{}
		var fernTestValueHref *string
		obj.SetHref(fernTestValueHref)
		assert.Equal(t, fernTestValueHref, obj.Href)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMethod", func(t *testing.T) {
		obj := &LandingPageLinksItem{}
		var fernTestValueMethod *LandingPageLinksItemMethod
		obj.SetMethod(fernTestValueMethod)
		assert.Equal(t, fernTestValueMethod, obj.Method)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRel", func(t *testing.T) {
		obj := &LandingPageLinksItem{}
		var fernTestValueRel *string
		obj.SetRel(fernTestValueRel)
		assert.Equal(t, fernTestValueRel, obj.Rel)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSchema", func(t *testing.T) {
		obj := &LandingPageLinksItem{}
		var fernTestValueSchema *string
		obj.SetSchema(fernTestValueSchema)
		assert.Equal(t, fernTestValueSchema, obj.Schema)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTargetSchema", func(t *testing.T) {
		obj := &LandingPageLinksItem{}
		var fernTestValueTargetSchema *string
		obj.SetTargetSchema(fernTestValueTargetSchema)
		assert.Equal(t, fernTestValueTargetSchema, obj.TargetSchema)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersLandingPageLinksItem(t *testing.T) {
	t.Run("GetHref", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPageLinksItem{}
		var expected *string
		obj.Href = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetHref(), "getter should return the property value")
	})

	t.Run("GetHref_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPageLinksItem{}
		obj.Href = nil

		// Act & Assert
		assert.Nil(t, obj.GetHref(), "getter should return nil when property is nil")
	})

	t.Run("GetHref_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *LandingPageLinksItem
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
		obj := &LandingPageLinksItem{}
		var expected *LandingPageLinksItemMethod
		obj.Method = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMethod(), "getter should return the property value")
	})

	t.Run("GetMethod_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPageLinksItem{}
		obj.Method = nil

		// Act & Assert
		assert.Nil(t, obj.GetMethod(), "getter should return nil when property is nil")
	})

	t.Run("GetMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *LandingPageLinksItem
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
		obj := &LandingPageLinksItem{}
		var expected *string
		obj.Rel = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetRel(), "getter should return the property value")
	})

	t.Run("GetRel_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPageLinksItem{}
		obj.Rel = nil

		// Act & Assert
		assert.Nil(t, obj.GetRel(), "getter should return nil when property is nil")
	})

	t.Run("GetRel_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *LandingPageLinksItem
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
		obj := &LandingPageLinksItem{}
		var expected *string
		obj.Schema = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSchema(), "getter should return the property value")
	})

	t.Run("GetSchema_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPageLinksItem{}
		obj.Schema = nil

		// Act & Assert
		assert.Nil(t, obj.GetSchema(), "getter should return nil when property is nil")
	})

	t.Run("GetSchema_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *LandingPageLinksItem
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
		obj := &LandingPageLinksItem{}
		var expected *string
		obj.TargetSchema = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTargetSchema(), "getter should return the property value")
	})

	t.Run("GetTargetSchema_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPageLinksItem{}
		obj.TargetSchema = nil

		// Act & Assert
		assert.Nil(t, obj.GetTargetSchema(), "getter should return nil when property is nil")
	})

	t.Run("GetTargetSchema_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *LandingPageLinksItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTargetSchema() // Should return zero value
	})

}

func TestSettersMarkExplicitLandingPageLinksItem(t *testing.T) {
	t.Run("SetHref_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPageLinksItem{}
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
		obj := &LandingPageLinksItem{}
		var fernTestValueMethod *LandingPageLinksItemMethod

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
		obj := &LandingPageLinksItem{}
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
		obj := &LandingPageLinksItem{}
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
		obj := &LandingPageLinksItem{}
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

func TestSettersLandingPageTracking(t *testing.T) {
	t.Run("SetEnableRestrictedDataProcessing", func(t *testing.T) {
		obj := &LandingPageTracking{}
		var fernTestValueEnableRestrictedDataProcessing *bool
		obj.SetEnableRestrictedDataProcessing(fernTestValueEnableRestrictedDataProcessing)
		assert.Equal(t, fernTestValueEnableRestrictedDataProcessing, obj.EnableRestrictedDataProcessing)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTrackWithMailchimp", func(t *testing.T) {
		obj := &LandingPageTracking{}
		var fernTestValueTrackWithMailchimp *bool
		obj.SetTrackWithMailchimp(fernTestValueTrackWithMailchimp)
		assert.Equal(t, fernTestValueTrackWithMailchimp, obj.TrackWithMailchimp)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersLandingPageTracking(t *testing.T) {
	t.Run("GetEnableRestrictedDataProcessing", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPageTracking{}
		var expected *bool
		obj.EnableRestrictedDataProcessing = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEnableRestrictedDataProcessing(), "getter should return the property value")
	})

	t.Run("GetEnableRestrictedDataProcessing_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPageTracking{}
		obj.EnableRestrictedDataProcessing = nil

		// Act & Assert
		assert.Nil(t, obj.GetEnableRestrictedDataProcessing(), "getter should return nil when property is nil")
	})

	t.Run("GetEnableRestrictedDataProcessing_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *LandingPageTracking
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEnableRestrictedDataProcessing() // Should return zero value
	})

	t.Run("GetTrackWithMailchimp", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPageTracking{}
		var expected *bool
		obj.TrackWithMailchimp = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTrackWithMailchimp(), "getter should return the property value")
	})

	t.Run("GetTrackWithMailchimp_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPageTracking{}
		obj.TrackWithMailchimp = nil

		// Act & Assert
		assert.Nil(t, obj.GetTrackWithMailchimp(), "getter should return nil when property is nil")
	})

	t.Run("GetTrackWithMailchimp_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *LandingPageTracking
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTrackWithMailchimp() // Should return zero value
	})

}

func TestSettersMarkExplicitLandingPageTracking(t *testing.T) {
	t.Run("SetEnableRestrictedDataProcessing_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPageTracking{}
		var fernTestValueEnableRestrictedDataProcessing *bool

		// Act
		obj.SetEnableRestrictedDataProcessing(fernTestValueEnableRestrictedDataProcessing)

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

	t.Run("SetTrackWithMailchimp_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPageTracking{}
		var fernTestValueTrackWithMailchimp *bool

		// Act
		obj.SetTrackWithMailchimp(fernTestValueTrackWithMailchimp)

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

func TestSettersCreateLandingPagesRequestTracking(t *testing.T) {
	t.Run("SetEnableRestrictedDataProcessing", func(t *testing.T) {
		obj := &CreateLandingPagesRequestTracking{}
		var fernTestValueEnableRestrictedDataProcessing *bool
		obj.SetEnableRestrictedDataProcessing(fernTestValueEnableRestrictedDataProcessing)
		assert.Equal(t, fernTestValueEnableRestrictedDataProcessing, obj.EnableRestrictedDataProcessing)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTrackWithMailchimp", func(t *testing.T) {
		obj := &CreateLandingPagesRequestTracking{}
		var fernTestValueTrackWithMailchimp *bool
		obj.SetTrackWithMailchimp(fernTestValueTrackWithMailchimp)
		assert.Equal(t, fernTestValueTrackWithMailchimp, obj.TrackWithMailchimp)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersCreateLandingPagesRequestTracking(t *testing.T) {
	t.Run("GetEnableRestrictedDataProcessing", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateLandingPagesRequestTracking{}
		var expected *bool
		obj.EnableRestrictedDataProcessing = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEnableRestrictedDataProcessing(), "getter should return the property value")
	})

	t.Run("GetEnableRestrictedDataProcessing_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateLandingPagesRequestTracking{}
		obj.EnableRestrictedDataProcessing = nil

		// Act & Assert
		assert.Nil(t, obj.GetEnableRestrictedDataProcessing(), "getter should return nil when property is nil")
	})

	t.Run("GetEnableRestrictedDataProcessing_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateLandingPagesRequestTracking
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEnableRestrictedDataProcessing() // Should return zero value
	})

	t.Run("GetTrackWithMailchimp", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateLandingPagesRequestTracking{}
		var expected *bool
		obj.TrackWithMailchimp = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTrackWithMailchimp(), "getter should return the property value")
	})

	t.Run("GetTrackWithMailchimp_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateLandingPagesRequestTracking{}
		obj.TrackWithMailchimp = nil

		// Act & Assert
		assert.Nil(t, obj.GetTrackWithMailchimp(), "getter should return nil when property is nil")
	})

	t.Run("GetTrackWithMailchimp_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateLandingPagesRequestTracking
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTrackWithMailchimp() // Should return zero value
	})

}

func TestSettersMarkExplicitCreateLandingPagesRequestTracking(t *testing.T) {
	t.Run("SetEnableRestrictedDataProcessing_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateLandingPagesRequestTracking{}
		var fernTestValueEnableRestrictedDataProcessing *bool

		// Act
		obj.SetEnableRestrictedDataProcessing(fernTestValueEnableRestrictedDataProcessing)

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

	t.Run("SetTrackWithMailchimp_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateLandingPagesRequestTracking{}
		var fernTestValueTrackWithMailchimp *bool

		// Act
		obj.SetTrackWithMailchimp(fernTestValueTrackWithMailchimp)

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

func TestSettersListContentLandingPagesResponse(t *testing.T) {
	t.Run("SetLinks", func(t *testing.T) {
		obj := &ListContentLandingPagesResponse{}
		var fernTestValueLinks []*ListContentLandingPagesResponseLinksItem
		obj.SetLinks(fernTestValueLinks)
		assert.Equal(t, fernTestValueLinks, obj.Links)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetHTML", func(t *testing.T) {
		obj := &ListContentLandingPagesResponse{}
		var fernTestValueHTML *string
		obj.SetHTML(fernTestValueHTML)
		assert.Equal(t, fernTestValueHTML, obj.HTML)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetJSON", func(t *testing.T) {
		obj := &ListContentLandingPagesResponse{}
		var fernTestValueJSON *string
		obj.SetJSON(fernTestValueJSON)
		assert.Equal(t, fernTestValueJSON, obj.JSON)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersListContentLandingPagesResponse(t *testing.T) {
	t.Run("GetLinks", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListContentLandingPagesResponse{}
		var expected []*ListContentLandingPagesResponseLinksItem
		obj.Links = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetLinks(), "getter should return the property value")
	})

	t.Run("GetLinks_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListContentLandingPagesResponse{}
		obj.Links = nil

		// Act & Assert
		assert.Nil(t, obj.GetLinks(), "getter should return nil when property is nil")
	})

	t.Run("GetLinks_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListContentLandingPagesResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetLinks() // Should return zero value
	})

	t.Run("GetHTML", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListContentLandingPagesResponse{}
		var expected *string
		obj.HTML = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetHTML(), "getter should return the property value")
	})

	t.Run("GetHTML_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListContentLandingPagesResponse{}
		obj.HTML = nil

		// Act & Assert
		assert.Nil(t, obj.GetHTML(), "getter should return nil when property is nil")
	})

	t.Run("GetHTML_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListContentLandingPagesResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetHTML() // Should return zero value
	})

	t.Run("GetJSON", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListContentLandingPagesResponse{}
		var expected *string
		obj.JSON = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetJSON(), "getter should return the property value")
	})

	t.Run("GetJSON_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListContentLandingPagesResponse{}
		obj.JSON = nil

		// Act & Assert
		assert.Nil(t, obj.GetJSON(), "getter should return nil when property is nil")
	})

	t.Run("GetJSON_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListContentLandingPagesResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetJSON() // Should return zero value
	})

}

func TestSettersMarkExplicitListContentLandingPagesResponse(t *testing.T) {
	t.Run("SetLinks_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListContentLandingPagesResponse{}
		var fernTestValueLinks []*ListContentLandingPagesResponseLinksItem

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

	t.Run("SetHTML_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListContentLandingPagesResponse{}
		var fernTestValueHTML *string

		// Act
		obj.SetHTML(fernTestValueHTML)

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

	t.Run("SetJSON_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListContentLandingPagesResponse{}
		var fernTestValueJSON *string

		// Act
		obj.SetJSON(fernTestValueJSON)

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

func TestSettersListContentLandingPagesResponseLinksItem(t *testing.T) {
	t.Run("SetHref", func(t *testing.T) {
		obj := &ListContentLandingPagesResponseLinksItem{}
		var fernTestValueHref *string
		obj.SetHref(fernTestValueHref)
		assert.Equal(t, fernTestValueHref, obj.Href)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMethod", func(t *testing.T) {
		obj := &ListContentLandingPagesResponseLinksItem{}
		var fernTestValueMethod *ListContentLandingPagesResponseLinksItemMethod
		obj.SetMethod(fernTestValueMethod)
		assert.Equal(t, fernTestValueMethod, obj.Method)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRel", func(t *testing.T) {
		obj := &ListContentLandingPagesResponseLinksItem{}
		var fernTestValueRel *string
		obj.SetRel(fernTestValueRel)
		assert.Equal(t, fernTestValueRel, obj.Rel)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSchema", func(t *testing.T) {
		obj := &ListContentLandingPagesResponseLinksItem{}
		var fernTestValueSchema *string
		obj.SetSchema(fernTestValueSchema)
		assert.Equal(t, fernTestValueSchema, obj.Schema)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTargetSchema", func(t *testing.T) {
		obj := &ListContentLandingPagesResponseLinksItem{}
		var fernTestValueTargetSchema *string
		obj.SetTargetSchema(fernTestValueTargetSchema)
		assert.Equal(t, fernTestValueTargetSchema, obj.TargetSchema)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersListContentLandingPagesResponseLinksItem(t *testing.T) {
	t.Run("GetHref", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListContentLandingPagesResponseLinksItem{}
		var expected *string
		obj.Href = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetHref(), "getter should return the property value")
	})

	t.Run("GetHref_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListContentLandingPagesResponseLinksItem{}
		obj.Href = nil

		// Act & Assert
		assert.Nil(t, obj.GetHref(), "getter should return nil when property is nil")
	})

	t.Run("GetHref_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListContentLandingPagesResponseLinksItem
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
		obj := &ListContentLandingPagesResponseLinksItem{}
		var expected *ListContentLandingPagesResponseLinksItemMethod
		obj.Method = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMethod(), "getter should return the property value")
	})

	t.Run("GetMethod_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListContentLandingPagesResponseLinksItem{}
		obj.Method = nil

		// Act & Assert
		assert.Nil(t, obj.GetMethod(), "getter should return nil when property is nil")
	})

	t.Run("GetMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListContentLandingPagesResponseLinksItem
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
		obj := &ListContentLandingPagesResponseLinksItem{}
		var expected *string
		obj.Rel = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetRel(), "getter should return the property value")
	})

	t.Run("GetRel_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListContentLandingPagesResponseLinksItem{}
		obj.Rel = nil

		// Act & Assert
		assert.Nil(t, obj.GetRel(), "getter should return nil when property is nil")
	})

	t.Run("GetRel_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListContentLandingPagesResponseLinksItem
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
		obj := &ListContentLandingPagesResponseLinksItem{}
		var expected *string
		obj.Schema = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSchema(), "getter should return the property value")
	})

	t.Run("GetSchema_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListContentLandingPagesResponseLinksItem{}
		obj.Schema = nil

		// Act & Assert
		assert.Nil(t, obj.GetSchema(), "getter should return nil when property is nil")
	})

	t.Run("GetSchema_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListContentLandingPagesResponseLinksItem
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
		obj := &ListContentLandingPagesResponseLinksItem{}
		var expected *string
		obj.TargetSchema = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTargetSchema(), "getter should return the property value")
	})

	t.Run("GetTargetSchema_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListContentLandingPagesResponseLinksItem{}
		obj.TargetSchema = nil

		// Act & Assert
		assert.Nil(t, obj.GetTargetSchema(), "getter should return nil when property is nil")
	})

	t.Run("GetTargetSchema_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListContentLandingPagesResponseLinksItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTargetSchema() // Should return zero value
	})

}

func TestSettersMarkExplicitListContentLandingPagesResponseLinksItem(t *testing.T) {
	t.Run("SetHref_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListContentLandingPagesResponseLinksItem{}
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
		obj := &ListContentLandingPagesResponseLinksItem{}
		var fernTestValueMethod *ListContentLandingPagesResponseLinksItemMethod

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
		obj := &ListContentLandingPagesResponseLinksItem{}
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
		obj := &ListContentLandingPagesResponseLinksItem{}
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
		obj := &ListContentLandingPagesResponseLinksItem{}
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

func TestSettersListLandingPagesResponse(t *testing.T) {
	t.Run("SetLinks", func(t *testing.T) {
		obj := &ListLandingPagesResponse{}
		var fernTestValueLinks []*ListLandingPagesResponseLinksItem
		obj.SetLinks(fernTestValueLinks)
		assert.Equal(t, fernTestValueLinks, obj.Links)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetLandingPages", func(t *testing.T) {
		obj := &ListLandingPagesResponse{}
		var fernTestValueLandingPages []*LandingPage
		obj.SetLandingPages(fernTestValueLandingPages)
		assert.Equal(t, fernTestValueLandingPages, obj.LandingPages)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTotalItems", func(t *testing.T) {
		obj := &ListLandingPagesResponse{}
		var fernTestValueTotalItems *int
		obj.SetTotalItems(fernTestValueTotalItems)
		assert.Equal(t, fernTestValueTotalItems, obj.TotalItems)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersListLandingPagesResponse(t *testing.T) {
	t.Run("GetLinks", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListLandingPagesResponse{}
		var expected []*ListLandingPagesResponseLinksItem
		obj.Links = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetLinks(), "getter should return the property value")
	})

	t.Run("GetLinks_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListLandingPagesResponse{}
		obj.Links = nil

		// Act & Assert
		assert.Nil(t, obj.GetLinks(), "getter should return nil when property is nil")
	})

	t.Run("GetLinks_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListLandingPagesResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetLinks() // Should return zero value
	})

	t.Run("GetLandingPages", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListLandingPagesResponse{}
		var expected []*LandingPage
		obj.LandingPages = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetLandingPages(), "getter should return the property value")
	})

	t.Run("GetLandingPages_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListLandingPagesResponse{}
		obj.LandingPages = nil

		// Act & Assert
		assert.Nil(t, obj.GetLandingPages(), "getter should return nil when property is nil")
	})

	t.Run("GetLandingPages_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListLandingPagesResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetLandingPages() // Should return zero value
	})

	t.Run("GetTotalItems", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListLandingPagesResponse{}
		var expected *int
		obj.TotalItems = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTotalItems(), "getter should return the property value")
	})

	t.Run("GetTotalItems_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListLandingPagesResponse{}
		obj.TotalItems = nil

		// Act & Assert
		assert.Nil(t, obj.GetTotalItems(), "getter should return nil when property is nil")
	})

	t.Run("GetTotalItems_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListLandingPagesResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTotalItems() // Should return zero value
	})

}

func TestSettersMarkExplicitListLandingPagesResponse(t *testing.T) {
	t.Run("SetLinks_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListLandingPagesResponse{}
		var fernTestValueLinks []*ListLandingPagesResponseLinksItem

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

	t.Run("SetLandingPages_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListLandingPagesResponse{}
		var fernTestValueLandingPages []*LandingPage

		// Act
		obj.SetLandingPages(fernTestValueLandingPages)

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
		obj := &ListLandingPagesResponse{}
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

func TestSettersListLandingPagesResponseLinksItem(t *testing.T) {
	t.Run("SetHref", func(t *testing.T) {
		obj := &ListLandingPagesResponseLinksItem{}
		var fernTestValueHref *string
		obj.SetHref(fernTestValueHref)
		assert.Equal(t, fernTestValueHref, obj.Href)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMethod", func(t *testing.T) {
		obj := &ListLandingPagesResponseLinksItem{}
		var fernTestValueMethod *ListLandingPagesResponseLinksItemMethod
		obj.SetMethod(fernTestValueMethod)
		assert.Equal(t, fernTestValueMethod, obj.Method)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRel", func(t *testing.T) {
		obj := &ListLandingPagesResponseLinksItem{}
		var fernTestValueRel *string
		obj.SetRel(fernTestValueRel)
		assert.Equal(t, fernTestValueRel, obj.Rel)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSchema", func(t *testing.T) {
		obj := &ListLandingPagesResponseLinksItem{}
		var fernTestValueSchema *string
		obj.SetSchema(fernTestValueSchema)
		assert.Equal(t, fernTestValueSchema, obj.Schema)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTargetSchema", func(t *testing.T) {
		obj := &ListLandingPagesResponseLinksItem{}
		var fernTestValueTargetSchema *string
		obj.SetTargetSchema(fernTestValueTargetSchema)
		assert.Equal(t, fernTestValueTargetSchema, obj.TargetSchema)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersListLandingPagesResponseLinksItem(t *testing.T) {
	t.Run("GetHref", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListLandingPagesResponseLinksItem{}
		var expected *string
		obj.Href = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetHref(), "getter should return the property value")
	})

	t.Run("GetHref_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListLandingPagesResponseLinksItem{}
		obj.Href = nil

		// Act & Assert
		assert.Nil(t, obj.GetHref(), "getter should return nil when property is nil")
	})

	t.Run("GetHref_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListLandingPagesResponseLinksItem
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
		obj := &ListLandingPagesResponseLinksItem{}
		var expected *ListLandingPagesResponseLinksItemMethod
		obj.Method = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMethod(), "getter should return the property value")
	})

	t.Run("GetMethod_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListLandingPagesResponseLinksItem{}
		obj.Method = nil

		// Act & Assert
		assert.Nil(t, obj.GetMethod(), "getter should return nil when property is nil")
	})

	t.Run("GetMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListLandingPagesResponseLinksItem
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
		obj := &ListLandingPagesResponseLinksItem{}
		var expected *string
		obj.Rel = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetRel(), "getter should return the property value")
	})

	t.Run("GetRel_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListLandingPagesResponseLinksItem{}
		obj.Rel = nil

		// Act & Assert
		assert.Nil(t, obj.GetRel(), "getter should return nil when property is nil")
	})

	t.Run("GetRel_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListLandingPagesResponseLinksItem
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
		obj := &ListLandingPagesResponseLinksItem{}
		var expected *string
		obj.Schema = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSchema(), "getter should return the property value")
	})

	t.Run("GetSchema_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListLandingPagesResponseLinksItem{}
		obj.Schema = nil

		// Act & Assert
		assert.Nil(t, obj.GetSchema(), "getter should return nil when property is nil")
	})

	t.Run("GetSchema_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListLandingPagesResponseLinksItem
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
		obj := &ListLandingPagesResponseLinksItem{}
		var expected *string
		obj.TargetSchema = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTargetSchema(), "getter should return the property value")
	})

	t.Run("GetTargetSchema_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListLandingPagesResponseLinksItem{}
		obj.TargetSchema = nil

		// Act & Assert
		assert.Nil(t, obj.GetTargetSchema(), "getter should return nil when property is nil")
	})

	t.Run("GetTargetSchema_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListLandingPagesResponseLinksItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTargetSchema() // Should return zero value
	})

}

func TestSettersMarkExplicitListLandingPagesResponseLinksItem(t *testing.T) {
	t.Run("SetHref_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListLandingPagesResponseLinksItem{}
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
		obj := &ListLandingPagesResponseLinksItem{}
		var fernTestValueMethod *ListLandingPagesResponseLinksItemMethod

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
		obj := &ListLandingPagesResponseLinksItem{}
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
		obj := &ListLandingPagesResponseLinksItem{}
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
		obj := &ListLandingPagesResponseLinksItem{}
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

func TestSettersUpdateLandingPagesRequestTracking(t *testing.T) {
	t.Run("SetEnableRestrictedDataProcessing", func(t *testing.T) {
		obj := &UpdateLandingPagesRequestTracking{}
		var fernTestValueEnableRestrictedDataProcessing *bool
		obj.SetEnableRestrictedDataProcessing(fernTestValueEnableRestrictedDataProcessing)
		assert.Equal(t, fernTestValueEnableRestrictedDataProcessing, obj.EnableRestrictedDataProcessing)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTrackWithMailchimp", func(t *testing.T) {
		obj := &UpdateLandingPagesRequestTracking{}
		var fernTestValueTrackWithMailchimp *bool
		obj.SetTrackWithMailchimp(fernTestValueTrackWithMailchimp)
		assert.Equal(t, fernTestValueTrackWithMailchimp, obj.TrackWithMailchimp)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersUpdateLandingPagesRequestTracking(t *testing.T) {
	t.Run("GetEnableRestrictedDataProcessing", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &UpdateLandingPagesRequestTracking{}
		var expected *bool
		obj.EnableRestrictedDataProcessing = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEnableRestrictedDataProcessing(), "getter should return the property value")
	})

	t.Run("GetEnableRestrictedDataProcessing_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &UpdateLandingPagesRequestTracking{}
		obj.EnableRestrictedDataProcessing = nil

		// Act & Assert
		assert.Nil(t, obj.GetEnableRestrictedDataProcessing(), "getter should return nil when property is nil")
	})

	t.Run("GetEnableRestrictedDataProcessing_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *UpdateLandingPagesRequestTracking
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEnableRestrictedDataProcessing() // Should return zero value
	})

	t.Run("GetTrackWithMailchimp", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &UpdateLandingPagesRequestTracking{}
		var expected *bool
		obj.TrackWithMailchimp = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTrackWithMailchimp(), "getter should return the property value")
	})

	t.Run("GetTrackWithMailchimp_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &UpdateLandingPagesRequestTracking{}
		obj.TrackWithMailchimp = nil

		// Act & Assert
		assert.Nil(t, obj.GetTrackWithMailchimp(), "getter should return nil when property is nil")
	})

	t.Run("GetTrackWithMailchimp_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *UpdateLandingPagesRequestTracking
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTrackWithMailchimp() // Should return zero value
	})

}

func TestSettersMarkExplicitUpdateLandingPagesRequestTracking(t *testing.T) {
	t.Run("SetEnableRestrictedDataProcessing_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &UpdateLandingPagesRequestTracking{}
		var fernTestValueEnableRestrictedDataProcessing *bool

		// Act
		obj.SetEnableRestrictedDataProcessing(fernTestValueEnableRestrictedDataProcessing)

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

	t.Run("SetTrackWithMailchimp_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &UpdateLandingPagesRequestTracking{}
		var fernTestValueTrackWithMailchimp *bool

		// Act
		obj.SetTrackWithMailchimp(fernTestValueTrackWithMailchimp)

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

func TestSettersUpdateLandingPagesRequest(t *testing.T) {
	t.Run("SetPageID", func(t *testing.T) {
		obj := &UpdateLandingPagesRequest{}
		var fernTestValuePageID string
		obj.SetPageID(fernTestValuePageID)
		assert.Equal(t, fernTestValuePageID, obj.PageID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetDescription", func(t *testing.T) {
		obj := &UpdateLandingPagesRequest{}
		var fernTestValueDescription *string
		obj.SetDescription(fernTestValueDescription)
		assert.Equal(t, fernTestValueDescription, obj.Description)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetListID", func(t *testing.T) {
		obj := &UpdateLandingPagesRequest{}
		var fernTestValueListID *string
		obj.SetListID(fernTestValueListID)
		assert.Equal(t, fernTestValueListID, obj.ListID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetName", func(t *testing.T) {
		obj := &UpdateLandingPagesRequest{}
		var fernTestValueName *string
		obj.SetName(fernTestValueName)
		assert.Equal(t, fernTestValueName, obj.Name)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStoreID", func(t *testing.T) {
		obj := &UpdateLandingPagesRequest{}
		var fernTestValueStoreID *string
		obj.SetStoreID(fernTestValueStoreID)
		assert.Equal(t, fernTestValueStoreID, obj.StoreID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTitle", func(t *testing.T) {
		obj := &UpdateLandingPagesRequest{}
		var fernTestValueTitle *string
		obj.SetTitle(fernTestValueTitle)
		assert.Equal(t, fernTestValueTitle, obj.Title)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTracking", func(t *testing.T) {
		obj := &UpdateLandingPagesRequest{}
		var fernTestValueTracking *UpdateLandingPagesRequestTracking
		obj.SetTracking(fernTestValueTracking)
		assert.Equal(t, fernTestValueTracking, obj.Tracking)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitUpdateLandingPagesRequest(t *testing.T) {
	t.Run("SetPageID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &UpdateLandingPagesRequest{}
		var fernTestValuePageID string

		// Act
		obj.SetPageID(fernTestValuePageID)

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

	t.Run("SetDescription_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &UpdateLandingPagesRequest{}
		var fernTestValueDescription *string

		// Act
		obj.SetDescription(fernTestValueDescription)

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

	t.Run("SetListID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &UpdateLandingPagesRequest{}
		var fernTestValueListID *string

		// Act
		obj.SetListID(fernTestValueListID)

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

	t.Run("SetName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &UpdateLandingPagesRequest{}
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

	t.Run("SetStoreID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &UpdateLandingPagesRequest{}
		var fernTestValueStoreID *string

		// Act
		obj.SetStoreID(fernTestValueStoreID)

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

	t.Run("SetTitle_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &UpdateLandingPagesRequest{}
		var fernTestValueTitle *string

		// Act
		obj.SetTitle(fernTestValueTitle)

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

	t.Run("SetTracking_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &UpdateLandingPagesRequest{}
		var fernTestValueTracking *UpdateLandingPagesRequestTracking

		// Act
		obj.SetTracking(fernTestValueTracking)

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

func TestJSONMarshalingCreateLandingPagesRequestTracking(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateLandingPagesRequestTracking{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled CreateLandingPagesRequestTracking
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj CreateLandingPagesRequestTracking
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj CreateLandingPagesRequestTracking
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingLandingPage(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPage{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled LandingPage
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj LandingPage
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj LandingPage
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingLandingPageLinksItem(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPageLinksItem{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled LandingPageLinksItem
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj LandingPageLinksItem
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj LandingPageLinksItem
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingLandingPageTracking(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &LandingPageTracking{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled LandingPageTracking
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj LandingPageTracking
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj LandingPageTracking
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingListContentLandingPagesResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListContentLandingPagesResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ListContentLandingPagesResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ListContentLandingPagesResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ListContentLandingPagesResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingListContentLandingPagesResponseLinksItem(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListContentLandingPagesResponseLinksItem{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ListContentLandingPagesResponseLinksItem
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ListContentLandingPagesResponseLinksItem
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ListContentLandingPagesResponseLinksItem
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingListLandingPagesResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListLandingPagesResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ListLandingPagesResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ListLandingPagesResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ListLandingPagesResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingListLandingPagesResponseLinksItem(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListLandingPagesResponseLinksItem{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ListLandingPagesResponseLinksItem
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ListLandingPagesResponseLinksItem
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ListLandingPagesResponseLinksItem
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingUpdateLandingPagesRequestTracking(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &UpdateLandingPagesRequestTracking{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled UpdateLandingPagesRequestTracking
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj UpdateLandingPagesRequestTracking
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj UpdateLandingPagesRequestTracking
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestStringCreateLandingPagesRequestTracking(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &CreateLandingPagesRequestTracking{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *CreateLandingPagesRequestTracking
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringLandingPage(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &LandingPage{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *LandingPage
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringLandingPageLinksItem(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &LandingPageLinksItem{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *LandingPageLinksItem
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringLandingPageTracking(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &LandingPageTracking{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *LandingPageTracking
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringListContentLandingPagesResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ListContentLandingPagesResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListContentLandingPagesResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringListContentLandingPagesResponseLinksItem(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ListContentLandingPagesResponseLinksItem{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListContentLandingPagesResponseLinksItem
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringListLandingPagesResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ListLandingPagesResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListLandingPagesResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringListLandingPagesResponseLinksItem(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ListLandingPagesResponseLinksItem{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListLandingPagesResponseLinksItem
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringUpdateLandingPagesRequestTracking(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &UpdateLandingPagesRequestTracking{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *UpdateLandingPagesRequestTracking
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestEnumCreateLandingPagesRequestType(t *testing.T) {
	t.Run("NewFromString_signup", func(t *testing.T) {
		t.Parallel()
		val, err := NewCreateLandingPagesRequestTypeFromString("signup")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, CreateLandingPagesRequestType("signup"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_product", func(t *testing.T) {
		t.Parallel()
		val, err := NewCreateLandingPagesRequestTypeFromString("product")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, CreateLandingPagesRequestType("product"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewCreateLandingPagesRequestTypeFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewCreateLandingPagesRequestTypeFromString("signup")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumLandingPageLinksItemMethod(t *testing.T) {
	t.Run("NewFromString_GET", func(t *testing.T) {
		t.Parallel()
		val, err := NewLandingPageLinksItemMethodFromString("GET")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, LandingPageLinksItemMethod("GET"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_POST", func(t *testing.T) {
		t.Parallel()
		val, err := NewLandingPageLinksItemMethodFromString("POST")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, LandingPageLinksItemMethod("POST"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_PUT", func(t *testing.T) {
		t.Parallel()
		val, err := NewLandingPageLinksItemMethodFromString("PUT")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, LandingPageLinksItemMethod("PUT"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_PATCH", func(t *testing.T) {
		t.Parallel()
		val, err := NewLandingPageLinksItemMethodFromString("PATCH")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, LandingPageLinksItemMethod("PATCH"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_DELETE", func(t *testing.T) {
		t.Parallel()
		val, err := NewLandingPageLinksItemMethodFromString("DELETE")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, LandingPageLinksItemMethod("DELETE"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_OPTIONS", func(t *testing.T) {
		t.Parallel()
		val, err := NewLandingPageLinksItemMethodFromString("OPTIONS")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, LandingPageLinksItemMethod("OPTIONS"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_HEAD", func(t *testing.T) {
		t.Parallel()
		val, err := NewLandingPageLinksItemMethodFromString("HEAD")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, LandingPageLinksItemMethod("HEAD"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewLandingPageLinksItemMethodFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewLandingPageLinksItemMethodFromString("GET")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumLandingPageStatus(t *testing.T) {
	t.Run("NewFromString_published", func(t *testing.T) {
		t.Parallel()
		val, err := NewLandingPageStatusFromString("published")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, LandingPageStatus("published"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_unpublished", func(t *testing.T) {
		t.Parallel()
		val, err := NewLandingPageStatusFromString("unpublished")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, LandingPageStatus("unpublished"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_draft", func(t *testing.T) {
		t.Parallel()
		val, err := NewLandingPageStatusFromString("draft")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, LandingPageStatus("draft"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewLandingPageStatusFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewLandingPageStatusFromString("published")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumListContentLandingPagesResponseLinksItemMethod(t *testing.T) {
	t.Run("NewFromString_GET", func(t *testing.T) {
		t.Parallel()
		val, err := NewListContentLandingPagesResponseLinksItemMethodFromString("GET")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListContentLandingPagesResponseLinksItemMethod("GET"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_POST", func(t *testing.T) {
		t.Parallel()
		val, err := NewListContentLandingPagesResponseLinksItemMethodFromString("POST")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListContentLandingPagesResponseLinksItemMethod("POST"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_PUT", func(t *testing.T) {
		t.Parallel()
		val, err := NewListContentLandingPagesResponseLinksItemMethodFromString("PUT")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListContentLandingPagesResponseLinksItemMethod("PUT"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_PATCH", func(t *testing.T) {
		t.Parallel()
		val, err := NewListContentLandingPagesResponseLinksItemMethodFromString("PATCH")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListContentLandingPagesResponseLinksItemMethod("PATCH"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_DELETE", func(t *testing.T) {
		t.Parallel()
		val, err := NewListContentLandingPagesResponseLinksItemMethodFromString("DELETE")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListContentLandingPagesResponseLinksItemMethod("DELETE"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_OPTIONS", func(t *testing.T) {
		t.Parallel()
		val, err := NewListContentLandingPagesResponseLinksItemMethodFromString("OPTIONS")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListContentLandingPagesResponseLinksItemMethod("OPTIONS"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_HEAD", func(t *testing.T) {
		t.Parallel()
		val, err := NewListContentLandingPagesResponseLinksItemMethodFromString("HEAD")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListContentLandingPagesResponseLinksItemMethod("HEAD"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewListContentLandingPagesResponseLinksItemMethodFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewListContentLandingPagesResponseLinksItemMethodFromString("GET")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumListLandingPagesRequestSortDir(t *testing.T) {
	t.Run("NewFromString_ASC", func(t *testing.T) {
		t.Parallel()
		val, err := NewListLandingPagesRequestSortDirFromString("ASC")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListLandingPagesRequestSortDir("ASC"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_DESC", func(t *testing.T) {
		t.Parallel()
		val, err := NewListLandingPagesRequestSortDirFromString("DESC")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListLandingPagesRequestSortDir("DESC"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewListLandingPagesRequestSortDirFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewListLandingPagesRequestSortDirFromString("ASC")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumListLandingPagesRequestSortField(t *testing.T) {
	t.Run("NewFromString_created_at", func(t *testing.T) {
		t.Parallel()
		val, err := NewListLandingPagesRequestSortFieldFromString("created_at")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListLandingPagesRequestSortField("created_at"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_updated_at", func(t *testing.T) {
		t.Parallel()
		val, err := NewListLandingPagesRequestSortFieldFromString("updated_at")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListLandingPagesRequestSortField("updated_at"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewListLandingPagesRequestSortFieldFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewListLandingPagesRequestSortFieldFromString("created_at")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumListLandingPagesResponseLinksItemMethod(t *testing.T) {
	t.Run("NewFromString_GET", func(t *testing.T) {
		t.Parallel()
		val, err := NewListLandingPagesResponseLinksItemMethodFromString("GET")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListLandingPagesResponseLinksItemMethod("GET"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_POST", func(t *testing.T) {
		t.Parallel()
		val, err := NewListLandingPagesResponseLinksItemMethodFromString("POST")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListLandingPagesResponseLinksItemMethod("POST"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_PUT", func(t *testing.T) {
		t.Parallel()
		val, err := NewListLandingPagesResponseLinksItemMethodFromString("PUT")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListLandingPagesResponseLinksItemMethod("PUT"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_PATCH", func(t *testing.T) {
		t.Parallel()
		val, err := NewListLandingPagesResponseLinksItemMethodFromString("PATCH")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListLandingPagesResponseLinksItemMethod("PATCH"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_DELETE", func(t *testing.T) {
		t.Parallel()
		val, err := NewListLandingPagesResponseLinksItemMethodFromString("DELETE")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListLandingPagesResponseLinksItemMethod("DELETE"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_OPTIONS", func(t *testing.T) {
		t.Parallel()
		val, err := NewListLandingPagesResponseLinksItemMethodFromString("OPTIONS")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListLandingPagesResponseLinksItemMethod("OPTIONS"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_HEAD", func(t *testing.T) {
		t.Parallel()
		val, err := NewListLandingPagesResponseLinksItemMethodFromString("HEAD")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListLandingPagesResponseLinksItemMethod("HEAD"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewListLandingPagesResponseLinksItemMethodFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewListLandingPagesResponseLinksItemMethodFromString("GET")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestExtraPropertiesCreateLandingPagesRequestTracking(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &CreateLandingPagesRequestTracking{}
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
		var obj *CreateLandingPagesRequestTracking
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesLandingPage(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &LandingPage{}
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
		var obj *LandingPage
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesLandingPageLinksItem(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &LandingPageLinksItem{}
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
		var obj *LandingPageLinksItem
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesLandingPageTracking(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &LandingPageTracking{}
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
		var obj *LandingPageTracking
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesListContentLandingPagesResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ListContentLandingPagesResponse{}
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
		var obj *ListContentLandingPagesResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesListContentLandingPagesResponseLinksItem(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ListContentLandingPagesResponseLinksItem{}
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
		var obj *ListContentLandingPagesResponseLinksItem
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesListLandingPagesResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ListLandingPagesResponse{}
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
		var obj *ListLandingPagesResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesListLandingPagesResponseLinksItem(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ListLandingPagesResponseLinksItem{}
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
		var obj *ListLandingPagesResponseLinksItem
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesUpdateLandingPagesRequestTracking(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &UpdateLandingPagesRequestTracking{}
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
		var obj *UpdateLandingPagesRequestTracking
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}
