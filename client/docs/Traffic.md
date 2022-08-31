# Traffic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** |  | 
**Uniques** | **int32** |  | 
**Count** | **int32** |  | 

## Methods

### NewTraffic

`func NewTraffic(timestamp time.Time, uniques int32, count int32, ) *Traffic`

NewTraffic instantiates a new Traffic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficWithDefaults

`func NewTrafficWithDefaults() *Traffic`

NewTrafficWithDefaults instantiates a new Traffic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *Traffic) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Traffic) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Traffic) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetUniques

`func (o *Traffic) GetUniques() int32`

GetUniques returns the Uniques field if non-nil, zero value otherwise.

### GetUniquesOk

`func (o *Traffic) GetUniquesOk() (*int32, bool)`

GetUniquesOk returns a tuple with the Uniques field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniques

`func (o *Traffic) SetUniques(v int32)`

SetUniques sets Uniques field to given value.


### GetCount

`func (o *Traffic) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Traffic) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Traffic) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


