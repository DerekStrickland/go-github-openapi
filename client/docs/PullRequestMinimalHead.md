# PullRequestMinimalHead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | **string** |  | 
**Sha** | **string** |  | 
**Repo** | [**PullRequestMinimalHeadRepo**](PullRequestMinimalHeadRepo.md) |  | 

## Methods

### NewPullRequestMinimalHead

`func NewPullRequestMinimalHead(ref string, sha string, repo PullRequestMinimalHeadRepo, ) *PullRequestMinimalHead`

NewPullRequestMinimalHead instantiates a new PullRequestMinimalHead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullRequestMinimalHeadWithDefaults

`func NewPullRequestMinimalHeadWithDefaults() *PullRequestMinimalHead`

NewPullRequestMinimalHeadWithDefaults instantiates a new PullRequestMinimalHead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *PullRequestMinimalHead) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *PullRequestMinimalHead) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *PullRequestMinimalHead) SetRef(v string)`

SetRef sets Ref field to given value.


### GetSha

`func (o *PullRequestMinimalHead) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *PullRequestMinimalHead) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *PullRequestMinimalHead) SetSha(v string)`

SetSha sets Sha field to given value.


### GetRepo

`func (o *PullRequestMinimalHead) GetRepo() PullRequestMinimalHeadRepo`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *PullRequestMinimalHead) GetRepoOk() (*PullRequestMinimalHeadRepo, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *PullRequestMinimalHead) SetRepo(v PullRequestMinimalHeadRepo)`

SetRepo sets Repo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


