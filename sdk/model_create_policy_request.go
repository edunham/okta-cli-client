/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

API version: 5.1.0
Contact: devex-public@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
	"fmt"
)

// model_oneof.mustache
// CreatePolicyRequest - struct for CreatePolicyRequest
type CreatePolicyRequest struct {
	AccessPolicy                *AccessPolicy
	MultifactorEnrollmentPolicy *MultifactorEnrollmentPolicy
	OktaSignOnPolicy            *OktaSignOnPolicy
	PasswordPolicy              *PasswordPolicy
	ProfileEnrollmentPolicy     *ProfileEnrollmentPolicy
}

// AccessPolicyAsCreatePolicyRequest is a convenience function that returns AccessPolicy wrapped in CreatePolicyRequest
func AccessPolicyAsCreatePolicyRequest(v *AccessPolicy) CreatePolicyRequest {
	return CreatePolicyRequest{
		AccessPolicy: v,
	}
}

// MultifactorEnrollmentPolicyAsCreatePolicyRequest is a convenience function that returns MultifactorEnrollmentPolicy wrapped in CreatePolicyRequest
func MultifactorEnrollmentPolicyAsCreatePolicyRequest(v *MultifactorEnrollmentPolicy) CreatePolicyRequest {
	return CreatePolicyRequest{
		MultifactorEnrollmentPolicy: v,
	}
}

// OktaSignOnPolicyAsCreatePolicyRequest is a convenience function that returns OktaSignOnPolicy wrapped in CreatePolicyRequest
func OktaSignOnPolicyAsCreatePolicyRequest(v *OktaSignOnPolicy) CreatePolicyRequest {
	return CreatePolicyRequest{
		OktaSignOnPolicy: v,
	}
}

// PasswordPolicyAsCreatePolicyRequest is a convenience function that returns PasswordPolicy wrapped in CreatePolicyRequest
func PasswordPolicyAsCreatePolicyRequest(v *PasswordPolicy) CreatePolicyRequest {
	return CreatePolicyRequest{
		PasswordPolicy: v,
	}
}

// ProfileEnrollmentPolicyAsCreatePolicyRequest is a convenience function that returns ProfileEnrollmentPolicy wrapped in CreatePolicyRequest
func ProfileEnrollmentPolicyAsCreatePolicyRequest(v *ProfileEnrollmentPolicy) CreatePolicyRequest {
	return CreatePolicyRequest{
		ProfileEnrollmentPolicy: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *CreatePolicyRequest) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'ACCESS_POLICY'
	if jsonDict["type"] == "ACCESS_POLICY" {
		// try to unmarshal JSON data into AccessPolicy
		err = json.Unmarshal(data, &dst.AccessPolicy)
		if err == nil {
			return nil // data stored in dst.AccessPolicy, return on the first match
		} else {
			dst.AccessPolicy = nil
			return fmt.Errorf("Failed to unmarshal CreatePolicyRequest as AccessPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AccessPolicy'
	if jsonDict["type"] == "AccessPolicy" {
		// try to unmarshal JSON data into AccessPolicy
		err = json.Unmarshal(data, &dst.AccessPolicy)
		if err == nil {
			return nil // data stored in dst.AccessPolicy, return on the first match
		} else {
			dst.AccessPolicy = nil
			return fmt.Errorf("Failed to unmarshal CreatePolicyRequest as AccessPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MFA_ENROLL'
	if jsonDict["type"] == "MFA_ENROLL" {
		// try to unmarshal JSON data into MultifactorEnrollmentPolicy
		err = json.Unmarshal(data, &dst.MultifactorEnrollmentPolicy)
		if err == nil {
			return nil // data stored in dst.MultifactorEnrollmentPolicy, return on the first match
		} else {
			dst.MultifactorEnrollmentPolicy = nil
			return fmt.Errorf("Failed to unmarshal CreatePolicyRequest as MultifactorEnrollmentPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MultifactorEnrollmentPolicy'
	if jsonDict["type"] == "MultifactorEnrollmentPolicy" {
		// try to unmarshal JSON data into MultifactorEnrollmentPolicy
		err = json.Unmarshal(data, &dst.MultifactorEnrollmentPolicy)
		if err == nil {
			return nil // data stored in dst.MultifactorEnrollmentPolicy, return on the first match
		} else {
			dst.MultifactorEnrollmentPolicy = nil
			return fmt.Errorf("Failed to unmarshal CreatePolicyRequest as MultifactorEnrollmentPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'OKTA_SIGN_ON'
	if jsonDict["type"] == "OKTA_SIGN_ON" {
		// try to unmarshal JSON data into OktaSignOnPolicy
		err = json.Unmarshal(data, &dst.OktaSignOnPolicy)
		if err == nil {
			return nil // data stored in dst.OktaSignOnPolicy, return on the first match
		} else {
			dst.OktaSignOnPolicy = nil
			return fmt.Errorf("Failed to unmarshal CreatePolicyRequest as OktaSignOnPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'OktaSignOnPolicy'
	if jsonDict["type"] == "OktaSignOnPolicy" {
		// try to unmarshal JSON data into OktaSignOnPolicy
		err = json.Unmarshal(data, &dst.OktaSignOnPolicy)
		if err == nil {
			return nil // data stored in dst.OktaSignOnPolicy, return on the first match
		} else {
			dst.OktaSignOnPolicy = nil
			return fmt.Errorf("Failed to unmarshal CreatePolicyRequest as OktaSignOnPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PASSWORD'
	if jsonDict["type"] == "PASSWORD" {
		// try to unmarshal JSON data into PasswordPolicy
		err = json.Unmarshal(data, &dst.PasswordPolicy)
		if err == nil {
			return nil // data stored in dst.PasswordPolicy, return on the first match
		} else {
			dst.PasswordPolicy = nil
			return fmt.Errorf("Failed to unmarshal CreatePolicyRequest as PasswordPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PROFILE_ENROLLMENT'
	if jsonDict["type"] == "PROFILE_ENROLLMENT" {
		// try to unmarshal JSON data into ProfileEnrollmentPolicy
		err = json.Unmarshal(data, &dst.ProfileEnrollmentPolicy)
		if err == nil {
			return nil // data stored in dst.ProfileEnrollmentPolicy, return on the first match
		} else {
			dst.ProfileEnrollmentPolicy = nil
			return fmt.Errorf("Failed to unmarshal CreatePolicyRequest as ProfileEnrollmentPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PasswordPolicy'
	if jsonDict["type"] == "PasswordPolicy" {
		// try to unmarshal JSON data into PasswordPolicy
		err = json.Unmarshal(data, &dst.PasswordPolicy)
		if err == nil {
			return nil // data stored in dst.PasswordPolicy, return on the first match
		} else {
			dst.PasswordPolicy = nil
			return fmt.Errorf("Failed to unmarshal CreatePolicyRequest as PasswordPolicy: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ProfileEnrollmentPolicy'
	if jsonDict["type"] == "ProfileEnrollmentPolicy" {
		// try to unmarshal JSON data into ProfileEnrollmentPolicy
		err = json.Unmarshal(data, &dst.ProfileEnrollmentPolicy)
		if err == nil {
			return nil // data stored in dst.ProfileEnrollmentPolicy, return on the first match
		} else {
			dst.ProfileEnrollmentPolicy = nil
			return fmt.Errorf("Failed to unmarshal CreatePolicyRequest as ProfileEnrollmentPolicy: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreatePolicyRequest) MarshalJSON() ([]byte, error) {
	if src.AccessPolicy != nil {
		return json.Marshal(&src.AccessPolicy)
	}

	if src.MultifactorEnrollmentPolicy != nil {
		return json.Marshal(&src.MultifactorEnrollmentPolicy)
	}

	if src.OktaSignOnPolicy != nil {
		return json.Marshal(&src.OktaSignOnPolicy)
	}

	if src.PasswordPolicy != nil {
		return json.Marshal(&src.PasswordPolicy)
	}

	if src.ProfileEnrollmentPolicy != nil {
		return json.Marshal(&src.ProfileEnrollmentPolicy)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreatePolicyRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AccessPolicy != nil {
		return obj.AccessPolicy
	}

	if obj.MultifactorEnrollmentPolicy != nil {
		return obj.MultifactorEnrollmentPolicy
	}

	if obj.OktaSignOnPolicy != nil {
		return obj.OktaSignOnPolicy
	}

	if obj.PasswordPolicy != nil {
		return obj.PasswordPolicy
	}

	if obj.ProfileEnrollmentPolicy != nil {
		return obj.ProfileEnrollmentPolicy
	}

	// all schemas are nil
	return nil
}

type NullableCreatePolicyRequest struct {
	value *CreatePolicyRequest
	isSet bool
}

func (v NullableCreatePolicyRequest) Get() *CreatePolicyRequest {
	return v.value
}

func (v *NullableCreatePolicyRequest) Set(val *CreatePolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePolicyRequest(val *CreatePolicyRequest) *NullableCreatePolicyRequest {
	return &NullableCreatePolicyRequest{value: val, isSet: true}
}

func (v NullableCreatePolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
