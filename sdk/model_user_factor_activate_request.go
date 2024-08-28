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
)

// UserFactorActivateRequest struct for UserFactorActivateRequest
type UserFactorActivateRequest struct {
	Attestation          *string `json:"attestation,omitempty"`
	ClientData           *string `json:"clientData,omitempty"`
	PassCode             *string `json:"passCode,omitempty"`
	RegistrationData     *string `json:"registrationData,omitempty"`
	StateToken           *string `json:"stateToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorActivateRequest UserFactorActivateRequest

// NewUserFactorActivateRequest instantiates a new UserFactorActivateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorActivateRequest() *UserFactorActivateRequest {
	this := UserFactorActivateRequest{}
	return &this
}

// NewUserFactorActivateRequestWithDefaults instantiates a new UserFactorActivateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorActivateRequestWithDefaults() *UserFactorActivateRequest {
	this := UserFactorActivateRequest{}
	return &this
}

// GetAttestation returns the Attestation field value if set, zero value otherwise.
func (o *UserFactorActivateRequest) GetAttestation() string {
	if o == nil || o.Attestation == nil {
		var ret string
		return ret
	}
	return *o.Attestation
}

// GetAttestationOk returns a tuple with the Attestation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorActivateRequest) GetAttestationOk() (*string, bool) {
	if o == nil || o.Attestation == nil {
		return nil, false
	}
	return o.Attestation, true
}

// HasAttestation returns a boolean if a field has been set.
func (o *UserFactorActivateRequest) HasAttestation() bool {
	if o != nil && o.Attestation != nil {
		return true
	}

	return false
}

// SetAttestation gets a reference to the given string and assigns it to the Attestation field.
func (o *UserFactorActivateRequest) SetAttestation(v string) {
	o.Attestation = &v
}

// GetClientData returns the ClientData field value if set, zero value otherwise.
func (o *UserFactorActivateRequest) GetClientData() string {
	if o == nil || o.ClientData == nil {
		var ret string
		return ret
	}
	return *o.ClientData
}

// GetClientDataOk returns a tuple with the ClientData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorActivateRequest) GetClientDataOk() (*string, bool) {
	if o == nil || o.ClientData == nil {
		return nil, false
	}
	return o.ClientData, true
}

// HasClientData returns a boolean if a field has been set.
func (o *UserFactorActivateRequest) HasClientData() bool {
	if o != nil && o.ClientData != nil {
		return true
	}

	return false
}

// SetClientData gets a reference to the given string and assigns it to the ClientData field.
func (o *UserFactorActivateRequest) SetClientData(v string) {
	o.ClientData = &v
}

// GetPassCode returns the PassCode field value if set, zero value otherwise.
func (o *UserFactorActivateRequest) GetPassCode() string {
	if o == nil || o.PassCode == nil {
		var ret string
		return ret
	}
	return *o.PassCode
}

// GetPassCodeOk returns a tuple with the PassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorActivateRequest) GetPassCodeOk() (*string, bool) {
	if o == nil || o.PassCode == nil {
		return nil, false
	}
	return o.PassCode, true
}

// HasPassCode returns a boolean if a field has been set.
func (o *UserFactorActivateRequest) HasPassCode() bool {
	if o != nil && o.PassCode != nil {
		return true
	}

	return false
}

// SetPassCode gets a reference to the given string and assigns it to the PassCode field.
func (o *UserFactorActivateRequest) SetPassCode(v string) {
	o.PassCode = &v
}

// GetRegistrationData returns the RegistrationData field value if set, zero value otherwise.
func (o *UserFactorActivateRequest) GetRegistrationData() string {
	if o == nil || o.RegistrationData == nil {
		var ret string
		return ret
	}
	return *o.RegistrationData
}

// GetRegistrationDataOk returns a tuple with the RegistrationData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorActivateRequest) GetRegistrationDataOk() (*string, bool) {
	if o == nil || o.RegistrationData == nil {
		return nil, false
	}
	return o.RegistrationData, true
}

// HasRegistrationData returns a boolean if a field has been set.
func (o *UserFactorActivateRequest) HasRegistrationData() bool {
	if o != nil && o.RegistrationData != nil {
		return true
	}

	return false
}

// SetRegistrationData gets a reference to the given string and assigns it to the RegistrationData field.
func (o *UserFactorActivateRequest) SetRegistrationData(v string) {
	o.RegistrationData = &v
}

// GetStateToken returns the StateToken field value if set, zero value otherwise.
func (o *UserFactorActivateRequest) GetStateToken() string {
	if o == nil || o.StateToken == nil {
		var ret string
		return ret
	}
	return *o.StateToken
}

// GetStateTokenOk returns a tuple with the StateToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorActivateRequest) GetStateTokenOk() (*string, bool) {
	if o == nil || o.StateToken == nil {
		return nil, false
	}
	return o.StateToken, true
}

// HasStateToken returns a boolean if a field has been set.
func (o *UserFactorActivateRequest) HasStateToken() bool {
	if o != nil && o.StateToken != nil {
		return true
	}

	return false
}

// SetStateToken gets a reference to the given string and assigns it to the StateToken field.
func (o *UserFactorActivateRequest) SetStateToken(v string) {
	o.StateToken = &v
}

func (o UserFactorActivateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attestation != nil {
		toSerialize["attestation"] = o.Attestation
	}
	if o.ClientData != nil {
		toSerialize["clientData"] = o.ClientData
	}
	if o.PassCode != nil {
		toSerialize["passCode"] = o.PassCode
	}
	if o.RegistrationData != nil {
		toSerialize["registrationData"] = o.RegistrationData
	}
	if o.StateToken != nil {
		toSerialize["stateToken"] = o.StateToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserFactorActivateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varUserFactorActivateRequest := _UserFactorActivateRequest{}

	err = json.Unmarshal(bytes, &varUserFactorActivateRequest)
	if err == nil {
		*o = UserFactorActivateRequest(varUserFactorActivateRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "attestation")
		delete(additionalProperties, "clientData")
		delete(additionalProperties, "passCode")
		delete(additionalProperties, "registrationData")
		delete(additionalProperties, "stateToken")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserFactorActivateRequest struct {
	value *UserFactorActivateRequest
	isSet bool
}

func (v NullableUserFactorActivateRequest) Get() *UserFactorActivateRequest {
	return v.value
}

func (v *NullableUserFactorActivateRequest) Set(val *UserFactorActivateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorActivateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorActivateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorActivateRequest(val *UserFactorActivateRequest) *NullableUserFactorActivateRequest {
	return &NullableUserFactorActivateRequest{value: val, isSet: true}
}

func (v NullableUserFactorActivateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorActivateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
