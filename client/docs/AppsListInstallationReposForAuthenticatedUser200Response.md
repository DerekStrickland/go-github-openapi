# AppsListInstallationReposForAuthenticatedUser200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | 
**RepositorySelection** | Pointer to **string** |  | [optional] 
**Repositories** | [**[]Repository**](Repository.md) |  | 

## Methods

### NewAppsListInstallationReposForAuthenticatedUser200Response

`func NewAppsListInstallationReposForAuthenticatedUser200Response(totalCount int32, repositories []Repository, ) *AppsListInstallationReposForAuthenticatedUser200Response`

NewAppsListInstallationReposForAuthenticatedUser200Response instantiates a new AppsListInstallationReposForAuthenticatedUser200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppsListInstallationReposForAuthenticatedUser200ResponseWithDefaults

`func NewAppsListInstallationReposForAuthenticatedUser200ResponseWithDefaults() *AppsListInstallationReposForAuthenticatedUser200Response`

NewAppsListInstallationReposForAuthenticatedUser200ResponseWithDefaults instantiates a new AppsListInstallationReposForAuthenticatedUser200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *AppsListInstallationReposForAuthenticatedUser200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AppsListInstallationReposForAuthenticatedUser200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AppsListInstallationReposForAuthenticatedUser200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetRepositorySelection

`func (o *AppsListInstallationReposForAuthenticatedUser200Response) GetRepositorySelection() string`

GetRepositorySelection returns the RepositorySelection field if non-nil, zero value otherwise.

### GetRepositorySelectionOk

`func (o *AppsListInstallationReposForAuthenticatedUser200Response) GetRepositorySelectionOk() (*string, bool)`

GetRepositorySelectionOk returns a tuple with the RepositorySelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositorySelection

`func (o *AppsListInstallationReposForAuthenticatedUser200Response) SetRepositorySelection(v string)`

SetRepositorySelection sets RepositorySelection field to given value.

### HasRepositorySelection

`func (o *AppsListInstallationReposForAuthenticatedUser200Response) HasRepositorySelection() bool`

HasRepositorySelection returns a boolean if a field has been set.

### GetRepositories

`func (o *AppsListInstallationReposForAuthenticatedUser200Response) GetRepositories() []Repository`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *AppsListInstallationReposForAuthenticatedUser200Response) GetRepositoriesOk() (*[]Repository, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *AppsListInstallationReposForAuthenticatedUser200Response) SetRepositories(v []Repository)`

SetRepositories sets Repositories field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


