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

// PolicyAccountLinkFilterGroups struct for PolicyAccountLinkFilterGroups
type PolicyAccountLinkFilterGroups struct {
	Include              []string `json:"include,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyAccountLinkFilterGroups PolicyAccountLinkFilterGroups

// NewPolicyAccountLinkFilterGroups instantiates a new PolicyAccountLinkFilterGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyAccountLinkFilterGroups() *PolicyAccountLinkFilterGroups {
	this := PolicyAccountLinkFilterGroups{}
	return &this
}

// NewPolicyAccountLinkFilterGroupsWithDefaults instantiates a new PolicyAccountLinkFilterGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyAccountLinkFilterGroupsWithDefaults() *PolicyAccountLinkFilterGroups {
	this := PolicyAccountLinkFilterGroups{}
	return &this
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *PolicyAccountLinkFilterGroups) GetInclude() []string {
	if o == nil || o.Include == nil {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAccountLinkFilterGroups) GetIncludeOk() ([]string, bool) {
	if o == nil || o.Include == nil {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *PolicyAccountLinkFilterGroups) HasInclude() bool {
	if o != nil && o.Include != nil {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *PolicyAccountLinkFilterGroups) SetInclude(v []string) {
	o.Include = v
}

func (o PolicyAccountLinkFilterGroups) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Include != nil {
		toSerialize["include"] = o.Include
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PolicyAccountLinkFilterGroups) UnmarshalJSON(bytes []byte) (err error) {
	varPolicyAccountLinkFilterGroups := _PolicyAccountLinkFilterGroups{}

	err = json.Unmarshal(bytes, &varPolicyAccountLinkFilterGroups)
	if err == nil {
		*o = PolicyAccountLinkFilterGroups(varPolicyAccountLinkFilterGroups)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "include")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePolicyAccountLinkFilterGroups struct {
	value *PolicyAccountLinkFilterGroups
	isSet bool
}

func (v NullablePolicyAccountLinkFilterGroups) Get() *PolicyAccountLinkFilterGroups {
	return v.value
}

func (v *NullablePolicyAccountLinkFilterGroups) Set(val *PolicyAccountLinkFilterGroups) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyAccountLinkFilterGroups) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyAccountLinkFilterGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyAccountLinkFilterGroups(val *PolicyAccountLinkFilterGroups) *NullablePolicyAccountLinkFilterGroups {
	return &NullablePolicyAccountLinkFilterGroups{value: val, isSet: true}
}

func (v NullablePolicyAccountLinkFilterGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyAccountLinkFilterGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
