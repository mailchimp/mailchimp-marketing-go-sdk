// Code generated from our API definition. DO NOT EDIT.

package api

import (
	json "encoding/json"
	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
	testing "testing"
	time "time"
)

func TestSettersListRootRequest(t *testing.T) {
	t.Run("SetFields", func(t *testing.T) {
		obj := &ListRootRequest{}
		var fernTestValueFields []*string
		obj.SetFields(fernTestValueFields)
		assert.Equal(t, fernTestValueFields, obj.Fields)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetExcludeFields", func(t *testing.T) {
		obj := &ListRootRequest{}
		var fernTestValueExcludeFields []*string
		obj.SetExcludeFields(fernTestValueExcludeFields)
		assert.Equal(t, fernTestValueExcludeFields, obj.ExcludeFields)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestSettersMarkExplicitListRootRequest(t *testing.T) {
	t.Run("SetFields_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootRequest{}
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
		obj := &ListRootRequest{}
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

func TestSettersListRootResponse(t *testing.T) {
	t.Run("SetLinks", func(t *testing.T) {
		obj := &ListRootResponse{}
		var fernTestValueLinks []*ListRootResponseLinksItem
		obj.SetLinks(fernTestValueLinks)
		assert.Equal(t, fernTestValueLinks, obj.Links)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetAccountID", func(t *testing.T) {
		obj := &ListRootResponse{}
		var fernTestValueAccountID *string
		obj.SetAccountID(fernTestValueAccountID)
		assert.Equal(t, fernTestValueAccountID, obj.AccountID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetAccountIndustry", func(t *testing.T) {
		obj := &ListRootResponse{}
		var fernTestValueAccountIndustry *string
		obj.SetAccountIndustry(fernTestValueAccountIndustry)
		assert.Equal(t, fernTestValueAccountIndustry, obj.AccountIndustry)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetAccountName", func(t *testing.T) {
		obj := &ListRootResponse{}
		var fernTestValueAccountName *string
		obj.SetAccountName(fernTestValueAccountName)
		assert.Equal(t, fernTestValueAccountName, obj.AccountName)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetAccountTimezone", func(t *testing.T) {
		obj := &ListRootResponse{}
		var fernTestValueAccountTimezone *string
		obj.SetAccountTimezone(fernTestValueAccountTimezone)
		assert.Equal(t, fernTestValueAccountTimezone, obj.AccountTimezone)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetAvatarURL", func(t *testing.T) {
		obj := &ListRootResponse{}
		var fernTestValueAvatarURL *string
		obj.SetAvatarURL(fernTestValueAvatarURL)
		assert.Equal(t, fernTestValueAvatarURL, obj.AvatarURL)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetContact", func(t *testing.T) {
		obj := &ListRootResponse{}
		var fernTestValueContact *ListRootResponseContact
		obj.SetContact(fernTestValueContact)
		assert.Equal(t, fernTestValueContact, obj.Contact)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetEmail", func(t *testing.T) {
		obj := &ListRootResponse{}
		var fernTestValueEmail *string
		obj.SetEmail(fernTestValueEmail)
		assert.Equal(t, fernTestValueEmail, obj.Email)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetFirstName", func(t *testing.T) {
		obj := &ListRootResponse{}
		var fernTestValueFirstName *string
		obj.SetFirstName(fernTestValueFirstName)
		assert.Equal(t, fernTestValueFirstName, obj.FirstName)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetFirstPayment", func(t *testing.T) {
		obj := &ListRootResponse{}
		var fernTestValueFirstPayment *ListRootResponseFirstPayment
		obj.SetFirstPayment(fernTestValueFirstPayment)
		assert.Equal(t, fernTestValueFirstPayment, obj.FirstPayment)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetIndustryStats", func(t *testing.T) {
		obj := &ListRootResponse{}
		var fernTestValueIndustryStats *ListRootResponseIndustryStats
		obj.SetIndustryStats(fernTestValueIndustryStats)
		assert.Equal(t, fernTestValueIndustryStats, obj.IndustryStats)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetLastLogin", func(t *testing.T) {
		obj := &ListRootResponse{}
		var fernTestValueLastLogin *time.Time
		obj.SetLastLogin(fernTestValueLastLogin)
		assert.Equal(t, fernTestValueLastLogin, obj.LastLogin)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetLastName", func(t *testing.T) {
		obj := &ListRootResponse{}
		var fernTestValueLastName *string
		obj.SetLastName(fernTestValueLastName)
		assert.Equal(t, fernTestValueLastName, obj.LastName)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetLoginID", func(t *testing.T) {
		obj := &ListRootResponse{}
		var fernTestValueLoginID *string
		obj.SetLoginID(fernTestValueLoginID)
		assert.Equal(t, fernTestValueLoginID, obj.LoginID)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMemberSince", func(t *testing.T) {
		obj := &ListRootResponse{}
		var fernTestValueMemberSince *time.Time
		obj.SetMemberSince(fernTestValueMemberSince)
		assert.Equal(t, fernTestValueMemberSince, obj.MemberSince)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetPricingPlanType", func(t *testing.T) {
		obj := &ListRootResponse{}
		var fernTestValuePricingPlanType *ListRootResponsePricingPlanType
		obj.SetPricingPlanType(fernTestValuePricingPlanType)
		assert.Equal(t, fernTestValuePricingPlanType, obj.PricingPlanType)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetProEnabled", func(t *testing.T) {
		obj := &ListRootResponse{}
		var fernTestValueProEnabled *bool
		obj.SetProEnabled(fernTestValueProEnabled)
		assert.Equal(t, fernTestValueProEnabled, obj.ProEnabled)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRole", func(t *testing.T) {
		obj := &ListRootResponse{}
		var fernTestValueRole *string
		obj.SetRole(fernTestValueRole)
		assert.Equal(t, fernTestValueRole, obj.Role)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTotalSubscribers", func(t *testing.T) {
		obj := &ListRootResponse{}
		var fernTestValueTotalSubscribers *int
		obj.SetTotalSubscribers(fernTestValueTotalSubscribers)
		assert.Equal(t, fernTestValueTotalSubscribers, obj.TotalSubscribers)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetUsername", func(t *testing.T) {
		obj := &ListRootResponse{}
		var fernTestValueUsername *string
		obj.SetUsername(fernTestValueUsername)
		assert.Equal(t, fernTestValueUsername, obj.Username)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersListRootResponse(t *testing.T) {
	t.Run("GetLinks", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var expected []*ListRootResponseLinksItem
		obj.Links = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetLinks(), "getter should return the property value")
	})

	t.Run("GetLinks_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		obj.Links = nil

		// Act & Assert
		assert.Nil(t, obj.GetLinks(), "getter should return nil when property is nil")
	})

	t.Run("GetLinks_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetLinks() // Should return zero value
	})

	t.Run("GetAccountID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var expected *string
		obj.AccountID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetAccountID(), "getter should return the property value")
	})

	t.Run("GetAccountID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		obj.AccountID = nil

		// Act & Assert
		assert.Nil(t, obj.GetAccountID(), "getter should return nil when property is nil")
	})

	t.Run("GetAccountID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetAccountID() // Should return zero value
	})

	t.Run("GetAccountIndustry", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var expected *string
		obj.AccountIndustry = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetAccountIndustry(), "getter should return the property value")
	})

	t.Run("GetAccountIndustry_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		obj.AccountIndustry = nil

		// Act & Assert
		assert.Nil(t, obj.GetAccountIndustry(), "getter should return nil when property is nil")
	})

	t.Run("GetAccountIndustry_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetAccountIndustry() // Should return zero value
	})

	t.Run("GetAccountName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var expected *string
		obj.AccountName = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetAccountName(), "getter should return the property value")
	})

	t.Run("GetAccountName_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		obj.AccountName = nil

		// Act & Assert
		assert.Nil(t, obj.GetAccountName(), "getter should return nil when property is nil")
	})

	t.Run("GetAccountName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetAccountName() // Should return zero value
	})

	t.Run("GetAccountTimezone", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var expected *string
		obj.AccountTimezone = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetAccountTimezone(), "getter should return the property value")
	})

	t.Run("GetAccountTimezone_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		obj.AccountTimezone = nil

		// Act & Assert
		assert.Nil(t, obj.GetAccountTimezone(), "getter should return nil when property is nil")
	})

	t.Run("GetAccountTimezone_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetAccountTimezone() // Should return zero value
	})

	t.Run("GetAvatarURL", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var expected *string
		obj.AvatarURL = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetAvatarURL(), "getter should return the property value")
	})

	t.Run("GetAvatarURL_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		obj.AvatarURL = nil

		// Act & Assert
		assert.Nil(t, obj.GetAvatarURL(), "getter should return nil when property is nil")
	})

	t.Run("GetAvatarURL_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetAvatarURL() // Should return zero value
	})

	t.Run("GetContact", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var expected *ListRootResponseContact
		obj.Contact = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetContact(), "getter should return the property value")
	})

	t.Run("GetContact_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		obj.Contact = nil

		// Act & Assert
		assert.Nil(t, obj.GetContact(), "getter should return nil when property is nil")
	})

	t.Run("GetContact_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetContact() // Should return zero value
	})

	t.Run("GetEmail", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var expected *string
		obj.Email = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetEmail(), "getter should return the property value")
	})

	t.Run("GetEmail_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		obj.Email = nil

		// Act & Assert
		assert.Nil(t, obj.GetEmail(), "getter should return nil when property is nil")
	})

	t.Run("GetEmail_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetEmail() // Should return zero value
	})

	t.Run("GetFirstName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var expected *string
		obj.FirstName = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetFirstName(), "getter should return the property value")
	})

	t.Run("GetFirstName_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		obj.FirstName = nil

		// Act & Assert
		assert.Nil(t, obj.GetFirstName(), "getter should return nil when property is nil")
	})

	t.Run("GetFirstName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetFirstName() // Should return zero value
	})

	t.Run("GetFirstPayment", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var expected *ListRootResponseFirstPayment
		obj.FirstPayment = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetFirstPayment(), "getter should return the property value")
	})

	t.Run("GetFirstPayment_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		obj.FirstPayment = nil

		// Act & Assert
		assert.Nil(t, obj.GetFirstPayment(), "getter should return nil when property is nil")
	})

	t.Run("GetFirstPayment_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetFirstPayment() // Should return zero value
	})

	t.Run("GetIndustryStats", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var expected *ListRootResponseIndustryStats
		obj.IndustryStats = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetIndustryStats(), "getter should return the property value")
	})

	t.Run("GetIndustryStats_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		obj.IndustryStats = nil

		// Act & Assert
		assert.Nil(t, obj.GetIndustryStats(), "getter should return nil when property is nil")
	})

	t.Run("GetIndustryStats_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetIndustryStats() // Should return zero value
	})

	t.Run("GetLastLogin", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var expected *time.Time
		obj.LastLogin = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetLastLogin(), "getter should return the property value")
	})

	t.Run("GetLastLogin_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		obj.LastLogin = nil

		// Act & Assert
		assert.Nil(t, obj.GetLastLogin(), "getter should return nil when property is nil")
	})

	t.Run("GetLastLogin_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetLastLogin() // Should return zero value
	})

	t.Run("GetLastName", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var expected *string
		obj.LastName = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetLastName(), "getter should return the property value")
	})

	t.Run("GetLastName_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		obj.LastName = nil

		// Act & Assert
		assert.Nil(t, obj.GetLastName(), "getter should return nil when property is nil")
	})

	t.Run("GetLastName_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetLastName() // Should return zero value
	})

	t.Run("GetLoginID", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var expected *string
		obj.LoginID = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetLoginID(), "getter should return the property value")
	})

	t.Run("GetLoginID_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		obj.LoginID = nil

		// Act & Assert
		assert.Nil(t, obj.GetLoginID(), "getter should return nil when property is nil")
	})

	t.Run("GetLoginID_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetLoginID() // Should return zero value
	})

	t.Run("GetMemberSince", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var expected *time.Time
		obj.MemberSince = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMemberSince(), "getter should return the property value")
	})

	t.Run("GetMemberSince_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		obj.MemberSince = nil

		// Act & Assert
		assert.Nil(t, obj.GetMemberSince(), "getter should return nil when property is nil")
	})

	t.Run("GetMemberSince_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetMemberSince() // Should return zero value
	})

	t.Run("GetPricingPlanType", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var expected *ListRootResponsePricingPlanType
		obj.PricingPlanType = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetPricingPlanType(), "getter should return the property value")
	})

	t.Run("GetPricingPlanType_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		obj.PricingPlanType = nil

		// Act & Assert
		assert.Nil(t, obj.GetPricingPlanType(), "getter should return nil when property is nil")
	})

	t.Run("GetPricingPlanType_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetPricingPlanType() // Should return zero value
	})

	t.Run("GetProEnabled", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var expected *bool
		obj.ProEnabled = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetProEnabled(), "getter should return the property value")
	})

	t.Run("GetProEnabled_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		obj.ProEnabled = nil

		// Act & Assert
		assert.Nil(t, obj.GetProEnabled(), "getter should return nil when property is nil")
	})

	t.Run("GetProEnabled_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetProEnabled() // Should return zero value
	})

	t.Run("GetRole", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var expected *string
		obj.Role = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetRole(), "getter should return the property value")
	})

	t.Run("GetRole_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		obj.Role = nil

		// Act & Assert
		assert.Nil(t, obj.GetRole(), "getter should return nil when property is nil")
	})

	t.Run("GetRole_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetRole() // Should return zero value
	})

	t.Run("GetTotalSubscribers", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var expected *int
		obj.TotalSubscribers = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTotalSubscribers(), "getter should return the property value")
	})

	t.Run("GetTotalSubscribers_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		obj.TotalSubscribers = nil

		// Act & Assert
		assert.Nil(t, obj.GetTotalSubscribers(), "getter should return nil when property is nil")
	})

	t.Run("GetTotalSubscribers_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTotalSubscribers() // Should return zero value
	})

	t.Run("GetUsername", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var expected *string
		obj.Username = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetUsername(), "getter should return the property value")
	})

	t.Run("GetUsername_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		obj.Username = nil

		// Act & Assert
		assert.Nil(t, obj.GetUsername(), "getter should return nil when property is nil")
	})

	t.Run("GetUsername_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponse
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetUsername() // Should return zero value
	})

}

func TestSettersMarkExplicitListRootResponse(t *testing.T) {
	t.Run("SetLinks_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var fernTestValueLinks []*ListRootResponseLinksItem

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

	t.Run("SetAccountID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var fernTestValueAccountID *string

		// Act
		obj.SetAccountID(fernTestValueAccountID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetAccountIndustry_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var fernTestValueAccountIndustry *string

		// Act
		obj.SetAccountIndustry(fernTestValueAccountIndustry)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetAccountName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var fernTestValueAccountName *string

		// Act
		obj.SetAccountName(fernTestValueAccountName)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetAccountTimezone_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var fernTestValueAccountTimezone *string

		// Act
		obj.SetAccountTimezone(fernTestValueAccountTimezone)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetAvatarURL_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var fernTestValueAvatarURL *string

		// Act
		obj.SetAvatarURL(fernTestValueAvatarURL)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetContact_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var fernTestValueContact *ListRootResponseContact

		// Act
		obj.SetContact(fernTestValueContact)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
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
		obj := &ListRootResponse{}
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

	t.Run("SetFirstName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var fernTestValueFirstName *string

		// Act
		obj.SetFirstName(fernTestValueFirstName)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetFirstPayment_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var fernTestValueFirstPayment *ListRootResponseFirstPayment

		// Act
		obj.SetFirstPayment(fernTestValueFirstPayment)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetIndustryStats_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var fernTestValueIndustryStats *ListRootResponseIndustryStats

		// Act
		obj.SetIndustryStats(fernTestValueIndustryStats)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetLastLogin_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var fernTestValueLastLogin *time.Time

		// Act
		obj.SetLastLogin(fernTestValueLastLogin)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetLastName_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var fernTestValueLastName *string

		// Act
		obj.SetLastName(fernTestValueLastName)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetLoginID_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var fernTestValueLoginID *string

		// Act
		obj.SetLoginID(fernTestValueLoginID)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetMemberSince_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var fernTestValueMemberSince *time.Time

		// Act
		obj.SetMemberSince(fernTestValueMemberSince)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetPricingPlanType_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var fernTestValuePricingPlanType *ListRootResponsePricingPlanType

		// Act
		obj.SetPricingPlanType(fernTestValuePricingPlanType)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetProEnabled_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var fernTestValueProEnabled *bool

		// Act
		obj.SetProEnabled(fernTestValueProEnabled)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetRole_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var fernTestValueRole *string

		// Act
		obj.SetRole(fernTestValueRole)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetTotalSubscribers_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var fernTestValueTotalSubscribers *int

		// Act
		obj.SetTotalSubscribers(fernTestValueTotalSubscribers)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetUsername_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}
		var fernTestValueUsername *string

		// Act
		obj.SetUsername(fernTestValueUsername)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
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

func TestSettersListRootResponseContact(t *testing.T) {
	t.Run("SetAddr1", func(t *testing.T) {
		obj := &ListRootResponseContact{}
		var fernTestValueAddr1 *string
		obj.SetAddr1(fernTestValueAddr1)
		assert.Equal(t, fernTestValueAddr1, obj.Addr1)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetAddr2", func(t *testing.T) {
		obj := &ListRootResponseContact{}
		var fernTestValueAddr2 *string
		obj.SetAddr2(fernTestValueAddr2)
		assert.Equal(t, fernTestValueAddr2, obj.Addr2)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCity", func(t *testing.T) {
		obj := &ListRootResponseContact{}
		var fernTestValueCity *string
		obj.SetCity(fernTestValueCity)
		assert.Equal(t, fernTestValueCity, obj.City)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCompany", func(t *testing.T) {
		obj := &ListRootResponseContact{}
		var fernTestValueCompany *string
		obj.SetCompany(fernTestValueCompany)
		assert.Equal(t, fernTestValueCompany, obj.Company)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetCountry", func(t *testing.T) {
		obj := &ListRootResponseContact{}
		var fernTestValueCountry *string
		obj.SetCountry(fernTestValueCountry)
		assert.Equal(t, fernTestValueCountry, obj.Country)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetState", func(t *testing.T) {
		obj := &ListRootResponseContact{}
		var fernTestValueState *string
		obj.SetState(fernTestValueState)
		assert.Equal(t, fernTestValueState, obj.State)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetZip", func(t *testing.T) {
		obj := &ListRootResponseContact{}
		var fernTestValueZip *string
		obj.SetZip(fernTestValueZip)
		assert.Equal(t, fernTestValueZip, obj.Zip)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersListRootResponseContact(t *testing.T) {
	t.Run("GetAddr1", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseContact{}
		var expected *string
		obj.Addr1 = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetAddr1(), "getter should return the property value")
	})

	t.Run("GetAddr1_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseContact{}
		obj.Addr1 = nil

		// Act & Assert
		assert.Nil(t, obj.GetAddr1(), "getter should return nil when property is nil")
	})

	t.Run("GetAddr1_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponseContact
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
		obj := &ListRootResponseContact{}
		var expected *string
		obj.Addr2 = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetAddr2(), "getter should return the property value")
	})

	t.Run("GetAddr2_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseContact{}
		obj.Addr2 = nil

		// Act & Assert
		assert.Nil(t, obj.GetAddr2(), "getter should return nil when property is nil")
	})

	t.Run("GetAddr2_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponseContact
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
		obj := &ListRootResponseContact{}
		var expected *string
		obj.City = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCity(), "getter should return the property value")
	})

	t.Run("GetCity_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseContact{}
		obj.City = nil

		// Act & Assert
		assert.Nil(t, obj.GetCity(), "getter should return nil when property is nil")
	})

	t.Run("GetCity_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponseContact
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCity() // Should return zero value
	})

	t.Run("GetCompany", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseContact{}
		var expected *string
		obj.Company = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCompany(), "getter should return the property value")
	})

	t.Run("GetCompany_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseContact{}
		obj.Company = nil

		// Act & Assert
		assert.Nil(t, obj.GetCompany(), "getter should return nil when property is nil")
	})

	t.Run("GetCompany_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponseContact
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCompany() // Should return zero value
	})

	t.Run("GetCountry", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseContact{}
		var expected *string
		obj.Country = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetCountry(), "getter should return the property value")
	})

	t.Run("GetCountry_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseContact{}
		obj.Country = nil

		// Act & Assert
		assert.Nil(t, obj.GetCountry(), "getter should return nil when property is nil")
	})

	t.Run("GetCountry_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponseContact
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetCountry() // Should return zero value
	})

	t.Run("GetState", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseContact{}
		var expected *string
		obj.State = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetState(), "getter should return the property value")
	})

	t.Run("GetState_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseContact{}
		obj.State = nil

		// Act & Assert
		assert.Nil(t, obj.GetState(), "getter should return nil when property is nil")
	})

	t.Run("GetState_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponseContact
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
		obj := &ListRootResponseContact{}
		var expected *string
		obj.Zip = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetZip(), "getter should return the property value")
	})

	t.Run("GetZip_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseContact{}
		obj.Zip = nil

		// Act & Assert
		assert.Nil(t, obj.GetZip(), "getter should return nil when property is nil")
	})

	t.Run("GetZip_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponseContact
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetZip() // Should return zero value
	})

}

func TestSettersMarkExplicitListRootResponseContact(t *testing.T) {
	t.Run("SetAddr1_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseContact{}
		var fernTestValueAddr1 *string

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
		obj := &ListRootResponseContact{}
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
		obj := &ListRootResponseContact{}
		var fernTestValueCity *string

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

	t.Run("SetCompany_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseContact{}
		var fernTestValueCompany *string

		// Act
		obj.SetCompany(fernTestValueCompany)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
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
		obj := &ListRootResponseContact{}
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

	t.Run("SetState_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseContact{}
		var fernTestValueState *string

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
		obj := &ListRootResponseContact{}
		var fernTestValueZip *string

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

}

func TestGettersListRootResponseFirstPayment(t *testing.T) {
	t.Run("GetDateTime", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseFirstPayment{}
		var expected time.Time
		obj.DateTime = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetDateTime(), "getter should return the property value")
	})

	t.Run("GetDateTime_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponseFirstPayment
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetDateTime() // Should return zero value
	})

	t.Run("GetListRootResponseFirstPaymentOne", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseFirstPayment{}
		var expected ListRootResponseFirstPaymentOne
		obj.ListRootResponseFirstPaymentOne = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetListRootResponseFirstPaymentOne(), "getter should return the property value")
	})

	t.Run("GetListRootResponseFirstPaymentOne_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponseFirstPayment
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetListRootResponseFirstPaymentOne() // Should return zero value
	})

}

func TestSettersListRootResponseIndustryStats(t *testing.T) {
	t.Run("SetBounceRate", func(t *testing.T) {
		obj := &ListRootResponseIndustryStats{}
		var fernTestValueBounceRate *float64
		obj.SetBounceRate(fernTestValueBounceRate)
		assert.Equal(t, fernTestValueBounceRate, obj.BounceRate)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetClickRate", func(t *testing.T) {
		obj := &ListRootResponseIndustryStats{}
		var fernTestValueClickRate *float64
		obj.SetClickRate(fernTestValueClickRate)
		assert.Equal(t, fernTestValueClickRate, obj.ClickRate)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetOpenRate", func(t *testing.T) {
		obj := &ListRootResponseIndustryStats{}
		var fernTestValueOpenRate *float64
		obj.SetOpenRate(fernTestValueOpenRate)
		assert.Equal(t, fernTestValueOpenRate, obj.OpenRate)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersListRootResponseIndustryStats(t *testing.T) {
	t.Run("GetBounceRate", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseIndustryStats{}
		var expected *float64
		obj.BounceRate = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetBounceRate(), "getter should return the property value")
	})

	t.Run("GetBounceRate_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseIndustryStats{}
		obj.BounceRate = nil

		// Act & Assert
		assert.Nil(t, obj.GetBounceRate(), "getter should return nil when property is nil")
	})

	t.Run("GetBounceRate_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponseIndustryStats
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetBounceRate() // Should return zero value
	})

	t.Run("GetClickRate", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseIndustryStats{}
		var expected *float64
		obj.ClickRate = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetClickRate(), "getter should return the property value")
	})

	t.Run("GetClickRate_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseIndustryStats{}
		obj.ClickRate = nil

		// Act & Assert
		assert.Nil(t, obj.GetClickRate(), "getter should return nil when property is nil")
	})

	t.Run("GetClickRate_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponseIndustryStats
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetClickRate() // Should return zero value
	})

	t.Run("GetOpenRate", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseIndustryStats{}
		var expected *float64
		obj.OpenRate = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetOpenRate(), "getter should return the property value")
	})

	t.Run("GetOpenRate_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseIndustryStats{}
		obj.OpenRate = nil

		// Act & Assert
		assert.Nil(t, obj.GetOpenRate(), "getter should return nil when property is nil")
	})

	t.Run("GetOpenRate_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponseIndustryStats
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetOpenRate() // Should return zero value
	})

}

func TestSettersMarkExplicitListRootResponseIndustryStats(t *testing.T) {
	t.Run("SetBounceRate_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseIndustryStats{}
		var fernTestValueBounceRate *float64

		// Act
		obj.SetBounceRate(fernTestValueBounceRate)

		// Assert - object with explicitly set field can be marshaled/unmarshaled
		bytes, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed for test setup")

		// This test ensures JSON marshaling and unmarshaling succeed when the field has a zero/nil value
		// Detect if marshaled JSON is an object or primitive to use correct unmarshal target
		if len(bytes) > 0 && bytes[0] == '{' {
			// JSON object - unmarshal into map
			var unmarshaled map[string]interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		} else {
			// JSON primitive (string, number, boolean, null) - unmarshal into interface{}
			var unmarshaled interface{}
			err = json.Unmarshal(bytes, &unmarshaled)
			require.NoError(t, err, "unmarshaling should succeed for test verification")
		}

		// Note: This does not explicitly assert the presence of a specific JSON field
		// It verifies that setting a field via setter allows successful JSON round-trip
	})

	t.Run("SetClickRate_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseIndustryStats{}
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

	t.Run("SetOpenRate_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseIndustryStats{}
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

}

func TestSettersListRootResponseLinksItem(t *testing.T) {
	t.Run("SetHref", func(t *testing.T) {
		obj := &ListRootResponseLinksItem{}
		var fernTestValueHref *string
		obj.SetHref(fernTestValueHref)
		assert.Equal(t, fernTestValueHref, obj.Href)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetMethod", func(t *testing.T) {
		obj := &ListRootResponseLinksItem{}
		var fernTestValueMethod *ListRootResponseLinksItemMethod
		obj.SetMethod(fernTestValueMethod)
		assert.Equal(t, fernTestValueMethod, obj.Method)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetRel", func(t *testing.T) {
		obj := &ListRootResponseLinksItem{}
		var fernTestValueRel *string
		obj.SetRel(fernTestValueRel)
		assert.Equal(t, fernTestValueRel, obj.Rel)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetSchema", func(t *testing.T) {
		obj := &ListRootResponseLinksItem{}
		var fernTestValueSchema *string
		obj.SetSchema(fernTestValueSchema)
		assert.Equal(t, fernTestValueSchema, obj.Schema)
		assert.NotNil(t, obj.explicitFields)
	})

	t.Run("SetTargetSchema", func(t *testing.T) {
		obj := &ListRootResponseLinksItem{}
		var fernTestValueTargetSchema *string
		obj.SetTargetSchema(fernTestValueTargetSchema)
		assert.Equal(t, fernTestValueTargetSchema, obj.TargetSchema)
		assert.NotNil(t, obj.explicitFields)
	})

}

func TestGettersListRootResponseLinksItem(t *testing.T) {
	t.Run("GetHref", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseLinksItem{}
		var expected *string
		obj.Href = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetHref(), "getter should return the property value")
	})

	t.Run("GetHref_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseLinksItem{}
		obj.Href = nil

		// Act & Assert
		assert.Nil(t, obj.GetHref(), "getter should return nil when property is nil")
	})

	t.Run("GetHref_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponseLinksItem
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
		obj := &ListRootResponseLinksItem{}
		var expected *ListRootResponseLinksItemMethod
		obj.Method = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetMethod(), "getter should return the property value")
	})

	t.Run("GetMethod_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseLinksItem{}
		obj.Method = nil

		// Act & Assert
		assert.Nil(t, obj.GetMethod(), "getter should return nil when property is nil")
	})

	t.Run("GetMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponseLinksItem
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
		obj := &ListRootResponseLinksItem{}
		var expected *string
		obj.Rel = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetRel(), "getter should return the property value")
	})

	t.Run("GetRel_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseLinksItem{}
		obj.Rel = nil

		// Act & Assert
		assert.Nil(t, obj.GetRel(), "getter should return nil when property is nil")
	})

	t.Run("GetRel_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponseLinksItem
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
		obj := &ListRootResponseLinksItem{}
		var expected *string
		obj.Schema = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetSchema(), "getter should return the property value")
	})

	t.Run("GetSchema_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseLinksItem{}
		obj.Schema = nil

		// Act & Assert
		assert.Nil(t, obj.GetSchema(), "getter should return nil when property is nil")
	})

	t.Run("GetSchema_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponseLinksItem
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
		obj := &ListRootResponseLinksItem{}
		var expected *string
		obj.TargetSchema = expected

		// Act & Assert
		assert.Equal(t, expected, obj.GetTargetSchema(), "getter should return the property value")
	})

	t.Run("GetTargetSchema_NilValue", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseLinksItem{}
		obj.TargetSchema = nil

		// Act & Assert
		assert.Nil(t, obj.GetTargetSchema(), "getter should return nil when property is nil")
	})

	t.Run("GetTargetSchema_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponseLinksItem
		// Should not panic - getters should handle nil receiver gracefully
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Getter panicked on nil receiver: %v", r)
			}
		}()
		_ = obj.GetTargetSchema() // Should return zero value
	})

}

func TestSettersMarkExplicitListRootResponseLinksItem(t *testing.T) {
	t.Run("SetHref_MarksExplicit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseLinksItem{}
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
		obj := &ListRootResponseLinksItem{}
		var fernTestValueMethod *ListRootResponseLinksItemMethod

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
		obj := &ListRootResponseLinksItem{}
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
		obj := &ListRootResponseLinksItem{}
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
		obj := &ListRootResponseLinksItem{}
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

func TestJSONMarshalingListRootResponse(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponse{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ListRootResponse
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ListRootResponse
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ListRootResponse
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingListRootResponseContact(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseContact{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ListRootResponseContact
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ListRootResponseContact
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ListRootResponseContact
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingListRootResponseIndustryStats(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseIndustryStats{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ListRootResponseIndustryStats
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ListRootResponseIndustryStats
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ListRootResponseIndustryStats
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestJSONMarshalingListRootResponseLinksItem(t *testing.T) {
	t.Run("MarshalUnmarshal", func(t *testing.T) {
		t.Parallel()
		// Arrange
		obj := &ListRootResponseLinksItem{}

		// Act - Marshal to JSON
		data, err := json.Marshal(obj)
		require.NoError(t, err, "marshaling should succeed")
		assert.NotNil(t, data, "marshaled data should not be nil")
		assert.NotEmpty(t, data, "marshaled data should not be empty")

		// Unmarshal back and verify round-trip
		var unmarshaled ListRootResponseLinksItem
		err = json.Unmarshal(data, &unmarshaled)
		assert.NoError(t, err, "round-trip unmarshal should succeed")
	})

	t.Run("UnmarshalInvalidJSON", func(t *testing.T) {
		t.Parallel()
		var obj ListRootResponseLinksItem
		err := json.Unmarshal([]byte(`{invalid json}`), &obj)
		assert.Error(t, err, "unmarshaling invalid JSON should return an error")
	})

	t.Run("UnmarshalEmptyObject", func(t *testing.T) {
		t.Parallel()
		var obj ListRootResponseLinksItem
		err := json.Unmarshal([]byte(`{}`), &obj)
		assert.NoError(t, err, "unmarshaling empty object should succeed")
	})
}

func TestStringListRootResponse(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ListRootResponse{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponse
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringListRootResponseContact(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ListRootResponseContact{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponseContact
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringListRootResponseIndustryStats(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ListRootResponseIndustryStats{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponseIndustryStats
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestStringListRootResponseLinksItem(t *testing.T) {
	t.Run("StringMethod", func(t *testing.T) {
		t.Parallel()
		obj := &ListRootResponseLinksItem{}
		result := obj.String()
		assert.NotEmpty(t, result, "String() should return a non-empty representation")
	})

	t.Run("StringMethod_NilReceiver", func(t *testing.T) {
		t.Parallel()
		var obj *ListRootResponseLinksItem
		result := obj.String()
		assert.Equal(t, "<nil>", result, "String() should return <nil> for nil receiver")
	})
}

func TestEnumListRootResponseFirstPaymentOne(t *testing.T) {
	t.Run("NewFromString_empty_string", func(t *testing.T) {
		t.Parallel()
		val, err := NewListRootResponseFirstPaymentOneFromString("")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListRootResponseFirstPaymentOne(""), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewListRootResponseFirstPaymentOneFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewListRootResponseFirstPaymentOneFromString("")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumListRootResponseLinksItemMethod(t *testing.T) {
	t.Run("NewFromString_GET", func(t *testing.T) {
		t.Parallel()
		val, err := NewListRootResponseLinksItemMethodFromString("GET")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListRootResponseLinksItemMethod("GET"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_POST", func(t *testing.T) {
		t.Parallel()
		val, err := NewListRootResponseLinksItemMethodFromString("POST")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListRootResponseLinksItemMethod("POST"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_PUT", func(t *testing.T) {
		t.Parallel()
		val, err := NewListRootResponseLinksItemMethodFromString("PUT")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListRootResponseLinksItemMethod("PUT"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_PATCH", func(t *testing.T) {
		t.Parallel()
		val, err := NewListRootResponseLinksItemMethodFromString("PATCH")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListRootResponseLinksItemMethod("PATCH"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_DELETE", func(t *testing.T) {
		t.Parallel()
		val, err := NewListRootResponseLinksItemMethodFromString("DELETE")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListRootResponseLinksItemMethod("DELETE"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_OPTIONS", func(t *testing.T) {
		t.Parallel()
		val, err := NewListRootResponseLinksItemMethodFromString("OPTIONS")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListRootResponseLinksItemMethod("OPTIONS"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_HEAD", func(t *testing.T) {
		t.Parallel()
		val, err := NewListRootResponseLinksItemMethodFromString("HEAD")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListRootResponseLinksItemMethod("HEAD"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewListRootResponseLinksItemMethodFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewListRootResponseLinksItemMethodFromString("GET")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestEnumListRootResponsePricingPlanType(t *testing.T) {
	t.Run("NewFromString_monthly", func(t *testing.T) {
		t.Parallel()
		val, err := NewListRootResponsePricingPlanTypeFromString("monthly")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListRootResponsePricingPlanType("monthly"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_pay_as_you_go", func(t *testing.T) {
		t.Parallel()
		val, err := NewListRootResponsePricingPlanTypeFromString("pay_as_you_go")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListRootResponsePricingPlanType("pay_as_you_go"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_forever_free", func(t *testing.T) {
		t.Parallel()
		val, err := NewListRootResponsePricingPlanTypeFromString("forever_free")
		assert.NoError(t, err, "valid enum value should not return error")
		assert.Equal(t, ListRootResponsePricingPlanType("forever_free"), val, "enum value should match expected wire value")
	})

	t.Run("NewFromString_Invalid", func(t *testing.T) {
		_, err := NewListRootResponsePricingPlanTypeFromString("invalid_value_that_does_not_exist")
		assert.Error(t, err)
	})

	t.Run("Ptr", func(t *testing.T) {
		val, err := NewListRootResponsePricingPlanTypeFromString("monthly")
		assert.NoError(t, err)
		ptr := val.Ptr()
		assert.NotNil(t, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestExtraPropertiesListRootResponse(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ListRootResponse{}
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
		var obj *ListRootResponse
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesListRootResponseContact(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ListRootResponseContact{}
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
		var obj *ListRootResponseContact
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesListRootResponseIndustryStats(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ListRootResponseIndustryStats{}
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
		var obj *ListRootResponseIndustryStats
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}

func TestExtraPropertiesListRootResponseLinksItem(t *testing.T) {
	t.Run("GetExtraProperties", func(t *testing.T) {
		t.Parallel()
		obj := &ListRootResponseLinksItem{}
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
		var obj *ListRootResponseLinksItem
		extraProps := obj.GetExtraProperties()
		assert.Nil(t, extraProps, "nil receiver should return nil without panicking")
	})
}
