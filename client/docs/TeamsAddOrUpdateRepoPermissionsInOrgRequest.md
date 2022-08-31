# TeamsAddOrUpdateRepoPermissionsInOrgRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permission** | Pointer to **string** | The permission to grant the team on this repository. We accept the following permissions to be set: &#x60;pull&#x60;, &#x60;triage&#x60;, &#x60;push&#x60;, &#x60;maintain&#x60;, &#x60;admin&#x60; and you can also specify a custom repository role name, if the owning organization has defined any. If no permission is specified, the team&#39;s &#x60;permission&#x60; attribute will be used to determine what permission to grant the team on this repository. | [optional] [default to "push"]

## Methods

### NewTeamsAddOrUpdateRepoPermissionsInOrgRequest

`func NewTeamsAddOrUpdateRepoPermissionsInOrgRequest() *TeamsAddOrUpdateRepoPermissionsInOrgRequest`

NewTeamsAddOrUpdateRepoPermissionsInOrgRequest instantiates a new TeamsAddOrUpdateRepoPermissionsInOrgRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamsAddOrUpdateRepoPermissionsInOrgRequestWithDefaults

`func NewTeamsAddOrUpdateRepoPermissionsInOrgRequestWithDefaults() *TeamsAddOrUpdateRepoPermissionsInOrgRequest`

NewTeamsAddOrUpdateRepoPermissionsInOrgRequestWithDefaults instantiates a new TeamsAddOrUpdateRepoPermissionsInOrgRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermission

`func (o *TeamsAddOrUpdateRepoPermissionsInOrgRequest) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *TeamsAddOrUpdateRepoPermissionsInOrgRequest) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *TeamsAddOrUpdateRepoPermissionsInOrgRequest) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *TeamsAddOrUpdateRepoPermissionsInOrgRequest) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


