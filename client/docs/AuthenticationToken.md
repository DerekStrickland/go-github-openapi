# AuthenticationToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | The token used for authentication | 
**ExpiresAt** | **time.Time** | The time this token expires | 
**Permissions** | Pointer to **map[string]interface{}** |  | [optional] 
**Repositories** | Pointer to [**[]Repository**](Repository.md) | The repositories this token has access to | [optional] 
**SingleFile** | Pointer to **NullableString** |  | [optional] 
**RepositorySelection** | Pointer to **string** | Describe whether all repositories have been selected or there&#39;s a selection involved | [optional] 

## Methods

### NewAuthenticationToken

`func NewAuthenticationToken(token string, expiresAt time.Time, ) *AuthenticationToken`

NewAuthenticationToken instantiates a new AuthenticationToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationTokenWithDefaults

`func NewAuthenticationTokenWithDefaults() *AuthenticationToken`

NewAuthenticationTokenWithDefaults instantiates a new AuthenticationToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *AuthenticationToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthenticationToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthenticationToken) SetToken(v string)`

SetToken sets Token field to given value.


### GetExpiresAt

`func (o *AuthenticationToken) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AuthenticationToken) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AuthenticationToken) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetPermissions

`func (o *AuthenticationToken) GetPermissions() map[string]interface{}`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AuthenticationToken) GetPermissionsOk() (*map[string]interface{}, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AuthenticationToken) SetPermissions(v map[string]interface{})`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AuthenticationToken) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetRepositories

`func (o *AuthenticationToken) GetRepositories() []Repository`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *AuthenticationToken) GetRepositoriesOk() (*[]Repository, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *AuthenticationToken) SetRepositories(v []Repository)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *AuthenticationToken) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.

### GetSingleFile

`func (o *AuthenticationToken) GetSingleFile() string`

GetSingleFile returns the SingleFile field if non-nil, zero value otherwise.

### GetSingleFileOk

`func (o *AuthenticationToken) GetSingleFileOk() (*string, bool)`

GetSingleFileOk returns a tuple with the SingleFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleFile

`func (o *AuthenticationToken) SetSingleFile(v string)`

SetSingleFile sets SingleFile field to given value.

### HasSingleFile

`func (o *AuthenticationToken) HasSingleFile() bool`

HasSingleFile returns a boolean if a field has been set.

### SetSingleFileNil

`func (o *AuthenticationToken) SetSingleFileNil(b bool)`

 SetSingleFileNil sets the value for SingleFile to be an explicit nil

### UnsetSingleFile
`func (o *AuthenticationToken) UnsetSingleFile()`

UnsetSingleFile ensures that no value is present for SingleFile, not even an explicit nil
### GetRepositorySelection

`func (o *AuthenticationToken) GetRepositorySelection() string`

GetRepositorySelection returns the RepositorySelection field if non-nil, zero value otherwise.

### GetRepositorySelectionOk

`func (o *AuthenticationToken) GetRepositorySelectionOk() (*string, bool)`

GetRepositorySelectionOk returns a tuple with the RepositorySelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositorySelection

`func (o *AuthenticationToken) SetRepositorySelection(v string)`

SetRepositorySelection sets RepositorySelection field to given value.

### HasRepositorySelection

`func (o *AuthenticationToken) HasRepositorySelection() bool`

HasRepositorySelection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


