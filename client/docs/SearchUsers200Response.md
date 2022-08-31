# SearchUsers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | 
**IncompleteResults** | **bool** |  | 
**Items** | [**[]UserSearchResultItem**](UserSearchResultItem.md) |  | 

## Methods

### NewSearchUsers200Response

`func NewSearchUsers200Response(totalCount int32, incompleteResults bool, items []UserSearchResultItem, ) *SearchUsers200Response`

NewSearchUsers200Response instantiates a new SearchUsers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchUsers200ResponseWithDefaults

`func NewSearchUsers200ResponseWithDefaults() *SearchUsers200Response`

NewSearchUsers200ResponseWithDefaults instantiates a new SearchUsers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *SearchUsers200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SearchUsers200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SearchUsers200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetIncompleteResults

`func (o *SearchUsers200Response) GetIncompleteResults() bool`

GetIncompleteResults returns the IncompleteResults field if non-nil, zero value otherwise.

### GetIncompleteResultsOk

`func (o *SearchUsers200Response) GetIncompleteResultsOk() (*bool, bool)`

GetIncompleteResultsOk returns a tuple with the IncompleteResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompleteResults

`func (o *SearchUsers200Response) SetIncompleteResults(v bool)`

SetIncompleteResults sets IncompleteResults field to given value.


### GetItems

`func (o *SearchUsers200Response) GetItems() []UserSearchResultItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SearchUsers200Response) GetItemsOk() (*[]UserSearchResultItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SearchUsers200Response) SetItems(v []UserSearchResultItem)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


