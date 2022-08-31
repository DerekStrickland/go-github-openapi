# GistHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**CommittedAt** | Pointer to **time.Time** |  | [optional] 
**ChangeStatus** | Pointer to [**GistHistoryChangeStatus**](GistHistoryChangeStatus.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewGistHistory

`func NewGistHistory() *GistHistory`

NewGistHistory instantiates a new GistHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGistHistoryWithDefaults

`func NewGistHistoryWithDefaults() *GistHistory`

NewGistHistoryWithDefaults instantiates a new GistHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *GistHistory) GetUser() NullableSimpleUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GistHistory) GetUserOk() (*NullableSimpleUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GistHistory) SetUser(v NullableSimpleUser)`

SetUser sets User field to given value.

### HasUser

`func (o *GistHistory) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *GistHistory) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *GistHistory) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetVersion

`func (o *GistHistory) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GistHistory) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GistHistory) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GistHistory) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCommittedAt

`func (o *GistHistory) GetCommittedAt() time.Time`

GetCommittedAt returns the CommittedAt field if non-nil, zero value otherwise.

### GetCommittedAtOk

`func (o *GistHistory) GetCommittedAtOk() (*time.Time, bool)`

GetCommittedAtOk returns a tuple with the CommittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommittedAt

`func (o *GistHistory) SetCommittedAt(v time.Time)`

SetCommittedAt sets CommittedAt field to given value.

### HasCommittedAt

`func (o *GistHistory) HasCommittedAt() bool`

HasCommittedAt returns a boolean if a field has been set.

### GetChangeStatus

`func (o *GistHistory) GetChangeStatus() GistHistoryChangeStatus`

GetChangeStatus returns the ChangeStatus field if non-nil, zero value otherwise.

### GetChangeStatusOk

`func (o *GistHistory) GetChangeStatusOk() (*GistHistoryChangeStatus, bool)`

GetChangeStatusOk returns a tuple with the ChangeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeStatus

`func (o *GistHistory) SetChangeStatus(v GistHistoryChangeStatus)`

SetChangeStatus sets ChangeStatus field to given value.

### HasChangeStatus

`func (o *GistHistory) HasChangeStatus() bool`

HasChangeStatus returns a boolean if a field has been set.

### GetUrl

`func (o *GistHistory) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GistHistory) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GistHistory) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GistHistory) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


