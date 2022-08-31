# CombinedBillingUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaysLeftInBillingCycle** | **int32** | Numbers of days left in billing cycle. | 
**EstimatedPaidStorageForMonth** | **int32** | Estimated storage space (GB) used in billing cycle. | 
**EstimatedStorageForMonth** | **int32** | Estimated sum of free and paid storage space (GB) used in billing cycle. | 

## Methods

### NewCombinedBillingUsage

`func NewCombinedBillingUsage(daysLeftInBillingCycle int32, estimatedPaidStorageForMonth int32, estimatedStorageForMonth int32, ) *CombinedBillingUsage`

NewCombinedBillingUsage instantiates a new CombinedBillingUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCombinedBillingUsageWithDefaults

`func NewCombinedBillingUsageWithDefaults() *CombinedBillingUsage`

NewCombinedBillingUsageWithDefaults instantiates a new CombinedBillingUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaysLeftInBillingCycle

`func (o *CombinedBillingUsage) GetDaysLeftInBillingCycle() int32`

GetDaysLeftInBillingCycle returns the DaysLeftInBillingCycle field if non-nil, zero value otherwise.

### GetDaysLeftInBillingCycleOk

`func (o *CombinedBillingUsage) GetDaysLeftInBillingCycleOk() (*int32, bool)`

GetDaysLeftInBillingCycleOk returns a tuple with the DaysLeftInBillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysLeftInBillingCycle

`func (o *CombinedBillingUsage) SetDaysLeftInBillingCycle(v int32)`

SetDaysLeftInBillingCycle sets DaysLeftInBillingCycle field to given value.


### GetEstimatedPaidStorageForMonth

`func (o *CombinedBillingUsage) GetEstimatedPaidStorageForMonth() int32`

GetEstimatedPaidStorageForMonth returns the EstimatedPaidStorageForMonth field if non-nil, zero value otherwise.

### GetEstimatedPaidStorageForMonthOk

`func (o *CombinedBillingUsage) GetEstimatedPaidStorageForMonthOk() (*int32, bool)`

GetEstimatedPaidStorageForMonthOk returns a tuple with the EstimatedPaidStorageForMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedPaidStorageForMonth

`func (o *CombinedBillingUsage) SetEstimatedPaidStorageForMonth(v int32)`

SetEstimatedPaidStorageForMonth sets EstimatedPaidStorageForMonth field to given value.


### GetEstimatedStorageForMonth

`func (o *CombinedBillingUsage) GetEstimatedStorageForMonth() int32`

GetEstimatedStorageForMonth returns the EstimatedStorageForMonth field if non-nil, zero value otherwise.

### GetEstimatedStorageForMonthOk

`func (o *CombinedBillingUsage) GetEstimatedStorageForMonthOk() (*int32, bool)`

GetEstimatedStorageForMonthOk returns a tuple with the EstimatedStorageForMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedStorageForMonth

`func (o *CombinedBillingUsage) SetEstimatedStorageForMonth(v int32)`

SetEstimatedStorageForMonth sets EstimatedStorageForMonth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


