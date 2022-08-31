# CodespaceGitStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ahead** | Pointer to **int32** | The number of commits the local repository is ahead of the remote. | [optional] 
**Behind** | Pointer to **int32** | The number of commits the local repository is behind the remote. | [optional] 
**HasUnpushedChanges** | Pointer to **bool** | Whether the local repository has unpushed changes. | [optional] 
**HasUncommittedChanges** | Pointer to **bool** | Whether the local repository has uncommitted changes. | [optional] 
**Ref** | Pointer to **string** | The current branch (or SHA if in detached HEAD state) of the local repository. | [optional] 

## Methods

### NewCodespaceGitStatus

`func NewCodespaceGitStatus() *CodespaceGitStatus`

NewCodespaceGitStatus instantiates a new CodespaceGitStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodespaceGitStatusWithDefaults

`func NewCodespaceGitStatusWithDefaults() *CodespaceGitStatus`

NewCodespaceGitStatusWithDefaults instantiates a new CodespaceGitStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAhead

`func (o *CodespaceGitStatus) GetAhead() int32`

GetAhead returns the Ahead field if non-nil, zero value otherwise.

### GetAheadOk

`func (o *CodespaceGitStatus) GetAheadOk() (*int32, bool)`

GetAheadOk returns a tuple with the Ahead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAhead

`func (o *CodespaceGitStatus) SetAhead(v int32)`

SetAhead sets Ahead field to given value.

### HasAhead

`func (o *CodespaceGitStatus) HasAhead() bool`

HasAhead returns a boolean if a field has been set.

### GetBehind

`func (o *CodespaceGitStatus) GetBehind() int32`

GetBehind returns the Behind field if non-nil, zero value otherwise.

### GetBehindOk

`func (o *CodespaceGitStatus) GetBehindOk() (*int32, bool)`

GetBehindOk returns a tuple with the Behind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehind

`func (o *CodespaceGitStatus) SetBehind(v int32)`

SetBehind sets Behind field to given value.

### HasBehind

`func (o *CodespaceGitStatus) HasBehind() bool`

HasBehind returns a boolean if a field has been set.

### GetHasUnpushedChanges

`func (o *CodespaceGitStatus) GetHasUnpushedChanges() bool`

GetHasUnpushedChanges returns the HasUnpushedChanges field if non-nil, zero value otherwise.

### GetHasUnpushedChangesOk

`func (o *CodespaceGitStatus) GetHasUnpushedChangesOk() (*bool, bool)`

GetHasUnpushedChangesOk returns a tuple with the HasUnpushedChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUnpushedChanges

`func (o *CodespaceGitStatus) SetHasUnpushedChanges(v bool)`

SetHasUnpushedChanges sets HasUnpushedChanges field to given value.

### HasHasUnpushedChanges

`func (o *CodespaceGitStatus) HasHasUnpushedChanges() bool`

HasHasUnpushedChanges returns a boolean if a field has been set.

### GetHasUncommittedChanges

`func (o *CodespaceGitStatus) GetHasUncommittedChanges() bool`

GetHasUncommittedChanges returns the HasUncommittedChanges field if non-nil, zero value otherwise.

### GetHasUncommittedChangesOk

`func (o *CodespaceGitStatus) GetHasUncommittedChangesOk() (*bool, bool)`

GetHasUncommittedChangesOk returns a tuple with the HasUncommittedChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUncommittedChanges

`func (o *CodespaceGitStatus) SetHasUncommittedChanges(v bool)`

SetHasUncommittedChanges sets HasUncommittedChanges field to given value.

### HasHasUncommittedChanges

`func (o *CodespaceGitStatus) HasHasUncommittedChanges() bool`

HasHasUncommittedChanges returns a boolean if a field has been set.

### GetRef

`func (o *CodespaceGitStatus) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *CodespaceGitStatus) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *CodespaceGitStatus) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *CodespaceGitStatus) HasRef() bool`

HasRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


