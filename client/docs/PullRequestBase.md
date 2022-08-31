# PullRequestBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** |  | 
**Ref** | **string** |  | 
**Repo** | [**PullRequestBaseRepo**](PullRequestBaseRepo.md) |  | 
**Sha** | **string** |  | 
**User** | [**PullRequestHeadRepoOwner**](PullRequestHeadRepoOwner.md) |  | 

## Methods

### NewPullRequestBase

`func NewPullRequestBase(label string, ref string, repo PullRequestBaseRepo, sha string, user PullRequestHeadRepoOwner, ) *PullRequestBase`

NewPullRequestBase instantiates a new PullRequestBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullRequestBaseWithDefaults

`func NewPullRequestBaseWithDefaults() *PullRequestBase`

NewPullRequestBaseWithDefaults instantiates a new PullRequestBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *PullRequestBase) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PullRequestBase) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PullRequestBase) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetRef

`func (o *PullRequestBase) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *PullRequestBase) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *PullRequestBase) SetRef(v string)`

SetRef sets Ref field to given value.


### GetRepo

`func (o *PullRequestBase) GetRepo() PullRequestBaseRepo`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *PullRequestBase) GetRepoOk() (*PullRequestBaseRepo, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *PullRequestBase) SetRepo(v PullRequestBaseRepo)`

SetRepo sets Repo field to given value.


### GetSha

`func (o *PullRequestBase) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *PullRequestBase) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *PullRequestBase) SetSha(v string)`

SetSha sets Sha field to given value.


### GetUser

`func (o *PullRequestBase) GetUser() PullRequestHeadRepoOwner`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PullRequestBase) GetUserOk() (*PullRequestHeadRepoOwner, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PullRequestBase) SetUser(v PullRequestHeadRepoOwner)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


