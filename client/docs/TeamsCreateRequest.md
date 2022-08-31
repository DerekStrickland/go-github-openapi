# TeamsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the team. | 
**Description** | Pointer to **string** | The description of the team. | [optional] 
**Maintainers** | Pointer to **[]string** | List GitHub IDs for organization members who will become team maintainers. | [optional] 
**RepoNames** | Pointer to **[]string** | The full name (e.g., \&quot;organization-name/repository-name\&quot;) of repositories to add the team to. | [optional] 
**Privacy** | Pointer to **string** | The level of privacy this team should have. The options are:   **For a non-nested team:**   \\* &#x60;secret&#x60; - only visible to organization owners and members of this team.   \\* &#x60;closed&#x60; - visible to all members of this organization.   Default: &#x60;secret&#x60;   **For a parent or child team:**   \\* &#x60;closed&#x60; - visible to all members of this organization.   Default for child team: &#x60;closed&#x60; | [optional] 
**Permission** | Pointer to **string** | **Deprecated**. The permission that new repositories will be added to the team with when none is specified. | [optional] [default to "pull"]
**ParentTeamId** | Pointer to **int32** | The ID of a team to set as the parent team. | [optional] 

## Methods

### NewTeamsCreateRequest

`func NewTeamsCreateRequest(name string, ) *TeamsCreateRequest`

NewTeamsCreateRequest instantiates a new TeamsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamsCreateRequestWithDefaults

`func NewTeamsCreateRequestWithDefaults() *TeamsCreateRequest`

NewTeamsCreateRequestWithDefaults instantiates a new TeamsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TeamsCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamsCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamsCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TeamsCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TeamsCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TeamsCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TeamsCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaintainers

`func (o *TeamsCreateRequest) GetMaintainers() []string`

GetMaintainers returns the Maintainers field if non-nil, zero value otherwise.

### GetMaintainersOk

`func (o *TeamsCreateRequest) GetMaintainersOk() (*[]string, bool)`

GetMaintainersOk returns a tuple with the Maintainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainers

`func (o *TeamsCreateRequest) SetMaintainers(v []string)`

SetMaintainers sets Maintainers field to given value.

### HasMaintainers

`func (o *TeamsCreateRequest) HasMaintainers() bool`

HasMaintainers returns a boolean if a field has been set.

### GetRepoNames

`func (o *TeamsCreateRequest) GetRepoNames() []string`

GetRepoNames returns the RepoNames field if non-nil, zero value otherwise.

### GetRepoNamesOk

`func (o *TeamsCreateRequest) GetRepoNamesOk() (*[]string, bool)`

GetRepoNamesOk returns a tuple with the RepoNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoNames

`func (o *TeamsCreateRequest) SetRepoNames(v []string)`

SetRepoNames sets RepoNames field to given value.

### HasRepoNames

`func (o *TeamsCreateRequest) HasRepoNames() bool`

HasRepoNames returns a boolean if a field has been set.

### GetPrivacy

`func (o *TeamsCreateRequest) GetPrivacy() string`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *TeamsCreateRequest) GetPrivacyOk() (*string, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *TeamsCreateRequest) SetPrivacy(v string)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *TeamsCreateRequest) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetPermission

`func (o *TeamsCreateRequest) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *TeamsCreateRequest) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *TeamsCreateRequest) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *TeamsCreateRequest) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetParentTeamId

`func (o *TeamsCreateRequest) GetParentTeamId() int32`

GetParentTeamId returns the ParentTeamId field if non-nil, zero value otherwise.

### GetParentTeamIdOk

`func (o *TeamsCreateRequest) GetParentTeamIdOk() (*int32, bool)`

GetParentTeamIdOk returns a tuple with the ParentTeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTeamId

`func (o *TeamsCreateRequest) SetParentTeamId(v int32)`

SetParentTeamId sets ParentTeamId field to given value.

### HasParentTeamId

`func (o *TeamsCreateRequest) HasParentTeamId() bool`

HasParentTeamId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


