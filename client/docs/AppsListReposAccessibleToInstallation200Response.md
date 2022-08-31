# AppsListReposAccessibleToInstallation200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | 
**Repositories** | [**[]Repository**](Repository.md) |  | 
**RepositorySelection** | Pointer to **string** |  | [optional] 

## Methods

### NewAppsListReposAccessibleToInstallation200Response

`func NewAppsListReposAccessibleToInstallation200Response(totalCount int32, repositories []Repository, ) *AppsListReposAccessibleToInstallation200Response`

NewAppsListReposAccessibleToInstallation200Response instantiates a new AppsListReposAccessibleToInstallation200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppsListReposAccessibleToInstallation200ResponseWithDefaults

`func NewAppsListReposAccessibleToInstallation200ResponseWithDefaults() *AppsListReposAccessibleToInstallation200Response`

NewAppsListReposAccessibleToInstallation200ResponseWithDefaults instantiates a new AppsListReposAccessibleToInstallation200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *AppsListReposAccessibleToInstallation200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AppsListReposAccessibleToInstallation200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AppsListReposAccessibleToInstallation200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetRepositories

`func (o *AppsListReposAccessibleToInstallation200Response) GetRepositories() []Repository`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *AppsListReposAccessibleToInstallation200Response) GetRepositoriesOk() (*[]Repository, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *AppsListReposAccessibleToInstallation200Response) SetRepositories(v []Repository)`

SetRepositories sets Repositories field to given value.


### GetRepositorySelection

`func (o *AppsListReposAccessibleToInstallation200Response) GetRepositorySelection() string`

GetRepositorySelection returns the RepositorySelection field if non-nil, zero value otherwise.

### GetRepositorySelectionOk

`func (o *AppsListReposAccessibleToInstallation200Response) GetRepositorySelectionOk() (*string, bool)`

GetRepositorySelectionOk returns a tuple with the RepositorySelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositorySelection

`func (o *AppsListReposAccessibleToInstallation200Response) SetRepositorySelection(v string)`

SetRepositorySelection sets RepositorySelection field to given value.

### HasRepositorySelection

`func (o *AppsListReposAccessibleToInstallation200Response) HasRepositorySelection() bool`

HasRepositorySelection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


