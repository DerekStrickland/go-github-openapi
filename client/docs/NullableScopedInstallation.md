# NullableScopedInstallation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | [**AppPermissions**](AppPermissions.md) |  | 
**RepositorySelection** | **string** | Describe whether all repositories have been selected or there&#39;s a selection involved | 
**SingleFileName** | **NullableString** |  | 
**HasMultipleSingleFiles** | Pointer to **bool** |  | [optional] 
**SingleFilePaths** | Pointer to **[]string** |  | [optional] 
**RepositoriesUrl** | **string** |  | 
**Account** | [**SimpleUser**](SimpleUser.md) |  | 

## Methods

### NewNullableScopedInstallation

`func NewNullableScopedInstallation(permissions AppPermissions, repositorySelection string, singleFileName NullableString, repositoriesUrl string, account SimpleUser, ) *NullableScopedInstallation`

NewNullableScopedInstallation instantiates a new NullableScopedInstallation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNullableScopedInstallationWithDefaults

`func NewNullableScopedInstallationWithDefaults() *NullableScopedInstallation`

NewNullableScopedInstallationWithDefaults instantiates a new NullableScopedInstallation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *NullableScopedInstallation) GetPermissions() AppPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *NullableScopedInstallation) GetPermissionsOk() (*AppPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *NullableScopedInstallation) SetPermissions(v AppPermissions)`

SetPermissions sets Permissions field to given value.


### GetRepositorySelection

`func (o *NullableScopedInstallation) GetRepositorySelection() string`

GetRepositorySelection returns the RepositorySelection field if non-nil, zero value otherwise.

### GetRepositorySelectionOk

`func (o *NullableScopedInstallation) GetRepositorySelectionOk() (*string, bool)`

GetRepositorySelectionOk returns a tuple with the RepositorySelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositorySelection

`func (o *NullableScopedInstallation) SetRepositorySelection(v string)`

SetRepositorySelection sets RepositorySelection field to given value.


### GetSingleFileName

`func (o *NullableScopedInstallation) GetSingleFileName() string`

GetSingleFileName returns the SingleFileName field if non-nil, zero value otherwise.

### GetSingleFileNameOk

`func (o *NullableScopedInstallation) GetSingleFileNameOk() (*string, bool)`

GetSingleFileNameOk returns a tuple with the SingleFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleFileName

`func (o *NullableScopedInstallation) SetSingleFileName(v string)`

SetSingleFileName sets SingleFileName field to given value.


### SetSingleFileNameNil

`func (o *NullableScopedInstallation) SetSingleFileNameNil(b bool)`

 SetSingleFileNameNil sets the value for SingleFileName to be an explicit nil

### UnsetSingleFileName
`func (o *NullableScopedInstallation) UnsetSingleFileName()`

UnsetSingleFileName ensures that no value is present for SingleFileName, not even an explicit nil
### GetHasMultipleSingleFiles

`func (o *NullableScopedInstallation) GetHasMultipleSingleFiles() bool`

GetHasMultipleSingleFiles returns the HasMultipleSingleFiles field if non-nil, zero value otherwise.

### GetHasMultipleSingleFilesOk

`func (o *NullableScopedInstallation) GetHasMultipleSingleFilesOk() (*bool, bool)`

GetHasMultipleSingleFilesOk returns a tuple with the HasMultipleSingleFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMultipleSingleFiles

`func (o *NullableScopedInstallation) SetHasMultipleSingleFiles(v bool)`

SetHasMultipleSingleFiles sets HasMultipleSingleFiles field to given value.

### HasHasMultipleSingleFiles

`func (o *NullableScopedInstallation) HasHasMultipleSingleFiles() bool`

HasHasMultipleSingleFiles returns a boolean if a field has been set.

### GetSingleFilePaths

`func (o *NullableScopedInstallation) GetSingleFilePaths() []string`

GetSingleFilePaths returns the SingleFilePaths field if non-nil, zero value otherwise.

### GetSingleFilePathsOk

`func (o *NullableScopedInstallation) GetSingleFilePathsOk() (*[]string, bool)`

GetSingleFilePathsOk returns a tuple with the SingleFilePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleFilePaths

`func (o *NullableScopedInstallation) SetSingleFilePaths(v []string)`

SetSingleFilePaths sets SingleFilePaths field to given value.

### HasSingleFilePaths

`func (o *NullableScopedInstallation) HasSingleFilePaths() bool`

HasSingleFilePaths returns a boolean if a field has been set.

### GetRepositoriesUrl

`func (o *NullableScopedInstallation) GetRepositoriesUrl() string`

GetRepositoriesUrl returns the RepositoriesUrl field if non-nil, zero value otherwise.

### GetRepositoriesUrlOk

`func (o *NullableScopedInstallation) GetRepositoriesUrlOk() (*string, bool)`

GetRepositoriesUrlOk returns a tuple with the RepositoriesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoriesUrl

`func (o *NullableScopedInstallation) SetRepositoriesUrl(v string)`

SetRepositoriesUrl sets RepositoriesUrl field to given value.


### GetAccount

`func (o *NullableScopedInstallation) GetAccount() SimpleUser`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NullableScopedInstallation) GetAccountOk() (*SimpleUser, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NullableScopedInstallation) SetAccount(v SimpleUser)`

SetAccount sets Account field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


