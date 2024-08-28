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

// AccessPolicyRuleCustomCondition struct for AccessPolicyRuleCustomCondition
type AccessPolicyRuleCustomCondition struct {
	Condition *string `json:"condition,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessPolicyRuleCustomCondition AccessPolicyRuleCustomCondition

// NewAccessPolicyRuleCustomCondition instantiates a new AccessPolicyRuleCustomCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessPolicyRuleCustomCondition() *AccessPolicyRuleCustomCondition {
	this := AccessPolicyRuleCustomCondition{}
	return &this
}

// NewAccessPolicyRuleCustomConditionWithDefaults instantiates a new AccessPolicyRuleCustomCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessPolicyRuleCustomConditionWithDefaults() *AccessPolicyRuleCustomCondition {
	this := AccessPolicyRuleCustomCondition{}
	return &this
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *AccessPolicyRuleCustomCondition) GetCondition() string {
	if o == nil || o.Condition == nil {
		var ret string
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleCustomCondition) GetConditionOk() (*string, bool) {
	if o == nil || o.Condition == nil {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *AccessPolicyRuleCustomCondition) HasCondition() bool {
	if o != nil && o.Condition != nil {
		return true
	}

	return false
}

// SetCondition gets a reference to the given string and assigns it to the Condition field.
func (o *AccessPolicyRuleCustomCondition) SetCondition(v string) {
	o.Condition = &v
}

func (o AccessPolicyRuleCustomCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Condition != nil {
		toSerialize["condition"] = o.Condition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccessPolicyRuleCustomCondition) UnmarshalJSON(bytes []byte) (err error) {
	varAccessPolicyRuleCustomCondition := _AccessPolicyRuleCustomCondition{}

	err = json.Unmarshal(bytes, &varAccessPolicyRuleCustomCondition)
	if err == nil {
		*o = AccessPolicyRuleCustomCondition(varAccessPolicyRuleCustomCondition)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "condition")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAccessPolicyRuleCustomCondition struct {
	value *AccessPolicyRuleCustomCondition
	isSet bool
}

func (v NullableAccessPolicyRuleCustomCondition) Get() *AccessPolicyRuleCustomCondition {
	return v.value
}

func (v *NullableAccessPolicyRuleCustomCondition) Set(val *AccessPolicyRuleCustomCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessPolicyRuleCustomCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessPolicyRuleCustomCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessPolicyRuleCustomCondition(val *AccessPolicyRuleCustomCondition) *NullableAccessPolicyRuleCustomCondition {
	return &NullableAccessPolicyRuleCustomCondition{value: val, isSet: true}
}

func (v NullableAccessPolicyRuleCustomCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessPolicyRuleCustomCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

