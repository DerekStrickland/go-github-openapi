# InstallationToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** |  | 
**ExpiresAt** | **string** |  | 
**Permissions** | Pointer to [**AppPermissions**](AppPermissions.md) |  | [optional] 
**RepositorySelection** | Pointer to **string** |  | [optional] 
**Repositories** | Pointer to [**[]Repository**](Repository.md) |  | [optional] 
**SingleFile** | Pointer to **string** |  | [optional] 
**HasMultipleSingleFiles** | Pointer to **bool** |  | [optional] 
**SingleFilePaths** | Pointer to **[]string** |  | [optional] 

## Methods

### NewInstallationToken

`func NewInstallationToken(token string, expiresAt string, ) *InstallationToken`

NewInstallationToken instantiates a new InstallationToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallationTokenWithDefaults

`func NewInstallationTokenWithDefaults() *InstallationToken`

NewInstallationTokenWithDefaults instantiates a new InstallationToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *InstallationToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *InstallationToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *InstallationToken) SetToken(v string)`

SetToken sets Token field to given value.


### GetExpiresAt

`func (o *InstallationToken) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *InstallationToken) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *InstallationToken) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### GetPermissions

`func (o *InstallationToken) GetPermissions() AppPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *InstallationToken) GetPermissionsOk() (*AppPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *InstallationToken) SetPermissions(v AppPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *InstallationToken) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetRepositorySelection

`func (o *InstallationToken) GetRepositorySelection() string`

GetRepositorySelection returns the RepositorySelection field if non-nil, zero value otherwise.

### GetRepositorySelectionOk

`func (o *InstallationToken) GetRepositorySelectionOk() (*string, bool)`

GetRepositorySelectionOk returns a tuple with the RepositorySelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositorySelection

`func (o *InstallationToken) SetRepositorySelection(v string)`

SetRepositorySelection sets RepositorySelection field to given value.

### HasRepositorySelection

`func (o *InstallationToken) HasRepositorySelection() bool`

HasRepositorySelection returns a boolean if a field has been set.

### GetRepositories

`func (o *InstallationToken) GetRepositories() []Repository`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *InstallationToken) GetRepositoriesOk() (*[]Repository, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *InstallationToken) SetRepositories(v []Repository)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *InstallationToken) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.

### GetSingleFile

`func (o *InstallationToken) GetSingleFile() string`

GetSingleFile returns the SingleFile field if non-nil, zero value otherwise.

### GetSingleFileOk

`func (o *InstallationToken) GetSingleFileOk() (*string, bool)`

GetSingleFileOk returns a tuple with the SingleFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleFile

`func (o *InstallationToken) SetSingleFile(v string)`

SetSingleFile sets SingleFile field to given value.

### HasSingleFile

`func (o *InstallationToken) HasSingleFile() bool`

HasSingleFile returns a boolean if a field has been set.

### GetHasMultipleSingleFiles

`func (o *InstallationToken) GetHasMultipleSingleFiles() bool`

GetHasMultipleSingleFiles returns the HasMultipleSingleFiles field if non-nil, zero value otherwise.

### GetHasMultipleSingleFilesOk

`func (o *InstallationToken) GetHasMultipleSingleFilesOk() (*bool, bool)`

GetHasMultipleSingleFilesOk returns a tuple with the HasMultipleSingleFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMultipleSingleFiles

`func (o *InstallationToken) SetHasMultipleSingleFiles(v bool)`

SetHasMultipleSingleFiles sets HasMultipleSingleFiles field to given value.

### HasHasMultipleSingleFiles

`func (o *InstallationToken) HasHasMultipleSingleFiles() bool`

HasHasMultipleSingleFiles returns a boolean if a field has been set.

### GetSingleFilePaths

`func (o *InstallationToken) GetSingleFilePaths() []string`

GetSingleFilePaths returns the SingleFilePaths field if non-nil, zero value otherwise.

### GetSingleFilePathsOk

`func (o *InstallationToken) GetSingleFilePathsOk() (*[]string, bool)`

GetSingleFilePathsOk returns a tuple with the SingleFilePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleFilePaths

`func (o *InstallationToken) SetSingleFilePaths(v []string)`

SetSingleFilePaths sets SingleFilePaths field to given value.

### HasSingleFilePaths

`func (o *InstallationToken) HasSingleFilePaths() bool`

HasSingleFilePaths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


