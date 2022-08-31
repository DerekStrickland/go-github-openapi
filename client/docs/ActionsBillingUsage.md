# ActionsBillingUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalMinutesUsed** | **int32** | The sum of the free and paid GitHub Actions minutes used. | 
**TotalPaidMinutesUsed** | **int32** | The total paid GitHub Actions minutes used. | 
**IncludedMinutes** | **int32** | The amount of free GitHub Actions minutes available. | 
**MinutesUsedBreakdown** | [**ActionsBillingUsageMinutesUsedBreakdown**](ActionsBillingUsageMinutesUsedBreakdown.md) |  | 

## Methods

### NewActionsBillingUsage

`func NewActionsBillingUsage(totalMinutesUsed int32, totalPaidMinutesUsed int32, includedMinutes int32, minutesUsedBreakdown ActionsBillingUsageMinutesUsedBreakdown, ) *ActionsBillingUsage`

NewActionsBillingUsage instantiates a new ActionsBillingUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsBillingUsageWithDefaults

`func NewActionsBillingUsageWithDefaults() *ActionsBillingUsage`

NewActionsBillingUsageWithDefaults instantiates a new ActionsBillingUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalMinutesUsed

`func (o *ActionsBillingUsage) GetTotalMinutesUsed() int32`

GetTotalMinutesUsed returns the TotalMinutesUsed field if non-nil, zero value otherwise.

### GetTotalMinutesUsedOk

`func (o *ActionsBillingUsage) GetTotalMinutesUsedOk() (*int32, bool)`

GetTotalMinutesUsedOk returns a tuple with the TotalMinutesUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMinutesUsed

`func (o *ActionsBillingUsage) SetTotalMinutesUsed(v int32)`

SetTotalMinutesUsed sets TotalMinutesUsed field to given value.


### GetTotalPaidMinutesUsed

`func (o *ActionsBillingUsage) GetTotalPaidMinutesUsed() int32`

GetTotalPaidMinutesUsed returns the TotalPaidMinutesUsed field if non-nil, zero value otherwise.

### GetTotalPaidMinutesUsedOk

`func (o *ActionsBillingUsage) GetTotalPaidMinutesUsedOk() (*int32, bool)`

GetTotalPaidMinutesUsedOk returns a tuple with the TotalPaidMinutesUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPaidMinutesUsed

`func (o *ActionsBillingUsage) SetTotalPaidMinutesUsed(v int32)`

SetTotalPaidMinutesUsed sets TotalPaidMinutesUsed field to given value.


### GetIncludedMinutes

`func (o *ActionsBillingUsage) GetIncludedMinutes() int32`

GetIncludedMinutes returns the IncludedMinutes field if non-nil, zero value otherwise.

### GetIncludedMinutesOk

`func (o *ActionsBillingUsage) GetIncludedMinutesOk() (*int32, bool)`

GetIncludedMinutesOk returns a tuple with the IncludedMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedMinutes

`func (o *ActionsBillingUsage) SetIncludedMinutes(v int32)`

SetIncludedMinutes sets IncludedMinutes field to given value.


### GetMinutesUsedBreakdown

`func (o *ActionsBillingUsage) GetMinutesUsedBreakdown() ActionsBillingUsageMinutesUsedBreakdown`

GetMinutesUsedBreakdown returns the MinutesUsedBreakdown field if non-nil, zero value otherwise.

### GetMinutesUsedBreakdownOk

`func (o *ActionsBillingUsage) GetMinutesUsedBreakdownOk() (*ActionsBillingUsageMinutesUsedBreakdown, bool)`

GetMinutesUsedBreakdownOk returns a tuple with the MinutesUsedBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutesUsedBreakdown

`func (o *ActionsBillingUsage) SetMinutesUsedBreakdown(v ActionsBillingUsageMinutesUsedBreakdown)`

SetMinutesUsedBreakdown sets MinutesUsedBreakdown field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


