# CodespacesCreateForAuthenticatedUserRequestOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepositoryId** | **int32** | Repository id for this codespace | 
**Ref** | Pointer to **string** | Git ref (typically a branch name) for this codespace | [optional] 
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

### NewCodespacesCreateForAuthenticatedUserRequestOneOf

`func NewCodespacesCreateForAuthenticatedUserRequestOneOf(repositoryId int32, ) *CodespacesCreateForAuthenticatedUserRequestOneOf`

NewCodespacesCreateForAuthenticatedUserRequestOneOf instantiates a new CodespacesCreateForAuthenticatedUserRequestOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodespacesCreateForAuthenticatedUserRequestOneOfWithDefaults

`func NewCodespacesCreateForAuthenticatedUserRequestOneOfWithDefaults() *CodespacesCreateForAuthenticatedUserRequestOneOf`

NewCodespacesCreateForAuthenticatedUserRequestOneOfWithDefaults instantiates a new CodespacesCreateForAuthenticatedUserRequestOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepositoryId

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) GetRepositoryId() int32`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) GetRepositoryIdOk() (*int32, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) SetRepositoryId(v int32)`

SetRepositoryId sets RepositoryId field to given value.


### GetRef

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetLocation

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetClientIp

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetMachine

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetDevcontainerPath

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) GetDevcontainerPath() string`

GetDevcontainerPath returns the DevcontainerPath field if non-nil, zero value otherwise.

### GetDevcontainerPathOk

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) GetDevcontainerPathOk() (*string, bool)`

GetDevcontainerPathOk returns a tuple with the DevcontainerPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevcontainerPath

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) SetDevcontainerPath(v string)`

SetDevcontainerPath sets DevcontainerPath field to given value.

### HasDevcontainerPath

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) HasDevcontainerPath() bool`

HasDevcontainerPath returns a boolean if a field has been set.

### GetMultiRepoPermissionsOptOut

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) GetMultiRepoPermissionsOptOut() bool`

GetMultiRepoPermissionsOptOut returns the MultiRepoPermissionsOptOut field if non-nil, zero value otherwise.

### GetMultiRepoPermissionsOptOutOk

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) GetMultiRepoPermissionsOptOutOk() (*bool, bool)`

GetMultiRepoPermissionsOptOutOk returns a tuple with the MultiRepoPermissionsOptOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiRepoPermissionsOptOut

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) SetMultiRepoPermissionsOptOut(v bool)`

SetMultiRepoPermissionsOptOut sets MultiRepoPermissionsOptOut field to given value.

### HasMultiRepoPermissionsOptOut

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) HasMultiRepoPermissionsOptOut() bool`

HasMultiRepoPermissionsOptOut returns a boolean if a field has been set.

### GetWorkingDirectory

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) GetWorkingDirectory() string`

GetWorkingDirectory returns the WorkingDirectory field if non-nil, zero value otherwise.

### GetWorkingDirectoryOk

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) GetWorkingDirectoryOk() (*string, bool)`

GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirectory

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) SetWorkingDirectory(v string)`

SetWorkingDirectory sets WorkingDirectory field to given value.

### HasWorkingDirectory

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) HasWorkingDirectory() bool`

HasWorkingDirectory returns a boolean if a field has been set.

### GetIdleTimeoutMinutes

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) GetIdleTimeoutMinutes() int32`

GetIdleTimeoutMinutes returns the IdleTimeoutMinutes field if non-nil, zero value otherwise.

### GetIdleTimeoutMinutesOk

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) GetIdleTimeoutMinutesOk() (*int32, bool)`

GetIdleTimeoutMinutesOk returns a tuple with the IdleTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeoutMinutes

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) SetIdleTimeoutMinutes(v int32)`

SetIdleTimeoutMinutes sets IdleTimeoutMinutes field to given value.

### HasIdleTimeoutMinutes

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) HasIdleTimeoutMinutes() bool`

HasIdleTimeoutMinutes returns a boolean if a field has been set.

### GetDisplayName

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetRetentionPeriodMinutes

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) GetRetentionPeriodMinutes() int32`

GetRetentionPeriodMinutes returns the RetentionPeriodMinutes field if non-nil, zero value otherwise.

### GetRetentionPeriodMinutesOk

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) GetRetentionPeriodMinutesOk() (*int32, bool)`

GetRetentionPeriodMinutesOk returns a tuple with the RetentionPeriodMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriodMinutes

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) SetRetentionPeriodMinutes(v int32)`

SetRetentionPeriodMinutes sets RetentionPeriodMinutes field to given value.

### HasRetentionPeriodMinutes

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf) HasRetentionPeriodMinutes() bool`

HasRetentionPeriodMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


