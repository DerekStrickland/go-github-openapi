# TeamsAddOrUpdateProjectPermissionsInOrgRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permission** | Pointer to **string** | The permission to grant to the team for this project. Default: the team&#39;s &#x60;permission&#x60; attribute will be used to determine what permission to grant the team on this project. Note that, if you choose not to pass any parameters, you&#39;ll need to set &#x60;Content-Length&#x60; to zero when calling this endpoint. For more information, see \&quot;[HTTP verbs](https://docs.github.com/rest/overview/resources-in-the-rest-api#http-verbs).\&quot; | [optional] 

## Methods

### NewTeamsAddOrUpdateProjectPermissionsInOrgRequest

`func NewTeamsAddOrUpdateProjectPermissionsInOrgRequest() *TeamsAddOrUpdateProjectPermissionsInOrgRequest`

NewTeamsAddOrUpdateProjectPermissionsInOrgRequest instantiates a new TeamsAddOrUpdateProjectPermissionsInOrgRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamsAddOrUpdateProjectPermissionsInOrgRequestWithDefaults

`func NewTeamsAddOrUpdateProjectPermissionsInOrgRequestWithDefaults() *TeamsAddOrUpdateProjectPermissionsInOrgRequest`

NewTeamsAddOrUpdateProjectPermissionsInOrgRequestWithDefaults instantiates a new TeamsAddOrUpdateProjectPermissionsInOrgRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermission

`func (o *TeamsAddOrUpdateProjectPermissionsInOrgRequest) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *TeamsAddOrUpdateProjectPermissionsInOrgRequest) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *TeamsAddOrUpdateProjectPermissionsInOrgRequest) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *TeamsAddOrUpdateProjectPermissionsInOrgRequest) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


