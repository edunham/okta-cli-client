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

// AccessPolicyConstraints struct for AccessPolicyConstraints
type AccessPolicyConstraints struct {
	Knowledge            *KnowledgeConstraint  `json:"knowledge,omitempty"`
	Possession           *PossessionConstraint `json:"possession,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessPolicyConstraints AccessPolicyConstraints

// NewAccessPolicyConstraints instantiates a new AccessPolicyConstraints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessPolicyConstraints() *AccessPolicyConstraints {
	this := AccessPolicyConstraints{}
	return &this
}

// NewAccessPolicyConstraintsWithDefaults instantiates a new AccessPolicyConstraints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessPolicyConstraintsWithDefaults() *AccessPolicyConstraints {
	this := AccessPolicyConstraints{}
	return &this
}

// GetKnowledge returns the Knowledge field value if set, zero value otherwise.
func (o *AccessPolicyConstraints) GetKnowledge() KnowledgeConstraint {
	if o == nil || o.Knowledge == nil {
		var ret KnowledgeConstraint
		return ret
	}
	return *o.Knowledge
}

// GetKnowledgeOk returns a tuple with the Knowledge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyConstraints) GetKnowledgeOk() (*KnowledgeConstraint, bool) {
	if o == nil || o.Knowledge == nil {
		return nil, false
	}
	return o.Knowledge, true
}

// HasKnowledge returns a boolean if a field has been set.
func (o *AccessPolicyConstraints) HasKnowledge() bool {
	if o != nil && o.Knowledge != nil {
		return true
	}

	return false
}

// SetKnowledge gets a reference to the given KnowledgeConstraint and assigns it to the Knowledge field.
func (o *AccessPolicyConstraints) SetKnowledge(v KnowledgeConstraint) {
	o.Knowledge = &v
}

// GetPossession returns the Possession field value if set, zero value otherwise.
func (o *AccessPolicyConstraints) GetPossession() PossessionConstraint {
	if o == nil || o.Possession == nil {
		var ret PossessionConstraint
		return ret
	}
	return *o.Possession
}

// GetPossessionOk returns a tuple with the Possession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyConstraints) GetPossessionOk() (*PossessionConstraint, bool) {
	if o == nil || o.Possession == nil {
		return nil, false
	}
	return o.Possession, true
}

// HasPossession returns a boolean if a field has been set.
func (o *AccessPolicyConstraints) HasPossession() bool {
	if o != nil && o.Possession != nil {
		return true
	}

	return false
}

// SetPossession gets a reference to the given PossessionConstraint and assigns it to the Possession field.
func (o *AccessPolicyConstraints) SetPossession(v PossessionConstraint) {
	o.Possession = &v
}

func (o AccessPolicyConstraints) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Knowledge != nil {
		toSerialize["knowledge"] = o.Knowledge
	}
	if o.Possession != nil {
		toSerialize["possession"] = o.Possession
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccessPolicyConstraints) UnmarshalJSON(bytes []byte) (err error) {
	varAccessPolicyConstraints := _AccessPolicyConstraints{}

	err = json.Unmarshal(bytes, &varAccessPolicyConstraints)
	if err == nil {
		*o = AccessPolicyConstraints(varAccessPolicyConstraints)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "knowledge")
		delete(additionalProperties, "possession")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAccessPolicyConstraints struct {
	value *AccessPolicyConstraints
	isSet bool
}

func (v NullableAccessPolicyConstraints) Get() *AccessPolicyConstraints {
	return v.value
}

func (v *NullableAccessPolicyConstraints) Set(val *AccessPolicyConstraints) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessPolicyConstraints) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessPolicyConstraints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessPolicyConstraints(val *AccessPolicyConstraints) *NullableAccessPolicyConstraints {
	return &NullableAccessPolicyConstraints{value: val, isSet: true}
}

func (v NullableAccessPolicyConstraints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessPolicyConstraints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
