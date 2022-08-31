# MarketplacePurchaseMarketplacePurchase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingCycle** | Pointer to **string** |  | [optional] 
**NextBillingDate** | Pointer to **NullableString** |  | [optional] 
**IsInstalled** | Pointer to **bool** |  | [optional] 
**UnitCount** | Pointer to **NullableInt32** |  | [optional] 
**OnFreeTrial** | Pointer to **bool** |  | [optional] 
**FreeTrialEndsOn** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to [**MarketplaceListingPlan**](MarketplaceListingPlan.md) |  | [optional] 

## Methods

### NewMarketplacePurchaseMarketplacePurchase

`func NewMarketplacePurchaseMarketplacePurchase() *MarketplacePurchaseMarketplacePurchase`

NewMarketplacePurchaseMarketplacePurchase instantiates a new MarketplacePurchaseMarketplacePurchase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplacePurchaseMarketplacePurchaseWithDefaults

`func NewMarketplacePurchaseMarketplacePurchaseWithDefaults() *MarketplacePurchaseMarketplacePurchase`

NewMarketplacePurchaseMarketplacePurchaseWithDefaults instantiates a new MarketplacePurchaseMarketplacePurchase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingCycle

`func (o *MarketplacePurchaseMarketplacePurchase) GetBillingCycle() string`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *MarketplacePurchaseMarketplacePurchase) GetBillingCycleOk() (*string, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *MarketplacePurchaseMarketplacePurchase) SetBillingCycle(v string)`

SetBillingCycle sets BillingCycle field to given value.

### HasBillingCycle

`func (o *MarketplacePurchaseMarketplacePurchase) HasBillingCycle() bool`

HasBillingCycle returns a boolean if a field has been set.

### GetNextBillingDate

`func (o *MarketplacePurchaseMarketplacePurchase) GetNextBillingDate() string`

GetNextBillingDate returns the NextBillingDate field if non-nil, zero value otherwise.

### GetNextBillingDateOk

`func (o *MarketplacePurchaseMarketplacePurchase) GetNextBillingDateOk() (*string, bool)`

GetNextBillingDateOk returns a tuple with the NextBillingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextBillingDate

`func (o *MarketplacePurchaseMarketplacePurchase) SetNextBillingDate(v string)`

SetNextBillingDate sets NextBillingDate field to given value.

### HasNextBillingDate

`func (o *MarketplacePurchaseMarketplacePurchase) HasNextBillingDate() bool`

HasNextBillingDate returns a boolean if a field has been set.

### SetNextBillingDateNil

`func (o *MarketplacePurchaseMarketplacePurchase) SetNextBillingDateNil(b bool)`

 SetNextBillingDateNil sets the value for NextBillingDate to be an explicit nil

### UnsetNextBillingDate
`func (o *MarketplacePurchaseMarketplacePurchase) UnsetNextBillingDate()`

UnsetNextBillingDate ensures that no value is present for NextBillingDate, not even an explicit nil
### GetIsInstalled

`func (o *MarketplacePurchaseMarketplacePurchase) GetIsInstalled() bool`

GetIsInstalled returns the IsInstalled field if non-nil, zero value otherwise.

### GetIsInstalledOk

`func (o *MarketplacePurchaseMarketplacePurchase) GetIsInstalledOk() (*bool, bool)`

GetIsInstalledOk returns a tuple with the IsInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInstalled

`func (o *MarketplacePurchaseMarketplacePurchase) SetIsInstalled(v bool)`

SetIsInstalled sets IsInstalled field to given value.

### HasIsInstalled

`func (o *MarketplacePurchaseMarketplacePurchase) HasIsInstalled() bool`

HasIsInstalled returns a boolean if a field has been set.

### GetUnitCount

`func (o *MarketplacePurchaseMarketplacePurchase) GetUnitCount() int32`

GetUnitCount returns the UnitCount field if non-nil, zero value otherwise.

### GetUnitCountOk

`func (o *MarketplacePurchaseMarketplacePurchase) GetUnitCountOk() (*int32, bool)`

GetUnitCountOk returns a tuple with the UnitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitCount

`func (o *MarketplacePurchaseMarketplacePurchase) SetUnitCount(v int32)`

SetUnitCount sets UnitCount field to given value.

### HasUnitCount

`func (o *MarketplacePurchaseMarketplacePurchase) HasUnitCount() bool`

HasUnitCount returns a boolean if a field has been set.

### SetUnitCountNil

`func (o *MarketplacePurchaseMarketplacePurchase) SetUnitCountNil(b bool)`

 SetUnitCountNil sets the value for UnitCount to be an explicit nil

### UnsetUnitCount
`func (o *MarketplacePurchaseMarketplacePurchase) UnsetUnitCount()`

UnsetUnitCount ensures that no value is present for UnitCount, not even an explicit nil
### GetOnFreeTrial

`func (o *MarketplacePurchaseMarketplacePurchase) GetOnFreeTrial() bool`

GetOnFreeTrial returns the OnFreeTrial field if non-nil, zero value otherwise.

### GetOnFreeTrialOk

`func (o *MarketplacePurchaseMarketplacePurchase) GetOnFreeTrialOk() (*bool, bool)`

GetOnFreeTrialOk returns a tuple with the OnFreeTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFreeTrial

`func (o *MarketplacePurchaseMarketplacePurchase) SetOnFreeTrial(v bool)`

SetOnFreeTrial sets OnFreeTrial field to given value.

### HasOnFreeTrial

`func (o *MarketplacePurchaseMarketplacePurchase) HasOnFreeTrial() bool`

HasOnFreeTrial returns a boolean if a field has been set.

### GetFreeTrialEndsOn

`func (o *MarketplacePurchaseMarketplacePurchase) GetFreeTrialEndsOn() string`

GetFreeTrialEndsOn returns the FreeTrialEndsOn field if non-nil, zero value otherwise.

### GetFreeTrialEndsOnOk

`func (o *MarketplacePurchaseMarketplacePurchase) GetFreeTrialEndsOnOk() (*string, bool)`

GetFreeTrialEndsOnOk returns a tuple with the FreeTrialEndsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeTrialEndsOn

`func (o *MarketplacePurchaseMarketplacePurchase) SetFreeTrialEndsOn(v string)`

SetFreeTrialEndsOn sets FreeTrialEndsOn field to given value.

### HasFreeTrialEndsOn

`func (o *MarketplacePurchaseMarketplacePurchase) HasFreeTrialEndsOn() bool`

HasFreeTrialEndsOn returns a boolean if a field has been set.

### SetFreeTrialEndsOnNil

`func (o *MarketplacePurchaseMarketplacePurchase) SetFreeTrialEndsOnNil(b bool)`

 SetFreeTrialEndsOnNil sets the value for FreeTrialEndsOn to be an explicit nil

### UnsetFreeTrialEndsOn
`func (o *MarketplacePurchaseMarketplacePurchase) UnsetFreeTrialEndsOn()`

UnsetFreeTrialEndsOn ensures that no value is present for FreeTrialEndsOn, not even an explicit nil
### GetUpdatedAt

`func (o *MarketplacePurchaseMarketplacePurchase) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MarketplacePurchaseMarketplacePurchase) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MarketplacePurchaseMarketplacePurchase) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MarketplacePurchaseMarketplacePurchase) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetPlan

`func (o *MarketplacePurchaseMarketplacePurchase) GetPlan() MarketplaceListingPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *MarketplacePurchaseMarketplacePurchase) GetPlanOk() (*MarketplaceListingPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *MarketplacePurchaseMarketplacePurchase) SetPlan(v MarketplaceListingPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *MarketplacePurchaseMarketplacePurchase) HasPlan() bool`

HasPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


