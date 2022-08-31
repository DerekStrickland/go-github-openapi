# ActionsGetActionsCacheUsageByRepoForOrg200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | 
**RepositoryCacheUsages** | [**[]ActionsCacheUsageByRepository**](ActionsCacheUsageByRepository.md) |  | 

## Methods

### NewActionsGetActionsCacheUsageByRepoForOrg200Response

`func NewActionsGetActionsCacheUsageByRepoForOrg200Response(totalCount int32, repositoryCacheUsages []ActionsCacheUsageByRepository, ) *ActionsGetActionsCacheUsageByRepoForOrg200Response`

NewActionsGetActionsCacheUsageByRepoForOrg200Response instantiates a new ActionsGetActionsCacheUsageByRepoForOrg200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsGetActionsCacheUsageByRepoForOrg200ResponseWithDefaults

`func NewActionsGetActionsCacheUsageByRepoForOrg200ResponseWithDefaults() *ActionsGetActionsCacheUsageByRepoForOrg200Response`

NewActionsGetActionsCacheUsageByRepoForOrg200ResponseWithDefaults instantiates a new ActionsGetActionsCacheUsageByRepoForOrg200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *ActionsGetActionsCacheUsageByRepoForOrg200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ActionsGetActionsCacheUsageByRepoForOrg200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ActionsGetActionsCacheUsageByRepoForOrg200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetRepositoryCacheUsages

`func (o *ActionsGetActionsCacheUsageByRepoForOrg200Response) GetRepositoryCacheUsages() []ActionsCacheUsageByRepository`

GetRepositoryCacheUsages returns the RepositoryCacheUsages field if non-nil, zero value otherwise.

### GetRepositoryCacheUsagesOk

`func (o *ActionsGetActionsCacheUsageByRepoForOrg200Response) GetRepositoryCacheUsagesOk() (*[]ActionsCacheUsageByRepository, bool)`

GetRepositoryCacheUsagesOk returns a tuple with the RepositoryCacheUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryCacheUsages

`func (o *ActionsGetActionsCacheUsageByRepoForOrg200Response) SetRepositoryCacheUsages(v []ActionsCacheUsageByRepository)`

SetRepositoryCacheUsages sets RepositoryCacheUsages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


