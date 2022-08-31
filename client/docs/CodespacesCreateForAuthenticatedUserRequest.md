# CodespacesCreateForAuthenticatedUserRequest

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
**PullRequest** | [**CodespacesCreateForAuthenticatedUserRequestOneOf1PullRequest**](CodespacesCreateForAuthenticatedUserRequestOneOf1PullRequest.md) |  | 

## Methods

### NewCodespacesCreateForAuthenticatedUserRequest

`func NewCodespacesCreateForAuthenticatedUserRequest(repositoryId int32, pullRequest CodespacesCreateForAuthenticatedUserRequestOneOf1PullRequest, ) *CodespacesCreateForAuthenticatedUserRequest`

NewCodespacesCreateForAuthenticatedUserRequest instantiates a new CodespacesCreateForAuthenticatedUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodespacesCreateForAuthenticatedUserRequestWithDefaults

`func NewCodespacesCreateForAuthenticatedUserRequestWithDefaults() *CodespacesCreateForAuthenticatedUserRequest`

NewCodespacesCreateForAuthenticatedUserRequestWithDefaults instantiates a new CodespacesCreateForAuthenticatedUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepositoryId

`func (o *CodespacesCreateForAuthenticatedUserRequest) GetRepositoryId() int32`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *CodespacesCreateForAuthenticatedUserRequest) GetRepositoryIdOk() (*int32, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *CodespacesCreateForAuthenticatedUserRequest) SetRepositoryId(v int32)`

SetRepositoryId sets RepositoryId field to given value.


### GetRef

`func (o *CodespacesCreateForAuthenticatedUserRequest) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *CodespacesCreateForAuthenticatedUserRequest) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *CodespacesCreateForAuthenticatedUserRequest) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *CodespacesCreateForAuthenticatedUserRequest) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetLocation

`func (o *CodespacesCreateForAuthenticatedUserRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CodespacesCreateForAuthenticatedUserRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CodespacesCreateForAuthenticatedUserRequest) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CodespacesCreateForAuthenticatedUserRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetClientIp

`func (o *CodespacesCreateForAuthenticatedUserRequest) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *CodespacesCreateForAuthenticatedUserRequest) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *CodespacesCreateForAuthenticatedUserRequest) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *CodespacesCreateForAuthenticatedUserRequest) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetMachine

`func (o *CodespacesCreateForAuthenticatedUserRequest) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *CodespacesCreateForAuthenticatedUserRequest) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *CodespacesCreateForAuthenticatedUserRequest) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *CodespacesCreateForAuthenticatedUserRequest) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetDevcontainerPath

`func (o *CodespacesCreateForAuthenticatedUserRequest) GetDevcontainerPath() string`

GetDevcontainerPath returns the DevcontainerPath field if non-nil, zero value otherwise.

### GetDevcontainerPathOk

`func (o *CodespacesCreateForAuthenticatedUserRequest) GetDevcontainerPathOk() (*string, bool)`

GetDevcontainerPathOk returns a tuple with the DevcontainerPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevcontainerPath

`func (o *CodespacesCreateForAuthenticatedUserRequest) SetDevcontainerPath(v string)`

SetDevcontainerPath sets DevcontainerPath field to given value.

### HasDevcontainerPath

`func (o *CodespacesCreateForAuthenticatedUserRequest) HasDevcontainerPath() bool`

HasDevcontainerPath returns a boolean if a field has been set.

### GetMultiRepoPermissionsOptOut

`func (o *CodespacesCreateForAuthenticatedUserRequest) GetMultiRepoPermissionsOptOut() bool`

GetMultiRepoPermissionsOptOut returns the MultiRepoPermissionsOptOut field if non-nil, zero value otherwise.

### GetMultiRepoPermissionsOptOutOk

`func (o *CodespacesCreateForAuthenticatedUserRequest) GetMultiRepoPermissionsOptOutOk() (*bool, bool)`

GetMultiRepoPermissionsOptOutOk returns a tuple with the MultiRepoPermissionsOptOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiRepoPermissionsOptOut

`func (o *CodespacesCreateForAuthenticatedUserRequest) SetMultiRepoPermissionsOptOut(v bool)`

SetMultiRepoPermissionsOptOut sets MultiRepoPermissionsOptOut field to given value.

### HasMultiRepoPermissionsOptOut

`func (o *CodespacesCreateForAuthenticatedUserRequest) HasMultiRepoPermissionsOptOut() bool`

HasMultiRepoPermissionsOptOut returns a boolean if a field has been set.

### GetWorkingDirectory

`func (o *CodespacesCreateForAuthenticatedUserRequest) GetWorkingDirectory() string`

GetWorkingDirectory returns the WorkingDirectory field if non-nil, zero value otherwise.

### GetWorkingDirectoryOk

`func (o *CodespacesCreateForAuthenticatedUserRequest) GetWorkingDirectoryOk() (*string, bool)`

GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirectory

`func (o *CodespacesCreateForAuthenticatedUserRequest) SetWorkingDirectory(v string)`

SetWorkingDirectory sets WorkingDirectory field to given value.

### HasWorkingDirectory

`func (o *CodespacesCreateForAuthenticatedUserRequest) HasWorkingDirectory() bool`

HasWorkingDirectory returns a boolean if a field has been set.

### GetIdleTimeoutMinutes

`func (o *CodespacesCreateForAuthenticatedUserRequest) GetIdleTimeoutMinutes() int32`

GetIdleTimeoutMinutes returns the IdleTimeoutMinutes field if non-nil, zero value otherwise.

### GetIdleTimeoutMinutesOk

`func (o *CodespacesCreateForAuthenticatedUserRequest) GetIdleTimeoutMinutesOk() (*int32, bool)`

GetIdleTimeoutMinutesOk returns a tuple with the IdleTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeoutMinutes

`func (o *CodespacesCreateForAuthenticatedUserRequest) SetIdleTimeoutMinutes(v int32)`

SetIdleTimeoutMinutes sets IdleTimeoutMinutes field to given value.

### HasIdleTimeoutMinutes

`func (o *CodespacesCreateForAuthenticatedUserRequest) HasIdleTimeoutMinutes() bool`

HasIdleTimeoutMinutes returns a boolean if a field has been set.

### GetDisplayName

`func (o *CodespacesCreateForAuthenticatedUserRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CodespacesCreateForAuthenticatedUserRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CodespacesCreateForAuthenticatedUserRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CodespacesCreateForAuthenticatedUserRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetRetentionPeriodMinutes

`func (o *CodespacesCreateForAuthenticatedUserRequest) GetRetentionPeriodMinutes() int32`

GetRetentionPeriodMinutes returns the RetentionPeriodMinutes field if non-nil, zero value otherwise.

### GetRetentionPeriodMinutesOk

`func (o *CodespacesCreateForAuthenticatedUserRequest) GetRetentionPeriodMinutesOk() (*int32, bool)`

GetRetentionPeriodMinutesOk returns a tuple with the RetentionPeriodMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriodMinutes

`func (o *CodespacesCreateForAuthenticatedUserRequest) SetRetentionPeriodMinutes(v int32)`

SetRetentionPeriodMinutes sets RetentionPeriodMinutes field to given value.

### HasRetentionPeriodMinutes

`func (o *CodespacesCreateForAuthenticatedUserRequest) HasRetentionPeriodMinutes() bool`

HasRetentionPeriodMinutes returns a boolean if a field has been set.

### GetPullRequest

`func (o *CodespacesCreateForAuthenticatedUserRequest) GetPullRequest() CodespacesCreateForAuthenticatedUserRequestOneOf1PullRequest`

GetPullRequest returns the PullRequest field if non-nil, zero value otherwise.

### GetPullRequestOk

`func (o *CodespacesCreateForAuthenticatedUserRequest) GetPullRequestOk() (*CodespacesCreateForAuthenticatedUserRequestOneOf1PullRequest, bool)`

GetPullRequestOk returns a tuple with the PullRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequest

`func (o *CodespacesCreateForAuthenticatedUserRequest) SetPullRequest(v CodespacesCreateForAuthenticatedUserRequestOneOf1PullRequest)`

SetPullRequest sets PullRequest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


