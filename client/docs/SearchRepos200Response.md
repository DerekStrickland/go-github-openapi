# SearchRepos200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | 
**IncompleteResults** | **bool** |  | 
**Items** | [**[]RepoSearchResultItem**](RepoSearchResultItem.md) |  | 

## Methods

### NewSearchRepos200Response

`func NewSearchRepos200Response(totalCount int32, incompleteResults bool, items []RepoSearchResultItem, ) *SearchRepos200Response`

NewSearchRepos200Response instantiates a new SearchRepos200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchRepos200ResponseWithDefaults

`func NewSearchRepos200ResponseWithDefaults() *SearchRepos200Response`

NewSearchRepos200ResponseWithDefaults instantiates a new SearchRepos200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *SearchRepos200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SearchRepos200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SearchRepos200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetIncompleteResults

`func (o *SearchRepos200Response) GetIncompleteResults() bool`

GetIncompleteResults returns the IncompleteResults field if non-nil, zero value otherwise.

### GetIncompleteResultsOk

`func (o *SearchRepos200Response) GetIncompleteResultsOk() (*bool, bool)`

GetIncompleteResultsOk returns a tuple with the IncompleteResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompleteResults

`func (o *SearchRepos200Response) SetIncompleteResults(v bool)`

SetIncompleteResults sets IncompleteResults field to given value.


### GetItems

`func (o *SearchRepos200Response) GetItems() []RepoSearchResultItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SearchRepos200Response) GetItemsOk() (*[]RepoSearchResultItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SearchRepos200Response) SetItems(v []RepoSearchResultItem)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


