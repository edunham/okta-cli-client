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

// CapabilitiesInboundProvisioningObject Defines the configuration for the INBOUND_PROVISIONING feature
type CapabilitiesInboundProvisioningObject struct {
	ImportRules CapabilitiesImportRulesObject `json:"importRules"`
	ImportSettings CapabilitiesImportSettingsObject `json:"importSettings"`
	AdditionalProperties map[string]interface{}
}

type _CapabilitiesInboundProvisioningObject CapabilitiesInboundProvisioningObject

// NewCapabilitiesInboundProvisioningObject instantiates a new CapabilitiesInboundProvisioningObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilitiesInboundProvisioningObject(importRules CapabilitiesImportRulesObject, importSettings CapabilitiesImportSettingsObject) *CapabilitiesInboundProvisioningObject {
	this := CapabilitiesInboundProvisioningObject{}
	this.ImportRules = importRules
	this.ImportSettings = importSettings
	return &this
}

// NewCapabilitiesInboundProvisioningObjectWithDefaults instantiates a new CapabilitiesInboundProvisioningObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilitiesInboundProvisioningObjectWithDefaults() *CapabilitiesInboundProvisioningObject {
	this := CapabilitiesInboundProvisioningObject{}
	return &this
}

// GetImportRules returns the ImportRules field value
func (o *CapabilitiesInboundProvisioningObject) GetImportRules() CapabilitiesImportRulesObject {
	if o == nil {
		var ret CapabilitiesImportRulesObject
		return ret
	}

	return o.ImportRules
}

// GetImportRulesOk returns a tuple with the ImportRules field value
// and a boolean to check if the value has been set.
func (o *CapabilitiesInboundProvisioningObject) GetImportRulesOk() (*CapabilitiesImportRulesObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImportRules, true
}

// SetImportRules sets field value
func (o *CapabilitiesInboundProvisioningObject) SetImportRules(v CapabilitiesImportRulesObject) {
	o.ImportRules = v
}

// GetImportSettings returns the ImportSettings field value
func (o *CapabilitiesInboundProvisioningObject) GetImportSettings() CapabilitiesImportSettingsObject {
	if o == nil {
		var ret CapabilitiesImportSettingsObject
		return ret
	}

	return o.ImportSettings
}

// GetImportSettingsOk returns a tuple with the ImportSettings field value
// and a boolean to check if the value has been set.
func (o *CapabilitiesInboundProvisioningObject) GetImportSettingsOk() (*CapabilitiesImportSettingsObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImportSettings, true
}

// SetImportSettings sets field value
func (o *CapabilitiesInboundProvisioningObject) SetImportSettings(v CapabilitiesImportSettingsObject) {
	o.ImportSettings = v
}

func (o CapabilitiesInboundProvisioningObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["importRules"] = o.ImportRules
	}
	if true {
		toSerialize["importSettings"] = o.ImportSettings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilitiesInboundProvisioningObject) UnmarshalJSON(bytes []byte) (err error) {
	varCapabilitiesInboundProvisioningObject := _CapabilitiesInboundProvisioningObject{}

	err = json.Unmarshal(bytes, &varCapabilitiesInboundProvisioningObject)
	if err == nil {
		*o = CapabilitiesInboundProvisioningObject(varCapabilitiesInboundProvisioningObject)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "importRules")
		delete(additionalProperties, "importSettings")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCapabilitiesInboundProvisioningObject struct {
	value *CapabilitiesInboundProvisioningObject
	isSet bool
}

func (v NullableCapabilitiesInboundProvisioningObject) Get() *CapabilitiesInboundProvisioningObject {
	return v.value
}

func (v *NullableCapabilitiesInboundProvisioningObject) Set(val *CapabilitiesInboundProvisioningObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitiesInboundProvisioningObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitiesInboundProvisioningObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitiesInboundProvisioningObject(val *CapabilitiesInboundProvisioningObject) *NullableCapabilitiesInboundProvisioningObject {
	return &NullableCapabilitiesInboundProvisioningObject{value: val, isSet: true}
}

func (v NullableCapabilitiesInboundProvisioningObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitiesInboundProvisioningObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

