# CodespacesCreateWithPrForAuthenticatedUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to **string** | Location for this codespace. Assigned by IP if not provided | [optional] 
**ClientIp** | Pointer to **string** | IP for location auto-detection when proxying a request | [optional] 
**Machine** | Pointer to **string** | Machine type to use for this codespace | [optional] 
**DevcontainerPath** | Pointer to **string** | Path to devcontainer.json config to use for this codespace | [optional] 
**MultiRepoPermissionsOptOut** | Pointer to **bool** | Whether to authorize requested permissions from devcontainer.json | [optional] 
**WorkingDirectory** | Pointer to **string** | Working directory for this codespace | [optional] 
**IdleTimeoutMinutes** | Pointer to **int32** | Time in minutes before codespace stops from inactivity | [optional] 
**DisplayName** | Pointer to **string** | Display name for this codespace | [optional] 
**RetentionPeriodMinutes** | Pointer to **int32** | Duration in minutes after codespace has gone idle in which it will be deleted. Must be integer minutes between 0 and 43200 (30 days). | [optional] 

## Methods

### NewCodespacesCreateWithPrForAuthenticatedUserRequest

`func NewCodespacesCreateWithPrForAuthenticatedUserRequest() *CodespacesCreateWithPrForAuthenticatedUserRequest`

NewCodespacesCreateWithPrForAuthenticatedUserRequest instantiates a new CodespacesCreateWithPrForAuthenticatedUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodespacesCreateWithPrForAuthenticatedUserRequestWithDefaults

`func NewCodespacesCreateWithPrForAuthenticatedUserRequestWithDefaults() *CodespacesCreateWithPrForAuthenticatedUserRequest`

NewCodespacesCreateWithPrForAuthenticatedUserRequestWithDefaults instantiates a new CodespacesCreateWithPrForAuthenticatedUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetClientIp

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetMachine

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetDevcontainerPath

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) GetDevcontainerPath() string`

GetDevcontainerPath returns the DevcontainerPath field if non-nil, zero value otherwise.

### GetDevcontainerPathOk

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) GetDevcontainerPathOk() (*string, bool)`

GetDevcontainerPathOk returns a tuple with the DevcontainerPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevcontainerPath

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) SetDevcontainerPath(v string)`

SetDevcontainerPath sets DevcontainerPath field to given value.

### HasDevcontainerPath

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) HasDevcontainerPath() bool`

HasDevcontainerPath returns a boolean if a field has been set.

### GetMultiRepoPermissionsOptOut

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) GetMultiRepoPermissionsOptOut() bool`

GetMultiRepoPermissionsOptOut returns the MultiRepoPermissionsOptOut field if non-nil, zero value otherwise.

### GetMultiRepoPermissionsOptOutOk

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) GetMultiRepoPermissionsOptOutOk() (*bool, bool)`

GetMultiRepoPermissionsOptOutOk returns a tuple with the MultiRepoPermissionsOptOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiRepoPermissionsOptOut

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) SetMultiRepoPermissionsOptOut(v bool)`

SetMultiRepoPermissionsOptOut sets MultiRepoPermissionsOptOut field to given value.

### HasMultiRepoPermissionsOptOut

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) HasMultiRepoPermissionsOptOut() bool`

HasMultiRepoPermissionsOptOut returns a boolean if a field has been set.

### GetWorkingDirectory

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) GetWorkingDirectory() string`

GetWorkingDirectory returns the WorkingDirectory field if non-nil, zero value otherwise.

### GetWorkingDirectoryOk

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) GetWorkingDirectoryOk() (*string, bool)`

GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirectory

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) SetWorkingDirectory(v string)`

SetWorkingDirectory sets WorkingDirectory field to given value.

### HasWorkingDirectory

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) HasWorkingDirectory() bool`

HasWorkingDirectory returns a boolean if a field has been set.

### GetIdleTimeoutMinutes

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) GetIdleTimeoutMinutes() int32`

GetIdleTimeoutMinutes returns the IdleTimeoutMinutes field if non-nil, zero value otherwise.

### GetIdleTimeoutMinutesOk

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) GetIdleTimeoutMinutesOk() (*int32, bool)`

GetIdleTimeoutMinutesOk returns a tuple with the IdleTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeoutMinutes

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) SetIdleTimeoutMinutes(v int32)`

SetIdleTimeoutMinutes sets IdleTimeoutMinutes field to given value.

### HasIdleTimeoutMinutes

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) HasIdleTimeoutMinutes() bool`

HasIdleTimeoutMinutes returns a boolean if a field has been set.

### GetDisplayName

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetRetentionPeriodMinutes

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) GetRetentionPeriodMinutes() int32`

GetRetentionPeriodMinutes returns the RetentionPeriodMinutes field if non-nil, zero value otherwise.

### GetRetentionPeriodMinutesOk

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) GetRetentionPeriodMinutesOk() (*int32, bool)`

GetRetentionPeriodMinutesOk returns a tuple with the RetentionPeriodMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriodMinutes

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) SetRetentionPeriodMinutes(v int32)`

SetRetentionPeriodMinutes sets RetentionPeriodMinutes field to given value.

### HasRetentionPeriodMinutes

`func (o *CodespacesCreateWithPrForAuthenticatedUserRequest) HasRetentionPeriodMinutes() bool`

HasRetentionPeriodMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


