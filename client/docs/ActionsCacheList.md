# ActionsCacheList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** | Total number of caches | 
**ActionsCaches** | [**[]ActionsCacheListActionsCachesInner**](ActionsCacheListActionsCachesInner.md) | Array of caches | 

## Methods

### NewActionsCacheList

`func NewActionsCacheList(totalCount int32, actionsCaches []ActionsCacheListActionsCachesInner, ) *ActionsCacheList`

NewActionsCacheList instantiates a new ActionsCacheList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsCacheListWithDefaults

`func NewActionsCacheListWithDefaults() *ActionsCacheList`

NewActionsCacheListWithDefaults instantiates a new ActionsCacheList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *ActionsCacheList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ActionsCacheList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ActionsCacheList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetActionsCaches

`func (o *ActionsCacheList) GetActionsCaches() []ActionsCacheListActionsCachesInner`

GetActionsCaches returns the ActionsCaches field if non-nil, zero value otherwise.

### GetActionsCachesOk

`func (o *ActionsCacheList) GetActionsCachesOk() (*[]ActionsCacheListActionsCachesInner, bool)`

GetActionsCachesOk returns a tuple with the ActionsCaches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionsCaches

`func (o *ActionsCacheList) SetActionsCaches(v []ActionsCacheListActionsCachesInner)`

SetActionsCaches sets ActionsCaches field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


