# ChecksSetSuitesPreferencesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoTriggerChecks** | Pointer to [**[]ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner**](ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner.md) | Enables or disables automatic creation of CheckSuite events upon pushes to the repository. Enabled by default. See the [&#x60;auto_trigger_checks&#x60; object](https://docs.github.com/rest/reference/checks#auto_trigger_checks-object) description for details. | [optional] 

## Methods

### NewChecksSetSuitesPreferencesRequest

`func NewChecksSetSuitesPreferencesRequest() *ChecksSetSuitesPreferencesRequest`

NewChecksSetSuitesPreferencesRequest instantiates a new ChecksSetSuitesPreferencesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChecksSetSuitesPreferencesRequestWithDefaults

`func NewChecksSetSuitesPreferencesRequestWithDefaults() *ChecksSetSuitesPreferencesRequest`

NewChecksSetSuitesPreferencesRequestWithDefaults instantiates a new ChecksSetSuitesPreferencesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoTriggerChecks

`func (o *ChecksSetSuitesPreferencesRequest) GetAutoTriggerChecks() []ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner`

GetAutoTriggerChecks returns the AutoTriggerChecks field if non-nil, zero value otherwise.

### GetAutoTriggerChecksOk

`func (o *ChecksSetSuitesPreferencesRequest) GetAutoTriggerChecksOk() (*[]ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner, bool)`

GetAutoTriggerChecksOk returns a tuple with the AutoTriggerChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTriggerChecks

`func (o *ChecksSetSuitesPreferencesRequest) SetAutoTriggerChecks(v []ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner)`

SetAutoTriggerChecks sets AutoTriggerChecks field to given value.

### HasAutoTriggerChecks

`func (o *ChecksSetSuitesPreferencesRequest) HasAutoTriggerChecks() bool`

HasAutoTriggerChecks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


