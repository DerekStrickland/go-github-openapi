# PullRequestHead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** |  | 
**Ref** | **string** |  | 
**Repo** | [**NullablePullRequestHeadRepo**](PullRequestHeadRepo.md) |  | 
**Sha** | **string** |  | 
**User** | [**PullRequestHeadRepoOwner**](PullRequestHeadRepoOwner.md) |  | 

## Methods

### NewPullRequestHead

`func NewPullRequestHead(label string, ref string, repo NullablePullRequestHeadRepo, sha string, user PullRequestHeadRepoOwner, ) *PullRequestHead`

NewPullRequestHead instantiates a new PullRequestHead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullRequestHeadWithDefaults

`func NewPullRequestHeadWithDefaults() *PullRequestHead`

NewPullRequestHeadWithDefaults instantiates a new PullRequestHead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *PullRequestHead) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PullRequestHead) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PullRequestHead) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetRef

`func (o *PullRequestHead) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *PullRequestHead) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *PullRequestHead) SetRef(v string)`

SetRef sets Ref field to given value.


### GetRepo

`func (o *PullRequestHead) GetRepo() PullRequestHeadRepo`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *PullRequestHead) GetRepoOk() (*PullRequestHeadRepo, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *PullRequestHead) SetRepo(v PullRequestHeadRepo)`

SetRepo sets Repo field to given value.


### SetRepoNil

`func (o *PullRequestHead) SetRepoNil(b bool)`

 SetRepoNil sets the value for Repo to be an explicit nil

### UnsetRepo
`func (o *PullRequestHead) UnsetRepo()`

UnsetRepo ensures that no value is present for Repo, not even an explicit nil
### GetSha

`func (o *PullRequestHead) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *PullRequestHead) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *PullRequestHead) SetSha(v string)`

SetSha sets Sha field to given value.


### GetUser

`func (o *PullRequestHead) GetUser() PullRequestHeadRepoOwner`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PullRequestHead) GetUserOk() (*PullRequestHeadRepoOwner, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PullRequestHead) SetUser(v PullRequestHeadRepoOwner)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


