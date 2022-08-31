# GistCommit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Version** | **string** |  | 
**User** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**ChangeStatus** | [**GistHistoryChangeStatus**](GistHistoryChangeStatus.md) |  | 
**CommittedAt** | **time.Time** |  | 

## Methods

### NewGistCommit

`func NewGistCommit(url string, version string, user NullableNullableSimpleUser, changeStatus GistHistoryChangeStatus, committedAt time.Time, ) *GistCommit`

NewGistCommit instantiates a new GistCommit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGistCommitWithDefaults

`func NewGistCommitWithDefaults() *GistCommit`

NewGistCommitWithDefaults instantiates a new GistCommit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *GistCommit) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GistCommit) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GistCommit) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetVersion

`func (o *GistCommit) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GistCommit) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GistCommit) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetUser

`func (o *GistCommit) GetUser() NullableSimpleUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GistCommit) GetUserOk() (*NullableSimpleUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GistCommit) SetUser(v NullableSimpleUser)`

SetUser sets User field to given value.


### SetUserNil

`func (o *GistCommit) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *GistCommit) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetChangeStatus

`func (o *GistCommit) GetChangeStatus() GistHistoryChangeStatus`

GetChangeStatus returns the ChangeStatus field if non-nil, zero value otherwise.

### GetChangeStatusOk

`func (o *GistCommit) GetChangeStatusOk() (*GistHistoryChangeStatus, bool)`

GetChangeStatusOk returns a tuple with the ChangeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeStatus

`func (o *GistCommit) SetChangeStatus(v GistHistoryChangeStatus)`

SetChangeStatus sets ChangeStatus field to given value.


### GetCommittedAt

`func (o *GistCommit) GetCommittedAt() time.Time`

GetCommittedAt returns the CommittedAt field if non-nil, zero value otherwise.

### GetCommittedAtOk

`func (o *GistCommit) GetCommittedAtOk() (*time.Time, bool)`

GetCommittedAtOk returns a tuple with the CommittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommittedAt

`func (o *GistCommit) SetCommittedAt(v time.Time)`

SetCommittedAt sets CommittedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


