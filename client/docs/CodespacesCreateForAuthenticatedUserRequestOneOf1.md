# CodespacesCreateForAuthenticatedUserRequestOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PullRequest** | [**CodespacesCreateForAuthenticatedUserRequestOneOf1PullRequest**](CodespacesCreateForAuthenticatedUserRequestOneOf1PullRequest.md) |  | 
**Location** | Pointer to **string** | Location for this codespace. Assigned by IP if not provided | [optional] 
**Machine** | Pointer to **string** | Machine type to use for this codespace | [optional] 
**DevcontainerPath** | Pointer to **string** | Path to devcontainer.json config to use for this codespace | [optional] 
**WorkingDirectory** | Pointer to **string** | Working directory for this codespace | [optional] 
**IdleTimeoutMinutes** | Pointer to **int32** | Time in minutes before codespace stops from inactivity | [optional] 

## Methods

### NewCodespacesCreateForAuthenticatedUserRequestOneOf1

`func NewCodespacesCreateForAuthenticatedUserRequestOneOf1(pullRequest CodespacesCreateForAuthenticatedUserRequestOneOf1PullRequest, ) *CodespacesCreateForAuthenticatedUserRequestOneOf1`

NewCodespacesCreateForAuthenticatedUserRequestOneOf1 instantiates a new CodespacesCreateForAuthenticatedUserRequestOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodespacesCreateForAuthenticatedUserRequestOneOf1WithDefaults

`func NewCodespacesCreateForAuthenticatedUserRequestOneOf1WithDefaults() *CodespacesCreateForAuthenticatedUserRequestOneOf1`

NewCodespacesCreateForAuthenticatedUserRequestOneOf1WithDefaults instantiates a new CodespacesCreateForAuthenticatedUserRequestOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPullRequest

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf1) GetPullRequest() CodespacesCreateForAuthenticatedUserRequestOneOf1PullRequest`

GetPullRequest returns the PullRequest field if non-nil, zero value otherwise.

### GetPullRequestOk

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf1) GetPullRequestOk() (*CodespacesCreateForAuthenticatedUserRequestOneOf1PullRequest, bool)`

GetPullRequestOk returns a tuple with the PullRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequest

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf1) SetPullRequest(v CodespacesCreateForAuthenticatedUserRequestOneOf1PullRequest)`

SetPullRequest sets PullRequest field to given value.


### GetLocation

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf1) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf1) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf1) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf1) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMachine

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf1) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf1) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf1) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf1) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetDevcontainerPath

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf1) GetDevcontainerPath() string`

GetDevcontainerPath returns the DevcontainerPath field if non-nil, zero value otherwise.

### GetDevcontainerPathOk

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf1) GetDevcontainerPathOk() (*string, bool)`

GetDevcontainerPathOk returns a tuple with the DevcontainerPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevcontainerPath

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf1) SetDevcontainerPath(v string)`

SetDevcontainerPath sets DevcontainerPath field to given value.

### HasDevcontainerPath

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf1) HasDevcontainerPath() bool`

HasDevcontainerPath returns a boolean if a field has been set.

### GetWorkingDirectory

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf1) GetWorkingDirectory() string`

GetWorkingDirectory returns the WorkingDirectory field if non-nil, zero value otherwise.

### GetWorkingDirectoryOk

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf1) GetWorkingDirectoryOk() (*string, bool)`

GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirectory

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf1) SetWorkingDirectory(v string)`

SetWorkingDirectory sets WorkingDirectory field to given value.

### HasWorkingDirectory

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf1) HasWorkingDirectory() bool`

HasWorkingDirectory returns a boolean if a field has been set.

### GetIdleTimeoutMinutes

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf1) GetIdleTimeoutMinutes() int32`

GetIdleTimeoutMinutes returns the IdleTimeoutMinutes field if non-nil, zero value otherwise.

### GetIdleTimeoutMinutesOk

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf1) GetIdleTimeoutMinutesOk() (*int32, bool)`

GetIdleTimeoutMinutesOk returns a tuple with the IdleTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeoutMinutes

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf1) SetIdleTimeoutMinutes(v int32)`

SetIdleTimeoutMinutes sets IdleTimeoutMinutes field to given value.

### HasIdleTimeoutMinutes

`func (o *CodespacesCreateForAuthenticatedUserRequestOneOf1) HasIdleTimeoutMinutes() bool`

HasIdleTimeoutMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


