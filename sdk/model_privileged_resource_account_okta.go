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

// PrivilegedResourceAccountOkta struct for PrivilegedResourceAccountOkta
type PrivilegedResourceAccountOkta struct {
	PrivilegedResource
	// The user ID associated with the Okta privileged resource
	ResourceId           string                                         `json:"resourceId"`
	Credentials          *PrivilegedResourceAccountOktaAllOfCredentials `json:"credentials,omitempty"`
	Links                *UserLink                                      `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrivilegedResourceAccountOkta PrivilegedResourceAccountOkta

// NewPrivilegedResourceAccountOkta instantiates a new PrivilegedResourceAccountOkta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivilegedResourceAccountOkta(resourceId string) *PrivilegedResourceAccountOkta {
	this := PrivilegedResourceAccountOkta{}
	this.ResourceId = resourceId
	return &this
}

// NewPrivilegedResourceAccountOktaWithDefaults instantiates a new PrivilegedResourceAccountOkta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivilegedResourceAccountOktaWithDefaults() *PrivilegedResourceAccountOkta {
	this := PrivilegedResourceAccountOkta{}
	return &this
}

// GetResourceId returns the ResourceId field value
func (o *PrivilegedResourceAccountOkta) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *PrivilegedResourceAccountOkta) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *PrivilegedResourceAccountOkta) SetResourceId(v string) {
	o.ResourceId = v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *PrivilegedResourceAccountOkta) GetCredentials() PrivilegedResourceAccountOktaAllOfCredentials {
	if o == nil || o.Credentials == nil {
		var ret PrivilegedResourceAccountOktaAllOfCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegedResourceAccountOkta) GetCredentialsOk() (*PrivilegedResourceAccountOktaAllOfCredentials, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *PrivilegedResourceAccountOkta) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given PrivilegedResourceAccountOktaAllOfCredentials and assigns it to the Credentials field.
func (o *PrivilegedResourceAccountOkta) SetCredentials(v PrivilegedResourceAccountOktaAllOfCredentials) {
	o.Credentials = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PrivilegedResourceAccountOkta) GetLinks() UserLink {
	if o == nil || o.Links == nil {
		var ret UserLink
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegedResourceAccountOkta) GetLinksOk() (*UserLink, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PrivilegedResourceAccountOkta) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given UserLink and assigns it to the Links field.
func (o *PrivilegedResourceAccountOkta) SetLinks(v UserLink) {
	o.Links = &v
}

func (o PrivilegedResourceAccountOkta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPrivilegedResource, errPrivilegedResource := json.Marshal(o.PrivilegedResource)
	if errPrivilegedResource != nil {
		return []byte{}, errPrivilegedResource
	}
	errPrivilegedResource = json.Unmarshal([]byte(serializedPrivilegedResource), &toSerialize)
	if errPrivilegedResource != nil {
		return []byte{}, errPrivilegedResource
	}
	if true {
		toSerialize["resourceId"] = o.ResourceId
	}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PrivilegedResourceAccountOkta) UnmarshalJSON(bytes []byte) (err error) {
	type PrivilegedResourceAccountOktaWithoutEmbeddedStruct struct {
		// The user ID associated with the Okta privileged resource
		ResourceId  string                                         `json:"resourceId"`
		Credentials *PrivilegedResourceAccountOktaAllOfCredentials `json:"credentials,omitempty"`
		Links       *UserLink                                      `json:"_links,omitempty"`
	}

	varPrivilegedResourceAccountOktaWithoutEmbeddedStruct := PrivilegedResourceAccountOktaWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPrivilegedResourceAccountOktaWithoutEmbeddedStruct)
	if err == nil {
		varPrivilegedResourceAccountOkta := _PrivilegedResourceAccountOkta{}
		varPrivilegedResourceAccountOkta.ResourceId = varPrivilegedResourceAccountOktaWithoutEmbeddedStruct.ResourceId
		varPrivilegedResourceAccountOkta.Credentials = varPrivilegedResourceAccountOktaWithoutEmbeddedStruct.Credentials
		varPrivilegedResourceAccountOkta.Links = varPrivilegedResourceAccountOktaWithoutEmbeddedStruct.Links
		*o = PrivilegedResourceAccountOkta(varPrivilegedResourceAccountOkta)
	} else {
		return err
	}

	varPrivilegedResourceAccountOkta := _PrivilegedResourceAccountOkta{}

	err = json.Unmarshal(bytes, &varPrivilegedResourceAccountOkta)
	if err == nil {
		o.PrivilegedResource = varPrivilegedResourceAccountOkta.PrivilegedResource
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "resourceId")
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "_links")

		// remove fields from embedded structs
		reflectPrivilegedResource := reflect.ValueOf(o.PrivilegedResource)
		for i := 0; i < reflectPrivilegedResource.Type().NumField(); i++ {
			t := reflectPrivilegedResource.Type().Field(i)

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

type NullablePrivilegedResourceAccountOkta struct {
	value *PrivilegedResourceAccountOkta
	isSet bool
}

func (v NullablePrivilegedResourceAccountOkta) Get() *PrivilegedResourceAccountOkta {
	return v.value
}

func (v *NullablePrivilegedResourceAccountOkta) Set(val *PrivilegedResourceAccountOkta) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivilegedResourceAccountOkta) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivilegedResourceAccountOkta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivilegedResourceAccountOkta(val *PrivilegedResourceAccountOkta) *NullablePrivilegedResourceAccountOkta {
	return &NullablePrivilegedResourceAccountOkta{value: val, isSet: true}
}

func (v NullablePrivilegedResourceAccountOkta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivilegedResourceAccountOkta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
