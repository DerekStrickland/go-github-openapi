# Codespace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** | Automatically generated name of this codespace. | 
**DisplayName** | Pointer to **NullableString** | Display name for this codespace. | [optional] 
**EnvironmentId** | **NullableString** | UUID identifying this codespace&#39;s environment. | 
**Owner** | [**SimpleUser**](SimpleUser.md) |  | 
**BillableOwner** | [**SimpleUser**](SimpleUser.md) |  | 
**Repository** | [**MinimalRepository**](MinimalRepository.md) |  | 
**Machine** | [**NullableNullableCodespaceMachine**](NullableCodespaceMachine.md) |  | 
**DevcontainerPath** | Pointer to **NullableString** | Path to devcontainer.json from repo root used to create Codespace. | [optional] 
**Prebuild** | **NullableBool** | Whether the codespace was created from a prebuild. | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**LastUsedAt** | **time.Time** | Last known time this codespace was started. | 
**State** | **string** | State of this codespace. | 
**Url** | **string** | API URL for this codespace. | 
**GitStatus** | [**CodespaceGitStatus**](CodespaceGitStatus.md) |  | 
**Location** | **string** | The Azure region where this codespace is located. | 
**IdleTimeoutMinutes** | **NullableInt32** | The number of minutes of inactivity after which this codespace will be automatically stopped. | 
**WebUrl** | **string** | URL to access this codespace on the web. | 
**MachinesUrl** | **string** | API URL to access available alternate machine types for this codespace. | 
**StartUrl** | **string** | API URL to start this codespace. | 
**StopUrl** | **string** | API URL to stop this codespace. | 
**PullsUrl** | **NullableString** | API URL for the Pull Request associated with this codespace, if any. | 
**RecentFolders** | **[]string** |  | 
**RuntimeConstraints** | Pointer to [**CodespaceRuntimeConstraints**](CodespaceRuntimeConstraints.md) |  | [optional] 
**PendingOperation** | Pointer to **NullableBool** | Whether or not a codespace has a pending async operation. This would mean that the codespace is temporarily unavailable. The only thing that you can do with a codespace in this state is delete it. | [optional] 
**PendingOperationDisabledReason** | Pointer to **NullableString** | Text to show user when codespace is disabled by a pending operation | [optional] 
**IdleTimeoutNotice** | Pointer to **NullableString** | Text to show user when codespace idle timeout minutes has been overriden by an organization policy | [optional] 
**RetentionPeriodMinutes** | Pointer to **NullableInt32** | Duration in minutes after codespace has gone idle in which it will be deleted. Must be integer minutes between 0 and 43200 (30 days). | [optional] 
**RetentionExpiresAt** | Pointer to **NullableTime** | When a codespace will be auto-deleted based on the \&quot;retention_period_minutes\&quot; and \&quot;last_used_at\&quot; | [optional] 
**LastKnownStopNotice** | Pointer to **NullableString** | The text to display to a user when a codespace has been stopped for a potentially actionable reason. | [optional] 

## Methods

### NewCodespace

`func NewCodespace(id int32, name string, environmentId NullableString, owner SimpleUser, billableOwner SimpleUser, repository MinimalRepository, machine NullableNullableCodespaceMachine, prebuild NullableBool, createdAt time.Time, updatedAt time.Time, lastUsedAt time.Time, state string, url string, gitStatus CodespaceGitStatus, location string, idleTimeoutMinutes NullableInt32, webUrl string, machinesUrl string, startUrl string, stopUrl string, pullsUrl NullableString, recentFolders []string, ) *Codespace`

NewCodespace instantiates a new Codespace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodespaceWithDefaults

`func NewCodespaceWithDefaults() *Codespace`

NewCodespaceWithDefaults instantiates a new Codespace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Codespace) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Codespace) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Codespace) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Codespace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Codespace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Codespace) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *Codespace) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Codespace) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Codespace) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Codespace) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *Codespace) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *Codespace) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEnvironmentId

`func (o *Codespace) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *Codespace) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *Codespace) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### SetEnvironmentIdNil

`func (o *Codespace) SetEnvironmentIdNil(b bool)`

 SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil

### UnsetEnvironmentId
`func (o *Codespace) UnsetEnvironmentId()`

UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil
### GetOwner

`func (o *Codespace) GetOwner() SimpleUser`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Codespace) GetOwnerOk() (*SimpleUser, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Codespace) SetOwner(v SimpleUser)`

SetOwner sets Owner field to given value.


### GetBillableOwner

`func (o *Codespace) GetBillableOwner() SimpleUser`

GetBillableOwner returns the BillableOwner field if non-nil, zero value otherwise.

### GetBillableOwnerOk

`func (o *Codespace) GetBillableOwnerOk() (*SimpleUser, bool)`

GetBillableOwnerOk returns a tuple with the BillableOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableOwner

`func (o *Codespace) SetBillableOwner(v SimpleUser)`

SetBillableOwner sets BillableOwner field to given value.


### GetRepository

`func (o *Codespace) GetRepository() MinimalRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *Codespace) GetRepositoryOk() (*MinimalRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *Codespace) SetRepository(v MinimalRepository)`

SetRepository sets Repository field to given value.


### GetMachine

`func (o *Codespace) GetMachine() NullableCodespaceMachine`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *Codespace) GetMachineOk() (*NullableCodespaceMachine, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *Codespace) SetMachine(v NullableCodespaceMachine)`

SetMachine sets Machine field to given value.


### SetMachineNil

`func (o *Codespace) SetMachineNil(b bool)`

 SetMachineNil sets the value for Machine to be an explicit nil

### UnsetMachine
`func (o *Codespace) UnsetMachine()`

UnsetMachine ensures that no value is present for Machine, not even an explicit nil
### GetDevcontainerPath

`func (o *Codespace) GetDevcontainerPath() string`

GetDevcontainerPath returns the DevcontainerPath field if non-nil, zero value otherwise.

### GetDevcontainerPathOk

`func (o *Codespace) GetDevcontainerPathOk() (*string, bool)`

GetDevcontainerPathOk returns a tuple with the DevcontainerPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevcontainerPath

`func (o *Codespace) SetDevcontainerPath(v string)`

SetDevcontainerPath sets DevcontainerPath field to given value.

### HasDevcontainerPath

`func (o *Codespace) HasDevcontainerPath() bool`

HasDevcontainerPath returns a boolean if a field has been set.

### SetDevcontainerPathNil

`func (o *Codespace) SetDevcontainerPathNil(b bool)`

 SetDevcontainerPathNil sets the value for DevcontainerPath to be an explicit nil

### UnsetDevcontainerPath
`func (o *Codespace) UnsetDevcontainerPath()`

UnsetDevcontainerPath ensures that no value is present for DevcontainerPath, not even an explicit nil
### GetPrebuild

`func (o *Codespace) GetPrebuild() bool`

GetPrebuild returns the Prebuild field if non-nil, zero value otherwise.

### GetPrebuildOk

`func (o *Codespace) GetPrebuildOk() (*bool, bool)`

GetPrebuildOk returns a tuple with the Prebuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrebuild

`func (o *Codespace) SetPrebuild(v bool)`

SetPrebuild sets Prebuild field to given value.


### SetPrebuildNil

`func (o *Codespace) SetPrebuildNil(b bool)`

 SetPrebuildNil sets the value for Prebuild to be an explicit nil

### UnsetPrebuild
`func (o *Codespace) UnsetPrebuild()`

UnsetPrebuild ensures that no value is present for Prebuild, not even an explicit nil
### GetCreatedAt

`func (o *Codespace) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Codespace) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Codespace) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Codespace) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Codespace) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Codespace) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLastUsedAt

`func (o *Codespace) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *Codespace) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *Codespace) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.


### GetState

`func (o *Codespace) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Codespace) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Codespace) SetState(v string)`

SetState sets State field to given value.


### GetUrl

`func (o *Codespace) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Codespace) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Codespace) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetGitStatus

`func (o *Codespace) GetGitStatus() CodespaceGitStatus`

GetGitStatus returns the GitStatus field if non-nil, zero value otherwise.

### GetGitStatusOk

`func (o *Codespace) GetGitStatusOk() (*CodespaceGitStatus, bool)`

GetGitStatusOk returns a tuple with the GitStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitStatus

`func (o *Codespace) SetGitStatus(v CodespaceGitStatus)`

SetGitStatus sets GitStatus field to given value.


### GetLocation

`func (o *Codespace) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Codespace) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Codespace) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetIdleTimeoutMinutes

`func (o *Codespace) GetIdleTimeoutMinutes() int32`

GetIdleTimeoutMinutes returns the IdleTimeoutMinutes field if non-nil, zero value otherwise.

### GetIdleTimeoutMinutesOk

`func (o *Codespace) GetIdleTimeoutMinutesOk() (*int32, bool)`

GetIdleTimeoutMinutesOk returns a tuple with the IdleTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeoutMinutes

`func (o *Codespace) SetIdleTimeoutMinutes(v int32)`

SetIdleTimeoutMinutes sets IdleTimeoutMinutes field to given value.


### SetIdleTimeoutMinutesNil

`func (o *Codespace) SetIdleTimeoutMinutesNil(b bool)`

 SetIdleTimeoutMinutesNil sets the value for IdleTimeoutMinutes to be an explicit nil

### UnsetIdleTimeoutMinutes
`func (o *Codespace) UnsetIdleTimeoutMinutes()`

UnsetIdleTimeoutMinutes ensures that no value is present for IdleTimeoutMinutes, not even an explicit nil
### GetWebUrl

`func (o *Codespace) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *Codespace) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *Codespace) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.


### GetMachinesUrl

`func (o *Codespace) GetMachinesUrl() string`

GetMachinesUrl returns the MachinesUrl field if non-nil, zero value otherwise.

### GetMachinesUrlOk

`func (o *Codespace) GetMachinesUrlOk() (*string, bool)`

GetMachinesUrlOk returns a tuple with the MachinesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachinesUrl

`func (o *Codespace) SetMachinesUrl(v string)`

SetMachinesUrl sets MachinesUrl field to given value.


### GetStartUrl

`func (o *Codespace) GetStartUrl() string`

GetStartUrl returns the StartUrl field if non-nil, zero value otherwise.

### GetStartUrlOk

`func (o *Codespace) GetStartUrlOk() (*string, bool)`

GetStartUrlOk returns a tuple with the StartUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartUrl

`func (o *Codespace) SetStartUrl(v string)`

SetStartUrl sets StartUrl field to given value.


### GetStopUrl

`func (o *Codespace) GetStopUrl() string`

GetStopUrl returns the StopUrl field if non-nil, zero value otherwise.

### GetStopUrlOk

`func (o *Codespace) GetStopUrlOk() (*string, bool)`

GetStopUrlOk returns a tuple with the StopUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopUrl

`func (o *Codespace) SetStopUrl(v string)`

SetStopUrl sets StopUrl field to given value.


### GetPullsUrl

`func (o *Codespace) GetPullsUrl() string`

GetPullsUrl returns the PullsUrl field if non-nil, zero value otherwise.

### GetPullsUrlOk

`func (o *Codespace) GetPullsUrlOk() (*string, bool)`

GetPullsUrlOk returns a tuple with the PullsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullsUrl

`func (o *Codespace) SetPullsUrl(v string)`

SetPullsUrl sets PullsUrl field to given value.


### SetPullsUrlNil

`func (o *Codespace) SetPullsUrlNil(b bool)`

 SetPullsUrlNil sets the value for PullsUrl to be an explicit nil

### UnsetPullsUrl
`func (o *Codespace) UnsetPullsUrl()`

UnsetPullsUrl ensures that no value is present for PullsUrl, not even an explicit nil
### GetRecentFolders

`func (o *Codespace) GetRecentFolders() []string`

GetRecentFolders returns the RecentFolders field if non-nil, zero value otherwise.

### GetRecentFoldersOk

`func (o *Codespace) GetRecentFoldersOk() (*[]string, bool)`

GetRecentFoldersOk returns a tuple with the RecentFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentFolders

`func (o *Codespace) SetRecentFolders(v []string)`

SetRecentFolders sets RecentFolders field to given value.


### GetRuntimeConstraints

`func (o *Codespace) GetRuntimeConstraints() CodespaceRuntimeConstraints`

GetRuntimeConstraints returns the RuntimeConstraints field if non-nil, zero value otherwise.

### GetRuntimeConstraintsOk

`func (o *Codespace) GetRuntimeConstraintsOk() (*CodespaceRuntimeConstraints, bool)`

GetRuntimeConstraintsOk returns a tuple with the RuntimeConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeConstraints

`func (o *Codespace) SetRuntimeConstraints(v CodespaceRuntimeConstraints)`

SetRuntimeConstraints sets RuntimeConstraints field to given value.

### HasRuntimeConstraints

`func (o *Codespace) HasRuntimeConstraints() bool`

HasRuntimeConstraints returns a boolean if a field has been set.

### GetPendingOperation

`func (o *Codespace) GetPendingOperation() bool`

GetPendingOperation returns the PendingOperation field if non-nil, zero value otherwise.

### GetPendingOperationOk

`func (o *Codespace) GetPendingOperationOk() (*bool, bool)`

GetPendingOperationOk returns a tuple with the PendingOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingOperation

`func (o *Codespace) SetPendingOperation(v bool)`

SetPendingOperation sets PendingOperation field to given value.

### HasPendingOperation

`func (o *Codespace) HasPendingOperation() bool`

HasPendingOperation returns a boolean if a field has been set.

### SetPendingOperationNil

`func (o *Codespace) SetPendingOperationNil(b bool)`

 SetPendingOperationNil sets the value for PendingOperation to be an explicit nil

### UnsetPendingOperation
`func (o *Codespace) UnsetPendingOperation()`

UnsetPendingOperation ensures that no value is present for PendingOperation, not even an explicit nil
### GetPendingOperationDisabledReason

`func (o *Codespace) GetPendingOperationDisabledReason() string`

GetPendingOperationDisabledReason returns the PendingOperationDisabledReason field if non-nil, zero value otherwise.

### GetPendingOperationDisabledReasonOk

`func (o *Codespace) GetPendingOperationDisabledReasonOk() (*string, bool)`

GetPendingOperationDisabledReasonOk returns a tuple with the PendingOperationDisabledReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingOperationDisabledReason

`func (o *Codespace) SetPendingOperationDisabledReason(v string)`

SetPendingOperationDisabledReason sets PendingOperationDisabledReason field to given value.

### HasPendingOperationDisabledReason

`func (o *Codespace) HasPendingOperationDisabledReason() bool`

HasPendingOperationDisabledReason returns a boolean if a field has been set.

### SetPendingOperationDisabledReasonNil

`func (o *Codespace) SetPendingOperationDisabledReasonNil(b bool)`

 SetPendingOperationDisabledReasonNil sets the value for PendingOperationDisabledReason to be an explicit nil

### UnsetPendingOperationDisabledReason
`func (o *Codespace) UnsetPendingOperationDisabledReason()`

UnsetPendingOperationDisabledReason ensures that no value is present for PendingOperationDisabledReason, not even an explicit nil
### GetIdleTimeoutNotice

`func (o *Codespace) GetIdleTimeoutNotice() string`

GetIdleTimeoutNotice returns the IdleTimeoutNotice field if non-nil, zero value otherwise.

### GetIdleTimeoutNoticeOk

`func (o *Codespace) GetIdleTimeoutNoticeOk() (*string, bool)`

GetIdleTimeoutNoticeOk returns a tuple with the IdleTimeoutNotice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeoutNotice

`func (o *Codespace) SetIdleTimeoutNotice(v string)`

SetIdleTimeoutNotice sets IdleTimeoutNotice field to given value.

### HasIdleTimeoutNotice

`func (o *Codespace) HasIdleTimeoutNotice() bool`

HasIdleTimeoutNotice returns a boolean if a field has been set.

### SetIdleTimeoutNoticeNil

`func (o *Codespace) SetIdleTimeoutNoticeNil(b bool)`

 SetIdleTimeoutNoticeNil sets the value for IdleTimeoutNotice to be an explicit nil

### UnsetIdleTimeoutNotice
`func (o *Codespace) UnsetIdleTimeoutNotice()`

UnsetIdleTimeoutNotice ensures that no value is present for IdleTimeoutNotice, not even an explicit nil
### GetRetentionPeriodMinutes

`func (o *Codespace) GetRetentionPeriodMinutes() int32`

GetRetentionPeriodMinutes returns the RetentionPeriodMinutes field if non-nil, zero value otherwise.

### GetRetentionPeriodMinutesOk

`func (o *Codespace) GetRetentionPeriodMinutesOk() (*int32, bool)`

GetRetentionPeriodMinutesOk returns a tuple with the RetentionPeriodMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriodMinutes

`func (o *Codespace) SetRetentionPeriodMinutes(v int32)`

SetRetentionPeriodMinutes sets RetentionPeriodMinutes field to given value.

### HasRetentionPeriodMinutes

`func (o *Codespace) HasRetentionPeriodMinutes() bool`

HasRetentionPeriodMinutes returns a boolean if a field has been set.

### SetRetentionPeriodMinutesNil

`func (o *Codespace) SetRetentionPeriodMinutesNil(b bool)`

 SetRetentionPeriodMinutesNil sets the value for RetentionPeriodMinutes to be an explicit nil

### UnsetRetentionPeriodMinutes
`func (o *Codespace) UnsetRetentionPeriodMinutes()`

UnsetRetentionPeriodMinutes ensures that no value is present for RetentionPeriodMinutes, not even an explicit nil
### GetRetentionExpiresAt

`func (o *Codespace) GetRetentionExpiresAt() time.Time`

GetRetentionExpiresAt returns the RetentionExpiresAt field if non-nil, zero value otherwise.

### GetRetentionExpiresAtOk

`func (o *Codespace) GetRetentionExpiresAtOk() (*time.Time, bool)`

GetRetentionExpiresAtOk returns a tuple with the RetentionExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionExpiresAt

`func (o *Codespace) SetRetentionExpiresAt(v time.Time)`

SetRetentionExpiresAt sets RetentionExpiresAt field to given value.

### HasRetentionExpiresAt

`func (o *Codespace) HasRetentionExpiresAt() bool`

HasRetentionExpiresAt returns a boolean if a field has been set.

### SetRetentionExpiresAtNil

`func (o *Codespace) SetRetentionExpiresAtNil(b bool)`

 SetRetentionExpiresAtNil sets the value for RetentionExpiresAt to be an explicit nil

### UnsetRetentionExpiresAt
`func (o *Codespace) UnsetRetentionExpiresAt()`

UnsetRetentionExpiresAt ensures that no value is present for RetentionExpiresAt, not even an explicit nil
### GetLastKnownStopNotice

`func (o *Codespace) GetLastKnownStopNotice() string`

GetLastKnownStopNotice returns the LastKnownStopNotice field if non-nil, zero value otherwise.

### GetLastKnownStopNoticeOk

`func (o *Codespace) GetLastKnownStopNoticeOk() (*string, bool)`

GetLastKnownStopNoticeOk returns a tuple with the LastKnownStopNotice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastKnownStopNotice

`func (o *Codespace) SetLastKnownStopNotice(v string)`

SetLastKnownStopNotice sets LastKnownStopNotice field to given value.

### HasLastKnownStopNotice

`func (o *Codespace) HasLastKnownStopNotice() bool`

HasLastKnownStopNotice returns a boolean if a field has been set.

### SetLastKnownStopNoticeNil

`func (o *Codespace) SetLastKnownStopNoticeNil(b bool)`

 SetLastKnownStopNoticeNil sets the value for LastKnownStopNotice to be an explicit nil

### UnsetLastKnownStopNotice
`func (o *Codespace) UnsetLastKnownStopNotice()`

UnsetLastKnownStopNotice ensures that no value is present for LastKnownStopNotice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


