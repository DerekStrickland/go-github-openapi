# SearchIssuesAndPullRequests200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | 
**IncompleteResults** | **bool** |  | 
**Items** | [**[]IssueSearchResultItem**](IssueSearchResultItem.md) |  | 

## Methods

### NewSearchIssuesAndPullRequests200Response

`func NewSearchIssuesAndPullRequests200Response(totalCount int32, incompleteResults bool, items []IssueSearchResultItem, ) *SearchIssuesAndPullRequests200Response`

NewSearchIssuesAndPullRequests200Response instantiates a new SearchIssuesAndPullRequests200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchIssuesAndPullRequests200ResponseWithDefaults

`func NewSearchIssuesAndPullRequests200ResponseWithDefaults() *SearchIssuesAndPullRequests200Response`

NewSearchIssuesAndPullRequests200ResponseWithDefaults instantiates a new SearchIssuesAndPullRequests200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *SearchIssuesAndPullRequests200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SearchIssuesAndPullRequests200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SearchIssuesAndPullRequests200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetIncompleteResults

`func (o *SearchIssuesAndPullRequests200Response) GetIncompleteResults() bool`

GetIncompleteResults returns the IncompleteResults field if non-nil, zero value otherwise.

### GetIncompleteResultsOk

`func (o *SearchIssuesAndPullRequests200Response) GetIncompleteResultsOk() (*bool, bool)`

GetIncompleteResultsOk returns a tuple with the IncompleteResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompleteResults

`func (o *SearchIssuesAndPullRequests200Response) SetIncompleteResults(v bool)`

SetIncompleteResults sets IncompleteResults field to given value.


### GetItems

`func (o *SearchIssuesAndPullRequests200Response) GetItems() []IssueSearchResultItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SearchIssuesAndPullRequests200Response) GetItemsOk() (*[]IssueSearchResultItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SearchIssuesAndPullRequests200Response) SetItems(v []IssueSearchResultItem)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


