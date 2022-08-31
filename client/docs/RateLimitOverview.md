# RateLimitOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | [**RateLimitOverviewResources**](RateLimitOverviewResources.md) |  | 
**Rate** | [**RateLimit**](RateLimit.md) |  | 

## Methods

### NewRateLimitOverview

`func NewRateLimitOverview(resources RateLimitOverviewResources, rate RateLimit, ) *RateLimitOverview`

NewRateLimitOverview instantiates a new RateLimitOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitOverviewWithDefaults

`func NewRateLimitOverviewWithDefaults() *RateLimitOverview`

NewRateLimitOverviewWithDefaults instantiates a new RateLimitOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResources

`func (o *RateLimitOverview) GetResources() RateLimitOverviewResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *RateLimitOverview) GetResourcesOk() (*RateLimitOverviewResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *RateLimitOverview) SetResources(v RateLimitOverviewResources)`

SetResources sets Resources field to given value.


### GetRate

`func (o *RateLimitOverview) GetRate() RateLimit`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *RateLimitOverview) GetRateOk() (*RateLimit, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *RateLimitOverview) SetRate(v RateLimit)`

SetRate sets Rate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


