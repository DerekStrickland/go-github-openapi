# TeamsUpdateLegacyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the team. | 
**Description** | Pointer to **string** | The description of the team. | [optional] 
**Privacy** | Pointer to **string** | The level of privacy this team should have. Editing teams without specifying this parameter leaves &#x60;privacy&#x60; intact. The options are:   **For a non-nested team:**   \\* &#x60;secret&#x60; - only visible to organization owners and members of this team.   \\* &#x60;closed&#x60; - visible to all members of this organization.   **For a parent or child team:**   \\* &#x60;closed&#x60; - visible to all members of this organization. | [optional] 
**Permission** | Pointer to **string** | **Deprecated**. The permission that new repositories will be added to the team with when none is specified. | [optional] [default to "pull"]
**ParentTeamId** | Pointer to **NullableInt32** | The ID of a team to set as the parent team. | [optional] 

## Methods

### NewTeamsUpdateLegacyRequest

`func NewTeamsUpdateLegacyRequest(name string, ) *TeamsUpdateLegacyRequest`

NewTeamsUpdateLegacyRequest instantiates a new TeamsUpdateLegacyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamsUpdateLegacyRequestWithDefaults

`func NewTeamsUpdateLegacyRequestWithDefaults() *TeamsUpdateLegacyRequest`

NewTeamsUpdateLegacyRequestWithDefaults instantiates a new TeamsUpdateLegacyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TeamsUpdateLegacyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamsUpdateLegacyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamsUpdateLegacyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TeamsUpdateLegacyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TeamsUpdateLegacyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TeamsUpdateLegacyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TeamsUpdateLegacyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrivacy

`func (o *TeamsUpdateLegacyRequest) GetPrivacy() string`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *TeamsUpdateLegacyRequest) GetPrivacyOk() (*string, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *TeamsUpdateLegacyRequest) SetPrivacy(v string)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *TeamsUpdateLegacyRequest) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetPermission

`func (o *TeamsUpdateLegacyRequest) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *TeamsUpdateLegacyRequest) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *TeamsUpdateLegacyRequest) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *TeamsUpdateLegacyRequest) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetParentTeamId

`func (o *TeamsUpdateLegacyRequest) GetParentTeamId() int32`

GetParentTeamId returns the ParentTeamId field if non-nil, zero value otherwise.

### GetParentTeamIdOk

`func (o *TeamsUpdateLegacyRequest) GetParentTeamIdOk() (*int32, bool)`

GetParentTeamIdOk returns a tuple with the ParentTeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTeamId

`func (o *TeamsUpdateLegacyRequest) SetParentTeamId(v int32)`

SetParentTeamId sets ParentTeamId field to given value.

### HasParentTeamId

`func (o *TeamsUpdateLegacyRequest) HasParentTeamId() bool`

HasParentTeamId returns a boolean if a field has been set.

### SetParentTeamIdNil

`func (o *TeamsUpdateLegacyRequest) SetParentTeamIdNil(b bool)`

 SetParentTeamIdNil sets the value for ParentTeamId to be an explicit nil

### UnsetParentTeamId
`func (o *TeamsUpdateLegacyRequest) UnsetParentTeamId()`

UnsetParentTeamId ensures that no value is present for ParentTeamId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


