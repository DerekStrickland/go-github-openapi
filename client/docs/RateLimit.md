# RateLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** |  | 
**Remaining** | **int32** |  | 
**Reset** | **int32** |  | 
**Used** | **int32** |  | 

## Methods

### NewRateLimit

`func NewRateLimit(limit int32, remaining int32, reset int32, used int32, ) *RateLimit`

NewRateLimit instantiates a new RateLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitWithDefaults

`func NewRateLimitWithDefaults() *RateLimit`

NewRateLimitWithDefaults instantiates a new RateLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *RateLimit) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RateLimit) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RateLimit) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetRemaining

`func (o *RateLimit) GetRemaining() int32`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *RateLimit) GetRemainingOk() (*int32, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *RateLimit) SetRemaining(v int32)`

SetRemaining sets Remaining field to given value.


### GetReset

`func (o *RateLimit) GetReset() int32`

GetReset returns the Reset field if non-nil, zero value otherwise.

### GetResetOk

`func (o *RateLimit) GetResetOk() (*int32, bool)`

GetResetOk returns a tuple with the Reset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReset

`func (o *RateLimit) SetReset(v int32)`

SetReset sets Reset field to given value.


### GetUsed

`func (o *RateLimit) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *RateLimit) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *RateLimit) SetUsed(v int32)`

SetUsed sets Used field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


