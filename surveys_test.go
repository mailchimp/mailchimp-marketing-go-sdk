// Code generated from our API definition. DO NOT EDIT.

package api

import (
	json "encoding/json"
	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
	testing "testing"
)

func TestSettersCreateListSurveyActionCreateEmailSurveysRequest(t *testing.T) {
	t.Run("SetListID", func(t *testing.T) {
		obj := &CreateListSurveyActionCreateEmailSurveysRequest{}
		var fernTestValueListID string
		obj.SetListID(fernTestValueListID)
		assert.Equal(t, fernTestValueListID, obj.ListID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSurveyID", func(t *testing.T) {
		obj := &CreateListSurveyActionCreateEmailSurveysRequest{}
		var fernTestValueSurveyID string
		obj.SetSurveyID(fernTestValueSurveyID)
		assert.Equal(t, fernTestValueSurveyID, obj.SurveyID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitCreateListSurveyActionCreateEmailSurveysRequest(t *testing.T) {
	t.Run("SetListID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateListSurveyActionCreateEmailSurveysRequest{}
		var fernTestValueListID string

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

	t.Run("SetSurveyID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateListSurveyActionCreateEmailSurveysRequest{}
		var fernTestValueSurveyID string

		// Act
		obj.SetSurveyID(fernTestValueSurveyID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
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

func TestSettersCreateListSurveyActionPublishSurveysRequest(t *testing.T) {
	t.Run("SetListID", func(t *testing.T) {
		obj := &CreateListSurveyActionPublishSurveysRequest{}
		var fernTestValueListID string
		obj.SetListID(fernTestValueListID)
		assert.Equal(t, fernTestValueListID, obj.ListID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSurveyID", func(t *testing.T) {
		obj := &CreateListSurveyActionPublishSurveysRequest{}
		var fernTestValueSurveyID string
		obj.SetSurveyID(fernTestValueSurveyID)
		assert.Equal(t, fernTestValueSurveyID, obj.SurveyID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitCreateListSurveyActionPublishSurveysRequest(t *testing.T) {
	t.Run("SetListID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateListSurveyActionPublishSurveysRequest{}
		var fernTestValueListID string

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

	t.Run("SetSurveyID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateListSurveyActionPublishSurveysRequest{}
		var fernTestValueSurveyID string

		// Act
		obj.SetSurveyID(fernTestValueSurveyID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
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

func TestSettersCreateListSurveyActionUnpublishSurveysRequest(t *testing.T) {
	t.Run("SetListID", func(t *testing.T) {
		obj := &CreateListSurveyActionUnpublishSurveysRequest{}
		var fernTestValueListID string
		obj.SetListID(fernTestValueListID)
		assert.Equal(t, fernTestValueListID, obj.ListID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSurveyID", func(t *testing.T) {
		obj := &CreateListSurveyActionUnpublishSurveysRequest{}
		var fernTestValueSurveyID string
		obj.SetSurveyID(fernTestValueSurveyID)
		assert.Equal(t, fernTestValueSurveyID, obj.SurveyID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitCreateListSurveyActionUnpublishSurveysRequest(t *testing.T) {
	t.Run("SetListID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateListSurveyActionUnpublishSurveysRequest{}
		var fernTestValueListID string

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

	t.Run("SetSurveyID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateListSurveyActionUnpublishSurveysRequest{}
		var fernTestValueSurveyID string

		// Act
		obj.SetSurveyID(fernTestValueSurveyID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
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
