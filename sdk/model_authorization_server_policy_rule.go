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
	"reflect"
	"strings"
)

// AuthorizationServerPolicyRule struct for AuthorizationServerPolicyRule
type AuthorizationServerPolicyRule struct {
	PolicyRule
	Actions *AuthorizationServerPolicyRuleActions `json:"actions,omitempty"`
	Conditions *AuthorizationServerPolicyRuleConditions `json:"conditions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthorizationServerPolicyRule AuthorizationServerPolicyRule

// NewAuthorizationServerPolicyRule instantiates a new AuthorizationServerPolicyRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationServerPolicyRule() *AuthorizationServerPolicyRule {
	this := AuthorizationServerPolicyRule{}
	var system bool = false
	this.System = &system
	return &this
}

// NewAuthorizationServerPolicyRuleWithDefaults instantiates a new AuthorizationServerPolicyRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationServerPolicyRuleWithDefaults() *AuthorizationServerPolicyRule {
	this := AuthorizationServerPolicyRule{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyRule) GetActions() AuthorizationServerPolicyRuleActions {
	if o == nil || o.Actions == nil {
		var ret AuthorizationServerPolicyRuleActions
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyRule) GetActionsOk() (*AuthorizationServerPolicyRuleActions, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyRule) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given AuthorizationServerPolicyRuleActions and assigns it to the Actions field.
func (o *AuthorizationServerPolicyRule) SetActions(v AuthorizationServerPolicyRuleActions) {
	o.Actions = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyRule) GetConditions() AuthorizationServerPolicyRuleConditions {
	if o == nil || o.Conditions == nil {
		var ret AuthorizationServerPolicyRuleConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyRule) GetConditionsOk() (*AuthorizationServerPolicyRuleConditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyRule) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given AuthorizationServerPolicyRuleConditions and assigns it to the Conditions field.
func (o *AuthorizationServerPolicyRule) SetConditions(v AuthorizationServerPolicyRuleConditions) {
	o.Conditions = &v
}

func (o AuthorizationServerPolicyRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyRule, errPolicyRule := json.Marshal(o.PolicyRule)
	if errPolicyRule != nil {
		return []byte{}, errPolicyRule
	}
	errPolicyRule = json.Unmarshal([]byte(serializedPolicyRule), &toSerialize)
	if errPolicyRule != nil {
		return []byte{}, errPolicyRule
	}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthorizationServerPolicyRule) UnmarshalJSON(bytes []byte) (err error) {
	type AuthorizationServerPolicyRuleWithoutEmbeddedStruct struct {
		Actions *AuthorizationServerPolicyRuleActions `json:"actions,omitempty"`
		Conditions *AuthorizationServerPolicyRuleConditions `json:"conditions,omitempty"`
	}

	varAuthorizationServerPolicyRuleWithoutEmbeddedStruct := AuthorizationServerPolicyRuleWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAuthorizationServerPolicyRuleWithoutEmbeddedStruct)
	if err == nil {
		varAuthorizationServerPolicyRule := _AuthorizationServerPolicyRule{}
		varAuthorizationServerPolicyRule.Actions = varAuthorizationServerPolicyRuleWithoutEmbeddedStruct.Actions
		varAuthorizationServerPolicyRule.Conditions = varAuthorizationServerPolicyRuleWithoutEmbeddedStruct.Conditions
		*o = AuthorizationServerPolicyRule(varAuthorizationServerPolicyRule)
	} else {
		return err
	}

	varAuthorizationServerPolicyRule := _AuthorizationServerPolicyRule{}

	err = json.Unmarshal(bytes, &varAuthorizationServerPolicyRule)
	if err == nil {
		o.PolicyRule = varAuthorizationServerPolicyRule.PolicyRule
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "actions")
		delete(additionalProperties, "conditions")

		// remove fields from embedded structs
		reflectPolicyRule := reflect.ValueOf(o.PolicyRule)
		for i := 0; i < reflectPolicyRule.Type().NumField(); i++ {
			t := reflectPolicyRule.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthorizationServerPolicyRule struct {
	value *AuthorizationServerPolicyRule
	isSet bool
}

func (v NullableAuthorizationServerPolicyRule) Get() *AuthorizationServerPolicyRule {
	return v.value
}

func (v *NullableAuthorizationServerPolicyRule) Set(val *AuthorizationServerPolicyRule) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationServerPolicyRule) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationServerPolicyRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationServerPolicyRule(val *AuthorizationServerPolicyRule) *NullableAuthorizationServerPolicyRule {
	return &NullableAuthorizationServerPolicyRule{value: val, isSet: true}
}

func (v NullableAuthorizationServerPolicyRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationServerPolicyRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

