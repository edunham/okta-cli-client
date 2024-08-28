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

// LogStreamSettingsAws Specifies the configuration for the `aws_eventbridge` Log Stream type. This configuration can't be modified after creation.
type LogStreamSettingsAws struct {
	// Your AWS account ID
	AccountId string `json:"accountId"`
	// An alphanumeric name (no spaces) to identify this event source in AWS EventBridge
	EventSourceName string `json:"eventSourceName"`
	// The destination AWS region where your event source is located
	Region string `json:"region"`
	AdditionalProperties map[string]interface{}
}

type _LogStreamSettingsAws LogStreamSettingsAws

// NewLogStreamSettingsAws instantiates a new LogStreamSettingsAws object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogStreamSettingsAws(accountId string, eventSourceName string, region string) *LogStreamSettingsAws {
	this := LogStreamSettingsAws{}
	this.AccountId = accountId
	this.EventSourceName = eventSourceName
	this.Region = region
	return &this
}

// NewLogStreamSettingsAwsWithDefaults instantiates a new LogStreamSettingsAws object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogStreamSettingsAwsWithDefaults() *LogStreamSettingsAws {
	this := LogStreamSettingsAws{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *LogStreamSettingsAws) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *LogStreamSettingsAws) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *LogStreamSettingsAws) SetAccountId(v string) {
	o.AccountId = v
}

// GetEventSourceName returns the EventSourceName field value
func (o *LogStreamSettingsAws) GetEventSourceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventSourceName
}

// GetEventSourceNameOk returns a tuple with the EventSourceName field value
// and a boolean to check if the value has been set.
func (o *LogStreamSettingsAws) GetEventSourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventSourceName, true
}

// SetEventSourceName sets field value
func (o *LogStreamSettingsAws) SetEventSourceName(v string) {
	o.EventSourceName = v
}

// GetRegion returns the Region field value
func (o *LogStreamSettingsAws) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *LogStreamSettingsAws) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *LogStreamSettingsAws) SetRegion(v string) {
	o.Region = v
}

func (o LogStreamSettingsAws) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["eventSourceName"] = o.EventSourceName
	}
	if true {
		toSerialize["region"] = o.Region
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LogStreamSettingsAws) UnmarshalJSON(bytes []byte) (err error) {
	varLogStreamSettingsAws := _LogStreamSettingsAws{}

	err = json.Unmarshal(bytes, &varLogStreamSettingsAws)
	if err == nil {
		*o = LogStreamSettingsAws(varLogStreamSettingsAws)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "accountId")
		delete(additionalProperties, "eventSourceName")
		delete(additionalProperties, "region")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLogStreamSettingsAws struct {
	value *LogStreamSettingsAws
	isSet bool
}

func (v NullableLogStreamSettingsAws) Get() *LogStreamSettingsAws {
	return v.value
}

func (v *NullableLogStreamSettingsAws) Set(val *LogStreamSettingsAws) {
	v.value = val
	v.isSet = true
}

func (v NullableLogStreamSettingsAws) IsSet() bool {
	return v.isSet
}

func (v *NullableLogStreamSettingsAws) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogStreamSettingsAws(val *LogStreamSettingsAws) *NullableLogStreamSettingsAws {
	return &NullableLogStreamSettingsAws{value: val, isSet: true}
}

func (v NullableLogStreamSettingsAws) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogStreamSettingsAws) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

