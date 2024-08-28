# CreatePolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the Policy was created | [optional] [readonly] 
**Description** | Pointer to **string** | Policy description | [optional] 
**Id** | Pointer to **string** | Policy ID | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the Policy was last updated | [optional] [readonly] 
**Name** | Pointer to **string** | Policy name | [optional] 
**Priority** | Pointer to **int32** | Specifies the order in which this Policy is evaluated in relation to the other policies | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**System** | Pointer to **bool** | Specifies whether Okta created the Policy | [optional] 
**Type** | Pointer to **string** | All Okta orgs contain only one IdP Discovery Policy with an immutable default Rule routing to your org&#39;s sign-in page. Creating or replacing a policy with &#x60;IDP_DISCOVERY&#x60; type isn&#39;t supported. | [optional] 
**Embedded** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**Links** | Pointer to [**PolicyLinks**](PolicyLinks.md) |  | [optional] 
**Conditions** | Pointer to [**PolicyRuleConditions**](PolicyRuleConditions.md) |  | [optional] 
**Settings** | Pointer to [**PasswordPolicySettings**](PasswordPolicySettings.md) |  | [optional] 

## Methods

### NewCreatePolicyRequest

`func NewCreatePolicyRequest() *CreatePolicyRequest`

NewCreatePolicyRequest instantiates a new CreatePolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePolicyRequestWithDefaults

`func NewCreatePolicyRequestWithDefaults() *CreatePolicyRequest`

NewCreatePolicyRequestWithDefaults instantiates a new CreatePolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *CreatePolicyRequest) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CreatePolicyRequest) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CreatePolicyRequest) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CreatePolicyRequest) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *CreatePolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatePolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatePolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreatePolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *CreatePolicyRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreatePolicyRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreatePolicyRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreatePolicyRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *CreatePolicyRequest) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CreatePolicyRequest) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CreatePolicyRequest) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *CreatePolicyRequest) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *CreatePolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePolicyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreatePolicyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *CreatePolicyRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreatePolicyRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreatePolicyRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CreatePolicyRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStatus

`func (o *CreatePolicyRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreatePolicyRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreatePolicyRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreatePolicyRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystem

`func (o *CreatePolicyRequest) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *CreatePolicyRequest) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *CreatePolicyRequest) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *CreatePolicyRequest) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetType

`func (o *CreatePolicyRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreatePolicyRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreatePolicyRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreatePolicyRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEmbedded

`func (o *CreatePolicyRequest) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *CreatePolicyRequest) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *CreatePolicyRequest) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *CreatePolicyRequest) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *CreatePolicyRequest) GetLinks() PolicyLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreatePolicyRequest) GetLinksOk() (*PolicyLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreatePolicyRequest) SetLinks(v PolicyLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreatePolicyRequest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetConditions

`func (o *CreatePolicyRequest) GetConditions() PolicyRuleConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *CreatePolicyRequest) GetConditionsOk() (*PolicyRuleConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *CreatePolicyRequest) SetConditions(v PolicyRuleConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *CreatePolicyRequest) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetSettings

`func (o *CreatePolicyRequest) GetSettings() PasswordPolicySettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CreatePolicyRequest) GetSettingsOk() (*PasswordPolicySettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CreatePolicyRequest) SetSettings(v PasswordPolicySettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *CreatePolicyRequest) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


