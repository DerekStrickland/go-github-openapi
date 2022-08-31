# ActionsListSelectedReposForOrgSecret200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | 
**Repositories** | [**[]MinimalRepository**](MinimalRepository.md) |  | 

## Methods

### NewActionsListSelectedReposForOrgSecret200Response

`func NewActionsListSelectedReposForOrgSecret200Response(totalCount int32, repositories []MinimalRepository, ) *ActionsListSelectedReposForOrgSecret200Response`

NewActionsListSelectedReposForOrgSecret200Response instantiates a new ActionsListSelectedReposForOrgSecret200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsListSelectedReposForOrgSecret200ResponseWithDefaults

`func NewActionsListSelectedReposForOrgSecret200ResponseWithDefaults() *ActionsListSelectedReposForOrgSecret200Response`

NewActionsListSelectedReposForOrgSecret200ResponseWithDefaults instantiates a new ActionsListSelectedReposForOrgSecret200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *ActionsListSelectedReposForOrgSecret200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ActionsListSelectedReposForOrgSecret200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ActionsListSelectedReposForOrgSecret200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetRepositories

`func (o *ActionsListSelectedReposForOrgSecret200Response) GetRepositories() []MinimalRepository`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *ActionsListSelectedReposForOrgSecret200Response) GetRepositoriesOk() (*[]MinimalRepository, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *ActionsListSelectedReposForOrgSecret200Response) SetRepositories(v []MinimalRepository)`

SetRepositories sets Repositories field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


