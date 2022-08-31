# AppsScopeTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | The OAuth access token used to authenticate to the GitHub API. | 
**Target** | Pointer to **string** | The name of the user or organization to scope the user-to-server access token to. **Required** unless &#x60;target_id&#x60; is specified. | [optional] 
**TargetId** | Pointer to **int32** | The ID of the user or organization to scope the user-to-server access token to. **Required** unless &#x60;target&#x60; is specified. | [optional] 
**Repositories** | Pointer to **[]string** | The list of repository names to scope the user-to-server access token to. &#x60;repositories&#x60; may not be specified if &#x60;repository_ids&#x60; is specified. | [optional] 
**RepositoryIds** | Pointer to **[]int32** | The list of repository IDs to scope the user-to-server access token to. &#x60;repository_ids&#x60; may not be specified if &#x60;repositories&#x60; is specified. | [optional] 
**Permissions** | Pointer to [**AppPermissions**](AppPermissions.md) |  | [optional] 

## Methods

### NewAppsScopeTokenRequest

`func NewAppsScopeTokenRequest(accessToken string, ) *AppsScopeTokenRequest`

NewAppsScopeTokenRequest instantiates a new AppsScopeTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppsScopeTokenRequestWithDefaults

`func NewAppsScopeTokenRequestWithDefaults() *AppsScopeTokenRequest`

NewAppsScopeTokenRequestWithDefaults instantiates a new AppsScopeTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *AppsScopeTokenRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AppsScopeTokenRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AppsScopeTokenRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetTarget

`func (o *AppsScopeTokenRequest) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AppsScopeTokenRequest) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AppsScopeTokenRequest) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *AppsScopeTokenRequest) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTargetId

`func (o *AppsScopeTokenRequest) GetTargetId() int32`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *AppsScopeTokenRequest) GetTargetIdOk() (*int32, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *AppsScopeTokenRequest) SetTargetId(v int32)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *AppsScopeTokenRequest) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetRepositories

`func (o *AppsScopeTokenRequest) GetRepositories() []string`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *AppsScopeTokenRequest) GetRepositoriesOk() (*[]string, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *AppsScopeTokenRequest) SetRepositories(v []string)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *AppsScopeTokenRequest) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.

### GetRepositoryIds

`func (o *AppsScopeTokenRequest) GetRepositoryIds() []int32`

GetRepositoryIds returns the RepositoryIds field if non-nil, zero value otherwise.

### GetRepositoryIdsOk

`func (o *AppsScopeTokenRequest) GetRepositoryIdsOk() (*[]int32, bool)`

GetRepositoryIdsOk returns a tuple with the RepositoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryIds

`func (o *AppsScopeTokenRequest) SetRepositoryIds(v []int32)`

SetRepositoryIds sets RepositoryIds field to given value.

### HasRepositoryIds

`func (o *AppsScopeTokenRequest) HasRepositoryIds() bool`

HasRepositoryIds returns a boolean if a field has been set.

### GetPermissions

`func (o *AppsScopeTokenRequest) GetPermissions() AppPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AppsScopeTokenRequest) GetPermissionsOk() (*AppPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AppsScopeTokenRequest) SetPermissions(v AppPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AppsScopeTokenRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


