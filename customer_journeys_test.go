// Code generated from our API definition. DO NOT EDIT.

package api

import (
	json "encoding/json"
	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
	testing "testing"
)

func TestSettersCreateJourneyStepActionTriggerCustomerJourneysRequest(t *testing.T) {
	t.Run("SetJourneyID", func(t *testing.T) {
		obj := &CreateJourneyStepActionTriggerCustomerJourneysRequest{}
		var fernTestValueJourneyID int
		obj.SetJourneyID(fernTestValueJourneyID)
		assert.Equal(t, fernTestValueJourneyID, obj.JourneyID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStepID", func(t *testing.T) {
		obj := &CreateJourneyStepActionTriggerCustomerJourneysRequest{}
		var fernTestValueStepID int
		obj.SetStepID(fernTestValueStepID)
		assert.Equal(t, fernTestValueStepID, obj.StepID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetEmailAddress", func(t *testing.T) {
		obj := &CreateJourneyStepActionTriggerCustomerJourneysRequest{}
		var fernTestValueEmailAddress string
		obj.SetEmailAddress(fernTestValueEmailAddress)
		assert.Equal(t, fernTestValueEmailAddress, obj.EmailAddress)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitCreateJourneyStepActionTriggerCustomerJourneysRequest(t *testing.T) {
	t.Run("SetJourneyID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateJourneyStepActionTriggerCustomerJourneysRequest{}
		var fernTestValueJourneyID int

		// Act
		obj.SetJourneyID(fernTestValueJourneyID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetStepID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateJourneyStepActionTriggerCustomerJourneysRequest{}
		var fernTestValueStepID int

		// Act
		obj.SetStepID(fernTestValueStepID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetEmailAddress_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &CreateJourneyStepActionTriggerCustomerJourneysRequest{}
		var fernTestValueEmailAddress string

		// Act
		obj.SetEmailAddress(fernTestValueEmailAddress)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
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
