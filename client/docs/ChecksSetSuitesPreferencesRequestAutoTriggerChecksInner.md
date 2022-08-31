# ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | **int32** | The &#x60;id&#x60; of the GitHub App. | 
**Setting** | **bool** | Set to &#x60;true&#x60; to enable automatic creation of CheckSuite events upon pushes to the repository, or &#x60;false&#x60; to disable them. | [default to true]

## Methods

### NewChecksSetSuitesPreferencesRequestAutoTriggerChecksInner

`func NewChecksSetSuitesPreferencesRequestAutoTriggerChecksInner(appId int32, setting bool, ) *ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner`

NewChecksSetSuitesPreferencesRequestAutoTriggerChecksInner instantiates a new ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChecksSetSuitesPreferencesRequestAutoTriggerChecksInnerWithDefaults

`func NewChecksSetSuitesPreferencesRequestAutoTriggerChecksInnerWithDefaults() *ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner`

NewChecksSetSuitesPreferencesRequestAutoTriggerChecksInnerWithDefaults instantiates a new ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner) SetAppId(v int32)`

SetAppId sets AppId field to given value.


### GetSetting

`func (o *ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner) GetSetting() bool`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner) GetSettingOk() (*bool, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner) SetSetting(v bool)`

SetSetting sets Setting field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


