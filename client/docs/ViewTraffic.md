# ViewTraffic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Uniques** | **int32** |  | 
**Views** | [**[]Traffic**](Traffic.md) |  | 

## Methods

### NewViewTraffic

`func NewViewTraffic(count int32, uniques int32, views []Traffic, ) *ViewTraffic`

NewViewTraffic instantiates a new ViewTraffic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewTrafficWithDefaults

`func NewViewTrafficWithDefaults() *ViewTraffic`

NewViewTrafficWithDefaults instantiates a new ViewTraffic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ViewTraffic) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ViewTraffic) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ViewTraffic) SetCount(v int32)`

SetCount sets Count field to given value.


### GetUniques

`func (o *ViewTraffic) GetUniques() int32`

GetUniques returns the Uniques field if non-nil, zero value otherwise.

### GetUniquesOk

`func (o *ViewTraffic) GetUniquesOk() (*int32, bool)`

GetUniquesOk returns a tuple with the Uniques field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniques

`func (o *ViewTraffic) SetUniques(v int32)`

SetUniques sets Uniques field to given value.


### GetViews

`func (o *ViewTraffic) GetViews() []Traffic`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *ViewTraffic) GetViewsOk() (*[]Traffic, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *ViewTraffic) SetViews(v []Traffic)`

SetViews sets Views field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


