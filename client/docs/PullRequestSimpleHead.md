# PullRequestSimpleHead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** |  | 
**Ref** | **string** |  | 
**Repo** | [**Repository**](Repository.md) |  | 
**Sha** | **string** |  | 
**User** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 

## Methods

### NewPullRequestSimpleHead

`func NewPullRequestSimpleHead(label string, ref string, repo Repository, sha string, user NullableNullableSimpleUser, ) *PullRequestSimpleHead`

NewPullRequestSimpleHead instantiates a new PullRequestSimpleHead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullRequestSimpleHeadWithDefaults

`func NewPullRequestSimpleHeadWithDefaults() *PullRequestSimpleHead`

NewPullRequestSimpleHeadWithDefaults instantiates a new PullRequestSimpleHead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *PullRequestSimpleHead) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PullRequestSimpleHead) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PullRequestSimpleHead) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetRef

`func (o *PullRequestSimpleHead) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *PullRequestSimpleHead) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *PullRequestSimpleHead) SetRef(v string)`

SetRef sets Ref field to given value.


### GetRepo

`func (o *PullRequestSimpleHead) GetRepo() Repository`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *PullRequestSimpleHead) GetRepoOk() (*Repository, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *PullRequestSimpleHead) SetRepo(v Repository)`

SetRepo sets Repo field to given value.


### GetSha

`func (o *PullRequestSimpleHead) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *PullRequestSimpleHead) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *PullRequestSimpleHead) SetSha(v string)`

SetSha sets Sha field to given value.


### GetUser

`func (o *PullRequestSimpleHead) GetUser() NullableSimpleUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PullRequestSimpleHead) GetUserOk() (*NullableSimpleUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PullRequestSimpleHead) SetUser(v NullableSimpleUser)`

SetUser sets User field to given value.


### SetUserNil

`func (o *PullRequestSimpleHead) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *PullRequestSimpleHead) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


