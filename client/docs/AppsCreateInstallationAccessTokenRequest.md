# AppsCreateInstallationAccessTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repositories** | Pointer to **[]string** | List of repository names that the token should have access to | [optional] 
**RepositoryIds** | Pointer to **[]int32** | List of repository IDs that the token should have access to | [optional] 
**Permissions** | Pointer to [**AppPermissions**](AppPermissions.md) |  | [optional] 

## Methods

### NewAppsCreateInstallationAccessTokenRequest

`func NewAppsCreateInstallationAccessTokenRequest() *AppsCreateInstallationAccessTokenRequest`

NewAppsCreateInstallationAccessTokenRequest instantiates a new AppsCreateInstallationAccessTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppsCreateInstallationAccessTokenRequestWithDefaults

`func NewAppsCreateInstallationAccessTokenRequestWithDefaults() *AppsCreateInstallationAccessTokenRequest`

NewAppsCreateInstallationAccessTokenRequestWithDefaults instantiates a new AppsCreateInstallationAccessTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepositories

`func (o *AppsCreateInstallationAccessTokenRequest) GetRepositories() []string`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *AppsCreateInstallationAccessTokenRequest) GetRepositoriesOk() (*[]string, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *AppsCreateInstallationAccessTokenRequest) SetRepositories(v []string)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *AppsCreateInstallationAccessTokenRequest) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.

### GetRepositoryIds

`func (o *AppsCreateInstallationAccessTokenRequest) GetRepositoryIds() []int32`

GetRepositoryIds returns the RepositoryIds field if non-nil, zero value otherwise.

### GetRepositoryIdsOk

`func (o *AppsCreateInstallationAccessTokenRequest) GetRepositoryIdsOk() (*[]int32, bool)`

GetRepositoryIdsOk returns a tuple with the RepositoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryIds

`func (o *AppsCreateInstallationAccessTokenRequest) SetRepositoryIds(v []int32)`

SetRepositoryIds sets RepositoryIds field to given value.

### HasRepositoryIds

`func (o *AppsCreateInstallationAccessTokenRequest) HasRepositoryIds() bool`

HasRepositoryIds returns a boolean if a field has been set.

### GetPermissions

`func (o *AppsCreateInstallationAccessTokenRequest) GetPermissions() AppPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AppsCreateInstallationAccessTokenRequest) GetPermissionsOk() (*AppPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AppsCreateInstallationAccessTokenRequest) SetPermissions(v AppPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AppsCreateInstallationAccessTokenRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


