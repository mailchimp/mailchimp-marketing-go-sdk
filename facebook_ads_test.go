// Code generated from our API definition. DO NOT EDIT.

package api

import (
	json "encoding/json"
	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
	testing "testing"
	time "time"
)

func TestSettersGetFacebookAdsRequest(t *testing.T) {
	t.Run("SetOutreachID", func(t *testing.T) {
		obj := &GetFacebookAdsRequest{}
		var fernTestValueOutreachID string
		obj.SetOutreachID(fernTestValueOutreachID)
		assert.Equal(t, fernTestValueOutreachID, obj.OutreachID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetFields", func(t *testing.T) {
		obj := &GetFacebookAdsRequest{}
		var fernTestValueFields []*string
		obj.SetFields(fernTestValueFields)
		assert.Equal(t, fernTestValueFields, obj.Fields)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetExcludeFields", func(t *testing.T) {
		obj := &GetFacebookAdsRequest{}
		var fernTestValueExcludeFields []*string
		obj.SetExcludeFields(fernTestValueExcludeFields)
		assert.Equal(t, fernTestValueExcludeFields, obj.ExcludeFields)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitGetFacebookAdsRequest(t *testing.T) {
	t.Run("SetOutreachID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &GetFacebookAdsRequest{}
		var fernTestValueOutreachID string

		// Act
		obj.SetOutreachID(fernTestValueOutreachID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
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
		obj := &GetFacebookAdsRequest{}
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
		obj := &GetFacebookAdsRequest{}
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

func TestSettersListFacebookAdsRequest(t *testing.T) {
	t.Run("SetFields", func(t *testing.T) {
		obj := &ListFacebookAdsRequest{}
		var fernTestValueFields []*string
		obj.SetFields(fernTestValueFields)
		assert.Equal(t, fernTestValueFields, obj.Fields)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetExcludeFields", func(t *testing.T) {
		obj := &ListFacebookAdsRequest{}
		var fernTestValueExcludeFields []*string
		obj.SetExcludeFields(fernTestValueExcludeFields)
		assert.Equal(t, fernTestValueExcludeFields, obj.ExcludeFields)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCount", func(t *testing.T) {
		obj := &ListFacebookAdsRequest{}
		var fernTestValueCount *int
		obj.SetCount(fernTestValueCount)
		assert.Equal(t, fernTestValueCount, obj.Count)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetOffset", func(t *testing.T) {
		obj := &ListFacebookAdsRequest{}
		var fernTestValueOffset *int
		obj.SetOffset(fernTestValueOffset)
		assert.Equal(t, fernTestValueOffset, obj.Offset)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSortField", func(t *testing.T) {
		obj := &ListFacebookAdsRequest{}
		var fernTestValueSortField *ListFacebookAdsRequestSortField
		obj.SetSortField(fernTestValueSortField)
		assert.Equal(t, fernTestValueSortField, obj.SortField)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSortDir", func(t *testing.T) {
		obj := &ListFacebookAdsRequest{}
		var fernTestValueSortDir *ListFacebookAdsRequestSortDir
		obj.SetSortDir(fernTestValueSortDir)
		assert.Equal(t, fernTestValueSortDir, obj.SortDir)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitListFacebookAdsRequest(t *testing.T) {
	t.Run("SetFields_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListFacebookAdsRequest{}
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
		obj := &ListFacebookAdsRequest{}
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
		obj := &ListFacebookAdsRequest{}
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
		obj := &ListFacebookAdsRequest{}
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

	t.Run("SetSortField_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListFacebookAdsRequest{}
		var fernTestValueSortField *ListFacebookAdsRequestSortField

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
		obj := &ListFacebookAdsRequest{}
		var fernTestValueSortDir *ListFacebookAdsRequestSortDir

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

func TestSettersFacebookAd(t *testing.T) {
	t.Run("SetCanceledAt", func(t *testing.T) {
		obj := &FacebookAd{}
		var fernTestValueCanceledAt *time.Time
		obj.SetCanceledAt(fernTestValueCanceledAt)
		assert.Equal(t, fernTestValueCanceledAt, obj.CanceledAt)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCreateTime", func(t *testing.T) {
		obj := &FacebookAd{}
		var fernTestValueCreateTime *time.Time
		obj.SetCreateTime(fernTestValueCreateTime)
		assert.Equal(t, fernTestValueCreateTime, obj.CreateTime)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetHasSegment", func(t *testing.T) {
		obj := &FacebookAd{}
		var fernTestValueHasSegment *bool
		obj.SetHasSegment(fernTestValueHasSegment)
		assert.Equal(t, fernTestValueHasSegment, obj.HasSegment)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetID", func(t *testing.T) {
		obj := &FacebookAd{}
		var fernTestValueID *string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetName", func(t *testing.T) {
		obj := &FacebookAd{}
		var fernTestValueName *string
		obj.SetName(fernTestValueName)
		assert.Equal(t, fernTestValueName, obj.Name)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetPublishedTime", func(t *testing.T) {
		obj := &FacebookAd{}
		var fernTestValuePublishedTime *time.Time
		obj.SetPublishedTime(fernTestValuePublishedTime)
		assert.Equal(t, fernTestValuePublishedTime, obj.PublishedTime)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRecipients", func(t *testing.T) {
		obj := &FacebookAd{}
		var fernTestValueRecipients *FacebookAdRecipients
		obj.SetRecipients(fernTestValueRecipients)
		assert.Equal(t, fernTestValueRecipients, obj.Recipients)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetReportSummary", func(t *testing.T) {
		obj := &FacebookAd{}
		var fernTestValueReportSummary *FacebookAdReportSummary
		obj.SetReportSummary(fernTestValueReportSummary)
		assert.Equal(t, fernTestValueReportSummary, obj.ReportSummary)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetShowReport", func(t *testing.T) {
		obj := &FacebookAd{}
		var fernTestValueShowReport *bool
		obj.SetShowReport(fernTestValueShowReport)
		assert.Equal(t, fernTestValueShowReport, obj.ShowReport)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStartTime", func(t *testing.T) {
		obj := &FacebookAd{}
		var fernTestValueStartTime *time.Time
		obj.SetStartTime(fernTestValueStartTime)
		assert.Equal(t, fernTestValueStartTime, obj.StartTime)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStatus", func(t *testing.T) {
		obj := &FacebookAd{}
		var fernTestValueStatus *FacebookAdStatus
		obj.SetStatus(fernTestValueStatus)
		assert.Equal(t, fernTestValueStatus, obj.Status)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetType", func(t *testing.T) {
		obj := &FacebookAd{}
		var fernTestValueType *FacebookAdType
		obj.SetType(fernTestValueType)
		assert.Equal(t, fernTestValueType, obj.Type)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetUpdatedAt", func(t *testing.T) {
		obj := &FacebookAd{}
		var fernTestValueUpdatedAt *time.Time
		obj.SetUpdatedAt(fernTestValueUpdatedAt)
		assert.Equal(t, fernTestValueUpdatedAt, obj.UpdatedAt)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetWebID", func(t *testing.T) {
		obj := &FacebookAd{}
		var fernTestValueWebID *int
		obj.SetWebID(fernTestValueWebID)
		assert.Equal(t, fernTestValueWebID, obj.WebID)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersFacebookAd(t *testing.T) {
	t.Run("GetCanceledAt", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		var expected *time.Time
		obj.CanceledAt = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCanceledAt(), "getter should return the property value")
	})

	t.Run("GetCanceledAt_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		obj.CanceledAt = nil

		// Act & Assert
		assert.Nil(t, obj.GetCanceledAt(), "getter should return nil when property is nil")
	})

	t.Run("GetCanceledAt_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAd
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCanceledAt() // Should return zero value
	})

	t.Run("GetCreateTime", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		var expected *time.Time
		obj.CreateTime = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCreateTime(), "getter should return the property value")
	})

	t.Run("GetCreateTime_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		obj.CreateTime = nil

		// Act & Assert
		assert.Nil(t, obj.GetCreateTime(), "getter should return nil when property is nil")
	})

	t.Run("GetCreateTime_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAd
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCreateTime() // Should return zero value
	})

	t.Run("GetHasSegment", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		var expected *bool
		obj.HasSegment = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetHasSegment(), "getter should return the property value")
	})

	t.Run("GetHasSegment_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		obj.HasSegment = nil

		// Act & Assert
		assert.Nil(t, obj.GetHasSegment(), "getter should return nil when property is nil")
	})

	t.Run("GetHasSegment_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAd
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetHasSegment() // Should return zero value
	})

	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		var expected *string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		obj.ID = nil

		// Act & Assert
		assert.Nil(t, obj.GetID(), "getter should return nil when property is nil")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAd
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetID() // Should return zero value
	})

	t.Run("GetName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		var expected *string
		obj.Name = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetName(), "getter should return the property value")
	})

	t.Run("GetName_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		obj.Name = nil

		// Act & Assert
		assert.Nil(t, obj.GetName(), "getter should return nil when property is nil")
	})

	t.Run("GetName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAd
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetName() // Should return zero value
	})

	t.Run("GetPublishedTime", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		var expected *time.Time
		obj.PublishedTime = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetPublishedTime(), "getter should return the property value")
	})

	t.Run("GetPublishedTime_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		obj.PublishedTime = nil

		// Act & Assert
		assert.Nil(t, obj.GetPublishedTime(), "getter should return nil when property is nil")
	})

	t.Run("GetPublishedTime_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAd
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetPublishedTime() // Should return zero value
	})

	t.Run("GetRecipients", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		var expected *FacebookAdRecipients
		obj.Recipients = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetRecipients(), "getter should return the property value")
	})

	t.Run("GetRecipients_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		obj.Recipients = nil

		// Act & Assert
		assert.Nil(t, obj.GetRecipients(), "getter should return nil when property is nil")
	})

	t.Run("GetRecipients_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAd
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetRecipients() // Should return zero value
	})

	t.Run("GetReportSummary", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		var expected *FacebookAdReportSummary
		obj.ReportSummary = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetReportSummary(), "getter should return the property value")
	})

	t.Run("GetReportSummary_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		obj.ReportSummary = nil

		// Act & Assert
		assert.Nil(t, obj.GetReportSummary(), "getter should return nil when property is nil")
	})

	t.Run("GetReportSummary_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAd
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetReportSummary() // Should return zero value
	})

	t.Run("GetShowReport", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		var expected *bool
		obj.ShowReport = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetShowReport(), "getter should return the property value")
	})

	t.Run("GetShowReport_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		obj.ShowReport = nil

		// Act & Assert
		assert.Nil(t, obj.GetShowReport(), "getter should return nil when property is nil")
	})

	t.Run("GetShowReport_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAd
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetShowReport() // Should return zero value
	})

	t.Run("GetStartTime", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		var expected *time.Time
		obj.StartTime = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStartTime(), "getter should return the property value")
	})

	t.Run("GetStartTime_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		obj.StartTime = nil

		// Act & Assert
		assert.Nil(t, obj.GetStartTime(), "getter should return nil when property is nil")
	})

	t.Run("GetStartTime_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAd
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStartTime() // Should return zero value
	})

	t.Run("GetStatus", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		var expected *FacebookAdStatus
		obj.Status = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStatus(), "getter should return the property value")
	})

	t.Run("GetStatus_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		obj.Status = nil

		// Act & Assert
		assert.Nil(t, obj.GetStatus(), "getter should return nil when property is nil")
	})

	t.Run("GetStatus_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAd
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStatus() // Should return zero value
	})

	t.Run("GetType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		var expected *FacebookAdType
		obj.Type = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetType(), "getter should return the property value")
	})

	t.Run("GetType_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		obj.Type = nil

		// Act & Assert
		assert.Nil(t, obj.GetType(), "getter should return nil when property is nil")
	})

	t.Run("GetType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAd
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetType() // Should return zero value
	})

	t.Run("GetUpdatedAt", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		var expected *time.Time
		obj.UpdatedAt = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetUpdatedAt(), "getter should return the property value")
	})

	t.Run("GetUpdatedAt_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		obj.UpdatedAt = nil

		// Act & Assert
		assert.Nil(t, obj.GetUpdatedAt(), "getter should return nil when property is nil")
	})

	t.Run("GetUpdatedAt_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAd
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetUpdatedAt() // Should return zero value
	})

	t.Run("GetWebID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		var expected *int
		obj.WebID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetWebID(), "getter should return the property value")
	})

	t.Run("GetWebID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		obj.WebID = nil

		// Act & Assert
		assert.Nil(t, obj.GetWebID(), "getter should return nil when property is nil")
	})

	t.Run("GetWebID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAd
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetWebID() // Should return zero value
	})

}

func TestSettersMarkExplicitFacebookAd(t *testing.T) {
	t.Run("SetCanceledAt_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		var fernTestValueCanceledAt *time.Time

		// Act
		obj.SetCanceledAt(fernTestValueCanceledAt)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetCreateTime_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		var fernTestValueCreateTime *time.Time

		// Act
		obj.SetCreateTime(fernTestValueCreateTime)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetHasSegment_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		var fernTestValueHasSegment *bool

		// Act
		obj.SetHasSegment(fernTestValueHasSegment)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
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
		obj := &FacebookAd{}
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

	t.Run("SetName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
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

	t.Run("SetPublishedTime_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		var fernTestValuePublishedTime *time.Time

		// Act
		obj.SetPublishedTime(fernTestValuePublishedTime)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetRecipients_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		var fernTestValueRecipients *FacebookAdRecipients

		// Act
		obj.SetRecipients(fernTestValueRecipients)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetReportSummary_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		var fernTestValueReportSummary *FacebookAdReportSummary

		// Act
		obj.SetReportSummary(fernTestValueReportSummary)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetShowReport_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		var fernTestValueShowReport *bool

		// Act
		obj.SetShowReport(fernTestValueShowReport)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetStartTime_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		var fernTestValueStartTime *time.Time

		// Act
		obj.SetStartTime(fernTestValueStartTime)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
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
		obj := &FacebookAd{}
		var fernTestValueStatus *FacebookAdStatus

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

	t.Run("SetType_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
		var fernTestValueType *FacebookAdType

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

	t.Run("SetUpdatedAt_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
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

	t.Run("SetWebID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}
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

func TestSettersFacebookAdReportSummary(t *testing.T) {
	t.Run("SetClickRate", func(t *testing.T) {
		obj := &FacebookAdReportSummary{}
		var fernTestValueClickRate *float64
		obj.SetClickRate(fernTestValueClickRate)
		assert.Equal(t, fernTestValueClickRate, obj.ClickRate)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetClicks", func(t *testing.T) {
		obj := &FacebookAdReportSummary{}
		var fernTestValueClicks *int
		obj.SetClicks(fernTestValueClicks)
		assert.Equal(t, fernTestValueClicks, obj.Clicks)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetConversionRate", func(t *testing.T) {
		obj := &FacebookAdReportSummary{}
		var fernTestValueConversionRate *float64
		obj.SetConversionRate(fernTestValueConversionRate)
		assert.Equal(t, fernTestValueConversionRate, obj.ConversionRate)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetEcommerce", func(t *testing.T) {
		obj := &FacebookAdReportSummary{}
		var fernTestValueEcommerce *FacebookAdReportSummaryEcommerce
		obj.SetEcommerce(fernTestValueEcommerce)
		assert.Equal(t, fernTestValueEcommerce, obj.Ecommerce)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetEngagements", func(t *testing.T) {
		obj := &FacebookAdReportSummary{}
		var fernTestValueEngagements *int
		obj.SetEngagements(fernTestValueEngagements)
		assert.Equal(t, fernTestValueEngagements, obj.Engagements)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetImpressions", func(t *testing.T) {
		obj := &FacebookAdReportSummary{}
		var fernTestValueImpressions *float64
		obj.SetImpressions(fernTestValueImpressions)
		assert.Equal(t, fernTestValueImpressions, obj.Impressions)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetOpenRate", func(t *testing.T) {
		obj := &FacebookAdReportSummary{}
		var fernTestValueOpenRate *float64
		obj.SetOpenRate(fernTestValueOpenRate)
		assert.Equal(t, fernTestValueOpenRate, obj.OpenRate)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetOpens", func(t *testing.T) {
		obj := &FacebookAdReportSummary{}
		var fernTestValueOpens *int
		obj.SetOpens(fernTestValueOpens)
		assert.Equal(t, fernTestValueOpens, obj.Opens)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetProxyExcludedOpenRate", func(t *testing.T) {
		obj := &FacebookAdReportSummary{}
		var fernTestValueProxyExcludedOpenRate *float64
		obj.SetProxyExcludedOpenRate(fernTestValueProxyExcludedOpenRate)
		assert.Equal(t, fernTestValueProxyExcludedOpenRate, obj.ProxyExcludedOpenRate)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetProxyExcludedOpens", func(t *testing.T) {
		obj := &FacebookAdReportSummary{}
		var fernTestValueProxyExcludedOpens *int
		obj.SetProxyExcludedOpens(fernTestValueProxyExcludedOpens)
		assert.Equal(t, fernTestValueProxyExcludedOpens, obj.ProxyExcludedOpens)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetProxyExcludedUniqueOpens", func(t *testing.T) {
		obj := &FacebookAdReportSummary{}
		var fernTestValueProxyExcludedUniqueOpens *int
		obj.SetProxyExcludedUniqueOpens(fernTestValueProxyExcludedUniqueOpens)
		assert.Equal(t, fernTestValueProxyExcludedUniqueOpens, obj.ProxyExcludedUniqueOpens)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetReach", func(t *testing.T) {
		obj := &FacebookAdReportSummary{}
		var fernTestValueReach *int
		obj.SetReach(fernTestValueReach)
		assert.Equal(t, fernTestValueReach, obj.Reach)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSubscriberClicks", func(t *testing.T) {
		obj := &FacebookAdReportSummary{}
		var fernTestValueSubscriberClicks *int
		obj.SetSubscriberClicks(fernTestValueSubscriberClicks)
		assert.Equal(t, fernTestValueSubscriberClicks, obj.SubscriberClicks)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSubscribes", func(t *testing.T) {
		obj := &FacebookAdReportSummary{}
		var fernTestValueSubscribes *int
		obj.SetSubscribes(fernTestValueSubscribes)
		assert.Equal(t, fernTestValueSubscribes, obj.Subscribes)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTotalSent", func(t *testing.T) {
		obj := &FacebookAdReportSummary{}
		var fernTestValueTotalSent *int
		obj.SetTotalSent(fernTestValueTotalSent)
		assert.Equal(t, fernTestValueTotalSent, obj.TotalSent)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetUniqueOpens", func(t *testing.T) {
		obj := &FacebookAdReportSummary{}
		var fernTestValueUniqueOpens *int
		obj.SetUniqueOpens(fernTestValueUniqueOpens)
		assert.Equal(t, fernTestValueUniqueOpens, obj.UniqueOpens)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetUniqueVisits", func(t *testing.T) {
		obj := &FacebookAdReportSummary{}
		var fernTestValueUniqueVisits *int
		obj.SetUniqueVisits(fernTestValueUniqueVisits)
		assert.Equal(t, fernTestValueUniqueVisits, obj.UniqueVisits)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetVisits", func(t *testing.T) {
		obj := &FacebookAdReportSummary{}
		var fernTestValueVisits *int
		obj.SetVisits(fernTestValueVisits)
		assert.Equal(t, fernTestValueVisits, obj.Visits)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersFacebookAdReportSummary(t *testing.T) {
	t.Run("GetClickRate", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var expected *float64
		obj.ClickRate = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetClickRate(), "getter should return the property value")
	})

	t.Run("GetClickRate_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		obj.ClickRate = nil

		// Act & Assert
		assert.Nil(t, obj.GetClickRate(), "getter should return nil when property is nil")
	})

	t.Run("GetClickRate_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdReportSummary
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetClickRate() // Should return zero value
	})

	t.Run("GetClicks", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var expected *int
		obj.Clicks = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetClicks(), "getter should return the property value")
	})

	t.Run("GetClicks_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		obj.Clicks = nil

		// Act & Assert
		assert.Nil(t, obj.GetClicks(), "getter should return nil when property is nil")
	})

	t.Run("GetClicks_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdReportSummary
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetClicks() // Should return zero value
	})

	t.Run("GetConversionRate", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var expected *float64
		obj.ConversionRate = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetConversionRate(), "getter should return the property value")
	})

	t.Run("GetConversionRate_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		obj.ConversionRate = nil

		// Act & Assert
		assert.Nil(t, obj.GetConversionRate(), "getter should return nil when property is nil")
	})

	t.Run("GetConversionRate_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdReportSummary
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetConversionRate() // Should return zero value
	})

	t.Run("GetEcommerce", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var expected *FacebookAdReportSummaryEcommerce
		obj.Ecommerce = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEcommerce(), "getter should return the property value")
	})

	t.Run("GetEcommerce_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		obj.Ecommerce = nil

		// Act & Assert
		assert.Nil(t, obj.GetEcommerce(), "getter should return nil when property is nil")
	})

	t.Run("GetEcommerce_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdReportSummary
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEcommerce() // Should return zero value
	})

	t.Run("GetEngagements", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var expected *int
		obj.Engagements = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEngagements(), "getter should return the property value")
	})

	t.Run("GetEngagements_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		obj.Engagements = nil

		// Act & Assert
		assert.Nil(t, obj.GetEngagements(), "getter should return nil when property is nil")
	})

	t.Run("GetEngagements_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdReportSummary
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEngagements() // Should return zero value
	})

	t.Run("GetImpressions", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var expected *float64
		obj.Impressions = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetImpressions(), "getter should return the property value")
	})

	t.Run("GetImpressions_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		obj.Impressions = nil

		// Act & Assert
		assert.Nil(t, obj.GetImpressions(), "getter should return nil when property is nil")
	})

	t.Run("GetImpressions_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdReportSummary
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetImpressions() // Should return zero value
	})

	t.Run("GetOpenRate", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var expected *float64
		obj.OpenRate = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetOpenRate(), "getter should return the property value")
	})

	t.Run("GetOpenRate_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		obj.OpenRate = nil

		// Act & Assert
		assert.Nil(t, obj.GetOpenRate(), "getter should return nil when property is nil")
	})

	t.Run("GetOpenRate_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdReportSummary
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetOpenRate() // Should return zero value
	})

	t.Run("GetOpens", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var expected *int
		obj.Opens = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetOpens(), "getter should return the property value")
	})

	t.Run("GetOpens_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		obj.Opens = nil

		// Act & Assert
		assert.Nil(t, obj.GetOpens(), "getter should return nil when property is nil")
	})

	t.Run("GetOpens_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdReportSummary
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetOpens() // Should return zero value
	})

	t.Run("GetProxyExcludedOpenRate", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var expected *float64
		obj.ProxyExcludedOpenRate = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetProxyExcludedOpenRate(), "getter should return the property value")
	})

	t.Run("GetProxyExcludedOpenRate_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		obj.ProxyExcludedOpenRate = nil

		// Act & Assert
		assert.Nil(t, obj.GetProxyExcludedOpenRate(), "getter should return nil when property is nil")
	})

	t.Run("GetProxyExcludedOpenRate_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdReportSummary
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetProxyExcludedOpenRate() // Should return zero value
	})

	t.Run("GetProxyExcludedOpens", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var expected *int
		obj.ProxyExcludedOpens = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetProxyExcludedOpens(), "getter should return the property value")
	})

	t.Run("GetProxyExcludedOpens_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		obj.ProxyExcludedOpens = nil

		// Act & Assert
		assert.Nil(t, obj.GetProxyExcludedOpens(), "getter should return nil when property is nil")
	})

	t.Run("GetProxyExcludedOpens_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdReportSummary
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetProxyExcludedOpens() // Should return zero value
	})

	t.Run("GetProxyExcludedUniqueOpens", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var expected *int
		obj.ProxyExcludedUniqueOpens = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetProxyExcludedUniqueOpens(), "getter should return the property value")
	})

	t.Run("GetProxyExcludedUniqueOpens_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		obj.ProxyExcludedUniqueOpens = nil

		// Act & Assert
		assert.Nil(t, obj.GetProxyExcludedUniqueOpens(), "getter should return nil when property is nil")
	})

	t.Run("GetProxyExcludedUniqueOpens_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdReportSummary
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetProxyExcludedUniqueOpens() // Should return zero value
	})

	t.Run("GetReach", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var expected *int
		obj.Reach = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetReach(), "getter should return the property value")
	})

	t.Run("GetReach_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		obj.Reach = nil

		// Act & Assert
		assert.Nil(t, obj.GetReach(), "getter should return nil when property is nil")
	})

	t.Run("GetReach_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdReportSummary
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetReach() // Should return zero value
	})

	t.Run("GetSubscriberClicks", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var expected *int
		obj.SubscriberClicks = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSubscriberClicks(), "getter should return the property value")
	})

	t.Run("GetSubscriberClicks_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		obj.SubscriberClicks = nil

		// Act & Assert
		assert.Nil(t, obj.GetSubscriberClicks(), "getter should return nil when property is nil")
	})

	t.Run("GetSubscriberClicks_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdReportSummary
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSubscriberClicks() // Should return zero value
	})

	t.Run("GetSubscribes", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var expected *int
		obj.Subscribes = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSubscribes(), "getter should return the property value")
	})

	t.Run("GetSubscribes_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		obj.Subscribes = nil

		// Act & Assert
		assert.Nil(t, obj.GetSubscribes(), "getter should return nil when property is nil")
	})

	t.Run("GetSubscribes_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdReportSummary
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSubscribes() // Should return zero value
	})

	t.Run("GetTotalSent", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var expected *int
		obj.TotalSent = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTotalSent(), "getter should return the property value")
	})

	t.Run("GetTotalSent_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		obj.TotalSent = nil

		// Act & Assert
		assert.Nil(t, obj.GetTotalSent(), "getter should return nil when property is nil")
	})

	t.Run("GetTotalSent_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdReportSummary
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTotalSent() // Should return zero value
	})

	t.Run("GetUniqueOpens", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var expected *int
		obj.UniqueOpens = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetUniqueOpens(), "getter should return the property value")
	})

	t.Run("GetUniqueOpens_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		obj.UniqueOpens = nil

		// Act & Assert
		assert.Nil(t, obj.GetUniqueOpens(), "getter should return nil when property is nil")
	})

	t.Run("GetUniqueOpens_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdReportSummary
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetUniqueOpens() // Should return zero value
	})

	t.Run("GetUniqueVisits", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var expected *int
		obj.UniqueVisits = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetUniqueVisits(), "getter should return the property value")
	})

	t.Run("GetUniqueVisits_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		obj.UniqueVisits = nil

		// Act & Assert
		assert.Nil(t, obj.GetUniqueVisits(), "getter should return nil when property is nil")
	})

	t.Run("GetUniqueVisits_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdReportSummary
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetUniqueVisits() // Should return zero value
	})

	t.Run("GetVisits", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var expected *int
		obj.Visits = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetVisits(), "getter should return the property value")
	})

	t.Run("GetVisits_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		obj.Visits = nil

		// Act & Assert
		assert.Nil(t, obj.GetVisits(), "getter should return nil when property is nil")
	})

	t.Run("GetVisits_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdReportSummary
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetVisits() // Should return zero value
	})

}

func TestSettersMarkExplicitFacebookAdReportSummary(t *testing.T) {
	t.Run("SetClickRate_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var fernTestValueClickRate *float64

		// Act
		obj.SetClickRate(fernTestValueClickRate)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetClicks_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var fernTestValueClicks *int

		// Act
		obj.SetClicks(fernTestValueClicks)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetConversionRate_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var fernTestValueConversionRate *float64

		// Act
		obj.SetConversionRate(fernTestValueConversionRate)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetEcommerce_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var fernTestValueEcommerce *FacebookAdReportSummaryEcommerce

		// Act
		obj.SetEcommerce(fernTestValueEcommerce)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetEngagements_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var fernTestValueEngagements *int

		// Act
		obj.SetEngagements(fernTestValueEngagements)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetImpressions_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var fernTestValueImpressions *float64

		// Act
		obj.SetImpressions(fernTestValueImpressions)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetOpenRate_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var fernTestValueOpenRate *float64

		// Act
		obj.SetOpenRate(fernTestValueOpenRate)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetOpens_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var fernTestValueOpens *int

		// Act
		obj.SetOpens(fernTestValueOpens)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetProxyExcludedOpenRate_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var fernTestValueProxyExcludedOpenRate *float64

		// Act
		obj.SetProxyExcludedOpenRate(fernTestValueProxyExcludedOpenRate)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetProxyExcludedOpens_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var fernTestValueProxyExcludedOpens *int

		// Act
		obj.SetProxyExcludedOpens(fernTestValueProxyExcludedOpens)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetProxyExcludedUniqueOpens_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var fernTestValueProxyExcludedUniqueOpens *int

		// Act
		obj.SetProxyExcludedUniqueOpens(fernTestValueProxyExcludedUniqueOpens)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetReach_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var fernTestValueReach *int

		// Act
		obj.SetReach(fernTestValueReach)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetSubscriberClicks_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var fernTestValueSubscriberClicks *int

		// Act
		obj.SetSubscriberClicks(fernTestValueSubscriberClicks)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetSubscribes_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var fernTestValueSubscribes *int

		// Act
		obj.SetSubscribes(fernTestValueSubscribes)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetTotalSent_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var fernTestValueTotalSent *int

		// Act
		obj.SetTotalSent(fernTestValueTotalSent)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetUniqueOpens_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var fernTestValueUniqueOpens *int

		// Act
		obj.SetUniqueOpens(fernTestValueUniqueOpens)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetUniqueVisits_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var fernTestValueUniqueVisits *int

		// Act
		obj.SetUniqueVisits(fernTestValueUniqueVisits)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetVisits_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}
		var fernTestValueVisits *int

		// Act
		obj.SetVisits(fernTestValueVisits)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
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

func TestSettersFacebookAdReportSummaryEcommerce(t *testing.T) {
	t.Run("SetAverageOrderRevenue", func(t *testing.T) {
		obj := &FacebookAdReportSummaryEcommerce{}
		var fernTestValueAverageOrderRevenue *float64
		obj.SetAverageOrderRevenue(fernTestValueAverageOrderRevenue)
		assert.Equal(t, fernTestValueAverageOrderRevenue, obj.AverageOrderRevenue)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCurrencyCode", func(t *testing.T) {
		obj := &FacebookAdReportSummaryEcommerce{}
		var fernTestValueCurrencyCode *string
		obj.SetCurrencyCode(fernTestValueCurrencyCode)
		assert.Equal(t, fernTestValueCurrencyCode, obj.CurrencyCode)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTotalRevenue", func(t *testing.T) {
		obj := &FacebookAdReportSummaryEcommerce{}
		var fernTestValueTotalRevenue *float64
		obj.SetTotalRevenue(fernTestValueTotalRevenue)
		assert.Equal(t, fernTestValueTotalRevenue, obj.TotalRevenue)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersFacebookAdReportSummaryEcommerce(t *testing.T) {
	t.Run("GetAverageOrderRevenue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummaryEcommerce{}
		var expected *float64
		obj.AverageOrderRevenue = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetAverageOrderRevenue(), "getter should return the property value")
	})

	t.Run("GetAverageOrderRevenue_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummaryEcommerce{}
		obj.AverageOrderRevenue = nil

		// Act & Assert
		assert.Nil(t, obj.GetAverageOrderRevenue(), "getter should return nil when property is nil")
	})

	t.Run("GetAverageOrderRevenue_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdReportSummaryEcommerce
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetAverageOrderRevenue() // Should return zero value
	})

	t.Run("GetCurrencyCode", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummaryEcommerce{}
		var expected *string
		obj.CurrencyCode = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCurrencyCode(), "getter should return the property value")
	})

	t.Run("GetCurrencyCode_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummaryEcommerce{}
		obj.CurrencyCode = nil

		// Act & Assert
		assert.Nil(t, obj.GetCurrencyCode(), "getter should return nil when property is nil")
	})

	t.Run("GetCurrencyCode_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdReportSummaryEcommerce
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCurrencyCode() // Should return zero value
	})

	t.Run("GetTotalRevenue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummaryEcommerce{}
		var expected *float64
		obj.TotalRevenue = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTotalRevenue(), "getter should return the property value")
	})

	t.Run("GetTotalRevenue_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummaryEcommerce{}
		obj.TotalRevenue = nil

		// Act & Assert
		assert.Nil(t, obj.GetTotalRevenue(), "getter should return nil when property is nil")
	})

	t.Run("GetTotalRevenue_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdReportSummaryEcommerce
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTotalRevenue() // Should return zero value
	})

}

func TestSettersMarkExplicitFacebookAdReportSummaryEcommerce(t *testing.T) {
	t.Run("SetAverageOrderRevenue_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummaryEcommerce{}
		var fernTestValueAverageOrderRevenue *float64

		// Act
		obj.SetAverageOrderRevenue(fernTestValueAverageOrderRevenue)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetCurrencyCode_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummaryEcommerce{}
		var fernTestValueCurrencyCode *string

		// Act
		obj.SetCurrencyCode(fernTestValueCurrencyCode)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetTotalRevenue_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummaryEcommerce{}
		var fernTestValueTotalRevenue *float64

		// Act
		obj.SetTotalRevenue(fernTestValueTotalRevenue)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
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

func TestSettersFacebookAds(t *testing.T) {
	t.Run("SetCanceledAt", func(t *testing.T) {
		obj := &FacebookAds{}
		var fernTestValueCanceledAt *time.Time
		obj.SetCanceledAt(fernTestValueCanceledAt)
		assert.Equal(t, fernTestValueCanceledAt, obj.CanceledAt)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCreateTime", func(t *testing.T) {
		obj := &FacebookAds{}
		var fernTestValueCreateTime *time.Time
		obj.SetCreateTime(fernTestValueCreateTime)
		assert.Equal(t, fernTestValueCreateTime, obj.CreateTime)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetHasSegment", func(t *testing.T) {
		obj := &FacebookAds{}
		var fernTestValueHasSegment *bool
		obj.SetHasSegment(fernTestValueHasSegment)
		assert.Equal(t, fernTestValueHasSegment, obj.HasSegment)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetID", func(t *testing.T) {
		obj := &FacebookAds{}
		var fernTestValueID *string
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetName", func(t *testing.T) {
		obj := &FacebookAds{}
		var fernTestValueName *string
		obj.SetName(fernTestValueName)
		assert.Equal(t, fernTestValueName, obj.Name)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetPublishedTime", func(t *testing.T) {
		obj := &FacebookAds{}
		var fernTestValuePublishedTime *time.Time
		obj.SetPublishedTime(fernTestValuePublishedTime)
		assert.Equal(t, fernTestValuePublishedTime, obj.PublishedTime)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRecipients", func(t *testing.T) {
		obj := &FacebookAds{}
		var fernTestValueRecipients *FacebookAdRecipients
		obj.SetRecipients(fernTestValueRecipients)
		assert.Equal(t, fernTestValueRecipients, obj.Recipients)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetReportSummary", func(t *testing.T) {
		obj := &FacebookAds{}
		var fernTestValueReportSummary *FacebookAdReportSummary
		obj.SetReportSummary(fernTestValueReportSummary)
		assert.Equal(t, fernTestValueReportSummary, obj.ReportSummary)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetShowReport", func(t *testing.T) {
		obj := &FacebookAds{}
		var fernTestValueShowReport *bool
		obj.SetShowReport(fernTestValueShowReport)
		assert.Equal(t, fernTestValueShowReport, obj.ShowReport)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStartTime", func(t *testing.T) {
		obj := &FacebookAds{}
		var fernTestValueStartTime *time.Time
		obj.SetStartTime(fernTestValueStartTime)
		assert.Equal(t, fernTestValueStartTime, obj.StartTime)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetStatus", func(t *testing.T) {
		obj := &FacebookAds{}
		var fernTestValueStatus *FacebookAdStatus
		obj.SetStatus(fernTestValueStatus)
		assert.Equal(t, fernTestValueStatus, obj.Status)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetType", func(t *testing.T) {
		obj := &FacebookAds{}
		var fernTestValueType *FacebookAdType
		obj.SetType(fernTestValueType)
		assert.Equal(t, fernTestValueType, obj.Type)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetUpdatedAt", func(t *testing.T) {
		obj := &FacebookAds{}
		var fernTestValueUpdatedAt *time.Time
		obj.SetUpdatedAt(fernTestValueUpdatedAt)
		assert.Equal(t, fernTestValueUpdatedAt, obj.UpdatedAt)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetWebID", func(t *testing.T) {
		obj := &FacebookAds{}
		var fernTestValueWebID *int
		obj.SetWebID(fernTestValueWebID)
		assert.Equal(t, fernTestValueWebID, obj.WebID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetEmailSourceName", func(t *testing.T) {
		obj := &FacebookAds{}
		var fernTestValueEmailSourceName *string
		obj.SetEmailSourceName(fernTestValueEmailSourceName)
		assert.Equal(t, fernTestValueEmailSourceName, obj.EmailSourceName)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetEndTime", func(t *testing.T) {
		obj := &FacebookAds{}
		var fernTestValueEndTime *time.Time
		obj.SetEndTime(fernTestValueEndTime)
		assert.Equal(t, fernTestValueEndTime, obj.EndTime)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetNeedsAttention", func(t *testing.T) {
		obj := &FacebookAds{}
		var fernTestValueNeedsAttention *bool
		obj.SetNeedsAttention(fernTestValueNeedsAttention)
		assert.Equal(t, fernTestValueNeedsAttention, obj.NeedsAttention)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetPausedAt", func(t *testing.T) {
		obj := &FacebookAds{}
		var fernTestValuePausedAt *time.Time
		obj.SetPausedAt(fernTestValuePausedAt)
		assert.Equal(t, fernTestValuePausedAt, obj.PausedAt)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetThumbnail", func(t *testing.T) {
		obj := &FacebookAds{}
		var fernTestValueThumbnail *string
		obj.SetThumbnail(fernTestValueThumbnail)
		assert.Equal(t, fernTestValueThumbnail, obj.Thumbnail)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetWasCanceledByFacebook", func(t *testing.T) {
		obj := &FacebookAds{}
		var fernTestValueWasCanceledByFacebook *bool
		obj.SetWasCanceledByFacebook(fernTestValueWasCanceledByFacebook)
		assert.Equal(t, fernTestValueWasCanceledByFacebook, obj.WasCanceledByFacebook)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetAudience", func(t *testing.T) {
		obj := &FacebookAds{}
		var fernTestValueAudience *FacebookAdsAudience
		obj.SetAudience(fernTestValueAudience)
		assert.Equal(t, fernTestValueAudience, obj.Audience)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetBudget", func(t *testing.T) {
		obj := &FacebookAds{}
		var fernTestValueBudget *FacebookAdsBudget
		obj.SetBudget(fernTestValueBudget)
		assert.Equal(t, fernTestValueBudget, obj.Budget)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetChannel", func(t *testing.T) {
		obj := &FacebookAds{}
		var fernTestValueChannel *FacebookAdsChannel
		obj.SetChannel(fernTestValueChannel)
		assert.Equal(t, fernTestValueChannel, obj.Channel)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetContent", func(t *testing.T) {
		obj := &FacebookAds{}
		var fernTestValueContent *FacebookAdsContent
		obj.SetContent(fernTestValueContent)
		assert.Equal(t, fernTestValueContent, obj.Content)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetFeedback", func(t *testing.T) {
		obj := &FacebookAds{}
		var fernTestValueFeedback *FacebookAdsFeedback
		obj.SetFeedback(fernTestValueFeedback)
		assert.Equal(t, fernTestValueFeedback, obj.Feedback)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetHasAudience", func(t *testing.T) {
		obj := &FacebookAds{}
		var fernTestValueHasAudience *bool
		obj.SetHasAudience(fernTestValueHasAudience)
		assert.Equal(t, fernTestValueHasAudience, obj.HasAudience)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetHasContent", func(t *testing.T) {
		obj := &FacebookAds{}
		var fernTestValueHasContent *bool
		obj.SetHasContent(fernTestValueHasContent)
		assert.Equal(t, fernTestValueHasContent, obj.HasContent)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetIsConnected", func(t *testing.T) {
		obj := &FacebookAds{}
		var fernTestValueIsConnected *bool
		obj.SetIsConnected(fernTestValueIsConnected)
		assert.Equal(t, fernTestValueIsConnected, obj.IsConnected)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSite", func(t *testing.T) {
		obj := &FacebookAds{}
		var fernTestValueSite *FacebookAdsSite
		obj.SetSite(fernTestValueSite)
		assert.Equal(t, fernTestValueSite, obj.Site)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetLinks", func(t *testing.T) {
		obj := &FacebookAds{}
		var fernTestValueLinks []*FacebookAdsLinksItem
		obj.SetLinks(fernTestValueLinks)
		assert.Equal(t, fernTestValueLinks, obj.Links)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersFacebookAds(t *testing.T) {
	t.Run("GetCanceledAt", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var expected *time.Time
		obj.CanceledAt = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCanceledAt(), "getter should return the property value")
	})

	t.Run("GetCanceledAt_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		obj.CanceledAt = nil

		// Act & Assert
		assert.Nil(t, obj.GetCanceledAt(), "getter should return nil when property is nil")
	})

	t.Run("GetCanceledAt_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCanceledAt() // Should return zero value
	})

	t.Run("GetCreateTime", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var expected *time.Time
		obj.CreateTime = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCreateTime(), "getter should return the property value")
	})

	t.Run("GetCreateTime_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		obj.CreateTime = nil

		// Act & Assert
		assert.Nil(t, obj.GetCreateTime(), "getter should return nil when property is nil")
	})

	t.Run("GetCreateTime_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCreateTime() // Should return zero value
	})

	t.Run("GetHasSegment", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var expected *bool
		obj.HasSegment = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetHasSegment(), "getter should return the property value")
	})

	t.Run("GetHasSegment_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		obj.HasSegment = nil

		// Act & Assert
		assert.Nil(t, obj.GetHasSegment(), "getter should return nil when property is nil")
	})

	t.Run("GetHasSegment_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetHasSegment() // Should return zero value
	})

	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var expected *string
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		obj.ID = nil

		// Act & Assert
		assert.Nil(t, obj.GetID(), "getter should return nil when property is nil")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetID() // Should return zero value
	})

	t.Run("GetName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var expected *string
		obj.Name = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetName(), "getter should return the property value")
	})

	t.Run("GetName_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		obj.Name = nil

		// Act & Assert
		assert.Nil(t, obj.GetName(), "getter should return nil when property is nil")
	})

	t.Run("GetName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetName() // Should return zero value
	})

	t.Run("GetPublishedTime", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var expected *time.Time
		obj.PublishedTime = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetPublishedTime(), "getter should return the property value")
	})

	t.Run("GetPublishedTime_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		obj.PublishedTime = nil

		// Act & Assert
		assert.Nil(t, obj.GetPublishedTime(), "getter should return nil when property is nil")
	})

	t.Run("GetPublishedTime_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetPublishedTime() // Should return zero value
	})

	t.Run("GetRecipients", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var expected *FacebookAdRecipients
		obj.Recipients = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetRecipients(), "getter should return the property value")
	})

	t.Run("GetRecipients_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		obj.Recipients = nil

		// Act & Assert
		assert.Nil(t, obj.GetRecipients(), "getter should return nil when property is nil")
	})

	t.Run("GetRecipients_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetRecipients() // Should return zero value
	})

	t.Run("GetReportSummary", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var expected *FacebookAdReportSummary
		obj.ReportSummary = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetReportSummary(), "getter should return the property value")
	})

	t.Run("GetReportSummary_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		obj.ReportSummary = nil

		// Act & Assert
		assert.Nil(t, obj.GetReportSummary(), "getter should return nil when property is nil")
	})

	t.Run("GetReportSummary_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetReportSummary() // Should return zero value
	})

	t.Run("GetShowReport", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var expected *bool
		obj.ShowReport = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetShowReport(), "getter should return the property value")
	})

	t.Run("GetShowReport_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		obj.ShowReport = nil

		// Act & Assert
		assert.Nil(t, obj.GetShowReport(), "getter should return nil when property is nil")
	})

	t.Run("GetShowReport_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetShowReport() // Should return zero value
	})

	t.Run("GetStartTime", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var expected *time.Time
		obj.StartTime = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStartTime(), "getter should return the property value")
	})

	t.Run("GetStartTime_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		obj.StartTime = nil

		// Act & Assert
		assert.Nil(t, obj.GetStartTime(), "getter should return nil when property is nil")
	})

	t.Run("GetStartTime_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStartTime() // Should return zero value
	})

	t.Run("GetStatus", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var expected *FacebookAdStatus
		obj.Status = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetStatus(), "getter should return the property value")
	})

	t.Run("GetStatus_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		obj.Status = nil

		// Act & Assert
		assert.Nil(t, obj.GetStatus(), "getter should return nil when property is nil")
	})

	t.Run("GetStatus_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetStatus() // Should return zero value
	})

	t.Run("GetType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var expected *FacebookAdType
		obj.Type = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetType(), "getter should return the property value")
	})

	t.Run("GetType_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		obj.Type = nil

		// Act & Assert
		assert.Nil(t, obj.GetType(), "getter should return nil when property is nil")
	})

	t.Run("GetType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetType() // Should return zero value
	})

	t.Run("GetUpdatedAt", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var expected *time.Time
		obj.UpdatedAt = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetUpdatedAt(), "getter should return the property value")
	})

	t.Run("GetUpdatedAt_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		obj.UpdatedAt = nil

		// Act & Assert
		assert.Nil(t, obj.GetUpdatedAt(), "getter should return nil when property is nil")
	})

	t.Run("GetUpdatedAt_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetUpdatedAt() // Should return zero value
	})

	t.Run("GetWebID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var expected *int
		obj.WebID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetWebID(), "getter should return the property value")
	})

	t.Run("GetWebID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		obj.WebID = nil

		// Act & Assert
		assert.Nil(t, obj.GetWebID(), "getter should return nil when property is nil")
	})

	t.Run("GetWebID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetWebID() // Should return zero value
	})

	t.Run("GetEmailSourceName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var expected *string
		obj.EmailSourceName = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEmailSourceName(), "getter should return the property value")
	})

	t.Run("GetEmailSourceName_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		obj.EmailSourceName = nil

		// Act & Assert
		assert.Nil(t, obj.GetEmailSourceName(), "getter should return nil when property is nil")
	})

	t.Run("GetEmailSourceName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEmailSourceName() // Should return zero value
	})

	t.Run("GetEndTime", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var expected *time.Time
		obj.EndTime = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEndTime(), "getter should return the property value")
	})

	t.Run("GetEndTime_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		obj.EndTime = nil

		// Act & Assert
		assert.Nil(t, obj.GetEndTime(), "getter should return nil when property is nil")
	})

	t.Run("GetEndTime_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEndTime() // Should return zero value
	})

	t.Run("GetNeedsAttention", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var expected *bool
		obj.NeedsAttention = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetNeedsAttention(), "getter should return the property value")
	})

	t.Run("GetNeedsAttention_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		obj.NeedsAttention = nil

		// Act & Assert
		assert.Nil(t, obj.GetNeedsAttention(), "getter should return nil when property is nil")
	})

	t.Run("GetNeedsAttention_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetNeedsAttention() // Should return zero value
	})

	t.Run("GetPausedAt", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var expected *time.Time
		obj.PausedAt = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetPausedAt(), "getter should return the property value")
	})

	t.Run("GetPausedAt_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		obj.PausedAt = nil

		// Act & Assert
		assert.Nil(t, obj.GetPausedAt(), "getter should return nil when property is nil")
	})

	t.Run("GetPausedAt_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetPausedAt() // Should return zero value
	})

	t.Run("GetThumbnail", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var expected *string
		obj.Thumbnail = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetThumbnail(), "getter should return the property value")
	})

	t.Run("GetThumbnail_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		obj.Thumbnail = nil

		// Act & Assert
		assert.Nil(t, obj.GetThumbnail(), "getter should return nil when property is nil")
	})

	t.Run("GetThumbnail_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetThumbnail() // Should return zero value
	})

	t.Run("GetWasCanceledByFacebook", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var expected *bool
		obj.WasCanceledByFacebook = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetWasCanceledByFacebook(), "getter should return the property value")
	})

	t.Run("GetWasCanceledByFacebook_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		obj.WasCanceledByFacebook = nil

		// Act & Assert
		assert.Nil(t, obj.GetWasCanceledByFacebook(), "getter should return nil when property is nil")
	})

	t.Run("GetWasCanceledByFacebook_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetWasCanceledByFacebook() // Should return zero value
	})

	t.Run("GetAudience", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var expected *FacebookAdsAudience
		obj.Audience = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetAudience(), "getter should return the property value")
	})

	t.Run("GetAudience_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		obj.Audience = nil

		// Act & Assert
		assert.Nil(t, obj.GetAudience(), "getter should return nil when property is nil")
	})

	t.Run("GetAudience_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetAudience() // Should return zero value
	})

	t.Run("GetBudget", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var expected *FacebookAdsBudget
		obj.Budget = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetBudget(), "getter should return the property value")
	})

	t.Run("GetBudget_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		obj.Budget = nil

		// Act & Assert
		assert.Nil(t, obj.GetBudget(), "getter should return nil when property is nil")
	})

	t.Run("GetBudget_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetBudget() // Should return zero value
	})

	t.Run("GetChannel", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var expected *FacebookAdsChannel
		obj.Channel = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetChannel(), "getter should return the property value")
	})

	t.Run("GetChannel_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		obj.Channel = nil

		// Act & Assert
		assert.Nil(t, obj.GetChannel(), "getter should return nil when property is nil")
	})

	t.Run("GetChannel_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetChannel() // Should return zero value
	})

	t.Run("GetContent", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var expected *FacebookAdsContent
		obj.Content = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetContent(), "getter should return the property value")
	})

	t.Run("GetContent_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		obj.Content = nil

		// Act & Assert
		assert.Nil(t, obj.GetContent(), "getter should return nil when property is nil")
	})

	t.Run("GetContent_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetContent() // Should return zero value
	})

	t.Run("GetFeedback", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var expected *FacebookAdsFeedback
		obj.Feedback = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetFeedback(), "getter should return the property value")
	})

	t.Run("GetFeedback_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		obj.Feedback = nil

		// Act & Assert
		assert.Nil(t, obj.GetFeedback(), "getter should return nil when property is nil")
	})

	t.Run("GetFeedback_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetFeedback() // Should return zero value
	})

	t.Run("GetHasAudience", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var expected *bool
		obj.HasAudience = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetHasAudience(), "getter should return the property value")
	})

	t.Run("GetHasAudience_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		obj.HasAudience = nil

		// Act & Assert
		assert.Nil(t, obj.GetHasAudience(), "getter should return nil when property is nil")
	})

	t.Run("GetHasAudience_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetHasAudience() // Should return zero value
	})

	t.Run("GetHasContent", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var expected *bool
		obj.HasContent = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetHasContent(), "getter should return the property value")
	})

	t.Run("GetHasContent_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		obj.HasContent = nil

		// Act & Assert
		assert.Nil(t, obj.GetHasContent(), "getter should return nil when property is nil")
	})

	t.Run("GetHasContent_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetHasContent() // Should return zero value
	})

	t.Run("GetIsConnected", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var expected *bool
		obj.IsConnected = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetIsConnected(), "getter should return the property value")
	})

	t.Run("GetIsConnected_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		obj.IsConnected = nil

		// Act & Assert
		assert.Nil(t, obj.GetIsConnected(), "getter should return nil when property is nil")
	})

	t.Run("GetIsConnected_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetIsConnected() // Should return zero value
	})

	t.Run("GetSite", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var expected *FacebookAdsSite
		obj.Site = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSite(), "getter should return the property value")
	})

	t.Run("GetSite_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		obj.Site = nil

		// Act & Assert
		assert.Nil(t, obj.GetSite(), "getter should return nil when property is nil")
	})

	t.Run("GetSite_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSite() // Should return zero value
	})

	t.Run("GetLinks", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var expected []*FacebookAdsLinksItem
		obj.Links = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetLinks(), "getter should return the property value")
	})

	t.Run("GetLinks_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		obj.Links = nil

		// Act & Assert
		assert.Nil(t, obj.GetLinks(), "getter should return nil when property is nil")
	})

	t.Run("GetLinks_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetLinks() // Should return zero value
	})

}

func TestSettersMarkExplicitFacebookAds(t *testing.T) {
	t.Run("SetCanceledAt_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var fernTestValueCanceledAt *time.Time

		// Act
		obj.SetCanceledAt(fernTestValueCanceledAt)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetCreateTime_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var fernTestValueCreateTime *time.Time

		// Act
		obj.SetCreateTime(fernTestValueCreateTime)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetHasSegment_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var fernTestValueHasSegment *bool

		// Act
		obj.SetHasSegment(fernTestValueHasSegment)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
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
		obj := &FacebookAds{}
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

	t.Run("SetName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
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

	t.Run("SetPublishedTime_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var fernTestValuePublishedTime *time.Time

		// Act
		obj.SetPublishedTime(fernTestValuePublishedTime)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetRecipients_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var fernTestValueRecipients *FacebookAdRecipients

		// Act
		obj.SetRecipients(fernTestValueRecipients)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetReportSummary_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var fernTestValueReportSummary *FacebookAdReportSummary

		// Act
		obj.SetReportSummary(fernTestValueReportSummary)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetShowReport_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var fernTestValueShowReport *bool

		// Act
		obj.SetShowReport(fernTestValueShowReport)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetStartTime_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var fernTestValueStartTime *time.Time

		// Act
		obj.SetStartTime(fernTestValueStartTime)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
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
		obj := &FacebookAds{}
		var fernTestValueStatus *FacebookAdStatus

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

	t.Run("SetType_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var fernTestValueType *FacebookAdType

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

	t.Run("SetUpdatedAt_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
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

	t.Run("SetWebID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
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

	t.Run("SetEmailSourceName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var fernTestValueEmailSourceName *string

		// Act
		obj.SetEmailSourceName(fernTestValueEmailSourceName)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetEndTime_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var fernTestValueEndTime *time.Time

		// Act
		obj.SetEndTime(fernTestValueEndTime)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetNeedsAttention_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var fernTestValueNeedsAttention *bool

		// Act
		obj.SetNeedsAttention(fernTestValueNeedsAttention)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetPausedAt_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var fernTestValuePausedAt *time.Time

		// Act
		obj.SetPausedAt(fernTestValuePausedAt)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetThumbnail_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var fernTestValueThumbnail *string

		// Act
		obj.SetThumbnail(fernTestValueThumbnail)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetWasCanceledByFacebook_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var fernTestValueWasCanceledByFacebook *bool

		// Act
		obj.SetWasCanceledByFacebook(fernTestValueWasCanceledByFacebook)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetAudience_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var fernTestValueAudience *FacebookAdsAudience

		// Act
		obj.SetAudience(fernTestValueAudience)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetBudget_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var fernTestValueBudget *FacebookAdsBudget

		// Act
		obj.SetBudget(fernTestValueBudget)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetChannel_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var fernTestValueChannel *FacebookAdsChannel

		// Act
		obj.SetChannel(fernTestValueChannel)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetContent_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var fernTestValueContent *FacebookAdsContent

		// Act
		obj.SetContent(fernTestValueContent)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetFeedback_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var fernTestValueFeedback *FacebookAdsFeedback

		// Act
		obj.SetFeedback(fernTestValueFeedback)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetHasAudience_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var fernTestValueHasAudience *bool

		// Act
		obj.SetHasAudience(fernTestValueHasAudience)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetHasContent_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var fernTestValueHasContent *bool

		// Act
		obj.SetHasContent(fernTestValueHasContent)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetIsConnected_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var fernTestValueIsConnected *bool

		// Act
		obj.SetIsConnected(fernTestValueIsConnected)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetSite_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}
		var fernTestValueSite *FacebookAdsSite

		// Act
		obj.SetSite(fernTestValueSite)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
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
		obj := &FacebookAds{}
		var fernTestValueLinks []*FacebookAdsLinksItem

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

func TestSettersFacebookAdsAudience(t *testing.T) {
	t.Run("SetEmailSource", func(t *testing.T) {
		obj := &FacebookAdsAudience{}
		var fernTestValueEmailSource *FacebookAdsAudienceEmailSource
		obj.SetEmailSource(fernTestValueEmailSource)
		assert.Equal(t, fernTestValueEmailSource, obj.EmailSource)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetIncludeSourceInTarget", func(t *testing.T) {
		obj := &FacebookAdsAudience{}
		var fernTestValueIncludeSourceInTarget *bool
		obj.SetIncludeSourceInTarget(fernTestValueIncludeSourceInTarget)
		assert.Equal(t, fernTestValueIncludeSourceInTarget, obj.IncludeSourceInTarget)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetLookalikeCountryCode", func(t *testing.T) {
		obj := &FacebookAdsAudience{}
		var fernTestValueLookalikeCountryCode *string
		obj.SetLookalikeCountryCode(fernTestValueLookalikeCountryCode)
		assert.Equal(t, fernTestValueLookalikeCountryCode, obj.LookalikeCountryCode)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSourceType", func(t *testing.T) {
		obj := &FacebookAdsAudience{}
		var fernTestValueSourceType *FacebookAdsAudienceSourceType
		obj.SetSourceType(fernTestValueSourceType)
		assert.Equal(t, fernTestValueSourceType, obj.SourceType)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTargetingSpecs", func(t *testing.T) {
		obj := &FacebookAdsAudience{}
		var fernTestValueTargetingSpecs *FacebookAdsAudienceTargetingSpecs
		obj.SetTargetingSpecs(fernTestValueTargetingSpecs)
		assert.Equal(t, fernTestValueTargetingSpecs, obj.TargetingSpecs)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetType", func(t *testing.T) {
		obj := &FacebookAdsAudience{}
		var fernTestValueType *FacebookAdsAudienceType
		obj.SetType(fernTestValueType)
		assert.Equal(t, fernTestValueType, obj.Type)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersFacebookAdsAudience(t *testing.T) {
	t.Run("GetEmailSource", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudience{}
		var expected *FacebookAdsAudienceEmailSource
		obj.EmailSource = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEmailSource(), "getter should return the property value")
	})

	t.Run("GetEmailSource_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudience{}
		obj.EmailSource = nil

		// Act & Assert
		assert.Nil(t, obj.GetEmailSource(), "getter should return nil when property is nil")
	})

	t.Run("GetEmailSource_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsAudience
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEmailSource() // Should return zero value
	})

	t.Run("GetIncludeSourceInTarget", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudience{}
		var expected *bool
		obj.IncludeSourceInTarget = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetIncludeSourceInTarget(), "getter should return the property value")
	})

	t.Run("GetIncludeSourceInTarget_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudience{}
		obj.IncludeSourceInTarget = nil

		// Act & Assert
		assert.Nil(t, obj.GetIncludeSourceInTarget(), "getter should return nil when property is nil")
	})

	t.Run("GetIncludeSourceInTarget_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsAudience
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetIncludeSourceInTarget() // Should return zero value
	})

	t.Run("GetLookalikeCountryCode", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudience{}
		var expected *string
		obj.LookalikeCountryCode = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetLookalikeCountryCode(), "getter should return the property value")
	})

	t.Run("GetLookalikeCountryCode_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudience{}
		obj.LookalikeCountryCode = nil

		// Act & Assert
		assert.Nil(t, obj.GetLookalikeCountryCode(), "getter should return nil when property is nil")
	})

	t.Run("GetLookalikeCountryCode_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsAudience
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetLookalikeCountryCode() // Should return zero value
	})

	t.Run("GetSourceType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudience{}
		var expected *FacebookAdsAudienceSourceType
		obj.SourceType = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSourceType(), "getter should return the property value")
	})

	t.Run("GetSourceType_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudience{}
		obj.SourceType = nil

		// Act & Assert
		assert.Nil(t, obj.GetSourceType(), "getter should return nil when property is nil")
	})

	t.Run("GetSourceType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsAudience
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSourceType() // Should return zero value
	})

	t.Run("GetTargetingSpecs", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudience{}
		var expected *FacebookAdsAudienceTargetingSpecs
		obj.TargetingSpecs = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTargetingSpecs(), "getter should return the property value")
	})

	t.Run("GetTargetingSpecs_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudience{}
		obj.TargetingSpecs = nil

		// Act & Assert
		assert.Nil(t, obj.GetTargetingSpecs(), "getter should return nil when property is nil")
	})

	t.Run("GetTargetingSpecs_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsAudience
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTargetingSpecs() // Should return zero value
	})

	t.Run("GetType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudience{}
		var expected *FacebookAdsAudienceType
		obj.Type = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetType(), "getter should return the property value")
	})

	t.Run("GetType_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudience{}
		obj.Type = nil

		// Act & Assert
		assert.Nil(t, obj.GetType(), "getter should return nil when property is nil")
	})

	t.Run("GetType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsAudience
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetType() // Should return zero value
	})

}

func TestSettersMarkExplicitFacebookAdsAudience(t *testing.T) {
	t.Run("SetEmailSource_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudience{}
		var fernTestValueEmailSource *FacebookAdsAudienceEmailSource

		// Act
		obj.SetEmailSource(fernTestValueEmailSource)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetIncludeSourceInTarget_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudience{}
		var fernTestValueIncludeSourceInTarget *bool

		// Act
		obj.SetIncludeSourceInTarget(fernTestValueIncludeSourceInTarget)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetLookalikeCountryCode_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudience{}
		var fernTestValueLookalikeCountryCode *string

		// Act
		obj.SetLookalikeCountryCode(fernTestValueLookalikeCountryCode)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetSourceType_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudience{}
		var fernTestValueSourceType *FacebookAdsAudienceSourceType

		// Act
		obj.SetSourceType(fernTestValueSourceType)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetTargetingSpecs_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudience{}
		var fernTestValueTargetingSpecs *FacebookAdsAudienceTargetingSpecs

		// Act
		obj.SetTargetingSpecs(fernTestValueTargetingSpecs)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
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
		obj := &FacebookAdsAudience{}
		var fernTestValueType *FacebookAdsAudienceType

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

func TestSettersFacebookAdsAudienceEmailSource(t *testing.T) {
	t.Run("SetIsSegment", func(t *testing.T) {
		obj := &FacebookAdsAudienceEmailSource{}
		var fernTestValueIsSegment *bool
		obj.SetIsSegment(fernTestValueIsSegment)
		assert.Equal(t, fernTestValueIsSegment, obj.IsSegment)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetListName", func(t *testing.T) {
		obj := &FacebookAdsAudienceEmailSource{}
		var fernTestValueListName *string
		obj.SetListName(fernTestValueListName)
		assert.Equal(t, fernTestValueListName, obj.ListName)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetName", func(t *testing.T) {
		obj := &FacebookAdsAudienceEmailSource{}
		var fernTestValueName *string
		obj.SetName(fernTestValueName)
		assert.Equal(t, fernTestValueName, obj.Name)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSegmentType", func(t *testing.T) {
		obj := &FacebookAdsAudienceEmailSource{}
		var fernTestValueSegmentType *string
		obj.SetSegmentType(fernTestValueSegmentType)
		assert.Equal(t, fernTestValueSegmentType, obj.SegmentType)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetType", func(t *testing.T) {
		obj := &FacebookAdsAudienceEmailSource{}
		var fernTestValueType *string
		obj.SetType(fernTestValueType)
		assert.Equal(t, fernTestValueType, obj.Type)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersFacebookAdsAudienceEmailSource(t *testing.T) {
	t.Run("GetIsSegment", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceEmailSource{}
		var expected *bool
		obj.IsSegment = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetIsSegment(), "getter should return the property value")
	})

	t.Run("GetIsSegment_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceEmailSource{}
		obj.IsSegment = nil

		// Act & Assert
		assert.Nil(t, obj.GetIsSegment(), "getter should return nil when property is nil")
	})

	t.Run("GetIsSegment_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsAudienceEmailSource
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetIsSegment() // Should return zero value
	})

	t.Run("GetListName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceEmailSource{}
		var expected *string
		obj.ListName = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetListName(), "getter should return the property value")
	})

	t.Run("GetListName_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceEmailSource{}
		obj.ListName = nil

		// Act & Assert
		assert.Nil(t, obj.GetListName(), "getter should return nil when property is nil")
	})

	t.Run("GetListName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsAudienceEmailSource
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetListName() // Should return zero value
	})

	t.Run("GetName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceEmailSource{}
		var expected *string
		obj.Name = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetName(), "getter should return the property value")
	})

	t.Run("GetName_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceEmailSource{}
		obj.Name = nil

		// Act & Assert
		assert.Nil(t, obj.GetName(), "getter should return nil when property is nil")
	})

	t.Run("GetName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsAudienceEmailSource
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetName() // Should return zero value
	})

	t.Run("GetSegmentType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceEmailSource{}
		var expected *string
		obj.SegmentType = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSegmentType(), "getter should return the property value")
	})

	t.Run("GetSegmentType_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceEmailSource{}
		obj.SegmentType = nil

		// Act & Assert
		assert.Nil(t, obj.GetSegmentType(), "getter should return nil when property is nil")
	})

	t.Run("GetSegmentType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsAudienceEmailSource
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetSegmentType() // Should return zero value
	})

	t.Run("GetType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceEmailSource{}
		var expected *string
		obj.Type = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetType(), "getter should return the property value")
	})

	t.Run("GetType_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceEmailSource{}
		obj.Type = nil

		// Act & Assert
		assert.Nil(t, obj.GetType(), "getter should return nil when property is nil")
	})

	t.Run("GetType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsAudienceEmailSource
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetType() // Should return zero value
	})

}

func TestSettersMarkExplicitFacebookAdsAudienceEmailSource(t *testing.T) {
	t.Run("SetIsSegment_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceEmailSource{}
		var fernTestValueIsSegment *bool

		// Act
		obj.SetIsSegment(fernTestValueIsSegment)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetListName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceEmailSource{}
		var fernTestValueListName *string

		// Act
		obj.SetListName(fernTestValueListName)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
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
		obj := &FacebookAdsAudienceEmailSource{}
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

	t.Run("SetSegmentType_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceEmailSource{}
		var fernTestValueSegmentType *string

		// Act
		obj.SetSegmentType(fernTestValueSegmentType)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
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
		obj := &FacebookAdsAudienceEmailSource{}
		var fernTestValueType *string

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

func TestSettersFacebookAdsAudienceTargetingSpecs(t *testing.T) {
	t.Run("SetGender", func(t *testing.T) {
		obj := &FacebookAdsAudienceTargetingSpecs{}
		var fernTestValueGender *int
		obj.SetGender(fernTestValueGender)
		assert.Equal(t, fernTestValueGender, obj.Gender)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetInterests", func(t *testing.T) {
		obj := &FacebookAdsAudienceTargetingSpecs{}
		var fernTestValueInterests []*FacebookAdsAudienceTargetingSpecsInterestsItem
		obj.SetInterests(fernTestValueInterests)
		assert.Equal(t, fernTestValueInterests, obj.Interests)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetLocations", func(t *testing.T) {
		obj := &FacebookAdsAudienceTargetingSpecs{}
		var fernTestValueLocations *FacebookAdsAudienceTargetingSpecsLocations
		obj.SetLocations(fernTestValueLocations)
		assert.Equal(t, fernTestValueLocations, obj.Locations)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMaxAge", func(t *testing.T) {
		obj := &FacebookAdsAudienceTargetingSpecs{}
		var fernTestValueMaxAge *int
		obj.SetMaxAge(fernTestValueMaxAge)
		assert.Equal(t, fernTestValueMaxAge, obj.MaxAge)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMinAge", func(t *testing.T) {
		obj := &FacebookAdsAudienceTargetingSpecs{}
		var fernTestValueMinAge *int
		obj.SetMinAge(fernTestValueMinAge)
		assert.Equal(t, fernTestValueMinAge, obj.MinAge)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersFacebookAdsAudienceTargetingSpecs(t *testing.T) {
	t.Run("GetGender", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecs{}
		var expected *int
		obj.Gender = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetGender(), "getter should return the property value")
	})

	t.Run("GetGender_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecs{}
		obj.Gender = nil

		// Act & Assert
		assert.Nil(t, obj.GetGender(), "getter should return nil when property is nil")
	})

	t.Run("GetGender_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsAudienceTargetingSpecs
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetGender() // Should return zero value
	})

	t.Run("GetInterests", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecs{}
		var expected []*FacebookAdsAudienceTargetingSpecsInterestsItem
		obj.Interests = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetInterests(), "getter should return the property value")
	})

	t.Run("GetInterests_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecs{}
		obj.Interests = nil

		// Act & Assert
		assert.Nil(t, obj.GetInterests(), "getter should return nil when property is nil")
	})

	t.Run("GetInterests_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsAudienceTargetingSpecs
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetInterests() // Should return zero value
	})

	t.Run("GetLocations", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecs{}
		var expected *FacebookAdsAudienceTargetingSpecsLocations
		obj.Locations = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetLocations(), "getter should return the property value")
	})

	t.Run("GetLocations_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecs{}
		obj.Locations = nil

		// Act & Assert
		assert.Nil(t, obj.GetLocations(), "getter should return nil when property is nil")
	})

	t.Run("GetLocations_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsAudienceTargetingSpecs
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetLocations() // Should return zero value
	})

	t.Run("GetMaxAge", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecs{}
		var expected *int
		obj.MaxAge = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMaxAge(), "getter should return the property value")
	})

	t.Run("GetMaxAge_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecs{}
		obj.MaxAge = nil

		// Act & Assert
		assert.Nil(t, obj.GetMaxAge(), "getter should return nil when property is nil")
	})

	t.Run("GetMaxAge_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsAudienceTargetingSpecs
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetMaxAge() // Should return zero value
	})

	t.Run("GetMinAge", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecs{}
		var expected *int
		obj.MinAge = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMinAge(), "getter should return the property value")
	})

	t.Run("GetMinAge_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecs{}
		obj.MinAge = nil

		// Act & Assert
		assert.Nil(t, obj.GetMinAge(), "getter should return nil when property is nil")
	})

	t.Run("GetMinAge_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsAudienceTargetingSpecs
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetMinAge() // Should return zero value
	})

}

func TestSettersMarkExplicitFacebookAdsAudienceTargetingSpecs(t *testing.T) {
	t.Run("SetGender_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecs{}
		var fernTestValueGender *int

		// Act
		obj.SetGender(fernTestValueGender)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetInterests_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecs{}
		var fernTestValueInterests []*FacebookAdsAudienceTargetingSpecsInterestsItem

		// Act
		obj.SetInterests(fernTestValueInterests)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetLocations_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecs{}
		var fernTestValueLocations *FacebookAdsAudienceTargetingSpecsLocations

		// Act
		obj.SetLocations(fernTestValueLocations)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetMaxAge_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecs{}
		var fernTestValueMaxAge *int

		// Act
		obj.SetMaxAge(fernTestValueMaxAge)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetMinAge_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecs{}
		var fernTestValueMinAge *int

		// Act
		obj.SetMinAge(fernTestValueMinAge)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
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

func TestSettersFacebookAdsAudienceTargetingSpecsInterestsItem(t *testing.T) {
	t.Run("SetName", func(t *testing.T) {
		obj := &FacebookAdsAudienceTargetingSpecsInterestsItem{}
		var fernTestValueName *string
		obj.SetName(fernTestValueName)
		assert.Equal(t, fernTestValueName, obj.Name)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersFacebookAdsAudienceTargetingSpecsInterestsItem(t *testing.T) {
	t.Run("GetName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecsInterestsItem{}
		var expected *string
		obj.Name = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetName(), "getter should return the property value")
	})

	t.Run("GetName_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecsInterestsItem{}
		obj.Name = nil

		// Act & Assert
		assert.Nil(t, obj.GetName(), "getter should return nil when property is nil")
	})

	t.Run("GetName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsAudienceTargetingSpecsInterestsItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetName() // Should return zero value
	})

}

func TestSettersMarkExplicitFacebookAdsAudienceTargetingSpecsInterestsItem(t *testing.T) {
	t.Run("SetName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecsInterestsItem{}
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

func TestSettersFacebookAdsAudienceTargetingSpecsLocations(t *testing.T) {
	t.Run("SetCities", func(t *testing.T) {
		obj := &FacebookAdsAudienceTargetingSpecsLocations{}
		var fernTestValueCities []string
		obj.SetCities(fernTestValueCities)
		assert.Equal(t, fernTestValueCities, obj.Cities)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCountries", func(t *testing.T) {
		obj := &FacebookAdsAudienceTargetingSpecsLocations{}
		var fernTestValueCountries []string
		obj.SetCountries(fernTestValueCountries)
		assert.Equal(t, fernTestValueCountries, obj.Countries)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRegions", func(t *testing.T) {
		obj := &FacebookAdsAudienceTargetingSpecsLocations{}
		var fernTestValueRegions []string
		obj.SetRegions(fernTestValueRegions)
		assert.Equal(t, fernTestValueRegions, obj.Regions)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetZips", func(t *testing.T) {
		obj := &FacebookAdsAudienceTargetingSpecsLocations{}
		var fernTestValueZips []string
		obj.SetZips(fernTestValueZips)
		assert.Equal(t, fernTestValueZips, obj.Zips)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersFacebookAdsAudienceTargetingSpecsLocations(t *testing.T) {
	t.Run("GetCities", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecsLocations{}
		var expected []string
		obj.Cities = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCities(), "getter should return the property value")
	})

	t.Run("GetCities_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecsLocations{}
		obj.Cities = nil

		// Act & Assert
		assert.Nil(t, obj.GetCities(), "getter should return nil when property is nil")
	})

	t.Run("GetCities_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsAudienceTargetingSpecsLocations
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCities() // Should return zero value
	})

	t.Run("GetCountries", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecsLocations{}
		var expected []string
		obj.Countries = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCountries(), "getter should return the property value")
	})

	t.Run("GetCountries_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecsLocations{}
		obj.Countries = nil

		// Act & Assert
		assert.Nil(t, obj.GetCountries(), "getter should return nil when property is nil")
	})

	t.Run("GetCountries_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsAudienceTargetingSpecsLocations
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCountries() // Should return zero value
	})

	t.Run("GetRegions", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecsLocations{}
		var expected []string
		obj.Regions = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetRegions(), "getter should return the property value")
	})

	t.Run("GetRegions_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecsLocations{}
		obj.Regions = nil

		// Act & Assert
		assert.Nil(t, obj.GetRegions(), "getter should return nil when property is nil")
	})

	t.Run("GetRegions_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsAudienceTargetingSpecsLocations
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetRegions() // Should return zero value
	})

	t.Run("GetZips", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecsLocations{}
		var expected []string
		obj.Zips = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetZips(), "getter should return the property value")
	})

	t.Run("GetZips_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecsLocations{}
		obj.Zips = nil

		// Act & Assert
		assert.Nil(t, obj.GetZips(), "getter should return nil when property is nil")
	})

	t.Run("GetZips_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsAudienceTargetingSpecsLocations
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetZips() // Should return zero value
	})

}

func TestSettersMarkExplicitFacebookAdsAudienceTargetingSpecsLocations(t *testing.T) {
	t.Run("SetCities_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecsLocations{}
		var fernTestValueCities []string

		// Act
		obj.SetCities(fernTestValueCities)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetCountries_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecsLocations{}
		var fernTestValueCountries []string

		// Act
		obj.SetCountries(fernTestValueCountries)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetRegions_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecsLocations{}
		var fernTestValueRegions []string

		// Act
		obj.SetRegions(fernTestValueRegions)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetZips_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecsLocations{}
		var fernTestValueZips []string

		// Act
		obj.SetZips(fernTestValueZips)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
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

func TestSettersFacebookAdsBudget(t *testing.T) {
	t.Run("SetCurrencyCode", func(t *testing.T) {
		obj := &FacebookAdsBudget{}
		var fernTestValueCurrencyCode *string
		obj.SetCurrencyCode(fernTestValueCurrencyCode)
		assert.Equal(t, fernTestValueCurrencyCode, obj.CurrencyCode)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetDuration", func(t *testing.T) {
		obj := &FacebookAdsBudget{}
		var fernTestValueDuration *int
		obj.SetDuration(fernTestValueDuration)
		assert.Equal(t, fernTestValueDuration, obj.Duration)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTotalAmount", func(t *testing.T) {
		obj := &FacebookAdsBudget{}
		var fernTestValueTotalAmount *float64
		obj.SetTotalAmount(fernTestValueTotalAmount)
		assert.Equal(t, fernTestValueTotalAmount, obj.TotalAmount)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersFacebookAdsBudget(t *testing.T) {
	t.Run("GetCurrencyCode", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsBudget{}
		var expected *string
		obj.CurrencyCode = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCurrencyCode(), "getter should return the property value")
	})

	t.Run("GetCurrencyCode_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsBudget{}
		obj.CurrencyCode = nil

		// Act & Assert
		assert.Nil(t, obj.GetCurrencyCode(), "getter should return nil when property is nil")
	})

	t.Run("GetCurrencyCode_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsBudget
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCurrencyCode() // Should return zero value
	})

	t.Run("GetDuration", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsBudget{}
		var expected *int
		obj.Duration = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetDuration(), "getter should return the property value")
	})

	t.Run("GetDuration_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsBudget{}
		obj.Duration = nil

		// Act & Assert
		assert.Nil(t, obj.GetDuration(), "getter should return nil when property is nil")
	})

	t.Run("GetDuration_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsBudget
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetDuration() // Should return zero value
	})

	t.Run("GetTotalAmount", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsBudget{}
		var expected *float64
		obj.TotalAmount = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTotalAmount(), "getter should return the property value")
	})

	t.Run("GetTotalAmount_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsBudget{}
		obj.TotalAmount = nil

		// Act & Assert
		assert.Nil(t, obj.GetTotalAmount(), "getter should return nil when property is nil")
	})

	t.Run("GetTotalAmount_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsBudget
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTotalAmount() // Should return zero value
	})

}

func TestSettersMarkExplicitFacebookAdsBudget(t *testing.T) {
	t.Run("SetCurrencyCode_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsBudget{}
		var fernTestValueCurrencyCode *string

		// Act
		obj.SetCurrencyCode(fernTestValueCurrencyCode)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetDuration_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsBudget{}
		var fernTestValueDuration *int

		// Act
		obj.SetDuration(fernTestValueDuration)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetTotalAmount_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsBudget{}
		var fernTestValueTotalAmount *float64

		// Act
		obj.SetTotalAmount(fernTestValueTotalAmount)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
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

func TestSettersFacebookAdsChannel(t *testing.T) {
	t.Run("SetFbPlacementAudience", func(t *testing.T) {
		obj := &FacebookAdsChannel{}
		var fernTestValueFbPlacementAudience *bool
		obj.SetFbPlacementAudience(fernTestValueFbPlacementAudience)
		assert.Equal(t, fernTestValueFbPlacementAudience, obj.FbPlacementAudience)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetFbPlacementFeed", func(t *testing.T) {
		obj := &FacebookAdsChannel{}
		var fernTestValueFbPlacementFeed *bool
		obj.SetFbPlacementFeed(fernTestValueFbPlacementFeed)
		assert.Equal(t, fernTestValueFbPlacementFeed, obj.FbPlacementFeed)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetIgPlacementFeed", func(t *testing.T) {
		obj := &FacebookAdsChannel{}
		var fernTestValueIgPlacementFeed *bool
		obj.SetIgPlacementFeed(fernTestValueIgPlacementFeed)
		assert.Equal(t, fernTestValueIgPlacementFeed, obj.IgPlacementFeed)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersFacebookAdsChannel(t *testing.T) {
	t.Run("GetFbPlacementAudience", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsChannel{}
		var expected *bool
		obj.FbPlacementAudience = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetFbPlacementAudience(), "getter should return the property value")
	})

	t.Run("GetFbPlacementAudience_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsChannel{}
		obj.FbPlacementAudience = nil

		// Act & Assert
		assert.Nil(t, obj.GetFbPlacementAudience(), "getter should return nil when property is nil")
	})

	t.Run("GetFbPlacementAudience_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsChannel
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetFbPlacementAudience() // Should return zero value
	})

	t.Run("GetFbPlacementFeed", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsChannel{}
		var expected *bool
		obj.FbPlacementFeed = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetFbPlacementFeed(), "getter should return the property value")
	})

	t.Run("GetFbPlacementFeed_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsChannel{}
		obj.FbPlacementFeed = nil

		// Act & Assert
		assert.Nil(t, obj.GetFbPlacementFeed(), "getter should return nil when property is nil")
	})

	t.Run("GetFbPlacementFeed_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsChannel
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetFbPlacementFeed() // Should return zero value
	})

	t.Run("GetIgPlacementFeed", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsChannel{}
		var expected *bool
		obj.IgPlacementFeed = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetIgPlacementFeed(), "getter should return the property value")
	})

	t.Run("GetIgPlacementFeed_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsChannel{}
		obj.IgPlacementFeed = nil

		// Act & Assert
		assert.Nil(t, obj.GetIgPlacementFeed(), "getter should return nil when property is nil")
	})

	t.Run("GetIgPlacementFeed_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsChannel
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetIgPlacementFeed() // Should return zero value
	})

}

func TestSettersMarkExplicitFacebookAdsChannel(t *testing.T) {
	t.Run("SetFbPlacementAudience_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsChannel{}
		var fernTestValueFbPlacementAudience *bool

		// Act
		obj.SetFbPlacementAudience(fernTestValueFbPlacementAudience)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetFbPlacementFeed_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsChannel{}
		var fernTestValueFbPlacementFeed *bool

		// Act
		obj.SetFbPlacementFeed(fernTestValueFbPlacementFeed)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetIgPlacementFeed_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsChannel{}
		var fernTestValueIgPlacementFeed *bool

		// Act
		obj.SetIgPlacementFeed(fernTestValueIgPlacementFeed)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
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

func TestSettersFacebookAdsContent(t *testing.T) {
	t.Run("SetAttachments", func(t *testing.T) {
		obj := &FacebookAdsContent{}
		var fernTestValueAttachments []*FacebookAdsContentAttachmentsItem
		obj.SetAttachments(fernTestValueAttachments)
		assert.Equal(t, fernTestValueAttachments, obj.Attachments)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCallToAction", func(t *testing.T) {
		obj := &FacebookAdsContent{}
		var fernTestValueCallToAction *string
		obj.SetCallToAction(fernTestValueCallToAction)
		assert.Equal(t, fernTestValueCallToAction, obj.CallToAction)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetDescription", func(t *testing.T) {
		obj := &FacebookAdsContent{}
		var fernTestValueDescription *string
		obj.SetDescription(fernTestValueDescription)
		assert.Equal(t, fernTestValueDescription, obj.Description)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetImageURL", func(t *testing.T) {
		obj := &FacebookAdsContent{}
		var fernTestValueImageURL *string
		obj.SetImageURL(fernTestValueImageURL)
		assert.Equal(t, fernTestValueImageURL, obj.ImageURL)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetLinkURL", func(t *testing.T) {
		obj := &FacebookAdsContent{}
		var fernTestValueLinkURL *string
		obj.SetLinkURL(fernTestValueLinkURL)
		assert.Equal(t, fernTestValueLinkURL, obj.LinkURL)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMessage", func(t *testing.T) {
		obj := &FacebookAdsContent{}
		var fernTestValueMessage *string
		obj.SetMessage(fernTestValueMessage)
		assert.Equal(t, fernTestValueMessage, obj.Message)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTitle", func(t *testing.T) {
		obj := &FacebookAdsContent{}
		var fernTestValueTitle *string
		obj.SetTitle(fernTestValueTitle)
		assert.Equal(t, fernTestValueTitle, obj.Title)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersFacebookAdsContent(t *testing.T) {
	t.Run("GetAttachments", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContent{}
		var expected []*FacebookAdsContentAttachmentsItem
		obj.Attachments = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetAttachments(), "getter should return the property value")
	})

	t.Run("GetAttachments_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContent{}
		obj.Attachments = nil

		// Act & Assert
		assert.Nil(t, obj.GetAttachments(), "getter should return nil when property is nil")
	})

	t.Run("GetAttachments_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsContent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetAttachments() // Should return zero value
	})

	t.Run("GetCallToAction", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContent{}
		var expected *string
		obj.CallToAction = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCallToAction(), "getter should return the property value")
	})

	t.Run("GetCallToAction_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContent{}
		obj.CallToAction = nil

		// Act & Assert
		assert.Nil(t, obj.GetCallToAction(), "getter should return nil when property is nil")
	})

	t.Run("GetCallToAction_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsContent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCallToAction() // Should return zero value
	})

	t.Run("GetDescription", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContent{}
		var expected *string
		obj.Description = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetDescription(), "getter should return the property value")
	})

	t.Run("GetDescription_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContent{}
		obj.Description = nil

		// Act & Assert
		assert.Nil(t, obj.GetDescription(), "getter should return nil when property is nil")
	})

	t.Run("GetDescription_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsContent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetDescription() // Should return zero value
	})

	t.Run("GetImageURL", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContent{}
		var expected *string
		obj.ImageURL = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetImageURL(), "getter should return the property value")
	})

	t.Run("GetImageURL_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContent{}
		obj.ImageURL = nil

		// Act & Assert
		assert.Nil(t, obj.GetImageURL(), "getter should return nil when property is nil")
	})

	t.Run("GetImageURL_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsContent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetImageURL() // Should return zero value
	})

	t.Run("GetLinkURL", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContent{}
		var expected *string
		obj.LinkURL = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetLinkURL(), "getter should return the property value")
	})

	t.Run("GetLinkURL_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContent{}
		obj.LinkURL = nil

		// Act & Assert
		assert.Nil(t, obj.GetLinkURL(), "getter should return nil when property is nil")
	})

	t.Run("GetLinkURL_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsContent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetLinkURL() // Should return zero value
	})

	t.Run("GetMessage", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContent{}
		var expected *string
		obj.Message = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMessage(), "getter should return the property value")
	})

	t.Run("GetMessage_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContent{}
		obj.Message = nil

		// Act & Assert
		assert.Nil(t, obj.GetMessage(), "getter should return nil when property is nil")
	})

	t.Run("GetMessage_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsContent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetMessage() // Should return zero value
	})

	t.Run("GetTitle", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContent{}
		var expected *string
		obj.Title = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTitle(), "getter should return the property value")
	})

	t.Run("GetTitle_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContent{}
		obj.Title = nil

		// Act & Assert
		assert.Nil(t, obj.GetTitle(), "getter should return nil when property is nil")
	})

	t.Run("GetTitle_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsContent
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTitle() // Should return zero value
	})

}

func TestSettersMarkExplicitFacebookAdsContent(t *testing.T) {
	t.Run("SetAttachments_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContent{}
		var fernTestValueAttachments []*FacebookAdsContentAttachmentsItem

		// Act
		obj.SetAttachments(fernTestValueAttachments)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetCallToAction_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContent{}
		var fernTestValueCallToAction *string

		// Act
		obj.SetCallToAction(fernTestValueCallToAction)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
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
		obj := &FacebookAdsContent{}
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

	t.Run("SetImageURL_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContent{}
		var fernTestValueImageURL *string

		// Act
		obj.SetImageURL(fernTestValueImageURL)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetLinkURL_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContent{}
		var fernTestValueLinkURL *string

		// Act
		obj.SetLinkURL(fernTestValueLinkURL)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetMessage_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContent{}
		var fernTestValueMessage *string

		// Act
		obj.SetMessage(fernTestValueMessage)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
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
		obj := &FacebookAdsContent{}
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

}

func TestSettersFacebookAdsContentAttachmentsItem(t *testing.T) {
	t.Run("SetCallToAction", func(t *testing.T) {
		obj := &FacebookAdsContentAttachmentsItem{}
		var fernTestValueCallToAction *string
		obj.SetCallToAction(fernTestValueCallToAction)
		assert.Equal(t, fernTestValueCallToAction, obj.CallToAction)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetDescription", func(t *testing.T) {
		obj := &FacebookAdsContentAttachmentsItem{}
		var fernTestValueDescription *string
		obj.SetDescription(fernTestValueDescription)
		assert.Equal(t, fernTestValueDescription, obj.Description)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetImageURL", func(t *testing.T) {
		obj := &FacebookAdsContentAttachmentsItem{}
		var fernTestValueImageURL *string
		obj.SetImageURL(fernTestValueImageURL)
		assert.Equal(t, fernTestValueImageURL, obj.ImageURL)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetLinkURL", func(t *testing.T) {
		obj := &FacebookAdsContentAttachmentsItem{}
		var fernTestValueLinkURL *string
		obj.SetLinkURL(fernTestValueLinkURL)
		assert.Equal(t, fernTestValueLinkURL, obj.LinkURL)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetName", func(t *testing.T) {
		obj := &FacebookAdsContentAttachmentsItem{}
		var fernTestValueName *string
		obj.SetName(fernTestValueName)
		assert.Equal(t, fernTestValueName, obj.Name)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersFacebookAdsContentAttachmentsItem(t *testing.T) {
	t.Run("GetCallToAction", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContentAttachmentsItem{}
		var expected *string
		obj.CallToAction = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCallToAction(), "getter should return the property value")
	})

	t.Run("GetCallToAction_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContentAttachmentsItem{}
		obj.CallToAction = nil

		// Act & Assert
		assert.Nil(t, obj.GetCallToAction(), "getter should return nil when property is nil")
	})

	t.Run("GetCallToAction_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsContentAttachmentsItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCallToAction() // Should return zero value
	})

	t.Run("GetDescription", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContentAttachmentsItem{}
		var expected *string
		obj.Description = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetDescription(), "getter should return the property value")
	})

	t.Run("GetDescription_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContentAttachmentsItem{}
		obj.Description = nil

		// Act & Assert
		assert.Nil(t, obj.GetDescription(), "getter should return nil when property is nil")
	})

	t.Run("GetDescription_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsContentAttachmentsItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetDescription() // Should return zero value
	})

	t.Run("GetImageURL", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContentAttachmentsItem{}
		var expected *string
		obj.ImageURL = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetImageURL(), "getter should return the property value")
	})

	t.Run("GetImageURL_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContentAttachmentsItem{}
		obj.ImageURL = nil

		// Act & Assert
		assert.Nil(t, obj.GetImageURL(), "getter should return nil when property is nil")
	})

	t.Run("GetImageURL_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsContentAttachmentsItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetImageURL() // Should return zero value
	})

	t.Run("GetLinkURL", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContentAttachmentsItem{}
		var expected *string
		obj.LinkURL = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetLinkURL(), "getter should return the property value")
	})

	t.Run("GetLinkURL_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContentAttachmentsItem{}
		obj.LinkURL = nil

		// Act & Assert
		assert.Nil(t, obj.GetLinkURL(), "getter should return nil when property is nil")
	})

	t.Run("GetLinkURL_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsContentAttachmentsItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetLinkURL() // Should return zero value
	})

	t.Run("GetName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContentAttachmentsItem{}
		var expected *string
		obj.Name = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetName(), "getter should return the property value")
	})

	t.Run("GetName_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContentAttachmentsItem{}
		obj.Name = nil

		// Act & Assert
		assert.Nil(t, obj.GetName(), "getter should return nil when property is nil")
	})

	t.Run("GetName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsContentAttachmentsItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetName() // Should return zero value
	})

}

func TestSettersMarkExplicitFacebookAdsContentAttachmentsItem(t *testing.T) {
	t.Run("SetCallToAction_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContentAttachmentsItem{}
		var fernTestValueCallToAction *string

		// Act
		obj.SetCallToAction(fernTestValueCallToAction)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
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
		obj := &FacebookAdsContentAttachmentsItem{}
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

	t.Run("SetImageURL_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContentAttachmentsItem{}
		var fernTestValueImageURL *string

		// Act
		obj.SetImageURL(fernTestValueImageURL)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetLinkURL_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContentAttachmentsItem{}
		var fernTestValueLinkURL *string

		// Act
		obj.SetLinkURL(fernTestValueLinkURL)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
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
		obj := &FacebookAdsContentAttachmentsItem{}
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

func TestSettersFacebookAdsFeedback(t *testing.T) {
	t.Run("SetAudience", func(t *testing.T) {
		obj := &FacebookAdsFeedback{}
		var fernTestValueAudience *string
		obj.SetAudience(fernTestValueAudience)
		assert.Equal(t, fernTestValueAudience, obj.Audience)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetBudget", func(t *testing.T) {
		obj := &FacebookAdsFeedback{}
		var fernTestValueBudget *string
		obj.SetBudget(fernTestValueBudget)
		assert.Equal(t, fernTestValueBudget, obj.Budget)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCompliance", func(t *testing.T) {
		obj := &FacebookAdsFeedback{}
		var fernTestValueCompliance *string
		obj.SetCompliance(fernTestValueCompliance)
		assert.Equal(t, fernTestValueCompliance, obj.Compliance)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetContent", func(t *testing.T) {
		obj := &FacebookAdsFeedback{}
		var fernTestValueContent *string
		obj.SetContent(fernTestValueContent)
		assert.Equal(t, fernTestValueContent, obj.Content)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersFacebookAdsFeedback(t *testing.T) {
	t.Run("GetAudience", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsFeedback{}
		var expected *string
		obj.Audience = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetAudience(), "getter should return the property value")
	})

	t.Run("GetAudience_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsFeedback{}
		obj.Audience = nil

		// Act & Assert
		assert.Nil(t, obj.GetAudience(), "getter should return nil when property is nil")
	})

	t.Run("GetAudience_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsFeedback
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetAudience() // Should return zero value
	})

	t.Run("GetBudget", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsFeedback{}
		var expected *string
		obj.Budget = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetBudget(), "getter should return the property value")
	})

	t.Run("GetBudget_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsFeedback{}
		obj.Budget = nil

		// Act & Assert
		assert.Nil(t, obj.GetBudget(), "getter should return nil when property is nil")
	})

	t.Run("GetBudget_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsFeedback
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetBudget() // Should return zero value
	})

	t.Run("GetCompliance", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsFeedback{}
		var expected *string
		obj.Compliance = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCompliance(), "getter should return the property value")
	})

	t.Run("GetCompliance_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsFeedback{}
		obj.Compliance = nil

		// Act & Assert
		assert.Nil(t, obj.GetCompliance(), "getter should return nil when property is nil")
	})

	t.Run("GetCompliance_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsFeedback
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCompliance() // Should return zero value
	})

	t.Run("GetContent", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsFeedback{}
		var expected *string
		obj.Content = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetContent(), "getter should return the property value")
	})

	t.Run("GetContent_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsFeedback{}
		obj.Content = nil

		// Act & Assert
		assert.Nil(t, obj.GetContent(), "getter should return nil when property is nil")
	})

	t.Run("GetContent_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsFeedback
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetContent() // Should return zero value
	})

}

func TestSettersMarkExplicitFacebookAdsFeedback(t *testing.T) {
	t.Run("SetAudience_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsFeedback{}
		var fernTestValueAudience *string

		// Act
		obj.SetAudience(fernTestValueAudience)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetBudget_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsFeedback{}
		var fernTestValueBudget *string

		// Act
		obj.SetBudget(fernTestValueBudget)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetCompliance_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsFeedback{}
		var fernTestValueCompliance *string

		// Act
		obj.SetCompliance(fernTestValueCompliance)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetContent_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsFeedback{}
		var fernTestValueContent *string

		// Act
		obj.SetContent(fernTestValueContent)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
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

func TestSettersFacebookAdsLinksItem(t *testing.T) {
	t.Run("SetHref", func(t *testing.T) {
		obj := &FacebookAdsLinksItem{}
		var fernTestValueHref *string
		obj.SetHref(fernTestValueHref)
		assert.Equal(t, fernTestValueHref, obj.Href)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMethod", func(t *testing.T) {
		obj := &FacebookAdsLinksItem{}
		var fernTestValueMethod *FacebookAdsLinksItemMethod
		obj.SetMethod(fernTestValueMethod)
		assert.Equal(t, fernTestValueMethod, obj.Method)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRel", func(t *testing.T) {
		obj := &FacebookAdsLinksItem{}
		var fernTestValueRel *string
		obj.SetRel(fernTestValueRel)
		assert.Equal(t, fernTestValueRel, obj.Rel)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSchema", func(t *testing.T) {
		obj := &FacebookAdsLinksItem{}
		var fernTestValueSchema *string
		obj.SetSchema(fernTestValueSchema)
		assert.Equal(t, fernTestValueSchema, obj.Schema)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTargetSchema", func(t *testing.T) {
		obj := &FacebookAdsLinksItem{}
		var fernTestValueTargetSchema *string
		obj.SetTargetSchema(fernTestValueTargetSchema)
		assert.Equal(t, fernTestValueTargetSchema, obj.TargetSchema)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersFacebookAdsLinksItem(t *testing.T) {
	t.Run("GetHref", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsLinksItem{}
		var expected *string
		obj.Href = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetHref(), "getter should return the property value")
	})

	t.Run("GetHref_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsLinksItem{}
		obj.Href = nil

		// Act & Assert
		assert.Nil(t, obj.GetHref(), "getter should return nil when property is nil")
	})

	t.Run("GetHref_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsLinksItem
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
		obj := &FacebookAdsLinksItem{}
		var expected *FacebookAdsLinksItemMethod
		obj.Method = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMethod(), "getter should return the property value")
	})

	t.Run("GetMethod_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsLinksItem{}
		obj.Method = nil

		// Act & Assert
		assert.Nil(t, obj.GetMethod(), "getter should return nil when property is nil")
	})

	t.Run("GetMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsLinksItem
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
		obj := &FacebookAdsLinksItem{}
		var expected *string
		obj.Rel = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetRel(), "getter should return the property value")
	})

	t.Run("GetRel_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsLinksItem{}
		obj.Rel = nil

		// Act & Assert
		assert.Nil(t, obj.GetRel(), "getter should return nil when property is nil")
	})

	t.Run("GetRel_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsLinksItem
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
		obj := &FacebookAdsLinksItem{}
		var expected *string
		obj.Schema = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSchema(), "getter should return the property value")
	})

	t.Run("GetSchema_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsLinksItem{}
		obj.Schema = nil

		// Act & Assert
		assert.Nil(t, obj.GetSchema(), "getter should return nil when property is nil")
	})

	t.Run("GetSchema_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsLinksItem
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
		obj := &FacebookAdsLinksItem{}
		var expected *string
		obj.TargetSchema = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTargetSchema(), "getter should return the property value")
	})

	t.Run("GetTargetSchema_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsLinksItem{}
		obj.TargetSchema = nil

		// Act & Assert
		assert.Nil(t, obj.GetTargetSchema(), "getter should return nil when property is nil")
	})

	t.Run("GetTargetSchema_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsLinksItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTargetSchema() // Should return zero value
	})

}

func TestSettersMarkExplicitFacebookAdsLinksItem(t *testing.T) {
	t.Run("SetHref_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsLinksItem{}
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
		obj := &FacebookAdsLinksItem{}
		var fernTestValueMethod *FacebookAdsLinksItemMethod

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
		obj := &FacebookAdsLinksItem{}
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
		obj := &FacebookAdsLinksItem{}
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
		obj := &FacebookAdsLinksItem{}
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

func TestSettersFacebookAdsSite(t *testing.T) {
	t.Run("SetID", func(t *testing.T) {
		obj := &FacebookAdsSite{}
		var fernTestValueID *int
		obj.SetID(fernTestValueID)
		assert.Equal(t, fernTestValueID, obj.ID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetName", func(t *testing.T) {
		obj := &FacebookAdsSite{}
		var fernTestValueName *string
		obj.SetName(fernTestValueName)
		assert.Equal(t, fernTestValueName, obj.Name)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetURL", func(t *testing.T) {
		obj := &FacebookAdsSite{}
		var fernTestValueURL *string
		obj.SetURL(fernTestValueURL)
		assert.Equal(t, fernTestValueURL, obj.URL)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersFacebookAdsSite(t *testing.T) {
	t.Run("GetID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsSite{}
		var expected *int
		obj.ID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetID(), "getter should return the property value")
	})

	t.Run("GetID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsSite{}
		obj.ID = nil

		// Act & Assert
		assert.Nil(t, obj.GetID(), "getter should return nil when property is nil")
	})

	t.Run("GetID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsSite
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetID() // Should return zero value
	})

	t.Run("GetName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsSite{}
		var expected *string
		obj.Name = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetName(), "getter should return the property value")
	})

	t.Run("GetName_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsSite{}
		obj.Name = nil

		// Act & Assert
		assert.Nil(t, obj.GetName(), "getter should return nil when property is nil")
	})

	t.Run("GetName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsSite
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetName() // Should return zero value
	})

	t.Run("GetURL", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsSite{}
		var expected *string
		obj.URL = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetURL(), "getter should return the property value")
	})

	t.Run("GetURL_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsSite{}
		obj.URL = nil

		// Act & Assert
		assert.Nil(t, obj.GetURL(), "getter should return nil when property is nil")
	})

	t.Run("GetURL_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsSite
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetURL() // Should return zero value
	})

}

func TestSettersMarkExplicitFacebookAdsSite(t *testing.T) {
	t.Run("SetID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsSite{}
		var fernTestValueID *int

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

	t.Run("SetName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsSite{}
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

	t.Run("SetURL_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsSite{}
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

func TestSettersListFacebookAdsResponse(t *testing.T) {
	t.Run("SetLinks", func(t *testing.T) {
		obj := &ListFacebookAdsResponse{}
		var fernTestValueLinks []*ListFacebookAdsResponseLinksItem
		obj.SetLinks(fernTestValueLinks)
		assert.Equal(t, fernTestValueLinks, obj.Links)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetFacebookAds", func(t *testing.T) {
		obj := &ListFacebookAdsResponse{}
		var fernTestValueFacebookAds []*FacebookAds
		obj.SetFacebookAds(fernTestValueFacebookAds)
		assert.Equal(t, fernTestValueFacebookAds, obj.FacebookAds)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTotalItems", func(t *testing.T) {
		obj := &ListFacebookAdsResponse{}
		var fernTestValueTotalItems *int
		obj.SetTotalItems(fernTestValueTotalItems)
		assert.Equal(t, fernTestValueTotalItems, obj.TotalItems)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersListFacebookAdsResponse(t *testing.T) {
	t.Run("GetLinks", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListFacebookAdsResponse{}
		var expected []*ListFacebookAdsResponseLinksItem
		obj.Links = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetLinks(), "getter should return the property value")
	})

	t.Run("GetLinks_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListFacebookAdsResponse{}
		obj.Links = nil

		// Act & Assert
		assert.Nil(t, obj.GetLinks(), "getter should return nil when property is nil")
	})

	t.Run("GetLinks_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListFacebookAdsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetLinks() // Should return zero value
	})

	t.Run("GetFacebookAds", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListFacebookAdsResponse{}
		var expected []*FacebookAds
		obj.FacebookAds = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetFacebookAds(), "getter should return the property value")
	})

	t.Run("GetFacebookAds_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListFacebookAdsResponse{}
		obj.FacebookAds = nil

		// Act & Assert
		assert.Nil(t, obj.GetFacebookAds(), "getter should return nil when property is nil")
	})

	t.Run("GetFacebookAds_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListFacebookAdsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetFacebookAds() // Should return zero value
	})

	t.Run("GetTotalItems", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListFacebookAdsResponse{}
		var expected *int
		obj.TotalItems = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTotalItems(), "getter should return the property value")
	})

	t.Run("GetTotalItems_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListFacebookAdsResponse{}
		obj.TotalItems = nil

		// Act & Assert
		assert.Nil(t, obj.GetTotalItems(), "getter should return nil when property is nil")
	})

	t.Run("GetTotalItems_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListFacebookAdsResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTotalItems() // Should return zero value
	})

}

func TestSettersMarkExplicitListFacebookAdsResponse(t *testing.T) {
	t.Run("SetLinks_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListFacebookAdsResponse{}
		var fernTestValueLinks []*ListFacebookAdsResponseLinksItem

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

	t.Run("SetFacebookAds_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListFacebookAdsResponse{}
		var fernTestValueFacebookAds []*FacebookAds

		// Act
		obj.SetFacebookAds(fernTestValueFacebookAds)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
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
		obj := &ListFacebookAdsResponse{}
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

func TestSettersListFacebookAdsResponseLinksItem(t *testing.T) {
	t.Run("SetHref", func(t *testing.T) {
		obj := &ListFacebookAdsResponseLinksItem{}
		var fernTestValueHref *string
		obj.SetHref(fernTestValueHref)
		assert.Equal(t, fernTestValueHref, obj.Href)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMethod", func(t *testing.T) {
		obj := &ListFacebookAdsResponseLinksItem{}
		var fernTestValueMethod *ListFacebookAdsResponseLinksItemMethod
		obj.SetMethod(fernTestValueMethod)
		assert.Equal(t, fernTestValueMethod, obj.Method)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRel", func(t *testing.T) {
		obj := &ListFacebookAdsResponseLinksItem{}
		var fernTestValueRel *string
		obj.SetRel(fernTestValueRel)
		assert.Equal(t, fernTestValueRel, obj.Rel)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSchema", func(t *testing.T) {
		obj := &ListFacebookAdsResponseLinksItem{}
		var fernTestValueSchema *string
		obj.SetSchema(fernTestValueSchema)
		assert.Equal(t, fernTestValueSchema, obj.Schema)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTargetSchema", func(t *testing.T) {
		obj := &ListFacebookAdsResponseLinksItem{}
		var fernTestValueTargetSchema *string
		obj.SetTargetSchema(fernTestValueTargetSchema)
		assert.Equal(t, fernTestValueTargetSchema, obj.TargetSchema)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersListFacebookAdsResponseLinksItem(t *testing.T) {
	t.Run("GetHref", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListFacebookAdsResponseLinksItem{}
		var expected *string
		obj.Href = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetHref(), "getter should return the property value")
	})

	t.Run("GetHref_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListFacebookAdsResponseLinksItem{}
		obj.Href = nil

		// Act & Assert
		assert.Nil(t, obj.GetHref(), "getter should return nil when property is nil")
	})

	t.Run("GetHref_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListFacebookAdsResponseLinksItem
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
		obj := &ListFacebookAdsResponseLinksItem{}
		var expected *ListFacebookAdsResponseLinksItemMethod
		obj.Method = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMethod(), "getter should return the property value")
	})

	t.Run("GetMethod_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListFacebookAdsResponseLinksItem{}
		obj.Method = nil

		// Act & Assert
		assert.Nil(t, obj.GetMethod(), "getter should return nil when property is nil")
	})

	t.Run("GetMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListFacebookAdsResponseLinksItem
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
		obj := &ListFacebookAdsResponseLinksItem{}
		var expected *string
		obj.Rel = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetRel(), "getter should return the property value")
	})

	t.Run("GetRel_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListFacebookAdsResponseLinksItem{}
		obj.Rel = nil

		// Act & Assert
		assert.Nil(t, obj.GetRel(), "getter should return nil when property is nil")
	})

	t.Run("GetRel_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListFacebookAdsResponseLinksItem
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
		obj := &ListFacebookAdsResponseLinksItem{}
		var expected *string
		obj.Schema = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSchema(), "getter should return the property value")
	})

	t.Run("GetSchema_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListFacebookAdsResponseLinksItem{}
		obj.Schema = nil

		// Act & Assert
		assert.Nil(t, obj.GetSchema(), "getter should return nil when property is nil")
	})

	t.Run("GetSchema_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListFacebookAdsResponseLinksItem
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
		obj := &ListFacebookAdsResponseLinksItem{}
		var expected *string
		obj.TargetSchema = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTargetSchema(), "getter should return the property value")
	})

	t.Run("GetTargetSchema_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListFacebookAdsResponseLinksItem{}
		obj.TargetSchema = nil

		// Act & Assert
		assert.Nil(t, obj.GetTargetSchema(), "getter should return nil when property is nil")
	})

	t.Run("GetTargetSchema_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListFacebookAdsResponseLinksItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTargetSchema() // Should return zero value
	})

}

func TestSettersMarkExplicitListFacebookAdsResponseLinksItem(t *testing.T) {
	t.Run("SetHref_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListFacebookAdsResponseLinksItem{}
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
		obj := &ListFacebookAdsResponseLinksItem{}
		var fernTestValueMethod *ListFacebookAdsResponseLinksItemMethod

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
		obj := &ListFacebookAdsResponseLinksItem{}
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
		obj := &ListFacebookAdsResponseLinksItem{}
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
		obj := &ListFacebookAdsResponseLinksItem{}
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

func TestJSONMarshalingFacebookAd(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAd{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled FacebookAd
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAd
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAd
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingFacebookAdReportSummary(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummary{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled FacebookAdReportSummary
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAdReportSummary
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAdReportSummary
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingFacebookAdReportSummaryEcommerce(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdReportSummaryEcommerce{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled FacebookAdReportSummaryEcommerce
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAdReportSummaryEcommerce
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAdReportSummaryEcommerce
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingFacebookAds(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAds{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled FacebookAds
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAds
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAds
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingFacebookAdsAudience(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudience{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled FacebookAdsAudience
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAdsAudience
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAdsAudience
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingFacebookAdsAudienceEmailSource(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceEmailSource{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled FacebookAdsAudienceEmailSource
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAdsAudienceEmailSource
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAdsAudienceEmailSource
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingFacebookAdsAudienceTargetingSpecs(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecs{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled FacebookAdsAudienceTargetingSpecs
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAdsAudienceTargetingSpecs
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAdsAudienceTargetingSpecs
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingFacebookAdsAudienceTargetingSpecsInterestsItem(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecsInterestsItem{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled FacebookAdsAudienceTargetingSpecsInterestsItem
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAdsAudienceTargetingSpecsInterestsItem
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAdsAudienceTargetingSpecsInterestsItem
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingFacebookAdsAudienceTargetingSpecsLocations(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsAudienceTargetingSpecsLocations{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled FacebookAdsAudienceTargetingSpecsLocations
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAdsAudienceTargetingSpecsLocations
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAdsAudienceTargetingSpecsLocations
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingFacebookAdsBudget(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsBudget{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled FacebookAdsBudget
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAdsBudget
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAdsBudget
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingFacebookAdsChannel(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsChannel{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled FacebookAdsChannel
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAdsChannel
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAdsChannel
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingFacebookAdsContent(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContent{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled FacebookAdsContent
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAdsContent
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAdsContent
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingFacebookAdsContentAttachmentsItem(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsContentAttachmentsItem{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled FacebookAdsContentAttachmentsItem
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAdsContentAttachmentsItem
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAdsContentAttachmentsItem
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingFacebookAdsFeedback(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsFeedback{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled FacebookAdsFeedback
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAdsFeedback
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAdsFeedback
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingFacebookAdsLinksItem(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsLinksItem{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled FacebookAdsLinksItem
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAdsLinksItem
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAdsLinksItem
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingFacebookAdsSite(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &FacebookAdsSite{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled FacebookAdsSite
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAdsSite
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj FacebookAdsSite
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingListFacebookAdsResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListFacebookAdsResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ListFacebookAdsResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ListFacebookAdsResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ListFacebookAdsResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingListFacebookAdsResponseLinksItem(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListFacebookAdsResponseLinksItem{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ListFacebookAdsResponseLinksItem
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ListFacebookAdsResponseLinksItem
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ListFacebookAdsResponseLinksItem
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestStringFacebookAd(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAd{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAd
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringFacebookAdReportSummary(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAdReportSummary{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdReportSummary
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringFacebookAdReportSummaryEcommerce(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAdReportSummaryEcommerce{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdReportSummaryEcommerce
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringFacebookAds(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAds{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAds
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringFacebookAdsAudience(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAdsAudience{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsAudience
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringFacebookAdsAudienceEmailSource(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAdsAudienceEmailSource{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsAudienceEmailSource
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringFacebookAdsAudienceTargetingSpecs(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAdsAudienceTargetingSpecs{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsAudienceTargetingSpecs
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringFacebookAdsAudienceTargetingSpecsInterestsItem(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAdsAudienceTargetingSpecsInterestsItem{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsAudienceTargetingSpecsInterestsItem
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringFacebookAdsAudienceTargetingSpecsLocations(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAdsAudienceTargetingSpecsLocations{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsAudienceTargetingSpecsLocations
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringFacebookAdsBudget(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAdsBudget{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsBudget
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringFacebookAdsChannel(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAdsChannel{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsChannel
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringFacebookAdsContent(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAdsContent{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsContent
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringFacebookAdsContentAttachmentsItem(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAdsContentAttachmentsItem{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsContentAttachmentsItem
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringFacebookAdsFeedback(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAdsFeedback{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsFeedback
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringFacebookAdsLinksItem(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAdsLinksItem{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsLinksItem
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringFacebookAdsSite(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAdsSite{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *FacebookAdsSite
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringListFacebookAdsResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ListFacebookAdsResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListFacebookAdsResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringListFacebookAdsResponseLinksItem(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ListFacebookAdsResponseLinksItem{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListFacebookAdsResponseLinksItem
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestEnumFacebookAdsAudienceSourceType(t *testing.T) {
	t.Run("NewFromString_facebook", func(t *testing.T) {
		t.Parallel()
		val, err := NewFacebookAdsAudienceSourceTypeFromString("facebook")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, FacebookAdsAudienceSourceType("facebook"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_list", func(t *testing.T) {
		t.Parallel()
		val, err := NewFacebookAdsAudienceSourceTypeFromString("list")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, FacebookAdsAudienceSourceType("list"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewFacebookAdsAudienceSourceTypeFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewFacebookAdsAudienceSourceTypeFromString("facebook")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumFacebookAdsAudienceType(t *testing.T) {
	t.Run("NewFromString_Custom_Audience", func(t *testing.T) {
		t.Parallel()
		val, err := NewFacebookAdsAudienceTypeFromString("Custom Audience")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, FacebookAdsAudienceType("Custom Audience"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Lookalike_Audience", func(t *testing.T) {
		t.Parallel()
		val, err := NewFacebookAdsAudienceTypeFromString("Lookalike Audience")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, FacebookAdsAudienceType("Lookalike Audience"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Interest_based_Audience", func(t *testing.T) {
		t.Parallel()
		val, err := NewFacebookAdsAudienceTypeFromString("Interest-based Audience")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, FacebookAdsAudienceType("Interest-based Audience"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewFacebookAdsAudienceTypeFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewFacebookAdsAudienceTypeFromString("Custom Audience")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumFacebookAdsLinksItemMethod(t *testing.T) {
	t.Run("NewFromString_GET", func(t *testing.T) {
		t.Parallel()
		val, err := NewFacebookAdsLinksItemMethodFromString("GET")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, FacebookAdsLinksItemMethod("GET"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_POST", func(t *testing.T) {
		t.Parallel()
		val, err := NewFacebookAdsLinksItemMethodFromString("POST")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, FacebookAdsLinksItemMethod("POST"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_PUT", func(t *testing.T) {
		t.Parallel()
		val, err := NewFacebookAdsLinksItemMethodFromString("PUT")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, FacebookAdsLinksItemMethod("PUT"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_PATCH", func(t *testing.T) {
		t.Parallel()
		val, err := NewFacebookAdsLinksItemMethodFromString("PATCH")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, FacebookAdsLinksItemMethod("PATCH"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_DELETE", func(t *testing.T) {
		t.Parallel()
		val, err := NewFacebookAdsLinksItemMethodFromString("DELETE")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, FacebookAdsLinksItemMethod("DELETE"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_OPTIONS", func(t *testing.T) {
		t.Parallel()
		val, err := NewFacebookAdsLinksItemMethodFromString("OPTIONS")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, FacebookAdsLinksItemMethod("OPTIONS"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_HEAD", func(t *testing.T) {
		t.Parallel()
		val, err := NewFacebookAdsLinksItemMethodFromString("HEAD")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, FacebookAdsLinksItemMethod("HEAD"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewFacebookAdsLinksItemMethodFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewFacebookAdsLinksItemMethodFromString("GET")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumListFacebookAdsRequestSortDir(t *testing.T) {
	t.Run("NewFromString_ASC", func(t *testing.T) {
		t.Parallel()
		val, err := NewListFacebookAdsRequestSortDirFromString("ASC")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListFacebookAdsRequestSortDir("ASC"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_DESC", func(t *testing.T) {
		t.Parallel()
		val, err := NewListFacebookAdsRequestSortDirFromString("DESC")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListFacebookAdsRequestSortDir("DESC"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewListFacebookAdsRequestSortDirFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewListFacebookAdsRequestSortDirFromString("ASC")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumListFacebookAdsRequestSortField(t *testing.T) {
	t.Run("NewFromString_created_at", func(t *testing.T) {
		t.Parallel()
		val, err := NewListFacebookAdsRequestSortFieldFromString("created_at")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListFacebookAdsRequestSortField("created_at"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_updated_at", func(t *testing.T) {
		t.Parallel()
		val, err := NewListFacebookAdsRequestSortFieldFromString("updated_at")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListFacebookAdsRequestSortField("updated_at"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_end_time", func(t *testing.T) {
		t.Parallel()
		val, err := NewListFacebookAdsRequestSortFieldFromString("end_time")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListFacebookAdsRequestSortField("end_time"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewListFacebookAdsRequestSortFieldFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewListFacebookAdsRequestSortFieldFromString("created_at")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumListFacebookAdsResponseLinksItemMethod(t *testing.T) {
	t.Run("NewFromString_GET", func(t *testing.T) {
		t.Parallel()
		val, err := NewListFacebookAdsResponseLinksItemMethodFromString("GET")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListFacebookAdsResponseLinksItemMethod("GET"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_POST", func(t *testing.T) {
		t.Parallel()
		val, err := NewListFacebookAdsResponseLinksItemMethodFromString("POST")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListFacebookAdsResponseLinksItemMethod("POST"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_PUT", func(t *testing.T) {
		t.Parallel()
		val, err := NewListFacebookAdsResponseLinksItemMethodFromString("PUT")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListFacebookAdsResponseLinksItemMethod("PUT"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_PATCH", func(t *testing.T) {
		t.Parallel()
		val, err := NewListFacebookAdsResponseLinksItemMethodFromString("PATCH")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListFacebookAdsResponseLinksItemMethod("PATCH"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_DELETE", func(t *testing.T) {
		t.Parallel()
		val, err := NewListFacebookAdsResponseLinksItemMethodFromString("DELETE")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListFacebookAdsResponseLinksItemMethod("DELETE"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_OPTIONS", func(t *testing.T) {
		t.Parallel()
		val, err := NewListFacebookAdsResponseLinksItemMethodFromString("OPTIONS")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListFacebookAdsResponseLinksItemMethod("OPTIONS"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_HEAD", func(t *testing.T) {
		t.Parallel()
		val, err := NewListFacebookAdsResponseLinksItemMethodFromString("HEAD")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListFacebookAdsResponseLinksItemMethod("HEAD"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewListFacebookAdsResponseLinksItemMethodFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewListFacebookAdsResponseLinksItemMethodFromString("GET")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestExtraPropertiesFacebookAd(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAd{}
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
		var obj *FacebookAd
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesFacebookAdReportSummary(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAdReportSummary{}
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
		var obj *FacebookAdReportSummary
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesFacebookAdReportSummaryEcommerce(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAdReportSummaryEcommerce{}
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
		var obj *FacebookAdReportSummaryEcommerce
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesFacebookAds(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAds{}
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
		var obj *FacebookAds
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesFacebookAdsAudience(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAdsAudience{}
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
		var obj *FacebookAdsAudience
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesFacebookAdsAudienceEmailSource(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAdsAudienceEmailSource{}
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
		var obj *FacebookAdsAudienceEmailSource
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesFacebookAdsAudienceTargetingSpecs(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAdsAudienceTargetingSpecs{}
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
		var obj *FacebookAdsAudienceTargetingSpecs
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesFacebookAdsAudienceTargetingSpecsInterestsItem(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAdsAudienceTargetingSpecsInterestsItem{}
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
		var obj *FacebookAdsAudienceTargetingSpecsInterestsItem
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesFacebookAdsAudienceTargetingSpecsLocations(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAdsAudienceTargetingSpecsLocations{}
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
		var obj *FacebookAdsAudienceTargetingSpecsLocations
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesFacebookAdsBudget(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAdsBudget{}
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
		var obj *FacebookAdsBudget
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesFacebookAdsChannel(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAdsChannel{}
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
		var obj *FacebookAdsChannel
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesFacebookAdsContent(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAdsContent{}
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
		var obj *FacebookAdsContent
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesFacebookAdsContentAttachmentsItem(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAdsContentAttachmentsItem{}
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
		var obj *FacebookAdsContentAttachmentsItem
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesFacebookAdsFeedback(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAdsFeedback{}
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
		var obj *FacebookAdsFeedback
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesFacebookAdsLinksItem(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAdsLinksItem{}
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
		var obj *FacebookAdsLinksItem
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesFacebookAdsSite(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &FacebookAdsSite{}
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
		var obj *FacebookAdsSite
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesListFacebookAdsResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ListFacebookAdsResponse{}
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
		var obj *ListFacebookAdsResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesListFacebookAdsResponseLinksItem(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ListFacebookAdsResponseLinksItem{}
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
		var obj *ListFacebookAdsResponseLinksItem
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}
